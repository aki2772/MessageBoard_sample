/*
	Created by Kobayashi Atsuki.
	infraパッケージ。メッセージのDBへの保存を行う。
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

// インターフェースを満たしているかの確認
var _ repository.Message = new(MySQLMessageRepository)

// コンストラクタ
// 構造体のフィールドがprivateの場合、コンストラクタで仲介し外部からアクセスする
func NewMySQLMessageRepository(db *sql.DB) *MySQLMessageRepository {
	return &MySQLMessageRepository{db: db}
}

type MySQLMessageRepository struct {
	db *sql.DB
}

func (mr MySQLMessageRepository) Save(message *model.Message) error {
	// テーブル作成
	ins, err := mr.db.Prepare("INSERT INTO message_tb (name, message, time) VALUES(?, ?, ?)")
	if err != nil {
		log.Println(err)
		return err
	}
	defer ins.Close()

	// クエリ実行
	ins.Exec(message.Name, message.Message, message.Time)

	return nil
}

func (mr MySQLMessageRepository) List() ([]*model.Message, error) {
	rows, err := mr.db.Query("SELECT * FROM message_tb")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var ret []*model.Message
	for rows.Next() {
		// メッセージを取得
		var msg model.Message
		// カラムをスキャン
		err := rows.Scan(&msg.Name, &msg.Message, &msg.Time)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		ret = append(ret, &msg)
	}
	return ret, nil
}
