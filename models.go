package models

// Car represents the data model for a car
type Car struct {
	ID    int    `json:"id"`
	Model string `json:"model"`
	Make  string `json:"make"`
	Year  int    `json:"year"`
	// Add more fields as needed
}
