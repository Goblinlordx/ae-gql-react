package main

import (
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
)

func init() {
	http.HandleFunc("/", root)
}

type QueryRequest struct {
	Query string `json:"query"`
}

func root(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Not Found", 404)
		return
	}
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	t := QueryRequest{}
	err := decoder.Decode(&t)
	if err != nil {
		http.Error(w, "Invalid Request", 400)
		return
	}

	s, err := schema()
	if err != nil {
		panic(err)
	}
	result := graphql.Do(graphql.Params{
		Schema:        *s,
		RequestString: t.Query,
	})
	resj, _ := json.Marshal(result)
	w.Write(resj)
}
