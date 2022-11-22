/*
Copyright Avast Software. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package oidc4ci

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/cucumber/godog"
	"github.com/google/uuid"
	"github.com/hyperledger/aries-framework-go-ext/component/vdr/orb"
	docdid "github.com/hyperledger/aries-framework-go/pkg/doc/did"
	"github.com/hyperledger/aries-framework-go/pkg/doc/jose/jwk/jwksupport"
	"github.com/hyperledger/aries-framework-go/pkg/doc/jwt"
	vdrapi "github.com/hyperledger/aries-framework-go/pkg/framework/aries/api/vdr"
	vdrpkg "github.com/hyperledger/aries-framework-go/pkg/vdr"
	"golang.org/x/oauth2"

	vccrypto "github.com/trustbloc/vcs/pkg/doc/vc/crypto"
	profileapi "github.com/trustbloc/vcs/pkg/profile"
	"github.com/trustbloc/vcs/test/bdd/pkg/bddutil"
	bddcontext "github.com/trustbloc/vcs/test/bdd/pkg/context"
)

const (
	vcsAPIGateway                       = "https://localhost:4455"
	initiateCredentialIssuanceURLFormat = vcsAPIGateway + "/issuer/profiles/%s/interactions/initiate-oidc"
	vcsAuthorizeEndpoint                = vcsAPIGateway + "/oidc/authorize"
	vcsTokenEndpoint                    = vcsAPIGateway + "/oidc/token"
	vcsCredentialEndpoint               = vcsAPIGateway + "/oidc/credential"
	oidcProviderURL                     = "https://localhost:4444"
	loginPageURL                        = "https://localhost:8099/login"
	claimDataURL                        = "https://mock-login-consent.example.com:8099/claim-data"
	didDomain                           = "https://testnet.orb.local"
	didServiceAuthToken                 = "tk1"
)

// Steps defines context for OIDC4CI scenario steps.
type Steps struct {
	bddContext          *bddcontext.BDDContext
	issuerProfile       *profileapi.Issuer
	oauthClient         *oauth2.Config // oauthClient is a public client to vcs oidc provider
	cookie              *cookiejar.Jar
	debug               bool
	initiateIssuanceURL string
	authCode            string
	token               *oauth2.Token
	credential          interface{}
}

// NewSteps returns new Steps context.
func NewSteps(ctx *bddcontext.BDDContext) (*Steps, error) {
	jar, err := cookiejar.New(&cookiejar.Options{})
	if err != nil {
		return nil, fmt.Errorf("init cookie jar: %w", err)
	}

	return &Steps{
		bddContext: ctx,
		cookie:     jar,
		debug:      false, // set to true to get request/response dumps
	}, nil
}

// RegisterSteps registers OIDC4VC scenario steps.
func (s *Steps) RegisterSteps(sc *godog.ScenarioContext) {
	sc.Step(`^issuer with id "([^"]*)" authorized as a profile user$`, s.authorizeIssuer)
	sc.Step(`^client registered as a public client to vcs oidc$`, s.registerPublicClient)
	sc.Step(`^issuer initiates credential issuance using authorization code flow$`, s.initiateCredentialIssuance)
	sc.Step(`^initiate issuance URL is returned$`, s.checkInitiateIssuanceURL)
	sc.Step(`^client requests an authorization code using data from initiate issuance URL$`, s.getAuthCode)
	sc.Step(`^user authenticates on issuer IdP$`, s.authenticateUser)
	sc.Step(`^client receives an authorization code$`, s.checkAuthCode)
	sc.Step(`^client exchanges authorization code for an access token$`, s.exchangeCodeForToken)
	sc.Step(`^client receives an access token$`, s.checkAccessToken)
	sc.Step(`^client requests credential for claim data$`, s.getCredential)
	sc.Step(`^client receives a valid credential$`, s.checkCredential)
}

func (s *Steps) authorizeIssuer(id string) error {
	issuer, ok := s.bddContext.IssuerProfiles[id]
	if !ok {
		return fmt.Errorf("issuer profile '%s' not found", id)
	}

	if issuer.OIDCConfig == nil {
		return fmt.Errorf("oidc config not set for issuer profile '%s'", id)
	}

	accessToken, err := bddutil.IssueAccessToken(context.Background(), oidcProviderURL,
		issuer.OrganizationID, "test-org-secret", []string{"org_admin"})
	if err != nil {
		return err
	}

	s.issuerProfile = issuer
	s.bddContext.Args[getOrgAuthTokenKey(issuer.OrganizationID)] = accessToken

	return nil
}

func (s *Steps) registerPublicClient() error {
	// TODO: Implement API to register public clients to vcs oidc or add support for dynamic client registration.
	// for now, oauth clients are imported into vcs from the file (./fixtures/oauth-clients/clients.json)
	s.oauthClient = &oauth2.Config{
		ClientID:    "oidc4vc_client",
		RedirectURL: "https://client.example.com/oauth/redirect",
		Scopes:      []string{"openid", "profile"},
		Endpoint: oauth2.Endpoint{
			AuthURL:   vcsAuthorizeEndpoint,
			TokenURL:  vcsTokenEndpoint,
			AuthStyle: oauth2.AuthStyleInHeader,
		},
	}

	return nil
}

func (s *Steps) initiateCredentialIssuance() error {
	endpointURL := fmt.Sprintf(initiateCredentialIssuanceURLFormat, s.issuerProfile.ID)
	token := s.bddContext.Args[getOrgAuthTokenKey(s.issuerProfile.OrganizationID)]

	reqBody, err := json.Marshal(&initiateOIDC4CIRequest{
		CredentialTemplateId: "templateID",
		GrantType:            "authorization_code",
		OpState:              uuid.New().String(),
		ResponseType:         "code",
		Scope:                []string{"openid", "profile"},
		ClaimEndpoint:        claimDataURL,
	})
	if err != nil {
		return fmt.Errorf("marshal initiate oidc4ci req: %w", err)
	}

	resp, err := bddutil.HTTPSDo(http.MethodPost, endpointURL, "application/json", token, bytes.NewReader(reqBody),
		s.bddContext.TLSConfig)
	if err != nil {
		return fmt.Errorf("https do: %w", err)
	}

	defer bddutil.CloseResponseBody(resp.Body)

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return bddutil.ExpectedStatusCodeError(http.StatusOK, resp.StatusCode, b)
	}

	var r initiateOIDC4CIResponse

	if err = json.Unmarshal(b, &r); err != nil {
		return fmt.Errorf("unmarshal initiate oidc4ci resp: %w", err)
	}

	s.initiateIssuanceURL = r.InitiateIssuanceUrl

	return nil
}

func (s *Steps) checkInitiateIssuanceURL() error {
	if s.initiateIssuanceURL == "" {
		return fmt.Errorf("initiate issuance URL is empty")
	}

	if _, err := url.Parse(s.initiateIssuanceURL); err != nil {
		return fmt.Errorf("parse initiate issuance URL: %w", err)
	}

	return nil
}

func (s *Steps) getAuthCode() error {
	httpClient := &http.Client{
		Jar:       s.cookie,
		Transport: &http.Transport{TLSClientConfig: s.bddContext.TLSConfig},
	}

	if s.debug {
		httpClient.Transport = &DumpTransport{httpClient.Transport}
	}

	u, err := url.Parse(s.initiateIssuanceURL)
	if err != nil {
		return fmt.Errorf("parse initiate issuance URL: %w", err)
	}

	opState := u.Query().Get("op_state")
	state := uuid.New().String()

	resp, err := httpClient.Get(
		s.oauthClient.AuthCodeURL(state,
			oauth2.SetAuthURLParam("op_state", opState),
			oauth2.SetAuthURLParam("code_challenge", "MLSjJIlPzeRQoN9YiIsSzziqEuBSmS4kDgI3NDjbfF8"),
			oauth2.SetAuthURLParam("code_challenge_method", "S256"),
			oauth2.SetAuthURLParam("authorization_details", `{"type":"openid_credential","credential_type":"PermanentResidentCard","format":"ldp_vc"}`), //nolint:lll
		),
	)
	if err != nil {
		return fmt.Errorf("get auth code request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	}

	return nil
}

func (s *Steps) authenticateUser() error {
	httpClient := &http.Client{
		Jar:       s.cookie,
		Transport: &http.Transport{TLSClientConfig: s.bddContext.TLSConfig},
		CheckRedirect: func(req *http.Request, via []*http.Request) error { // hijack redirects
			return http.ErrUseLastResponse
		},
	}

	if s.debug {
		httpClient.Transport = &DumpTransport{httpClient.Transport}
	}

	// authenticate user
	resp, err := httpClient.Post(loginPageURL, "", http.NoBody)
	if err != nil {
		return err
	}
	_ = resp.Body.Close()

	// redirect back to third-party oidc provider after login
	resp, err = httpClient.Get(resp.Header.Get("Location"))
	if err != nil {
		return fmt.Errorf("redirect to third-party oidc provider: %w", err)
	}
	_ = resp.Body.Close()

	// redirect to consent page
	resp, err = httpClient.Get(resp.Header.Get("Location"))
	if err != nil {
		return fmt.Errorf("redirect to consent page: %w", err)
	}
	_ = resp.Body.Close()

	// redirect back to third-party oidc provider with consent verifier
	resp, err = httpClient.Get(resp.Header.Get("Location"))
	if err != nil {
		return fmt.Errorf("redirect back to auth after consent: %w", err)
	}
	_ = resp.Body.Close()

	// redirect to public vcs public /oidc/redirect
	resp, err = httpClient.Get(resp.Header.Get("Location"))
	if err != nil {
		return fmt.Errorf("redirect to public oidc redirect: %w", err)
	}
	_ = resp.Body.Close()

	u, err := url.Parse(resp.Header.Get("Location"))
	if err != nil {
		return fmt.Errorf("parse client redirect url: %w", err)
	}

	if !strings.HasPrefix(u.String(), s.oauthClient.RedirectURL) {
		return fmt.Errorf("invalid client redirect url")
	}

	s.authCode = u.Query().Get("code")

	return nil
}

func (s *Steps) checkAuthCode() error {
	if s.authCode == "" {
		return fmt.Errorf("auth code is empty")
	}

	return nil
}

func (s *Steps) exchangeCodeForToken() error {
	httpClient := &http.Client{
		Transport: &http.Transport{TLSClientConfig: s.bddContext.TLSConfig},
	}

	if s.debug {
		httpClient.Transport = &DumpTransport{httpClient.Transport}
	}

	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, httpClient)

	token, err := s.oauthClient.Exchange(ctx, s.authCode,
		oauth2.SetAuthURLParam("code_verifier", "xalsLDydJtHwIQZukUyj6boam5vMUaJRWv-BnGCAzcZi3ZTs"),
	)
	if err != nil {
		return fmt.Errorf("exchange code for token: %w", err)
	}

	s.token = token

	return nil
}

func (s *Steps) checkAccessToken() error {
	if s.token.AccessToken == "" {
		return fmt.Errorf("access token is empty")
	}

	return nil
}

func (s *Steps) getCredential() error {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return fmt.Errorf("generate key pair: %w", err)
	}

	diddoc, err := s.createDID(publicKey)
	if err != nil {
		return fmt.Errorf("create did: %w", err)
	}

	jws, err := s.createProof(privateKey, diddoc.VerificationMethod[0].ID)
	if err != nil {
		return fmt.Errorf("create proof: %w", err)
	}

	b, err := json.Marshal(credentialRequest{
		DID:    diddoc.ID,
		Format: "jwt_vc",
		Type:   "UniversityDegreeCredential",
		Proof: jwtProof{
			ProofType: "jwt",
			JWT:       jws,
		},
	})
	if err != nil {
		return fmt.Errorf("marshal credential request: %w", err)
	}

	var transport http.RoundTripper = &http.Transport{TLSClientConfig: s.bddContext.TLSConfig}

	if s.debug {
		transport = &DumpTransport{transport}
	}

	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{Transport: transport})

	httpClient := s.oauthClient.Client(ctx, s.token)

	resp, err := httpClient.Post(vcsCredentialEndpoint, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return fmt.Errorf("get credential: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("get credential: status %s", resp.Status)
	}

	var credentialResp credentialResponse

	if err = json.NewDecoder(resp.Body).Decode(&credentialResp); err != nil {
		return fmt.Errorf("decode credential response: %w", err)
	}

	s.credential = credentialResp.Credential

	return nil
}

func (s *Steps) createDID(pub ed25519.PublicKey) (*docdid.Doc, error) {
	vdr, err := orb.New(nil, orb.WithDomain(didDomain), orb.WithTLSConfig(s.bddContext.TLSConfig),
		orb.WithAuthToken(didServiceAuthToken))
	if err != nil {
		return nil, fmt.Errorf("create orb vdr: %w", err)
	}

	jwk, err := jwksupport.JWKFromKey(pub)
	if err != nil {
		return nil, fmt.Errorf("create jwk from key: %w", err)
	}

	docID := uuid.NewString()
	keyID := uuid.NewString()

	vm, err := docdid.NewVerificationMethodFromJWK(docID+"#"+keyID, vccrypto.JSONWebKey2020, "", jwk)
	if err != nil {
		return nil, fmt.Errorf("create verification method: %w", err)
	}

	doc := &docdid.Doc{
		ID:              docID,
		Authentication:  []docdid.Verification{*docdid.NewReferencedVerification(vm, docdid.Authentication)},
		AssertionMethod: []docdid.Verification{*docdid.NewReferencedVerification(vm, docdid.AssertionMethod)},
	}

	vdrRegistry := vdrpkg.New(vdrpkg.WithVDR(vdr))

	updateKey, _, err := ed25519.GenerateKey(rand.Reader)
	recoverKey, _, err := ed25519.GenerateKey(rand.Reader)

	docResolution, err := vdrRegistry.Create(
		orb.DIDMethod,
		doc,
		vdrapi.WithOption(orb.UpdatePublicKeyOpt, updateKey),
		vdrapi.WithOption(orb.RecoveryPublicKeyOpt, recoverKey),
	)
	if err != nil {
		return nil, fmt.Errorf("register did in vdr: %w", err)
	}

	return docResolution.DIDDocument, nil
}

func (s *Steps) createProof(privateKey ed25519.PrivateKey, verificationKID string) (string, error) {
	jwtSigner := jwt.NewEd25519Signer(privateKey)

	claims := &jwtProofClaims{
		Issuer:   s.oauthClient.ClientID,
		IssuedAt: time.Now().Unix(),
		Nonce:    s.token.Extra("c_nonce").(string),
	}

	jwtHeaders := map[string]interface{}{
		"alg": "EdDSA",
		"kid": verificationKID,
	}

	signedJWT, err := jwt.NewSigned(claims, jwtHeaders, jwtSigner)
	if err != nil {
		return "", fmt.Errorf("create signed jwt: %w", err)
	}

	jws, err := signedJWT.Serialize(false)
	if err != nil {
		return "", fmt.Errorf("serialize signed jwt: %w", err)
	}

	return jws, nil
}

func (s *Steps) checkCredential() error {
	if s.credential == nil {
		return fmt.Errorf("credential is empty")
	}

	return nil
}

func getOrgAuthTokenKey(org string) string {
	return org + "-accessToken"
}

// DumpTransport is http.RoundTripper that dumps request/response.
type DumpTransport struct {
	r http.RoundTripper
}

// RoundTrip implements the RoundTripper interface.
func (d *DumpTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	reqDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		return nil, fmt.Errorf("failed to dump request: %w", err)
	}

	fmt.Printf("REQUEST:\n%s\n", string(reqDump))

	resp, err := d.r.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	respDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return nil, fmt.Errorf("failed to dump response: %w", err)
	}

	fmt.Printf("RESPONSE:\n%s\n", string(respDump))

	return resp, err
}
