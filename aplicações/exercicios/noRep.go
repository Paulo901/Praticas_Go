package exercicios

type User struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ordenaPorIdade []User

func (a ordenaPorIdade) Len() int           { return len(a) }
func (a ordenaPorIdade) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ordenaPorIdade) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ordenaPorSobrenome []User

func (a ordenaPorSobrenome) Len() int           { return len(a) }
func (a ordenaPorSobrenome) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ordenaPorSobrenome) Less(i, j int) bool { return a[i].Age < a[j].Age }
