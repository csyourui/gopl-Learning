package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"time"
)

//https://api.github.com/users/csyourui/repos
const GithubApiURL = "https://api.github.com/"

type SearchResult []Repository

type Repository struct {
	Id        int
	Name      string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Language  string
	Private   bool
	Owner     User
}
type User struct {
	Login string
	Id    int
}

func SearchRepo(terms []string) (*SearchResult, error) {
	q := terms[0]
	resp, err := http.Get(GithubApiURL + "users/" + q + "/repos")
	if err != nil {
		return nil, err
	}
	//如果返回值不是http的200 OK则
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Search queary failed :%s", resp.Status)
	}
	var result SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, err
}
func (s SearchResult) Len() int           { return len(s) }
func (s SearchResult) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SearchResult) Less(i, j int) bool { return s[i].UpdatedAt.Unix() < s[j].UpdatedAt.Unix() }

func main() {
	result, err := SearchRepo(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	sort.Sort(*result)
	fmt.Printf("%d Repository: \n", len(*result))
	for _, repo := range *result {
		fmt.Printf("%-5d %-35s %v\n", repo.Id, repo.Name, repo.CreatedAt)
	}
}
