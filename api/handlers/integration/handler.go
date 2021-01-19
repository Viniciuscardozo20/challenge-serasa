package integration

import (
	"challenge-serasa/api/controller"

	httping "github.com/ednailson/httping-go"
)

type Handler struct {
	ctl controller.Controller
}

func NewHandler(ctl controller.Controller) *Handler {
	return &Handler{ctl: ctl}
}

func (c *Handler) Handle(request httping.HttpRequest) httping.IResponse {
	err := c.ctl.UpdateNegativations()
	if err != nil {
		return httping.InternalServerError(map[string]string{"error": err.Error()})
	}
	return httping.OK(map[string]interface{}{"status": "success"})
}
