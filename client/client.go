// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddUserToOrganizationalUnitsRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The organization IDs. You can add an account to a maximum of 100 organizations.
	//
	// This parameter is required.
	OrganizationalUnitIds []*string `json:"OrganizationalUnitIds,omitempty" xml:"OrganizationalUnitIds,omitempty" type:"Repeated"`
	// The account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddUserToOrganizationalUnitsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUserToOrganizationalUnitsRequest) GoString() string {
	return s.String()
}

func (s *AddUserToOrganizationalUnitsRequest) SetInstanceId(v string) *AddUserToOrganizationalUnitsRequest {
	s.InstanceId = &v
	return s
}

func (s *AddUserToOrganizationalUnitsRequest) SetOrganizationalUnitIds(v []*string) *AddUserToOrganizationalUnitsRequest {
	s.OrganizationalUnitIds = v
	return s
}

func (s *AddUserToOrganizationalUnitsRequest) SetUserId(v string) *AddUserToOrganizationalUnitsRequest {
	s.UserId = &v
	return s
}

type AddUserToOrganizationalUnitsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUserToOrganizationalUnitsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUserToOrganizationalUnitsResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserToOrganizationalUnitsResponseBody) SetRequestId(v string) *AddUserToOrganizationalUnitsResponseBody {
	s.RequestId = &v
	return s
}

type AddUserToOrganizationalUnitsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserToOrganizationalUnitsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUserToOrganizationalUnitsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUserToOrganizationalUnitsResponse) GoString() string {
	return s.String()
}

func (s *AddUserToOrganizationalUnitsResponse) SetHeaders(v map[string]*string) *AddUserToOrganizationalUnitsResponse {
	s.Headers = v
	return s
}

func (s *AddUserToOrganizationalUnitsResponse) SetStatusCode(v int32) *AddUserToOrganizationalUnitsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserToOrganizationalUnitsResponse) SetBody(v *AddUserToOrganizationalUnitsResponseBody) *AddUserToOrganizationalUnitsResponse {
	s.Body = v
	return s
}

type AddUsersToGroupRequest struct {
	// The group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The account IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// [ou_001]
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s AddUsersToGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUsersToGroupRequest) GoString() string {
	return s.String()
}

func (s *AddUsersToGroupRequest) SetGroupId(v string) *AddUsersToGroupRequest {
	s.GroupId = &v
	return s
}

func (s *AddUsersToGroupRequest) SetInstanceId(v string) *AddUsersToGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *AddUsersToGroupRequest) SetUserIds(v []*string) *AddUsersToGroupRequest {
	s.UserIds = v
	return s
}

type AddUsersToGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUsersToGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUsersToGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddUsersToGroupResponseBody) SetRequestId(v string) *AddUsersToGroupResponseBody {
	s.RequestId = &v
	return s
}

type AddUsersToGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUsersToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUsersToGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUsersToGroupResponse) GoString() string {
	return s.String()
}

func (s *AddUsersToGroupResponse) SetHeaders(v map[string]*string) *AddUsersToGroupResponse {
	s.Headers = v
	return s
}

func (s *AddUsersToGroupResponse) SetStatusCode(v int32) *AddUsersToGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUsersToGroupResponse) SetBody(v *AddUsersToGroupResponseBody) *AddUsersToGroupResponse {
	s.Body = v
	return s
}

type AuthorizeApplicationToGroupsRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The group IDs. You can specify up to 100 group IDs at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_miu8e4t4d7i4u7uwezgr54xxxx
	GroupIds []*string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk2676xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s AuthorizeApplicationToGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeApplicationToGroupsRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeApplicationToGroupsRequest) SetApplicationId(v string) *AuthorizeApplicationToGroupsRequest {
	s.ApplicationId = &v
	return s
}

func (s *AuthorizeApplicationToGroupsRequest) SetGroupIds(v []*string) *AuthorizeApplicationToGroupsRequest {
	s.GroupIds = v
	return s
}

func (s *AuthorizeApplicationToGroupsRequest) SetInstanceId(v string) *AuthorizeApplicationToGroupsRequest {
	s.InstanceId = &v
	return s
}

type AuthorizeApplicationToGroupsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuthorizeApplicationToGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeApplicationToGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeApplicationToGroupsResponseBody) SetRequestId(v string) *AuthorizeApplicationToGroupsResponseBody {
	s.RequestId = &v
	return s
}

type AuthorizeApplicationToGroupsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeApplicationToGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeApplicationToGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeApplicationToGroupsResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeApplicationToGroupsResponse) SetHeaders(v map[string]*string) *AuthorizeApplicationToGroupsResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeApplicationToGroupsResponse) SetStatusCode(v int32) *AuthorizeApplicationToGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeApplicationToGroupsResponse) SetBody(v *AuthorizeApplicationToGroupsResponseBody) *AuthorizeApplicationToGroupsResponse {
	s.Body = v
	return s
}

type AuthorizeApplicationToOrganizationalUnitsRequest struct {
	// The ID of the application on which you want to grant permissions.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk2676xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IDs of the organizations to which you want to grant permissions. You can grant permissions to a maximum of 100 organizations at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitIds []*string `json:"OrganizationalUnitIds,omitempty" xml:"OrganizationalUnitIds,omitempty" type:"Repeated"`
}

func (s AuthorizeApplicationToOrganizationalUnitsRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeApplicationToOrganizationalUnitsRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeApplicationToOrganizationalUnitsRequest) SetApplicationId(v string) *AuthorizeApplicationToOrganizationalUnitsRequest {
	s.ApplicationId = &v
	return s
}

func (s *AuthorizeApplicationToOrganizationalUnitsRequest) SetInstanceId(v string) *AuthorizeApplicationToOrganizationalUnitsRequest {
	s.InstanceId = &v
	return s
}

func (s *AuthorizeApplicationToOrganizationalUnitsRequest) SetOrganizationalUnitIds(v []*string) *AuthorizeApplicationToOrganizationalUnitsRequest {
	s.OrganizationalUnitIds = v
	return s
}

type AuthorizeApplicationToOrganizationalUnitsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuthorizeApplicationToOrganizationalUnitsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeApplicationToOrganizationalUnitsResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeApplicationToOrganizationalUnitsResponseBody) SetRequestId(v string) *AuthorizeApplicationToOrganizationalUnitsResponseBody {
	s.RequestId = &v
	return s
}

type AuthorizeApplicationToOrganizationalUnitsResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeApplicationToOrganizationalUnitsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeApplicationToOrganizationalUnitsResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeApplicationToOrganizationalUnitsResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeApplicationToOrganizationalUnitsResponse) SetHeaders(v map[string]*string) *AuthorizeApplicationToOrganizationalUnitsResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeApplicationToOrganizationalUnitsResponse) SetStatusCode(v int32) *AuthorizeApplicationToOrganizationalUnitsResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeApplicationToOrganizationalUnitsResponse) SetBody(v *AuthorizeApplicationToOrganizationalUnitsResponseBody) *AuthorizeApplicationToOrganizationalUnitsResponse {
	s.Body = v
	return s
}

type AuthorizeApplicationToUsersRequest struct {
	// The ID of the application on which you want to grant permissions.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk2676xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IDs of the accounts to which you want to grant permissions. You can grant permissions to a maximum of 100 accounts at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s AuthorizeApplicationToUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeApplicationToUsersRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeApplicationToUsersRequest) SetApplicationId(v string) *AuthorizeApplicationToUsersRequest {
	s.ApplicationId = &v
	return s
}

func (s *AuthorizeApplicationToUsersRequest) SetInstanceId(v string) *AuthorizeApplicationToUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *AuthorizeApplicationToUsersRequest) SetUserIds(v []*string) *AuthorizeApplicationToUsersRequest {
	s.UserIds = v
	return s
}

type AuthorizeApplicationToUsersResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuthorizeApplicationToUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeApplicationToUsersResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeApplicationToUsersResponseBody) SetRequestId(v string) *AuthorizeApplicationToUsersResponseBody {
	s.RequestId = &v
	return s
}

type AuthorizeApplicationToUsersResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeApplicationToUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeApplicationToUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeApplicationToUsersResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeApplicationToUsersResponse) SetHeaders(v map[string]*string) *AuthorizeApplicationToUsersResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeApplicationToUsersResponse) SetStatusCode(v int32) *AuthorizeApplicationToUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeApplicationToUsersResponse) SetBody(v *AuthorizeApplicationToUsersResponseBody) *AuthorizeApplicationToUsersResponse {
	s.Body = v
	return s
}

type CreateApplicationRequest struct {
	// The name of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// Ram Account SSO
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The type of the application source. Valid values:
	//
	// 	- urn:alibaba:idaas:app:source:template: application template
	//
	// 	- urn:alibaba:idaas:app:source:standard: standard protocol
	//
	// This parameter is required.
	//
	// example:
	//
	// urn:alibaba:idaas:app:source:standard
	ApplicationSourceType *string `json:"ApplicationSourceType,omitempty" xml:"ApplicationSourceType,omitempty"`
	// The ID of the application template. This parameter is required if you set the ApplicationSourceType parameter to urn:alibaba:idaas:app:source:template.
	//
	// example:
	//
	// template_cloud_ram
	ApplicationTemplateId *string `json:"ApplicationTemplateId,omitempty" xml:"ApplicationTemplateId,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// RAM user SSO application
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk2676xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The URL of the application logo.
	//
	// example:
	//
	// https://oss.cn-hangzhou.aliyuncs.com/logo.png
	LogoUrl *string `json:"LogoUrl,omitempty" xml:"LogoUrl,omitempty"`
	// The SSO protocol. Valid values:
	//
	// 	- saml2: the SAML 2.0 protocol.
	//
	// 	- oidc: the OpenID Connect protocol.
	//
	// This parameter is required.
	//
	// example:
	//
	// saml2
	SsoType *string `json:"SsoType,omitempty" xml:"SsoType,omitempty"`
}

func (s CreateApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequest) SetApplicationName(v string) *CreateApplicationRequest {
	s.ApplicationName = &v
	return s
}

func (s *CreateApplicationRequest) SetApplicationSourceType(v string) *CreateApplicationRequest {
	s.ApplicationSourceType = &v
	return s
}

func (s *CreateApplicationRequest) SetApplicationTemplateId(v string) *CreateApplicationRequest {
	s.ApplicationTemplateId = &v
	return s
}

func (s *CreateApplicationRequest) SetDescription(v string) *CreateApplicationRequest {
	s.Description = &v
	return s
}

func (s *CreateApplicationRequest) SetInstanceId(v string) *CreateApplicationRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateApplicationRequest) SetLogoUrl(v string) *CreateApplicationRequest {
	s.LogoUrl = &v
	return s
}

func (s *CreateApplicationRequest) SetSsoType(v string) *CreateApplicationRequest {
	s.SsoType = &v
	return s
}

type CreateApplicationResponseBody struct {
	// The ID of the application.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mnkom
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBody) SetApplicationId(v string) *CreateApplicationResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *CreateApplicationResponseBody) SetRequestId(v string) *CreateApplicationResponseBody {
	s.RequestId = &v
	return s
}

type CreateApplicationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponse) SetHeaders(v map[string]*string) *CreateApplicationResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationResponse) SetStatusCode(v int32) *CreateApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationResponse) SetBody(v *CreateApplicationResponseBody) *CreateApplicationResponse {
	s.Body = v
	return s
}

type CreateApplicationClientSecretRequest struct {
	// The ID of the application for which you want to create a client key.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateApplicationClientSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationClientSecretRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationClientSecretRequest) SetApplicationId(v string) *CreateApplicationClientSecretRequest {
	s.ApplicationId = &v
	return s
}

func (s *CreateApplicationClientSecretRequest) SetInstanceId(v string) *CreateApplicationClientSecretRequest {
	s.InstanceId = &v
	return s
}

type CreateApplicationClientSecretResponseBody struct {
	// The information about the client key.
	ApplicationClientSecret *CreateApplicationClientSecretResponseBodyApplicationClientSecret `json:"ApplicationClientSecret,omitempty" xml:"ApplicationClientSecret,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApplicationClientSecretResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationClientSecretResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationClientSecretResponseBody) SetApplicationClientSecret(v *CreateApplicationClientSecretResponseBodyApplicationClientSecret) *CreateApplicationClientSecretResponseBody {
	s.ApplicationClientSecret = v
	return s
}

func (s *CreateApplicationClientSecretResponseBody) SetRequestId(v string) *CreateApplicationClientSecretResponseBody {
	s.RequestId = &v
	return s
}

type CreateApplicationClientSecretResponseBodyApplicationClientSecret struct {
	// The client ID of the application.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The client key secret of the application.
	//
	// example:
	//
	// CSEHDcHcrUKHw1CuxkJEHPveWRXBGqVqRsxxxx
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	// The client key ID of the application.
	//
	// example:
	//
	// sci_k52x2ru63rlkflina5utgkxxxx
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
}

func (s CreateApplicationClientSecretResponseBodyApplicationClientSecret) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationClientSecretResponseBodyApplicationClientSecret) GoString() string {
	return s.String()
}

func (s *CreateApplicationClientSecretResponseBodyApplicationClientSecret) SetClientId(v string) *CreateApplicationClientSecretResponseBodyApplicationClientSecret {
	s.ClientId = &v
	return s
}

func (s *CreateApplicationClientSecretResponseBodyApplicationClientSecret) SetClientSecret(v string) *CreateApplicationClientSecretResponseBodyApplicationClientSecret {
	s.ClientSecret = &v
	return s
}

func (s *CreateApplicationClientSecretResponseBodyApplicationClientSecret) SetSecretId(v string) *CreateApplicationClientSecretResponseBodyApplicationClientSecret {
	s.SecretId = &v
	return s
}

type CreateApplicationClientSecretResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApplicationClientSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApplicationClientSecretResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationClientSecretResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationClientSecretResponse) SetHeaders(v map[string]*string) *CreateApplicationClientSecretResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationClientSecretResponse) SetStatusCode(v int32) *CreateApplicationClientSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationClientSecretResponse) SetBody(v *CreateApplicationClientSecretResponseBody) *CreateApplicationClientSecretResponse {
	s.Body = v
	return s
}

type CreateDomainRequest struct {
	// 域名。最大长度限制255，格式由数字、字母、横线（-）点（.）组成;
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// 备案信息参数。
	Filing *CreateDomainRequestFiling `json:"Filing,omitempty" xml:"Filing,omitempty" type:"Struct"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainRequest) SetDomain(v string) *CreateDomainRequest {
	s.Domain = &v
	return s
}

func (s *CreateDomainRequest) SetFiling(v *CreateDomainRequestFiling) *CreateDomainRequest {
	s.Filing = v
	return s
}

func (s *CreateDomainRequest) SetInstanceId(v string) *CreateDomainRequest {
	s.InstanceId = &v
	return s
}

type CreateDomainRequestFiling struct {
	// 域名关联的备案号，长度最大限制64。
	//
	// example:
	//
	// 浙xx-xxxxxx
	IcpNumber *string `json:"IcpNumber,omitempty" xml:"IcpNumber,omitempty"`
}

func (s CreateDomainRequestFiling) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainRequestFiling) GoString() string {
	return s.String()
}

func (s *CreateDomainRequestFiling) SetIcpNumber(v string) *CreateDomainRequestFiling {
	s.IcpNumber = &v
	return s
}

type CreateDomainResponseBody struct {
	// example:
	//
	// dm_mtohn6mltdz3ibtly2rxvnvxxx
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDomainResponseBody) SetDomainId(v string) *CreateDomainResponseBody {
	s.DomainId = &v
	return s
}

func (s *CreateDomainResponseBody) SetRequestId(v string) *CreateDomainResponseBody {
	s.RequestId = &v
	return s
}

type CreateDomainResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainResponse) GoString() string {
	return s.String()
}

func (s *CreateDomainResponse) SetHeaders(v map[string]*string) *CreateDomainResponse {
	s.Headers = v
	return s
}

func (s *CreateDomainResponse) SetStatusCode(v int32) *CreateDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDomainResponse) SetBody(v *CreateDomainResponseBody) *CreateDomainResponse {
	s.Body = v
	return s
}

type CreateDomainProxyTokenRequest struct {
	// 域名ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// dm_examplexxxxx
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateDomainProxyTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainProxyTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainProxyTokenRequest) SetDomainId(v string) *CreateDomainProxyTokenRequest {
	s.DomainId = &v
	return s
}

func (s *CreateDomainProxyTokenRequest) SetInstanceId(v string) *CreateDomainProxyTokenRequest {
	s.InstanceId = &v
	return s
}

type CreateDomainProxyTokenResponseBody struct {
	// example:
	//
	// pt_mtohn73423stghoivjmi4jwxxx
	DomainProxyTokenId *string `json:"DomainProxyTokenId,omitempty" xml:"DomainProxyTokenId,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDomainProxyTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainProxyTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDomainProxyTokenResponseBody) SetDomainProxyTokenId(v string) *CreateDomainProxyTokenResponseBody {
	s.DomainProxyTokenId = &v
	return s
}

func (s *CreateDomainProxyTokenResponseBody) SetRequestId(v string) *CreateDomainProxyTokenResponseBody {
	s.RequestId = &v
	return s
}

type CreateDomainProxyTokenResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDomainProxyTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDomainProxyTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainProxyTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateDomainProxyTokenResponse) SetHeaders(v map[string]*string) *CreateDomainProxyTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateDomainProxyTokenResponse) SetStatusCode(v int32) *CreateDomainProxyTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDomainProxyTokenResponse) SetBody(v *CreateDomainProxyTokenResponseBody) *CreateDomainProxyTokenResponse {
	s.Body = v
	return s
}

type CreateGroupRequest struct {
	// The description of the group. The value can be up to 256 characters in length.
	//
	// example:
	//
	// this is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The external ID of the group, which can be used to associate the group with an external system. By default, the external ID is the group ID. The value can be up to 64 characters in length.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupExternalId *string `json:"GroupExternalId,omitempty" xml:"GroupExternalId,omitempty"`
	// The name of the group. The name can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// name_test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupRequest) SetDescription(v string) *CreateGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupRequest) SetGroupExternalId(v string) *CreateGroupRequest {
	s.GroupExternalId = &v
	return s
}

func (s *CreateGroupRequest) SetGroupName(v string) *CreateGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateGroupRequest) SetInstanceId(v string) *CreateGroupRequest {
	s.InstanceId = &v
	return s
}

type CreateGroupResponseBody struct {
	// The group ID.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBody) SetGroupId(v string) *CreateGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateGroupResponseBody) SetRequestId(v string) *CreateGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupResponse) SetHeaders(v map[string]*string) *CreateGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupResponse) SetStatusCode(v int32) *CreateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupResponse) SetBody(v *CreateGroupResponseBody) *CreateGroupResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	// The description of the instance. The description can be up to 128 characters in length.
	//
	// example:
	//
	// instance_for_test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetDescription(v string) *CreateInstanceRequest {
	s.Description = &v
	return s
}

type CreateInstanceResponseBody struct {
	// The ID of the instance that is created.
	//
	// example:
	//
	// idaas_wj5htncdvoc4q5xxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetInstanceId(v string) *CreateInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

type CreateInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponse) SetHeaders(v map[string]*string) *CreateInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceResponse) SetStatusCode(v int32) *CreateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
	s.Body = v
	return s
}

type CreateNetworkAccessEndpointRequest struct {
	// 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。ClientToken只支持ASCII字符，且不能超过64个字符。
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 专属网络端点名称。
	//
	// This parameter is required.
	//
	// example:
	//
	// xx业务VPC访问端点
	NetworkAccessEndpointName *string `json:"NetworkAccessEndpointName,omitempty" xml:"NetworkAccessEndpointName,omitempty"`
	// 专属网络端点连接的指定vSwitch。
	//
	// example:
	//
	// vsw-examplexxx
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// 专属网络端点连接的VpcID。
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-examplexxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// 专属网络端点连接的VpcID所属地域，该地域取值必须在ListNetworkAccessEndpointAvailableRegions接口中返回。
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	VpcRegionId *string `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
}

func (s CreateNetworkAccessEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkAccessEndpointRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkAccessEndpointRequest) SetClientToken(v string) *CreateNetworkAccessEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateNetworkAccessEndpointRequest) SetInstanceId(v string) *CreateNetworkAccessEndpointRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateNetworkAccessEndpointRequest) SetNetworkAccessEndpointName(v string) *CreateNetworkAccessEndpointRequest {
	s.NetworkAccessEndpointName = &v
	return s
}

func (s *CreateNetworkAccessEndpointRequest) SetVSwitchIds(v []*string) *CreateNetworkAccessEndpointRequest {
	s.VSwitchIds = v
	return s
}

func (s *CreateNetworkAccessEndpointRequest) SetVpcId(v string) *CreateNetworkAccessEndpointRequest {
	s.VpcId = &v
	return s
}

func (s *CreateNetworkAccessEndpointRequest) SetVpcRegionId(v string) *CreateNetworkAccessEndpointRequest {
	s.VpcRegionId = &v
	return s
}

type CreateNetworkAccessEndpointResponseBody struct {
	// example:
	//
	// nae_examplexxxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNetworkAccessEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkAccessEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkAccessEndpointResponseBody) SetNetworkAccessEndpointId(v string) *CreateNetworkAccessEndpointResponseBody {
	s.NetworkAccessEndpointId = &v
	return s
}

func (s *CreateNetworkAccessEndpointResponseBody) SetRequestId(v string) *CreateNetworkAccessEndpointResponseBody {
	s.RequestId = &v
	return s
}

type CreateNetworkAccessEndpointResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNetworkAccessEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNetworkAccessEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNetworkAccessEndpointResponse) GoString() string {
	return s.String()
}

func (s *CreateNetworkAccessEndpointResponse) SetHeaders(v map[string]*string) *CreateNetworkAccessEndpointResponse {
	s.Headers = v
	return s
}

func (s *CreateNetworkAccessEndpointResponse) SetStatusCode(v int32) *CreateNetworkAccessEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNetworkAccessEndpointResponse) SetBody(v *CreateNetworkAccessEndpointResponseBody) *CreateNetworkAccessEndpointResponse {
	s.Body = v
	return s
}

type CreateOrganizationalUnitRequest struct {
	// The description of the organization. The value can be up to 256 characters in length.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The external ID of the organization, which can be used to associate the organization with an external system. By default, the external ID is the organization ID. The value can be up to 64 characters in length.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitExternalId *string `json:"OrganizationalUnitExternalId,omitempty" xml:"OrganizationalUnitExternalId,omitempty"`
	// The name of the organization. The name can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_ou_name
	OrganizationalUnitName *string `json:"OrganizationalUnitName,omitempty" xml:"OrganizationalUnitName,omitempty"`
	// The parent organization ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s CreateOrganizationalUnitRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *CreateOrganizationalUnitRequest) SetDescription(v string) *CreateOrganizationalUnitRequest {
	s.Description = &v
	return s
}

func (s *CreateOrganizationalUnitRequest) SetInstanceId(v string) *CreateOrganizationalUnitRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateOrganizationalUnitRequest) SetOrganizationalUnitExternalId(v string) *CreateOrganizationalUnitRequest {
	s.OrganizationalUnitExternalId = &v
	return s
}

func (s *CreateOrganizationalUnitRequest) SetOrganizationalUnitName(v string) *CreateOrganizationalUnitRequest {
	s.OrganizationalUnitName = &v
	return s
}

func (s *CreateOrganizationalUnitRequest) SetParentId(v string) *CreateOrganizationalUnitRequest {
	s.ParentId = &v
	return s
}

type CreateOrganizationalUnitResponseBody struct {
	// The organization ID.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrganizationalUnitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrganizationalUnitResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrganizationalUnitResponseBody) SetOrganizationalUnitId(v string) *CreateOrganizationalUnitResponseBody {
	s.OrganizationalUnitId = &v
	return s
}

func (s *CreateOrganizationalUnitResponseBody) SetRequestId(v string) *CreateOrganizationalUnitResponseBody {
	s.RequestId = &v
	return s
}

type CreateOrganizationalUnitResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrganizationalUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrganizationalUnitResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *CreateOrganizationalUnitResponse) SetHeaders(v map[string]*string) *CreateOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *CreateOrganizationalUnitResponse) SetStatusCode(v int32) *CreateOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrganizationalUnitResponse) SetBody(v *CreateOrganizationalUnitResponseBody) *CreateOrganizationalUnitResponse {
	s.Body = v
	return s
}

type CreateUserRequest struct {
	// The extended fields.
	//
	// example:
	//
	// description
	CustomFields []*CreateUserRequestCustomFields `json:"CustomFields,omitempty" xml:"CustomFields,omitempty" type:"Repeated"`
	// The description of the organizational unit. The description can be up to 256 characters in length.
	//
	// example:
	//
	// description text
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the account. The display name can be up to 64 characters in length.
	//
	// example:
	//
	// name_001
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the user who owns the account. The email address prefix can contain letters, digits, underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// example@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Specifies whether the email address is a trusted email address. This parameter is required if the Email parameter is specified. If you have no special business requirements, set this parameter to true.
	//
	// example:
	//
	// true
	EmailVerified *bool `json:"EmailVerified,omitempty" xml:"EmailVerified,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IDs of organizational units to which the account belongs. An account can belong to multiple organizational units.
	OrganizationalUnitIds []*string `json:"OrganizationalUnitIds,omitempty" xml:"OrganizationalUnitIds,omitempty" type:"Repeated"`
	// The password of the account. For more information, view the password policy of the instance in the IDaaS console.
	//
	// example:
	//
	// 123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The configurations for password initialization.
	PasswordInitializationConfig *CreateUserRequestPasswordInitializationConfig `json:"PasswordInitializationConfig,omitempty" xml:"PasswordInitializationConfig,omitempty" type:"Struct"`
	// The mobile phone number, which contains 6 to 15 digits.
	//
	// example:
	//
	// 12345678901
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Specifies whether the mobile phone number is a trusted mobile phone number. This parameter is required if the PhoneNumber parameter is specified. If you have no special business requirements, set this parameter to true.
	//
	// example:
	//
	// true
	PhoneNumberVerified *bool `json:"PhoneNumberVerified,omitempty" xml:"PhoneNumberVerified,omitempty"`
	// The country code of the mobile phone number. The country code contains only digits and does not contain a plus sign (+).
	//
	// example:
	//
	// 86
	PhoneRegion *string `json:"PhoneRegion,omitempty" xml:"PhoneRegion,omitempty"`
	// The ID of the primary organizational unit to which the account belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	PrimaryOrganizationalUnitId *string `json:"PrimaryOrganizationalUnitId,omitempty" xml:"PrimaryOrganizationalUnitId,omitempty"`
	// The external ID of the account. The external ID can be used to associate the account with an external system. The external ID can be up to 64 characters in length. If you do not specify an external ID for the account, the ID of the account is used as the external ID by default.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserExternalId *string `json:"UserExternalId,omitempty" xml:"UserExternalId,omitempty"`
	// The name of the account. The name can be up to 64 characters in length and can contain letters, digits, underscores (_), periods (.), at signs (@), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// user_001
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) SetCustomFields(v []*CreateUserRequestCustomFields) *CreateUserRequest {
	s.CustomFields = v
	return s
}

func (s *CreateUserRequest) SetDescription(v string) *CreateUserRequest {
	s.Description = &v
	return s
}

func (s *CreateUserRequest) SetDisplayName(v string) *CreateUserRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateUserRequest) SetEmail(v string) *CreateUserRequest {
	s.Email = &v
	return s
}

func (s *CreateUserRequest) SetEmailVerified(v bool) *CreateUserRequest {
	s.EmailVerified = &v
	return s
}

func (s *CreateUserRequest) SetInstanceId(v string) *CreateUserRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateUserRequest) SetOrganizationalUnitIds(v []*string) *CreateUserRequest {
	s.OrganizationalUnitIds = v
	return s
}

func (s *CreateUserRequest) SetPassword(v string) *CreateUserRequest {
	s.Password = &v
	return s
}

func (s *CreateUserRequest) SetPasswordInitializationConfig(v *CreateUserRequestPasswordInitializationConfig) *CreateUserRequest {
	s.PasswordInitializationConfig = v
	return s
}

func (s *CreateUserRequest) SetPhoneNumber(v string) *CreateUserRequest {
	s.PhoneNumber = &v
	return s
}

func (s *CreateUserRequest) SetPhoneNumberVerified(v bool) *CreateUserRequest {
	s.PhoneNumberVerified = &v
	return s
}

func (s *CreateUserRequest) SetPhoneRegion(v string) *CreateUserRequest {
	s.PhoneRegion = &v
	return s
}

func (s *CreateUserRequest) SetPrimaryOrganizationalUnitId(v string) *CreateUserRequest {
	s.PrimaryOrganizationalUnitId = &v
	return s
}

func (s *CreateUserRequest) SetUserExternalId(v string) *CreateUserRequest {
	s.UserExternalId = &v
	return s
}

func (s *CreateUserRequest) SetUsername(v string) *CreateUserRequest {
	s.Username = &v
	return s
}

type CreateUserRequestCustomFields struct {
	// The name of the extended field. You must create the extended field in advance. To create an extended field, log on to the IDaaS console. In the left-side navigation pane, choose Accounts > Extended Fields, and then click Create Field on the Extended Fields page.
	//
	// example:
	//
	// age
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The value of the extended field. The value follows the limits on the properties of the extended field.
	//
	// example:
	//
	// 10
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
}

func (s CreateUserRequestCustomFields) String() string {
	return tea.Prettify(s)
}

func (s CreateUserRequestCustomFields) GoString() string {
	return s.String()
}

func (s *CreateUserRequestCustomFields) SetFieldName(v string) *CreateUserRequestCustomFields {
	s.FieldName = &v
	return s
}

func (s *CreateUserRequestCustomFields) SetFieldValue(v string) *CreateUserRequestCustomFields {
	s.FieldValue = &v
	return s
}

type CreateUserRequestPasswordInitializationConfig struct {
	// Specifies whether to forcibly change the password status. Default value: disabled. Valid values:
	//
	// 	- enabled: forcibly changes the password status.
	//
	// 	- disabled: does not forcibly change the password status.
	//
	// example:
	//
	// enabled
	PasswordForcedUpdateStatus *string `json:"PasswordForcedUpdateStatus,omitempty" xml:"PasswordForcedUpdateStatus,omitempty"`
	// The priority of the password initialization policy. By default, this parameter does not take effect. Valid values:
	//
	// 	- global: The password initialization policy globally takes effect.
	//
	// 	- custom: The password initialization policy takes effect based on custom settings.
	//
	// example:
	//
	// global
	PasswordInitializationPolicyPriority *string `json:"PasswordInitializationPolicyPriority,omitempty" xml:"PasswordInitializationPolicyPriority,omitempty"`
	// The password initialization method. Set the value to random,
	//
	// 	- which indicates that the password is randomly generated.
	//
	// example:
	//
	// random
	PasswordInitializationType *string `json:"PasswordInitializationType,omitempty" xml:"PasswordInitializationType,omitempty"`
	// The password notification methods.
	//
	// example:
	//
	// sms
	UserNotificationChannels []*string `json:"UserNotificationChannels,omitempty" xml:"UserNotificationChannels,omitempty" type:"Repeated"`
}

func (s CreateUserRequestPasswordInitializationConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateUserRequestPasswordInitializationConfig) GoString() string {
	return s.String()
}

func (s *CreateUserRequestPasswordInitializationConfig) SetPasswordForcedUpdateStatus(v string) *CreateUserRequestPasswordInitializationConfig {
	s.PasswordForcedUpdateStatus = &v
	return s
}

func (s *CreateUserRequestPasswordInitializationConfig) SetPasswordInitializationPolicyPriority(v string) *CreateUserRequestPasswordInitializationConfig {
	s.PasswordInitializationPolicyPriority = &v
	return s
}

func (s *CreateUserRequestPasswordInitializationConfig) SetPasswordInitializationType(v string) *CreateUserRequestPasswordInitializationConfig {
	s.PasswordInitializationType = &v
	return s
}

func (s *CreateUserRequestPasswordInitializationConfig) SetUserNotificationChannels(v []*string) *CreateUserRequestPasswordInitializationConfig {
	s.UserNotificationChannels = v
	return s
}

type CreateUserResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the account.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBody) SetRequestId(v string) *CreateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserResponseBody) SetUserId(v string) *CreateUserResponseBody {
	s.UserId = &v
	return s
}

type CreateUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponse) GoString() string {
	return s.String()
}

func (s *CreateUserResponse) SetHeaders(v map[string]*string) *CreateUserResponse {
	s.Headers = v
	return s
}

func (s *CreateUserResponse) SetStatusCode(v int32) *CreateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserResponse) SetBody(v *CreateUserResponseBody) *CreateUserResponse {
	s.Body = v
	return s
}

type DeleteApplicationRequest struct {
	// The ID of the application that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationRequest) SetApplicationId(v string) *DeleteApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *DeleteApplicationRequest) SetInstanceId(v string) *DeleteApplicationRequest {
	s.InstanceId = &v
	return s
}

type DeleteApplicationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationResponseBody) SetRequestId(v string) *DeleteApplicationResponseBody {
	s.RequestId = &v
	return s
}

type DeleteApplicationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationResponse) SetHeaders(v map[string]*string) *DeleteApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationResponse) SetStatusCode(v int32) *DeleteApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApplicationResponse) SetBody(v *DeleteApplicationResponseBody) *DeleteApplicationResponse {
	s.Body = v
	return s
}

type DeleteApplicationClientSecretRequest struct {
	// The ID of the application for which you want to delete a client key.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the client key that you want to delete for the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// sci_k52x2ru63rlkflina5utgkxxxx
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
}

func (s DeleteApplicationClientSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationClientSecretRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationClientSecretRequest) SetApplicationId(v string) *DeleteApplicationClientSecretRequest {
	s.ApplicationId = &v
	return s
}

func (s *DeleteApplicationClientSecretRequest) SetInstanceId(v string) *DeleteApplicationClientSecretRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteApplicationClientSecretRequest) SetSecretId(v string) *DeleteApplicationClientSecretRequest {
	s.SecretId = &v
	return s
}

type DeleteApplicationClientSecretResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApplicationClientSecretResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationClientSecretResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationClientSecretResponseBody) SetRequestId(v string) *DeleteApplicationClientSecretResponseBody {
	s.RequestId = &v
	return s
}

type DeleteApplicationClientSecretResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApplicationClientSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApplicationClientSecretResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationClientSecretResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationClientSecretResponse) SetHeaders(v map[string]*string) *DeleteApplicationClientSecretResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationClientSecretResponse) SetStatusCode(v int32) *DeleteApplicationClientSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApplicationClientSecretResponse) SetBody(v *DeleteApplicationClientSecretResponseBody) *DeleteApplicationClientSecretResponse {
	s.Body = v
	return s
}

type DeleteDomainRequest struct {
	// 域名ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// dm_examplexxxx
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainRequest) SetDomainId(v string) *DeleteDomainRequest {
	s.DomainId = &v
	return s
}

func (s *DeleteDomainRequest) SetInstanceId(v string) *DeleteDomainRequest {
	s.InstanceId = &v
	return s
}

type DeleteDomainResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponseBody) SetRequestId(v string) *DeleteDomainResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDomainResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponse) SetHeaders(v map[string]*string) *DeleteDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainResponse) SetStatusCode(v int32) *DeleteDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDomainResponse) SetBody(v *DeleteDomainResponseBody) *DeleteDomainResponse {
	s.Body = v
	return s
}

type DeleteDomainProxyTokenRequest struct {
	// 域名ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// dm_examplexxxxx
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// 域名代理Token ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// pt_examplexxxx
	DomainProxyTokenId *string `json:"DomainProxyTokenId,omitempty" xml:"DomainProxyTokenId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteDomainProxyTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainProxyTokenRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainProxyTokenRequest) SetDomainId(v string) *DeleteDomainProxyTokenRequest {
	s.DomainId = &v
	return s
}

func (s *DeleteDomainProxyTokenRequest) SetDomainProxyTokenId(v string) *DeleteDomainProxyTokenRequest {
	s.DomainProxyTokenId = &v
	return s
}

func (s *DeleteDomainProxyTokenRequest) SetInstanceId(v string) *DeleteDomainProxyTokenRequest {
	s.InstanceId = &v
	return s
}

type DeleteDomainProxyTokenResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDomainProxyTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainProxyTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainProxyTokenResponseBody) SetRequestId(v string) *DeleteDomainProxyTokenResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDomainProxyTokenResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDomainProxyTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDomainProxyTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainProxyTokenResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainProxyTokenResponse) SetHeaders(v map[string]*string) *DeleteDomainProxyTokenResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainProxyTokenResponse) SetStatusCode(v int32) *DeleteDomainProxyTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDomainProxyTokenResponse) SetBody(v *DeleteDomainProxyTokenResponseBody) *DeleteDomainProxyTokenResponse {
	s.Body = v
	return s
}

type DeleteGroupRequest struct {
	// The group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupRequest) SetGroupId(v string) *DeleteGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteGroupRequest) SetInstanceId(v string) *DeleteGroupRequest {
	s.InstanceId = &v
	return s
}

type DeleteGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGroupResponseBody) SetRequestId(v string) *DeleteGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupResponse) SetHeaders(v map[string]*string) *DeleteGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteGroupResponse) SetStatusCode(v int32) *DeleteGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGroupResponse) SetBody(v *DeleteGroupResponseBody) *DeleteGroupResponse {
	s.Body = v
	return s
}

type DeleteInstanceRequest struct {
	// The ID of the instance to be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) SetInstanceId(v string) *DeleteInstanceRequest {
	s.InstanceId = &v
	return s
}

type DeleteInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponse) SetHeaders(v map[string]*string) *DeleteInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceResponse) SetStatusCode(v int32) *DeleteInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceResponse) SetBody(v *DeleteInstanceResponseBody) *DeleteInstanceResponse {
	s.Body = v
	return s
}

type DeleteNetworkAccessEndpointRequest struct {
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 专属网络端点ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// nae_examplexxxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
}

func (s DeleteNetworkAccessEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkAccessEndpointRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkAccessEndpointRequest) SetInstanceId(v string) *DeleteNetworkAccessEndpointRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteNetworkAccessEndpointRequest) SetNetworkAccessEndpointId(v string) *DeleteNetworkAccessEndpointRequest {
	s.NetworkAccessEndpointId = &v
	return s
}

type DeleteNetworkAccessEndpointResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNetworkAccessEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkAccessEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNetworkAccessEndpointResponseBody) SetRequestId(v string) *DeleteNetworkAccessEndpointResponseBody {
	s.RequestId = &v
	return s
}

type DeleteNetworkAccessEndpointResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNetworkAccessEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNetworkAccessEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNetworkAccessEndpointResponse) GoString() string {
	return s.String()
}

func (s *DeleteNetworkAccessEndpointResponse) SetHeaders(v map[string]*string) *DeleteNetworkAccessEndpointResponse {
	s.Headers = v
	return s
}

func (s *DeleteNetworkAccessEndpointResponse) SetStatusCode(v int32) *DeleteNetworkAccessEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNetworkAccessEndpointResponse) SetBody(v *DeleteNetworkAccessEndpointResponseBody) *DeleteNetworkAccessEndpointResponse {
	s.Body = v
	return s
}

type DeleteOrganizationalUnitRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The organization ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
}

func (s DeleteOrganizationalUnitRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *DeleteOrganizationalUnitRequest) SetInstanceId(v string) *DeleteOrganizationalUnitRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteOrganizationalUnitRequest) SetOrganizationalUnitId(v string) *DeleteOrganizationalUnitRequest {
	s.OrganizationalUnitId = &v
	return s
}

type DeleteOrganizationalUnitResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOrganizationalUnitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteOrganizationalUnitResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOrganizationalUnitResponseBody) SetRequestId(v string) *DeleteOrganizationalUnitResponseBody {
	s.RequestId = &v
	return s
}

type DeleteOrganizationalUnitResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOrganizationalUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOrganizationalUnitResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *DeleteOrganizationalUnitResponse) SetHeaders(v map[string]*string) *DeleteOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *DeleteOrganizationalUnitResponse) SetStatusCode(v int32) *DeleteOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOrganizationalUnitResponse) SetBody(v *DeleteOrganizationalUnitResponseBody) *DeleteOrganizationalUnitResponse {
	s.Body = v
	return s
}

type DeleteUserRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserRequest) SetInstanceId(v string) *DeleteUserRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteUserRequest) SetUserId(v string) *DeleteUserRequest {
	s.UserId = &v
	return s
}

type DeleteUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserResponseBody) SetRequestId(v string) *DeleteUserResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserResponse) SetHeaders(v map[string]*string) *DeleteUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserResponse) SetStatusCode(v int32) *DeleteUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserResponse) SetBody(v *DeleteUserResponseBody) *DeleteUserResponse {
	s.Body = v
	return s
}

type DisableApplicationRequest struct {
	// The ID of the application that you want to disable.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableApplicationRequest) GoString() string {
	return s.String()
}

func (s *DisableApplicationRequest) SetApplicationId(v string) *DisableApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *DisableApplicationRequest) SetInstanceId(v string) *DisableApplicationRequest {
	s.InstanceId = &v
	return s
}

type DisableApplicationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DisableApplicationResponseBody) SetRequestId(v string) *DisableApplicationResponseBody {
	s.RequestId = &v
	return s
}

type DisableApplicationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableApplicationResponse) GoString() string {
	return s.String()
}

func (s *DisableApplicationResponse) SetHeaders(v map[string]*string) *DisableApplicationResponse {
	s.Headers = v
	return s
}

func (s *DisableApplicationResponse) SetStatusCode(v int32) *DisableApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableApplicationResponse) SetBody(v *DisableApplicationResponseBody) *DisableApplicationResponse {
	s.Body = v
	return s
}

type DisableApplicationApiInvokeRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableApplicationApiInvokeRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableApplicationApiInvokeRequest) GoString() string {
	return s.String()
}

func (s *DisableApplicationApiInvokeRequest) SetApplicationId(v string) *DisableApplicationApiInvokeRequest {
	s.ApplicationId = &v
	return s
}

func (s *DisableApplicationApiInvokeRequest) SetInstanceId(v string) *DisableApplicationApiInvokeRequest {
	s.InstanceId = &v
	return s
}

type DisableApplicationApiInvokeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableApplicationApiInvokeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableApplicationApiInvokeResponseBody) GoString() string {
	return s.String()
}

func (s *DisableApplicationApiInvokeResponseBody) SetRequestId(v string) *DisableApplicationApiInvokeResponseBody {
	s.RequestId = &v
	return s
}

type DisableApplicationApiInvokeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableApplicationApiInvokeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableApplicationApiInvokeResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableApplicationApiInvokeResponse) GoString() string {
	return s.String()
}

func (s *DisableApplicationApiInvokeResponse) SetHeaders(v map[string]*string) *DisableApplicationApiInvokeResponse {
	s.Headers = v
	return s
}

func (s *DisableApplicationApiInvokeResponse) SetStatusCode(v int32) *DisableApplicationApiInvokeResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableApplicationApiInvokeResponse) SetBody(v *DisableApplicationApiInvokeResponseBody) *DisableApplicationApiInvokeResponse {
	s.Body = v
	return s
}

type DisableApplicationClientSecretRequest struct {
	// The ID of the application for which you want to disable a client key.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The client key ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// sci_k52x2ru63rlkflina5utgkxxxx
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
}

func (s DisableApplicationClientSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableApplicationClientSecretRequest) GoString() string {
	return s.String()
}

func (s *DisableApplicationClientSecretRequest) SetApplicationId(v string) *DisableApplicationClientSecretRequest {
	s.ApplicationId = &v
	return s
}

func (s *DisableApplicationClientSecretRequest) SetInstanceId(v string) *DisableApplicationClientSecretRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableApplicationClientSecretRequest) SetSecretId(v string) *DisableApplicationClientSecretRequest {
	s.SecretId = &v
	return s
}

type DisableApplicationClientSecretResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableApplicationClientSecretResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableApplicationClientSecretResponseBody) GoString() string {
	return s.String()
}

func (s *DisableApplicationClientSecretResponseBody) SetRequestId(v string) *DisableApplicationClientSecretResponseBody {
	s.RequestId = &v
	return s
}

type DisableApplicationClientSecretResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableApplicationClientSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableApplicationClientSecretResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableApplicationClientSecretResponse) GoString() string {
	return s.String()
}

func (s *DisableApplicationClientSecretResponse) SetHeaders(v map[string]*string) *DisableApplicationClientSecretResponse {
	s.Headers = v
	return s
}

func (s *DisableApplicationClientSecretResponse) SetStatusCode(v int32) *DisableApplicationClientSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableApplicationClientSecretResponse) SetBody(v *DisableApplicationClientSecretResponseBody) *DisableApplicationClientSecretResponse {
	s.Body = v
	return s
}

type DisableApplicationProvisioningRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableApplicationProvisioningRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableApplicationProvisioningRequest) GoString() string {
	return s.String()
}

func (s *DisableApplicationProvisioningRequest) SetApplicationId(v string) *DisableApplicationProvisioningRequest {
	s.ApplicationId = &v
	return s
}

func (s *DisableApplicationProvisioningRequest) SetInstanceId(v string) *DisableApplicationProvisioningRequest {
	s.InstanceId = &v
	return s
}

type DisableApplicationProvisioningResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableApplicationProvisioningResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableApplicationProvisioningResponseBody) GoString() string {
	return s.String()
}

func (s *DisableApplicationProvisioningResponseBody) SetRequestId(v string) *DisableApplicationProvisioningResponseBody {
	s.RequestId = &v
	return s
}

type DisableApplicationProvisioningResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableApplicationProvisioningResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableApplicationProvisioningResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableApplicationProvisioningResponse) GoString() string {
	return s.String()
}

func (s *DisableApplicationProvisioningResponse) SetHeaders(v map[string]*string) *DisableApplicationProvisioningResponse {
	s.Headers = v
	return s
}

func (s *DisableApplicationProvisioningResponse) SetStatusCode(v int32) *DisableApplicationProvisioningResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableApplicationProvisioningResponse) SetBody(v *DisableApplicationProvisioningResponseBody) *DisableApplicationProvisioningResponse {
	s.Body = v
	return s
}

type DisableApplicationSsoRequest struct {
	// IDaaS的应用主键id
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// IDaaS EIAM的实例id
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableApplicationSsoRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableApplicationSsoRequest) GoString() string {
	return s.String()
}

func (s *DisableApplicationSsoRequest) SetApplicationId(v string) *DisableApplicationSsoRequest {
	s.ApplicationId = &v
	return s
}

func (s *DisableApplicationSsoRequest) SetInstanceId(v string) *DisableApplicationSsoRequest {
	s.InstanceId = &v
	return s
}

type DisableApplicationSsoResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableApplicationSsoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableApplicationSsoResponseBody) GoString() string {
	return s.String()
}

func (s *DisableApplicationSsoResponseBody) SetRequestId(v string) *DisableApplicationSsoResponseBody {
	s.RequestId = &v
	return s
}

type DisableApplicationSsoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableApplicationSsoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableApplicationSsoResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableApplicationSsoResponse) GoString() string {
	return s.String()
}

func (s *DisableApplicationSsoResponse) SetHeaders(v map[string]*string) *DisableApplicationSsoResponse {
	s.Headers = v
	return s
}

func (s *DisableApplicationSsoResponse) SetStatusCode(v int32) *DisableApplicationSsoResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableApplicationSsoResponse) SetBody(v *DisableApplicationSsoResponseBody) *DisableApplicationSsoResponse {
	s.Body = v
	return s
}

type DisableDomainProxyTokenRequest struct {
	// 域名ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// dm_examplexxxxx
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// 域名代理Token ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// pt_examplexxxx
	DomainProxyTokenId *string `json:"DomainProxyTokenId,omitempty" xml:"DomainProxyTokenId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableDomainProxyTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableDomainProxyTokenRequest) GoString() string {
	return s.String()
}

func (s *DisableDomainProxyTokenRequest) SetDomainId(v string) *DisableDomainProxyTokenRequest {
	s.DomainId = &v
	return s
}

func (s *DisableDomainProxyTokenRequest) SetDomainProxyTokenId(v string) *DisableDomainProxyTokenRequest {
	s.DomainProxyTokenId = &v
	return s
}

func (s *DisableDomainProxyTokenRequest) SetInstanceId(v string) *DisableDomainProxyTokenRequest {
	s.InstanceId = &v
	return s
}

type DisableDomainProxyTokenResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableDomainProxyTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableDomainProxyTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DisableDomainProxyTokenResponseBody) SetRequestId(v string) *DisableDomainProxyTokenResponseBody {
	s.RequestId = &v
	return s
}

type DisableDomainProxyTokenResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableDomainProxyTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableDomainProxyTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableDomainProxyTokenResponse) GoString() string {
	return s.String()
}

func (s *DisableDomainProxyTokenResponse) SetHeaders(v map[string]*string) *DisableDomainProxyTokenResponse {
	s.Headers = v
	return s
}

func (s *DisableDomainProxyTokenResponse) SetStatusCode(v int32) *DisableDomainProxyTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableDomainProxyTokenResponse) SetBody(v *DisableDomainProxyTokenResponseBody) *DisableDomainProxyTokenResponse {
	s.Body = v
	return s
}

type DisableInitDomainAutoRedirectRequest struct {
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableInitDomainAutoRedirectRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableInitDomainAutoRedirectRequest) GoString() string {
	return s.String()
}

func (s *DisableInitDomainAutoRedirectRequest) SetInstanceId(v string) *DisableInitDomainAutoRedirectRequest {
	s.InstanceId = &v
	return s
}

type DisableInitDomainAutoRedirectResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableInitDomainAutoRedirectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableInitDomainAutoRedirectResponseBody) GoString() string {
	return s.String()
}

func (s *DisableInitDomainAutoRedirectResponseBody) SetRequestId(v string) *DisableInitDomainAutoRedirectResponseBody {
	s.RequestId = &v
	return s
}

type DisableInitDomainAutoRedirectResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableInitDomainAutoRedirectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableInitDomainAutoRedirectResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableInitDomainAutoRedirectResponse) GoString() string {
	return s.String()
}

func (s *DisableInitDomainAutoRedirectResponse) SetHeaders(v map[string]*string) *DisableInitDomainAutoRedirectResponse {
	s.Headers = v
	return s
}

func (s *DisableInitDomainAutoRedirectResponse) SetStatusCode(v int32) *DisableInitDomainAutoRedirectResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableInitDomainAutoRedirectResponse) SetBody(v *DisableInitDomainAutoRedirectResponseBody) *DisableInitDomainAutoRedirectResponse {
	s.Body = v
	return s
}

type DisableUserRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DisableUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableUserRequest) GoString() string {
	return s.String()
}

func (s *DisableUserRequest) SetInstanceId(v string) *DisableUserRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableUserRequest) SetUserId(v string) *DisableUserRequest {
	s.UserId = &v
	return s
}

type DisableUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableUserResponseBody) GoString() string {
	return s.String()
}

func (s *DisableUserResponseBody) SetRequestId(v string) *DisableUserResponseBody {
	s.RequestId = &v
	return s
}

type DisableUserResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableUserResponse) GoString() string {
	return s.String()
}

func (s *DisableUserResponse) SetHeaders(v map[string]*string) *DisableUserResponse {
	s.Headers = v
	return s
}

func (s *DisableUserResponse) SetStatusCode(v int32) *DisableUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableUserResponse) SetBody(v *DisableUserResponseBody) *DisableUserResponse {
	s.Body = v
	return s
}

type EnableApplicationRequest struct {
	// The ID of the application that you want to enable.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableApplicationRequest) GoString() string {
	return s.String()
}

func (s *EnableApplicationRequest) SetApplicationId(v string) *EnableApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *EnableApplicationRequest) SetInstanceId(v string) *EnableApplicationRequest {
	s.InstanceId = &v
	return s
}

type EnableApplicationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *EnableApplicationResponseBody) SetRequestId(v string) *EnableApplicationResponseBody {
	s.RequestId = &v
	return s
}

type EnableApplicationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableApplicationResponse) GoString() string {
	return s.String()
}

func (s *EnableApplicationResponse) SetHeaders(v map[string]*string) *EnableApplicationResponse {
	s.Headers = v
	return s
}

func (s *EnableApplicationResponse) SetStatusCode(v int32) *EnableApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableApplicationResponse) SetBody(v *EnableApplicationResponseBody) *EnableApplicationResponse {
	s.Body = v
	return s
}

type EnableApplicationApiInvokeRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableApplicationApiInvokeRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableApplicationApiInvokeRequest) GoString() string {
	return s.String()
}

func (s *EnableApplicationApiInvokeRequest) SetApplicationId(v string) *EnableApplicationApiInvokeRequest {
	s.ApplicationId = &v
	return s
}

func (s *EnableApplicationApiInvokeRequest) SetInstanceId(v string) *EnableApplicationApiInvokeRequest {
	s.InstanceId = &v
	return s
}

type EnableApplicationApiInvokeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableApplicationApiInvokeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableApplicationApiInvokeResponseBody) GoString() string {
	return s.String()
}

func (s *EnableApplicationApiInvokeResponseBody) SetRequestId(v string) *EnableApplicationApiInvokeResponseBody {
	s.RequestId = &v
	return s
}

type EnableApplicationApiInvokeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableApplicationApiInvokeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableApplicationApiInvokeResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableApplicationApiInvokeResponse) GoString() string {
	return s.String()
}

func (s *EnableApplicationApiInvokeResponse) SetHeaders(v map[string]*string) *EnableApplicationApiInvokeResponse {
	s.Headers = v
	return s
}

func (s *EnableApplicationApiInvokeResponse) SetStatusCode(v int32) *EnableApplicationApiInvokeResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableApplicationApiInvokeResponse) SetBody(v *EnableApplicationApiInvokeResponseBody) *EnableApplicationApiInvokeResponse {
	s.Body = v
	return s
}

type EnableApplicationClientSecretRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The client key ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// sci_k52x2ru63rlkflina5utgkxxxx
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
}

func (s EnableApplicationClientSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableApplicationClientSecretRequest) GoString() string {
	return s.String()
}

func (s *EnableApplicationClientSecretRequest) SetApplicationId(v string) *EnableApplicationClientSecretRequest {
	s.ApplicationId = &v
	return s
}

func (s *EnableApplicationClientSecretRequest) SetInstanceId(v string) *EnableApplicationClientSecretRequest {
	s.InstanceId = &v
	return s
}

func (s *EnableApplicationClientSecretRequest) SetSecretId(v string) *EnableApplicationClientSecretRequest {
	s.SecretId = &v
	return s
}

type EnableApplicationClientSecretResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableApplicationClientSecretResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableApplicationClientSecretResponseBody) GoString() string {
	return s.String()
}

func (s *EnableApplicationClientSecretResponseBody) SetRequestId(v string) *EnableApplicationClientSecretResponseBody {
	s.RequestId = &v
	return s
}

type EnableApplicationClientSecretResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableApplicationClientSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableApplicationClientSecretResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableApplicationClientSecretResponse) GoString() string {
	return s.String()
}

func (s *EnableApplicationClientSecretResponse) SetHeaders(v map[string]*string) *EnableApplicationClientSecretResponse {
	s.Headers = v
	return s
}

func (s *EnableApplicationClientSecretResponse) SetStatusCode(v int32) *EnableApplicationClientSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableApplicationClientSecretResponse) SetBody(v *EnableApplicationClientSecretResponseBody) *EnableApplicationClientSecretResponse {
	s.Body = v
	return s
}

type EnableApplicationProvisioningRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableApplicationProvisioningRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableApplicationProvisioningRequest) GoString() string {
	return s.String()
}

func (s *EnableApplicationProvisioningRequest) SetApplicationId(v string) *EnableApplicationProvisioningRequest {
	s.ApplicationId = &v
	return s
}

func (s *EnableApplicationProvisioningRequest) SetInstanceId(v string) *EnableApplicationProvisioningRequest {
	s.InstanceId = &v
	return s
}

type EnableApplicationProvisioningResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableApplicationProvisioningResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableApplicationProvisioningResponseBody) GoString() string {
	return s.String()
}

func (s *EnableApplicationProvisioningResponseBody) SetRequestId(v string) *EnableApplicationProvisioningResponseBody {
	s.RequestId = &v
	return s
}

type EnableApplicationProvisioningResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableApplicationProvisioningResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableApplicationProvisioningResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableApplicationProvisioningResponse) GoString() string {
	return s.String()
}

func (s *EnableApplicationProvisioningResponse) SetHeaders(v map[string]*string) *EnableApplicationProvisioningResponse {
	s.Headers = v
	return s
}

func (s *EnableApplicationProvisioningResponse) SetStatusCode(v int32) *EnableApplicationProvisioningResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableApplicationProvisioningResponse) SetBody(v *EnableApplicationProvisioningResponseBody) *EnableApplicationProvisioningResponse {
	s.Body = v
	return s
}

type EnableApplicationSsoRequest struct {
	// IDaaS的应用主键id
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// IDaaS EIAM的实例id
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableApplicationSsoRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableApplicationSsoRequest) GoString() string {
	return s.String()
}

func (s *EnableApplicationSsoRequest) SetApplicationId(v string) *EnableApplicationSsoRequest {
	s.ApplicationId = &v
	return s
}

func (s *EnableApplicationSsoRequest) SetInstanceId(v string) *EnableApplicationSsoRequest {
	s.InstanceId = &v
	return s
}

type EnableApplicationSsoResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableApplicationSsoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableApplicationSsoResponseBody) GoString() string {
	return s.String()
}

func (s *EnableApplicationSsoResponseBody) SetRequestId(v string) *EnableApplicationSsoResponseBody {
	s.RequestId = &v
	return s
}

type EnableApplicationSsoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableApplicationSsoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableApplicationSsoResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableApplicationSsoResponse) GoString() string {
	return s.String()
}

func (s *EnableApplicationSsoResponse) SetHeaders(v map[string]*string) *EnableApplicationSsoResponse {
	s.Headers = v
	return s
}

func (s *EnableApplicationSsoResponse) SetStatusCode(v int32) *EnableApplicationSsoResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableApplicationSsoResponse) SetBody(v *EnableApplicationSsoResponseBody) *EnableApplicationSsoResponse {
	s.Body = v
	return s
}

type EnableDomainProxyTokenRequest struct {
	// 域名ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// dm_examplexxxxx
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// 域名代理Token ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// pt_examplexxxx
	DomainProxyTokenId *string `json:"DomainProxyTokenId,omitempty" xml:"DomainProxyTokenId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableDomainProxyTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableDomainProxyTokenRequest) GoString() string {
	return s.String()
}

func (s *EnableDomainProxyTokenRequest) SetDomainId(v string) *EnableDomainProxyTokenRequest {
	s.DomainId = &v
	return s
}

func (s *EnableDomainProxyTokenRequest) SetDomainProxyTokenId(v string) *EnableDomainProxyTokenRequest {
	s.DomainProxyTokenId = &v
	return s
}

func (s *EnableDomainProxyTokenRequest) SetInstanceId(v string) *EnableDomainProxyTokenRequest {
	s.InstanceId = &v
	return s
}

type EnableDomainProxyTokenResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableDomainProxyTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableDomainProxyTokenResponseBody) GoString() string {
	return s.String()
}

func (s *EnableDomainProxyTokenResponseBody) SetRequestId(v string) *EnableDomainProxyTokenResponseBody {
	s.RequestId = &v
	return s
}

type EnableDomainProxyTokenResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableDomainProxyTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableDomainProxyTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableDomainProxyTokenResponse) GoString() string {
	return s.String()
}

func (s *EnableDomainProxyTokenResponse) SetHeaders(v map[string]*string) *EnableDomainProxyTokenResponse {
	s.Headers = v
	return s
}

func (s *EnableDomainProxyTokenResponse) SetStatusCode(v int32) *EnableDomainProxyTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableDomainProxyTokenResponse) SetBody(v *EnableDomainProxyTokenResponseBody) *EnableDomainProxyTokenResponse {
	s.Body = v
	return s
}

type EnableInitDomainAutoRedirectRequest struct {
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableInitDomainAutoRedirectRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableInitDomainAutoRedirectRequest) GoString() string {
	return s.String()
}

func (s *EnableInitDomainAutoRedirectRequest) SetInstanceId(v string) *EnableInitDomainAutoRedirectRequest {
	s.InstanceId = &v
	return s
}

type EnableInitDomainAutoRedirectResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableInitDomainAutoRedirectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableInitDomainAutoRedirectResponseBody) GoString() string {
	return s.String()
}

func (s *EnableInitDomainAutoRedirectResponseBody) SetRequestId(v string) *EnableInitDomainAutoRedirectResponseBody {
	s.RequestId = &v
	return s
}

type EnableInitDomainAutoRedirectResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableInitDomainAutoRedirectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableInitDomainAutoRedirectResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableInitDomainAutoRedirectResponse) GoString() string {
	return s.String()
}

func (s *EnableInitDomainAutoRedirectResponse) SetHeaders(v map[string]*string) *EnableInitDomainAutoRedirectResponse {
	s.Headers = v
	return s
}

func (s *EnableInitDomainAutoRedirectResponse) SetStatusCode(v int32) *EnableInitDomainAutoRedirectResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableInitDomainAutoRedirectResponse) SetBody(v *EnableInitDomainAutoRedirectResponseBody) *EnableInitDomainAutoRedirectResponse {
	s.Body = v
	return s
}

type EnableUserRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s EnableUserRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableUserRequest) GoString() string {
	return s.String()
}

func (s *EnableUserRequest) SetInstanceId(v string) *EnableUserRequest {
	s.InstanceId = &v
	return s
}

func (s *EnableUserRequest) SetUserId(v string) *EnableUserRequest {
	s.UserId = &v
	return s
}

type EnableUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableUserResponseBody) GoString() string {
	return s.String()
}

func (s *EnableUserResponseBody) SetRequestId(v string) *EnableUserResponseBody {
	s.RequestId = &v
	return s
}

type EnableUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableUserResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableUserResponse) GoString() string {
	return s.String()
}

func (s *EnableUserResponse) SetHeaders(v map[string]*string) *EnableUserResponse {
	s.Headers = v
	return s
}

func (s *EnableUserResponse) SetStatusCode(v int32) *EnableUserResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableUserResponse) SetBody(v *EnableUserResponseBody) *EnableUserResponse {
	s.Body = v
	return s
}

type GetApplicationRequest struct {
	// The ID of the application that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationRequest) SetApplicationId(v string) *GetApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationRequest) SetInstanceId(v string) *GetApplicationRequest {
	s.InstanceId = &v
	return s
}

type GetApplicationResponseBody struct {
	// The details of the application.
	Application *GetApplicationResponseBodyApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBody) SetApplication(v *GetApplicationResponseBodyApplication) *GetApplicationResponseBody {
	s.Application = v
	return s
}

func (s *GetApplicationResponseBody) SetRequestId(v string) *GetApplicationResponseBody {
	s.RequestId = &v
	return s
}

type GetApplicationResponseBodyApplication struct {
	// The status of the Developer API feature. Valid values:
	//
	// 	- Enabled: The Developer API feature is enabled.
	//
	// 	- Disabled: The Developer API feature is disabled.
	//
	// example:
	//
	// disabled
	ApiInvokeStatus *string `json:"ApiInvokeStatus,omitempty" xml:"ApiInvokeStatus,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// SAML Application
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The origin of the application. Valid values:
	//
	// 	- urn:alibaba:idaas:app:source:template: The application is created based on a template.
	//
	// 	- urn:alibaba:idaas: The application is created based on the standard protocol.
	//
	// example:
	//
	// urn:alibaba:idaas:app:source:template
	ApplicationSourceType *string `json:"ApplicationSourceType,omitempty" xml:"ApplicationSourceType,omitempty"`
	// The ID of the template based on which the application is created. This parameter is returned only if the application is created based on a template.
	//
	// example:
	//
	// apt_rpa_tdsxxx
	ApplicationTemplateId *string `json:"ApplicationTemplateId,omitempty" xml:"ApplicationTemplateId,omitempty"`
	// The authorization type of the EIAM application. Valid values:
	//
	// 	- authorize_required: Only the user with explicit authorization can access the application.
	//
	// 	- default_all: By default, all users can access the application.
	//
	// example:
	//
	// authorize_required
	AuthorizationType *string `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
	// The client ID of the application.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The time when the application was created. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1649830226000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// The application is applicable to the test environment.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The features that are supported by the application. The value is a JSON array. Valid values:
	//
	// 	- sso: The application supports SSO.
	//
	// 	- provision: The application supports account synchronization.
	//
	// 	- api_invoke: The application supports custom APIs.
	//
	// example:
	//
	// ["sso", "provision"]
	Features *string `json:"Features,omitempty" xml:"Features,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The URL of the application icon.
	//
	// example:
	//
	// https://img.alicdn.com/imgextra/i4/O1CN01lvYwpv1aGowQXDML9_!!6000000003303-0-tps-580-580.jpg
	LogoUrl *string `json:"LogoUrl,omitempty" xml:"LogoUrl,omitempty"`
	// The service code of the cloud service that manages the application template.
	//
	// example:
	//
	// rpa
	ManagedServiceCode *string `json:"ManagedServiceCode,omitempty" xml:"ManagedServiceCode,omitempty"`
	// Indicates whether the application template is managed by a cloud service.
	//
	// example:
	//
	// true
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The type of the single sign-on (SSO) protocol. Valid values:
	//
	// 	- saml2: the Security Assertion Markup Language (SAML) 2.0 protocol.
	//
	// 	- oidc: the OpenID Connect (OIDC) protocol.
	//
	// example:
	//
	// saml2
	SsoType *string `json:"SsoType,omitempty" xml:"SsoType,omitempty"`
	// The status of the application. Valid values:
	//
	// 	- Enabled: The application is enabled.
	//
	// 	- Disabled: The application is disabled.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the application was last updated. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1649830226000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetApplicationResponseBodyApplication) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseBodyApplication) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyApplication) SetApiInvokeStatus(v string) *GetApplicationResponseBodyApplication {
	s.ApiInvokeStatus = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetApplicationId(v string) *GetApplicationResponseBodyApplication {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetApplicationName(v string) *GetApplicationResponseBodyApplication {
	s.ApplicationName = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetApplicationSourceType(v string) *GetApplicationResponseBodyApplication {
	s.ApplicationSourceType = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetApplicationTemplateId(v string) *GetApplicationResponseBodyApplication {
	s.ApplicationTemplateId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetAuthorizationType(v string) *GetApplicationResponseBodyApplication {
	s.AuthorizationType = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetClientId(v string) *GetApplicationResponseBodyApplication {
	s.ClientId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetCreateTime(v int64) *GetApplicationResponseBodyApplication {
	s.CreateTime = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetDescription(v string) *GetApplicationResponseBodyApplication {
	s.Description = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetFeatures(v string) *GetApplicationResponseBodyApplication {
	s.Features = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetInstanceId(v string) *GetApplicationResponseBodyApplication {
	s.InstanceId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetLogoUrl(v string) *GetApplicationResponseBodyApplication {
	s.LogoUrl = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetManagedServiceCode(v string) *GetApplicationResponseBodyApplication {
	s.ManagedServiceCode = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetServiceManaged(v bool) *GetApplicationResponseBodyApplication {
	s.ServiceManaged = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetSsoType(v string) *GetApplicationResponseBodyApplication {
	s.SsoType = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetStatus(v string) *GetApplicationResponseBodyApplication {
	s.Status = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetUpdateTime(v int64) *GetApplicationResponseBodyApplication {
	s.UpdateTime = &v
	return s
}

type GetApplicationResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationResponse) SetHeaders(v map[string]*string) *GetApplicationResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationResponse) SetStatusCode(v int32) *GetApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationResponse) SetBody(v *GetApplicationResponseBody) *GetApplicationResponse {
	s.Body = v
	return s
}

type GetApplicationGrantScopeRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetApplicationGrantScopeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationGrantScopeRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationGrantScopeRequest) SetApplicationId(v string) *GetApplicationGrantScopeRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationGrantScopeRequest) SetInstanceId(v string) *GetApplicationGrantScopeRequest {
	s.InstanceId = &v
	return s
}

type GetApplicationGrantScopeResponseBody struct {
	// The permissions of the Developer API feature.
	ApplicationGrantScope *GetApplicationGrantScopeResponseBodyApplicationGrantScope `json:"ApplicationGrantScope,omitempty" xml:"ApplicationGrantScope,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationGrantScopeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationGrantScopeResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationGrantScopeResponseBody) SetApplicationGrantScope(v *GetApplicationGrantScopeResponseBodyApplicationGrantScope) *GetApplicationGrantScopeResponseBody {
	s.ApplicationGrantScope = v
	return s
}

func (s *GetApplicationGrantScopeResponseBody) SetRequestId(v string) *GetApplicationGrantScopeResponseBody {
	s.RequestId = &v
	return s
}

type GetApplicationGrantScopeResponseBodyApplicationGrantScope struct {
	// The permissions of the Developer API feature.
	GrantScopes []*string `json:"GrantScopes,omitempty" xml:"GrantScopes,omitempty" type:"Repeated"`
}

func (s GetApplicationGrantScopeResponseBodyApplicationGrantScope) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationGrantScopeResponseBodyApplicationGrantScope) GoString() string {
	return s.String()
}

func (s *GetApplicationGrantScopeResponseBodyApplicationGrantScope) SetGrantScopes(v []*string) *GetApplicationGrantScopeResponseBodyApplicationGrantScope {
	s.GrantScopes = v
	return s
}

type GetApplicationGrantScopeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationGrantScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationGrantScopeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationGrantScopeResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationGrantScopeResponse) SetHeaders(v map[string]*string) *GetApplicationGrantScopeResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationGrantScopeResponse) SetStatusCode(v int32) *GetApplicationGrantScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationGrantScopeResponse) SetBody(v *GetApplicationGrantScopeResponseBody) *GetApplicationGrantScopeResponse {
	s.Body = v
	return s
}

type GetApplicationProvisioningConfigRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetApplicationProvisioningConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationProvisioningConfigRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningConfigRequest) SetApplicationId(v string) *GetApplicationProvisioningConfigRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationProvisioningConfigRequest) SetInstanceId(v string) *GetApplicationProvisioningConfigRequest {
	s.InstanceId = &v
	return s
}

type GetApplicationProvisioningConfigResponseBody struct {
	// The configuration of the account synchronization feature for the application.
	ApplicationProvisioningConfig *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig `json:"ApplicationProvisioningConfig,omitempty" xml:"ApplicationProvisioningConfig,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationProvisioningConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationProvisioningConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningConfigResponseBody) SetApplicationProvisioningConfig(v *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) *GetApplicationProvisioningConfigResponseBody {
	s.ApplicationProvisioningConfig = v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBody) SetRequestId(v string) *GetApplicationProvisioningConfigResponseBody {
	s.RequestId = &v
	return s
}

type GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig struct {
	// The ID of the application.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The configuration of the custom event callback protocol of IDaaS.
	CallbackProvisioningConfig *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig `json:"CallbackProvisioningConfig,omitempty" xml:"CallbackProvisioningConfig,omitempty" type:"Struct"`
	// The rendering mode of the account synchronization page. Valid values:
	//
	// 	- standard: standard mode
	//
	// 	- template: template mode
	//
	// example:
	//
	// standard
	ConfigOperateMode *string `json:"ConfigOperateMode,omitempty" xml:"ConfigOperateMode,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The public key endpoint for signature verification of the synchronization callback information.
	//
	// example:
	//
	// https://eiam-api-cn-hangzhou.aliyuncs.com/v2/idaas_ue2jvisn35ea5lmthk267xxxxx/app_mkv7rgt4d7i4u7zqtzev2mxxxx/provisioning/jwks
	ProvisionJwksEndpoint *string `json:"ProvisionJwksEndpoint,omitempty" xml:"ProvisionJwksEndpoint,omitempty"`
	// Indicates whether the password is synchronized in IDaaS user event callbacks. Valid values:
	//
	// 	- true: The password is synchronized.
	//
	// 	- false: The password is not synchronized.
	//
	// example:
	//
	// true
	ProvisionPassword *bool `json:"ProvisionPassword,omitempty" xml:"ProvisionPassword,omitempty"`
	// The synchronization protocol type of the application. Valid values:
	//
	// 	- idaas_callback: custom event callback protocol of IDaaS.
	//
	// 	- scim2: System for Cross-domain Identity Management (SCIM) protocol.
	//
	// example:
	//
	// idaas_callback
	ProvisionProtocolType *string `json:"ProvisionProtocolType,omitempty" xml:"ProvisionProtocolType,omitempty"`
	// The configuration of SCIM-based IDaaS synchronization.
	ScimProvisioningConfig *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig `json:"ScimProvisioningConfig,omitempty" xml:"ScimProvisioningConfig,omitempty" type:"Struct"`
	// The status of the IDaaS account synchronization feature. Valid values:
	//
	// 	- enabled: The feature is enabled.
	//
	// 	- disabled: The feature is disabled.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) SetApplicationId(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) SetCallbackProvisioningConfig(v *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig {
	s.CallbackProvisioningConfig = v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) SetConfigOperateMode(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig {
	s.ConfigOperateMode = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) SetInstanceId(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig {
	s.InstanceId = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) SetProvisionJwksEndpoint(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig {
	s.ProvisionJwksEndpoint = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) SetProvisionPassword(v bool) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig {
	s.ProvisionPassword = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) SetProvisionProtocolType(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig {
	s.ProvisionProtocolType = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) SetScimProvisioningConfig(v *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig {
	s.ScimProvisioningConfig = v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig) SetStatus(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfig {
	s.Status = &v
	return s
}

type GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig struct {
	// The URL that the application uses to receive IDaaS event callbacks.
	//
	// example:
	//
	// https://example.com/event/callback
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// The symmetric key for IDaaS event callbacks. The key is an AES-256 encryption key in the HEX format.
	//
	// example:
	//
	// 1adfdfdfd******111
	EncryptKey *string `json:"EncryptKey,omitempty" xml:"EncryptKey,omitempty"`
	// Indicates whether IDaaS event callback messages are encrypted. Valid values:
	//
	// 	- true: The messages are encrypted.
	//
	// 	- false: The messages are transmitted in plaintext.
	//
	// example:
	//
	// true
	EncryptRequired *bool `json:"EncryptRequired,omitempty" xml:"EncryptRequired,omitempty"`
	// The list of types of IDaaS event callback messages that are supported by the listener.
	ListenEventScopes []*string `json:"ListenEventScopes,omitempty" xml:"ListenEventScopes,omitempty" type:"Repeated"`
}

func (s GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig) SetCallbackUrl(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig {
	s.CallbackUrl = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig) SetEncryptKey(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig {
	s.EncryptKey = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig) SetEncryptRequired(v bool) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig {
	s.EncryptRequired = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig) SetListenEventScopes(v []*string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigCallbackProvisioningConfig {
	s.ListenEventScopes = v
	return s
}

type GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig struct {
	// The configuration parameters related to SCIM-based synchronization.
	AuthnConfiguration *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfiguration `json:"AuthnConfiguration,omitempty" xml:"AuthnConfiguration,omitempty" type:"Struct"`
	// The full synchronization scope of the SCIM protocol. Valid value:
	//
	// 	- urn:alibaba:idaas:app:scim:User:PUSH: full account data synchronization.
	FullPushScopes []*string `json:"FullPushScopes,omitempty" xml:"FullPushScopes,omitempty" type:"Repeated"`
	// The resource operations of the SCIM protocol. Valid values:
	//
	// 	- urn:alibaba:idaas:app:scim:User:CREATE: account creation.
	//
	// 	- urn:alibaba:idaas:app:scim:User:UPDATE: account update.
	//
	// 	- urn:alibaba:idaas:app:scim:User:DELETE: account deletion.
	ProvisioningActions []*string `json:"ProvisioningActions,omitempty" xml:"ProvisioningActions,omitempty" type:"Repeated"`
	// The base URL that the application uses to receive the SCIM protocol for IDaaS synchronization.
	//
	// example:
	//
	// https://example.com/scim
	ScimBaseUrl *string `json:"ScimBaseUrl,omitempty" xml:"ScimBaseUrl,omitempty"`
}

func (s GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig) SetAuthnConfiguration(v *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfiguration) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig {
	s.AuthnConfiguration = v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig) SetFullPushScopes(v []*string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig {
	s.FullPushScopes = v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig) SetProvisioningActions(v []*string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig {
	s.ProvisioningActions = v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig) SetScimBaseUrl(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfig {
	s.ScimBaseUrl = &v
	return s
}

type GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfiguration struct {
	// The authentication mode of the SCIM protocol. Valid value:
	//
	// 	- oauth2: OAuth2.0 mode.
	//
	// example:
	//
	// oauth2
	AuthnMode *string `json:"AuthnMode,omitempty" xml:"AuthnMode,omitempty"`
	// The configuration parameters related to authorization.
	//
	// 	- If the GrantType parameter is set to client_credentials, the configuration parameters ClientId, ClientSecret, and AuthnMethod are returned.
	//
	// 	- If the GrantType parameter is set to bearer_token, the configuration parameter AccessToken is returned.
	AuthnParam *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam `json:"AuthnParam,omitempty" xml:"AuthnParam,omitempty" type:"Struct"`
	// The grant type of the SCIM protocol. Valid values:
	//
	// 	- client_credentials: client mode.
	//
	// 	- bearer_token: key mode.
	//
	// example:
	//
	// bearer_token
	GrantType *string `json:"GrantType,omitempty" xml:"GrantType,omitempty"`
}

func (s GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfiguration) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfiguration) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfiguration) SetAuthnMode(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfiguration {
	s.AuthnMode = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfiguration) SetAuthnParam(v *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfiguration {
	s.AuthnParam = v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfiguration) SetGrantType(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfiguration {
	s.GrantType = &v
	return s
}

type GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam struct {
	// The access token. This parameter is returned when the GrantType parameter is set to bearer_token.
	//
	// example:
	//
	// k52x2ru63rlkflina5utgkxxxx
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// The authentication mode of the SCIM protocol. Valid values:
	//
	// 	- client_secret_basic: The client secret is passed in the request header.
	//
	// 	- client_secret_post: The client secret is passed in the request body.
	//
	// example:
	//
	// client_secret_basic
	AuthnMethod *string `json:"AuthnMethod,omitempty" xml:"AuthnMethod,omitempty"`
	// The client ID of the application.
	//
	// example:
	//
	// mkv7rgt4d7i4u7zqtzev2mxxxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The client secret of the application.
	//
	// example:
	//
	// CSEHDcHcrUKHw1CuxkJEHPveWRXBGqVqRsxxxx
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	// The token endpoint.
	//
	// example:
	//
	// https://www.example.com/oauth/token
	TokenEndpoint *string `json:"TokenEndpoint,omitempty" xml:"TokenEndpoint,omitempty"`
}

func (s GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam) SetAccessToken(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam {
	s.AccessToken = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam) SetAuthnMethod(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam {
	s.AuthnMethod = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam) SetClientId(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam {
	s.ClientId = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam) SetClientSecret(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam {
	s.ClientSecret = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam) SetTokenEndpoint(v string) *GetApplicationProvisioningConfigResponseBodyApplicationProvisioningConfigScimProvisioningConfigAuthnConfigurationAuthnParam {
	s.TokenEndpoint = &v
	return s
}

type GetApplicationProvisioningConfigResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationProvisioningConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationProvisioningConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationProvisioningConfigResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningConfigResponse) SetHeaders(v map[string]*string) *GetApplicationProvisioningConfigResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationProvisioningConfigResponse) SetStatusCode(v int32) *GetApplicationProvisioningConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationProvisioningConfigResponse) SetBody(v *GetApplicationProvisioningConfigResponseBody) *GetApplicationProvisioningConfigResponse {
	s.Body = v
	return s
}

type GetApplicationProvisioningScopeRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetApplicationProvisioningScopeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationProvisioningScopeRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningScopeRequest) SetApplicationId(v string) *GetApplicationProvisioningScopeRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationProvisioningScopeRequest) SetInstanceId(v string) *GetApplicationProvisioningScopeRequest {
	s.InstanceId = &v
	return s
}

type GetApplicationProvisioningScopeResponseBody struct {
	// The scope of account synchronization.
	ApplicationProvisioningScope *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope `json:"ApplicationProvisioningScope,omitempty" xml:"ApplicationProvisioningScope,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationProvisioningScopeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationProvisioningScopeResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningScopeResponseBody) SetApplicationProvisioningScope(v *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope) *GetApplicationProvisioningScopeResponseBody {
	s.ApplicationProvisioningScope = v
	return s
}

func (s *GetApplicationProvisioningScopeResponseBody) SetRequestId(v string) *GetApplicationProvisioningScopeResponseBody {
	s.RequestId = &v
	return s
}

type GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope struct {
	// 同步授权的组列表
	GroupIds []*string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// 租户最大授权主体quota数量
	MaxQuota *int32 `json:"MaxQuota,omitempty" xml:"MaxQuota,omitempty"`
	// The list of organizational units that are authorized for account synchronization.
	OrganizationalUnitIds []*string `json:"OrganizationalUnitIds,omitempty" xml:"OrganizationalUnitIds,omitempty" type:"Repeated"`
	// 已使用授权主体quota数量
	UsedQuota *int32 `json:"UsedQuota,omitempty" xml:"UsedQuota,omitempty"`
}

func (s GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope) SetGroupIds(v []*string) *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope {
	s.GroupIds = v
	return s
}

func (s *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope) SetMaxQuota(v int32) *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope {
	s.MaxQuota = &v
	return s
}

func (s *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope) SetOrganizationalUnitIds(v []*string) *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope {
	s.OrganizationalUnitIds = v
	return s
}

func (s *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope) SetUsedQuota(v int32) *GetApplicationProvisioningScopeResponseBodyApplicationProvisioningScope {
	s.UsedQuota = &v
	return s
}

type GetApplicationProvisioningScopeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationProvisioningScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationProvisioningScopeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationProvisioningScopeResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningScopeResponse) SetHeaders(v map[string]*string) *GetApplicationProvisioningScopeResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationProvisioningScopeResponse) SetStatusCode(v int32) *GetApplicationProvisioningScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationProvisioningScopeResponse) SetBody(v *GetApplicationProvisioningScopeResponseBody) *GetApplicationProvisioningScopeResponse {
	s.Body = v
	return s
}

type GetApplicationSsoConfigRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetApplicationSsoConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationSsoConfigRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationSsoConfigRequest) SetApplicationId(v string) *GetApplicationSsoConfigRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationSsoConfigRequest) SetInstanceId(v string) *GetApplicationSsoConfigRequest {
	s.InstanceId = &v
	return s
}

type GetApplicationSsoConfigResponseBody struct {
	// The SSO configuration information of the application.
	ApplicationSsoConfig *GetApplicationSsoConfigResponseBodyApplicationSsoConfig `json:"ApplicationSsoConfig,omitempty" xml:"ApplicationSsoConfig,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationSsoConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationSsoConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationSsoConfigResponseBody) SetApplicationSsoConfig(v *GetApplicationSsoConfigResponseBodyApplicationSsoConfig) *GetApplicationSsoConfigResponseBody {
	s.ApplicationSsoConfig = v
	return s
}

func (s *GetApplicationSsoConfigResponseBody) SetRequestId(v string) *GetApplicationSsoConfigResponseBody {
	s.RequestId = &v
	return s
}

type GetApplicationSsoConfigResponseBodyApplicationSsoConfig struct {
	// The initial SSO method. Valid values:
	//
	// 	- only_app_init_sso: Only application-initiated SSO is allowed. This method is selected by default when the SSO protocol of the application is an OIDC protocol. If this method is selected when the SSO protocol of the application is SAML, the InitLoginUrl parameter is required.
	//
	// 	- idaas_or_app_init_sso: IDaaS-initiated SSO and application-initiated SSO are allowed. This method is selected by default when the SSO protocol of the application is SAML. If this method is selected when the SSO protocol of the application is an OIDC protocol, the InitLoginUrl parameter is required.
	//
	// example:
	//
	// only_app_init_sso
	InitLoginType *string `json:"InitLoginType,omitempty" xml:"InitLoginType,omitempty"`
	// The initial webhook URL of SSO. This parameter is required when the SSO protocol of the application is an OIDC protocol and the InitLoginType parameters is set to idaas_or_app_init_sso or when the SSO protocol of the application is SAML and the InitLoginType parameter is set to only_app_init_sso.
	//
	// example:
	//
	// http://127.0.0.1:8000/start_login?enterprise_code=ABCDEF
	InitLoginUrl *string `json:"InitLoginUrl,omitempty" xml:"InitLoginUrl,omitempty"`
	// The Open ID Connect (OIDC)-based SSO configuration attributes of the application. This parameter is returned only when the SSO protocol of the application is an OIDC protocol.
	OidcSsoConfig *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig `json:"OidcSsoConfig,omitempty" xml:"OidcSsoConfig,omitempty" type:"Struct"`
	// The configuration of the metadata endpoint provided by the application.
	ProtocolEndpointDomain *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain `json:"ProtocolEndpointDomain,omitempty" xml:"ProtocolEndpointDomain,omitempty" type:"Struct"`
	// The Security Assertion Markup Language (SAML)-based SSO configuration attributes of the application. This parameter is returned only when the SSO protocol of the application is SAML 2.0.
	SamlSsoConfig *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig `json:"SamlSsoConfig,omitempty" xml:"SamlSsoConfig,omitempty" type:"Struct"`
	// The SSO feature status of the application. Valid values:
	//
	// 	- enabled: The feature is enabled.
	//
	// 	- disabled: The feature is disabled.
	//
	// example:
	//
	// enabled
	SsoStatus *string `json:"SsoStatus,omitempty" xml:"SsoStatus,omitempty"`
}

func (s GetApplicationSsoConfigResponseBodyApplicationSsoConfig) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationSsoConfigResponseBodyApplicationSsoConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfig) SetInitLoginType(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfig {
	s.InitLoginType = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfig) SetInitLoginUrl(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfig {
	s.InitLoginUrl = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfig) SetOidcSsoConfig(v *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) *GetApplicationSsoConfigResponseBodyApplicationSsoConfig {
	s.OidcSsoConfig = v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfig) SetProtocolEndpointDomain(v *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) *GetApplicationSsoConfigResponseBodyApplicationSsoConfig {
	s.ProtocolEndpointDomain = v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfig) SetSamlSsoConfig(v *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) *GetApplicationSsoConfigResponseBodyApplicationSsoConfig {
	s.SamlSsoConfig = v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfig) SetSsoStatus(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfig {
	s.SsoStatus = &v
	return s
}

type GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig struct {
	// The validity period of the issued access token. Unit: seconds. Default value: 1200.
	//
	// example:
	//
	// 1200
	AccessTokenEffectiveTime *int64 `json:"AccessTokenEffectiveTime,omitempty" xml:"AccessTokenEffectiveTime,omitempty"`
	// The validity period of the issued code. Unit: seconds. Default value: 60.
	//
	// example:
	//
	// 60
	CodeEffectiveTime *int64 `json:"CodeEffectiveTime,omitempty" xml:"CodeEffectiveTime,omitempty"`
	// The custom claims that are returned for the ID token.
	CustomClaims []*GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfigCustomClaims `json:"CustomClaims,omitempty" xml:"CustomClaims,omitempty" type:"Repeated"`
	// The scopes of user attributes that can be returned for the UserInfo endpoint or ID token.
	//
	// example:
	//
	// profile，email
	GrantScopes []*string `json:"GrantScopes,omitempty" xml:"GrantScopes,omitempty" type:"Repeated"`
	// The list of grant types that are supported for OIDC protocols.
	//
	// example:
	//
	// authorization_code
	GrantTypes []*string `json:"GrantTypes,omitempty" xml:"GrantTypes,omitempty" type:"Repeated"`
	// The validity period of the issued ID token. Unit: seconds. Default value: 300.
	//
	// example:
	//
	// 1200
	IdTokenEffectiveTime *int64 `json:"IdTokenEffectiveTime,omitempty" xml:"IdTokenEffectiveTime,omitempty"`
	// The ID of the identity authentication source in password mode. This parameter is returned only when the value of the GrantTypes parameter includes the password mode.
	//
	// example:
	//
	// ia_password
	PasswordAuthenticationSourceId *string `json:"PasswordAuthenticationSourceId,omitempty" xml:"PasswordAuthenticationSourceId,omitempty"`
	// Indicates whether time-based one-time password (TOTP) authentication is required in password mode. This parameter is returned only when the value of the GrantTypes parameter includes the password mode.
	//
	// example:
	//
	// true
	PasswordTotpMfaRequired *bool `json:"PasswordTotpMfaRequired,omitempty" xml:"PasswordTotpMfaRequired,omitempty"`
	// The algorithms that are used to calculate the code challenge for PKCE.
	//
	// example:
	//
	// S256
	PkceChallengeMethods []*string `json:"PkceChallengeMethods,omitempty" xml:"PkceChallengeMethods,omitempty" type:"Repeated"`
	// Indicates whether the SSO of the application requires Proof Key for Code Exchange (PKCE) (RFC 7636).
	//
	// example:
	//
	// true
	PkceRequired *bool `json:"PkceRequired,omitempty" xml:"PkceRequired,omitempty"`
	// The list of logout redirect URIs that are supported by the application.
	PostLogoutRedirectUris []*string `json:"PostLogoutRedirectUris,omitempty" xml:"PostLogoutRedirectUris,omitempty" type:"Repeated"`
	// The list of redirect URIs that are supported by the application.
	RedirectUris []*string `json:"RedirectUris,omitempty" xml:"RedirectUris,omitempty" type:"Repeated"`
	// The validity period of the issued refresh token. Unit: seconds. Default value: 86400.
	//
	// example:
	//
	// 86400
	RefreshTokenEffective *int64 `json:"RefreshTokenEffective,omitempty" xml:"RefreshTokenEffective,omitempty"`
	// The response types that are supported by the application. This parameter is returned when the value of the GrantTypes parameter includes the implicit mode.
	//
	// example:
	//
	// token id_token
	ResponseTypes []*string `json:"ResponseTypes,omitempty" xml:"ResponseTypes,omitempty" type:"Repeated"`
	// The custom expression that is used to generate the subject ID returned for the ID token.
	//
	// example:
	//
	// user.userid
	SubjectIdExpression *string `json:"SubjectIdExpression,omitempty" xml:"SubjectIdExpression,omitempty"`
}

func (s GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetAccessTokenEffectiveTime(v int64) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.AccessTokenEffectiveTime = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetCodeEffectiveTime(v int64) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.CodeEffectiveTime = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetCustomClaims(v []*GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfigCustomClaims) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.CustomClaims = v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetGrantScopes(v []*string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.GrantScopes = v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetGrantTypes(v []*string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.GrantTypes = v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetIdTokenEffectiveTime(v int64) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.IdTokenEffectiveTime = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetPasswordAuthenticationSourceId(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.PasswordAuthenticationSourceId = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetPasswordTotpMfaRequired(v bool) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.PasswordTotpMfaRequired = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetPkceChallengeMethods(v []*string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.PkceChallengeMethods = v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetPkceRequired(v bool) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.PkceRequired = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetPostLogoutRedirectUris(v []*string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.PostLogoutRedirectUris = v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetRedirectUris(v []*string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.RedirectUris = v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetRefreshTokenEffective(v int64) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.RefreshTokenEffective = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetResponseTypes(v []*string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.ResponseTypes = v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig) SetSubjectIdExpression(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfig {
	s.SubjectIdExpression = &v
	return s
}

type GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfigCustomClaims struct {
	// The claim name.
	//
	// example:
	//
	// userOuIds
	ClaimName *string `json:"ClaimName,omitempty" xml:"ClaimName,omitempty"`
	// The expression that is used to generate the value of the claim.
	//
	// example:
	//
	// ObjectToJsonString(user.organizationalUnits)
	ClaimValueExpression *string `json:"ClaimValueExpression,omitempty" xml:"ClaimValueExpression,omitempty"`
}

func (s GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfigCustomClaims) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfigCustomClaims) GoString() string {
	return s.String()
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfigCustomClaims) SetClaimName(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfigCustomClaims {
	s.ClaimName = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfigCustomClaims) SetClaimValueExpression(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigOidcSsoConfigCustomClaims {
	s.ClaimValueExpression = &v
	return s
}

type GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain struct {
	// The OAuth2.0 authorization endpoint. This parameter is returned only when the SSO protocol of the application is an OIDC protocol.
	//
	// example:
	//
	// https://l1seshcn.aliyunidaas.com/login/app/app_mltta64q65enci54slingvvsgq/oauth2/authorize
	Oauth2AuthorizationEndpoint *string `json:"Oauth2AuthorizationEndpoint,omitempty" xml:"Oauth2AuthorizationEndpoint,omitempty"`
	// The OAuth2.0 device authorization endpoint. This parameter is returned only when the SSO protocol of the application is an OIDC protocol.
	//
	// example:
	//
	// https://eiam-api-cn-hangzhou.aliyuncs.com/v2/idaas_ue2jvisn35ea5lmthk2676rypm/app_mltta64q65enci54slingvvsgq/oauth2/device/code
	Oauth2DeviceAuthorizationEndpoint *string `json:"Oauth2DeviceAuthorizationEndpoint,omitempty" xml:"Oauth2DeviceAuthorizationEndpoint,omitempty"`
	// The OAuth2.0 token revocation endpoint. This parameter is returned only when the SSO protocol of the application is an OIDC protocol.
	//
	// example:
	//
	// https://eiam-api-cn-hangzhou.aliyuncs.com/v2/idaas_ue2jvisn35ea5lmthk2676rypm/app_mltta64q65enci54slingvvsgq/oauth2/revoke
	Oauth2RevokeEndpoint *string `json:"Oauth2RevokeEndpoint,omitempty" xml:"Oauth2RevokeEndpoint,omitempty"`
	// The OAuth2.0 token endpoint. This parameter is returned only when the SSO protocol of the application is an OIDC protocol.
	//
	// example:
	//
	// https://eiam-api-cn-hangzhou.aliyuncs.com/v2/idaas_ue2jvisn35ea5lmthk2676rypm/app_mltta64q65enci54slingvvsgq/oauth2/token
	Oauth2TokenEndpoint *string `json:"Oauth2TokenEndpoint,omitempty" xml:"Oauth2TokenEndpoint,omitempty"`
	// The OIDC UserInfo endpoint. This parameter is returned only when the SSO protocol of the application is an OIDC protocol.
	//
	// example:
	//
	// https://eiam-api-cn-hangzhou.aliyuncs.com/v2/idaas_ue2jvisn35ea5lmthk2676rypm/app_mltta64q65enci54slingvvsgq/oauth2/userinfo
	Oauth2UserinfoEndpoint *string `json:"Oauth2UserinfoEndpoint,omitempty" xml:"Oauth2UserinfoEndpoint,omitempty"`
	// The information about the OIDC issuer. This parameter is returned only when the SSO protocol of the application is an OIDC protocol.
	//
	// example:
	//
	// https://eiam-api-cn-hangzhou.aliyuncs.com/v2/idaas_ue2jvisn35ea5lmthk2676rypm/app_mltta64q65enci54slingvvsgq/oidc
	OidcIssuer *string `json:"OidcIssuer,omitempty" xml:"OidcIssuer,omitempty"`
	// The JSON Web Key Set (JWKS) URL of the OIDC issuer. This parameter is returned only when the SSO protocol of the application is an OIDC protocol.
	//
	// example:
	//
	// https://eiam-api-cn-hangzhou.aliyuncs.com/v2/idaas_ue2jvisn35ea5lmthk2676rypm/app_mltta64q65enci54slingvvsgq/oidc/jwks
	OidcJwksEndpoint *string `json:"OidcJwksEndpoint,omitempty" xml:"OidcJwksEndpoint,omitempty"`
	// The OIDC relying party (RP)-initiated logout endpoint. This parameter is returned only when the SSO protocol of the application is an OIDC protocol.
	//
	// example:
	//
	// https://l1seshcn.aliyunidaas.com/login/app/app_mltta64q65enci54slingvvsgq/oauth2/logout
	OidcLogoutEndpoint *string `json:"OidcLogoutEndpoint,omitempty" xml:"OidcLogoutEndpoint,omitempty"`
	// The metadata URL of the SAML protocol. This parameter is returned only when the SSO protocol of the application is SAML 2.0.
	//
	// example:
	//
	// https://l1seshcn.aliyunidaas.com/api/v2/app_mltuxdwd4lq4eer6tmtlmaxm5e/saml2/meta
	SamlMetaEndpoint *string `json:"SamlMetaEndpoint,omitempty" xml:"SamlMetaEndpoint,omitempty"`
	// The request receiving URL of the SAML protocol. This parameter is returned only when the SSO protocol of the application is SAML 2.0.
	//
	// example:
	//
	// https://l1seshcn.aliyunidaas.com/login/app/app_mltuxdwd4lq4eer6tmtlmaxm5e/saml2/sso
	SamlSsoEndpoint *string `json:"SamlSsoEndpoint,omitempty" xml:"SamlSsoEndpoint,omitempty"`
}

func (s GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) GoString() string {
	return s.String()
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) SetOauth2AuthorizationEndpoint(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain {
	s.Oauth2AuthorizationEndpoint = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) SetOauth2DeviceAuthorizationEndpoint(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain {
	s.Oauth2DeviceAuthorizationEndpoint = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) SetOauth2RevokeEndpoint(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain {
	s.Oauth2RevokeEndpoint = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) SetOauth2TokenEndpoint(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain {
	s.Oauth2TokenEndpoint = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) SetOauth2UserinfoEndpoint(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain {
	s.Oauth2UserinfoEndpoint = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) SetOidcIssuer(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain {
	s.OidcIssuer = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) SetOidcJwksEndpoint(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain {
	s.OidcJwksEndpoint = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) SetOidcLogoutEndpoint(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain {
	s.OidcLogoutEndpoint = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) SetSamlMetaEndpoint(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain {
	s.SamlMetaEndpoint = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain) SetSamlSsoEndpoint(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigProtocolEndpointDomain {
	s.SamlSsoEndpoint = &v
	return s
}

type GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig struct {
	// assertion是否签名
	//
	// example:
	//
	// true
	AssertionSigned *bool `json:"AssertionSigned,omitempty" xml:"AssertionSigned,omitempty"`
	// The additional user attributes in the SAML assertion.
	AttributeStatements []*GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigAttributeStatements `json:"AttributeStatements,omitempty" xml:"AttributeStatements,omitempty" type:"Repeated"`
	// The default value of the RelayState attribute. If the SSO request is initiated in EIAM, the RelayState attribute in the SAML response is set to this default value.
	//
	// example:
	//
	// https://home.console.aliyun.com
	DefaultRelayState *string `json:"DefaultRelayState,omitempty" xml:"DefaultRelayState,omitempty"`
	// The Format attribute of the NameID element in the SAML assertion. Valid values:
	//
	// 	- urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified: No format is specified. How to resolve the NameID element depends on the application.
	//
	// 	- urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress: The NameID element must be an email address.
	//
	// 	- urn:oasis:names:tc:SAML:2.0:nameid-format:persistent: The NameID element must be persistent.
	//
	// 	- urn:oasis:names:tc:SAML:2.0:nameid-format:transient: The NameID element must be transient.
	//
	// example:
	//
	// urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified
	NameIdFormat *string `json:"NameIdFormat,omitempty" xml:"NameIdFormat,omitempty"`
	// The expression that is used to generate the value of NameID in the SAML assertion.
	//
	// example:
	//
	// user.username
	NameIdValueExpression *string `json:"NameIdValueExpression,omitempty" xml:"NameIdValueExpression,omitempty"`
	// response是否签名
	//
	// example:
	//
	// true
	ResponseSigned *bool `json:"ResponseSigned,omitempty" xml:"ResponseSigned,omitempty"`
	// The algorithm that is used to calculate the signature for the SAML assertion.
	//
	// example:
	//
	// RSA-SHA256
	SignatureAlgorithm *string `json:"SignatureAlgorithm,omitempty" xml:"SignatureAlgorithm,omitempty"`
	// The entity ID of the application in SAML. The application assumes the role of service provider.
	//
	// example:
	//
	// urn:alibaba:cloudcomputing
	SpEntityId *string `json:"SpEntityId,omitempty" xml:"SpEntityId,omitempty"`
	// The Assertion Consumer Service (ACS) URL of the application in SAML. The application assumes the role of service provider.
	//
	// example:
	//
	// https://signin.aliyun.com/saml-role/sso
	SpSsoAcsUrl *string `json:"SpSsoAcsUrl,omitempty" xml:"SpSsoAcsUrl,omitempty"`
}

func (s GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) SetAssertionSigned(v bool) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig {
	s.AssertionSigned = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) SetAttributeStatements(v []*GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigAttributeStatements) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig {
	s.AttributeStatements = v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) SetDefaultRelayState(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig {
	s.DefaultRelayState = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) SetNameIdFormat(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig {
	s.NameIdFormat = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) SetNameIdValueExpression(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig {
	s.NameIdValueExpression = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) SetResponseSigned(v bool) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig {
	s.ResponseSigned = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) SetSignatureAlgorithm(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig {
	s.SignatureAlgorithm = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) SetSpEntityId(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig {
	s.SpEntityId = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig) SetSpSsoAcsUrl(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfig {
	s.SpSsoAcsUrl = &v
	return s
}

type GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigAttributeStatements struct {
	// The attribute name.
	//
	// example:
	//
	// https://www.aliyun.com/SAML-Role/Attributes/RoleSessionName
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	// The expression that is used to generate the value of the attribute.
	//
	// example:
	//
	// user.username
	AttributeValueExpression *string `json:"AttributeValueExpression,omitempty" xml:"AttributeValueExpression,omitempty"`
}

func (s GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigAttributeStatements) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigAttributeStatements) GoString() string {
	return s.String()
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigAttributeStatements) SetAttributeName(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigAttributeStatements {
	s.AttributeName = &v
	return s
}

func (s *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigAttributeStatements) SetAttributeValueExpression(v string) *GetApplicationSsoConfigResponseBodyApplicationSsoConfigSamlSsoConfigAttributeStatements {
	s.AttributeValueExpression = &v
	return s
}

type GetApplicationSsoConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationSsoConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationSsoConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationSsoConfigResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationSsoConfigResponse) SetHeaders(v map[string]*string) *GetApplicationSsoConfigResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationSsoConfigResponse) SetStatusCode(v int32) *GetApplicationSsoConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationSsoConfigResponse) SetBody(v *GetApplicationSsoConfigResponseBody) *GetApplicationSsoConfigResponse {
	s.Body = v
	return s
}

type GetDomainRequest struct {
	// 域名ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// dm_examplexxxxx
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDomainRequest) GoString() string {
	return s.String()
}

func (s *GetDomainRequest) SetDomainId(v string) *GetDomainRequest {
	s.DomainId = &v
	return s
}

func (s *GetDomainRequest) SetInstanceId(v string) *GetDomainRequest {
	s.InstanceId = &v
	return s
}

type GetDomainResponseBody struct {
	Domain *GetDomainResponseBodyDomain `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Struct"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDomainResponseBody) GoString() string {
	return s.String()
}

func (s *GetDomainResponseBody) SetDomain(v *GetDomainResponseBodyDomain) *GetDomainResponseBody {
	s.Domain = v
	return s
}

func (s *GetDomainResponseBody) SetRequestId(v string) *GetDomainResponseBody {
	s.RequestId = &v
	return s
}

type GetDomainResponseBodyDomain struct {
	// 域名创建时间，Unix时间戳格式，单位为毫秒。
	//
	// example:
	//
	// 1649830226000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 是否默认域名。true表示实例默认域名，false表示非默认域名
	//
	// example:
	//
	// false
	DefaultDomain *bool `json:"DefaultDomain,omitempty" xml:"DefaultDomain,omitempty"`
	// 域名。
	//
	// example:
	//
	// login.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// 域名ID。
	//
	// example:
	//
	// dm_examplexxxxx
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// 域名类型。枚举取值:system_init(系统初始化)、user_custom(用户自定义)。
	//
	// example:
	//
	// system_init
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// 域名备案信息。
	Filing *GetDomainResponseBodyDomainFiling `json:"Filing,omitempty" xml:"Filing,omitempty" type:"Struct"`
	// 实例ID。
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 域名锁定状态。枚举取值:unlock(正常)、lockByLicense(因License限制不可用)。
	//
	// example:
	//
	// unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// 域名最近更新时间，Unix时间戳格式，单位为毫秒。
	//
	// example:
	//
	// 1649830226000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetDomainResponseBodyDomain) String() string {
	return tea.Prettify(s)
}

func (s GetDomainResponseBodyDomain) GoString() string {
	return s.String()
}

func (s *GetDomainResponseBodyDomain) SetCreateTime(v int64) *GetDomainResponseBodyDomain {
	s.CreateTime = &v
	return s
}

func (s *GetDomainResponseBodyDomain) SetDefaultDomain(v bool) *GetDomainResponseBodyDomain {
	s.DefaultDomain = &v
	return s
}

func (s *GetDomainResponseBodyDomain) SetDomain(v string) *GetDomainResponseBodyDomain {
	s.Domain = &v
	return s
}

func (s *GetDomainResponseBodyDomain) SetDomainId(v string) *GetDomainResponseBodyDomain {
	s.DomainId = &v
	return s
}

func (s *GetDomainResponseBodyDomain) SetDomainType(v string) *GetDomainResponseBodyDomain {
	s.DomainType = &v
	return s
}

func (s *GetDomainResponseBodyDomain) SetFiling(v *GetDomainResponseBodyDomainFiling) *GetDomainResponseBodyDomain {
	s.Filing = v
	return s
}

func (s *GetDomainResponseBodyDomain) SetInstanceId(v string) *GetDomainResponseBodyDomain {
	s.InstanceId = &v
	return s
}

func (s *GetDomainResponseBodyDomain) SetLockMode(v string) *GetDomainResponseBodyDomain {
	s.LockMode = &v
	return s
}

func (s *GetDomainResponseBodyDomain) SetUpdateTime(v int64) *GetDomainResponseBodyDomain {
	s.UpdateTime = &v
	return s
}

type GetDomainResponseBodyDomainFiling struct {
	// 域名关联的备案号, 长度最大限制64。
	IcpNumber *string `json:"IcpNumber,omitempty" xml:"IcpNumber,omitempty"`
}

func (s GetDomainResponseBodyDomainFiling) String() string {
	return tea.Prettify(s)
}

func (s GetDomainResponseBodyDomainFiling) GoString() string {
	return s.String()
}

func (s *GetDomainResponseBodyDomainFiling) SetIcpNumber(v string) *GetDomainResponseBodyDomainFiling {
	s.IcpNumber = &v
	return s
}

type GetDomainResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDomainResponse) GoString() string {
	return s.String()
}

func (s *GetDomainResponse) SetHeaders(v map[string]*string) *GetDomainResponse {
	s.Headers = v
	return s
}

func (s *GetDomainResponse) SetStatusCode(v int32) *GetDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDomainResponse) SetBody(v *GetDomainResponseBody) *GetDomainResponse {
	s.Body = v
	return s
}

type GetDomainDnsChallengeRequest struct {
	// 域名。
	//
	// This parameter is required.
	//
	// example:
	//
	// dm_examplexxxx
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetDomainDnsChallengeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDomainDnsChallengeRequest) GoString() string {
	return s.String()
}

func (s *GetDomainDnsChallengeRequest) SetDomain(v string) *GetDomainDnsChallengeRequest {
	s.Domain = &v
	return s
}

func (s *GetDomainDnsChallengeRequest) SetInstanceId(v string) *GetDomainDnsChallengeRequest {
	s.InstanceId = &v
	return s
}

type GetDomainDnsChallengeResponseBody struct {
	DomainDnsChallenge *GetDomainDnsChallengeResponseBodyDomainDnsChallenge `json:"DomainDnsChallenge,omitempty" xml:"DomainDnsChallenge,omitempty" type:"Struct"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDomainDnsChallengeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDomainDnsChallengeResponseBody) GoString() string {
	return s.String()
}

func (s *GetDomainDnsChallengeResponseBody) SetDomainDnsChallenge(v *GetDomainDnsChallengeResponseBodyDomainDnsChallenge) *GetDomainDnsChallengeResponseBody {
	s.DomainDnsChallenge = v
	return s
}

func (s *GetDomainDnsChallengeResponseBody) SetRequestId(v string) *GetDomainDnsChallengeResponseBody {
	s.RequestId = &v
	return s
}

type GetDomainDnsChallengeResponseBodyDomainDnsChallenge struct {
	// DNS challenge名称。
	//
	// example:
	//
	// _idaas-challenge.${domain}
	DnsChallengeName *string `json:"DnsChallengeName,omitempty" xml:"DnsChallengeName,omitempty"`
	// DNS challenge值。
	//
	// example:
	//
	// exmple123xxx
	DnsChallengeValue *string `json:"DnsChallengeValue,omitempty" xml:"DnsChallengeValue,omitempty"`
	// DNS记录类型。
	//
	// example:
	//
	// 枚举，目前只支持TXT类型
	DnsType *string `json:"DnsType,omitempty" xml:"DnsType,omitempty"`
}

func (s GetDomainDnsChallengeResponseBodyDomainDnsChallenge) String() string {
	return tea.Prettify(s)
}

func (s GetDomainDnsChallengeResponseBodyDomainDnsChallenge) GoString() string {
	return s.String()
}

func (s *GetDomainDnsChallengeResponseBodyDomainDnsChallenge) SetDnsChallengeName(v string) *GetDomainDnsChallengeResponseBodyDomainDnsChallenge {
	s.DnsChallengeName = &v
	return s
}

func (s *GetDomainDnsChallengeResponseBodyDomainDnsChallenge) SetDnsChallengeValue(v string) *GetDomainDnsChallengeResponseBodyDomainDnsChallenge {
	s.DnsChallengeValue = &v
	return s
}

func (s *GetDomainDnsChallengeResponseBodyDomainDnsChallenge) SetDnsType(v string) *GetDomainDnsChallengeResponseBodyDomainDnsChallenge {
	s.DnsType = &v
	return s
}

type GetDomainDnsChallengeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDomainDnsChallengeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDomainDnsChallengeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDomainDnsChallengeResponse) GoString() string {
	return s.String()
}

func (s *GetDomainDnsChallengeResponse) SetHeaders(v map[string]*string) *GetDomainDnsChallengeResponse {
	s.Headers = v
	return s
}

func (s *GetDomainDnsChallengeResponse) SetStatusCode(v int32) *GetDomainDnsChallengeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDomainDnsChallengeResponse) SetBody(v *GetDomainDnsChallengeResponseBody) *GetDomainDnsChallengeResponse {
	s.Body = v
	return s
}

type GetForgetPasswordConfigurationRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetForgetPasswordConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetForgetPasswordConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetForgetPasswordConfigurationRequest) SetInstanceId(v string) *GetForgetPasswordConfigurationRequest {
	s.InstanceId = &v
	return s
}

type GetForgetPasswordConfigurationResponseBody struct {
	// The forgot password configurations.
	OpenForgetPasswordConfiguration *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration `json:"OpenForgetPasswordConfiguration,omitempty" xml:"OpenForgetPasswordConfiguration,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetForgetPasswordConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetForgetPasswordConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetForgetPasswordConfigurationResponseBody) SetOpenForgetPasswordConfiguration(v *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration) *GetForgetPasswordConfigurationResponseBody {
	s.OpenForgetPasswordConfiguration = v
	return s
}

func (s *GetForgetPasswordConfigurationResponseBody) SetRequestId(v string) *GetForgetPasswordConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration struct {
	// 表示忘记密码认证渠道。枚举取值:email(邮件)、sms(短信)
	AuthenticationChannels []*string `json:"AuthenticationChannels,omitempty" xml:"AuthenticationChannels,omitempty" type:"Repeated"`
	// Indicates whether the forgot password feature is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Indicates whether email authentication is enabled for the forgot password feature.
	//
	// example:
	//
	// true
	EnableEmail *bool `json:"EnableEmail,omitempty" xml:"EnableEmail,omitempty"`
	// Indicates whether Short Message Service (SMS) authentication is enabled for the forgot password feature.
	//
	// example:
	//
	// true
	EnableSms *bool `json:"EnableSms,omitempty" xml:"EnableSms,omitempty"`
	// 表示忘记密码配置状态。枚举取值:enabled(开启)、disabled(禁用)
	ForgetPasswordStatus *string `json:"ForgetPasswordStatus,omitempty" xml:"ForgetPasswordStatus,omitempty"`
}

func (s GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration) String() string {
	return tea.Prettify(s)
}

func (s GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration) GoString() string {
	return s.String()
}

func (s *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration) SetAuthenticationChannels(v []*string) *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration {
	s.AuthenticationChannels = v
	return s
}

func (s *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration) SetEnable(v bool) *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration {
	s.Enable = &v
	return s
}

func (s *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration) SetEnableEmail(v bool) *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration {
	s.EnableEmail = &v
	return s
}

func (s *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration) SetEnableSms(v bool) *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration {
	s.EnableSms = &v
	return s
}

func (s *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration) SetForgetPasswordStatus(v string) *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration {
	s.ForgetPasswordStatus = &v
	return s
}

type GetForgetPasswordConfigurationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetForgetPasswordConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetForgetPasswordConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetForgetPasswordConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetForgetPasswordConfigurationResponse) SetHeaders(v map[string]*string) *GetForgetPasswordConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetForgetPasswordConfigurationResponse) SetStatusCode(v int32) *GetForgetPasswordConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetForgetPasswordConfigurationResponse) SetBody(v *GetForgetPasswordConfigurationResponseBody) *GetForgetPasswordConfigurationResponse {
	s.Body = v
	return s
}

type GetGroupRequest struct {
	// The group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGroupRequest) GoString() string {
	return s.String()
}

func (s *GetGroupRequest) SetGroupId(v string) *GetGroupRequest {
	s.GroupId = &v
	return s
}

func (s *GetGroupRequest) SetInstanceId(v string) *GetGroupRequest {
	s.InstanceId = &v
	return s
}

type GetGroupResponseBody struct {
	// The information about the account group.
	Group *GetGroupResponseBodyGroup `json:"Group,omitempty" xml:"Group,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupResponseBody) SetGroup(v *GetGroupResponseBodyGroup) *GetGroupResponseBody {
	s.Group = v
	return s
}

func (s *GetGroupResponseBody) SetRequestId(v string) *GetGroupResponseBody {
	s.RequestId = &v
	return s
}

type GetGroupResponseBodyGroup struct {
	// The time at which the group was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1652085686179
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the group.
	//
	// example:
	//
	// test_group
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The external ID of the group, which can be used to associate the group with an external system. By default, the external ID is the group ID.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupExternalId *string `json:"GroupExternalId,omitempty" xml:"GroupExternalId,omitempty"`
	// The group ID.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the group.
	//
	// example:
	//
	// group_name
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The source ID of the group. By default, the source ID is the instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	GroupSourceId *string `json:"GroupSourceId,omitempty" xml:"GroupSourceId,omitempty"`
	// The source type of the group. Only build_in may be returned, which indicates that the group was created in IDaaS.
	//
	// *build_in:Create By Self.
	//
	// example:
	//
	// build_in
	GroupSourceType *string `json:"GroupSourceType,omitempty" xml:"GroupSourceType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time at which the group was last updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1652085686179
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetGroupResponseBodyGroup) String() string {
	return tea.Prettify(s)
}

func (s GetGroupResponseBodyGroup) GoString() string {
	return s.String()
}

func (s *GetGroupResponseBodyGroup) SetCreateTime(v int64) *GetGroupResponseBodyGroup {
	s.CreateTime = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetDescription(v string) *GetGroupResponseBodyGroup {
	s.Description = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetGroupExternalId(v string) *GetGroupResponseBodyGroup {
	s.GroupExternalId = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetGroupId(v string) *GetGroupResponseBodyGroup {
	s.GroupId = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetGroupName(v string) *GetGroupResponseBodyGroup {
	s.GroupName = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetGroupSourceId(v string) *GetGroupResponseBodyGroup {
	s.GroupSourceId = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetGroupSourceType(v string) *GetGroupResponseBodyGroup {
	s.GroupSourceType = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetInstanceId(v string) *GetGroupResponseBodyGroup {
	s.InstanceId = &v
	return s
}

func (s *GetGroupResponseBodyGroup) SetUpdateTime(v int64) *GetGroupResponseBodyGroup {
	s.UpdateTime = &v
	return s
}

type GetGroupResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGroupResponse) GoString() string {
	return s.String()
}

func (s *GetGroupResponse) SetHeaders(v map[string]*string) *GetGroupResponse {
	s.Headers = v
	return s
}

func (s *GetGroupResponse) SetStatusCode(v int32) *GetGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGroupResponse) SetBody(v *GetGroupResponseBody) *GetGroupResponse {
	s.Body = v
	return s
}

type GetInstanceRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceRequest) SetInstanceId(v string) *GetInstanceRequest {
	s.InstanceId = &v
	return s
}

type GetInstanceResponseBody struct {
	// The details of the instance.
	Instance *GetInstanceResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) SetInstance(v *GetInstanceResponseBodyInstance) *GetInstanceResponseBody {
	s.Instance = v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

type GetInstanceResponseBodyInstance struct {
	// The time when the instance was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1550115455000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The default endpoint of the instance.
	DefaultEndpoint *GetInstanceResponseBodyInstanceDefaultEndpoint `json:"DefaultEndpoint,omitempty" xml:"DefaultEndpoint,omitempty" type:"Struct"`
	// The description of the instance.
	//
	// example:
	//
	// test_description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The default domain of the instance.
	DomainConfig *GetInstanceResponseBodyInstanceDomainConfig `json:"DomainConfig,omitempty" xml:"DomainConfig,omitempty" type:"Struct"`
	// The outbound public CIDR blocks of the instance. For example, when you synchronize Active Directory (AD) accounts, the IDaaS EIAM instance accesses your AD service by using the outbound public CIDR blocks.
	EgressAddresses []*string `json:"EgressAddresses,omitempty" xml:"EgressAddresses,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// example:
	//
	// idaas_abt3pfwojojcq323si6g5xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- creating
	//
	// 	- running
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetInstanceResponseBodyInstance) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstance) SetCreateTime(v int64) *GetInstanceResponseBodyInstance {
	s.CreateTime = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetDefaultEndpoint(v *GetInstanceResponseBodyInstanceDefaultEndpoint) *GetInstanceResponseBodyInstance {
	s.DefaultEndpoint = v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetDescription(v string) *GetInstanceResponseBodyInstance {
	s.Description = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetDomainConfig(v *GetInstanceResponseBodyInstanceDomainConfig) *GetInstanceResponseBodyInstance {
	s.DomainConfig = v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetEgressAddresses(v []*string) *GetInstanceResponseBodyInstance {
	s.EgressAddresses = v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetInstanceId(v string) *GetInstanceResponseBodyInstance {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetStatus(v string) *GetInstanceResponseBodyInstance {
	s.Status = &v
	return s
}

type GetInstanceResponseBodyInstanceDefaultEndpoint struct {
	// The endpoint of the instance.
	//
	// example:
	//
	// example-xxx.aliyunidaas.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The status of the endpoint. Valid values:
	//
	// 	- resolved
	//
	// 	- unresolved
	//
	// example:
	//
	// resolved
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetInstanceResponseBodyInstanceDefaultEndpoint) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyInstanceDefaultEndpoint) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstanceDefaultEndpoint) SetEndpoint(v string) *GetInstanceResponseBodyInstanceDefaultEndpoint {
	s.Endpoint = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceDefaultEndpoint) SetStatus(v string) *GetInstanceResponseBodyInstanceDefaultEndpoint {
	s.Status = &v
	return s
}

type GetInstanceResponseBodyInstanceDomainConfig struct {
	// The default domain of the instance.
	//
	// example:
	//
	// example-xxx.example.com
	DefaultDomain *string `json:"DefaultDomain,omitempty" xml:"DefaultDomain,omitempty"`
	// The init domain of the instance.
	//
	// example:
	//
	// example-xxx.aliyunidaas.com
	InitDomain *string `json:"InitDomain,omitempty" xml:"InitDomain,omitempty"`
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	InitDomainAutoRedirectStatus *string `json:"InitDomainAutoRedirectStatus,omitempty" xml:"InitDomainAutoRedirectStatus,omitempty"`
}

func (s GetInstanceResponseBodyInstanceDomainConfig) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyInstanceDomainConfig) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstanceDomainConfig) SetDefaultDomain(v string) *GetInstanceResponseBodyInstanceDomainConfig {
	s.DefaultDomain = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceDomainConfig) SetInitDomain(v string) *GetInstanceResponseBodyInstanceDomainConfig {
	s.InitDomain = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceDomainConfig) SetInitDomainAutoRedirectStatus(v string) *GetInstanceResponseBodyInstanceDomainConfig {
	s.InitDomainAutoRedirectStatus = &v
	return s
}

type GetInstanceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceResponse) SetHeaders(v map[string]*string) *GetInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceResponse) SetStatusCode(v int32) *GetInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceResponse) SetBody(v *GetInstanceResponseBody) *GetInstanceResponse {
	s.Body = v
	return s
}

type GetNetworkAccessEndpointRequest struct {
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 专属网络端点ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// nae-examplexxxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
}

func (s GetNetworkAccessEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkAccessEndpointRequest) GoString() string {
	return s.String()
}

func (s *GetNetworkAccessEndpointRequest) SetInstanceId(v string) *GetNetworkAccessEndpointRequest {
	s.InstanceId = &v
	return s
}

func (s *GetNetworkAccessEndpointRequest) SetNetworkAccessEndpointId(v string) *GetNetworkAccessEndpointRequest {
	s.NetworkAccessEndpointId = &v
	return s
}

type GetNetworkAccessEndpointResponseBody struct {
	NetworkAccessEndpoint *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint `json:"NetworkAccessEndpoint,omitempty" xml:"NetworkAccessEndpoint,omitempty" type:"Struct"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNetworkAccessEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkAccessEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *GetNetworkAccessEndpointResponseBody) SetNetworkAccessEndpoint(v *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) *GetNetworkAccessEndpointResponseBody {
	s.NetworkAccessEndpoint = v
	return s
}

func (s *GetNetworkAccessEndpointResponseBody) SetRequestId(v string) *GetNetworkAccessEndpointResponseBody {
	s.RequestId = &v
	return s
}

type GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint struct {
	// 专属网络端点创建时间，Unix时间戳格式，单位为毫秒。
	//
	// example:
	//
	// 1649830226000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 网络访问端私网出口IP地址列表。
	//
	// example:
	//
	// 172.168.x.x
	EgressPrivateIpAddresses []*string `json:"EgressPrivateIpAddresses,omitempty" xml:"EgressPrivateIpAddresses,omitempty" type:"Repeated"`
	// 网络访问端点公网出口IP地址段
	//
	// example:
	//
	// 8.xx.xx.xxx/27
	EgressPublicIpAddresses []*string `json:"EgressPublicIpAddresses,omitempty" xml:"EgressPublicIpAddresses,omitempty" type:"Repeated"`
	// 实例ID。
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 专属网络端点ID。
	//
	// example:
	//
	// nae_examplexxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
	// 专属网络端点名称。
	//
	// example:
	//
	// xx业务VPC访问端点
	NetworkAccessEndpointName *string `json:"NetworkAccessEndpointName,omitempty" xml:"NetworkAccessEndpointName,omitempty"`
	// 专属网络端点连接的类型。
	//
	// example:
	//
	// private
	NetworkAccessEndpointType *string `json:"NetworkAccessEndpointType,omitempty" xml:"NetworkAccessEndpointType,omitempty"`
	// 专属网络端点使用的安全组ID。
	//
	// example:
	//
	// sg-examplexxx
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// 专属网络端点状态。
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 专属网络端点最近更新时间，Unix时间戳格式，单位为毫秒。
	//
	// example:
	//
	// 1649830226000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 专属网络端点连接的指定vSwitch列表。
	//
	// example:
	//
	// vsw-examplexxx
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// 专属网络端点连接的VpcID。
	//
	// example:
	//
	// vpc-examplexxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// 专属网络端点连接的Vpc所属地域。
	//
	// example:
	//
	// cn-hangzhou
	VpcRegionId *string `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
}

func (s GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) GoString() string {
	return s.String()
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) SetCreateTime(v int64) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	s.CreateTime = &v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) SetEgressPrivateIpAddresses(v []*string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	s.EgressPrivateIpAddresses = v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) SetEgressPublicIpAddresses(v []*string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	s.EgressPublicIpAddresses = v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) SetInstanceId(v string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	s.InstanceId = &v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) SetNetworkAccessEndpointId(v string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	s.NetworkAccessEndpointId = &v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) SetNetworkAccessEndpointName(v string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	s.NetworkAccessEndpointName = &v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) SetNetworkAccessEndpointType(v string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	s.NetworkAccessEndpointType = &v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) SetSecurityGroupId(v string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	s.SecurityGroupId = &v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) SetStatus(v string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	s.Status = &v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) SetUpdateTime(v int64) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	s.UpdateTime = &v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) SetVSwitchIds(v []*string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	s.VSwitchIds = v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) SetVpcId(v string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	s.VpcId = &v
	return s
}

func (s *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint) SetVpcRegionId(v string) *GetNetworkAccessEndpointResponseBodyNetworkAccessEndpoint {
	s.VpcRegionId = &v
	return s
}

type GetNetworkAccessEndpointResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNetworkAccessEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNetworkAccessEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNetworkAccessEndpointResponse) GoString() string {
	return s.String()
}

func (s *GetNetworkAccessEndpointResponse) SetHeaders(v map[string]*string) *GetNetworkAccessEndpointResponse {
	s.Headers = v
	return s
}

func (s *GetNetworkAccessEndpointResponse) SetStatusCode(v int32) *GetNetworkAccessEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNetworkAccessEndpointResponse) SetBody(v *GetNetworkAccessEndpointResponseBody) *GetNetworkAccessEndpointResponse {
	s.Body = v
	return s
}

type GetOrganizationalUnitRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the organizational unit.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
}

func (s GetOrganizationalUnitRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitRequest) SetInstanceId(v string) *GetOrganizationalUnitRequest {
	s.InstanceId = &v
	return s
}

func (s *GetOrganizationalUnitRequest) SetOrganizationalUnitId(v string) *GetOrganizationalUnitRequest {
	s.OrganizationalUnitId = &v
	return s
}

type GetOrganizationalUnitResponseBody struct {
	// The data object of the organizational unit.
	OrganizationalUnit *GetOrganizationalUnitResponseBodyOrganizationalUnit `json:"OrganizationalUnit,omitempty" xml:"OrganizationalUnit,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOrganizationalUnitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationalUnitResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitResponseBody) SetOrganizationalUnit(v *GetOrganizationalUnitResponseBodyOrganizationalUnit) *GetOrganizationalUnitResponseBody {
	s.OrganizationalUnit = v
	return s
}

func (s *GetOrganizationalUnitResponseBody) SetRequestId(v string) *GetOrganizationalUnitResponseBody {
	s.RequestId = &v
	return s
}

type GetOrganizationalUnitResponseBodyOrganizationalUnit struct {
	// The time when the organizational unit was created. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the organizational unit.
	//
	// example:
	//
	// Test organizational unit
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the node is a leaf node.
	//
	// example:
	//
	// false
	Leaf *bool `json:"Leaf,omitempty" xml:"Leaf,omitempty"`
	// The external ID of the organizational unit. The external ID can be used by external data to map the data of the organizational unit in IDaaS EIAM. By default, the external ID is the organizational unit ID.
	//
	// For organizational units with the same source type and source ID, each organizational unit has a unique external ID.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitExternalId *string `json:"OrganizationalUnitExternalId,omitempty" xml:"OrganizationalUnitExternalId,omitempty"`
	// The ID of the organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
	// 组织名称。
	//
	// example:
	//
	// test_organizationalUnit_name
	OrganizationalUnitName *string `json:"OrganizationalUnitName,omitempty" xml:"OrganizationalUnitName,omitempty"`
	// The source ID of the organizational unit.
	//
	// If the organizational unit was created in IDaaS, its source ID is the ID of the IDaaS instance. If the organizational unit was imported, its source ID is the enterprise ID in the source. For example, if the organizational unit was imported from DingTalk, its source ID is the corpId value of the enterprise in DingTalk.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	OrganizationalUnitSourceId *string `json:"OrganizationalUnitSourceId,omitempty" xml:"OrganizationalUnitSourceId,omitempty"`
	// The source type of the organizational unit. Valid values:
	//
	// 	- build_in: The organizational unit was created in IDaaS.
	//
	// 	- ding_talk: The organizational unit was imported from DingTalk.
	//
	// 	- ad: The organizational unit was imported from Microsoft Active Directory (AD).
	//
	// 	- ldap: The organizational unit was imported from a Lightweight Directory Access Protocol (LDAP) service.
	//
	// example:
	//
	// build_in
	OrganizationalUnitSourceType *string `json:"OrganizationalUnitSourceType,omitempty" xml:"OrganizationalUnitSourceType,omitempty"`
	// The ID of the parent organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The time when the organizational unit was last updated. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetOrganizationalUnitResponseBodyOrganizationalUnit) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationalUnitResponseBodyOrganizationalUnit) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) SetCreateTime(v int64) *GetOrganizationalUnitResponseBodyOrganizationalUnit {
	s.CreateTime = &v
	return s
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) SetDescription(v string) *GetOrganizationalUnitResponseBodyOrganizationalUnit {
	s.Description = &v
	return s
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) SetInstanceId(v string) *GetOrganizationalUnitResponseBodyOrganizationalUnit {
	s.InstanceId = &v
	return s
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) SetLeaf(v bool) *GetOrganizationalUnitResponseBodyOrganizationalUnit {
	s.Leaf = &v
	return s
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) SetOrganizationalUnitExternalId(v string) *GetOrganizationalUnitResponseBodyOrganizationalUnit {
	s.OrganizationalUnitExternalId = &v
	return s
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) SetOrganizationalUnitId(v string) *GetOrganizationalUnitResponseBodyOrganizationalUnit {
	s.OrganizationalUnitId = &v
	return s
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) SetOrganizationalUnitName(v string) *GetOrganizationalUnitResponseBodyOrganizationalUnit {
	s.OrganizationalUnitName = &v
	return s
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) SetOrganizationalUnitSourceId(v string) *GetOrganizationalUnitResponseBodyOrganizationalUnit {
	s.OrganizationalUnitSourceId = &v
	return s
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) SetOrganizationalUnitSourceType(v string) *GetOrganizationalUnitResponseBodyOrganizationalUnit {
	s.OrganizationalUnitSourceType = &v
	return s
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) SetParentId(v string) *GetOrganizationalUnitResponseBodyOrganizationalUnit {
	s.ParentId = &v
	return s
}

func (s *GetOrganizationalUnitResponseBodyOrganizationalUnit) SetUpdateTime(v int64) *GetOrganizationalUnitResponseBodyOrganizationalUnit {
	s.UpdateTime = &v
	return s
}

type GetOrganizationalUnitResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOrganizationalUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrganizationalUnitResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *GetOrganizationalUnitResponse) SetHeaders(v map[string]*string) *GetOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *GetOrganizationalUnitResponse) SetStatusCode(v int32) *GetOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrganizationalUnitResponse) SetBody(v *GetOrganizationalUnitResponseBody) *GetOrganizationalUnitResponse {
	s.Body = v
	return s
}

type GetPasswordComplexityConfigurationRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetPasswordComplexityConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordComplexityConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetPasswordComplexityConfigurationRequest) SetInstanceId(v string) *GetPasswordComplexityConfigurationRequest {
	s.InstanceId = &v
	return s
}

type GetPasswordComplexityConfigurationResponseBody struct {
	// The password complexity configurations.
	PasswordComplexityConfiguration *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration `json:"PasswordComplexityConfiguration,omitempty" xml:"PasswordComplexityConfiguration,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPasswordComplexityConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordComplexityConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetPasswordComplexityConfigurationResponseBody) SetPasswordComplexityConfiguration(v *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration) *GetPasswordComplexityConfigurationResponseBody {
	s.PasswordComplexityConfiguration = v
	return s
}

func (s *GetPasswordComplexityConfigurationResponseBody) SetRequestId(v string) *GetPasswordComplexityConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration struct {
	// The password complexity rules.
	PasswordComplexityRules []*GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfigurationPasswordComplexityRules `json:"PasswordComplexityRules,omitempty" xml:"PasswordComplexityRules,omitempty" type:"Repeated"`
	// The minimum number of characters in a password.
	//
	// example:
	//
	// 3
	PasswordMinLength *int32 `json:"PasswordMinLength,omitempty" xml:"PasswordMinLength,omitempty"`
}

func (s GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration) GoString() string {
	return s.String()
}

func (s *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration) SetPasswordComplexityRules(v []*GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfigurationPasswordComplexityRules) *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration {
	s.PasswordComplexityRules = v
	return s
}

func (s *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration) SetPasswordMinLength(v int32) *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration {
	s.PasswordMinLength = &v
	return s
}

type GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfigurationPasswordComplexityRules struct {
	// The type of the password check. Valid values:
	//
	// 	- inclusion_upper_case: The password must contain uppercase letters.
	//
	// 	- inclusion_lower_case: The password must contain lowercase letters.
	//
	// 	- inclusion_special_case: The password must contain one or more of the following special characters: @ % + \\ / \\" ! # $ ^ ? : , ( ) { } [ ] ~ - _ .
	//
	// 	- inclusion_number: The password must contain digits.
	//
	// 	- exclusion_username: The password cannot contain a username.
	//
	// 	- exclusion_email: The password cannot contain an email prefix.
	//
	// 	- exclusion_phone_number: The password cannot contain a mobile number.
	//
	// 	- exclusion_display_name: The password cannot contain a display name.
	//
	// example:
	//
	// inclusion_upper_case
	PasswordCheckType *string `json:"PasswordCheckType,omitempty" xml:"PasswordCheckType,omitempty"`
}

func (s GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfigurationPasswordComplexityRules) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfigurationPasswordComplexityRules) GoString() string {
	return s.String()
}

func (s *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfigurationPasswordComplexityRules) SetPasswordCheckType(v string) *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfigurationPasswordComplexityRules {
	s.PasswordCheckType = &v
	return s
}

type GetPasswordComplexityConfigurationResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPasswordComplexityConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPasswordComplexityConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordComplexityConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetPasswordComplexityConfigurationResponse) SetHeaders(v map[string]*string) *GetPasswordComplexityConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetPasswordComplexityConfigurationResponse) SetStatusCode(v int32) *GetPasswordComplexityConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPasswordComplexityConfigurationResponse) SetBody(v *GetPasswordComplexityConfigurationResponseBody) *GetPasswordComplexityConfigurationResponse {
	s.Body = v
	return s
}

type GetPasswordExpirationConfigurationRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetPasswordExpirationConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordExpirationConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetPasswordExpirationConfigurationRequest) SetInstanceId(v string) *GetPasswordExpirationConfigurationRequest {
	s.InstanceId = &v
	return s
}

type GetPasswordExpirationConfigurationResponseBody struct {
	// The password expiration configurations.
	PasswordExpirationConfiguration *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration `json:"PasswordExpirationConfiguration,omitempty" xml:"PasswordExpirationConfiguration,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPasswordExpirationConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordExpirationConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetPasswordExpirationConfigurationResponseBody) SetPasswordExpirationConfiguration(v *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) *GetPasswordExpirationConfigurationResponseBody {
	s.PasswordExpirationConfiguration = v
	return s
}

func (s *GetPasswordExpirationConfigurationResponseBody) SetRequestId(v string) *GetPasswordExpirationConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration struct {
	// The action to take when a password expires. Valid values:
	//
	// 	- forbid_login: Prohibit the user from using the password to log on to IDaaS.
	//
	// 	- force_update_password: Force the user to change the password.
	//
	// 	- remind_update_password: Remind the user to change the password.
	//
	// example:
	//
	// forbid_login
	PasswordExpirationAction *string `json:"PasswordExpirationAction,omitempty" xml:"PasswordExpirationAction,omitempty"`
	// The methods for receiving password expiration notifications.
	//
	// example:
	//
	// login
	PasswordExpirationNotificationChannels []*string `json:"PasswordExpirationNotificationChannels,omitempty" xml:"PasswordExpirationNotificationChannels,omitempty" type:"Repeated"`
	// The number of days before the expiration date during which password expiration notifications are sent. Unit: day.
	//
	// example:
	//
	// 7
	PasswordExpirationNotificationDuration *int32 `json:"PasswordExpirationNotificationDuration,omitempty" xml:"PasswordExpirationNotificationDuration,omitempty"`
	// Indicates whether the password expiration notification feature is enabled. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// example:
	//
	// enabled
	PasswordExpirationNotificationStatus *string `json:"PasswordExpirationNotificationStatus,omitempty" xml:"PasswordExpirationNotificationStatus,omitempty"`
	// Indicates whether the password expiration feature is enabled. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// example:
	//
	// enabled
	PasswordExpirationStatus *string `json:"PasswordExpirationStatus,omitempty" xml:"PasswordExpirationStatus,omitempty"`
	// The number of days before which users must change the password to prevent password expiration. Unit: day.
	//
	// example:
	//
	// 3
	PasswordForcedUpdateDuration *int32 `json:"PasswordForcedUpdateDuration,omitempty" xml:"PasswordForcedUpdateDuration,omitempty"`
	// The validity period of a password. Unit: day.
	//
	// example:
	//
	// 180
	PasswordValidMaxDay *int32 `json:"PasswordValidMaxDay,omitempty" xml:"PasswordValidMaxDay,omitempty"`
}

func (s GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) GoString() string {
	return s.String()
}

func (s *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) SetPasswordExpirationAction(v string) *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration {
	s.PasswordExpirationAction = &v
	return s
}

func (s *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) SetPasswordExpirationNotificationChannels(v []*string) *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration {
	s.PasswordExpirationNotificationChannels = v
	return s
}

func (s *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) SetPasswordExpirationNotificationDuration(v int32) *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration {
	s.PasswordExpirationNotificationDuration = &v
	return s
}

func (s *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) SetPasswordExpirationNotificationStatus(v string) *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration {
	s.PasswordExpirationNotificationStatus = &v
	return s
}

func (s *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) SetPasswordExpirationStatus(v string) *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration {
	s.PasswordExpirationStatus = &v
	return s
}

func (s *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) SetPasswordForcedUpdateDuration(v int32) *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration {
	s.PasswordForcedUpdateDuration = &v
	return s
}

func (s *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) SetPasswordValidMaxDay(v int32) *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration {
	s.PasswordValidMaxDay = &v
	return s
}

type GetPasswordExpirationConfigurationResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPasswordExpirationConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPasswordExpirationConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordExpirationConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetPasswordExpirationConfigurationResponse) SetHeaders(v map[string]*string) *GetPasswordExpirationConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetPasswordExpirationConfigurationResponse) SetStatusCode(v int32) *GetPasswordExpirationConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPasswordExpirationConfigurationResponse) SetBody(v *GetPasswordExpirationConfigurationResponseBody) *GetPasswordExpirationConfigurationResponse {
	s.Body = v
	return s
}

type GetPasswordHistoryConfigurationRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetPasswordHistoryConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordHistoryConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetPasswordHistoryConfigurationRequest) SetInstanceId(v string) *GetPasswordHistoryConfigurationRequest {
	s.InstanceId = &v
	return s
}

type GetPasswordHistoryConfigurationResponseBody struct {
	// The password history configurations.
	PasswordHistoryConfiguration *GetPasswordHistoryConfigurationResponseBodyPasswordHistoryConfiguration `json:"PasswordHistoryConfiguration,omitempty" xml:"PasswordHistoryConfiguration,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPasswordHistoryConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordHistoryConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetPasswordHistoryConfigurationResponseBody) SetPasswordHistoryConfiguration(v *GetPasswordHistoryConfigurationResponseBodyPasswordHistoryConfiguration) *GetPasswordHistoryConfigurationResponseBody {
	s.PasswordHistoryConfiguration = v
	return s
}

func (s *GetPasswordHistoryConfigurationResponseBody) SetRequestId(v string) *GetPasswordHistoryConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type GetPasswordHistoryConfigurationResponseBodyPasswordHistoryConfiguration struct {
	// The maximum number of recent passwords that are retained.
	//
	// example:
	//
	// 3
	PasswordHistoryMaxRetention *int32 `json:"PasswordHistoryMaxRetention,omitempty" xml:"PasswordHistoryMaxRetention,omitempty"`
	// Indicates whether the password history feature is enabled. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// example:
	//
	// enabled
	PasswordHistoryStatus *string `json:"PasswordHistoryStatus,omitempty" xml:"PasswordHistoryStatus,omitempty"`
}

func (s GetPasswordHistoryConfigurationResponseBodyPasswordHistoryConfiguration) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordHistoryConfigurationResponseBodyPasswordHistoryConfiguration) GoString() string {
	return s.String()
}

func (s *GetPasswordHistoryConfigurationResponseBodyPasswordHistoryConfiguration) SetPasswordHistoryMaxRetention(v int32) *GetPasswordHistoryConfigurationResponseBodyPasswordHistoryConfiguration {
	s.PasswordHistoryMaxRetention = &v
	return s
}

func (s *GetPasswordHistoryConfigurationResponseBodyPasswordHistoryConfiguration) SetPasswordHistoryStatus(v string) *GetPasswordHistoryConfigurationResponseBodyPasswordHistoryConfiguration {
	s.PasswordHistoryStatus = &v
	return s
}

type GetPasswordHistoryConfigurationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPasswordHistoryConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPasswordHistoryConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordHistoryConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetPasswordHistoryConfigurationResponse) SetHeaders(v map[string]*string) *GetPasswordHistoryConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetPasswordHistoryConfigurationResponse) SetStatusCode(v int32) *GetPasswordHistoryConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPasswordHistoryConfigurationResponse) SetBody(v *GetPasswordHistoryConfigurationResponseBody) *GetPasswordHistoryConfigurationResponse {
	s.Body = v
	return s
}

type GetPasswordInitializationConfigurationRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetPasswordInitializationConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordInitializationConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetPasswordInitializationConfigurationRequest) SetInstanceId(v string) *GetPasswordInitializationConfigurationRequest {
	s.InstanceId = &v
	return s
}

type GetPasswordInitializationConfigurationResponseBody struct {
	// The password initialization configurations.
	PasswordInitializationConfiguration *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration `json:"PasswordInitializationConfiguration,omitempty" xml:"PasswordInitializationConfiguration,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPasswordInitializationConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordInitializationConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetPasswordInitializationConfigurationResponseBody) SetPasswordInitializationConfiguration(v *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration) *GetPasswordInitializationConfigurationResponseBody {
	s.PasswordInitializationConfiguration = v
	return s
}

func (s *GetPasswordInitializationConfigurationResponseBody) SetRequestId(v string) *GetPasswordInitializationConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration struct {
	// Indicates whether forcible password change upon first logon is enabled. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// example:
	//
	// enabled
	PasswordForcedUpdateStatus *string `json:"PasswordForcedUpdateStatus,omitempty" xml:"PasswordForcedUpdateStatus,omitempty"`
	// The methods for receiving password initialization notifications.
	//
	// example:
	//
	// email
	PasswordInitializationNotificationChannels []*string `json:"PasswordInitializationNotificationChannels,omitempty" xml:"PasswordInitializationNotificationChannels,omitempty" type:"Repeated"`
	// Indicates whether the password initialization feature is enabled. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// example:
	//
	// enabled
	PasswordInitializationStatus *string `json:"PasswordInitializationStatus,omitempty" xml:"PasswordInitializationStatus,omitempty"`
	// The password initialization method. Set the value to random.
	//
	// 	- random: A randomly generated password is used.
	//
	// example:
	//
	// random
	PasswordInitializationType *string `json:"PasswordInitializationType,omitempty" xml:"PasswordInitializationType,omitempty"`
}

func (s GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration) GoString() string {
	return s.String()
}

func (s *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration) SetPasswordForcedUpdateStatus(v string) *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration {
	s.PasswordForcedUpdateStatus = &v
	return s
}

func (s *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration) SetPasswordInitializationNotificationChannels(v []*string) *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration {
	s.PasswordInitializationNotificationChannels = v
	return s
}

func (s *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration) SetPasswordInitializationStatus(v string) *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration {
	s.PasswordInitializationStatus = &v
	return s
}

func (s *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration) SetPasswordInitializationType(v string) *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration {
	s.PasswordInitializationType = &v
	return s
}

type GetPasswordInitializationConfigurationResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPasswordInitializationConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPasswordInitializationConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordInitializationConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetPasswordInitializationConfigurationResponse) SetHeaders(v map[string]*string) *GetPasswordInitializationConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetPasswordInitializationConfigurationResponse) SetStatusCode(v int32) *GetPasswordInitializationConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPasswordInitializationConfigurationResponse) SetBody(v *GetPasswordInitializationConfigurationResponseBody) *GetPasswordInitializationConfigurationResponse {
	s.Body = v
	return s
}

type GetRootOrganizationalUnitRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetRootOrganizationalUnitRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRootOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *GetRootOrganizationalUnitRequest) SetInstanceId(v string) *GetRootOrganizationalUnitRequest {
	s.InstanceId = &v
	return s
}

type GetRootOrganizationalUnitResponseBody struct {
	// The data object of the organizational unit.
	OrganizationalUnit *GetRootOrganizationalUnitResponseBodyOrganizationalUnit `json:"OrganizationalUnit,omitempty" xml:"OrganizationalUnit,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRootOrganizationalUnitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRootOrganizationalUnitResponseBody) GoString() string {
	return s.String()
}

func (s *GetRootOrganizationalUnitResponseBody) SetOrganizationalUnit(v *GetRootOrganizationalUnitResponseBodyOrganizationalUnit) *GetRootOrganizationalUnitResponseBody {
	s.OrganizationalUnit = v
	return s
}

func (s *GetRootOrganizationalUnitResponseBody) SetRequestId(v string) *GetRootOrganizationalUnitResponseBody {
	s.RequestId = &v
	return s
}

type GetRootOrganizationalUnitResponseBodyOrganizationalUnit struct {
	// The time when the organizational unit was created. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the organizational unit.
	//
	// example:
	//
	// Test organizational unit
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
	// The name of the organization.
	//
	// example:
	//
	// name001
	OrganizationalUnitName *string `json:"OrganizationalUnitName,omitempty" xml:"OrganizationalUnitName,omitempty"`
	// The time when the organizational unit was last updated. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetRootOrganizationalUnitResponseBodyOrganizationalUnit) String() string {
	return tea.Prettify(s)
}

func (s GetRootOrganizationalUnitResponseBodyOrganizationalUnit) GoString() string {
	return s.String()
}

func (s *GetRootOrganizationalUnitResponseBodyOrganizationalUnit) SetCreateTime(v int64) *GetRootOrganizationalUnitResponseBodyOrganizationalUnit {
	s.CreateTime = &v
	return s
}

func (s *GetRootOrganizationalUnitResponseBodyOrganizationalUnit) SetDescription(v string) *GetRootOrganizationalUnitResponseBodyOrganizationalUnit {
	s.Description = &v
	return s
}

func (s *GetRootOrganizationalUnitResponseBodyOrganizationalUnit) SetInstanceId(v string) *GetRootOrganizationalUnitResponseBodyOrganizationalUnit {
	s.InstanceId = &v
	return s
}

func (s *GetRootOrganizationalUnitResponseBodyOrganizationalUnit) SetOrganizationalUnitId(v string) *GetRootOrganizationalUnitResponseBodyOrganizationalUnit {
	s.OrganizationalUnitId = &v
	return s
}

func (s *GetRootOrganizationalUnitResponseBodyOrganizationalUnit) SetOrganizationalUnitName(v string) *GetRootOrganizationalUnitResponseBodyOrganizationalUnit {
	s.OrganizationalUnitName = &v
	return s
}

func (s *GetRootOrganizationalUnitResponseBodyOrganizationalUnit) SetUpdateTime(v int64) *GetRootOrganizationalUnitResponseBodyOrganizationalUnit {
	s.UpdateTime = &v
	return s
}

type GetRootOrganizationalUnitResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRootOrganizationalUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRootOrganizationalUnitResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRootOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *GetRootOrganizationalUnitResponse) SetHeaders(v map[string]*string) *GetRootOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *GetRootOrganizationalUnitResponse) SetStatusCode(v int32) *GetRootOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRootOrganizationalUnitResponse) SetBody(v *GetRootOrganizationalUnitResponseBody) *GetRootOrganizationalUnitResponse {
	s.Body = v
	return s
}

type GetSynchronizationJobRequest struct {
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 同步任务ID
	//
	// This parameter is required.
	//
	// example:
	//
	// sync_0000347vjovtcf41li0fgsd98gn24q9njxxxxx
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
}

func (s GetSynchronizationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobRequest) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobRequest) SetInstanceId(v string) *GetSynchronizationJobRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSynchronizationJobRequest) SetSynchronizationJobId(v string) *GetSynchronizationJobRequest {
	s.SynchronizationJobId = &v
	return s
}

type GetSynchronizationJobResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId          *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SynchronizationJob *GetSynchronizationJobResponseBodySynchronizationJob `json:"SynchronizationJob,omitempty" xml:"SynchronizationJob,omitempty" type:"Struct"`
}

func (s GetSynchronizationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBody) SetRequestId(v string) *GetSynchronizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSynchronizationJobResponseBody) SetSynchronizationJob(v *GetSynchronizationJobResponseBodySynchronizationJob) *GetSynchronizationJobResponseBody {
	s.SynchronizationJob = v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJob struct {
	// 同步任务方向
	//
	// example:
	//
	// ingress
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// 同步结束时间
	//
	// example:
	//
	// 1649830226000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 同步任务结果
	Result *GetSynchronizationJobResponseBodySynchronizationJobResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// 同步开始时间
	//
	// example:
	//
	// 1649830226000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 同步任务状态
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 同步任务ID
	//
	// example:
	//
	// sync_0000347vjovtcf41li0fgsd98gn24q9nj9xxxxx
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	// 同步目标ID
	//
	// example:
	//
	// idp_my664lwkhpicbyzirog3nxxxxx
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// 同步目标类型
	//
	// example:
	//
	// identity_provider
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// 同步触发类型
	//
	// example:
	//
	// auto
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJob) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJob) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) SetDirection(v string) *GetSynchronizationJobResponseBodySynchronizationJob {
	s.Direction = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) SetEndTime(v int64) *GetSynchronizationJobResponseBodySynchronizationJob {
	s.EndTime = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) SetResult(v *GetSynchronizationJobResponseBodySynchronizationJobResult) *GetSynchronizationJobResponseBodySynchronizationJob {
	s.Result = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) SetStartTime(v int64) *GetSynchronizationJobResponseBodySynchronizationJob {
	s.StartTime = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) SetStatus(v string) *GetSynchronizationJobResponseBodySynchronizationJob {
	s.Status = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) SetSynchronizationJobId(v string) *GetSynchronizationJobResponseBodySynchronizationJob {
	s.SynchronizationJobId = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) SetTargetId(v string) *GetSynchronizationJobResponseBodySynchronizationJob {
	s.TargetId = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) SetTargetType(v string) *GetSynchronizationJobResponseBodySynchronizationJob {
	s.TargetType = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) SetTriggerType(v string) *GetSynchronizationJobResponseBodySynchronizationJob {
	s.TriggerType = &v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResult struct {
	// 同步结果错误码
	//
	// example:
	//
	// ErrorCodeNotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 同步结果错误信息描述
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 组成员同步结果统计
	GroupMemberStatistics *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics `json:"GroupMemberStatistics,omitempty" xml:"GroupMemberStatistics,omitempty" type:"Struct"`
	// 组同步结果统计
	GroupStatistics *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics `json:"GroupStatistics,omitempty" xml:"GroupStatistics,omitempty" type:"Struct"`
	// 组织同步结果统计
	OrganizationalUnitStatistics *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics `json:"OrganizationalUnitStatistics,omitempty" xml:"OrganizationalUnitStatistics,omitempty" type:"Struct"`
	// 用户同步结果统计
	UserStatistics *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics `json:"UserStatistics,omitempty" xml:"UserStatistics,omitempty" type:"Struct"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResult) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResult) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResult) SetErrorCode(v string) *GetSynchronizationJobResponseBodySynchronizationJobResult {
	s.ErrorCode = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResult) SetErrorMessage(v string) *GetSynchronizationJobResponseBodySynchronizationJobResult {
	s.ErrorMessage = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResult) SetGroupMemberStatistics(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics) *GetSynchronizationJobResponseBodySynchronizationJobResult {
	s.GroupMemberStatistics = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResult) SetGroupStatistics(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics) *GetSynchronizationJobResponseBodySynchronizationJobResult {
	s.GroupStatistics = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResult) SetOrganizationalUnitStatistics(v *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics) *GetSynchronizationJobResponseBodySynchronizationJobResult {
	s.OrganizationalUnitStatistics = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResult) SetUserStatistics(v *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics) *GetSynchronizationJobResponseBodySynchronizationJobResult {
	s.UserStatistics = v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics struct {
	// 绑定结果统计
	Binded *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded `json:"Binded,omitempty" xml:"Binded,omitempty" type:"Struct"`
	// 创建结果统计
	Created *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated `json:"Created,omitempty" xml:"Created,omitempty" type:"Struct"`
	// 删除结果统计
	Deleted *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted `json:"Deleted,omitempty" xml:"Deleted,omitempty" type:"Struct"`
	// 推送结果统计
	Pushed *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed `json:"Pushed,omitempty" xml:"Pushed,omitempty" type:"Struct"`
	// 相同结果统计
	Same *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame `json:"Same,omitempty" xml:"Same,omitempty" type:"Struct"`
	// 更新结果统计
	Updated *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated `json:"Updated,omitempty" xml:"Updated,omitempty" type:"Struct"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics) SetBinded(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics {
	s.Binded = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics) SetCreated(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics {
	s.Created = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics) SetDeleted(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics {
	s.Deleted = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics) SetPushed(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics {
	s.Pushed = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics) SetSame(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics {
	s.Same = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics) SetUpdated(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics {
	s.Updated = v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded struct {
	// 失败数目
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded {
	s.Total = &v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated struct {
	// 失败数目
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated {
	s.Total = &v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted struct {
	// 失败数目
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted {
	s.Total = &v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed struct {
	// 失败数目
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed {
	s.Total = &v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame struct {
	// 失败数目
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame {
	s.Total = &v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated struct {
	// 失败数目
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated {
	s.Total = &v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics struct {
	// 绑定结果统计
	Binded *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded `json:"Binded,omitempty" xml:"Binded,omitempty" type:"Struct"`
	// 创建结果统计
	Created *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated `json:"Created,omitempty" xml:"Created,omitempty" type:"Struct"`
	// 删除结果统计
	Deleted *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted `json:"Deleted,omitempty" xml:"Deleted,omitempty" type:"Struct"`
	// 推送结果统计
	Pushed *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed `json:"Pushed,omitempty" xml:"Pushed,omitempty" type:"Struct"`
	// 相同结果统计
	Same *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame `json:"Same,omitempty" xml:"Same,omitempty" type:"Struct"`
	// 更新结果统计
	Updated *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated `json:"Updated,omitempty" xml:"Updated,omitempty" type:"Struct"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics) SetBinded(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics {
	s.Binded = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics) SetCreated(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics {
	s.Created = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics) SetDeleted(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics {
	s.Deleted = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics) SetPushed(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics {
	s.Pushed = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics) SetSame(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics {
	s.Same = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics) SetUpdated(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics {
	s.Updated = v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded struct {
	// 失败数目
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded {
	s.Total = &v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated struct {
	// 失败数目
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated {
	s.Total = &v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted struct {
	// 失败数目
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted {
	s.Total = &v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed struct {
	// 失败数目
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed {
	s.Total = &v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame struct {
	// 失败数目
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame {
	s.Total = &v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated struct {
	// 失败数目
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated {
	s.Total = &v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics struct {
	// 绑定结果统计
	Binded *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded `json:"Binded,omitempty" xml:"Binded,omitempty" type:"Struct"`
	// 创建结果统计
	Created *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated `json:"Created,omitempty" xml:"Created,omitempty" type:"Struct"`
	// 删除结果统计
	Deleted *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted `json:"Deleted,omitempty" xml:"Deleted,omitempty" type:"Struct"`
	// 推送结果统计
	Pushed *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed `json:"Pushed,omitempty" xml:"Pushed,omitempty" type:"Struct"`
	// 相同结果统计
	Same *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame `json:"Same,omitempty" xml:"Same,omitempty" type:"Struct"`
	// 更新结果统计
	Updated *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated `json:"Updated,omitempty" xml:"Updated,omitempty" type:"Struct"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics) SetBinded(v *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics {
	s.Binded = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics) SetCreated(v *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics {
	s.Created = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics) SetDeleted(v *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics {
	s.Deleted = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics) SetPushed(v *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics {
	s.Pushed = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics) SetSame(v *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics {
	s.Same = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics) SetUpdated(v *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics {
	s.Updated = v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded struct {
	// 失败数目
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded {
	s.Total = &v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated struct {
	// 失败数目
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated {
	s.Total = &v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted struct {
	// 失败数目
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted {
	s.Total = &v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed struct {
	// 失败数目
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed {
	s.Total = &v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame struct {
	// 失败数目
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame {
	s.Total = &v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated struct {
	// 失败数目
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated {
	s.Total = &v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics struct {
	// 绑定结果统计
	Binded *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded `json:"Binded,omitempty" xml:"Binded,omitempty" type:"Struct"`
	// 创建结果统计
	Created *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated `json:"Created,omitempty" xml:"Created,omitempty" type:"Struct"`
	// 删除结果统计
	Deleted *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted `json:"Deleted,omitempty" xml:"Deleted,omitempty" type:"Struct"`
	// 推送结果统计
	Pushed *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed `json:"Pushed,omitempty" xml:"Pushed,omitempty" type:"Struct"`
	// 相同结果统计
	Same *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame `json:"Same,omitempty" xml:"Same,omitempty" type:"Struct"`
	// 更新结果统计
	Updated *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated `json:"Updated,omitempty" xml:"Updated,omitempty" type:"Struct"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics) SetBinded(v *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics {
	s.Binded = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics) SetCreated(v *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics {
	s.Created = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics) SetDeleted(v *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics {
	s.Deleted = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics) SetPushed(v *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics {
	s.Pushed = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics) SetSame(v *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics {
	s.Same = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics) SetUpdated(v *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics {
	s.Updated = v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded struct {
	// 失败数目
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded {
	s.Total = &v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated struct {
	// 失败数目
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated {
	s.Total = &v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted struct {
	// 失败数目
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted {
	s.Total = &v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed struct {
	// 失败数目
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed {
	s.Total = &v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame struct {
	// 失败数目
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame {
	s.Total = &v
	return s
}

type GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated struct {
	// 失败数目
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated {
	s.Total = &v
	return s
}

type GetSynchronizationJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSynchronizationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSynchronizationJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSynchronizationJobResponse) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponse) SetHeaders(v map[string]*string) *GetSynchronizationJobResponse {
	s.Headers = v
	return s
}

func (s *GetSynchronizationJobResponse) SetStatusCode(v int32) *GetSynchronizationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSynchronizationJobResponse) SetBody(v *GetSynchronizationJobResponseBody) *GetSynchronizationJobResponse {
	s.Body = v
	return s
}

type GetUserRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) SetInstanceId(v string) *GetUserRequest {
	s.InstanceId = &v
	return s
}

func (s *GetUserRequest) SetUserId(v string) *GetUserRequest {
	s.UserId = &v
	return s
}

type GetUserResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The data object of the account.
	User *GetUserResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s GetUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserResponseBody) SetUser(v *GetUserResponseBodyUser) *GetUserResponseBody {
	s.User = v
	return s
}

type GetUserResponseBodyUser struct {
	// The time when the account expires. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	AccountExpireTime *int64 `json:"AccountExpireTime,omitempty" xml:"AccountExpireTime,omitempty"`
	// The time when the account was created. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The list of custom fields that describe the account.
	CustomFields []*GetUserResponseBodyUserCustomFields `json:"CustomFields,omitempty" xml:"CustomFields,omitempty" type:"Repeated"`
	// The description of the account.
	//
	// example:
	//
	// Test account
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the account.
	//
	// example:
	//
	// display_name001
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the user who owns the account.
	//
	// example:
	//
	// user@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Indicates whether the email address has been verified. A value of true indicates that the email address has been verified by the user or has been set to the verified status by the administrator. A value of false indicates that the email address has not been verified.
	//
	// example:
	//
	// true
	EmailVerified *bool `json:"EmailVerified,omitempty" xml:"EmailVerified,omitempty"`
	// The organizational units to which the account belongs.
	Groups []*GetUserResponseBodyUserGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The ID of the instance
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the account lock expires. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	LockExpireTime *int64 `json:"LockExpireTime,omitempty" xml:"LockExpireTime,omitempty"`
	// The organizational units to which the account belongs.
	OrganizationalUnits []*GetUserResponseBodyUserOrganizationalUnits `json:"OrganizationalUnits,omitempty" xml:"OrganizationalUnits,omitempty" type:"Repeated"`
	// The time when the password of the account expires. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// 	- If the value -1 is returned, the password does not expire.
	//
	// 	- If no value is returned, the password does not expire.
	//
	// 	- If a UNIX timestamp is returned, the password expires at the indicated point of time.
	//
	// example:
	//
	// 1652085686179
	PasswordExpireTime *int64 `json:"PasswordExpireTime,omitempty" xml:"PasswordExpireTime,omitempty"`
	// Indicates whether a password is set.
	//
	// example:
	//
	// false
	PasswordSet *bool `json:"PasswordSet,omitempty" xml:"PasswordSet,omitempty"`
	// The mobile number of the user who owns the account.
	//
	// example:
	//
	// 156xxxxxxx
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Indicates whether the mobile number has been verified. A value of true indicates that the mobile number has been verified by the user or has been set to the verified status by the administrator. A value of false indicates that the mobile number has not been verified.
	//
	// example:
	//
	// true
	PhoneNumberVerified *bool `json:"PhoneNumberVerified,omitempty" xml:"PhoneNumberVerified,omitempty"`
	// The country code of the mobile number. For example, the country code of China is 86 without 00 or +.
	//
	// example:
	//
	// 86
	PhoneRegion *string `json:"PhoneRegion,omitempty" xml:"PhoneRegion,omitempty"`
	// example:
	//
	// en-US
	PreferredLanguage *string `json:"PreferredLanguage,omitempty" xml:"PreferredLanguage,omitempty"`
	// The ID of the primary organizational unit to which the account belongs.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	PrimaryOrganizationalUnitId *string `json:"PrimaryOrganizationalUnitId,omitempty" xml:"PrimaryOrganizationalUnitId,omitempty"`
	// The time when the account was registered. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	RegisterTime *int64 `json:"RegisterTime,omitempty" xml:"RegisterTime,omitempty"`
	// The status of the account. Valid values:
	//
	// 	- enabled: The account is enabled.
	//
	// 	- disabled: The account is disabled.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the account was last updated. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The external ID of the account. The external ID can be used by external data to map the data of the account in IDaaS EIAM. By default, the external ID is the account ID.
	//
	// For accounts with the same source type and source ID, each account has a unique external ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserExternalId *string `json:"UserExternalId,omitempty" xml:"UserExternalId,omitempty"`
	// The ID of the account.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The source ID of the account.
	//
	// If the account was created in IDaaS, its source ID is the ID of the IDaaS instance. If the account was imported, its source ID is the enterprise ID in the source. For example, if the account was imported from DingTalk, its source ID is the corpId value of the enterprise in DingTalk.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	UserSourceId *string `json:"UserSourceId,omitempty" xml:"UserSourceId,omitempty"`
	// The source type of the account. Valid values:
	//
	// 	- build_in: The account was created in IDaaS.
	//
	// 	- ding_talk: The account was imported from DingTalk.
	//
	// 	- ad: The account was imported from Microsoft Active Directory (AD).
	//
	// 	- ldap: The account was imported from a Lightweight Directory Access Protocol (LDAP) service.
	//
	// example:
	//
	// build_in
	UserSourceType *string `json:"UserSourceType,omitempty" xml:"UserSourceType,omitempty"`
	// The username of the account.
	//
	// example:
	//
	// name001
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetUserResponseBodyUser) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyUser) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUser) SetAccountExpireTime(v int64) *GetUserResponseBodyUser {
	s.AccountExpireTime = &v
	return s
}

func (s *GetUserResponseBodyUser) SetCreateTime(v int64) *GetUserResponseBodyUser {
	s.CreateTime = &v
	return s
}

func (s *GetUserResponseBodyUser) SetCustomFields(v []*GetUserResponseBodyUserCustomFields) *GetUserResponseBodyUser {
	s.CustomFields = v
	return s
}

func (s *GetUserResponseBodyUser) SetDescription(v string) *GetUserResponseBodyUser {
	s.Description = &v
	return s
}

func (s *GetUserResponseBodyUser) SetDisplayName(v string) *GetUserResponseBodyUser {
	s.DisplayName = &v
	return s
}

func (s *GetUserResponseBodyUser) SetEmail(v string) *GetUserResponseBodyUser {
	s.Email = &v
	return s
}

func (s *GetUserResponseBodyUser) SetEmailVerified(v bool) *GetUserResponseBodyUser {
	s.EmailVerified = &v
	return s
}

func (s *GetUserResponseBodyUser) SetGroups(v []*GetUserResponseBodyUserGroups) *GetUserResponseBodyUser {
	s.Groups = v
	return s
}

func (s *GetUserResponseBodyUser) SetInstanceId(v string) *GetUserResponseBodyUser {
	s.InstanceId = &v
	return s
}

func (s *GetUserResponseBodyUser) SetLockExpireTime(v int64) *GetUserResponseBodyUser {
	s.LockExpireTime = &v
	return s
}

func (s *GetUserResponseBodyUser) SetOrganizationalUnits(v []*GetUserResponseBodyUserOrganizationalUnits) *GetUserResponseBodyUser {
	s.OrganizationalUnits = v
	return s
}

func (s *GetUserResponseBodyUser) SetPasswordExpireTime(v int64) *GetUserResponseBodyUser {
	s.PasswordExpireTime = &v
	return s
}

func (s *GetUserResponseBodyUser) SetPasswordSet(v bool) *GetUserResponseBodyUser {
	s.PasswordSet = &v
	return s
}

func (s *GetUserResponseBodyUser) SetPhoneNumber(v string) *GetUserResponseBodyUser {
	s.PhoneNumber = &v
	return s
}

func (s *GetUserResponseBodyUser) SetPhoneNumberVerified(v bool) *GetUserResponseBodyUser {
	s.PhoneNumberVerified = &v
	return s
}

func (s *GetUserResponseBodyUser) SetPhoneRegion(v string) *GetUserResponseBodyUser {
	s.PhoneRegion = &v
	return s
}

func (s *GetUserResponseBodyUser) SetPreferredLanguage(v string) *GetUserResponseBodyUser {
	s.PreferredLanguage = &v
	return s
}

func (s *GetUserResponseBodyUser) SetPrimaryOrganizationalUnitId(v string) *GetUserResponseBodyUser {
	s.PrimaryOrganizationalUnitId = &v
	return s
}

func (s *GetUserResponseBodyUser) SetRegisterTime(v int64) *GetUserResponseBodyUser {
	s.RegisterTime = &v
	return s
}

func (s *GetUserResponseBodyUser) SetStatus(v string) *GetUserResponseBodyUser {
	s.Status = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUpdateTime(v int64) *GetUserResponseBodyUser {
	s.UpdateTime = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserExternalId(v string) *GetUserResponseBodyUser {
	s.UserExternalId = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserId(v string) *GetUserResponseBodyUser {
	s.UserId = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserSourceId(v string) *GetUserResponseBodyUser {
	s.UserSourceId = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserSourceType(v string) *GetUserResponseBodyUser {
	s.UserSourceType = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUsername(v string) *GetUserResponseBodyUser {
	s.Username = &v
	return s
}

type GetUserResponseBodyUserCustomFields struct {
	// The identifier of the custom field.
	//
	// example:
	//
	// age
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The value of the custom field.
	//
	// example:
	//
	// 10
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
}

func (s GetUserResponseBodyUserCustomFields) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyUserCustomFields) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUserCustomFields) SetFieldName(v string) *GetUserResponseBodyUserCustomFields {
	s.FieldName = &v
	return s
}

func (s *GetUserResponseBodyUserCustomFields) SetFieldValue(v string) *GetUserResponseBodyUserCustomFields {
	s.FieldValue = &v
	return s
}

type GetUserResponseBodyUserGroups struct {
	// The description of the organizational unit.
	//
	// example:
	//
	// this is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the organizational unit.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the organizational unit.
	//
	// example:
	//
	// group_test_name
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s GetUserResponseBodyUserGroups) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyUserGroups) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUserGroups) SetDescription(v string) *GetUserResponseBodyUserGroups {
	s.Description = &v
	return s
}

func (s *GetUserResponseBodyUserGroups) SetGroupId(v string) *GetUserResponseBodyUserGroups {
	s.GroupId = &v
	return s
}

func (s *GetUserResponseBodyUserGroups) SetGroupName(v string) *GetUserResponseBodyUserGroups {
	s.GroupName = &v
	return s
}

type GetUserResponseBodyUserOrganizationalUnits struct {
	// The ID of the organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
	// The name of the organizational unit.
	//
	// example:
	//
	// test_ou_name
	OrganizationalUnitName *string `json:"OrganizationalUnitName,omitempty" xml:"OrganizationalUnitName,omitempty"`
	// Indicates whether the organization is the primary organization.
	//
	// example:
	//
	// true
	Primary *bool `json:"Primary,omitempty" xml:"Primary,omitempty"`
}

func (s GetUserResponseBodyUserOrganizationalUnits) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyUserOrganizationalUnits) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUserOrganizationalUnits) SetOrganizationalUnitId(v string) *GetUserResponseBodyUserOrganizationalUnits {
	s.OrganizationalUnitId = &v
	return s
}

func (s *GetUserResponseBodyUserOrganizationalUnits) SetOrganizationalUnitName(v string) *GetUserResponseBodyUserOrganizationalUnits {
	s.OrganizationalUnitName = &v
	return s
}

func (s *GetUserResponseBodyUserOrganizationalUnits) SetPrimary(v bool) *GetUserResponseBodyUserOrganizationalUnits {
	s.Primary = &v
	return s
}

type GetUserResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponse) GoString() string {
	return s.String()
}

func (s *GetUserResponse) SetHeaders(v map[string]*string) *GetUserResponse {
	s.Headers = v
	return s
}

func (s *GetUserResponse) SetStatusCode(v int32) *GetUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserResponse) SetBody(v *GetUserResponseBody) *GetUserResponse {
	s.Body = v
	return s
}

type ListApplicationClientSecretsRequest struct {
	// The ID of the application that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListApplicationClientSecretsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationClientSecretsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationClientSecretsRequest) SetApplicationId(v string) *ListApplicationClientSecretsRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationClientSecretsRequest) SetInstanceId(v string) *ListApplicationClientSecretsRequest {
	s.InstanceId = &v
	return s
}

type ListApplicationClientSecretsResponseBody struct {
	// The information about the client keys.
	ApplicationClientSecrets []*ListApplicationClientSecretsResponseBodyApplicationClientSecrets `json:"ApplicationClientSecrets,omitempty" xml:"ApplicationClientSecrets,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationClientSecretsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationClientSecretsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationClientSecretsResponseBody) SetApplicationClientSecrets(v []*ListApplicationClientSecretsResponseBodyApplicationClientSecrets) *ListApplicationClientSecretsResponseBody {
	s.ApplicationClientSecrets = v
	return s
}

func (s *ListApplicationClientSecretsResponseBody) SetRequestId(v string) *ListApplicationClientSecretsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationClientSecretsResponseBody) SetTotalCount(v int64) *ListApplicationClientSecretsResponseBody {
	s.TotalCount = &v
	return s
}

type ListApplicationClientSecretsResponseBodyApplicationClientSecrets struct {
	// The ID of the application that you want to query.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The client ID of the application.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The client key secret of the application. The value is not masked.
	//
	// example:
	//
	// eyJh*****************************************************************************************************OQ
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// idaas_wdziy4vnjt33ehhf7z2o2nxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the client key was last used. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1649830226000
	LastUsedTime *int64 `json:"LastUsedTime,omitempty" xml:"LastUsedTime,omitempty"`
	// The client key ID of the application.
	//
	// example:
	//
	// sci_k52x2ru63rlkflina5utgkxxxx
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The status of the client key. Valid values:
	//
	// 	- Enabled: The client key is enabled.
	//
	// 	- Disabled: The client key is disabled.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListApplicationClientSecretsResponseBodyApplicationClientSecrets) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationClientSecretsResponseBodyApplicationClientSecrets) GoString() string {
	return s.String()
}

func (s *ListApplicationClientSecretsResponseBodyApplicationClientSecrets) SetApplicationId(v string) *ListApplicationClientSecretsResponseBodyApplicationClientSecrets {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationClientSecretsResponseBodyApplicationClientSecrets) SetClientId(v string) *ListApplicationClientSecretsResponseBodyApplicationClientSecrets {
	s.ClientId = &v
	return s
}

func (s *ListApplicationClientSecretsResponseBodyApplicationClientSecrets) SetClientSecret(v string) *ListApplicationClientSecretsResponseBodyApplicationClientSecrets {
	s.ClientSecret = &v
	return s
}

func (s *ListApplicationClientSecretsResponseBodyApplicationClientSecrets) SetInstanceId(v string) *ListApplicationClientSecretsResponseBodyApplicationClientSecrets {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationClientSecretsResponseBodyApplicationClientSecrets) SetLastUsedTime(v int64) *ListApplicationClientSecretsResponseBodyApplicationClientSecrets {
	s.LastUsedTime = &v
	return s
}

func (s *ListApplicationClientSecretsResponseBodyApplicationClientSecrets) SetSecretId(v string) *ListApplicationClientSecretsResponseBodyApplicationClientSecrets {
	s.SecretId = &v
	return s
}

func (s *ListApplicationClientSecretsResponseBodyApplicationClientSecrets) SetStatus(v string) *ListApplicationClientSecretsResponseBodyApplicationClientSecrets {
	s.Status = &v
	return s
}

type ListApplicationClientSecretsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationClientSecretsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationClientSecretsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationClientSecretsResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationClientSecretsResponse) SetHeaders(v map[string]*string) *ListApplicationClientSecretsResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationClientSecretsResponse) SetStatusCode(v int32) *ListApplicationClientSecretsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationClientSecretsResponse) SetBody(v *ListApplicationClientSecretsResponseBody) *ListApplicationClientSecretsResponse {
	s.Body = v
	return s
}

type ListApplicationsRequest struct {
	// The IDs of the applications.
	//
	// example:
	//
	// Ram Account SSO
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// The name of the application. Only fuzzy match from the leftmost character is supported.
	//
	// example:
	//
	// Ram Account SSO
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The authorization of the application. Valid values:
	//
	// 	- authorize_required: Only the user with explicit authorization can access the application.
	//
	// 	- default_all: By default, all users can access the application.
	//
	// example:
	//
	// authorize_required
	AuthorizationType *string `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of the application. Valid values:
	//
	// 	- Enabled: The application is enabled.
	//
	// 	- Disabled: The application is disabled.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListApplicationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsRequest) SetApplicationIds(v []*string) *ListApplicationsRequest {
	s.ApplicationIds = v
	return s
}

func (s *ListApplicationsRequest) SetApplicationName(v string) *ListApplicationsRequest {
	s.ApplicationName = &v
	return s
}

func (s *ListApplicationsRequest) SetAuthorizationType(v string) *ListApplicationsRequest {
	s.AuthorizationType = &v
	return s
}

func (s *ListApplicationsRequest) SetInstanceId(v string) *ListApplicationsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationsRequest) SetPageNumber(v int64) *ListApplicationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationsRequest) SetPageSize(v int64) *ListApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicationsRequest) SetStatus(v string) *ListApplicationsRequest {
	s.Status = &v
	return s
}

type ListApplicationsResponseBody struct {
	// The details of the applications.
	Applications []*ListApplicationsResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the returned entries.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBody) SetApplications(v []*ListApplicationsResponseBodyApplications) *ListApplicationsResponseBody {
	s.Applications = v
	return s
}

func (s *ListApplicationsResponseBody) SetRequestId(v string) *ListApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsResponseBody) SetTotalCount(v int64) *ListApplicationsResponseBody {
	s.TotalCount = &v
	return s
}

type ListApplicationsResponseBodyApplications struct {
	// The ID of the application.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// SAML Application
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The origin of the application. Valid values:
	//
	// 	- urn:alibaba:idaas:app:source:template: The application is created based on a template.
	//
	// 	- urn:alibaba:idaas: The application is created based on the standard protocol.
	//
	// example:
	//
	// urn:alibaba:idaas:app:source:standard
	ApplicationSourceType *string `json:"ApplicationSourceType,omitempty" xml:"ApplicationSourceType,omitempty"`
	// 应用模板ID
	ApplicationTemplateId *string `json:"ApplicationTemplateId,omitempty" xml:"ApplicationTemplateId,omitempty"`
	// The client ID of the application.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The time when the application was created. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1649830226000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// A single application. The code is pkces.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The features that are supported by the application. The value is a JSON array. Valid values:
	//
	// 	- sso: The application supports SSO.
	//
	// 	- provision: The application supports account synchronization.
	//
	// 	- api_invoke: The application supports custom APIs.
	//
	// example:
	//
	// ["sso", "provision"]
	Features *string `json:"Features,omitempty" xml:"Features,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The URL of the application icon.
	//
	// example:
	//
	// https://img.alicdn.com/imgextra/i4/O1CN01lvYwpv1aGowQXDML9_!!6000000003303-0-tps-580-580.jpg
	LogoUrl *string `json:"LogoUrl,omitempty" xml:"LogoUrl,omitempty"`
	// The service code of the cloud service that manages the application template.
	//
	// example:
	//
	// rpa
	ManagedServiceCode *string `json:"ManagedServiceCode,omitempty" xml:"ManagedServiceCode,omitempty"`
	// Indicates whether the application template is managed by a cloud service.
	//
	// example:
	//
	// true
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The type of the single sign-on (SSO) protocol. Valid values:
	//
	// 	- saml2: the Security Assertion Markup Language (SAML) 2.0 protocol.
	//
	// 	- oidc: the OpenID Connect (OIDC) protocol.
	//
	// example:
	//
	// saml2
	SsoType *string `json:"SsoType,omitempty" xml:"SsoType,omitempty"`
	// The status of the application. Valid values:
	//
	// 	- Enabled: The application is enabled.
	//
	// 	- Disabled: The application is disabled.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the application was last updated. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1649830226000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListApplicationsResponseBodyApplications) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplications) SetApplicationId(v string) *ListApplicationsResponseBodyApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetApplicationName(v string) *ListApplicationsResponseBodyApplications {
	s.ApplicationName = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetApplicationSourceType(v string) *ListApplicationsResponseBodyApplications {
	s.ApplicationSourceType = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetApplicationTemplateId(v string) *ListApplicationsResponseBodyApplications {
	s.ApplicationTemplateId = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetClientId(v string) *ListApplicationsResponseBodyApplications {
	s.ClientId = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetCreateTime(v int64) *ListApplicationsResponseBodyApplications {
	s.CreateTime = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetDescription(v string) *ListApplicationsResponseBodyApplications {
	s.Description = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetFeatures(v string) *ListApplicationsResponseBodyApplications {
	s.Features = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetInstanceId(v string) *ListApplicationsResponseBodyApplications {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetLogoUrl(v string) *ListApplicationsResponseBodyApplications {
	s.LogoUrl = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetManagedServiceCode(v string) *ListApplicationsResponseBodyApplications {
	s.ManagedServiceCode = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetServiceManaged(v bool) *ListApplicationsResponseBodyApplications {
	s.ServiceManaged = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetSsoType(v string) *ListApplicationsResponseBodyApplications {
	s.SsoType = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetStatus(v string) *ListApplicationsResponseBodyApplications {
	s.Status = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetUpdateTime(v int64) *ListApplicationsResponseBodyApplications {
	s.UpdateTime = &v
	return s
}

type ListApplicationsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponse) SetHeaders(v map[string]*string) *ListApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationsResponse) SetStatusCode(v int32) *ListApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationsResponse) SetBody(v *ListApplicationsResponseBody) *ListApplicationsResponse {
	s.Body = v
	return s
}

type ListApplicationsForOrganizationalUnitRequest struct {
	// The IDs of the applications that the EIAM organization can access. You can query a maximum of 100 application IDs at a time.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the EIAM organization.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListApplicationsForOrganizationalUnitRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsForOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsForOrganizationalUnitRequest) SetApplicationIds(v []*string) *ListApplicationsForOrganizationalUnitRequest {
	s.ApplicationIds = v
	return s
}

func (s *ListApplicationsForOrganizationalUnitRequest) SetInstanceId(v string) *ListApplicationsForOrganizationalUnitRequest {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationsForOrganizationalUnitRequest) SetOrganizationalUnitId(v string) *ListApplicationsForOrganizationalUnitRequest {
	s.OrganizationalUnitId = &v
	return s
}

func (s *ListApplicationsForOrganizationalUnitRequest) SetPageNumber(v int64) *ListApplicationsForOrganizationalUnitRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationsForOrganizationalUnitRequest) SetPageSize(v int64) *ListApplicationsForOrganizationalUnitRequest {
	s.PageSize = &v
	return s
}

type ListApplicationsForOrganizationalUnitResponseBody struct {
	// The applications that the EIAM organization can access.
	Applications []*ListApplicationsForOrganizationalUnitResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the returned entries.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationsForOrganizationalUnitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsForOrganizationalUnitResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsForOrganizationalUnitResponseBody) SetApplications(v []*ListApplicationsForOrganizationalUnitResponseBodyApplications) *ListApplicationsForOrganizationalUnitResponseBody {
	s.Applications = v
	return s
}

func (s *ListApplicationsForOrganizationalUnitResponseBody) SetRequestId(v string) *ListApplicationsForOrganizationalUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsForOrganizationalUnitResponseBody) SetTotalCount(v int64) *ListApplicationsForOrganizationalUnitResponseBody {
	s.TotalCount = &v
	return s
}

type ListApplicationsForOrganizationalUnitResponseBodyApplications struct {
	// The ID of the application that the EIAM organization can access.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s ListApplicationsForOrganizationalUnitResponseBodyApplications) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsForOrganizationalUnitResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsForOrganizationalUnitResponseBodyApplications) SetApplicationId(v string) *ListApplicationsForOrganizationalUnitResponseBodyApplications {
	s.ApplicationId = &v
	return s
}

type ListApplicationsForOrganizationalUnitResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationsForOrganizationalUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationsForOrganizationalUnitResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsForOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationsForOrganizationalUnitResponse) SetHeaders(v map[string]*string) *ListApplicationsForOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationsForOrganizationalUnitResponse) SetStatusCode(v int32) *ListApplicationsForOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationsForOrganizationalUnitResponse) SetBody(v *ListApplicationsForOrganizationalUnitResponseBody) *ListApplicationsForOrganizationalUnitResponse {
	s.Body = v
	return s
}

type ListApplicationsForUserRequest struct {
	// The IDs of the applications that the EIAM account can access. You can query a maximum of 100 application IDs at a time.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query mode. Default value: **OnlyDirect**. Valid values:
	//
	// 	- OnlyDirect: Only the direct permissions are queried. Direct permissions are the permissions that are directly granted to the account.
	//
	// 	- IncludeInherit: Both the permissions that are directly granted to the account and the inherited permissions are queried. Inherited permissions are the permissions that an account inherits from the parent organization or the group to which the account belongs.
	//
	// example:
	//
	// OnlyDirect
	QueryMode *string `json:"QueryMode,omitempty" xml:"QueryMode,omitempty"`
	// The ID of the EIAM account.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListApplicationsForUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsForUserRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsForUserRequest) SetApplicationIds(v []*string) *ListApplicationsForUserRequest {
	s.ApplicationIds = v
	return s
}

func (s *ListApplicationsForUserRequest) SetInstanceId(v string) *ListApplicationsForUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationsForUserRequest) SetPageNumber(v int64) *ListApplicationsForUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationsForUserRequest) SetPageSize(v int64) *ListApplicationsForUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicationsForUserRequest) SetQueryMode(v string) *ListApplicationsForUserRequest {
	s.QueryMode = &v
	return s
}

func (s *ListApplicationsForUserRequest) SetUserId(v string) *ListApplicationsForUserRequest {
	s.UserId = &v
	return s
}

type ListApplicationsForUserResponseBody struct {
	// The applications that the EIAM account can access.
	Applications []*ListApplicationsForUserResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the returned entries.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationsForUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsForUserResponseBody) SetApplications(v []*ListApplicationsForUserResponseBodyApplications) *ListApplicationsForUserResponseBody {
	s.Applications = v
	return s
}

func (s *ListApplicationsForUserResponseBody) SetRequestId(v string) *ListApplicationsForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsForUserResponseBody) SetTotalCount(v int64) *ListApplicationsForUserResponseBody {
	s.TotalCount = &v
	return s
}

type ListApplicationsForUserResponseBodyApplications struct {
	// The ID of the application that the EIAM account can access.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// Indicates whether the EIAM account has direct permissions on the application. Valid values:
	//
	// 	- true: The EIAM account has direct permissions on the application.
	//
	// 	- false: The EIAM account does not have direct permissions on the application.
	//
	// example:
	//
	// true
	HasDirectAuthorization *bool `json:"HasDirectAuthorization,omitempty" xml:"HasDirectAuthorization,omitempty"`
	// Indicates whether the EIAM account has inherited permissions on the application. Valid values:
	//
	// 	- true: A parent organization or an organization to which the EIAM account belongs has direct permissions on the application.
	//
	// 	- false: A parent organization or an organization to which the EIAM account belongs does not have direct permissions on the application.
	//
	// example:
	//
	// false
	HasInheritAuthorization *bool `json:"HasInheritAuthorization,omitempty" xml:"HasInheritAuthorization,omitempty"`
}

func (s ListApplicationsForUserResponseBodyApplications) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsForUserResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsForUserResponseBodyApplications) SetApplicationId(v string) *ListApplicationsForUserResponseBodyApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationsForUserResponseBodyApplications) SetHasDirectAuthorization(v bool) *ListApplicationsForUserResponseBodyApplications {
	s.HasDirectAuthorization = &v
	return s
}

func (s *ListApplicationsForUserResponseBodyApplications) SetHasInheritAuthorization(v bool) *ListApplicationsForUserResponseBodyApplications {
	s.HasInheritAuthorization = &v
	return s
}

type ListApplicationsForUserResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationsForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationsForUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsForUserResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationsForUserResponse) SetHeaders(v map[string]*string) *ListApplicationsForUserResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationsForUserResponse) SetStatusCode(v int32) *ListApplicationsForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationsForUserResponse) SetBody(v *ListApplicationsForUserResponseBody) *ListApplicationsForUserResponse {
	s.Body = v
	return s
}

type ListDomainProxyTokensRequest struct {
	// 域名ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// dm_examplexxxxx
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListDomainProxyTokensRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDomainProxyTokensRequest) GoString() string {
	return s.String()
}

func (s *ListDomainProxyTokensRequest) SetDomainId(v string) *ListDomainProxyTokensRequest {
	s.DomainId = &v
	return s
}

func (s *ListDomainProxyTokensRequest) SetInstanceId(v string) *ListDomainProxyTokensRequest {
	s.InstanceId = &v
	return s
}

type ListDomainProxyTokensResponseBody struct {
	DomainProxyTokens []*ListDomainProxyTokensResponseBodyDomainProxyTokens `json:"DomainProxyTokens,omitempty" xml:"DomainProxyTokens,omitempty" type:"Repeated"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDomainProxyTokensResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDomainProxyTokensResponseBody) GoString() string {
	return s.String()
}

func (s *ListDomainProxyTokensResponseBody) SetDomainProxyTokens(v []*ListDomainProxyTokensResponseBodyDomainProxyTokens) *ListDomainProxyTokensResponseBody {
	s.DomainProxyTokens = v
	return s
}

func (s *ListDomainProxyTokensResponseBody) SetRequestId(v string) *ListDomainProxyTokensResponseBody {
	s.RequestId = &v
	return s
}

type ListDomainProxyTokensResponseBodyDomainProxyTokens struct {
	// 域名代理Token创建时间，Unix时间戳格式，单位为毫秒。
	//
	// example:
	//
	// 1649830226000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 域名ID。
	//
	// example:
	//
	// dm_examplexxxx
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// 域名代理Token。
	//
	// example:
	//
	// PTxxxxxxxx
	DomainProxyToken *string `json:"DomainProxyToken,omitempty" xml:"DomainProxyToken,omitempty"`
	// 域名代理Token ID。
	//
	// example:
	//
	// pt_examplexxxx
	DomainProxyTokenId *string `json:"DomainProxyTokenId,omitempty" xml:"DomainProxyTokenId,omitempty"`
	// 实例ID。
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 域名代理Token最近使用时间，Unix时间戳格式，单位为毫秒。
	//
	// example:
	//
	// 1649830226000
	LastUsedTime *int64 `json:"LastUsedTime,omitempty" xml:"LastUsedTime,omitempty"`
	// token状态，枚举类型：(enabled）启用,（disabled）禁用。
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 域名代理Token最近更新时间，Unix时间戳格式，单位为毫秒。
	//
	// example:
	//
	// 1649830226000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListDomainProxyTokensResponseBodyDomainProxyTokens) String() string {
	return tea.Prettify(s)
}

func (s ListDomainProxyTokensResponseBodyDomainProxyTokens) GoString() string {
	return s.String()
}

func (s *ListDomainProxyTokensResponseBodyDomainProxyTokens) SetCreateTime(v int64) *ListDomainProxyTokensResponseBodyDomainProxyTokens {
	s.CreateTime = &v
	return s
}

func (s *ListDomainProxyTokensResponseBodyDomainProxyTokens) SetDomainId(v string) *ListDomainProxyTokensResponseBodyDomainProxyTokens {
	s.DomainId = &v
	return s
}

func (s *ListDomainProxyTokensResponseBodyDomainProxyTokens) SetDomainProxyToken(v string) *ListDomainProxyTokensResponseBodyDomainProxyTokens {
	s.DomainProxyToken = &v
	return s
}

func (s *ListDomainProxyTokensResponseBodyDomainProxyTokens) SetDomainProxyTokenId(v string) *ListDomainProxyTokensResponseBodyDomainProxyTokens {
	s.DomainProxyTokenId = &v
	return s
}

func (s *ListDomainProxyTokensResponseBodyDomainProxyTokens) SetInstanceId(v string) *ListDomainProxyTokensResponseBodyDomainProxyTokens {
	s.InstanceId = &v
	return s
}

func (s *ListDomainProxyTokensResponseBodyDomainProxyTokens) SetLastUsedTime(v int64) *ListDomainProxyTokensResponseBodyDomainProxyTokens {
	s.LastUsedTime = &v
	return s
}

func (s *ListDomainProxyTokensResponseBodyDomainProxyTokens) SetStatus(v string) *ListDomainProxyTokensResponseBodyDomainProxyTokens {
	s.Status = &v
	return s
}

func (s *ListDomainProxyTokensResponseBodyDomainProxyTokens) SetUpdateTime(v int64) *ListDomainProxyTokensResponseBodyDomainProxyTokens {
	s.UpdateTime = &v
	return s
}

type ListDomainProxyTokensResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDomainProxyTokensResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDomainProxyTokensResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDomainProxyTokensResponse) GoString() string {
	return s.String()
}

func (s *ListDomainProxyTokensResponse) SetHeaders(v map[string]*string) *ListDomainProxyTokensResponse {
	s.Headers = v
	return s
}

func (s *ListDomainProxyTokensResponse) SetStatusCode(v int32) *ListDomainProxyTokensResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDomainProxyTokensResponse) SetBody(v *ListDomainProxyTokensResponseBody) *ListDomainProxyTokensResponse {
	s.Body = v
	return s
}

type ListDomainsRequest struct {
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListDomainsRequest) SetInstanceId(v string) *ListDomainsRequest {
	s.InstanceId = &v
	return s
}

type ListDomainsResponseBody struct {
	Domains []*ListDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBody) SetDomains(v []*ListDomainsResponseBodyDomains) *ListDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *ListDomainsResponseBody) SetRequestId(v string) *ListDomainsResponseBody {
	s.RequestId = &v
	return s
}

type ListDomainsResponseBodyDomains struct {
	// 域名创建时间，Unix时间戳格式，单位为毫秒。
	//
	// example:
	//
	// 1649830226000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 是否默认域名。true表示实例默认域名，false表示非默认域名
	//
	// example:
	//
	// false
	DefaultDomain *bool `json:"DefaultDomain,omitempty" xml:"DefaultDomain,omitempty"`
	// 域名。
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// 域名ID。
	//
	// example:
	//
	// dm_examplexxxxx
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// 域名类型。枚举取值:system_init(系统初始化)、user_custom(用户自定义)。
	//
	// example:
	//
	// system_init
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// 域名备案信息。
	Filing *ListDomainsResponseBodyDomainsFiling `json:"Filing,omitempty" xml:"Filing,omitempty" type:"Struct"`
	// 实例ID。
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 域名锁定状态。枚举取值:unlock(正常)、lockByLicense(因License限制不可用)。
	//
	// example:
	//
	// unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// 域名最近更新时间，Unix时间戳格式，单位为毫秒。
	//
	// example:
	//
	// 1649830226000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListDomainsResponseBodyDomains) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBodyDomains) SetCreateTime(v int64) *ListDomainsResponseBodyDomains {
	s.CreateTime = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetDefaultDomain(v bool) *ListDomainsResponseBodyDomains {
	s.DefaultDomain = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetDomain(v string) *ListDomainsResponseBodyDomains {
	s.Domain = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetDomainId(v string) *ListDomainsResponseBodyDomains {
	s.DomainId = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetDomainType(v string) *ListDomainsResponseBodyDomains {
	s.DomainType = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetFiling(v *ListDomainsResponseBodyDomainsFiling) *ListDomainsResponseBodyDomains {
	s.Filing = v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetInstanceId(v string) *ListDomainsResponseBodyDomains {
	s.InstanceId = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetLockMode(v string) *ListDomainsResponseBodyDomains {
	s.LockMode = &v
	return s
}

func (s *ListDomainsResponseBodyDomains) SetUpdateTime(v int64) *ListDomainsResponseBodyDomains {
	s.UpdateTime = &v
	return s
}

type ListDomainsResponseBodyDomainsFiling struct {
	// 域名关联的备案号, 长度最大限制64。
	IcpNumber *string `json:"IcpNumber,omitempty" xml:"IcpNumber,omitempty"`
}

func (s ListDomainsResponseBodyDomainsFiling) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsResponseBodyDomainsFiling) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBodyDomainsFiling) SetIcpNumber(v string) *ListDomainsResponseBodyDomainsFiling {
	s.IcpNumber = &v
	return s
}

type ListDomainsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListDomainsResponse) SetHeaders(v map[string]*string) *ListDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListDomainsResponse) SetStatusCode(v int32) *ListDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDomainsResponse) SetBody(v *ListDomainsResponseBody) *ListDomainsResponse {
	s.Body = v
	return s
}

type ListEiamInstancesRequest struct {
	// 实例ID列表，支持0到100个
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// 实例所属Region
	//
	// example:
	//
	// cn-hangzhou
	InstanceRegionId *string `json:"InstanceRegionId,omitempty" xml:"InstanceRegionId,omitempty"`
}

func (s ListEiamInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEiamInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListEiamInstancesRequest) SetInstanceIds(v []*string) *ListEiamInstancesRequest {
	s.InstanceIds = v
	return s
}

func (s *ListEiamInstancesRequest) SetInstanceRegionId(v string) *ListEiamInstancesRequest {
	s.InstanceRegionId = &v
	return s
}

type ListEiamInstancesResponseBody struct {
	Instances []*ListEiamInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEiamInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEiamInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEiamInstancesResponseBody) SetInstances(v []*ListEiamInstancesResponseBodyInstances) *ListEiamInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListEiamInstancesResponseBody) SetRequestId(v string) *ListEiamInstancesResponseBody {
	s.RequestId = &v
	return s
}

type ListEiamInstancesResponseBodyInstances struct {
	// 实例描述信息
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 实例developer私网域名地址
	//
	// example:
	//
	// eiam-developerapi-cn.vpc-proxy.aliyuncs.com
	DeveloperAPIPrivateDomain *string `json:"DeveloperAPIPrivateDomain,omitempty" xml:"DeveloperAPIPrivateDomain,omitempty"`
	// 实例developer公网域名地址
	//
	// example:
	//
	// eiam-developerapi.cn-hangzhou.aliyuncs.com
	DeveloperAPIPublicDomain *string `json:"DeveloperAPIPublicDomain,omitempty" xml:"DeveloperAPIPublicDomain,omitempty"`
	// 实例id
	//
	// example:
	//
	// idaas_eypq6ljgyeuwmlw672sulxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例状态，Pending(初始状态)、Creating(创建中)、Running(运行中)、Disabled(禁用)、CreateFailed(创建失败)
	//
	// example:
	//
	// RUNNING
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// 实例版本，EIAM2.0/ EIAM1.0
	//
	// example:
	//
	// EIAM 2.0
	InstanceVersion *string `json:"InstanceVersion,omitempty" xml:"InstanceVersion,omitempty"`
	// 实例openApi私网域名地址
	//
	// example:
	//
	// eiam-cn.vpc-proxy.aliyuncs.com
	OpenAPIPrivateDomain *string `json:"OpenAPIPrivateDomain,omitempty" xml:"OpenAPIPrivateDomain,omitempty"`
	// 实例openApi公网域名地址
	//
	// example:
	//
	// eiam.cn-hangzhou.aliyuncs.com
	OpenAPIPublicDomain *string `json:"OpenAPIPublicDomain,omitempty" xml:"OpenAPIPublicDomain,omitempty"`
	// 实例域名地址
	//
	// example:
	//
	// xxxx.aliyunidaas.com
	SSODomain *string `json:"SSODomain,omitempty" xml:"SSODomain,omitempty"`
	// 实例的创建时间
	//
	// example:
	//
	// 1677810869300
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListEiamInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s ListEiamInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListEiamInstancesResponseBodyInstances) SetDescription(v string) *ListEiamInstancesResponseBodyInstances {
	s.Description = &v
	return s
}

func (s *ListEiamInstancesResponseBodyInstances) SetDeveloperAPIPrivateDomain(v string) *ListEiamInstancesResponseBodyInstances {
	s.DeveloperAPIPrivateDomain = &v
	return s
}

func (s *ListEiamInstancesResponseBodyInstances) SetDeveloperAPIPublicDomain(v string) *ListEiamInstancesResponseBodyInstances {
	s.DeveloperAPIPublicDomain = &v
	return s
}

func (s *ListEiamInstancesResponseBodyInstances) SetInstanceId(v string) *ListEiamInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListEiamInstancesResponseBodyInstances) SetInstanceStatus(v string) *ListEiamInstancesResponseBodyInstances {
	s.InstanceStatus = &v
	return s
}

func (s *ListEiamInstancesResponseBodyInstances) SetInstanceVersion(v string) *ListEiamInstancesResponseBodyInstances {
	s.InstanceVersion = &v
	return s
}

func (s *ListEiamInstancesResponseBodyInstances) SetOpenAPIPrivateDomain(v string) *ListEiamInstancesResponseBodyInstances {
	s.OpenAPIPrivateDomain = &v
	return s
}

func (s *ListEiamInstancesResponseBodyInstances) SetOpenAPIPublicDomain(v string) *ListEiamInstancesResponseBodyInstances {
	s.OpenAPIPublicDomain = &v
	return s
}

func (s *ListEiamInstancesResponseBodyInstances) SetSSODomain(v string) *ListEiamInstancesResponseBodyInstances {
	s.SSODomain = &v
	return s
}

func (s *ListEiamInstancesResponseBodyInstances) SetStartTime(v int64) *ListEiamInstancesResponseBodyInstances {
	s.StartTime = &v
	return s
}

type ListEiamInstancesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEiamInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEiamInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEiamInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListEiamInstancesResponse) SetHeaders(v map[string]*string) *ListEiamInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListEiamInstancesResponse) SetStatusCode(v int32) *ListEiamInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEiamInstancesResponse) SetBody(v *ListEiamInstancesResponseBody) *ListEiamInstancesResponse {
	s.Body = v
	return s
}

type ListEiamRegionsResponseBody struct {
	Regions []*ListEiamRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEiamRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEiamRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEiamRegionsResponseBody) SetRegions(v []*ListEiamRegionsResponseBodyRegions) *ListEiamRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListEiamRegionsResponseBody) SetRequestId(v string) *ListEiamRegionsResponseBody {
	s.RequestId = &v
	return s
}

type ListEiamRegionsResponseBodyRegions struct {
	// 地域名称
	//
	// example:
	//
	// 华东1（杭州）
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// 地域ID
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListEiamRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s ListEiamRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListEiamRegionsResponseBodyRegions) SetLocalName(v string) *ListEiamRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *ListEiamRegionsResponseBodyRegions) SetRegionId(v string) *ListEiamRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type ListEiamRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEiamRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEiamRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEiamRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListEiamRegionsResponse) SetHeaders(v map[string]*string) *ListEiamRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListEiamRegionsResponse) SetStatusCode(v int32) *ListEiamRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEiamRegionsResponse) SetBody(v *ListEiamRegionsResponseBody) *ListEiamRegionsResponse {
	s.Body = v
	return s
}

type ListGroupsRequest struct {
	// The external ID of the group.
	//
	// example:
	//
	// group_external_id
	GroupExternalId *string `json:"GroupExternalId,omitempty" xml:"GroupExternalId,omitempty"`
	// The group IDs.
	GroupIds []*string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// The name of the group. If you specify this parameter, the query is based on an exact match.
	//
	// example:
	//
	// name_test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The prefix of the group name. If you specify this parameter, the query follows the leftmost matching principle.
	//
	// example:
	//
	// name
	GroupNameStartsWith *string `json:"GroupNameStartsWith,omitempty" xml:"GroupNameStartsWith,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsRequest) SetGroupExternalId(v string) *ListGroupsRequest {
	s.GroupExternalId = &v
	return s
}

func (s *ListGroupsRequest) SetGroupIds(v []*string) *ListGroupsRequest {
	s.GroupIds = v
	return s
}

func (s *ListGroupsRequest) SetGroupName(v string) *ListGroupsRequest {
	s.GroupName = &v
	return s
}

func (s *ListGroupsRequest) SetGroupNameStartsWith(v string) *ListGroupsRequest {
	s.GroupNameStartsWith = &v
	return s
}

func (s *ListGroupsRequest) SetInstanceId(v string) *ListGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListGroupsRequest) SetPageNumber(v int64) *ListGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGroupsRequest) SetPageSize(v int64) *ListGroupsRequest {
	s.PageSize = &v
	return s
}

type ListGroupsResponseBody struct {
	// The queried account groups.
	Groups []*ListGroupsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned. The maximum number of entries returned at a time depends on the value of PageSize.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBody) SetGroups(v []*ListGroupsResponseBodyGroups) *ListGroupsResponseBody {
	s.Groups = v
	return s
}

func (s *ListGroupsResponseBody) SetRequestId(v string) *ListGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupsResponseBody) SetTotalCount(v int64) *ListGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListGroupsResponseBodyGroups struct {
	// The time at which the group was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1652085686179
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the group.
	//
	// example:
	//
	// test group
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The external ID of the group, which can be used to associate the group with an external system. By default, the external ID is the group ID.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupExternalId *string `json:"GroupExternalId,omitempty" xml:"GroupExternalId,omitempty"`
	// The group ID.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the group.
	//
	// example:
	//
	// group_name
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The source ID of the group. If the group was imported from other services, this value indicates the external source ID. By default, the source ID is the instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	GroupSourceId *string `json:"GroupSourceId,omitempty" xml:"GroupSourceId,omitempty"`
	// The source type of the group. Only build_in may be returned, which indicates that the group was created in IDaaS.
	//
	// example:
	//
	// build_in
	GroupSourceType *string `json:"GroupSourceType,omitempty" xml:"GroupSourceType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time at which the group was last updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1652085686179
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListGroupsResponseBodyGroups) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBodyGroups) SetCreateTime(v int64) *ListGroupsResponseBodyGroups {
	s.CreateTime = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) SetDescription(v string) *ListGroupsResponseBodyGroups {
	s.Description = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) SetGroupExternalId(v string) *ListGroupsResponseBodyGroups {
	s.GroupExternalId = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) SetGroupId(v string) *ListGroupsResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) SetGroupName(v string) *ListGroupsResponseBodyGroups {
	s.GroupName = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) SetGroupSourceId(v string) *ListGroupsResponseBodyGroups {
	s.GroupSourceId = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) SetGroupSourceType(v string) *ListGroupsResponseBodyGroups {
	s.GroupSourceType = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) SetInstanceId(v string) *ListGroupsResponseBodyGroups {
	s.InstanceId = &v
	return s
}

func (s *ListGroupsResponseBodyGroups) SetUpdateTime(v int64) *ListGroupsResponseBodyGroups {
	s.UpdateTime = &v
	return s
}

type ListGroupsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListGroupsResponse) SetHeaders(v map[string]*string) *ListGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListGroupsResponse) SetStatusCode(v int32) *ListGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupsResponse) SetBody(v *ListGroupsResponseBody) *ListGroupsResponse {
	s.Body = v
	return s
}

type ListGroupsForApplicationRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The group IDs. You can specify up to 100 group IDs at a time.
	//
	// example:
	//
	// group_miu8e4t4d7i4u7uwezgr54xxxx
	GroupIds []*string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListGroupsForApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsForApplicationRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsForApplicationRequest) SetApplicationId(v string) *ListGroupsForApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListGroupsForApplicationRequest) SetGroupIds(v []*string) *ListGroupsForApplicationRequest {
	s.GroupIds = v
	return s
}

func (s *ListGroupsForApplicationRequest) SetInstanceId(v string) *ListGroupsForApplicationRequest {
	s.InstanceId = &v
	return s
}

func (s *ListGroupsForApplicationRequest) SetPageNumber(v int64) *ListGroupsForApplicationRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGroupsForApplicationRequest) SetPageSize(v int64) *ListGroupsForApplicationRequest {
	s.PageSize = &v
	return s
}

type ListGroupsForApplicationResponseBody struct {
	// The group IDs.
	Groups []*ListGroupsForApplicationResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGroupsForApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsForApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsForApplicationResponseBody) SetGroups(v []*ListGroupsForApplicationResponseBodyGroups) *ListGroupsForApplicationResponseBody {
	s.Groups = v
	return s
}

func (s *ListGroupsForApplicationResponseBody) SetRequestId(v string) *ListGroupsForApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupsForApplicationResponseBody) SetTotalCount(v int64) *ListGroupsForApplicationResponseBody {
	s.TotalCount = &v
	return s
}

type ListGroupsForApplicationResponseBodyGroups struct {
	// The group ID.
	//
	// example:
	//
	// group_miu8e4t4d7i4u7uwezgr54xxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s ListGroupsForApplicationResponseBodyGroups) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsForApplicationResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *ListGroupsForApplicationResponseBodyGroups) SetGroupId(v string) *ListGroupsForApplicationResponseBodyGroups {
	s.GroupId = &v
	return s
}

type ListGroupsForApplicationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupsForApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupsForApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsForApplicationResponse) GoString() string {
	return s.String()
}

func (s *ListGroupsForApplicationResponse) SetHeaders(v map[string]*string) *ListGroupsForApplicationResponse {
	s.Headers = v
	return s
}

func (s *ListGroupsForApplicationResponse) SetStatusCode(v int32) *ListGroupsForApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupsForApplicationResponse) SetBody(v *ListGroupsForApplicationResponseBody) *ListGroupsForApplicationResponse {
	s.Body = v
	return s
}

type ListGroupsForUserRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListGroupsForUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsForUserRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserRequest) SetInstanceId(v string) *ListGroupsForUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListGroupsForUserRequest) SetPageNumber(v int64) *ListGroupsForUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGroupsForUserRequest) SetPageSize(v int64) *ListGroupsForUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListGroupsForUserRequest) SetUserId(v string) *ListGroupsForUserRequest {
	s.UserId = &v
	return s
}

type ListGroupsForUserResponseBody struct {
	// The queried account groups.
	Groups []*ListGroupsForUserResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned. The maximum number of entries returned at a time depends on the value of PageSize.
	//
	// example:
	//
	// 1000
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGroupsForUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserResponseBody) SetGroups(v []*ListGroupsForUserResponseBodyGroups) *ListGroupsForUserResponseBody {
	s.Groups = v
	return s
}

func (s *ListGroupsForUserResponseBody) SetRequestId(v string) *ListGroupsForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupsForUserResponseBody) SetTotalCount(v int64) *ListGroupsForUserResponseBody {
	s.TotalCount = &v
	return s
}

type ListGroupsForUserResponseBodyGroups struct {
	// The group ID.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupId                       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupMemberRelationSourceId   *string `json:"GroupMemberRelationSourceId,omitempty" xml:"GroupMemberRelationSourceId,omitempty"`
	GroupMemberRelationSourceType *string `json:"GroupMemberRelationSourceType,omitempty" xml:"GroupMemberRelationSourceType,omitempty"`
}

func (s ListGroupsForUserResponseBodyGroups) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsForUserResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserResponseBodyGroups) SetGroupId(v string) *ListGroupsForUserResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *ListGroupsForUserResponseBodyGroups) SetGroupMemberRelationSourceId(v string) *ListGroupsForUserResponseBodyGroups {
	s.GroupMemberRelationSourceId = &v
	return s
}

func (s *ListGroupsForUserResponseBodyGroups) SetGroupMemberRelationSourceType(v string) *ListGroupsForUserResponseBodyGroups {
	s.GroupMemberRelationSourceType = &v
	return s
}

type ListGroupsForUserResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupsForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupsForUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsForUserResponse) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserResponse) SetHeaders(v map[string]*string) *ListGroupsForUserResponse {
	s.Headers = v
	return s
}

func (s *ListGroupsForUserResponse) SetStatusCode(v int32) *ListGroupsForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupsForUserResponse) SetBody(v *ListGroupsForUserResponseBody) *ListGroupsForUserResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	// The list of instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- creating
	//
	// 	- running
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetInstanceIds(v []*string) *ListInstancesRequest {
	s.InstanceIds = v
	return s
}

func (s *ListInstancesRequest) SetPageNumber(v int64) *ListInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int64) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequest) SetStatus(v string) *ListInstancesRequest {
	s.Status = &v
	return s
}

type ListInstancesResponseBody struct {
	// The information of instances.
	Instances []*ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetTotalCount(v int64) *ListInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListInstancesResponseBodyInstances struct {
	// The time when the instance was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1550115455000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The default endpoint of the instance.
	DefaultEndpoint *ListInstancesResponseBodyInstancesDefaultEndpoint `json:"DefaultEndpoint,omitempty" xml:"DefaultEndpoint,omitempty" type:"Struct"`
	// The description of the instance.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_eypq6ljgyeuwmlw672sulxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- creating
	//
	// 	- running
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstances) SetCreateTime(v int64) *ListInstancesResponseBodyInstances {
	s.CreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetDefaultEndpoint(v *ListInstancesResponseBodyInstancesDefaultEndpoint) *ListInstancesResponseBodyInstances {
	s.DefaultEndpoint = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetDescription(v string) *ListInstancesResponseBodyInstances {
	s.Description = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceId(v string) *ListInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetStatus(v string) *ListInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

type ListInstancesResponseBodyInstancesDefaultEndpoint struct {
	// The endpoint of the instance.
	//
	// example:
	//
	// example-xxx.aliyunidaas.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The status of the endpoint. Valid values:
	//
	// 	- resolved
	//
	// 	- unresolved
	//
	// example:
	//
	// resolved
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstancesResponseBodyInstancesDefaultEndpoint) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesDefaultEndpoint) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesDefaultEndpoint) SetEndpoint(v string) *ListInstancesResponseBodyInstancesDefaultEndpoint {
	s.Endpoint = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDefaultEndpoint) SetStatus(v string) *ListInstancesResponseBodyInstancesDefaultEndpoint {
	s.Status = &v
	return s
}

type ListInstancesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesResponse) SetHeaders(v map[string]*string) *ListInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesResponse) SetStatusCode(v int32) *ListInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesResponse) SetBody(v *ListInstancesResponseBody) *ListInstancesResponse {
	s.Body = v
	return s
}

type ListNetworkAccessEndpointAvailableRegionsResponseBody struct {
	Regions []*ListNetworkAccessEndpointAvailableRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNetworkAccessEndpointAvailableRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkAccessEndpointAvailableRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessEndpointAvailableRegionsResponseBody) SetRegions(v []*ListNetworkAccessEndpointAvailableRegionsResponseBodyRegions) *ListNetworkAccessEndpointAvailableRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListNetworkAccessEndpointAvailableRegionsResponseBody) SetRequestId(v string) *ListNetworkAccessEndpointAvailableRegionsResponseBody {
	s.RequestId = &v
	return s
}

type ListNetworkAccessEndpointAvailableRegionsResponseBodyRegions struct {
	// 地域名称。
	//
	// example:
	//
	// 华东1（杭州）
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// 地域ID。
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListNetworkAccessEndpointAvailableRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkAccessEndpointAvailableRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessEndpointAvailableRegionsResponseBodyRegions) SetLocalName(v string) *ListNetworkAccessEndpointAvailableRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *ListNetworkAccessEndpointAvailableRegionsResponseBodyRegions) SetRegionId(v string) *ListNetworkAccessEndpointAvailableRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type ListNetworkAccessEndpointAvailableRegionsResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNetworkAccessEndpointAvailableRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNetworkAccessEndpointAvailableRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkAccessEndpointAvailableRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessEndpointAvailableRegionsResponse) SetHeaders(v map[string]*string) *ListNetworkAccessEndpointAvailableRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListNetworkAccessEndpointAvailableRegionsResponse) SetStatusCode(v int32) *ListNetworkAccessEndpointAvailableRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNetworkAccessEndpointAvailableRegionsResponse) SetBody(v *ListNetworkAccessEndpointAvailableRegionsResponseBody) *ListNetworkAccessEndpointAvailableRegionsResponse {
	s.Body = v
	return s
}

type ListNetworkAccessEndpointAvailableZonesRequest struct {
	// 专属网络端点支持的地域
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	NaeRegionId *string `json:"NaeRegionId,omitempty" xml:"NaeRegionId,omitempty"`
}

func (s ListNetworkAccessEndpointAvailableZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkAccessEndpointAvailableZonesRequest) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessEndpointAvailableZonesRequest) SetNaeRegionId(v string) *ListNetworkAccessEndpointAvailableZonesRequest {
	s.NaeRegionId = &v
	return s
}

type ListNetworkAccessEndpointAvailableZonesResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Zones     []*ListNetworkAccessEndpointAvailableZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s ListNetworkAccessEndpointAvailableZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkAccessEndpointAvailableZonesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessEndpointAvailableZonesResponseBody) SetRequestId(v string) *ListNetworkAccessEndpointAvailableZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNetworkAccessEndpointAvailableZonesResponseBody) SetZones(v []*ListNetworkAccessEndpointAvailableZonesResponseBodyZones) *ListNetworkAccessEndpointAvailableZonesResponseBody {
	s.Zones = v
	return s
}

type ListNetworkAccessEndpointAvailableZonesResponseBodyZones struct {
	// 可用区名称。
	//
	// example:
	//
	// 华东1（杭州）可用区J
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// 可用区ID。
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListNetworkAccessEndpointAvailableZonesResponseBodyZones) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkAccessEndpointAvailableZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessEndpointAvailableZonesResponseBodyZones) SetLocalName(v string) *ListNetworkAccessEndpointAvailableZonesResponseBodyZones {
	s.LocalName = &v
	return s
}

func (s *ListNetworkAccessEndpointAvailableZonesResponseBodyZones) SetZoneId(v string) *ListNetworkAccessEndpointAvailableZonesResponseBodyZones {
	s.ZoneId = &v
	return s
}

type ListNetworkAccessEndpointAvailableZonesResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNetworkAccessEndpointAvailableZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNetworkAccessEndpointAvailableZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkAccessEndpointAvailableZonesResponse) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessEndpointAvailableZonesResponse) SetHeaders(v map[string]*string) *ListNetworkAccessEndpointAvailableZonesResponse {
	s.Headers = v
	return s
}

func (s *ListNetworkAccessEndpointAvailableZonesResponse) SetStatusCode(v int32) *ListNetworkAccessEndpointAvailableZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNetworkAccessEndpointAvailableZonesResponse) SetBody(v *ListNetworkAccessEndpointAvailableZonesResponseBody) *ListNetworkAccessEndpointAvailableZonesResponse {
	s.Body = v
	return s
}

type ListNetworkAccessEndpointsRequest struct {
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 分页查询时每页行数。默认值为20，最大值为100。
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 专属网络端点连接的状态。NetworkAccessEndpointType取值为shared时不生效。
	//
	// example:
	//
	// running
	NetworkAccessEndpointStatus *string `json:"NetworkAccessEndpointStatus,omitempty" xml:"NetworkAccessEndpointStatus,omitempty"`
	// 专属网络端点连接的类型。取值可选范围：1. private - 专属网络端点；2. shared - 共享网络端点
	//
	// example:
	//
	// private
	NetworkAccessEndpointType *string `json:"NetworkAccessEndpointType,omitempty" xml:"NetworkAccessEndpointType,omitempty"`
	// 查询凭证（Token），取值为上一次API调用返回的NextToken参数值。
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 专属网络端点连接的Vpc ID。NetworkAccessEndpointType取值为shared时不生效。
	//
	// example:
	//
	// vpc-examplexxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// 专属网络端点连接的Vpc所属地域，该地域取值必须在ListNetworkAccessEndpointAvailableRegions接口中返回。NetworkAccessEndpointType取值为shared时不生效。
	//
	// example:
	//
	// cn-hangzhou
	VpcRegionId *string `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
}

func (s ListNetworkAccessEndpointsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkAccessEndpointsRequest) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessEndpointsRequest) SetInstanceId(v string) *ListNetworkAccessEndpointsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListNetworkAccessEndpointsRequest) SetMaxResults(v int64) *ListNetworkAccessEndpointsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNetworkAccessEndpointsRequest) SetNetworkAccessEndpointStatus(v string) *ListNetworkAccessEndpointsRequest {
	s.NetworkAccessEndpointStatus = &v
	return s
}

func (s *ListNetworkAccessEndpointsRequest) SetNetworkAccessEndpointType(v string) *ListNetworkAccessEndpointsRequest {
	s.NetworkAccessEndpointType = &v
	return s
}

func (s *ListNetworkAccessEndpointsRequest) SetNextToken(v string) *ListNetworkAccessEndpointsRequest {
	s.NextToken = &v
	return s
}

func (s *ListNetworkAccessEndpointsRequest) SetVpcId(v string) *ListNetworkAccessEndpointsRequest {
	s.VpcId = &v
	return s
}

func (s *ListNetworkAccessEndpointsRequest) SetVpcRegionId(v string) *ListNetworkAccessEndpointsRequest {
	s.VpcRegionId = &v
	return s
}

type ListNetworkAccessEndpointsResponseBody struct {
	NetworkAccessEndpoints []*ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints `json:"NetworkAccessEndpoints,omitempty" xml:"NetworkAccessEndpoints,omitempty" type:"Repeated"`
	// 本次调用返回的查询凭证（Token）值，用于下一次翻页查询。
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNetworkAccessEndpointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkAccessEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessEndpointsResponseBody) SetNetworkAccessEndpoints(v []*ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) *ListNetworkAccessEndpointsResponseBody {
	s.NetworkAccessEndpoints = v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBody) SetNextToken(v string) *ListNetworkAccessEndpointsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBody) SetRequestId(v string) *ListNetworkAccessEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBody) SetTotalCount(v int64) *ListNetworkAccessEndpointsResponseBody {
	s.TotalCount = &v
	return s
}

type ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints struct {
	// 专属网络端点创建时间，Unix时间戳格式，单位为毫秒。
	//
	// example:
	//
	// 1649830226000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 实例ID。
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 专属网络端点ID。
	//
	// example:
	//
	// nae_examplexxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
	// 专属网络端点名称。
	//
	// example:
	//
	// xx业务VPC访问端点
	NetworkAccessEndpointName *string `json:"NetworkAccessEndpointName,omitempty" xml:"NetworkAccessEndpointName,omitempty"`
	// 专属网络端点连接的类型。
	//
	// example:
	//
	// private
	NetworkAccessEndpointType *string `json:"NetworkAccessEndpointType,omitempty" xml:"NetworkAccessEndpointType,omitempty"`
	// 专属网络端点使用的安全组ID。
	//
	// example:
	//
	// sg-examplexxx
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// 专属网络端点状态。
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 专属网络端点最近更新时间，Unix时间戳格式，单位为毫秒。
	//
	// example:
	//
	// 1649830226000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 专属网络端点连接的指定vSwitch列表。
	//
	// example:
	//
	// vsw-examplexxx
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// 专属网络端点连接的VpcID。
	//
	// example:
	//
	// vpc-examplexxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// 专属网络端点连接的Vpc所属地域。
	//
	// example:
	//
	// cn-hangzhou
	VpcRegionId *string `json:"VpcRegionId,omitempty" xml:"VpcRegionId,omitempty"`
}

func (s ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetCreateTime(v int64) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.CreateTime = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetInstanceId(v string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.InstanceId = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetNetworkAccessEndpointId(v string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.NetworkAccessEndpointId = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetNetworkAccessEndpointName(v string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.NetworkAccessEndpointName = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetNetworkAccessEndpointType(v string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.NetworkAccessEndpointType = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetSecurityGroupId(v string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.SecurityGroupId = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetStatus(v string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.Status = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetUpdateTime(v int64) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.UpdateTime = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetVSwitchIds(v []*string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.VSwitchIds = v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetVpcId(v string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.VpcId = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints) SetVpcRegionId(v string) *ListNetworkAccessEndpointsResponseBodyNetworkAccessEndpoints {
	s.VpcRegionId = &v
	return s
}

type ListNetworkAccessEndpointsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNetworkAccessEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNetworkAccessEndpointsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkAccessEndpointsResponse) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessEndpointsResponse) SetHeaders(v map[string]*string) *ListNetworkAccessEndpointsResponse {
	s.Headers = v
	return s
}

func (s *ListNetworkAccessEndpointsResponse) SetStatusCode(v int32) *ListNetworkAccessEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponse) SetBody(v *ListNetworkAccessEndpointsResponseBody) *ListNetworkAccessEndpointsResponse {
	s.Body = v
	return s
}

type ListNetworkAccessPathsRequest struct {
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 专属网络端点ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// nae_examplexxxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
}

func (s ListNetworkAccessPathsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkAccessPathsRequest) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessPathsRequest) SetInstanceId(v string) *ListNetworkAccessPathsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListNetworkAccessPathsRequest) SetNetworkAccessEndpointId(v string) *ListNetworkAccessPathsRequest {
	s.NetworkAccessEndpointId = &v
	return s
}

type ListNetworkAccessPathsResponseBody struct {
	NetworkAccessPaths []*ListNetworkAccessPathsResponseBodyNetworkAccessPaths `json:"NetworkAccessPaths,omitempty" xml:"NetworkAccessPaths,omitempty" type:"Repeated"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNetworkAccessPathsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkAccessPathsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessPathsResponseBody) SetNetworkAccessPaths(v []*ListNetworkAccessPathsResponseBodyNetworkAccessPaths) *ListNetworkAccessPathsResponseBody {
	s.NetworkAccessPaths = v
	return s
}

func (s *ListNetworkAccessPathsResponseBody) SetRequestId(v string) *ListNetworkAccessPathsResponseBody {
	s.RequestId = &v
	return s
}

type ListNetworkAccessPathsResponseBodyNetworkAccessPaths struct {
	// 专属网络端点访问路径创建时间，Unix时间戳格式，单位为毫秒。
	//
	// example:
	//
	// 1649830226000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 实例ID。
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 专属网络端点ID。
	//
	// example:
	//
	// nae_examplexxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
	// 专属网络端点访问路径ID。
	//
	// example:
	//
	// nap_examplexxx
	NetworkAccessPathId *string `json:"NetworkAccessPathId,omitempty" xml:"NetworkAccessPathId,omitempty"`
	// 专属网络端点访问路径使用的ENI ID。
	//
	// example:
	//
	// eni-examplexxx
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// 专属网络端点访问路径使用的ENI私网地址。
	//
	// example:
	//
	// cn-hangzhou
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// 专属网络端点访问路径状态。
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 专属网络端点访问路径最近更新时间，Unix时间戳格式，单位为毫秒。
	//
	// example:
	//
	// 1649830226000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 专属网络端点访问路径的ENI归属的交换机ID。
	//
	// example:
	//
	// vsw-examplexxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s ListNetworkAccessPathsResponseBodyNetworkAccessPaths) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkAccessPathsResponseBodyNetworkAccessPaths) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) SetCreateTime(v int64) *ListNetworkAccessPathsResponseBodyNetworkAccessPaths {
	s.CreateTime = &v
	return s
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) SetInstanceId(v string) *ListNetworkAccessPathsResponseBodyNetworkAccessPaths {
	s.InstanceId = &v
	return s
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) SetNetworkAccessEndpointId(v string) *ListNetworkAccessPathsResponseBodyNetworkAccessPaths {
	s.NetworkAccessEndpointId = &v
	return s
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) SetNetworkAccessPathId(v string) *ListNetworkAccessPathsResponseBodyNetworkAccessPaths {
	s.NetworkAccessPathId = &v
	return s
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) SetNetworkInterfaceId(v string) *ListNetworkAccessPathsResponseBodyNetworkAccessPaths {
	s.NetworkInterfaceId = &v
	return s
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) SetPrivateIpAddress(v string) *ListNetworkAccessPathsResponseBodyNetworkAccessPaths {
	s.PrivateIpAddress = &v
	return s
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) SetStatus(v string) *ListNetworkAccessPathsResponseBodyNetworkAccessPaths {
	s.Status = &v
	return s
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) SetUpdateTime(v int64) *ListNetworkAccessPathsResponseBodyNetworkAccessPaths {
	s.UpdateTime = &v
	return s
}

func (s *ListNetworkAccessPathsResponseBodyNetworkAccessPaths) SetVSwitchId(v string) *ListNetworkAccessPathsResponseBodyNetworkAccessPaths {
	s.VSwitchId = &v
	return s
}

type ListNetworkAccessPathsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNetworkAccessPathsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNetworkAccessPathsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNetworkAccessPathsResponse) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessPathsResponse) SetHeaders(v map[string]*string) *ListNetworkAccessPathsResponse {
	s.Headers = v
	return s
}

func (s *ListNetworkAccessPathsResponse) SetStatusCode(v int32) *ListNetworkAccessPathsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNetworkAccessPathsResponse) SetBody(v *ListNetworkAccessPathsResponseBody) *ListNetworkAccessPathsResponse {
	s.Body = v
	return s
}

type ListOrganizationalUnitParentsRequest struct {
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 组织ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
}

func (s ListOrganizationalUnitParentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitParentsRequest) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitParentsRequest) SetInstanceId(v string) *ListOrganizationalUnitParentsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOrganizationalUnitParentsRequest) SetOrganizationalUnitId(v string) *ListOrganizationalUnitParentsRequest {
	s.OrganizationalUnitId = &v
	return s
}

type ListOrganizationalUnitParentsResponseBody struct {
	Parents []*ListOrganizationalUnitParentsResponseBodyParents `json:"Parents,omitempty" xml:"Parents,omitempty" type:"Repeated"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListOrganizationalUnitParentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitParentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitParentsResponseBody) SetParents(v []*ListOrganizationalUnitParentsResponseBodyParents) *ListOrganizationalUnitParentsResponseBody {
	s.Parents = v
	return s
}

func (s *ListOrganizationalUnitParentsResponseBody) SetRequestId(v string) *ListOrganizationalUnitParentsResponseBody {
	s.RequestId = &v
	return s
}

type ListOrganizationalUnitParentsResponseBodyParents struct {
	// 组织ID
	//
	// example:
	//
	// ou_4lag76zc2km5ssg5vsmm2lznvu
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
	// 父组织ID
	//
	// example:
	//
	// ou_x3beoyepv2ls5iwuge3xhjkwbm
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s ListOrganizationalUnitParentsResponseBodyParents) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitParentsResponseBodyParents) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitParentsResponseBodyParents) SetOrganizationalUnitId(v string) *ListOrganizationalUnitParentsResponseBodyParents {
	s.OrganizationalUnitId = &v
	return s
}

func (s *ListOrganizationalUnitParentsResponseBodyParents) SetParentId(v string) *ListOrganizationalUnitParentsResponseBodyParents {
	s.ParentId = &v
	return s
}

type ListOrganizationalUnitParentsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOrganizationalUnitParentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOrganizationalUnitParentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitParentsResponse) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitParentsResponse) SetHeaders(v map[string]*string) *ListOrganizationalUnitParentsResponse {
	s.Headers = v
	return s
}

func (s *ListOrganizationalUnitParentsResponse) SetStatusCode(v int32) *ListOrganizationalUnitParentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOrganizationalUnitParentsResponse) SetBody(v *ListOrganizationalUnitParentsResponseBody) *ListOrganizationalUnitParentsResponse {
	s.Body = v
	return s
}

type ListOrganizationalUnitsRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 组织ID列表。size限制最大100。
	//
	// example:
	//
	// [ou_wovwffm62xifdziem7an7xxxxx]
	OrganizationalUnitIds []*string `json:"OrganizationalUnitIds,omitempty" xml:"OrganizationalUnitIds,omitempty" type:"Repeated"`
	// The name of the organizational unit.
	//
	// example:
	//
	// name_001
	OrganizationalUnitName *string `json:"OrganizationalUnitName,omitempty" xml:"OrganizationalUnitName,omitempty"`
	// 组织名称，左匹配
	//
	// example:
	//
	// name
	OrganizationalUnitNameStartsWith *string `json:"OrganizationalUnitNameStartsWith,omitempty" xml:"OrganizationalUnitNameStartsWith,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the parent organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s ListOrganizationalUnitsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitsRequest) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsRequest) SetInstanceId(v string) *ListOrganizationalUnitsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOrganizationalUnitsRequest) SetOrganizationalUnitIds(v []*string) *ListOrganizationalUnitsRequest {
	s.OrganizationalUnitIds = v
	return s
}

func (s *ListOrganizationalUnitsRequest) SetOrganizationalUnitName(v string) *ListOrganizationalUnitsRequest {
	s.OrganizationalUnitName = &v
	return s
}

func (s *ListOrganizationalUnitsRequest) SetOrganizationalUnitNameStartsWith(v string) *ListOrganizationalUnitsRequest {
	s.OrganizationalUnitNameStartsWith = &v
	return s
}

func (s *ListOrganizationalUnitsRequest) SetPageNumber(v int64) *ListOrganizationalUnitsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOrganizationalUnitsRequest) SetPageSize(v int64) *ListOrganizationalUnitsRequest {
	s.PageSize = &v
	return s
}

func (s *ListOrganizationalUnitsRequest) SetParentId(v string) *ListOrganizationalUnitsRequest {
	s.ParentId = &v
	return s
}

type ListOrganizationalUnitsResponseBody struct {
	// The list of data objects of organizational units.
	OrganizationalUnits []*ListOrganizationalUnitsResponseBodyOrganizationalUnits `json:"OrganizationalUnits,omitempty" xml:"OrganizationalUnits,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries in the list.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOrganizationalUnitsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsResponseBody) SetOrganizationalUnits(v []*ListOrganizationalUnitsResponseBodyOrganizationalUnits) *ListOrganizationalUnitsResponseBody {
	s.OrganizationalUnits = v
	return s
}

func (s *ListOrganizationalUnitsResponseBody) SetRequestId(v string) *ListOrganizationalUnitsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBody) SetTotalCount(v int64) *ListOrganizationalUnitsResponseBody {
	s.TotalCount = &v
	return s
}

type ListOrganizationalUnitsResponseBodyOrganizationalUnits struct {
	// The time when the organizational unit was created. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the organizational unit.
	//
	// example:
	//
	// Test organizational unit
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the node is a leaf node.
	//
	// example:
	//
	// false
	Leaf *bool `json:"Leaf,omitempty" xml:"Leaf,omitempty"`
	// The external ID of the organizational unit. The external ID can be used by external data to map the data of the organizational unit in IDaaS EIAM. By default, the external ID is the organizational unit ID.
	//
	// For organizational units with the same source type and source ID, each organizational unit has a unique external ID.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitExternalId *string `json:"OrganizationalUnitExternalId,omitempty" xml:"OrganizationalUnitExternalId,omitempty"`
	// The ID of the organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
	// 组织名称。
	//
	// example:
	//
	// test_organizationalUnit_name
	OrganizationalUnitName *string `json:"OrganizationalUnitName,omitempty" xml:"OrganizationalUnitName,omitempty"`
	// The source ID of the organizational unit.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	OrganizationalUnitSourceId *string `json:"OrganizationalUnitSourceId,omitempty" xml:"OrganizationalUnitSourceId,omitempty"`
	// The source type of the organizational unit. Valid values:
	//
	// 	- build_in: The organizational unit was created in IDaaS.
	//
	// 	- ding_talk: The organizational unit was imported from DingTalk.
	//
	// 	- ad: The organizational unit was imported from Microsoft Active Directory (AD).
	//
	// 	- ldap: The organizational unit was imported from a Lightweight Directory Access Protocol (LDAP) service.
	//
	// example:
	//
	// build_in
	OrganizationalUnitSourceType *string `json:"OrganizationalUnitSourceType,omitempty" xml:"OrganizationalUnitSourceType,omitempty"`
	// The ID of the parent organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The time when the organizational unit was last updated. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListOrganizationalUnitsResponseBodyOrganizationalUnits) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitsResponseBodyOrganizationalUnits) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) SetCreateTime(v int64) *ListOrganizationalUnitsResponseBodyOrganizationalUnits {
	s.CreateTime = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) SetDescription(v string) *ListOrganizationalUnitsResponseBodyOrganizationalUnits {
	s.Description = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) SetInstanceId(v string) *ListOrganizationalUnitsResponseBodyOrganizationalUnits {
	s.InstanceId = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) SetLeaf(v bool) *ListOrganizationalUnitsResponseBodyOrganizationalUnits {
	s.Leaf = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) SetOrganizationalUnitExternalId(v string) *ListOrganizationalUnitsResponseBodyOrganizationalUnits {
	s.OrganizationalUnitExternalId = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) SetOrganizationalUnitId(v string) *ListOrganizationalUnitsResponseBodyOrganizationalUnits {
	s.OrganizationalUnitId = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) SetOrganizationalUnitName(v string) *ListOrganizationalUnitsResponseBodyOrganizationalUnits {
	s.OrganizationalUnitName = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) SetOrganizationalUnitSourceId(v string) *ListOrganizationalUnitsResponseBodyOrganizationalUnits {
	s.OrganizationalUnitSourceId = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) SetOrganizationalUnitSourceType(v string) *ListOrganizationalUnitsResponseBodyOrganizationalUnits {
	s.OrganizationalUnitSourceType = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) SetParentId(v string) *ListOrganizationalUnitsResponseBodyOrganizationalUnits {
	s.ParentId = &v
	return s
}

func (s *ListOrganizationalUnitsResponseBodyOrganizationalUnits) SetUpdateTime(v int64) *ListOrganizationalUnitsResponseBodyOrganizationalUnits {
	s.UpdateTime = &v
	return s
}

type ListOrganizationalUnitsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOrganizationalUnitsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOrganizationalUnitsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitsResponse) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsResponse) SetHeaders(v map[string]*string) *ListOrganizationalUnitsResponse {
	s.Headers = v
	return s
}

func (s *ListOrganizationalUnitsResponse) SetStatusCode(v int32) *ListOrganizationalUnitsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOrganizationalUnitsResponse) SetBody(v *ListOrganizationalUnitsResponseBody) *ListOrganizationalUnitsResponse {
	s.Body = v
	return s
}

type ListOrganizationalUnitsForApplicationRequest struct {
	// The ID of the application that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IDs of the organizations that are allowed to access the application. You can query a maximum of 100 organization IDs at a time.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitIds []*string `json:"OrganizationalUnitIds,omitempty" xml:"OrganizationalUnitIds,omitempty" type:"Repeated"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListOrganizationalUnitsForApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitsForApplicationRequest) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsForApplicationRequest) SetApplicationId(v string) *ListOrganizationalUnitsForApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListOrganizationalUnitsForApplicationRequest) SetInstanceId(v string) *ListOrganizationalUnitsForApplicationRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOrganizationalUnitsForApplicationRequest) SetOrganizationalUnitIds(v []*string) *ListOrganizationalUnitsForApplicationRequest {
	s.OrganizationalUnitIds = v
	return s
}

func (s *ListOrganizationalUnitsForApplicationRequest) SetPageNumber(v int64) *ListOrganizationalUnitsForApplicationRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOrganizationalUnitsForApplicationRequest) SetPageSize(v int64) *ListOrganizationalUnitsForApplicationRequest {
	s.PageSize = &v
	return s
}

type ListOrganizationalUnitsForApplicationResponseBody struct {
	// The IDs of the organizations that are allowed to access the application.
	OrganizationalUnits []*ListOrganizationalUnitsForApplicationResponseBodyOrganizationalUnits `json:"OrganizationalUnits,omitempty" xml:"OrganizationalUnits,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the returned entries.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOrganizationalUnitsForApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitsForApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsForApplicationResponseBody) SetOrganizationalUnits(v []*ListOrganizationalUnitsForApplicationResponseBodyOrganizationalUnits) *ListOrganizationalUnitsForApplicationResponseBody {
	s.OrganizationalUnits = v
	return s
}

func (s *ListOrganizationalUnitsForApplicationResponseBody) SetRequestId(v string) *ListOrganizationalUnitsForApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrganizationalUnitsForApplicationResponseBody) SetTotalCount(v int64) *ListOrganizationalUnitsForApplicationResponseBody {
	s.TotalCount = &v
	return s
}

type ListOrganizationalUnitsForApplicationResponseBodyOrganizationalUnits struct {
	// The ID of the organization that is allowed to access the application.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
}

func (s ListOrganizationalUnitsForApplicationResponseBodyOrganizationalUnits) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitsForApplicationResponseBodyOrganizationalUnits) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsForApplicationResponseBodyOrganizationalUnits) SetOrganizationalUnitId(v string) *ListOrganizationalUnitsForApplicationResponseBodyOrganizationalUnits {
	s.OrganizationalUnitId = &v
	return s
}

type ListOrganizationalUnitsForApplicationResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOrganizationalUnitsForApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOrganizationalUnitsForApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrganizationalUnitsForApplicationResponse) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsForApplicationResponse) SetHeaders(v map[string]*string) *ListOrganizationalUnitsForApplicationResponse {
	s.Headers = v
	return s
}

func (s *ListOrganizationalUnitsForApplicationResponse) SetStatusCode(v int32) *ListOrganizationalUnitsForApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOrganizationalUnitsForApplicationResponse) SetBody(v *ListOrganizationalUnitsForApplicationResponseBody) *ListOrganizationalUnitsForApplicationResponse {
	s.Body = v
	return s
}

type ListRegionsResponseBody struct {
	// The supported regions.
	Regions []*ListRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) SetRegions(v []*ListRegionsResponseBodyRegions) *ListRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

type ListRegionsResponseBodyRegions struct {
	// The name of the region.
	//
	// example:
	//
	// China (Hangzhou)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The endpoint of the region.
	//
	// example:
	//
	// eiam.cn-hangzhou.aliyuncs.com
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyRegions) SetLocalName(v string) *ListRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *ListRegionsResponseBodyRegions) SetRegionEndpoint(v string) *ListRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *ListRegionsResponseBodyRegions) SetRegionId(v string) *ListRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type ListRegionsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListRegionsResponse) SetHeaders(v map[string]*string) *ListRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListRegionsResponse) SetStatusCode(v int32) *ListRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRegionsResponse) SetBody(v *ListRegionsResponseBody) *ListRegionsResponse {
	s.Body = v
	return s
}

type ListSynchronizationJobsRequest struct {
	// 同步方向[ingress,egress]
	//
	// example:
	//
	// ingress
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// 同步结束时间
	//
	// example:
	//
	// 1649830226000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 分页查询时每页行数。默认值为20，最大值为100。
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 查询凭证（Token），取值为上一次API调用返回的NextToken参数值。
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 当前查询的列表页码，默认为1。
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 当前查询的列表页码，默认为20。
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 同步开始时间
	//
	// example:
	//
	// 1649830226000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 同步状态[pending,running,suspending,failed,partial_success,success]
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 同步目标ID
	//
	// example:
	//
	// target_001
	TargetIds []*string `json:"TargetIds,omitempty" xml:"TargetIds,omitempty" type:"Repeated"`
	// 同步目标类型[identity_provider,organizational_unit,application,user]
	//
	// example:
	//
	// identity_provider
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListSynchronizationJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsRequest) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsRequest) SetDirection(v string) *ListSynchronizationJobsRequest {
	s.Direction = &v
	return s
}

func (s *ListSynchronizationJobsRequest) SetEndTime(v int64) *ListSynchronizationJobsRequest {
	s.EndTime = &v
	return s
}

func (s *ListSynchronizationJobsRequest) SetInstanceId(v string) *ListSynchronizationJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSynchronizationJobsRequest) SetMaxResults(v int64) *ListSynchronizationJobsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSynchronizationJobsRequest) SetNextToken(v string) *ListSynchronizationJobsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSynchronizationJobsRequest) SetPageNumber(v int64) *ListSynchronizationJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSynchronizationJobsRequest) SetPageSize(v int64) *ListSynchronizationJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSynchronizationJobsRequest) SetStartTime(v int64) *ListSynchronizationJobsRequest {
	s.StartTime = &v
	return s
}

func (s *ListSynchronizationJobsRequest) SetStatus(v string) *ListSynchronizationJobsRequest {
	s.Status = &v
	return s
}

func (s *ListSynchronizationJobsRequest) SetTargetIds(v []*string) *ListSynchronizationJobsRequest {
	s.TargetIds = v
	return s
}

func (s *ListSynchronizationJobsRequest) SetTargetType(v string) *ListSynchronizationJobsRequest {
	s.TargetType = &v
	return s
}

type ListSynchronizationJobsResponseBody struct {
	// 本次调用返回的查询凭证（Token）值，用于下一次翻页查询。
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId           *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SynchronizationJobs []*ListSynchronizationJobsResponseBodySynchronizationJobs `json:"SynchronizationJobs,omitempty" xml:"SynchronizationJobs,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSynchronizationJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBody) SetNextToken(v string) *ListSynchronizationJobsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSynchronizationJobsResponseBody) SetRequestId(v string) *ListSynchronizationJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSynchronizationJobsResponseBody) SetSynchronizationJobs(v []*ListSynchronizationJobsResponseBodySynchronizationJobs) *ListSynchronizationJobsResponseBody {
	s.SynchronizationJobs = v
	return s
}

func (s *ListSynchronizationJobsResponseBody) SetTotalCount(v int64) *ListSynchronizationJobsResponseBody {
	s.TotalCount = &v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobs struct {
	// 同步任务方向
	//
	// example:
	//
	// ingress
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// 同步结束时间
	//
	// example:
	//
	// 1649830226000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 同步任务结果
	Result *ListSynchronizationJobsResponseBodySynchronizationJobsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// 同步开始时间
	//
	// example:
	//
	// 1649830226000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 同步任务状态
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 同步任务ID
	//
	// example:
	//
	// sync_0000347vjovtcf41li0fgsd98gn24q9njxxxxx
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	// 同步目标ID
	//
	// example:
	//
	// idp_my664lwkhpicbyzirog3xxxxx
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// 同步目标类型
	//
	// example:
	//
	// identity_provider
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// 同步触发类型
	//
	// example:
	//
	// auto
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobs) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobs) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) SetDirection(v string) *ListSynchronizationJobsResponseBodySynchronizationJobs {
	s.Direction = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) SetEndTime(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobs {
	s.EndTime = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) SetResult(v *ListSynchronizationJobsResponseBodySynchronizationJobsResult) *ListSynchronizationJobsResponseBodySynchronizationJobs {
	s.Result = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) SetStartTime(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobs {
	s.StartTime = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) SetStatus(v string) *ListSynchronizationJobsResponseBodySynchronizationJobs {
	s.Status = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) SetSynchronizationJobId(v string) *ListSynchronizationJobsResponseBodySynchronizationJobs {
	s.SynchronizationJobId = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) SetTargetId(v string) *ListSynchronizationJobsResponseBodySynchronizationJobs {
	s.TargetId = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) SetTargetType(v string) *ListSynchronizationJobsResponseBodySynchronizationJobs {
	s.TargetType = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) SetTriggerType(v string) *ListSynchronizationJobsResponseBodySynchronizationJobs {
	s.TriggerType = &v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResult struct {
	// 同步结果错误码
	//
	// example:
	//
	// ErrorCodeNotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 同步结果错误信息描述
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 组成员同步结果统计
	GroupMemberStatistics *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics `json:"GroupMemberStatistics,omitempty" xml:"GroupMemberStatistics,omitempty" type:"Struct"`
	// 组同步结果统计
	GroupStatistics *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics `json:"GroupStatistics,omitempty" xml:"GroupStatistics,omitempty" type:"Struct"`
	// 组织同步结果统计
	OrganizationalUnitStatistics *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics `json:"OrganizationalUnitStatistics,omitempty" xml:"OrganizationalUnitStatistics,omitempty" type:"Struct"`
	// 用户同步结果统计
	UserStatistics *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics `json:"UserStatistics,omitempty" xml:"UserStatistics,omitempty" type:"Struct"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResult) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResult) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResult) SetErrorCode(v string) *ListSynchronizationJobsResponseBodySynchronizationJobsResult {
	s.ErrorCode = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResult) SetErrorMessage(v string) *ListSynchronizationJobsResponseBodySynchronizationJobsResult {
	s.ErrorMessage = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResult) SetGroupMemberStatistics(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics) *ListSynchronizationJobsResponseBodySynchronizationJobsResult {
	s.GroupMemberStatistics = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResult) SetGroupStatistics(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics) *ListSynchronizationJobsResponseBodySynchronizationJobsResult {
	s.GroupStatistics = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResult) SetOrganizationalUnitStatistics(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics) *ListSynchronizationJobsResponseBodySynchronizationJobsResult {
	s.OrganizationalUnitStatistics = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResult) SetUserStatistics(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics) *ListSynchronizationJobsResponseBodySynchronizationJobsResult {
	s.UserStatistics = v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics struct {
	// 绑定结果统计
	Binded *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded `json:"Binded,omitempty" xml:"Binded,omitempty" type:"Struct"`
	// 创建结果统计
	Created *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated `json:"Created,omitempty" xml:"Created,omitempty" type:"Struct"`
	// 删除结果统计
	Deleted *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted `json:"Deleted,omitempty" xml:"Deleted,omitempty" type:"Struct"`
	// 推送结果统计
	Pushed *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed `json:"Pushed,omitempty" xml:"Pushed,omitempty" type:"Struct"`
	// 相同结果统计
	Same *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame `json:"Same,omitempty" xml:"Same,omitempty" type:"Struct"`
	// 更新结果统计
	Updated *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated `json:"Updated,omitempty" xml:"Updated,omitempty" type:"Struct"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics) SetBinded(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics {
	s.Binded = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics) SetCreated(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics {
	s.Created = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics) SetDeleted(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics {
	s.Deleted = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics) SetPushed(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics {
	s.Pushed = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics) SetSame(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics {
	s.Same = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics) SetUpdated(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics {
	s.Updated = v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded {
	s.Total = &v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated {
	s.Total = &v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted {
	s.Total = &v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed {
	s.Total = &v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame {
	s.Total = &v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated {
	s.Total = &v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics struct {
	// 绑定结果统计
	Binded *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded `json:"Binded,omitempty" xml:"Binded,omitempty" type:"Struct"`
	// 创建结果统计
	Created *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated `json:"Created,omitempty" xml:"Created,omitempty" type:"Struct"`
	// 删除结果统计
	Deleted *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted `json:"Deleted,omitempty" xml:"Deleted,omitempty" type:"Struct"`
	// 推送结果统计
	Pushed *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed `json:"Pushed,omitempty" xml:"Pushed,omitempty" type:"Struct"`
	// 相同结果统计
	Same *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame `json:"Same,omitempty" xml:"Same,omitempty" type:"Struct"`
	// 更新结果统计
	Updated *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated `json:"Updated,omitempty" xml:"Updated,omitempty" type:"Struct"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics) SetBinded(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics {
	s.Binded = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics) SetCreated(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics {
	s.Created = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics) SetDeleted(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics {
	s.Deleted = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics) SetPushed(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics {
	s.Pushed = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics) SetSame(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics {
	s.Same = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics) SetUpdated(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics {
	s.Updated = v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded {
	s.Total = &v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated {
	s.Total = &v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted {
	s.Total = &v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed {
	s.Total = &v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame {
	s.Total = &v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated {
	s.Total = &v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics struct {
	// 绑定结果统计
	Binded *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded `json:"Binded,omitempty" xml:"Binded,omitempty" type:"Struct"`
	// 创建结果统计
	Created *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated `json:"Created,omitempty" xml:"Created,omitempty" type:"Struct"`
	// 删除结果统计
	Deleted *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted `json:"Deleted,omitempty" xml:"Deleted,omitempty" type:"Struct"`
	// 推送结果统计
	Pushed *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed `json:"Pushed,omitempty" xml:"Pushed,omitempty" type:"Struct"`
	// 相同结果统计
	Same *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame `json:"Same,omitempty" xml:"Same,omitempty" type:"Struct"`
	// 更新结果统计
	Updated *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated `json:"Updated,omitempty" xml:"Updated,omitempty" type:"Struct"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics) SetBinded(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics {
	s.Binded = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics) SetCreated(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics {
	s.Created = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics) SetDeleted(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics {
	s.Deleted = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics) SetPushed(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics {
	s.Pushed = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics) SetSame(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics {
	s.Same = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics) SetUpdated(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics {
	s.Updated = v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded {
	s.Total = &v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated {
	s.Total = &v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted {
	s.Total = &v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed {
	s.Total = &v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame {
	s.Total = &v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated {
	s.Total = &v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics struct {
	// 绑定结果统计
	Binded *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded `json:"Binded,omitempty" xml:"Binded,omitempty" type:"Struct"`
	// 创建结果统计
	Created *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated `json:"Created,omitempty" xml:"Created,omitempty" type:"Struct"`
	// 删除结果统计
	Deleted *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted `json:"Deleted,omitempty" xml:"Deleted,omitempty" type:"Struct"`
	// 推送结果统计
	Pushed *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed `json:"Pushed,omitempty" xml:"Pushed,omitempty" type:"Struct"`
	// 相同结果统计
	Same *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame `json:"Same,omitempty" xml:"Same,omitempty" type:"Struct"`
	// 更新结果统计
	Updated *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated `json:"Updated,omitempty" xml:"Updated,omitempty" type:"Struct"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics) SetBinded(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics {
	s.Binded = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics) SetCreated(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics {
	s.Created = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics) SetDeleted(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics {
	s.Deleted = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics) SetPushed(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics {
	s.Pushed = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics) SetSame(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics {
	s.Same = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics) SetUpdated(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics {
	s.Updated = v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded {
	s.Total = &v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated {
	s.Total = &v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted {
	s.Total = &v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed {
	s.Total = &v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame {
	s.Total = &v
	return s
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated {
	s.Total = &v
	return s
}

type ListSynchronizationJobsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSynchronizationJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSynchronizationJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSynchronizationJobsResponse) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponse) SetHeaders(v map[string]*string) *ListSynchronizationJobsResponse {
	s.Headers = v
	return s
}

func (s *ListSynchronizationJobsResponse) SetStatusCode(v int32) *ListSynchronizationJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSynchronizationJobsResponse) SetBody(v *ListSynchronizationJobsResponseBody) *ListSynchronizationJobsResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	// 账户展示名，模糊匹配
	//
	// example:
	//
	// name_001
	DisplayNameStartsWith *string `json:"DisplayNameStartsWith,omitempty" xml:"DisplayNameStartsWith,omitempty"`
	// The email address of the user who owns the account.
	//
	// example:
	//
	// user@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the organizational unit.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The mobile number of the user who owns the account.
	//
	// example:
	//
	// 156xxxxxxx
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The country code of the mobile number. For example, the country code of China is 86 without 00 or +.
	//
	// example:
	//
	// 86
	PhoneRegion *string `json:"PhoneRegion,omitempty" xml:"PhoneRegion,omitempty"`
	// The status of the account. Valid values:
	//
	// 	- enabled: The account is enabled.
	//
	// 	- disabled: The account is disabled.
	//
	// example:
	//
	// enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The external ID of the account. The external ID can be used by external data to map the data of the account in IDaaS EIAM.
	//
	// For accounts with the same source type and source ID, each account has a unique external ID.
	//
	// example:
	//
	// id_wovwffm62xifdziem7an7xxxxx
	UserExternalId *string `json:"UserExternalId,omitempty" xml:"UserExternalId,omitempty"`
	// 账户的ID集合
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
	// The source ID of the account.
	//
	// If the account was created in IDaaS, its source ID is the ID of the IDaaS instance. If the account was imported, its source ID is the enterprise ID in the source. For example, if the account was imported from DingTalk, its source ID is the corpId value of the enterprise in DingTalk.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	UserSourceId *string `json:"UserSourceId,omitempty" xml:"UserSourceId,omitempty"`
	// The source type of the account. Valid values:
	//
	// 	- build_in: The account was created in IDaaS.
	//
	// 	- ding_talk: The account was imported from DingTalk.
	//
	// 	- ad: The account was imported from Microsoft Active Directory (AD).
	//
	// 	- ldap: The account was imported from a Lightweight Directory Access Protocol (LDAP) service.
	//
	// example:
	//
	// build_in
	UserSourceType *string `json:"UserSourceType,omitempty" xml:"UserSourceType,omitempty"`
	// 账户名，左模糊匹配
	//
	// example:
	//
	// name_001
	UsernameStartsWith *string `json:"UsernameStartsWith,omitempty" xml:"UsernameStartsWith,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetDisplayNameStartsWith(v string) *ListUsersRequest {
	s.DisplayNameStartsWith = &v
	return s
}

func (s *ListUsersRequest) SetEmail(v string) *ListUsersRequest {
	s.Email = &v
	return s
}

func (s *ListUsersRequest) SetInstanceId(v string) *ListUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUsersRequest) SetOrganizationalUnitId(v string) *ListUsersRequest {
	s.OrganizationalUnitId = &v
	return s
}

func (s *ListUsersRequest) SetPageNumber(v int64) *ListUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v int64) *ListUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListUsersRequest) SetPhoneNumber(v string) *ListUsersRequest {
	s.PhoneNumber = &v
	return s
}

func (s *ListUsersRequest) SetPhoneRegion(v string) *ListUsersRequest {
	s.PhoneRegion = &v
	return s
}

func (s *ListUsersRequest) SetStatus(v string) *ListUsersRequest {
	s.Status = &v
	return s
}

func (s *ListUsersRequest) SetUserExternalId(v string) *ListUsersRequest {
	s.UserExternalId = &v
	return s
}

func (s *ListUsersRequest) SetUserIds(v []*string) *ListUsersRequest {
	s.UserIds = v
	return s
}

func (s *ListUsersRequest) SetUserSourceId(v string) *ListUsersRequest {
	s.UserSourceId = &v
	return s
}

func (s *ListUsersRequest) SetUserSourceType(v string) *ListUsersRequest {
	s.UserSourceType = &v
	return s
}

func (s *ListUsersRequest) SetUsernameStartsWith(v string) *ListUsersRequest {
	s.UsernameStartsWith = &v
	return s
}

type ListUsersResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries in the list.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of data objects of accounts.
	Users []*ListUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetTotalCount(v int64) *ListUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUsersResponseBody) SetUsers(v []*ListUsersResponseBodyUsers) *ListUsersResponseBody {
	s.Users = v
	return s
}

type ListUsersResponseBodyUsers struct {
	// The time when the account expires. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	AccountExpireTime *int64 `json:"AccountExpireTime,omitempty" xml:"AccountExpireTime,omitempty"`
	// The time when the account was created. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the account.
	//
	// example:
	//
	// Test account
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the account.
	//
	// example:
	//
	// display_name001
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address of the user who owns the account.
	//
	// example:
	//
	// user@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Indicates whether the email address has been verified. A value of true indicates that the email address has been verified by the user or has been set to the verified status by the administrator. A value of false indicates that the email address has not been verified.
	//
	// example:
	//
	// true
	EmailVerified *bool `json:"EmailVerified,omitempty" xml:"EmailVerified,omitempty"`
	// The ID of the instance
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the account lock expires. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	LockExpireTime *int64 `json:"LockExpireTime,omitempty" xml:"LockExpireTime,omitempty"`
	// Time When Password Expires
	//
	// example:
	//
	// 1652085686179
	PasswordExpireTime *int64 `json:"PasswordExpireTime,omitempty" xml:"PasswordExpireTime,omitempty"`
	// Indicates whether a password is set.
	//
	// example:
	//
	// false
	PasswordSet *bool `json:"PasswordSet,omitempty" xml:"PasswordSet,omitempty"`
	// The mobile number of the user who owns the account.
	//
	// example:
	//
	// 156xxxxxxx
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Indicates whether the mobile number has been verified. A value of true indicates that the mobile number has been verified by the user or has been set to the verified status by the administrator. A value of false indicates that the mobile number has not been verified.
	//
	// example:
	//
	// true
	PhoneNumberVerified *bool `json:"PhoneNumberVerified,omitempty" xml:"PhoneNumberVerified,omitempty"`
	// The country code of the mobile number. For example, the country code of China is 86 without 00 or +.
	//
	// example:
	//
	// 86
	PhoneRegion *string `json:"PhoneRegion,omitempty" xml:"PhoneRegion,omitempty"`
	// The time when the account was registered. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	RegisterTime *int64 `json:"RegisterTime,omitempty" xml:"RegisterTime,omitempty"`
	// The status of the account. Valid values:
	//
	// 	- enabled: The account is enabled.
	//
	// 	- disabled: The account is disabled.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the account was last updated. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1652085686179
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The external ID of the account. The external ID can be used by external data to map the data of the account in IDaaS EIAM. By default, the external ID is the account ID.
	//
	// For accounts with the same source type and source ID, each account has a unique external ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserExternalId *string `json:"UserExternalId,omitempty" xml:"UserExternalId,omitempty"`
	// The ID of the account.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The source ID of the account.
	//
	// If the account was created in IDaaS, its source ID is the ID of the IDaaS instance. If the account was imported, its source ID is the enterprise ID in the source. For example, if the account was imported from DingTalk, its source ID is the corpId value of the enterprise in DingTalk.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	UserSourceId *string `json:"UserSourceId,omitempty" xml:"UserSourceId,omitempty"`
	// The source type of the account. Valid values:
	//
	// 	- build_in: The account was created in IDaaS.
	//
	// 	- ding_talk: The account was imported from DingTalk.
	//
	// 	- ad: The account was imported from Microsoft Active Directory (AD).
	//
	// 	- ldap: The account was imported from a Lightweight Directory Access Protocol (LDAP) service.
	//
	// example:
	//
	// build_in
	UserSourceType *string `json:"UserSourceType,omitempty" xml:"UserSourceType,omitempty"`
	// The username of the account.
	//
	// example:
	//
	// name001
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsers) SetAccountExpireTime(v int64) *ListUsersResponseBodyUsers {
	s.AccountExpireTime = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetCreateTime(v int64) *ListUsersResponseBodyUsers {
	s.CreateTime = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetDescription(v string) *ListUsersResponseBodyUsers {
	s.Description = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetDisplayName(v string) *ListUsersResponseBodyUsers {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetEmail(v string) *ListUsersResponseBodyUsers {
	s.Email = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetEmailVerified(v bool) *ListUsersResponseBodyUsers {
	s.EmailVerified = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetInstanceId(v string) *ListUsersResponseBodyUsers {
	s.InstanceId = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetLockExpireTime(v int64) *ListUsersResponseBodyUsers {
	s.LockExpireTime = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetPasswordExpireTime(v int64) *ListUsersResponseBodyUsers {
	s.PasswordExpireTime = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetPasswordSet(v bool) *ListUsersResponseBodyUsers {
	s.PasswordSet = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetPhoneNumber(v string) *ListUsersResponseBodyUsers {
	s.PhoneNumber = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetPhoneNumberVerified(v bool) *ListUsersResponseBodyUsers {
	s.PhoneNumberVerified = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetPhoneRegion(v string) *ListUsersResponseBodyUsers {
	s.PhoneRegion = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetRegisterTime(v int64) *ListUsersResponseBodyUsers {
	s.RegisterTime = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetStatus(v string) *ListUsersResponseBodyUsers {
	s.Status = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUpdateTime(v int64) *ListUsersResponseBodyUsers {
	s.UpdateTime = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUserExternalId(v string) *ListUsersResponseBodyUsers {
	s.UserExternalId = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUserId(v string) *ListUsersResponseBodyUsers {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUserSourceId(v string) *ListUsersResponseBodyUsers {
	s.UserSourceId = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUserSourceType(v string) *ListUsersResponseBodyUsers {
	s.UserSourceType = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUsername(v string) *ListUsersResponseBodyUsers {
	s.Username = &v
	return s
}

type ListUsersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponse) GoString() string {
	return s.String()
}

func (s *ListUsersResponse) SetHeaders(v map[string]*string) *ListUsersResponse {
	s.Headers = v
	return s
}

func (s *ListUsersResponse) SetStatusCode(v int32) *ListUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type ListUsersForApplicationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The IDs of the accounts. You can query a maximum of 100 accounts at a time.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s ListUsersForApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersForApplicationRequest) GoString() string {
	return s.String()
}

func (s *ListUsersForApplicationRequest) SetApplicationId(v string) *ListUsersForApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListUsersForApplicationRequest) SetInstanceId(v string) *ListUsersForApplicationRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUsersForApplicationRequest) SetPageNumber(v int64) *ListUsersForApplicationRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersForApplicationRequest) SetPageSize(v int64) *ListUsersForApplicationRequest {
	s.PageSize = &v
	return s
}

func (s *ListUsersForApplicationRequest) SetUserIds(v []*string) *ListUsersForApplicationRequest {
	s.UserIds = v
	return s
}

type ListUsersForApplicationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The IDs of the accounts.
	Users []*ListUsersForApplicationResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListUsersForApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersForApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersForApplicationResponseBody) SetRequestId(v string) *ListUsersForApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersForApplicationResponseBody) SetTotalCount(v int64) *ListUsersForApplicationResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUsersForApplicationResponseBody) SetUsers(v []*ListUsersForApplicationResponseBodyUsers) *ListUsersForApplicationResponseBody {
	s.Users = v
	return s
}

type ListUsersForApplicationResponseBodyUsers struct {
	// The ID of the account.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListUsersForApplicationResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListUsersForApplicationResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersForApplicationResponseBodyUsers) SetUserId(v string) *ListUsersForApplicationResponseBodyUsers {
	s.UserId = &v
	return s
}

type ListUsersForApplicationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUsersForApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUsersForApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersForApplicationResponse) GoString() string {
	return s.String()
}

func (s *ListUsersForApplicationResponse) SetHeaders(v map[string]*string) *ListUsersForApplicationResponse {
	s.Headers = v
	return s
}

func (s *ListUsersForApplicationResponse) SetStatusCode(v int32) *ListUsersForApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersForApplicationResponse) SetBody(v *ListUsersForApplicationResponseBody) *ListUsersForApplicationResponse {
	s.Body = v
	return s
}

type ListUsersForGroupRequest struct {
	// The group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The account IDs. A maximum of 100 accounts can be queried.
	//
	// example:
	//
	// [ou_001]
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s ListUsersForGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersForGroupRequest) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupRequest) SetGroupId(v string) *ListUsersForGroupRequest {
	s.GroupId = &v
	return s
}

func (s *ListUsersForGroupRequest) SetInstanceId(v string) *ListUsersForGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUsersForGroupRequest) SetPageNumber(v int64) *ListUsersForGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersForGroupRequest) SetPageSize(v int64) *ListUsersForGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListUsersForGroupRequest) SetUserIds(v []*string) *ListUsersForGroupRequest {
	s.UserIds = v
	return s
}

type ListUsersForGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned. The maximum number of entries that can be returned per page is specified by PageSize.
	//
	// example:
	//
	// 1000
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The information about accounts.
	Users []*ListUsersForGroupResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListUsersForGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersForGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupResponseBody) SetRequestId(v string) *ListUsersForGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersForGroupResponseBody) SetTotalCount(v int64) *ListUsersForGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUsersForGroupResponseBody) SetUsers(v []*ListUsersForGroupResponseBodyUsers) *ListUsersForGroupResponseBody {
	s.Users = v
	return s
}

type ListUsersForGroupResponseBodyUsers struct {
	GroupMemberRelationSourceId   *string `json:"GroupMemberRelationSourceId,omitempty" xml:"GroupMemberRelationSourceId,omitempty"`
	GroupMemberRelationSourceType *string `json:"GroupMemberRelationSourceType,omitempty" xml:"GroupMemberRelationSourceType,omitempty"`
	// The account ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListUsersForGroupResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListUsersForGroupResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupResponseBodyUsers) SetGroupMemberRelationSourceId(v string) *ListUsersForGroupResponseBodyUsers {
	s.GroupMemberRelationSourceId = &v
	return s
}

func (s *ListUsersForGroupResponseBodyUsers) SetGroupMemberRelationSourceType(v string) *ListUsersForGroupResponseBodyUsers {
	s.GroupMemberRelationSourceType = &v
	return s
}

func (s *ListUsersForGroupResponseBodyUsers) SetUserId(v string) *ListUsersForGroupResponseBodyUsers {
	s.UserId = &v
	return s
}

type ListUsersForGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUsersForGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUsersForGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersForGroupResponse) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupResponse) SetHeaders(v map[string]*string) *ListUsersForGroupResponse {
	s.Headers = v
	return s
}

func (s *ListUsersForGroupResponse) SetStatusCode(v int32) *ListUsersForGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersForGroupResponse) SetBody(v *ListUsersForGroupResponseBody) *ListUsersForGroupResponse {
	s.Body = v
	return s
}

type ObtainApplicationClientSecretRequest struct {
	// The ID of the application whose client key you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The client key ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// sci_k52x2ru63rlkflina5utgkxxxx
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
}

func (s ObtainApplicationClientSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s ObtainApplicationClientSecretRequest) GoString() string {
	return s.String()
}

func (s *ObtainApplicationClientSecretRequest) SetApplicationId(v string) *ObtainApplicationClientSecretRequest {
	s.ApplicationId = &v
	return s
}

func (s *ObtainApplicationClientSecretRequest) SetInstanceId(v string) *ObtainApplicationClientSecretRequest {
	s.InstanceId = &v
	return s
}

func (s *ObtainApplicationClientSecretRequest) SetSecretId(v string) *ObtainApplicationClientSecretRequest {
	s.SecretId = &v
	return s
}

type ObtainApplicationClientSecretResponseBody struct {
	// The information about the client key.
	ApplicationClientSecret *ObtainApplicationClientSecretResponseBodyApplicationClientSecret `json:"ApplicationClientSecret,omitempty" xml:"ApplicationClientSecret,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ObtainApplicationClientSecretResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ObtainApplicationClientSecretResponseBody) GoString() string {
	return s.String()
}

func (s *ObtainApplicationClientSecretResponseBody) SetApplicationClientSecret(v *ObtainApplicationClientSecretResponseBodyApplicationClientSecret) *ObtainApplicationClientSecretResponseBody {
	s.ApplicationClientSecret = v
	return s
}

func (s *ObtainApplicationClientSecretResponseBody) SetRequestId(v string) *ObtainApplicationClientSecretResponseBody {
	s.RequestId = &v
	return s
}

type ObtainApplicationClientSecretResponseBodyApplicationClientSecret struct {
	// The ID of the application whose client key you want to query.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The client ID of the application.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The client key secret of the application.
	//
	// example:
	//
	// CSEHDcHcrUKHw1CuxkJEHPveWRXBGqVqRsxxxx
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the client key was last used. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1649830226000
	LastUsedTime *int64 `json:"LastUsedTime,omitempty" xml:"LastUsedTime,omitempty"`
	// The client key ID of the application.
	//
	// example:
	//
	// sci_k52x2ru63rlkflina5utgkxxxx
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The status of the client key. Valid values:
	//
	// 	- Enabled: The client key is enabled.
	//
	// 	- Disabled: The client key is disabled.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ObtainApplicationClientSecretResponseBodyApplicationClientSecret) String() string {
	return tea.Prettify(s)
}

func (s ObtainApplicationClientSecretResponseBodyApplicationClientSecret) GoString() string {
	return s.String()
}

func (s *ObtainApplicationClientSecretResponseBodyApplicationClientSecret) SetApplicationId(v string) *ObtainApplicationClientSecretResponseBodyApplicationClientSecret {
	s.ApplicationId = &v
	return s
}

func (s *ObtainApplicationClientSecretResponseBodyApplicationClientSecret) SetClientId(v string) *ObtainApplicationClientSecretResponseBodyApplicationClientSecret {
	s.ClientId = &v
	return s
}

func (s *ObtainApplicationClientSecretResponseBodyApplicationClientSecret) SetClientSecret(v string) *ObtainApplicationClientSecretResponseBodyApplicationClientSecret {
	s.ClientSecret = &v
	return s
}

func (s *ObtainApplicationClientSecretResponseBodyApplicationClientSecret) SetInstanceId(v string) *ObtainApplicationClientSecretResponseBodyApplicationClientSecret {
	s.InstanceId = &v
	return s
}

func (s *ObtainApplicationClientSecretResponseBodyApplicationClientSecret) SetLastUsedTime(v int64) *ObtainApplicationClientSecretResponseBodyApplicationClientSecret {
	s.LastUsedTime = &v
	return s
}

func (s *ObtainApplicationClientSecretResponseBodyApplicationClientSecret) SetSecretId(v string) *ObtainApplicationClientSecretResponseBodyApplicationClientSecret {
	s.SecretId = &v
	return s
}

func (s *ObtainApplicationClientSecretResponseBodyApplicationClientSecret) SetStatus(v string) *ObtainApplicationClientSecretResponseBodyApplicationClientSecret {
	s.Status = &v
	return s
}

type ObtainApplicationClientSecretResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ObtainApplicationClientSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ObtainApplicationClientSecretResponse) String() string {
	return tea.Prettify(s)
}

func (s ObtainApplicationClientSecretResponse) GoString() string {
	return s.String()
}

func (s *ObtainApplicationClientSecretResponse) SetHeaders(v map[string]*string) *ObtainApplicationClientSecretResponse {
	s.Headers = v
	return s
}

func (s *ObtainApplicationClientSecretResponse) SetStatusCode(v int32) *ObtainApplicationClientSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *ObtainApplicationClientSecretResponse) SetBody(v *ObtainApplicationClientSecretResponseBody) *ObtainApplicationClientSecretResponse {
	s.Body = v
	return s
}

type ObtainDomainProxyTokenRequest struct {
	// 域名ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// dm_examplexxxxx
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// 域名代理Token ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// pt_examplexxxx
	DomainProxyTokenId *string `json:"DomainProxyTokenId,omitempty" xml:"DomainProxyTokenId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ObtainDomainProxyTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s ObtainDomainProxyTokenRequest) GoString() string {
	return s.String()
}

func (s *ObtainDomainProxyTokenRequest) SetDomainId(v string) *ObtainDomainProxyTokenRequest {
	s.DomainId = &v
	return s
}

func (s *ObtainDomainProxyTokenRequest) SetDomainProxyTokenId(v string) *ObtainDomainProxyTokenRequest {
	s.DomainProxyTokenId = &v
	return s
}

func (s *ObtainDomainProxyTokenRequest) SetInstanceId(v string) *ObtainDomainProxyTokenRequest {
	s.InstanceId = &v
	return s
}

type ObtainDomainProxyTokenResponseBody struct {
	DomainProxyToken *ObtainDomainProxyTokenResponseBodyDomainProxyToken `json:"DomainProxyToken,omitempty" xml:"DomainProxyToken,omitempty" type:"Struct"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ObtainDomainProxyTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ObtainDomainProxyTokenResponseBody) GoString() string {
	return s.String()
}

func (s *ObtainDomainProxyTokenResponseBody) SetDomainProxyToken(v *ObtainDomainProxyTokenResponseBodyDomainProxyToken) *ObtainDomainProxyTokenResponseBody {
	s.DomainProxyToken = v
	return s
}

func (s *ObtainDomainProxyTokenResponseBody) SetRequestId(v string) *ObtainDomainProxyTokenResponseBody {
	s.RequestId = &v
	return s
}

type ObtainDomainProxyTokenResponseBodyDomainProxyToken struct {
	// 域名代理Token创建时间，Unix时间戳格式，单位为毫秒。
	//
	// example:
	//
	// 1649830226000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 域名ID。
	//
	// example:
	//
	// dm_examplexxxx
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// 域名代理Token。
	//
	// example:
	//
	// PTxxxxxxxx
	DomainProxyToken *string `json:"DomainProxyToken,omitempty" xml:"DomainProxyToken,omitempty"`
	// 域名代理Token ID。
	//
	// example:
	//
	// pt_examplexxxx
	DomainProxyTokenId *string `json:"DomainProxyTokenId,omitempty" xml:"DomainProxyTokenId,omitempty"`
	// 实例ID。
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 域名代理Token最近使用时间，Unix时间戳格式，单位为毫秒。
	//
	// example:
	//
	// 1649830226000
	LastUsedTime *int64 `json:"LastUsedTime,omitempty" xml:"LastUsedTime,omitempty"`
	// token状态，枚举类型：(enabled）启用,（disabled）禁用。
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 域名代理Token最近更新时间，Unix时间戳格式，单位为毫秒。
	//
	// example:
	//
	// 1649830226000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ObtainDomainProxyTokenResponseBodyDomainProxyToken) String() string {
	return tea.Prettify(s)
}

func (s ObtainDomainProxyTokenResponseBodyDomainProxyToken) GoString() string {
	return s.String()
}

func (s *ObtainDomainProxyTokenResponseBodyDomainProxyToken) SetCreateTime(v int64) *ObtainDomainProxyTokenResponseBodyDomainProxyToken {
	s.CreateTime = &v
	return s
}

func (s *ObtainDomainProxyTokenResponseBodyDomainProxyToken) SetDomainId(v string) *ObtainDomainProxyTokenResponseBodyDomainProxyToken {
	s.DomainId = &v
	return s
}

func (s *ObtainDomainProxyTokenResponseBodyDomainProxyToken) SetDomainProxyToken(v string) *ObtainDomainProxyTokenResponseBodyDomainProxyToken {
	s.DomainProxyToken = &v
	return s
}

func (s *ObtainDomainProxyTokenResponseBodyDomainProxyToken) SetDomainProxyTokenId(v string) *ObtainDomainProxyTokenResponseBodyDomainProxyToken {
	s.DomainProxyTokenId = &v
	return s
}

func (s *ObtainDomainProxyTokenResponseBodyDomainProxyToken) SetInstanceId(v string) *ObtainDomainProxyTokenResponseBodyDomainProxyToken {
	s.InstanceId = &v
	return s
}

func (s *ObtainDomainProxyTokenResponseBodyDomainProxyToken) SetLastUsedTime(v int64) *ObtainDomainProxyTokenResponseBodyDomainProxyToken {
	s.LastUsedTime = &v
	return s
}

func (s *ObtainDomainProxyTokenResponseBodyDomainProxyToken) SetStatus(v string) *ObtainDomainProxyTokenResponseBodyDomainProxyToken {
	s.Status = &v
	return s
}

func (s *ObtainDomainProxyTokenResponseBodyDomainProxyToken) SetUpdateTime(v int64) *ObtainDomainProxyTokenResponseBodyDomainProxyToken {
	s.UpdateTime = &v
	return s
}

type ObtainDomainProxyTokenResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ObtainDomainProxyTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ObtainDomainProxyTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s ObtainDomainProxyTokenResponse) GoString() string {
	return s.String()
}

func (s *ObtainDomainProxyTokenResponse) SetHeaders(v map[string]*string) *ObtainDomainProxyTokenResponse {
	s.Headers = v
	return s
}

func (s *ObtainDomainProxyTokenResponse) SetStatusCode(v int32) *ObtainDomainProxyTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *ObtainDomainProxyTokenResponse) SetBody(v *ObtainDomainProxyTokenResponseBody) *ObtainDomainProxyTokenResponse {
	s.Body = v
	return s
}

type RemoveUserFromOrganizationalUnitsRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The organization IDs. You can remove an account from a maximum of 100 organizations.
	//
	// This parameter is required.
	OrganizationalUnitIds []*string `json:"OrganizationalUnitIds,omitempty" xml:"OrganizationalUnitIds,omitempty" type:"Repeated"`
	// The account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RemoveUserFromOrganizationalUnitsRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromOrganizationalUnitsRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserFromOrganizationalUnitsRequest) SetInstanceId(v string) *RemoveUserFromOrganizationalUnitsRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveUserFromOrganizationalUnitsRequest) SetOrganizationalUnitIds(v []*string) *RemoveUserFromOrganizationalUnitsRequest {
	s.OrganizationalUnitIds = v
	return s
}

func (s *RemoveUserFromOrganizationalUnitsRequest) SetUserId(v string) *RemoveUserFromOrganizationalUnitsRequest {
	s.UserId = &v
	return s
}

type RemoveUserFromOrganizationalUnitsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveUserFromOrganizationalUnitsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromOrganizationalUnitsResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUserFromOrganizationalUnitsResponseBody) SetRequestId(v string) *RemoveUserFromOrganizationalUnitsResponseBody {
	s.RequestId = &v
	return s
}

type RemoveUserFromOrganizationalUnitsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveUserFromOrganizationalUnitsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveUserFromOrganizationalUnitsResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromOrganizationalUnitsResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserFromOrganizationalUnitsResponse) SetHeaders(v map[string]*string) *RemoveUserFromOrganizationalUnitsResponse {
	s.Headers = v
	return s
}

func (s *RemoveUserFromOrganizationalUnitsResponse) SetStatusCode(v int32) *RemoveUserFromOrganizationalUnitsResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUserFromOrganizationalUnitsResponse) SetBody(v *RemoveUserFromOrganizationalUnitsResponseBody) *RemoveUserFromOrganizationalUnitsResponse {
	s.Body = v
	return s
}

type RemoveUsersFromGroupRequest struct {
	// The group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The account IDs. A maximum of 100 accounts can be removed from a group.
	//
	// This parameter is required.
	//
	// example:
	//
	// [ou_001]
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s RemoveUsersFromGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersFromGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveUsersFromGroupRequest) SetGroupId(v string) *RemoveUsersFromGroupRequest {
	s.GroupId = &v
	return s
}

func (s *RemoveUsersFromGroupRequest) SetInstanceId(v string) *RemoveUsersFromGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveUsersFromGroupRequest) SetUserIds(v []*string) *RemoveUsersFromGroupRequest {
	s.UserIds = v
	return s
}

type RemoveUsersFromGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveUsersFromGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersFromGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUsersFromGroupResponseBody) SetRequestId(v string) *RemoveUsersFromGroupResponseBody {
	s.RequestId = &v
	return s
}

type RemoveUsersFromGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveUsersFromGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveUsersFromGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersFromGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveUsersFromGroupResponse) SetHeaders(v map[string]*string) *RemoveUsersFromGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveUsersFromGroupResponse) SetStatusCode(v int32) *RemoveUsersFromGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUsersFromGroupResponse) SetBody(v *RemoveUsersFromGroupResponseBody) *RemoveUsersFromGroupResponse {
	s.Body = v
	return s
}

type RevokeApplicationFromGroupsRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The group IDs. You can specify up to 100 group IDs at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_miu8e4t4d7i4u7uwezgr54xxxx
	GroupIds []*string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RevokeApplicationFromGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeApplicationFromGroupsRequest) GoString() string {
	return s.String()
}

func (s *RevokeApplicationFromGroupsRequest) SetApplicationId(v string) *RevokeApplicationFromGroupsRequest {
	s.ApplicationId = &v
	return s
}

func (s *RevokeApplicationFromGroupsRequest) SetGroupIds(v []*string) *RevokeApplicationFromGroupsRequest {
	s.GroupIds = v
	return s
}

func (s *RevokeApplicationFromGroupsRequest) SetInstanceId(v string) *RevokeApplicationFromGroupsRequest {
	s.InstanceId = &v
	return s
}

type RevokeApplicationFromGroupsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeApplicationFromGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeApplicationFromGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeApplicationFromGroupsResponseBody) SetRequestId(v string) *RevokeApplicationFromGroupsResponseBody {
	s.RequestId = &v
	return s
}

type RevokeApplicationFromGroupsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeApplicationFromGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeApplicationFromGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeApplicationFromGroupsResponse) GoString() string {
	return s.String()
}

func (s *RevokeApplicationFromGroupsResponse) SetHeaders(v map[string]*string) *RevokeApplicationFromGroupsResponse {
	s.Headers = v
	return s
}

func (s *RevokeApplicationFromGroupsResponse) SetStatusCode(v int32) *RevokeApplicationFromGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeApplicationFromGroupsResponse) SetBody(v *RevokeApplicationFromGroupsResponseBody) *RevokeApplicationFromGroupsResponse {
	s.Body = v
	return s
}

type RevokeApplicationFromOrganizationalUnitsRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IDs of the organizations. You can revoke the access permissions from a maximum of 100 organizations at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitIds []*string `json:"OrganizationalUnitIds,omitempty" xml:"OrganizationalUnitIds,omitempty" type:"Repeated"`
}

func (s RevokeApplicationFromOrganizationalUnitsRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeApplicationFromOrganizationalUnitsRequest) GoString() string {
	return s.String()
}

func (s *RevokeApplicationFromOrganizationalUnitsRequest) SetApplicationId(v string) *RevokeApplicationFromOrganizationalUnitsRequest {
	s.ApplicationId = &v
	return s
}

func (s *RevokeApplicationFromOrganizationalUnitsRequest) SetInstanceId(v string) *RevokeApplicationFromOrganizationalUnitsRequest {
	s.InstanceId = &v
	return s
}

func (s *RevokeApplicationFromOrganizationalUnitsRequest) SetOrganizationalUnitIds(v []*string) *RevokeApplicationFromOrganizationalUnitsRequest {
	s.OrganizationalUnitIds = v
	return s
}

type RevokeApplicationFromOrganizationalUnitsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeApplicationFromOrganizationalUnitsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeApplicationFromOrganizationalUnitsResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeApplicationFromOrganizationalUnitsResponseBody) SetRequestId(v string) *RevokeApplicationFromOrganizationalUnitsResponseBody {
	s.RequestId = &v
	return s
}

type RevokeApplicationFromOrganizationalUnitsResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeApplicationFromOrganizationalUnitsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeApplicationFromOrganizationalUnitsResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeApplicationFromOrganizationalUnitsResponse) GoString() string {
	return s.String()
}

func (s *RevokeApplicationFromOrganizationalUnitsResponse) SetHeaders(v map[string]*string) *RevokeApplicationFromOrganizationalUnitsResponse {
	s.Headers = v
	return s
}

func (s *RevokeApplicationFromOrganizationalUnitsResponse) SetStatusCode(v int32) *RevokeApplicationFromOrganizationalUnitsResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeApplicationFromOrganizationalUnitsResponse) SetBody(v *RevokeApplicationFromOrganizationalUnitsResponseBody) *RevokeApplicationFromOrganizationalUnitsResponse {
	s.Body = v
	return s
}

type RevokeApplicationFromUsersRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IDs of the accounts. You can revoke the access permissions from a maximum of 100 accounts at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s RevokeApplicationFromUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeApplicationFromUsersRequest) GoString() string {
	return s.String()
}

func (s *RevokeApplicationFromUsersRequest) SetApplicationId(v string) *RevokeApplicationFromUsersRequest {
	s.ApplicationId = &v
	return s
}

func (s *RevokeApplicationFromUsersRequest) SetInstanceId(v string) *RevokeApplicationFromUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *RevokeApplicationFromUsersRequest) SetUserIds(v []*string) *RevokeApplicationFromUsersRequest {
	s.UserIds = v
	return s
}

type RevokeApplicationFromUsersResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeApplicationFromUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeApplicationFromUsersResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeApplicationFromUsersResponseBody) SetRequestId(v string) *RevokeApplicationFromUsersResponseBody {
	s.RequestId = &v
	return s
}

type RevokeApplicationFromUsersResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeApplicationFromUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeApplicationFromUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeApplicationFromUsersResponse) GoString() string {
	return s.String()
}

func (s *RevokeApplicationFromUsersResponse) SetHeaders(v map[string]*string) *RevokeApplicationFromUsersResponse {
	s.Headers = v
	return s
}

func (s *RevokeApplicationFromUsersResponse) SetStatusCode(v int32) *RevokeApplicationFromUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeApplicationFromUsersResponse) SetBody(v *RevokeApplicationFromUsersResponseBody) *RevokeApplicationFromUsersResponse {
	s.Body = v
	return s
}

type RunSynchronizationJobRequest struct {
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 同步目标ID
	//
	// This parameter is required.
	//
	// example:
	//
	// idp_my664lwkhpicbyzirog3ngxxxxx
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// 同步目标类型
	//
	// This parameter is required.
	//
	// example:
	//
	// identity_provider
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s RunSynchronizationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s RunSynchronizationJobRequest) GoString() string {
	return s.String()
}

func (s *RunSynchronizationJobRequest) SetInstanceId(v string) *RunSynchronizationJobRequest {
	s.InstanceId = &v
	return s
}

func (s *RunSynchronizationJobRequest) SetTargetId(v string) *RunSynchronizationJobRequest {
	s.TargetId = &v
	return s
}

func (s *RunSynchronizationJobRequest) SetTargetType(v string) *RunSynchronizationJobRequest {
	s.TargetType = &v
	return s
}

type RunSynchronizationJobResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// sync_0000347vjovtcf41li0fgsd98gn24q9nj9xxxxx
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
}

func (s RunSynchronizationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunSynchronizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *RunSynchronizationJobResponseBody) SetRequestId(v string) *RunSynchronizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunSynchronizationJobResponseBody) SetSynchronizationJobId(v string) *RunSynchronizationJobResponseBody {
	s.SynchronizationJobId = &v
	return s
}

type RunSynchronizationJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunSynchronizationJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunSynchronizationJobResponse) String() string {
	return tea.Prettify(s)
}

func (s RunSynchronizationJobResponse) GoString() string {
	return s.String()
}

func (s *RunSynchronizationJobResponse) SetHeaders(v map[string]*string) *RunSynchronizationJobResponse {
	s.Headers = v
	return s
}

func (s *RunSynchronizationJobResponse) SetStatusCode(v int32) *RunSynchronizationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *RunSynchronizationJobResponse) SetBody(v *RunSynchronizationJobResponseBody) *RunSynchronizationJobResponse {
	s.Body = v
	return s
}

type SetApplicationGrantScopeRequest struct {
	// The ID of the application that you want to configure.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The permissions of the Developer API feature.
	GrantScopes []*string `json:"GrantScopes,omitempty" xml:"GrantScopes,omitempty" type:"Repeated"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s SetApplicationGrantScopeRequest) String() string {
	return tea.Prettify(s)
}

func (s SetApplicationGrantScopeRequest) GoString() string {
	return s.String()
}

func (s *SetApplicationGrantScopeRequest) SetApplicationId(v string) *SetApplicationGrantScopeRequest {
	s.ApplicationId = &v
	return s
}

func (s *SetApplicationGrantScopeRequest) SetGrantScopes(v []*string) *SetApplicationGrantScopeRequest {
	s.GrantScopes = v
	return s
}

func (s *SetApplicationGrantScopeRequest) SetInstanceId(v string) *SetApplicationGrantScopeRequest {
	s.InstanceId = &v
	return s
}

type SetApplicationGrantScopeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetApplicationGrantScopeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetApplicationGrantScopeResponseBody) GoString() string {
	return s.String()
}

func (s *SetApplicationGrantScopeResponseBody) SetRequestId(v string) *SetApplicationGrantScopeResponseBody {
	s.RequestId = &v
	return s
}

type SetApplicationGrantScopeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetApplicationGrantScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetApplicationGrantScopeResponse) String() string {
	return tea.Prettify(s)
}

func (s SetApplicationGrantScopeResponse) GoString() string {
	return s.String()
}

func (s *SetApplicationGrantScopeResponse) SetHeaders(v map[string]*string) *SetApplicationGrantScopeResponse {
	s.Headers = v
	return s
}

func (s *SetApplicationGrantScopeResponse) SetStatusCode(v int32) *SetApplicationGrantScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetApplicationGrantScopeResponse) SetBody(v *SetApplicationGrantScopeResponseBody) *SetApplicationGrantScopeResponse {
	s.Body = v
	return s
}

type SetApplicationProvisioningConfigRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The configuration of event callback synchronization. This parameter is required when the ProvisionProtocolType parameter is set to idaas_callback.
	CallbackProvisioningConfig *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig `json:"CallbackProvisioningConfig,omitempty" xml:"CallbackProvisioningConfig,omitempty" type:"Struct"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to synchronize the password in IDaaS user event callbacks. Valid values:
	//
	// 	- true: synchronize the password.
	//
	// 	- false: do not synchronize the password.
	//
	// example:
	//
	// true
	ProvisionPassword *bool `json:"ProvisionPassword,omitempty" xml:"ProvisionPassword,omitempty"`
	// The synchronization protocol type of the application. Valid values:
	//
	// 	- idaas_callback: custom event callback protocol of IDaaS.
	//
	// 	- scim2: System for Cross-domain Identity Management (SCIM) protocol.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_callback
	ProvisionProtocolType *string `json:"ProvisionProtocolType,omitempty" xml:"ProvisionProtocolType,omitempty"`
	// The configuration of SCIM-based IDaaS synchronization. This parameter is required when the ProvisionProtocolType parameter is set to scim2.
	ScimProvisioningConfig *SetApplicationProvisioningConfigRequestScimProvisioningConfig `json:"ScimProvisioningConfig,omitempty" xml:"ScimProvisioningConfig,omitempty" type:"Struct"`
}

func (s SetApplicationProvisioningConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetApplicationProvisioningConfigRequest) GoString() string {
	return s.String()
}

func (s *SetApplicationProvisioningConfigRequest) SetApplicationId(v string) *SetApplicationProvisioningConfigRequest {
	s.ApplicationId = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequest) SetCallbackProvisioningConfig(v *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig) *SetApplicationProvisioningConfigRequest {
	s.CallbackProvisioningConfig = v
	return s
}

func (s *SetApplicationProvisioningConfigRequest) SetInstanceId(v string) *SetApplicationProvisioningConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequest) SetProvisionPassword(v bool) *SetApplicationProvisioningConfigRequest {
	s.ProvisionPassword = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequest) SetProvisionProtocolType(v string) *SetApplicationProvisioningConfigRequest {
	s.ProvisionProtocolType = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequest) SetScimProvisioningConfig(v *SetApplicationProvisioningConfigRequestScimProvisioningConfig) *SetApplicationProvisioningConfigRequest {
	s.ScimProvisioningConfig = v
	return s
}

type SetApplicationProvisioningConfigRequestCallbackProvisioningConfig struct {
	// The URL that the application uses to receive IDaaS event callbacks.
	//
	// example:
	//
	// https://example.com/event/callback
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// The symmetric key for IDaaS event callbacks. The key is an AES-256 encryption key in the HEX format.
	//
	// example:
	//
	// ad3b248**************************b3561a73d7
	EncryptKey *string `json:"EncryptKey,omitempty" xml:"EncryptKey,omitempty"`
	// Specifies whether to encrypt IDaaS event callback messages. Valid values:
	//
	// 	- true: encrypt the messages.
	//
	// 	- false: transmit the messages in plaintext.
	//
	// example:
	//
	// true
	EncryptRequired *bool `json:"EncryptRequired,omitempty" xml:"EncryptRequired,omitempty"`
	// The list of types of IDaaS event callback messages that are supported by the listener.
	ListenEventScopes []*string `json:"ListenEventScopes,omitempty" xml:"ListenEventScopes,omitempty" type:"Repeated"`
}

func (s SetApplicationProvisioningConfigRequestCallbackProvisioningConfig) String() string {
	return tea.Prettify(s)
}

func (s SetApplicationProvisioningConfigRequestCallbackProvisioningConfig) GoString() string {
	return s.String()
}

func (s *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig) SetCallbackUrl(v string) *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig {
	s.CallbackUrl = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig) SetEncryptKey(v string) *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig {
	s.EncryptKey = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig) SetEncryptRequired(v bool) *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig {
	s.EncryptRequired = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig) SetListenEventScopes(v []*string) *SetApplicationProvisioningConfigRequestCallbackProvisioningConfig {
	s.ListenEventScopes = v
	return s
}

type SetApplicationProvisioningConfigRequestScimProvisioningConfig struct {
	// The configuration parameters related to SCIM-based synchronization.
	AuthnConfiguration *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfiguration `json:"AuthnConfiguration,omitempty" xml:"AuthnConfiguration,omitempty" type:"Struct"`
	// The full synchronization scope of the SCIM protocol. Valid value:
	//
	// 	- urn:alibaba:idaas:app:scim:User:PUSH: full account data synchronization.
	FullPushScopes []*string `json:"FullPushScopes,omitempty" xml:"FullPushScopes,omitempty" type:"Repeated"`
	// The resource operations of the SCIM protocol. Valid values:
	//
	// 	- urn:alibaba:idaas:app:scim:User:CREATE: account creation.
	//
	// 	- urn:alibaba:idaas:app:scim:User:UPDATE: account update.
	//
	// 	- urn:alibaba:idaas:app:scim:User:DELETE: account deletion.
	ProvisioningActions []*string `json:"ProvisioningActions,omitempty" xml:"ProvisioningActions,omitempty" type:"Repeated"`
	// The base URL that the application uses to receive the SCIM protocol for IDaaS synchronization.
	//
	// example:
	//
	// https://example.com/scim
	ScimBaseUrl *string `json:"ScimBaseUrl,omitempty" xml:"ScimBaseUrl,omitempty"`
}

func (s SetApplicationProvisioningConfigRequestScimProvisioningConfig) String() string {
	return tea.Prettify(s)
}

func (s SetApplicationProvisioningConfigRequestScimProvisioningConfig) GoString() string {
	return s.String()
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfig) SetAuthnConfiguration(v *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfiguration) *SetApplicationProvisioningConfigRequestScimProvisioningConfig {
	s.AuthnConfiguration = v
	return s
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfig) SetFullPushScopes(v []*string) *SetApplicationProvisioningConfigRequestScimProvisioningConfig {
	s.FullPushScopes = v
	return s
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfig) SetProvisioningActions(v []*string) *SetApplicationProvisioningConfigRequestScimProvisioningConfig {
	s.ProvisioningActions = v
	return s
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfig) SetScimBaseUrl(v string) *SetApplicationProvisioningConfigRequestScimProvisioningConfig {
	s.ScimBaseUrl = &v
	return s
}

type SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfiguration struct {
	// The authentication mode of the SCIM protocol. Valid value:
	//
	// 	- oauth2: OAuth2.0 mode.
	//
	// example:
	//
	// oauth2
	AuthnMode *string `json:"AuthnMode,omitempty" xml:"AuthnMode,omitempty"`
	// The configuration parameters related to authorization.
	//
	// 	- If the GrantType parameter is set to client_credentials, you can set the configuration parameters ClientId, ClientSecret, and AuthnMethod.
	//
	// 	- If the GrantType parameter is set to bearer_token, you can set the configuration parameter AccessToken.
	AuthnParam *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam `json:"AuthnParam,omitempty" xml:"AuthnParam,omitempty" type:"Struct"`
	// The grant type of the SCIM protocol. Valid values:
	//
	// 	- client_credentials: client mode.
	//
	// 	- bearer_token: key mode.
	//
	// example:
	//
	// bearer_token
	GrantType *string `json:"GrantType,omitempty" xml:"GrantType,omitempty"`
}

func (s SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfiguration) String() string {
	return tea.Prettify(s)
}

func (s SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfiguration) GoString() string {
	return s.String()
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfiguration) SetAuthnMode(v string) *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfiguration {
	s.AuthnMode = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfiguration) SetAuthnParam(v *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam) *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfiguration {
	s.AuthnParam = v
	return s
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfiguration) SetGrantType(v string) *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfiguration {
	s.GrantType = &v
	return s
}

type SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam struct {
	// The access token. If the GrantType parameter is set to bearer_token, you can set this parameter.
	//
	// example:
	//
	// k52x2ru63rlkflina5utgkxxxx
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// The authentication mode of the SCIM protocol. Valid values:
	//
	// 	- client_secret_basic: The client secret is passed in the request header.
	//
	// 	- client_secret_post: The client secret is passed in the request body.
	//
	// example:
	//
	// client_secret_basic
	AuthnMethod *string `json:"AuthnMethod,omitempty" xml:"AuthnMethod,omitempty"`
	// The client ID of the application.
	//
	// example:
	//
	// mkv7rgt4d7i4u7zqtzev2mxxxx
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The client secret of the application.
	//
	// example:
	//
	// CSEHDcHcrUKHw1CuxkJEHPveWRXBGqVqRsxxxx
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	// The token endpoint.
	//
	// example:
	//
	// https://www.example.com/oauth/token
	TokenEndpoint *string `json:"TokenEndpoint,omitempty" xml:"TokenEndpoint,omitempty"`
}

func (s SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam) String() string {
	return tea.Prettify(s)
}

func (s SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam) GoString() string {
	return s.String()
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam) SetAccessToken(v string) *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam {
	s.AccessToken = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam) SetAuthnMethod(v string) *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam {
	s.AuthnMethod = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam) SetClientId(v string) *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam {
	s.ClientId = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam) SetClientSecret(v string) *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam {
	s.ClientSecret = &v
	return s
}

func (s *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam) SetTokenEndpoint(v string) *SetApplicationProvisioningConfigRequestScimProvisioningConfigAuthnConfigurationAuthnParam {
	s.TokenEndpoint = &v
	return s
}

type SetApplicationProvisioningConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetApplicationProvisioningConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetApplicationProvisioningConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetApplicationProvisioningConfigResponseBody) SetRequestId(v string) *SetApplicationProvisioningConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetApplicationProvisioningConfigResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetApplicationProvisioningConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetApplicationProvisioningConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetApplicationProvisioningConfigResponse) GoString() string {
	return s.String()
}

func (s *SetApplicationProvisioningConfigResponse) SetHeaders(v map[string]*string) *SetApplicationProvisioningConfigResponse {
	s.Headers = v
	return s
}

func (s *SetApplicationProvisioningConfigResponse) SetStatusCode(v int32) *SetApplicationProvisioningConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetApplicationProvisioningConfigResponse) SetBody(v *SetApplicationProvisioningConfigResponseBody) *SetApplicationProvisioningConfigResponse {
	s.Body = v
	return s
}

type SetApplicationProvisioningScopeRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// 授权同步出的组列表
	GroupIds []*string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of organizational units that are authorized for account synchronization.
	OrganizationalUnitIds []*string `json:"OrganizationalUnitIds,omitempty" xml:"OrganizationalUnitIds,omitempty" type:"Repeated"`
}

func (s SetApplicationProvisioningScopeRequest) String() string {
	return tea.Prettify(s)
}

func (s SetApplicationProvisioningScopeRequest) GoString() string {
	return s.String()
}

func (s *SetApplicationProvisioningScopeRequest) SetApplicationId(v string) *SetApplicationProvisioningScopeRequest {
	s.ApplicationId = &v
	return s
}

func (s *SetApplicationProvisioningScopeRequest) SetGroupIds(v []*string) *SetApplicationProvisioningScopeRequest {
	s.GroupIds = v
	return s
}

func (s *SetApplicationProvisioningScopeRequest) SetInstanceId(v string) *SetApplicationProvisioningScopeRequest {
	s.InstanceId = &v
	return s
}

func (s *SetApplicationProvisioningScopeRequest) SetOrganizationalUnitIds(v []*string) *SetApplicationProvisioningScopeRequest {
	s.OrganizationalUnitIds = v
	return s
}

type SetApplicationProvisioningScopeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetApplicationProvisioningScopeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetApplicationProvisioningScopeResponseBody) GoString() string {
	return s.String()
}

func (s *SetApplicationProvisioningScopeResponseBody) SetRequestId(v string) *SetApplicationProvisioningScopeResponseBody {
	s.RequestId = &v
	return s
}

type SetApplicationProvisioningScopeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetApplicationProvisioningScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetApplicationProvisioningScopeResponse) String() string {
	return tea.Prettify(s)
}

func (s SetApplicationProvisioningScopeResponse) GoString() string {
	return s.String()
}

func (s *SetApplicationProvisioningScopeResponse) SetHeaders(v map[string]*string) *SetApplicationProvisioningScopeResponse {
	s.Headers = v
	return s
}

func (s *SetApplicationProvisioningScopeResponse) SetStatusCode(v int32) *SetApplicationProvisioningScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetApplicationProvisioningScopeResponse) SetBody(v *SetApplicationProvisioningScopeResponseBody) *SetApplicationProvisioningScopeResponse {
	s.Body = v
	return s
}

type SetApplicationSsoConfigRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The initial SSO method. Valid values:
	//
	// 	- only_app_init_sso: Only application-initiated SSO is allowed. This method is selected by default when the SSO protocol of the application is an OIDC protocol. If this method is selected when the SSO protocol of the application is SAML, the InitLoginUrl parameter is required.
	//
	// 	- idaas_or_app_init_sso: IDaaS-initiated SSO and application-initiated SSO are allowed. This method is selected by default when the SSO protocol of the application is SAML. If this method is selected when the SSO protocol of the application is an OIDC protocol, the InitLoginUrl parameter is required.
	//
	// example:
	//
	// only_app_init_sso
	InitLoginType *string `json:"InitLoginType,omitempty" xml:"InitLoginType,omitempty"`
	// The initial webhook URL of SSO. This parameter is required when the SSO protocol of the application is an OIDC protocol and the InitLoginType parameters is set to idaas_or_app_init_sso or when the SSO protocol of the application is SAML and the InitLoginType parameter is set to only_app_init_sso.
	//
	// example:
	//
	// http://127.0.0.1:8000/start_login?enterprise_code=ABCDEF
	InitLoginUrl *string `json:"InitLoginUrl,omitempty" xml:"InitLoginUrl,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The Open ID Connect (OIDC)-based SSO configuration attributes of the application.
	OidcSsoConfig *SetApplicationSsoConfigRequestOidcSsoConfig `json:"OidcSsoConfig,omitempty" xml:"OidcSsoConfig,omitempty" type:"Struct"`
	// The Security Assertion Markup Language (SAML)-based SSO configuration attributes of the application.
	SamlSsoConfig *SetApplicationSsoConfigRequestSamlSsoConfig `json:"SamlSsoConfig,omitempty" xml:"SamlSsoConfig,omitempty" type:"Struct"`
}

func (s SetApplicationSsoConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetApplicationSsoConfigRequest) GoString() string {
	return s.String()
}

func (s *SetApplicationSsoConfigRequest) SetApplicationId(v string) *SetApplicationSsoConfigRequest {
	s.ApplicationId = &v
	return s
}

func (s *SetApplicationSsoConfigRequest) SetInitLoginType(v string) *SetApplicationSsoConfigRequest {
	s.InitLoginType = &v
	return s
}

func (s *SetApplicationSsoConfigRequest) SetInitLoginUrl(v string) *SetApplicationSsoConfigRequest {
	s.InitLoginUrl = &v
	return s
}

func (s *SetApplicationSsoConfigRequest) SetInstanceId(v string) *SetApplicationSsoConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *SetApplicationSsoConfigRequest) SetOidcSsoConfig(v *SetApplicationSsoConfigRequestOidcSsoConfig) *SetApplicationSsoConfigRequest {
	s.OidcSsoConfig = v
	return s
}

func (s *SetApplicationSsoConfigRequest) SetSamlSsoConfig(v *SetApplicationSsoConfigRequestSamlSsoConfig) *SetApplicationSsoConfigRequest {
	s.SamlSsoConfig = v
	return s
}

type SetApplicationSsoConfigRequestOidcSsoConfig struct {
	// The validity period of the issued access token. Unit: seconds. Default value: 1200.
	//
	// example:
	//
	// 1200
	AccessTokenEffectiveTime *int64 `json:"AccessTokenEffectiveTime,omitempty" xml:"AccessTokenEffectiveTime,omitempty"`
	// The validity period of the issued code. Unit: seconds. Default value: 60.
	//
	// example:
	//
	// 60
	CodeEffectiveTime *int64 `json:"CodeEffectiveTime,omitempty" xml:"CodeEffectiveTime,omitempty"`
	// The custom claims that are returned for the ID token.
	CustomClaims []*SetApplicationSsoConfigRequestOidcSsoConfigCustomClaims `json:"CustomClaims,omitempty" xml:"CustomClaims,omitempty" type:"Repeated"`
	// The scopes of user attributes that can be returned for the UserInfo endpoint or ID token.
	//
	// example:
	//
	// profile，email
	GrantScopes []*string `json:"GrantScopes,omitempty" xml:"GrantScopes,omitempty" type:"Repeated"`
	// The list of grant types that are supported for OIDC protocols.
	//
	// example:
	//
	// authorization_code
	GrantTypes []*string `json:"GrantTypes,omitempty" xml:"GrantTypes,omitempty" type:"Repeated"`
	// The validity period of the issued ID token. Unit: seconds. Default value: 300.
	//
	// example:
	//
	// 300
	IdTokenEffectiveTime *int64 `json:"IdTokenEffectiveTime,omitempty" xml:"IdTokenEffectiveTime,omitempty"`
	// The ID of the identity authentication source in password mode. Specify this parameter only when the value of the GrantTypes parameter includes the password mode.
	//
	// example:
	//
	// ia_password
	PasswordAuthenticationSourceId *string `json:"PasswordAuthenticationSourceId,omitempty" xml:"PasswordAuthenticationSourceId,omitempty"`
	// Specifies whether time-based one-time password (TOTP) authentication is required in password mode. Specify this parameter only when the value of the GrantTypes parameter includes the password mode.
	//
	// example:
	//
	// true
	PasswordTotpMfaRequired *bool `json:"PasswordTotpMfaRequired,omitempty" xml:"PasswordTotpMfaRequired,omitempty"`
	// The algorithms that are used to calculate the code challenge for PKCE.
	//
	// example:
	//
	// S256
	PkceChallengeMethods []*string `json:"PkceChallengeMethods,omitempty" xml:"PkceChallengeMethods,omitempty" type:"Repeated"`
	// Specifies whether the SSO of the application requires Proof Key for Code Exchange (PKCE) (RFC 7636).
	//
	// example:
	//
	// true
	PkceRequired *bool `json:"PkceRequired,omitempty" xml:"PkceRequired,omitempty"`
	// The list of logout redirect URIs that are supported by the application.
	PostLogoutRedirectUris []*string `json:"PostLogoutRedirectUris,omitempty" xml:"PostLogoutRedirectUris,omitempty" type:"Repeated"`
	// The list of redirect URIs that are supported by the application.
	RedirectUris []*string `json:"RedirectUris,omitempty" xml:"RedirectUris,omitempty" type:"Repeated"`
	// The validity period of the issued refresh token. Unit: seconds. Default value: 86400.
	//
	// example:
	//
	// 86400
	RefreshTokenEffective *int64 `json:"RefreshTokenEffective,omitempty" xml:"RefreshTokenEffective,omitempty"`
	// The response types that are supported by the application. Specify this parameter when the value of the GrantTypes parameter includes the implicit mode.
	//
	// example:
	//
	// token id_token
	ResponseTypes []*string `json:"ResponseTypes,omitempty" xml:"ResponseTypes,omitempty" type:"Repeated"`
	// The custom expression that is used to generate the subject ID returned for the ID token.
	//
	// example:
	//
	// user.userid
	SubjectIdExpression *string `json:"SubjectIdExpression,omitempty" xml:"SubjectIdExpression,omitempty"`
}

func (s SetApplicationSsoConfigRequestOidcSsoConfig) String() string {
	return tea.Prettify(s)
}

func (s SetApplicationSsoConfigRequestOidcSsoConfig) GoString() string {
	return s.String()
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetAccessTokenEffectiveTime(v int64) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.AccessTokenEffectiveTime = &v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetCodeEffectiveTime(v int64) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.CodeEffectiveTime = &v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetCustomClaims(v []*SetApplicationSsoConfigRequestOidcSsoConfigCustomClaims) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.CustomClaims = v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetGrantScopes(v []*string) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.GrantScopes = v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetGrantTypes(v []*string) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.GrantTypes = v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetIdTokenEffectiveTime(v int64) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.IdTokenEffectiveTime = &v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetPasswordAuthenticationSourceId(v string) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.PasswordAuthenticationSourceId = &v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetPasswordTotpMfaRequired(v bool) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.PasswordTotpMfaRequired = &v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetPkceChallengeMethods(v []*string) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.PkceChallengeMethods = v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetPkceRequired(v bool) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.PkceRequired = &v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetPostLogoutRedirectUris(v []*string) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.PostLogoutRedirectUris = v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetRedirectUris(v []*string) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.RedirectUris = v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetRefreshTokenEffective(v int64) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.RefreshTokenEffective = &v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetResponseTypes(v []*string) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.ResponseTypes = v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfig) SetSubjectIdExpression(v string) *SetApplicationSsoConfigRequestOidcSsoConfig {
	s.SubjectIdExpression = &v
	return s
}

type SetApplicationSsoConfigRequestOidcSsoConfigCustomClaims struct {
	// The claim name.
	//
	// example:
	//
	// "Role"
	ClaimName *string `json:"ClaimName,omitempty" xml:"ClaimName,omitempty"`
	// The expression that is used to generate the value of the claim.
	//
	// example:
	//
	// user.dict.applicationRole
	ClaimValueExpression *string `json:"ClaimValueExpression,omitempty" xml:"ClaimValueExpression,omitempty"`
}

func (s SetApplicationSsoConfigRequestOidcSsoConfigCustomClaims) String() string {
	return tea.Prettify(s)
}

func (s SetApplicationSsoConfigRequestOidcSsoConfigCustomClaims) GoString() string {
	return s.String()
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfigCustomClaims) SetClaimName(v string) *SetApplicationSsoConfigRequestOidcSsoConfigCustomClaims {
	s.ClaimName = &v
	return s
}

func (s *SetApplicationSsoConfigRequestOidcSsoConfigCustomClaims) SetClaimValueExpression(v string) *SetApplicationSsoConfigRequestOidcSsoConfigCustomClaims {
	s.ClaimValueExpression = &v
	return s
}

type SetApplicationSsoConfigRequestSamlSsoConfig struct {
	// assertion是否签名
	//
	// example:
	//
	// true
	AssertionSigned *bool `json:"AssertionSigned,omitempty" xml:"AssertionSigned,omitempty"`
	// The additional user attributes in the SAML assertion.
	AttributeStatements []*SetApplicationSsoConfigRequestSamlSsoConfigAttributeStatements `json:"AttributeStatements,omitempty" xml:"AttributeStatements,omitempty" type:"Repeated"`
	// The default value of the RelayState attribute. If the SSO request is initiated in EIAM, the RelayState attribute in the SAML response is set to this default value.
	//
	// example:
	//
	// https://home.console.aliyun.com
	DefaultRelayState *string `json:"DefaultRelayState,omitempty" xml:"DefaultRelayState,omitempty"`
	// The Format attribute of the NameID element in the SAML assertion. Valid values:
	//
	// 	- urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified: No format is specified. How to resolve the NameID element depends on the application.
	//
	// 	- urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress: The NameID element must be an email address.
	//
	// 	- urn:oasis:names:tc:SAML:2.0:nameid-format:persistent: The NameID element must be persistent.
	//
	// 	- urn:oasis:names:tc:SAML:2.0:nameid-format:transient: The NameID element must be transient.
	//
	// example:
	//
	// urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified
	NameIdFormat *string `json:"NameIdFormat,omitempty" xml:"NameIdFormat,omitempty"`
	// The expression that is used to generate the value of NameID in the SAML assertion.
	//
	// example:
	//
	// user.email
	NameIdValueExpression *string `json:"NameIdValueExpression,omitempty" xml:"NameIdValueExpression,omitempty"`
	// response是否签名
	//
	// example:
	//
	// true
	ResponseSigned *bool `json:"ResponseSigned,omitempty" xml:"ResponseSigned,omitempty"`
	// The algorithm that is used to calculate the signature for the SAML assertion.
	//
	// example:
	//
	// RSA-SHA256
	SignatureAlgorithm *string `json:"SignatureAlgorithm,omitempty" xml:"SignatureAlgorithm,omitempty"`
	// The entity ID of the application in SAML. The application assumes the role of service provider.
	//
	// example:
	//
	// urn:alibaba:cloudcomputing
	SpEntityId *string `json:"SpEntityId,omitempty" xml:"SpEntityId,omitempty"`
	// The Assertion Consumer Service (ACS) URL of the application in SAML. The application assumes the role of service provider.
	//
	// example:
	//
	// https://signin.aliyun.com/saml-role/sso
	SpSsoAcsUrl *string `json:"SpSsoAcsUrl,omitempty" xml:"SpSsoAcsUrl,omitempty"`
}

func (s SetApplicationSsoConfigRequestSamlSsoConfig) String() string {
	return tea.Prettify(s)
}

func (s SetApplicationSsoConfigRequestSamlSsoConfig) GoString() string {
	return s.String()
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) SetAssertionSigned(v bool) *SetApplicationSsoConfigRequestSamlSsoConfig {
	s.AssertionSigned = &v
	return s
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) SetAttributeStatements(v []*SetApplicationSsoConfigRequestSamlSsoConfigAttributeStatements) *SetApplicationSsoConfigRequestSamlSsoConfig {
	s.AttributeStatements = v
	return s
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) SetDefaultRelayState(v string) *SetApplicationSsoConfigRequestSamlSsoConfig {
	s.DefaultRelayState = &v
	return s
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) SetNameIdFormat(v string) *SetApplicationSsoConfigRequestSamlSsoConfig {
	s.NameIdFormat = &v
	return s
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) SetNameIdValueExpression(v string) *SetApplicationSsoConfigRequestSamlSsoConfig {
	s.NameIdValueExpression = &v
	return s
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) SetResponseSigned(v bool) *SetApplicationSsoConfigRequestSamlSsoConfig {
	s.ResponseSigned = &v
	return s
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) SetSignatureAlgorithm(v string) *SetApplicationSsoConfigRequestSamlSsoConfig {
	s.SignatureAlgorithm = &v
	return s
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) SetSpEntityId(v string) *SetApplicationSsoConfigRequestSamlSsoConfig {
	s.SpEntityId = &v
	return s
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfig) SetSpSsoAcsUrl(v string) *SetApplicationSsoConfigRequestSamlSsoConfig {
	s.SpSsoAcsUrl = &v
	return s
}

type SetApplicationSsoConfigRequestSamlSsoConfigAttributeStatements struct {
	// The attribute name.
	//
	// example:
	//
	// https://www.aliyun.com/SAML-Role/Attributes/RoleSessionName
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	// The expression that is used to generate the value of the attribute.
	//
	// example:
	//
	// user.username
	AttributeValueExpression *string `json:"AttributeValueExpression,omitempty" xml:"AttributeValueExpression,omitempty"`
}

func (s SetApplicationSsoConfigRequestSamlSsoConfigAttributeStatements) String() string {
	return tea.Prettify(s)
}

func (s SetApplicationSsoConfigRequestSamlSsoConfigAttributeStatements) GoString() string {
	return s.String()
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfigAttributeStatements) SetAttributeName(v string) *SetApplicationSsoConfigRequestSamlSsoConfigAttributeStatements {
	s.AttributeName = &v
	return s
}

func (s *SetApplicationSsoConfigRequestSamlSsoConfigAttributeStatements) SetAttributeValueExpression(v string) *SetApplicationSsoConfigRequestSamlSsoConfigAttributeStatements {
	s.AttributeValueExpression = &v
	return s
}

type SetApplicationSsoConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetApplicationSsoConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetApplicationSsoConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetApplicationSsoConfigResponseBody) SetRequestId(v string) *SetApplicationSsoConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetApplicationSsoConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetApplicationSsoConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetApplicationSsoConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetApplicationSsoConfigResponse) GoString() string {
	return s.String()
}

func (s *SetApplicationSsoConfigResponse) SetHeaders(v map[string]*string) *SetApplicationSsoConfigResponse {
	s.Headers = v
	return s
}

func (s *SetApplicationSsoConfigResponse) SetStatusCode(v int32) *SetApplicationSsoConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetApplicationSsoConfigResponse) SetBody(v *SetApplicationSsoConfigResponseBody) *SetApplicationSsoConfigResponse {
	s.Body = v
	return s
}

type SetDefaultDomainRequest struct {
	// 域名ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// dm_examplexxxxx
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s SetDefaultDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultDomainRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultDomainRequest) SetDomainId(v string) *SetDefaultDomainRequest {
	s.DomainId = &v
	return s
}

func (s *SetDefaultDomainRequest) SetInstanceId(v string) *SetDefaultDomainRequest {
	s.InstanceId = &v
	return s
}

type SetDefaultDomainResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDefaultDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultDomainResponseBody) GoString() string {
	return s.String()
}

func (s *SetDefaultDomainResponseBody) SetRequestId(v string) *SetDefaultDomainResponseBody {
	s.RequestId = &v
	return s
}

type SetDefaultDomainResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDefaultDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDefaultDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultDomainResponse) GoString() string {
	return s.String()
}

func (s *SetDefaultDomainResponse) SetHeaders(v map[string]*string) *SetDefaultDomainResponse {
	s.Headers = v
	return s
}

func (s *SetDefaultDomainResponse) SetStatusCode(v int32) *SetDefaultDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDefaultDomainResponse) SetBody(v *SetDefaultDomainResponseBody) *SetDefaultDomainResponse {
	s.Body = v
	return s
}

type SetForgetPasswordConfigurationRequest struct {
	// 身份认证渠道。枚举取值:email(邮件)、sms(短信)
	//
	// example:
	//
	// email
	AuthenticationChannels []*string `json:"AuthenticationChannels,omitempty" xml:"AuthenticationChannels,omitempty" type:"Repeated"`
	// 忘记密码配置状态。枚举取值:enabled(开启)、disabled(禁用)
	//
	// This parameter is required.
	//
	// example:
	//
	// enabled
	ForgetPasswordStatus *string `json:"ForgetPasswordStatus,omitempty" xml:"ForgetPasswordStatus,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// eiam-111ccc1111
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s SetForgetPasswordConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s SetForgetPasswordConfigurationRequest) GoString() string {
	return s.String()
}

func (s *SetForgetPasswordConfigurationRequest) SetAuthenticationChannels(v []*string) *SetForgetPasswordConfigurationRequest {
	s.AuthenticationChannels = v
	return s
}

func (s *SetForgetPasswordConfigurationRequest) SetForgetPasswordStatus(v string) *SetForgetPasswordConfigurationRequest {
	s.ForgetPasswordStatus = &v
	return s
}

func (s *SetForgetPasswordConfigurationRequest) SetInstanceId(v string) *SetForgetPasswordConfigurationRequest {
	s.InstanceId = &v
	return s
}

type SetForgetPasswordConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetForgetPasswordConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetForgetPasswordConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *SetForgetPasswordConfigurationResponseBody) SetRequestId(v string) *SetForgetPasswordConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type SetForgetPasswordConfigurationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetForgetPasswordConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetForgetPasswordConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s SetForgetPasswordConfigurationResponse) GoString() string {
	return s.String()
}

func (s *SetForgetPasswordConfigurationResponse) SetHeaders(v map[string]*string) *SetForgetPasswordConfigurationResponse {
	s.Headers = v
	return s
}

func (s *SetForgetPasswordConfigurationResponse) SetStatusCode(v int32) *SetForgetPasswordConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *SetForgetPasswordConfigurationResponse) SetBody(v *SetForgetPasswordConfigurationResponseBody) *SetForgetPasswordConfigurationResponse {
	s.Body = v
	return s
}

type SetPasswordComplexityConfigurationRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The password complexity rules.
	PasswordComplexityRules []*SetPasswordComplexityConfigurationRequestPasswordComplexityRules `json:"PasswordComplexityRules,omitempty" xml:"PasswordComplexityRules,omitempty" type:"Repeated"`
	// The minimum number of characters in a password.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PasswordMinLength *int32 `json:"PasswordMinLength,omitempty" xml:"PasswordMinLength,omitempty"`
}

func (s SetPasswordComplexityConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s SetPasswordComplexityConfigurationRequest) GoString() string {
	return s.String()
}

func (s *SetPasswordComplexityConfigurationRequest) SetInstanceId(v string) *SetPasswordComplexityConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *SetPasswordComplexityConfigurationRequest) SetPasswordComplexityRules(v []*SetPasswordComplexityConfigurationRequestPasswordComplexityRules) *SetPasswordComplexityConfigurationRequest {
	s.PasswordComplexityRules = v
	return s
}

func (s *SetPasswordComplexityConfigurationRequest) SetPasswordMinLength(v int32) *SetPasswordComplexityConfigurationRequest {
	s.PasswordMinLength = &v
	return s
}

type SetPasswordComplexityConfigurationRequestPasswordComplexityRules struct {
	// The type of the password check. Valid values:
	//
	// 	- inclusion_upper_case: The password must contain uppercase letters.
	//
	// 	- inclusion_lower_case: The password must contain lowercase letters.
	//
	// 	- inclusion_special_case: The password must contain one or more of the following special characters: @ % + \\ / \\" ! # $ ^ ? : , ( ) { } [ ] ~ - _ .
	//
	// 	- inclusion_number: The password must contain digits.
	//
	// 	- exclusion_username: The password cannot contain a username.
	//
	// 	- exclusion_email: The password cannot contain an email prefix.
	//
	// 	- exclusion_phone_number: The password cannot contain a mobile number.
	//
	// 	- exclusion_display_name: The password cannot contain a display name.
	//
	// example:
	//
	// inclusion_upper_case
	PasswordCheckType *string `json:"PasswordCheckType,omitempty" xml:"PasswordCheckType,omitempty"`
}

func (s SetPasswordComplexityConfigurationRequestPasswordComplexityRules) String() string {
	return tea.Prettify(s)
}

func (s SetPasswordComplexityConfigurationRequestPasswordComplexityRules) GoString() string {
	return s.String()
}

func (s *SetPasswordComplexityConfigurationRequestPasswordComplexityRules) SetPasswordCheckType(v string) *SetPasswordComplexityConfigurationRequestPasswordComplexityRules {
	s.PasswordCheckType = &v
	return s
}

type SetPasswordComplexityConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPasswordComplexityConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetPasswordComplexityConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *SetPasswordComplexityConfigurationResponseBody) SetRequestId(v string) *SetPasswordComplexityConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type SetPasswordComplexityConfigurationResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetPasswordComplexityConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetPasswordComplexityConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s SetPasswordComplexityConfigurationResponse) GoString() string {
	return s.String()
}

func (s *SetPasswordComplexityConfigurationResponse) SetHeaders(v map[string]*string) *SetPasswordComplexityConfigurationResponse {
	s.Headers = v
	return s
}

func (s *SetPasswordComplexityConfigurationResponse) SetStatusCode(v int32) *SetPasswordComplexityConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *SetPasswordComplexityConfigurationResponse) SetBody(v *SetPasswordComplexityConfigurationResponseBody) *SetPasswordComplexityConfigurationResponse {
	s.Body = v
	return s
}

type SetPasswordExpirationConfigurationRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The action to take upon password expiration. This parameter must be specified when PasswordExpirationStatus is set to enabled. Valid values:
	//
	// 	- forbid_login: Users cannot log on to IDaaS.
	//
	// 	- force_update_password: Users must change the password.
	//
	// 	- remind_update_password: IDaaS reminds users to change the password upon each logon.
	//
	// example:
	//
	// force_update_password
	PasswordExpirationAction *string `json:"PasswordExpirationAction,omitempty" xml:"PasswordExpirationAction,omitempty"`
	// The methods for receiving password expiration notifications. This parameter must be specified when PasswordExpirationNotificationStatus is set to enabled.
	//
	// example:
	//
	// login
	PasswordExpirationNotificationChannels []*string `json:"PasswordExpirationNotificationChannels,omitempty" xml:"PasswordExpirationNotificationChannels,omitempty" type:"Repeated"`
	// The number of days before the expiration date during which password expiration notifications are sent. Unit: day. This parameter must be specified when PasswordExpirationNotificationStatus is set to enabled.
	//
	// example:
	//
	// 7
	PasswordExpirationNotificationDuration *int32 `json:"PasswordExpirationNotificationDuration,omitempty" xml:"PasswordExpirationNotificationDuration,omitempty"`
	// Specifies whether to enable the password expiration notification feature. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// example:
	//
	// enabled
	PasswordExpirationNotificationStatus *string `json:"PasswordExpirationNotificationStatus,omitempty" xml:"PasswordExpirationNotificationStatus,omitempty"`
	// Specifies whether to enable the password expiration feature. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// This parameter is required.
	//
	// example:
	//
	// enabled
	PasswordExpirationStatus *string `json:"PasswordExpirationStatus,omitempty" xml:"PasswordExpirationStatus,omitempty"`
	// The number of days before which users must change the password to prevent password expiration. Unit: day. You must set this parameter to a value greater than the value of PasswordExpirationNotificationDuration.
	//
	// example:
	//
	// 7
	PasswordForcedUpdateDuration *int32 `json:"PasswordForcedUpdateDuration,omitempty" xml:"PasswordForcedUpdateDuration,omitempty"`
	// The validity period of a password. Unit: day. This parameter must be specified when PasswordExpirationStatus is set to enabled.
	//
	// example:
	//
	// 180
	PasswordValidMaxDay *int32 `json:"PasswordValidMaxDay,omitempty" xml:"PasswordValidMaxDay,omitempty"`
}

func (s SetPasswordExpirationConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s SetPasswordExpirationConfigurationRequest) GoString() string {
	return s.String()
}

func (s *SetPasswordExpirationConfigurationRequest) SetInstanceId(v string) *SetPasswordExpirationConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *SetPasswordExpirationConfigurationRequest) SetPasswordExpirationAction(v string) *SetPasswordExpirationConfigurationRequest {
	s.PasswordExpirationAction = &v
	return s
}

func (s *SetPasswordExpirationConfigurationRequest) SetPasswordExpirationNotificationChannels(v []*string) *SetPasswordExpirationConfigurationRequest {
	s.PasswordExpirationNotificationChannels = v
	return s
}

func (s *SetPasswordExpirationConfigurationRequest) SetPasswordExpirationNotificationDuration(v int32) *SetPasswordExpirationConfigurationRequest {
	s.PasswordExpirationNotificationDuration = &v
	return s
}

func (s *SetPasswordExpirationConfigurationRequest) SetPasswordExpirationNotificationStatus(v string) *SetPasswordExpirationConfigurationRequest {
	s.PasswordExpirationNotificationStatus = &v
	return s
}

func (s *SetPasswordExpirationConfigurationRequest) SetPasswordExpirationStatus(v string) *SetPasswordExpirationConfigurationRequest {
	s.PasswordExpirationStatus = &v
	return s
}

func (s *SetPasswordExpirationConfigurationRequest) SetPasswordForcedUpdateDuration(v int32) *SetPasswordExpirationConfigurationRequest {
	s.PasswordForcedUpdateDuration = &v
	return s
}

func (s *SetPasswordExpirationConfigurationRequest) SetPasswordValidMaxDay(v int32) *SetPasswordExpirationConfigurationRequest {
	s.PasswordValidMaxDay = &v
	return s
}

type SetPasswordExpirationConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPasswordExpirationConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetPasswordExpirationConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *SetPasswordExpirationConfigurationResponseBody) SetRequestId(v string) *SetPasswordExpirationConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type SetPasswordExpirationConfigurationResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetPasswordExpirationConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetPasswordExpirationConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s SetPasswordExpirationConfigurationResponse) GoString() string {
	return s.String()
}

func (s *SetPasswordExpirationConfigurationResponse) SetHeaders(v map[string]*string) *SetPasswordExpirationConfigurationResponse {
	s.Headers = v
	return s
}

func (s *SetPasswordExpirationConfigurationResponse) SetStatusCode(v int32) *SetPasswordExpirationConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *SetPasswordExpirationConfigurationResponse) SetBody(v *SetPasswordExpirationConfigurationResponseBody) *SetPasswordExpirationConfigurationResponse {
	s.Body = v
	return s
}

type SetPasswordHistoryConfigurationRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of recent passwords that can be retained. This parameter must be specified when PasswordHistoryStatus is set to enabled.
	//
	// example:
	//
	// 3
	PasswordHistoryMaxRetention *int32 `json:"PasswordHistoryMaxRetention,omitempty" xml:"PasswordHistoryMaxRetention,omitempty"`
	// Specifies whether to enable the password history feature. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// This parameter is required.
	//
	// example:
	//
	// enabled
	PasswordHistoryStatus *string `json:"PasswordHistoryStatus,omitempty" xml:"PasswordHistoryStatus,omitempty"`
}

func (s SetPasswordHistoryConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s SetPasswordHistoryConfigurationRequest) GoString() string {
	return s.String()
}

func (s *SetPasswordHistoryConfigurationRequest) SetInstanceId(v string) *SetPasswordHistoryConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *SetPasswordHistoryConfigurationRequest) SetPasswordHistoryMaxRetention(v int32) *SetPasswordHistoryConfigurationRequest {
	s.PasswordHistoryMaxRetention = &v
	return s
}

func (s *SetPasswordHistoryConfigurationRequest) SetPasswordHistoryStatus(v string) *SetPasswordHistoryConfigurationRequest {
	s.PasswordHistoryStatus = &v
	return s
}

type SetPasswordHistoryConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPasswordHistoryConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetPasswordHistoryConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *SetPasswordHistoryConfigurationResponseBody) SetRequestId(v string) *SetPasswordHistoryConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type SetPasswordHistoryConfigurationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetPasswordHistoryConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetPasswordHistoryConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s SetPasswordHistoryConfigurationResponse) GoString() string {
	return s.String()
}

func (s *SetPasswordHistoryConfigurationResponse) SetHeaders(v map[string]*string) *SetPasswordHistoryConfigurationResponse {
	s.Headers = v
	return s
}

func (s *SetPasswordHistoryConfigurationResponse) SetStatusCode(v int32) *SetPasswordHistoryConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *SetPasswordHistoryConfigurationResponse) SetBody(v *SetPasswordHistoryConfigurationResponseBody) *SetPasswordHistoryConfigurationResponse {
	s.Body = v
	return s
}

type SetPasswordInitializationConfigurationRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to enable forcible password change upon first logon. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// example:
	//
	// enabled
	PasswordForcedUpdateStatus *string `json:"PasswordForcedUpdateStatus,omitempty" xml:"PasswordForcedUpdateStatus,omitempty"`
	// The methods for receiving password initialization notifications.
	//
	// example:
	//
	// email
	PasswordInitializationNotificationChannels []*string `json:"PasswordInitializationNotificationChannels,omitempty" xml:"PasswordInitializationNotificationChannels,omitempty" type:"Repeated"`
	// Specifies whether to enable password initialization. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// This parameter is required.
	//
	// example:
	//
	// enabled
	PasswordInitializationStatus *string `json:"PasswordInitializationStatus,omitempty" xml:"PasswordInitializationStatus,omitempty"`
	// The password initialization method. This parameter is required when PasswordInitializationStatus is set to enabled. Set the value to random.
	//
	// 	- random: A randomly generated password is used.
	//
	// example:
	//
	// random
	PasswordInitializationType *string `json:"PasswordInitializationType,omitempty" xml:"PasswordInitializationType,omitempty"`
}

func (s SetPasswordInitializationConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s SetPasswordInitializationConfigurationRequest) GoString() string {
	return s.String()
}

func (s *SetPasswordInitializationConfigurationRequest) SetInstanceId(v string) *SetPasswordInitializationConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *SetPasswordInitializationConfigurationRequest) SetPasswordForcedUpdateStatus(v string) *SetPasswordInitializationConfigurationRequest {
	s.PasswordForcedUpdateStatus = &v
	return s
}

func (s *SetPasswordInitializationConfigurationRequest) SetPasswordInitializationNotificationChannels(v []*string) *SetPasswordInitializationConfigurationRequest {
	s.PasswordInitializationNotificationChannels = v
	return s
}

func (s *SetPasswordInitializationConfigurationRequest) SetPasswordInitializationStatus(v string) *SetPasswordInitializationConfigurationRequest {
	s.PasswordInitializationStatus = &v
	return s
}

func (s *SetPasswordInitializationConfigurationRequest) SetPasswordInitializationType(v string) *SetPasswordInitializationConfigurationRequest {
	s.PasswordInitializationType = &v
	return s
}

type SetPasswordInitializationConfigurationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPasswordInitializationConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetPasswordInitializationConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *SetPasswordInitializationConfigurationResponseBody) SetRequestId(v string) *SetPasswordInitializationConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type SetPasswordInitializationConfigurationResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetPasswordInitializationConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetPasswordInitializationConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s SetPasswordInitializationConfigurationResponse) GoString() string {
	return s.String()
}

func (s *SetPasswordInitializationConfigurationResponse) SetHeaders(v map[string]*string) *SetPasswordInitializationConfigurationResponse {
	s.Headers = v
	return s
}

func (s *SetPasswordInitializationConfigurationResponse) SetStatusCode(v int32) *SetPasswordInitializationConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *SetPasswordInitializationConfigurationResponse) SetBody(v *SetPasswordInitializationConfigurationResponseBody) *SetPasswordInitializationConfigurationResponse {
	s.Body = v
	return s
}

type SetUserPrimaryOrganizationalUnitRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the new primary organizational unit.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
	// The ID of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SetUserPrimaryOrganizationalUnitRequest) String() string {
	return tea.Prettify(s)
}

func (s SetUserPrimaryOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *SetUserPrimaryOrganizationalUnitRequest) SetInstanceId(v string) *SetUserPrimaryOrganizationalUnitRequest {
	s.InstanceId = &v
	return s
}

func (s *SetUserPrimaryOrganizationalUnitRequest) SetOrganizationalUnitId(v string) *SetUserPrimaryOrganizationalUnitRequest {
	s.OrganizationalUnitId = &v
	return s
}

func (s *SetUserPrimaryOrganizationalUnitRequest) SetUserId(v string) *SetUserPrimaryOrganizationalUnitRequest {
	s.UserId = &v
	return s
}

type SetUserPrimaryOrganizationalUnitResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetUserPrimaryOrganizationalUnitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetUserPrimaryOrganizationalUnitResponseBody) GoString() string {
	return s.String()
}

func (s *SetUserPrimaryOrganizationalUnitResponseBody) SetRequestId(v string) *SetUserPrimaryOrganizationalUnitResponseBody {
	s.RequestId = &v
	return s
}

type SetUserPrimaryOrganizationalUnitResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetUserPrimaryOrganizationalUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetUserPrimaryOrganizationalUnitResponse) String() string {
	return tea.Prettify(s)
}

func (s SetUserPrimaryOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *SetUserPrimaryOrganizationalUnitResponse) SetHeaders(v map[string]*string) *SetUserPrimaryOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *SetUserPrimaryOrganizationalUnitResponse) SetStatusCode(v int32) *SetUserPrimaryOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *SetUserPrimaryOrganizationalUnitResponse) SetBody(v *SetUserPrimaryOrganizationalUnitResponseBody) *SetUserPrimaryOrganizationalUnitResponse {
	s.Body = v
	return s
}

type UnlockUserRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UnlockUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UnlockUserRequest) GoString() string {
	return s.String()
}

func (s *UnlockUserRequest) SetInstanceId(v string) *UnlockUserRequest {
	s.InstanceId = &v
	return s
}

func (s *UnlockUserRequest) SetUserId(v string) *UnlockUserRequest {
	s.UserId = &v
	return s
}

type UnlockUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnlockUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnlockUserResponseBody) GoString() string {
	return s.String()
}

func (s *UnlockUserResponseBody) SetRequestId(v string) *UnlockUserResponseBody {
	s.RequestId = &v
	return s
}

type UnlockUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnlockUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnlockUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UnlockUserResponse) GoString() string {
	return s.String()
}

func (s *UnlockUserResponse) SetHeaders(v map[string]*string) *UnlockUserResponse {
	s.Headers = v
	return s
}

func (s *UnlockUserResponse) SetStatusCode(v int32) *UnlockUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UnlockUserResponse) SetBody(v *UnlockUserResponseBody) *UnlockUserResponse {
	s.Body = v
	return s
}

type UpdateApplicationAuthorizationTypeRequest struct {
	// The ID of the application that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The authorization type of the application. Valid values:
	//
	// 	- authorize_required: Only the user with explicit authorization can access the application.
	//
	// 	- default_all: By default, all users can access the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// authorize_required
	AuthorizationType *string `json:"AuthorizationType,omitempty" xml:"AuthorizationType,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateApplicationAuthorizationTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationAuthorizationTypeRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationAuthorizationTypeRequest) SetApplicationId(v string) *UpdateApplicationAuthorizationTypeRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdateApplicationAuthorizationTypeRequest) SetAuthorizationType(v string) *UpdateApplicationAuthorizationTypeRequest {
	s.AuthorizationType = &v
	return s
}

func (s *UpdateApplicationAuthorizationTypeRequest) SetInstanceId(v string) *UpdateApplicationAuthorizationTypeRequest {
	s.InstanceId = &v
	return s
}

type UpdateApplicationAuthorizationTypeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApplicationAuthorizationTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationAuthorizationTypeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationAuthorizationTypeResponseBody) SetRequestId(v string) *UpdateApplicationAuthorizationTypeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateApplicationAuthorizationTypeResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationAuthorizationTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationAuthorizationTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationAuthorizationTypeResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationAuthorizationTypeResponse) SetHeaders(v map[string]*string) *UpdateApplicationAuthorizationTypeResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationAuthorizationTypeResponse) SetStatusCode(v int32) *UpdateApplicationAuthorizationTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationAuthorizationTypeResponse) SetBody(v *UpdateApplicationAuthorizationTypeResponseBody) *UpdateApplicationAuthorizationTypeResponse {
	s.Body = v
	return s
}

type UpdateApplicationDescriptionRequest struct {
	// The ID of the application that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// A demo application that is used for test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk2676xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateApplicationDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationDescriptionRequest) SetApplicationId(v string) *UpdateApplicationDescriptionRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdateApplicationDescriptionRequest) SetDescription(v string) *UpdateApplicationDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateApplicationDescriptionRequest) SetInstanceId(v string) *UpdateApplicationDescriptionRequest {
	s.InstanceId = &v
	return s
}

type UpdateApplicationDescriptionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApplicationDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationDescriptionResponseBody) SetRequestId(v string) *UpdateApplicationDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateApplicationDescriptionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationDescriptionResponse) SetHeaders(v map[string]*string) *UpdateApplicationDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationDescriptionResponse) SetStatusCode(v int32) *UpdateApplicationDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationDescriptionResponse) SetBody(v *UpdateApplicationDescriptionResponseBody) *UpdateApplicationDescriptionResponse {
	s.Body = v
	return s
}

type UpdateGroupRequest struct {
	// The external ID of the group.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupExternalId *string `json:"GroupExternalId,omitempty" xml:"GroupExternalId,omitempty"`
	// The group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the group.
	//
	// example:
	//
	// name_test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupRequest) SetGroupExternalId(v string) *UpdateGroupRequest {
	s.GroupExternalId = &v
	return s
}

func (s *UpdateGroupRequest) SetGroupId(v string) *UpdateGroupRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateGroupRequest) SetGroupName(v string) *UpdateGroupRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateGroupRequest) SetInstanceId(v string) *UpdateGroupRequest {
	s.InstanceId = &v
	return s
}

type UpdateGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupResponseBody) SetRequestId(v string) *UpdateGroupResponseBody {
	s.RequestId = &v
	return s
}

type UpdateGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupResponse) SetHeaders(v map[string]*string) *UpdateGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupResponse) SetStatusCode(v int32) *UpdateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGroupResponse) SetBody(v *UpdateGroupResponseBody) *UpdateGroupResponse {
	s.Body = v
	return s
}

type UpdateGroupDescriptionRequest struct {
	// The description of the account group. The value can be up to 256 characters in length.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the account group.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_d6sbsuumeta4h66ec3il7yxxxx
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateGroupDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupDescriptionRequest) SetDescription(v string) *UpdateGroupDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateGroupDescriptionRequest) SetGroupId(v string) *UpdateGroupDescriptionRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateGroupDescriptionRequest) SetInstanceId(v string) *UpdateGroupDescriptionRequest {
	s.InstanceId = &v
	return s
}

type UpdateGroupDescriptionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGroupDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupDescriptionResponseBody) SetRequestId(v string) *UpdateGroupDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateGroupDescriptionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGroupDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGroupDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupDescriptionResponse) SetHeaders(v map[string]*string) *UpdateGroupDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupDescriptionResponse) SetStatusCode(v int32) *UpdateGroupDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGroupDescriptionResponse) SetBody(v *UpdateGroupDescriptionResponseBody) *UpdateGroupDescriptionResponse {
	s.Body = v
	return s
}

type UpdateInstanceDescriptionRequest struct {
	// The new description of the instance.
	//
	// example:
	//
	// 测试实例
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance whose description you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateInstanceDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceDescriptionRequest) SetDescription(v string) *UpdateInstanceDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateInstanceDescriptionRequest) SetInstanceId(v string) *UpdateInstanceDescriptionRequest {
	s.InstanceId = &v
	return s
}

type UpdateInstanceDescriptionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceDescriptionResponseBody) SetRequestId(v string) *UpdateInstanceDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateInstanceDescriptionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceDescriptionResponse) SetHeaders(v map[string]*string) *UpdateInstanceDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceDescriptionResponse) SetStatusCode(v int32) *UpdateInstanceDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceDescriptionResponse) SetBody(v *UpdateInstanceDescriptionResponseBody) *UpdateInstanceDescriptionResponse {
	s.Body = v
	return s
}

type UpdateNetworkAccessEndpointNameRequest struct {
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 专属网络端点ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// nae_examplexxxx
	NetworkAccessEndpointId *string `json:"NetworkAccessEndpointId,omitempty" xml:"NetworkAccessEndpointId,omitempty"`
	// 专属网络端点名称。
	//
	// This parameter is required.
	//
	// example:
	//
	// xx业务VPC访问端点
	NetworkAccessEndpointName *string `json:"NetworkAccessEndpointName,omitempty" xml:"NetworkAccessEndpointName,omitempty"`
}

func (s UpdateNetworkAccessEndpointNameRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNetworkAccessEndpointNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateNetworkAccessEndpointNameRequest) SetInstanceId(v string) *UpdateNetworkAccessEndpointNameRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateNetworkAccessEndpointNameRequest) SetNetworkAccessEndpointId(v string) *UpdateNetworkAccessEndpointNameRequest {
	s.NetworkAccessEndpointId = &v
	return s
}

func (s *UpdateNetworkAccessEndpointNameRequest) SetNetworkAccessEndpointName(v string) *UpdateNetworkAccessEndpointNameRequest {
	s.NetworkAccessEndpointName = &v
	return s
}

type UpdateNetworkAccessEndpointNameResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateNetworkAccessEndpointNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateNetworkAccessEndpointNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNetworkAccessEndpointNameResponseBody) SetRequestId(v string) *UpdateNetworkAccessEndpointNameResponseBody {
	s.RequestId = &v
	return s
}

type UpdateNetworkAccessEndpointNameResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNetworkAccessEndpointNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNetworkAccessEndpointNameResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateNetworkAccessEndpointNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateNetworkAccessEndpointNameResponse) SetHeaders(v map[string]*string) *UpdateNetworkAccessEndpointNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateNetworkAccessEndpointNameResponse) SetStatusCode(v int32) *UpdateNetworkAccessEndpointNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNetworkAccessEndpointNameResponse) SetBody(v *UpdateNetworkAccessEndpointNameResponseBody) *UpdateNetworkAccessEndpointNameResponse {
	s.Body = v
	return s
}

type UpdateOrganizationalUnitRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The organization ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
	// The name of the organization. The name can be up to 64 characters in length and must be unique in the same parent organization.
	//
	// example:
	//
	// ou_name
	OrganizationalUnitName *string `json:"OrganizationalUnitName,omitempty" xml:"OrganizationalUnitName,omitempty"`
}

func (s UpdateOrganizationalUnitRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationalUnitRequest) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationalUnitRequest) SetInstanceId(v string) *UpdateOrganizationalUnitRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateOrganizationalUnitRequest) SetOrganizationalUnitId(v string) *UpdateOrganizationalUnitRequest {
	s.OrganizationalUnitId = &v
	return s
}

func (s *UpdateOrganizationalUnitRequest) SetOrganizationalUnitName(v string) *UpdateOrganizationalUnitRequest {
	s.OrganizationalUnitName = &v
	return s
}

type UpdateOrganizationalUnitResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOrganizationalUnitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationalUnitResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationalUnitResponseBody) SetRequestId(v string) *UpdateOrganizationalUnitResponseBody {
	s.RequestId = &v
	return s
}

type UpdateOrganizationalUnitResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOrganizationalUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOrganizationalUnitResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationalUnitResponse) SetHeaders(v map[string]*string) *UpdateOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *UpdateOrganizationalUnitResponse) SetStatusCode(v int32) *UpdateOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOrganizationalUnitResponse) SetBody(v *UpdateOrganizationalUnitResponseBody) *UpdateOrganizationalUnitResponse {
	s.Body = v
	return s
}

type UpdateOrganizationalUnitDescriptionRequest struct {
	// The description of the organization. The value can be up to 256 characters in length.
	//
	// example:
	//
	// organizationalUnit_test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The organization ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
}

func (s UpdateOrganizationalUnitDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationalUnitDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationalUnitDescriptionRequest) SetDescription(v string) *UpdateOrganizationalUnitDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateOrganizationalUnitDescriptionRequest) SetInstanceId(v string) *UpdateOrganizationalUnitDescriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateOrganizationalUnitDescriptionRequest) SetOrganizationalUnitId(v string) *UpdateOrganizationalUnitDescriptionRequest {
	s.OrganizationalUnitId = &v
	return s
}

type UpdateOrganizationalUnitDescriptionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOrganizationalUnitDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationalUnitDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationalUnitDescriptionResponseBody) SetRequestId(v string) *UpdateOrganizationalUnitDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateOrganizationalUnitDescriptionResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOrganizationalUnitDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOrganizationalUnitDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationalUnitDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationalUnitDescriptionResponse) SetHeaders(v map[string]*string) *UpdateOrganizationalUnitDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateOrganizationalUnitDescriptionResponse) SetStatusCode(v int32) *UpdateOrganizationalUnitDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOrganizationalUnitDescriptionResponse) SetBody(v *UpdateOrganizationalUnitDescriptionResponseBody) *UpdateOrganizationalUnitDescriptionResponse {
	s.Body = v
	return s
}

type UpdateOrganizationalUnitParentIdRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The organization ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	OrganizationalUnitId *string `json:"OrganizationalUnitId,omitempty" xml:"OrganizationalUnitId,omitempty"`
	// The parent organization ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s UpdateOrganizationalUnitParentIdRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationalUnitParentIdRequest) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationalUnitParentIdRequest) SetInstanceId(v string) *UpdateOrganizationalUnitParentIdRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateOrganizationalUnitParentIdRequest) SetOrganizationalUnitId(v string) *UpdateOrganizationalUnitParentIdRequest {
	s.OrganizationalUnitId = &v
	return s
}

func (s *UpdateOrganizationalUnitParentIdRequest) SetParentId(v string) *UpdateOrganizationalUnitParentIdRequest {
	s.ParentId = &v
	return s
}

type UpdateOrganizationalUnitParentIdResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOrganizationalUnitParentIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationalUnitParentIdResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationalUnitParentIdResponseBody) SetRequestId(v string) *UpdateOrganizationalUnitParentIdResponseBody {
	s.RequestId = &v
	return s
}

type UpdateOrganizationalUnitParentIdResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOrganizationalUnitParentIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOrganizationalUnitParentIdResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOrganizationalUnitParentIdResponse) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationalUnitParentIdResponse) SetHeaders(v map[string]*string) *UpdateOrganizationalUnitParentIdResponse {
	s.Headers = v
	return s
}

func (s *UpdateOrganizationalUnitParentIdResponse) SetStatusCode(v int32) *UpdateOrganizationalUnitParentIdResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOrganizationalUnitParentIdResponse) SetBody(v *UpdateOrganizationalUnitParentIdResponseBody) *UpdateOrganizationalUnitParentIdResponse {
	s.Body = v
	return s
}

type UpdateUserRequest struct {
	// The custom extended fields.
	CustomFields []*UpdateUserRequestCustomFields `json:"CustomFields,omitempty" xml:"CustomFields,omitempty" type:"Repeated"`
	// The display name of the account. The display name can be up to 64 characters in length.
	//
	// example:
	//
	// test_name
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address. The prefix of the email address can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// example@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Specifies whether the email address is verified. This parameter must be specified if you specify Email. You can set this parameter to true if you have no special business requirements.
	//
	// example:
	//
	// true
	EmailVerified *bool `json:"EmailVerified,omitempty" xml:"EmailVerified,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The mobile number. The mobile number must be 6 to 15 digits in length.
	//
	// example:
	//
	// 156xxxxxxxxx
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Specifies whether the mobile number is verified. This parameter must be specified if you specify PhoneNumber. You can set this parameter to true if you have no special business requirements.
	//
	// example:
	//
	// true
	PhoneNumberVerified *bool `json:"PhoneNumberVerified,omitempty" xml:"PhoneNumberVerified,omitempty"`
	// The area code of the mobile number. For example, the area code of a mobile number in the Chinese mainland is 86 without 00 or the plus sign (+). This parameter must be specified if you specify PhoneNumber.
	//
	// example:
	//
	// 86
	PhoneRegion *string `json:"PhoneRegion,omitempty" xml:"PhoneRegion,omitempty"`
	// The account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The name of the account. The name can be up to 64 characters in length. It can contain letters, digits, and the following special characters: _ . @ -
	//
	// example:
	//
	// username_test
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) SetCustomFields(v []*UpdateUserRequestCustomFields) *UpdateUserRequest {
	s.CustomFields = v
	return s
}

func (s *UpdateUserRequest) SetDisplayName(v string) *UpdateUserRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateUserRequest) SetEmail(v string) *UpdateUserRequest {
	s.Email = &v
	return s
}

func (s *UpdateUserRequest) SetEmailVerified(v bool) *UpdateUserRequest {
	s.EmailVerified = &v
	return s
}

func (s *UpdateUserRequest) SetInstanceId(v string) *UpdateUserRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateUserRequest) SetPhoneNumber(v string) *UpdateUserRequest {
	s.PhoneNumber = &v
	return s
}

func (s *UpdateUserRequest) SetPhoneNumberVerified(v bool) *UpdateUserRequest {
	s.PhoneNumberVerified = &v
	return s
}

func (s *UpdateUserRequest) SetPhoneRegion(v string) *UpdateUserRequest {
	s.PhoneRegion = &v
	return s
}

func (s *UpdateUserRequest) SetUserId(v string) *UpdateUserRequest {
	s.UserId = &v
	return s
}

func (s *UpdateUserRequest) SetUsername(v string) *UpdateUserRequest {
	s.Username = &v
	return s
}

type UpdateUserRequestCustomFields struct {
	// The name of the extended field. You must create an extended field before you specify this parameter. To create an extended field, go to the Extended Fields page of the specified EIAM instance in the IDaaS console.
	//
	// example:
	//
	// nick_name
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The value of the extended field. The value follows the limits on the properties of the extended field.
	//
	// example:
	//
	// test_value
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	// The operation type of the extended field. Valid values:
	//
	// 	- add: adds a value to the extended field of the account.
	//
	// 	- replace: replaces the existing value of the extended field of the account. If the existing value to be replaced does not exist, this operation changes to the add operation.
	//
	// 	- remove: removes a value from the extended field of the account.
	//
	// example:
	//
	// add
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
}

func (s UpdateUserRequestCustomFields) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserRequestCustomFields) GoString() string {
	return s.String()
}

func (s *UpdateUserRequestCustomFields) SetFieldName(v string) *UpdateUserRequestCustomFields {
	s.FieldName = &v
	return s
}

func (s *UpdateUserRequestCustomFields) SetFieldValue(v string) *UpdateUserRequestCustomFields {
	s.FieldValue = &v
	return s
}

func (s *UpdateUserRequestCustomFields) SetOperation(v string) *UpdateUserRequestCustomFields {
	s.Operation = &v
	return s
}

type UpdateUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBody) SetRequestId(v string) *UpdateUserResponseBody {
	s.RequestId = &v
	return s
}

type UpdateUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserResponse) SetHeaders(v map[string]*string) *UpdateUserResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserResponse) SetStatusCode(v int32) *UpdateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserResponse) SetBody(v *UpdateUserResponseBody) *UpdateUserResponse {
	s.Body = v
	return s
}

type UpdateUserDescriptionRequest struct {
	// The description of the account. The value can be up to 256 characters in length.
	//
	// example:
	//
	// this is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateUserDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserDescriptionRequest) SetDescription(v string) *UpdateUserDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateUserDescriptionRequest) SetInstanceId(v string) *UpdateUserDescriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateUserDescriptionRequest) SetUserId(v string) *UpdateUserDescriptionRequest {
	s.UserId = &v
	return s
}

type UpdateUserDescriptionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserDescriptionResponseBody) SetRequestId(v string) *UpdateUserDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateUserDescriptionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserDescriptionResponse) SetHeaders(v map[string]*string) *UpdateUserDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserDescriptionResponse) SetStatusCode(v int32) *UpdateUserDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserDescriptionResponse) SetBody(v *UpdateUserDescriptionResponseBody) *UpdateUserDescriptionResponse {
	s.Body = v
	return s
}

type UpdateUserPasswordRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The new password of the account. For more information about the password format, see the "Password Policies" topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether to enable forcible password change upon first logon. Default value: disabled. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// example:
	//
	// enabled
	PasswordForcedUpdateStatus *string `json:"PasswordForcedUpdateStatus,omitempty" xml:"PasswordForcedUpdateStatus,omitempty"`
	// The account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The methods for receiving password notifications.
	//
	// example:
	//
	// sms
	UserNotificationChannels []*string `json:"UserNotificationChannels,omitempty" xml:"UserNotificationChannels,omitempty" type:"Repeated"`
}

func (s UpdateUserPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserPasswordRequest) SetInstanceId(v string) *UpdateUserPasswordRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateUserPasswordRequest) SetPassword(v string) *UpdateUserPasswordRequest {
	s.Password = &v
	return s
}

func (s *UpdateUserPasswordRequest) SetPasswordForcedUpdateStatus(v string) *UpdateUserPasswordRequest {
	s.PasswordForcedUpdateStatus = &v
	return s
}

func (s *UpdateUserPasswordRequest) SetUserId(v string) *UpdateUserPasswordRequest {
	s.UserId = &v
	return s
}

func (s *UpdateUserPasswordRequest) SetUserNotificationChannels(v []*string) *UpdateUserPasswordRequest {
	s.UserNotificationChannels = v
	return s
}

type UpdateUserPasswordResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserPasswordResponseBody) SetRequestId(v string) *UpdateUserPasswordResponseBody {
	s.RequestId = &v
	return s
}

type UpdateUserPasswordResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserPasswordResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserPasswordResponse) SetHeaders(v map[string]*string) *UpdateUserPasswordResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserPasswordResponse) SetStatusCode(v int32) *UpdateUserPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserPasswordResponse) SetBody(v *UpdateUserPasswordResponseBody) *UpdateUserPasswordResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("eiam"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an Employee Identity and Access Management (EIAM) account to multiple EIAM organizations of Identity as a Service (IDaaS). If the account already exists in the organizational unit, the system directly returns a success response.
//
// @param request - AddUserToOrganizationalUnitsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUserToOrganizationalUnitsResponse
func (client *Client) AddUserToOrganizationalUnitsWithOptions(request *AddUserToOrganizationalUnitsRequest, runtime *util.RuntimeOptions) (_result *AddUserToOrganizationalUnitsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitIds)) {
		query["OrganizationalUnitIds"] = request.OrganizationalUnitIds
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUserToOrganizationalUnits"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddUserToOrganizationalUnitsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an Employee Identity and Access Management (EIAM) account to multiple EIAM organizations of Identity as a Service (IDaaS). If the account already exists in the organizational unit, the system directly returns a success response.
//
// @param request - AddUserToOrganizationalUnitsRequest
//
// @return AddUserToOrganizationalUnitsResponse
func (client *Client) AddUserToOrganizationalUnits(request *AddUserToOrganizationalUnitsRequest) (_result *AddUserToOrganizationalUnitsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUserToOrganizationalUnitsResponse{}
	_body, _err := client.AddUserToOrganizationalUnitsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds Employee Identity and Access Management (EIAM) accounts to an EIAM group of Identity as a Service (IDaaS).
//
// @param request - AddUsersToGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddUsersToGroupResponse
func (client *Client) AddUsersToGroupWithOptions(request *AddUsersToGroupRequest, runtime *util.RuntimeOptions) (_result *AddUsersToGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		query["UserIds"] = request.UserIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUsersToGroup"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddUsersToGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds Employee Identity and Access Management (EIAM) accounts to an EIAM group of Identity as a Service (IDaaS).
//
// @param request - AddUsersToGroupRequest
//
// @return AddUsersToGroupResponse
func (client *Client) AddUsersToGroup(request *AddUsersToGroupRequest) (_result *AddUsersToGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUsersToGroupResponse{}
	_body, _err := client.AddUsersToGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants the permissions to access an application to multiple account groups at a time in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - AuthorizeApplicationToGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeApplicationToGroupsResponse
func (client *Client) AuthorizeApplicationToGroupsWithOptions(request *AuthorizeApplicationToGroupsRequest, runtime *util.RuntimeOptions) (_result *AuthorizeApplicationToGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupIds)) {
		query["GroupIds"] = request.GroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AuthorizeApplicationToGroups"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AuthorizeApplicationToGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants the permissions to access an application to multiple account groups at a time in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - AuthorizeApplicationToGroupsRequest
//
// @return AuthorizeApplicationToGroupsResponse
func (client *Client) AuthorizeApplicationToGroups(request *AuthorizeApplicationToGroupsRequest) (_result *AuthorizeApplicationToGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AuthorizeApplicationToGroupsResponse{}
	_body, _err := client.AuthorizeApplicationToGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants the access permissions on an application to multiple Employee Identity and Access Management (EIAM) organizations at a time.
//
// @param request - AuthorizeApplicationToOrganizationalUnitsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeApplicationToOrganizationalUnitsResponse
func (client *Client) AuthorizeApplicationToOrganizationalUnitsWithOptions(request *AuthorizeApplicationToOrganizationalUnitsRequest, runtime *util.RuntimeOptions) (_result *AuthorizeApplicationToOrganizationalUnitsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitIds)) {
		query["OrganizationalUnitIds"] = request.OrganizationalUnitIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AuthorizeApplicationToOrganizationalUnits"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AuthorizeApplicationToOrganizationalUnitsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants the access permissions on an application to multiple Employee Identity and Access Management (EIAM) organizations at a time.
//
// @param request - AuthorizeApplicationToOrganizationalUnitsRequest
//
// @return AuthorizeApplicationToOrganizationalUnitsResponse
func (client *Client) AuthorizeApplicationToOrganizationalUnits(request *AuthorizeApplicationToOrganizationalUnitsRequest) (_result *AuthorizeApplicationToOrganizationalUnitsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AuthorizeApplicationToOrganizationalUnitsResponse{}
	_body, _err := client.AuthorizeApplicationToOrganizationalUnitsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants the access permissions on an application to multiple Employee Identity and Access Management (EIAM) accounts at a time.
//
// @param request - AuthorizeApplicationToUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeApplicationToUsersResponse
func (client *Client) AuthorizeApplicationToUsersWithOptions(request *AuthorizeApplicationToUsersRequest, runtime *util.RuntimeOptions) (_result *AuthorizeApplicationToUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		query["UserIds"] = request.UserIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AuthorizeApplicationToUsers"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AuthorizeApplicationToUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants the access permissions on an application to multiple Employee Identity and Access Management (EIAM) accounts at a time.
//
// @param request - AuthorizeApplicationToUsersRequest
//
// @return AuthorizeApplicationToUsersResponse
func (client *Client) AuthorizeApplicationToUsers(request *AuthorizeApplicationToUsersRequest) (_result *AuthorizeApplicationToUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AuthorizeApplicationToUsersResponse{}
	_body, _err := client.AuthorizeApplicationToUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an application to an Enterprise Identity Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// Description:
//
// IDaaS EIAM supports the following two standard single sign-on (SSO) protocols for adding applications: SAML 2.0 and OIDC. You can select an SSO protocol based on your business requirements when you add an application. You cannot change the SSO protocol that you selected after the application is added.
//
// @param request - CreateApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationResponse
func (client *Client) CreateApplicationWithOptions(request *CreateApplicationRequest, runtime *util.RuntimeOptions) (_result *CreateApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationName)) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationSourceType)) {
		query["ApplicationSourceType"] = request.ApplicationSourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationTemplateId)) {
		query["ApplicationTemplateId"] = request.ApplicationTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LogoUrl)) {
		query["LogoUrl"] = request.LogoUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SsoType)) {
		query["SsoType"] = request.SsoType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApplication"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an application to an Enterprise Identity Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// Description:
//
// IDaaS EIAM supports the following two standard single sign-on (SSO) protocols for adding applications: SAML 2.0 and OIDC. You can select an SSO protocol based on your business requirements when you add an application. You cannot change the SSO protocol that you selected after the application is added.
//
// @param request - CreateApplicationRequest
//
// @return CreateApplicationResponse
func (client *Client) CreateApplication(request *CreateApplicationRequest) (_result *CreateApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateApplicationResponse{}
	_body, _err := client.CreateApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a client key for an Employee Identity and Access Management (EIAM) application. An EIAM application can have up to two client keys.
//
// @param request - CreateApplicationClientSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationClientSecretResponse
func (client *Client) CreateApplicationClientSecretWithOptions(request *CreateApplicationClientSecretRequest, runtime *util.RuntimeOptions) (_result *CreateApplicationClientSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApplicationClientSecret"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateApplicationClientSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a client key for an Employee Identity and Access Management (EIAM) application. An EIAM application can have up to two client keys.
//
// @param request - CreateApplicationClientSecretRequest
//
// @return CreateApplicationClientSecretResponse
func (client *Client) CreateApplicationClientSecret(request *CreateApplicationClientSecretRequest) (_result *CreateApplicationClientSecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateApplicationClientSecretResponse{}
	_body, _err := client.CreateApplicationClientSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建域名。
//
// @param request - CreateDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDomainResponse
func (client *Client) CreateDomainWithOptions(request *CreateDomainRequest, runtime *util.RuntimeOptions) (_result *CreateDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Filing)) {
		query["Filing"] = request.Filing
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDomain"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建域名。
//
// @param request - CreateDomainRequest
//
// @return CreateDomainResponse
func (client *Client) CreateDomain(request *CreateDomainRequest) (_result *CreateDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDomainResponse{}
	_body, _err := client.CreateDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建域名代理Token。
//
// @param request - CreateDomainProxyTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDomainProxyTokenResponse
func (client *Client) CreateDomainProxyTokenWithOptions(request *CreateDomainProxyTokenRequest, runtime *util.RuntimeOptions) (_result *CreateDomainProxyTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainId)) {
		query["DomainId"] = request.DomainId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDomainProxyToken"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDomainProxyTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建域名代理Token。
//
// @param request - CreateDomainProxyTokenRequest
//
// @return CreateDomainProxyTokenResponse
func (client *Client) CreateDomainProxyToken(request *CreateDomainProxyTokenRequest) (_result *CreateDomainProxyTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDomainProxyTokenResponse{}
	_body, _err := client.CreateDomainProxyTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an account group in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - CreateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupResponse
func (client *Client) CreateGroupWithOptions(request *CreateGroupRequest, runtime *util.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupExternalId)) {
		query["GroupExternalId"] = request.GroupExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGroup"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an account group in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - CreateGroupRequest
//
// @return CreateGroupResponse
func (client *Client) CreateGroup(request *CreateGroupRequest) (_result *CreateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGroupResponse{}
	_body, _err := client.CreateGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an instance based on which all capabilities of Identity as a Service (IDaaS) Enterprise Identity and Access Management (EIAM) are provided.
//
// @param request - CreateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstance"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an instance based on which all capabilities of Identity as a Service (IDaaS) Enterprise Identity and Access Management (EIAM) are provided.
//
// @param request - CreateInstanceRequest
//
// @return CreateInstanceResponse
func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建一个专属网络端点。
//
// @param request - CreateNetworkAccessEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNetworkAccessEndpointResponse
func (client *Client) CreateNetworkAccessEndpointWithOptions(request *CreateNetworkAccessEndpointRequest, runtime *util.RuntimeOptions) (_result *CreateNetworkAccessEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkAccessEndpointName)) {
		query["NetworkAccessEndpointName"] = request.NetworkAccessEndpointName
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchIds)) {
		query["VSwitchIds"] = request.VSwitchIds
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcRegionId)) {
		query["VpcRegionId"] = request.VpcRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNetworkAccessEndpoint"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNetworkAccessEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建一个专属网络端点。
//
// @param request - CreateNetworkAccessEndpointRequest
//
// @return CreateNetworkAccessEndpointResponse
func (client *Client) CreateNetworkAccessEndpoint(request *CreateNetworkAccessEndpointRequest) (_result *CreateNetworkAccessEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNetworkAccessEndpointResponse{}
	_body, _err := client.CreateNetworkAccessEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an organization in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - CreateOrganizationalUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrganizationalUnitResponse
func (client *Client) CreateOrganizationalUnitWithOptions(request *CreateOrganizationalUnitRequest, runtime *util.RuntimeOptions) (_result *CreateOrganizationalUnitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitExternalId)) {
		query["OrganizationalUnitExternalId"] = request.OrganizationalUnitExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitName)) {
		query["OrganizationalUnitName"] = request.OrganizationalUnitName
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		query["ParentId"] = request.ParentId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOrganizationalUnit"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOrganizationalUnitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an organization in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - CreateOrganizationalUnitRequest
//
// @return CreateOrganizationalUnitResponse
func (client *Client) CreateOrganizationalUnit(request *CreateOrganizationalUnitRequest) (_result *CreateOrganizationalUnitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOrganizationalUnitResponse{}
	_body, _err := client.CreateOrganizationalUnitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an account in an Identity as a Service (IDaaS) Enterprise Identity Access Management (EIAM) instance.
//
// @param request - CreateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserResponse
func (client *Client) CreateUserWithOptions(request *CreateUserRequest, runtime *util.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomFields)) {
		query["CustomFields"] = request.CustomFields
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.EmailVerified)) {
		query["EmailVerified"] = request.EmailVerified
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitIds)) {
		query["OrganizationalUnitIds"] = request.OrganizationalUnitIds
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordInitializationConfig)) {
		query["PasswordInitializationConfig"] = request.PasswordInitializationConfig
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumberVerified)) {
		query["PhoneNumberVerified"] = request.PhoneNumberVerified
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneRegion)) {
		query["PhoneRegion"] = request.PhoneRegion
	}

	if !tea.BoolValue(util.IsUnset(request.PrimaryOrganizationalUnitId)) {
		query["PrimaryOrganizationalUnitId"] = request.PrimaryOrganizationalUnitId
	}

	if !tea.BoolValue(util.IsUnset(request.UserExternalId)) {
		query["UserExternalId"] = request.UserExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUser"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an account in an Identity as a Service (IDaaS) Enterprise Identity Access Management (EIAM) instance.
//
// @param request - CreateUserRequest
//
// @return CreateUserResponse
func (client *Client) CreateUser(request *CreateUserRequest) (_result *CreateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserResponse{}
	_body, _err := client.CreateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an Employee Identity and Access Management (EIAM) application.
//
// Description:
//
// Make sure that the EIAM application that you want to delete is not used before you delete the EIAM application. After you delete the EIAM application, all configurations are deleted and cannot be restored.
//
// @param request - DeleteApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationResponse
func (client *Client) DeleteApplicationWithOptions(request *DeleteApplicationRequest, runtime *util.RuntimeOptions) (_result *DeleteApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApplication"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an Employee Identity and Access Management (EIAM) application.
//
// Description:
//
// Make sure that the EIAM application that you want to delete is not used before you delete the EIAM application. After you delete the EIAM application, all configurations are deleted and cannot be restored.
//
// @param request - DeleteApplicationRequest
//
// @return DeleteApplicationResponse
func (client *Client) DeleteApplication(request *DeleteApplicationRequest) (_result *DeleteApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteApplicationResponse{}
	_body, _err := client.DeleteApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a client key for an Employee Identity and Access Management (EIAM) application.
//
// @param request - DeleteApplicationClientSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationClientSecretResponse
func (client *Client) DeleteApplicationClientSecretWithOptions(request *DeleteApplicationClientSecretRequest, runtime *util.RuntimeOptions) (_result *DeleteApplicationClientSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretId)) {
		query["SecretId"] = request.SecretId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApplicationClientSecret"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteApplicationClientSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a client key for an Employee Identity and Access Management (EIAM) application.
//
// @param request - DeleteApplicationClientSecretRequest
//
// @return DeleteApplicationClientSecretResponse
func (client *Client) DeleteApplicationClientSecret(request *DeleteApplicationClientSecretRequest) (_result *DeleteApplicationClientSecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteApplicationClientSecretResponse{}
	_body, _err := client.DeleteApplicationClientSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除域名。
//
// @param request - DeleteDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainResponse
func (client *Client) DeleteDomainWithOptions(request *DeleteDomainRequest, runtime *util.RuntimeOptions) (_result *DeleteDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainId)) {
		query["DomainId"] = request.DomainId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDomain"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除域名。
//
// @param request - DeleteDomainRequest
//
// @return DeleteDomainResponse
func (client *Client) DeleteDomain(request *DeleteDomainRequest) (_result *DeleteDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDomainResponse{}
	_body, _err := client.DeleteDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定域名代理Token，删除之前请保证代理Token处于禁用状态。
//
// @param request - DeleteDomainProxyTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainProxyTokenResponse
func (client *Client) DeleteDomainProxyTokenWithOptions(request *DeleteDomainProxyTokenRequest, runtime *util.RuntimeOptions) (_result *DeleteDomainProxyTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainId)) {
		query["DomainId"] = request.DomainId
	}

	if !tea.BoolValue(util.IsUnset(request.DomainProxyTokenId)) {
		query["DomainProxyTokenId"] = request.DomainProxyTokenId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDomainProxyToken"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDomainProxyTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定域名代理Token，删除之前请保证代理Token处于禁用状态。
//
// @param request - DeleteDomainProxyTokenRequest
//
// @return DeleteDomainProxyTokenResponse
func (client *Client) DeleteDomainProxyToken(request *DeleteDomainProxyTokenRequest) (_result *DeleteDomainProxyTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDomainProxyTokenResponse{}
	_body, _err := client.DeleteDomainProxyTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the information of an account group in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - DeleteGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGroupResponse
func (client *Client) DeleteGroupWithOptions(request *DeleteGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGroup"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the information of an account group in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - DeleteGroupRequest
//
// @return DeleteGroupResponse
func (client *Client) DeleteGroup(request *DeleteGroupRequest) (_result *DeleteGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGroupResponse{}
	_body, _err := client.DeleteGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an Enterprise Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS) that you do not need.
//
// Description:
//
// Make sure that the instance to be deleted is no longer used. If the instance is deleted, all data related to the instance will be deleted.
//
// @param request - DeleteInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstance"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an Enterprise Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS) that you do not need.
//
// Description:
//
// Make sure that the instance to be deleted is no longer used. If the instance is deleted, all data related to the instance will be deleted.
//
// @param request - DeleteInstanceRequest
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstance(request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除一个专属网络端点。
//
// @param request - DeleteNetworkAccessEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNetworkAccessEndpointResponse
func (client *Client) DeleteNetworkAccessEndpointWithOptions(request *DeleteNetworkAccessEndpointRequest, runtime *util.RuntimeOptions) (_result *DeleteNetworkAccessEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkAccessEndpointId)) {
		query["NetworkAccessEndpointId"] = request.NetworkAccessEndpointId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNetworkAccessEndpoint"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNetworkAccessEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除一个专属网络端点。
//
// @param request - DeleteNetworkAccessEndpointRequest
//
// @return DeleteNetworkAccessEndpointResponse
func (client *Client) DeleteNetworkAccessEndpoint(request *DeleteNetworkAccessEndpointRequest) (_result *DeleteNetworkAccessEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNetworkAccessEndpointResponse{}
	_body, _err := client.DeleteNetworkAccessEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an organization in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM). If the organization has EIAM accounts or child organizations, the delete operation fails.
//
// @param request - DeleteOrganizationalUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOrganizationalUnitResponse
func (client *Client) DeleteOrganizationalUnitWithOptions(request *DeleteOrganizationalUnitRequest, runtime *util.RuntimeOptions) (_result *DeleteOrganizationalUnitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitId)) {
		query["OrganizationalUnitId"] = request.OrganizationalUnitId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteOrganizationalUnit"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteOrganizationalUnitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an organization in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM). If the organization has EIAM accounts or child organizations, the delete operation fails.
//
// @param request - DeleteOrganizationalUnitRequest
//
// @return DeleteOrganizationalUnitResponse
func (client *Client) DeleteOrganizationalUnit(request *DeleteOrganizationalUnitRequest) (_result *DeleteOrganizationalUnitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteOrganizationalUnitResponse{}
	_body, _err := client.DeleteOrganizationalUnitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an Employee Identity and Access Management (EIAM) account of Identity as a Service (IDaaS). The information related to the account is cleared.
//
// @param request - DeleteUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserResponse
func (client *Client) DeleteUserWithOptions(request *DeleteUserRequest, runtime *util.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUser"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an Employee Identity and Access Management (EIAM) account of Identity as a Service (IDaaS). The information related to the account is cleared.
//
// @param request - DeleteUserRequest
//
// @return DeleteUserResponse
func (client *Client) DeleteUser(request *DeleteUserRequest) (_result *DeleteUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserResponse{}
	_body, _err := client.DeleteUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables an enabled Employee Identity and Access Management (EIAM) application. All features of the EIAM application cannot be used if you disable the EIAM application.
//
// Description:
//
// All features of the EIAM application cannot be used if you disable the EIAM application, such as single sign-on (SSO) and account synchronization. Make sure that you acknowledge the risks of the delete operation.
//
// @param request - DisableApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableApplicationResponse
func (client *Client) DisableApplicationWithOptions(request *DisableApplicationRequest, runtime *util.RuntimeOptions) (_result *DisableApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableApplication"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables an enabled Employee Identity and Access Management (EIAM) application. All features of the EIAM application cannot be used if you disable the EIAM application.
//
// Description:
//
// All features of the EIAM application cannot be used if you disable the EIAM application, such as single sign-on (SSO) and account synchronization. Make sure that you acknowledge the risks of the delete operation.
//
// @param request - DisableApplicationRequest
//
// @return DisableApplicationResponse
func (client *Client) DisableApplication(request *DisableApplicationRequest) (_result *DisableApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableApplicationResponse{}
	_body, _err := client.DisableApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the Developer API feature for an Employee Identity and Access Management (EIAM) application.
//
// @param request - DisableApplicationApiInvokeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableApplicationApiInvokeResponse
func (client *Client) DisableApplicationApiInvokeWithOptions(request *DisableApplicationApiInvokeRequest, runtime *util.RuntimeOptions) (_result *DisableApplicationApiInvokeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableApplicationApiInvoke"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableApplicationApiInvokeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the Developer API feature for an Employee Identity and Access Management (EIAM) application.
//
// @param request - DisableApplicationApiInvokeRequest
//
// @return DisableApplicationApiInvokeResponse
func (client *Client) DisableApplicationApiInvoke(request *DisableApplicationApiInvokeRequest) (_result *DisableApplicationApiInvokeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableApplicationApiInvokeResponse{}
	_body, _err := client.DisableApplicationApiInvokeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables a client key of an Employee Identity and Access Management (EIAM) application.
//
// @param request - DisableApplicationClientSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableApplicationClientSecretResponse
func (client *Client) DisableApplicationClientSecretWithOptions(request *DisableApplicationClientSecretRequest, runtime *util.RuntimeOptions) (_result *DisableApplicationClientSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretId)) {
		query["SecretId"] = request.SecretId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableApplicationClientSecret"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableApplicationClientSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a client key of an Employee Identity and Access Management (EIAM) application.
//
// @param request - DisableApplicationClientSecretRequest
//
// @return DisableApplicationClientSecretResponse
func (client *Client) DisableApplicationClientSecret(request *DisableApplicationClientSecretRequest) (_result *DisableApplicationClientSecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableApplicationClientSecretResponse{}
	_body, _err := client.DisableApplicationClientSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables the account synchronization feature for an application in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// @param request - DisableApplicationProvisioningRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableApplicationProvisioningResponse
func (client *Client) DisableApplicationProvisioningWithOptions(request *DisableApplicationProvisioningRequest, runtime *util.RuntimeOptions) (_result *DisableApplicationProvisioningResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableApplicationProvisioning"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableApplicationProvisioningResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the account synchronization feature for an application in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// @param request - DisableApplicationProvisioningRequest
//
// @return DisableApplicationProvisioningResponse
func (client *Client) DisableApplicationProvisioning(request *DisableApplicationProvisioningRequest) (_result *DisableApplicationProvisioningResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableApplicationProvisioningResponse{}
	_body, _err := client.DisableApplicationProvisioningWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 禁用应用SSO能力
//
// @param request - DisableApplicationSsoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableApplicationSsoResponse
func (client *Client) DisableApplicationSsoWithOptions(request *DisableApplicationSsoRequest, runtime *util.RuntimeOptions) (_result *DisableApplicationSsoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableApplicationSso"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableApplicationSsoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 禁用应用SSO能力
//
// @param request - DisableApplicationSsoRequest
//
// @return DisableApplicationSsoResponse
func (client *Client) DisableApplicationSso(request *DisableApplicationSsoRequest) (_result *DisableApplicationSsoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableApplicationSsoResponse{}
	_body, _err := client.DisableApplicationSsoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 禁用指定域名代理Token。
//
// @param request - DisableDomainProxyTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableDomainProxyTokenResponse
func (client *Client) DisableDomainProxyTokenWithOptions(request *DisableDomainProxyTokenRequest, runtime *util.RuntimeOptions) (_result *DisableDomainProxyTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainId)) {
		query["DomainId"] = request.DomainId
	}

	if !tea.BoolValue(util.IsUnset(request.DomainProxyTokenId)) {
		query["DomainProxyTokenId"] = request.DomainProxyTokenId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableDomainProxyToken"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableDomainProxyTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 禁用指定域名代理Token。
//
// @param request - DisableDomainProxyTokenRequest
//
// @return DisableDomainProxyTokenResponse
func (client *Client) DisableDomainProxyToken(request *DisableDomainProxyTokenRequest) (_result *DisableDomainProxyTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableDomainProxyTokenResponse{}
	_body, _err := client.DisableDomainProxyTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 关闭初始化域名自动跳转。
//
// @param request - DisableInitDomainAutoRedirectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableInitDomainAutoRedirectResponse
func (client *Client) DisableInitDomainAutoRedirectWithOptions(request *DisableInitDomainAutoRedirectRequest, runtime *util.RuntimeOptions) (_result *DisableInitDomainAutoRedirectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableInitDomainAutoRedirect"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableInitDomainAutoRedirectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关闭初始化域名自动跳转。
//
// @param request - DisableInitDomainAutoRedirectRequest
//
// @return DisableInitDomainAutoRedirectResponse
func (client *Client) DisableInitDomainAutoRedirect(request *DisableInitDomainAutoRedirectRequest) (_result *DisableInitDomainAutoRedirectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableInitDomainAutoRedirectResponse{}
	_body, _err := client.DisableInitDomainAutoRedirectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables an Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM) account. If the account is disabled, a success message is returned.
//
// @param request - DisableUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableUserResponse
func (client *Client) DisableUserWithOptions(request *DisableUserRequest, runtime *util.RuntimeOptions) (_result *DisableUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableUser"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables an Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM) account. If the account is disabled, a success message is returned.
//
// @param request - DisableUserRequest
//
// @return DisableUserResponse
func (client *Client) DisableUser(request *DisableUserRequest) (_result *DisableUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableUserResponse{}
	_body, _err := client.DisableUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables a disabled Employee Identity and Access Management (EIAM) application.
//
// @param request - EnableApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableApplicationResponse
func (client *Client) EnableApplicationWithOptions(request *EnableApplicationRequest, runtime *util.RuntimeOptions) (_result *EnableApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableApplication"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a disabled Employee Identity and Access Management (EIAM) application.
//
// @param request - EnableApplicationRequest
//
// @return EnableApplicationResponse
func (client *Client) EnableApplication(request *EnableApplicationRequest) (_result *EnableApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableApplicationResponse{}
	_body, _err := client.EnableApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the Developer API feature for an Employee Identity and Access Management (EIAM) application.
//
// @param request - EnableApplicationApiInvokeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableApplicationApiInvokeResponse
func (client *Client) EnableApplicationApiInvokeWithOptions(request *EnableApplicationApiInvokeRequest, runtime *util.RuntimeOptions) (_result *EnableApplicationApiInvokeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableApplicationApiInvoke"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableApplicationApiInvokeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the Developer API feature for an Employee Identity and Access Management (EIAM) application.
//
// @param request - EnableApplicationApiInvokeRequest
//
// @return EnableApplicationApiInvokeResponse
func (client *Client) EnableApplicationApiInvoke(request *EnableApplicationApiInvokeRequest) (_result *EnableApplicationApiInvokeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableApplicationApiInvokeResponse{}
	_body, _err := client.EnableApplicationApiInvokeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the client key of an application in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// @param request - EnableApplicationClientSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableApplicationClientSecretResponse
func (client *Client) EnableApplicationClientSecretWithOptions(request *EnableApplicationClientSecretRequest, runtime *util.RuntimeOptions) (_result *EnableApplicationClientSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretId)) {
		query["SecretId"] = request.SecretId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableApplicationClientSecret"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableApplicationClientSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the client key of an application in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// @param request - EnableApplicationClientSecretRequest
//
// @return EnableApplicationClientSecretResponse
func (client *Client) EnableApplicationClientSecret(request *EnableApplicationClientSecretRequest) (_result *EnableApplicationClientSecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableApplicationClientSecretResponse{}
	_body, _err := client.EnableApplicationClientSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the account synchronization feature for an application in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// @param request - EnableApplicationProvisioningRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableApplicationProvisioningResponse
func (client *Client) EnableApplicationProvisioningWithOptions(request *EnableApplicationProvisioningRequest, runtime *util.RuntimeOptions) (_result *EnableApplicationProvisioningResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableApplicationProvisioning"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableApplicationProvisioningResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the account synchronization feature for an application in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// @param request - EnableApplicationProvisioningRequest
//
// @return EnableApplicationProvisioningResponse
func (client *Client) EnableApplicationProvisioning(request *EnableApplicationProvisioningRequest) (_result *EnableApplicationProvisioningResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableApplicationProvisioningResponse{}
	_body, _err := client.EnableApplicationProvisioningWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启用应用SSO能力
//
// @param request - EnableApplicationSsoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableApplicationSsoResponse
func (client *Client) EnableApplicationSsoWithOptions(request *EnableApplicationSsoRequest, runtime *util.RuntimeOptions) (_result *EnableApplicationSsoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableApplicationSso"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableApplicationSsoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启用应用SSO能力
//
// @param request - EnableApplicationSsoRequest
//
// @return EnableApplicationSsoResponse
func (client *Client) EnableApplicationSso(request *EnableApplicationSsoRequest) (_result *EnableApplicationSsoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableApplicationSsoResponse{}
	_body, _err := client.EnableApplicationSsoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启用指定域名代理Token。
//
// @param request - EnableDomainProxyTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableDomainProxyTokenResponse
func (client *Client) EnableDomainProxyTokenWithOptions(request *EnableDomainProxyTokenRequest, runtime *util.RuntimeOptions) (_result *EnableDomainProxyTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainId)) {
		query["DomainId"] = request.DomainId
	}

	if !tea.BoolValue(util.IsUnset(request.DomainProxyTokenId)) {
		query["DomainProxyTokenId"] = request.DomainProxyTokenId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableDomainProxyToken"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableDomainProxyTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启用指定域名代理Token。
//
// @param request - EnableDomainProxyTokenRequest
//
// @return EnableDomainProxyTokenResponse
func (client *Client) EnableDomainProxyToken(request *EnableDomainProxyTokenRequest) (_result *EnableDomainProxyTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableDomainProxyTokenResponse{}
	_body, _err := client.EnableDomainProxyTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开启初始化域名自动跳转。开启后，访问初始化域名将会自动跳转至默认域名。
//
// @param request - EnableInitDomainAutoRedirectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableInitDomainAutoRedirectResponse
func (client *Client) EnableInitDomainAutoRedirectWithOptions(request *EnableInitDomainAutoRedirectRequest, runtime *util.RuntimeOptions) (_result *EnableInitDomainAutoRedirectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableInitDomainAutoRedirect"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableInitDomainAutoRedirectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启初始化域名自动跳转。开启后，访问初始化域名将会自动跳转至默认域名。
//
// @param request - EnableInitDomainAutoRedirectRequest
//
// @return EnableInitDomainAutoRedirectResponse
func (client *Client) EnableInitDomainAutoRedirect(request *EnableInitDomainAutoRedirectRequest) (_result *EnableInitDomainAutoRedirectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableInitDomainAutoRedirectResponse{}
	_body, _err := client.EnableInitDomainAutoRedirectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables an Employee Identity and Access Management (EIAM) account of Identity as a Service (IDaaS).
//
// @param request - EnableUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableUserResponse
func (client *Client) EnableUserWithOptions(request *EnableUserRequest, runtime *util.RuntimeOptions) (_result *EnableUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableUser"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables an Employee Identity and Access Management (EIAM) account of Identity as a Service (IDaaS).
//
// @param request - EnableUserRequest
//
// @return EnableUserResponse
func (client *Client) EnableUser(request *EnableUserRequest) (_result *EnableUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableUserResponse{}
	_body, _err := client.EnableUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an Employee Identity and Access Management (EIAM) application.
//
// @param request - GetApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationResponse
func (client *Client) GetApplicationWithOptions(request *GetApplicationRequest, runtime *util.RuntimeOptions) (_result *GetApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetApplication"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an Employee Identity and Access Management (EIAM) application.
//
// @param request - GetApplicationRequest
//
// @return GetApplicationResponse
func (client *Client) GetApplication(request *GetApplicationRequest) (_result *GetApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetApplicationResponse{}
	_body, _err := client.GetApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the permissions of the Developer API feature for an Employee Identity and Access Management (EIAM) application.
//
// @param request - GetApplicationGrantScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationGrantScopeResponse
func (client *Client) GetApplicationGrantScopeWithOptions(request *GetApplicationGrantScopeRequest, runtime *util.RuntimeOptions) (_result *GetApplicationGrantScopeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetApplicationGrantScope"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetApplicationGrantScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the permissions of the Developer API feature for an Employee Identity and Access Management (EIAM) application.
//
// @param request - GetApplicationGrantScopeRequest
//
// @return GetApplicationGrantScopeResponse
func (client *Client) GetApplicationGrantScope(request *GetApplicationGrantScopeRequest) (_result *GetApplicationGrantScopeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetApplicationGrantScopeResponse{}
	_body, _err := client.GetApplicationGrantScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configuration of the account synchronization feature for an application in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - GetApplicationProvisioningConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationProvisioningConfigResponse
func (client *Client) GetApplicationProvisioningConfigWithOptions(request *GetApplicationProvisioningConfigRequest, runtime *util.RuntimeOptions) (_result *GetApplicationProvisioningConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetApplicationProvisioningConfig"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetApplicationProvisioningConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration of the account synchronization feature for an application in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - GetApplicationProvisioningConfigRequest
//
// @return GetApplicationProvisioningConfigResponse
func (client *Client) GetApplicationProvisioningConfig(request *GetApplicationProvisioningConfigRequest) (_result *GetApplicationProvisioningConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetApplicationProvisioningConfigResponse{}
	_body, _err := client.GetApplicationProvisioningConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the account synchronization scope of applications in Identity as a Service (IDaaS) Employee IAM (EIAM). This scope is the same as the scope within which developers can call the DeveloperAPI to query and manage accounts.
//
// @param request - GetApplicationProvisioningScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationProvisioningScopeResponse
func (client *Client) GetApplicationProvisioningScopeWithOptions(request *GetApplicationProvisioningScopeRequest, runtime *util.RuntimeOptions) (_result *GetApplicationProvisioningScopeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetApplicationProvisioningScope"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetApplicationProvisioningScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the account synchronization scope of applications in Identity as a Service (IDaaS) Employee IAM (EIAM). This scope is the same as the scope within which developers can call the DeveloperAPI to query and manage accounts.
//
// @param request - GetApplicationProvisioningScopeRequest
//
// @return GetApplicationProvisioningScopeResponse
func (client *Client) GetApplicationProvisioningScope(request *GetApplicationProvisioningScopeRequest) (_result *GetApplicationProvisioningScopeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetApplicationProvisioningScopeResponse{}
	_body, _err := client.GetApplicationProvisioningScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the single sign-on (SSO) configuration attributes of an application in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// @param request - GetApplicationSsoConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationSsoConfigResponse
func (client *Client) GetApplicationSsoConfigWithOptions(request *GetApplicationSsoConfigRequest, runtime *util.RuntimeOptions) (_result *GetApplicationSsoConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetApplicationSsoConfig"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetApplicationSsoConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the single sign-on (SSO) configuration attributes of an application in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// @param request - GetApplicationSsoConfigRequest
//
// @return GetApplicationSsoConfigResponse
func (client *Client) GetApplicationSsoConfig(request *GetApplicationSsoConfigRequest) (_result *GetApplicationSsoConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetApplicationSsoConfigResponse{}
	_body, _err := client.GetApplicationSsoConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询一个域名的详细信息。
//
// @param request - GetDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDomainResponse
func (client *Client) GetDomainWithOptions(request *GetDomainRequest, runtime *util.RuntimeOptions) (_result *GetDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainId)) {
		query["DomainId"] = request.DomainId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDomain"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询一个域名的详细信息。
//
// @param request - GetDomainRequest
//
// @return GetDomainResponse
func (client *Client) GetDomain(request *GetDomainRequest) (_result *GetDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDomainResponse{}
	_body, _err := client.GetDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看域名的DNS Challenge记录。
//
// @param request - GetDomainDnsChallengeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDomainDnsChallengeResponse
func (client *Client) GetDomainDnsChallengeWithOptions(request *GetDomainDnsChallengeRequest, runtime *util.RuntimeOptions) (_result *GetDomainDnsChallengeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDomainDnsChallenge"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDomainDnsChallengeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看域名的DNS Challenge记录。
//
// @param request - GetDomainDnsChallengeRequest
//
// @return GetDomainDnsChallengeResponse
func (client *Client) GetDomainDnsChallenge(request *GetDomainDnsChallengeRequest) (_result *GetDomainDnsChallengeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDomainDnsChallengeResponse{}
	_body, _err := client.GetDomainDnsChallengeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the forgot password configurations of an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - GetForgetPasswordConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetForgetPasswordConfigurationResponse
func (client *Client) GetForgetPasswordConfigurationWithOptions(request *GetForgetPasswordConfigurationRequest, runtime *util.RuntimeOptions) (_result *GetForgetPasswordConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetForgetPasswordConfiguration"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetForgetPasswordConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the forgot password configurations of an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - GetForgetPasswordConfigurationRequest
//
// @return GetForgetPasswordConfigurationResponse
func (client *Client) GetForgetPasswordConfiguration(request *GetForgetPasswordConfigurationRequest) (_result *GetForgetPasswordConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetForgetPasswordConfigurationResponse{}
	_body, _err := client.GetForgetPasswordConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of an account group in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - GetGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGroupResponse
func (client *Client) GetGroupWithOptions(request *GetGroupRequest, runtime *util.RuntimeOptions) (_result *GetGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGroup"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of an account group in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - GetGroupRequest
//
// @return GetGroupResponse
func (client *Client) GetGroup(request *GetGroupRequest) (_result *GetGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGroupResponse{}
	_body, _err := client.GetGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of an Enterprise Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - GetInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithOptions(request *GetInstanceRequest, runtime *util.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstance"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of an Enterprise Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - GetInstanceRequest
//
// @return GetInstanceResponse
func (client *Client) GetInstance(request *GetInstanceRequest) (_result *GetInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询一个专属网络端点的详细信息。
//
// @param request - GetNetworkAccessEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNetworkAccessEndpointResponse
func (client *Client) GetNetworkAccessEndpointWithOptions(request *GetNetworkAccessEndpointRequest, runtime *util.RuntimeOptions) (_result *GetNetworkAccessEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkAccessEndpointId)) {
		query["NetworkAccessEndpointId"] = request.NetworkAccessEndpointId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNetworkAccessEndpoint"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNetworkAccessEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询一个专属网络端点的详细信息。
//
// @param request - GetNetworkAccessEndpointRequest
//
// @return GetNetworkAccessEndpointResponse
func (client *Client) GetNetworkAccessEndpoint(request *GetNetworkAccessEndpointRequest) (_result *GetNetworkAccessEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNetworkAccessEndpointResponse{}
	_body, _err := client.GetNetworkAccessEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an organizational unit in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// @param request - GetOrganizationalUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrganizationalUnitResponse
func (client *Client) GetOrganizationalUnitWithOptions(request *GetOrganizationalUnitRequest, runtime *util.RuntimeOptions) (_result *GetOrganizationalUnitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitId)) {
		query["OrganizationalUnitId"] = request.OrganizationalUnitId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOrganizationalUnit"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOrganizationalUnitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an organizational unit in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// @param request - GetOrganizationalUnitRequest
//
// @return GetOrganizationalUnitResponse
func (client *Client) GetOrganizationalUnit(request *GetOrganizationalUnitRequest) (_result *GetOrganizationalUnitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOrganizationalUnitResponse{}
	_body, _err := client.GetOrganizationalUnitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the password complexity configurations of an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - GetPasswordComplexityConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPasswordComplexityConfigurationResponse
func (client *Client) GetPasswordComplexityConfigurationWithOptions(request *GetPasswordComplexityConfigurationRequest, runtime *util.RuntimeOptions) (_result *GetPasswordComplexityConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPasswordComplexityConfiguration"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPasswordComplexityConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the password complexity configurations of an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - GetPasswordComplexityConfigurationRequest
//
// @return GetPasswordComplexityConfigurationResponse
func (client *Client) GetPasswordComplexityConfiguration(request *GetPasswordComplexityConfigurationRequest) (_result *GetPasswordComplexityConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPasswordComplexityConfigurationResponse{}
	_body, _err := client.GetPasswordComplexityConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the password expiration configurations of an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - GetPasswordExpirationConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPasswordExpirationConfigurationResponse
func (client *Client) GetPasswordExpirationConfigurationWithOptions(request *GetPasswordExpirationConfigurationRequest, runtime *util.RuntimeOptions) (_result *GetPasswordExpirationConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPasswordExpirationConfiguration"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPasswordExpirationConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the password expiration configurations of an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - GetPasswordExpirationConfigurationRequest
//
// @return GetPasswordExpirationConfigurationResponse
func (client *Client) GetPasswordExpirationConfiguration(request *GetPasswordExpirationConfigurationRequest) (_result *GetPasswordExpirationConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPasswordExpirationConfigurationResponse{}
	_body, _err := client.GetPasswordExpirationConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the password history configurations of an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - GetPasswordHistoryConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPasswordHistoryConfigurationResponse
func (client *Client) GetPasswordHistoryConfigurationWithOptions(request *GetPasswordHistoryConfigurationRequest, runtime *util.RuntimeOptions) (_result *GetPasswordHistoryConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPasswordHistoryConfiguration"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPasswordHistoryConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the password history configurations of an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - GetPasswordHistoryConfigurationRequest
//
// @return GetPasswordHistoryConfigurationResponse
func (client *Client) GetPasswordHistoryConfiguration(request *GetPasswordHistoryConfigurationRequest) (_result *GetPasswordHistoryConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPasswordHistoryConfigurationResponse{}
	_body, _err := client.GetPasswordHistoryConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the password initialization configurations of an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - GetPasswordInitializationConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPasswordInitializationConfigurationResponse
func (client *Client) GetPasswordInitializationConfigurationWithOptions(request *GetPasswordInitializationConfigurationRequest, runtime *util.RuntimeOptions) (_result *GetPasswordInitializationConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPasswordInitializationConfiguration"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPasswordInitializationConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the password initialization configurations of an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - GetPasswordInitializationConfigurationRequest
//
// @return GetPasswordInitializationConfigurationResponse
func (client *Client) GetPasswordInitializationConfiguration(request *GetPasswordInitializationConfigurationRequest) (_result *GetPasswordInitializationConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPasswordInitializationConfigurationResponse{}
	_body, _err := client.GetPasswordInitializationConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the root organizational unit in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// @param request - GetRootOrganizationalUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRootOrganizationalUnitResponse
func (client *Client) GetRootOrganizationalUnitWithOptions(request *GetRootOrganizationalUnitRequest, runtime *util.RuntimeOptions) (_result *GetRootOrganizationalUnitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRootOrganizationalUnit"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRootOrganizationalUnitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the root organizational unit in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// @param request - GetRootOrganizationalUnitRequest
//
// @return GetRootOrganizationalUnitResponse
func (client *Client) GetRootOrganizationalUnit(request *GetRootOrganizationalUnitRequest) (_result *GetRootOrganizationalUnitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRootOrganizationalUnitResponse{}
	_body, _err := client.GetRootOrganizationalUnitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询同步任务
//
// @param request - GetSynchronizationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSynchronizationJobResponse
func (client *Client) GetSynchronizationJobWithOptions(request *GetSynchronizationJobRequest, runtime *util.RuntimeOptions) (_result *GetSynchronizationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationJobId)) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSynchronizationJob"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSynchronizationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询同步任务
//
// @param request - GetSynchronizationJobRequest
//
// @return GetSynchronizationJobResponse
func (client *Client) GetSynchronizationJob(request *GetSynchronizationJobRequest) (_result *GetSynchronizationJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSynchronizationJobResponse{}
	_body, _err := client.GetSynchronizationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an account in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// @param request - GetUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserResponse
func (client *Client) GetUserWithOptions(request *GetUserRequest, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUser"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an account in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// @param request - GetUserRequest
//
// @return GetUserResponse
func (client *Client) GetUser(request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all client keys of an Employee Identity and Access Management (EIAM) application. The returned key secret is not masked. If you want to query the key secret that is masked, call the ObtainApplicationClientSecret operation.
//
// @param request - ListApplicationClientSecretsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationClientSecretsResponse
func (client *Client) ListApplicationClientSecretsWithOptions(request *ListApplicationClientSecretsRequest, runtime *util.RuntimeOptions) (_result *ListApplicationClientSecretsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApplicationClientSecrets"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListApplicationClientSecretsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all client keys of an Employee Identity and Access Management (EIAM) application. The returned key secret is not masked. If you want to query the key secret that is masked, call the ObtainApplicationClientSecret operation.
//
// @param request - ListApplicationClientSecretsRequest
//
// @return ListApplicationClientSecretsResponse
func (client *Client) ListApplicationClientSecrets(request *ListApplicationClientSecretsRequest) (_result *ListApplicationClientSecretsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApplicationClientSecretsResponse{}
	_body, _err := client.ListApplicationClientSecretsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about one or multiple Employee Identity and Access Management (EIAM) applications by page.
//
// @param request - ListApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsResponse
func (client *Client) ListApplicationsWithOptions(request *ListApplicationsRequest, runtime *util.RuntimeOptions) (_result *ListApplicationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationIds)) {
		query["ApplicationIds"] = request.ApplicationIds
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationName)) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizationType)) {
		query["AuthorizationType"] = request.AuthorizationType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApplications"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about one or multiple Employee Identity and Access Management (EIAM) applications by page.
//
// @param request - ListApplicationsRequest
//
// @return ListApplicationsResponse
func (client *Client) ListApplications(request *ListApplicationsRequest) (_result *ListApplicationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApplicationsResponse{}
	_body, _err := client.ListApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the applications that an Employee Identity and Access Management (EIAM) organization can access. The return result includes the IDs of the applications. If you want to obtain the details of the applications, call the GetApplication operation.
//
// Description:
//
// You can only query the permissions that are directly granted to the EIAM organization by calling the ListApplicationsForOrganizationalUnit operation. You can filter applications by configuring the **ApplicationIds*	- parameter when you call this operation.
//
// @param request - ListApplicationsForOrganizationalUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsForOrganizationalUnitResponse
func (client *Client) ListApplicationsForOrganizationalUnitWithOptions(request *ListApplicationsForOrganizationalUnitRequest, runtime *util.RuntimeOptions) (_result *ListApplicationsForOrganizationalUnitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationIds)) {
		query["ApplicationIds"] = request.ApplicationIds
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitId)) {
		query["OrganizationalUnitId"] = request.OrganizationalUnitId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApplicationsForOrganizationalUnit"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListApplicationsForOrganizationalUnitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the applications that an Employee Identity and Access Management (EIAM) organization can access. The return result includes the IDs of the applications. If you want to obtain the details of the applications, call the GetApplication operation.
//
// Description:
//
// You can only query the permissions that are directly granted to the EIAM organization by calling the ListApplicationsForOrganizationalUnit operation. You can filter applications by configuring the **ApplicationIds*	- parameter when you call this operation.
//
// @param request - ListApplicationsForOrganizationalUnitRequest
//
// @return ListApplicationsForOrganizationalUnitResponse
func (client *Client) ListApplicationsForOrganizationalUnit(request *ListApplicationsForOrganizationalUnitRequest) (_result *ListApplicationsForOrganizationalUnitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApplicationsForOrganizationalUnitResponse{}
	_body, _err := client.ListApplicationsForOrganizationalUnitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the applications that an Employee Identity and Access Management (EIAM) account can access. The return result includes the IDs of the applications. If you want to obtain the details of the applications, call the GetApplication operation.
//
// @param request - ListApplicationsForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsForUserResponse
func (client *Client) ListApplicationsForUserWithOptions(request *ListApplicationsForUserRequest, runtime *util.RuntimeOptions) (_result *ListApplicationsForUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationIds)) {
		query["ApplicationIds"] = request.ApplicationIds
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryMode)) {
		query["QueryMode"] = request.QueryMode
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApplicationsForUser"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListApplicationsForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the applications that an Employee Identity and Access Management (EIAM) account can access. The return result includes the IDs of the applications. If you want to obtain the details of the applications, call the GetApplication operation.
//
// @param request - ListApplicationsForUserRequest
//
// @return ListApplicationsForUserResponse
func (client *Client) ListApplicationsForUser(request *ListApplicationsForUserRequest) (_result *ListApplicationsForUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApplicationsForUserResponse{}
	_body, _err := client.ListApplicationsForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列表查询域名代理Token信息。
//
// @param request - ListDomainProxyTokensRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDomainProxyTokensResponse
func (client *Client) ListDomainProxyTokensWithOptions(request *ListDomainProxyTokensRequest, runtime *util.RuntimeOptions) (_result *ListDomainProxyTokensResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainId)) {
		query["DomainId"] = request.DomainId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDomainProxyTokens"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDomainProxyTokensResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列表查询域名代理Token信息。
//
// @param request - ListDomainProxyTokensRequest
//
// @return ListDomainProxyTokensResponse
func (client *Client) ListDomainProxyTokens(request *ListDomainProxyTokensRequest) (_result *ListDomainProxyTokensResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDomainProxyTokensResponse{}
	_body, _err := client.ListDomainProxyTokensWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列表查询域名记录。
//
// @param request - ListDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDomainsResponse
func (client *Client) ListDomainsWithOptions(request *ListDomainsRequest, runtime *util.RuntimeOptions) (_result *ListDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDomains"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列表查询域名记录。
//
// @param request - ListDomainsRequest
//
// @return ListDomainsResponse
func (client *Client) ListDomains(request *ListDomainsRequest) (_result *ListDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDomainsResponse{}
	_body, _err := client.ListDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询EIAM2.0/EIAM1.0实例列表
//
// @param request - ListEiamInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEiamInstancesResponse
func (client *Client) ListEiamInstancesWithOptions(request *ListEiamInstancesRequest, runtime *util.RuntimeOptions) (_result *ListEiamInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceRegionId)) {
		query["InstanceRegionId"] = request.InstanceRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEiamInstances"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEiamInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询EIAM2.0/EIAM1.0实例列表
//
// @param request - ListEiamInstancesRequest
//
// @return ListEiamInstancesResponse
func (client *Client) ListEiamInstances(request *ListEiamInstancesRequest) (_result *ListEiamInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEiamInstancesResponse{}
	_body, _err := client.ListEiamInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询EIAM2.0/EIAM1.0地域列表
//
// @param request - ListEiamRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEiamRegionsResponse
func (client *Client) ListEiamRegionsWithOptions(runtime *util.RuntimeOptions) (_result *ListEiamRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListEiamRegions"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEiamRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询EIAM2.0/EIAM1.0地域列表
//
// @return ListEiamRegionsResponse
func (client *Client) ListEiamRegions() (_result *ListEiamRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEiamRegionsResponse{}
	_body, _err := client.ListEiamRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of account groups in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - ListGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsResponse
func (client *Client) ListGroupsWithOptions(request *ListGroupsRequest, runtime *util.RuntimeOptions) (_result *ListGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupExternalId)) {
		query["GroupExternalId"] = request.GroupExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupIds)) {
		query["GroupIds"] = request.GroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupNameStartsWith)) {
		query["GroupNameStartsWith"] = request.GroupNameStartsWith
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGroups"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of account groups in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - ListGroupsRequest
//
// @return ListGroupsResponse
func (client *Client) ListGroups(request *ListGroupsRequest) (_result *ListGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGroupsResponse{}
	_body, _err := client.ListGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of account groups to which the permissions to access an application are granted. The returned results contain the group IDs. You can call the GetGroup operation to query the information about an account group based on the group ID.
//
// @param request - ListGroupsForApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsForApplicationResponse
func (client *Client) ListGroupsForApplicationWithOptions(request *ListGroupsForApplicationRequest, runtime *util.RuntimeOptions) (_result *ListGroupsForApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupIds)) {
		query["GroupIds"] = request.GroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGroupsForApplication"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGroupsForApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of account groups to which the permissions to access an application are granted. The returned results contain the group IDs. You can call the GetGroup operation to query the information about an account group based on the group ID.
//
// @param request - ListGroupsForApplicationRequest
//
// @return ListGroupsForApplicationResponse
func (client *Client) ListGroupsForApplication(request *ListGroupsForApplicationRequest) (_result *ListGroupsForApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGroupsForApplicationResponse{}
	_body, _err := client.ListGroupsForApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of account groups to which an Employee Identity and Access Management (EIAM) account of Identity as a Service (IDaaS) belongs.
//
// @param request - ListGroupsForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupsForUserResponse
func (client *Client) ListGroupsForUserWithOptions(request *ListGroupsForUserRequest, runtime *util.RuntimeOptions) (_result *ListGroupsForUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGroupsForUser"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGroupsForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of account groups to which an Employee Identity and Access Management (EIAM) account of Identity as a Service (IDaaS) belongs.
//
// @param request - ListGroupsForUserRequest
//
// @return ListGroupsForUserResponse
func (client *Client) ListGroupsForUser(request *ListGroupsForUserRequest) (_result *ListGroupsForUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGroupsForUserResponse{}
	_body, _err := client.ListGroupsForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of one or more Enterprise Identity and Access Management (EIAM) instances of Identity as a Service (IDaaS).
//
// @param request - ListInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstances"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of one or more Enterprise Identity and Access Management (EIAM) instances of Identity as a Service (IDaaS).
//
// @param request - ListInstancesRequest
//
// @return ListInstancesResponse
func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取支持专属端点的region列表
//
// @param request - ListNetworkAccessEndpointAvailableRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNetworkAccessEndpointAvailableRegionsResponse
func (client *Client) ListNetworkAccessEndpointAvailableRegionsWithOptions(runtime *util.RuntimeOptions) (_result *ListNetworkAccessEndpointAvailableRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListNetworkAccessEndpointAvailableRegions"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNetworkAccessEndpointAvailableRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取支持专属端点的region列表
//
// @return ListNetworkAccessEndpointAvailableRegionsResponse
func (client *Client) ListNetworkAccessEndpointAvailableRegions() (_result *ListNetworkAccessEndpointAvailableRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNetworkAccessEndpointAvailableRegionsResponse{}
	_body, _err := client.ListNetworkAccessEndpointAvailableRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取支持NAE的可用区列表
//
// @param request - ListNetworkAccessEndpointAvailableZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNetworkAccessEndpointAvailableZonesResponse
func (client *Client) ListNetworkAccessEndpointAvailableZonesWithOptions(request *ListNetworkAccessEndpointAvailableZonesRequest, runtime *util.RuntimeOptions) (_result *ListNetworkAccessEndpointAvailableZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NaeRegionId)) {
		query["NaeRegionId"] = request.NaeRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNetworkAccessEndpointAvailableZones"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNetworkAccessEndpointAvailableZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取支持NAE的可用区列表
//
// @param request - ListNetworkAccessEndpointAvailableZonesRequest
//
// @return ListNetworkAccessEndpointAvailableZonesResponse
func (client *Client) ListNetworkAccessEndpointAvailableZones(request *ListNetworkAccessEndpointAvailableZonesRequest) (_result *ListNetworkAccessEndpointAvailableZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNetworkAccessEndpointAvailableZonesResponse{}
	_body, _err := client.ListNetworkAccessEndpointAvailableZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列表查询专属网络端点。
//
// @param request - ListNetworkAccessEndpointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNetworkAccessEndpointsResponse
func (client *Client) ListNetworkAccessEndpointsWithOptions(request *ListNetworkAccessEndpointsRequest, runtime *util.RuntimeOptions) (_result *ListNetworkAccessEndpointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkAccessEndpointStatus)) {
		query["NetworkAccessEndpointStatus"] = request.NetworkAccessEndpointStatus
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkAccessEndpointType)) {
		query["NetworkAccessEndpointType"] = request.NetworkAccessEndpointType
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcRegionId)) {
		query["VpcRegionId"] = request.VpcRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNetworkAccessEndpoints"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNetworkAccessEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列表查询专属网络端点。
//
// @param request - ListNetworkAccessEndpointsRequest
//
// @return ListNetworkAccessEndpointsResponse
func (client *Client) ListNetworkAccessEndpoints(request *ListNetworkAccessEndpointsRequest) (_result *ListNetworkAccessEndpointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNetworkAccessEndpointsResponse{}
	_body, _err := client.ListNetworkAccessEndpointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列表某个网络访问端点下的访问路径。
//
// @param request - ListNetworkAccessPathsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNetworkAccessPathsResponse
func (client *Client) ListNetworkAccessPathsWithOptions(request *ListNetworkAccessPathsRequest, runtime *util.RuntimeOptions) (_result *ListNetworkAccessPathsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkAccessEndpointId)) {
		query["NetworkAccessEndpointId"] = request.NetworkAccessEndpointId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNetworkAccessPaths"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNetworkAccessPathsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列表某个网络访问端点下的访问路径。
//
// @param request - ListNetworkAccessPathsRequest
//
// @return ListNetworkAccessPathsResponse
func (client *Client) ListNetworkAccessPaths(request *ListNetworkAccessPathsRequest) (_result *ListNetworkAccessPathsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNetworkAccessPathsResponse{}
	_body, _err := client.ListNetworkAccessPathsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询组织的所有父级路径
//
// @param request - ListOrganizationalUnitParentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOrganizationalUnitParentsResponse
func (client *Client) ListOrganizationalUnitParentsWithOptions(request *ListOrganizationalUnitParentsRequest, runtime *util.RuntimeOptions) (_result *ListOrganizationalUnitParentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitId)) {
		query["OrganizationalUnitId"] = request.OrganizationalUnitId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOrganizationalUnitParents"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOrganizationalUnitParentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询组织的所有父级路径
//
// @param request - ListOrganizationalUnitParentsRequest
//
// @return ListOrganizationalUnitParentsResponse
func (client *Client) ListOrganizationalUnitParents(request *ListOrganizationalUnitParentsRequest) (_result *ListOrganizationalUnitParentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOrganizationalUnitParentsResponse{}
	_body, _err := client.ListOrganizationalUnitParentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about organizational units in Identity as a Service (IDaaS) Employee IAM (EIAM) by page.
//
// @param request - ListOrganizationalUnitsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOrganizationalUnitsResponse
func (client *Client) ListOrganizationalUnitsWithOptions(request *ListOrganizationalUnitsRequest, runtime *util.RuntimeOptions) (_result *ListOrganizationalUnitsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitIds)) {
		query["OrganizationalUnitIds"] = request.OrganizationalUnitIds
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitName)) {
		query["OrganizationalUnitName"] = request.OrganizationalUnitName
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitNameStartsWith)) {
		query["OrganizationalUnitNameStartsWith"] = request.OrganizationalUnitNameStartsWith
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		query["ParentId"] = request.ParentId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOrganizationalUnits"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOrganizationalUnitsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about organizational units in Identity as a Service (IDaaS) Employee IAM (EIAM) by page.
//
// @param request - ListOrganizationalUnitsRequest
//
// @return ListOrganizationalUnitsResponse
func (client *Client) ListOrganizationalUnits(request *ListOrganizationalUnitsRequest) (_result *ListOrganizationalUnitsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOrganizationalUnitsResponse{}
	_body, _err := client.ListOrganizationalUnitsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the organizations that are allowed to access an Employee Identity and Access Management (EIAM) application by page. The return result includes the IDs of the organizations. If you want to obtain the details of the organizations, call the GetOrganizationalUnit operation.
//
// @param request - ListOrganizationalUnitsForApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOrganizationalUnitsForApplicationResponse
func (client *Client) ListOrganizationalUnitsForApplicationWithOptions(request *ListOrganizationalUnitsForApplicationRequest, runtime *util.RuntimeOptions) (_result *ListOrganizationalUnitsForApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitIds)) {
		query["OrganizationalUnitIds"] = request.OrganizationalUnitIds
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOrganizationalUnitsForApplication"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOrganizationalUnitsForApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the organizations that are allowed to access an Employee Identity and Access Management (EIAM) application by page. The return result includes the IDs of the organizations. If you want to obtain the details of the organizations, call the GetOrganizationalUnit operation.
//
// @param request - ListOrganizationalUnitsForApplicationRequest
//
// @return ListOrganizationalUnitsForApplicationResponse
func (client *Client) ListOrganizationalUnitsForApplication(request *ListOrganizationalUnitsForApplicationRequest) (_result *ListOrganizationalUnitsForApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOrganizationalUnitsForApplicationResponse{}
	_body, _err := client.ListOrganizationalUnitsForApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the supported Alibaba Cloud regions.
//
// @param request - ListRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegionsResponse
func (client *Client) ListRegionsWithOptions(runtime *util.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListRegions"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the supported Alibaba Cloud regions.
//
// @return ListRegionsResponse
func (client *Client) ListRegions() (_result *ListRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询同步任务
//
// @param request - ListSynchronizationJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSynchronizationJobsResponse
func (client *Client) ListSynchronizationJobsWithOptions(request *ListSynchronizationJobsRequest, runtime *util.RuntimeOptions) (_result *ListSynchronizationJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TargetIds)) {
		query["TargetIds"] = request.TargetIds
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSynchronizationJobs"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSynchronizationJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询同步任务
//
// @param request - ListSynchronizationJobsRequest
//
// @return ListSynchronizationJobsResponse
func (client *Client) ListSynchronizationJobs(request *ListSynchronizationJobsRequest) (_result *ListSynchronizationJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSynchronizationJobsResponse{}
	_body, _err := client.ListSynchronizationJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of accounts in Identity as a Service (IDaaS) Employee IAM (EIAM) by page.
//
// @param request - ListUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisplayNameStartsWith)) {
		query["DisplayNameStartsWith"] = request.DisplayNameStartsWith
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitId)) {
		query["OrganizationalUnitId"] = request.OrganizationalUnitId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneRegion)) {
		query["PhoneRegion"] = request.PhoneRegion
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UserExternalId)) {
		query["UserExternalId"] = request.UserExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		query["UserIds"] = request.UserIds
	}

	if !tea.BoolValue(util.IsUnset(request.UserSourceId)) {
		query["UserSourceId"] = request.UserSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserSourceType)) {
		query["UserSourceType"] = request.UserSourceType
	}

	if !tea.BoolValue(util.IsUnset(request.UsernameStartsWith)) {
		query["UsernameStartsWith"] = request.UsernameStartsWith
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsers"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of accounts in Identity as a Service (IDaaS) Employee IAM (EIAM) by page.
//
// @param request - ListUsersRequest
//
// @return ListUsersResponse
func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the accounts that are allowed to access an Employee Identity and Access Management (EIAM) application. The return results include the IDs of the accounts. If you need to obtain the details of the accounts, call the GetUser operation.
//
// @param request - ListUsersForApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersForApplicationResponse
func (client *Client) ListUsersForApplicationWithOptions(request *ListUsersForApplicationRequest, runtime *util.RuntimeOptions) (_result *ListUsersForApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		query["UserIds"] = request.UserIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsersForApplication"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsersForApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the accounts that are allowed to access an Employee Identity and Access Management (EIAM) application. The return results include the IDs of the accounts. If you need to obtain the details of the accounts, call the GetUser operation.
//
// @param request - ListUsersForApplicationRequest
//
// @return ListUsersForApplicationResponse
func (client *Client) ListUsersForApplication(request *ListUsersForApplicationRequest) (_result *ListUsersForApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsersForApplicationResponse{}
	_body, _err := client.ListUsersForApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of accounts in an Employee Identity and Access Management (EIAM) group of Identity as a Service (IDaaS).
//
// @param request - ListUsersForGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersForGroupResponse
func (client *Client) ListUsersForGroupWithOptions(request *ListUsersForGroupRequest, runtime *util.RuntimeOptions) (_result *ListUsersForGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		query["UserIds"] = request.UserIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsersForGroup"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsersForGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of accounts in an Employee Identity and Access Management (EIAM) group of Identity as a Service (IDaaS).
//
// @param request - ListUsersForGroupRequest
//
// @return ListUsersForGroupResponse
func (client *Client) ListUsersForGroup(request *ListUsersForGroupRequest) (_result *ListUsersForGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsersForGroupResponse{}
	_body, _err := client.ListUsersForGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a client key of an Employee Identity and Access Management (EIAM) application. The returned key secret is masked. If you want to query the key secret that is not masked, call the ListApplicationClientSecrets operation.
//
// @param request - ObtainApplicationClientSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ObtainApplicationClientSecretResponse
func (client *Client) ObtainApplicationClientSecretWithOptions(request *ObtainApplicationClientSecretRequest, runtime *util.RuntimeOptions) (_result *ObtainApplicationClientSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretId)) {
		query["SecretId"] = request.SecretId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ObtainApplicationClientSecret"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ObtainApplicationClientSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a client key of an Employee Identity and Access Management (EIAM) application. The returned key secret is masked. If you want to query the key secret that is not masked, call the ListApplicationClientSecrets operation.
//
// @param request - ObtainApplicationClientSecretRequest
//
// @return ObtainApplicationClientSecretResponse
func (client *Client) ObtainApplicationClientSecret(request *ObtainApplicationClientSecretRequest) (_result *ObtainApplicationClientSecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ObtainApplicationClientSecretResponse{}
	_body, _err := client.ObtainApplicationClientSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看指定域名安全代理Token。
//
// @param request - ObtainDomainProxyTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ObtainDomainProxyTokenResponse
func (client *Client) ObtainDomainProxyTokenWithOptions(request *ObtainDomainProxyTokenRequest, runtime *util.RuntimeOptions) (_result *ObtainDomainProxyTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainId)) {
		query["DomainId"] = request.DomainId
	}

	if !tea.BoolValue(util.IsUnset(request.DomainProxyTokenId)) {
		query["DomainProxyTokenId"] = request.DomainProxyTokenId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ObtainDomainProxyToken"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ObtainDomainProxyTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看指定域名安全代理Token。
//
// @param request - ObtainDomainProxyTokenRequest
//
// @return ObtainDomainProxyTokenResponse
func (client *Client) ObtainDomainProxyToken(request *ObtainDomainProxyTokenRequest) (_result *ObtainDomainProxyTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ObtainDomainProxyTokenResponse{}
	_body, _err := client.ObtainDomainProxyTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes an Employee Identity and Access Management (EIAM) account from multiple EIAM organizations of Identity as a Service (IDaaS). You cannot remove an account from a primary organization.
//
// @param request - RemoveUserFromOrganizationalUnitsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserFromOrganizationalUnitsResponse
func (client *Client) RemoveUserFromOrganizationalUnitsWithOptions(request *RemoveUserFromOrganizationalUnitsRequest, runtime *util.RuntimeOptions) (_result *RemoveUserFromOrganizationalUnitsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitIds)) {
		query["OrganizationalUnitIds"] = request.OrganizationalUnitIds
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveUserFromOrganizationalUnits"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveUserFromOrganizationalUnitsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes an Employee Identity and Access Management (EIAM) account from multiple EIAM organizations of Identity as a Service (IDaaS). You cannot remove an account from a primary organization.
//
// @param request - RemoveUserFromOrganizationalUnitsRequest
//
// @return RemoveUserFromOrganizationalUnitsResponse
func (client *Client) RemoveUserFromOrganizationalUnits(request *RemoveUserFromOrganizationalUnitsRequest) (_result *RemoveUserFromOrganizationalUnitsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveUserFromOrganizationalUnitsResponse{}
	_body, _err := client.RemoveUserFromOrganizationalUnitsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes Employee Identity and Access Management (EIAM) accounts from an EIAM group of Identity as a Service (IDaaS).
//
// @param request - RemoveUsersFromGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUsersFromGroupResponse
func (client *Client) RemoveUsersFromGroupWithOptions(request *RemoveUsersFromGroupRequest, runtime *util.RuntimeOptions) (_result *RemoveUsersFromGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		query["UserIds"] = request.UserIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveUsersFromGroup"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveUsersFromGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes Employee Identity and Access Management (EIAM) accounts from an EIAM group of Identity as a Service (IDaaS).
//
// @param request - RemoveUsersFromGroupRequest
//
// @return RemoveUsersFromGroupResponse
func (client *Client) RemoveUsersFromGroup(request *RemoveUsersFromGroupRequest) (_result *RemoveUsersFromGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveUsersFromGroupResponse{}
	_body, _err := client.RemoveUsersFromGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes the permissions to access an application from multiple account groups at a time in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - RevokeApplicationFromGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeApplicationFromGroupsResponse
func (client *Client) RevokeApplicationFromGroupsWithOptions(request *RevokeApplicationFromGroupsRequest, runtime *util.RuntimeOptions) (_result *RevokeApplicationFromGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupIds)) {
		query["GroupIds"] = request.GroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RevokeApplicationFromGroups"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RevokeApplicationFromGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the permissions to access an application from multiple account groups at a time in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - RevokeApplicationFromGroupsRequest
//
// @return RevokeApplicationFromGroupsResponse
func (client *Client) RevokeApplicationFromGroups(request *RevokeApplicationFromGroupsRequest) (_result *RevokeApplicationFromGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevokeApplicationFromGroupsResponse{}
	_body, _err := client.RevokeApplicationFromGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes the permissions to access an application from multiple Employee Identity and Access Management (EIAM) organizations at a time.
//
// @param request - RevokeApplicationFromOrganizationalUnitsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeApplicationFromOrganizationalUnitsResponse
func (client *Client) RevokeApplicationFromOrganizationalUnitsWithOptions(request *RevokeApplicationFromOrganizationalUnitsRequest, runtime *util.RuntimeOptions) (_result *RevokeApplicationFromOrganizationalUnitsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitIds)) {
		query["OrganizationalUnitIds"] = request.OrganizationalUnitIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RevokeApplicationFromOrganizationalUnits"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RevokeApplicationFromOrganizationalUnitsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the permissions to access an application from multiple Employee Identity and Access Management (EIAM) organizations at a time.
//
// @param request - RevokeApplicationFromOrganizationalUnitsRequest
//
// @return RevokeApplicationFromOrganizationalUnitsResponse
func (client *Client) RevokeApplicationFromOrganizationalUnits(request *RevokeApplicationFromOrganizationalUnitsRequest) (_result *RevokeApplicationFromOrganizationalUnitsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevokeApplicationFromOrganizationalUnitsResponse{}
	_body, _err := client.RevokeApplicationFromOrganizationalUnitsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes the permissions to access an application from multiple Employee Identity and Access Management (EIAM) accounts at a time.
//
// @param request - RevokeApplicationFromUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeApplicationFromUsersResponse
func (client *Client) RevokeApplicationFromUsersWithOptions(request *RevokeApplicationFromUsersRequest, runtime *util.RuntimeOptions) (_result *RevokeApplicationFromUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		query["UserIds"] = request.UserIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RevokeApplicationFromUsers"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RevokeApplicationFromUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the permissions to access an application from multiple Employee Identity and Access Management (EIAM) accounts at a time.
//
// @param request - RevokeApplicationFromUsersRequest
//
// @return RevokeApplicationFromUsersResponse
func (client *Client) RevokeApplicationFromUsers(request *RevokeApplicationFromUsersRequest) (_result *RevokeApplicationFromUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevokeApplicationFromUsersResponse{}
	_body, _err := client.RevokeApplicationFromUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 运行同步任务
//
// @param request - RunSynchronizationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunSynchronizationJobResponse
func (client *Client) RunSynchronizationJobWithOptions(request *RunSynchronizationJobRequest, runtime *util.RuntimeOptions) (_result *RunSynchronizationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["TargetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RunSynchronizationJob"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunSynchronizationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 运行同步任务
//
// @param request - RunSynchronizationJobRequest
//
// @return RunSynchronizationJobResponse
func (client *Client) RunSynchronizationJob(request *RunSynchronizationJobRequest) (_result *RunSynchronizationJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunSynchronizationJobResponse{}
	_body, _err := client.RunSynchronizationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the permissions of the Developer API feature of an Employee Identity and Access Management (EIAM) application.
//
// @param request - SetApplicationGrantScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetApplicationGrantScopeResponse
func (client *Client) SetApplicationGrantScopeWithOptions(request *SetApplicationGrantScopeRequest, runtime *util.RuntimeOptions) (_result *SetApplicationGrantScopeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.GrantScopes)) {
		query["GrantScopes"] = request.GrantScopes
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetApplicationGrantScope"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetApplicationGrantScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the permissions of the Developer API feature of an Employee Identity and Access Management (EIAM) application.
//
// @param request - SetApplicationGrantScopeRequest
//
// @return SetApplicationGrantScopeResponse
func (client *Client) SetApplicationGrantScope(request *SetApplicationGrantScopeRequest) (_result *SetApplicationGrantScopeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetApplicationGrantScopeResponse{}
	_body, _err := client.SetApplicationGrantScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the account synchronization feature for an application in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - SetApplicationProvisioningConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetApplicationProvisioningConfigResponse
func (client *Client) SetApplicationProvisioningConfigWithOptions(request *SetApplicationProvisioningConfigRequest, runtime *util.RuntimeOptions) (_result *SetApplicationProvisioningConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackProvisioningConfig)) {
		query["CallbackProvisioningConfig"] = request.CallbackProvisioningConfig
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ProvisionPassword)) {
		query["ProvisionPassword"] = request.ProvisionPassword
	}

	if !tea.BoolValue(util.IsUnset(request.ProvisionProtocolType)) {
		query["ProvisionProtocolType"] = request.ProvisionProtocolType
	}

	if !tea.BoolValue(util.IsUnset(request.ScimProvisioningConfig)) {
		query["ScimProvisioningConfig"] = request.ScimProvisioningConfig
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetApplicationProvisioningConfig"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetApplicationProvisioningConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the account synchronization feature for an application in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM).
//
// @param request - SetApplicationProvisioningConfigRequest
//
// @return SetApplicationProvisioningConfigResponse
func (client *Client) SetApplicationProvisioningConfig(request *SetApplicationProvisioningConfigRequest) (_result *SetApplicationProvisioningConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetApplicationProvisioningConfigResponse{}
	_body, _err := client.SetApplicationProvisioningConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the account synchronization scope of applications in Identity as a Service (IDaaS) Employee IAM (EIAM). This scope is the same as the scope within which developers can call the DeveloperAPI to query and manage accounts.
//
// @param request - SetApplicationProvisioningScopeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetApplicationProvisioningScopeResponse
func (client *Client) SetApplicationProvisioningScopeWithOptions(request *SetApplicationProvisioningScopeRequest, runtime *util.RuntimeOptions) (_result *SetApplicationProvisioningScopeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupIds)) {
		query["GroupIds"] = request.GroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitIds)) {
		query["OrganizationalUnitIds"] = request.OrganizationalUnitIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetApplicationProvisioningScope"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetApplicationProvisioningScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the account synchronization scope of applications in Identity as a Service (IDaaS) Employee IAM (EIAM). This scope is the same as the scope within which developers can call the DeveloperAPI to query and manage accounts.
//
// @param request - SetApplicationProvisioningScopeRequest
//
// @return SetApplicationProvisioningScopeResponse
func (client *Client) SetApplicationProvisioningScope(request *SetApplicationProvisioningScopeRequest) (_result *SetApplicationProvisioningScopeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetApplicationProvisioningScopeResponse{}
	_body, _err := client.SetApplicationProvisioningScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Specifies the single sign-on (SSO) configuration attributes of an application in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// Description:
//
// In IDaaS EIAM, the application management feature supports multiple SSO protocols for applications, including SAML 2.0 and OIDC protocols. Each application supports only one protocol, and the protocol cannot be changed after the application is created. You can specify the SSO configuration attributes of an application based on the supported SSO protocol.
//
// @param request - SetApplicationSsoConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetApplicationSsoConfigResponse
func (client *Client) SetApplicationSsoConfigWithOptions(request *SetApplicationSsoConfigRequest, runtime *util.RuntimeOptions) (_result *SetApplicationSsoConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InitLoginType)) {
		query["InitLoginType"] = request.InitLoginType
	}

	if !tea.BoolValue(util.IsUnset(request.InitLoginUrl)) {
		query["InitLoginUrl"] = request.InitLoginUrl
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OidcSsoConfig)) {
		query["OidcSsoConfig"] = request.OidcSsoConfig
	}

	if !tea.BoolValue(util.IsUnset(request.SamlSsoConfig)) {
		query["SamlSsoConfig"] = request.SamlSsoConfig
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetApplicationSsoConfig"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetApplicationSsoConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Specifies the single sign-on (SSO) configuration attributes of an application in Identity as a Service (IDaaS) Employee IAM (EIAM).
//
// Description:
//
// In IDaaS EIAM, the application management feature supports multiple SSO protocols for applications, including SAML 2.0 and OIDC protocols. Each application supports only one protocol, and the protocol cannot be changed after the application is created. You can specify the SSO configuration attributes of an application based on the supported SSO protocol.
//
// @param request - SetApplicationSsoConfigRequest
//
// @return SetApplicationSsoConfigResponse
func (client *Client) SetApplicationSsoConfig(request *SetApplicationSsoConfigRequest) (_result *SetApplicationSsoConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetApplicationSsoConfigResponse{}
	_body, _err := client.SetApplicationSsoConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置默认域名。
//
// @param request - SetDefaultDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDefaultDomainResponse
func (client *Client) SetDefaultDomainWithOptions(request *SetDefaultDomainRequest, runtime *util.RuntimeOptions) (_result *SetDefaultDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainId)) {
		query["DomainId"] = request.DomainId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDefaultDomain"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDefaultDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置默认域名。
//
// @param request - SetDefaultDomainRequest
//
// @return SetDefaultDomainResponse
func (client *Client) SetDefaultDomain(request *SetDefaultDomainRequest) (_result *SetDefaultDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDefaultDomainResponse{}
	_body, _err := client.SetDefaultDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures a forgot password policy for an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - SetForgetPasswordConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetForgetPasswordConfigurationResponse
func (client *Client) SetForgetPasswordConfigurationWithOptions(request *SetForgetPasswordConfigurationRequest, runtime *util.RuntimeOptions) (_result *SetForgetPasswordConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthenticationChannels)) {
		query["AuthenticationChannels"] = request.AuthenticationChannels
	}

	if !tea.BoolValue(util.IsUnset(request.ForgetPasswordStatus)) {
		query["ForgetPasswordStatus"] = request.ForgetPasswordStatus
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetForgetPasswordConfiguration"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetForgetPasswordConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a forgot password policy for an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - SetForgetPasswordConfigurationRequest
//
// @return SetForgetPasswordConfigurationResponse
func (client *Client) SetForgetPasswordConfiguration(request *SetForgetPasswordConfigurationRequest) (_result *SetForgetPasswordConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetForgetPasswordConfigurationResponse{}
	_body, _err := client.SetForgetPasswordConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures a password complexity policy for an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - SetPasswordComplexityConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPasswordComplexityConfigurationResponse
func (client *Client) SetPasswordComplexityConfigurationWithOptions(request *SetPasswordComplexityConfigurationRequest, runtime *util.RuntimeOptions) (_result *SetPasswordComplexityConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordComplexityRules)) {
		query["PasswordComplexityRules"] = request.PasswordComplexityRules
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordMinLength)) {
		query["PasswordMinLength"] = request.PasswordMinLength
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetPasswordComplexityConfiguration"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetPasswordComplexityConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a password complexity policy for an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - SetPasswordComplexityConfigurationRequest
//
// @return SetPasswordComplexityConfigurationResponse
func (client *Client) SetPasswordComplexityConfiguration(request *SetPasswordComplexityConfigurationRequest) (_result *SetPasswordComplexityConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetPasswordComplexityConfigurationResponse{}
	_body, _err := client.SetPasswordComplexityConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures a password expiration policy for an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - SetPasswordExpirationConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPasswordExpirationConfigurationResponse
func (client *Client) SetPasswordExpirationConfigurationWithOptions(request *SetPasswordExpirationConfigurationRequest, runtime *util.RuntimeOptions) (_result *SetPasswordExpirationConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordExpirationAction)) {
		query["PasswordExpirationAction"] = request.PasswordExpirationAction
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordExpirationNotificationChannels)) {
		query["PasswordExpirationNotificationChannels"] = request.PasswordExpirationNotificationChannels
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordExpirationNotificationDuration)) {
		query["PasswordExpirationNotificationDuration"] = request.PasswordExpirationNotificationDuration
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordExpirationNotificationStatus)) {
		query["PasswordExpirationNotificationStatus"] = request.PasswordExpirationNotificationStatus
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordExpirationStatus)) {
		query["PasswordExpirationStatus"] = request.PasswordExpirationStatus
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordForcedUpdateDuration)) {
		query["PasswordForcedUpdateDuration"] = request.PasswordForcedUpdateDuration
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordValidMaxDay)) {
		query["PasswordValidMaxDay"] = request.PasswordValidMaxDay
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetPasswordExpirationConfiguration"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetPasswordExpirationConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a password expiration policy for an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - SetPasswordExpirationConfigurationRequest
//
// @return SetPasswordExpirationConfigurationResponse
func (client *Client) SetPasswordExpirationConfiguration(request *SetPasswordExpirationConfigurationRequest) (_result *SetPasswordExpirationConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetPasswordExpirationConfigurationResponse{}
	_body, _err := client.SetPasswordExpirationConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures a password history policy for an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - SetPasswordHistoryConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPasswordHistoryConfigurationResponse
func (client *Client) SetPasswordHistoryConfigurationWithOptions(request *SetPasswordHistoryConfigurationRequest, runtime *util.RuntimeOptions) (_result *SetPasswordHistoryConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordHistoryMaxRetention)) {
		query["PasswordHistoryMaxRetention"] = request.PasswordHistoryMaxRetention
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordHistoryStatus)) {
		query["PasswordHistoryStatus"] = request.PasswordHistoryStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetPasswordHistoryConfiguration"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetPasswordHistoryConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a password history policy for an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - SetPasswordHistoryConfigurationRequest
//
// @return SetPasswordHistoryConfigurationResponse
func (client *Client) SetPasswordHistoryConfiguration(request *SetPasswordHistoryConfigurationRequest) (_result *SetPasswordHistoryConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetPasswordHistoryConfigurationResponse{}
	_body, _err := client.SetPasswordHistoryConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the password initialization configurations for an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - SetPasswordInitializationConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPasswordInitializationConfigurationResponse
func (client *Client) SetPasswordInitializationConfigurationWithOptions(request *SetPasswordInitializationConfigurationRequest, runtime *util.RuntimeOptions) (_result *SetPasswordInitializationConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordForcedUpdateStatus)) {
		query["PasswordForcedUpdateStatus"] = request.PasswordForcedUpdateStatus
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordInitializationNotificationChannels)) {
		query["PasswordInitializationNotificationChannels"] = request.PasswordInitializationNotificationChannels
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordInitializationStatus)) {
		query["PasswordInitializationStatus"] = request.PasswordInitializationStatus
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordInitializationType)) {
		query["PasswordInitializationType"] = request.PasswordInitializationType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetPasswordInitializationConfiguration"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetPasswordInitializationConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the password initialization configurations for an Employee Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - SetPasswordInitializationConfigurationRequest
//
// @return SetPasswordInitializationConfigurationResponse
func (client *Client) SetPasswordInitializationConfiguration(request *SetPasswordInitializationConfigurationRequest) (_result *SetPasswordInitializationConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetPasswordInitializationConfigurationResponse{}
	_body, _err := client.SetPasswordInitializationConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the primary organizational unit to which an Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM) account belongs. This account will be removed from the previous primary organizational unit and added to the new primary organization.
//
// @param request - SetUserPrimaryOrganizationalUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetUserPrimaryOrganizationalUnitResponse
func (client *Client) SetUserPrimaryOrganizationalUnitWithOptions(request *SetUserPrimaryOrganizationalUnitRequest, runtime *util.RuntimeOptions) (_result *SetUserPrimaryOrganizationalUnitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitId)) {
		query["OrganizationalUnitId"] = request.OrganizationalUnitId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetUserPrimaryOrganizationalUnit"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetUserPrimaryOrganizationalUnitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the primary organizational unit to which an Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM) account belongs. This account will be removed from the previous primary organizational unit and added to the new primary organization.
//
// @param request - SetUserPrimaryOrganizationalUnitRequest
//
// @return SetUserPrimaryOrganizationalUnitResponse
func (client *Client) SetUserPrimaryOrganizationalUnit(request *SetUserPrimaryOrganizationalUnitRequest) (_result *SetUserPrimaryOrganizationalUnitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetUserPrimaryOrganizationalUnitResponse{}
	_body, _err := client.SetUserPrimaryOrganizationalUnitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unlocks an Employee Identity and Access Management (EIAM) account of Identity as a Service (IDaaS) that is locked.
//
// @param request - UnlockUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnlockUserResponse
func (client *Client) UnlockUserWithOptions(request *UnlockUserRequest, runtime *util.RuntimeOptions) (_result *UnlockUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnlockUser"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnlockUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unlocks an Employee Identity and Access Management (EIAM) account of Identity as a Service (IDaaS) that is locked.
//
// @param request - UnlockUserRequest
//
// @return UnlockUserResponse
func (client *Client) UnlockUser(request *UnlockUserRequest) (_result *UnlockUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnlockUserResponse{}
	_body, _err := client.UnlockUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the authorization type of an Employee Identity and Access Management (EIAM) application.
//
// @param request - UpdateApplicationAuthorizationTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationAuthorizationTypeResponse
func (client *Client) UpdateApplicationAuthorizationTypeWithOptions(request *UpdateApplicationAuthorizationTypeRequest, runtime *util.RuntimeOptions) (_result *UpdateApplicationAuthorizationTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.AuthorizationType)) {
		query["AuthorizationType"] = request.AuthorizationType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateApplicationAuthorizationType"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateApplicationAuthorizationTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the authorization type of an Employee Identity and Access Management (EIAM) application.
//
// @param request - UpdateApplicationAuthorizationTypeRequest
//
// @return UpdateApplicationAuthorizationTypeResponse
func (client *Client) UpdateApplicationAuthorizationType(request *UpdateApplicationAuthorizationTypeRequest) (_result *UpdateApplicationAuthorizationTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateApplicationAuthorizationTypeResponse{}
	_body, _err := client.UpdateApplicationAuthorizationTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description of an Employee Identity and Access Management (EIAM) application.
//
// @param request - UpdateApplicationDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationDescriptionResponse
func (client *Client) UpdateApplicationDescriptionWithOptions(request *UpdateApplicationDescriptionRequest, runtime *util.RuntimeOptions) (_result *UpdateApplicationDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateApplicationDescription"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateApplicationDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of an Employee Identity and Access Management (EIAM) application.
//
// @param request - UpdateApplicationDescriptionRequest
//
// @return UpdateApplicationDescriptionResponse
func (client *Client) UpdateApplicationDescription(request *UpdateApplicationDescriptionRequest) (_result *UpdateApplicationDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateApplicationDescriptionResponse{}
	_body, _err := client.UpdateApplicationDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about an account group in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM). If the information is empty, the information is not updated by default.
//
// @param request - UpdateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupResponse
func (client *Client) UpdateGroupWithOptions(request *UpdateGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupExternalId)) {
		query["GroupExternalId"] = request.GroupExternalId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGroup"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about an account group in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM). If the information is empty, the information is not updated by default.
//
// @param request - UpdateGroupRequest
//
// @return UpdateGroupResponse
func (client *Client) UpdateGroup(request *UpdateGroupRequest) (_result *UpdateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGroupResponse{}
	_body, _err := client.UpdateGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the description of an Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM) account group.
//
// @param request - UpdateGroupDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupDescriptionResponse
func (client *Client) UpdateGroupDescriptionWithOptions(request *UpdateGroupDescriptionRequest, runtime *util.RuntimeOptions) (_result *UpdateGroupDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGroupDescription"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGroupDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the description of an Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM) account group.
//
// @param request - UpdateGroupDescriptionRequest
//
// @return UpdateGroupDescriptionResponse
func (client *Client) UpdateGroupDescription(request *UpdateGroupDescriptionRequest) (_result *UpdateGroupDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGroupDescriptionResponse{}
	_body, _err := client.UpdateGroupDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description of an Enterprise Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - UpdateInstanceDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceDescriptionResponse
func (client *Client) UpdateInstanceDescriptionWithOptions(request *UpdateInstanceDescriptionRequest, runtime *util.RuntimeOptions) (_result *UpdateInstanceDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceDescription"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of an Enterprise Identity and Access Management (EIAM) instance of Identity as a Service (IDaaS).
//
// @param request - UpdateInstanceDescriptionRequest
//
// @return UpdateInstanceDescriptionResponse
func (client *Client) UpdateInstanceDescription(request *UpdateInstanceDescriptionRequest) (_result *UpdateInstanceDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateInstanceDescriptionResponse{}
	_body, _err := client.UpdateInstanceDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新一个专属网络端点的名称。
//
// @param request - UpdateNetworkAccessEndpointNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNetworkAccessEndpointNameResponse
func (client *Client) UpdateNetworkAccessEndpointNameWithOptions(request *UpdateNetworkAccessEndpointNameRequest, runtime *util.RuntimeOptions) (_result *UpdateNetworkAccessEndpointNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkAccessEndpointId)) {
		query["NetworkAccessEndpointId"] = request.NetworkAccessEndpointId
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkAccessEndpointName)) {
		query["NetworkAccessEndpointName"] = request.NetworkAccessEndpointName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateNetworkAccessEndpointName"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateNetworkAccessEndpointNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新一个专属网络端点的名称。
//
// @param request - UpdateNetworkAccessEndpointNameRequest
//
// @return UpdateNetworkAccessEndpointNameResponse
func (client *Client) UpdateNetworkAccessEndpointName(request *UpdateNetworkAccessEndpointNameRequest) (_result *UpdateNetworkAccessEndpointNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateNetworkAccessEndpointNameResponse{}
	_body, _err := client.UpdateNetworkAccessEndpointNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the basic information about an Employee Identity and Access Management (EIAM) organization. The basic information about the organization is not updated by default if no parameter is specified.
//
// @param request - UpdateOrganizationalUnitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOrganizationalUnitResponse
func (client *Client) UpdateOrganizationalUnitWithOptions(request *UpdateOrganizationalUnitRequest, runtime *util.RuntimeOptions) (_result *UpdateOrganizationalUnitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitId)) {
		query["OrganizationalUnitId"] = request.OrganizationalUnitId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitName)) {
		query["OrganizationalUnitName"] = request.OrganizationalUnitName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateOrganizationalUnit"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateOrganizationalUnitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic information about an Employee Identity and Access Management (EIAM) organization. The basic information about the organization is not updated by default if no parameter is specified.
//
// @param request - UpdateOrganizationalUnitRequest
//
// @return UpdateOrganizationalUnitResponse
func (client *Client) UpdateOrganizationalUnit(request *UpdateOrganizationalUnitRequest) (_result *UpdateOrganizationalUnitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateOrganizationalUnitResponse{}
	_body, _err := client.UpdateOrganizationalUnitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description of an Employee Identity and Access Management (EIAM) organization.
//
// @param request - UpdateOrganizationalUnitDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOrganizationalUnitDescriptionResponse
func (client *Client) UpdateOrganizationalUnitDescriptionWithOptions(request *UpdateOrganizationalUnitDescriptionRequest, runtime *util.RuntimeOptions) (_result *UpdateOrganizationalUnitDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitId)) {
		query["OrganizationalUnitId"] = request.OrganizationalUnitId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateOrganizationalUnitDescription"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateOrganizationalUnitDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of an Employee Identity and Access Management (EIAM) organization.
//
// @param request - UpdateOrganizationalUnitDescriptionRequest
//
// @return UpdateOrganizationalUnitDescriptionResponse
func (client *Client) UpdateOrganizationalUnitDescription(request *UpdateOrganizationalUnitDescriptionRequest) (_result *UpdateOrganizationalUnitDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateOrganizationalUnitDescriptionResponse{}
	_body, _err := client.UpdateOrganizationalUnitDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the parent organization ID of an organization in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM). In this case, the organization is moved from a parent node to a new node.
//
// @param request - UpdateOrganizationalUnitParentIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOrganizationalUnitParentIdResponse
func (client *Client) UpdateOrganizationalUnitParentIdWithOptions(request *UpdateOrganizationalUnitParentIdRequest, runtime *util.RuntimeOptions) (_result *UpdateOrganizationalUnitParentIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrganizationalUnitId)) {
		query["OrganizationalUnitId"] = request.OrganizationalUnitId
	}

	if !tea.BoolValue(util.IsUnset(request.ParentId)) {
		query["ParentId"] = request.ParentId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateOrganizationalUnitParentId"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateOrganizationalUnitParentIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the parent organization ID of an organization in Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM). In this case, the organization is moved from a parent node to a new node.
//
// @param request - UpdateOrganizationalUnitParentIdRequest
//
// @return UpdateOrganizationalUnitParentIdResponse
func (client *Client) UpdateOrganizationalUnitParentId(request *UpdateOrganizationalUnitParentIdRequest) (_result *UpdateOrganizationalUnitParentIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateOrganizationalUnitParentIdResponse{}
	_body, _err := client.UpdateOrganizationalUnitParentIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the basic information about an Employee Identity and Access Management (EIAM) account of Identity as a Service (IDaaS).
//
// @param request - UpdateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserResponse
func (client *Client) UpdateUserWithOptions(request *UpdateUserRequest, runtime *util.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomFields)) {
		query["CustomFields"] = request.CustomFields
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.EmailVerified)) {
		query["EmailVerified"] = request.EmailVerified
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumberVerified)) {
		query["PhoneNumberVerified"] = request.PhoneNumberVerified
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneRegion)) {
		query["PhoneRegion"] = request.PhoneRegion
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUser"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic information about an Employee Identity and Access Management (EIAM) account of Identity as a Service (IDaaS).
//
// @param request - UpdateUserRequest
//
// @return UpdateUserResponse
func (client *Client) UpdateUser(request *UpdateUserRequest) (_result *UpdateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserResponse{}
	_body, _err := client.UpdateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description of an Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM) account.
//
// @param request - UpdateUserDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserDescriptionResponse
func (client *Client) UpdateUserDescriptionWithOptions(request *UpdateUserDescriptionRequest, runtime *util.RuntimeOptions) (_result *UpdateUserDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUserDescription"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of an Identity as a Service (IDaaS) Employee Identity and Access Management (EIAM) account.
//
// @param request - UpdateUserDescriptionRequest
//
// @return UpdateUserDescriptionResponse
func (client *Client) UpdateUserDescription(request *UpdateUserDescriptionRequest) (_result *UpdateUserDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserDescriptionResponse{}
	_body, _err := client.UpdateUserDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the password information of an Employee Identity and Access Management (EIAM) account of Identity as a Service (IDaaS). The password must meet the requirements of the password policies that are configured in the IDaaS console.
//
// @param request - UpdateUserPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserPasswordResponse
func (client *Client) UpdateUserPasswordWithOptions(request *UpdateUserPasswordRequest, runtime *util.RuntimeOptions) (_result *UpdateUserPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordForcedUpdateStatus)) {
		query["PasswordForcedUpdateStatus"] = request.PasswordForcedUpdateStatus
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserNotificationChannels)) {
		query["UserNotificationChannels"] = request.UserNotificationChannels
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUserPassword"),
		Version:     tea.String("2021-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the password information of an Employee Identity and Access Management (EIAM) account of Identity as a Service (IDaaS). The password must meet the requirements of the password policies that are configured in the IDaaS console.
//
// @param request - UpdateUserPasswordRequest
//
// @return UpdateUserPasswordResponse
func (client *Client) UpdateUserPassword(request *UpdateUserPasswordRequest) (_result *UpdateUserPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserPasswordResponse{}
	_body, _err := client.UpdateUserPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
