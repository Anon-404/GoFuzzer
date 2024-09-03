# GoFuzzer 🔎
**Web directory, file and path finder tool**  
*Written in Go*

<div align="center">
  <img src="https://raw.githubusercontent.com/Anon-404/My-assets/main/GoMapper/GoMapper.jpg" alt="GoMapper Logo" width="200"/>
</div>

## Features ✨
- 🔍 **Network Scanning and OS Detection**
- ⚡ **Faster Scanning:** Scans 65,535 ports quicker than Nmap
- 🌐 **DNS Lookup**
- 🕵️‍♂️ **Whois Lookup**
- 🎯 **User-Friendly Interface**
- 📖 **Comprehensive User Manual**

## Installation 🛠️

### Method 1: Using `go install`

#### Step 1: Install Go
- **Arch-based Linux:**
  ```bash
  sudo pacman -S go
  ```
- **Debian-based Linux:**
  ```bash
  sudo apt install golang -y
  ```
- **Fedora:**
  ```bash
  sudo dnf install golang
  ```
- **Termux:**
  ```bash
  pkg install golang -y
  ```

- **OpenSUSE:**
  ```bash
  sudo zypper install go
  ```
- **Void Linux:**
  ```bash
  sudo xbps-install go
  ```

#### Step 2: Install GoFuzzer 
- **For Linux:**
  ```bash
  go install -v github.com/Anon-404/GoFuzzer@latest
  sudo cp $HOME/go/bin/GoFuzzer /usr/bin
  ```
- **For Termux:**
  ```bash
  go install -v github.com/Anon-404/GoFuzzer@latest
  cp $HOME/go/bin/GoFuzzer ../usr/bin
  ```

### Method 2: Cloning the Repository

#### Step 1: Clone and Build
- **For Linux:**
  ```bash
  git clone https://github.com/Anon-404/GoFuzzer 
  cd GoFuzzer 
  go build -o GoFuzzer 
  sudo cp GoFuzzer /usr/bin
  ```
- **For Termux:**
  ```bash
  git clone https://github.com/Anon-404/GoFuzzer 
  cd GoFuzzer 
  go build -o GoFuzzer 
  cp GoFuzzer ../../usr/bin
  ```

## Usage 🧑‍💻

```bash
GoMapper <option> <domain/ip>
```

### Main Options:
- **`-a`, `--all [domain]`**  
  🔗 **Perform all actions**  
  Executes network scan, DNS lookup, and WHOIS lookup in a single command.

- **`-n`, `--networkScan [domain]`**  
  🌐 **Network Scan**  
  Performs a thorough network scan, including IP, port scanning, and OS detection.

- **`-d`, `--dnslookup [domain]`**  
  🛠 **DNS Lookup**  
  Retrieves detailed DNS records for the specified domain.

- **`-w`, `--whoislookup [domain]`**  
  🔍 **Whois Lookup**  
  Gathers WHOIS registration data for the specified domain.

### Additional Options:
- **`-h`, `--help`**  
  📝 **Help**  
  Displays this help page with descriptions of all available commands.

- **`-v`, `--version`**  
  🆚 **Version**  
  Prints the current version number of GoMapper.


## Contributions 🤝
Contributions are welcome! Feel free to open issues or submit pull requests.
