package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/o1egl/paseto"
	"github.com/spinales/ASIA-backend/util"
)

func (s *Server) AcademicRecordHandler(w http.ResponseWriter, r *http.Request) {
	ID := r.Context().Value(PasetoKey).(*paseto.JSONToken).Issuer
	user, err := s.service.UserService.UserByTuition(ID)
	if err != nil {
		log.Println("Error in academic record Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	if strings.ToLower(user.Role) != "estudiante" || user.Role != "Estudiante" {
		util.RespondwithJSON(w, http.StatusUnauthorized, map[string]interface{}{"message": "you do not have permission to this functionality."})
		return
	}

	student, err := s.service.StudentService.StudentByUserID(user.ID)
	if err != nil {
		log.Println("Error in academic record Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	rats, err := s.service.RatingService.Ratings(student.ID)
	if err != nil {
		log.Println("Error in academic record Handler\n", err)
		util.RespondWithError(w, http.StatusInternalServerError, "An error has occurred on the server, try later.")
		return
	}

	util.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"message": "ok", "ratings": rats})
}
