package main

import ( 
	"log"

	"net/http"
	"gopkg.in/mgo.v2"
)

var (
	clean = false
	characterCollection *mgo.Collection
	counterCollection *mgo.Collection
	counter Counter
)

func main() {

	session, err := mgo.Dial("mongodb://isaac:sucks@ds021650.mlab.com:21650/aepi-ak-booth-2016")

	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	// When clean is true, it resets the DB on run
	if clean {
		err = session.DB("aepi-ak-booth-2016").DropDatabase()
		if err != nil {
			panic(err)
		}
	}

	characterCollection = session.DB("aepi-ak-booth-2016").C("people")
	counterCollection = session.DB("aepi-ak-booth-2016").C("counter")

	// Initialize Counter if starting a clean DB
	if clean { 
		err = counterCollection.Insert(&Counter{ID: "isaacsucks", Seq: 0})
	}

	if err != nil {
		panic(err)
	}


	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}
