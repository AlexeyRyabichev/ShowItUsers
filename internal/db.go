package internal

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var username = os.Getenv("DBUSER")
var password = os.Getenv("DBPASS")

func IsUserExists(user *UserHttp) bool {
	connStr := fmt.Sprintf("user=%s password=%s dbname=showit sslmode=disable", username, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("ERR\t%v", err)
	}
	defer db.Close()

	var login string
	err = db.QueryRow(fmt.Sprintf(`select login from users where login='%s' and password='%s'`, user.Login, user.Password)).Scan(&login)
	if err != nil {
		log.Printf("ERR\tcannot get user from db: %v", err)
		return false
	}

	if login != "" {
		return true
	} else {
		return false
	}
}

func IsUserLoginOrEmailExists(user *UserHttp) bool {
	connStr := fmt.Sprintf("user=%s password=%s dbname=showit sslmode=disable", username, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("ERR\t%v", err)
	}
	defer db.Close()

	var login string
	err = db.QueryRow(fmt.Sprintf(`select login from users where login='%s' or email='%s'`, user.Login, user.Email)).Scan(&login)
	if err != nil {
		log.Printf("ERR\tcannot get user from db: %v", err)
		return false
	}

	if login != "" {
		return true
	} else {
		return false
	}
}

func GetUserInfo(user *UserHttp) UserHttp {
	connStr := fmt.Sprintf("user=%s password=%s dbname=showit sslmode=disable", username, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("ERR\t%v", err)
	}
	defer db.Close()

	var userDB UserDB
	query := fmt.Sprintf("select * from users where login = '%s'", user.Login)

	err = db.QueryRow(query).Scan(&userDB.Login, &userDB.Email, &userDB.Password, &userDB.FirstName, &userDB.LastName,
		&userDB.TotalSeries, &userDB.TotalSeenEpisodes, &userDB.TotalUnseenEpisodes, &userDB.TotalSeenMovies, &userDB.TotalUnseenMovies,
		&userDB.TotalTimeSpent, &userDB.YearActivity)
	if err != nil {
		log.Printf("ERR\tcannot get user from db: %v", err)
		return UserHttp{}
	}
	return DBtoHTTP(userDB)
}

func InsertUser(user *UserHttp) bool {
	connStr := fmt.Sprintf("user=%s password=%s dbname=showit sslmode=disable", username, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("ERR\t%v", err)
	}
	defer db.Close()

	query := fmt.Sprintf("insert into users (login, email, password, first_name, last_name) values ('%s', '%s', '%s', '%s', '%s')", user.Login, user.Email, user.Password, user.FirstName, user.LastName)
	result, err := db.Exec(query)
	if err != nil {
		log.Printf("ERR\t%v", err)
		return false
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Printf("%v", err)
		return false
	}

	log.Printf("user %s inserted, rows affected: %d", user.Login, rows)
	return true
}

func InsertUserFull(user *UserHttp) bool {
	connStr := fmt.Sprintf("user=%s password=%s dbname=showit sslmode=disable", username, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("ERR\t%v", err)
	}
	defer db.Close()

	userDB := HTTPtoDB(*user)

	query := fmt.Sprintf("insert into users "+
		"(login, email, password, first_name, last_name, total_series, total_seen_episodes, total_unseen_episodes, total_seen_movies, total_unseen_movies, total_time_spent, year_activity) "+
		"values "+
		"('%s', '%s', '%s', '%s', '%s', %d, %d, %d, %d, %d, '%s', '%s')",
		userDB.Login, userDB.Email, userDB.Password, userDB.FirstName, userDB.LastName,
		userDB.TotalSeries, userDB.TotalSeenEpisodes, userDB.TotalUnseenEpisodes, userDB.TotalSeenMovies, userDB.TotalUnseenMovies,
		userDB.TotalTimeSpent, userDB.YearActivity)

	result, err := db.Exec(query)
	if err != nil {
		log.Printf("ERR\t%v", err)
		return false
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Printf("ERR\t%v", err)
		return false
	}

	log.Printf("DB\t%s user inserted, rows affected: %d", user.Login, rows)
	return true
}
