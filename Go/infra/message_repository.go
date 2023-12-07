/*
	Created by Kobayashi Atsuki.
	infraパッケージ。メッセージの永続化を行う。
*/

package infra

import (
	"bufio"
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
func (mr MessageRepository) Save(message []string) error {
	// 書き込み先のファイルを開く
	f, err := os.OpenFile(mr.FilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	// 存在しないならエラー
	if err != nil {
		log.Fatal(err)
		return err
	}
	// 関数終了時にファイルを閉じる
	defer f.Close()

	// 名前・メッセージ・時間・区切り文字を連結し、それぞれを改行で区切る
	// リスト取得時に行ごとにデータを復元し、区切り文字で1データと識別するため。
	d := []byte(message[0] + "\n" + message[1] + "\n" + message[2] + "\n" + "***" + "\n")

	// 書き込み
	n, err := f.Write(d)
	// 失敗したらエラー
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
func (mr MessageRepository) List() ([]string, error) {
	// ファイルを開く
	f, err := os.Open(mr.FilePath)
	// 存在しないならエラー
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// 関数終了時にファイルを閉じる
	defer f.Close()

	// メッセージのリスト
	messages := []string{}
	// ファイルを1行ずつ読み込む
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// リストに追加
		messages = append(messages, scanner.Text())
	}

	// ファイルの終端まで読み込んだら終了
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return messages, nil
}
