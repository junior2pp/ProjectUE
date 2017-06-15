package main

import (
	"net/http"
	"text/template"
	"log"
	"github.com/gorilla/mux"
	"time"
	"fmt"
	"database/sql"
	"os"
)

func main() {
	r := mux.NewRouter()
	h := r.PathPrefix("/").Subrouter()
	h.HandleFunc("/", home)
	h.HandleFunc("/introduccion",introduccion)
	h.HandleFunc("/objetivos",objetivos)
	h.HandleFunc("/metodologia",metodologia)
	h.HandleFunc("/resultados",resultados)
	h.HandleFunc("/presupuesto",presupuesto)
	h.HandleFunc("/instituciones",instituciones)
	h.HandleFunc("/herramientas",herramientas)
	h.HandleFunc("/login",login)
	h.HandleFunc("/principal",principal)
	h.HandleFunc("/parrafo", parrafo)
	h.HandleFunc("/go", golang)
	h.HandleFunc("/medioambiente", medioambiete)

	//Subenrutador de noticias
	n := r.PathPrefix("/noticia").Subrouter()
	n.HandleFunc("/{id:[0-9]+}", noticia)
	n.HandleFunc("/nueva", nueva)
	n.HandleFunc("/", lista)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}

func home(w http.ResponseWriter, r *http.Request)  {
	t,err := template.ParseFiles("./public/html/home.html")
	if err != nil{
		log.Println(err)
	}
	err = t.Execute(w,nil)
	if err != nil{
		log.Println(err)
	}
}

type Noticia struct {
	Id int
	Titulo string
	Cuerpo string
	Fecha string
	Autor string
	Correo string
}

type ListaNoticia struct {
	Noticias []Noticia
}

func noticia(w http.ResponseWriter, r *http.Request)  {


	vars := mux.Vars(r)
	identificacion := vars["id"]

	db, err:= sql.Open("sqlite3", "./datos.db")
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("SELECT * FROM noticias where id=?")
	if err != nil {
		panic(err)
	}

	rows, err := stmt.Query(identificacion)
	if err != nil {
		panic(err)
	}

	var D Noticia

	for rows.Next() {
		err = rows.Scan(&D.Id, &D.Titulo, &D.Cuerpo, &D.Fecha, &D.Autor, &D.Correo)
		if err != nil {
			panic(err)
		}
	}
	defer rows.Close()
	if D.Id == 0 {
		t, err := template.ParseFiles("./public/html/noticiaError.html")
		if err != nil{
			log.Println(err)
		}
		err = t.Execute(w, identificacion)
		if err != nil{
			log.Println(err)
		}
	}else {
		t,err := template.ParseFiles("./public/html/noticias.html")
		if err != nil{
			log.Println(err)
		}
		err = t.Execute(w, D)
		if err != nil{
			log.Println(err)
		}
	}


}

func nueva(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, " Creando una nueva noticia")
}

func lista(w http.ResponseWriter, r *http.Request)  {

	db, err:= sql.Open("sqlite3", "./datos.db")
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM noticias")
	if err != nil{
		log.Println(err)
	}

	var L ListaNoticia
	var (
		Id int
		Titulo string
		Cuerpo string
		Fecha string
		Autor string
		Correo string
	)

	for rows.Next() {

		err = rows.Scan(&Id, &Titulo, &Cuerpo, &Fecha, &Autor, &Correo)
		if err != nil {
			panic(err)
		}
		L.Noticias = append(L.Noticias, Noticia{
			Id: Id,
			Titulo: Titulo,
			Cuerpo: Cuerpo,
			Fecha: Fecha,
			Autor: Autor,
			Correo: Correo,
		})
	}
	rows.Close()

	t, err := template.ParseFiles("./public/html/listaNoticia.html")
	if err != nil{
		log.Println(err)
	}
	t.Execute(w, L)
	if err != nil{
		log.Println(err)
	}

}
func parrafo(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("./public/html/parrafo.html")
	if err != nil{
		log.Println(err)
	}
	t.Execute(w, nil)
	if err != nil{
		log.Println(err)
	}
}

func golang(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("./public/html/go.html")
	if err != nil{
		log.Println(err)
	}
	t.Execute(w, nil)
	if err != nil{
		log.Println(err)
	}
}

func medioambiete(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("./public/html/medio_ambiente.html")
	if err != nil{
		log.Println(err)
	}
	t.Execute(w, nil)
	if err != nil{
		log.Println(err)
	}
}
func introduccion(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("./public/html/introduccion.html")
	if err != nil{
		log.Println(err)
	}
	t.Execute(w, nil)
	if err != nil{
		log.Println(err)
	}
}
func objetivos(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("./public/html/objetivos.html")
	if err != nil{
		log.Println(err)
	}
	t.Execute(w, nil)
	if err != nil{
		log.Println(err)
	}
}
func metodologia(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("./public/html/metodologia.html")
	if err != nil{
		log.Println(err)
	}
	t.Execute(w, nil)
	if err != nil{
		log.Println(err)
	}
}
func resultados(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("./public/html/resultados.html")
	if err != nil{
		log.Println(err)
	}
	t.Execute(w, nil)
	if err != nil{
		log.Println(err)
	}
}
func presupuesto(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("./public/html/presupuesto.html")
	if err != nil{
		log.Println(err)
	}
	t.Execute(w, nil)
	if err != nil{
		log.Println(err)
	}
}
func instituciones(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("./public/html/instituciones.html")
	if err != nil{
		log.Println(err)
	}
	t.Execute(w, nil)
	if err != nil{
		log.Println(err)
	}
}
func principal(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("./public/html/principal.html")
	if err != nil{
		log.Println(err)
	}
	t.Execute(w, nil)
	if err != nil{
		log.Println(err)
	}
}
func herramientas(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("./public/html/herramientas.html")
	if err != nil{
		log.Println(err)
	}
	t.Execute(w, nil)
	if err != nil{
		log.Println(err)
	}
}
func login(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("./public/html/login.html")
	if err != nil{
		log.Println(err)
	}
	t.Execute(w, nil)
	if err != nil{
		log.Println(err)
	}
}
func fecha()  {
	año, mes, dia:= time.Now().Date()
	fecha := fmt.Sprintf("%d/%d/%d",año, mes, dia)
	fmt.Println(fecha)
}
