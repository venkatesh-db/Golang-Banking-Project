package main 

import (
	"encoding/json"
	  "github.com/go-sql-driver/mysql"
	"net/http"
	"fmt"
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"os"
)

type  KycInput struct {
	 Name     string        `form:"name"`
        Dob      string              `form:"dob"`
        Pan      string                `form:"pan"`
	Address  string       `form:"address"`
	Pincode   string   `form:"pincode"`
}

type KycResponse struct{
	Message string
}

var db *sql.DB

func API(   res   http.ResponseWriter , req   *http.Request  ){

	var inputs KycInput

	id := req.URL.Query().Get("name")

	if len(inputs.Name) <0 {
		        json.NewEncoder(res).Encode("invalid jamon")
	}
 
	if len(inputs.Pan) >10 {
		     	json.NewEncoder(res).Encode("invalid momos ")
	}

       db.Exec("INSERT INTO kyc_dbs (name) VALUES (?)",id)

	var resposne KycResponse = KycResponse{ Message:"partner push completed succeesfuly" }

  	json.NewEncoder(res).Encode(resposne)

}

     // 3. connect to db
	 // 4. create table 
func init() {

 // Capture connection properties.

    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "pgbus",
	AllowNativePasswords: true,
    }

    // Get a database handle.
    var err error

    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    
    fmt.Println("Connected happy brek year 2022 !")

}

func main(){

	//  1. server 
	//  2. register api to server 
	 
 	 
	 ret := mux.NewRouter() // Register api to server

	// ret.HandleFunc("/kycs",API).Methods("GET")

	ret.HandleFunc("/kycs",API).Methods("POST")

//	 ret.HandleFunc("/kycs/{name}/{dob}/{pan}/{address}/{pincode}",API).Methods("POST")
	
	 http.ListenAndServe(":8080", ret) // server


}