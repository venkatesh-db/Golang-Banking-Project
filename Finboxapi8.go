
package main

import (
	"log"
	"net/http"
)

func middlewareOne(next http.Handler) http.Handler {

                                                 // interface            // interface


	log.Print(" prabhas ")

	return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {

		log.Print("Executing middlewareOne")

		               next.ServeHTTP(w, r)

		  // interface object       // method     

		log.Print("Executing middlewareOne again")

	})
}

func middlewareTwo(next http.Handler) http.Handler {

         log.Print(" venkatesh villain for studnets ")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Print("Executing middlewareTwo")

		if r.URL.Path == "/foo" {
			return
		}

		next.ServeHTTP(w, r)
		log.Print("Executing middlewareTwo again")
	})
}

func final(w http.ResponseWriter, r *http.Request) {

      log.Print(" bhramnandam - balaram ")
	log.Print("Executing finalHandler")
	w.Write([]byte("OK"))
}

func main() {

	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final)

	mux.Handle("/", middlewareOne(middlewareTwo(finalHandler)))

	log.Print("Listening on :3000...")

	err := http.ListenAndServe(":3000", mux)

	log.Fatal(err)

}	