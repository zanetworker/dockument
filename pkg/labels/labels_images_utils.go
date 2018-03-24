package labels

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	dockerclient "github.com/docker/docker/client"
)

func getImageLabels(imageName string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	d, err := dockerclient.NewEnvClient()
	if err != nil {
		return err
	}

	images, err := d.ImageList(ctx, types.ImageListOptions{All: false})
	if err != nil {
		return err
	}

	for _, image := range images {
		fmt.Printf("%#v\n", image)
	}
	return nil
}
