package controllers

import(
	"github.com/revel/revel"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type BookController struct {
	*revel.Controller
}

type Book struct {
	Title string
}

func InitBookDB() {
	revel.INFO.Printf("Attempting to InitBookDB")
	var dbLoc = "db/books.db"
	sqliteDB, err := sql.Open("sqlite3", dbLoc)
	if err != nil {
		revel.ERROR.Printf("Failed to open sqlite3 database: %s", err)
	}
	/*} else {
		if perr := sqliteDB.Ping(); perr != nil {
			revel.ERROR.Printf("Failed to ping db: %s", perr)
		}
	}
	*/
	defer sqliteDB.Close()
}

func (c BookController) HandleBookUpload(title string) revel.Result {
	revel.INFO.Printf("Uploading Book titled: %s", title)
	return c.RenderText(title)
}
