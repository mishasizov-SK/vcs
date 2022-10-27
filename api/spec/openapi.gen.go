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
