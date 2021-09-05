package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")


	r.GET("/", func(c *gin.Context){
		c.HTML(http.StatusOK, "index.html", gin.H{"title":"プロ野球有名選手ガチャ","message":"何回引く？"})
	})
	r.POST("",func(c *gin.Context){
		number:=c.PostForm("field")
		c.HTML(http.StatusOK, "index.html", gin.H{"title":"プロ野球有名選手ガチャ","message":result(number)})

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
func result(number string) string{
	//numberInt,_ :=strconv.Atoi(number)
	message:=number


	return message;
}
/*
func handler(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

 */
