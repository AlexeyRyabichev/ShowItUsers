package internal

type User struct {
	Login         string  `json:"login"`
	Password      string  `json:"password"`
	Name          string  `json:"name"`
	MoviesViewed  []Movie `json:"movies_viewed"`
	MoviesToWatch []Movie `json:"movies_to_watch"`
}

type Movie struct {
	Id    string `json:"id"`
	Score string `json:"score"`
}
