package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	// "os"
)

var tpls *template.Template

func quadratic(a, b, c int) (int, int) {
return a+b , b+c
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

	outPUT := make(map[string]int)

	
	
	
result1 , result2 := quadratic(a,b,c)


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
