package rest

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/taimoor99/assignment/utills"
)

func NewRouter(r *chi.Mux, ctrl MessageControllers){

	r.Group(func(rr chi.Router) {
		rr.Use(middleware.Logger)
		rr.Get("/message/{id}", ctrl.GetMessageDetailsByIdHandler)
		rr.Post("/message", ctrl.PostMessageHandler)
		rr.Delete("/message/{id}", ctrl.DeleteMessageHandler)
		rr.Get("/messages/{limit}/{offset}", ctrl.GetAllMessagesHandler)
		rr.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			utills.WriteJsonRes(writer, 200, nil, "Qlik's assignment!")
		})
	})

	return
}
