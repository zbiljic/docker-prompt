package cmd

import (
	docker "github.com/fsouza/go-dockerclient"
)

var c client

func getClient() client {
	if c != nil {
		return c
	}
	c, _ = newEnvClient()
	return c
}

// client is an interface specifying the subset of
// github.com/fsouza/go-dockerclient.Client that the application uses.
type client interface {
	ListImages(opts docker.ListImagesOptions) ([]docker.APIImages, error)
}

type dockerClient struct {
	client *docker.Client
}

func newEnvClient() (client, error) {
	client, err := docker.NewClientFromEnv()
	if err != nil {
		return nil, err
	}
	return &dockerClient{client}, nil
}

func newClient(endpoint string, cert, key, ca string) (client, error) {
	client, err := docker.NewTLSClient(endpoint, cert, key, ca)
	if err != nil {
		return nil, err
	}
	return &dockerClient{client}, nil
}

func (c *dockerClient) ListImages(opts docker.ListImagesOptions) ([]docker.APIImages, error) {
	return c.client.ListImages(opts)
}
