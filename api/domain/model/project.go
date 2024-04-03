package model

type ProjectID string

type Project struct {
	ID         ProjectID
	UserID     UserID
	Name       string
	Color      string
	IsArchived bool
}

// ContainsTask はプロジェクトがタスクを含むかを返す
func (p Project) ContainsTask(t Task) bool {
	return p.ID == t.ProjectID
}
