package internal

type User struct {
	Login     string `json:"login"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	//MoviesViewed  []Movie `json:"movies_viewed"`
	//MoviesToWatch []Movie `json:"movies_to_watch"`
}

type Movie struct {
	Id    string `json:"id"`
	Score string `json:"score"`
}
