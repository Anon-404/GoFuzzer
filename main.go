package main

import (
    "bufio"
    "fmt"
    "net/http"
    "net"
    "os"
    "sync"
    "time"

    "github.com/fatih/color"
)

var i = 0

// color blur
var blue = color.New(color.FgBlue)
var boldBlue = blue.Add(color.Bold)

// color red
var red = color.New(color.FgRed)
var boldRed = red.Add(color.Bold)

// color green
var Green = color.New(color.FgGreen)
var boldGreen = Green.Add(color.Bold)

// color yellow
var Yellow = color.New(color.FgYellow)
var boldYellow = Yellow.Add(color.Bold)

//color cyan
var cyan = color.New(color.FgCyan)
var boldCyan = cyan.Add(color.Bold)

//color magenta
var magenta = color.New(color.FgMagenta)
var boldMagenta = magenta.Add(color.Bold)

// color white
var white = color.New(color.FgWhite)
var boldWhite = white.Add(color.Bold)

var wg sync.WaitGroup
var mutex sync.Mutex

func banner() {
    banner := `
   ______      ____                         
  / ____/___  / __/_  __________  ___  _____
 / / __/ __ \/ /_/ / / /_  /_  / / _ \/ ___/
/ /_/ / /_/ / __/ /_/ / / /_/ /_/  __/ /    
\____/\____/_/  \__,_/ /___/___/\___/_/     
    `
    boldCyan.Println(banner)
    boldGreen.Println("William Steven(Anon404)\n")
}

func version() {
	boldGreen.Println("version 1.5")
	return
}

func help() {
    boldBlue.Println("╔════════════════════════════════════════════════════════════╗")
    boldBlue.Println("║                         GoFuzzer                          ║")
    boldBlue.Println("║            Web Directory Brute-Forcing Utility             ║")
    boldBlue.Println("╚════════════════════════════════════════════════════════════╝")
    boldBlue.Println()

    boldBlue.Println("Usage:")                                                                                                                                                       
    boldWhite.Println("  GoFuzzer <option> <domain> [wordlist]")
    boldBlue.Println()

    boldBlue.Println("Options:")                                                                                                                                                  
    boldGreen.Print("  -a ")
    boldBlue.Print(" --all          ")
    boldWhite.Println(": Brute-force all responses")                                                                                                                             
    boldGreen.Print("  -s ")
    boldBlue.Print(" --success      ")
    boldWhite.Println(": Brute-force 200 status responses")

    boldGreen.Print("  -r ")
    boldBlue.Print(" --redirect     ")
    boldWhite.Println(": Brute-force redirections (3xx)")                                                                                                                        
    boldGreen.Print("  -i ")
    boldBlue.Print(" --information  ")
    boldWhite.Println(": Brute-force informational responses")

    boldGreen.Print("  -d ")
    boldBlue.Print(" --debug        ")
    boldWhite.Println(": Brute-force server errors (5xx)")

    boldGreen.Print("  -h ")
    boldBlue.Print(" --help         ")
    boldWhite.Println(": Show this help manual")

    boldGreen.Print("  -v ")
    boldBlue.Print(" --version      ")
    boldWhite.Println(": Show version info")
    fmt.Println()

    boldBlue.Println("Examples:")
    boldWhite.Println("  GoFuzzer -a example.com              : For all responses on example.com")
    boldWhite.Println("  GoFuzzer -s example.com              : For successful responses (200)")
    boldWhite.Println("  GoFuzzer -r example.com              : For redirection responses (3xx)")
    boldWhite.Println("  GoFuzzer -d example.com              : For server errors (5xx)")
    boldWhite.Println("  GoFuzzer -a example.com mylist.txt   : Use 'mylist.txt' as the wordlist")
    boldBlue.Println()

    boldGreen.Println("Note: If no wordlist is provided, 'default.txt' will be used by default.")
    boldBlue.Println()
    boldBlue.Println("╔════════════════════════════════════════════════════════════╗")
    boldBlue.Println("║                    Thank you for using GoFuzzer!          ║")
    boldBlue.Println("╚════════════════════════════════════════════════════════════╝")
}

func secret(wg *sync.WaitGroup, domain string, sword string){
	defer wg.Done()
    	dom := "https://" + domain + "/" + sword
    	conn, err := http.Get(dom)
    	time.Sleep(500 * time.Millisecond)

    	if err != nil {
        	//fmt.Println(err)
        	return
    	}

    	defer conn.Body.Close()

    	if conn.StatusCode >= 200 && conn.StatusCode <= 399 {
        	boldCyan.Println("[!] Found :", dom)
        }	
}

func dbgbrute(wg *sync.WaitGroup, domain string, word string) {
    defer wg.Done()
    dom := "https://" + domain + "/" + word
    conn, err := http.Get(dom)
    time.Sleep(500 * time.Millisecond)

    if err != nil {
        //fmt.Println(err)
        return
    }

    defer conn.Body.Close()

    if conn.StatusCode >= 100 && conn.StatusCode <= 199 {
        boldYellow.Printf("%-40s status : %d\n", dom, conn.StatusCode)
		i++
        // 200-299
    } else if conn.StatusCode >= 200 && conn.StatusCode <= 299 {
    	boldGreen.Printf("%-40s status : %d\n", dom, conn.StatusCode)
		i++
        // 300-399
    } else if conn.StatusCode >= 300 && conn.StatusCode <= 399 {
    	boldBlue.Printf("%-40s status : %d\n", dom, conn.StatusCode)

		i++
        // 400-499
    } else if conn.StatusCode >= 400 && conn.StatusCode <= 499 {
        boldRed.Printf("%-40s status : %d\n", dom, conn.StatusCode)

        // 500-599
    } else if conn.StatusCode >= 500 && conn.StatusCode <= 599 {
        boldMagenta.Printf("%-40s status : %d\n", dom, conn.StatusCode)
        i++
    }
}



func allbrute(wg *sync.WaitGroup, domain string, word string) {
    defer wg.Done()
    dom := "https://" + domain + "/" + word
    conn, err := http.Get(dom)
    time.Sleep(500 * time.Millisecond)

    if err != nil {
        //fmt.Println(err)
        return
    }

    defer conn.Body.Close()

    if conn.StatusCode >= 100 && conn.StatusCode <= 199 {
    	boldYellow.Printf("%-40s status : %d\n", dom, conn.StatusCode)

		i++
	// 200-299
	} else if conn.StatusCode >= 200 && conn.StatusCode <= 299 {
    	boldGreen.Printf("%-40s status : %d\n", dom, conn.StatusCode)
		i++
	// 300-399
	} else if conn.StatusCode >= 300 && conn.StatusCode <= 399 {
    	boldBlue.Printf("%-40s status : %d\n", dom, conn.StatusCode)
		i++
	// 500-599
	} else if conn.StatusCode >= 500 && conn.StatusCode <= 599 {
    	boldRed.Printf("%-40s status : %d\n", dom, conn.StatusCode)
		i++
	}
}


func successbrute(wg *sync.WaitGroup, domain string, word string) {
    defer wg.Done()
    dom := "https://" + domain + "/" + word
    conn, err := http.Get(dom)
    time.Sleep(500 * time.Millisecond)

    if err != nil {
        //fmt.Println(err)
        return
    }

    defer conn.Body.Close()

    if conn.StatusCode >= 200 && conn.StatusCode <= 299 {
        boldGreen.Printf("%-40s status : %d\n", dom, conn.StatusCode)
	i++
	}
}



func redirectbrute(wg *sync.WaitGroup, domain string, word string) {
    defer wg.Done()
    dom := "https://" + domain + "/" + word
    conn, err := http.Get(dom)
    time.Sleep(500 * time.Millisecond)

    if err != nil {
        //fmt.Println(err)
        return
    }

    defer conn.Body.Close()

    if conn.StatusCode >= 300 && conn.StatusCode <= 399 {
        boldBlue.Printf("%-40s status : %d\n", dom, conn.StatusCode)
	i++
	}

}

func infobrute(wg *sync.WaitGroup, domain string, word string) {
    defer wg.Done()
    dom := "https://" + domain + "/" + word
    conn, err := http.Get(dom)
    time.Sleep(500 * time.Millisecond)

    if err != nil {
        //fmt.Println(err)
        return
    }

    defer conn.Body.Close()
    
        // 200-299
    if conn.StatusCode >= 100 && conn.StatusCode <= 199 {
    
        boldGreen.Printf("%-40s status : %d\n", dom, conn.StatusCode)

    i++
	}
}


func main() {
    banner()

    if len(os.Args) < 2 {
        boldBlue.Println("Usage: GoFuzzer <option> [domain]")
        boldBlue.Println("Manual: GoFuzzer -h")
        boldBlue.Println("Version: GoFuzzer -v")
        return
    }


    option := os.Args[1]
    if option == "-h" || option == "--help" {
        help()
        return
    } else if option == "-v" || option == "--version" {
    	version()
    	return
    }

    if len(os.Args) < 3 {
        boldRed.Println("Error: Domain missing")
        boldBlue.Println("Usage: GoFuzzer <option> <domain/ip>")
        boldBlue.Println("Manual: GoFuzzer -h")
        boldBlue.Println("Version: GoFuzzer -v")
        return
    }

    domain := os.Args[2]

    _, err := net.LookupIP(domain)
    if err != nil {
    	boldRed.Println("[-] Error :",err)
    	// fmt.Println("[-] check internet connection")
    	return
    }

    filename := "default.txt"
    wordlist := "default"

    if len(os.Args) > 3 {
    	filename = os.Args[3]
    	wordlist = os.Args[3]
    }

	boldBlue.Println("[+] target :",domain)
	boldBlue.Println("[+] wordlist :",wordlist)
    boldGreen.Println("[+] status : demo version\n")
    
    boldYellow.Println("[+] Looking for secret path/file")
    scret := []string{
		"robots.txt",      // Often lists directories or files meant to be hidden from search engines
		".htaccess",       // Apache configuration file
		".env",            // Environment variables file
		"admin",           // Admin interface directory
		"config.php",      // PHP configuration file
		"db-config.php",   // Database configuration file
		"backup",          // Backup files directory
		"includes",        // Includes directory for code
		"uploads",         // File uploads directory
		"temp",            // Temporary files directory
		".git",            // Git directory, exposing repository files
		".svn",            // SVN (Subversion) directory
		"install",         // Installation directory
		"logs",            // Log files directory
		"scripts",         // Scripts directory
		"cache",           // Cache files directory
		"private",         // Private or restricted directory
		"cgi-bin",         // CGI scripts directory
		"phpinfo.php",     // PHP information page
		"status",          // Status page or file
		"readme.txt",      // README file
		"error.log",       // Error log file
		"debug.log",       // Debugging log file
		"webdav",          // WebDAV directory
		"old",             // Old or legacy files directory
		"dump.sql",        // SQL database dump file
		"backup.sql",      // Backup of SQL database
		"user/.ssh",       // SSH configuration directory
		"settings.php",    // Settings file for applications
		"metadata",        // Metadata directory
		"sensitive-data",  // Sensitive data directory
		"config.json",     // JSON configuration file
		"database.yml",    // YAML database configuration file
		"appsettings.json",// Application settings file
		"web.config",      // IIS Web configuration file
		".bashrc",         // Bash configuration file
		".bash_profile",   // Bash profile file
		"site_backup.zip", // Compressed site backup
		".htpasswd",       // Password file for basic HTTP authentication
		"index.php.bak",   // Backup of index.php
		"wp-config.php",   // WordPress configuration file
		"local.xml",       // Magento configuration file
		"main.js.bak",     // Backup of main JavaScript file
		".DS_Store",       // macOS directory attribute file
		"composer.json",   // Composer dependency manager file
		"package-lock.json", // NPM package lock file
		"yarn.lock",       // Yarn lock file
		"swagger.json",    // Swagger API definition file
		"backup.tar.gz",   // Tarball backup file
		"license.txt",     // License file
		"CONTRIBUTING.md", // Contribution guidelines file
		"keys.json",       // API keys or credentials file
		"ftpconfig.json",  // FTP configuration file
		"error_log",
		"sitemap.xml",
	}

    for _, sword := range scret {
        wg.Add(1)
        go secret(&wg, domain, sword)
    }
    wg.Wait()

	boldYellow.Println("\n[+] Starting main process\n")



	file, err := os.Open(filename)
    if err != nil {
        boldRed.Println("[-] Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    // Implement a worker pool
    const numWorkers = 10
    semaphore := make(chan struct{}, numWorkers)

    processWord := func(word string) {
        defer wg.Done()
        semaphore <- struct{}{}
        defer func() { <-semaphore }()

        switch option {
        case "-a", "--all":
            allbrute(&wg, domain, word)
        case "-s", "--success":
            successbrute(&wg, domain, word)
        case "-r", "--redirect":
            redirectbrute(&wg, domain, word)
        case "-i", "--information":
            infobrute(&wg, domain, word)
        case "-d", "--debug":
            dbgbrute(&wg, domain, word)
        default:
            red := color.New(color.FgRed)
            boldRed := red.Add(color.Bold)
            boldRed.Println("[-] Unknown option", option)
            return
        }
    }

    for scanner.Scan() {
        word := scanner.Text()
        wg.Add(1)
        go processWord(word)
    }
    wg.Wait()

    if err := scanner.Err(); err != nil {
        boldRed.Println("[-] Error reading file:", err)
    }

    boldBlue.Println("[+] Total found :", i)
}