package main

import (
	"io"
	"net/http"
	"strings"
	//"fmt"
)

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello World")

}

func nGrams(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	data := request.Form.Get("data")
	splitted := strings.Split(data, " ")
	ngrams_str := "[["
	for i := 0; i < len(splitted); i++ {
		if i != len(splitted)-1 {
			var tmp_arr = []string{ngrams_str, strings.Join([]string{"'", "'"}, splitted[i])}
			ngrams_str = strings.Join(tmp_arr, "")
			tmp_arr[0] = ngrams_str
			tmp_arr[1] = strings.Join([]string{"'", "'"}, splitted[i+1])
			ngrams_str = strings.Join(tmp_arr, ", ")
			tmp_arr[0] = ngrams_str
			tmp_arr[1] = "]"
			ngrams_str = strings.Join(tmp_arr, "")
			if i < len(splitted)-2 {
				tmp_arr[0] = ngrams_str
				tmp_arr[1] = ", ["
				ngrams_str = strings.Join(tmp_arr, "")
			}
		}
	}
	var tmp = []string{ngrams_str, "]"}
	ngrams_str = strings.Join(tmp, "")
	io.WriteString(writer, ngrams_str)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/ngrams", nGrams)
	http.ListenAndServe(":8080", nil)
}
