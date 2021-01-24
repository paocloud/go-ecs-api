package main

import (
  "fmt"
  "net/http"
  "encoding/json"
)

type healthz struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type addressBook struct {
    Firstname string  `json:"firstname"`
    Lastname  string  `json:"lastname"`
    Email     string  `json:"email"`
    Company   string  `json:"company"`
    Phone     string  `json:"phone"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "go-ecs-api-example")
}

func getContact(w http.ResponseWriter, r *http.Request) {
   addBook := addressBook{
                Firstname: "Payungsak",
                Lastname:  "Klinchampa",
                Email:     "payungsak.k@paocloud.co.th",
                Company:   "PAOCLOUD COMPANY LIMITED",
                Phone:     "093-7738265",
              }
   json.NewEncoder(w).Encode(addBook)
}

func getHealthz(w http.ResponseWriter, r *http.Request) {
   healthzMsg := healthz{
                  Message: "Healthz !!!",
                  Code: 200,
              }
   json.NewEncoder(w).Encode(healthzMsg)
}

func handleRequest() {
  http.HandleFunc("/", homePage)
  http.HandleFunc("/contact", getContact)
  http.HandleFunc("/healthz", getHealthz)
  http.ListenAndServe(":8080", nil)
}

func main() {
  handleRequest()
}

