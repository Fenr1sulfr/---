package repository

import (
	"architecture_go/services/contact/internal/domain"
	"context"
)

type ContactRepository interface {
	// Contact model
	CreateContact(ctx context.Context, contact domain.Contact) (int, error)
	GetContact(ctx context.Context, id int) (*domain.Contact, error)
	UpdateContact(ctx context.Context, contact domain.Contact) error
	DeleteContact(ctx context.Context, id int) error

	// Group model
	CreateGroup(ctx context.Context, group domain.Group) (int, error)
	GetGroup(ctx context.Context, id int) (*domain.Group, error)
	UpdateGroup(ctx context.Context, group domain.Group) error
	DeleteGroup(ctx context.Context, id int) error

	// ContactGroup model
	AddContactToGroup(ctx context.Context, contactID, groupID int) error
	RemoveContactFromGroup(ctx context.Context, contactID, groupID int) error
	GetContactsByGroup(ctx context.Context, groupID int) ([]*domain.Contact, error)
}
