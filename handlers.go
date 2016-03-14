package main

import (
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    // "log"
    "net/http"
    "strconv"
    
    "github.com/gorilla/mux"
    "github.com/davecgh/go-spew/spew"
    "gopkg.in/yaml.v2"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to your Todo Service!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request){    
    w.Header().Set("Content-type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    
    if len(todos) != 0 {
        if err := json.NewEncoder(w).Encode(todos); err != nil {
            panic(err)
        }
        return
    }
    
    if err := json.NewEncoder(w).Encode("No ToDos!"); err != nil {
        panic(err)
    }
}

func AddFile(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    
    if err != nil {
        panic(err)
    }
    
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    // source, err := ioutil.ReadFile("../file4.yml")
    spew.Dump(body)
    var compose Compose
    error := yaml.Unmarshal(body, &compose)
    
    spew.Dump(error)
    
    spew.Dump(compose)
    
}

func ExtractTodoId(vars map[string]string) int {
    var todoId int
    var err error
    
    if todoId, err = strconv.Atoi(vars["todoId"]); err != nil {
        panic(err)
    }
    return todoId
}

func TodoRemove(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := ExtractTodoId(vars)
    todoRemove := RepoFindTodo(todoId)
    todo := RepoRemoveTodo(todoId)
    if todo != nil {
        w.Header().Set("Content-type", "application/json;charset=UTF-8")
        w.WriteHeader(http.StatusNotFound)
        
        if err := json.NewEncoder(w).Encode(todo); err != nil {
            panic(err)
        }
        return
    }
    
    w.Header().Set("Content-type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(todoRemove); err != nil {
        panic(err)
    }
    
}

func TodoComplete(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := ExtractTodoId(vars)
    todo := RepoCompleteTodo(todoId)
    
    if todo.Id > 0 {
        w.Header().Set("Content-type", "application/json;charset=UTF-8")
        w.WriteHeader(http.StatusOK)
        if err := json.NewEncoder(w).Encode(todo); err != nil {
            panic(err)
        }
        return
    }
    
    w.Header().Set("Content-type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusNotFound)

    if err := json.NewEncoder(w).Encode("Not Found"); err != nil {
        panic(err)
    }
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := ExtractTodoId(vars)
    todo := RepoFindTodo(todoId)
    if todo.Id > 0 {
        w.Header().Set("Content-type", "application/json;charset=UTF-8")
        w.WriteHeader(http.StatusOK)
        if err := json.NewEncoder(w).Encode(todos); err != nil {
            panic(err)
        }
        return
    }
    w.Header().Set("Content-type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusNotFound)

    if err := json.NewEncoder(w).Encode("Not Found"); err != nil {
        panic(err)
    }
}

func TodoCreate(w http.ResponseWriter, r * http.Request){
    var todo Todo
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    
    if err != nil {
        panic(err)
    }
    
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    
    if err := json.Unmarshal(body, &todo); err != nil {
        w.Header().Set("Content-type", "application/json;charset=UTF-8")
        w.WriteHeader(422)
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }
    
    t := RepoCreateTodo(todo)
    w.Header().Set("Content-type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(t); err != nil {
        panic(err)
    }
}
