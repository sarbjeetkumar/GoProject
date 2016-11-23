package main

import (
	"../github.com/go-martini/martini"
	"../github.com/martini-contrib/render"
	"../github.com/martini-contrib/binding"
	"../gopkg.in/mgo.v2"
	"../gopkg.in/mgo.v2/bson"
	"log"
	"fmt"
	"net/http"
	"io/ioutil"
)

type Register struct {
	Name  		string `form:"name"`
	Email 		string `form:"email"`
	Username 	string `form:"username"`
	Password 	string `form:"password"`
	Age 		string `form:"age"`
	Sex 		string `form:"sex"`
}

func find() martini.Handler{
	session, err := mgo.Dial("localhost:27017") // connect to mongo localhost
	//session, err := mgo.Dial("mongodb://51.141.9.220:27017") // connect to mongo on azure

	c := session.DB("GoDropBox").C("users")

	var results []Register
	err = c.Find(bson.M{"name": "sean"}).All(&results)

	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)

	return func(c martini.Context) {
		s := session.Clone()
		//Database GoDropBox
		c.Map(s.DB("GoDropBox")) // local
		defer s.Close()
		c.Next()
	}

}

func DB() martini.Handler {
	session, err := mgo.Dial("localhost:27017") // connect to mongo localhost
	//session, err := mgo.Dial("mongodb://51.141.9.220:27017") // connect to mongo on azure

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

func All(db *mgo.Database) []Register {


	var register []Register
	db.C("users").Find(nil).All(&register)
	return register
}

func upload(w http.ResponseWriter, req *http.Request) {
	// Adapted from: http://stackoverflow.com/questions/22159665/store-uploaded-file-in-mongodb-gridfs-using-mgo-without-saving-to-memory
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
	panic(err)
	}

	defer session.Close()

	// Retrieve the form file
	file, handler, err := req.FormFile("uploadfile")
	//Check if there is an error
	if err != nil {
	fmt.Println(err)
	}
	// Read the file into memory
	data, err := ioutil.ReadAll(file)
	if err != nil {
	fmt.Println(err)
	}

	// Specify the Mongodb database
	my_db := session.DB("Files")

	// Set the filename as the uploadfile name
	filename := handler.Filename

	// Create the file in the Mongodb Gridfs instance
	my_file, err := my_db.GridFS("fs").Create(filename)
	if err != nil {
	fmt.Println(err)
	}
	// Write the file to the Mongodb Gridfs instance
	n, err := my_file.Write(data)
	if err != nil {
	fmt.Println(err)
	}

	file_id := my_file.Id().(bson.ObjectId).Hex()
	response := file_id

	// Close the file
	err = my_file.Close()
	if err != nil {
	fmt.Println(err)
	}

	// Write a log type message
	fmt.Printf("%d bytes written to the Mongodb instance\n", n)
	fmt.Println(response)

}

func main() {

	m := martini.Classic()

	m.Use(render.Renderer(render.Options {
		Directory: "public/pages",
	}))

	m.Use(DB())

	m.Post("/upload", upload)

	m.Get("user/:userID", func(r render.Render , p martini.Params) {

		/*
		t := template.New("fieldname example")
		t, _ = t.Parse("hello {{.UserName}}!")
		p := Register{Username: "Astaxie"}
		t.Execute(os.Stdout, p)
		*/


		var retData struct{
			ID string
		}
		retData.ID = p["userID"]
		r.HTML(http.StatusOK, "user", retData)

	})

	m.Post("/", binding.Form(Register{}), func(register Register, r render.Render, db *mgo.Database) {
		if err := db.C("users").Insert(register); err != nil {
			if mgo.IsDup(err) {
			}
		}
	})

	//m.Run() uses the default port 3000
	m.RunOnAddr(":8080")//specifies port 8080
}