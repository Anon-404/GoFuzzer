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

var wg sync.WaitGroup

func banner() {
    banner := `
   ________             __    ___       __  __
  / ___/ _ \___  _____ / /__ / _ \___ _/ /_/ /
 / /__/ // / _ \/  /__/   _// ___/ _  / __/ _ \
 \___/____/\_,_/_ /  /_/\_\/_/   \_,_/\__/_//_/
    `
    boldCyan.Println(banner)
    boldGreen.Println("William Steven(Anon404)\n")
}

func version() {
	boldGreen.Println("version 1.5")
	return
}

func help() {
	fmt.Println("comming soon")
	return
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

		i++
        // 500-599
        } else if conn.StatusCode >= 500 && conn.StatusCode <= 599 {
        boldMagenta.Printf("%-40s status : %d\n", dom, conn.StatusCode)
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
        boldBlue.Println("Usage: GDarkPath <option> [domain]")
        boldBlue.Println("Manual: GDarkPath -h")
        boldBlue.Println("Version: GDarkPath -v")
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
        boldBlue.Println("Usage: GoDarkPath <option> <domain/ip>")
        boldBlue.Println("Manual: GoDarkPath -h")
        boldBlue.Println("Version: GoDarkPath -v")
        return
    }

    domain := os.Args[2]

    _, err := net.LookupIP(domain)
    if err != nil {
    	boldRed.Println("[-] Error :",err)
    	// fmt.Println("[-] check internet connection")
    	return
    }

	file, err := os.Open("dirs.txt")
    if err != nil {
       	fmt.Println("Error opening file:", err)
       	return
   	}

    scanner := bufio.NewScanner(file)

    if option == "-a" || option == "--all" {

    	for scanner.Scan() {
    		word := scanner.Text()
    		wg.Add(1)
    		go allbrute(&wg, domain, word)
    	}
    	wg.Wait()
    
	} else if option == "-s" || option == "--success" {

		for scanner.Scan() {
			word := scanner.Text()
			wg.Add(1)
			go successbrute(&wg, domain, word)
		}
		wg.Wait()
		boldBlue.Println("[+] Total found :",i)

	} else if option == "-r" || option == "--redirect" {

        for scanner.Scan() {
         	word := scanner.Text()
            wg.Add(1)
			go redirectbrute(&wg, domain, word)
        }
        wg.Wait()
		//boldBlue.Println("[+] Total found :",i)
	} else if option == "-i" || option == "--information" {

        for scanner.Scan() {
            word := scanner.Text()
            wg.Add(1)
			go infobrute(&wg, domain, word)
        }
        wg.Wait()
		//boldBlue.Println("[+] Total found :",i)


	} else if option == "-d" || option == "--debug" {

        for scanner.Scan() {
            word := scanner.Text()
            wg.Add(1)
            go dbgbrute(&wg, domain, word)
        }
        wg.Wait()
                //boldBlue.Println("[+] Total found :",i)
	} else {
    red := color.New(color.FgRed)
	    boldRed := red.Add(color.Bold)
    	boldRed.Println("[-] Unknown option", option)
    	return
	}
    
}
