package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ivanglzr/PaintBackend/db"
	"github.com/ivanglzr/PaintBackend/models"
	"github.com/ivanglzr/PaintBackend/utils"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hash), err
}

func comparePassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	isValid := err == nil

	return isValid
}

func LogIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var login models.Login

	err := json.NewDecoder(r.Body).Decode(&login)

	if err != nil {
		utils.JSONResponse(w, 422, "Login wasn't valid")

		return
	}

	var hash string

	err = db.DB.QueryRow("SELECT password FROM users WHERE email = $1", login.Email).Scan(&hash)

	if err != nil {
		utils.JSONResponse(w, 500, "An error ocurred while finding the user")

		return
	}

	isPasswordValid := comparePassword(hash, login.Password)

	if !isPasswordValid {
		utils.JSONResponse(w, 401, "Log in unauthorized")

		return
	}

	//TODO: send an auth cookie
	utils.JSONResponse(w, 200, "Log in authorized")
}
