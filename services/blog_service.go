package services

import (
	"bytes"
	"fmt"
	"blog/config"
	dto_output "blog/dto/output"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/yuin/goldmark"
)

type BlogService struct {
}

func NewBlogService() *BlogService {
	return &BlogService{}
}

func (s BlogService) GetPostsMetadata(page int) ([]*dto_output.PostMetadata, error) {
	files, err := filepath.Glob("posts/*")
	if err != nil {
		return nil, fmt.Errorf("failed to get files in posts: %w", err)
	}

	var posts []*dto_output.PostMetadata

	for _, path := range files {
		raw, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("read file %q: %w", path, err)
		}

		lines := strings.Split(string(raw), "\n")
		if len(lines) < 2 {
			return nil, fmt.Errorf("file %q: not enough lines for metadata", path)
		}

		title_line := lines[0]
		date_line := lines[1]

		const title_prefix = "title:"
		const date_prefix = "date:"

		if !strings.HasPrefix(title_line, title_prefix) {
			return nil, fmt.Errorf("file %q: first line must start with %q", path, title_prefix)
		}
		if !strings.HasPrefix(date_line, date_prefix) {
			return nil, fmt.Errorf("file %q: second line must start with %q", path, date_prefix)
		}

		title := strings.TrimPrefix(title_line, title_prefix)
		date := strings.TrimPrefix(date_line, date_prefix)

		parsed_date, err := time.Parse(config.DATE_FORMAT, date)
		if err != nil {
			return nil, fmt.Errorf("file %q: cannot parse date %q: %w", path, date, err)
		}

		posts = append(posts, &dto_output.PostMetadata{
			Title:    title,
			Date:     parsed_date,
			Filename: filepath.Base(path),
		})
	}

	return posts, nil
}

func (s BlogService) GetBlogPost(path string) (*dto_output.Post, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file %q: %w", path, err)
	}

	lines := strings.Split(string(raw), "\n")
	if len(lines) < 2 {
		return nil, fmt.Errorf("file %q: not enough lines for metadata", path)
	}

	titleLine := strings.TrimSpace(lines[0])
	dateLine := strings.TrimSpace(lines[1])

	const titlePrefix = "title:"
	const datePrefix = "date:"

	if !strings.HasPrefix(titleLine, titlePrefix) {
		return nil, fmt.Errorf("file %q: first line must start with %q", path, titlePrefix)
	}
	if !strings.HasPrefix(dateLine, datePrefix) {
		return nil, fmt.Errorf("file %q: second line must start with %q", path, datePrefix)
	}

	title := strings.TrimPrefix(lines[0], titlePrefix)
	dateStr := strings.TrimPrefix(lines[1], datePrefix)

	parsedDate, err := time.Parse("2006-01-02 15:04", dateStr)
	if err != nil {
		return nil, fmt.Errorf("file %q: cannot parse date %q (use layout YYYY-MM-DD HH:MM): %w", path, dateStr, err)
	}

	var content bytes.Buffer
	err = goldmark.Convert([]byte(strings.Join(lines[2:], "\n")), &content)
	if err != nil {
		return nil, err
	}

	return &dto_output.Post{
		Title:   title,
		Date:    parsedDate,
		Content: content.String(),
	}, nil
}
