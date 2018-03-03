package docker

import (
	"context"
	"fmt"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// https://docs.docker.com/develop/sdk/examples/#list-all-images
func TestImage_List(t *testing.T) {
	c, err := client.NewEnvClient()
	if err != nil {
		t.Fatalf("failed to create docker client %v", err)
	}
	images, err := c.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		t.Fatalf("failed to list images %v", err)
	}
	for _, image := range images {
		// NOTE: the size is a smaller than docker images command
		// because it is using / 1000 / 1000 instead / 1024 / 1024
		fmt.Printf("id %s repo %s size %dMB\n", image.ID, image.RepoTags, image.Size/1024/1024)
	}
}
