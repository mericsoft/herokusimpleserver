 package main

import(
"io"
"net/http"
"os)

func hello(w http.ResponseWriter, r *http.Request){
io.WriteString(w,"hello world")
}

func main(){
port:=os.Getenv("PORT")
http.HandleFunc("/",hello)
http.ListAndServe(":"+port,nil)
}
