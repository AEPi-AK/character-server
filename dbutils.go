package main

import (
	"strconv"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Gets the next auto-incremented player num.
func GetNextCharacterNum() int {
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
func UpdateCharacter(request UpdateRequest) Character {

	char := bson.M{"_id": request.ID}
	if request.Gold != 0 {
		change := bson.M{"$set": bson.M{"gold": request.Gold}}
		err := DB.C("characters").Update(char, change)

		if err != nil {
			panic(err)

		}
	}

	if request.ProID != "" {
		change := bson.M{"$set": bson.M{"pro_id": HashString(request.ProID)}}
		err := DB.C("characters").Update(char, change)

		if err != nil {
			panic(err)

		}
	}

	if request.Experience != 0 {
		change := bson.M{"$set": bson.M{"experience": request.Experience}}
		err := DB.C("characters").Update(char, change)

		if err != nil {
			panic(err)

		}
	}

	character, _ := FindCharacter(request.ID)

	return character

}

// Returns the player number given some id (pro or regular). If the id is not
// in the DB, we return an error.
func PlayerNumForID(id string) (int, error) {

	result := Character{}
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
func CreateNewCharacter(data string) Character {
	char := Character{ID: HashString(data), PlayerNum: GetNextCharacterNum()}
	err := DB.C("characters").Insert(&char)

	if err != nil {
		panic(err)
	}

	return char
}

// Finds a character given some identifier string. This string can either
// be the _id, pro_id, or num of the character. If none exists, an error
// is returned.
func FindCharacter(identifier string) (Character, error) {

	result := Character{}
	err := DB.C("characters").Find(bson.M{"_id": identifier}).One(&result)
	if err != nil {
		i, err := strconv.Atoi(identifier)

		if err != nil {
			return result, err
		}

		err = DB.C("characters").Find(bson.M{"num": i}).One(&result)
		if err != nil {
			return result, err
		}
	}

	return result, nil
}
