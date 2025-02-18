package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	gosxnotifier "github.com/deckarep/gosx-notifier"
)

const (
	checkInterval = 30 * time.Minute
	silksongURL   = "https://issilksongout.com/"
	githubAPI     = "https://api.github.com/repos/Araraura/IsSilksongOut/commits/master"
	appName       = "silksong-monitor"
)

var stateFile string

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Could not determine home directory:", err)
	}

	// Create path like ~/Library/Application Support/silksong-monitor/
	appDir := filepath.Join(homeDir, "Library", "Application Support", appName)
	if err := os.MkdirAll(appDir, 0755); err != nil {
		log.Fatal("Could not create application directory:", err)
	}

	stateFile = filepath.Join(appDir, "lastcommit.txt")
}

type CommitResponse struct {
	SHA string `json:"sha"`
}

func getLatestCommit() (string, error) {
	req, err := http.NewRequest("GET", githubAPI, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "Silksong-Monitor")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GitHub API returned status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var commit CommitResponse
	if err := json.Unmarshal(body, &commit); err != nil {
		return "", err
	}

	return commit.SHA, nil
}

func loadLastCommit() string {
	data, err := os.ReadFile(stateFile)
	if err != nil {
		return ""
	}
	return string(data)
}

func saveLastCommit(sha string) error {
	return os.WriteFile(stateFile, []byte(sha), 0644)
}

func sendNotification() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("could not get home directory: %v", err)
	}

	note := gosxnotifier.NewNotification("Silksong status page has been updated!")
	note.Title = "Silksong News"
	note.Sound = gosxnotifier.Default
	note.Link = silksongURL
	note.AppIcon = filepath.Join(homeDir, "Library", "Applications", "silksong-monitor", "logo.png")

	return note.Push()
}

func monitorRepository() {
	lastCommit := loadLastCommit()
	log.Printf("Starting monitor with last commit: %s", lastCommit)

	ticker := time.NewTicker(checkInterval)
	defer ticker.Stop()

	for {
		currentCommit, err := getLatestCommit()
		if err != nil {
			log.Printf("Error checking latest commit: %v", err)
			time.Sleep(checkInterval)
			continue
		}

		if currentCommit != "" && currentCommit != lastCommit {
			if lastCommit != "" { // Don't notify on first run
				log.Printf("Repository updated! Old commit: %s, New commit: %s", lastCommit, currentCommit)

				if err := sendNotification(); err != nil {
					log.Printf("Error sending notification: %v", err)
				}
			}

			if err := saveLastCommit(currentCommit); err != nil {
				log.Printf("Error saving commit SHA: %v", err)
			}
			lastCommit = currentCommit
		}

		time.Sleep(checkInterval)
	}
}

func main() {
	monitorRepository()
}
