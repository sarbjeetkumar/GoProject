
# GoDropBox
4th Year Go Project

GoDropbox is a Web application used for the storing of files. The application can be accessed by downloading the project from GitHub. GoDropbox can be used to store and retrieve files, and allows you to add and delete files of any type. The application uses a MongoDb document database. As with all Nosql databases, Mongo puts no constraints on your data, which allows for excellent retrieval times. The front end of the application was developed using a combination of AngualrJs, Bootstrap, CSS and Jquery. Golang is the language of choice for middle ware, with Martini being the framework which was used in development.

##User Guide. 

####Needed Software
In order to use the application you must first have Go Lang (<https://golang.org/doc/install>) enviroment installed and set up on your pc. Follow the link provided for instructions.  
Next you will need Mongodb installed and working, you can find instalation instructions at [the Mongodb website]( https://www.mongodb.com/download-center?jmp=nav#community). If you perfer to follow a video [Derek Banas on Youtube has an exellent installation guide](https://www.youtube.com/watch?v=-0X8mr6Q8Ew&list=PLGLfVvz_LVvRfdt8_W0dV311Xa8SayfCY&index=1&t=172s)  

####Git Clone
Once the rquired software is installed, you then need to git clone the master branch of the GoProject repository or download a zip file copy.
```
 git clone https://github.com/sarbjeetkumar/GoProject.git
```
####Unzip Package
When the application is on your machine, navigate to the src directory. 
```
In the src directory navigate to gopkg.in directory and unzip the mgo.v2.zip
replace the existing mgo.v2 with the recently unziped.
```
####Running Mongo
To start the mongo database. Open up 2 seperate windows in your terminal comand prompt.  
In terminal 1 type
```
mongod
```
In terminal 2 type
```
mongo
```

####Set your go path:

For Linux Run:

```
export GOPATH=$HOME/path/to/project

```
For Windows Run:
```
set GOPATH=c:\path\to\project
```

####Run Go 
Build the code 
```
go build
```
Run the application 
```
go buid main.go
```
