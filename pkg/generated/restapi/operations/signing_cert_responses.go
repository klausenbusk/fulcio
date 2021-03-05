// Code generated by go-swagger; DO NOT EDIT.

// /*
// Copyright The Fulcio Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/sigstore/fulcio/pkg/generated/models"
)

// SigningCertCreatedCode is the HTTP code returned for type SigningCertCreated
const SigningCertCreatedCode int = 201

/*SigningCertCreated Successful CSR Submit

swagger:response signingCertCreated
*/
type SigningCertCreated struct {

	/*
	  In: Body
	*/
	Payload *models.SubmitSuccess `json:"body,omitempty"`
}

// NewSigningCertCreated creates SigningCertCreated with default headers values
func NewSigningCertCreated() *SigningCertCreated {

	return &SigningCertCreated{}
}

// WithPayload adds the payload to the signing cert created response
func (o *SigningCertCreated) WithPayload(payload *models.SubmitSuccess) *SigningCertCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the signing cert created response
func (o *SigningCertCreated) SetPayload(payload *models.SubmitSuccess) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SigningCertCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SigningCertBadRequestCode is the HTTP code returned for type SigningCertBadRequest
const SigningCertBadRequestCode int = 400

/*SigningCertBadRequest Bad Request

swagger:response signingCertBadRequest
*/
type SigningCertBadRequest struct {
}

// NewSigningCertBadRequest creates SigningCertBadRequest with default headers values
func NewSigningCertBadRequest() *SigningCertBadRequest {

	return &SigningCertBadRequest{}
}

// WriteResponse to the client
func (o *SigningCertBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// SigningCertUnauthorizedCode is the HTTP code returned for type SigningCertUnauthorized
const SigningCertUnauthorizedCode int = 401

/*SigningCertUnauthorized Unauthorized

swagger:response signingCertUnauthorized
*/
type SigningCertUnauthorized struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewSigningCertUnauthorized creates SigningCertUnauthorized with default headers values
func NewSigningCertUnauthorized() *SigningCertUnauthorized {

	return &SigningCertUnauthorized{}
}

// WithPayload adds the payload to the signing cert unauthorized response
func (o *SigningCertUnauthorized) WithPayload(payload string) *SigningCertUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the signing cert unauthorized response
func (o *SigningCertUnauthorized) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SigningCertUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// SigningCertInternalServerErrorCode is the HTTP code returned for type SigningCertInternalServerError
const SigningCertInternalServerErrorCode int = 500

/*SigningCertInternalServerError Server error

swagger:response signingCertInternalServerError
*/
type SigningCertInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewSigningCertInternalServerError creates SigningCertInternalServerError with default headers values
func NewSigningCertInternalServerError() *SigningCertInternalServerError {

	return &SigningCertInternalServerError{}
}

// WithPayload adds the payload to the signing cert internal server error response
func (o *SigningCertInternalServerError) WithPayload(payload string) *SigningCertInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the signing cert internal server error response
func (o *SigningCertInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SigningCertInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}