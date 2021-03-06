// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package domain_name

import (
	"context"

	ackv1alpha1 "github.com/aws/aws-controllers-k8s/apis/core/v1alpha1"
	ackcompare "github.com/aws/aws-controllers-k8s/pkg/compare"
	ackerr "github.com/aws/aws-controllers-k8s/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws/aws-controllers-k8s/services/apigatewayv2/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = &svcsdk.ApiGatewayV2{}
	_ = &svcapitypes.DomainName{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.GetDomainNameWithContext(ctx, input)
	if respErr != nil {
		if awsErr, ok := ackerr.AWSError(respErr); ok && awsErr.Code() == "NotFoundException" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.ApiMappingSelectionExpression != nil {
		ko.Status.APIMappingSelectionExpression = resp.ApiMappingSelectionExpression
	}

	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required by not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Spec.DomainName == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.GetDomainNameInput, error) {
	res := &svcsdk.GetDomainNameInput{}

	if r.ko.Spec.DomainName != nil {
		res.SetDomainName(*r.ko.Spec.DomainName)
	}

	return res, nil
}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.GetDomainNamesInput, error) {
	res := &svcsdk.GetDomainNamesInput{}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a new resource with any fields in the Status field filled in
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newCreateRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.CreateDomainNameWithContext(ctx, input)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.ApiMappingSelectionExpression != nil {
		ko.Status.APIMappingSelectionExpression = resp.ApiMappingSelectionExpression
	}

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	ko.Status.Conditions = []*ackv1alpha1.Condition{}
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	r *resource,
) (*svcsdk.CreateDomainNameInput, error) {
	res := &svcsdk.CreateDomainNameInput{}

	if r.ko.Spec.DomainName != nil {
		res.SetDomainName(*r.ko.Spec.DomainName)
	}
	if r.ko.Spec.DomainNameConfigurations != nil {
		f1 := []*svcsdk.DomainNameConfiguration{}
		for _, f1iter := range r.ko.Spec.DomainNameConfigurations {
			f1elem := &svcsdk.DomainNameConfiguration{}
			if f1iter.APIGatewayDomainName != nil {
				f1elem.SetApiGatewayDomainName(*f1iter.APIGatewayDomainName)
			}
			if f1iter.CertificateARN != nil {
				f1elem.SetCertificateArn(*f1iter.CertificateARN)
			}
			if f1iter.CertificateName != nil {
				f1elem.SetCertificateName(*f1iter.CertificateName)
			}
			if f1iter.CertificateUploadDate != nil {
				f1elem.SetCertificateUploadDate(f1iter.CertificateUploadDate.Time)
			}
			if f1iter.DomainNameStatus != nil {
				f1elem.SetDomainNameStatus(*f1iter.DomainNameStatus)
			}
			if f1iter.DomainNameStatusMessage != nil {
				f1elem.SetDomainNameStatusMessage(*f1iter.DomainNameStatusMessage)
			}
			if f1iter.EndpointType != nil {
				f1elem.SetEndpointType(*f1iter.EndpointType)
			}
			if f1iter.HostedZoneID != nil {
				f1elem.SetHostedZoneId(*f1iter.HostedZoneID)
			}
			if f1iter.SecurityPolicy != nil {
				f1elem.SetSecurityPolicy(*f1iter.SecurityPolicy)
			}
			f1 = append(f1, f1elem)
		}
		res.SetDomainNameConfigurations(f1)
	}
	if r.ko.Spec.MutualTLSAuthentication != nil {
		f2 := &svcsdk.MutualTlsAuthenticationInput{}
		if r.ko.Spec.MutualTLSAuthentication.TruststoreURI != nil {
			f2.SetTruststoreUri(*r.ko.Spec.MutualTLSAuthentication.TruststoreURI)
		}
		if r.ko.Spec.MutualTLSAuthentication.TruststoreVersion != nil {
			f2.SetTruststoreVersion(*r.ko.Spec.MutualTLSAuthentication.TruststoreVersion)
		}
		res.SetMutualTlsAuthentication(f2)
	}
	if r.ko.Spec.Tags != nil {
		f3 := map[string]*string{}
		for f3key, f3valiter := range r.ko.Spec.Tags {
			var f3val string
			f3val = *f3valiter
			f3[f3key] = &f3val
		}
		res.SetTags(f3)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	diffReporter *ackcompare.Reporter,
) (*resource, error) {

	input, err := rm.newUpdateRequestPayload(desired)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.UpdateDomainNameWithContext(ctx, input)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.ApiMappingSelectionExpression != nil {
		ko.Status.APIMappingSelectionExpression = resp.ApiMappingSelectionExpression
	}

	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	r *resource,
) (*svcsdk.UpdateDomainNameInput, error) {
	res := &svcsdk.UpdateDomainNameInput{}

	if r.ko.Spec.DomainName != nil {
		res.SetDomainName(*r.ko.Spec.DomainName)
	}
	if r.ko.Spec.DomainNameConfigurations != nil {
		f1 := []*svcsdk.DomainNameConfiguration{}
		for _, f1iter := range r.ko.Spec.DomainNameConfigurations {
			f1elem := &svcsdk.DomainNameConfiguration{}
			if f1iter.APIGatewayDomainName != nil {
				f1elem.SetApiGatewayDomainName(*f1iter.APIGatewayDomainName)
			}
			if f1iter.CertificateARN != nil {
				f1elem.SetCertificateArn(*f1iter.CertificateARN)
			}
			if f1iter.CertificateName != nil {
				f1elem.SetCertificateName(*f1iter.CertificateName)
			}
			if f1iter.CertificateUploadDate != nil {
				f1elem.SetCertificateUploadDate(f1iter.CertificateUploadDate.Time)
			}
			if f1iter.DomainNameStatus != nil {
				f1elem.SetDomainNameStatus(*f1iter.DomainNameStatus)
			}
			if f1iter.DomainNameStatusMessage != nil {
				f1elem.SetDomainNameStatusMessage(*f1iter.DomainNameStatusMessage)
			}
			if f1iter.EndpointType != nil {
				f1elem.SetEndpointType(*f1iter.EndpointType)
			}
			if f1iter.HostedZoneID != nil {
				f1elem.SetHostedZoneId(*f1iter.HostedZoneID)
			}
			if f1iter.SecurityPolicy != nil {
				f1elem.SetSecurityPolicy(*f1iter.SecurityPolicy)
			}
			f1 = append(f1, f1elem)
		}
		res.SetDomainNameConfigurations(f1)
	}
	if r.ko.Spec.MutualTLSAuthentication != nil {
		f2 := &svcsdk.MutualTlsAuthenticationInput{}
		if r.ko.Spec.MutualTLSAuthentication.TruststoreURI != nil {
			f2.SetTruststoreUri(*r.ko.Spec.MutualTLSAuthentication.TruststoreURI)
		}
		if r.ko.Spec.MutualTLSAuthentication.TruststoreVersion != nil {
			f2.SetTruststoreVersion(*r.ko.Spec.MutualTLSAuthentication.TruststoreVersion)
		}
		res.SetMutualTlsAuthentication(f2)
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) error {
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return err
	}
	_, respErr := rm.sdkapi.DeleteDomainNameWithContext(ctx, input)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteDomainNameInput, error) {
	res := &svcsdk.DeleteDomainNameInput{}

	if r.ko.Spec.DomainName != nil {
		res.SetDomainName(*r.ko.Spec.DomainName)
	}

	return res, nil
}
