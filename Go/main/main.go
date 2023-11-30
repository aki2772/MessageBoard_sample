/*
	Created by Kobayashi Atsuki.
	メインパッケージ。ユーザーの入力を受け取り、メッセージを保管する。
*/

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aki2772/MessageBoard_sample/Go/model"
)

func main() {

	// コマンドライン引数が2つでなければ終了
	if len(os.Args) != 2 {
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
	fmt.Println("list")
}

func New() {
	var name, message string
	fmt.Print("Enter your name... >")
	fmt.Scan(&name)
	fmt.Print("Enter message... >")
	fmt.Scan(&message)

	msgStruct := model.Message{
		Name:    name,
		Message: message,
		Time:    time.Now(),
	}

	fmt.Println(msgStruct)
}
