package updater

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/walk"
)

func TestNewUpdaterWithPaths(t *testing.T) {
	items := []UpdaterItem{
		{Path: "path1", Type: File},
		{Path: "path2", Type: File},
	}
	updater, err := NewUpdater("test", items)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(updater.Items) != 2 {
		t.Errorf("expected 2 items, got %d", len(updater.Items))
	}
	if updater.Items[0].Type != File || updater.Items[1].Type != File {
		t.Errorf("expected UpdateType to be File, got %v and %v", updater.Items[0].Type, updater.Items[1].Type)
	}
	if updater.LastTimeStamp == 0 {
		t.Errorf("expected LastTimeStamp to be non-zero, got %d", updater.LastTimeStamp)
	}
}

func TestNewUpdaterWithDuration(t *testing.T) {
	duration := 10 * time.Second
	items := []UpdaterItem{
		{Duration: duration, Type: Timer},
	}
	updater, err := NewUpdater("test", items)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(updater.Items) != 1 {
		t.Errorf("expected 1 item, got %d", len(updater.Items))
	}
	if updater.Items[0].Type != Timer {
		t.Errorf("expected UpdateType to be Timer, got %v", updater.Items[0].Type)
	}
	if updater.Items[0].Duration != duration {
		t.Errorf("expected duration %v, got %v", duration, updater.Items[0].Duration)
	}
	if updater.LastTimeStamp == 0 {
		t.Errorf("expected LastTimeStamp to be non-zero, got %d", updater.LastTimeStamp)
	}
}

func TestNewUpdaterWithPathsAndDuration(t *testing.T) {
	duration := 15 * time.Second
	items := []UpdaterItem{
		{Path: "path1", Type: Folder, Duration: duration},
	}
	updater, err := NewUpdater("test", items)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(updater.Items) != 1 {
		t.Errorf("expected 1 item, got %d", len(updater.Items))
	}
	if updater.Items[0].Path != "path1" {
		t.Errorf("expected path 'path1', got %s", updater.Items[0].Path)
	}
	if updater.Items[0].Type != Folder {
		t.Errorf("expected UpdateType to be Folder, got %v", updater.Items[0].Type)
	}
	if updater.Items[0].Duration != duration {
		t.Errorf("expected duration %v, got %v", duration, updater.Items[0].Duration)
	}
	if updater.LastTimeStamp == 0 {
		t.Errorf("expected LastTimeStamp to be non-zero, got %d", updater.LastTimeStamp)
	}
}

func TestNeedsUpdateWithDuration(t *testing.T) {
	duration := 10 * time.Second
	items := []UpdaterItem{
		{Duration: duration, Type: Timer},
	}
	updater, err := NewUpdater("test", items)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	updater.LastTimeStamp = time.Now().Add(-20 * time.Second).Unix()
	updatedUpdater, needsUpdate, err := updater.NeedsUpdate()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !needsUpdate {
		t.Errorf("expected needsUpdate to be true")
	}
	if updatedUpdater.LastTimeStamp <= updater.LastTimeStamp {
		t.Errorf("expected updated LastTimeStamp to be greater than original")
	}
}

func TestSetChain(t *testing.T) {
	items := []UpdaterItem{
		{Path: "/oldChain/path1", Type: File},
		{Path: "/oldChain/path2", Type: File},
	}
	updater, err := NewUpdater("test", items)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	updater.SetChain("oldChain", "newChain")
	expectedPaths := []string{"/newChain/path1", "/newChain/path2"}
	for i, item := range updater.Items {
		if item.Path != expectedPaths[i] {
			t.Errorf("expected path %s, got %s", expectedPaths[i], item.Path)
		}
	}
}

func TestReset(t *testing.T) {
	items := []UpdaterItem{
		{Path: "path1", Type: File},
	}
	updater, err := NewUpdater("test", items)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	updater.LastTimeStamp = time.Now().Unix()
	updater.Reset()
	if updater.LastTimeStamp != 0 {
		t.Errorf("expected LastTimeStamp to be reset to 0, got %d", updater.LastTimeStamp)
	}
	if updater.LastTotalSize != 0 {
		t.Errorf("expected LastTotalSize to be reset to 0, got %d", updater.LastTotalSize)
	}
}

func TestNewUpdaterWithSize(t *testing.T) {
	items := []UpdaterItem{
		{Path: "path1", Type: FileSize},
	}
	updater, err := NewUpdater("test", items)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(updater.Items) != 1 {
		t.Errorf("expected 1 item, got %d", len(updater.Items))
	}
	if updater.Items[0].Type != FileSize {
		t.Errorf("expected UpdateType to be FileSize, got %v", updater.Items[0].Type)
	}
	if updater.LastTimeStamp != -1 {
		t.Errorf("expected LastTimeStamp to be -1, got %d", updater.LastTimeStamp)
	}
	if updater.LastTotalSize != -1 {
		t.Errorf("expected LastTotalSize to be -1, got %d", updater.LastTotalSize)
	}
}

func TestNeedsUpdateWithShortDuration(t *testing.T) {
	duration := 1 * time.Second
	items := []UpdaterItem{
		{Duration: duration, Type: Timer},
	}
	updater, err := NewUpdater("test", items)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Set LastTimeStamp to a past time
	updater.LastTimeStamp = time.Now().Add(-2 * time.Second).Unix()

	// Wait for a short duration to ensure the timer condition is met
	time.Sleep(2 * time.Second)

	updatedUpdater, needsUpdate, err := updater.NeedsUpdate()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !needsUpdate {
		t.Errorf("expected needsUpdate to be true")
	}
	if updatedUpdater.LastTimeStamp <= updater.LastTimeStamp {
		t.Errorf("expected updated LastTimeStamp to be greater than original")
	}
}

func TestNeedsUpdateWithSize(t *testing.T) {
	// Create a temporary file
	tempFile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up

	// Write some data to the file to set its size
	if _, err := tempFile.Write([]byte("initial data")); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	items := []UpdaterItem{
		{Path: tempFile.Name(), Type: FileSize},
	}
	updater, err := NewUpdater("test", items)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Set LastTotalSize to the current size of the file
	updater.LastTotalSize = fileSize(tempFile.Name())

	// Write more data to the file to increase its size
	if _, err := tempFile.Write([]byte(" more data")); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	updatedUpdater, needsUpdate, err := updater.NeedsUpdate()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !needsUpdate {
		t.Errorf("expected needsUpdate to be true")
	}
	if updatedUpdater.LastTotalSize <= updater.LastTotalSize {
		t.Errorf("expected updated LastTotalSize to be greater than original")
	}
}

// Helper function to get the size of a file
func fileSize(path string) int64 {
	info, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return info.Size()
}

func TestNeedsUpdateWithFile(t *testing.T) {
	// Create a temporary file
	tempFile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up

	// Write some data to the file to set its modification time
	if _, err := tempFile.Write([]byte("initial data")); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	items := []UpdaterItem{
		{Path: tempFile.Name(), Type: File},
	}
	updater, err := NewUpdater("test", items)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Set LastTimeStamp to the current modification time of the file
	updater.LastTimeStamp = fileModTime(tempFile.Name())

	// Modify the file to update its modification time
	if _, err := tempFile.Write([]byte(" more data")); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Explicitly update the file's modification time
	newModTime := time.Now().Add(time.Second)
	if err := os.Chtimes(tempFile.Name(), newModTime, newModTime); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	updatedUpdater, needsUpdate, err := updater.NeedsUpdate()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !needsUpdate {
		t.Errorf("expected needsUpdate to be true")
	}
	if updatedUpdater.LastTimeStamp <= updater.LastTimeStamp {
		t.Errorf("expected updated LastTimeStamp to be greater than original")
	}
}

// Helper function to get the modification time of a file
func fileModTime(path string) int64 {
	info, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return info.ModTime().Unix()
}

func TestNeedsUpdateWithFolder(t *testing.T) {
	// Create a temporary directory
	tempDir, err := ioutil.TempDir("", "testdir")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up

	// Create a temporary file in the directory
	tempFile, err := ioutil.TempFile(tempDir, "testfile")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Write some data to the file to set its modification time
	if _, err := tempFile.Write([]byte("initial data")); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	items := []UpdaterItem{
		{Path: tempDir, Type: Folder},
	}
	updater, err := NewUpdater("test", items)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Set LastTimeStamp to the current modification time of the file
	updater.LastTimeStamp = fileModTime(tempFile.Name())

	// Modify the file to update its modification time
	if _, err := tempFile.Write([]byte(" more data")); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Explicitly update the file's modification time
	newModTime := time.Now().Add(time.Second)
	if err := os.Chtimes(tempFile.Name(), newModTime, newModTime); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	updatedUpdater, needsUpdate, err := updater.NeedsUpdate()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !needsUpdate {
		t.Errorf("expected needsUpdate to be true")
	}
	if updatedUpdater.LastTimeStamp <= updater.LastTimeStamp {
		t.Errorf("expected updated LastTimeStamp to be greater than original")
	}
}

func TestNeedsUpdateWithError(t *testing.T) {
	// Create a temporary file
	tempFile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up

	// Write some data to the file to set its size
	if _, err := tempFile.Write([]byte("initial data")); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	items := []UpdaterItem{
		{Path: tempFile.Name(), Type: File},
	}
	updater, err := NewUpdater("test", items)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Simulate an error by setting an invalid path
	updater.Items[0].Path = "/invalid/path"

	_, _, err = updater.NeedsUpdate()
	if err == nil {
		t.Errorf("expected an error, got nil")
	}
}

func TestNeedsUpdateWithFolderSize(t *testing.T) {
	// Create a temporary directory
	tempDir, err := ioutil.TempDir("", "testdir")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up

	// Create a temporary file in the directory
	tempFile, err := ioutil.TempFile(tempDir, "testfile")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Write some data to the file to set its size
	if _, err := tempFile.Write([]byte("initial data")); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	items := []UpdaterItem{
		{Path: tempDir, Type: FolderSize},
	}
	updater, err := NewUpdater("test", items)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Set LastTotalSize to the current size of the folder
	updater.LastTotalSize = folderSize(tempDir)

	// Write more data to the file to increase its size
	if _, err := tempFile.Write([]byte(" more data")); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	updatedUpdater, needsUpdate, err := updater.NeedsUpdate()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !needsUpdate {
		t.Errorf("expected needsUpdate to be true")
	}
	if updatedUpdater.LastTotalSize <= updater.LastTotalSize {
		t.Errorf("expected updated LastTotalSize to be greater than original")
	}
}

// Helper function to get the total size of all files in a folder
func folderSize(path string) int64 {
	var totalSize int64
	err := walk.ForEveryFileInFolder(path, func(filePath string, _ any) (bool, error) {
		info, err := os.Stat(filePath)
		if err != nil {
			return false, err
		}
		totalSize += info.Size()
		return true, nil
	}, nil)
	if err != nil {
		return 0
	}
	return totalSize
}

func TestNeedsUpdateWithSizeAndFolderSize(t *testing.T) {
	// Create a temporary file
	tempFile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up

	// Write some data to the file to set its size
	if _, err := tempFile.Write([]byte("initial data")); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Create a temporary directory
	tempDir, err := ioutil.TempDir("", "testdir")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up

	// Create a temporary file in the directory
	tempFileInDir, err := ioutil.TempFile(tempDir, "testfile")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Write some data to the file in the directory to set its size
	if _, err := tempFileInDir.Write([]byte("initial data")); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	items := []UpdaterItem{
		{Path: tempFile.Name(), Type: FileSize},
		{Path: tempDir, Type: FolderSize},
	}
	updater, err := NewUpdater("test", items)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Set LastTotalSize to the current size of the file and folder
	updater.LastTotalSize = fileSize(tempFile.Name()) + folderSize(tempDir)

	// Write more data to the file to increase its size
	if _, err := tempFile.Write([]byte(" more data")); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Write more data to the file in the directory to increase its size
	if _, err := tempFileInDir.Write([]byte(" more data")); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	updatedUpdater, needsUpdate, err := updater.NeedsUpdate()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !needsUpdate {
		t.Errorf("expected needsUpdate to be true")
	}
	if updatedUpdater.LastTotalSize <= updater.LastTotalSize {
		t.Errorf("expected updated LastTotalSize to be greater than original")
	}

	// Verify that LastTotalSize is updated correctly
	if updatedUpdater.LastTotalSize != fileSize(tempFile.Name())+folderSize(tempDir) {
		t.Errorf("expected LastTotalSize to be updated correctly")
	}

	// Check for unexpected resets
	if updatedUpdater.LastTotalSize == 0 {
		t.Errorf("unexpected reset of LastTotalSize")
	}
}
