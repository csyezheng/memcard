package services

import (
	"github.com/csyezheng/memcard/internal/auth/models"
	"github.com/csyezheng/memcard/pkg/database"
	"github.com/csyezheng/memcard/pkg/httputils"
	"github.com/csyezheng/memcard/pkg/logging"
	"net/http"
	"time"
)

func (s *Service) Signup(w http.ResponseWriter, r *http.Request) {
	type RequestData struct {
		Username  string `forms:"username"`
		Email     string `forms:"email"`
		FirstName string `forms:"firstName"`
		LastName  string `forms:"lastName"`
		Password  string `forms:"password"`
	}

	ctx := r.Context()
	logger := logging.FromContext(ctx)

	var data RequestData
	if statusCode, err := httputils.BindJSON(w, r, &data); err != nil {
		logger.Error(err.Error())
		s.JSON(w, statusCode, err)
		return
	}

	// See if the user already exists and use that record.
	user, err := s.authDB.FindUserByEmailAndUsername(ctx, data.Username, data.Email)
	if err != nil {
		if !database.IsNotFound(err) {
			s.JSON(w, http.StatusInternalServerError, err)
			return
		}
	} else {
		// User already exist
		s.JSON(w, http.StatusConflict, struct {
			Msg string `json:"msg"`
		}{Msg: "user already exist"})
		return
	}
	// User does not exist, create a new one.
	user = &models.User{
		Username:   data.Username,
		Email:      data.Email,
		FirstName:  data.FirstName,
		LastName:   data.LastName,
		Password:   data.Password,
		DateJoined: time.Now(),
	}
	if err := s.authDB.SaveUser(ctx, user); err != nil {
		if database.IsValidationError(err) {
			s.JSON(w, http.StatusUnprocessableEntity, err)
			return
		}
		s.JSON(w, http.StatusInternalServerError, err)
		return
	}
	s.JSON(w, http.StatusOK, nil)
}
