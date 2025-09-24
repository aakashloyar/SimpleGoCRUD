package main
import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)
type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`

}
type Director struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)

}
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			//fmt.Println(item)
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
func addMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	if err:=json.NewDecoder(r.Body).Decode(&movie); err!=nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		return 
	}
	movies=append(movies,movie)
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index,item:= range movies {
		if item.ID==params["id"] {
			movies=append(movies[:index],movies[index+1:]...)
			return;
		}
	}
}
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application.json")
	params:=mux.Vars(r)
	for index,item:=range movies {
		if item.ID==params["id"] {
			movies=append(movies[:index],movies[index+1:]...)
			var movie Movie
			if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			movies = append(movies, movie)
		}
	}
}


var movies []Movie
func main() {
	movies = append(movies, Movie{
		ID:       "1",
		Isbn:     "544599",
		Title:    "Movie 1",
		Director: &Director{FirstName: "Aakash", LastName: "Loyar"},
	})
	movies = append(movies, Movie{
		ID:       "2",
		Isbn:     "544699",
		Title:    "Movie 2",
		Director: &Director{FirstName: "Krishan", LastName: "Loyar"},
	})
	r:=mux.NewRouter()
	r.HandleFunc("/movies",getMovies).Methods("Get")
	r.HandleFunc("/movies/{id}",getMovie).Methods("Get")
	r.HandleFunc("/movies/",addMovie).Methods("Post")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("Delete")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("Put")
	fmt.Println("Server starting at Port 8000")
	log.Fatal(http.ListenAndServe(":8000",r))
}