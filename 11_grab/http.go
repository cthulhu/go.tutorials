package main

import (
    "http"
    "fmt"
    "io/ioutil"    
)

func main() {
    repsonce := new( http.Response )
    var readed_bytes []byte
    repsonce, _ , err := http.Get( "http://google.com" )
    if err != nil {
        fmt.Println("Error to read url ", err );
    } else {
         readed_bytes , _ = ioutil.ReadAll(repsonce.Body)
         repsonce.Body.Close();
         fmt.Println("Readed:\n", string( readed_bytes ) );
    }
}

