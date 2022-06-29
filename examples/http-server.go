package examples

import (
	"fmt"
	"log"
	"net/http"
)

func HttpServer(port string) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Content-Type", "text/plain")
		fmt.Fprintf(w, "Hello")

	})

	fmt.Printf("Starting web server on port %v\n", port)

	error := http.ListenAndServe(":"+port, nil)

	if error != nil {
		log.Fatal("Cant start http server ", error)
	}

}
