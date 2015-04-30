package result

import (
	"net/http"

	"github.com/anonx/ok/action"
)

// HTML is a result that is returned from actions by default.
// This struct implements action.Result interface.
type HTML struct {
}

// Apply writes to response the result received from action.
func (t *HTML) Apply(w http.ResponseWriter, r *http.Request) {
}

// RenderTemplate initializes and returns HTML type that implements Result interface.
func (t *Middleware) RenderTemplate(templatePath string) action.Result {
	return &HTML{}
}
