package main 

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "gorm.io/gorm"
)

type  KycInput struct {
	 Name     string        `form:"name"`
        Dob      string              `form:"dob"`
        Pan      string                `form:"pan"`
	Address  string       `form:"address"`
	Pincode   string   `form:"pincode"`
}

type  KycDB struct {
	Name     string      
    Dob      string            
    Pan      string             
	Address  string   
	Pincode   string    
}

type KycResponse struct{
	Message string
}

var db *gorm.DB

func API(   res   http.ResponseWriter , req   *http.Request  ){

	// 5. Inside api 
       // above 3 steps   
 

	//rets := mux.Vars(req) 

	var inputs KycInput

	id := req.URL.Query().Get("name")

	/inputs.Name = rets["name"]
	//inputs.Dob = rets["dob"]
	//inputs.Pan = rets["pan"]
	//inputs.Address = rets["address"]
	//inputs.Pincode = rets["pincode"]


	if len(inputs.Name) <0 {
		        json.NewEncoder(res).Encode("invalid jamon")
	}
 
	if len(inputs.Pan) >10 {
		     	json.NewEncoder(res).Encode("invalid momos ")
	}

         var dbinsert KycDB 

         dbinsert.Name = id

	db.Create(&dbinsert)

	var resposne KycResponse = KycResponse{ Message:"partner push completed succeesfuly" }

  	json.NewEncoder(res).Encode(resposne)

}

     // 3. connect to db
	 // 4. create table 

	 func init() {

	var err error

	db, err = gorm.Open("mysql", "root:jvt123@tcp(127.0.0.1:3306)/pgbus?charset=utf8&parseTime=True")

	if err != nil {
		fmt.Print("jvt", err)
		panic("db not connected")

	}
	db.AutoMigrate(&KycDB{})

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