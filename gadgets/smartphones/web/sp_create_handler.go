package web

import (
	"encoding/json"
	"net/http"
	"phones-review/gadgets/smartphones/gateway"
	"phones-review/gadgets/smartphones/models"
	"phones-review/internal/database"
)

func (h *CreateSmartphoneHandler) SaveSmartphoneHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res, err := h.Create(parseRequest(r))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		m := map[string]interface{}{
			"message": "error in create smartphone",
		}
		_ = json.NewEncoder(w).Encode(m)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(res)
}

type CreateSmartphoneHandler struct {
	gateway.SmartphoneCreateGateway
}

func NewCreateSmartphoneHandler(client *database.MySQLClient) *CreateSmartphoneHandler {
	return &CreateSmartphoneHandler{gateway.NewSmartphoneCreateGateway(client)}
}

func parseRequest(r *http.Request) *models.CreateSmartphoneCMD {
	body := r.Body
	defer body.Close()

	var cmd models.CreateSmartphoneCMD
	_ = json.NewDecoder(body).Decode(&cmd)

	return &cmd
}
