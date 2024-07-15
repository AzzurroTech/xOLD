package main

import (
  "net/http"
)

func main(){
  //BELOW WE ADD OTHER offerings avaliable through Docker
  //root server handling
  http.Handle("/", veni.Load(http.FileServer(http.Dir("./root"))))
  //Other server handling
  http.Handle("/data/", veni.Load(http.StripPrefix("/data/", http.FileServer(http.Dir("./data")))))
  //serve the server
  http.ListenAndServe(":3000", nil)
}
