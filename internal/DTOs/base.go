package DTOs

import (
	"github.com/brianvoe/gofakeit/v6"
	"math/rand"
	"strings"
)

func init() {

	gofakeit.AddFuncLookup("custom_string", gofakeit.Info{
		Category:    "custom",
		Description: "custom data",
		Example:     "custom",
		Output:      "string",
		Params: []gofakeit.Param{
			{Field: "word", Type: "string", Description: "custom string"},
		},
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
			word, err := info.GetString(m, "word")
			if err != nil {
				return nil, err
			}

			return word, nil
		},
	})

	gofakeit.AddFuncLookup("custom_enum", gofakeit.Info{
		Category:    "custom",
		Description: "custom data",
		Example:     "custom",
		Output:      "string",
		Params: []gofakeit.Param{
			{Field: "word", Type: "string", Description: "custom string"},
		},
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
			word, err := info.GetString(m, "word")
			if err != nil {
				return nil, err
			}
			split := strings.Split(word, ",")
			gofakeit.ShuffleStrings(split)
			return split[0], nil
		},
	})

	gofakeit.AddFuncLookup("custom_uint64", gofakeit.Info{
		Category:    "custom",
		Description: "custom data",
		Example:     "custom",
		Output:      "string",
		Params: []gofakeit.Param{
			{Field: "word", Type: "uint64", Description: "custom string"},
		},
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
			word, err := info.GetString(m, "word")
			if err != nil {
				return nil, err
			}

			return word, nil
		},
	})

	gofakeit.AddFuncLookup("custom_uint32", gofakeit.Info{
		Category:    "custom",
		Description: "custom data",
		Example:     "custom",
		Output:      "string",
		Params: []gofakeit.Param{
			{Field: "word", Type: "uint32", Description: "custom string"},
		},
		Generate: func(r *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (interface{}, error) {
			word, err := info.GetString(m, "word")
			if err != nil {
				return nil, err
			}

			return word, nil
		},
	})
}

type Index struct {
	Search   string `form:"search" json:"search"`
	Page     uint32 `form:"page" json:"page" validate:"numeric"`
	PageSize uint32 `form:"page_size" json:"page_size" validate:"numeric"`
	Sort     string `form:"sort" json:"sort"`
}
