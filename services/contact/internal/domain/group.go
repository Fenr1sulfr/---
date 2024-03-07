package domain

import "fmt"

type Group struct {
	ID   int
	Name string
}

func NewGroup(id int, name string) (*Group, error) {
	if len(name) > 250 {
		return nil, fmt.Errorf("name is too long: %d characters (max 250)", len(name))
	}

	return &Group{
		ID:   id,
		Name: name,
	}, nil
}
