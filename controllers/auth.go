package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ivanglzr/PaintBackend/db"
	"github.com/ivanglzr/PaintBackend/models"
	"github.com/ivanglzr/PaintBackend/utils"
)

func LogIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var login models.Login

	err := json.NewDecoder(r.Body).Decode(&login)

	if err != nil {
		utils.JSONResponse(w, 422, "Login wasn't valid")

		return
	}

	row := db.DB.QueryRow("SELECT id, password FROM users WHERE email = $1", login.Email)

	var id string
	var hash string

	err = row.Scan(&id, &hash)

	if err != nil {
		if err == sql.ErrNoRows {
			utils.JSONResponse(w, 404, "User not found")

			return
		}

		utils.JSONResponse(w, 500, "An error ocurred while finding the user")

		return
	}

	isPasswordValid := utils.ComparePassword(hash, login.Password)

	if !isPasswordValid {
		utils.JSONResponse(w, 401, "Log in unauthorized")

		return
	}

	token, err := utils.GenerateToken(id)

	cookie := utils.GenerateCookie(token)

	http.SetCookie(w, &cookie)

	utils.JSONResponse(w, 200, "Log in authorized")
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		utils.JSONResponse(w, 422, "User wasn't valid")

		return
	}

	hash, err := utils.HashPassword(user.Password)

	if err != nil {
		utils.JSONResponse(w, 500, "An error ocurred while hashing the password")

		return
	}

	err = db.DB.QueryRow("SELECT * FROM users WHERE email = $1", user.Email).Scan()

	if err != sql.ErrNoRows {
		utils.JSONResponse(w, 409, "User already exists")

		return
	}

	var id string

	err = db.DB.QueryRow("INSERT INTO users (fullname, email, password) VALUES ($1, $2, $3) RETURNING ID", user.Fullname, user.Email, hash).Scan(&id)

	log.Println(err)

	if err != nil {
		utils.JSONResponse(w, 500, "An error ocurred while creating the user")

		return
	}

	token, err := utils.GenerateToken(id)

	if err != nil {
		utils.JSONResponse(w, 500, "An error ocurred while creating the token")

		return
	}

	cookie := utils.GenerateCookie(token)

	http.SetCookie(w, &cookie)

	utils.JSONResponse(w, 201, "User created")
}
