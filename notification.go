package main

import (
	"errors"
	"os/exec"
)

func findTerminalNotifier() string {
	terminalNotifierPaths := []string{
		"/usr/local/bin/terminal-notifier",
		"/opt/homebrew/bin/terminal-notifier",
	}

	for _, path := range terminalNotifierPaths {
		if _, err := exec.LookPath(path); err == nil {
			return path
		}
	}
	return ""
}

func sendNotification(title, message, icon, url string) error {
	notifierPath := findTerminalNotifier()
	if notifierPath == "" {
		return errors.New("terminal-notifier not found in expected paths")
	}

	cmd := exec.Command(notifierPath,
		"-title", title,
		"-message", message,
		"-contentImage", icon,
		"-open", url,
		"-appIcon", icon,
		"-sound", "default",
		"-group", "com.conormkelly.silksong-monitor")

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
