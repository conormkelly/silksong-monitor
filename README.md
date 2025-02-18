# Silksong Monitor

A simple macOS application that monitors the [issilksongout.com](https://issilksongout.com) website for updates and sends you a notification when changes are detected.

## Features

- Checks for updates every 30 minutes
- Runs silently in the background
- Starts automatically on login
- Native macOS notifications with URL support
- Clickable notifications that take you directly to the website
- Custom notification sound and icon

## Requirements

- macOS
- Go 1.24 or later
- Homebrew (for installing terminal-notifier)

## Installation

```bash
git clone https://github.com/conormkelly/silksong-monitor
cd silksong-monitor
make install
```

The installer will automatically:

- Install terminal-notifier if needed
- Set up the application directories
- Configure the launch agent
- Start the monitor

The application will be installed to:

- Binary and assets: `~/Library/Applications/silksong-monitor/`
- Application data: `~/Library/Application Support/silksong-monitor/`
- Logs: `~/Library/Logs/silksong-monitor.log`

To uninstall:

```bash
make uninstall
```

## How it works

The application monitors the GitHub repository behind IsSilksongOut for any changes. When a change is detected, it sends a native macOS notification using terminal-notifier. Clicking the notification will take you directly to the website.

## Testing

To simulate a notification (useful for testing):

```bash
make test
```

## Logs

You can view the application logs at any time:

```bash
tail -f ~/Library/Logs/silksong-monitor.log
```

## Credits

- Application icon by [AndonovMarko](https://www.deviantart.com/andonovmarko/art/Hollow-Knight-Silksong-Icon-804805724) ([gallery](https://www.deviantart.com/andonovmarko/gallery))
- IsSilksongOut website by [Araraura](https://github.com/Araraura/IsSilksongOut)
- Notifications powered by [terminal-notifier](https://github.com/julienXX/terminal-notifier)
- Hollow Knight and Silksong are created by Team Cherry

## License

Feel free to use, modify, and distribute this code. Just remember to credit the original icon creator if you use the included logo.
