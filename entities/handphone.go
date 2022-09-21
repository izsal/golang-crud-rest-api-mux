package entities

type Handphone struct {
	ID        uint    `json:"id"`
	ModelName string  `json:"modelname"`
	Price     float64 `json:"price"`
	Spek      string  `json:"spec"`
}
