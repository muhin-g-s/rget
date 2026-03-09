package domain

import (
	"testing"
)

func TestRemoteFiles(t *testing.T) {
	t.Run("NewRemoteFiles creates empty set", func(t *testing.T) {
		rf := NewRemoteFiles()
		urls := rf.GetUrls()
		if len(urls) != 0 {
			t.Errorf("NewRemoteFiles() expected empty set, got %d urls", len(urls))
		}
	})

	t.Run("Add adds single file", func(t *testing.T) {
		rf := NewRemoteFiles()
		addr, _ := NewRemoteFileAddr("https://example.com/file.txt")
		rf.Add(addr)

		urls := rf.GetUrls()
		if len(urls) != 1 {
			t.Errorf("Add() expected 1 url, got %d", len(urls))
		}
		if urls[0].Value() != "https://example.com/file.txt" {
			t.Errorf("Add() expected url https://example.com/file.txt, got %s", urls[0].Value())
		}
	})

	t.Run("AddAll adds multiple files", func(t *testing.T) {
		rf := NewRemoteFiles()
		addr1, _ := NewRemoteFileAddr("https://example.com/file1.txt")
		addr2, _ := NewRemoteFileAddr("https://example.com/file2.txt")
		rf.AddAll([]RemoteFileAddr{addr1, addr2})

		urls := rf.GetUrls()
		if len(urls) != 2 {
			t.Errorf("AddAll() expected 2 urls, got %d", len(urls))
		}
	})

	t.Run("Add does not duplicate same url", func(t *testing.T) {
		rf := NewRemoteFiles()
		addr, _ := NewRemoteFileAddr("https://example.com/file.txt")
		rf.Add(addr)
		rf.Add(addr)

		urls := rf.GetUrls()
		if len(urls) != 1 {
			t.Errorf("Add() should not duplicate, expected 1 url, got %d", len(urls))
		}
	})
}
