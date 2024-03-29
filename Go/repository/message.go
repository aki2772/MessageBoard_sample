/*
	Created by Kobayashi Atsuki.
	repositoryパッケージ。メッセージの保管を行う。
*/

package repository

import "github.com/aki2772/MessageBoard_sample/model"

// ここはrepositorypackage内のMessageという書き方で良い
type Message interface {
	Save(*model.Message) error       // メッセージを保管する
	List() ([]*model.Message, error) // 保管されているメッセージのリストを引き出す
}
