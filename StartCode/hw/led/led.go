package led

type LED interface {
	Toggle()
	String() string
}
