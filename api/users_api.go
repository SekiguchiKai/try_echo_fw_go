package api

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/SekiguchiKai/try_echo_fw_go/app"

	"github.com/SekiguchiKai/try_echo_fw_go/model"
)

// 各エンドポイント毎にinitを作成することができる
func init() {
	// エンドポイントの設定
	// endpoint groupとして設定
	g := app.E.Group("/users")

	empty := ""
	g.POST(empty, createUser)

}

// ユーザーを新規生成
// echo.Contextは、現在のHTTP requestのcontextを表す。
// echo.Contextは、request objects、response objects、 path、 path parameters、 data、registered handlerを持っている。
func createUser(c echo.Context) error {
	u := model.NewUser()




	// Bindは、リクエストボディを与えられた引数と結びつける。
	if err := c.Bind(u); err != nil {
		return err
	}
	// 引数で与えられたstatus codeと共に、構造体をJSONにして返す
	return c.JSON(http.StatusCreated, u)
}

// ユーザーを取得
func getUser(c echo.Context) {




}
