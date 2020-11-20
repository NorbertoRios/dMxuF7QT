package interfaces

import (
	"net/http"
)

//IRequest ...
type IRequest interface {
	Request() *http.Request
}
