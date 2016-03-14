package main
import "fmt"
import "log"

var currentId int
var todos Todos

func init() {
    
}

func RepoFindTodo(id int) Todo {
    for _, t := range todos {
        if t.Id == id {
            return t
        }
    }
    return Todo{}
}

func RepoCreateTodo(t Todo) Todo {
    currentId += 1
    t.Id = currentId
    todos = append(todos, t)
    return t
}

func RepoCompleteTodo(id int) Todo {
    for i, t := range todos {
        if t.Id == id {
            log.Printf("todoId: %d", id)
            todos[i].Completed = true
            return t
        }
    }
    return Todo{}
}

func RepoRemoveTodo(id int) error {
    for i, t := range todos {
        if t.Id == id {
            todos = append(todos[:i], todos[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}