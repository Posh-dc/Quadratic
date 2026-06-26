package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
	// "os"
)

var tpls *template.Template

func quadratic(a, b, c float64) (float64, float64) {

	det := math.Sqrt(((b) * (b)) - (4 * a * c))

	result1 := (-1*b + det) / (2 * a)
	result2 := (-1*b - det) / (2 * a)

	return result1, result2
}

func homehandler(w http.ResponseWriter, r *http.Request) {

	tpl, _ := template.ParseFiles("index.html")

	tpl.Execute(w, "index.html")
}

func quadHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "home")
	// w.Write([]byte("hello"))
	// body := r.ParseForm()
	a, _ := strconv.Atoi(r.FormValue("a"))
	b, _ := strconv.Atoi(r.FormValue("b"))
	c, _ := strconv.Atoi(r.FormValue("c"))

	outPUT := make(map[string]float64)

	result1, result2 := quadratic(float64(a), float64(b), float64(c))

	outPUT["result1"] = result1
	outPUT["result2"] = result2

	json.NewEncoder(w).Encode(outPUT)
	fmt.Println("i got clicked")
}

func main() {

	fs := http.FileServer(http.Dir("style"))
	http.Handle("/style/", http.StripPrefix("/style/", fs))

	http.HandleFunc("/", homehandler)
	port := ":9090"
	fmt.Println("server listening on port", port)

	http.HandleFunc("/Result", quadHandler)

	http.ListenAndServe(port, nil)
}
