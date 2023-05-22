package main 

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"github.com/gorilla/mux"
)

type  KycInput struct {
	 Redirecturl     string       
}

type KycResponse struct{
	Url string   `json:"url"`
}


func API(   res   http.ResponseWriter , req   *http.Request  ){

	var inputs KycInput

	id := req.URL.Query().Get("redirecturl")

        fmt.Println( id )

          u,err := url.Parse(id)   // http://localhost:8080/session/callback

	  if err!=nil{
                  json.NewEncoder(res).Encode("abhiram jamon")
		  return
	  }

	  if u.Scheme != "http" && u.Host  != "localhost"{
                      json.NewEncoder(res).Encode("invalid jamon")
		      return
	  }

         inputs.Redirecturl = id

	var resposne KycResponse = KycResponse{ Url:"https://coderrange.com" }

        fmt.Println( resposne )

  	json.NewEncoder(res).Encode(resposne)

}

 

func main(){

	//  1. server 
	//  2. register api to server 
	 
 	 
	 ret := mux.NewRouter() // Register api to server

	ret.HandleFunc("/kycs",API).Methods("POST")
	
	 http.ListenAndServe(":8080", ret) // server


}