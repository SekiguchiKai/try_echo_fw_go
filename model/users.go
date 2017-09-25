package model

import (
	"time"
	"github.com/satori/go.uuid"
	"github.com/labstack/echo"
	"github.com/mjibson/goon"
	"context"
)

// ユーザー
type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
}

// 新規のUserインスタンスの生成
func NewUser() *User {
	u := new(User)
	u.ID = uuid.NewV4().String()
	u.CreatedAt = time.Now()

	return u
}

// DatastoreにUserを登録する
func (u *User) Post(c echo.Context)(*User, error){
	// echo.Context => context.Contextに変換した後、goonをContextから生成
	g := goon.FromContext(context.Context(c))
	// DatastoreにUserのEntityを登録
	if _, err := g.Put(u); err != nil {
		return u, nil
	}

	return u, nil
}

//  DatastoreからUserを取得する
func(u *User)Get(c echo.Context) (*User, error) {
	// echo.Context => context.Contextに変換した後、goonをContextから生成
	g := goon.FromContext(context.Context(c))
	// key(userID)を指定してDatastoreからEntityを取得する、エラーの場合のハンドリングを同時に行う
	if err := g.Get(u); err != nil {
		return nil, err
	}

	return u, nil
}