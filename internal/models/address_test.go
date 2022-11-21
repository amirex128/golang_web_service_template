package models

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/gin-gonic/gin"
	"regexp"
	"testing"
)

func TestCreateAddress(t *testing.T) {
	mockManager := NewMysqlMockManager()
	mockManager.Mock.MatchExpectationsInOrder(false)

	mockManager.Mock.ExpectBegin()

	mockManager.Mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "students" ("id","name")
					VALUES ($1,$2) RETURNING "students"."id"`)).
		WithArgs(s.student.ID, s.student.Name).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(s.student.ID))

	mockManager.Mock.ExpectCommit()

	_, err := mockManager.CreateAddress(&gin.Context{}, context.Background(), DTOs.CreateAddress{})
	if err != nil {

	}
	err = mockManager.Mock.ExpectationsWereMet()

}
