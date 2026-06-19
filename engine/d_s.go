package engine

type UniqueElements struct {
	data map[string]bool
}

func NewUniqueElements() *UniqueElements {
	return &UniqueElements{data: map[string]bool{}}
}

func (ds *UniqueElements) AddLiteral(items ...string) {
	for _, item := range items {
		ds.data[item] = true
	}
}

func (ds *UniqueElements) Intersection(comparableArr map[string]bool) map[string]bool {
	var intersection map[string]bool = map[string]bool{}

	if len(ds.data) > len(comparableArr) {
		for item := range comparableArr {
			if ds.data[item] {
				intersection[item] = true
			}
		}
	} else {
		for item := range ds.data {
			if comparableArr[item] {
				intersection[item] = true
			}
		}
	}

	return intersection
}

type DataSet struct {
	Categories []Category `json:"categories"`
}

func (ds *DataSet) Refine() *RefinedDataSet {
	var categories []RefinedCategory
	for _, cat := range ds.Categories {
		var uniquePatterns map[string]bool = map[string]bool{}
		for _, pattern := range cat.Patterns {
			uniquePatterns[pattern] = true
		}
		categories = append(
			categories,
			RefinedCategory{
				Label:    cat.Label,
				Patterns: uniquePatterns,
			})
	}
	return &RefinedDataSet{Categories: categories}
}

type RefinedDataSet struct {
	Categories []RefinedCategory
}

type KnownData struct {
	sentences []Sentence
}

func (ds *KnownData) AddSentence(sentence Sentence) {
	ds.sentences = append(ds.sentences, sentence)
}

// Find the main context of the known
// []Sentence
func (ds *KnownData) GetMainConext() string {
	var contexts []string
	var mainContext string
	for _, s := range ds.sentences {
		contexts = append(contexts, s.mainContext)
	}
	var freqMap map[string]int = map[string]int{}

	for _, ctx := range contexts {
		if _, exists := freqMap[ctx]; !exists {
			freqMap[ctx] = 1
		} else {
			freqMap[ctx] += 1
		}
	}

	var greatestCxt string
	var greatestFreq int
	for ctx, freq := range freqMap {
		if freq > greatestFreq {
			greatestCxt = ctx
			greatestFreq = freq
		}
	}

	mainContext = greatestCxt

	return mainContext
}

func (ds *KnownData) PresentData() []Sentence {
	return ds.sentences
}

type UnknownData struct {
	sentences []Sentence
}

func (ds *UnknownData) AddSentence(sentence Sentence) {
	ds.sentences = append(ds.sentences, sentence)
}

func (ds *UnknownData) PresentData() []Sentence {
	return ds.sentences
}
