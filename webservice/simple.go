package webservice

import (
	"fmt"
	"log"
	"net/http"
)

func RunSimpleWebservice() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Follow The Pattern!")
	})
	fmt.Println("Server running!")

	log.Fatal(http.ListenAndServe(":8081", nil))
}
