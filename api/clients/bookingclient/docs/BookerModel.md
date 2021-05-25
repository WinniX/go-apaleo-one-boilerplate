# BookerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Title of the booker | [optional] 
**Gender** | Pointer to **string** | Gender of the booker | [optional] 
**FirstName** | Pointer to **string** | First name of the booker | [optional] 
**MiddleInitial** | Pointer to **string** | Middle initial of the booker | [optional] 
**LastName** | **string** | Last name of the booker | 
**Email** | Pointer to **string** | Email address of the booker | [optional] 
**Phone** | Pointer to **string** | Phone number of the booker | [optional] 
**Address** | Pointer to [**PersonAddressModel**](PersonAddressModel.md) |  | [optional] 
**NationalityCountryCode** | Pointer to **string** | The booker&#39;s nationality, in ISO 3166-1 alpha-2 code | [optional] 
**IdentificationNumber** | Pointer to **string** | The booker&#39;s identification number for the given identificationType. | [optional] 
**IdentificationIssueDate** | Pointer to **string** | The issue date of the booker&#39;s identification document. | [optional] 
**IdentificationType** | Pointer to **string** | The type of the identificationNumber | [optional] 
**Company** | Pointer to [**PersonCompanyModel**](PersonCompanyModel.md) |  | [optional] 
**PreferredLanguage** | Pointer to **string** | Preferred contact two-letter language code (ISO Alpha-2) | [optional] 
**BirthDate** | Pointer to **string** | Birth date | [optional] 
**BirthPlace** | Pointer to **string** | The place of birth | [optional] 

## Methods

### NewBookerModel

`func NewBookerModel(lastName string, ) *BookerModel`

NewBookerModel instantiates a new BookerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookerModelWithDefaults

`func NewBookerModelWithDefaults() *BookerModel`

NewBookerModelWithDefaults instantiates a new BookerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *BookerModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BookerModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BookerModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BookerModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetGender

`func (o *BookerModel) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *BookerModel) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *BookerModel) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *BookerModel) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetFirstName

`func (o *BookerModel) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *BookerModel) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *BookerModel) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *BookerModel) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetMiddleInitial

`func (o *BookerModel) GetMiddleInitial() string`

GetMiddleInitial returns the MiddleInitial field if non-nil, zero value otherwise.

### GetMiddleInitialOk

`func (o *BookerModel) GetMiddleInitialOk() (*string, bool)`

GetMiddleInitialOk returns a tuple with the MiddleInitial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleInitial

`func (o *BookerModel) SetMiddleInitial(v string)`

SetMiddleInitial sets MiddleInitial field to given value.

### HasMiddleInitial

`func (o *BookerModel) HasMiddleInitial() bool`

HasMiddleInitial returns a boolean if a field has been set.

### GetLastName

`func (o *BookerModel) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *BookerModel) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *BookerModel) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetEmail

`func (o *BookerModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BookerModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BookerModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BookerModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *BookerModel) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *BookerModel) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *BookerModel) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *BookerModel) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetAddress

`func (o *BookerModel) GetAddress() PersonAddressModel`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BookerModel) GetAddressOk() (*PersonAddressModel, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BookerModel) SetAddress(v PersonAddressModel)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *BookerModel) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetNationalityCountryCode

`func (o *BookerModel) GetNationalityCountryCode() string`

GetNationalityCountryCode returns the NationalityCountryCode field if non-nil, zero value otherwise.

### GetNationalityCountryCodeOk

`func (o *BookerModel) GetNationalityCountryCodeOk() (*string, bool)`

GetNationalityCountryCodeOk returns a tuple with the NationalityCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalityCountryCode

`func (o *BookerModel) SetNationalityCountryCode(v string)`

SetNationalityCountryCode sets NationalityCountryCode field to given value.

### HasNationalityCountryCode

`func (o *BookerModel) HasNationalityCountryCode() bool`

HasNationalityCountryCode returns a boolean if a field has been set.

### GetIdentificationNumber

`func (o *BookerModel) GetIdentificationNumber() string`

GetIdentificationNumber returns the IdentificationNumber field if non-nil, zero value otherwise.

### GetIdentificationNumberOk

`func (o *BookerModel) GetIdentificationNumberOk() (*string, bool)`

GetIdentificationNumberOk returns a tuple with the IdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationNumber

`func (o *BookerModel) SetIdentificationNumber(v string)`

SetIdentificationNumber sets IdentificationNumber field to given value.

### HasIdentificationNumber

`func (o *BookerModel) HasIdentificationNumber() bool`

HasIdentificationNumber returns a boolean if a field has been set.

### GetIdentificationIssueDate

`func (o *BookerModel) GetIdentificationIssueDate() string`

GetIdentificationIssueDate returns the IdentificationIssueDate field if non-nil, zero value otherwise.

### GetIdentificationIssueDateOk

`func (o *BookerModel) GetIdentificationIssueDateOk() (*string, bool)`

GetIdentificationIssueDateOk returns a tuple with the IdentificationIssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationIssueDate

`func (o *BookerModel) SetIdentificationIssueDate(v string)`

SetIdentificationIssueDate sets IdentificationIssueDate field to given value.

### HasIdentificationIssueDate

`func (o *BookerModel) HasIdentificationIssueDate() bool`

HasIdentificationIssueDate returns a boolean if a field has been set.

### GetIdentificationType

`func (o *BookerModel) GetIdentificationType() string`

GetIdentificationType returns the IdentificationType field if non-nil, zero value otherwise.

### GetIdentificationTypeOk

`func (o *BookerModel) GetIdentificationTypeOk() (*string, bool)`

GetIdentificationTypeOk returns a tuple with the IdentificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationType

`func (o *BookerModel) SetIdentificationType(v string)`

SetIdentificationType sets IdentificationType field to given value.

### HasIdentificationType

`func (o *BookerModel) HasIdentificationType() bool`

HasIdentificationType returns a boolean if a field has been set.

### GetCompany

`func (o *BookerModel) GetCompany() PersonCompanyModel`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *BookerModel) GetCompanyOk() (*PersonCompanyModel, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *BookerModel) SetCompany(v PersonCompanyModel)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *BookerModel) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetPreferredLanguage

`func (o *BookerModel) GetPreferredLanguage() string`

GetPreferredLanguage returns the PreferredLanguage field if non-nil, zero value otherwise.

### GetPreferredLanguageOk

`func (o *BookerModel) GetPreferredLanguageOk() (*string, bool)`

GetPreferredLanguageOk returns a tuple with the PreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLanguage

`func (o *BookerModel) SetPreferredLanguage(v string)`

SetPreferredLanguage sets PreferredLanguage field to given value.

### HasPreferredLanguage

`func (o *BookerModel) HasPreferredLanguage() bool`

HasPreferredLanguage returns a boolean if a field has been set.

### GetBirthDate

`func (o *BookerModel) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *BookerModel) GetBirthDateOk() (*string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *BookerModel) SetBirthDate(v string)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *BookerModel) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### GetBirthPlace

`func (o *BookerModel) GetBirthPlace() string`

GetBirthPlace returns the BirthPlace field if non-nil, zero value otherwise.

### GetBirthPlaceOk

`func (o *BookerModel) GetBirthPlaceOk() (*string, bool)`

GetBirthPlaceOk returns a tuple with the BirthPlace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthPlace

`func (o *BookerModel) SetBirthPlace(v string)`

SetBirthPlace sets BirthPlace field to given value.

### HasBirthPlace

`func (o *BookerModel) HasBirthPlace() bool`

HasBirthPlace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


