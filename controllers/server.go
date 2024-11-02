package Controllers

import (

	//"reflect"

	"fmt"
	"net/http"
	"os"
	"time"

	model "IAM/models"

	"github.com/gorilla/mux"
	//	"github.com/gorilla/sessions"
)

//var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func StartServer() error {
	//srvAddr, _ := os.LookupEnv("MB_CONFIG_SRV_HOST")
	//srvPort, _ := os.LookupEnv("MB_CONFIG_DB_PORT")
	srvAddr := "localhost"
	srvPort := "8081"
	err := model.Logsave("id", "123")
	//mux := http.NewServeMux()
	router := mux.NewRouter()
	//feature toggle авторизация клиента
	ftAuth, _ := os.LookupEnv("FT_ROUTES_AUTH")
	if ftAuth == "X" {
		router.HandleFunc("/auth", userAuthCheckHandler)
	}
	//feature toggle создание нового клиента
	ftUserCreate, _ := os.LookupEnv("FT_ROUTES_user_create")
	ftUserCreate = "X"
	if ftUserCreate == "X" {
		router.HandleFunc("/user/create", userCreateHandler)
	}
	ftUserRead, _ := os.LookupEnv("FT_ROUTES_user_read")
	if ftUserRead == "X" {
		// router.HandleFunc("/user/read", userReadHandler)
	}
	ftUserUpdate, _ := os.LookupEnv("FT_ROUTES_user_update")
	if ftUserUpdate == "X" {
		// router.HandleFunc("/user/update", userUpdateHandler)
	}
	ftUserDelete, _ := os.LookupEnv("FT_ROUTES_user_delete")
	if ftUserDelete == "X" {
		// router.HandleFunc("/user/delete", userDeleteHandler)
	}
	handler := LoggingHandler(router)

	err = http.ListenAndServeTLS(fmt.Sprintf("%s:%s", srvAddr, srvPort), "cert/cert.pem", "cert/key.pem", handler)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//

func LoggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		err := model.Logsave("id", fmt.Sprintf("%s %s %s", req.Method, req.RequestURI, time.Since(start)))
		if err != nil {
			fmt.Println("Error log save")
		}

	})
}

func userRegistration(w http.ResponseWriter, r *http.Request) {

}

// func session(w http.ResponseWriter, r *http.Request) {
// 	// Get a session. We're ignoring the error resulted from decoding an
// 	// existing session: Get() always returns a session, even if empty.
// 	session, _ := store.Get(r, "session-name")
// 	// Set some session values.
// 	session.Values["foo"] = "bar"
// 	session.Values[42] = 43
// 	// Save it before we write to the response/return from the handler.
// 	err := session.Save(r, w)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

func userAuthCheckHandler(w http.ResponseWriter, r *http.Request) {

}

func userCreateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now(), ": user")
}

func userReadHandler(w http.ResponseWriter, r *http.Request) {
}
func userUpdateHandler(w http.ResponseWriter, r *http.Request) {
}

func userDeleteHandler(w http.ResponseWriter, r *http.Request) {
}
