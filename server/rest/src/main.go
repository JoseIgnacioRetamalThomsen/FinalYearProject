package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome World Places And Cities.")
}

func main() {

	// read configuration
	args := os.Args[1]
	readConfig(args)

	fmt.Print(configuration.Auth[0])
	//conect to profiles server
	dbserverCtx, err := newProfilesContextLoadBalancing() //newProfilesServiceContext(configuration.Profiles[0])
	if err != nil {
		log.Fatal(err)
	}
	s2 := &profileServer{dbserverCtx}
	ProfSerConn = *s2

	//conect to  post server
	dbserverCtx1, err := newPostContextLoadBalancing() //newPostServiceContext(configuration.Post[0])
	if err != nil {
		log.Fatal(err)
	}
	s1 := &postService{dbserverCtx1}
	serviceConn = *s1

	//connect to photo service
	dbserverCtx2, err := newPhotosServiceContext(configuration.Photo[0])
	if err != nil {
		log.Fatal(err)
	}
	s3 := &photosServer{dbserverCtx2}
	photoConn = *s3

	//connect to auth service
	dbserverCtx3, err := newAuthServiceContext(configuration.Auth[0])
	if err != nil {
		log.Fatal(err)
	}
	s4 := &authService{dbserverCtx3}
	authConn = *s4

	log.Printf("Received: %v", "here")
	//start load balance connection
	dbserverCtxLB, err := newDBContextLoadBalancing()
	if err != nil {
		log.Fatal(err)
	}
	s5 := &clientDBLoadBalancing{dbserverCtxLB}
	dbConnLB = *s5

	router := mux.NewRouter().StrictSlash(true)
	//auth
	router.HandleFunc("/user", GetUser).Methods("GET")
	router.HandleFunc("/user", CreateNewUser).Methods("POST")
	router.HandleFunc("/user", UptadeUserEndPoint).Methods("PUT")
	router.HandleFunc("/login", LoginEndPoint).Methods("POST")
	router.HandleFunc("/login", LogoutEndPoint).Methods("DELETE")

	//profiles
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/city", CreateCityRequest).Methods("POST")
	router.HandleFunc("/city", GetAllCityEndPoint).Methods("GET")
	router.HandleFunc("/city", GetAllCityEndPoint).Methods("OPTIONS")
	//.Queries("search", "{search}")
	router.HandleFunc("/city", GetAllCityEndPoint).Queries("search", "{search}").Methods("GET")
	router.HandleFunc("/city/{country}/{name}/", GetCityRequest).Methods("GET")

	//put not working
	router.HandleFunc("/city", UpdateCityRequest).Methods("PUT")

	router.HandleFunc("/place", CreatePlaceRequest).Methods("POST")
	router.HandleFunc("/place/{country}/{city}", GetCityPlacesEndPoint).Methods("GET")

	router.HandleFunc("/visitcity/{id}", VisitCityEndPoint).Methods("POST")

	router.HandleFunc("/place/{country}/{city}/{name}/", GetPlaceRequest).Methods("GET")

	router.HandleFunc("/posts/city/{indexid}/", GetCityPostRequest).Methods("GET")

	router.HandleFunc("/posts/place/{indexid}/", GetPlacePostRequest).Methods("GET")

	router.Use(commonMiddleware)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Accept", "Accept-Language", "Content-Language", "Origin", "Access-Control-Request-Headers", "*", "token", "email"})
	//originsOk := handlers.AllowedOrigins([]string{os.Getenv("*")})
	originsOk1 := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "*"})
	log.Printf("Server started at : %v", configuration.Port)
	log.Fatal(http.ListenAndServe(configuration.Port, handlers.CORS( /*originsOk,*/ headersOk, methodsOk, originsOk1)(router)))

}

// add common headers
func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
