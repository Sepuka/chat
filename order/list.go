package order

import (
	"os/exec"
)

type List struct {

}

func (o *List) Run() {
	c := exec.Command(`touch`, `/tmp/work`)
	c.Run()
}