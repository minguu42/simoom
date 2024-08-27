package model

import "context"

type UserID string

type User struct {
	ID       UserID
	Name     string
	Email    string
	Password string
}

// HasProject はユーザがプロジェクトを所有しているかを返す
func (u User) HasProject(p Project) bool {
	return u.ID == p.UserID
}

// HasStep はユーザがステップを所有しているかを返す
func (u User) HasStep(s Step) bool {
	return u.ID == s.UserID
}

// HasTag はユーザがタグを所有しているかを返す
func (u User) HasTag(t Tag) bool {
	return u.ID == t.UserID
}

// HasTask はユーザがタスクを所有しているかを返す
func (u User) HasTask(t Task) bool {
	return u.ID == t.UserID
}

type userKey struct{}

// ContextWithUser は ctx に model.User をセットする
func ContextWithUser(ctx context.Context, u User) context.Context {
	return context.WithValue(ctx, userKey{}, u)
}

// UserFromContext は ctx から model.User を取り出す
// ctx にユーザがセットされていない場合は空の構造体値を返す
func UserFromContext(ctx context.Context) User {
	return ctx.Value(userKey{}).(User)
}
