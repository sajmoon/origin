package template

import (
	"testing"
	"time"

	kubeapi "github.com/GoogleCloudPlatform/kubernetes/pkg/api"
)

func TestNewStorageInvalidType(t *testing.T) {
	storage := NewStorage()
	_, err := storage.Create("string")
	if err == nil {
		t.Errorf("Expected type error.")
	}
}

func TestStorageNotImplementedFunctions(t *testing.T) {
	storage := NewStorage()

	if _, err := storage.List(nil); err == nil {
		t.Errorf("Expected not implemented error.")
	}

	if _, err := storage.Get(""); err == nil {
		t.Errorf("Expected not implemented error.")
	}

	if _, err := storage.Update(nil); err == nil {
		t.Errorf("Expected not implemented error.")
	}

	channel, err := storage.Delete("")
	if err != nil {
		t.Errorf("Unexpected error when deleting: %v", err)
	}
	select {
	case result := <-channel:
		status, ok := result.(*kubeapi.Status)
		if !ok || status.Status != kubeapi.StatusFailure {
			t.Errorf("Expected not implemented error.")
		}
	case <-time.After(time.Millisecond * 100):
		t.Error("Unexpected timeout from async channel")
	}
}
