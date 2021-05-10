package responses

import "backend/structures"

type TokenResponse struct {

	Status string `json:"status"`
	Error Error `json:"error"`
	Token structures.AuthToken `json:"token"`

}
