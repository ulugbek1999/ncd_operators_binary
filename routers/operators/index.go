package operators

import (
	"github.com/gorilla/context"
	m "ncd_operators/models"
	"net/http"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	c := context.Get(r, "user")
	user := c.(*m.User)
	if user.Operator.Operator == 1 {
		http.Redirect(w, r, "/operator-1", 302)
		return
	}
	if user.Operator.Operator == 2 {
		http.Redirect(w, r, "/operator-2-1", 302)
		return
	}
	if user.Operator.Operator == 3 {
		http.Redirect(w, r, "/operator-3-2", 302)
		return
	}
	if user.Operator.Operator == 4 {
		http.Redirect(w, r, "/operator-4-3", 302)
		return
	}
}
