package errorx

import (
	"github.com/amirex128/selloora_backend/internal/providers"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type Errorx struct {
	Message string        `json:"message"`
	Type    string        `json:"type"`
	Args    []interface{} `json:"args"`
	Err     error         `json:"error"`
}

func (e *Errorx) Error() string {
	return e.Err.Error()
}

func ResponseErrorx(c *gin.Context, err error) {
	logrusProvider := do.MustInvoke[*providers.LogrusProvider](providers.Injector)
	e := err.(*Errorx)

	if strings.Contains(e.Type, ":panic") {
		//tr := apm.DefaultTracer().Recovered(e)
		//tr.SetTransaction(apm.TransactionFromContext(c.Request.Context()))
		//tr.Send()
	}

	if e.Err != nil {
		//if strings.Contains(e.Error(), "record not found") {
		//	e.Message = "موردی یافت نشد"
		//}

		logrus.Errorf("Errorx: %s, %s, %s, %s", e.Message, e.Type, e.Args, e.Error())
		logrusProvider.Log.WithFields(logrus.Fields{
			"message": e.Message,
			"type":    e.Type,
			"error":   e.Error(),
			"args":    e.Args,
		}).Error("Errorx")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": e.Message,
			"type":    e.Type,
			"args":    e.Args,
			"error":   e.Err.Error(),
		})
	} else {
		logrus.Errorf("Errorx: %s, %s, %s", e.Message, e.Type)
		logrusProvider.Log.WithFields(logrus.Fields{
			"message": e.Message,
			"type":    e.Type,
			"error":   "",
			"args":    e.Args,
		}).Error("Errorx")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": e.Message,
			"type":    e.Type,
			"args":    e.Args,
			"error":   "",
		})
	}
}

func New(message, typ string, sendError error, args ...interface{}) *Errorx {
	return &Errorx{
		Message: message,
		Type:    typ,
		Args:    args,
		Err:     sendError,
	}
}
