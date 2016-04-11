package main

import (
	"strconv"

	"github.com/AEPi-AK/character-server/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Gets the next auto-incremented player num.
func GetNextCharacterNum() int {
	var counter models.Counter

	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"count": 1}},
		ReturnNew: true,
	}

	_, err := DB.C("counter").Find(bson.M{"_id": "isaacsucks"}).Apply(change, &counter)

	if err != nil {
		panic(err)
	}

	return counter.Seq
}

// Updates a character given some update request
func UpdateCharacter(request UpdateRequest) models.Character {
	character, err := FindCharacter(request.ID)

	if err != nil {
		panic(err)
	}

	if request.Gold != 0 {
		character.Gold = request.Gold
	}

	if request.ProID != "" {
		character.ProID = HashString(request.ProID)
	}

	if request.Experience != 0 {
		character.Experience = request.Experience
	}

	err = DB.C("characters").Update(bson.M{"_id": character.ID}, character)
	if err != nil {
		panic(err)
	}

	return character
}

// Returns the player number given some id (pro or regular). If the id is not
// in the DB, we return an error.
func PlayerNumForID(id string) (int, error) {
	result := models.Character{}
	err := DB.C("characters").Find(bson.M{"_id": id}).One(&result)
	if err != nil {
		err = DB.C("characters").Find(bson.M{"pro_id": id}).One(&result)
		if err != nil {
			return 0, err
		}
	}

	return result.PlayerNum, nil
}

// Creates a new character given some string of data. The data is used to
// generate the hash for the _id.
func CreateNewCharacter(data string) models.Character {
	char := models.Character{ID: HashString(data), PlayerNum: GetNextCharacterNum()}
	err := DB.C("characters").Insert(&char)

	if err != nil {
		panic(err)
	}

	return char
}

// Finds a character given some identifier string. This string can either
// be the _id, pro_id, or num of the character. If none exists, an error
// is returned.
func FindCharacter(identifier string) (models.Character, error) {
	var result models.Character

	i, err := strconv.Atoi(identifier)

	// If identifier is an integer, find character with that player number
	if err == nil {
		err = DB.C("characters").Find(bson.M{"num": i}).One(&result)
	} else {
		err = DB.C("characters").Find(bson.M{"_id": identifier}).One(&result)
		if err != nil {
			err = DB.C("characters").Find(bson.M{"pro_id": identifier}).One(&result)
		}
	}

	if err != nil {
		panic(err)
	}

	return result, nil
}