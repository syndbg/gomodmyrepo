package cmd

import (
	"fmt"
	"github.com/syndbg/gomodmyrepo/mig"
	cmd2 "github.com/syndbg/gomodanotherrepo/cmd"
)

func PrintExample() {
	cmd2.PrintExample()
	fmt.Println(mig.Example)
}
