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
type Repository struct {
	Name    string `json:"name"`
	TagsUrl string `json:"tags_url"`
}

func getRepos(user string) (list []Repository, err error) {
	reposUrl := fmt.Sprintf("https://api.github.com/users/%s/repos", user)

	req, err := http.NewRequest("GET", reposUrl, nil)
	if nil != err {
		return
	}

	if os.Getenv("USER") != "" && os.Getenv("GITHUB_TOKEN") != "" {
		req.SetBasicAuth(os.Getenv("USER"), os.Getenv("GITHUB_TOKEN"))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		responseBody, _ := ioutil.ReadAll(resp.Body)
		err = errors.New(fmt.Sprintf("Unexpected http response code %d response %s", resp.StatusCode, responseBody))
		return
	}
	json.NewDecoder(resp.Body).Decode(&list)
	return
}

// END OMIT

// START OMIT2

type Tag struct {
	Name string `json:"name"`
}

func getTagsForRepo(url string) (list []Tag, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	json.NewDecoder(resp.Body).Decode(&list)
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
		log.Println("no repos for this user (or did you hit the rate limitter?)")
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
