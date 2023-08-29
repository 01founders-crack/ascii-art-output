package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func compareFiles(t *testing.T, actual, expected string) {
	actualBytes, err := os.ReadFile(actual)
	if err != nil {
		t.Fatalf("Error reading actual output file: %v", err)
	}

	expectedBytes, err := os.ReadFile(expected)
	if err != nil {
		t.Fatalf("Error reading expected output file: %v", err)
	}

	if !bytes.Equal(actualBytes, expectedBytes) {
		t.Errorf("Output mismatch for %s\nExpected:\n%s\nActual:\n%s", actual, expectedBytes, actualBytes)
	}
}

func TestMain(t *testing.T) {
	// Compile the program
	cmd := exec.Command("go", "build", "-o", "test_program", "your_program_name.go")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error compiling program:", err)
		os.Exit(1)
	}
	defer os.Remove("test_program")

	testCases := []struct {
		args            []string
		expectedOutput  string
		expectedFile    string
	}{
		{[]string{"--output=test01.txt", "hello", "standard"}, "expected_output_01.txt", "test01.txt"},
	}

	for _, tc := range testCases {
		t.Run(strings.Join(tc.args, " "), func(t *testing.T) {
			cmd := exec.Command("./test_program", tc.args...)
			if err := cmd.Run(); err != nil {
				t.Fatalf("Error running program: %v", err)
			}
			compareFiles(t, tc.expectedFile, tc.expectedOutput)
		})
	}

}
