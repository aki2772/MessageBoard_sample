
/*
	Created by Kobayashi Atsuki.
	メインパッケージ。ユーザーの入力を受け取り、メッセージを保管する。
*/

package main

import "fmt"

func main() {
	var str string
	fmt.Print("Enter your name >")
	fmt.Scan(&str)
	fmt.Println("Hello " + str + ".")
}
