package models

type Input struct {
	Startdate string `json:"startdate"`
	Enddate   string `json:"enddate"`
}

type TotalRevenuebyproduct struct {
	Productname string  `json:"productname"`
	Revenue     float64 `json:"revenue"`
}

type TotalRevenueByCategory struct {
	Category string  `json:"category"`
	Revenue  float64 `json:"revenue"`
}
