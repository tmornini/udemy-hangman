package interfaces

import "net/http"

type Endpointable interface {
	RespondsToPathOf(*http.Request) bool
	RespondsToMethodOf(*http.Request) bool
	Process(*http.Request) (Responsible, error)
}
