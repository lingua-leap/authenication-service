package handler

type MainHandler interface {
	AuthenticationHandler() AuthenticationHandler
}
