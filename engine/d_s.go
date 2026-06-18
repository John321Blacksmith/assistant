package engine

type UniqueElements struct {
	data map[string]bool
}

func NewUniqueElements() *UniqueElements {
	return &UniqueElements{data: map[string]bool{}}
}

func (entity *UniqueElements) AddLiteral(items ...string) {
	for _, item := range items {
		entity.data[item] = true
	}
}

func (entity *UniqueElements) Intersection(comparableArr map[string]bool) map[string]bool {
	var intersection map[string]bool = map[string]bool{}

	if len(entity.data) > len(comparableArr) {
		for item := range comparableArr {
			if entity.data[item] {
				intersection[item] = true
			}
		}
	} else {
		for item := range entity.data {
			if comparableArr[item] {
				intersection[item] = true
			}
		}
	}

	return intersection
}
