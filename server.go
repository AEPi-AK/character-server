package main

import (
	"log"

	"net/http"

	"gopkg.in/mgo.v2"
)

var (
	clean   = false
	DB      *mgo.Database
	counter Counter
)

func main() {

	session, err := mgo.Dial("mongodb://isaac:sucks@ds021650.mlab.com:21650/aepi-ak-booth-2016")

	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	DB = session.DB("aepi-ak-booth-2016")

	// If clean is true, drop the database and reset the auto-incrementing sequence
	if clean {
		err = DB.DropDatabase()
		if err != nil {
			panic(err)
		}

		err = DB.C("counter").Insert(&Counter{ID: "isaacsucks", Seq: 0})
		if err != nil {
			panic(err)
		}
	}

	if err != nil {
		panic(err)
	}

	log.Fatal(http.ListenAndServe(":8000", NewRouter()))
}
