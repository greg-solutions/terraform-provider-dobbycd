package client

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/mitchellh/mapstructure"
	"github.com/vadimDidenko/terraform-provider-dobbycd/client/model"

	"gitlab.com/gregsolutions/dobby-cd/common/api"
)

// Loggers
var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

type DobbyCDApi struct {
	client *Requester
}

func NewApi(url string, user, password string) *DobbyCDApi {
	return &DobbyCDApi{
		client: &Requester{
			BasicAuth: &BasicAuth{
				Username: user,
				Password: password,
			},
			Base:   url,
			Client: http.DefaultClient,
		},
	}
}

func (d *DobbyCDApi) CreateProject(name string) (*model.ProjectResponse, error) {
	request := &model.ProjectRequest{Name: name}
	data, _ := request.MarshalBinary()
	response := &api.Response{}

	_, err := d.client.Post("/project", bytes.NewBuffer(data), response, nil)
	if err != nil {
		return nil, err
	}
	mod := model.ProjectResponse{}

	return &mod, mapstructure.Decode(response.Payload, &mod)
}

func (d *DobbyCDApi) GetProject(id string) (*model.ProjectResponse, error) {
	response := &api.Response{}
	_, err := d.client.Get(fmt.Sprintf("/project/%s", id), response, nil)
	if err != nil {
		return nil, err
	}

	mod := model.ProjectResponse{}

	return &mod, mapstructure.Decode(response.Payload, &mod)
}

func (d *DobbyCDApi) UpdateProject(id, name string) (*model.ProjectResponse, error) {

	request := &model.ProjectRequest{Name: name}
	data, _ := request.MarshalBinary()

	response := &api.Response{}
	_, err := d.client.Put(fmt.Sprintf("/project/%s", id), bytes.NewBuffer(data), response, nil)
	if err != nil {
		return nil, err
	}

	mod := model.ProjectResponse{}

	return &mod, mapstructure.Decode(response.Payload, &mod)
}
func (d *DobbyCDApi) DeleteProject(id string) error {

	response := &api.Response{}
	_, err := d.client.Delete(fmt.Sprintf("/project/%s", id), response, nil)
	if err != nil {
		return err
	}

	return nil
}

func (d *DobbyCDApi) CreatePipeline(projectId string, pipeline *model.PipelineRequest) (*model.PipelineResponse, error) {

	data, _ := pipeline.MarshalBinary()

	response := &api.Response{}

	requestString := fmt.Sprintf("/project/%s/pipeline", projectId)
	_, err := d.client.Post(requestString, bytes.NewBuffer(data), response, nil)
	if err != nil {
		return nil, err
	}

	mod := &model.PipelineResponse{}

	return mod, mapstructure.Decode(response.Payload, mod)
}

func (d *DobbyCDApi) GetGlobalPermissions() (*model.PermissionResponse, error) {
	response := &api.Response{}
	_, err := d.client.Get("/admin/permission", response, nil)
	if err != nil {
		return nil, err
	}

	mod := model.PermissionResponse{
		Permits:make([]model.Permit,0),
	}

	return &mod, mapstructure.Decode(response.Payload, &mod.Permits)
}
