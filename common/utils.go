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
package common

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"os/exec"
)

// ReadFile reads the specified file line-by-line executing func f for each line.
func ReadFile(fileName string, f func(line string)) {
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

// ExecuteCommand execute the specified command and return its output.
func ExecuteCommand(writer *bufio.Writer, cmd string, arg ...string) (string, error) {
	var b bytes.Buffer
	var stdout, stderr bytes.Buffer
	c := exec.Command(cmd, arg...)
	c.Stdout = &stdout
	c.Stderr = &stderr

	err := c.Run()
	b.WriteString(string(stdout.Bytes()))
	if err != nil {
		b.WriteString(string(stderr.Bytes()))
		b.WriteRune('\n')
		b.WriteString(err.Error())
	}
	return b.String(), err
}

// SplitCommand splits a command line and return its main command and parameters.
func SplitCommand(cmd string) (mainCmd string, args []string) {
	mainCmd = ""
	foundQuote := false
	var b bytes.Buffer
	args = make([]string, 0)
	writeArg := func(arg string) {
		if len(mainCmd) == 0 {
			mainCmd = arg
		} else {
			args = append(args, arg)
		}
	}
	for i := 0; i < len(cmd); i++ {
		if cmd[i] == ' ' || cmd[i] == '"' || cmd[i] == '\'' {
			if cmd[i] == '"' || cmd[i] == '\'' {
				if foundQuote {
					writeArg(b.String())
					foundQuote = false
				} else {
					foundQuote = true
				}
				b.Reset()
			} else if !foundQuote {
				if b.Len() > 0 {
					writeArg(b.String())
					b.Reset()
				}
			} else {
				b.WriteByte(cmd[i])
			}
		} else {
			b.WriteByte(cmd[i])
		}
	}
	if b.Len() > 0 {
		writeArg(b.String())
	}
	return mainCmd, args
}
