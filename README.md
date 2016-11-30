
# GoProject
4th Year Go Project



Technologies 


Technologies For front end , Middle ware and Backend.
Front end

Bootstrap
http://getbootstrap.com/

We used bootstrap because: 

- Speed development 
- Responsiveness

AngularJS
https://angularjs.org/

- REST Easy.
- MVVM to the Rescue
- Data Binding and Dependency Injection
- Makes HTML your Template

JQuery
https://jquery.com/

- we use Jquery because we wanted to make a single page application and we did not want page to refresh but we want jQuery to work in the background . 

##Setting up the front-end 
Created a public folder, which holds all the HTML, Javascript and CSS.  
Website template was taken from http://getbootstrap.com/, added to the project and adobted.  
The package manager Bower was used to install any dependancies needed in the web development.   
By using the following commaand    
```linux
$bower install angular/bootstrap/jquery  
```
The bower-componts folder in the public directory was created and the dependancies added.  
The rquired links were then added to the html to include Angularjs, Bootstrap and Jquery.


MiddleWare (API)

Must be written in Go

##Setting up Go Environment Windows command line
```
**If you want to use unix commands in windows command-line**
http://lifehacker.com/362316/use-unix-commands-in-windows-built-in-command-prompt

cd desktop - Change Directory
mkdir src, mkdir pkg, mkdir bin, mkdir public - mkdir: will create a new directory. http://www.slackbook.org/html/file-commands-creation.html

set GOPATH=location of project. Example set GOPATH=C:\Users\John Doe\Desktop\Project
set GOBIN=location of project. Example set GOBIN=C:\Users\John Doe\Desktop\Project\bin
**Notice windows uses backslashes**

cd src - Change Directory
touch main.go - touch: is used to change the timestamp on a file. If the file specified does not exist, touch will create a zero length file with the name specified. http://www.slackbook.org/html/file-commands-creation.html

Install IntelliJ IDEA. https://www.jetbrains.com/idea/download/
*Set up Intellij for Go* - https://rootpd.com/2016/02/04/setting-up-intellij-idea-for-your-first-golang-project/
```
##Setting up Go Environment Git Bash
```
https://git-scm.com/downloads
*Difference between Git Bash and Windows command-line*
https://www.quora.com/What-is-difference-between-GIT-GUI-GIT-BASH-and-GIT-CMD

cd desktop - Change Directory
mkdir src, mkdir pkg, mkdir bin, mkdir public - mkdir: will create a new directory. http://www.slackbook.org/html/file-commands-creation.html

export GOPATH=location of project. Example set GOPATH=C:/Users/John\ Doe/Desktop/Project

export GOBIN=location of bin. Example set GOPATH=C:/Users/John\ Doe/Desktop/Project/bin

**Notice Bash uses forward slashes and backslash to ignore the space **

cd src - Change Directory
touch main.go - touch: is used to change the timestamp on a file. If the file specified does not exist, touch will create a zero length file with the name specified. http://www.slackbook.org/html/file-commands-creation.html

Install IntelliJ IDEA. https://www.jetbrains.com/idea/download/
*Set up Intellij for Go* - https://rootpd.com/2016/02/04/setting-up-intellij-idea-for-your-first-golang-project/
```


https://golang.org/

Backend(Database)

https://www.mongodb.com/
```
**We decide to use MongoDB because its a doucument based data base and save data in Json structure . Can query database through MongoDB 
query language , MongoDB uses dynamic schemas, meaning that you can create records without first defining the structure, such as the fields or the types of their values.

quick guide how to install MongoDB on Windows :-

*https://docs.mongodb.com/manual/tutorial/install-mongodb-on-windows/



```



Frameworks

Martini 
------

```
**Martini is framwork which is was created for go Language .

How to install Martini through command promt ....#

First install the Martini folder in project folder .

And type this command in your command promt and Martini will install itself.

*

```


