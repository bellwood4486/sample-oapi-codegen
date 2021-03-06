package oapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// BlogAPI implements OpenAPI-based endpoints.
type BlogAPI struct {
	posts  map[int]Post
	lock   sync.Mutex
	nextID int
}

func NewBlogAPI() *BlogAPI {
	return &BlogAPI{
		posts:  make(map[int]Post),
		nextID: 1,
	}
}

func (b *BlogAPI) FindPosts(w http.ResponseWriter, r *http.Request, params FindPostsParams) {
	_, _ = fmt.Fprintln(w, "success")
	//result := make([]Post, 0)
	//for _, v := range b.posts {
	//	result = append(result, v)
	//}
	//w.WriteHeader(http.StatusOK)
	//_ = json.NewEncoder(w).Encode(result)
}

func (b *BlogAPI) GetPost(w http.ResponseWriter, r *http.Request, id int) {
	// do nothing
	fmt.Printf("called GetPost with id:%d\n", id)
	_, _ = fmt.Fprintln(w, "success")
}

func (b *BlogAPI) AddPost(w http.ResponseWriter, r *http.Request) {
	var newPost NewPost
	if err := json.NewDecoder(r.Body).Decode(&newPost); err != nil {
		log.Println(err)
		b.sendBlogAPIError(w, http.StatusBadRequest, "Invalid format for NewPost")
		return
	}
	_, _ = fmt.Fprintln(w, "success")
	//b.lock.Lock()
	//defer b.lock.Unlock()
	//
	//var post Post
	//post.Title = newPost.Title
	//post.Content = newPost.Content
	//post.Id = b.nextID
	//b.nextID++
	//
	//b.posts[post.Id] = post
	//
	//w.WriteHeader(http.StatusCreated)
	//_ = json.NewEncoder(w).Encode(post)
}

func (b *BlogAPI) sendBlogAPIError(w http.ResponseWriter, code int, message string) {
	err := Error{
		Code:    code,
		Message: message,
	}
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(err)
}
