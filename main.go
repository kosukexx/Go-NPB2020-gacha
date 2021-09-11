package main

import (
	"database/sql"//DB周り，もっと使いやすいフレームワークがあるらしいが今回は標準パッケージで
	"github.com/gin-gonic/gin"//webフレームワーク
	_ "github.com/go-sql-driver/mysql"//mysql使うとき
	"math/rand"//乱数生成時
	"net/http"//標準パッケージ
	"strconv"//strをintにするときに使う
	"time"//乱数生成時
)

type Player struct {
	ID   int
	Name string
	Rank string
	average float32
	HR float32
	OPS float32
	SB float32
}
//ユーザ情報を使うとき用，まだ実装できてない
type User struct{
	ID int
	Name string
	Jewel int
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
	defer db.Close()//どこで書いても関数の最後に実行される．基本db:=hogeの後ろ

	//DBからデータ取得，本当は都度必要なところだけ取り出したいけど今回は全部取り出す
	rows, err := db.Query("SELECT * FROM players")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()//どこで書いても関数の最後に実行される．基本db:=hogeの後ろ

	players :=[]Player{}
	for rows.Next() {
		var player Player
		err := rows.Scan(&player.ID, &player.Name, &player.Rank, &player.average,&player.HR,&player.OPS,&player.SB)
		if err != nil {
			panic(err.Error())
		}
		players = append (players, player)
		//fmt.Println(&player.ID, &player.Name, &player.Rank, &player.average,&player.HR,&player.OPS,&player.SB)

	}
	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/index.html")

	r.GET("/", func(c *gin.Context){
		c.HTML(http.StatusOK, "index.html", gin.H{"title":"プロ野球有名選手ガチャ","message":"何回引く？"})
	})

	r.POST("/",func(c *gin.Context){
		number,_:=strconv.Atoi(c.PostForm("field"))//フォームに入力された数字はstringとして読み込まれてるから，intにパースしてあげる
		var postJson bool//JSONで問い合わせされたときはJSONで返す，フォーム入力で問い合わせされたらHTMLで返すための変数
		if number==0{//以下webページからの入力でなく，JSONでの問い合わせのとき(API本来の機能)，
			var receiveNumber receiveNumber
			c.BindJSON(&receiveNumber)
			number,_=strconv.Atoi(receiveNumber.Number)
			postJson=true
		}

		reality:=[]string{}
		result :=[]Player{}
		var getNumber []int
		for i:=0;i<number;i++{
			reality=append(reality,getReality())//入力された数字の回数だけレアリティををreality配列に入れる
			switch{//あるレアリティの時にどのidの選手を排出するか，ここは汎用性がないので改良したい
			case reality[i]=="SR":
				getNumber=append(getNumber,rand.Intn(12))
			case reality[i]=="R":
				getNumber=append(getNumber,12+rand.Intn(21))
			default://reality[i]=="N":
				getNumber=append(getNumber,33+rand.Intn(19))
			}
			result=append(result,players[getNumber[i]])
		}



		if postJson==true{
			c.JSON(200,result)//APIとしてならこれ，構造体とかスライスのまま返せる
		}else{
			c.HTML(http.StatusOK, "index.html", gin.H{"title":"プロ野球有名選手ガチャ","message":result})//HTML+JSONで返すとき
		}

		//data,_:=json.Marshal(players)//c.JSONで返すなら必要ないが，JSON形式にしてなんか処理するときは必要


	})

	r.Run(":8080")

}


//スライスを引数で渡したときの挙動https://christina04.hatenablog.com/entry/2017/09/26/190000
/*
func resultNumbers(number int,players []Player,reality []string) []int{
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