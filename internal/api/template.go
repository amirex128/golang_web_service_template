package api

import (
	"fmt"
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func pongo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		name := stringFromContext(c, "template")
		data, _ := c.Get("data")

		if name == "" {
			return
		}

		template := pongo2.Must(pongo2.FromFile(fmt.Sprintf("%s/%s", "./templates", name)))
		err := template.ExecuteWriter(convertContext(data), c.Writer)
		if err != nil {

			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		}
	}
}

func stringFromContext(c *gin.Context, input string) string {
	raw, ok := c.Get(input)
	if ok {
		strVal, ok := raw.(string)
		if ok {
			return strVal
		}
	}
	return ""
}

func convertContext(thing interface{}) pongo2.Context {
	if thing != nil {
		context, isMap := thing.(map[string]interface{})
		if isMap {
			return context
		}
	}
	return nil
}
