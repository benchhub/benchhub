package docker

// TODO: get container stats
import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/dyweb/gommon/util/testutil"
)

// https://docs.docker.com/develop/sdk/examples/#list-all-images
func TestContainer_List(t *testing.T) {
	testutil.RunIf(t, testutil.IsTravis())

	c, err := client.NewEnvClient()
	if err != nil {
		t.Fatalf("failed to create docker client %v", err)
	}
	containers, err := c.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		t.Fatalf("failed to list images %v", err)
	}
	for _, container := range containers {
		fmt.Printf("id %s image %s cmd %s created %d status %s ports %v name %v \n",
			container.ID, container.Image, container.Command, container.Created, container.Status, container.Ports, container.Names)
	}
	// get stats
	for _, container := range containers {
		// TODO: it's possible to use stream and use json Decoder, result is one line with \n
		// TODO: it seems this request took a while to finish ....
		res, err := c.ContainerStats(context.Background(), container.ID, false)
		if err != nil {
			t.Fatalf("failed to get stats %v", err)
		}
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("failed to read stats body %v", err)
		}
		var stats types.StatsJSON
		if err := json.Unmarshal(b, &stats); err != nil {
			t.Fatalf("failed to unmarshal JSON %v", err)
		}
		fmt.Println(string(b))
	}
}
