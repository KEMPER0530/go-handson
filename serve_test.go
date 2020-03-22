package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("before test serve_test.go")
	code := m.Run()
	fmt.Println("after test serve_test.go")
	os.Exit(code)
}

func Testserve(t *testing.T) {
	serve("8080")
	t.Log("Testserve 終了")
}
