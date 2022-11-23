package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.PathPrefix("/public").Handler(http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/formproject", formproject).Methods("GET")
	route.HandleFunc("/detail", detail).Methods("GET")
	route.HandleFunc("/add-project", addproject).Methods("POST")
	

	fmt.Println("Server running on port 5000")
	http.ListenAndServe("Localhost:5000", route)
}


func home(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Contain type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/index.html")

	if err != nil{
		w.Write([]byte("Mesage :"+ err.Error()))
		return
	}
	tmpt.Execute(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Contain type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/form.html")

	if err != nil{
		w.Write([]byte("Mesage :"+ err.Error()))
		return
	}
	tmpt.Execute(w, nil)
}

func formproject(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Contain type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/formproject.html")

	if err != nil{
		w.Write([]byte("Mesage :"+ err.Error()))
		return
	}
	tmpt.Execute(w, nil)
}

func detail(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Contain type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/detail.html")

	if err != nil{
		w.Write([]byte("Mesage :"+ err.Error()))
		return
	}
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// object golang
	data := map[string]interface{}{
		"Title":   "Pasar Coding Dari Dumbways",
		"Content": "REPUBLIKA.CO.ID, JAKARTA -- Ketimpangan sumber daya manusia (SDM) di sektor digital masih menjadi isu yang belum terpecahkan.",
		"Id":      id,
	}

	tmpt.Execute(w, data)
}

func addproject(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}
		fmt.Println("Title : " + r.PostForm.Get("title"))
		fmt.Println("StartDate : " + r.PostForm.Get("startDate"))
		fmt.Println("EndDate : " + r.PostForm.Get("endDate"))
		fmt.Println("Description : " + r.PostForm.Get("description"))
		fmt.Println("Checkbox1 : " + r.PostForm.Get("node"))
		fmt.Println("Checkbox2 : " + r.PostForm.Get("next"))
		fmt.Println("Checkbox3 : " + r.PostForm.Get("react"))
		fmt.Println("Checkbox4 : " + r.PostForm.Get("typescript"))

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		log.Println(" --> Data Berhasil di Kirim")
}