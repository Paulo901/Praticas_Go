package testes

import (
	"fmt"
	"sort"
)

func Padroes() {

	a := []string{"de", "abc", "fghi"}
	b := []int{1, 3, 2, 4, 6, 5, 7}

	sort.Strings(a)
	fmt.Println(a)
	//------------------
	sort.Ints(b)
	fmt.Println(b)
}

type PC struct {
	PotenProcess float64
	MemoriaRam   int
}
type ordenaPorRam []PC

// esses são os métodos de Sort, que implementei para poder atribuir à interface
func (a ordenaPorRam) Len() int           { return len(a) }
func (a ordenaPorRam) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ordenaPorRam) Less(i, j int) bool { return a[i].MemoriaRam < a[j].MemoriaRam }

func Forjados() {

	pcs := []PC{{1.2, 8}, {PotenProcess: 3.1, MemoriaRam: 4}}

	sort.Sort(ordenaPorRam(pcs))
	fmt.Println(pcs)

}
