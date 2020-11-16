package models

// Smartphone model structure for smartphone
type Smartphone struct {
	Id            int64
	Name          string
	Price         int
	CountryOrigin string
	OS            string
}

// CreateSmartphoneCMD to create a new smartphone
type CreateSmartphoneCMD struct {
	Name          string `json:"name"`
	Price         int    `json:"price"`
	CountryOrigin string `json:"country_origin"`
	OS            string `json:"os"`
}
