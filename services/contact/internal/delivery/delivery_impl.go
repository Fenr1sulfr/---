package delivery

import (
	"architecture_go/services/contact/internal/domain"
	"architecture_go/services/contact/internal/usecase"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type ContactDeliveryImpl struct {
	usecase usecase.ContactUseCase
}

func NewContactDelivery(usecase usecase.ContactUseCase) ContactDelivery {
	return &ContactDeliveryImpl{
		usecase: usecase,
	}
}

func (c *ContactDeliveryImpl) CreateContact(w http.ResponseWriter, r *http.Request) {
	var contact domain.Contact
    
    ctx := r.Context()
    log.Default().Printf("[CreateContact] Request recieved, contact: %+v", contact)

    if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
        log.Fatal(err);
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    id, err := c.usecase.CreateContact(ctx, contact)

	if err != nil {
        log.Fatal(err);
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

    log.Default().Printf("[CreateContact] Contact created with id: %d", id)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

func (cd *ContactDeliveryImpl) GetContact(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/contacts/")
    id, err := strconv.Atoi(idStr)

    log.Default().Printf("[GetContact] Request recieved, contact id: %d", id)

    ctx := r.Context()

    if err != nil {
        log.Fatal(err);
        http.Error(w, "Contact ID should be integer value!", http.StatusBadRequest)
        return
    }

    contact, err := cd.usecase.GetContact(ctx, id)
    if err != nil {
        log.Fatal(err);
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    log.Default().Printf("[GetContact] Contact sent: %+v", contact)

    json.NewEncoder(w).Encode(contact)
}

func (cd *ContactDeliveryImpl) UpdateContact(w http.ResponseWriter, r *http.Request) {
    var contact domain.Contact
    ctx := r.Context()

    log.Default().Printf("[UpdateContact] Request recieved, contact: %+v", contact)

    if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
        log.Fatal(err);
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err := cd.usecase.UpdateContact(ctx, contact)

    if err != nil {
        log.Fatal(err);
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    log.Default().Printf("[UpdateContact] Contact updated: %+v", contact)

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(contact)
}

func (cd *ContactDeliveryImpl) DeleteContact(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/contacts/")
    id, err := strconv.Atoi(idStr)

    ctx := r.Context()

    log.Default().Printf("[DeleteContact] Request recieved, contact id: %d", id)

    if err != nil {
        log.Fatal(err);
        http.Error(w, "Contact ID should be integer value!", http.StatusBadRequest)
        return
    }

    err = cd.usecase.DeleteContact(ctx, id)

    if err != nil {
        log.Fatal(err);
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    log.Default().Printf("[DeleteContact] Contact with id: %d deleted.", id)

    w.WriteHeader(http.StatusNoContent)
}

func (cd *ContactDeliveryImpl) CreateGroup(w http.ResponseWriter, r *http.Request) {
    var group domain.Group

    ctx := r.Context()

    log.Default().Printf("[CreateGroup] Request recieved, group: %+v", group)

    if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
        log.Fatal(err);
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    id, err := cd.usecase.CreateGroup(ctx, group)

    if err != nil {
        log.Fatal(err);
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    log.Default().Printf("[CreateGroup] Group created with id: %d", id)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]int{"id": id})
}

func (cd *ContactDeliveryImpl) GetGroup(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/groups/")
    id, err := strconv.Atoi(idStr)

    ctx := r.Context()

    log.Default().Printf("[GetGroup] Request recieved, group id: %d", id)

    if err != nil {
        log.Fatal(err);
        http.Error(w, "Invalid group ID", http.StatusBadRequest)
        return
    }

    group, err := cd.usecase.GetGroup(ctx, id)

    if err != nil {
        log.Fatal(err);
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    log.Default().Printf("[GetGroup] Group sent: %+v", group)

    json.NewEncoder(w).Encode(group)
}

func (cd *ContactDeliveryImpl) AddContactToGroup(w http.ResponseWriter, r *http.Request) {
    var request struct {
        ContactID int `json:"contact_id"`
        GroupID   int `json:"group_id"`
    }

    ctx := r.Context()

    log.Default().Printf("[AddContactToGroup] Request recieved, contact_id: %d, group_id: %d", request.ContactID, request.GroupID)

    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        log.Fatal(err);
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err := cd.usecase.AddContactToGroup(ctx, request.ContactID, request.GroupID)

    if err != nil {
        log.Fatal(err);
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    log.Default().Printf("[AddContactToGroup] Contact added to group successfully")

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"result": "Contact added to group successfully"})
}
