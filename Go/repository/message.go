/*
	Created by Kobayashi Atsuki.
	repositoryパッケージ。メッセージの保管を行う。
*/

package repository

import (
	"github.com/aki2772/MessageBoard_sample/Go/model" // 独自パッケージ

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MessageRepository interface {
	/*Save(*model.Message) error       // メッセージを保管する
	List() ([]*model.Message, error) // 保管されているメッセージのリストを引き出す*/

	DBSave(*model.Message, *sql.DB) error     // メッセージを保管する
	DBList(*sql.DB) ([]*model.Message, error) // 保管されているメッセージのリストを引き出す
}
