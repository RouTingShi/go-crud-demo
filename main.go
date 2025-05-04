package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/member", memberHandle)
	http.HandleFunc("/members", membersHandle)
	fmt.Println("連線")
	http.ListenAndServe(":8080", nil)
}
func membersHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "不支援此種方法", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)

}
func memberHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		idStr := r.FormValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "id必須是數字", http.StatusBadRequest)
			return
		} else if id > len(members)-1 || id < 0 {
			http.Error(w, "超出範圍", http.StatusBadRequest)
			return
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(members[id])
		}

	case http.MethodPost:
		var newMember Member
		err := json.NewDecoder(r.Body).Decode(&newMember)
		if err != nil {
			http.Error(w, "Json格式錯誤", http.StatusBadRequest)
			return
		}
		members = append(members, newMember)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newMember)

	case http.MethodDelete:
		idStr := r.FormValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "id必須是數字", http.StatusBadRequest)
			return
		} else if id > len(members)-1 || id < 0 {
			http.Error(w, "超出範圍", http.StatusBadRequest)
			return
		}
		members = append(members[:id], members[id+1:]...)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, "刪除成功")

	case http.MethodPut:
		idStr := r.FormValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "id必須是數字", http.StatusBadRequest)
			return
		} else if id > len(members)-1 || id < 0 {
			http.Error(w, "超出範圍", http.StatusBadRequest)
			return
		}
		var update Member
		err2 := json.NewDecoder(r.Body).Decode(&update)
		if err2 != nil {
			http.Error(w, "json格式錯誤", http.StatusBadRequest)
			return
		}
		members[id] = update
		fmt.Fprintln(w, "更新成功")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(update)

	default:
		http.Error(w, "不支援此種方法", http.StatusMethodNotAllowed)
		return
	}

}

type Member struct {
	Name  string `json:"name"`
	Point int    `json:"point"`
}

var members = []Member{
	{Name: "Alice", Point: 120},
	{Name: "Bob", Point: 90},
}
