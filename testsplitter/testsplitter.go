// Copyright (c) 2018 ef-ds
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//
// This tool splits a benchmark file into one benchmark file per bechmark
// test inside the file. The name of the new file is the name of the test and
// the test name written inside the file is written hard-coded as "Benchmark".
// Below benchmark file will be split in below two files.
// This is useful to run benchstat to compare different test names (TestQueue1 and TestQueue2).
//
// ORIGINAL FILE CONTENT
// BenchmarkFillSliceQueue/0-4         	30000000	        44.3 ns/op	      32 B/op	       1 allocs/op
// BenchmarkFillSliceQueue/0-4         	30000000	        49.2 ns/op	      32 B/op	       1 allocs/op
// BenchmarkFillSliceQueue/0-4         	30000000	        44.3 ns/op	      32 B/op	       1 allocs/op
// BenchmarkRefillFullSliceQueue/10000-4        	      30	  41071216 ns/op	52044800 B/op	 1000250 allocs/op
//
// SPLIT FILE ONE NAME: BenchmarkFillSliceQueue.txt
// SPLIT FILE ONE CONTENT
// Benchmark/0-4         	30000000	        44.3 ns/op	      32 B/op	       1 allocs/op
// Benchmark/0-4         	30000000	        49.2 ns/op	      32 B/op	       1 allocs/op
// Benchmark/0-4         	30000000	        44.3 ns/op	      32 B/op	       1 allocs/op
//
// SPLIT FILE ONE NAME: BenchmarkRefillFullSliceQueue.txt
// SPLIT FILE ONE CONTENT
// Benchmark/10000-4        	      30	  41071216 ns/op	52044800 B/op	 1000250 allocs/op
//
// To run it, from the tool main directory, use below command.
// go run *.go --file PATH_T0_BENCH_FILE_TO_SPLIT
//
// After running this, benchstat can be used to compare the split files. Below are useful benchstat
// commands for deque vs other queues.
//
// Deque vs list
// benchstat testdata/BenchmarkFillDequeQueue.txt testdata/BenchmarkFillListQueue.txt
// benchstat testdata/BenchmarkRefillDequeQueue.txt testdata/BenchmarkRefillListQueue.txt
// benchstat testdata/BenchmarkRefillFullDequeQueue.txt testdata/BenchmarkRefillFullListQueue.txt
// benchstat testdata/BenchmarkSlowIncreaseDequeQueue.txt testdata/BenchmarkSlowIncreaseListQueue.txt
// benchstat testdata/BenchmarkSlowDecreaseDequeQueue.txt testdata/BenchmarkSlowDecreaseListQueue.txt
// benchstat testdata/BenchmarkStableDequeQueue.txt testdata/BenchmarkStableListQueue.txt
// benchstat testdata/BenchmarkMicroserviceDequeQueue.txt testdata/BenchmarkMicroserviceListQueue.txt
// benchstat testdata/BenchmarkFillDequeStack.txt testdata/BenchmarkFillListStack.txt
// benchstat testdata/BenchmarkRefillDequeStack.txt testdata/BenchmarkRefillListStack.txt
// benchstat testdata/BenchmarkRefillFullDequeStack.txt testdata/BenchmarkRefillFullListStack.txt
// benchstat testdata/BenchmarkSlowIncreaseDequeStack.txt testdata/BenchmarkSlowIncreaseListStack.txt
// benchstat testdata/BenchmarkSlowDecreaseDequeStack.txt testdata/BenchmarkSlowDecreaseListStack.txt
// benchstat testdata/BenchmarkStableDequeStack.txt testdata/BenchmarkStableListStack.txt
// benchstat testdata/BenchmarkMicroserviceDequeStack.txt testdata/BenchmarkMicroserviceListStack.txt

// Deque vs slice
// benchstat testdata/BenchmarkFillDequeQueue.txt testdata/BenchmarkFillSliceQueue.txt
// benchstat testdata/BenchmarkRefillDequeQueue.txt testdata/BenchmarkRefillSliceQueue.txt
// benchstat testdata/BenchmarkRefillFullDequeQueue.txt testdata/BenchmarkRefillFullSliceQueue.txt
// benchstat testdata/BenchmarkSlowIncreaseDequeQueue.txt testdata/BenchmarkSlowIncreaseSliceQueue.txt
// benchstat testdata/BenchmarkSlowDecreaseDequeQueue.txt testdata/BenchmarkSlowDecreaseSliceQueue.txt
// benchstat testdata/BenchmarkStableDequeQueue.txt testdata/BenchmarkStableSliceQueue.txt
// benchstat testdata/BenchmarkMicroserviceDequeQueue.txt testdata/BenchmarkMicroserviceSliceQueue.txt
// benchstat testdata/BenchmarkFillDequeStack.txt testdata/BenchmarkFillSliceStack.txt
// benchstat testdata/BenchmarkRefillDequeStack.txt testdata/BenchmarkRefillSliceStack.txt
// benchstat testdata/BenchmarkRefillFullDequeStack.txt testdata/BenchmarkRefillFullSliceStack.txt
// benchstat testdata/BenchmarkSlowIncreaseDequeStack.txt testdata/BenchmarkSlowIncreaseSliceStack.txt
// benchstat testdata/BenchmarkSlowDecreaseDequeStack.txt testdata/BenchmarkSlowDecreaseSliceStack.txt
// benchstat testdata/BenchmarkStableDequeStack.txt testdata/BenchmarkStableSliceStack.txt
// benchstat testdata/BenchmarkMicroserviceDequeStack.txt testdata/BenchmarkMicroserviceSliceStack.txt

// Deque vs gammazero
// benchstat testdata/BenchmarkFillDequeQueue.txt testdata/BenchmarkFillGammazeroQueue.txt
// benchstat testdata/BenchmarkRefillDequeQueue.txt testdata/BenchmarkRefillGammazeroQueue.txt
// benchstat testdata/BenchmarkRefillFullDequeQueue.txt testdata/BenchmarkRefillFullGammazeroQueue.txt
// benchstat testdata/BenchmarkSlowIncreaseDequeQueue.txt testdata/BenchmarkSlowIncreaseGammazeroQueue.txt
// benchstat testdata/BenchmarkSlowDecreaseDequeQueue.txt testdata/BenchmarkSlowDecreaseGammazeroQueue.txt
// benchstat testdata/BenchmarkStableDequeQueue.txt testdata/BenchmarkStableGammazeroQueue.txt
// benchstat testdata/BenchmarkMicroserviceDequeQueue.txt testdata/BenchmarkMicroserviceGammazeroQueue.txt
// benchstat testdata/BenchmarkFillDequeStack.txt testdata/BenchmarkFillGammazeroStack.txt
// benchstat testdata/BenchmarkRefillDequeStack.txt testdata/BenchmarkRefillGammazeroStack.txt
// benchstat testdata/BenchmarkRefillFullDequeStack.txt testdata/BenchmarkRefillFullGammazeroStack.txt
// benchstat testdata/BenchmarkSlowIncreaseDequeStack.txt testdata/BenchmarkSlowIncreaseGammazeroStack.txt
// benchstat testdata/BenchmarkSlowDecreaseDequeStack.txt testdata/BenchmarkSlowDecreaseGammazeroStack.txt
// benchstat testdata/BenchmarkStableDequeStack.txt testdata/BenchmarkStableGammazeroStack.txt
// benchstat testdata/BenchmarkMicroserviceDequeStack.txt testdata/BenchmarkMicroserviceGammazeroStack.txt

// Deque vs phf
// benchstat testdata/BenchmarkFillDequeQueue.txt testdata/BenchmarkFillPhfQueue.txt
// benchstat testdata/BenchmarkRefillDequeQueue.txt testdata/BenchmarkRefillPhfQueue.txt
// benchstat testdata/BenchmarkRefillFullDequeQueue.txt testdata/BenchmarkRefillFullPhfQueue.txt
// benchstat testdata/BenchmarkSlowIncreaseDequeQueue.txt testdata/BenchmarkSlowIncreasePhfQueue.txt
// benchstat testdata/BenchmarkSlowDecreaseDequeQueue.txt testdata/BenchmarkSlowDecreasePhfQueue.txt
// benchstat testdata/BenchmarkStableDequeQueue.txt testdata/BenchmarkStablePhfQueue.txt
// benchstat testdata/BenchmarkMicroserviceDequeQueue.txt testdata/BenchmarkMicroservicePhfQueue.txt
// benchstat testdata/BenchmarkFillDequeStack.txt testdata/BenchmarkFillPhfStack.txt
// benchstat testdata/BenchmarkRefillDequeStack.txt testdata/BenchmarkRefillPhfStack.txt
// benchstat testdata/BenchmarkRefillFullDequeStack.txt testdata/BenchmarkRefillFullPhfStack.txt
// benchstat testdata/BenchmarkSlowIncreaseDequeStack.txt testdata/BenchmarkSlowIncreasePhfStack.txt
// benchstat testdata/BenchmarkSlowDecreaseDequeStack.txt testdata/BenchmarkSlowDecreasePhfStack.txt
// benchstat testdata/BenchmarkStableDequeStack.txt testdata/BenchmarkStablePhfStack.txt
// benchstat testdata/BenchmarkMicroserviceDequeStack.txt testdata/BenchmarkMicroservicePhfStack.txt

// Deque vs cookiejar
// benchstat testdata/BenchmarkFillDequeQueue.txt testdata/BenchmarkFillCookiejarQueue.txt
// benchstat testdata/BenchmarkRefillDequeQueue.txt testdata/BenchmarkRefillCookiejarQueue.txt
// benchstat testdata/BenchmarkRefillFullDequeQueue.txt testdata/BenchmarkRefillFullCookiejarQueue.txt
// benchstat testdata/BenchmarkSlowIncreaseDequeQueue.txt testdata/BenchmarkSlowIncreaseCookiejarQueue.txt
// benchstat testdata/BenchmarkSlowDecreaseDequeQueue.txt testdata/BenchmarkSlowDecreaseCookiejarQueue.txt
// benchstat testdata/BenchmarkStableDequeQueue.txt testdata/BenchmarkStableCookiejarQueue.txt
// benchstat testdata/BenchmarkMicroserviceDequeQueue.txt testdata/BenchmarkMicroserviceCookiejarQueue.txt
// benchstat testdata/BenchmarkFillDequeStack.txt testdata/BenchmarkFillCookiejarStack.txt
// benchstat testdata/BenchmarkRefillDequeStack.txt testdata/BenchmarkRefillCookiejarStack.txt
// benchstat testdata/BenchmarkRefillFullDequeStack.txt testdata/BenchmarkRefillFullCookiejarStack.txt
// benchstat testdata/BenchmarkSlowIncreaseDequeStack.txt testdata/BenchmarkSlowIncreaseCookiejarStack.txt
// benchstat testdata/BenchmarkSlowDecreaseDequeStack.txt testdata/BenchmarkSlowDecreaseCookiejarStack.txt
// benchstat testdata/BenchmarkStableDequeStack.txt testdata/BenchmarkStableCookiejarStack.txt
// benchstat testdata/BenchmarkMicroserviceDequeStack.txt testdata/BenchmarkMicroserviceCookiejarStack.txt

// Deque vs impl7
// benchstat testdata/BenchmarkFillDequeQueue.txt testdata/BenchmarkFillImpl7Queue.txt
// benchstat testdata/BenchmarkRefillDequeQueue.txt testdata/BenchmarkRefillImpl7Queue.txt
// benchstat testdata/BenchmarkRefillFullDequeQueue.txt testdata/BenchmarkRefillFullImpl7Queue.txt
// benchstat testdata/BenchmarkSlowIncreaseDequeQueue.txt testdata/BenchmarkSlowIncreaseImpl7Queue.txt
// benchstat testdata/BenchmarkSlowDecreaseDequeQueue.txt testdata/BenchmarkSlowDecreaseImpl7Queue.txt
// benchstat testdata/BenchmarkStableDequeQueue.txt testdata/BenchmarkStableImpl7Queue.txt
// benchstat testdata/BenchmarkMicroserviceDequeQueue.txt testdata/BenchmarkMicroserviceImpl7Queue.txt

// Deque vs deque2 (useful to test changes to deque)
// benchstat testdata/BenchmarkFillDequeQueue.txt testdata/BenchmarkFillDequeQueue2.txt
// benchstat testdata/BenchmarkRefillDequeQueue.txt testdata/BenchmarkRefillDequeQueue2.txt
// benchstat testdata/BenchmarkRefillFullDequeQueue.txt testdata/BenchmarkRefillFullDequeQueue2.txt
// benchstat testdata/BenchmarkSlowIncreaseDequeQueue.txt testdata/BenchmarkSlowIncreaseDequeQueue2.txt
// benchstat testdata/BenchmarkSlowDecreaseDequeQueue.txt testdata/BenchmarkSlowDecreaseDequeQueue2.txt
// benchstat testdata/BenchmarkStableDequeQueue.txt testdata/BenchmarkStableDequeQueue2.txt
// benchstat testdata/BenchmarkMicroserviceDequeQueue.txt testdata/BenchmarkMicroserviceDequeQueue2.txt
// benchstat testdata/BenchmarkFillDequeStack.txt testdata/BenchmarkFillDequeStack2.txt
// benchstat testdata/BenchmarkRefillDequeStack.txt testdata/BenchmarkRefillDequeStack2.txt
// benchstat testdata/BenchmarkRefillFullDequeStack.txt testdata/BenchmarkRefillFullDequeStack2.txt
// benchstat testdata/BenchmarkSlowIncreaseDequeStack.txt testdata/BenchmarkSlowIncreaseDequeStack2.txt
// benchstat testdata/BenchmarkSlowIncreaseDequeStack.txt testdata/BenchmarkSlowIncreaseDequeStack2.txt
// benchstat testdata/BenchmarkStableDequeStack.txt testdata/BenchmarkStableDequeStack2.txt
// benchstat testdata/BenchmarkMicroserviceDequeStack.txt testdata/BenchmarkMicroserviceDequeStack2.txt

package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

func main() {
	file := flag.String("file", "", "The path to the test file to split.")
	suffix := flag.String("suffix", "", "The suffix to be added to each split test file.")
	flag.Parse()
	if len(*file) == 0 {
		fmt.Println("--file parameter not specified.")
		os.Exit(1)
	}

	fmt.Println("Splitting file '" + *file + "'")
	fmt.Println("Files suffix '" + *suffix + "'")
	splitFile(*file, *suffix)
	fmt.Println("File split successfully.")
}

func splitFile(file, suffix string) {
	var writer *bufio.Writer
	var fileHandle *os.File
	var previousTestName string
	const benchmark = "Benchmark"
	getTestName := func(line string) (string, error) {
		var b bytes.Buffer
		for i := 0; i < len(line); i++ {
			if line[i] == ' ' || line[i] == '/' {
				return b.String(), nil
			}
			b.WriteByte(line[i])
		}
		return "", errors.New("test name not found")
	}

	readFile(file, func(line string) {
		if strings.HasPrefix(line, benchmark) {
			if testName, err := getTestName(line); err == nil {
				if testName != previousTestName {
					if writer != nil {
						writer.Flush()
						fileHandle.Close()
					}
					fileHandle, _ = os.Create(path.Dir(file) + "/" + testName + suffix + ".txt")
					writer = bufio.NewWriter(fileHandle)
					previousTestName = testName
				}
				writer.WriteString(benchmark)
				writer.WriteString(line[len(testName):])
				writer.WriteRune('\n')
			}
		}
	})
	writer.Flush()
	fileHandle.Close()
}

func readFile(fileName string, f func(line string)) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
