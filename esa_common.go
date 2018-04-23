package main

import (
  "os"
  "fmt"
  "strings"
  "strconv"
  "io/ioutil"
  "net/url"
  "path/filepath"
  "github.com/upamune/go-esa/esa"
)

func get_post(client *esa.Client, team string, id int) *esa.PostResponse {
  post_res, err := client.Post.GetPost(team, id)
  if err != nil {
    panic("client.Post.GetPost Failed: " + err.Error())
  }
  return post_res
}

func get_posts(client *esa.Client, team string, user string) []esa.PostResponse {
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

func read_array(path string) []string {
  data, err := ioutil.ReadFile(path)
  if err != nil {
    panic("ioutil.ReadFile Failed: " + err.Error())
  }
  return strings.Split(string(data), "\n")
}

func read_bool(path string) bool {
  data, err := ioutil.ReadFile(path)
  if err != nil {
    panic("ioutil.ReadFile Failed: " + err.Error())
  }

  boolean, err:= strconv.ParseBool(string(data))
  if err != nil {
    panic("strconv.ParseBool Failed: " + err.Error())
  }
  return boolean
}

func read_string(path string) string {
  data, err := ioutil.ReadFile(path)
  if err != nil {
    panic("ioutil.ReadFile Failed: " + err.Error())
  }
  return string(data)
}

func update_post(client *esa.Client, team string, path string) {
  post_id, err := strconv.Atoi(filepath.Base(path))
  if err != nil {
    panic("strconv.Atoi Failed: " + err.Error())
  }

  post := esa.Post {
    BodyMd:   read_string(path + "/BodyMd"),
    Category: read_string(path + "/Category"),
    Message:  read_string(path + "/Message"),
    Name:     read_string(path + "/Name"),
    Tags:     read_array(path + "/Tags"),
    Wip:      read_bool(path + "/Wip"),
  }

  if _, err := client.Post.Update(team, post_id, post); err != nil {
    panic("client.Post.Update Failed: " + err.Error())
  }
}

func new_post(client *esa.Client, team string, category string, name string) *esa.PostResponse {
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

func write_bool(path string, data bool) {
  if err := ioutil.WriteFile(path, []byte(fmt.Sprint(data)), 0644); err != nil {
    panic("ioutil.WriteFile Failed: " + path + err.Error())
  }
}

func write_array(path string, data []string) {
  if err := ioutil.WriteFile(path, []byte(strings.Join(data, "\n")), 0644); err != nil {
    panic("ioutil.WriteFile Failed: " + path + err.Error())
  }
}

func write_string(path string, data string) {
  if err := ioutil.WriteFile(path, []byte(data), 0644); err != nil {
    panic("ioutil.WriteFile Failed: " + path + err.Error())
  }
}

func write_post(post_res *esa.PostResponse) {
  prefix := fmt.Sprint(post_res.Number)

  if err := os.MkdirAll(prefix, 0755); err != nil {
    panic("os.Mkdir Failed: " + err.Error())
  }

  write_string(prefix + "/BodyMd", post_res.BodyMd)
  write_string(prefix + "/Category", post_res.Category)
  write_string(prefix + "/Message", post_res.Message)
  write_string(prefix + "/Name", post_res.Name)
  write_array(prefix + "/Tags", post_res.Tags)
  write_bool(prefix + "/Wip", post_res.Wip)
}

func write_posts(posts []esa.PostResponse) {
  for _, post := range posts {
    write_post(&post)
  }
}
