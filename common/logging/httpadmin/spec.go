package httpadmin

type Logging interface {
	ActivateSpec(spec string) error
	Spec() string
}

type LogSpec struct {
	Spec string `json:"spec,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SpecHandler struct {
	Logging Logging
	
}