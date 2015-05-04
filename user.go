package deezer

type User struct {
	Id               int    `json:"id"`               //	The user's Deezer ID
	Name             string `json:"name"`             //	The user's Deezer nickname
	Lastname         string `json:"lastname"`         //	The user's last name
	Firstname        string `json:"firstname"`        //	The user's first name
	Email            string `json:"email"`            //	The user's email
	Status           int    `json:"status"`           //	The user's status
	Birthday         int    `json:"birthday"`         //	The user's birthday
	Inscription_date int    `json:"inscription_date"` //	The user's inscription date
	Gender           string `json:"gender"`           //	The user's gender : F or M
	Link             string `json:"link"`             //	The url of the profil for the user on Deezer
	Picture          string `json:"picture"`          //	The url of the user's profil picture.
	Country          string `json:"country"`          //	The user's country
	Lang             string `json:"lang"`             //	The user's language
	Tracklist        string `json:"tracklist"`        //	API Link to the flow of this user
}

type UserList struct {
	Data []User
}
