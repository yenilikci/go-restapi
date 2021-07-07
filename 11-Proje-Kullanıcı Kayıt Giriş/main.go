package main

import (
	"fmt"
	"net/http"

	helper "./helpers"
)

func main() {

	uName, email, pwd, pwdConfirm := "", "", "", ""

	mux := http.NewServeMux()

	//signup
	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		/*
			for key, value := range r.Form {
				fmt.Printf("%s = %s\n", key, value)
			}
		*/
		//gelen key-value
		uName = r.FormValue("username")     //username-key->value
		email = r.FormValue("email")        //email-key->value
		pwd = r.FormValue("password")       //password-key->value
		pwdConfirm = r.FormValue("confirm") //confirm-key->value

		//veri bütünlüğü - boş mu veri var mı
		uNameCheck := helper.IsEmpty(uName)
		emailCheck := helper.IsEmpty(email)
		pwdCheck := helper.IsEmpty(pwd)
		pwdConfirmCheck := helper.IsEmpty(pwdConfirm)

		if uNameCheck || emailCheck || pwdCheck || pwdConfirmCheck {
			fmt.Fprintf(w, "There is empty data!")
			return
		}
		if pwd == pwdConfirm {
			fmt.Fprintf(w, "Registration succesful!")
		} else {
			fmt.Fprintf(w, "Password information must be tha same!")
		}

	})

	//login
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		email = r.FormValue("email")
		pwd = r.FormValue("password")

		emailCheck := helper.IsEmpty(email)
		pwdCheck := helper.IsEmpty(pwd)

		if emailCheck || pwdCheck {
			fmt.Fprintf(w, "There is empty data!")
			return
		}

		dbPwd := "12345!*."
		dbEmail := "melih@github.com"

		if email == dbEmail && pwd == dbPwd {
			fmt.Fprintln(w, "Login succesful!")
		} else {
			fmt.Fprintln(w, "Login failed!")
		}

	})

	http.ListenAndServe(":8080", mux)

}
