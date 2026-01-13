# 🐦 Canary

A lightweight, high-performance audio player library for use cases in the terminal. Plays file from the system's path or URLs pointing to audio files.

## Features

- 🎵 Lightweight and fast audio playback
- 🎚️ Volume and playback controls
- 🔧 Simple CLI interface
- 📦 Easy installation and setup
- 🌐 Cross-platform support

## Installation
## Prerequisites

Before installation, ensure you have the following dependencies installed:

- `mplayer` - audio playback engine
- `go` - programming language (1.16 or higher)

Then proceed with one of the installation methods below:

### Manually

```bash
make
```

The binary will be generated in the `bin/` folder.

### Automated

```bash
sudo ./installer.sh
```

This script calls `make` and installs the binary to `/usr/local/bin`.

## Quick Start

```bash
canary --help
canary play
canary volume up
canary stop
```

## License

MIT
