package handler

import (
	"encoding/json"
	"fmt"
	"mygo/http/oapi"
	"net/http"
)

// GetMyTasks Get: /my_tasks
func (h *Handler) GetMyTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, err := h.db.User.Get(ctx, 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	tasks, err := user.QueryAssignTasks().All(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprint(w, tasks)
}

func (h *Handler) PostTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var p oapi.Task
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.useCase.CreateTask(ctx, &p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "ok")
}