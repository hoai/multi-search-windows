package main

import (
  "path/filepath"
  "os"
  "flag"
  "fmt"

  "html/template"
  "log"
  "net/http"
  "strings"
  
  "github.com/stretchr/powerwalk"


)
func check(e error) {
    if e != nil {
        panic(e)
    }
}





func main() {
	
	
  http.HandleFunc("/", search) // setting router rule
    //http.HandleFunc("/login", login)
    err := http.ListenAndServe(":9090", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
   
 
  
}

func search(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        // logic part of log in
       
		root := r.Form["dirname"]
		query := r.Form["query"]
		
		fmt.Fprintf(w, root[0])
        fmt.Fprintf(w, query[0])
        
        
        flag.Parse()
        
        file, err := os.Create("result.html")
		check(err)
		defer file.Close()
		
		err = filepath.Walk(root[0], func (path string, f os.FileInfo, err error) error {

			//n3, err := file.WriteString(path)
		
		    
		    if(strings.Contains(path,  query[0])){
				fmt.Fprintf(w, path)
				fmt.Fprintf(w, "\n")
			}
			return nil
		})
		
		
        
    }
}

