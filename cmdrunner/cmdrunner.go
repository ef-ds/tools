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
// This tool runs all the commands provided in the --input parameter and
// outputs the results to the console or into the supplied --output file.
//
// Each command is limited to a single line.
//
// Lines that starts with '//' (comment) are not executed, but their content
// is written either to the console or the the specified output file.
//
// The result of each command is output with "```" before and after the result.
//
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ef-ds/tools/common"
)

func main() {
	input := flag.String("input", "", "The path to the file that contains the commands to run.")
	output := flag.String("output", "", "The path to the file where the results will be written to.")
	flag.Parse()
	if len(*input) == 0 {
		fmt.Println("--input parameter not specified.")
		os.Exit(1)
	}

	fmt.Println("Input file '" + *input + "'")
	fmt.Println("Output file '" + *output + "'")
	runCommands(*input, *output)
	fmt.Println("\nFile processed successfully.")
}

func runCommands(input, output string) {
	const comment = "//"
	var outputWriter *bufio.Writer
	var outputFileHandle *os.File
	if len(output) > 0 {
		outputFileHandle, _ = os.Create(output)
	} else {
		outputFileHandle = os.Stdout
	}
	outputWriter = bufio.NewWriter(outputFileHandle)
	outputWriter.WriteRune('\n')

	common.ReadFile(input, func(line string) {
		if len(line) > 0 {
			if strings.HasPrefix(line, comment) {
				outputWriter.WriteString(line[strings.Index(line, comment)+len(comment):])
			} else {
				mainCmd, args := common.SplitCommand(line)
				outputWriter.WriteString("```")
				outputWriter.WriteRune('\n')
				outputWriter.WriteString(line)
				outputWriter.WriteRune('\n')
				output, _ := common.ExecuteCommand(outputWriter, mainCmd, args...)
				outputWriter.WriteString(output)
				outputWriter.WriteString("```")
			}
		}
		outputWriter.WriteRune('\n')
	})
	outputWriter.Flush()
	if len(output) > 0 {
		outputFileHandle.Close()
	}
}
