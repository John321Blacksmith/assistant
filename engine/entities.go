package engine

type Category struct {
	Label    string   `json:"label"`
	Patterns []string `json:"patterns"`
}

type RefinedCategory struct {
	Label    string
	Patterns *UniqueElements
}

type Sentence struct {
	Data        *UniqueElements
	MainContext string
}

func (entity *Sentence) GetData() map[string]bool {
	return entity.Data.data
}

func (entity *Sentence) GetMainContext() string {
	return entity.MainContext
}

func (entity *Sentence) SetMainContext(context string) {
	entity.MainContext = context
}
