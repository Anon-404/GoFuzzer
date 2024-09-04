# GoFuzzer 🔎
**Web directory, file and path finder tool**  
*Written in Go*

<div align="center">
  <img src="https://raw.githubusercontent.com/Anon-404/My-assets/main/GoFuzzer/GoFuzzer.jpg" alt="GoFuzzer Logo" width="200"/>
</div>

## Features ✨

- 🔍 **Web Directory Brute-Forcing**  
  Performs targeted brute-forcing to discover hidden directories and files on a web server.

- ⚡ **Efficient Scanning:**  
  Utilizes concurrent requests to speed up the scanning process, allowing for faster discovery of paths and files.

- 🌐 **Customizable Options:**  
  Choose between different scanning modes such as showing all responses, success responses, redirections, or informational responses.

- 🕵️‍♂️ **Detailed Status Reporting:**  
  Provides comprehensive feedback on server responses with color-coded status codes for easy identification.

- 🎯 **User-Friendly Interface:**  
  Includes a clear and concise command-line interface with helpful usage instructions and examples.

- 📖 **Comprehensive User Manual:**  
  Well-documented help and version information to assist users in navigating and utilizing the tool effectively.

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
GoFuzzer <option> <domain> <wordlist>
```

### Main Options:

- **`-a`, `--all [domain]`**  
  🔗 **Show All Responses**  
  Displays all responses from the server, including status codes ranging from 100 to 599.

- **`-s`, `--success [domain]`**  
  ✅ **Show Success Responses**  
  Filters and displays only successful responses with status codes in the range of 200 to 299.

- **`-r`, `--redirect [domain]`**  
  🔄 **Show Redirections**  
  Shows responses that indicate redirections with status codes between 300 and 399.

- **`-i`, `--information [domain]`**  
  ℹ️ **Show Informational Responses**  
  Displays responses with informational status codes, ranging from 100 to 199.

- **`-d`, `--debug [domain]`**  
  🛠 **Debug Mode**  
  Enables debugging mode to show detailed responses for all status codes from 100 to 599.

### Additional Options:

- **`-h`, `--help`**  
  📝 **Help**  
  Displays this help page with descriptions of all available commands and options.

- **`-v`, `--version`**  
  🆚 **Version**  
  Prints the current version number of GoFuzzer.

### Examples:

- **Show all responses:**
  ```bash
  GoFuzzer -a example.com wordlist.txt
  ```

- **Show only successful responses:**
  ```bash
  GoFuzzer -s example.com wordlist.txt
  ```

- **Show only redirections:**
  ```bash
  GoFuzzer -r example.com wordlist.txt
  ```

- **Enable debugging mode:**
  ```bash
  GoFuzzer -d example.com wordlist.txt
  ```

For more information or assistance, use the `-h` option to display help page.


## Contributions 🤝
Contributions are welcome! Feel free to open issues or submit pull requests.
