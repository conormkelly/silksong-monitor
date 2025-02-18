BINARY_NAME=silksong-monitor
BUILD_DIR=build
INSTALL_DIR=~/Library/Applications/$(BINARY_NAME)
SUPPORT_DIR=~/Library/Application\ Support/$(BINARY_NAME)
AGENTS_DIR=~/Library/LaunchAgents
PLIST_FILE=com.conormkelly.silksong-monitor.plist

.PHONY: all build clean install uninstall test check-deps install-deps

all: build

check-deps:
	@which terminal-notifier > /dev/null || (echo "Installing terminal-notifier..." && brew install terminal-notifier)

build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME)

clean:
	@rm -rf $(BUILD_DIR)

install: check-deps build
	@echo "Creating application directories..."
	@mkdir -p $(INSTALL_DIR)
	@mkdir -p $(SUPPORT_DIR)
	@mkdir -p $(AGENTS_DIR)
	
	@echo "Installing binary and assets..."
	@cp $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_DIR)/
	@cp logo.png $(INSTALL_DIR)/ 2>/dev/null || (echo "Warning: logo.png not found" && exit 0)
	
	@echo "Installing and loading launch agent..."
	@sed "s|HOMEDIR|$$HOME|g" $(PLIST_FILE) > $(AGENTS_DIR)/$(PLIST_FILE)
	@launchctl unload $(AGENTS_DIR)/$(PLIST_FILE) 2>/dev/null || true
	@launchctl load $(AGENTS_DIR)/$(PLIST_FILE)
	
	@echo "Installation complete. The monitor will start automatically on next login."
	@echo "Logs will be available at: ~/Library/Logs/silksong-monitor.log"

uninstall:
	@echo "Stopping and removing launch agent..."
	@launchctl unload $(AGENTS_DIR)/$(PLIST_FILE) 2>/dev/null || true
	@rm -f $(AGENTS_DIR)/$(PLIST_FILE)
	
	@echo "Removing application files..."
	@rm -rf $(INSTALL_DIR)
	@echo "Note: Application data in $(SUPPORT_DIR) has been preserved."
	@echo "To remove it completely, run: rm -rf $(SUPPORT_DIR)"
	
	@echo "Uninstallation complete"

test:
	@launchctl unload ~/Library/LaunchAgents/com.conormkelly.silksong-monitor.plist
	@echo "testcommit123" > ~/Library/Application\ Support/silksong-monitor/lastcommit.txt
	@launchctl load ~/Library/LaunchAgents/com.conormkelly.silksong-monitor.plist
