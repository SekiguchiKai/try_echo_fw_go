// +build !appengine

package app

import (
	"github.com/labstack/echo"
	"net/http"
)

// echoを設定して返す
// GAEのStandard環境用の設定
func createMux() *echo.Echo {
	// Echoのインスタンスを生成
	e := echo.New()

	http.Handle("/", e)
	return e
}
