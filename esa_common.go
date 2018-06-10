package main

import (
	"fmt"
	"github.com/uobikiemukot/go-esa/esa"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func getPost(client *esa.Client, team string, id int) (*esa.PostResponse, error) {
	post, err := client.Post.GetPost(team, id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func getPosts(client *esa.Client, team string, user string) ([]esa.PostResponse, error) {
	var posts []esa.PostResponse

	query := url.Values{}
	query.Add("per_page", "100")
	query.Add("user", user)
	query.Set("page", "1")

	for {
		res, err := client.Post.GetPosts(team, query)
		if err != nil {
			return nil, err
		}

		for _, p := range res.Posts {
			posts = append(posts, p)
		}

		if res.NextPage == nil {
			break
		}
		query.Set("page", fmt.Sprint(res.NextPage))
	}
	return posts, nil
}

func readArray(path string) []string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ioutil.ReadFile Failed: %v\n", err)
		return []string{}
	}
	return strings.Split(string(data), "\n")
}

func readBool(path string) bool {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ioutil.ReadFile Failed: %v\n", err)
		return true
	}

	b, err := strconv.ParseBool(string(data))
	if err != nil {
		fmt.Fprintf(os.Stderr, "ioutil.ReadFile Failed: %v\n", err)
		return true
	}
	return b
}

func readString(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ioutil.ReadFile Failed: %v\n", err)
		return ""
	}
	return string(data)
}

func updatePost(client *esa.Client, team string, path string) error {

	id, err := strconv.Atoi(filepath.Base(path))
	if err != nil {
		return err
	}

	post := esa.Post{
		BodyMd:   readString(path + "/BodyMd"),
		Category: readString(path + "/Category"),
		Message:  readString(path + "/Message"),
		Name:     readString(path + "/Name"),
		Tags:     readArray(path + "/Tags"),
		Wip:      readBool(path + "/Wip"),
	}

	if _, err := client.Post.Update(team, id, post); err != nil {
		return err
	}

	return nil
}

func newPost(client *esa.Client, team string, category string, name string) (*esa.PostResponse, error) {
	post := esa.Post{
		BodyMd:   "",
		Category: category,
		Message:  "",
		Name:     name,
		Tags:     []string{},
		Wip:      true,
	}

	res, err := client.Post.Create(team, post)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func writeBool(path string, data bool) {
	if err := ioutil.WriteFile(path, []byte(fmt.Sprint(data)), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "ioutil.WriteFile Failed: %v\n", err)
	}
}

func writeArray(path string, data []string) {
	if err := ioutil.WriteFile(path, []byte(strings.Join(data, "\n")), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "ioutil.WriteFile Failed: %v\n", err)
	}
}

func writeString(path string, data string) {
	if err := ioutil.WriteFile(path, []byte(data), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "ioutil.WriteFile Failed: %v\n", err)
	}
}

func writePost(post_res *esa.PostResponse) error {
	prefix := fmt.Sprint(post_res.Number)

	if err := os.MkdirAll(prefix, 0755); err != nil {
		return err
	}

	writeString(prefix+"/BodyMd", post_res.BodyMd)
	writeString(prefix+"/Category", post_res.Category)
	writeString(prefix+"/Message", post_res.Message)
	writeString(prefix+"/Name", post_res.Name)
	writeArray(prefix+"/Tags", post_res.Tags)
	writeBool(prefix+"/Wip", post_res.Wip)

	return nil
}

func writePosts(posts []esa.PostResponse) error {
	for _, post := range posts {
		if err := writePost(&post); err != nil {
			return err
		}
	}
	return nil
}
