package main

import (
 // "path/filepath"
  "os"
  "flag"
  "fmt"

  "html/template"
  "log"
  "net/http"

  
  "github.com/stretchr/powerwalk"
  "runtime"
  
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  


)
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
func check(e error) {
    if e != nil {
        panic(e)
    }
}





func main() {
	 
    
 
	
  http.HandleFunc("/", search) // setting router rule
    //http.HandleFunc("/login", login)
    err = http.ListenAndServe(":9090", nil) // setting listening port
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
       db, err := sql.Open("mysql", "root:12345,rabbit@/search_windows?charset=utf8")
	   
	   checkErr(err)

		defer db.Close()
		
        // Prepare statement for inserting data
    stmtIns, err := db.Prepare("INSERT INTO dir VALUES( ?, ?, ? )") // ? = placeholder
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

   
		root := r.Form["dirname"]
		query := r.Form["query"]
		
		fmt.Fprintf(w, root[0])
        fmt.Fprintf(w, query[0])
        
        runtime.GOMAXPROCS(runtime.NumCPU())
        
        flag.Parse()
        
        file, err := os.Create("result.txt")
		check(err)
		defer file.Close()
	
		err = powerwalk.Walk(root[0], func (path string, f os.FileInfo, err error) error {
            
           
		   file.Write([]byte(path))
		   file.WriteString("\n")
		   //split path to dir and check exist
		   
		   _, err = stmtIns.Exec(i, (i * i)) // Insert tuples (i, i^2)
           if err != nil {
              panic(err.Error()) // proper error handling instead of panic in your app
           }
		    /*if(strings.Contains(path,  query[0])){
				fmt.Fprintf(w, path)
				fmt.Fprintf(w, "\n")
			}*/
			return nil
		})
		
	    
    }
}

