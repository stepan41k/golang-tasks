package main

import (
	"time"
	"fmt"
)

func handleRequest() error {
	errChan := make(chan error, 1)
	
	go func() {
		err := doExternalCall() // допустим, это долгий сетевой запрос
		errChan <- err
	}()

	select {
	case err := <-errChan:
		return err
	case <-time.After(500 * time.Millisecond):
		return fmt.Errorf("timeout")
	}
}