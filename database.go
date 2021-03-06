package main

import (
	"strconv"
	"time"

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
func UpdateCharacter(request UpdateRequest) (models.Character, error) {
	character, err := FindCharacter(request.ID)

	if err != nil {
		return character, err
	}

	if request.ProID != "" {
		character.ProID = request.ProID
	}

	if request.Experience != 0 {
		character.Experience = request.Experience
	}

	if request.Name != "" {
		character.Name = request.Name
	}

	err = DB.C("characters").Update(bson.M{"_id": character.ID}, character)
	if err != nil {
		return character, err
	}

	return character, nil
}

func CharactersInPointOrder() ([]models.Character, error){
	var results []models.Character
	err := DB.C("characters").Find(nil).Sort("-points").All(&results)
	if err != nil {
		return results, err
	}

	return results, err
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
// generate the for the _id.
func CreateNewCharacter(requestData CreateRequest) (models.Character, error) {
	num := GetNextCharacterNum()
	num_string := strconv.Itoa(num)
	char := models.Character{
		ID:        requestData.ID,
		Name:	   "Player " + num_string,	
		PlayerNum: num,
		Race:      requestData.Race,
		Strength:  requestData.Strength,
		Dexterity: requestData.Dexterity,
		Wisdom:    requestData.Wisdom,
		CreatedAt: time.Now(),
	}
	err := DB.C("characters").Insert(&char)

	if err != nil {
		return char, err
	}

	return char, nil
}

// Finds a character given some identifier string. This string can either
// be the _id, pro_id, or num of the character. If none exists, an error
// is returned.
func FindCharacter(identifier string) (models.Character, error) {
	var result models.Character


	// If identifier is an integer, find character with that player number
	err := DB.C("characters").Find(bson.M{"_id": identifier}).One(&result)
	if err != nil {
		err = DB.C("characters").Find(bson.M{"pro_id": identifier}).One(&result)
		if err != nil {
			i, _ := strconv.Atoi(identifier)
			err = DB.C("characters").Find(bson.M{"number": i}).One(&result)
			if err != nil {
				return result, err
			}
		}
	}

	if err != nil {
		return result, err
	}

	return result, nil
}
