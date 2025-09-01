package blogposts

import (
	"bufio"
	"io"
	"io/fs"
)

type Post struct {
	Title       string
	Description string
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

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	scanner.Scan()
	titleLine := scanner.Text()

	scanner.Scan()
	descriptionLine := scanner.Text()

	post := Post{Title: titleLine[7:], Description: descriptionLine[13:]}
	return post, nil
}
