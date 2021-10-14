package main

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.octolab.org/safe"
)

func TestExecution(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		exit = func(code int) { assert.Equal(t, 0, code) }
		stderr, stdout = io.Discard, io.Discard
		os.Args = []string{"root", "version"}
		main()
	})

	t.Run("failure", func(t *testing.T) {
		exit = func(code int) { assert.Equal(t, 1, code) }
		stderr, stdout = io.Discard, io.Discard
		os.Args = []string{"root", "unknown"}
		main()
	})

	t.Run("shutdown with panic", func(t *testing.T) {
		exit = func(code int) { assert.Equal(t, 1, code) }
		stderr, stdout = io.Discard, io.Discard
		safe.Do(func() error { panic("test") }, shutdown)
	})
}
