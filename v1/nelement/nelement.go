package nelement

type NElement struct {
	children []NElement
	parent   *NElement
	name     string
	root     *NElement
	isRoot   bool
}




