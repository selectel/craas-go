package repository

import "time"

// Repository represents an unmarshalled repository from API responses.
type Repository struct {
	// Name is the name of the repository.
	Name string `json:"name"`

	// UpdatedAt is the timestamp in UTC timezone of when the repository has been updated.
	UpdatedAt time.Time `json:"updatedAt"`

	// Size is the size of the repository layers in bytes.
	Size int64 `json:"size"`
}

// Image represents an unmarshalled image from API responses.
type Image struct {
	// Digest is the digest of the image.
	Digest string `json:"digest"`

	// CreatedAt is the timestamp in UTC timezone of when the image has been created.
	CreatedAt time.Time `json:"createdAt"`

	// Tags is the list of tags of the image.
	Tags []string `json:"tags"`

	// Size is the size of the image layers in bytes.
	Size int64 `json:"size"`

	// Layers is the list of layers of the image.
	Layers []Layer `json:"layers"`
}

// Layer represents an unmarshalled layer from API responses.
type Layer struct {
	// Digest is the digest of the layer.
	Digest string `json:"digest"`

	// Size is the size of the layer in bytes.
	Size int64 `json:"size"`
}
