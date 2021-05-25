# GuestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Title of the guest | [optional] 
**Gender** | Pointer to **string** | Gender of the booker | [optional] 
**FirstName** | Pointer to **string** | First name of the guest | [optional] 
**MiddleInitial** | Pointer to **string** | Middle initial of the guest | [optional] 
**LastName** | **string** | Last name of the guest | 
**Email** | Pointer to **string** | Email address of the guest | [optional] 
**Phone** | Pointer to **string** | Phone number of the guest | [optional] 
**Address** | Pointer to [**PersonAddressModel**](PersonAddressModel.md) |  | [optional] 
**NationalityCountryCode** | Pointer to **string** | The guest&#39;s nationality, in ISO 3166-1 alpha-2 code | [optional] 
**IdentificationNumber** | Pointer to **string** | The guest&#39;s identification number for the given identificationType. | [optional] 
**IdentificationIssueDate** | Pointer to **string** | The issue date of the guest&#39;s identification document. | [optional] 
**IdentificationType** | Pointer to **string** | The type of the identificationNumber | [optional] 
**Company** | Pointer to [**PersonCompanyModel**](PersonCompanyModel.md) |  | [optional] 
**PreferredLanguage** | Pointer to **string** | Two-letter code (ISO Alpha-2) of a language preferred for contact | [optional] 
**BirthDate** | Pointer to **string** | Guest&#39;s birthdate | [optional] 
**BirthPlace** | Pointer to **string** | Guest&#39;s place of birth | [optional] 

## Methods

### NewGuestModel

`func NewGuestModel(lastName string, ) *GuestModel`

NewGuestModel instantiates a new GuestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuestModelWithDefaults

`func NewGuestModelWithDefaults() *GuestModel`

NewGuestModelWithDefaults instantiates a new GuestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *GuestModel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GuestModel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GuestModel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GuestModel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetGender

`func (o *GuestModel) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *GuestModel) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *GuestModel) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *GuestModel) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetFirstName

`func (o *GuestModel) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *GuestModel) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *GuestModel) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *GuestModel) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetMiddleInitial

`func (o *GuestModel) GetMiddleInitial() string`

GetMiddleInitial returns the MiddleInitial field if non-nil, zero value otherwise.

### GetMiddleInitialOk

`func (o *GuestModel) GetMiddleInitialOk() (*string, bool)`

GetMiddleInitialOk returns a tuple with the MiddleInitial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleInitial

`func (o *GuestModel) SetMiddleInitial(v string)`

SetMiddleInitial sets MiddleInitial field to given value.

### HasMiddleInitial

`func (o *GuestModel) HasMiddleInitial() bool`

HasMiddleInitial returns a boolean if a field has been set.

### GetLastName

`func (o *GuestModel) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *GuestModel) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *GuestModel) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetEmail

`func (o *GuestModel) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GuestModel) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GuestModel) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GuestModel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *GuestModel) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *GuestModel) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *GuestModel) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *GuestModel) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetAddress

`func (o *GuestModel) GetAddress() PersonAddressModel`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GuestModel) GetAddressOk() (*PersonAddressModel, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GuestModel) SetAddress(v PersonAddressModel)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GuestModel) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetNationalityCountryCode

`func (o *GuestModel) GetNationalityCountryCode() string`

GetNationalityCountryCode returns the NationalityCountryCode field if non-nil, zero value otherwise.

### GetNationalityCountryCodeOk

`func (o *GuestModel) GetNationalityCountryCodeOk() (*string, bool)`

GetNationalityCountryCodeOk returns a tuple with the NationalityCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalityCountryCode

`func (o *GuestModel) SetNationalityCountryCode(v string)`

SetNationalityCountryCode sets NationalityCountryCode field to given value.

### HasNationalityCountryCode

`func (o *GuestModel) HasNationalityCountryCode() bool`

HasNationalityCountryCode returns a boolean if a field has been set.

### GetIdentificationNumber

`func (o *GuestModel) GetIdentificationNumber() string`

GetIdentificationNumber returns the IdentificationNumber field if non-nil, zero value otherwise.

### GetIdentificationNumberOk

`func (o *GuestModel) GetIdentificationNumberOk() (*string, bool)`

GetIdentificationNumberOk returns a tuple with the IdentificationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationNumber

`func (o *GuestModel) SetIdentificationNumber(v string)`

SetIdentificationNumber sets IdentificationNumber field to given value.

### HasIdentificationNumber

`func (o *GuestModel) HasIdentificationNumber() bool`

HasIdentificationNumber returns a boolean if a field has been set.

### GetIdentificationIssueDate

`func (o *GuestModel) GetIdentificationIssueDate() string`

GetIdentificationIssueDate returns the IdentificationIssueDate field if non-nil, zero value otherwise.

### GetIdentificationIssueDateOk

`func (o *GuestModel) GetIdentificationIssueDateOk() (*string, bool)`

GetIdentificationIssueDateOk returns a tuple with the IdentificationIssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationIssueDate

`func (o *GuestModel) SetIdentificationIssueDate(v string)`

SetIdentificationIssueDate sets IdentificationIssueDate field to given value.

### HasIdentificationIssueDate

`func (o *GuestModel) HasIdentificationIssueDate() bool`

HasIdentificationIssueDate returns a boolean if a field has been set.

### GetIdentificationType

`func (o *GuestModel) GetIdentificationType() string`

GetIdentificationType returns the IdentificationType field if non-nil, zero value otherwise.

### GetIdentificationTypeOk

`func (o *GuestModel) GetIdentificationTypeOk() (*string, bool)`

GetIdentificationTypeOk returns a tuple with the IdentificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationType

`func (o *GuestModel) SetIdentificationType(v string)`

SetIdentificationType sets IdentificationType field to given value.

### HasIdentificationType

`func (o *GuestModel) HasIdentificationType() bool`

HasIdentificationType returns a boolean if a field has been set.

### GetCompany

`func (o *GuestModel) GetCompany() PersonCompanyModel`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *GuestModel) GetCompanyOk() (*PersonCompanyModel, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *GuestModel) SetCompany(v PersonCompanyModel)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *GuestModel) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetPreferredLanguage

`func (o *GuestModel) GetPreferredLanguage() string`

GetPreferredLanguage returns the PreferredLanguage field if non-nil, zero value otherwise.

### GetPreferredLanguageOk

`func (o *GuestModel) GetPreferredLanguageOk() (*string, bool)`

GetPreferredLanguageOk returns a tuple with the PreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLanguage

`func (o *GuestModel) SetPreferredLanguage(v string)`

SetPreferredLanguage sets PreferredLanguage field to given value.

### HasPreferredLanguage

`func (o *GuestModel) HasPreferredLanguage() bool`

HasPreferredLanguage returns a boolean if a field has been set.

### GetBirthDate

`func (o *GuestModel) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *GuestModel) GetBirthDateOk() (*string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *GuestModel) SetBirthDate(v string)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *GuestModel) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### GetBirthPlace

`func (o *GuestModel) GetBirthPlace() string`

GetBirthPlace returns the BirthPlace field if non-nil, zero value otherwise.

### GetBirthPlaceOk

`func (o *GuestModel) GetBirthPlaceOk() (*string, bool)`

GetBirthPlaceOk returns a tuple with the BirthPlace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthPlace

`func (o *GuestModel) SetBirthPlace(v string)`

SetBirthPlace sets BirthPlace field to given value.

### HasBirthPlace

`func (o *GuestModel) HasBirthPlace() bool`

HasBirthPlace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


