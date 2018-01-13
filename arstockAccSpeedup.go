


/*

golang arstock speedup acc 11.1.2018



*/

// Copyright 2017 George Loo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.


//jj jj
package main

import (
    "bufio"
    "bytes"
    "fmt"
    "io"
    "os"
    "log"
    "strings"
    "strconv"

)

const (
	kModelsFile = "D:\\agloo\\dev\\arstockdb\\databasefiles\\models.txt"
	kDataFile = "D:\\agloo\\dev\\arstockdb\\databasefiles\\database.txt"
	kModelnameIdx = 5
	kQuantityField = 6
	kFromField = 7
	kCompanyName = "AR2000"
	kToField = 8

)

type modelsIndicesObj struct {
	modelidx int 
	length int
	indices []int

}

type accDataObj struct {
	modelsList []string 
	recordsList []string

}

var (
	arstock accDataObj
	modelsIdx []modelsIndicesObj

)

func (a *accDataObj) loadmodels() {
    file, err := os.Open(kModelsFile)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
        a.modelsList = append(a.modelsList, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func (a *accDataObj) loadfile(fn string, listb *[]string) {
    file, err := os.Open(fn)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        //fmt.Println(scanner.Text())
        *listb = append(*listb, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    fmt.Println("Loadfile ends", fn)
}

func bruteforceProcessing() {
	var (
		i, last, j, lastRec, qty int
		//rq int 

	)
	i = 0
	j = 0
	qty = 1
	last = len(arstock.modelsList)
	last = 10
	lastRec = len(arstock.recordsList)
	//lastRec = 10
	for i < last - 1 {  // like a while loop
		//fmt.Println(arstock.modelsList[i])
		j = 0
		qty = 0
		for j < lastRec - 1 {
			aRecordSlice := strings.Split(arstock.recordsList[j], ",")
			//fmt.Println(aRecordSlice[kModelnameIdx])
			if aRecordSlice[kModelnameIdx] == arstock.modelsList[i] {
				//fmt.Println(arstock.recordsList[j])
				rq, err := strconv.Atoi(aRecordSlice[kQuantityField])
				if err != nil {
					fmt.Println("error qty!", aRecordSlice[kQuantityField])
				}
				if aRecordSlice[kFromField] == kCompanyName {
					qty -= rq
				} else if aRecordSlice[kToField] == kCompanyName {
					qty += rq
				}
				
			}
			
			j += 1
		}
		fmt.Println(arstock.modelsList[i]," = ", qty)
		i += 1
	}

}

func test1() {

	var arr []int
	var brr []int

	arr = append(arr,37)
	brr = append(brr,1965)

	modelsIdx = append(modelsIdx, modelsIndicesObj{ 10, 123, arr })
	modelsIdx = append(modelsIdx, modelsIndicesObj{ 12, 456, brr })

	fmt.Println("test1")
	fmt.Println(modelsIdx[0].modelidx)
	fmt.Println(modelsIdx[0].length)
	fmt.Println(modelsIdx[0].indices[0])

	//modelsIdx[1].indices[0] = 88
	fmt.Println(modelsIdx[1].modelidx)
	fmt.Println(modelsIdx[1].length)
	fmt.Println(modelsIdx[1].indices[0])
	fmt.Println(modelsIdx[0].indices[0])


}


func main() {
	var i int

	//arstock.loadmodels()
	arstock.loadfile(kModelsFile, &arstock.modelsList)
	fmt.Println("hello --------------- jibai")
	fmt.Println(arstock.modelsList[0])
	last := len(arstock.modelsList)
	fmt.Println(last)
	fmt.Println(arstock.modelsList[last-1])

	arstock.loadfile(kDataFile, &arstock.recordsList)
	fmt.Println("hello --------------- jibai")
	fmt.Println(arstock.recordsList[0])
	last = len(arstock.recordsList)
	fmt.Println(last)
	fmt.Println(arstock.recordsList[last-1])

	aRecordSlice := strings.Split(arstock.recordsList[last-1], ",")
	i = 0
	for i < 11 {  // like a while loop
		fmt.Println(aRecordSlice[i])
		i += 1
	
	}

	test1()
	bruteforceProcessing()

}

// from https://stackoverflow.com/questions/8757389/reading-file-line-by-line-in-go
func readFileWithReadString(fn string) (err error) {
    fmt.Println("readFileWithReadString")

    file, err := os.Open(fn)
    defer file.Close()

    if err != nil {
        return err
    }

    // Start reading from the file with a reader.
    reader := bufio.NewReader(file)

    var line string
    for {
        line, err = reader.ReadString('\n')

        fmt.Printf(" > Read %d characters\n", len(line))

        // Process the line here.
        fmt.Println(" > > " + limitLength(line, 50))

        if err != nil {
            break
        }
    }

    if err != io.EOF {
        fmt.Printf(" > Failed!: %v\n", err)
    }

    return
}

func readFileWithScanner(fn string) (err error) {
    fmt.Println("readFileWithScanner - this will fail!")

    // Don't use this, it doesn't work with long lines...

    file, err := os.Open(fn)
    defer file.Close()

    if err != nil {
        return err
    }

    // Start reading from the file using a scanner.
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()

        fmt.Printf(" > Read %d characters\n", len(line))

        // Process the line here.
        fmt.Println(" > > " + limitLength(line, 50))
    }

    if scanner.Err() != nil {
        fmt.Printf(" > Failed!: %v\n", scanner.Err())
    }

    return
}

func readFileWithReadLine(fn string) (err error) {
    fmt.Println("readFileWithReadLine")

    file, err := os.Open(fn)
    defer file.Close()

    if err != nil {
        return err
    }

    // Start reading from the file with a reader.
    reader := bufio.NewReader(file)

    for {
        var buffer bytes.Buffer

        var l []byte
        var isPrefix bool
        for {
            l, isPrefix, err = reader.ReadLine()
            buffer.Write(l)

            // If we've reached the end of the line, stop reading.
            if !isPrefix {
                break
            }

            // If we're just at the EOF, break
            if err != nil {
                break
            }
        }

        if err == io.EOF {
            break
        }

        line := buffer.String()

        fmt.Printf(" > Read %d characters\n", len(line))

        // Process the line here.
        fmt.Println(" > > " + limitLength(line, 50))
    }

    if err != io.EOF {
        fmt.Printf(" > Failed!: %v\n", err)
    }

    return
}

func main00() {
    //testLongLines()
    //testLinesThatDoNotFinishWithALinebreak()

    //readFileWithReadLine(kModelsFile)
    readFileWithReadString(kModelsFile)
}

func testLongLines() {
    fmt.Println("Long lines")
    fmt.Println()

    createFileWithLongLine("longline.txt")
    readFileWithReadString("longline.txt")
    fmt.Println()
    readFileWithScanner("longline.txt")
    fmt.Println()
    readFileWithReadLine("longline.txt")
    fmt.Println()
}

func testLinesThatDoNotFinishWithALinebreak() {
    fmt.Println("No linebreak")
    fmt.Println()

    createFileThatDoesNotEndWithALineBreak("nolinebreak.txt")
    readFileWithReadString("nolinebreak.txt")
    fmt.Println()
    readFileWithScanner("nolinebreak.txt")
    fmt.Println()
    readFileWithReadLine("nolinebreak.txt")
    fmt.Println()
}

func createFileThatDoesNotEndWithALineBreak(fn string) (err error) {
    file, err := os.Create(fn)
    defer file.Close()

    if err != nil {
        return err
    }

    w := bufio.NewWriter(file)
    w.WriteString("Does not end with linebreak.")
    w.Flush()

    return
}

func createFileWithLongLine(fn string) (err error) {
    file, err := os.Create(fn)
    defer file.Close()

    if err != nil {
        return err
    }

    w := bufio.NewWriter(file)

    fs := 1024 * 1024 * 4 // 4MB

    // Create a 4MB long line consisting of the letter a.
    for i := 0; i < fs; i++ {
        w.WriteRune('a')
    }

    // Terminate the line with a break.
    w.WriteRune('\n')

    // Put in a second line, which doesn't have a linebreak.
    w.WriteString("Second line.")

    w.Flush()

    return
}

func limitLength(s string, length int) string {
    if len(s) < length {
        return s
    }

    return s[:length]
}