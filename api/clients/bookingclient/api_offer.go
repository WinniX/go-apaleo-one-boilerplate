/*
 * apaleo Booking API
 *
 * Resources and methods to manage guest journeys.
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bookingclient

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// OfferApiService OfferApi service
type OfferApiService service

type ApiBookingOfferIndexGetRequest struct {
	ctx _context.Context
	ApiService *OfferApiService
	ratePlanId *string
	from *string
	to *string
	channelCode *string
	pageNumber *int32
	pageSize *int32
}

func (r ApiBookingOfferIndexGetRequest) RatePlanId(ratePlanId string) ApiBookingOfferIndexGetRequest {
	r.ratePlanId = &ratePlanId
	return r
}
func (r ApiBookingOfferIndexGetRequest) From(from string) ApiBookingOfferIndexGetRequest {
	r.from = &from
	return r
}
func (r ApiBookingOfferIndexGetRequest) To(to string) ApiBookingOfferIndexGetRequest {
	r.to = &to
	return r
}
func (r ApiBookingOfferIndexGetRequest) ChannelCode(channelCode string) ApiBookingOfferIndexGetRequest {
	r.channelCode = &channelCode
	return r
}
func (r ApiBookingOfferIndexGetRequest) PageNumber(pageNumber int32) ApiBookingOfferIndexGetRequest {
	r.pageNumber = &pageNumber
	return r
}
func (r ApiBookingOfferIndexGetRequest) PageSize(pageSize int32) ApiBookingOfferIndexGetRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiBookingOfferIndexGetRequest) Execute() (TimeSliceListModel, *_nethttp.Response, error) {
	return r.ApiService.BookingOfferIndexGetExecute(r)
}

/*
 * BookingOfferIndexGet Returns offers with rates and availabilities for the specified range.
 * Calculates and returns offers per time slice for a specific rate plan, arrival and departure date.<br>You must have at least one of these scopes: 'offer-index.read, offers.read'.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiBookingOfferIndexGetRequest
 */
func (a *OfferApiService) BookingOfferIndexGet(ctx _context.Context) ApiBookingOfferIndexGetRequest {
	return ApiBookingOfferIndexGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return TimeSliceListModel
 */
func (a *OfferApiService) BookingOfferIndexGetExecute(r ApiBookingOfferIndexGetRequest) (TimeSliceListModel, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TimeSliceListModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OfferApiService.BookingOfferIndexGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/booking/v1/offer-index"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.ratePlanId == nil {
		return localVarReturnValue, nil, reportError("ratePlanId is required and must be specified")
	}
	if r.from == nil {
		return localVarReturnValue, nil, reportError("from is required and must be specified")
	}
	if r.to == nil {
		return localVarReturnValue, nil, reportError("to is required and must be specified")
	}
	if r.channelCode == nil {
		return localVarReturnValue, nil, reportError("channelCode is required and must be specified")
	}

	localVarQueryParams.Add("ratePlanId", parameterToString(*r.ratePlanId, ""))
	localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	localVarQueryParams.Add("channelCode", parameterToString(*r.channelCode, ""))
	if r.pageNumber != nil {
		localVarQueryParams.Add("pageNumber", parameterToString(*r.pageNumber, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v MessageItemCollection
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiBookingOffersGetRequest struct {
	ctx _context.Context
	ApiService *OfferApiService
	propertyId *string
	arrival *string
	departure *string
	adults *int32
	timeSliceTemplate *string
	timeSliceDefinitionIds *[]string
	unitGroupIds *[]string
	unitGroupTypes *[]string
	channelCode *string
	promoCode *string
	corporateCode *string
	childrenAges *[]int32
	includeUnavailable *bool
}

func (r ApiBookingOffersGetRequest) PropertyId(propertyId string) ApiBookingOffersGetRequest {
	r.propertyId = &propertyId
	return r
}
func (r ApiBookingOffersGetRequest) Arrival(arrival string) ApiBookingOffersGetRequest {
	r.arrival = &arrival
	return r
}
func (r ApiBookingOffersGetRequest) Departure(departure string) ApiBookingOffersGetRequest {
	r.departure = &departure
	return r
}
func (r ApiBookingOffersGetRequest) Adults(adults int32) ApiBookingOffersGetRequest {
	r.adults = &adults
	return r
}
func (r ApiBookingOffersGetRequest) TimeSliceTemplate(timeSliceTemplate string) ApiBookingOffersGetRequest {
	r.timeSliceTemplate = &timeSliceTemplate
	return r
}
func (r ApiBookingOffersGetRequest) TimeSliceDefinitionIds(timeSliceDefinitionIds []string) ApiBookingOffersGetRequest {
	r.timeSliceDefinitionIds = &timeSliceDefinitionIds
	return r
}
func (r ApiBookingOffersGetRequest) UnitGroupIds(unitGroupIds []string) ApiBookingOffersGetRequest {
	r.unitGroupIds = &unitGroupIds
	return r
}
func (r ApiBookingOffersGetRequest) UnitGroupTypes(unitGroupTypes []string) ApiBookingOffersGetRequest {
	r.unitGroupTypes = &unitGroupTypes
	return r
}
func (r ApiBookingOffersGetRequest) ChannelCode(channelCode string) ApiBookingOffersGetRequest {
	r.channelCode = &channelCode
	return r
}
func (r ApiBookingOffersGetRequest) PromoCode(promoCode string) ApiBookingOffersGetRequest {
	r.promoCode = &promoCode
	return r
}
func (r ApiBookingOffersGetRequest) CorporateCode(corporateCode string) ApiBookingOffersGetRequest {
	r.corporateCode = &corporateCode
	return r
}
func (r ApiBookingOffersGetRequest) ChildrenAges(childrenAges []int32) ApiBookingOffersGetRequest {
	r.childrenAges = &childrenAges
	return r
}
func (r ApiBookingOffersGetRequest) IncludeUnavailable(includeUnavailable bool) ApiBookingOffersGetRequest {
	r.includeUnavailable = &includeUnavailable
	return r
}

func (r ApiBookingOffersGetRequest) Execute() (StayOffersModel, *_nethttp.Response, error) {
	return r.ApiService.BookingOffersGetExecute(r)
}

/*
 * BookingOffersGet Returns offers for one specific stay.
 * Calculates and returns available offers for a specific property, arrival and departure date.<br>You must have at least one of these scopes: 'offers.read, reservations.manage'.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiBookingOffersGetRequest
 */
func (a *OfferApiService) BookingOffersGet(ctx _context.Context) ApiBookingOffersGetRequest {
	return ApiBookingOffersGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return StayOffersModel
 */
func (a *OfferApiService) BookingOffersGetExecute(r ApiBookingOffersGetRequest) (StayOffersModel, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  StayOffersModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OfferApiService.BookingOffersGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/booking/v1/offers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.propertyId == nil {
		return localVarReturnValue, nil, reportError("propertyId is required and must be specified")
	}
	if r.arrival == nil {
		return localVarReturnValue, nil, reportError("arrival is required and must be specified")
	}
	if r.departure == nil {
		return localVarReturnValue, nil, reportError("departure is required and must be specified")
	}
	if r.adults == nil {
		return localVarReturnValue, nil, reportError("adults is required and must be specified")
	}

	localVarQueryParams.Add("propertyId", parameterToString(*r.propertyId, ""))
	localVarQueryParams.Add("arrival", parameterToString(*r.arrival, ""))
	localVarQueryParams.Add("departure", parameterToString(*r.departure, ""))
	if r.timeSliceTemplate != nil {
		localVarQueryParams.Add("timeSliceTemplate", parameterToString(*r.timeSliceTemplate, ""))
	}
	if r.timeSliceDefinitionIds != nil {
		localVarQueryParams.Add("timeSliceDefinitionIds", parameterToString(*r.timeSliceDefinitionIds, "csv"))
	}
	if r.unitGroupIds != nil {
		localVarQueryParams.Add("unitGroupIds", parameterToString(*r.unitGroupIds, "csv"))
	}
	if r.unitGroupTypes != nil {
		localVarQueryParams.Add("unitGroupTypes", parameterToString(*r.unitGroupTypes, "csv"))
	}
	if r.channelCode != nil {
		localVarQueryParams.Add("channelCode", parameterToString(*r.channelCode, ""))
	}
	if r.promoCode != nil {
		localVarQueryParams.Add("promoCode", parameterToString(*r.promoCode, ""))
	}
	if r.corporateCode != nil {
		localVarQueryParams.Add("corporateCode", parameterToString(*r.corporateCode, ""))
	}
	localVarQueryParams.Add("adults", parameterToString(*r.adults, ""))
	if r.childrenAges != nil {
		localVarQueryParams.Add("childrenAges", parameterToString(*r.childrenAges, "csv"))
	}
	if r.includeUnavailable != nil {
		localVarQueryParams.Add("includeUnavailable", parameterToString(*r.includeUnavailable, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v MessageItemCollection
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiBookingRatePlanOffersGetRequest struct {
	ctx _context.Context
	ApiService *OfferApiService
	ratePlanId *string
	arrival *string
	departure *string
	adults *int32
	channelCode *string
	childrenAges *[]int32
	includeUnavailable *bool
}

func (r ApiBookingRatePlanOffersGetRequest) RatePlanId(ratePlanId string) ApiBookingRatePlanOffersGetRequest {
	r.ratePlanId = &ratePlanId
	return r
}
func (r ApiBookingRatePlanOffersGetRequest) Arrival(arrival string) ApiBookingRatePlanOffersGetRequest {
	r.arrival = &arrival
	return r
}
func (r ApiBookingRatePlanOffersGetRequest) Departure(departure string) ApiBookingRatePlanOffersGetRequest {
	r.departure = &departure
	return r
}
func (r ApiBookingRatePlanOffersGetRequest) Adults(adults int32) ApiBookingRatePlanOffersGetRequest {
	r.adults = &adults
	return r
}
func (r ApiBookingRatePlanOffersGetRequest) ChannelCode(channelCode string) ApiBookingRatePlanOffersGetRequest {
	r.channelCode = &channelCode
	return r
}
func (r ApiBookingRatePlanOffersGetRequest) ChildrenAges(childrenAges []int32) ApiBookingRatePlanOffersGetRequest {
	r.childrenAges = &childrenAges
	return r
}
func (r ApiBookingRatePlanOffersGetRequest) IncludeUnavailable(includeUnavailable bool) ApiBookingRatePlanOffersGetRequest {
	r.includeUnavailable = &includeUnavailable
	return r
}

func (r ApiBookingRatePlanOffersGetRequest) Execute() (StayOffersModel, *_nethttp.Response, error) {
	return r.ApiService.BookingRatePlanOffersGetExecute(r)
}

/*
 * BookingRatePlanOffersGet Returns offers for a specific rate plan.
 * Calculates and returns available offers for a specific rate plan, arrival and departure date.<br>You must have at least one of these scopes: 'offers.read, reservations.manage'.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiBookingRatePlanOffersGetRequest
 */
func (a *OfferApiService) BookingRatePlanOffersGet(ctx _context.Context) ApiBookingRatePlanOffersGetRequest {
	return ApiBookingRatePlanOffersGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return StayOffersModel
 */
func (a *OfferApiService) BookingRatePlanOffersGetExecute(r ApiBookingRatePlanOffersGetRequest) (StayOffersModel, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  StayOffersModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OfferApiService.BookingRatePlanOffersGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/booking/v1/rate-plan-offers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.ratePlanId == nil {
		return localVarReturnValue, nil, reportError("ratePlanId is required and must be specified")
	}
	if r.arrival == nil {
		return localVarReturnValue, nil, reportError("arrival is required and must be specified")
	}
	if r.departure == nil {
		return localVarReturnValue, nil, reportError("departure is required and must be specified")
	}
	if r.adults == nil {
		return localVarReturnValue, nil, reportError("adults is required and must be specified")
	}

	localVarQueryParams.Add("ratePlanId", parameterToString(*r.ratePlanId, ""))
	localVarQueryParams.Add("arrival", parameterToString(*r.arrival, ""))
	localVarQueryParams.Add("departure", parameterToString(*r.departure, ""))
	if r.channelCode != nil {
		localVarQueryParams.Add("channelCode", parameterToString(*r.channelCode, ""))
	}
	localVarQueryParams.Add("adults", parameterToString(*r.adults, ""))
	if r.childrenAges != nil {
		localVarQueryParams.Add("childrenAges", parameterToString(*r.childrenAges, "csv"))
	}
	if r.includeUnavailable != nil {
		localVarQueryParams.Add("includeUnavailable", parameterToString(*r.includeUnavailable, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v MessageItemCollection
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiBookingServiceOffersGetRequest struct {
	ctx _context.Context
	ApiService *OfferApiService
	ratePlanId *string
	arrival *string
	departure *string
	adults *int32
	channelCode *string
	childrenAges *[]int32
	onlyDefaultDates *bool
}

func (r ApiBookingServiceOffersGetRequest) RatePlanId(ratePlanId string) ApiBookingServiceOffersGetRequest {
	r.ratePlanId = &ratePlanId
	return r
}
func (r ApiBookingServiceOffersGetRequest) Arrival(arrival string) ApiBookingServiceOffersGetRequest {
	r.arrival = &arrival
	return r
}
func (r ApiBookingServiceOffersGetRequest) Departure(departure string) ApiBookingServiceOffersGetRequest {
	r.departure = &departure
	return r
}
func (r ApiBookingServiceOffersGetRequest) Adults(adults int32) ApiBookingServiceOffersGetRequest {
	r.adults = &adults
	return r
}
func (r ApiBookingServiceOffersGetRequest) ChannelCode(channelCode string) ApiBookingServiceOffersGetRequest {
	r.channelCode = &channelCode
	return r
}
func (r ApiBookingServiceOffersGetRequest) ChildrenAges(childrenAges []int32) ApiBookingServiceOffersGetRequest {
	r.childrenAges = &childrenAges
	return r
}
func (r ApiBookingServiceOffersGetRequest) OnlyDefaultDates(onlyDefaultDates bool) ApiBookingServiceOffersGetRequest {
	r.onlyDefaultDates = &onlyDefaultDates
	return r
}

func (r ApiBookingServiceOffersGetRequest) Execute() (ServiceOffersModel, *_nethttp.Response, error) {
	return r.ApiService.BookingServiceOffersGetExecute(r)
}

/*
 * BookingServiceOffersGet Returns service offers for one specific stay.
 * <br>You must have at least one of these scopes: 'offers.read, reservations.manage'.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiBookingServiceOffersGetRequest
 */
func (a *OfferApiService) BookingServiceOffersGet(ctx _context.Context) ApiBookingServiceOffersGetRequest {
	return ApiBookingServiceOffersGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return ServiceOffersModel
 */
func (a *OfferApiService) BookingServiceOffersGetExecute(r ApiBookingServiceOffersGetRequest) (ServiceOffersModel, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ServiceOffersModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OfferApiService.BookingServiceOffersGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/booking/v1/service-offers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.ratePlanId == nil {
		return localVarReturnValue, nil, reportError("ratePlanId is required and must be specified")
	}
	if r.arrival == nil {
		return localVarReturnValue, nil, reportError("arrival is required and must be specified")
	}
	if r.departure == nil {
		return localVarReturnValue, nil, reportError("departure is required and must be specified")
	}
	if r.adults == nil {
		return localVarReturnValue, nil, reportError("adults is required and must be specified")
	}

	localVarQueryParams.Add("ratePlanId", parameterToString(*r.ratePlanId, ""))
	localVarQueryParams.Add("arrival", parameterToString(*r.arrival, ""))
	localVarQueryParams.Add("departure", parameterToString(*r.departure, ""))
	if r.channelCode != nil {
		localVarQueryParams.Add("channelCode", parameterToString(*r.channelCode, ""))
	}
	localVarQueryParams.Add("adults", parameterToString(*r.adults, ""))
	if r.childrenAges != nil {
		localVarQueryParams.Add("childrenAges", parameterToString(*r.childrenAges, "csv"))
	}
	if r.onlyDefaultDates != nil {
		localVarQueryParams.Add("onlyDefaultDates", parameterToString(*r.onlyDefaultDates, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v MessageItemCollection
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
