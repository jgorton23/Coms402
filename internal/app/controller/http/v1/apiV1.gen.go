// Package http provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package http

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
)

// Certification defines model for Certification.
type Certification struct {
	// CompanyUUID uuid of the company
	CompanyUUID string `json:"companyUUID"`

	// ImageUUID uuid of the image
	ImageUUID *string `json:"imageUUID,omitempty"`

	// ItemBatchUUID uuid of the item batch
	ItemBatchUUID string `json:"itemBatchUUID"`

	// PrimaryAttribute primaryAttribute
	PrimaryAttribute string `json:"primaryAttribute"`

	// TemplateUUID uuid of the template
	TemplateUUID *string `json:"templateUUID,omitempty"`

	// Uuid uuid of the item batch
	Uuid string `json:"uuid"`
}

// Company defines model for Company.
type Company struct {
	// Name Name of the company
	Name string `json:"name"`

	// Uuid uuid of the company
	Uuid string `json:"uuid"`
}

// Error defines model for Error.
type Error struct {
	// Code Error code
	Code int32 `json:"code"`

	// Message Error message
	Message string `json:"message"`
}

// ItemBatch defines model for ItemBatch.
type ItemBatch struct {
	// CompanyUuid uuid of the company
	CompanyUuid string `json:"company_uuid"`

	// Description description of item batch
	Description *string `json:"description,omitempty"`

	// ItemNumber item number of the batch
	ItemNumber string `json:"itemNumber"`

	// Uuid uuid of the item batch
	Uuid string `json:"uuid"`
}

// NewCertification defines model for NewCertification.
type NewCertification struct {
	// CompanyUUID uuid of the company
	CompanyUUID string `json:"companyUUID"`

	// ImageUUID uuid of the image
	ImageUUID *string `json:"imageUUID,omitempty"`

	// ItemBatchUUID uuid of the item batch
	ItemBatchUUID string `json:"itemBatchUUID"`

	// PrimaryAttribute primaryAttribute
	PrimaryAttribute string `json:"primaryAttribute"`
}

// NewCompany defines model for NewCompany.
type NewCompany struct {
	// Name Name of the company
	Name string `json:"name"`
}

// NewItemBatch defines model for NewItemBatch.
type NewItemBatch struct {
	// CompanyUuid uuid of the company
	CompanyUuid string `json:"company_uuid"`

	// Description description of item batch
	Description *string `json:"description,omitempty"`

	// ItemNumber item number of the batch
	ItemNumber string `json:"itemNumber"`
}

// NewUserRole defines model for NewUserRole.
type NewUserRole struct {
	Approved    *bool   `json:"approved,omitempty"`
	CompanyUUID string  `json:"companyUUID"`
	RoleType    string  `json:"roleType"`
	UserUUID    *string `json:"userUUID,omitempty"`
}

// User defines model for User.
type User struct {
	// Created time user was created at
	Created string `json:"created"`

	// Email email of the user
	Email string `json:"email"`

	// Uuid uuid of the user
	Uuid string `json:"uuid"`
}

// UserRole defines model for UserRole.
type UserRole struct {
	Approved    *bool   `json:"approved,omitempty"`
	CompanyUUID string  `json:"companyUUID"`
	RoleType    string  `json:"roleType"`
	UserUUID    *string `json:"userUUID,omitempty"`
	Uuid        string  `json:"uuid"`
}

// DefaultErrorResponse defines model for DefaultErrorResponse.
type DefaultErrorResponse = Error

// DefaultUnauthenticatedErrorResponse defines model for DefaultUnauthenticatedErrorResponse.
type DefaultUnauthenticatedErrorResponse = Error

// DefaultUnauthorizedErrorResponse defines model for DefaultUnauthorizedErrorResponse.
type DefaultUnauthorizedErrorResponse = Error

// GetCertificationByParams defines parameters for GetCertificationBy.
type GetCertificationByParams struct {
	// CertificationUUID UUID of the item batch
	CertificationUUID *string `form:"certificationUUID,omitempty" json:"certificationUUID,omitempty"`

	// CompanyUUID uuid of company for which to get
	CompanyUUID *string `form:"companyUUID,omitempty" json:"companyUUID,omitempty"`
}

// GetItemBatchByParams defines parameters for GetItemBatchBy.
type GetItemBatchByParams struct {
	// ItemBatchUUID UUID of the item batch
	ItemBatchUUID *string `form:"itemBatchUUID,omitempty" json:"itemBatchUUID,omitempty"`

	// CompanyUUID uuid of company for which to get
	CompanyUUID *string `form:"companyUUID,omitempty" json:"companyUUID,omitempty"`
}

// AddSubItemsJSONBody defines parameters for AddSubItems.
type AddSubItemsJSONBody struct {
	Children []ItemBatch `json:"children"`
	Parent   ItemBatch   `json:"parent"`
}

// GetRolesByParams defines parameters for GetRolesBy.
type GetRolesByParams struct {
	// UserUUID UUID of user for which to get roles
	UserUUID *string `form:"userUUID,omitempty" json:"userUUID,omitempty"`

	// CompanyUUID uuid of company for which to get roles
	CompanyUUID *string `form:"companyUUID,omitempty" json:"companyUUID,omitempty"`
}

// ApproveRoleParams defines parameters for ApproveRole.
type ApproveRoleParams struct {
	// CompanyUUID uuid of the company for which the role belongs
	CompanyUUID string `form:"companyUUID" json:"companyUUID"`

	// UserUUID uuid of the user for which the role belongs
	UserUUID string `form:"userUUID" json:"userUUID"`
}

// GetUserByParams defines parameters for GetUserBy.
type GetUserByParams struct {
	// UserEmail email of user to return
	UserEmail *string `form:"userEmail,omitempty" json:"userEmail,omitempty"`

	// UserUUID uuid of user to return
	UserUUID *string `form:"userUUID,omitempty" json:"userUUID,omitempty"`
}

// AddCertificationJSONRequestBody defines body for AddCertification for application/json ContentType.
type AddCertificationJSONRequestBody = NewCertification

// UpdateCertificationJSONRequestBody defines body for UpdateCertification for application/json ContentType.
type UpdateCertificationJSONRequestBody = Certification

// AddCompanyJSONRequestBody defines body for AddCompany for application/json ContentType.
type AddCompanyJSONRequestBody = NewCompany

// UpdateCompanyJSONRequestBody defines body for UpdateCompany for application/json ContentType.
type UpdateCompanyJSONRequestBody = Company

// AddItemBatchJSONRequestBody defines body for AddItemBatch for application/json ContentType.
type AddItemBatchJSONRequestBody = NewItemBatch

// UpdateItemBatchJSONRequestBody defines body for UpdateItemBatch for application/json ContentType.
type UpdateItemBatchJSONRequestBody = ItemBatch

// AddSubItemsJSONRequestBody defines body for AddSubItems for application/json ContentType.
type AddSubItemsJSONRequestBody AddSubItemsJSONBody

// AddUserRoleJSONRequestBody defines body for AddUserRole for application/json ContentType.
type AddUserRoleJSONRequestBody = NewUserRole

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Find certification by *
	// (GET /certification)
	GetCertificationBy(w http.ResponseWriter, r *http.Request, params GetCertificationByParams)
	// Create a new item batch
	// (POST /certification)
	AddCertification(w http.ResponseWriter, r *http.Request)
	// Update an existing item batch
	// (PUT /certification)
	UpdateCertification(w http.ResponseWriter, r *http.Request)
	// Create a new Company
	// (POST /company)
	AddCompany(w http.ResponseWriter, r *http.Request)
	// Update an existing company
	// (PUT /company)
	UpdateCompany(w http.ResponseWriter, r *http.Request)
	// Find company by UUID
	// (GET /company/{companyUUID})
	GetCompanyByUUID(w http.ResponseWriter, r *http.Request, companyUUID string)
	// Find itembatch by *
	// (GET /itembatch)
	GetItemBatchBy(w http.ResponseWriter, r *http.Request, params GetItemBatchByParams)
	// Create a new item batch
	// (POST /itembatch)
	AddItemBatch(w http.ResponseWriter, r *http.Request)
	// Update an existing item batch
	// (PUT /itembatch)
	UpdateItemBatch(w http.ResponseWriter, r *http.Request)
	// Create subcomponent mapping
	// (POST /itembatch/children)
	AddSubItems(w http.ResponseWriter, r *http.Request)
	// Find roles by *
	// (GET /role)
	GetRolesBy(w http.ResponseWriter, r *http.Request, params GetRolesByParams)
	// Create a new User Role
	// (POST /role)
	AddUserRole(w http.ResponseWriter, r *http.Request)
	// Approve the given role
	// (PUT /role/approve)
	ApproveRole(w http.ResponseWriter, r *http.Request, params ApproveRoleParams)
	// Find user by *
	// (GET /user)
	GetUserBy(w http.ResponseWriter, r *http.Request, params GetUserByParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetCertificationBy operation middleware
func (siw *ServerInterfaceWrapper) GetCertificationBy(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetCertificationByParams

	// ------------- Optional query parameter "certificationUUID" -------------

	err = runtime.BindQueryParameter("form", true, false, "certificationUUID", r.URL.Query(), &params.CertificationUUID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "certificationUUID", Err: err})
		return
	}

	// ------------- Optional query parameter "companyUUID" -------------

	err = runtime.BindQueryParameter("form", true, false, "companyUUID", r.URL.Query(), &params.CompanyUUID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "companyUUID", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetCertificationBy(w, r, params)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AddCertification operation middleware
func (siw *ServerInterfaceWrapper) AddCertification(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddCertification(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateCertification operation middleware
func (siw *ServerInterfaceWrapper) UpdateCertification(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateCertification(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AddCompany operation middleware
func (siw *ServerInterfaceWrapper) AddCompany(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddCompany(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateCompany operation middleware
func (siw *ServerInterfaceWrapper) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateCompany(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetCompanyByUUID operation middleware
func (siw *ServerInterfaceWrapper) GetCompanyByUUID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "companyUUID" -------------
	var companyUUID string

	err = runtime.BindStyledParameterWithLocation("simple", false, "companyUUID", runtime.ParamLocationPath, chi.URLParam(r, "companyUUID"), &companyUUID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "companyUUID", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetCompanyByUUID(w, r, companyUUID)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetItemBatchBy operation middleware
func (siw *ServerInterfaceWrapper) GetItemBatchBy(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetItemBatchByParams

	// ------------- Optional query parameter "itemBatchUUID" -------------

	err = runtime.BindQueryParameter("form", true, false, "itemBatchUUID", r.URL.Query(), &params.ItemBatchUUID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "itemBatchUUID", Err: err})
		return
	}

	// ------------- Optional query parameter "companyUUID" -------------

	err = runtime.BindQueryParameter("form", true, false, "companyUUID", r.URL.Query(), &params.CompanyUUID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "companyUUID", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetItemBatchBy(w, r, params)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AddItemBatch operation middleware
func (siw *ServerInterfaceWrapper) AddItemBatch(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddItemBatch(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// UpdateItemBatch operation middleware
func (siw *ServerInterfaceWrapper) UpdateItemBatch(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateItemBatch(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AddSubItems operation middleware
func (siw *ServerInterfaceWrapper) AddSubItems(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddSubItems(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetRolesBy operation middleware
func (siw *ServerInterfaceWrapper) GetRolesBy(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetRolesByParams

	// ------------- Optional query parameter "userUUID" -------------

	err = runtime.BindQueryParameter("form", true, false, "userUUID", r.URL.Query(), &params.UserUUID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userUUID", Err: err})
		return
	}

	// ------------- Optional query parameter "companyUUID" -------------

	err = runtime.BindQueryParameter("form", true, false, "companyUUID", r.URL.Query(), &params.CompanyUUID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "companyUUID", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetRolesBy(w, r, params)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AddUserRole operation middleware
func (siw *ServerInterfaceWrapper) AddUserRole(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddUserRole(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// ApproveRole operation middleware
func (siw *ServerInterfaceWrapper) ApproveRole(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ApproveRoleParams

	// ------------- Required query parameter "companyUUID" -------------

	if paramValue := r.URL.Query().Get("companyUUID"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "companyUUID"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "companyUUID", r.URL.Query(), &params.CompanyUUID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "companyUUID", Err: err})
		return
	}

	// ------------- Required query parameter "userUUID" -------------

	if paramValue := r.URL.Query().Get("userUUID"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "userUUID"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "userUUID", r.URL.Query(), &params.UserUUID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userUUID", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ApproveRole(w, r, params)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetUserBy operation middleware
func (siw *ServerInterfaceWrapper) GetUserBy(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetUserByParams

	// ------------- Optional query parameter "userEmail" -------------

	err = runtime.BindQueryParameter("form", true, false, "userEmail", r.URL.Query(), &params.UserEmail)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userEmail", Err: err})
		return
	}

	// ------------- Optional query parameter "userUUID" -------------

	err = runtime.BindQueryParameter("form", true, false, "userUUID", r.URL.Query(), &params.UserUUID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userUUID", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUserBy(w, r, params)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshallingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshallingParamError) Error() string {
	return fmt.Sprintf("Error unmarshalling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshallingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/certification", wrapper.GetCertificationBy)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/certification", wrapper.AddCertification)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/certification", wrapper.UpdateCertification)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/company", wrapper.AddCompany)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/company", wrapper.UpdateCompany)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/company/{companyUUID}", wrapper.GetCompanyByUUID)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/itembatch", wrapper.GetItemBatchBy)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/itembatch", wrapper.AddItemBatch)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/itembatch", wrapper.UpdateItemBatch)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/itembatch/children", wrapper.AddSubItems)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/role", wrapper.GetRolesBy)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/role", wrapper.AddUserRole)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/role/approve", wrapper.ApproveRole)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/user", wrapper.GetUserBy)
	})

	return r
}
