package model

type Project struct {
	ID         string
	UserID     string
	Name       string
	Color      string
	IsArchived bool
}

// ContainsTask はプロジェクトがタスクを含むかを返す
func (p Project) ContainsTask(t Task) bool {
	return p.ID == t.ProjectID
}
