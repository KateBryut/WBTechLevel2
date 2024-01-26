package handler

import (
	"dev11/internal/models"
	"fmt"
	"net/http"
	"time"
)

func parceCreateEventRequest(w http.ResponseWriter, r *http.Request) (string, models.Event, error) {
	var event models.Event
	userId := r.PostFormValue("user_id")
	if userId == "" {
		res := responseError{Error: "parameter user_id missing"}
		JSONError(w, res, http.StatusBadRequest)
		return "", event, fmt.Errorf(res.Error)
	}
	name := r.PostFormValue("name")
	if name == "" {
		res := responseError{Error: "parameter name missing"}
		JSONError(w, res, http.StatusBadRequest)
		return "", event, fmt.Errorf(res.Error)
	}
	description := r.PostFormValue("description")
	if description == "" {
		res := responseError{Error: "parameter description missing"}
		JSONError(w, res, http.StatusBadRequest)
		return "", event, fmt.Errorf(res.Error)
	}
	dateString := r.PostFormValue("date")
	if dateString == "" {
		res := responseError{Error: "parameter date missing"}
		JSONError(w, res, http.StatusBadRequest)
		return "", event, fmt.Errorf(res.Error)
	}
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		res := responseError{Error: fmt.Sprintf("wrong format of date %s", err.Error())}
		JSONError(w, res, http.StatusBadRequest)
		return "", event, fmt.Errorf(res.Error)
	}
	event = models.Event{
		Name:        name,
		Date:        date,
		Description: description,
	}
	return userId, event, nil
}

func parceUpdateEventRequest(w http.ResponseWriter, r *http.Request) (string, models.Event, error) {
	var event models.Event
	userId := r.PostFormValue("user_id")
	if userId == "" {
		res := responseError{Error: "parameter user_id missing"}
		JSONError(w, res, http.StatusBadRequest)
		return "", event, fmt.Errorf(res.Error)
	}
	eventId := r.PostFormValue("event_id")
	if eventId == "" {
		res := responseError{Error: "parameter event_id missing"}
		JSONError(w, res, http.StatusBadRequest)
		return "", event, fmt.Errorf(res.Error)
	}
	name := r.PostFormValue("name")
	if name == "" {
		res := responseError{Error: "parameter name missing"}
		JSONError(w, res, http.StatusBadRequest)
		return "", event, fmt.Errorf(res.Error)
	}
	description := r.PostFormValue("description")
	if description == "" {
		res := responseError{Error: "parameter description missing"}
		JSONError(w, res, http.StatusBadRequest)
		return "", event, fmt.Errorf(res.Error)
	}
	dateString := r.PostFormValue("date")
	if dateString == "" {
		res := responseError{Error: "parameter date missing"}
		JSONError(w, res, http.StatusBadRequest)
		return "", event, fmt.Errorf(res.Error)
	}
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		res := responseError{Error: fmt.Sprintf("wrong format of date %s", err.Error())}
		JSONError(w, res, http.StatusBadRequest)
		return "", event, fmt.Errorf(res.Error)
	}
	event = models.Event{
		Id:          eventId,
		Name:        name,
		Date:        date,
		Description: description,
	}
	return userId, event, nil
}

func parceDeleteEventRequest(w http.ResponseWriter, r *http.Request) (string, models.Event, error) {
	var event models.Event
	userId := r.PostFormValue("user_id")
	if userId == "" {
		res := responseError{Error: "parameter user_id missing"}
		JSONError(w, res, http.StatusBadRequest)
		return "", event, fmt.Errorf(res.Error)
	}
	eventId := r.PostFormValue("event_id")
	if eventId == "" {
		res := responseError{Error: "parameter event_id missing"}
		JSONError(w, res, http.StatusBadRequest)
		return "", event, fmt.Errorf(res.Error)
	}
	event = models.Event{
		Id: eventId,
	}
	return userId, event, nil
}

func parceGetEvents(w http.ResponseWriter, r *http.Request) (string, time.Time, error) {
	var date time.Time
	userId := r.FormValue("user_id")
	if userId == "" {
		res := responseError{Error: "parameter user_id missing"}
		JSONError(w, res, http.StatusBadRequest)
		return "", date, fmt.Errorf(res.Error)
	}
	dateString := r.FormValue("date")
	if dateString == "" {
		res := responseError{Error: "parameter date missing"}
		JSONError(w, res, http.StatusBadRequest)
		return "", date, fmt.Errorf(res.Error)
	}
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		res := responseError{Error: fmt.Sprintf("wrong format of date %s", err.Error())}
		JSONError(w, res, http.StatusBadRequest)
		return "", date, fmt.Errorf(res.Error)
	}
	return userId, date, nil
}
