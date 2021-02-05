package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	username = "root"
	password = "admin123"
	portnum = 3306
	dbname = "8chan"
)

func main() {

	myRouter := mux.NewRouter().StrictSlash(true)

	type Post struct {
		Title string
	}

	type PostPageInfo struct {
		PageTitle string
		Posts []Post
	}

	config := username + ":" + password + "@(localhost:" + strconv.Itoa(portnum) + ")/" + dbname + "?parseTime=true"
	db, err := sql.Open("mysql", config)
	// db, err := sql.Open("mysql", "root:admin123@tcp(127.0.0.1:3306)/testdb")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("everything looks fine")
	}

	homePage := template.Must(template.ParseFiles("./templates/index.html"))

	myRouter.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {

		// insert, err := db.Query("INSERT INTO posts VALUES ('FIRST POST')")
		// if err != nil {
		// 	panic(err.Error())
		// }
		// defer insert.Close()
		// fmt.Print("Successfully inserted into post tables")

		posts, err := db.Query("SELECT title FROM posts")
		if err!=nil {
			fmt.Print(err)
		} else {
			fmt.Println("no error")
		}
		var post []Post

		for posts.Next(){
			var p Post
			err := posts.Scan(&p.Title)
			if err == nil {
				fmt.Print(p.Title)
				post = append(post, p)
			} else {
				panic(err.Error())
			}
		}
		fmt.Println(post)
		info := PostPageInfo {
			PageTitle: "Tabooooo",
			Posts: post,
		}
		fmt.Println(info)
		er := homePage.Execute(w, info)
		if er!=nil {
			fmt.Print(er)
		}
		defer posts.Close()
	})

	myRouter.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request){
			title := r.FormValue("title")
			// insert, err := db.Query("INSERT INTO posts VALUES (title)", title)
			_, err = db.Exec("insert into posts(title) values(?)", title)
			if err != nil {
				panic(err.Error())
			} else {
				fmt.Println("Post submitted Successfully")
				http.Redirect(w, r, "/posts", http.StatusSeeOther)
			}
	})

	log.Fatal(http.ListenAndServe(":8081", myRouter))


}
