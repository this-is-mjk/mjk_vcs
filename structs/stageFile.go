package structs

import "time"

type StageFile struct {
	DirPath string
	Name    string
	Hash    string
	Time    time.Time
	Size	int64
	FileType string
}
