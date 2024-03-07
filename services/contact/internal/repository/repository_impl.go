package repository

import (
	"architecture_go/services/contact/internal/domain"
	"fmt"

	"context"

	"github.com/jackc/pgx/v5"
)

type ContactGroupAssociation struct {
	ContactID int
	GroupID   int
}

type ContactRepositoryImpl struct {
	conn *pgx.Conn
}

func NewContactRepository(conn *pgx.Conn) ContactRepository {
	return &ContactRepositoryImpl{
		conn: conn,
	}
}

func (repo *ContactRepositoryImpl) CreateContact(ctx context.Context, contact domain.Contact) (int, error) {
	sql := `INSERT INTO contacts (first_name, middle_name, last_name, phone_number) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int
	err := repo.conn.QueryRow(ctx, sql, contact.FirstName, contact.MiddleName, contact.LastName, contact.PhoneNumber).Scan(&id)
	
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (repo *ContactRepositoryImpl) GetContact(ctx context.Context, id int) (*domain.Contact, error) {
	sql := `SELECT id, first_name, middle_name, last_name, phone_number FROM contacts WHERE id = $1`
	var contact domain.Contact
	err := repo.conn.QueryRow(ctx, sql, id).Scan(&contact.ID, &contact.FirstName, &contact.MiddleName, &contact.LastName, &contact.PhoneNumber)
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

func (repo *ContactRepositoryImpl) UpdateContact(ctx context.Context, contact domain.Contact) error {
	sql := `UPDATE contacts SET first_name = $1, middle_name = $2, last_name = $3, phone_number = $4 WHERE id = $5`
	cmdTag, err := repo.conn.Exec(ctx, sql, contact.FirstName, contact.MiddleName, contact.LastName, contact.PhoneNumber, contact.ID)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("no rows affected")
	}
	return nil
}

func (repo *ContactRepositoryImpl) DeleteContact(ctx context.Context, id int) error {
	sql := `DELETE FROM contacts WHERE id = $1`
	cmdTag, err := repo.conn.Exec(ctx, sql, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("no rows affected")
	}
	return nil
}

func (repo *ContactRepositoryImpl) CreateGroup(ctx context.Context, group domain.Group) (int, error) {
	sql := `INSERT INTO groups (name) VALUES ($1) RETURNING id`
	var id int
	err := repo.conn.QueryRow(ctx, sql, group.Name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (repo *ContactRepositoryImpl) GetGroup(ctx context.Context, id int) (*domain.Group, error) {
	sql := `SELECT id, name FROM groups WHERE id = $1`
	var group domain.Group
	err := repo.conn.QueryRow(ctx, sql, id).Scan(&group.ID, &group.Name)

	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (repo *ContactRepositoryImpl) UpdateGroup(ctx context.Context, group domain.Group) error {
	sql := `UPDATE groups SET name = $1 WHERE id = $2`
	cmdTag, err := repo.conn.Exec(ctx, sql, group.Name, group.ID)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("no rows affected")
	}
	return nil
}

func (repo *ContactRepositoryImpl) DeleteGroup(ctx context.Context, id int) error {
	sql := `DELETE FROM groups WHERE id = $1`
	cmdTag, err := repo.conn.Exec(ctx, sql, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("no rows affected")
	}
	return nil
}

func (repo *ContactRepositoryImpl) AddContactToGroup(ctx context.Context, contactID, groupID int) error {
	sql := `INSERT INTO contact_group (contact_id, group_id) VALUES ($1, $2)`
	_, err := repo.conn.Exec(ctx, sql, contactID, groupID)
	return err
}

func (repo *ContactRepositoryImpl) RemoveContactFromGroup(ctx context.Context, contactID, groupID int) error {
	sql := `DELETE FROM contact_group WHERE contact_id = $1 AND group_id = $2`
	_, err := repo.conn.Exec(ctx, sql, contactID, groupID)
	return err
}

func (repo *ContactRepositoryImpl) GetContactsByGroup(ctx context.Context, groupID int) ([]*domain.Contact, error) {
	sql := `SELECT c.id, c.first_name, c.middle_name, c.last_name, c.phone_number 
	FROM contacts c
	INNER JOIN contact_group cg ON c.id = cg.contact_id
	WHERE cg.group_id = $1`
	rows, err := repo.conn.Query(ctx, sql, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []*domain.Contact
	for rows.Next() {
		var contact domain.Contact
		err = rows.Scan(&contact.ID, &contact.FirstName, &contact.MiddleName, &contact.LastName, &contact.PhoneNumber)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, &contact)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return contacts, nil
}
