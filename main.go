package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	bootstrapFlag bool
)

func init() {
	flag.BoolVar(&bootstrapFlag, "bootstrap", true, "use bootstrap css?")
}

func main() {
	hp, root := func() (hp, root string) {
		lFlag := flag.Bool("l", false, "only use localhost")
		dFlag := flag.String("d", ".", "root dir path")
		root = *dFlag
		flag.Parse()
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		if *lFlag {
			hp = "localhost:" + port
		} else {
			hp = ":" + port
		}
		args := flag.Args()
		for i, arg := range args {
			switch i {
			case 0:
				if arg != "" {
					hp = arg
				}
			case 1:
				root = arg
			}
		}
		if !strings.Contains(hp, ":") {
			hp = ":" + hp
		}
		return
	}()

	info, err := os.Stat(root)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !info.IsDir() {
		fmt.Println(root, "is not dir!")
		return
	}
	fmt.Println("Http Static File Server")
	fmt.Println("Please access", hp)

	var handler http.Handler
	handler = http.FileServer(http.Dir(root))
	if bootstrapFlag {
		handler = FileServer(http.Dir(root))
	}
	if err = http.ListenAndServe(hp, handler); err != nil {
		log.Printf("Server Error: %v", err)
	}
}
