package models

import (
	"encoding/gob"
	"io"
)

type Gallery struct {
	ID        int    `json:"id"`
	ProductID int64  `json:"product_id"`
	Path      string `json:"path"`
	Caption   string `json:"caption"`
	Type      string `json:"type"` // image video cover
	Sort      uint   `json:"sort"`
	IsMain    byte   `json:"is_main"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type GalleryArr []Gallery

func (s GalleryArr) Len() int {
	return len(s)
}
func (s GalleryArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s GalleryArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *Gallery) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *Gallery) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}
