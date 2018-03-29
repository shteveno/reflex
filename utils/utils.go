// A file that contains functions for reading text files
// and basic string manipulation
package utils

import(
    "fmt"
    "bufio"
    "log"
    "os"
    "strings"
)

// Takes as input a .txt "input" file that lists of all
// reconstructions, the actual Greek words, and the glosses.

func Init_maps(protos map[string]string, glosses map[string]string, fileName string) {
	file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        bundle := strings.Split(scanner.Text(), ";")
        switch len(bundle) {
        case 3:
            proto := bundle[0]
            greek := bundle[1]
            gloss := bundle[2]
            protos[greek] = proto
            glosses[greek] = gloss
            continue
        default:
            fmt.Println("Parsing error with ", scanner.Text())
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}


// Waits for user input to process and parse into a tree.
func Wait_user() string {
	buf := bufio.NewReader(os.Stdin)
    fmt.Print("> ")
    rule, err := buf.ReadBytes('\n')
    if err != nil {
        fmt.Println(err)
		return ""
    } else {
        return strings.TrimSpace(string(rule))
    }
}
