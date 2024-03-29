/*
	Created by Kobayashi Atsuki.
	infraパッケージ。メッセージのローカルテキストファイルへの保存を行う。
*/

package infra

import (
	/*"bufio"
	"fmt"
	"os"
	"time"*/

	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aki2772/MessageBoard_sample/model"
	"github.com/aki2772/MessageBoard_sample/repository"
	_ "github.com/go-sql-driver/mysql"
)

// 時刻のフォーマット
var layout = "2006.01.02 15:04:05"

// インターフェースを満たしているかの確認
var _ repository.Message = new(FileMessageRepository)

// コンストラクタ
// 構造体のフィールドがprivateの場合、コンストラクタで仲介し外部からアクセスする
func NewFileMessageRepository() *FileMessageRepository {
	return &FileMessageRepository{filePath: "C:/Users/aki/Documents/GitHub/MessageBoard_sample2/Go/message.txt"}
}

type FileMessageRepository struct {
	filePath string
}

// / <summary>
// / メッセージを保管する。
// / </summary>
// / <param name="message">メッセージ</param>
// / <returns>エラー</returns>
func (mr FileMessageRepository) Save(message *model.Message) error {
	// 書き込み先のファイルを開く
	f, err := os.OpenFile(mr.filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	// 存在しないならエラー
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		return err
	}

	fmt.Printf("%d bytesの書き込み", n)
	return nil
}

// / <summary>
// / 保管されているメッセージのリストを引き出す。
// / </summary>
// / <returns>メッセージのリスト</returns>
func (mr FileMessageRepository) List() ([]*model.Message, error) {
	// ファイルを開く
	f, err := os.Open(mr.filePath)
	// 存在しないならエラー
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		return nil, err
	}

	return ret, nil
}
