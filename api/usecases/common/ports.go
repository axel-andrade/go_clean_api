package common

type OutputPort struct {
	StatusCode int16       `json:"-"`
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
}
