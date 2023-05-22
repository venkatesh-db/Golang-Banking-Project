package main 

import (
	"fmt"
       "github.com/gin-gonic/gin"
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

func API( req   *gin.Context   ){

	// 5. Inside api 
       // above 3 steps   
 
      var  inputs KycInput

       req.Bind(&inputs)

	if len(inputs.Name) <0 {
			    req.JSON(201,"invalid name")
	}
 
	if len(inputs.Pan) >10 {
		    req.JSON(201,"invalid pan")
	}

         var dbinsert KycDB 

         dbinsert.Name = inputs.Name

	db.Create(&dbinsert)

	var resposne KycResponse = KycResponse{ Message:"partner push completed succeesfuly" }

       req.JSON(200,resposne)

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
	 
           r := gin.Default()

	  r.POST("/kycs",API )
	
	 r.Run()


}