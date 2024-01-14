package handlers

import (
	"encoding/json"
	"fmt"
	"instagram/halper"
	"instagram/models"
	"instagram/timefunc"
	"net/http"
	"os"
)

func PostHendler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetAllPost(w, r)
		timefunc.Timefunc()
	case "POST":
		CreatePost(w, r)
		timefunc.Timefunc()
	case "PUT":
		UpdatePost(w, r)
		timefunc.Timefunc()
	case "DELETE":
		DeletePost(w, r)
		timefunc.Timefunc()
	}
}
func CreatePost(w http.ResponseWriter, r *http.Request) {
	var newPost models.PostModel
	json.NewDecoder(r.Body).Decode(&newPost)

	var PostData []models.PostModel

	PostByte, _ := os.ReadFile("db/post.json")
	json.Unmarshal(PostByte, &PostData)

	//------------------------------------

	var UserData []models.UserModel
	UserByte, _ := os.ReadFile("db/user.json")
	json.Unmarshal(UserByte, &UserData)
	
	var userFound bool
	for i := 0; i < len(UserData); i++ {
		if UserData[i].ID == newPost.UserID {
			userFound = true
			break
		}
	}
	if !userFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "User not found with ID", newPost.UserID)
		return
	}
	newPost.ID = halper.MaxIDPost(PostData)
	PostData = append(PostData, newPost)

	res, _ := json.Marshal(PostData)
	os.WriteFile("db/post.json", res, 0)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Psot created!")
	fmt.Println("Post created!")
}
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	var updatePost models.PostModel
	json.NewDecoder(r.Body).Decode(&updatePost)

	var PostData []models.PostModel

	PostByte, _ := os.ReadFile("db/post.json")
	json.Unmarshal(PostByte, &PostData)

	//------------------------------------

	var UserData []models.UserModel
	UserByte, _ := os.ReadFile("db/user.json")
	json.Unmarshal(UserByte, &UserData)
	var userFound bool
	for i := 0; i < len(UserData); i++ {
		if UserData[i].ID == updatePost.UserID {
			userFound = true
			break
		}
	}
	if !userFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "User not found with ID", updatePost.UserID)
		return
	}

	var postFound bool
	for i := 0; i < len(PostData); i++ {
		if PostData[i].ID == updatePost.ID {
			if updatePost.Title != "" {
				PostData[i].Title = updatePost.Title
			}
			if updatePost.Content != "" {
				PostData[i].Content = updatePost.Content
			}

			postFound = true
			break
		}

	}
	if !postFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Post not found with ID", updatePost.ID)
		return
	}
	res, _ := json.Marshal(PostData)
	os.WriteFile("db/post.json", res, 0)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Psot updated!")
	fmt.Println("Post updated!")
}
func DeletePost(w http.ResponseWriter, r *http.Request) {
	var deletePost models.PostModel
	json.NewDecoder(r.Body).Decode(&deletePost)

	var PostData []models.PostModel

	PostByte, _ := os.ReadFile("db/post.json")
	json.Unmarshal(PostByte, &PostData)

	//------------------------------------

	var UserData []models.UserModel
	UserByte, _ := os.ReadFile("db/user.json")
	json.Unmarshal(UserByte, &UserData)
	var userFound bool
	for i := 0; i < len(UserData); i++ {
		if UserData[i].ID == deletePost.UserID {
			userFound = true
			break
		}
	}
	if !userFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "User not found with ID", deletePost.UserID)
		return
	}

	var postFound bool
	for i := 0; i < len(PostData); i++ {
		if PostData[i].ID == deletePost.ID {
			PostData = append(PostData[:i], PostData[i+1:]...)
			postFound = true
			break
		}

	}
	if !postFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Post not found with ID", deletePost.ID)
		return
	}
	res, _ := json.Marshal(PostData)
	os.WriteFile("db/post.json", res, 0)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Psot deleted!")
	fmt.Println("Post deleted!")
}
func GetAllPost(w http.ResponseWriter, r *http.Request) {

	var PostData []models.PostModel
	PostByte, _ := os.ReadFile("db/post.json")
	json.Unmarshal(PostByte, &PostData)

	var UserData []models.UserModel
	UserByte, _ := os.ReadFile("db/user.json")
	json.Unmarshal(UserByte, &UserData)

	fmt.Fprintln(w, "__________________________________")
	for i := 0; i < len(PostData); i++ {
		fmt.Fprintln(w, "Post's ID:", PostData[i].ID)
		fmt.Fprintln(w, "Post's Title:", PostData[i].Title)
		fmt.Fprintln(w, "Post's Content:", PostData[i].Content)
		fmt.Fprintln(w, "Post's Likes count:", PostData[i].Likes)
		fmt.Fprintln(w, "__________________________________")
	}
	fmt.Fprintln(w, "__________________________________")

	w.WriteHeader(http.StatusOK)
	
}
