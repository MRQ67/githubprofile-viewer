
# GitHub Profile Viewer CLI

A simple command-line tool to fetch and display a GitHub user's profile information and list of repositories.

## Installation

1. Clone the repository:  
  
   ```bash  
   git clone https://github.com/MRQ67/github-profile-viewer.git  
   ```  

2. Navigate to the project directory:  
  
   ```bash  
   cd github-profile-viewer  
   ```  

## Usage

Run the tool and enter a GitHub username when prompted:  

```bash  
go run main.go
```  

You will be asked to input a GitHub username. After entering the username, the tool will display the user's profile information and a list of their repositories.

## Displayed Information

- **Profile Information**:  
  - Username  
  - Name  
  - Bio  
  - Followers  
  - Following  
- **Repositories**:  
  - Repository Name  
  - Description  

## Requirements

- Go 1.16 or higher  

## Example

```
$ go run main.go
Enter GitHub username: MRQ67
Profile Information:
- Username: MRQ67
- Name: Abdellah Qadi
- Bio: Building impactful solutions with code. And a healthy dose of late-night debugging.
- Public Repos: 9

Repositories:
- MRQ67 () ★ 1
  https://github.com/MRQ67/MRQ67
- My_Portfolio_Website () ★ 0
  https://github.com/MRQ67/My_Portfolio_Website
- blockchain (Python) ★ 2
  https://github.com/MRQ67/blockchain
- gif_generator (JavaScript) ★ 1
  https://github.com/MRQ67/gif_generator
- githubprofile-viewer (Go) ★ 0
  https://github.com/MRQ67/githubprofile-viewer
- java101 (Java) ★ 2
  https://github.com/MRQ67/java101
- my-weather-cli (Go) ★ 1
  https://github.com/MRQ67/my-weather-cli
- quotes_calls (Python) ★ 1
  https://github.com/MRQ67/quotes_calls
- url-shorten (Go) ★ 1
  https://github.com/MRQ67/url-shorten   
```

## How It Works

The tool uses the GitHub API to fetch the user's profile and repository information. It makes HTTP requests to the API endpoints and parses the JSON responses to display the relevant data.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
