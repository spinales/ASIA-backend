package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/spinales/ASIA-backend/models"
	"github.com/spinales/ASIA-backend/util"
)

func (s *Server) login(w http.ResponseWriter, r *http.Request) {
	var req models.User
	json.NewDecoder(r.Body).Decode(&req)

	user, err := s.service.UserService.UserByTuition(req.Tuition)
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
		return
	}

	if user.Password != req.Password {
		util.RespondWithError(w, http.StatusOK, "The user tuition or password is incorrect, try again.")
		return
	}

	token, err := s.tokenMaker.CreateToken(user.Tuition, s.config.AccessTokenDuration)
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
		return
	}

	util.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"token": token, "user": map[string]interface{}{
		"createdAt":     user.CreatedAt,
		"updatedAt":     user.UpdatedAt,
		"deletedAt":     user.DeletedAt,
		"tuition":       user.Tuition,
		"password":      user.Password,
		"Firstname":     user.Firstname,
		"Lastname":      user.Lastname,
		"Age":           user.Age,
		"InsituteEmail": user.InsituteEmail,
		"Status":        user.Status,
		"Role":          user.Role,
	}, "message": "OK"})
}

func (s *Server) register(w http.ResponseWriter, r *http.Request) {
	var req models.User
	json.NewDecoder(r.Body).Decode(&req)

	user, _ := s.service.UserService.UserByTuition(req.Tuition)
	if user.Password != req.Password {
		util.RespondWithError(w, http.StatusUnauthorized, "The user tuition or password is incorrect, try again.")
		return
	}

	if user.Status != "Pending" {
		util.RespondWithError(w, http.StatusOK, "You can't login using this method, use Log In.")
		return
	}

	token, err := s.tokenMaker.CreateToken(user.Tuition, s.config.AccessTokenDuration)
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
		return
	}

	if user.Status == "Pending" {
		util.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"token": token, "user": map[string]interface{}{
			"createdAt":     user.CreatedAt,
			"updatedAt":     user.UpdatedAt,
			"deletedAt":     user.DeletedAt,
			"tuition":       user.Tuition,
			"password":      user.Password,
			"Firstname":     user.Firstname,
			"Lastname":      user.Lastname,
			"Age":           user.Age,
			"InsituteEmail": user.InsituteEmail,
			"Status":        user.Status,
			"Role":          user.Role,
		}, "message": "Your user is incomplete, please complete the info in settings."})
		return
	}

	util.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"token": token, "user": map[string]interface{}{
		"createdAt":     user.CreatedAt,
		"updatedAt":     user.UpdatedAt,
		"deletedAt":     user.DeletedAt,
		"tuition":       user.Tuition,
		"password":      user.Password,
		"Firstname":     user.Firstname,
		"Lastname":      user.Lastname,
		"Age":           user.Age,
		"InsituteEmail": user.InsituteEmail,
		"Status":        user.Status,
		"Role":          user.Role,
	}, "message": "OK"})

	return
}
