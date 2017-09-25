package main

import (
	"time"
	"github.com/satori/go.uuid"
	"github.com/labstack/echo"
	"github.com/mjibson/goon"
	"google.golang.org/appengine"
)

// ユーザー
type User struct {
	ID        string    `json:"id" goon:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
}

// 新規のUserインスタンスの生成
func NewUser(name string) *User {
	u := new(User)
	u.ID = uuid.NewV4().String()
	u.Name = name
	u.CreatedAt = time.Now()

	return u
}

// DatastoreにUserを登録する
func (u *User) Post(c echo.Context)(*User, error){
	// echo.Context => context.Contextに変換した後、goonをContextから生成
	g := goon.FromContext(appengine.NewContext(c.Request()))
	// DatastoreにUserのEntityを登録
	if _, err := g.Put(u); err != nil {
		ErrorLog(appengine.NewContext(c.Request()), err.Error())
		return u, nil
	}

	return u, nil
}

//  DatastoreからUserを取得する
func(u *User)Get(c echo.Context) (*User, error) {
	// echo.Context => context.Contextに変換した後、goonをContextから生成
	g := goon.FromContext(appengine.NewContext(c.Request()))
	// key(userID)を指定してDatastoreからEntityを取得する、エラーの場合のハンドリングを同時に行う
	if err := g.Get(u); err != nil {
		ErrorLog(appengine.NewContext(c.Request()), err.Error())
		return nil, err
	}

	return u, nil
}