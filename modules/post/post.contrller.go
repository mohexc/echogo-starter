package post

import (
	"fmt"
	"strconv"
)

var posts = []Post{
	{ID: "1", Content: "This is a post 1"},
	{ID: "2", Content: "This is a post 2"},
	{ID: "3", Content: "This is a post 3"},
}

func createPost(newPost Post) Post {
	newPost.ID = strconv.Itoa(len(posts) + 1)
	posts = append(posts, newPost)
	return posts[len(posts)-1]
}

func findPosts() []Post {
	return posts
}

func findPost(id string) (Post, int, error) {
	var post Post
	var idx int
	fmt.Println("findPost", idx)
	for i := range posts {
		if posts[i].ID == id {
			post = posts[i]
			idx = i
		}
	}
	fmt.Println("findPost", idx)
	return post, idx, nil
}

func updatePost(id string, updatePost Post) Post {
	post, idx, _ := findPost(id)
	post.Content = updatePost.Content
	posts[idx] = post
	return post
}

func deletePost(id string) (string, error) {
	_, idx, _ := findPost(id)
	posts = append(posts[:idx], posts[idx+1:]...)
	return "delete sucess", nil
}
