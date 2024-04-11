// Package customcamera implements a camera where all methods are unimplemented.
// It extends the built-in resource subtype camera and implements methods to handle resource construction and attribute configuration.

package viamalpr

import (
	"bytes"
	"context"
	"errors"
	"image"
	"image/jpeg"
	"os"
	"strings"
	"sync"

	"github.com/openalpr/openalpr/src/bindings/go/openalpr"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/services/vision"
	vis "go.viam.com/rdk/vision"
	"go.viam.com/rdk/vision/classification"
	"go.viam.com/rdk/vision/objectdetection"
)

// Here is where we define your new model's colon-delimited-triplet (viam-labs:go-module-templates-camera:customcamera)
// viam-labs = namespace, go-module-templates-camera = repo-name, customcamera = model name.
// TODO: Change model namespace, family (often the repo-name), and model. For more information see https://docs.viam.com/registry/create/#name-your-new-resource-model
var (
	Model            = resource.NewModel("viam-soleng", "vision", "openalpr")
	errUnimplemented = errors.New("unimplemented")
	PrettyName       = "Viam openalpr vision service"
	Description      = "A Viam automatic license plate recognition module based upon OpenALPR"
)

func init() {
	resource.RegisterService(vision.API, Model,
		resource.Registration[vision.Service, *Config]{
			Constructor: newViamAlpr,
		},
	)
}

// TODO: Change the Config struct to contain any values that you would like to be able to configure from the attributes field in the component
// configuration. For more information see https://docs.viam.com/build/configure/#components
type Config struct {
	Country    string `json:"country"`
	ConfigFile string `json:"config_file"`
	RuntimeDir string `json:"runtime_dir"`
}

// Validate validates the config and returns implicit dependencies.
// TODO: Change the Validate function to validate any config variables.
func (cfg *Config) Validate(path string) ([]string, error) {
	/*
		if cfg.ArgumentOne == 0 {
			return nil, utils.NewConfigValidationFieldRequiredError(path, "one")
		}
	*/

	// TODO: return implicit dependencies if needed as the first value
	return []string{}, nil
}

// Constructor for a custom camera that creates and returns a customCamera.
// TODO: update the customCamera struct and the initialization.
func newViamAlpr(ctx context.Context, deps resource.Dependencies, rawConf resource.Config, logger logging.Logger) (vision.Service, error) {
	// This takes the generic resource.Config passed down from the parent and converts it to the
	// model-specific (aka "native") Config structure defined above, making it easier to directly access attributes.
	logger.Debugf("Starting %s %s", PrettyName)
	conf, err := resource.NativeConfig[*Config](rawConf)
	if err != nil {
		return nil, err
	}

	// Create a cancelable context for custom camera
	cancelCtx, cancelFunc := context.WithCancel(context.Background())

	viamAlpr := &viamAlpr{
		name:       rawConf.ResourceName(),
		logger:     logger,
		cfg:        conf,
		cancelCtx:  cancelCtx,
		cancelFunc: cancelFunc,
		mu:         sync.RWMutex{},
		done:       make(chan bool),
	}

	// TODO: If your custom component has dependencies, perform any checks you need to on them.

	// The Reconfigure() method changes the values on the customCamera based on the attributes in the component config
	if err := viamAlpr.Reconfigure(ctx, deps, rawConf); err != nil {
		logger.Error("Error configuring module with ", err)
		return nil, err
	}

	return viamAlpr, nil
}

// TODO: update the viamAlpr struct with any fields you require.
type viamAlpr struct {
	name   resource.Name
	logger logging.Logger
	cfg    *Config

	cancelCtx  context.Context
	cancelFunc func()
	country    string
	configFile string
	runtimeDir string
	mu         sync.RWMutex
	done       chan bool

	alpr openalpr.Alpr
}

// Name implements vision.Service.
func (va *viamAlpr) Name() resource.Name {
	panic("unimplemented")
}

// Reconfigures the model. Most models can be reconfigured in place without needing to rebuild. If you need to instead create a new instance of the camera, throw a NewMustBuildError.
func (va *viamAlpr) Reconfigure(ctx context.Context, deps resource.Dependencies, conf resource.Config) error {
	va.mu.Lock()
	defer va.mu.Unlock()

	// TODO: Make NewAlpr configurable
	newConf, err := resource.NativeConfig[*Config](conf)
	if err != nil {
		return err
	}
	if newConf.Country != "" {
		va.country = newConf.Country
	} else {
		va.country = "us"
	}
	if newConf.ConfigFile != "" {
		va.configFile = newConf.ConfigFile
	} else {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		va.configFile = strings.TrimSuffix(wd, "/bin") + "/openalpr/openalpr.conf.user"
		va.logger.Infof("openalpr config file: %s", va.configFile)
	}
	if newConf.RuntimeDir != "" {
		va.runtimeDir = newConf.RuntimeDir
	} else {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		va.runtimeDir = strings.TrimSuffix(wd, "/bin") + "/openalpr/runtime_data"
		va.logger.Infof("runtime_data directory: %s", va.runtimeDir)
	}
	va.alpr = *openalpr.NewAlpr(va.country, va.configFile, va.runtimeDir)
	if !va.alpr.IsLoaded() {
		return errors.New("openalpr failed to load")
	}
	va.alpr.SetTopN(20)
	va.logger.Debugf("openalpr version: %v", openalpr.GetVersion())
	va.name = conf.ResourceName()
	va.logger.Debug("**** Reconfigured ****")
	return nil
}

// Classifications implements vision.Service.
func (va *viamAlpr) Classifications(ctx context.Context, img image.Image, n int, extra map[string]interface{}) (classification.Classifications, error) {
	return nil, errUnimplemented
}

// ClassificationsFromCamera implements vision.Service.
func (va *viamAlpr) ClassificationsFromCamera(ctx context.Context, cameraName string, n int, extra map[string]interface{}) (classification.Classifications, error) {
	return nil, errUnimplemented
}

// Detections implements vision.Service.
func (va *viamAlpr) Detections(ctx context.Context, img image.Image, extra map[string]interface{}) ([]objectdetection.Detection, error) {
	detections, err := va.detectAlpr(img)
	if err != nil {
		return nil, err
	}
	return detections, nil
}

// DetectionsFromCamera implements vision.Service.
func (va *viamAlpr) DetectionsFromCamera(ctx context.Context, cameraName string, extra map[string]interface{}) ([]objectdetection.Detection, error) {
	va.detectAlpr(nil)
	return nil, nil
}

// GetObjectPointClouds implements vision.Service.
func (va *viamAlpr) GetObjectPointClouds(ctx context.Context, cameraName string, extra map[string]interface{}) ([]*vis.Object, error) {
	return nil, errUnimplemented
}

// DoCommand implements vision.Service.
func (va *viamAlpr) DoCommand(ctx context.Context, cmd map[string]interface{}) (map[string]interface{}, error) {
	return nil, errUnimplemented
}

// Close implements vision.Service.
func (va *viamAlpr) Close(ctx context.Context) error {
	va.logger.Debugf("Shutting down %s", PrettyName)
	va.alpr.Unload()
	return nil
}

func (va *viamAlpr) detectAlpr(img image.Image) ([]objectdetection.Detection, error) {
	/*
		resultFromFilePath, err := svc.alpr.RecognizeByFilePath("lp.jpg")
		if err != nil {
			fmt.Println(err)
		}
		svc.logger.Infof("Detections: %v", resultFromFilePath)
		//fmt.Printf("%+v\n", resultFromFilePath)
		//fmt.Printf("\n\n\n")
	*/
	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, img, nil)
	if err != nil {
		return nil, err
	}
	imageBytes := buf.Bytes()

	/*
		imageBytes, err := os.ReadFile("lp.jpg")
		if err != nil {
			fmt.Println(err)
		}
	*/
	resultFromBlob, err := va.alpr.RecognizeByBlob(imageBytes)
	if err != nil {
		return nil, err
	}
	va.logger.Debugf("%v", resultFromBlob)
	detections := []objectdetection.Detection{}
	for _, result := range resultFromBlob.Plates {
		minPoint := image.Point{result.PlatePoints[0].X, result.PlatePoints[0].Y}
		maxPoint := image.Point{result.PlatePoints[3].X, result.PlatePoints[3].Y}
		bbox := image.Rectangle{minPoint, maxPoint}
		detection := objectdetection.NewDetection(bbox, float64(result.TopNPlates[result.PlateIndex].OverallConfidence), result.BestPlate)
		detections = append(detections, detection)
	}
	return detections, nil
}
