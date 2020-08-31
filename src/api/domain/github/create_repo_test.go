package github

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestCreateRepoRequestAsJSON tests converting CreateRepoRequest to and from json properly
func TestCreateRepoRequestAsJSON(t *testing.T) {
	request := CreateRepoRequest{
		Name:        	"golang introduction",
		Description: 	"a golang introduction repository",
		Homepage:    	"https://github.com",
		Private:     	true,
		HasIssues:   	false,
		HasProjects: 	true,
		HasWiki:     	false,
	}

	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	var target CreateRepoRequest
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)
	assert.EqualValues(t, target.Name, request.Name)
	assert.EqualValues(t, target.HasIssues, request.HasIssues)
}