package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dileep9490/todoapp/Backend/database"
	"github.com/dileep9490/todoapp/Backend/models"
	"github.com/dileep9490/todoapp/Backend/utils"
	"github.com/dileep9490/todoapp/Backend/utils/types"
	"github.com/google/uuid"
)

func SignUP(response http.ResponseWriter, r *http.Request) {
	db := database.DB
	data := new(types.SignUpType)
	user := new(models.User)

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	hashpassword, err := utils.HashPassword(data.Password)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	} else {
		user.Password = hashpassword
	}

	user.Email = data.Email
	user.Name = data.Name
	user.ID = uuid.New()
	if err := db.Create(&user).Error; err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	userResponse, err := json.Marshal(user)
	if err != nil {
		fmt.Print("unable to marshal the user struct")
	}
	response.Header().Set("content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	response.Write(userResponse)

}

func Login(response http.ResponseWriter, r *http.Request) {
	data := new(types.LoginType)
	db := database.DB
	user := new(models.User)
	response.Header().Set("content-Type", "application/json")
	resp := make(map[string]string)

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		resp["error"] = "internal Server Error"
		response.WriteHeader(http.StatusInternalServerError)
		jsonResp, _ := json.Marshal(resp)
		response.Write(jsonResp)
		return
	}
	db.First(&user, "email=?", data.Email)

	if !(utils.ComparePassword(user.Password, data.Password)) {
		response.WriteHeader(http.StatusNotFound)
		resp["error"] = "credentials don't match"
		jsonResp, _ := json.Marshal(resp)
		response.Write(jsonResp)
		return
	}
	resp["apiKey"] = user.ID.String()
	response.WriteHeader(http.StatusOK)
	jsonResp, _ := json.Marshal(resp)
	response.Write(jsonResp)
}
