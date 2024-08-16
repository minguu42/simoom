package sqlc

type Steps []Step

func (ss Steps) StepsByTaskID() map[string]Steps {
	m := map[string]Steps{}
	for _, s := range ss {
		m[s.TaskID] = append(m[s.TaskID], s)
	}
	return m
}

type Tasks []Task

func (ts Tasks) IDs() []string {
	ids := make([]string, 0, len(ts))
	for _, t := range ts {
		ids = append(ids, t.ID)
	}
	return ids
}
