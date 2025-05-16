package nfile


type NFileType string

const (
	FileStreamType NFileType = "stream"
	FileDataType   NFileType = "data"
)

type NFile struct {
	isData   bool
	isStream bool
	name     string
	path     string
}

func (n *NFile) GetContent() ([]byte, error) {
	panic("unimplemented")
}

func (n *NFile) GetName() string {
	return n.name
}

func (n *NFile) GetPath() string {
	return n.path
}

func (n *NFile) GetSize() int64 {
	panic("unimplemented")
}
