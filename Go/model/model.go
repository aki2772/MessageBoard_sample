/*
	Created by Kobayashi Atsuki.
	modelパッケージ。メッセージの構造定義と生成を行う。
*/

package model

import (
	"time"
)

type Message struct {
	Name    string    // 名前
	Message string    // メッセージ
	Time    time.Time // タイムスタンプ
}
