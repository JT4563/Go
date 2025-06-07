package main

import (
  "encoding/json"
  "fmt"
  "net/http"
)

type User struct {
  Name  string `json:"name"`
  Email string `json:"email"`
}

func postHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST" {
    http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
    return
  }

  var user User
  err := json.NewDecoder(r.Body).Decode(&user)
  if err != nil {
    http.Error(w, "Invalid JSON", http.StatusBadRequest)
    return
  }

  fmt.Fprintf(w, "Received user: %s (%s)", user.Name, user.Email)
}

func main() {
  http.HandleFunc("/user", postHandler)
  fmt.Println("Server running on http://localhost:8080")
  http.ListenAndServe(":8080", nil)
}
