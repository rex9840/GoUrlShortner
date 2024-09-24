package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InnitializeStore()
}

func TestStoreInit(T *testing.T) {
	assert.True(T, testStoreService.redisClient != nil)

}

func TestInsertionAndRetrival(T *testing.T) {
        initialLink := "https://www.example.com"
        shortLink := "example"
        UserId := "eqweqwe-213-qwe"

	SaveURLMapping(shortLink, initialLink, UserId)  

	retriveUrl := RetriveOriginalUrl(shortLink)

	assert.Equal(T, initialLink, retriveUrl)
}
 
