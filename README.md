# Silksong Monitor

A simple macOS application that monitors the [issilksongout.com](https://issilksongout.com) website for updates and sends you a notification when changes are detected.

## Features

- Checks for updates every 30 minutes
- Runs silently in the background
- Starts automatically on login
- Clickable notifications that take you directly to the website

## Installation

```bash
git clone https://github.com/conormkelly/silksong-monitor
cd silksong-monitor
make install
```

The application will be installed to:

- Binary and assets: `~/Library/Applications/silksong-monitor/`
- Application data: `~/Library/Application Support/silksong-monitor/`
- Logs: `~/Library/Logs/silksong-monitor.log`

To uninstall:

```bash
make uninstall
```

## How it works

The application monitors the GitHub repository behind IsSilksongOut for any changes. When a change is detected, it sends a macOS notification that you can click to visit the website.

## Logs

You can view the application logs at any time:

```bash
tail -f ~/Library/Logs/silksong-monitor.log
```

## Credits

- Application icon by [AndonovMarko](https://www.deviantart.com/andonovmarko/art/Hollow-Knight-Silksong-Icon-804805724) ([gallery](https://www.deviantart.com/andonovmarko/gallery))
- IsSilksongOut website by [Araraura](https://github.com/Araraura/IsSilksongOut)
- Hollow Knight and Silksong are created by Team Cherry

## License

Feel free to use, modify, and distribute this code. Just remember to credit the original icon creator if you use the included logo.
