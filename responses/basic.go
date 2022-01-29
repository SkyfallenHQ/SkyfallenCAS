package responses

type Basic struct {

	Status string `json:"status"`
	Error Error `json:"error"`
	Result interface{} `json:"result"`

}
