package models

import (
	"encoding/gob"
	"io"
)

type Discount struct {
	ID        int64   `json:"id"`
	Code      string  `json:"code"`
	UserID    int64   `json:"user_id"`
	StartedAt int64   `json:"started_at"`
	EndedAt   int64   `json:"ended_at"`
	Type      string  `json:"type"` // percent, amount
	Value     float32 `json:"value"`
	Percent   float32 `json:"percent"`
	Status    byte    `json:"status"`
}
type DiscountArr []Discount

func (s DiscountArr) Len() int {
	return len(s)
}
func (s DiscountArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s DiscountArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *Discount) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *Discount) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}
