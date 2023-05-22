
package main 

import (
	"fmt"
	"net/http"
	"io"
       "github.com/gin-gonic/gin"

)

func API( req   *gin.Context   ){

  resp, err := http. ("http://localhost:9090/mobiles/kart?title=one plus")

  fmt.Println( "  happy happy year raey ",  resp ,err)

  body, err := io.ReadAll(resp.Body)

  fmt.Println( "  happy happy year raey ",  body ,err)

  fmt.Printf("%s", body)

}



func main(){

	//  1. server 
	//  2. register api to server 
	 
           r := gin.Default()

     v1 := r.Group("/happy")
	{

	  v1.GET("/kart",API )

	  }
	
	 r.Run(":8080")


}