package delivery

import (
	"net/http"
)

type ContactDelivery interface {
	CreateContact(w http.ResponseWriter, r *http.Request)
	GetContact(w http.ResponseWriter, r *http.Request)
	UpdateContact(w http.ResponseWriter, r *http.Request)
	DeleteContact(w http.ResponseWriter, r *http.Request)

	CreateGroup(w http.ResponseWriter, r *http.Request)
	GetGroup(w http.ResponseWriter, r *http.Request)
	AddContactToGroup(w http.ResponseWriter, r *http.Request)
}
