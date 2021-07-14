package http

type eventStartRequest struct {
	Type string `json:"type"`
}

type eventFinishRequest struct {
	Type string `json:"type"`
}
