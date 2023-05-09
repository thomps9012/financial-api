package models

import (
	"context"
	database "financial-api/db"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Category string

const (
	IOP            Category = "IOP"
	INTAKE         Category = "INTAKE"
	PEERS          Category = "PEERS"
	ACT_TEAM       Category = "ACT_TEAM"
	IHBT           Category = "IHBT"
	PERKINS        Category = "PERKINS"
	MENS_HOUSE     Category = "MENS_HOUSE"
	NEXT_STEP      Category = "NEXT_STEP"
	LORAIN         Category = "LORAIN"
	PREVENTION     Category = "PREVENTION"
	ADMINISTRATIVE Category = "ADMINISTRATIVE"
	FINANCE        Category = "FINANCE"
)

type Month int

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

type MonthlyRequestInput struct {
	Month Month `json:"month" bson:"month" validate:"required"`
	Year  int   `json:"year" bson:"year" validate:"required"`
}

type ApproveRejectRequest struct {
	RequestID string `json:"request_id" bson:"request_id" validate:"required"`
}

type CustomError struct {
	Status  int    `json:"status" bson:"status"`
	Message string `json:"message" bson:"message"`
}

func (e *CustomError) Error() string {
	return e.Message
}

type ErrorLog struct {
	ID           string    `json:"id" bson:"_id"`
	UserID       string    `json:"user_id" bson:"user_id" validate:"required"`
	Status       int       `json:"error_status" bson:"error_status" validate:"required"`
	Error        string    `json:"error" bson:"error" validate:"required"`
	ErrorPath    string    `json:"error_path" bson:"error_path" validate:"required"`
	ErrorMessage string    `json:"error_message" bson:"error_message" validate:"required"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
}

type ErrorLogOverview struct {
	ID           string    `json:"id" bson:"_id"`
	ErrorMessage string    `json:"error_message" bson:"error_message" validate:"required"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
}

func (el *ErrorLog) Save() (ErrorLogOverview, error) {
	el.ID = uuid.NewString()
	el.CreatedAt = time.Now()
	collection, err := database.Use("error_logs")
	if err != nil {
		return ErrorLogOverview{}, err
	}
	_, err = collection.InsertOne(context.TODO(), el)
	if err != nil {
		return ErrorLogOverview{}, err
	}

	return ErrorLogOverview{
		ID:           el.ID,
		ErrorMessage: el.ErrorMessage,
		CreatedAt:    el.CreatedAt,
	}, nil
}

func ToFixed(num float64, precision int) float64 {
	string_output := strconv.FormatFloat(num, 'f', precision+1, 64)
	string_array := strings.Split(string_output, ".")
	cents := string_array[1]
	dollars := string_array[0]
	if len(cents) < precision {
		cents = strings.Repeat("0", precision-len(cents)) + cents
	}
	if len(cents) == precision+1 {
		cents = cents[0:precision]
	}
	output, _ := strconv.ParseFloat(dollars+"."+cents, 64)
	return output
}
