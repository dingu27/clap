package model

//Clap model structure
type Clap struct {
	ImageName string `json:"imgName"`
	Claps     int    `json:"claps"`
}

//Claps Clap slice
type Claps []Clap
