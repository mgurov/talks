package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"errors"
	"io/ioutil"
	"os"
)

// repos
// https://api.github.com/users/bolcom/repos

// Tags
// https://api.github.com/repos/bolcom/kibana/tags

// START OMIT
func getJson(url string, v interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if nil != err {
		return err
	}
	if os.Getenv("USER") != "" && os.Getenv("GITHUB_TOKEN") != "" {
		req.SetBasicAuth(os.Getenv("USER"), os.Getenv("GITHUB_TOKEN"))
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		responseBody, _ := ioutil.ReadAll(resp.Body)
		return errors.New(fmt.Sprintf(
			"Unexpected http response code %d url %s response %s",
			resp.StatusCode, url, responseBody))
	}
	return json.NewDecoder(resp.Body).Decode(v)
}

// END OMIT

// START OMIT2

type Repository struct {
	Name    string `json:"name"`
	TagsUrl string `json:"tags_url"`
}

func getRepos(user string) (list []Repository, err error) {
	reposUrl := fmt.Sprintf("https://api.github.com/users/%s/repos", user)
	err = getJson(reposUrl, &list);
	return
}

type Tag struct {
	Name string `json:"name"`
}

func getTagsForRepo(url string) (list []Tag, err error) {
	err = getJson(url, &list);
	return
}

// END OMIT2

// START OMIT3

func main() {
	repos, err := getRepos("coreos")
	if err != nil {
		log.Fatalln(err)
	}
	if len(repos) == 0 {
		log.Println("no repos for this user")
	}
	for _, each := range repos {
		if list, err := getTagsForRepo(each.TagsUrl); err != nil {
			log.Fatalln(each, err)
		} else {
			if len(list) > 0 {
				fmt.Println(each.Name)
				fmt.Println(list)
			}
		}
	}
}

// END OMIT3
