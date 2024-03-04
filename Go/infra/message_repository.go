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
	ins, err := db.Prepare("INSERT INTO message_tb (name, message, time) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer ins.Close()

	// クエリ実行
	ins.Exec(message.Name, message.Message, message.Time)

	return nil
}

func (mr MessageRepository) DBList(db *sql.DB) ([]*model.Message, error) {
	rows, err := db.Query("SELECT * FROM message_tb")
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
