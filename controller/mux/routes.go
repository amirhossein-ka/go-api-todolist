package mux

import "net/http"



func (r *rest) routing() {
	r.router = r.router.StrictSlash(true)
	api := r.router.PathPrefix("/api/v1").Subrouter()
	{
		api.HandleFunc("/create/", r.handler.create).Methods(http.MethodPost)
        api.HandleFunc("/get/{_id}/", r.handler.getOne).Methods(http.MethodGet)
        api.HandleFunc("/get/", r.handler.getAll).Methods(http.MethodGet)
        api.HandleFunc("/update/{_id}/", r.handler.update).Methods(http.MethodPut)
        api.HandleFunc("/delete/{_id}/", r.handler.delete).Methods(http.MethodDelete)
	}
}
