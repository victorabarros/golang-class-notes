// https://www.hackerrank.com/challenges/ctci-ransom-note/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dictionaries-hashmaps
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the checkMagazine function below.
func checkMagazine(magazine []string, note []string) bool{
    for _, word := range note {
        success, idx := matchWordInSlice(word, magazine)
        if success {
            magazine = append(magazine[:idx], magazine[(idx+1):]...)
        } else {
            return false
        }
    }
    return true
}

func matchWordInSlice(targetWord string, words []string) (bool, int){
    for idx, word := range words {
        if targetWord == word { return true, idx}
    }
    return false, -1
}

func test() {
    fmt.Println(
        checkMagazine(
            []string{"give", "me", "one", "grand", "today", "night"},
            []string{"give", "one", "grand", "today"}))
}

func main() {
    // test()
    // return
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    mn := strings.Split(readLine(reader), " ")

    mTemp, err := strconv.ParseInt(mn[0], 10, 64)
    checkError(err)
    m := int32(mTemp)

    nTemp, err := strconv.ParseInt(mn[1], 10, 64)
    checkError(err)
    n := int32(nTemp)

    magazineTemp := strings.Split(readLine(reader), " ")

    var magazine []string

    for i := 0; i < int(m); i++ {
        magazineItem := magazineTemp[i]
        magazine = append(magazine, magazineItem)
    }

    noteTemp := strings.Split(readLine(reader), " ")

    var note []string

    for i := 0; i < int(n); i++ {
        noteItem := noteTemp[i]
        note = append(note, noteItem)
    }

    if checkMagazine(magazine, note) {
        fmt.Println("Yes")
    } else {
        fmt.Println("No")
    }
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
