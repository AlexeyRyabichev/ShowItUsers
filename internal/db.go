package internal

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	//"path/filepath"
)

var username = os.Getenv("DBUSER")
var password = os.Getenv("DBPASS")

type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	return n, nil
}

func IsUserExists(user *User) bool {
	connStr := fmt.Sprintf("user=%s password=%s dbname=showit sslmode=disable", username, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var login string
	err = db.QueryRow(fmt.Sprintf(`select login from users where login='%s' and password='%s'`, user.Login, user.Password)).Scan(&login)
	if err != nil {
		log.Printf("cannot get user from db: %v", err)
		return false
	}

	if login != "" {
		return true
	} else {
		return false
	}
}

func IsUserLoginOrEmailExists(user *User) bool {
	connStr := fmt.Sprintf("user=%s password=%s dbname=showit sslmode=disable", username, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var login string
	err = db.QueryRow(fmt.Sprintf(`select login from users where login='%s' or email='%s'`, user.Login, user.Email)).Scan(&login)
	if err != nil {
		log.Printf("cannot get user from db: %v", err)
		return false
	}

	if login != "" {
		return true
	} else {
		return false
	}
}

//func Upload(filename *FileName) {
//	log.Printf("%s upload started", filename.FileName())
//
//	connStr := fmt.Sprintf("user=%s password=%s dbname=showit sslmode=disable", username, password)
//	db, err := sql.Open("postgres", connStr)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//
//	path, err := filepath.Abs("/opt/" + filename.FileName())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	result, err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", filename.DBName()))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	result, err = db.Exec(fmt.Sprintf("COPY %s FROM '%s'", filename.DBName(), path))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	rows, err := result.RowsAffected()
//	log.Printf("%s upload finished, rows affected: %s", filename.FileName(), rows)
//}
