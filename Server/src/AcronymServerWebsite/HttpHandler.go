package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type ApiHandler struct {
	logger     *log.Logger
	repository AcronymRepository
}

type RouteHandler struct {
	Route   string
	Handler http.Handler
}

func NewRouteHandler(route string, handler http.Handler) *RouteHandler {
	routeHandler := &RouteHandler{
		Route:   route,
		Handler: handler,
	}

	return routeHandler
}

type RoutingApiHandler struct {
	logger *log.Logger
	routes []RouteHandler
}

func NewRoutingApiHandler(routes []RouteHandler, logger *log.Logger) *RoutingApiHandler {
	handler := new(RoutingApiHandler)
	handler.routes = routes
	handler.logger = logger

	return handler
}

func (h *RoutingApiHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	for _, element := range h.routes {
		route := element.Route
		handler := element.Handler

		isPrefix := strings.HasPrefix(request.URL.Path, route)
		if isPrefix {
			log.Printf("[RoutingApiHandler.ServerHTTP] Serving on %s: %s ", route, request.URL.Path)
			handler.ServeHTTP(writer, request)
			return
		}

	}

	log.Printf("[RoutingApiHandler.ServerHTTP] Serving: %s", request.URL.Path)
	badRequest(writer, errors.New("Handler not found"))
}

type AcronymApiHandler struct {
	controller *AcronymApiController
}

func NewAcronymApiHandler(controller *AcronymApiController) *AcronymApiHandler {
	handler := new(AcronymApiHandler)
	handler.controller = controller

	return handler
}

func (h *AcronymApiHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	out, err := h.controller.HandleRequest(request.URL.Path, nil)
	if err != nil {
		badRequest(writer, err)
		return
	}

	ok(writer, out)
}

func badRequest(writer http.ResponseWriter, err error) {
	writer.WriteHeader(500)
	fmt.Fprintf(writer, "%v", err)
}

func ok(writer http.ResponseWriter, out string) {
	writer.WriteHeader(200)
	fmt.Fprint(writer, out)
}
