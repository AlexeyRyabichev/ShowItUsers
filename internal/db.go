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

func GetUserInfo(user *UserHttp) UserHttp {
	connStr := fmt.Sprintf("user=%s password=%s dbname=showit sslmode=disable", username, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("%v", err)
	}
	defer db.Close()

	var userDB UserDB
	query := fmt.Sprintf("select * from users where login = '%s'", user.Login)
	err = db.QueryRow(query).Scan(&userDB.Login, &userDB.Email, &userDB.Password, &userDB.FirstName, &userDB.LastName,
		&userDB.TotalSeries, &userDB.TotalSeenEpisodes, &userDB.TotalUnseenEpisodes, &userDB.TotalSeenMovies, &userDB.TotalUnseenMovies,
		&userDB.TotalTimeSpent, &userDB.YearActivity)
	//err = row.Scan(&userDB.Login, &userDB.Email, &userDB.Password, &userDB.FirstName, &userDB.LastName, &userDB.TotalSeries, &userDB.TotalSeenEpisodes, &userDB.TotalUnseenEpisodes, &userDB.TotalSeenMovies, &userDB.TotalUnseenMovies, &userDB.TotalTimeSpent, &userDB.YearActivity)
	if err != nil {
		log.Printf("cannot get user from db: %v", err)
		return UserHttp{}
	}

	userToReturn := UserHttp{
		Login:               userDB.Login,
		Email:               userDB.Email,
		Password:            userDB.Password,
		FirstName:           userDB.FirstName,
		LastName:            userDB.LastName,
		TotalSeries:         userDB.TotalSeries,
		TotalSeenEpisodes:   userDB.TotalSeenEpisodes,
		TotalUnseenEpisodes: userDB.TotalUnseenEpisodes,
		TotalSeenMovies:     userDB.TotalSeenMovies,
		TotalUnseenMovies:   userDB.TotalUnseenMovies,
		TotalTimeSpent:      []byte(userDB.TotalTimeSpent),
		YearActivity:        []byte(userDB.YearActivity),
	}

	return userToReturn
}

func IsUserExists(user *UserHttp) bool {
	connStr := fmt.Sprintf("user=%s password=%s dbname=showit sslmode=disable", username, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("%v", err)
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

func IsUserLoginOrEmailExists(user *UserHttp) bool {
	connStr := fmt.Sprintf("user=%s password=%s dbname=showit sslmode=disable", username, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("%v", err)
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

func InsertUser(user *UserHttp) bool {
	connStr := fmt.Sprintf("user=%s password=%s dbname=showit sslmode=disable", username, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("%v", err)
	}
	defer db.Close()

	query := fmt.Sprintf("insert into users (login, email, password, first_name, last_name) values ('%s', '%s', '%s', '%s', '%s')", user.Login, user.Email, user.Password, user.FirstName, user.LastName)
	log.Print(query)

	result, err := db.Exec(query)
	if err != nil {
		log.Printf("%v", err)
		return false
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Printf("%v", err)
		return false
	}

	log.Printf("%s user inserted, rows affected: %d", user.Login, rows)
	return true
}

func InsertUserFull(user *UserHttp) bool {
	connStr := fmt.Sprintf("user=%s password=%s dbname=showit sslmode=disable", username, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("%v", err)
	}
	defer db.Close()

	//newTTS := strings.Replace(string(user.TotalTimeSpent), "\n", "", -1)
	//newTTS = strings.Replace(newTTS, "    ", "", -1)
	//
	//newYA := strings.Replace(string(user.YearActivity), "\n", "", -1)
	//newYA = strings.Replace(newYA, "    ", "", -1)

	userDB := UserDB{
		Login:               user.Login,
		Email:               user.Email,
		Password:            user.Password,
		FirstName:           user.FirstName,
		LastName:            user.LastName,
		TotalSeries:         user.TotalSeries,
		TotalSeenEpisodes:   user.TotalSeenEpisodes,
		TotalUnseenEpisodes: user.TotalUnseenEpisodes,
		TotalSeenMovies:     user.TotalSeenMovies,
		TotalUnseenMovies:   user.TotalUnseenMovies,
		TotalTimeSpent:      string(user.TotalTimeSpent),
		YearActivity:        string(user.YearActivity),
	}

	query := fmt.Sprintf("insert into users "+
		"(login, email, password, first_name, last_name, total_series, total_seen_episodes, total_unseen_episodes, total_seen_movies, total_unseen_movies, total_time_spent, year_activity) "+
		"values "+
		"('%s', '%s', '%s', '%s', '%s', %d, %d, %d, %d, %d, '%s', '%s')",
		userDB.Login, userDB.Email, userDB.Password, userDB.FirstName, userDB.LastName,
		userDB.TotalSeries, userDB.TotalSeenEpisodes, userDB.TotalUnseenEpisodes, userDB.TotalSeenMovies, userDB.TotalUnseenMovies,
		userDB.TotalTimeSpent, userDB.YearActivity)
	//query := fmt.Sprintf("insert into users (login, email, password, first_name, last_name) values ('%s', '%s', '%s', '%s', '%s')", user.Login, user.Email, user.Password, user.FirstName, user.LastName)
	//log.Print(query)

	result, err := db.Exec(query)
	if err != nil {
		log.Printf("%v", err)
		return false
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Printf("%v", err)
		return false
	}

	log.Printf("%s user inserted, rows affected: %d", user.Login, rows)
	return true
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
