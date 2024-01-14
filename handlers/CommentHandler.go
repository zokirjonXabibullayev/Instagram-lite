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

func CommentHendler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetAllCommet(w, r)
		timefunc.Timefunc()
	case "POST":
		CreateCommet(w, r)
		timefunc.Timefunc()
	case "PUT":
		UpdateCommet(w, r)
		timefunc.Timefunc()
	case "DELETE":
		DeleteCommet(w, r)
		timefunc.Timefunc()
	}
}
func CreateCommet(w http.ResponseWriter, r *http.Request) {
	var newComment models.CommentModel
	json.NewDecoder(r.Body).Decode(&newComment)

	var CommentData []models.CommentModel

	ByteComment, _ := os.ReadFile("db/comment.json")
	json.Unmarshal(ByteComment, &CommentData)

	//---------post

	var PostData []models.PostModel

	PostByte, _ := os.ReadFile("db/post.json")
	json.Unmarshal(PostByte, &PostData)

	//---------user

	var UserData []models.UserModel

	UserByte, _ := os.ReadFile("db/user.json")
	json.Unmarshal(UserByte, &UserData)

	var userFound bool
	for i := 0; i < len(UserData); i++ {
		if UserData[i].ID == newComment.UserID {
			userFound = true
			break
		}
	}
	if !userFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "User not found with ID", newComment.UserID)
		return
	}

	var postFound bool
	for i := 0; i < len(PostData); i++ {
		if PostData[i].ID == newComment.PostID {
			postFound = true
			break
		}

	}
	if !postFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Post not found with ID", newComment.ID)
		return
	}

	newComment.ID = halper.MaxIDComment(CommentData)
	CommentData = append(CommentData, newComment)

	res, _ := json.Marshal(CommentData)
	os.WriteFile("db/comment.json", res, 0)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Comment created!")
	fmt.Println("Comment created!")

}
func UpdateCommet(w http.ResponseWriter, r *http.Request) {
	var updateComment models.CommentModel
	json.NewDecoder(r.Body).Decode(&updateComment)

	var CommentData []models.CommentModel

	ByteComment, _ := os.ReadFile("db/comment.json")
	json.Unmarshal(ByteComment, &CommentData)

	//---------post

	var PostData []models.PostModel

	PostByte, _ := os.ReadFile("db/post.json")
	json.Unmarshal(PostByte, &PostData)

	//---------user

	var UserData []models.UserModel

	UserByte, _ := os.ReadFile("db/user.json")
	json.Unmarshal(UserByte, &UserData)

	var userFound bool
	for i := 0; i < len(UserData); i++ {
		if UserData[i].ID == updateComment.UserID {
			userFound = true
			break
		}
	}
	if !userFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "User not found with ID", updateComment.UserID)
		return
	}

	var postFound bool
	for i := 0; i < len(PostData); i++ {
		if PostData[i].ID == updateComment.PostID {
			postFound = true
			break
		}

	}
	if !postFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Post not found with ID", updateComment.ID)
		return
	}

	var commentFound bool
	for i := 0; i < len(CommentData); i++ {
		if CommentData[i].ID == updateComment.ID {
			if updateComment.Content != "" {
				CommentData[i].Content = updateComment.Content
			}
			commentFound = true
			break
		}
	}
	if !commentFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Comment not found with ID", updateComment.ID)
		return
	}
	res, _ := json.Marshal(CommentData)
	os.WriteFile("db/comment.json", res, 0)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Psot updated!")
	fmt.Println("Post updated!")
}
func DeleteCommet(w http.ResponseWriter, r *http.Request) {
	var deleteComment models.CommentModel
	json.NewDecoder(r.Body).Decode(&deleteComment)

	var CommentData []models.CommentModel
	ByteComment, _ := os.ReadFile("db/comment.json")
	json.Unmarshal(ByteComment, &CommentData)

	//---------post

	var PostData []models.PostModel

	PostByte, _ := os.ReadFile("db/post.json")
	json.Unmarshal(PostByte, &PostData)

	//---------user

	var UserData []models.UserModel

	UserByte, _ := os.ReadFile("db/user.json")
	json.Unmarshal(UserByte, &UserData)

	var userFound bool
	for i := 0; i < len(UserData); i++ {
		if UserData[i].ID == deleteComment.UserID {
			userFound = true
			break
		}
	}
	if !userFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "User not found with ID", deleteComment.UserID)
		return
	}

	var postFound bool
	for i := 0; i < len(PostData); i++ {
		if PostData[i].ID == deleteComment.PostID {
			postFound = true
			break
		}

	}
	if !postFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Post not found with ID", deleteComment.ID)
		return
	}

	var commentFound bool
	for i := 0; i < len(CommentData); i++ {
		if CommentData[i].ID == deleteComment.ID {
			CommentData = append(CommentData[:i], CommentData[i+1:]...)
			commentFound = true
			break
		}
	}
	if !commentFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Comment not found with ID", deleteComment.ID)
		return
	}
	res, _ := json.Marshal(CommentData)
	os.WriteFile("db/comment.json", res, 0)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Psot deleted!")
	fmt.Println("Post deleted!")
}
func GetAllCommet(w http.ResponseWriter, r *http.Request) {

	var CommentData []models.CommentModel
	ByteComment, _ := os.ReadFile("db/comment.json")
	json.Unmarshal(ByteComment, &CommentData)

	var PostData []models.PostModel
	PostByte, _ := os.ReadFile("db/post.json")
	json.Unmarshal(PostByte, &PostData)

	var UserData []models.UserModel
	UserByte, _ := os.ReadFile("db/user.json")
	json.Unmarshal(UserByte, &UserData)

	fmt.Fprintln(w, "__________________________________")
	for i := 0; i < len(CommentData); i++ {
		fmt.Fprintln(w, "Comment's ID", CommentData[i].ID)
		fmt.Fprintln(w, "Comment's content", CommentData[i].Content)
		fmt.Fprintln(w, "__________________________________")
	}
	fmt.Fprintln(w, "__________________________________")
	
	w.WriteHeader(http.StatusOK)
}
