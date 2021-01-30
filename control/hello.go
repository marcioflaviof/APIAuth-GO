package control

import "net/http"

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, welcome"))
}

func VerifiedHelloWorld(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("auth-token")
	if VerifyToken(tokenString) {
		w.Write([]byte("Hello, welcome"))
		return
	}

	w.Write([]byte("Not Verified"))
	return

}
