package led

type LED interface {
	//Create a LED interface for using methods in both fakeled and realled
	Toggle()
	String() string
}
