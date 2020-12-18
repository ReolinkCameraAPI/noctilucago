//go:generate swagger generate spec
package main

import "github.com/ReolinkCameraAPI/noctilucago/cmd"


func main() {
	cmd.Execute()
}