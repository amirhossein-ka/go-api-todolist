package mux

import (
	"context"
	"go-api-todolist/controller"
	"go-api-todolist/service"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type handler struct {
	service service.Service
}

type rest struct {
	handler *handler
	server  *http.Server
	router  *mux.Router
}

func (r *rest) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "application/json")
	r.router.ServeHTTP(w, req)
}

func New(s service.Service) controller.Rest {
	return &rest{
		router: mux.NewRouter(),
		handler: &handler{
			service: s,
		},
	}
}

// Start implements controller.Rest
func (r *rest) Start(addr string) error {
	log.Println("Starting server on:", addr)
	r.routing()
	r.server = &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	return r.server.ListenAndServe()
}

// Stop implements controller.Rest
func (r *rest) Stop() error {
	log.Println("Stopping server...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return r.server.Shutdown(ctx)
}
