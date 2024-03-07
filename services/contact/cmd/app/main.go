package main

import (
	"architecture_go/pkg/store/postgres"
	"architecture_go/services/contact/internal/delivery"
	"architecture_go/services/contact/internal/repository"
	"architecture_go/services/contact/internal/usecase"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting the application...")
	conn, err := postgres.Connect("localhost", 5432, "postgres", "1112", "Aizel")
	log.Println("Connected to the database...")

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	repository := repository.NewContactRepository(conn)
	useCase := usecase.NewContactUseCase(repository)
	delivery := delivery.NewContactDelivery(useCase)

	http.HandleFunc("/contacts/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			delivery.CreateContact(w, r)
		case http.MethodGet:
			delivery.GetContact(w, r)
		case http.MethodPut:
			delivery.UpdateContact(w, r)
		case http.MethodDelete:
			delivery.DeleteContact(w, r)
		}
	})

	http.HandleFunc("/groups/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			delivery.CreateGroup(w, r)
		} else if r.Method == http.MethodGet {
			delivery.GetGroup(w, r)
		}
	})

	http.HandleFunc("/addContactToGroup", delivery.AddContactToGroup)

	log.Println("Application is up and running on http://localhost:8080/...")
	http.ListenAndServe(":8080", nil)
}
