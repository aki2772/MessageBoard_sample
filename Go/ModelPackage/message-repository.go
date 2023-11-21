/*
	Created by Kobayashi Atsuki.
	modelパッケージ。メッセージの構造定義と生成を行う。
*/

package model

import (
	"time"
)

type Message struct {
	mame    string    // 名前
	message string    // メッセージ
	time    time.Time //タイムスタンプ
}

// 新規メッセージの生成
func NewMessage(name string, message string) *Message {
	return &Message{
		name,
		message,
		time.Now(),
	}
}
