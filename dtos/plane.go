// Package dtos
package dtos

type PlaneResponse struct {
	ID         string `json:"id"`
	Model      string `json:"model"`
	TotalSeats uint   `json:"totalSeats"`
}
