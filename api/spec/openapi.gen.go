// Package spec provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package spec

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	externalRef0 "github.com/trustbloc/vcs/pkg/restapi/v1/common"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	pathPrefix := path.Dir(pathToFile)

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(pathPrefix, "./common.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
