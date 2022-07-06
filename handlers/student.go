package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/o1egl/paseto"
	"github.com/spinales/ASIA-backend/models"
	"github.com/spinales/ASIA-backend/util"
)

func (s *Server) AddStudentHandler(w http.ResponseWriter, r *http.Request) {
	var req models.Student
	json.NewDecoder(r.Body).Decode(&req)

	ID := r.Context().Value(PasetoKey).(*paseto.JSONToken).Issuer
	user, err := s.service.UserService.UserByTuition(ID)
	if err != nil {
		log.Println("Error in add user Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	if strings.ToLower(user.Role) != "admin" || strings.ToLower(user.Role) != "administrador" {
		util.RespondwithJSON(w, http.StatusUnauthorized, map[string]interface{}{"message": "you do not have permission to this functionality."})
		return
	}

	u, err := s.service.StudentService.AddStudent(&req)
	if err != nil {
		log.Println("Error in add student Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	util.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"student": map[string]interface{}{
		"ID":                  u.ID,
		"UserID":              u.UserId,
		"Career":              u.Career,
		"TrimestrerCompleted": u.TrimesterCompleted,
		"Pensum":              u.Pensum,
		"State":               u.State,
		"QuartelyIndex":       u.QuarterlyIndex,
		"GeneralIndex":        u.GeneralIndex,
	}, "message": "ok"})
	return
}

func (s *Server) DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	ID := r.Context().Value(PasetoKey).(*paseto.JSONToken).Issuer
	user, err := s.service.UserService.UserByTuition(ID)
	if err != nil {
		log.Println("Error in delete student Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	if strings.ToLower(user.Role) != "admin" || strings.ToLower(user.Role) != "administrador" {
		util.RespondwithJSON(w, http.StatusUnauthorized, map[string]interface{}{"message": "you do not have permission to this functionality."})
		return
	}

	id := chi.URLParam(r, "id")
	num, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error in delete student Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	err = s.service.StudentService.DeleteStudent(uint(num))
	if err != nil {
		log.Println("Error in delete student Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	util.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"message": "ok"})
	return
}

func (s *Server) UpdateStudentHandler(w http.ResponseWriter, r *http.Request) {
	var req models.Student
	json.NewDecoder(r.Body).Decode(&req)

	ID := r.Context().Value(PasetoKey).(*paseto.JSONToken).Issuer
	user, err := s.service.UserService.UserByTuition(ID)
	if err != nil {
		log.Println("Error in update student Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	if strings.ToLower(user.Role) != "admin" || strings.ToLower(user.Role) != "administrador" {
		util.RespondwithJSON(w, http.StatusUnauthorized, map[string]interface{}{"message": "you do not have permission to this functionality."})
		return
	}

	id := chi.URLParam(r, "id")
	num, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error in update student Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	u, err := s.service.StudentService.UpdateStudent(&req, uint(num))
	if err != nil {
		log.Println("Error in update student Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	util.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"student": map[string]interface{}{
		"ID":                  u.ID,
		"UserID":              u.UserId,
		"Career":              u.Career,
		"TrimestrerCompleted": u.TrimesterCompleted,
		"Pensum":              u.Pensum,
		"State":               u.State,
		"QuartelyIndex":       u.QuarterlyIndex,
		"GeneralIndex":        u.GeneralIndex,
	}, "message": "ok"})
	return
}

func (s *Server) GetStudentsHandler(w http.ResponseWriter, r *http.Request) {
	ID := r.Context().Value(PasetoKey).(*paseto.JSONToken).Issuer
	user, err := s.service.UserService.UserByTuition(ID)
	if err != nil {
		log.Println("Error in get students Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	if strings.ToLower(user.Role) != "admin" || strings.ToLower(user.Role) != "administrador" {
		util.RespondwithJSON(w, http.StatusUnauthorized, map[string]interface{}{"message": "you do not have permission to this functionality."})
		return
	}

	stus, err := s.service.StudentService.Students()
	if err != nil {
		log.Println("Error in get students Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	util.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"message": "ok", "students": stus})
	return
}

func (s *Server) GetStudentByIDHandler(w http.ResponseWriter, r *http.Request) {
	ID := r.Context().Value(PasetoKey).(*paseto.JSONToken).Issuer
	user, err := s.service.UserService.UserByTuition(ID)
	if err != nil {
		log.Println("Error in get student by ID Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	if strings.ToLower(user.Role) != "admin" || strings.ToLower(user.Role) != "administrador" {
		util.RespondwithJSON(w, http.StatusUnauthorized, map[string]interface{}{"message": "you do not have permission to this functionality."})
		return
	}

	id := chi.URLParam(r, "id")
	num, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error in get student by id Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	u, err := s.service.StudentService.Student(uint(num))
	if err != nil {
		log.Println("Error in get student by ID Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	util.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"message": "ok", "student": map[string]interface{}{
		"ID":                  u.ID,
		"UserID":              u.UserId,
		"Career":              u.Career,
		"TrimestrerCompleted": u.TrimesterCompleted,
		"Pensum":              u.Pensum,
		"State":               u.State,
		"QuartelyIndex":       u.QuarterlyIndex,
		"GeneralIndex":        u.GeneralIndex,
	}})
	return
}

func (s *Server) GetRankingHandler(w http.ResponseWriter, r *http.Request) {
	ranking, err := s.service.StudentService.Ranking()
	if err != nil {
		log.Println("Error in get student by ID Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	util.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"message": "ok", "ranking": ranking})
	return
}
