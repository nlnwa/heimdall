package pdp

import (
	"fmt"
	"github.com/gobwas/glob"
	"github.com/nlnwa/whatwg-url/canonicalizer"
	"gopkg.in/yaml.v3"
	"os"
)

var policies []AccessPolicy

/*
CanAccess takes an AccessRequest and returns an AccessResponse.
The AccessResponse contains the permission to access the url.
The url get evaluated against the policies, before a response is returned.
*/
func CanAccess(accRec AccessRequest) AccessResponse {
	for _, p := range policies {
		if p.Name == getRoleFromToken(accRec.Token) {
			for _, r := range p.Rules {
				if evaluateUrl(r.Url, accRec.Url) {
					switch r.Policy {
					case "allow":
						return AccessResponse{Permission: Allow}
					case "deny":
						return AccessResponse{Permission: Deny}
					default:
						fmt.Println("Illegal policy", r.Policy)
						return AccessResponse{Permission: Deny}
					}
				}
			}
			if p.DefaultPolicy == "allow" {
				return AccessResponse{Permission: Allow}
			} else {
				return AccessResponse{Permission: Deny}
			}
		}
	}
	return AccessResponse{Permission: Deny}
}

/*
SetPolicies loads a list of policies from a yaml file.
*/
func SetPolicies(policyFile string) error {
	f, err := os.ReadFile(policyFile)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	if err := yaml.Unmarshal(f, &policies); err != nil {
		return nil
	}
	return err
}

/*
evaluateUrl takes a policy url and a request url and returns true if the request url matches the policy url.
The request url is canonicalized before comparison.
The policy url can contain wildcards.
*/
func evaluateUrl(policyUrl string, requestUrl string) bool {
	c := canonicalizer.New(canonicalizer.WithDefaultScheme("http"))
	reqUrl, err := c.Parse(requestUrl)
	if err != nil {
		fmt.Println("error parsing url", err)
		return false
	}
	g, err := glob.Compile(policyUrl)
	if err != nil {
		fmt.Println("error compiling glob", err)
		return false
	}
	return g.Match(reqUrl.Href(false))
}

/*
getRoleFromToken takes a token and returns the role of the user.
*/
func getRoleFromToken(token string) string {
	switch token {
	case "1":
		return "admin"
	case "2":
		return "curator"
	default:
		return "default"
	}
}

// AccessPolicy  describes the format of our policies/*
type AccessPolicy struct {
	Name          string `yaml:"name"`
	DefaultPolicy string `yaml:"defaultPolicy"`
	Rules         []Rule `yaml:"rules"`
}

// Rule describe rules in the AccessPolicy/*
type Rule struct {
	Url         string `yaml:"url"`
	Policy      string `yaml:"policy"`
	Description string `yaml:"description"`
}

// AccessRequest model info
// @Description Requests access to a url
// @Description for a user with a certain role given by token
type AccessRequest struct {
	Url   string `json:"url"`
	Token string `json:"token"`
}

// AccessResponse model info
// @Description Response to a request for access to a url
// @Description for a user with a certain role given by token
type AccessResponse struct {
	Permission Permission `json:"permission"`
}

// Permission model info
// @Description Permission to access a url
// @Description Used in AccessResponse
type Permission string

const (
	Allow Permission = "allow"
	Deny  Permission = "deny"
)
