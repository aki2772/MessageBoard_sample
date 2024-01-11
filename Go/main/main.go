/*
	Created by Kobayashi Atsuki.
	メインパッケージ。ユーザーの入力を受け取り、メッセージを保管する。
*/

package main

import (
	"bufio"
	"os"
	"time"

	"database/sql"
	"fmt"
	"log"

	"github.com/aki2772/MessageBoard_sample/Go/infra"      // 独自パッケージ
	"github.com/aki2772/MessageBoard_sample/Go/model"      // 独自パッケージ
	"github.com/aki2772/MessageBoard_sample/Go/repository" // 独自パッケージ

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

// ファイルパス
const filePath = "C:/Users/aki/Documents/GitHub/MessageBoard_sample/messages.txt"

// 時刻のフォーマット
var layout = "2006.01.02 15:04:05"

func main() {
	// コマンドライン引数が2つでなければ終了
	if len(os.Args) != 2 {
		fmt.Println("コマンドライン引数が不正です。")
		return
	}

	cmdLine := os.Args
	if cmdLine[1] != "list" && cmdLine[1] != "new" {
		fmt.Println("コマンドライン引数が不正です。")
		return
	}

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		// エラーハンドリング
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
		log.Fatal(err)
	}
	defer db.Close()

	// 接続確認
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("データベースに接続しました")
	fmt.Println()

	// 永続化関数を持つ構造体を生成
	mrStruct := infra.MessageRepository{
		FilePath: filePath, // string
	}

	if cmdLine[1] == "list" { // list
		List(mrStruct, db)
	} else { // new
		New(mrStruct, db)
	}
}

func New(mrStruct repository.MessageRepository, db *sql.DB) {
	// 標準入力のスキャナー
	nameScn := bufio.NewScanner(os.Stdin)
	messageScn := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter your name...")
	fmt.Print(">")
	// 名前入力
	nameScn.Scan()

	fmt.Println("Enter message...")
	fmt.Print(">")
	// メッセージ入力
	messageScn.Scan()

	// メッセージを生成
	msgStruct := &model.Message{
		Name:    nameScn.Text(),    // string
		Message: messageScn.Text(), // string
		Time:    time.Now(),        // time.Time
	}

	// メッセージを永続化
	err := mrStruct.DBSave(msgStruct, db)
	// 失敗したら終了
	if err != nil {
		fmt.Println("メッセージの永続化に失敗しました。")
	}
}

func List(mrStruct repository.MessageRepository, db *sql.DB) {
	// メッセージのリストを取得
	msgList, err := mrStruct.DBList(db)
	// 失敗したら終了
	if err != nil {
		fmt.Println("メッセージの取得に失敗しました。")
		return
	}

	for _, msg := range msgList {
		fmt.Println(msg.Name + ": " + msg.Message + " (" + msg.Time.Format(layout) + ")")
	}
}

// / <summary>
// / ローカルの.txtファイルにメッセージを保管する。
// / </summary>
/*
func main() {
	// コマンドライン引数が2つでなければ終了
	if len(os.Args) != 2 {
		fmt.Println("コマンドライン引数が不正です。")
		return
	}

	// 永続化関数を持つ構造体を生成
	mrStruct := infra.MessageRepository{
		FilePath: filePath, // string
	}

	cmdLine := os.Args
	if cmdLine[1] == "list" { // list
		List(mrStruct)
	} else if cmdLine[1] == "new" { // new
		New(mrStruct)
	} else {
		fmt.Println("コマンドライン引数が不正です。")
		return
	}
}

func List(mrStruct repository.MessageRepository) {
	// メッセージのリストを取得
	msgList, err := mrStruct.List()
	// 失敗したら終了
	if err != nil {
		fmt.Println("メッセージの取得に失敗しました。")
		return
	}

	for _, msg := range msgList {
		fmt.Println(msg.Name + ": " + msg.Message + " (" + msg.Time.Format(layout) + ")")
	}
}

// / <summary>
// / 新しくメッセージを生成する。
// / </summary>
func New(mrStruct repository.MessageRepository) {
	// 標準入力のスキャナー
	nameScn := bufio.NewScanner(os.Stdin)
	messageScn := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter your name...")
	fmt.Print(">")
	// 名前入力
	nameScn.Scan()

	fmt.Println("Enter message...")
	fmt.Print(">")
	// メッセージ入力
	messageScn.Scan()

	// メッセージを生成
	msgStruct := &model.Message{
		Name:    nameScn.Text(),    // string
		Message: messageScn.Text(), // string
		Time:    time.Now(),        // time.Time
	}

	// メッセージを永続化
	err := mrStruct.Save(msgStruct)
	// 失敗したら終了
	if err != nil {
		fmt.Println("メッセージの永続化に失敗しました。")
	}
}
*/
