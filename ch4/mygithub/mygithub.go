package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"text/template"
	"time"
)

//GithubAPIURLgo https://api.github.com/users/csyourui/repos
const GithubAPIURLgo = "https://api.github.com/"

//PublicReposResult 请求返回json对象
type PublicReposResult []Repository

//Repository 请求返回json对象中的repos对象
type Repository struct {
	ID        int
	Name      string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Private   bool
	Owner     User
	HTMLURL   string `json:"html_url"`
	Language  string
}

//User 请求返回json格式中User对象
type User struct {
	Login string
	ID    int
}

//SearchPublicRepo 接受用户名参数返回PublicReposResult对象
func SearchPublicRepo(userName string) (*PublicReposResult, error) {
	resp, err := http.Get(GithubAPIURLgo + "users" + userName + "/repos")
	if err != nil {
		return nil, err
	}
	//如果返回值不是http的200 OK则
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Search queary failed :%s", resp.Status)
	}
	//调用json.Decode解析resp.Body
	var result PublicReposResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, err
}

const templ = `{{. | getTotolCount}} Repositories:
{{range .}}----------------------------------------
ID :		{{.ID}}
Name :		{{.Name}}
Language :	{{.Language}}
User :		{{.Owner.Login}}
Age: 		{{.CreatedAt | daysAgo}} days
Update: 	{{.UpdatedAt | daysAgo}} days
{{end}}`

const htmltempl = `
<h1>{{. | getTotolCount}} Repositories</h1>
<table>
	<tr style='text-align: left'>
		<th>ID</th>
		<th>Name</th>
		<th>Language</th>
		<th>User</th>
		<th>Age</th>
		<th>Update</th>
	</tr>
	{{range .}}
		<tr>
			<td>{{.ID}}</td>
			<td><a href='{{.HTMLURL}}'>{{.Name}}</a></td>
			<td>{{.Language}}</td>
			<td>{{.Owner.Login}}</td>
			<td>{{.CreatedAt | daysAgo}}</td>
			<td>{{.UpdatedAt | daysAgo}}</td>
		</tr>
	{{end}}
</table>`

//重写sort的接口
func (s PublicReposResult) Len() int           { return len(s) }
func (s PublicReposResult) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s PublicReposResult) Less(i, j int) bool { return s[i].CreatedAt.Unix() < s[j].CreatedAt.Unix() }

func getTotolCount(result *PublicReposResult) int {
	return len(*result)
}
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("repos").
	Funcs(template.FuncMap{"getTotolCount": getTotolCount}).
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(htmltempl))

func printTemplate(result *PublicReposResult) {
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	result, err := SearchPublicRepo(r.URL.Path)
	if err != nil {
		log.Fatal(err)
	}
	sort.Sort(*result)

	if err := report.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}
func serverHelper() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func main() {

	// fmt.Printf("%d Repository: \n", len(*result))
	// for _, repo := range *result {
	// 	fmt.Printf("%-10d %-35s %-10s %v\n", repo.ID, repo.Name, repo.Language, repo.CreatedAt)
	// }
	serverHelper()
}
