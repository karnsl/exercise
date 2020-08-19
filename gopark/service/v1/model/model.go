package model

// Place a model of table places
type Place struct {
	ID   int16  `json:"id"`
	Name string `json:"name"`
}

// Lot a model of table lot
type Lot struct {
	ID       int16  `json:"id"`
	PlaceID  int16  `json:"place_id"`
	Building string `json:"building"`
	Floor    string `json:"floor"`
	Zone     string `json:"zone"`
	Number   int8   `json:"number"`
	Username string `json:"username"`
}

// Account a model of table accounts.
type Account struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Channel     string `json:"channel"`
	DisplayName string `json:"display_name"`
	Active      bool   `json:"active"`
}
