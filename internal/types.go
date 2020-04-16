package internal

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

	//"total_seen_episodes": 2603,
	//"total_unseen_episodes": 36,
	//"total_seen_movies": 205,
	//"total_unseen_movies": 36,

	//MoviesViewed  []Movie `json:"movies_viewed"`
	//MoviesToWatch []Movie `json:"movies_to_watch"`
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
