package main

import (
	"fmt"
	"godefaultvalue"
)

type Test struct {
	Name string `defaultV:"\"test\""`
	godefaultvalue.Godefault[Test]
}

func main() {
	t := Test{Godefault: godefaultvalue.Godefault[Test]{}}
	fmt.Println(t.GetDefault("Name"))
}
