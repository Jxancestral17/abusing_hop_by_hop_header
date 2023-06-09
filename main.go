package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Jxancestral17/abusing_hop_by_hop_header/headers"
)

func main() {
	headers.Url = os.Args[1]

	fmt.Println("[+]Read headers file")
	headers.Headers = headers.Readfile()
	for i := 0; i < len(headers.Headers); i++ {

		headers.Params1 = headers.GeneretorRandomString()
		headers.Params2 = headers.GeneretorRandomString()

		resp1 := headers.MakeRequestNoHeaders()
		resp2 := headers.MakeRequestWithHeaders(headers.Headers[i])

		if resp1 != nil && resp2 != nil {
			headers.Check(resp1, resp2, headers.Headers[i])
		} else {
			fmt.Println("[-]No poisoning detected")
		}
		time.Sleep(time.Millisecond * 300)
	}
	fmt.Println("")
	if headers.Result == 0 {
		fmt.Println("[-]No poisoning detected")
	} else {
		fmt.Println("[+]Poisoning detected open report.log")
	}
	fmt.Println("\n\nCreated by Jxancestral17 - 2023 - V 1.5")
}
