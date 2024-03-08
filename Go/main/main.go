/*
	Created by Kobayashi Atsuki.
	メインパッケージ。ユーザーの入力を受け取り、メッセージを保管する。

	echoを用いたWebアプリケーション。
	https://qiita.com/kubota_ndatacom/items/a45d89ab63530c640450
*/

package main

import (
	"fmt"
	"io"
	"log"
	"main/infra"
	"time"

	"database/sql"

	"github.com/aki2772/MessageBoard_sample/Go/model"      // 独自パッケージ
	"github.com/aki2772/MessageBoard_sample/Go/repository" // 独自パッケージ

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"

	"net/http"

	// テンプレート用パッケージ
	"html/template"
	// ミドルウェア用パッケージ
	"github.com/labstack/echo/v4/middleware"
	// メインのフレームワークにechoを使用
	"github.com/labstack/echo/v4"
)

// 時刻のフォーマット
var layout = "2006.01.02 15:04:05"

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// 共通使用するデータ用の構造体
type CommonData struct {
	// field名は大文字で始める
	Title string
}

func main() {

	// インスタンスを作成
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("C:/Users/aki/Documents/GitHub/MessageBoard_sample/Go/views/*.html")),
	}

	e.Renderer = t

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートを設定
	// http://localhost:1323/main にGETアクセスされるとviewMainPageハンドラーを実行する
	e.GET("/main", viewMainPage)

	// POSTにすると{"message":"Method Not Allowed"}というレスポンスが返ってくるためGETにしている
	// なぜGETでPOSTの処理ができるのかは不明
	// POSTにするとviewNewPageハンドラーが実行されない
	e.GET("/newPage", viewNewPage)

	e.GET("/listPage", viewListPage)

	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":1323"))
}

// データベース接続
func ConnectDB() *sql.DB {
	// タイムゾーンを設定
	jst, _ := time.LoadLocation("Asia/Tokyo")

	// データベースの設定
	cfg := mysql.Config{
		DBName:    "message-db",
		User:      "aki",
		Passwd:    "fafnirclear",
		Addr:      "127.0.0.1:3006",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	// データベースに接続
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	// 接続確認
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("データベースに接続しました")
	fmt.Println()

	return db
}

func New(mrStruct repository.MessageRepository, db *sql.DB, name string, message string) {
	// メッセージを作成
	msg := model.Message{
		Name:    name,
		Message: message,
		Time:    time.Now(),
	}

	// メッセージを保存
	err := mrStruct.DBSave(&msg, db)
	// 失敗したら終了
	if err != nil {
		fmt.Println("メッセージの保存に失敗しました。")
		return
	}
}

func List(mrStruct repository.MessageRepository, db *sql.DB) []*model.Message {
	// メッセージのリストを取得
	msgList, err := mrStruct.DBList(db)
	// 失敗したら終了
	if err != nil {
		fmt.Println("メッセージの取得に失敗しました。")
		return nil
	}

	return msgList
}

// メインページ表示ハンドラー
func viewMainPage(c echo.Context) error {
	// テンプレートに渡す値をセット
	var common = CommonData{
		"ホーム",
	}
	data := struct {
		// field名は大文字で始める
		CommonData
	}{
		CommonData: common,
	}

	// Renderでhtmlを表示
	return c.Render(http.StatusOK, "mainPage", data)
}

// 新規作成ページ表示ハンドラー
func viewNewPage(c echo.Context) error {
	// テンプレートに渡す値をセット
	var common = CommonData{
		"新規メッセージ作成",
	}
	data := struct {
		// field名は大文字で始める
		CommonData
	}{
		CommonData: common,
	}

	// フォームから送信されたテキストデータを取得
	name := c.FormValue("name")
	message := c.FormValue("message")

	// テキストデータが空の場合はエラーを返す
	if name == "" || message == "" {
		fmt.Println("名前またはメッセージが入力されていません")
		return c.Render(http.StatusNotFound, "newPage", data)
	}

	// ここでテキストデータを使用して必要な処理を実行
	// データベースに接続
	mrStruct := infra.MessageRepository{}
	db := ConnectDB()

	// メッセージを保存
	New(mrStruct, db, name, message)
	return c.Render(http.StatusOK, "newPage", data)
}

// 一覧ページ表示ハンドラー
func viewListPage(c echo.Context) error {
	// データベースに接続
	mrStruct := infra.MessageRepository{}
	db := ConnectDB()

	var times []string

	msgList := List(mrStruct, db)
	for _, msg := range msgList {
		times = append(times, msg.Time.Format(layout))
	}

	// テンプレートに渡す値をセット
	var common = CommonData{
		"メッセージリスト表示",
	}
	data := struct {
		// field名は大文字で始める
		CommonData
		MsgList []*model.Message
		Times   []string
	}{
		CommonData: common,
		MsgList:    msgList,
		Times:      times,
	}
	// Renderでhtmlを表示
	return c.Render(http.StatusOK, "listPage", data)
}
