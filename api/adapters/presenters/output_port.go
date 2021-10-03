package output

type OutputPort struct {
	StatusCode int16
	Data       interface{}
	Error      error
}
