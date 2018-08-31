package domain

type Element struct {
	Items []Item
}

type Item struct {
	Kind string
	Spec []Spec
}

type Spec struct {
	Hosts []string
}
