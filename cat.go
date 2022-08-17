package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	stdin := os.Stdin
	stat, _ := stdin.Stat()

	// Read data from STDIN and print it to STDOUT.
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		var maxBytes = 255

		stdin.SetReadDeadline(time.Now().Add(10 * time.Second))

		for {
			var fbs = make([]byte, maxBytes)
			_, err := stdin.Read(fbs)

			if err != nil {
				break
			}

			fmt.Println(maxBytes)
		}
	} else {
		// Read files from Args and print to STDOUT.
		files := os.Args[1:]
		for _, file := range files {
			contents, _ := os.ReadFile(file)
			fmt.Print(string(contents))
		}
	}
}

/*

Man docs from MacOS cat.
-------

The cat utility reads files sequentially, writing them to the standard
output.  The file operands are processed in command-line order.  If file is
a single dash (‘-’) or absent, cat reads from the standard input.  If file
is a UNIX domain socket, cat connects to it and then reads it until EOF.
This complements the UNIX domain binding capability available in inetd(8).

The options are as follows:

-b      Number the non-blank output lines, starting at 1.

-e      Display non-printing characters (see the -v option), and display a
				dollar sign (‘$’) at the end of each line.

-l      Set an exclusive advisory lock on the standard output file
				descriptor.  This lock is set using fcntl(2) with the F_SETLKW
				command.  If the output file is already locked, cat will block
				until the lock is acquired.

-n      Number the output lines, starting at 1.

-s      Squeeze multiple adjacent empty lines, causing the output to be
				single spaced.

-t      Display non-printing characters (see the -v option), and display
				tab characters as ‘^I’.

-u      Disable output buffering.

-v      Display non-printing characters so they are visible.  Control
				characters print as ‘^X’ for control-X; the delete character (octal
				0177) prints as ‘^?’.  Non-ASCII characters (with the high bit set)
				are printed as ‘M-’ (for meta) followed by the character for the
				low 7 bits.
*/
