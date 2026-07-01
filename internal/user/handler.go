package user

import (
	"github/morazss/rest-api/internal/handlers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var _ handlers.Handler = &handler{}

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.POST(usersURL, h.GetUserByUUID)
	router.GET(userURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartialUpdateUser)
	router.DELETE(userURL, h.DeleteUser)

}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("Hello world"))
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(201)
	w.Write([]byte("Hello world"))
}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("Hello world"))
}
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("Hello world"))
}
func (h *handler) PartialUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("Hello world"))
}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("Hello world"))
}
