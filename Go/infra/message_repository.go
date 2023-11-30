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

func (mr MessageRepository) Save(msg string) {
	fmt.Print(mr.FilePath)
	fmt.Println("メッセージ保存")
}
