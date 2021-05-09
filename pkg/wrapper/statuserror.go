package wrapper

import "errors"

type Status string

// StatOK describe okeeey
var StatOK Status = "OK"

// StatInvalidRequest describe invalid req
var StatInvalidRequest Status = "INVALID_REQUEST"

// StatUnexpectedError describe unexpected err
var StatUnexpectedError Status = "UNEXPECTED_ERROR"

// ErrInternalServer error
var ErrInternalServer = errors.New("Internal Server Error")

// ErrInternalServer error
var ErrInvalidRequest = errors.New("Invalid Request")
