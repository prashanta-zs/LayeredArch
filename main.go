package main

import (
	"log"
	"net/http"
	"time"

	service "github.com/prashant/layeredArc/services/customer"

	"github.com/gorilla/mux"
	"github.com/prashant/layeredArc/drivers"
	handler "github.com/prashant/layeredArc/handlers/customer"
	store "github.com/prashant/layeredArc/stores/customer"
)

func main() {

	db := drivers.ConnectToSQL()

	stores := store.New(db)
	services := service.New(stores)
	h := handler.New(services)
	r := mux.NewRouter()

	r.HandleFunc("/customers", h.Create).Methods(http.MethodPost)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatalln(srv.ListenAndServe())

}
