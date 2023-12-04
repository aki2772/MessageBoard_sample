/*
	Created by Kobayashi Atsuki.
	infraパッケージ。メッセージの永続化を行う。
*/

package infra

import (
	"fmt"

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
func (mr MessageRepository) Save(message string) string {
	// ファイルにメッセージを書き込む
	fmt.Println("Save")
	fmt.Println("message: " + message)
	return ""
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
