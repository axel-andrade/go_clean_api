package common

type OutputPort struct {
	StatusCode int16  `json:"-"`
	Data       any    `json:"data,omitempty"`
	Error      string `json:"error,omitempty"`
}
