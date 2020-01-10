package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Firefox struct {
	FF  string
	NO1 string
}

func main() {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("testhl").C("Firefox")
	err = c.Insert(&Firefox{"FF_v", "NO.1"})

	if err != nil {
		log.Fatal(err)
	}

	result := Firefox{}
	err = c.Find(bson.M{"ff": "FF_v"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Top: ", result.NO1, "\t", result.FF)
}
