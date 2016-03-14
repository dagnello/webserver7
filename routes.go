package main

import "net/http"

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        Name:        "Index",
        Method:      "GET",
        Pattern:     "/",
        HandlerFunc: Index,
    },
    Route{
        Name:        "TodoIndex",
        Method:      "GET",
        Pattern:     "/todos",
        HandlerFunc: TodoIndex,
    },
    Route{
        Name:        "TodoCreate",
        Method:      "POST",
        Pattern:     "/todos",
        HandlerFunc: TodoCreate,
    },
    Route{
        Name:        "AddFile",
        Method:      "POST",
        Pattern:     "/files",
        HandlerFunc: AddFile,
    },
    Route{
        Name:        "TodoRemove",
        Method:      "DELETE",
        Pattern:     "/todos/{todoId}",
        HandlerFunc: TodoRemove,
    },
    Route{
        Name:        "TodoComplete",
        Method:      "PUT",
        Pattern:     "/todos/{todoId}",
        HandlerFunc: TodoComplete,
    },
    Route{
        Name:        "Index",
        Method:      "GET",
        Pattern:     "/todos/{todoId}",
        HandlerFunc: TodoShow,
    },
}