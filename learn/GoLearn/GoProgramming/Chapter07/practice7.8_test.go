package Chapter01

import (
	"fmt"
	"sort"
	"testing"
)

type book struct {
	name   string
	price  float64
	author string
}

type byFunc func(i, j int) bool
type tableSlice struct {
	lists    []*book
	lessFunc []byFunc
}

func (ts tableSlice) Len() int {
	return len(ts.lists)
}

func (ts tableSlice) Swap(i, j int) {
	ts.lists[i], ts.lists[j] = ts.lists[j], ts.lists[i]
}

func (ts tableSlice) Less(i, j int) bool {
	for t := len(ts.lessFunc) - 1; t >= 0; t-- {
		if ts.lessFunc[t](i, j) {
			return true
		} else if !ts.lessFunc[t](j, i) {
			continue
		}
	}
	return false
}

func (ts tableSlice) byName(i, j int) bool {
	return ts.lists[i].name < ts.lists[j].name
}
func (ts tableSlice) byPrice(i, j int) bool {
	return ts.lists[i].price < ts.lists[j].price
}

func start() {
	book1 := book{"GoLang", 65.50, "Aideng"}
	book2 := book{"PHP", 45.50, "Sombody"}
	book3 := book{"C", 45.00, "Tan"}

	ts := tableSlice{
		lists: []*book{&book1, &book2, &book3},
	}

	ts.lessFunc = []byFunc{ts.byPrice, ts.byName}

	sort.Sort(ts)
	for _, book := range ts.lists {
		fmt.Println(*book)
	}

}

func Test078(t *testing.T) {
	start()
}
