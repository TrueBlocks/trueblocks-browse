package updater

import (
	"testing"
	"time"
)

func TestNewUpdaterWithPaths(t *testing.T) {
	updater, err := NewUpdater("test", []string{"path1", "path2"}, File)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(updater.Paths) != 2 {
		t.Errorf("expected 2 paths, got %d", len(updater.Paths))
	}
	if updater.UpdateType != File {
		t.Errorf("expected UpdateType to be File, got %v", updater.UpdateType)
	}
	if updater.Duration != 0 {
		t.Errorf("expected duration to be 0, got %v", updater.Duration)
	}
	if updater.LastTs == 0 {
		t.Errorf("expected LastTs to be non-zero, got %d", updater.LastTs)
	}
}

func TestNewUpdaterWithDuration(t *testing.T) {
	duration := 10 * time.Second
	updater, err := NewUpdater("test", []string{}, Timer, duration)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(updater.Paths) != 0 {
		t.Errorf("expected no paths, got %d", len(updater.Paths))
	}
	if updater.UpdateType != Timer {
		t.Errorf("expected UpdateType to be Timer, got %v", updater.UpdateType)
	}
	if updater.Duration != duration {
		t.Errorf("expected duration %v, got %v", duration, updater.Duration)
	}
	if updater.LastTs == 0 {
		t.Errorf("expected LastTs to be non-zero, got %d", updater.LastTs)
	}
}

func TestNewUpdaterWithPathsAndDuration(t *testing.T) {
	duration := 15 * time.Second
	updater, err := NewUpdater("test", []string{"path1"}, Folder, duration)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(updater.Paths) != 1 {
		t.Errorf("expected 1 path, got %d", len(updater.Paths))
	}
	if updater.Paths[0] != "path1" {
		t.Errorf("expected path 'path1', got %s", updater.Paths[0])
	}
	if updater.UpdateType != Folder {
		t.Errorf("expected UpdateType to be Folder, got %v", updater.UpdateType)
	}
	if updater.Duration != duration {
		t.Errorf("expected duration %v, got %v", duration, updater.Duration)
	}
	if updater.LastTs == 0 {
		t.Errorf("expected LastTs to be non-zero, got %d", updater.LastTs)
	}
}

func TestNeedsUpdateWithDuration(t *testing.T) {
	duration := 10 * time.Second
	updater, err := NewUpdater("test", []string{}, Timer, duration)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	updater.LastTs = time.Now().Add(-20 * time.Second).Unix()
	updatedUpdater, needsUpdate, err := updater.NeedsUpdate()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !needsUpdate {
		t.Errorf("expected needsUpdate to be true")
	}
	if updatedUpdater.LastTs <= updater.LastTs {
		t.Errorf("expected updated LastTs to be greater than original")
	}
}

func TestSetChain(t *testing.T) {
	updater, err := NewUpdater("test", []string{"/oldChain/path1", "/oldChain/path2"}, File)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	updater.SetChain("oldChain", "newChain")
	expectedPaths := []string{"/newChain/path1", "/newChain/path2"}
	for i, path := range updater.Paths {
		if path != expectedPaths[i] {
			t.Errorf("expected path %s, got %s", expectedPaths[i], path)
		}
	}
}

func TestReset(t *testing.T) {
	updater, err := NewUpdater("test", []string{"path1"}, File)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	updater.LastTs = time.Now().Unix()
	updater.Reset()
	if updater.LastTs != 0 {
		t.Errorf("expected LastTs to be reset to 0, got %d", updater.LastTs)
	}
}

func TestNewUpdaterWithSize(t *testing.T) {
	updater, err := NewUpdater("test", []string{"path1"}, Size)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(updater.Paths) != 1 {
		t.Errorf("expected 1 path, got %d", len(updater.Paths))
	}
	if updater.UpdateType != Size {
		t.Errorf("expected UpdateType to be Size, got %v", updater.UpdateType)
	}
	if updater.Duration != 0 {
		t.Errorf("expected duration to be 0, got %v", updater.Duration)
	}
	if updater.LastTs == 0 {
		t.Errorf("expected LastTs to be non-zero, got %d", updater.LastTs)
	}
}
