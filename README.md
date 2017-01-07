# javaVSgo

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

TODO

## Go code for it

install go and go-swagger:

    install go from https://golang.org/doc/install
    $ brew tap go-swagger/go-swagger
    $ brew install go-swagger

preparation:

* download swagger UI (https://github.com/swagger-api/swagger-ui)
* create a folder named swagger in folder go
* copy content of dist folder of swagger-ui into the folder folder named
swagger

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


## comparison

TODO

## bugs

swagger generate spec -o ./swagger/swagger.json -m does not work


