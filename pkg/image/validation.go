package image

import (
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api/errors"
	"github.com/openshift/origin/pkg/image/api"
)

// ValidateImage tests required fields for an Image.
func ValidateImage(image *api.Image) errors.ErrorList {
	result := errors.ErrorList{}

	if len(image.ID) == 0 {
		result = append(result, errors.NewFieldRequired("ID", image.ID))
	}

	if len(image.DockerImageReference) == 0 {
		result = append(result, errors.NewFieldRequired("DockerImageReference", image.DockerImageReference))
	}

	return result
}

// ValidateImageRepositoryMapping tests required fields for an ImageRepositoryMapping.
func ValidateImageRepositoryMapping(mapping *api.ImageRepositoryMapping) errors.ErrorList {
	result := errors.ErrorList{}

	if len(mapping.DockerImageRepository) == 0 {
		result = append(result, errors.NewFieldRequired("DockerImageRepository", mapping.DockerImageRepository))
	}

	if len(mapping.Tag) == 0 {
		result = append(result, errors.NewFieldRequired("Tag", mapping.Tag))
	}

	for _, err := range ValidateImage(&mapping.Image).Prefix("image") {
		result = append(result, err)
	}

	return result
}
