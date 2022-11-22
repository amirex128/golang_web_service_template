package errorx

import (
	"fmt"
	"github.com/amirex128/selloora_backend/internal/providers"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Errorx struct {
	Message string        `json:"message"`
	Type    string        `json:"type"`
	Args    []interface{} `json:"args"`
	Err     error         `json:"error"`
}

func (e *Errorx) Error() string {
	return fmt.Sprintf("Errorx: %s, %s, %s", e.Message, e.Type, e.Err.Error())
}

func ResponseErrorx(c *gin.Context, err error) {
	e := err.(*Errorx)
	if e.Err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": e.Message,
			"type":    e.Type,
			"args":    e.Args,
			"error":   e.Err.Error(),
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": e.Message,
			"type":    e.Type,
			"args":    e.Args,
			"error":   "",
		})
	}
}

func New(message, typ string, sendError error, args ...interface{}) *Errorx {
	logrusProvider := do.MustInvoke[*providers.LogrusProvider](providers.Injector)
	if sendError != nil {
		logrus.Errorf("Errorx: %s, %s, %s", message, typ, sendError.Error())
		logrusProvider.Log.WithFields(logrus.Fields{
			"message": message,
			"type":    typ,
			"error":   sendError.Error(),
			"args":    args,
		}).Error("Errorx")
	} else {
		logrus.Errorf("Errorx: %s, %s, %s", message, typ)
		logrusProvider.Log.WithFields(logrus.Fields{
			"message": message,
			"type":    typ,
			"error":   "",
			"args":    args,
		}).Error("Errorx")
	}

	return &Errorx{
		Message: message,
		Type:    typ,
		Args:    args,
		Err:     sendError,
	}
}
