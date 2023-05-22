
package main

import (
    "os"
    "fmt"
)

// set AUTH_USERNAME="balarm"

 type    auth struct {
        username string
        password string


func main() {

 //  party:= new (auth  )

  var happynewyear  auth

happynewyear.username = os.Getenv("AUTH_USERNAME")

happynewyear.password = os.Getenv("AUTH_PASSWORD")

fmt.Println( happynewyear)


}