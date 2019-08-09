package goose

import (
	"net/http"
	"regexp"
)

type GooseValidator struct {
	route       *GooseRoute
	regexRoutes map[string]*GooseRoute
	request     *http.Request
	message		*GooseMessage
	response    *GooseResponse
}

func (v GooseValidator) GetInstance() *GooseValidator {
	validator := new(GooseValidator)
	validator.route = v.route
	validator.request = v.request
	validator.message = v.message
	validator.response = v.response
	validator.regexRoutes = v.regexRoutes
	return validator
}
func (v *GooseValidator) validate() bool {
	if contains(v.route.methods, v.request.Method) {
		return true
	}
	return false
}
func (v *GooseValidator) regexValidate() (bool, *GooseRoute) {
	if route := v.getRouteForRegexURI(); route != nil {
		v.route = route
		return v.validate(), route
	}
	return false, &GooseRoute{}
}

func (v *GooseValidator) getRouteForRegexURI() *GooseRoute {
	for _, route := range v.regexRoutes {
		regex := regexp.MustCompile(route.uRLRegex)
		substrings := regex.FindAllStringSubmatch(v.request.URL.Path, 1)
		urlParams := map[string]string{}
		if len(substrings) > 0 && len(substrings[0]) > 1 {
			for i := 1; i < len(substrings[0]); i++ {
				key := route.dynamics[i-1]
				value := substrings[0][i]
				urlParams[key] = value
			}
			v.message.UrlParams = urlParams
			return route
		}
	}
	return nil
}
func (v *GooseValidator) getDefaultResponse(statusCode int) {
	switch statusCode {
	case 404:
		v.response.StatusCode = statusCode
		v.response.Response = PAGE_NOT_FOUND
		v.response.Headers = map[string]string{CONTENT_TYPE: PLAIN_TEXT}
	case 405:
		v.response.StatusCode = statusCode
		v.response.Response = METHOD_NOT_ALLOWED
		v.response.Headers = map[string]string{CONTENT_TYPE: PLAIN_TEXT}
	}
}
