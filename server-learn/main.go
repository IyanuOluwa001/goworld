package main

import (
	// "encoding/json"
	// "fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	// "strconv"
)

// func main2(){
// 	//http://localhost:8080/math-form/
// fs := http.FileServer(http.Dir("./static"))
// http.Handle("/math-form/", http.StripPrefix("/math-form", fs))

// 	http.HandleFunc("/add-form", func(w http.ResponseWriter, r *http.Request){
// 		if r.Method != http.MethodPost{
// 			http.Error(w, "Invalid Request", http.StatusMethodNotAllowed)
// 			return
// 		}
// 		if err := r.ParseForm(); err!= nil{
// 			http.Error(w, "Can't parse body", http.StatusBadRequest)
// 			return
// 		}

// 		num1Str := r.FormValue("num1")
// 		num2Str := r.FormValue("num2")

// 		num1Int, err1 := strconv.Atoi(num1Str)
// 		num2Int, err2 := strconv.Atoi(num2Str)

// 		if err1 != nil || err2 != nil{
// 			http.Error(w, "Invalid Input", http.StatusBadRequest)
// 		}

// 		result := num1Int + num2Int

// 		response := map[string]interface{}{
// 			"result": result,
// 		}

// 		w.Header().Set("Cntent-Type", "application/json")
// 		json.NewEncoder(w).Encode(response)
// 	})

// 	//http://localhost:8080/add?a=6&b=8
// 	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request){
// 		r.ParseForm();
// 		a := r.FormValue("a")
// 		b := r.FormValue("b")

// 		aInt, err1 := strconv.Atoi(a);
// 		bInt, err2 := strconv.Atoi(b);

// 		if err1 != nil || err2!=nil{
// 			http.Error(w, "Invalid Input", http.StatusBadRequest)
// 			return

// 		}
// 		fmt.Fprintf(w, "Addition of %d and %d is %d", aInt, bInt, aInt+bInt)

// 	})

// 	//Greet
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
// 		fmt.Fprint(w, "Greetings")
// 	})

// 	//Start the Server
// 	fmt.Println("server is running on port :8080")
// 	if err := http.ListenAndServe(":8080", nil); err != nil{
// 		fmt.Println("Server is not running")
// 	}
// }

// Handler, HandleFunc, Handle

// var templ = template.Must(template.ParseFiles("templates/index.html"))



type NumResult struct{
	Result int
}


var templ = template.Must(template.ParseGlob("templates/*.html"))


func main(){
	mux := http.NewServeMux()

    mux.Handle("/static/",http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	// mux.Handle("/static",http.FileServer(http.Dir("./style.css")))
	// mux.Handle("/static",http.FileServer(http.Dir("./app.js")))

   mux.HandleFunc("/", HomeHandler)
   mux.HandleFunc("/add-num", AddHandler)
//    mux.HandleFunc("ascii", Ascii)
  log.Println("Server listening on port 8080")
  http.ListenAndServe(":8080", mux)
 }


 func HomeHandler( w http.ResponseWriter, r *http.Request){
	
	// err := templ.Execute(w, nil)

	err := templ.ExecuteTemplate(w,"index.html",nil)
	if err != nil {
	   log.Println("error opening index.html")
	}

 }


 func AddHandler( w http.ResponseWriter, r *http.Request){
    r.ParseForm()

	num1, _ := strconv.Atoi(r.FormValue("num1")) 
	num2, _ := strconv.Atoi(r.FormValue("num2"))

	result := num1 + num2

	data := NumResult{
		Result: result,
	}

	err := templ.ExecuteTemplate(w,"index.html",data)
	if err != nil {
	   log.Println("error opening index.html")
	}
 }
