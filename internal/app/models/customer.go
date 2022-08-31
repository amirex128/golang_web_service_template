package models

import (
	"encoding/gob"
	"io"
)

type Customer struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Family     string `json:"family"`
	Mobile     string `json:"mobile"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Address    string `json:"address"`
	PostalCode string `json:"postal_code"`
	Lat        string `json:"lat"`
	Lng        string `json:"lng"`
	CreatedAt  string `json:"created_at"`
}
type CustomerArr []Customer

func (s CustomerArr) Len() int {
	return len(s)
}
func (s CustomerArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s CustomerArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *Customer) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *Customer) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}
