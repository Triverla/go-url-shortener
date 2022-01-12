package logic

import (
	"github.com/Triverla/go-url-shortener/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStoreService = &store.StorageService{}

func init() {
	testStoreService = store.InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.RedisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://google.com"
	userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortURL := "jTa4L57P"

	// Persist data mapping
	store.SaveUrlMapping(shortURL, initialLink, userUUId)

	// Retrieve initial URL
	retrievedUrl := store.RetrieveInitialUrl(shortURL)

	assert.Equal(t, initialLink, retrievedUrl)
}
