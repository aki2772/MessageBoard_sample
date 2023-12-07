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

	// 名前・メッセージ・時間を連結し、それぞれを改行で区切る。
	// リスト取得時に行ごとにデータを復元し、最後の改行で1行開け、1メッセージと識別する。
	d := []byte(message[0] + "\n" + message[1] + "\n" + message[2] + "\n" + "\n")

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

	// 読み込むデータのリスト
	texts := []string{}
	// 名前のリスト
	names := []string{}
	// メッセージのリスト
	messages := []string{}
	// 時間のリスト
	times := []string{}
	// ファイルを1行ずつ読み込む
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// リストに追加
		texts = append(texts, scanner.Text())
	}

	// リストを3つに分ける
	for i, text := range texts {
		if i%4 == 0 {
			names = append(names, text)
		} else if i%4 == 1 {
			messages = append(messages, text)
		} else if i%4 == 2 {
			times = append(times, text)
		} else {
			continue
		}
	}

	// リストを連結
	ret := []string{}
	// 1メッセージ4行なので、4で割る
	for i := 0; i < len(texts)/4; i++ {
		// 名前・メッセージ・時間を連結
		ret = append(ret, names[i]+": "+messages[i]+" ("+times[i]+")")
	}

	// ファイルの終端まで読み込んだら終了
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return ret, nil
}
