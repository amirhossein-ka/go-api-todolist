package mux

import (
	"encoding/json"
	"fmt"
	"go-api-todolist/models"
	"go-api-todolist/repository/mongo"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *handler) create(w http.ResponseWriter, r *http.Request) {
	var t models.Todo
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		writeJson(w, http.StatusBadRequest, map[string]any{"status": false, "error": err.Error()})
		return
	}

	id, err := h.service.CreateTask(r.Context(), t)
	if err != nil {
		writeJson(w, http.StatusInternalServerError, map[string]any{"status": false, "error": err.Error()})
		return
	}

	writeJson(w, http.StatusCreated, map[string]any{"status": "created", "data": map[string]any{"id": id}})
}

func (h *handler) getOne(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	id := v["_id"]

	task, err := h.service.GetTask(r.Context(), id)
	if err != nil {
		var code int
		if err == mongo.ErrNoDocument {
			code = http.StatusNotFound
		} else {
			code = http.StatusInternalServerError
		}

		writeJson(w, code, map[string]any{"status": false, "error": err.Error()})
		return
	}

	writeJson(w, http.StatusOK, map[string]any{"status": "ok", "data": task})
}

func (h *handler) getAll(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.GetAllTasks(r.Context())

	if err != nil {
		var code int
		if err == mongo.ErrNoDocument {
			code = http.StatusNotFound
		} else {
			code = http.StatusInternalServerError
		}

		writeJson(w, code, map[string]any{"status": false, "error": err.Error()})
		return
	}

	writeJson(w, http.StatusOK, map[string]any{"status": "ok", "data": tasks})
}

func (h *handler) update(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	id, ok := v["_id"]
	if !ok {
		writeJson(w, http.StatusBadRequest, map[string]any{"status": false, "error": fmt.Errorf("_id parameter not found.")})
		return
	}
	var task models.Todo
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJson(w, http.StatusBadRequest, map[string]any{"status": false, "error": err.Error()})
		return
	}

	err := h.service.UpdateTask(r.Context(), id, task)
	if err != nil {
		var code int
		if err == mongo.ErrNoDocument {
			code = http.StatusNotFound
		} else {
			code = http.StatusInternalServerError
		}

		writeJson(w, code, map[string]any{"status": false, "error": err.Error()})
		return
	}

	writeJson(w, http.StatusOK, map[string]any{"status": "updated"})
}

func (h *handler) delete(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	id, ok := v["_id"]
	if !ok {
		writeJson(w, http.StatusBadRequest, map[string]any{"error": fmt.Errorf("_id parameter not found.")})
		return
	}

	err := h.service.DeleteTask(r.Context(), id)
	if err != nil {
		var code int
		if err == mongo.ErrNoDocument {
			code = http.StatusNotFound
		} else {
			code = http.StatusInternalServerError
		}

		writeJson(w, code, map[string]any{"error": err.Error()})
		return
	}

	writeJson(w, http.StatusOK, map[string]any{"status": "deleted"})
}
