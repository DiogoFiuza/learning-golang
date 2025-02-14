package handlers

import (
	"encoding/json"
	"github.com/DiogoFiuza/learning-golang/APIs/internal/dto"
	"github.com/DiogoFiuza/learning-golang/APIs/internal/entity"
	"github.com/DiogoFiuza/learning-golang/APIs/internal/infra/database"
	"github.com/go-chi/jwtauth"
	"net/http"
	"time"
)

type UserHandler struct {
	UserDB      database.UserInterface
	JWT         *jwtauth.JWTAuth
	JWTExpireIn int
}

func NewUserHandler(db database.UserInterface, jwt *jwtauth.JWTAuth, JWTExpireIn int) *UserHandler {
	return &UserHandler{
		UserDB:      db,
		JWT:         jwt,
		JWTExpireIn: JWTExpireIn,
	}
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var jwt dto.GetJwtDTO
	err := json.NewDecoder(r.Body).Decode(&jwt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindByEmail(jwt.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !u.ValidatePassword(jwt.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := h.JWT.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JWTExpireIn)).Unix(),
	})

	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserDTO
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
