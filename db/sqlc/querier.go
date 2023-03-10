// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CountMembers(ctx context.Context) (int64, error)
	CreateMember(ctx context.Context, arg CreateMemberParams) (Member, error)
	DeleteMember(ctx context.Context, id uuid.UUID) error
	DeleteMembers(ctx context.Context, dollar_1 []uuid.UUID) error
	GetMember(ctx context.Context, id uuid.UUID) (Member, error)
	ListMembers(ctx context.Context, arg ListMembersParams) ([]Member, error)
	UpdateMember(ctx context.Context, arg UpdateMemberParams) (Member, error)
}

var _ Querier = (*Queries)(nil)
