/*
	Created by Kobayashi Atsuki.
	repositoryパッケージ。メッセージの保管を行う。
*/

package repository

type MessageRepository interface {
	Save(string) error       // メッセージを保管する
	List() ([]string, error) // 保管されているメッセージのリストを引き出す
}
