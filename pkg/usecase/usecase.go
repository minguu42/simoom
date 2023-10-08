package usecase

import "github.com/cockroachdb/errors"

const userID = "01DXF6DT000000000000000000"

var (
	ErrProjectNotFound = errors.New("the specified project is not found")
	ErrUnkown          = errors.New("an unintentional error occurred on the server")
)
