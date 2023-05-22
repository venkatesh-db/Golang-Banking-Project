

//  set DBUSER=root
// set DBPASS=jvt123


package main

import  ( 
         "github.com/go-sql-driver/mysql"
	 "database/sql"
	 "log"
	 "fmt"
	 "os"
	 )

var db *sql.DB

func main() {

    // Capture connection properties.

    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "pgbus",
	AllowNativePasswords: true,
    }

    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }

        result, err := db.Exec("INSERT INTO kyc_dbs (name) VALUES (?)", "anu")

	if err  !=nil{

	}

	if result  !=nil{

	}


       // stmt, err := db.Prepare("INSERT INTO  kyc_dbs(name) VALUES( "anu" )")

    fmt.Println("Connected!")
}