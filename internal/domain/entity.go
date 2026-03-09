package domain

import "maps"

type RemoteFilesSet map[RemoteFileAddr]interface{}

type RemoteFiles struct {
	urls RemoteFilesSet
}

func NewRemoteFiles() *RemoteFiles {
	return &RemoteFiles{
		urls: make(RemoteFilesSet),
	}
}

func (fa *RemoteFiles) Add(fileAddr RemoteFileAddr) {
	fa.urls[fileAddr] = nil
}

func (fa *RemoteFiles) AddAll(fileAddrs []RemoteFileAddr) {
	for _, fileAddr := range fileAddrs {
		fa.Add(fileAddr)
	}
}

func (fa *RemoteFiles) GetUrls() []RemoteFileAddr {
	result := make([]RemoteFileAddr, 0)
	for key := range maps.Keys(fa.urls) {
		result = append(result, key)
	}
	return result
}
