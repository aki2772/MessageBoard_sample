/*
	Created by Kobayashi Atsuki.
	メインパッケージ。ユーザーの入力を受け取り、メッセージを保管する。
*/

package main

import (
	"net/http"

	/* テンプレート用パッケージ */

	/* ミドルウェア用パッケージ */
	"github.com/labstack/echo/v4/middleware"
	/* メインのフレームワークにechoを使用 */
	"github.com/labstack/echo/v4"
)

func main() {
	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートを設定
	e.GET("/test", thanks) // ローカル環境の場合、http://localhost:1323/ にGETアクセスされるとthanksハンドラーを実行する
	e.GET("/", hello)      // ローカル環境の場合、http://localhost:1323/ にGETアクセスされるとhelloハンドラーを実行する

	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":1323"))
}

// ハンドラーを定義
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// ハンドラーを定義
func thanks(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Programing!")
}
