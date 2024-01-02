package main

import (
	//"fmt"
	"net/http"
	"strconv"
	"strings"
)
	
var counter = make(map[string]int64)

func gaugeHandler(w http.ResponseWriter, r *http.Request) {
	gauge := make(map[string]float64)
	textUrl := r.URL.Path
	var sliceText []string = strings.Split(textUrl, "/")
	if len(sliceText) < 5 {
		http.Error(w, "Запрос без метрики", http.StatusNotFound)
		return
	}
	value, _ := strconv.ParseFloat(sliceText[4], 64)

	if r.Method == http.MethodPost {
		gauge[sliceText[3]] = value
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		// for key, val := range gauge {
		// 	fmt.Println(key, val)
		// }
		// fmt.Println()
	} 
}

func counterHandler(w http.ResponseWriter, r *http.Request) {
	textUrl := r.URL.Path
	var sliceText []string = strings.Split(textUrl, "/")
	if len(sliceText) < 5 {
		http.Error(w, "Запрос без метрики", http.StatusNotFound)
		return
	}
	value, _ := strconv.Atoi(sliceText[4])

	if r.Method == http.MethodPost {
		counter[sliceText[3]] = int64(value)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		// for key, val := range counter {
		// 	fmt.Println(key, val)
		// }
		// fmt.Println()
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/update/gauge/", gaugeHandler)
	mux.HandleFunc("/update/counter/", counterHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}