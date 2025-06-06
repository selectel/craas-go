package svc

import "errors"

const errGotHTTPStatusCodeFmt = "craas-go: got the %d status code from the server"

var ErrEndpointVersionMismatch = errors.New("endpoint version mismatch")
