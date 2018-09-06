package domain

type Element struct {
	Items []Item
}

type Item struct {
	ApiVersion string
	Kind       string
	Spec       Spec
}

type Spec struct {
	Hosts []string
	Ports []Port
}

type Port struct {
	Number   string
	Protocol string
}
