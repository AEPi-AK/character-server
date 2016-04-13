package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

// Represents a character creation request.
type CreateRequest struct {
	ID        string `json:"id"`
	Race      string `json:"race"`
	Strength  int    `json:"strength"`
	Dexterity int    `json:"dexterity"`
	Wisdom    int    `json:"wisdom"`
}

type UpdateRequest struct {
	ID         string `json:"id"`
	ProID      string `json:"pro_id"`
	Experience int    `json:"points"`
	Gold       int    `json:"gold"`
	Name       string `json:"name"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// Given a message, responds with a JSON object containing that message 
// as an error string/
func RespondBadRequest(w http.ResponseWriter, message string) {
	log.WithFields(log.Fields{
		"time":    time.Now(),
		"message": message,
	}).Error("Received a bad request")
	errorResponse := ErrorResponse{Error: message}
	http.Error(w, "", http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(errorResponse)
}

func CharactersByPoints(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"time": time.Now(),
	}).Info("Received characters sorted by points request")

	characters, err := CharactersInPointOrder()

	if err != nil {
		RespondBadRequest(w, "unable to query characters by points")
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(characters); err != nil {
		RespondBadRequest(w, err.Error())
		return
	}

}

// Handler for character creation
// ENDPOINT: /characters/create
func CharacterCreate(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"time": time.Now(),
	}).Info("Received character create request")

	var requestData CreateRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		RespondBadRequest(w, err.Error())
		return
	}

	if err := r.Body.Close(); err != nil {
		RespondBadRequest(w, err.Error())
		return
	}

	if err := json.Unmarshal(body, &requestData); err != nil {
		w.WriteHeader(422) // unprocessable entity
		fmt.Println("HERE!")
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	character, err := CreateNewCharacter(requestData)

	if err != nil {
		RespondBadRequest(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(character); err != nil {
		RespondBadRequest(w, err.Error())
		return
	}

}

// Handler for character updating
// ENDPOINT: /characters/update
func CharacterUpdate(w http.ResponseWriter, r *http.Request) {
	var requestData UpdateRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		RespondBadRequest(w, err.Error())
		return
	}
	if err := r.Body.Close(); err != nil {
		RespondBadRequest(w, err.Error())
		return
	}
	if err := json.Unmarshal(body, &requestData); err != nil {
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			RespondBadRequest(w, "Bad character update request")
			return
		}
	}

	character, err := UpdateCharacter(requestData)
	if err != nil {
		RespondBadRequest(w, "Error updating character")
		return
	}
	log.WithFields(log.Fields{
		"time": time.Now(),
		"id":   character.ID,
	}).Info("Received character update request")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(character); err != nil {
		RespondBadRequest(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Handler for getting a character
// ENDPOINT: /character/{identifier}
func CharacterShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	identifier := vars["identifier"]

	log.WithFields(log.Fields{
		"time": time.Now(),
		"id":   identifier,
	}).Info("Received character display request")

	result, err := FindCharacter(identifier)

	if err != nil {
		log.WithFields(log.Fields{
			"time": time.Now(),
			"id":   identifier,
		}).Warn("Character not found")

		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resultString, _ := json.Marshal(result)
	fmt.Fprintln(w, "", string(resultString))
}
