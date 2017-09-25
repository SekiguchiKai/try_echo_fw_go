package main

import (
	"github.com/labstack/echo"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
	"google.golang.org/appengine"
)

// 各エンドポイント毎にinitを作成することができる
func init() {
	log.Print("a")
	// エンドポイントの設定
	// endpoint groupとして設定
	g := E.Group("/user")

	empty := ""
	g.POST(empty, createUser)
	g.GET("/:id", getUser)
	g.GET("", getUserA)

}

type UserName struct {
	Name string `json:"name"`
}

// ユーザーを新規生成
// echo.Contextは、現在のHTTP requestのcontextを表す。
// echo.Contextは、request objects、response objects、 path、 path parameters、 data、registered handlerを持っている。
func createUser(c echo.Context) error {
	log.Println("aa")

	ac := appengine.NewContext(c.Request())

	// request bodyからJSONを読み込む
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		ErrorLog(ac, err.Error())
		return err
	}
	defer c.Request().Body.Close()

	DebugLog(ac ,string(body))

	un := &UserName{}
	// JSONをUnmarshal
	if err := json.Unmarshal(body, un); err != nil {
		ErrorLog(ac, err.Error())
		return err
	}

	// Userのインスタンスを生成
	u := NewUser(un.Name)
	// DatastoreにUserを登録
	if _, err := u.Post(c); err != nil {
		ErrorLog(ac, err.Error())
		return err
	}
	DebugLog(ac ,u.Name)

	// Bindは、リクエストボディを与えられた引数と結びつける。
	//if err := c.Bind(u); err != nil {
	//	ErrorLog(ac, err.Error())
	//	return err
	//}
	// 引数で与えられたstatus codeと共に、構造体をJSONにして返す
	return c.JSON(http.StatusCreated, u)
}

// ユーザーを取得
func getUser(c echo.Context) error{
	ac := appengine.NewContext(c.Request())
	// DatastoreからUserをGetする際に、データを格納するためのインスタンスを生成
	u := new(User)
	// クエリパラメータからidを取得し、UserのIDに格納
	u.ID = c.Param("id")
	DebugLog(ac, u.ID)

	// DatastoreからUserのデータを取得
	u, err := u.Get(c)
	if err != nil {
		ErrorLog(c.Request().Context(), err.Error())
		return err
	}
	// 引数で与えられたstatus codeと共に、構造体をJSONにして返す
	return c.JSON(http.StatusOK, u)
}

func getUserA(c echo.Context) error{
	log.Print("a")

	return nil
}