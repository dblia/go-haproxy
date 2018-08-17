/*

  A sample module emulating a HAProxy prompt (aka interactive-like mode).

  The current sample displays a prompt ('>') and waits for commands to be entered
  on the line. It processes them, returns the response, and then it re-displays
  the prompt again waiting for a new command.

  Example:

      $ go run samples/prompt.go unix /run/haproxy.sock
      > show acl
      ...
      > show info
      ...
      >

*/

package main

import "bufio"
import "fmt"
import "os"
import "time"

import "go-haproxy/haproxy"

func main() {

	if len(os.Args) != 3 {
		fmt.Println("usage: go run samples/prompt.go <afnet> <addr>")
		os.Exit(1)
	}

	client := haproxy.HAProxyClient{
		AfNet:   os.Args[1],
		Address: os.Args[2],
		Timeout: time.Second * 2,
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		msg, _ := reader.ReadString('\n')
		if msg == "q\n" || msg == "quit\n" {
			fmt.Println("Exiting...")
			break
		}
		if msg == "prompt\n" {
			fmt.Println("[WARN] You are already in a prompt-like mode.\n")
			continue
		}
		resp, err := client.SendCommand(msg)
		if err != nil {
			fmt.Println("[ERROR]", err)
			os.Exit(1)
		}
		fmt.Printf(resp)
	}
}
