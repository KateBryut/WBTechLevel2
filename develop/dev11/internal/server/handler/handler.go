package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"dev11/internal/models"
	"dev11/internal/storage"
)

type responseError struct {
	Error string `json:"error"`
}

type responseSuccess struct {
	Result string `json:"result"`
}

type responseSuccessArray struct {
	Result []models.Event `json:"result"`
}

type Handler struct {
	repo storage.Storager
}

func NewHandler(r storage.Storager) *Handler {
	return &Handler{r}
}

func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", h.createEvent)
	mux.HandleFunc("/update_event", h.updateEvent)
	mux.HandleFunc("/delete_event", h.deleteEvent)
	mux.HandleFunc("/events_for_day", h.getEventsForDay)
	mux.HandleFunc("/events_for_week", h.getEventsForWeek)
	mux.HandleFunc("/events_for_month", h.getEventsForMonth)

	handler := Logging(mux)
	return handler
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	})
}

func JSONError(w http.ResponseWriter, err responseError, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}

func JSONSuccess(w http.ResponseWriter, res interface{}, code int) {
	output, err := json.Marshal(res)
	if err != nil {
		res := responseError{Error: err.Error()}
		JSONError(w, res, http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(output)
}

func (h *Handler) createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		res := responseError{Error: fmt.Sprintf("method %s not supported", r.Method)}
		JSONError(w, res, http.StatusInternalServerError)
		return
	}

	userId, event, err := parceCreateEventRequest(w, r)
	if err != nil {
		return
	}
	eventId, err := h.repo.CreateEvent(event, userId)
	if err != nil {
		res := responseError{Error: err.Error()}
		JSONError(w, res, http.StatusServiceUnavailable)
		return
	}

	res := responseSuccess{Result: fmt.Sprintf("event was created, id: %s", eventId)}
	JSONSuccess(w, res, http.StatusOK)
}

func (h *Handler) updateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		res := responseError{Error: fmt.Sprintf("method %s not supported", r.Method)}
		JSONError(w, res, http.StatusInternalServerError)
		return
	}
	userId, event, err := parceUpdateEventRequest(w, r)
	if err != nil {
		return
	}
	err = h.repo.UpdateEvent(event, userId)
	if err != nil {
		res := responseError{Error: err.Error()}
		JSONError(w, res, http.StatusServiceUnavailable)
		return
	}
	res := responseSuccess{Result: fmt.Sprintf("event was updated, id: %s", event.Id)}
	JSONSuccess(w, res, http.StatusOK)
}

func (h *Handler) deleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		res := responseError{Error: fmt.Sprintf("method %s not supported", r.Method)}
		JSONError(w, res, http.StatusInternalServerError)
		return
	}
	userId, event, err := parceDeleteEventRequest(w, r)
	if err != nil {
		return
	}
	err = h.repo.DeleteEvent(event, userId)
	if err != nil {
		res := responseError{Error: err.Error()}
		JSONError(w, res, http.StatusServiceUnavailable)
		return
	}
	res := responseSuccess{Result: fmt.Sprintf("event was deleted, id: %s", event.Id)}
	JSONSuccess(w, res, http.StatusOK)
}

func (h *Handler) getEventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		res := responseError{Error: fmt.Sprintf("method %s not supported", r.Method)}
		JSONError(w, res, http.StatusInternalServerError)
		return
	}
	userId, date, err := parceGetEvents(w, r)
	if err != nil {
		return
	}
	events, err := h.repo.GetEventForDay(userId, date)
	if err != nil {
		res := responseError{Error: err.Error()}
		JSONError(w, res, http.StatusServiceUnavailable)
		return
	}
	res := responseSuccessArray{Result: events}
	JSONSuccess(w, res, http.StatusOK)
}

func (h *Handler) getEventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		res := responseError{Error: fmt.Sprintf("method %s not supported", r.Method)}
		JSONError(w, res, http.StatusInternalServerError)
		return
	}
	userId, date, err := parceGetEvents(w, r)
	if err != nil {
		return
	}
	events, err := h.repo.GetEventForWeek(userId, date)
	if err != nil {
		res := responseError{Error: err.Error()}
		JSONError(w, res, http.StatusServiceUnavailable)
		return
	}
	res := responseSuccessArray{Result: events}
	JSONSuccess(w, res, http.StatusOK)
}

func (h *Handler) getEventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		res := responseError{Error: fmt.Sprintf("method %s not supported", r.Method)}
		JSONError(w, res, http.StatusInternalServerError)
		return
	}
	userId, date, err := parceGetEvents(w, r)
	if err != nil {
		return
	}
	events, err := h.repo.GetEventForMonth(userId, date)
	if err != nil {
		res := responseError{Error: err.Error()}
		JSONError(w, res, http.StatusServiceUnavailable)
		return
	}
	res := responseSuccessArray{Result: events}
	JSONSuccess(w, res, http.StatusOK)
}
