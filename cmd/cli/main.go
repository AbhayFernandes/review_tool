package main

import "flag"
import "fmt"

func main() {
    diffFile := flag.String("diff", "", "Path to the git diff file")
    flag.Parse()

    if *diffFile == "" {
        fmt.Println("Error: --diff flag is required")
        return
    }

    fmt.Printf("Uploading diff from %s...", *diffFile)
}

