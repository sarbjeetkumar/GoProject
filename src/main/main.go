package main

import (
	"../github.com/go-martini/martini"
	"../github.com/martini-contrib/render"
	"../github.com/martini-contrib/binding"
	"gopkg.in/mgo.v2"
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
	session, err := mgo.Dial("localhost:27017/users") // mongodb://localhost
	if err != nil {
		panic(err)
	}

	return func(c martini.Context) {
		s := session.Clone()
		c.Map(s.DB("users")) // local
		defer s.Close()
		c.Next()
	}
}

func main() {

	m := martini.Classic()

	m.Use(render.Renderer(render.Options {
		Directory: "public/pages", // Specify what path to load the templates from.
		//Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.


	}))
	m.Use(DB())

	m.Post("/", binding.Form(Register{}), func(register Register, r render.Render, db *mgo.Database) {
		db.C("users").Insert(register)
	})

	//m.Run() uses the default port 3000
	m.RunOnAddr(":8080")//specifies port 8080
}