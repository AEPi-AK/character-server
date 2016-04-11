package main

import (
	"net/http"
	"os"

	"github.com/AEPi-AK/character-server/models"
	log "github.com/Sirupsen/logrus"
	"github.com/rs/cors"
	"gopkg.in/mgo.v2"
)

var (
	DB *mgo.Database
)

func main() {

	session, err := mgo.Dial("mongodb://isaac:sucks@ds021650.mlab.com:21650/aepi-ak-booth-2016")

	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	DB = session.DB("aepi-ak-booth-2016")

	// If characters collection empty or non-existent, drop the database
	// and reset the auto-incrementing sequence collection.
	count, err := DB.C("characters").Count()
	if err != nil || count == 0 {
		log.Warn("character collection empty or nonexistent, resetting database")
		err = DB.DropDatabase()
		if err != nil {
			panic(err)
		}

		err = DB.C("counter").Insert(&models.Counter{ID: "isaacsucks", Seq: 0})
		if err != nil {
			panic(err)
		}
	}

	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	handler := cors.Default().Handler(NewRouter())
	log.Info("Listening on port ", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
