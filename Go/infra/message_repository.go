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

	"github.com/aki2772/MessageBoard_sample/model"
	"github.com/aki2772/MessageBoard_sample/repository"
	_ "github.com/go-sql-driver/mysql"
)

// 時刻のフォーマット
// var layout = "2006.01.02 15:04:05"

// インターフェースを満たしているかの確認
var _ repository.MessageRepository = new(MessageRepository)

// コンストラクタ
func NewMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{db: db, filePath: "C:/Users/aki/Documents/GitHub/MessageBoard_sample/Go/message.txt"}
}

type MessageRepository struct {
	// repository.MessageRepository // インターフェースの埋め込み
	db       *sql.DB
	filePath string
}

func (mr MessageRepository) DBSave(message *model.Message) error {
	// テーブル作成
	ins, err := mr.db.Prepare("INSERT INTO message_tb (name, message, time) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer ins.Close()

	// クエリ実行
	ins.Exec(message.Name, message.Message, message.Time)

	return nil
}

func (mr MessageRepository) DBList() ([]*model.Message, error) {
	rows, err := mr.db.Query("SELECT * FROM message_tb")
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
