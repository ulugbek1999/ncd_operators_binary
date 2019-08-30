package middlewares

import (
	"github.com/gorilla/context"
	m "ncd_operators/models"
	s "ncd_operators/pkg/settings"
	"net/http"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// exclude the urls for AuthMiddleware
		whiteListUrls := []string{"/login", "/logut"}
		for i := range whiteListUrls {
			if r.URL.Path == whiteListUrls[i] {
				next.ServeHTTP(w, r)
				return
			}
		}
		// get `sessionid` cookie
		cookie, err := r.Cookie("sessionid")
		if err != nil {
			// if `sessionid` cookie does not exist,
			// then redirect to login page
			http.Redirect(w, r, "/login", 302)
			return
		}
		// variable for `sessionid` cookie values
		value := make(map[string]string)
		// checking `sessionid` cookie for valid
		err = s.SecureCookie.Decode("sessionid", cookie.Value, &value)
		if err != nil {
			// if `sessionid` cookie invalid,
			// then redirect to login page
			http.Redirect(w, r, "/login", 302)
			return
		}
		q := `
	SELECT
		u.id, u.first_name, u.last_name, u.username, u.password,
	   	op.id AS "op.id", op.channel AS "op.channel", op.operator AS "op.operator", op.phone AS "op.phone",
	   	og.id AS "op.op_group.id", og.group_name AS "op.op_group.group_name",
	    dd.id AS "op.op_group.district.id", dd.name_ru AS "op.op_group.district.name_ru", dd.name_en AS "op.op_group.district.name_en",
	    dc.id AS "op.op_group.district.city.id", dc.name_ru AS "op.op_group.district.city.name_ru",
	       dc.name_en AS "op.op_group.district.city.name_en", dc.code AS "op.op_group.district.city.code"
	FROM
		auth_user u 
	INNER JOIN 
		operator op ON u.id = op.user_id
	INNER JOIN 
		operator_group og ON 
		       op.id = og.operator1_id
		    OR op.id = og.operator2_id
		    OR op.id = og.operator3_id
		    OR op.id = og.operator4_id
	INNER JOIN
		directory_district dd ON og.district_id = dd.id
	INNER JOIN 
		directory_city dc ON dd.city_id = dc.id
	WHERE 
	  	u.username = $1`
		//udb := s.DB.Unsafe()
		var user m.User
		err = s.UDB.Get(&user, q, value["username"])
		if err != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}
		context.Set(r, "user", &user)
		next.ServeHTTP(w, r)
		return
	})
}
