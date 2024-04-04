// Package customcamera implements a camera where all methods are unimplemented.
// It extends the built-in resource subtype camera and implements methods to handle resource construction and attribute configuration.

package viamalpr

import (
	"context"
	"errors"
	"image"

	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/services/vision"
	vis "go.viam.com/rdk/vision"
	"go.viam.com/rdk/vision/classification"
	"go.viam.com/rdk/vision/objectdetection"
	"go.viam.com/utils"
)

// Here is where we define your new model's colon-delimited-triplet (viam-labs:go-module-templates-camera:customcamera)
// viam-labs = namespace, go-module-templates-camera = repo-name, customcamera = model name.
// TODO: Change model namespace, family (often the repo-name), and model. For more information see https://docs.viam.com/registry/create/#name-your-new-resource-model
var (
	Model            = resource.NewModel("viam-soleng", "go-module-templates-camera", "customcamera")
	errUnimplemented = errors.New("unimplemented")
)

func init() {
	resource.RegisterComponent(vision.API, Model,
		resource.Registration[vision.Service, *Config]{
			Constructor: newViamAlpr,
		},
	)
}

// TODO: Change the Config struct to contain any values that you would like to be able to configure from the attributes field in the component
// configuration. For more information see https://docs.viam.com/build/configure/#components
type Config struct {
	ArgumentOne int    `json:"one"`
	ArgumentTwo string `json:"two"`
}

// Validate validates the config and returns implicit dependencies.
// TODO: Change the Validate function to validate any config variables.
func (cfg *Config) Validate(path string) ([]string, error) {
	if cfg.ArgumentOne == 0 {
		return nil, utils.NewConfigValidationFieldRequiredError(path, "one")
	}

	if cfg.ArgumentTwo == "" {
		return nil, utils.NewConfigValidationFieldRequiredError(path, "two")
	}

	// TODO: return implicit dependencies if needed as the first value
	return []string{}, nil
}

// Constructor for a custom camera that creates and returns a customCamera.
// TODO: update the customCamera struct and the initialization.
func newViamAlpr(ctx context.Context, deps resource.Dependencies, rawConf resource.Config, logger logging.Logger) (vision.Service, error) {
	// This takes the generic resource.Config passed down from the parent and converts it to the
	// model-specific (aka "native") Config structure defined above, making it easier to directly access attributes.
	conf, err := resource.NativeConfig[*Config](rawConf)
	if err != nil {
		return nil, err
	}

	// Create a cancelable context for custom camera
	cancelCtx, cancelFunc := context.WithCancel(context.Background())

	c := &viamAlpr{
		name:       rawConf.ResourceName(),
		logger:     logger,
		cfg:        conf,
		cancelCtx:  cancelCtx,
		cancelFunc: cancelFunc,
	}

	// TODO: If your custom component has dependencies, perform any checks you need to on them.

	// The Reconfigure() method changes the values on the customCamera based on the attributes in the component config
	if err := c.Reconfigure(ctx, deps, rawConf); err != nil {
		logger.Error("Error configuring module with ", err)
		return nil, err
	}

	return c, nil
}

// TODO: update the viamAlpr struct with any fields you require.
type viamAlpr struct {
	name   resource.Name
	logger logging.Logger
	cfg    *Config

	cancelCtx  context.Context
	cancelFunc func()

	argumentOne int
	argumentTwo string
}

// Name implements vision.Service.
func (c *viamAlpr) Name() resource.Name {
	panic("unimplemented")
}

// Reconfigures the model. Most models can be reconfigured in place without needing to rebuild. If you need to instead create a new instance of the camera, throw a NewMustBuildError.
func (c *viamAlpr) Reconfigure(ctx context.Context, deps resource.Dependencies, conf resource.Config) error {
	cameraConfig, err := resource.NativeConfig[*Config](conf)
	if err != nil {
		//c.logger.Warn("Error reconfiguring module with ", rawConf)
		return err
	}

	c.argumentOne = cameraConfig.ArgumentOne
	c.argumentTwo = cameraConfig.ArgumentTwo
	c.name = conf.ResourceName()
	c.logger.Info("one is now configured to: ", c.argumentOne)
	c.logger.Info("two is now configured to ", c.argumentTwo)

	return nil
}

// Classifications implements vision.Service.
func (c *viamAlpr) Classifications(ctx context.Context, img image.Image, n int, extra map[string]interface{}) (classification.Classifications, error) {
	return nil, errUnimplemented
}

// ClassificationsFromCamera implements vision.Service.
func (c *viamAlpr) ClassificationsFromCamera(ctx context.Context, cameraName string, n int, extra map[string]interface{}) (classification.Classifications, error) {
	return nil, errUnimplemented
}

// Detections implements vision.Service.
func (c *viamAlpr) Detections(ctx context.Context, img image.Image, extra map[string]interface{}) ([]objectdetection.Detection, error) {
	return nil, errUnimplemented
}

// DetectionsFromCamera implements vision.Service.
func (c *viamAlpr) DetectionsFromCamera(ctx context.Context, cameraName string, extra map[string]interface{}) ([]objectdetection.Detection, error) {
	return nil, errUnimplemented
}

// GetObjectPointClouds implements vision.Service.
func (c *viamAlpr) GetObjectPointClouds(ctx context.Context, cameraName string, extra map[string]interface{}) ([]*vis.Object, error) {
	return nil, errUnimplemented
}

// DoCommand implements vision.Service.
func (c *viamAlpr) DoCommand(ctx context.Context, cmd map[string]interface{}) (map[string]interface{}, error) {
	return nil, errUnimplemented
}

// Close implements vision.Service.
func (c *viamAlpr) Close(ctx context.Context) error {
	return errUnimplemented
}
