package app

import (
	"database/sql"

	"github.com/bmorgan5/beboptechnology/app/models"
	"github.com/revel/revel"
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	// register startup functions with OnAppStart
	// ( order dependent )
	revel.OnAppStart(InitBookDB)
	// revel.OnAppStart(FillCache)
}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

// InitBookDB initializes the BooksDB handle
func InitBookDB() {
	revel.INFO.Printf("Attempting to InitBookDB")
	var dbLoc = "db/books.db"
	var err error
	models.BooksDB, err = sql.Open("sqlite3", dbLoc)
	if err != nil {
		revel.ERROR.Printf("Failed to open sqlite3 database: %s", err)
	} else {
		if perr := models.BooksDB.Ping(); perr != nil {
			revel.ERROR.Printf("Failed to ping db: %s", perr)
		} else {
			if rows, err := models.BooksDB.Query("SELECT title, author FROM books"); err != nil {
				revel.ERROR.Printf("Failed to run query: %s", err)
			} else {
				defer rows.Close()
				for rows.Next() {
					var title, author string
					if err := rows.Scan(&title, &author); err != nil {
						revel.ERROR.Printf("Failed to scan row: %s", err)
					} else {
						revel.INFO.Printf("%s | %s", title, author)
					}
				}
			}
		}
	}

	//defer sqliteDB.Close()
}
