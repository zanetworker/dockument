package labels

import (
	"context"
	"strings"

	"github.com/docker/docker/api/types"
	dockerclient "github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"
	"github.com/zanetworker/dockument/pkg/utils"
)

func getImageLabels(imageName string) (map[string]string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	d, err := dockerclient.NewEnvClient()
	if err != nil {
		return nil, err
	}

	images, err := d.ImageList(ctx, types.ImageListOptions{All: false})
	if err != nil {
		return nil, err
	}

	for _, image := range images {
		if containsImageName(image.RepoTags, imageName) {
			return image.Labels, nil
		}
	}

	log.Warn(utils.ColorString("green", "Ooops! Could not find the Image you mentioned, please retry!"))
	return nil, nil
}

// Checks if RepoTags of an image contains the requested image tag
func containsImageName(imageRepoTags []string, requestedImage string) bool {
	for _, tag := range imageRepoTags {
		nameWithoutTag := strings.Split(tag, ":")[0]
		if tag == requestedImage || nameWithoutTag == requestedImage {
			return true
		}
	}
	return false
}

//ImageExists checks if a certain image exists in the local repo
func ImageExists(imageName string) bool {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	d, err := dockerclient.NewEnvClient()
	if err != nil {
		log.Fatalf("Failed to create docker socket, error: %s", err.Error())
	}

	images, err := d.ImageList(ctx, types.ImageListOptions{All: false})
	if err != nil {
		log.Fatalf("Failed to list docker images, error: %s", err.Error())
	}

	for _, image := range images {
		if containsImageName(image.RepoTags, imageName) {
			return true
		}
	}
	return false
}
