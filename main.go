package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
)

// GitHubUser represents a GitHub user's profile information.
type GitHubUser struct {
	Login       string `json:"login"`
	Name        string `json:"name"`
	Bio         string `json:"bio"`
	PublicRepos int    `json:"public_repos"`
}

// Repository represents a GitHub repository.
type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	HTMLURL     string `json:"html_url"`
	Language    string `json:"language"`
	Stars       int    `json:"stargazers_count"`
}

// fetchJSON makes an HTTP GET request to the given URL and decodes the JSON response into v.
func fetchJSON(url string, v interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received status code %d", resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(v)
}

// printUser displays the user's profile information in a structured format.
func printUser(user *GitHubUser) {
	fmt.Println("Profile Information:")
	fmt.Printf("- Username: %s\n", user.Login)
	fmt.Printf("- Name: %s\n", user.Name)
	fmt.Printf("- Bio: %s\n", user.Bio)
	fmt.Printf("- Public Repos: %d\n", user.PublicRepos)
	fmt.Println()
}

// printRepos displays the list of repositories, sorted alphabetically by name.
func printRepos(repos []Repository) {
	if len(repos) == 0 {
		fmt.Println("No public repositories found.")
		return
	}
	sort.Slice(repos, func(i, j int) bool {
		return repos[i].Name < repos[j].Name
	})
	fmt.Println("Repositories:")
	for _, repo := range repos {
		fmt.Printf("- %s (%s) ★ %d\n  %s\n", repo.Name, repo.Language, repo.Stars, repo.HTMLURL)
	}
}

func main() {
	var username string
	fmt.Print("Enter GitHub username: ")
	fmt.Scanln(&username)
	if username == "" {
		log.Fatal("Username cannot be empty")
	}

	profileURL := fmt.Sprintf("https://api.github.com/users/%s", username)
	reposURL := fmt.Sprintf("https://api.github.com/users/%s/repos", username)

	user := &GitHubUser{}
	if err := fetchJSON(profileURL, user); err != nil {
		log.Fatalf("Error fetching profile: %v", err)
	}

	repos := []Repository{}
	if err := fetchJSON(reposURL, &repos); err != nil {
		log.Fatalf("Error fetching repos: %v", err)
	}

	printUser(user)
	printRepos(repos)
}
