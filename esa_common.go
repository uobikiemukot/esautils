package main

import (
	"fmt"
	"github.com/upamune/go-esa/esa"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getPost(client *esa.Client, team string, id int) *esa.PostResponse {
	post_res, err := client.Post.GetPost(team, id)
	if err != nil {
		panic("client.Post.GetPost Failed: " + err.Error())
	}
	return post_res
}

func getPosts(client *esa.Client, team string, user string) []esa.PostResponse {
	var posts []esa.PostResponse

	query := url.Values{}
	query.Add("per_page", "100")
	query.Add("user", user)
	query.Set("page", "1")

	for {
		res, err := client.Post.GetPosts(team, query)
		if err != nil {
			panic("client.Post.GetPosts Failed: " + err.Error())
		}

		for _, post := range res.Posts {
			posts = append(posts, post)
		}

		if res.NextPage == nil {
			break
		}
		query.Set("page", fmt.Sprint(res.NextPage))
	}
	return posts
}

func readArray(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic("ioutil.ReadFile Failed: " + err.Error())
	}
	return strings.Split(string(data), "\n")
}

func readBool(path string) bool {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic("ioutil.ReadFile Failed: " + err.Error())
	}

	boolean, err := strconv.ParseBool(string(data))
	if err != nil {
		panic("strconv.ParseBool Failed: " + err.Error())
	}
	return boolean
}

func readString(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic("ioutil.ReadFile Failed: " + err.Error())
	}
	return string(data)
}

func updatePost(client *esa.Client, team string, path string) {
	post_id, err := strconv.Atoi(filepath.Base(path))
	if err != nil {
		panic("strconv.Atoi Failed: " + err.Error())
	}

	post := esa.Post{
		BodyMd:   readString(path + "/BodyMd"),
		Category: readString(path + "/Category"),
		Message:  readString(path + "/Message"),
		Name:     readString(path + "/Name"),
		Tags:     readArray(path + "/Tags"),
		Wip:      readBool(path + "/Wip"),
	}

	if _, err := client.Post.Update(team, post_id, post); err != nil {
		panic("client.Post.Update Failed: " + err.Error())
	}
}

func newPost(client *esa.Client, team string, category string, name string) *esa.PostResponse {
	post := esa.Post{
		BodyMd:   "",
		Category: category,
		Message:  "",
		Name:     name,
		Tags:     []string{},
		Wip:      true,
	}

	post_res, err := client.Post.Create(team, post)
	if err != nil {
		panic("cilent.Post.Create Failed: " + err.Error())
	}
	return post_res
}

func writeBool(path string, data bool) {
	if err := ioutil.WriteFile(path, []byte(fmt.Sprint(data)), 0644); err != nil {
		panic("ioutil.WriteFile Failed: " + path + err.Error())
	}
}

func writeArray(path string, data []string) {
	if err := ioutil.WriteFile(path, []byte(strings.Join(data, "\n")), 0644); err != nil {
		panic("ioutil.WriteFile Failed: " + path + err.Error())
	}
}

func writeString(path string, data string) {
	if err := ioutil.WriteFile(path, []byte(data), 0644); err != nil {
		panic("ioutil.WriteFile Failed: " + path + err.Error())
	}
}

func writePost(post_res *esa.PostResponse) {
	prefix := fmt.Sprint(post_res.Number)

	if err := os.MkdirAll(prefix, 0755); err != nil {
		panic("os.Mkdir Failed: " + err.Error())
	}

	writeString(prefix+"/BodyMd", post_res.BodyMd)
	writeString(prefix+"/Category", post_res.Category)
	writeString(prefix+"/Message", post_res.Message)
	writeString(prefix+"/Name", post_res.Name)
	writeArray(prefix+"/Tags", post_res.Tags)
	writeBool(prefix+"/Wip", post_res.Wip)
}

func writePosts(posts []esa.PostResponse) {
	for _, post := range posts {
		writePost(&post)
	}
}
