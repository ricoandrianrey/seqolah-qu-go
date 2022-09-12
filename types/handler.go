package types

import "context"

type Request struct {
	Params func(key string) string
	Body   func(data interface{})
	Query  func(data interface{})
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Handler func(r Request, c context.Context) (Response, error)
