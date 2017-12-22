package randint

type Rand struct {
	Add   int
	First int
	Mod   int
}

func (r *Rand) Int() int {
	r.First = (r.First + r.Add) % r.Mod
	return r.First
}
