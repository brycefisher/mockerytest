package test

type Proxy interface {
	Passthrough(s string) string
}
