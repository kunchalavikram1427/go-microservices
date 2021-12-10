package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	details "github.com/kunchalavikram1427/go-microservices/details"

	"github.com/gorilla/mux"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking application health")
	response := map[string]string{
		"status":    "UP",
		"timestamp": time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving the homepage")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application is up and running")
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching the details")
	hostname, err := details.GetHostname()
	if err != nil {
		panic(err)
	}
	IP, _ := details.GetIP()
	fmt.Println(hostname, IP)
	response := map[string]string{
		"hostname": hostname,
		"ip":       IP.String(),
	}
	json.NewEncoder(w).Encode(response)

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/details", detailsHandler)
	log.Println("Server has started!!!")
	log.Fatal(http.ListenAndServe(":80", r))
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func rootHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, you've requested: %s with token: %s\n", r.URL.Path, r.URL.Query().Get("token"))
// }

// func main() {
// 	http.HandleFunc("/", rootHandler)

// 	fs := http.FileServer(http.Dir("static/"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	log.Println("Web server has started")
// 	http.ListenAndServe(":80", nil)
// }

// package main

// import (
// 	"fmt"
// 	"unsafe"

// 	geo "github.com/kunchalavikram1427/go-microservices/geometry"

// 	"rsc.io/quote"
// )

// func rectProps(length, width float64) (area, perimeter float64) {
// 	area = length * width
// 	perimeter = 2 * (length + width)
// 	return
// }

// func main() {
// 	x := 10
// 	name := "DevOps"
// 	isWorking := false

// 	fmt.Println("Hello, World!")
// 	fmt.Println(quote.Go())
// 	fmt.Println(x, name, isWorking)
// 	fmt.Printf("Type of name %T and size is %d\n", name, unsafe.Sizeof(name))

// 	a, p := rectProps(1, 2)
// 	fmt.Printf("Area is %f and perimeter is %f", a, p)

// 	// var daysOfTheMonth map[string]int
// 	// daysOfTheMonth["Jan"] = 31

// 	var daysOfTheMonth = map[string]int{"Jan": 31, "Feb": 28}
// 	fmt.Println(daysOfTheMonth)

// 	area := geo.Area(1, 2)
// 	diag := geo.Diagonal(1, 2)
// 	fmt.Println(area, diag)

// }
