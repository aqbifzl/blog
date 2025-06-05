package dto_output

import "time"

type PostMetadata struct {
	Title    string
	Date     time.Time
	Filename string
}

type Post struct {
	Title   string
	Date    time.Time
	Content string
}
