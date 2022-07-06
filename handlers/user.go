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

func (s *Server) AddUserHandler(w http.ResponseWriter, r *http.Request) {
	var req models.User
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

	u, err := s.service.UserService.AddUser(&req)
	if err != nil {
		log.Println("Error in add user Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	util.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"user": map[string]interface{}{
		"ID":             u.ID,
		"Tuition":        u.Tuition,
		"Firstname":      u.Firstname,
		"Lastname":       u.Lastname,
		"Age":            u.Age,
		"InstituteEmail": u.InsituteEmail,
		"Status":         u.Status,
		"Role":           u.Role,
		"Nationality":    u.NationalityID,
	}, "message": "ok"})
	return
}

func (s *Server) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	ID := r.Context().Value(PasetoKey).(*paseto.JSONToken).Issuer
	user, err := s.service.UserService.UserByTuition(ID)
	if err != nil {
		log.Println("Error in delete user Handler\n", err)
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
		log.Println("Error in delete user Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	err = s.service.UserService.Deleteuser(uint(num))
	if err != nil {
		log.Println("Error in delete user Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	util.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"message": "ok"})
	return
}

func (s *Server) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var req models.User
	json.NewDecoder(r.Body).Decode(&req)

	ID := r.Context().Value(PasetoKey).(*paseto.JSONToken).Issuer
	user, err := s.service.UserService.UserByTuition(ID)
	if err != nil {
		log.Println("Error in update user Handler\n", err)
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
		log.Println("Error in update user Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	u, err := s.service.UserService.UpdateUser(&req, uint(num))
	if err != nil {
		log.Println("Error in update user Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	util.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"user": map[string]interface{}{
		"ID":             u.ID,
		"Tuition":        u.Tuition,
		"Firstname":      u.Firstname,
		"Lastname":       u.Lastname,
		"Age":            u.Age,
		"InstituteEmail": u.InsituteEmail,
		"Status":         u.Status,
		"Role":           u.Role,
		"Nationality":    u.NationalityID,
	}, "message": "ok"})
	return
}

func (s *Server) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	ID := r.Context().Value(PasetoKey).(*paseto.JSONToken).Issuer
	user, err := s.service.UserService.UserByTuition(ID)
	if err != nil {
		log.Println("Error in get user Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	if strings.ToLower(user.Role) != "admin" || strings.ToLower(user.Role) != "administrador" {
		util.RespondwithJSON(w, http.StatusUnauthorized, map[string]interface{}{"message": "you do not have permission to this functionality."})
		return
	}

	users, err := s.service.UserService.Users()
	if err != nil {
		log.Println("Error in get users Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	util.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"message": "ok", "users": users})
	return
}

func (s *Server) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	ID := r.Context().Value(PasetoKey).(*paseto.JSONToken).Issuer
	user, err := s.service.UserService.UserByTuition(ID)
	if err != nil {
		log.Println("Error in get user by ID Handler\n", err)
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
		log.Println("Error in get user by id Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	u, err := s.service.UserService.User(uint(num))
	if err != nil {
		log.Println("Error in get user by ID Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	util.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"message": "ok", "user": map[string]interface{}{
		"ID":             u.ID,
		"Tuition":        u.Tuition,
		"Firstname":      u.Firstname,
		"Lastname":       u.Lastname,
		"Age":            u.Age,
		"InstituteEmail": u.InsituteEmail,
		"Status":         u.Status,
		"Role":           u.Role,
		"Nationality":    u.NationalityID,
	}})
	return
}

func (s *Server) GetUserByFirstnameHandler(w http.ResponseWriter, r *http.Request) {
	ID := r.Context().Value(PasetoKey).(*paseto.JSONToken).Issuer
	user, err := s.service.UserService.UserByTuition(ID)
	if err != nil {
		log.Println("Error in get user by firstname Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	if strings.ToLower(user.Role) != "admin" || strings.ToLower(user.Role) != "administrador" {
		util.RespondwithJSON(w, http.StatusUnauthorized, map[string]interface{}{"message": "you do not have permission to this functionality."})
		return
	}

	firstname := chi.URLParam(r, "firstname")
	u, err := s.service.UserService.UserByFirstname(firstname)
	if err != nil {
		log.Println("Error in get user by firstname Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	util.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"message": "ok", "user": map[string]interface{}{
		"ID":             u.ID,
		"Tuition":        u.Tuition,
		"Firstname":      u.Firstname,
		"Lastname":       u.Lastname,
		"Age":            u.Age,
		"InstituteEmail": u.InsituteEmail,
		"Status":         u.Status,
		"Role":           u.Role,
		"Nationality":    u.NationalityID,
	}})
	return
}

func (s *Server) GetUserByTuitionHandler(w http.ResponseWriter, r *http.Request) {
	ID := r.Context().Value(PasetoKey).(*paseto.JSONToken).Issuer
	user, err := s.service.UserService.UserByTuition(ID)
	if err != nil {
		log.Println("Error in get user by tuition Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	if strings.ToLower(user.Role) != "admin" || strings.ToLower(user.Role) != "administrador" {
		util.RespondwithJSON(w, http.StatusUnauthorized, map[string]interface{}{"message": "you do not have permission to this functionality."})
		return
	}

	tuition := chi.URLParam(r, "tuition")
	u, err := s.service.UserService.UserByFirstname(tuition)
	if err != nil {
		log.Println("Error in get user by tuition Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	util.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"message": "ok", "user": map[string]interface{}{
		"ID":             u.ID,
		"Tuition":        u.Tuition,
		"Firstname":      u.Firstname,
		"Lastname":       u.Lastname,
		"Age":            u.Age,
		"InstituteEmail": u.InsituteEmail,
		"Status":         u.Status,
		"Role":           u.Role,
		"Nationality":    u.NationalityID,
	}})
	return
}
