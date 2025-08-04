package beginner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type InstagramData struct {
	Relationship_following []Wrapper `json:"relationships_following"`
}

type Wrapper struct {
	StringListData []StringData `json:"string_list_data"`
}

type StringData struct {
	Value string `json:"value"`
	Href  string `json:"href"`
}

func CekFollower() {

	// dir, _ := os.Getwd()
	// fmt.Println("Current working directory:", dir)

	file_followers, err := os.Open("1_beginner/followers_1.json")
	if err != nil {
		fmt.Println("Error opening file_followers:", err)
		return
	}
	defer file_followers.Close()

	file_following, err := os.Open("1_beginner/following.json")
	if err != nil {
		fmt.Println("Error opening file_following.json")
		return
	}
	defer file_following.Close()

	data_followers, err := ioutil.ReadAll(file_followers)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	data_following, err := ioutil.ReadAll(file_following)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var wrapper_followers []Wrapper
	err = json.Unmarshal([]byte(data_followers), &wrapper_followers)
	if err != nil {
		fmt.Println("Error unmarshaling JSON follower:", err)
		return
	}

	var wrapper_following InstagramData
	err = json.Unmarshal([]byte(data_following), &wrapper_following)
	if err != nil {
		fmt.Println("Error unmarshaling JSON following:", err)
		return
	}

	date_followers := make(map[string]bool)
	for _, w := range wrapper_followers {
		for _, s := range w.StringListData {
			date_followers[s.Value] = true
		}
	}

	date_following := make(map[string]bool)
	for _, w := range wrapper_following.Relationship_following {
		for _, s := range w.StringListData {
			date_following[s.Value] = true
		}
	}

	nomor := 1
	fmt.Println("Yang tidak follback")
	for i := range date_following {
		if !date_followers[i] {
			fmt.Println(nomor, ".", i)
			nomor++
		}
	}
}
