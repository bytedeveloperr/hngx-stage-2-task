package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bytedeveloperr/hng-stage-2/api"
	"github.com/bytedeveloperr/hng-stage-2/db"
	"github.com/gorilla/mux"
)

type Server struct {
	address string
}

func NewServer(address string) *Server {
	return &Server{
		address: address,
	}
}

func (s *Server) Run() {
	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var body map[string]interface{}
			if r.ContentLength != 0 {
				err := json.NewDecoder(r.Body).Decode(&body)

				if err != nil {
					fmt.Println(err)
				}
			}

			Db, err := db.OpenDB()
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(api.ErrorAPIResponse(err.Error()))
			} else {
				ctx := context.WithValue(r.Context(), "apiContext", &api.Context{
					Body:   body,
					DB:     Db,
					Params: mux.Vars(r),
				})
				next.ServeHTTP(w, r.WithContext(ctx))
			}
		})
	})

	router.HandleFunc("", api.RequestHandler(api.CreatePerson)).Methods("POST")
	router.HandleFunc("", api.RequestHandler(api.GetPersons)).Methods("GET")
	router.HandleFunc("/{id}", api.RequestHandler(api.GetPerson)).Methods("GET")
	router.HandleFunc("/{id}", api.RequestHandler(api.UpdatePerson)).Methods("PUT")
	router.HandleFunc("/{id}", api.RequestHandler(api.DeletePerson)).Methods("DELETE")

	http.ListenAndServe(s.address, router)
}
