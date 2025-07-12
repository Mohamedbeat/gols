# lo - A Colorful File Listing Utility

[![Go Version](https://img.shields.io/badge/Go-1.24.5+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

`gols` is a modern, colorful alternative to the traditional `ls` command, written in Go. It provides enhanced file listing with color-coded output, human-readable file sizes, and interactive features for large directories.

## ✨ Features

- **Color-coded output**: Different colors for directories, files, and hidden items
- **Human-readable file sizes**: Automatically formats sizes in KB, MB, GB, etc.
- **Interactive mode**: Prompts for confirmation when listing large directories (>100 items)
- **Hidden file support**: Option to show/hide hidden files
- **Full path display**: Option to show full file paths
- **Detailed information**: Shows permissions, size, name, and modification time
- **Cross-platform**: Works on Linux, macOS, and Windows

## 🚀 Installation

### Prerequisites

- Go 1.24.5 or higher

### Build from source

```bash
# Clone the repository
git clone https://github.com/mohamedbeat/gols.git
cd gols

# Build the binary
go build -o gols main.go

# Install to your system (optional)
sudo cp lo /usr/local/bin/
```

### Using go install

```bash
go install github.com/mohamedbeat/gols@latest
```

## 📖 Usage

### Basic Usage

```bash
# List files in current directory
gols

# List files with hidden files
gols -a

# List files with full paths
gols -p

# Combine flags
gols -a -p
```

### Command Line Options

| Flag | Description |
|------|-------------|
| `-a` | Show hidden files (files starting with `.`) |
| `-p` | Show full path for each file |

### Output Format

The output displays files in a tabular format with the following columns:

1. **Permissions** - File mode and permissions
2. **Size** - Human-readable file size
3. **Name** - File/directory name (color-coded)
4. **Modification Time** - Last modified date and time

### Color Scheme

- 🔵 **Blue**: Regular directories
- 🔴 **Red**: Hidden directories
- 🟢 **Green**: Regular files
- 🟡 **Yellow**: Hidden files

## 🎯 Examples

### Basic file listing
```bash
$ gols
drwxr-xr-x    4.0 KB  Documents    15-01-2024 10:30:45
-rw-r--r--    2.1 KB  README.md    14-01-2024 16:20:12
-rw-r--r--  156.3 KB  image.png    13-01-2024 09:15:30
```

### Showing hidden files
```bash
$ gols -a
drwxr-xr-x    4.0 KB  .git         15-01-2024 10:30:45
-rw-r--r--    2.1 KB  .gitignore   14-01-2024 16:20:12
drwxr-xr-x    4.0 KB  Documents    15-01-2024 10:30:45
-rw-r--r--    2.1 KB  README.md    14-01-2024 16:20:12
```

### Large directory handling
```bash
$ lo
show 150 results? y/n y
# Lists all files with confirmation
```

## 🛠️ Development

### Project Structure

```
lo/
├── main.go      # Main application code
├── go.mod       # Go module definition
├── go.sum       # Dependency checksums
└── README.md    # This file
```

### Dependencies

- `github.com/fatih/color` - For colored terminal output

### Building

```bash
# Build for current platform
go build -o gols main.go

# Build for specific platform
GOOS=linux GOARCH=amd64 go build -o lo main.go
```

## 🤝 Contributing

Contributions are welcome! 
## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Inspired by the traditional `ls` command
- Uses [fatih/color](https://github.com/fatih/color) for beautiful terminal colors
- Built with Go for cross-platform compatibility

## 📊 Project Status

- ✅ Core functionality implemented
- ✅ Color-coded output
- ✅ Interactive mode for large directories
- ✅ Hidden file support
- ✅ Full path option
- 🔄 Future enhancements planned

---

**Made with ❤️ using Go** 