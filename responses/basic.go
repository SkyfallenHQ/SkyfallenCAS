package responses

type Basic struct {

	Status string `json:"status"`
	Error Error `json:"error"`
	Result string `json:"result"`

}
