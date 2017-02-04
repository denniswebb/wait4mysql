package main

import (
	"database/sql"
	"os"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"
	"github.com/go-sql-driver/mysql"
)

var (
	username, password, host, port string
	timeout int
)

func init() {
	username = getenv("MYSQL_USERNAME","root")
	password = getenv("MYSQL_PASSWORD","")
	host = getenv("MYSQL_HOST","mysql")
	port = getenv("MYSQL_PORT","3306")

	var err	error
	timeout, err = strconv.Atoi(getenv("MYSQL_TIMEOUT", "60"))
	if err != nil {
		timeout = 60
	}

	mysql.SetLogger(log.New(ioutil.Discard, "[mysql] ", log.Ldate|log.Ltime|log.Lshortfile))
}

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/",username,password,host,port))
	if err != nil {
		fmt.Printf("Unexpected error: %s\n", err)
		os.Exit(1)
	}
	defer db.Close()

	t := 1
	for t < timeout {
		t += 1

		err = db.Ping()
		if err != nil {
			switch err.(type) {
			case *mysql.MySQLError:
				fmt.Printf("Connected.\n")
				os.Exit(0)
			}
		}
		fmt.Print(".")
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("\nOperation timed out in %d seconds.\n", timeout)
	os.Exit(1)
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}