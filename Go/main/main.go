/*
	Created by Kobayashi Atsuki.
	メインパッケージ。ユーザーの入力を受け取り、メッセージを保管する。
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/aki2772/MessageBoard_sample/Go/infra"      // 独自パッケージ
	"github.com/aki2772/MessageBoard_sample/Go/model"      // 独自パッケージ
	"github.com/aki2772/MessageBoard_sample/Go/repository" // 独自パッケージ
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

	/*// メッセージ文字列を結合して作成(時刻データは文字列に変換)
	msgStr := msgStruct.Name + msgStruct.Message + strconv.Itoa(year) + "." + strconv.Itoa(int(month)) + "." + strconv.Itoa(day) +
		" " + strconv.Itoa(hour) + ":" + strconv.Itoa(min)*/

	// メッセージを永続化
	err := mrStruct.Save(msgStruct)
	// 失敗したら終了
	if err != nil {
		fmt.Println("メッセージの永続化に失敗しました。")
	}
}
