// Package common provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package common

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Defines values for DIDMethod.
const (
	DIDMethodKey DIDMethod = "key"
	DIDMethodOrb DIDMethod = "orb"
	DIDMethodWeb DIDMethod = "web"
)

// Defines values for KMSConfigType.
const (
	KMSConfigTypeAws   KMSConfigType = "aws"
	KMSConfigTypeLocal KMSConfigType = "local"
	KMSConfigTypeWeb   KMSConfigType = "web"
)

// Defines values for VCFormat.
const (
	JwtVcJson   VCFormat = "jwt_vc_json"
	JwtVcJsonLd VCFormat = "jwt_vc_json-ld"
	LdpVc       VCFormat = "ldp_vc"
)

// Defines values for VPFormat.
const (
	JwtVp VPFormat = "jwt_vp"
	LdpVp VPFormat = "ldp_vp"
)

// Model to convey the details about the Credentials the Client wants to obtain.
type AuthorizationDetails struct {
	// REQUIRED when Format parameter is not present. String specifying a unique identifier of the Credential being described in the credential_configurations_supported map in the Credential Issuer Metadata. The referenced object in the credential_configurations_supported map conveys the details, such as the format, for issuance of the requested Credential. It MUST NOT be present if format parameter is present.
	CredentialConfigurationId *string `json:"credential_configuration_id,omitempty"`

	// Object containing the detailed description of the credential type.
	CredentialDefinition *CredentialDefinition `json:"credential_definition,omitempty"`

	// REQUIRED when CredentialConfigurationId parameter is not present. String identifying the format of the Credential the Wallet needs. This Credential format identifier determines further claims in the authorization details object needed to identify the Credential type in the requested format. It MUST NOT be present if credential_configuration_id parameter is present.
	Format *string `json:"format,omitempty"`

	// An array of strings that allows a client to specify the location of the resource server(s) allowing the Authorization Server to mint audience restricted access tokens.
	Locations *[]string `json:"locations,omitempty"`

	// String that determines the authorization details type. MUST be set to "openid_credential" for OIDC4VC.
	Type string `json:"type"`
}

// Object containing the detailed description of the credential type.
type CredentialDefinition struct {
	// For ldp_vc only. Array as defined in https://www.w3.org/TR/vc-data-model/#contexts.
	Context *[]string `json:"@context,omitempty"`

	// An object containing a list of name/value pairs, where each name identifies a claim offered in the Credential. The value can be another such object (nested data structures), or an array of such objects.
	CredentialSubject *map[string]interface{} `json:"credentialSubject,omitempty"`

	// Array designating the types a certain credential type supports
	Type []string `json:"type"`
}

// DID method of the DID to be used for signing.
type DIDMethod string

// Model for KMS configuration.
type KMSConfig struct {
	// Prefix of database used by local kms.
	DbPrefix *string `json:"dbPrefix,omitempty"`

	// Type of database used by local kms.
	DbType *string `json:"dbType,omitempty"`

	// URL to database used by local kms.
	DbURL *string `json:"dbURL,omitempty"`

	// KMS endpoint.
	Endpoint *string `json:"endpoint,omitempty"`

	// Path to secret lock used by local kms.
	SecretLockKeyPath *string `json:"secretLockKeyPath,omitempty"`

	// Type of kms used to create and store DID keys.
	Type KMSConfigType `json:"type"`
}

// Type of kms used to create and store DID keys.
type KMSConfigType string

// Supported VC formats.
type VCFormat string

// Supported VP formats.
type VPFormat string

// WalletInitiatedFlowData defines model for WalletInitiatedFlowData.
type WalletInitiatedFlowData struct {
	ClaimEndpoint        string    `json:"claim_endpoint"`
	CredentialTemplateId string    `json:"credential_template_id"`
	OpState              string    `json:"op_state"`
	ProfileId            string    `json:"profile_id"`
	ProfileVersion       string    `json:"profile_version"`
	Scopes               *[]string `json:"scopes"`
}

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
