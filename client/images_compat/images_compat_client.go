// Code generated by go-swagger; DO NOT EDIT.

package images_compat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new images compat API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for images compat API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	BuildImage(params *BuildImageParams) (*BuildImageOK, error)

	CreateImage(params *CreateImageParams) (*CreateImageOK, error)

	ExportImage(params *ExportImageParams, writer io.Writer) (*ExportImageOK, error)

	ImageHistory(params *ImageHistoryParams) (*ImageHistoryOK, error)

	ImportImage(params *ImportImageParams) (*ImportImageOK, error)

	InspectImage(params *InspectImageParams) (*InspectImageOK, error)

	ListImages(params *ListImagesParams) (*ListImagesOK, error)

	PruneImages(params *PruneImagesParams) (*PruneImagesOK, error)

	PushImage(params *PushImageParams, writer io.Writer) (*PushImageOK, error)

	RemoveImage(params *RemoveImageParams) (*RemoveImageOK, error)

	SearchImages(params *SearchImagesParams) (*SearchImagesOK, error)

	TagImage(params *TagImageParams) (*TagImageCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BuildImage creates image

  Build an image from the given Dockerfile(s)
*/
func (a *Client) BuildImage(params *BuildImageParams) (*BuildImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBuildImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "buildImage",
		Method:             "POST",
		PathPattern:        "/build",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BuildImageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BuildImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for buildImage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateImage creates an image

  Create an image by either pulling it from a registry or importing it.
*/
func (a *Client) CreateImage(params *CreateImageParams) (*CreateImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createImage",
		Method:             "POST",
		PathPattern:        "/images/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateImageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createImage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExportImage exports an image

  Export an image in tarball format
*/
func (a *Client) ExportImage(params *ExportImageParams, writer io.Writer) (*ExportImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "exportImage",
		Method:             "GET",
		PathPattern:        "/images/{name:.*}/get",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ExportImageReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ExportImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for exportImage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ImageHistory histories of an image

  Return parent layers of an image.
*/
func (a *Client) ImageHistory(params *ImageHistoryParams) (*ImageHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImageHistoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "imageHistory",
		Method:             "GET",
		PathPattern:        "/images/{name:.*}/history",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ImageHistoryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ImageHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for imageHistory: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ImportImage imports image

  Load a set of images and tags into a repository.
*/
func (a *Client) ImportImage(params *ImportImageParams) (*ImportImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "importImage",
		Method:             "POST",
		PathPattern:        "/images/load",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ImportImageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ImportImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for importImage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  InspectImage inspects an image

  Return low-level information about an image.
*/
func (a *Client) InspectImage(params *InspectImageParams) (*InspectImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInspectImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "inspectImage",
		Method:             "GET",
		PathPattern:        "/images/{name:.*}/json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InspectImageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*InspectImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for inspectImage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListImages lists images

  Returns a list of images on the server. Note that it uses a different, smaller representation of an image than inspecting a single image.
*/
func (a *Client) ListImages(params *ListImagesParams) (*ListImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListImagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listImages",
		Method:             "GET",
		PathPattern:        "/images/json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListImagesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listImages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PruneImages prunes unused images

  Remove images from local storage that are not being used by a container
*/
func (a *Client) PruneImages(params *PruneImagesParams) (*PruneImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPruneImagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pruneImages",
		Method:             "POST",
		PathPattern:        "/images/prune",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PruneImagesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PruneImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pruneImages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PushImage pushes image

  Push an image to a container registry
*/
func (a *Client) PushImage(params *PushImageParams, writer io.Writer) (*PushImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPushImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "pushImage",
		Method:             "POST",
		PathPattern:        "/images/{name:.*}/push",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PushImageReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PushImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pushImage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RemoveImage removes image

  Delete an image from local storage
*/
func (a *Client) RemoveImage(params *RemoveImageParams) (*RemoveImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "removeImage",
		Method:             "DELETE",
		PathPattern:        "/images/{name:.*}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveImageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RemoveImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeImage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SearchImages searches images

  Search registries for an image
*/
func (a *Client) SearchImages(params *SearchImagesParams) (*SearchImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchImagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "searchImages",
		Method:             "GET",
		PathPattern:        "/images/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchImagesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SearchImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for searchImages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TagImage tags an image

  Tag an image so that it becomes part of a repository.
*/
func (a *Client) TagImage(params *TagImageParams) (*TagImageCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTagImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tagImage",
		Method:             "POST",
		PathPattern:        "/images/{name:.*}/tag",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TagImageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TagImageCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tagImage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
