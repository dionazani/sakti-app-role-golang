package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func InsertAppRole(response http.ResponseWriter, request *http.Request) {

	var model appRoleModel

	err := json.NewDecoder(request.Body).Decode(&model)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	id := model.Id
	code := model.Code
	description := model.Description

	var myResponse myResponse

	if err != nil {
		myResponse.Status = "002"
		myResponse.Message = ""

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(myResponse)

		return
	}

	var result string = Insert(id, code, description)
	if len(result) > 0 {
		myResponse.Status = "001"
		myResponse.Message = result

	} else {
		myResponse.Status = "000"
		myResponse.Message = ""
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(myResponse)

	return
}

func GetAppRoleById(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	var myResponse myResponse
	myResponse.Status = "000"
	myResponse.Message = ""

	if err != nil {
		myResponse.Status = "002"
		myResponse.Message = ""

		fmt.Println(err)
	}

	var result = GetById(id)
	myResponse.Data = result

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(myResponse)
}
