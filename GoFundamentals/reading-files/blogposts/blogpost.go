package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func NewPostsFromFS(fsys fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fsys, ".")

	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fsys, f.Name())
		if err != nil {
			continue
		}

		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fsys fs.FS, fname string) (Post, error) {
	postFile, err := fsys.Open(fname)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	return newPost(postFile)

}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

func newPost(postBody io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postBody)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	return Post{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
		Tags:        strings.Split(readMetaLine(tagsSeparator), ", "),
		Body:        readBody(scanner),
	}, nil
}

func readBody(s *bufio.Scanner) string {
	s.Scan()
	b := bytes.Buffer{}
	for s.Scan() {
		fmt.Fprintln(&b, s.Text())
	}
	return strings.TrimSuffix(b.String(), "\n")
}
