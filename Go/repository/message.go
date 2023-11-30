/*
	Created by Kobayashi Atsuki.
	repositoryパッケージ。メッセージの保管を行う。
*/

package repository

type MessageRepository interface {
	Save(string) // メッセージを保管する
	List() // 保管されているメッセージのリストを引き出す
}
