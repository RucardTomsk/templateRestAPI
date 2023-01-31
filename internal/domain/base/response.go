package base

import "github.com/google/uuid"

type Status string
type Blame string

const (
	StatusOK            Status = "OK"
	StatusBadRequest    Status = "Bad Request"
	StatusInternalError Status = "Internal Error"
)

const (
	BlameUser     Blame = "User"
	BlamePostgres Blame = "Postgres"
	BlameServer   Blame = "Server"
)

// ResponseOK is a base OK response from server.
type ResponseOK struct {
	Status     Status `example:"OK"`
	TrackingID string `example:"12345678-1234-1234-1234-000000000000"`
}

// ResponseOKWithGUID is a base OK response from server with additional GUID in answer.
type ResponseOKWithGUID struct {
	Status     Status    `example:"OK"`
	TrackingID string    `example:"12345678-1234-1234-1234-000000000000"`
	GUID       uuid.UUID `example:"12345678-1234-1234-1234-000000000000"`
}

// ResponseFailure is a general error response from server.
type ResponseFailure struct {
	Status     Status `example:"Error"`
	Blame      Blame  `example:"Postgres"`
	TrackingID string `example:"12345678-1234-1234-1234-000000000000"`
	Message    string `example:"error occurred"`
}
