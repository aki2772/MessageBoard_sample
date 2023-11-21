/*
	Created by Kobayashi Atsuki.
	repositoryパッケージ。メッセージの保管を行う。
*/

package repository

type MessageRepository interface {
	Save(msg string) error   // 保管されているメッセージ
	List(msg []string) error // 保管されているメッセージのリスト
}
