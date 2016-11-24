package main

import (
	"../github.com/go-martini/martini"
	"../github.com/martini-contrib/render"
	"../github.com/martini-contrib/binding"
	"../gopkg.in/mgo.v2"
	//"../gopkg.in/mgo.v2/bson"
	"fmt"
	"log"
)

type Register struct {
	Name  		string `form:"name"`
	Email 		string `form:"email"`
	Username 	string `form:"username"`
	Password 	string `form:"password"`
	Age 		string `form:"age"`
	Sex 		string `form:"sex"`
}

func DB() martini.Handler {
	session, err := mgo.Dial("localhost:27017") // mongodb://localhost
	if err != nil {
		panic(err)
	}

	emailIndex(session)

	return func(c martini.Context) {
		s := session.Clone()
		//Database GoDropBox
		c.Map(s.DB("GoDropBox")) // local
		defer s.Close()
		c.Next()
	}
}


func emailIndex(s *mgo.Session) {
	//Modified from - http://goinbigdata.com/how-to-build-microservice-with-mongodb-in-golang/
	session := s.Copy()
	defer session.Close()

	c := session.DB("GoDropBox").C("users")

	index := mgo.Index{
		Key: []string{"email"},
		Unique: true,
		DropDups: true,
		Background: true,
		Sparse: true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	m := martini.Classic()

	m.Use(render.Renderer(render.Options {
		Directory: "public/",
	}))

	m.Use(DB())

	m.Post("/", binding.Form(Register{}), func(register Register, r render.Render, db *mgo.Database) {
		if err := db.C("users").Insert(register); err != nil {
			if mgo.IsDup(err) {
				fmt.Printf("%s Already exsists ", register.Email)
				// Is a duplicate key, but we don't know which one
			}
		}

	})

	//m.Run() uses the default port 3000
	m.RunOnAddr(":8080")//specifies port 8080
}