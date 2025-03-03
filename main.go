package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type GitHubUser struct {
	Login       string `json:"login"`
	Name        string `json:"name"`
	Bio         string `json:"bio"`
	PublicRepos int    `json:"public_repos"`
}

type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	HTMLURL     string `json:"html_url"`
	Language    string `json:"language"`
	Stars       int    `json:"stargazers_count"`
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

	user, err := fetchUser(profileURL)
	if err != nil {
		log.Fatalf("Error fetching profile: %v", err)
	}
	fmt.Printf("Username: %s\nName: %s\nBio: %s\nPublic Repos: %d\n\n",
		user.Login, user.Name, user.Bio, user.PublicRepos)

	repos, err := fetchRepos(reposURL)
	if err != nil {
		log.Fatalf("Error fetching repos: %v", err)
	}

	fmt.Println("Repositories:")
	for _, repo := range repos {
		fmt.Printf("- %s (%s) ★ %d\n  %s\n\n", repo.Name, repo.Language, repo.Stars, repo.HTMLURL)
	}
}

func fetchUser(url string) (*GitHubUser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received status code %d", resp.StatusCode)
	}
	var user GitHubUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func fetchRepos(url string) ([]Repository, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received status code %d", resp.StatusCode)
	}
	var repos []Repository
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return nil, err
	}
	return repos, nil
}
