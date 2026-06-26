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

func quadratic(a, b, c float64) (float64, float64, string) {
	// Not a quadratic equation
	if a == 0 {
		return 0, 0, "a cannot be zero"
	}

	discriminant := b*b - 4*a*c

	// No real roots
	if discriminant < 0 {
		return 0, 0, "equation has no real roots"
	}

	sqrtD := math.Sqrt(discriminant)

	x1 := (-b + sqrtD) / (2 * a)
	x2 := (-b - sqrtD) / (2 * a)

	return x1, x2, ""
}

func homehandler(w http.ResponseWriter, r *http.Request) {

	tpl, _ := template.ParseFiles("index.html")

	tpl.Execute(w, "index.html")
}

type comp struct {
	Result1 float64
	Result2 float64
	Error   string
}

func quadHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "home")
	// w.Write([]byte("hello"))
	// body := r.ParseForm()
	a, _ := strconv.Atoi(r.FormValue("a"))
	b, _ := strconv.Atoi(r.FormValue("b"))
	c, _ := strconv.Atoi(r.FormValue("c"))

	x1, x2, err := quadratic(float64(a), float64(b), float64(c))

	outPut := comp{x1, x2, err}

	json.NewEncoder(w).Encode(outPut)
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
