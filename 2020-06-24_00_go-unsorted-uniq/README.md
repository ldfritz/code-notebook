# Go unsorted uniq

I occasionally use this quick and dirty filter when merging files.  It prints each line of piped input only the first time that exact line is encountered.  I can `cat` CSV files together and as long as the headers match, only the first header will be printed.

On line 9, I attach to os.Stdin, instead of loading any files.  On line 16, I just print, instead of writing to a file.  Avoiding file handling in both these places keeps this script short.

On lines 8 and 11 through 15, I create and populate a map with each unique line as it is encountered.  I look for each line before writing it to the map, and if has been encountered already, I `continue` directly to the next line without printing.

I frequently repeat the pattern on lines 19 and 20.  I send formatted text to stderr and then exit with a non-zero code.  I like having this little bit of control over the exit and the formatting of the error.  This feels like very predictable behavior for this type of command-line utility.

 1	package main

 2	import (
 3		"bufio"
 4		"fmt"
 5		"os"
 6	)

 7	func main() {
 8		uniqRows := make(map[string]bool)
 9		scanner := bufio.NewScanner(os.Stdin)
10		for scanner.Scan() {
11			row := scanner.Text()
12			if uniqRows[row] {
13				continue
14			}
15			uniqRows[row] = true
16			fmt.Println(row)
17		}
18		if err := scanner.Err(); err != nil {
19			fmt.Fprintf(os.Stderr, "bufio.Scanner.Err: %v\n", err)
20			os.Exit(1)
21		}
22	}
