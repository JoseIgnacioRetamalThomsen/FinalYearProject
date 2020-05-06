package main

import (
	"fmt"
	//"strconv"

	//	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)


func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

/*

func UpdatePlaceRequest(w http.ResponseWriter, r *http.Request){
	log.Printf("Received: %v", "Update place")

}

func CreateCityPostRequest(w http.ResponseWriter, r *http.Request){
	log.Printf("Received: %v", "Craete city post")

	var post pb.CityPostCreateUser
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Wrong request, body must contain city data")
	}
	json.Unmarshal(reqBody, &post)

	 post.CreatorEmail = r.Header["Email"][0]

	fmt.Println(post)

	response , err := CreateCityPost(post)
	json.NewEncoder(w).Encode(response)
}

func GetCityPostRequest(w http.ResponseWriter, r *http.Request){
	log.Printf("Received: %v", "Get city post")

	i1, err := strconv.Atoi(mux.Vars(r)["indexid"])
	if err !=nil{

	}

	response , err := GetCityPosts(pb.PostsRequest{
		IndexId:              int32(i1),

	})

	json.NewEncoder(w).Encode(response)
}


func GetPlacePostRequest(w http.ResponseWriter, r *http.Request){
	log.Printf("Received: %v", "Get place post")

	i1, err := strconv.Atoi(mux.Vars(r)["indexid"])
	if err !=nil{

	}

	response , err := GetPlacePosts(pb.PostsRequest{
		IndexId:              int32(i1),

	})

	json.NewEncoder(w).Encode(response)
}


*/

func main() {

	// read configuration
	args := os.Args[1]
	readConfig(args)

	fmt.Print(configuration.Auth[0])
	//conect to profiles server
	dbserverCtx, err := newProfilesServiceContext(configuration.Profiles[0])
	if err != nil {
		log.Fatal(err)
	}
	s2 := &profileServer{dbserverCtx}
	ProfSerConn = *s2

	//conect to  post server
	dbserverCtx1, err := newPostServiceContext(configuration.Post[0])
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



	//temp, _ := GetAllCities(pb.GetAllRequest{
	//	Max:                  0,
	//
	//})
	//
	//fmt.Print(temp)
	//log.Printf("Received: %v", "here")
	//err, user:= loginUser(pb.UserRequest{
	//	Email:                "email116",
	//	HashPassword:         "password",
	//	Name:                 "user1",
	//
	//})
	//if err!=nil{
	//	panic(err)
	//}
	//log.Printf("Received: %v", user)


	//GetCity(tokenEmail,token,"san pedro","chile")




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
	router.HandleFunc("/city", GetAllCityEndPoint).Methods("GET")//.Queries("search", "{search}")
	router.HandleFunc("/city", GetAllCityEndPoint).Queries("search", "{search}").Methods("GET")
	router.HandleFunc("/city/{country}/{name}/", GetCityRequest).Methods("GET")

	//put not working
	router.HandleFunc("/city",UpdateCityRequest).Methods("PUT")


	router.HandleFunc("/place", CreatePlaceRequest).Methods("POST")
	router.HandleFunc("/place/{country}/{city}", GetCityPlacesEndPoint).Methods("GET")

	router.HandleFunc("/visitcity/{id}", VisitCityEndPoint).Methods("POST")

	router.HandleFunc("/place/{country}/{city}/{name}/", GetPlaceRequest).Methods("GET")
	/*router.HandleFunc("/place/{country}/{city}/{name}/",UpdatePlaceRequest).Methods("PUT")

	//post
	router.HandleFunc("/city/post", CreateCityPostRequest).Methods("POST")
	router.HandleFunc("/posts/city/{indexid}/", GetCityPostRequest).Methods("GET")

	router.HandleFunc("/posts/place/{indexid}/", GetPlacePostRequest).Methods("GET")

*/


	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With","Content-Type","Accept", "Accept-Language", "Content-Language", "Origin"})
	//originsOk := handlers.AllowedOrigins([]string{os.Getenv("*")})
	originsOk1 := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	//	headersOk := handlers.AllowedHeaders([]string{"*"})
////	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
//	originsOk := handlers.AllowedOrigins([]string{"*"})
//	//originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
//	//originsOk := handlers.AllowedOrigins([]string{os.Getenv("*")})
//	methodsOk := handlers.AllowedMethods([]string{"*"})
	//methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS","DELETE"})

	log.Printf("Server started at : %v", configuration.Port)
	log.Fatal(http.ListenAndServe(configuration.Port, handlers.CORS(/*originsOk,*/ headersOk, methodsOk,originsOk1)(router)))
	//log.Fatal(http.ListenAndServe(":8080", router))
}
















///********************* Profiles
const(
	//url = "0.0.0.0:60051"
	url="35.197.216.42:60051";
	//url = "35.234.146.99:5777"
	token ="a31e31a2fcdf2a9a230120ea620f3b24f7379d923fb122323d3cb9bc56fe6508"
	tokenEmail ="a@a.com"
	photoUrl = "35.197.216.42:30051"
)


//profile methos
const(
	DEADLINE = 5;
	)




/*
func GetPlace(request pb.PlaceRequestP)(*pb.PlaceResponseP,error){
	ctx, cancel := context.WithTimeout(context.Background(), DEADLINE*time.Second)
	defer cancel()
	r, err := ProfSerConn.context.dbClient.GetPlace(ctx,&request)
	if err != nil{
		return nil,err
	}

	return r,nil
}

*/



/*


*************************POSTTTT

 */


const (
	POST_url ="35.197.216.42:10051"
)





/**

Photo
 */





