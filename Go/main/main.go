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
	"time"

	"database/sql"

	"github.com/aki2772/MessageBoard_sample/infra"
	"github.com/aki2772/MessageBoard_sample/model"
	"github.com/aki2772/MessageBoard_sample/repository"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"

	"net/http"

	"html/template"
	// ミドルウェア用パッケージ
	"github.com/labstack/echo/v4/middleware"
	// メインのフレームワークにechoを使用
	"github.com/labstack/echo/v4"
)

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

	//この時点でDBに接続
	db, err := ConnectDB()
	if err != nil {
		fmt.Println("データベースに接続できませんでした")
		// エラーが発生したら終了
		return
	}

	// メッセージリポジトリの作成
	messageRepository := infra.NewMySQLMessageRepository(db)

	// インスタンスを作成
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("C:/Users/aki/Documents/GitHub/MessageBoard_sample/Go/views/*.html")),
	}

	e.Renderer = t

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// コントローラーを作成
	mc := NewMessageController(messageRepository)

	// ルートを設定
	// http://localhost:1323/main にGETアクセスされるとviewMainPageハンドラーを実行する
	e.GET("/main", mc.ViewMainPage)

	e.GET("/newPage", mc.ViewNewPage)

	e.GET("/listPage", mc.ViewListPage)

	e.POST("/api/new", mc.NewApi)

	e.GET("/api/list", mc.ListApi)

	// サーバーをポート番号1323で起動
	err = e.Start(":1323")
	if err != nil {
		fmt.Println("サーバーの起動に失敗しました")
		return
	}
}

// データベース接続
func ConnectDB() (*sql.DB, error) {
	// タイムゾーンを設定
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

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
		return nil, err
	}

	// 接続確認
	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}
	fmt.Println("データベースに接続しました")
	fmt.Println()

	return db, nil
}

func New(mrStruct repository.Message, name string, message string) error {
	// メッセージを作成
	msg := model.Message{
		Name:    name,
		Message: message,
		Time:    time.Now().In(time.FixedZone("Asia/Tokyo", 9*60*60)),
	}

	// メッセージを保存
	err := mrStruct.Save(&msg)
	if err != nil {
		fmt.Println("メッセージの保存に失敗しました。")
		return err
	}

	return nil
}

func List(mrStruct repository.Message) []*model.Message {
	// メッセージのリストを取得
	msgList, err := mrStruct.List()
	// 失敗したら終了
	if err != nil {
		fmt.Println("メッセージの取得に失敗しました。")
		return nil
	}

	return msgList
}

// コンストラクタ
func NewMessageController(mr repository.Message) *MessageController {
	return &MessageController{mr: mr}
}

type MessageController struct {
	mr repository.Message
}

// メインページ表示ハンドラー
func (mc *MessageController) ViewMainPage(c echo.Context) error {
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
func (mc *MessageController) ViewNewPage(c echo.Context) error {
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

	// Renderでhtmlを表示
	return c.Render(http.StatusOK, "newPage", data)
}

// 一覧ページ表示ハンドラー
func (mc *MessageController) ViewListPage(c echo.Context) error {
	// msgList := List(mc.mr)

	// テンプレートに渡す値をセット
	var common = CommonData{
		"メッセージリスト表示",
	}
	data := struct {
		// field名は大文字で始める
		CommonData
		MsgList []*model.Message
	}{
		CommonData: common,
		// MsgList:    msgList,
	}
	// Renderでhtmlを表示
	return c.Render(http.StatusOK, "listPage", data)
}

// 新規作成APIハンドラー
func (mc MessageController) NewApi(c echo.Context) error {
	// フォームから送信されたテキストデータを取得(htmlからのfetch)
	m := new(model.Message)
	if err := c.Bind(m); err != nil {
		return err
	}

	name := m.Name
	message := m.Message

	data := struct {
		Name    string `json:"name"`
		Message string `json:"message"`
	}{
		Name:    name,
		Message: message,
	}

	// メッセージを保存
	New(mc.mr, name, message)
	// JSONを返す
	return c.JSON(http.StatusOK, data)
}

// 一覧データ要求APIハンドラー
func (mc MessageController) ListApi(c echo.Context) error {
	// メッセージのリストを取得
	msgList := List(mc.mr)

	// JSONを返す
	return c.JSON(http.StatusOK, msgList)
}
