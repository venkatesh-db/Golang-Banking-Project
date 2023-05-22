package main

import ("fmt"
               "net/http"
	       )

func zoder() int {

    fmt.Println("zoder")

    return 18
}

func coder( a int ) int {

    fmt.Println("coder" , a)

    return 9
}



type HandlerFunc func(  res http.ResponseWriter,  req *http.Request)

func coding(  ) http. HandlerFunc  {

        return   func(  res http.ResponseWriter,  req *http.Request){

    
                        }
}

func codings()  func(  res http.ResponseWriter,  req *http.Request)   {

return  func(  res http.ResponseWriter,  req *http.Request){

    
                        }

}

// type HandlerFunc func(ResponseWriter, *Request)

func messageHandler(message string) http.Handler {
                                                                                 // interface

	return http.HandlerFunc(  func(w http.ResponseWriter, r *http.Request) {
		                                                         w.Write([]byte(message))
                                                       })

}


func main(){

    type  see  func(int) int

      var  jvt   see

	   jvt = coder

	   jvt(5)

	    fmt.Println("coderrange")

	   fmt.Println(  coder(  zoder()  ) )


}