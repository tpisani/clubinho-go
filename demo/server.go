package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello, %s! It is %s.", ps.ByName("username"), time.Now().Format(time.Kitchen))
}

func helloWorldJSONHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"username": ps.ByName("username"),
		"time":     time.Now(),
	})
}

func main() {
	addr := ":8000"

	r := httprouter.New()

	r.GET("/:username/", helloWorldHandler)
	r.GET("/:username/json/", helloWorldJSONHandler)

	log.Printf("Running at %s...", addr)
	fmt.Println("Try opening the following URLs on your browser:")
	fmt.Println("http://localhost:8000/<username>/")
	fmt.Println("http://localhost:8000/<username>/json/")
	http.ListenAndServe(":8000", r)
}
