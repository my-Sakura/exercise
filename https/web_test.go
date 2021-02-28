package main

import (
	"net/http"
	"testing"
	"time"
)

type situation struct {
	Name string
	A    []string
	Want []bool
}

func createTestCase(t *testing.T, c *situation) {
	t.Helper()
	t.Run(c.Name, func(t *testing.T) {
		got, err := BcjClient(c.A)
		if err != nil {
			t.Fatalf("runtime err %v", err)
		}
		if len(got) != len(c.Want) {
			t.Fatalf("A %v got %v want %v", c.A, got, c.Want)
		}

		for i, v := range c.Want {
			if v != got[i] {
				t.Fatalf("A %v got %v want %v", c.A, got, c.Want)
			}
		}
	})
}

func TestBcjClient(t *testing.T) {
	go func(t *testing.T) {
		http.HandleFunc("/", StringSliceHandler)
		err := http.ListenAndServeTLS(":1234", "./cert.pem", "./key.pem", nil)
		if err != nil {
			t.Fatalf("server start fail %v \n", err)
		}
	}(t)

	time.Sleep(time.Second)

	createTestCase(t, &situation{"normal", []string{"hi", "server"}, []bool{false, false}})
	createTestCase(t, &situation{"normal", []string{"hi", "server"}, []bool{true, true}})
	createTestCase(t, &situation{"mutiple", []string{"hello", "hello"}, []bool{false, true}})
}
