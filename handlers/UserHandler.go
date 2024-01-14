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

func UserHendler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetAllUser(w, r)
		timefunc.Timefunc()

	case "POST":
		CreateUser(w, r)
		timefunc.Timefunc()
	case "PUT":
		UpdateUser(w, r)
		timefunc.Timefunc()
	case "DELETE":
		DeleteUser(w, r)
		timefunc.Timefunc()
	}
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.UserModel
	json.NewDecoder(r.Body).Decode(&newUser)

	var UserData []models.UserModel

	UserByte, _ := os.ReadFile("db/user.json")
	json.Unmarshal(UserByte, &UserData)

	newUser.ID = halper.MaxIDUser(UserData)

	UserData = append(UserData, newUser)

	res, _ := json.Marshal(UserData)
	os.WriteFile("db/user.json", res, 0)

	
	
	
	

	fmt.Println("User yaratildi ID", newUser.ID,)

	fmt.Fprintln(w, "User yaratildi:")
	json.NewEncoder(w).Encode(newUser)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser models.UserModel
	json.NewDecoder(r.Body).Decode(&updateUser)

	var UserData []models.UserModel

	UserByte, _ := os.ReadFile("db/user.json")
	json.Unmarshal(UserByte, &UserData)

	var UserFound bool
	for i := 0; i < len(UserData); i++ {
		if UserData[i].ID == updateUser.ID {
			if updateUser.Firstname != "" {
				UserData[i].Firstname = updateUser.Firstname
			}
			if updateUser.Lastname != "" {
				UserData[i].Lastname = updateUser.Lastname
			}
			UserFound = true
			break
		}
	}
	if !UserFound {
		fmt.Println("User topilmadi ID", updateUser.ID)
		fmt.Fprintln(w, "User topilmadi ID", updateUser.ID)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res, _ := json.Marshal(UserData)
	os.WriteFile("db/user.json", res, 0)

	fmt.Println("User o'zgardi ID", updateUser.ID)
	fmt.Fprintln(w, "User o'zgardi:")
	json.NewEncoder(w).Encode(updateUser)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var DeleteUser models.UserModel
	json.NewDecoder(r.Body).Decode(&DeleteUser)

	var UserData []models.UserModel
	UserByte, _ := os.ReadFile("db/user.json")
	json.Unmarshal(UserByte, &UserData)

	var UserFound bool
	for i := 0; i < len(UserData); i++ {
		if UserData[i].ID == DeleteUser.ID {

			UserData = append(UserData[:i], UserData[i+1:]...)

			UserFound = true
			break
		}
	}
	if !UserFound {
		fmt.Println("User topilmadi ID", DeleteUser.ID)
		fmt.Fprintln(w, "User topilmadi ID", DeleteUser.ID)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res, _ := json.Marshal(UserData)
	os.WriteFile("db/user.json", res, 0)

	fmt.Println("User ochirildi ID", DeleteUser.ID)
	fmt.Fprintln(w, "User ochirildi:")
	json.NewEncoder(w).Encode(DeleteUser)
}
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	var UserData []models.UserModel
	UserByte, _ := os.ReadFile("db/user.json")
	json.Unmarshal(UserByte, &UserData)

	var PostData []models.PostModel
	PostByte, _ := os.ReadFile("db/post.json")
	json.Unmarshal(PostByte, &PostData)

	var CommentData []models.CommentModel
	ByteComment, _ := os.ReadFile("db/comment.json")
	json.Unmarshal(ByteComment, &CommentData)

	fmt.Fprintln(w, "__________________________________")
	for i := 0; i < len(UserData); i++ {
		fmt.Fprintln(w, "User's ID:", UserData[i].ID)
		fmt.Fprintln(w, "User's Fristname:", UserData[i].Firstname)
		fmt.Fprintln(w, "User's Lastname:", UserData[i].Lastname)
		fmt.Fprintln(w, "User's Posts:", UserData[i].Posts)
		fmt.Fprintln(w, "__________________________________")
		for j := 0; j < len(PostData); j++ {
			if PostData[j].UserID == UserData[i].ID {
				fmt.Fprintln(w, "  Post's ID:", PostData[j].ID)
				fmt.Fprintln(w, "  Post's Title:", PostData[j].Title)
				fmt.Fprintln(w, "  Post's Content:", PostData[j].Content)
				fmt.Fprintln(w, "  Post's Likes count:", PostData[j].Likes)
				fmt.Fprintln(w, "  __________________________________")
			}
			for l := 0; l < len(CommentData); l++ {
				if CommentData[l].UserID == UserData[i].ID && CommentData[l].PostID == PostData[j].ID {
					fmt.Fprintln(w, "      Comment's ID:", CommentData[i].ID)
					fmt.Fprintln(w, "      Comment's content:", CommentData[i].Content)
					fmt.Fprintln(w, "      __________________________________")
				}
			}
		}
	}
	fmt.Fprintln(w, "__________________________________")

	w.WriteHeader(http.StatusOK)

}
