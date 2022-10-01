package validations

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var (
	validate *validator.Validate
)

func init() {
	validate = validator.New()
}

type ValidationTags map[string]map[string]string

func validateTags(items ValidationTags, err error, c *gin.Context) error {

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "مقادیر ارسال شده نا درست میباشد",
				"error":   err.Error(),
			})
			return errors.New("validation error")

		}
		var validationErrors = make(map[string]map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			for s, m := range items {
				if err.StructField() == s {
					for t, v := range m {
						if err.Tag() == t {
							validationErrors[s] = map[string]string{
								t: v,
							}
						}
					}
				}
			}
		}

		if len(validationErrors) > 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "مقادیر ارسال شده نا درست میباشد",
				"errors":  validationErrors,
			})
			return err
		}
	}
	return nil

}
