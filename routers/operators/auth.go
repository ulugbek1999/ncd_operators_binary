package operators

import (
	"html/template"
	"log"
	m "ncd_operators/models"
	s "ncd_operators/pkg/settings"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/auth/auth.html"))
	if r.Method == http.MethodGet {
		err := tpl.Execute(w, nil)
		if err != nil {
			_, err = w.Write([]byte(err.Error()))
			if err != nil {
				log.Println(err.Error())
			}
		}
	} else if r.Method == http.MethodPost {
		u := r.FormValue("username")
		log.Println(u)
		p := r.FormValue("password")
		var err error
		var user m.User
		q := `
		SELECT 
		       u.id, u.username, u.password
		FROM 
		     auth_user u 
		WHERE
		     u.username = $1`
		err = s.DB.Get(&user, q, u)
		if err != nil {
			log.Println(err.Error())
			err := tpl.Execute(w, nil)
			if err != nil {
				_, err = w.Write([]byte(err.Error()))
				if err != nil {
					log.Println(err.Error())
				}
			}
		}
		if ok := user.CheckPassword(p); !ok {
			err = tpl.Execute(w, map[string]string{"error": "Username or password incorrect"})
			if err != nil {
				log.Println(err.Error())
				_, err = w.Write([]byte(err.Error()))
				if err != nil {
					log.Println(err.Error())
				}
			}
			return
		}
		value := map[string]string{
			"username": user.Username,
		}
		if encoded, err := s.SecureCookie.Encode("sessionid", value); err == nil {
			cookie := &http.Cookie{
				Name:  "sessionid",
				Value: encoded,
				Path:  "/",
				//Secure:   true,
				//HttpOnly: true,
			}
			http.SetCookie(w, cookie)
		}
		http.Redirect(w, r, "/", 302)
		return
	}

}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:   "sessionid",
		Value:  "",
		MaxAge: 0,
		Path:   "/",
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}
