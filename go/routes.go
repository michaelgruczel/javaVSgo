package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

    // swagger:route GET /messages list messages
    //
    // Lists messages.
    //
    // This will list all messages
    //
    //     Consumes:
    //     - application/json
    //
    //     Produces:
    //     - application/json
    //
    //     Schemes: http
    //
    //     
    //     Responses:
    //       default: genericError
    //       200: someResponse

    // swagger:route POST /message add message
    //
    // Adds a message.
    //
    // This will add a message
    //
    //     Consumes:
    //     - application/json
    //
    //     Produces:
    //     - application/json
    //
    //     Schemes: http
    //
    //     swagger: Message
    //
    //     Responses:
    //       default: genericError
    //       200: someResponse    

    
var routes = Routes {
	Route {
		"MessageIndex",
		"GET",
		"/messages",
		MessageIndex,
	},
	Route {
		"CreateMessage",
		"POST",
		"/message",
		CreateMessage,
	},
	
}
