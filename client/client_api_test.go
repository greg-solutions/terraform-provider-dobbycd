package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	model2 "gitlab.com/gregsolutions/dobby-cd/api/model"
)

var projectName string

func init() {
	projectName = "test-111"
}
func TestDobbyCDApi_CreateProject(t *testing.T) {
	client := NewApi("http://localhost:8080/v1", "user", "user")
	resp, err := client.CreateProject(projectName)
	assert.Nil(t, err)
	assert.NotNil(t, resp)

}

func TestDobbyCDApi_GetProject(t *testing.T) {
	client := NewApi("http://localhost:8080/v1", "user", "user")
	resp, err := client.GetProject("d6602a48-0c03-4a3f-b0cd-0c7c91025c0b")
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestDobbyCDApi_UpdateProject(t *testing.T) {
	client := NewApi("http://localhost:8080/v1", "user", "user")
	resp, err := client.UpdateProject("d6602a48-0c03-4a3f-b0cd-0c7c91025c0b", "test123123123")
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestDobbyCDApi_DeleteProject(t *testing.T) {
	client := NewApi("http://localhost:8080/v1", "user", "user")
	err := client.DeleteProject("d6602a48-0c03-4a3f-b0cd-0c7c91025c0b")
	assert.Nil(t, err)
}

func TestDobbyCDApi_CreatePipeline(t *testing.T) {
	client := NewApi("http://localhost:8080/v1", "user", "user")
	model := &model2.PipelineRequest{
		Path:       "/server/example/",
		Token:      "",
		Branch:     "develop",
		Repository: "",
	}
	projectId := "b9264506-e3a3-4b43-9574-b1da5d5744f1"
	resp, err := client.CreatePipeline(projectId, model)
	assert.Nil(t, err)
	assert.NotNil(t, resp)

}
