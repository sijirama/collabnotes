package utils

type Index map[string][]int

func (idx Index) Add(docs []document) {
	for _, doc := range docs {
		for _, token := range analyze(doc.Text) {
			ids := idx[token]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				//dont add same id twice my nigger
				continue
			}
			idx[token] = append(ids, doc.ID)
		}
	}
}

func intersection(a, b []int) []int {
	maxLen := len(a)

	if len(b) > len(a) {
		maxLen = len(b)
	}

	r := make([]int, 0, maxLen)

	var i, j int

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}

	return r
}

func (idx Index) Search(query string) []int {
	var r []int
	for _, token := range analyze(query) {
		if ids, ok := idx[token]; ok {
			if r == nil {
				r = ids
			} else {
				r = intersection(r, ids)
			}
		} else {
			return nil
		}
	}
	return r
}
