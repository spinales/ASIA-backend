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

func (s *Server) AddCourseHandler(w http.ResponseWriter, r *http.Request) {
	var req models.Course
	json.NewDecoder(r.Body).Decode(&req)

	ID := r.Context().Value(PasetoKey).(*paseto.JSONToken).Issuer
	user, err := s.service.UserService.UserByTuition(ID)
	if err != nil {
		log.Println("Error in add course Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	if strings.ToLower(user.Role) != "admin" || strings.ToLower(user.Role) != "administrador" {
		util.RespondwithJSON(w, http.StatusUnauthorized, map[string]interface{}{"message": "you do not have permission to this functionality."})
		return
	}

	c, err := s.service.CourseService.AddCourse(&req)
	if err != nil {
		log.Println("Error in add course Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	util.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"course": map[string]interface{}{
		"ID":           c.ID,
		"Code":         c.Code,
		"Name":         c.Name,
		"Career":       c.Career,
		"Credits":      c.Credits,
		"AcademicArea": c.AcademicArea,
	}, "message": "ok"})
	return
}

func (s *Server) DeleteCourseHandler(w http.ResponseWriter, r *http.Request) {
	ID := r.Context().Value(PasetoKey).(*paseto.JSONToken).Issuer
	user, err := s.service.UserService.UserByTuition(ID)
	if err != nil {
		log.Println("Error in delete course Handler\n", err)
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
		log.Println("Error in delete course Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	err = s.service.CourseService.DeleteCourse(uint(num))
	if err != nil {
		log.Println("Error in delete course Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	util.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"message": "ok"})
	return
}

func (s *Server) UpdateCourseHandler(w http.ResponseWriter, r *http.Request) {
	var req models.Course
	json.NewDecoder(r.Body).Decode(&req)

	ID := r.Context().Value(PasetoKey).(*paseto.JSONToken).Issuer
	user, err := s.service.UserService.UserByTuition(ID)
	if err != nil {
		log.Println("Error in update course Handler\n", err)
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
		log.Println("Error in update course Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	c, err := s.service.CourseService.UpdateCourse(&req, uint(num))
	if err != nil {
		log.Println("Error in update course Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	util.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"course": map[string]interface{}{
		"ID":           c.ID,
		"Code":         c.Code,
		"Name":         c.Name,
		"Career":       c.Career,
		"Credits":      c.Credits,
		"AcademicArea": c.AcademicArea,
	}, "message": "ok"})
	return
}

func (s *Server) GetCoursesHandler(w http.ResponseWriter, r *http.Request) {
	courses, err := s.service.CourseService.Courses()
	if err != nil {
		log.Println("Error in get courses Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	util.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"message": "ok", "courses": courses})
	return
}

func (s *Server) GetCourseByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	num, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error in get course by id Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	course, err := s.service.CourseService.Course(uint(num))
	if err != nil {
		log.Println("Error in update course Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	util.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"message": "ok", "course": course})
	return
}

func (s *Server) GetCourseByNameHandler(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	course, err := s.service.CourseService.CourseByName(name)
	if err != nil {
		log.Println("Error in get course by name Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	util.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"message": "ok", "course": course})
	return
}

func (s *Server) GetCourseByCodeHandler(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	course, err := s.service.CourseService.CourseByCode(code)
	if err != nil {
		log.Println("Error in get course by code Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	util.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"message": "ok", "course": course})
	return
}
