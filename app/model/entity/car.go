package entity

type Car struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Brand string `json:"brand,omitempty"`
	Year  string `json:"year,omitempty"`
	Price string `json:"price,omitempty"`
}
