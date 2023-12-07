/*
	Created by Kobayashi Atsuki.
	メインパッケージ。ユーザーの入力を受け取り、メッセージを保管する。
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aki2772/MessageBoard_sample/Go/infra" // 独自パッケージ
	"github.com/aki2772/MessageBoard_sample/Go/model" // 独自パッケージ
)

const filePath = "C:/Users/aki/Documents/GitHub/MessageBoard_sample/messages.txt" // ファイルパス

func main() {
	// コマンドライン引数が2つでなければ終了
	if len(os.Args) != 2 {
		fmt.Println("コマンドライン引数が不正です。")
		return
	}

	cmdLine := os.Args
	if cmdLine[1] == "list" { // list
		List()
	} else if cmdLine[1] == "new" { // new
		New()
	} else {
		fmt.Println("コマンドライン引数が不正です。")
		return
	}
}

func List() {
	// 永続化関数を持つ構造体を生成
	mrStruct := infra.MessageRepository{
		FilePath: filePath, // string
	}

	// メッセージのリストを取得
	msgList, err := mrStruct.List()
	// 失敗したら終了
	if err != nil {
		fmt.Println("メッセージの取得に失敗しました。")
		return
	}
	fmt.Print(msgList)
}

// / <summary>
// / 新しくメッセージを生成する。
// / </summary>
func New() {
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

	// 永続化関数を持つ構造体を生成
	mrStruct := infra.MessageRepository{
		FilePath: filePath, // string
	}

	// 時刻データを取得
	year, month, day := msgStruct.Time.Date()
	hour, min, _ := msgStruct.Time.Clock()

	// メッセージ文字列を結合して作成(時刻データは文字列に変換)
	msgComp := []string{msgStruct.Name, msgStruct.Message, strconv.Itoa(year) + "." + strconv.Itoa(int(month)) + "." + strconv.Itoa(day) +
		" " + strconv.Itoa(hour) + ":" + strconv.Itoa(min)}

	// メッセージを永続化
	err := mrStruct.Save(msgComp)
	// 失敗したら終了
	if err != nil {
		fmt.Println("メッセージの永続化に失敗しました。")
		return
	}
}
