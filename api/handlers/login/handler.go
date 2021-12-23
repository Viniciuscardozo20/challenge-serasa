package login

import (
	"challenge-serasa/api/controller"
	"encoding/json"

	httping "github.com/ednailson/httping-go"
)

type Handler struct {
	ctl controller.Controller
}

type user struct {
	CustomerDocument string `json:"customerDocument"`
}

func NewHandler(ctl controller.Controller) *Handler {
	return &Handler{ctl: ctl}
}

func (c *Handler) Handle(request httping.HttpRequest) httping.IResponse {
	var user user
	err := json.Unmarshal(request.Body, &user)
	if err != nil {
		return httping.BadRequest(map[string]string{"body": "invalid body"})
	}
	token, err := c.ctl.Login(user.CustomerDocument)
	if err != nil {
		return httping.InternalServerError(map[string]string{"error": err.Error()})
	}
	return httping.OK(map[string]interface{}{"status": "success", "token": token})
}
