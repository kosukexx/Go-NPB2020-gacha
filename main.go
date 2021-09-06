package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)
type User struct {
	ID   int
	Name string
	Rank string
	average float32
	HR float32
	OPS float32
	SB float32
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/sample")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM players")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	users :=[]User{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Rank, &user.average,&user.HR,&user.OPS,&user.SB)
		if err != nil {
			panic(err.Error())
		}
		users = append (users, user)
		//fmt.Println(user.ID, user.Name, user.Rank, user.average,user.HR,user.OPS,user.SB)

	}
	fmt.Println(users)
	fmt.Println(users[0])



	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")


	r.GET("/", func(c *gin.Context){
		c.HTML(http.StatusOK, "index.html", gin.H{"title":"プロ野球有名選手ガチャ","message":"何回引く？"})
	})
	r.POST("",func(c *gin.Context){
		number:=c.PostForm("field")
		c.HTML(http.StatusOK, "index.html", gin.H{"title":"プロ野球有名選手ガチャ","message":result(number,users)})
		c.JSON(200,gin.H{"name":users[1].Name})

	})

	/*
	r.GET("", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"name":"nakata",
		})
	})

	 */
	r.Run(":8080")

}
func result(number string,users []User) string{
	//numberInt,_ :=strconv.Atoi(number)
	//message:=number


	return (users[1].Name);
}
