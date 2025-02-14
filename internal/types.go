package internal

import "encoding/json"

type User struct {
	Login               string          `json:"login"`
	Email               string          `json:"email"`
	Password            string          `json:"password"`
	FirstName           string          `json:"first_name"`
	LastName            string          `json:"last_name"`
	TotalSeries         int             `json:"total_series"`
	TotalSeenEpisodes   int             `json:"total_seen_episodes"`
	TotalUnseenEpisodes int             `json:"total_unseen_episodes"`
	TotalSeenMovies     int             `json:"total_seen_movies"`
	TotalUnseenMovies   int             `json:"total_unseen_movies"`
	TotalTimeSpent      TimeSpent       `json:"total_time_spent"`
	YearActivity        []MonthActivity `json:"year_activity"`
}

type UserHttp struct {
	Login               string          `json:"login"`
	Email               string          `json:"email"`
	Password            string          `json:"password"`
	FirstName           string          `json:"first_name"`
	LastName            string          `json:"last_name"`
	TotalSeries         int             `json:"total_series"`
	TotalSeenEpisodes   int             `json:"total_seen_episodes"`
	TotalUnseenEpisodes int             `json:"total_unseen_episodes"`
	TotalSeenMovies     int             `json:"total_seen_movies"`
	TotalUnseenMovies   int             `json:"total_unseen_movies"`
	TotalTimeSpent      json.RawMessage `json:"total_time_spent"`
	YearActivity        json.RawMessage `json:"year_activity"`
}

type UserDB struct {
	Login               string `db:"login"`
	Email               string `db:"email"`
	Password            string `db:"password"`
	FirstName           string `db:"first_name"`
	LastName            string `db:"last_name"`
	TotalSeries         int    `db:"total_series"`
	TotalSeenEpisodes   int    `db:"total_seen_episodes"`
	TotalUnseenEpisodes int    `db:"total_unseen_episodes"`
	TotalSeenMovies     int    `db:"total_seen_movies"`
	TotalUnseenMovies   int    `db:"total_unseen_movies"`
	TotalTimeSpent      string `db:"total_time_spent"`
	YearActivity        string `db:"year_activity"`
}

type TimeSpent struct {
	Years  int `json:"years"`
	Months int `json:"months"`
	Days   int `json:"days"`
	Hours  int `json:"hours"`
}

type MonthActivity struct {
	Month string `json:"month"`
	Hours int    `json:"hours"`
}

type Movie struct {
	Id    string `json:"id"`
	Score string `json:"score"`
}

func HTTPtoDB(http UserHttp) UserDB {
	return UserDB{
		Login:               http.Login,
		Email:               http.Email,
		Password:            http.Password,
		FirstName:           http.FirstName,
		LastName:            http.LastName,
		TotalSeries:         http.TotalSeries,
		TotalSeenEpisodes:   http.TotalSeenEpisodes,
		TotalUnseenEpisodes: http.TotalUnseenEpisodes,
		TotalSeenMovies:     http.TotalSeenMovies,
		TotalUnseenMovies:   http.TotalUnseenMovies,
		TotalTimeSpent:      string(http.TotalTimeSpent),
		YearActivity:        string(http.YearActivity),
	}
}

func DBtoHTTP(db UserDB) UserHttp {
	return UserHttp{
		Login:               db.Login,
		Email:               db.Email,
		Password:            db.Password,
		FirstName:           db.FirstName,
		LastName:            db.LastName,
		TotalSeries:         db.TotalSeries,
		TotalSeenEpisodes:   db.TotalSeenEpisodes,
		TotalUnseenEpisodes: db.TotalUnseenEpisodes,
		TotalSeenMovies:     db.TotalSeenMovies,
		TotalUnseenMovies:   db.TotalUnseenMovies,
		TotalTimeSpent:      []byte(db.TotalTimeSpent),
		YearActivity:        []byte(db.YearActivity),
	}
}
