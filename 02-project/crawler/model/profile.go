package model

type Profile struct {
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Age        int    `json:"age"`
	ID         int    `json:"id"`
	Height     int    `json:"Height"`
	Weight     int    `json:"weight"`
	Income     string `json:"income"`
	Marriage   string `json:"Marriage"`
	Education  string `json:"Education"`
	Occupation string `json:"occupation"`
	Hongkou    string `json:"hongkou"`
	Xinzuo     string `json:"xinzuo"`
	Car        string `json:"car"`
	Info       string `json:"info"`
}
