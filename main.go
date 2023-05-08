package main

import (
	"fmt"
	"os"

	"github.com/Jxancestral17/abusing_hop_by_hop_header/headers"
)

func main() {
	headers.Url = os.Args[1]
	fmt.Println("[+]Read headers file")
	headers.Headers = headers.Readfile()
	// for i := 0; i < len(headers.Headers); i++ {

	// 	headers.Params1 = headers.GeneretorRandomString()
	// 	headers.Params2 = headers.GeneretorRandomString()

	// 	resp1 := headers.MakeRequestNoHeaders()
	// 	resp2 := headers.MakeRequestWithHeaders(headers.Headers[i])

	// 	if resp1 != nil && resp2 != nil {
	// 		headers.Check(resp1, resp2)
	// 	} else {
	// 		fmt.Println("[-]No poisoning detected")
	// 	}
	// 	time.Sleep(time.Millisecond * 300)
	// }

}
