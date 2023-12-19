/*
	Created by Kobayashi Atsuki.
	infraパッケージ。メッセージの永続化を行う。
*/

package infra

import (
	/*"bufio"
	"fmt"
	"os"
	"time"*/

	"database/sql"
	"log"

	"github.com/aki2772/MessageBoard_sample/Go/model"      // 独自パッケージ
	"github.com/aki2772/MessageBoard_sample/Go/repository" // 独自パッケージ

	_ "github.com/go-sql-driver/mysql"
)

// 時刻のフォーマット
var layout = "2006.01.02 15:04:05"

type MessageRepository struct {
	FilePath                     string // ファイルパス
	repository.MessageRepository        // インターフェースの埋め込み
}

func (mr MessageRepository) DBSave(message *model.Message, db *sql.DB) error {
	// テーブル作成
	ins, err := db.Prepare("INSERT INTO messages (name, message, time) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer ins.Close()

	// クエリ実行
	ins.Exec(`aki`, `Hello, World!`, `2006.01.02 15:04:05`)

	return nil
}

func (mr MessageRepository) DBList(db *sql.DB) ([]*model.Message, error) {
	rows, err := db.Query("SELECT * FROM messages")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// アウトプットメッセージ
	var ret []*model.Message
	for rows.Next() {
		model := model.Message{}
		err := rows.Scan(&model.Name, &model.Message, &model.Time)
		if err != nil {
			log.Fatal(err)
		}
		ret = append(ret, &model)
	}
	return ret, nil
}

/*
// / <summary>
// / メッセージを保管する。
// / </summary>
// / <param name="message">メッセージ</param>
// / <returns>エラー</returns>
func (mr MessageRepository) Save(message *model.Message) error {
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
	d := []byte(message.Name + "\n" + message.Message + "\n" + message.Time.Format(layout) + "\n" + "\n")

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
func (mr MessageRepository) List() ([]*model.Message, error) {
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

	// アウトプットメッセージ
	ret := []*model.Message{}

	// 1メッセージ4行なので、4で割る
	for i := 0; i < len(texts)/4; i++ {
		time, _ := time.Parse(layout, times[i])
		// メッセージを生成
		ret = append(ret, &model.Message{
			Name:    names[i],
			Message: messages[i],
			Time:    time,
		})
	}

	// ファイルの終端まで読み込んだら終了
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return ret, nil
}
*/
