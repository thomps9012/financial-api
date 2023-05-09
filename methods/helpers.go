package methods

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type MalformedRequest struct {
	Status int
	Msg    string
}

func (mr *MalformedRequest) Error() string {
	return mr.Msg
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateStruct(destination interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(destination)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func DecodeJSONBody(c *fiber.Ctx, destination interface{}) error {
	if c.Get("Content-Type") != "application/json" {
		msg := "Content-Type header is not application/json"
		return &MalformedRequest{Status: http.StatusUnsupportedMediaType, Msg: msg}
	}

	decoder := json.NewDecoder(bytes.NewReader(c.Body()))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&destination)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf(
				"Request body contains badly-formed JSON (at position %d)",
				syntaxError.Offset,
			)
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

		case errors.Is(err, io.ErrUnexpectedEOF):
			return &MalformedRequest{
				Status: http.StatusBadRequest,
				Msg:    "Request body contains badly-formed JSON",
			}

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf(
				"Request body contains an invalid value for the %q field (at position %d)",
				unmarshalTypeError.Field,
				unmarshalTypeError.Offset,
			)
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}

		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			return &MalformedRequest{Status: http.StatusRequestEntityTooLarge, Msg: msg}

		default:
			return err
		}
	}

	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		msg := "Request body must only contain a single JSON object"
		return &MalformedRequest{Status: http.StatusBadRequest, Msg: msg}
	}

	return nil
}

func Normalize(string_val string) []byte {
	return []byte(string_val)
}

func TimeNowFormat() string {
	return time.Now().Format("2017-09-07 2:3:5 PM")
}
