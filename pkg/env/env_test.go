package env

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestGetENV_Success(t *testing.T) {
	// Arrange
	want := "production"
	os.Setenv("GO_ENV", "Production")

	// Act
	got := GetENV()

	// Assert
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestGetENV_NotSet(t *testing.T) {
	// Arrange: unset env
	os.Unsetenv("GO_ENV")

	// Because GetENV() calls logger.Log.Fatalln, which os.Exit(1)'s,
	// we must test it in a subprocess (since os.Exit terminates the process).
	if os.Getenv("TEST_EXIT") == "1" {
		GetENV()
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestGetENV_NotSet")
	cmd.Env = append(os.Environ(), "TEST_EXIT=1")
	output, err := cmd.CombinedOutput()

	// Assert: expect process to exit with non-nil error
	if err == nil {
		t.Fatalf("expected fatal exit when GO_ENV not set, but test did not exit")
	}

	if !strings.Contains(string(output), "GO_ENV is not set") {
		t.Errorf("expected fatal log message about missing GO_ENV, got %s", output)
	}
}

func TestLoadENV_Success(t *testing.T) {
	// Arrange: create a temporary .env file
	tmpDir := t.TempDir()
	envPath := filepath.Join(tmpDir, ".env")
	if err := os.WriteFile(envPath, []byte("FOO=bar\n"), 0644); err != nil {
		t.Fatalf("failed to write temp .env: %v", err)
	}

	// Change working dir so godotenv.Load() finds the .env
	oldWd, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(oldWd)

	// Act: should not log fatal
	LoadENV(context.Background())

	// Assert: env var should now be loaded
	if got := os.Getenv("FOO"); got != "bar" {
		t.Errorf("expected env var FOO=bar, got %q", got)
	}
}

func TestLoadENV_NoEnvFile(t *testing.T) {
	// Arrange: ensure no .env file exists in current dir
	tmpDir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(oldWd)

	// Test fatal behavior in subprocess
	if os.Getenv("TEST_EXIT") == "1" {
		LoadENV(context.Background())
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestLoadENV_NoEnvFile")
	cmd.Env = append(os.Environ(), "TEST_EXIT=1")
	output, err := cmd.CombinedOutput()

	// Assert
	if err == nil {
		t.Fatalf("expected fatal exit when no .env file found, but test did not exit")
	}

	if !strings.Contains(string(output), "No .env file found") {
		t.Errorf("expected fatal log about missing .env, got %s", output)
	}
}
