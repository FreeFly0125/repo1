// Code generated by github.com/firstcontributions/matro. DO NOT EDIT.

package schema

import (
	"context"
	"errors"

	"github.com/firstcontributions/backend/internal/gateway/session"
	"github.com/firstcontributions/backend/internal/storemanager"
)

func (m *Resolver) CreateReaction(
	ctx context.Context,
	args struct {
		Reaction *CreateReactionInput
	},
) (*Reaction, error) {
	session := session.FromContext(ctx)
	if session == nil {
		return nil, errors.New("unauthorized")
	}

	reactionModelInput, err := args.Reaction.ToModel()
	if err != nil {
		return nil, err
	}
	reactionModelInput.CreatedBy = session.UserID()

	reaction, err := storemanager.FromContext(ctx).StoriesStore.CreateReaction(ctx, reactionModelInput)
	if err != nil {
		return nil, err
	}
	return NewReaction(reaction), nil
}
func (m *Resolver) UpdateReaction(
	ctx context.Context,
	args struct {
		Reaction *UpdateReactionInput
	},
) (*Reaction, error) {
	store := storemanager.FromContext(ctx)

	id, err := ParseGraphqlID(args.Reaction.ID)
	if err != nil {
		return nil, err
	}
	if err := store.StoriesStore.UpdateReaction(ctx, id.ID, args.Reaction.ToModel()); err != nil {
		return nil, err
	}
	reaction, err := store.StoriesStore.GetReactionByID(ctx, id.ID)
	if err != nil {
		return nil, err
	}
	return NewReaction(reaction), nil
}
