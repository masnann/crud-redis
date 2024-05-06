package helpers

import (
	"crud-redis/models"
	"time"

	"github.com/labstack/echo/v4"
)

func BindAndValidateStruct(ctx echo.Context, i interface{}) error {
	if err := ctx.Bind(i); err != nil {
		return err
	}
	if err := ctx.Validate(i); err != nil {
		return err
	}
	return nil
}

func ResponseJSON(success bool, code, message string, result interface{}) models.Response {
	response := models.Response{
		StatusCode:       code,
		Success:          success,
		Message:          message,
		ResponseDateTime: time.Now(),
		Result:           result,
	}

	return response
}
