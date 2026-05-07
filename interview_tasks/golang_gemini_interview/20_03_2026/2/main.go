package main

import (
	"time"
	"fmt"
)

func doExternalCall() error {
	return nil
}

func handleRequest() error {
	errChan := make(chan error, 1)
	
	go func() {
		err := doExternalCall()
		errChan <- err
	}()

	select {
	case err := <-errChan:
		return err
	case <-time.After(500 * time.Millisecond):
		return fmt.Errorf("timeout")
	}
}