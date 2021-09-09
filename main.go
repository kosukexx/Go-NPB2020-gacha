package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"net/http"
	"strconv"

	//"strconv"
	"time"
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

type receiveNumber struct {
	Number string `json:"number"`
}


func main() {
	rand.Seed(time.Now().Unix())//乱数生成

	//以下DB接続
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/sample")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//DBからデータ取得
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
	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", func(c *gin.Context){
		c.HTML(http.StatusOK, "index.html", gin.H{"title":"プロ野球有名選手ガチャ","message":"何回引く？"})
	})

	r.POST("/",func(c *gin.Context){
		//number:=c.PostForm("field")
		number,_:=strconv.Atoi(c.PostForm("field"))//フォームに入力された数字はstringとして読み込まれてるから，intにパースしてあげる
		if number==0{
			var hoge receiveNumber
			c.BindJSON(&hoge)
			number,_=strconv.Atoi(hoge.Number)
		}
		reality:=[]string{}
		var getNumber []int
		for i:=0;i<number;i++{
			reality=append(reality,getReality())//入力された数字の回数だけ// をreality配列に入れる
		}

		for i:=0;i<number;i++{//あるレアリティの時にどのidの選手を排出するか
			switch{
			case reality[i]=="SR":
				getNumber=append(getNumber,rand.Intn(12))
			case reality[i]=="R":
				getNumber=append(getNumber,12+rand.Intn(21))
			default:
				getNumber=append(getNumber,33+rand.Intn(19))
			}
		}
		result :=[]User{}
		for i:=0;i<number;i++{
			result=append(result,users[getNumber[i]])
		}



		//c.HTML(http.StatusOK, "index.html", gin.H{"title":"プロ野球有名選手ガチャ","message":reality[0]})
		//c.HTML(http.StatusOK, "index.html", gin.H{"title":"プロ野球有名選手ガチャ","message":result})

		//data,_:=json.Marshal(users)//c.JSONで返すなら必要ないが，JSON形式にしてなんか処理するときは必要
		c.JSON(200,result)//構造体とかスライスのまま返せる



	})



	r.Run(":8080")

}
//スライスを引数で渡したときの挙動https://christina04.hatenablog.com/entry/2017/09/26/190000
/*
func resultNumbers(number int,users []User,reality []string) []int{
	return
}



func srRankNumber(number int) int{
		num:=rand.Intn(12)
	return num
}
func rRankNumber(number int) []int{
	var num []int
	for i:=0;i<number;i++{
		num[i]=12+rand.Intn(21)
	}
	return num
}
func normalRankNumber(number int) []int{
	var num []int
	for i:=0;i<number;i++{
		num[i]=33+rand.Intn(19)
	}
	return num
}

func contains(arr []string, str string) bool{
	for _, v := range arr{
		if v == str{
			return true
		}
	}
	return false
}
*/

func getReality()string{
	num:=rand.Intn(100)
	switch{
	case num<80:
		return "N"
	case num<99:
		return "R"
	default:
		return "SR"
	}
}