/*
	Created by Kobayashi Atsuki.
	infraパッケージ。メッセージの永続化を行う。
*/

package infra

import (
	"fmt"
	"log"
	"os"

	"github.com/aki2772/MessageBoard_sample/Go/repository" // 独自パッケージ
)

type MessageRepository struct {
	FilePath                     string // ファイルパス
	repository.MessageRepository        // インターフェースの埋め込み
}

// / <summary>
// / メッセージを保管する。
// / </summary>
// / <param name="message">メッセージ</param>
// / <returns>エラー</returns>
func (mr MessageRepository) Save(message string) error {
	// 書き込み先のファイルを開く
	f, err := os.OpenFile(mr.FilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	// 存在しないならエラー
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer f.Close()

	// メッセージをバイト文字列に変換(改行込み)
	d := []byte(message + "\n")

	// 書き込み
	n, err := f.Write(d)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("%d bytesの書き込み", n)
	return nil
}

// / <summary>
// / 保管されているメッセージのリストを引き出す。
// / </summary>
// / <returns>メッセージのリスト</returns>
func (mr MessageRepository) List() (string, error) {
	// ファイルからメッセージを読み込む
	fmt.Println("List")
	return "", nil
}
