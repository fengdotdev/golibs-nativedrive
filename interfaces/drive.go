package interfaces

type ElementType = string

const (
	FolderType ElementType = "folder"
	FileType   ElementType = "file"
)

type Element interface {
	GetName() string
	GetPath() string
	Type() ElementType
	GetSize() int64
	GetChildren() []Element
	GetParent() Element
	IsFolder() bool
	IsFile() bool
	IsRoot() bool
}

type FileData interface {
	GetContent() ([]byte, error)
	GetName() string
	GetPath() string
	GetSize() int64

}
