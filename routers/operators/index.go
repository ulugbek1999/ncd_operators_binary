package operators

import (
	"fmt"
	m "ncd_operators/models"
	"net/http"

	"github.com/gorilla/context"
)

//IndexPage handles redirecting according to the operator
func IndexPage(w http.ResponseWriter, r *http.Request) {
	c := context.Get(r, "user")
	user := c.(*m.User)
	if user.Operator.Operator == 1 {
		fmt.Println(user.Operator.Operator)
		http.Redirect(w, r, "/operator-1", 302)
		return
	}
	if user.Operator.Operator == 2 {
		fmt.Println(user.Operator.Operator)
		http.Redirect(w, r, "/operator-2-1", 302)
		return
	}
	if user.Operator.Operator == 3 {
		fmt.Println(user.Operator.Operator)
		http.Redirect(w, r, "/operator-3-2", 302)
		return
	}
	if user.Operator.Operator == 4 {
		fmt.Println(user.Operator.Operator)
		http.Redirect(w, r, "/operator-4-3", 302)
		return
	}
	if user.Operator.Operator == 5 {
		fmt.Println(user.Operator.Operator)
		http.Redirect(w, r, "/operator-universal", 302)
		return
	}
}
