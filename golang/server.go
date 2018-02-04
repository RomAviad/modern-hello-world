package main

import (
	"io"
	"net/http"
	"strings"
	"encoding/json"
	"strconv"
)

type Result struct {
	Ngrams [][]string `json:"ngrams"`
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello World")

}

func nGrams(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	data := request.Form.Get("data")
	N, atoi_err := strconv.Atoi(request.Form.Get("n"))
	if N < 1 || atoi_err != nil {
		http.Error(writer, atoi_err.Error(), http.StatusInternalServerError)
	}
	splitted := strings.Split(data, " ")
	result := [][]string{}
	for i := 0; i < len(splitted) - N + 1; i++ {
		ngram := splitted[i:i+N]
		result = append(result, ngram)
	}
	res := Result{result}
	jsonified, err := json.Marshal(res)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(jsonified)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/ngrams", nGrams)
	http.ListenAndServe(":8080", nil)
}
