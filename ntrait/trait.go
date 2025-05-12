package ntrait

type SudoDrive[E any, F any, T any] interface {
	DirWorker
	Drive[E, F, T]
}

type ElementType = string

const (
	FolderType ElementType = "folder"
	FileType   ElementType = "file"
)

type Drive[E any, F any, T any] interface {
	Drive_Create[E, F, T]
	Drive_Read[E, F, T]
	Drive_Update[E, F, T]
	Drive_Delete[E, F, T]
}

type Drive_Create[E any, F any, T any] interface {
	CreateFolder(name string) error
	CreateFileEmpty(name string) error
	CreateFileWithContent(name string, content []byte) error
}

type Drive_Read[E any, F any, T any] interface {
	GetFolder(name string) (Folder[F, E, T], error)
	GetFile(name string) (File[T], error)
	GetList() ([]Element[E, F, T], error)
}

type Drive_Update[E any, F any, T any] interface {
	UpdateFolder(name string, newName string) error
	UpdateFile(name string, newName string) error
	UpdateFileContent(name string, content []byte) error
}

type Drive_Delete[E any, F any, T any] interface {
	DeleteFolder(name string) error
	DeleteFile(name string) error
}

type Element[E any, F any, T any] interface {
	GetName() string
	GetPath() string
	Type() ElementType
	GetSize() int64
}

type Folder[E any, F any, T any] interface {
	GetName() string
	GetPath() string
	GetSize() int64
	GetList() ([]Element[E, F, T], error)
}

type File[T any] interface {
	GetName() string
	GetPath() string
	GetSize() int64
	GetContent() ([]byte, error)
}
