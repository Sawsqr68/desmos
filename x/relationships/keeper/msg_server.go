package keeper

import (
	"context"
	"fmt"

	subspacestypes "github.com/desmos-labs/desmos/v7/x/subspaces/types"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v7/x/relationships/types"
)

var _ types.MsgServer = &msgServer{}

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the relationships MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{keeper}
}

// CreateRelationship defines a rpc method for MsgCreateRelationship
func (k msgServer) CreateRelationship(goCtx context.Context, msg *types.MsgCreateRelationship) (*types.MsgCreateRelationshipResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.DoesSubspaceExist(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the receiver has blocked the sender before
	if k.HasUserBlocked(ctx, msg.Counterparty, msg.Signer, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "%s is blocked by %s", msg.Signer, msg.Counterparty)
	}

	// Check if the relationship already exists
	if k.HasRelationship(ctx, msg.Signer, msg.Counterparty, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "relationship from %s to %s already exists inside subspace %d",
			msg.Signer, msg.Counterparty, msg.SubspaceID)
	}

	// Save the relationship
	k.SaveRelationship(ctx, types.NewRelationship(msg.Signer, msg.Counterparty, msg.SubspaceID))

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreatedRelationship,
			sdk.NewAttribute(types.AttributeRelationshipCreator, msg.Signer),
			sdk.NewAttribute(types.AttributeRelationshipCounterparty, msg.Counterparty),
			sdk.NewAttribute(subspacestypes.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
		),
	})

	return &types.MsgCreateRelationshipResponse{}, nil
}

// DeleteRelationship defines a rpc method for MsgDeleteRelationship
func (k msgServer) DeleteRelationship(goCtx context.Context, msg *types.MsgDeleteRelationship) (*types.MsgDeleteRelationshipResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.DoesSubspaceExist(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the relationship exists
	if !k.HasRelationship(ctx, msg.Signer, msg.Counterparty, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "relationship from %s to %s does not exist inside subspace %d",
			msg.Signer, msg.Counterparty, msg.SubspaceID)
	}

	k.Keeper.DeleteRelationship(ctx, msg.Signer, msg.Counterparty, msg.SubspaceID)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeDeletedRelationship,
			sdk.NewAttribute(types.AttributeRelationshipCreator, msg.Signer),
			sdk.NewAttribute(types.AttributeRelationshipCounterparty, msg.Counterparty),
			sdk.NewAttribute(subspacestypes.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
		),
	})

	return &types.MsgDeleteRelationshipResponse{}, nil
}

// BlockUser defines a rpc method for MsgBlockUser
func (k msgServer) BlockUser(goCtx context.Context, msg *types.MsgBlockUser) (*types.MsgBlockUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.DoesSubspaceExist(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the receiver has blocked the sender before
	if k.HasUserBlocked(ctx, msg.Blocker, msg.Blocked, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "%s has already blocked %s", msg.Blocker, msg.Blocked)
	}

	// Save the block
	k.SaveUserBlock(ctx, types.NewUserBlock(msg.Blocker, msg.Blocked, msg.Reason, msg.SubspaceID))

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeBlockedUser,
			sdk.NewAttribute(types.AttributeKeyUserBlockBlocker, msg.Blocker),
			sdk.NewAttribute(types.AttributeKeyUserBlockBlocked, msg.Blocked),
			sdk.NewAttribute(subspacestypes.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
		),
	})

	return &types.MsgBlockUserResponse{}, nil
}

// UnblockUser defines a rpc method for MsgUnblockUser
func (k msgServer) UnblockUser(goCtx context.Context, msg *types.MsgUnblockUser) (*types.MsgUnblockUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.DoesSubspaceExist(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the block exists
	if !k.HasUserBlocked(ctx, msg.Blocker, msg.Blocked, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "%s has not blocked %s", msg.Blocker, msg.Blocked)
	}

	// Delete the block
	k.DeleteUserBlock(ctx, msg.Blocker, msg.Blocked, msg.SubspaceID)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUnblockedUser,
			sdk.NewAttribute(types.AttributeKeyUserBlockBlocker, msg.Blocker),
			sdk.NewAttribute(types.AttributeKeyUserBlockBlocked, msg.Blocked),
			sdk.NewAttribute(subspacestypes.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
		),
	})

	return &types.MsgUnblockUserResponse{}, nil
}
