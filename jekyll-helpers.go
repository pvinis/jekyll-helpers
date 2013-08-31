package main

import (
	"fmt"
	"os"
	"strings"
)

func print_help() {
	fmt.Println("Commands:")
	fmt.Println(" n:\t<name>")
	fmt.Println("   \tcreates new draft. n test-draft will create a test-draft.markdown file inside _drafts")
	fmt.Println(" p:\tcreate post from draft")
	fmt.Println(" h:\tshow this help")
}

func main() {

	args := os.Args
	switch args[1] {
	case "h", "help":
		print_help()
	default:
		print_help()
	case "n":
		if len(args) < 3 {
			fmt.Println("forgot to give a name")
			break
		}

		os.Chdir("_drafts")
		pwd, _ := os.Getwd()
		if strings.Contains(pwd, "_drafts") {
			f, err := os.Create(args[2])
			if err != nil {
				panic(err)
			}
			defer func() {
				err := f.Close()
				if err != nil {
					panic(err)
				}
			}()
			fmt.Fprintln(f, "---");
			fmt.Fprintln(f, "layout: post");
			fmt.Fprintln(f, "title:", );
			fmt.Fprintln(f, "---");
			fmt.Fprintln(f, "---");

/// file here
		} else {
			fmt.Println("cant find the _drafts dir")
		}
///open it with vim
	case "p":
	}
}

