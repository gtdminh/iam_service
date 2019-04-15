package model

import (
	"github.com/julienschmidt/httprouter"
)

type Server struct {
	Router *httprouter.Router
}
