package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Employee struct {
	Emp_ID   int32  `json:"Emp_ID"`
	Emp_Name string `json:"Emp_Name"`
	Emp_Dept string `json:"Emp_Dept"`
}

func main() {
	http.HandleFunc("/", ReqeustHandler)
	http.ListenAndServe(":8080", nil)
}

func ReqeustHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	getEmp := Employee{
		Emp_ID:   11232,
		Emp_Name: "Vishal",
		Emp_Dept: "Technology",
	}

	switch req.Method {
	case "GET":
		err := json.NewEncoder(w).Encode(getEmp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal(err)
		}
	case "POST":
		postEmp := Employee{}
		err := json.NewDecoder(req.Body).Decode(&postEmp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal(err)
		}
		fmt.Fprintf(w, "The input recieved is=>%+v", postEmp)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
