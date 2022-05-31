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

func findPost(id string) Post {
	var post Post
	for i := range posts {
		if posts[i].ID == id {
			post = posts[i]
		}
	}
	return post
}

func updatePost(id string, updatePost Post) Post {
	post := findPost(id)
	post.Content = updatePost.Content
	fmt.Println("newContent", updatePost.Content)
	var idx int

	for i := range posts {
		if posts[i].ID == id {
			idx = i
			fmt.Println(posts[i])
		}
	}
	posts[idx] = post
	return post
}

// func deleteUser() {

// }
