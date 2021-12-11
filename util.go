// Create and maintain by Chaiyapong Lapliengtrakul (chaiyapong@3dsinteractive.com), All right reserved (2021 - Present)
package main

import (
	"fmt"
	"math/rand"
)

func randString() string {
	i := rand.Int()
	return fmt.Sprintf("%d", i)
}

func print(format string, val ...interface{}) {
	fmt.Printf(format+"\n", val...)
}

func logE(err error) error {
	fmt.Printf("error: %s", err.Error())
	return err
}
