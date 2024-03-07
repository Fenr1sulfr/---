package usecase

import (
	"architecture_go/services/contact/internal/domain"
	"architecture_go/services/contact/internal/repository"
	"context"
)

type ContactUseCaseImpl struct {
	repository repository.ContactRepository
}

func NewContactUseCase(repository repository.ContactRepository) ContactUseCase {
	return &ContactUseCaseImpl{
		repository: repository,
	}
}

func (c *ContactUseCaseImpl) CreateContact(ctx context.Context, contact domain.Contact) (int, error) {
	return c.repository.CreateContact(ctx, contact)
}

func (c *ContactUseCaseImpl) GetContact(ctx context.Context, id int) (*domain.Contact, error) {
	return c.repository.GetContact(ctx, id)
}

func (c *ContactUseCaseImpl) UpdateContact(ctx context.Context, contact domain.Contact) error {
	return c.repository.UpdateContact(ctx, contact)
}

func (c *ContactUseCaseImpl) DeleteContact(ctx context.Context, id int) error {
	return c.repository.DeleteContact(ctx, id)
}

func (c *ContactUseCaseImpl) CreateGroup(ctx context.Context, group domain.Group) (int, error) {
	return c.repository.CreateGroup(ctx, group)
}

func (c *ContactUseCaseImpl) GetGroup(ctx context.Context, id int) (*domain.Group, error) {
	return c.repository.GetGroup(ctx, id)
}

func (c *ContactUseCaseImpl) AddContactToGroup(ctx context.Context, contactID, groupID int) error {
	return c.repository.AddContactToGroup(ctx, contactID, groupID)
}
