# javaVSgo

![travis status](https://travis-ci.org/michaelgruczel/javaVSgo.svg?branch=master)

This repo is just a playground to compare go to java (spring boot) a little bit
by a very simple example. 

The mission is to create a rest-microservice as lean as possible,
and it should contains some things:

* add a message by post
* read all messages
* restful
* json
* embedded executable Swagger docu

So it's very simple but needs some embedding of libs.

## Java code for it

    install java 8
    $ cd java
    $ ./gradlew build
    $ ./gradlew bootRun
    
now you can get the list of existing messages by:

    $ curl -H "Content-Type: application/json" http://localhost:8080/messages

you can add new messages by

    $ curl -H "Content-Type: application/json" -d '{"author":"Bart", "content":"Eat my shorts"}' http://localhost:8080/message

you can open the executable swagger spec by open

* http://localhost:8080/swagger-ui.html

lets just run the binary:

    $ cp build/libs/java-example-service-0.1.0.jar ./../binaries/java-example-service-0.1.0.jar
    $ cd ./../binaries
    $ java -jar java-example-service-0.1.0.jar

## Go code for it

install go and go-swagger:

    install go from https://golang.org/doc/install
    $ brew tap go-swagger/go-swagger
    $ brew install go-swagger

lets start

    $ cd go
    $ swagger generate spec -o ./swagger/swagger.json
    $ go run *.go

now you can get the list of existing messages by:

    $ curl -H "Content-Type: application/json" http://localhost:8080/messages

you can add new messages by

    $ curl -H "Content-Type: application/json" -d '{"author":"Bart", "content":"Eat my shorts"}' http://localhost:8080/message

you can open the executable swagger spec by open

* http://localhost:8080/swagger/
* enter http://localhost:8080/swagger/swagger.json in swagger spec field

let's package it as application:

    $ go build -o ./../binaries/go-example-service .

## comparison

A comparison shows how lean and fast go is for micro services 
in comparison to microservices in Java with a framework like spring boot.

    $ cd binaries
    $ ls -lh 
   
show the different sizes of the apps   
   
    6.1M go-example-service
    408B swagger
    
    31M java-example-service-0.1.0.jar
    
the go-service-example with dependencies has a size of 6.1 MB and the java rest service with dependencies  has 31MB.
In this case the assets (swagger html) for go lying in the same folder like the go assembly.
There are included in the java assembly. 
There are some plugins to add the resources to the go assembly like for example https://github.com/jteeuwen/go-bindata
but I skipped it here. Booting the go service takes a second, the java service takes about 8 seconds on my mac.

The swagger spec lacks models, means the "add description" cannot be used by the spec.
There is a parameter (-models) which should generate the model description which only works if
the files are organized within a specific go structure. I skipped that here.
Building Frameworks are not so mature in go, the handling of assets or the swagger integration
are just examples for it. Go is great for simple Microservices, but with increasing complexity
the lack of tooling support makes other frameworks like java with spring boot more attractive. 

## Issues and TODOs

Issues:

swagger generate spec -o ./swagger/swagger.json -m does not work
https://github.com/go-swagger/go-swagger/issues/753
http://swagger.io/specification/#definitionsObject

TODOs:

* Integrate asset handling and swagger correctly
* build go example with service discovery, config management
* more detailed comparison, build tools, code, ...
