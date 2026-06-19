package engine

type Category struct {
	Label    string   `json:"label"`
	Patterns []string `json:"patterns"`
}

type RefinedCategory struct {
	Label    string
	Patterns map[string]bool
}

type Sentence struct {
	data        *UniqueElements
	mainContext string
}

func (entity *Sentence) GetData() map[string]bool {
	return entity.data.data
}

func (entity *Sentence) GetMainContext() string {
	return entity.mainContext
}

func (entity *Sentence) SetMainContext(context string) {
	entity.mainContext = context
}
