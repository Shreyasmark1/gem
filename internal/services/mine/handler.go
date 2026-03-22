package mine

import (
	middleware "gem/internal/middlewares"
	"gem/pkg/response"
	"net/http"
)

type MineHandler struct {
	mineService *MineService
}

func NewMinehHandler(service *MineService) *MineHandler {
	return &MineHandler{mineService: service}
}

func (h *MineHandler) SaveMine(w http.ResponseWriter, r *http.Request) error {

	if body, err := middleware.DecodeAndValidate[SaveMine](r); err != nil {
		return response.BadRequest("Invalid request body")
	} else {
		res, err := h.mineService.SaveMine(r.Context(), body)

		if err != nil {
			return err
		}

		return response.Send(w, res)
	}

}
