package model

type User struct {
	ID       string
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
