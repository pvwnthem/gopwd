package main

import (
	"fmt"

	"github.com/pvwnthem/gopwd/cmd"
)

func main() {
	version := GetVersion()
	cmd.Version = fmt.Sprintf("%d.%d.%d", version.Major, version.Minor, version.Patch)
	cmd.Execute()

}
