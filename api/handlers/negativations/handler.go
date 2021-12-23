package negativations

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
	if len(request.Headers["Token"]) == 0 {
		return httping.Unauthorized(map[string]string{"status": "unauthorized"})
	}
	if request.Params["customerDocument"] == "" {
		return httping.BadRequest(map[string]string{"customerDocument": "the field currency is required"})
	}
	err := c.ctl.TokenValid(request.Headers["Token"][0])
	if err != nil {
		return httping.Unauthorized(map[string]string{"error": "Token not valid"})
	}
	negativations, err := c.ctl.GetNegativationByCustomer(request.Params["customerDocument"])
	if err != nil {
		return httping.InternalServerError(map[string]string{"error": err.Error()})
	}
	return httping.OK(map[string]interface{}{"status": "success", "data": negativations})
}
