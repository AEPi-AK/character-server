package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// Represents a character creation request.
type CreateRequest struct {
	Data string `json:"data"`
	Race string `json:"race"`
	Strength int `json:"strength"`
	Dexterity int `json:"dexterity"`
	Wisdom int `json:"wisdom"`
}

type UpdateRequest struct {
	ID         string `json:"_id"`
	ProID      string `json:"pro_id"`
	Experience int    `json:"experience"`
	Gold       int    `json:"gold"`
}

func RespondBadRequest(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}

// Handler for character creation
// ENDPOINT: /characters/create
func CharacterCreate(w http.ResponseWriter, r *http.Request) {
	var requestData CreateRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		RespondBadRequest(w, err)
		return
	}
	if err := r.Body.Close(); err != nil {
		RespondBadRequest(w, err)
		return
	}
	if err := json.Unmarshal(body, &requestData); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	character, err := CreateNewCharacter(requestData)
	
	if err != nil {
		RespondBadRequest(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(character); err != nil {
		RespondBadRequest(w, err)
		return
	}
}

// Handler for character updating
// ENDPOINT: /characters/update
func CharacterUpdate(w http.ResponseWriter, r *http.Request) {
	var requestData UpdateRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		RespondBadRequest(w, err)
		return
	}
	if err := r.Body.Close(); err != nil {
		RespondBadRequest(w, err)
		return
	}
	if err := json.Unmarshal(body, &requestData); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	character, err := UpdateCharacter(requestData)

	if err != nil {
		RespondBadRequest(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(character); err != nil {
		RespondBadRequest(w, err)
		return
	}
}

// Handler for getting a character
// ENDPOINT: /character/{identifier}
func CharacterShow(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	identifier := vars["identifier"]

	result, err := FindCharacter(identifier)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resultString, _ := json.Marshal(result)
	fmt.Fprintln(w, "", string(resultString))
}
