package main

import (
    "encoding/json"
    "net/http"
    "strconv"
    "sync"
)

type Person struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}

var (
    people = make(map[int]Person)
    nextID = 1
    mu     sync.Mutex
)

func main() {
    http.HandleFunc("/people", peopleHandler)
    http.HandleFunc("/people/", personHandler)
    http.ListenAndServe(":8080", nil)
}

func peopleHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        getAllPeople(w, r)
    case http.MethodPost:
        createPerson(w, r)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func personHandler(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Path[len("/people/"):])
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    switch r.Method {
    case http.MethodGet:
        getPerson(w, r, id)
    case http.MethodPut:
        updatePerson(w, r, id)
    case http.MethodDelete:
        deletePerson(w, r, id)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func getAllPeople(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()

    var result []Person
    for _, person := range people {
        result = append(result, person)
    }
    json.NewEncoder(w).Encode(result)
}

func createPerson(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()

    var person Person
    if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    person.ID = nextID
    nextID++
    people[person.ID] = person
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(person)
}

func getPerson(w http.ResponseWriter, r *http.Request, id int) {
    mu.Lock()
    defer mu.Unlock()

    person, exists := people[id]
    if !exists {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(person)
}

func updatePerson(w http.ResponseWriter, r *http.Request, id int) {
    mu.Lock()
    defer mu.Unlock()

    var updatedPerson Person
    if err := json.NewDecoder(r.Body).Decode(&updatedPerson); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    if _, exists := people[id]; !exists {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    updatedPerson.ID = id
    people[id] = updatedPerson
    json.NewEncoder(w).Encode(updatedPerson)
}

func deletePerson(w http.ResponseWriter, r *http.Request, id int) {
    mu.Lock()
    defer mu.Unlock()

    if _, exists := people[id]; !exists {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    delete(people, id)
    w.WriteHeader(http.StatusNoContent)
}
