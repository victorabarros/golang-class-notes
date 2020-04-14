// https://www.hackerrank.com/challenges/minimum-swaps-2/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int) int {
    var changeRequest bool = true
    var changes int
    var newArr []int = arr

    for changeRequest == true {
        changeRequest = false

        for idx, elem := range arr {
            if idx + 1 != elem {
                changeRequest = true

                newArr[idx] = idx + 1  // the correct value

                targetIdx, _ := findIndex(arr, idx + 1)
                newArr[targetIdx] = elem
                changes++
                break;
            }
        }
        arr = newArr
    }
    return changes
}

func findIndex(arr []int, target int) (int, error) {
    for idx, elem := range arr {
        if elem == target {
            return idx, nil
        }
    }
    var err error
    return 0, err
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int(arrItemTemp)
        arr = append(arr, arrItem)
    }

    res := minimumSwaps(arr)
    fmt.Fprintf(writer, "%d\n", res)

    writer.Flush()
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
