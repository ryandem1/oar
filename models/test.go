package models

type Test struct {
	Name    string
	Outcome Outcome
	Action  ActionTaken
	Doc     map[string]any
}
