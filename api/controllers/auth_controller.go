package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Clareand/rest-api/api/auth"
	"github.com/Clareand/rest-api/api/helpers"
	"github.com/Clareand/rest-api/api/models"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Register(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.Error(w, http.StatusUnprocessableEntity, err)
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		helpers.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("")
	if err != nil {
		helpers.Error(w, http.StatusUnprocessableEntity, err)
	}
	userCreated, err := user.SaveUser(server.DB)

	if err != nil {
		formatedError := helpers.FormatError(err.Error())
		helpers.Error(w, http.StatusUnprocessableEntity, formatedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	helpers.JSON(w, http.StatusCreated, userCreated)

}

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		helpers.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("login")

	if err != nil {
		helpers.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formatedError := helpers.FormatError(err.Error())
		helpers.Error(w, http.StatusUnprocessableEntity, formatedError)
		return
	}
	data := map[string]string{
		"token": token,
		"email": user.Email,
		"name":  user.Name,
	}
	helpers.JSON(w, http.StatusOK, data)
}

func (server *Server) SignIn(email, password string) (string, error) {
	var err error
	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email=?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = helpers.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}
