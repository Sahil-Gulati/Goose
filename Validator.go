package goose

import "net/http"

type GooseValidator struct {
	route   GooseRoute
	request *http.Request
}

func (v GooseValidator) GetInstance(route GooseRoute, request *http.Request) *GooseValidator {
	validator := new(GooseValidator)
	validator.route = route
	validator.request = request
	return validator
}
func (v *GooseValidator) validate() bool {
	return v.checkMethod() && v.checkUrl()
}
func (v *GooseValidator) checkMethod() bool {
	if contains(v.route.methods, v.request.Method) {
		return true
	}
	return false

}
func (v *GooseValidator) checkUrl() bool {
	return true

}