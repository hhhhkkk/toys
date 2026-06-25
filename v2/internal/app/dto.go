package app

type CreateTask struct {
	FileName string `from: filename`
	FilePath string `from: filepath`
	Uid      int    `from: uid`
}
