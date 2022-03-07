package main

import (
	"database/sql"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var DB *sql.DB

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	database, err := sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+dbname)
	if err != nil {
		panic(err)
	}
	DB = database
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {

		var user = User{}
		var users []User

		rows, err := DB.Query("select id, name from users")
		if err != nil {
			return Error(c, err)
		}

		for rows.Next() {
			err := rows.Scan(&user.Id, &user.Name)
			if err != nil {
				return Error(c, err)
			}
			users = append(users, user)
		}

		return c.JSON(http.StatusOK, users)
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func Error(c echo.Context, e error) error {
	c.Logger().Error(e)
	return e
}
