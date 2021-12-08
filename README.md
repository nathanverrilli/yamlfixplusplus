openapi: 3.0.3

info:
title: XTRM Payments API
version: 4.8.0
description: >
The XTRM Payments API
termsOfService: "https://www.xtrm.com/web/Participant/Profile/TermsPopup.aspx"
contact:
email: apisupport@xtrm.com
name: XTRM API Help desk
license:
name: OFL

#########################
# RegEx to find descriptions on the same line
# and move them to a > format
#      // ^([\s]*)description:[\s]+([^\|^>].+)$
#      // \1description: > \n\1  \2
# note: will ALSO match description fields, as opposed to
#       OAS description parameters :-(
#########################

tags:
# new tags

- name: Authorization
  description: API Authorization API calls.

- name: Banks
  description: API calls dealing with bank connections and information

- name: Beneficiary Company
  description: API calls dealing with corporate beneficiaries

- name: Beneficiary User
  description: API calls dealing with personal beneficiaries

- name: OTP
  description: API calls to manage one time passwords

- name: Payment Methods
  description: API calls dealing with payment methods

- name: Programs
  description: API calls to manage remittance programs

- name: Transfer Funds
  description: API calls dealing with moving funds

- name: Wallets
  description: API calls dealing with wallets

- name: Advanced Profile
  description: API calls dealing with the advanced profile

- name: Deprecated
  description: API calls that are specifically deprecated

######################## OLD TAGS ##########################
#- name: AUTHORIZATION
# description: Authorization token API
#- name: XTRM AnySource
#  description: API Operations for XTRM AnySource
#- name: REMITTER API REFERENCE
#  description: API Operations for XTRM Remitters
#- name: BENEFICIARY API REFERENCE
#  description: API Operations for XTRM Beneficiaries
#- name: LINK BANK
#  description: API Operations for connecting remitters and beneficiaries with banks
#- name: COMPANY ADVANCED SERVICES
#  description: Advanced API services

#########################

security:
- bearerAuth: [ ]

# all requests require bearer-authorization
#########################

servers:
- url: "https://xapisandbox.xtrm.com/API/V4"
#- url: '{protocol}://{environment}.xtrm.com/API/v4'
#  variables:
#    protocol:
#      default: https
#      enum:
#        - https
#      description: ''
#    environment:
#      default: xapisandbox
#      enum:
#        - xapisandbox
#        - production
#        - staging
#        - devTrunk
#        - devFork
#      description: sandbox
#  description: sandbox authorization server payment
#########################

paths:
/oAuth/token:
servers:
- url: https://xapisandbox.xtrm.com
summary: Authorization token request
description: >
Fetch initial authorization token
post:
security:
- { }
tags:
- Authorization

      summary: Get or renew an authorization token
      description: |
        ## Initial authorization token request

        This API call has no security; it is the basis of
        the bearer authentication. Security is provided by
        the use of secure http (https) and the remitter’s
        API credentials.

        `grant_type` is either `password`
        to get a new session token, or `refresh_token`
        to renew a session.

        The `refresh_token` is present only when the
        `grant_type` is `refresh_token`.

      operationId: newAuthToken
      requestBody:
        description: >
          Get or renew an authorization token
        # required: true
        content:
          application/x-www-form-urlencoded:
            schema:
              $ref: "#/components/schemas/authorizationTokenRequest"
      responses:
        "200":
          description: >
            Authorization token and associated data as a JSON object
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/authorizationTokenResponse"
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"

/Register/GetCompanyAdvancedProfileDetails:
post:
operationId: GetCompanyAdvancedProfileDetails
summary: Get the advanced profile details for this or a connected company
description: >
This API fetches the advanced profile details for a company for
which your account has read access (either your company, or a
connected company). **This call’s parameters will be changed
in the next release as a breaking change.**
tags:
- Beneficiary Company
- Advanced Profile

      requestBody:
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/GetCompanyAdvancedProfileDetailRequest"

      responses:
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/GetCompanyAdvancedProfileDetailResponse"

        default:
          description: >
            Error response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"


/Register/CompanyAdvancedProfile:
post:
operationId: CompanyAdvancedProfile
summary: Set advanced profile
description: >

        This API calls allows you to fill and
        submit your or your connected company’s
        Advanced Profile. There are 4 sections in the advanced profile.
        Business Entity Information, Authorized Contact Information,
        Director Information, and Ownership Information.
        If you are completing this form for a US company,
        you must complete Business Entity Information and
        Authorized Contact Information; the rest of the sections
        are not required. If you are completing this form for
        a non US company, you must complete all the four sections
        (Business entity, Authorized contact, Director
        and Ownership information).
         **This call’s parameters will be changed
        in the next release as a breaking change.**
      tags:
        - Beneficiary Company
        - Advanced Profile

      requestBody:
        description: >
          Advanced profile information request
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/CompanyAdvancedProfileRequest"
            examples:
              canonical:
                value:
                  AdvancedProfile:
                    request:
                      issuerAccountNumber: SPN1122334
                      BeneficiaryAccountNumber: SPN5566778
                      Country_Of_Registration_Code2: 'US'
                      BusinessEntityInformation:
                        Legal_Company_Name: XYZ Inc
                        Fictitious_Name: Assumed
                        Website: https://www.xyz.com
                        Company_Type: Inc
                        Industry_Classification: Information Technology
                        Date_Of_Registration: '10/10/2010'
                        Region_Of_Incorporation: Centro
                        Ticker_Symbol: R
                        Tax_Identification_Number: 555-66-7777
                        Address_Line_1: The CommericalCo Home Office
                        Address_Line_2: 1 Calle Xerimba
                        Address_Line_3: Suite #314
                        City: Coimbra
                        Country_Code2: PT
                        Region_Code2: CT
                        Postal_Code: 11223344
                        Number_Of_Employees: 100
                        Expected_Monthly_Payments: 500
                        Expected_Monthly_Volume: 500000
                        Purpose_Of_Payments: Vendor Payments
                        Countries_Sending_Payments_To: United States of America
                        Countries_Receiving_Payments_From: Portugal
                      AuthorizedContactInformation:
                        Job_Title: CEO
                        First_Name: Jonathan
                        Middle_Name: Dough
                        Last_Name: Joe
                        Gender: Male
                        Email_Address: noreply@xtrm.com
                        Date_Of_Birth: '10/10/1990'
                        Citizenship: Portugal
                        Identification_Type: Passport
                        Identification_Number: 'XX123456'
                        Issue_Date: '10/10/2020'
                        Expiration_Date: '10/10/2030'
                        Issuing_Agency: XXX Inc.
                        Address_1: The CommericalCo Home Office
                        City: Coimbra
                        Region_Code2: CT
                        Country_Code2: PT
                        Postal_Code: 11223344
                        Business_Phone: 998877665544
                      DirectorInformation:
                        Directors:
                          - Job_Title: CEO
                            First_Name: Jonathan
                            Last_Name: Dough
                            Date_of_Birth: '10/10/1990'
                            Nationality: Portugal
                            Identification_Type: Passport
                            Identity_Document_Number: 'XX123456'
                            Identity_Document_Expiration: '10/10/2030'
                            Jurisdiction: Centro
                            Address_1: The CommericalCo Home Office
                            City: Coimbra
                            Country_Code2: PT
                            Region_Code2: CT
                            Postal_Code: 11223344
                      OwnerInformation:
                        IsPublicallyTraded: 'no'
                        Owners:
                          - Percentage_Owned: 25
                            First_Name: Jonathan
                            Last_Name: Dough
                            Occupation: Director
                            Source_of_Income: Business
                            Date_of_Birth: '10/10/1990'
                            Nationality: Portugal
                            Identity_Document_Type: Passport
                            Identity_Document_Number: 'XX123456'
                            Identity_Document_Expiration: '10/10/2030'
                            Jurisdiction: Centro
                            Address_1: The CommericalCo Home Office
                            City: Coimbra
                            Country_Code2: PT
                            Region_Code2: CT
                            Postal_Code: 11223344
      responses:
        "200":
          description: >
            success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/CompanyAdvancedProfileResponse"
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"

/Register/CreateUser:
post:
operationId: CreateUser
summary: Create a beneficiary account
description: >
## Create Beneficiary User Account

        This creates a beneficiary user account where the
        beneficiary can send and receive all payments.
        A wallet is also created for the user, holding the
        currency for the country specified during creation.
        These will be reviewed for KYC (*Know Your
        Customer*) compliance so information
        submitted must be real and accurate.
        **Do not create test users or emails on
        production servers.**

      tags:
        - Beneficiary User

      requestBody:
        description: >
          information about user
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/CreateUser"
            examples:
              canonical:
                value:
                  CreateUser:
                    request:
                      IssuerAccountNumber: SPN1122334
                      LegalFirstName: Jonathan
                      LegalLastName: Dough
                      EmailAddress: noreply@xtrm.com
                      MobilePhone: '14085551234'
                      EmailNotification: true
                      TaxId: 555-66-7777
                      DateOfBirth:
                        Day: '22'
                        Month: '12'
                        Year: '1933'
                      Address:
                        AddressLine1: The CommericalCo Home Office
                        AddressLine2: "1 Calle Xerimba"
                        AptSuitNum: "Suite #314"
                        City: Coimbra
                        Country: Portugal
                        CountryISO2: PT
                        Region: Centro

      responses:
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/CreateUserResponse"
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"

/Register/UpdateUser:
post:
operationId: UpdateUser
description: >
## Update Beneficiary User Account

        This will update a user’s beneficiary account
        or digital wallet.

      summary: Update a beneficiary user
      tags:
        - Beneficiary User

      requestBody:
        description: >
          updated information
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/UpdateUserRequest"

      responses:
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/UpdateUserResponse"
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"

/Beneficiary/CheckUserExist:
post:
operationId: CheckUserExist
description: >
# Check User Exist

        This returns some minimal details about a user,
        if the user exists.

      summary: Check beneficiary existence
      tags:
        - Beneficiary User

      requestBody:
        description: >
          Request and email to check
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/CheckUserExistRequest"

      responses:
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/CheckUserExistResponse"

/Fund/GetConnectedCompanyFundRequest:
post:
operationId: GetConnectedCompanyFundRequest
description: >
# There is no documentation on this API call.

      summary: Get connected company fund request

      tags:
        - Transfer Funds

      requestBody:
        description: >
          Request and pagination information
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/GetConnectedCompanyFundRequestRequest"

      responses:
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/GetConnectedCompanyFundRequestResponse"

/Beneficiary/GetPersonalBeneficiaries:
post:
operationId: GetPersonalBeneficiaries
description: >
# Get Personal Beneficiaries

        This allows you to get a list of your personal (non-company) beneficiaries.

      summary: Get personal beneficiaries

      tags:
        - Beneficiary User

      requestBody:
        description: >
          Request and pagination information
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/GetPersonalBeneficiariesRequest"

      responses:
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/GetPersonalBeneficiariesResponse"

/Payment/GetUserPaymentMethods:
post:
deprecated: true
operationId: GetUserPaymentMethods
summary: Get send payment
description: >
# Remitter Payment Methods

        ** Deprecated **

        **This has been replaced by GetPaymentMethods.**

        Choose the payment method.
        You can pay to the XTRM &lsquo;AnyPay™’
        digital wallet
        (recommended) where the user can choose how to
        transfer the funds out of XTRM.
        Alternatively you create a pass through payment
        directly to a number of end points such as
        bank account, gift cards, or a prepaid debit
        card.

        This call has an empty JSON body: `{ }`


      tags:
        - Payment Methods
        - Deprecated
      requestBody:
        description: >
          The body consists of the empty JSON string:<br />
          **`{ }`**
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/xEmpty"

      responses:
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/GetUserPaymentMethodsResponse"

/Payment/GetPaymentMethods:
post:
operationId: GetPaymentMethods
summary: Get send payment methods
description: >
# Remitter Payment Methods

        Choose the payment method.
        You can pay to the XTRM AnyPay™
        digital wallet
        (recommended) where the user can choose how to
        transfer the funds out of XTRM.
        Alternatively you create a pass through payment
        directly to a number of end points such as
        bank account, gift cards, or a prepaid debit
        card.

        The body of this request consist of the empty
        JSON string <b>` { } `</b>
      tags:
        - Payment Methods
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/xEmpty"

      responses:
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/GetPaymentMethodsResponse"

/Payment/GetBeneficiaryCompanyPaymentMethods:
post:
operationId: GetBeneficiaryCompanyPaymentMethods
summary: Get payment methods for a beneficary company
description: >
A beneficiary company may transfer to bank endpoints from
their XTRM AnyPay™ digital wallet.

        The body consists of the empty JSON string:
        **` { } `**
      tags:
        - Payment Methods

      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/xEmpty"

      responses:
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/GetBeneficiaryCompanyPaymentsResponse"

/Wallet/GetUserWalletBalance:
post:
operationId: GetUserWalletBalance
summary: Fetch the balance of a user’s wallet
description: >
Get an individual user’s wallet balance
for the specified currency
tags:
- Wallets
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/GetUserWalletBalanceRequest"

      responses:
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                anyOf:
                  - $ref: "#/components/schemas/UserWalletBalanceResponse"
                  - $ref: "#/components/schemas/UserWalletBalanceResponseDeprecated"
                  #containing field is misspelled

/Wallet/GetCompanyWallets:
post:
operationId: GetCompanyWallets
summary: Get company wallets
description: >
Get a list of wallets for a particular company
tags:
- Wallets

      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/GetCompanyWalletsRequest"

      responses:
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/GetCompanyWalletsResponse"

/Wallet/GetBeneficiaryWallets:
post:
operationId: GetBeneficiaryWallets
summary: Get wallets for a particular beneficary
description: >
Fetch the wallets belonging to a beneficiary
tags:
- Wallets
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/GetBeneficiaryWalletsRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/GetBeneficiaryWalletsResponse"

/Wallet/CreateUserWallet:
post:
operationId: CreateUserWallet
summary: Create user wallet
description: >
<p>Create a wallet for a beneficiary with a specific name and currency.</p>
tags:
- Wallets
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/CreateUserWalletRequest"

      responses:
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/CreateUserWalletResponse"

/Wallet/CreateCompanyWallet:
post:
operationId: CreateCompanyWallet
summary: Create company wallet
description: >
<p>Create a wallet with a specific name and currency.</p>
tags:
- Wallets
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/CreateCompanyWalletRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/CreateCompanyWalletResponse"

/Wallet/CreateBeneficiaryCompanyWallet:
post:
operationId: CreateBeneficiaryCompanyWallet
summary: Create beneficiary company wallet
tags:
- Wallets
description: >
create beneficiary company wallets and
specify the currency
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/CreateBeneficiaryCompanyWalletRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/CreateBeneficiaryCompanyWalletResponse"

/Wallet/UpdateUserWallet:
post:
operationId: UpdateUserWallet
summary: Update user wallet
description: >
Change the name of a user’s wallet. The other elements
of a wallet (the `WalletID` and currency) cannot be changed.
tags:
- Wallets
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/UpdateUserWalletRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/UpdateUserWalletResponse"

/Wallet/UpdateCompanyWallet:
post:
operationId: UpdateCompanyWallet
summary: Update Company Wallet
description: >
Change the name of a company’s wallet. The other elements
of a wallet (the `WalletID` and currency) cannot be changed.
tags:
- Wallets
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/UpdateCompanyWalletRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/UpdateCompanyWalletResponse"

/Beneficiary/UpdateBeneficiary:
post:
operationId: UpdateBeneficiary
summary: Update beneficiary
description: |
# THIS CALL LACKS DOCUMENTATION
`/Beneficiary/UpdateBeneficiary` lacks documentation.
tags:
- Beneficiary Company
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/UpdateBeneficiaryRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/UpdateBeneficiaryResponse"

/Wallet/UpdateBeneficiaryCompanyWallet:
post:
operationId: UpdateBeneficiaryCompanyWallet
summary: Update a beneficiary company wallet
description: >
Update a beneficiary company’s name.
The other elements of a wallet (the
`WalletID` and currency) cannot be changed.
tags:
- Wallets
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/UpdateBeneficiaryCompanyWalletRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/UpdateBeneficiaryCompanyWalletResponse"

/Wallet/GetUserWalletTransactions:
post:
operationId: GetUserWalletTransactions
summary: Get user wallet transactions
tags:
- Wallets
description: >
Get a list of transactions of a user’s wallet.
A user may have multiple wallets.
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/GetUserWalletTransactionsRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/GetUserWalletTransactionsResponse"

/Wallet/GetCompanyWalletTransactions:
post:
operationId: GetCompanyWalletTransactions
summary: Get company wallet transactions
tags:
- Wallets
description: >
Get all the wallet transactions. There can be one or more wallets.
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/GetCompanyWalletTransactionsRequest"

      responses:
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/GetCompanyWalletTransactionsResponse"

/Wallet/GetCompanyWalletTransactionDetails:
post:
operationId: GetCompanyWalletTransactionDetails
summary: Get company wallet transaction details
tags:
- Wallets
description: >
get the transaction details for the selected transaction ID
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/GetCompanyWalletTransactionDetailsRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/GetCompanyWalletTransactionDetailsResponse"

/Wallet/FundCompanyWalletUsingCreditCard:
post:
operationId: FundCompanyWalletUsingCreditCard
summary: Fund company wallet using credit card
description: >
companies may fund their own company wallets using
the company’s own credit card.
** *Please note:* This call is not intended to take payments
from consumers; it meant only for companies to
fund their own wallets.**
tags:
- Wallets
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/FundCompanyWalletUsingCreditCardRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/FundCompanyWalletUsingCreditCardResponse"

/Wallet/FundUserWalletUsingCreditCard:
post:
operationId: FundUserWalletUsingCreditCard
summary: Fund user wallet wallet using credit card
description: >
This allows companies to allow their connected users to fund
their wallets using their credit card. **Please contact
*support@xtrm.com* to enable this API call.**
tags:
- Wallets
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/FundUserWalletsUsingCreditCardRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/FundUserWalletsUsingCreditCardResponse"

/Wallet/FundWalletUsingACHDebit:
post:
operationId: FundWalletUsingACHDebit
summary: Fund wallet using ACH debit
tags:
- Wallets
description: >
This API call allows companies to fund their own wallets using
ACH (or systems similar to ACH) debit linked banks.
**Funding using ACH debit is instantaneous. However,
ACH debit may take from 3 to 5 business days to complete**
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/FundWalletUsingACHDebitRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/FundWalletUsingACHDebitResponse"

/Wallet/FundCompanyWalletUsingACHDebit:
post:
operationId: FundCompanyWalletUsingACHDebit
summary: Fund company wallet using ACH debit
tags:
- Wallets
description: >
This API call allows companies to fund their own wallets using
their ACH debit linked banks.
**Funding using ACH debit is instantaneous. However,
ACH debit may take from 3 to 5 business days to complete**
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/FundWalletUsingACHDebitRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/FundWalletUsingACHDebitResponse"

/Wallet/GetUserWallets:
post:
operationId: GetUserWallets
summary: Get user wallets
tags:
- Wallets
description: >
Get the wallets for a particular user
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/GetUserWalletsRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/GetUserWalletsResponse"

/Payment/GetExchangeRate:
post:
operationId: GetExchangeRate
summary: Get exchange rate
tags:
- Wallets
description: >
<p>Fetches the real-time exchange rate between two
currencies. This rate is real-time, and may
fluctuate slightly between a call to
`GetExchangeRate` and
`BookExchange`.
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/GetExchangeRateRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/GetExchangeRateResponse"

/Payment/BookExchange:
post:
operationId: BookExchange
summary: Book exchange (exchange currency)
description: >
Exchange currency. Funds in the original currency are
withdrawn from a wallet (that holds the original currency),
and exchanged for funds in a second currency. Those funds
are then deposited into a second wallet (that holds the new
currency).
tags:
- Wallets
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/BookExchangeRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/BookExchangeResponse"

/Wallet/GetUserWalletTransactionsByRemitter:
post:
operationId: GetUserWalletTransactionsByRemitter
summary: Get user transaction by remitter
description: >
Fetch a list of transactions of a specific user and remitter.
Records will contain amount sent by the remitter to the user.
tags:
- Wallets
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/GetUserWalletTransactionsByRemitterRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/GetUserWalletTransactionsByRemitterResponse"

/Fund/TransferFund:
post:
operationId: TransferFund
summary: Transfer funds
tags:
- Transfer Funds
description: >
This allows you to make transfers from a company wallet to
a beneficiary user wallet, (AnyPay™) or pass-through
payments to a beneficiaries Bank account or
Prepaid Virtual Visa. When you make transfers, if
the beneficiary does not have a wallet in that currency,
one will be created automatically.

        The `Transaction Detail` element provides information
        on when the transaction generates additional information
        (typically redemption details for gift cards or debit
        cards) that need to be displayed to the user in a web
        or mobile application.
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/TransferFundRequest"

      responses:
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/TransferFundResponse"

/Fund/UserWithdrawFund:
post:
operationId: UserWithdrawFund
summary: User withdraw fund
description: >
This allows beneficiary to withdraw funds using one of
the payment methods from their XTRM 'AnyPay™' Digital Wallet.
On UserWithdrawFund call a 6 digit one time password (OTP)
is generated and sent to user.
Get the 6 digit OTP value from the user in your application
and resubmit UserWithdrawFund with the OTP included.

      tags:
        - Transfer Funds

      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/UserWithdrawFundRequest"

      responses:
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/UserWithdrawFundResponse"

/Fund/TransferFundToCompany:
post:
operationId: TransferFundToCompany
summary: Transfer funds to company
tags:
- Transfer Funds
description: >
This allows you to make transfers from a company wallet
to a beneficiary company wallet, *&lsquo;AnyPay&trade’)
or pass-through payments to a beneficiaries Bank account
or Paypal. When you make transfers, if the beneficiary
does not have a wallet in that currency, one will
be created automatically.

        **Paypal transactions have been suspended. Please contact
        support@xtrm.com for Paypal connectivity requirements.**
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/TransferFundToCompanyRequest"
      responses:
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/TransferFundToCompanyResponse"

/Fund/BeneficiaryCompanyWithdrawFund:
post:
operationId: BeneficiaryCompanyWithdrawFund
summary: Beneficiary company withdraw funds
tags:
- Transfer Funds
description: |
<p>This allows beneficiary company to transfer funds
to their bank from their XTRM &lsquo;AnyPay&trade;’
Digital Wallet. On UserWithdrawFund call an OTP
(one time password)  is generated and sent to user.
Get the 6 digit OTP value from the user in your
application and resubmit `UserWithdrawFund` with
OTP included.</p>
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/BeneficiaryCompanyWithdrawFundRequest"

      responses:
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/BeneficiaryCompanyWithdrawFundResponse"

/Fund/TransferFundUsertoCompanyUsingCC:
post:
operationId: TransferFundUsertoCompanyUsingCC
summary: Transfer user funds to company using credit card
tags:
- Transfer Funds
description: >
This allows you to receive funds from a beneficiary user using
the beneficiary’s credit card.
This single call results in two sequential transactions.
First, funds are deposited in user account (the beneficiary’s wallet)
from the credit card.
Then, the funds are transferred from the beneficiary’s
wallet to the company’s wallet.
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/TransferFundUsertoCompanyUsingCCRequest"

      responses:
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/TransferFundUsertoCompanyUsingCCResponse"

/Fund/TransferFundWalletToWallet:
post:
operationId: TransferFundWalletToWallet
summary: Transfer Funds Wallet to Wallet
tags:
- Transfer Funds
description: >
This allows you to transfer funds between wallets
in any combination. i.e from a user wallet to
company wallet, user wallet to another user wallet,
company wallet to user wallet or company wallet or
another company wallet. On `TransferFundWallettoWallet`
call, a 6 digit one time password (OTP) is generated
and sent to the user from whom the fund is going to
be debited. Capture this 6 digit OTP value from the user
in your application and resubmit
`TransferFundWallettoWallet` with the OTP included.
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/TransferFundWalletToWalletRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/TransferFundWalletToWalletResponse"

/Fund/TransferFundDynamicAccountCreateUser:
post:
operationId: TransferFundDynamicAccountCreateUser
summary: Transfer funds to user’s default wallet (create user and wallet as needed)
tags:
- Transfer Funds
description: >
This allows the transfer of funds from a company wallet
to a new user or existing user’s wallet.
The user is created dynamically if the email id used
does not already exist within the system. The new user
is connected to the creating account.
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/TransferFundDynamicAccountCreateUserRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/TransferFundDynamicAccountCreateUserResponse"

/GiftCard/GetPrepaidCards:
post:
operationId: GetPrepaidCards
summary: Get prepaid cards
tags:
- Transfer Funds
description: >
<p>NO DOCUMENTATION</p>
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/GetPrepaidCardsRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/GetPrepaidCardsResponse"

/GiftCard/GetPrepaidCardDetails:
post:
operationId: GetPrepaidCardDetails
summary: Get prepaid card details
tags:
- Transfer Funds
description: >
<p>NO DOCUMENTATION</p>
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/GetPrepaidCardDetailsRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/GetPrepaidCardDetailsResponse"

/GiftCard/GetDigitalGiftCards:
post:
operationId: GetDigitalGiftCards
summary: Get digital gift cards (by currency)
tags:
- Transfer Funds
description: >
<p>Gets the list of digital gift cards supported by XTRM.
The list has a unique SKU which is used as input to the
`userwithdrawfund` API call.
Digital Gift Card is added as an additional
withdraw fund method for the user.
Available gift cards vary by currency.</p>
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/GetDigitalGiftCardsRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/GetDigitalGiftCardsResponse"

/GiftCard/GetGiftCardDetails:
post:
operationId: GetGiftCardDetails
summary: Get gift card details (by SKU)
tags:
- Transfer Funds
description: >
<p>Gets the specifics of a gift card supported by XTRM.
The list has a unique SKU which is used as input to the
`userwithdrawfund` API call. </p>

      requestBody:
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/GetGiftCardDetailsRequest"
      responses:
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/GetGiftCardDetailsResponse"

/Bank/GetLinkedBankAccounts:
post:
operationId: GetLinkedBankAccounts
summary: Get a beneficiary’s linked bank accounts
tags:
- Banks
description: >
Fetch a list of the beneficiary linked bank accounts.
This is used when making direct pass through payments
to a beneficiary’s bank account.
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/GetLinkedBankAccountsRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/GetLinkedBankAccountsResponse"

/Bank/GetACHDebitLinkedBankAccounts:
post:
operationId: GetACHDebitLinkedBankAccounts
summary: Get ACH debit linked bank accounts
description: >
Fetch a list of the ACH debit linked bank accounts
(or bank transfers systems analogous to ACH).
This is used to fund company wallets using the
ACH debit method.
tags:
- Banks
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/GetACHDebitLinkedBankAccountsRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/GetACHDebitLinkedBankAccountsResponse"

/Bank/GetBankWithdrawTypes:
post:
operationId: GetBankWithdrawTypes
summary: Get bank withdraw types
tags:
- Banks
description: >
Retrieves a list of bank withdrawal types for the bank
such as wire, ACH, and similar systems.
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/GetBankWithdrawTypesRequest"

      responses:
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/GetBankWithdrawTypesResponse"

/Bank/LinkBankBeneficiary:
post:
operationId: LinkBankBeneficiary
summary: Link beneficiary’s bank for transfers
tags:
- Banks
description: >
# Link Bank Beneficiary

        Linking a bank beneficiary ***may*** require multiple steps.
        Different jurisdictions require different or additional
        pieces of information.
        The first step involves submitting five always-required pieces of
        information.


        ### Required Information:

        - Bank account number

        - Bank country code

        - Beneficiary country code

        - Beneficiary name

        - Currency

        ### Additional Jurisdiction Information

        In some cases, additional information may be requested by the bank to
        complete the linkage depending on its regulatory and compliance
        environment. In such cases, the initial
        response may return with a 200 code, and the
        `LinkBankBeneficiary:request:OperationStatus: Success`
        may be **`false`**.
        The `Error` field will have an `ErrorMessage`,
        and report that an additional `FieldName` and
        `FieldLabel` are required
        (possibly **`"RoutingNumber"`** and
        **`"Account Number"`**, respectively.

        The call is then made a second time, filling out the
        required field (in this example, the field
        `LinkBankBeneficiary:request:Beneficiary:BankDetails:RoutingNumber`).

        Unusual data will be specified as going into `RemittanceLine3`
        or `RemittanceLine4`.

        ### Exotic Currency Requirements

        Some currencies will require currency-specific information to remit funds.
        This follows the two-step process above, specifying what additional
        data items are required to complete the linkage.

      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/LinkBankBeneficiaryRequest"
      responses:
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/LinkBankBeneficiaryResponse"

/Bank/LinkACHDebitBankBeneficiary:
post:
operationId: LinkACHDebitBankBeneficiary
summary: Link beneficiary’s Bank for ACH debit
description: >
# Link ACH Debit Bank Beneficiary

        Linking a bank beneficiary ***may*** require multiple steps.
        Different jurisdictions require different or additional
        pieces of information.
        The first step involves submitting five always-required pieces of
        information.


        ### Required Information:

        - Bank account number

        - Bank country code

        - Beneficiary country code

        - Beneficiary name

        - Currency

        ### Additional Jurisdiction Information

        In some cases, additional information may be requested by the bank to
        complete the linkage depending on its regulatory and compliance
        environment. In such cases, the initial
        response may return with a 200 code, and the
        `LinkBankBeneficiary:request:OperationStatus: Success`
        may be **`false`**.
        The `Error` field will have an `ErrorMessage`,
        and report that an additional `FieldName` and
        `FieldLabel` are required
        (possibly **`"RoutingNumber"`** and
        **`"Account Number"`**, respectively.

        The call is then made a second time, filling out the
        required field (in this example, the field
        `LinkACHDebitBankBeneficiary:request:Beneficiary:BankDetails:RoutingNumber`).

        Unusual data will be specified as going into `RemittanceLine3`
        or `RemittanceLine4`.

        ### Exotic Currency Requirements

        Some currencies will require currency-specific information to remit funds.
        This follows the two-step process above, specifying what additional
        data items are required to complete the linkage.
      tags:
        - Banks
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/LinkACHDebitBankBeneficiaryRequest"
      responses:
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/LinkACHDebitBankBeneficiaryResponse"

# NAME should this be DeleteBeneficiaryBank?
/Bank/DeleteBankBeneficiary:
post:
operationId: DeleteBankBeneficiary
summary: Delete a beneficiary’s linked bank
description: >
<p>Delete a beneficiary’s linked bank.
Get the required `BankBeneficiaryID` from the
`GetLinkedBankAccounts` API call.</p>
tags:
- Banks
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/DeleteBankBeneficiaryRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/DeleteBankBeneficiaryResponse"

/Card/GetLinkedCards:
post:
operationId: GetLinkedCards
summary: Get a beneficiary’s linked Card accounts
tags:
- Cards
description: >
Fetch a list of the beneficiary linked Card accounts.
This is used when making direct pass through payments
to a beneficiary’s Card account.
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/GetLinkedCardsRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/GetLinkedCardsResponse"

/Card/LinkCard:
post:
operationId: LinkCard
summary: Link beneficiary’s Card for transfers
tags:
- Cards
description: >
# Link Card

        Linking a Card for identity or transfer


        ### Required Information:

        - Card number

        - Card expiry date

        - Card CVV

        - Card name

        - Card Type

        ### Exotic Currency Requirements

        When link card type is transfer then require USD only.

      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/LinkCardRequest"
      responses:
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/LinkCardResponse"

/Card/DeleteCard:
post:
operationId: DeleteCard
summary: Delete a beneficiary’s linked Card
description: >
<p>Delete a beneficiary’s linked Card.
Get the required `CardBeneficiaryID` from the
`GetLinkedCards` API call.</p>
tags:
- Cards
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/DeleteCardRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/DeleteCardResponse"

/Beneficiary/GetBeneficiaries:
post:
operationId: GetBeneficiaries
summary: Get beneficiaries
tags:
- Beneficiary Company
description: >
Fetch a list of beneficiaries for the company.
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/GetBeneficiariesRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/GetBeneficiariesResponse"

/Beneficiary/CreateBeneficiary:
post:
operationId: CreateBeneficiary
summary: Create beneficiary company
description: |
<h2>Create a beneficiary company</h2>
<p>Beneficary companies will be reviewed for
KYC (<i>Know Your Customer</i>) compliance.
Ensure that the information is real and accurate.
<b>Do not create test companies or emails on
production servers.</b></p>
tags:
- Beneficiary Company
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/CreateBeneficiaryRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/CreateBeneficiaryResponse"

/Beneficiary/CheckBeneficiaryExist:
post:
operationId: CheckBeneficiaryExist
summary: Check company beneficiary existence
description: >
<p>Check to see if an existing email address is associated with an XTRM account</p>
<p>Often, companies have preexisting accounts
within the XTRM ecosystem. This API finds all
companies matching a particular name, and
returns some minimal information about them.</p>
tags:
- Beneficiary Company
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/CheckBeneficiaryExistRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/CheckBeneficiaryExistResponse"

/Register/GetNAICS:
post:
operationId: GetNAICS
summary: Get NAICS standard industry types
tags:
- Beneficiary Company
description: |
## Get standard industry job types
<p>This API call returns standard industry types</p>
<p>The body consists of the empty JSON object:
`{ }`</p>
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/xEmpty"

      responses:
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/GetNAICSResponse"

/Register/GetCompanyType:
post:
operationId: GetCompanyType
summary: Get company type (standard company types)
tags:
- Beneficiary Company
description: >
This API call returns standard company types.

        The body consists of the empty JSON object&colon;
        `{ }`
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/xEmpty"

      responses:
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/GetCompanyTypeResponse"

/Register/GetAdvancedContactJobTitles:
post:
operationId: GetAdvancedContactJobTitles
summary: Get standard job titles
tags:
- Beneficiary Company
description: |
<p>Fetch a list of standard job titles</p>
<p>&nbsp</p>
<p>This request takes an empty request body.</p>

      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/xEmpty"
      responses:
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/GetAdvancedContactJobTitlesResponse"

/Register/GetIdentificationType:
post:
operationId: GetIdentificationType
summary: Get identification types
tags:
- Beneficiary Company
description: |
<p>Get the list of identification types</p>
<p>&nbsp;</p>
<p>This request has an empty body</p>
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/xEmpty"

      responses:
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/GetIdentificationTypeResponse"

/Register/GetAdvancedContactCountry:
post:
operationId: GetAdvancedContactCountry
summary: Get countries and country codes
tags:
- Beneficiary Company
description: |
<p>Return the list of countries</p>
<p>&nbsp;</p>
<p>This request has an empty body</p>
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/xEmpty"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/GetAdvancedContactCountryResponse"

/Register/GetAdvancedContactState:
post:
operationId: GetAdvancedContactState
summary: Get states (regions) within a country
tags:
- Beneficiary Company
description: |
<p>For a particular country, get the list of
regions (states, provinces, regions, <i>etc</i></p>
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/GetAdvancedContactStateRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/GetAdvancedContactStateResponse"

/Register/CompanyAdvancedProfileStatus:
post:
operationId: CompanyAdvancedProfileStatus
summary: Get company profile status
tags:
- Beneficiary Company
- Advanced Profile
description: |
<p>This API call is used to get the status of the
remitter application using company advanced profile.
The status would be one of:</p>
<ul style="font-weight: bolder">
<li>Submitted</li> <li>Pending</li> <li>Approved</li>
</ul>

      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/CompanyAdvancedProfileStatusRequest"
      responses:
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/CompanyAdvancedProfileStatusResponse"

/Programs/GetPrograms:
post:
operationId: GetPrograms
summary: Get company created programs
tags:
- Programs
description: |
<p>This allows you to get a list of your company
payment programs or projects for use
with the Transfer Funds API request.</p>
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/GetProgramsRequest"

      responses:
        default:
          description: >
            Error Response
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/xErrorDefault"
        "200":
          description: >
            Success
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/GetProgramsResponse"

/Programs/ProgramCategory:
post:
operationId: ProgramCategory
summary: Fetch a list of program categories
tags:
- Programs
description: |
<p>Fetch a list of program categories.
Used along with the program type to
create new programs.</p>
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/ProgramCategoryRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/ProgramCategoryResponse"

/Programs/ProgramType:
post:
operationId: ProgramType
summary: Fetch a list of program types
tags:
- Programs
description: |
<p>Return the available program types</p>
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/ProgramTypeRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/ProgramTypeResponse"

/Programs/CreateProgram:
post:
operationId: CreateProgram
tags:
- Programs
summary: Create a company payment program
description: |
<p>Fetch a list of your company payment programs
or projects for use with the Transfer Funds
API request.</p>
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/CreateProgramRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/CreateProgramResponse"

/Programs/UpdatePrograms:
post:
operationId: UpdatePrograms
tags:
- Programs
description: |
<p>Update your company payment programs or projects</p>
summary: "Update company program"
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/UpdateProgramsRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/UpdateProgramsResponse"

/OTP/GetOTPAuthorizedVendor:
post:
operationId: GetOTPAuthorizedVendor
tags:
- OTP
summary: Generate OTP code for user
description: >
<p>This is used to generate an authentication one
time password to allow the vendor (remitter) to be
authorized by the user (beneficiary) to transfer funds,
access wallets, access wallet transactions and
withdraw funds.</p>
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/GetOTPAuthorizedVendorRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/GetOTPAuthorizedVendorResponse"

/OTP/ValidateOTPAuthorizeVendor:
post:
operationId: ValidateOTPAuthorizeVendor
summary: "Validate OTP"
tags:
- OTP
description: >
<p>This is used to validate the one time
password to allow the user to transfer funds
and is used in conjunction with
&lsquo;Authorize Vendor’</p>
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/ValidateOTPAuthorizeVendorRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/ValidateOTPAuthorizeVendorResponse"

/Wallet/GetUserWalletTransactionDetails:
post:
operationId: GetUserWalletTransactionDetails
summary: "Get details on a specific user transaction"
tags:
- Wallets
description: >
Get details of a specific transaction using the
unique transaction ID with
the `GetUserWalletTransactions` API call.
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/GetUserWalletTransactionDetailsRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/GetUserWalletTransactionDetailsResponse"

/Bank/SearchBank:
post:
operationId: SearchBank
summary: Search banks within a country
tags:
- Banks
description: >
Search banks within a country using bank name
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/SearchBankRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/SearchBankResponse"

/OTP/GetConnectedStatus:
post:
operationId: GetConnectedStatus
tags:
- OTP
summary: Get connected status
description: >
<p>This allows you to get the connected status
of the remitter and the user or company.
If they are not connected, you can use
`GetOTPAuthorizedVendor` to send a
one time password to the user or company to
connect and authorize the remitter followed
by `ValidateOTPAuthorizeVendor`
to validate the one time password and complete
the connection/authorization.</p>
requestBody:
required: true
content:
application/json:
schema:
$ref: "#/components/schemas/GetConnectedStatusRequest"
responses:
default:
description: >
Error Response
content:
application/json:
schema:
$ref: "#/components/schemas/xErrorDefault"
"200":
description: >
Success
content:
application/json:
schema:
$ref: "#/components/schemas/GetConnectedStatusResponse"

###########################################

components:
securitySchemes:
bearerAuth:
type: http
scheme: bearer

schemas:
xErrorDefault:
type: object
properties:
Errors:
description: >
The nature, type, and sort of error
type: array
items:
type: string


    xAuthorizedContactInformation:
      type: object
      required:
        - First_Name
        - Last_Name
        - Email_Address
        - Address_1
        - City
        - Region_Code2
        - Country_Code2
        - Postal_Code
      properties:
        Job_Title:
          $ref: "#/components/schemas/x255CharString"
        First_Name:
          $ref: "#/components/schemas/x255CharString"
        Last_Name:
          $ref: "#/components/schemas/x255CharString"
        Middle_Name:
          $ref: "#/components/schemas/x255CharString"
        Gender:
          type: string
          minLength: 5
          maxLength: 6
          example: "Female"
        Email_Address:
          $ref: "#/components/schemas/xEmail"
        Date_Of_Birth:
          $ref: "#/components/schemas/xNumericDateString"
        Citizenship:
          $ref: "#/components/schemas/x255CharString"
        Identification_Type:
          $ref: "#/components/schemas/x255CharString"
        Identification_Number:
          $ref: "#/components/schemas/x255CharString"
        Issue_Date:
          $ref: "#/components/schemas/xNumericDateString"
        Expiration_Date:
          $ref: "#/components/schemas/xNumericDateString"
        Issuing_Agency:
          $ref: "#/components/schemas/x255CharString"
        Address_1:
          $ref: "#/components/schemas/x255CharString"
        City:
          $ref: "#/components/schemas/x255CharString"
        Region_Code2:
          $ref: "#/components/schemas/xRegion_Code2"
        Country_Code2:
          $ref: "#/components/schemas/xCountryISO2"
        Postal_Code:
          $ref: "#/components/schemas/x255CharString"
        Business_Phone:
          $ref: "#/components/schemas/x255CharString"

    xWebUrl:
      type: string
      pattern: ^((ftp|http|https):\/\/)?(www.)?(?!.*(ftp|http|https|www.))[a-zA-Z0-9_-]+(\.[a-zA-Z]+)+((\/)[\w#]+)*(\/\w+\?[a-zA-Z0-9_]+=\w+(&[a-zA-Z0-9_]+=\w+)*)?
      example: "https://www.xtrm.com"
      minLength: 1
      maxLength: 100

    xBusinessEntityInformation:
      type: object
      required:
        - Region_Of_Incorporation
        - Country_Of_Incorporation_Code2
        - Tax_Indentification_Number
        - Address_Line_1
        - City
        - Country_Code2
        - Region_Code2
        - Postal_Code
      properties:
        Legal_Company_Name:
          $ref: "#/components/schemas/xCompanyName"
        FictitiousName:
          $ref: "#/components/schemas/xCompanyName"
        Website:
          $ref: "#/components/schemas/xWebUrl"
        Company_Type:
          $ref: "#/components/schemas/xHundredCharString"
        Industry_Classification:
          description: >
            A general taxonomic identification of the business (Accommodation, utilities, etc.)
          type: string
          example: "Use 'GetNAICS' to get acceptable industry information"
        Date_Of_Registration:
          $ref: "#/components/schemas/xNumericDateString"
        Region_Of_Incorporation:
          $ref: "#/components/schemas/xHundredCharString"
        Ticker_Symbol:
          $ref: "#/components/schemas/xFiftyCharString"
        Tax_Identification_Number:
          $ref: "#/components/schemas/xTaxId"
        Address_Line_1:
          $ref: "#/components/schemas/xAddressLine"
        Address_Line_2:
          $ref: "#/components/schemas/xAddressLine"
        Address_Line_3:
          $ref: "#/components/schemas/xAddressLine"
        City:
          $ref: "#/components/schemas/xCity"
        Country_Code2:
          $ref: "#/components/schemas/xCountryISO2"
        Region_Code2:
          $ref: "#/components/schemas/xRegion_Code2"
        Postal_Code:
          $ref: "#/components/schemas/xPostalCode"
        Number_Of_Employees:
          $ref: "#/components/schemas/xNumericStringSix"
        Expected_Monthly_Payments:
          $ref: "#/components/schemas/xNumericStringTen"
        Expected_Monthly_Volume:
          $ref: "#/components/schemas/xNumericStringTen"
        Purpose_Of_Payments:
          $ref: "#/components/schemas/x255CharString"

    xRegion_Code2:
      type: string
      minLength: 2
      maxLength: 2
      example: "CA"

    x255CharString:
      type: string
      minLength: 1
      maxLength: 255
      example: "A string of at most two hundred and fifty-five (255) characters"

    xNumericStringTen:
      type: string
      minLength: 1
      maxLength: 10
      pattern: ^\d{1,10}$
      example: "850000"

    xNumericStringSix:
      type: string
      minLength: 1
      maxLength: 6
      pattern: ^\d{1,6}$
      example: "850000"

    xName:
      description: >
        A string from 2 to 40 characters in length
      type: string
      minLength: 02
      maxLength: 40
      example: Genevieve

    xEmail:
      description: >
        A valid email address of 100 characters or less
      type: string
      minLength: 000
      maxLength: 100
      pattern: ^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$
      example: genevieveshefa@xtrm.com

    xEmailNotification:
      description: >
        Control some email alerts to beneficiary
      type: boolean
      example: true

    xAmount:
      description: >
        An amount representing some value in a specific
        currency, as a string
      type: string
      minLength: 01
      maxLength: 25
      pattern: ([0-9]*\.?[0-9]*)
      example: '341.25'

    xIAN:
      description: >
        the `AccountID`, `IssuerAccountNumber`, `UserID`
        (and possibly some others) is a unique XTRM
        account identifier, starting with `SPN` for
        companies (usually remitters), or `PAT`
        for individual beneficiaries.
      type: string
      minLength: 01
      maxLength: 20
      pattern: (SPN[1-9][0-9]*|PAT[1-9][0-9]*)
      example: SPN1234567

    xDateString:
      description: >
        A date and time string
      type: string
      example: '2017-10-16T04:10:05.397'

    xPhone:
      description: >
        A phone number, represented as
        a string of 1 to 25 characters
      type: string
      minLength: 01
      maxLength: 25
      example: "14085556245"

    xTaxId:
      description: >
        TaxId: A string such as a US SSN, or some other tax
        jurisdiction’s identifier, of 100 characters or
        less.
      type: string
      minLength: 000
      maxLength: 100
      example: 222-33-4444

    xNumericDay:
      description: >
        numeric day of month '1' to '31'
      type: string
      minLength: 1
      maxLength: 2
      pattern: (12]\d|3[01]|0?[1-9])
      example: '14'

    xNumericMonth:
      description: >
        numeric month of year '11 to '121
      type: string
      minLength: 1
      maxLength: 2
      pattern: ([1][012]|0?[1-9])
      example: '5'

    xNumericYear:
      description: >
        <p>year from 1930 forward through 2999</p>
      type: string
      minLength: 4
      maxLength: 4
      pattern: ^19[3-9]\d|2[0-9]{3}$
      example: '1942'

    xAddressLine:
      description: >
        Street name, number, and other delivery specifics
      type: string
      minLength: 001
      maxLength: 100
      example: 1234 Alhambra Street Way

    xCity:
      description: >
        City: Name of the city, from 1 to 25 characters
      type: string
      minLength: 01
      maxLength: 25
      example: Port Vila
      # Capital of Vanuatu

    xCountry:
      description: >
        Name of the country, from 1 to 255 characters
      type: string
      minLength: 01
      maxLength: 245
      example: Vanuatu

    xCountryISO2:
      description: >
        The two-character ISO country abbreviation
      type: string
      minLength: 02
      maxLength: 02
      pattern: ([a-zA-Z]{2})
      example: VU
      # ISO2 code for Vanuatu

    xCurrencyCode:
      description: >
        The three-letter ISO currency code
      type: string
      minLength: 3
      maxLength: 3
      pattern: ([a-zA-z]{3})
      example: VUV
      # Vanuatu's vatu currency

    xPostalCode:
      description: >
        ZIP code or other code for mail delivery
      type: string
      minLength: 01
      maxLength: 30
      example: 94553-1258

    xRegion:
      description: >
        A jurisdiction such as state, province, or region
      type: string
      minLength: 01
      maxLength: 50
      example: Shefa
      # Administrative Region containing the capital Port Vila

    xAddress:
      type: object
      required:
        - Country
      properties:
        AddressLine1:
          $ref: "#/components/schemas/xAddressLine"
        AddressLine2:
          $ref: "#/components/schemas/xAddressLine"
        AptSuitNum:
          $ref: "#/components/schemas/xAddressLine"
        City:
          $ref: "#/components/schemas/xCity"
        Country:
          $ref: "#/components/schemas/xCountry"
        CountryISO2:
          $ref: "#/components/schemas/xCountryISO2"
        PostalCode:
          $ref: "#/components/schemas/xPostalCode"
        Region:
          $ref: "#/components/schemas/xRegion"
      example:
        AddressLine1: Carpet Investments
        AddressLine2: 3555 Farnam St
        AptSuitNum: "Suite #832-C"
        City: Omaha
        Country: United States
        CountryISO2: US
        PostalCode: '69131'
        Region: Nebraska

    xImageUrl:
      type: object
      additionalProperties:
        type: string
      example:
        additionalProperties: https://www.cdn101.com/znarf/kbld03.png

    xItemDetail:
      type: object
      properties:
        sku:
          description: >
            XTRM ID for product
          type: string
          example: xtrm2322
        rewardName:
          description: >
            Description of reward
          type: string
          example: giftcard
        currencyCode:
          $ref: "#/components/schemas/xCurrencyCode"
        status:
          description: >
            Status of reward or item
          type: string
          example: pending
        valueType:
          description: >
            fixed, variable, or other type of value
          type: string
        rewardType:
          description: >
            Type of reward
          type: string
        maxValue:
          description: >
            Maximum value this reward can carry
          type: number
          example: 500.00
        minValue:
          description: >
            Minimum value this reward can carry
          type: number
          example: 10.00
        faceValue:
          description: >
            Value on the face of the card
          type: number
          example: 50.00
        countries:
          type: array
          items:
            $ref: "#/components/schemas/xCountryISO2"

    xGiftCardDetail:
      type: object
      properties:
        brandName:
          description: >
            Brand or name of company offering the redemption
          type: string
          example: FredCo Chocolates
        description:
          description: >
            A brief description of the nature of the redeemed award, experience, or item
          type: string
          example: Delicious Apricot and Lotus Chocolates
        disclaimer:
          description: >
            Notes or cautions regarding the award
          type: string
          example: may contain sulfites and/or forgotten Hungarian dieties
        terms:
          description: >
            Legal terms and boilerplate
          type: string
          example: legal terms here
        imageUrls:
          type: array
          items:
            $ref: "#/components/schemas/xImageUrl"
        items:
          type: array
          items:
            $ref: "#/components/schemas/xItemDetail"

    xDate:
      type: object
      properties:
        Day:
          $ref: "#/components/schemas/xNumericDay"
        Month:
          $ref: "#/components/schemas/xNumericMonth"
        Year:
          $ref: "#/components/schemas/xNumericYear"

    xOpStatus:
      description: >
        Errors are returned variously by different calls.
        Possibilities include a string, an array of string
        tuples, or the null value.
      type: object
      properties:
        Success:
          description: >
            Success or failure of operation (true / false)
          type: boolean
          example: true
        Errors:
          nullable: true
          type: object
          example: null

    xBeneficiary:
      type: object
      properties:
        AccountNo:
          $ref: "#/components/schemas/xIAN"
        Name:
          $ref: "#/components/schemas/xName"
        State:
          $ref: "#/components/schemas/xRegion"
        Country:
          description: >
            Country
          type: string
          example: Vanuatu
        Employer:
          description: >
            Name or type of employer
          type: string
          example: Vanauatuan Exports Ltd.

    xLinkedBankID:
      description: >
        Identification code for linked bank
      type: string
      minLength: 001
      maxLength: 100
      example: '982353390933219d9'

    xPagination:
      type: object
      description: >
        Record retrieval is paginated.
        `RecordsToSkip` refers to the desired page
        and `RecordsToTake` defines the number of
        records per page. Thus, `RecordsToSkip: '2'`
        and `RecordsToTake: '25'` will display
        the twenty-sixth record through the fiftieth
        record. **Please note that this is a 1-based
        count, and not a 0-based offset.**
      required:
        - RecordsToSkip
        - RecordsToTake
      properties:
        RecordsToSkip:
          description: >
            Page **1-based** number of results
          type: string
          pattern: (^[1-9]\d*$)
          example: '1'
        RecordsToTake:
          description: >
            Number of results per page
          type: string
          pattern: (^[1-9]\d*$)
          example: "25"

    xPaginationTotal:
      type: object
      properties:
        RecordsToSkip:
          description: >
            Page **1-based** number of results
          type: integer
          example: 1
        RecordsToTake:
          description: >
            Number of results per page
          type: integer
          example: 25
        RecordsTotal:
          description: >
            Total number of records to display
          type: integer
          example: 2981

    xPaymentMethodID:
      description: >
        Method of payment, fetched via API
      type: string
      minLength: 01
      maxLength: 25
      example: "Use API"

    xPaymentMethodName:
      description: >
        Name of payment method response
      type: string
      minLength: 001
      maxLength: 100
      example: "Wire"

    xPaymentDetail:
      type: object
      properties:
        PaymentMethodID:
          $ref: "#/components/schemas/xPaymentMethodID"
        PaymentMethodName:
          $ref: "#/components/schemas/xPaymentMethodName"

    xWalletName:
      description: >
        XTRM’s name for the wallet (unique only to that account’s set of wallets)
      type: string
      minLength: 001
      maxLength: 100
      example: "AnyPay VUV Wallet"

    xWalletType:
      type: string
      enum:
        - Standard
        - Accrual
      description: >
        *Standard* wallets are standard, whereas *Accrual* wallets are accrual
      example: Standard

    xWalletRequest:
      type: object
      required:
        - IssuerAccountNumber
        - WalletName
        - WalletCurrency
        - WalletType
      properties:
        IssuerAccountNumber:
          $ref: "#/components/schemas/xIAN"
        WalletName:
          $ref: "#/components/schemas/xWalletName"
        WalletCurrency:
          $ref: "#/components/schemas/xCurrencyCode"
        WalletType:
          $ref: "#/components/schemas/xWalletType"

    xWalletResponse:
      type: object
      properties:
        WalletID:
          description: >
            XTRM identifier for wallet (a globally unique identifier)
          type: integer
          example: 12345
        WalletName:
          $ref: "#/components/schemas/xWalletName"
        WalletCurrency:
          $ref: "#/components/schemas/xCurrencyCode"
        OperationStatus:
          $ref: "#/components/schemas/xOpStatus"

    xProgramDetails:
      type: object
      properties:
        ProgramId:
          $ref: "#/components/schemas/xProgramID"
        ProgramName:
          $ref: "#/components/schemas/xProgramName"
        ProgramCategory:
          description: >
            One of the available program categories
          type: string
          example: Performance
        ProgramType:
          description: >
            Type of the program for the customer
          type: string
          example: Bonus
        IsClaim:
          description: >
            Is this a claimed reward (beneficiary-requested compensation)
          type: string
          example: "N"

    xFullWalletDetails:
      type: object
      properties:
        WalletID:
          $ref: "#/components/schemas/xWalletID"
        WalletName:
          $ref: "#/components/schemas/xWalletName"
        WalletCurrency:
          $ref: "#/components/schemas/xCurrencyCode"
        WalletBalance:
          $ref: "#/components/schemas/xAmount"
        WalletType:
          $ref: "#/components/schemas/xWalletType"

    xProfileStatus:
      description: >
        Status is one of *Submitted*, *Pending*, or *Approved*
      type: string
      enum:
        - Submitted
        - Pending
        - Approved
      example: Submitted

    xNumericDateString:
      description: >
        *mm/dd/yyyy*, from 1930 forward. The month
        and day **must** be two (2) digits, and the year **must**
        be four digits.
      type: string
      minLength: 10
      maxLength: 10
      pattern: (^([1][012]|[1-9])\/(3[01]|[012]\d)\/(19[3-9]\d|20\d{2})$)
      example: "12/31/2020"

    xWallet:
      type: object
      properties:
        WalletID:
          "$ref": "#/components/schemas/xWalletID"
        WalletName:
          $ref: "#/components/schemas/xWalletName"
        WalletCurrency:
          $ref: "#/components/schemas/xCurrencyCode"

    xLimitedWalletDetails:
      type: object
      properties:
        Name:
          description: >
            Named of wallet (unique to account)
          type: string
          example: AnyPay Vatu Wallet
        ID:
          $ref: "#/components/schemas/xWalletID"
        Currency:
          $ref: "#/components/schemas/xCurrencyCode"

    xWalletID:
      description: >
        Name of wallet, unique within an account
      type: string
      minLength: 01
      maxLength: 20
      example: "289112"

    xCompanyName:
      description: >
        Name of company
      type: string
      example: "CarpetCo Investments LLC"

    xProgramName:
      description: >
        Customer-supplied name of program
      type: string
      example: "CarpetCo Jacquard Sales Spiff"

    xUserWalletTransaction:
      type: object
      properties:
        TransactionID:
          description: >
            Globally unique identifier for transaction and account
          type: integer
          example: 79821134
        CompanyName:
          $ref: "#/components/schemas/xCompanyName"
        ProgramName:
          $ref: "#/components/schemas/xProgramName"
        TransactionDate:
          $ref: "#/components/schemas/xDateString"
        description:
          description: >
            Brief description of transaction
          type: string
          example: "Selling jacquard rather than damask carpet"
        Amount:
          $ref: "#/components/schemas/xAmount"
        Type:
          description: >
            Type of transaction
          type: string
          example: "Debit"
        PaymentMethod:
          description: >
            Details of payment method (may be blank)
          nullable: true
          type: string
          example: ""

    xUpdateWalletResult:
      type: object
      properties:
        Status:
          description: >
            Status of update
          type: string
          example: "Updated successfully"
        OperationStatus:
          $ref: "#/components/schemas/xOpStatus"

    xUserPaymentDetails:
      type: object
      properties:
        IssuerAccountNumber:
          $ref: "#/components/schemas/xIAN"
        UserID:
          $ref: "#/components/schemas/xIAN"
        UserWalletID:
          $ref: "#/components/schemas/xWalletID"
        CompanyAccountNumber:
          $ref: "#/components/schemas/xIAN"
        CompanyWalletId:
          $ref: "#/components/schemas/xWalletID"
        Amount:
          $ref: "#/components/schemas/xAmount"
        CurrencyCode:
          $ref: "#/components/schemas/xCurrencyCode"
        description:
          $ref: "#/components/schemas/xDescription"

    xCompanyWalletTransaction:
      type: object
      properties:
        TransactionID:
          $ref: "#/components/schemas/xTransactionID"
        TransactionDate:
          $ref: "#/components/schemas/xDateString"
        Currency:
          $ref: "#/components/schemas/xCurrencyCode"
        ProgramName:
          $ref: "#/components/schemas/xProgramName"
        description:
          description: >
            Nature and explanation of transaction
          type: string
          example: "Currency Exchange"
        Type:
          description: >
            Type of transaction
          type: string
          example: "Debit"
        Amount:
          $ref: "#/components/schemas/xAmount"
        Balance:
          $ref: "#/components/schemas/xAmount"
        PaymentType:
          $ref: "#/components/schemas/xPaymentType"

    xTransactionID:
      description: >
        Globally unique identifier for transaction and account
      type: string
      minLength: 01
      maxLength: 25
      example: "324461"

    xCreditCardType:
      description: >
        Identifier or brand of credit card
      type: string
      minLength: 01
      maxLength: 25
      example: MASTERCARD

    xCreditCardNumber:
      description: >
        Credit card number
      type: string
      minLength: 13
      maxLength: 16
      pattern: (^[0-9]{13,16}$)
      example: "1112333455567778"

    xCVV:
      description: >
        Card Verification Value (**CVV**), a three or four digit number
        serving as a verifier for the credit card’s legitimacy
      type: string
      minLength: 3
      maxLength: 4
      pattern: (^\d{3,4}$)
      example: "456"

    xCreditCardDetails:
      type: object
      required:
        - ExpireMonth
        - ExpireYear
        - CreditCardNumber
        - CreditCardType
        - CVV
      properties:
        ExpireMonth:
          $ref: "#/components/schemas/xNumericMonth"
        ExpireYear:
          $ref: "#/components/schemas/xNumericYear"
        CreditCardNumber:
          $ref: "#/components/schemas/xCreditCardNumber"
        CreditCardType:
          $ref: "#/components/schemas/xCreditCardType"
        CVV:
          $ref: "#/components/schemas/xCVV"

    xSKU:
      description: >
        Identifier for purchase code
      type: string
      minLength: 7
      maxLength: 7
      example: U935268

    xFundUsingCreditCardResult:
      type: object
      properties:
        TransactionID:
          $ref: "#/components/schemas/xTransactionID"
        Amount:
          $ref: "#/components/schemas/xAmount"
        Fee:
          $ref: "#/components/schemas/xAmount"
        TotalAmount:
          $ref: "#/components/schemas/xAmount"
        OperationStatus:
          $ref: "#/components/schemas/xOpStatus"

    xPaymentDetails:
      type: object
      required:
        - IssuerAccountNumber
        - Amount
        - CurrencyCode
        - WalletID
      properties:
        IssuerAccountNumber:
          $ref: "#/components/schemas/xIAN"
        Amount:
          description: >
            Value of transacted currency
          type: number
        CurrencyCode:
          $ref: "#/components/schemas/xCurrencyCode"
        WalletID:
          $ref: "#/components/schemas/xWalletID"

    xPayerInformation:
      type: object
      required:
        - FirstName
        - LastName
      properties:
        FirstName:
          $ref: "#/components/schemas/xName"
        LastName:
          $ref: "#/components/schemas/xName"

    xPayerBillingAddress:
      type: object
      required:
        - Address1
        - City
        - State
        - CountryISO2
        - PostalCode
      properties:
        Address1:
          $ref: "#/components/schemas/xAddress"
        City:
          $ref: "#/components/schemas/xCity"
        State:
          $ref: "#/components/schemas/xRegion"
        CountryISO2:
          $ref: "#/components/schemas/xCountryISO2"
        PostalCode:
          $ref: "#/components/schemas/xPostalCode"

    xExchangeRateMethodsDetailItem:
      type: object
      properties:
        IssuerAccountNumber:
          $ref: "#/components/schemas/xIAN"
        FromCurrency:
          $ref: "#/components/schemas/xCurrencyCode"
        ToCurrency:
          $ref: "#/components/schemas/xCurrencyCode"
        Amount:
          $ref: "#/components/schemas/xAmount"
        ExchangeRate:
          $ref: "#/components/schemas/xCurrencyCode"

    xExchangeRate:
      description: >
        String consisting of the ratio of the starting currency to ending currency (the rate of exchange)
      type: string
      minLength: 01
      maxLength: 25
      pattern: (^[0-9]*\.?[0-9]*$)
      example: ".7312"

    xOTP:
      description: >
        A one-time password string consisting of exactly six digits.
      type: string
      minLength: 6
      maxLength: 6
      pattern: (^\d{6}$)
      example: '843599'

    xWalletTransaction:
      type: object
      properties:
        TransactionID:
          $ref: "#/components/schemas/xTransactionID"
        CompanyName:
          $ref: "#/components/schemas/xCompanyName"
        ProgramName:
          $ref: "#/components/schemas/xProgramName"
        TransactionDate:
          $ref: "#/components/schemas/xDateString"
        description:
          $ref: "#/components/schemas/xDescription"
        Amount:
          $ref: "#/components/schemas/xAmount"
        Currency:
          $ref: "#/components/schemas/xCurrencyCode"
        Type:
          $ref: "#/components/schemas/xPaymentType"
        PaymentMethod:
          description: >
            Method of payment
          type: string

    xTransactionDetails:
      type: object
      description: >
        `UserLinkedBankID` Mandatory if using bank payment

        `UserPayPalEmailID` Mandatory if using PayPal.
        **Support for PayPal transactions has been
        suspended.** Please contact *support@xtrm.com*
        to discuss PayPal connectivity.

        `UserPrepaidVisaEmailID` Mandatory if payment
        method is **Prepaid Virtual Visa**. Minimum amount is
        USD$5<sup style="text-decoration-line: underline">00</sup>,
        maximum amount is USD$1000.00</sup>.
      required:
        - IssuerTransactionId
        - PaymentAmount
        - PartnerAccountNumber
        - RecipientUserID
      properties:
        IssuerTransactionId:
          $ref: "#/components/schemas/xIssuerTransactionId"
        PaymentAmount:
          $ref: "#/components/schemas/xAmount"
        PartnerAccountNumber:
          $ref: "#/components/schemas/xIAN"
        RecipientUserID:
          $ref: "#/components/schemas/xIAN"
        UserLinkedBankID:
          $ref: "#/components/schemas/xLinkedBankID"
        UserPrepaidVisaEmailID:
          $ref: "#/components/schemas/xEmail"
        UserGiftCardEmailID:
          $ref: "#/components/schemas/xEmail"
        SKU:
          $ref: "#/components/schemas/xSKU"
        DealRegId:
          $ref: "#/components/schemas/xDealRegId"
        Comment:
          $ref: "#/components/schemas/xComment"

    xDealRegId:
      description: >
        Person or identifier for the deal registration
      type: string
      minLength: 01
      maxLength: 25
      example: "H. Smith - Canada"

    xComment:
      description: >
        A comment on the transaction, up to 500 characters
      type: string
      minLength: 001
      maxLength: 500
      example: "Long comment up to five hundred (500) characters"

    xIssuerTransactionId:
      description: >
        Issuer-supplied identifier
        that specifies this transaction.
        XTRM returns this identifier unchanged in
        the transaction result. XTRM does not use
        this identifier; it is for the use of the
        issuer.
      type: string
      minLength: 01
      maxLength: 25

    xPaymentType:
      description: >
        Type of payment (Personal, credits, debits ...)
      type: string
      minLength: 01
      maxLength: 50
      enum:
        - Personal
        - ALL
        - Credits
        - Debits
      example: Credits

    xDescription:
      description: >
        A field for explanatory remarks
      type: string
      minLength: 001
      maxLength: 300

    xProgramID:
      description: >
        Identifier for program
      type: string
      minLength: 01
      maxLength: 25
      pattern: (^[1-9]\d{0,24}$)
      example: '2314'

    xTransferFundTransactionDetail:
      type: object
      properties:
        IssuerTransactionId:
          $ref: "#/components/schemas/xIssuerTransactionId"
        PaymentTransactionId:
          $ref: "#/components/schemas/xTransactionID"
        RecipientUserId:
          $ref: "#/components/schemas/xIAN"
        PaymentDate:
          description: >
            String representing a date
          type: string
        PaymentStatus:
          $ref: "#/components/schemas/xPaymentStatus"
        Amount:
          $ref: "#/components/schemas/xAmount"
        Fee:
          $ref: "#/components/schemas/xAmount"
        TotalAmount:
          $ref: "#/components/schemas/xAmount"
        Currency:
          $ref: "#/components/schemas/xCurrencyCode"
        RedemptionDetails:
          $ref: "#/components/schemas/xRedemptionDetails"

    xBankBeneficiaryID:
      description: >
        Bank-dependent ID for beneficiary
      type: string
      minLength: 01
      maxLength: 100
      example: '12e6f8ac19804f90be485045f50ace57'

    xBankDetails:
      type: object
      properties:
        BeneficiaryBankInformation:
          type: object
          properties:
            AccountNumber:
              description: Bank account numer
              type: string
              example: '5498221'
            BankName:
              description: Bank Name
              type: string
              example: 'Axis Bank'
            BranchName:
              description: Bank branch identifier
              type: string
              example: "Madurai"
            BankBeneficiaryStatus:
              nullable: true
              description: >
                **optional**
              type: string
              example: null
            ACHDebitApprovalStatus:
              nullable: true
              description: >
                **optional**
              type: string
              example: null

    xBeneficiaryDetails:
      type: object
      properties:
        BeneficiaryId:
          description: >
            Identity code for beneficiary
          type: string
        BeneficiaryName:
          $ref: "#/components/schemas/xName"
        Currency:
          $ref: "#/components/schemas/xCurrencyCode"
        Country:
          $ref: "#/components/schemas/xCountry"
        TransferFee:
          description: Fee for transfer
          type: number
        PaymentMethods:
          type: object
          additionalProperties:
            type: string
          example:
            Payment: WIRE
            Method: SWIFTBIC
            SomeOtherProperty: SomeOtherValue
            YetAnotherProperty: YetAnotherValue
        BankDetails:
          $ref: "#/components/schemas/xBankDetails"

    xFiftyCharString:
      description: A string of at most fifty (50) characters
      type: string
      minLength: 01
      maxLength: 50
      example: "A string of at most fifty (50) characters"

    xRoutingNumber:
      description: >
        Financial institution routing number
      type: string
      minLength: 01
      maxLength: 15
      pattern: (^\d{1,15}$)
      example: '825331988'

    xAccountNumber:
      description: >
        Institution account identifier
      type: string
      minLength: 01
      maxLength: 50

    xSWIFTBIC:
      description: >
        SWIFT-BIC routing code
      type: string
      minLength: 01
      maxLength: 13

    xBeneficiaryBankInformation:
      type: object
      required:
        - InstitutionName
        - Currency
        - SWIFTBIC
        - AccountNumber
        - RoutingNumber
        - CountryISO2
      properties:
        InstitutionName:
          description: Name of institution
          type: string
          minLength: 01
          maxLength: 50
        Currency:
          $ref: "#/components/schemas/xCurrencyCode"
        SWIFTBIC:
          $ref: "#/components/schemas/xSWIFTBIC"
        AccountNumber:
          $ref: "#/components/schemas/xAccountNumber"
        RoutingNumber:
          $ref: "#/components/schemas/xRoutingNumber"
        CountryISO2:
          $ref: "#/components/schemas/xCountryISO2"
        RemittanceLine3:
          $ref: "#/components/schemas/xFiftyCharString"
        RemittanceLine4:
          $ref: "#/components/schemas/xFiftyCharString"

    xBankBeneficiaryInformation:
      type: object
      required:
        - ContactName
        - PhoneNumber
        - AddressLine1
        - City
        - Region
        - PostalCode
        - CountryISO2
      properties:
        ContactName:
          $ref: "#/components/schemas/xName"
        PhoneNumber:
          $ref: "#/components/schemas/xPhone"
        AddressLine1:
          $ref: "#/components/schemas/xAddressLine"
        AddressLine2:
          $ref: "#/components/schemas/xAddressLine"
        City:
          $ref: "#/components/schemas/xCity"
        Region:
          $ref: "#/components/schemas/xRegion"
        PostalCode:
          $ref: "#/components/schemas/xPostalCode"
        CountryISO2:
          $ref: "#/components/schemas/xCountryISO2"

    xRedemptionDetails:
      nullable: true
      type: object
      properties:
        RedemptionDetails:
          type: object
          properties:
            method:
              description: >
                Payment method
              type: string
              example: "Digital Gift Card"
            TransactionStatus:
              description: >
                Transaction pending, completed, failed
              type: string
              example: "Success"
            TransactionNumber:
              description: >
                ID for transaction
              type: string
              example: "WA210126-28669-40"
            Giftcard:
              description: >
                Name and type of gift card
              type: string
              example: '1-800-PUMPKINS.COM® Gift Card $10.00'
            GiftcardImageUrl:
              description: >
                URL for a displayable image of the card
              type: string
              example: '"https://www.example.com/images/example.png"'
            RecipientEmail:
              $ref: "#/components/schemas/xEmail"
            RecipientName:
              $ref: "#/components/schemas/xName"
            RecipientMobilePhoneNumber:
              $ref: "#/components/schemas/xPhone"
            Amount:
              $ref: "#/components/schemas/xAmount"
            TimeProcessed:
              $ref: "#/components/schemas/xTimestamp"
            Reward:
              $ref: "#/components/schemas/xReward"

    xReward:
      type: object
      properties:
        Credentials:
          $ref: "#/components/schemas/xCredentials"
        CredentialList:
          $ref: "#/components/schemas/xCredentialList"
        RedemptionInstructions:
          description: >
            Redemption instructions for the reward
          type: string
          example: "redemption instructions for the reward"

    xCredentials:
      type: object
      properties:
        ClaimCode:
          description: >
            ID for the claim
          type: string

    xCredentialList:
      type: array
      items:
        type: object
        properties:
          Label:
            description: >
              Name or term for the credential
            type: string
            example: "PIN"
          Value:
            description: >
              Credential value
            type: string
            example: '6363'
          Type:
            description: >
              Kind of data
            type: string
            example: "text"
          CredentialType:
            description: >
              Type of credential
            type: string
            example: "pin"

    xTimestamp:
      description: >
        Time and date information
      type: string
      example: '26 Jan 2021 07:20:04'
      # does not rule out all impossible dates/times
      # such as '39 Feb 1000 29:29:59'
      pattern: ([0123]?\d\s(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Sept|Nov|Dec)\s[12]\d{3}\s[012]?\d:[0-5]\d:[0-5]\d)

    xField:
      type: object
      properties:
        Name:
          description: >
            The name of the field (equivalent to *name*:value)
          type: string
        Value:
          description: >
            The value of the field (equivalent to name:*value*)
          type: string

    xEmpty:
      description: >
        No value or blank object
      type: object
      nullable: true

    xBankBeneficiary:
      #type: object
      #required:
      #  - BeneficiaryDetails
      #properties:
      # BeneficiaryDetails:
      type: object
      required:
        - BeneficiaryInformation
      properties:
        BeneficiaryInformation:
          $ref: "#/components/schemas/xBankBeneficiaryInformation"

    xBasicBeneficiaryDetail:
      type: object
      properties:
        BeneficiaryCompanyName:
          $ref: "#/components/schemas/xCompanyName"
        PartnerAccountManager:
          $ref: "#/components/schemas/xName"
        BeneficiaryID:
          $ref: "#/components/schemas/xIAN"
        BeneficiarySalesforceId:
          description: >
            Salesforce ID for beneficiary
          type: string
          example: 'sf56870921'
        OtherPartnerID:
          description: >
            Partner’s identification code
          type: string
          example: '832b996121cfl99'
        Region:
          $ref: "#/components/schemas/xRegion"

    xAdminDetails:
      type: object
      required:
        - AdminEmail
        - EmailNotification
        - AdminFirstName
        - AdminLastName
        - AdminMobileNumber
        - Country
      properties:
        AdminEmail:
          $ref: "#/components/schemas/xEmail"
        EmailNotification:
          $ref: "#/components/schemas/xEmailNotification"
        AdminFirstName:
          $ref: "#/components/schemas/xName"
        AdminLastName:
          $ref: "#/components/schemas/xName"
        AdminMobileNumber:
          $ref: "#/components/schemas/xPhone"
        City:
          $ref: "#/components/schemas/xCity"
        Country:
          $ref: "#/components/schemas/xCountry"
        Region:
          $ref: "#/components/schemas/xRegion"
        PostalCode:
          $ref: "#/components/schemas/xPostalCode"

    xHundredCharString:
      description: >
        A string from one (1) to one hundred (100) characters
      type: string
      minLength: 001
      maxLength: 100
      example: "A longer string up to one hundred (100) characters"

    xAPIID:
      description: >
        XTRM API User Identifier
      type: string
      example: 9999999_API_User

    xAPISecretKey:
      description: >
        XTRM-supplied secret key
      type: string
      example: dummyDUMMYdummyDUMMYdummyDUMMY/q7LNUdH4ks=

    xBeneficiaryExistDetail:
      type: object
      properties:
        Name:
          $ref: "#/components/schemas/xName"
        AccountNumber:
          $ref: "#/components/schemas/xIAN"
        Email:
          $ref: "#/components/schemas/xEmail"
        MasterAdminFirstName:
          $ref: "#/components/schemas/xName"
        MasterAdminLastName:
          $ref: "#/components/schemas/xName"

    xBeneficiaryAdminDetails:
      type: object
      required:
        - AdminFirstName
        - AdminLastName
        - City
        - Country
        - Region
        - PostalCode
      properties:
        AdminFirstName:
          $ref: "#/components/schemas/xName"
        AdminLastName:
          $ref: "#/components/schemas/xName"
        City:
          $ref: "#/components/schemas/xCity"
        Country:
          $ref: "#/components/schemas/xCountry"
        Region:
          $ref: "#/components/schemas/xRegion"
        PostalCode:
          $ref: "#/components/schemas/xPostalCode"

    xSalesProgramDetails:
      type: object
      required:
        - PartnerAccountManager
        - SalesforcePartnerID
        - OtherPartnerID
        - Region
        - ParnerLevel
      properties:
        PartnerAccountManager:
          description: >
            Name of partner’s account manager
          type: string
        SalesforcePartnerID:
          description: >
            Salesforce identifier for partner
          type: string
        OtherPartnerID:
          description: >
            Other identifying code for partner
          type: string
        Region:
          $ref: "#/components/schemas/xRegion"
        PartnerLevel:
          description: >
            Partner status or sales level
          type: string

    xWebAddress:
      description: >
        URL or other internet resource identifier
      type: string
      minLength: 8
      maxLength: 100
      pattern: ^((?:http(s)?:\/\/)?[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:/?#[\]@!\$&'\(\)\*\+,;=.]+$)$
      example: "https://www.xtrm.com"

    ###########################################################

    SearchBankRequest:
      type: object
      required:
        - SearchBank
      properties:
        SearchBank:
          type: object
          required:
            - Request
          properties:
            Request:
              type: object
              required:
                - IssuerAccountNumber
                - BankName
                - BankCountryISO2
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                BankName:
                  $ref: "#/components/schemas/xHundredCharString"
                BankCountryISO2:
                  $ref: "#/components/schemas/xCountryISO2"
                Pagination:
                  $ref: "#/components/schemas/xPagination"

    SearchBankResponse:
      type: object
      properties:
        SearchBankResponse:
          type: object
          properties:
            SearchBankResult:
              type: object
              properties:
                Banks:
                  type: array
                  items:
                    type: object
                    properties:
                      BankName:
                        $ref: "#/components/schemas/xHundredCharString"
                  example:
                    - BankName: "Wells Fargo Bank New Mexico Na"
                    - BankName: "Wells Fargo Bank, National Association"
                    - BankName: "Wells Fargo Clearing Services, LLC"
                    - BankName: "Wells Fargo Na"
                    - BankName: "Wells Fargo Securities International Limited"
                    - BankName: "Wells Fargo Securities, LLC"
                    - BankName: "Wells Fargo Trust Company, National Association"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"
                PaginationTotal:
                  $ref: "#/components/schemas/xPaginationTotal"

    CreateUser:
      type: object
      required:
        - CreateUser
      properties:
        CreateUser:
          type: object
          required:
            - request
          properties:
            request:
              $ref: "#/components/schemas/CreateUserRequest"

    CreateUserRequest:
      type: object
      required:
        - IssuerAccountNumber
        - LegalFirstName
        - LegalLastName
        - EmailAddress
        - EmailNotification

      properties:
        IssuerAccountNumber:
          $ref: "#/components/schemas/xIAN"

        LegalFirstName:
          $ref: "#/components/schemas/xName"

        LegalLastName:
          $ref: "#/components/schemas/xName"

        EmailAddress:
          $ref: "#/components/schemas/xEmail"

        EmailNotification:
          $ref: "#/components/schemas/xEmailNotification"

        MobilePhone:
          $ref: "#/components/schemas/xPhone"

        TaxId:
          $ref: "#/components/schemas/xTaxId"

        DateOfBirth:
          $ref: "#/components/schemas/xDate"

        Address:
          $ref: "#/components/schemas/xAddress"

    CreateUserResponse:
      type: object
      properties:
        CreateUserResponse:
          type: object
          properties:
            CreateUserResult:
              $ref: "#/components/schemas/UserResult"

    UserResult:
      type: object
      properties:
        UserID:
          $ref: "#/components/schemas/xIAN"
        AccountIdentityLevel:
          description: >
            Known, partially known, etc.
          type: string
          example: "Pending"
        OperationStatus:
          $ref: "#/components/schemas/xOpStatus"

    UpdateUserRequest:
      type: object
      required:
        - UpdateUser
      properties:
        UpdateUser:
          type: object
          required:
            - request
          properties:
            request:
              type: object
              required:
                - IssuerAccountNumber
                - UserId
                - LegalFirstName
                - LegalLastName
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"

                UserId:
                  $ref: "#/components/schemas/xIAN"

                LegalFirstName:
                  $ref: "#/components/schemas/xName"

                LegalLastName:
                  $ref: "#/components/schemas/xName"

                # EmailAddress:
                #  $ref: "#/components/schemas/xEmail"

                # EmailNotification:
                #  type: boolean

                # MobilePhone:
                #  $ref: "#/components/schemas/xPhone"

                TaxId:
                  $ref: "#/components/schemas/xTaxId"

                DateOfBirth:
                  $ref: "#/components/schemas/xDate"

                Address:
                  $ref: "#/components/schemas/xAddress"

    UpdateUserResponse:
      type: object
      properties:
        UpdateUserResponse:
          type: object
          properties:
            UpdateUserResult:
              $ref: "#/components/schemas/UserResult"

    CheckUserExistRequest:
      type: object
      properties:
        CheckUserExist:
          type: object
          properties:
            request:
              type: object
              required:
                - IssuerAccountNumber
                - Email
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"

                Email:
                  $ref: "#/components/schemas/xEmail"

    CheckUserExistResponse:
      type: object
      properties:
        CheckUserExistResponse:
          type: object
          properties:
            CheckUserExistResult:
              $ref: "#/components/schemas/CheckUserExistResult"

    CheckUserExistResult:
      type: object
      properties:
        Beneficiary:
          type: array
          items:
            $ref: "#/components/schemas/xBeneficiary"

        OperationStatus:
          $ref: "#/components/schemas/xOpStatus"

    GetConnectedCompanyFundRequestRequest:
      type: object
      required:
        - GetConnectedFundRequestDetails
      properties:
        GetConnectedCompanyFundRequestDetails:
          type: object
          required:
            - IssuerAccountNumber
            - BeneficiaryAccountNumber
            - PaymentType
            - FromDate
            - ToDate
            - Pagination
          properties:
            IssuerAccountNumber:
              $ref: "#/components/schemas/xIAN"
            BeneficiaryAccountNumber:
              $ref: "#/components/schemas/xIAN"
            PaymentType:
              $ref: "#/components/schemas/xPaymentType"
            FromDate:
              $ref: "#/components/schemas/xNumericDateString"
            ToDate:
              $ref: "#/components/schemas/xNumericDateString"
            Pagination:
              $ref: "#/components/schemas/xPagination"

    GetConnectedCompanyFundRequestResponse:
      type: string
      example: "GetConnectedCompanyFundRequestResponse is not documented."

    GetPersonalBeneficiariesRequest:
      type: object
      required:
        - GetPersonalBeneficiaries
      properties:
        GetPersonalBeneficiaries:
          type: object
          required:
            - request
          properties:
            request:
              $ref: "#/components/schemas/GetPersonalBeneficiariesRequestRequest"

    GetPersonalBeneficiariesRequestRequest:
      type: object
      required:
        - IssuerAccountNumber
        - Pagination
      properties:
        IssuerAccountNumber:
          $ref: "#/components/schemas/xIAN"
        Pagination:
          $ref: "#/components/schemas/xPagination"

    GetPersonalBeneficiariesResponse:
      type: object
      properties:
        GetPersonalBeneficiariesReponse:
          type: object
          properties:
            GetPersonalBeneficiariesResult:
              type: object
              properties:
                Beneficiary:
                  type: array
                  items:
                    $ref: "#/components/schemas/xBeneficiary"
            OperationStatus:
              $ref: "#/components/schemas/xOpStatus"
            PaginationTotal:
              $ref: "#/components/schemas/xPaginationTotal"

      # GetPaymentMethodsRequest:
      #  this is just {}, so no schema for it

    GetPaymentMethodsResponse:
      type: object
      properties:
        GetPaymentMethodsResponse:
          type: object
          properties:
            PaymentMethodResult:
              type: object
              properties:
                PaymentMethods:
                  type: object
                  properties:
                    PaymentMethodDetails:
                      type: array
                      items:
                        $ref: "#/components/schemas/xPaymentDetail"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    GetBeneficiaryCompanyPaymentsResponse:
      type: object
      properties:
        GetBeneficiaryCompanyPaymentsResponse:
          type: object
          properties:
            BeneficiaryCompanyPaymentMethodResult:
              type: object
              properties:
                BeneficiaryCompanyPaymentMethods:
                  type: object
                  properties:
                    BeneficiaryCompanyPaymentMethodDetails:
                      type: array
                      items:
                        type: object
                        properties:
                          BeneficiaryCompanyPaymentMethodID:
                            $ref: "#/components/schemas/xPaymentMethodID"
                          BeneficiaryCompanyPaymentMethodName:
                            $ref: "#/components/schemas/xPaymentMethodName"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    GetUserWalletBalanceRequest:
      type: object
      properties:
        GetUserWalletBalance:
          type: object
          properties:
            request:
              type: object
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                UserId:
                  $ref: "#/components/schemas/xIAN"
                Currency:
                  $ref: "#/components/schemas/xCurrencyCode"

    UserWalletResult:
      type: object
      properties:
        Balance:
          type: number
          example: 1398.62
        Currency:
          $ref: "#/components/schemas/xCurrencyCode"
        OperationStatus:
          $ref: "#/components/schemas/xOpStatus"

    UserWalletBalanceResponseDeprecated:
      deprecated: true
      type: object
      properties:
        UserWalletBalnceResponse:
          $ref: "#/components/schemas/UserWalletResult"

    UserWalletBalanceResponse:
      type: object
      properties:
        UserWalletBalanceResponse:
          $ref: "#/components/schemas/UserWalletResult"

    GetCompanyWalletsRequest:
      type: object
      properties:
        GetCompanyWallets:
          type: object
          properties:
            request:
              type: object
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"

    GetCompanyWalletsResponse:
      type: object
      properties:
        GetCompanyWalletResponse:
          type: object
          properties:
            CompanyWalletResult:
              type: object
              properties:
                Company Wallets:
                  type: object
                  properties:
                    CompanyWalletDetails:
                      type: array
                      items:
                        $ref: "#/components/schemas/xFullWalletDetails"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    GetBeneficiaryWalletsRequest:
      type: object
      properties:
        GetBeneficiaryWallets:
          type: object
          properties:
            request:
              type: object
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                BeneficiaryAccountNumber:
                  $ref: "#/components/schemas/xIAN"

    GetBeneficiaryWalletsResponse:
      type: object
      properties:
        GetBeneficiaryWalletResponse:
          type: object
          properties:
            GetBeneficiaryWalletsResult:
              type: object
              properties:
                Wallets:
                  type: array
                  items:
                    $ref: "#/components/schemas/xLimitedWalletDetails"

    CreateUserWalletRequest:
      type: object
      properties:
        CreateUserWallet:
          type: object
          properties:
            request:
              type: object
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                UserID:
                  $ref: "#/components/schemas/xIAN"
                WalletName:
                  $ref: "#/components/schemas/xWalletName"
                WalletCurrency:
                  $ref: "#/components/schemas/xCurrencyCode"

    CreateUserWalletResponse:
      type: object
      properties:
        CreateUserWalletResponse:
          type: object
          properties:
            CreateUserWalletResult:
              $ref: "#/components/schemas/xWalletResponse"

    CreateCompanyWalletRequest:
      type: object
      properties:
        CreateCompanyWallet:
          type: object
          properties:
            request:
              type: object
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                WalletName:
                  $ref: "#/components/schemas/xWalletName"
                WalletCurrency:
                  $ref: "#/components/schemas/xCurrencyCode"
                WalletType:
                  $ref: "#/components/schemas/xWalletType"
                AllowAccessAccountNumber:
                  $ref: "#/components/schemas/xIAN"

    CreateCompanyWalletResponse:
      type: object
      properties:
        CreateCompanyWalletResponse:
          type: object
          properties:
            WalletID:
              $ref: "#/components/schemas/xWalletID"
            WalletName:
              $ref: "#/components/schemas/xWalletName"
            WalletCurrency:
              $ref: "#/components/schemas/xCurrencyCode"
            OperationStatus:
              $ref: "#/components/schemas/xOpStatus"

    CreateBeneficiaryCompanyWalletRequest:
      type: object
      required:
        - CreateBeneficiaryCompanyWallet
      properties:
        CreateBeneficiaryCompanyWallet:
          type: object
          required:
            - request
          properties:
            request:
              $ref: "#/components/schemas/xWalletRequest"

    CreateBeneficiaryCompanyWalletResponse:
      type: object
      properties:
        CreateBeneficiaryCompanyWalletResult:
          $ref: "#/components/schemas/xWalletResponse"

    UpdateUserWalletRequest:
      type: object
      required:
        - UpdateUserWallet
      properties:
        UpdateUserWallet:
          type: object
          required:
            - Request
          properties:
            Request:
              type: object
              required:
                - IssuerAccountNumber
                - UserID
                - WalletID
                - WalletName
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                UserID:
                  $ref: "#/components/schemas/xIAN"
                WalletID:
                  $ref: "#/components/schemas/xWalletID"
                WalletName:
                  $ref: "#/components/schemas/xWalletName"

    UpdateUserWalletResponse:
      type: object
      properties:
        UpdateUserWalletResponse:
          type: object
          properties:
            UpdateUserWalletResult:
              $ref: "#/components/schemas/xUpdateWalletResult"

    UpdateCompanyWalletRequest:
      type: object
      required:
        - UpdateCompanyWallet
      properties:
        UpdateCompanyWallet:
          type: object
          required:
            - request
          properties:
            request:
              type: object
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                WalletID:
                  $ref: "#/components/schemas/xWalletID"
                WalletName:
                  $ref: "#/components/schemas/xWalletName"

    UpdateCompanyWalletResponse:
      type: object
      properties:
        UpdateCompanyWalletResponse:
          type: object
          properties:
            UpdateCompanyWalletResult:
              $ref: "#/components/schemas/xUpdateWalletResult"

    UpdateBeneficiaryCompanyWalletRequest:
      type: object
      required:
        - UpdateBeneficiaryCompanyWallet
      properties:
        UpdateBeneficiaryCompanyWallet:
          type: object
          required:
            - Request
          properties:
            Request:
              type: object
              required:
                - IssuerAccountNumber
                - BeneficiaryAccountNumber
                - WalletID
                - WalletName
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                BeneficiaryAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                WalletID:
                  $ref: "#/components/schemas/xWalletID"
                WalletName:
                  $ref: "#/components/schemas/xWalletName"

    UpdateBeneficiaryCompanyWalletResponse:
      type: object
      properties:
        UpdateBeneficiaryCompnayWalletResponse:
          type: object
          properties:
            UpdateBeneficiaryCompanyWalletResult:
              $ref: "#/components/schemas/xUpdateWalletResult"

    GetUserWalletTransactionsRequest:
      type: object
      required:
        - GetUserWalletTransactions
      properties:
        GetUserWalletTransactions:
          type: object
          required:
            - Request
          properties:
            Request:
              type: object
              required:
                - IssuerAccountNumber
                - UserID
                - WalletCurrency
                - Pagination
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                UserID:
                  $ref: "#/components/schemas/xIAN"
                WalletCurrency:
                  $ref: "#/components/schemas/xCurrencyCode"
                Pagination:
                  $ref: "#/components/schemas/xPagination"

    GetUserWalletTransactionsResponse:
      type: object
      required:
        - GetUserWalletTransactionsResponse
      properties:
        GetUserWalletTransactionsResponse:
          type: object
          properties:
            Transactions:
              type: array
              items:
                $ref: "#/components/schemas/xUserWalletTransaction"
            OperationStatus:
              $ref: "#/components/schemas/xOpStatus"
            PaginationTotal:
              $ref: "#/components/schemas/xPaginationTotal"

    GetCompanyWalletTransactionsRequest:
      type: object
      required:
        - GetCompanyWalletTransactions
      properties:
        GetCompanyWalletTransactions:
          type: object
          required:
            - Request
          properties:
            Request:
              type: object
              required:
                - IssuerAccountNumber
                - WalletID
                - Pagination
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                WalletID:
                  $ref: "#/components/schemas/xWalletID"
                Pagination:
                  $ref: "#/components/schemas/xPagination"

    GetCompanyWalletTransactionsResponse:
      type: object
      properties:
        GetCompanyWalletTransactionsResponse:
          type: object
          properties:
            GetCompanyWalletTransactionsResult:
              type: object
              properties:
                Transactions:
                  type: array
                  items:
                    $ref: "#/components/schemas/xCompanyWalletTransaction"

    GetCompanyWalletTransactionDetailsRequest:
      type: object
      required:
        - GetCompanyTransactionDetails
      properties:
        GetCompanyTransactionDetails:
          type: object
          required:
            - Request
          properties:
            Request:
              type: object
              required:
                - IssuerAccountNumber
                - TransactionID
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                TransactionID:
                  $ref: "#/components/schemas/xTransactionID"

    GetCompanyWalletTransactionDetailsResponse:
      type: object
      properties:
        GetCompanyTransactionDetailsResponse:
          type: object
          properties:
            GetCompanyTransactionDetailsResult:
              type: object
              properties:
                Field:
                  type: array
                  items:
                    $ref: '#/components/schemas/xField'

            OperationStatus:
              $ref: "#/components/schemas/xOpStatus"
          example:
            GetCompanyTransactionDetailsResponse:
              GetCompanyTransactionDetailsResult:
                - Field:
                    Name: Method
                    Value: CITI
                - Field:
                    Name: Status
                    Value: Transaction Pending
                - Field:
                    Name: Transferred to
                    Value: Roger Penske (xtrm2009@geemail.com)
                - Field:
                    Name: Program
                    Value: test
                - Field:
                    Name: Description
                    Value: test
                - Field:
                    Name: Comment
                    Value: "-"
                - Field:
                    Name: Submission File
                    Value: "0"
                - Field:
                    Name: Partner
                    Value: "-"
                - Field:
                    Name: Sale Date
                    Value: "Oct 09, 2015"
                - Field:
                    Name: Total Payment
                    Value: USD 23.00
                - Field:
                    Name: Budget Code
                    Value: "-"
            OperationStatus:
              Success: true
              Errors: null

    FundCompanyWalletUsingCreditCardRequest:
      type: object
      required:
        - FundCompanyWalletUsingCreditCardRequest
      properties:
        FundCompanyWalletUsingCreditCardRequest:
          type: object
          required:
            - Request
          properties:
            Request:
              type: object
              required:
                - PaymentDetails
                - PayerInformation
                - PayerBillingAddress
                - CreditCardDetails

              properties:
                PaymentDetails:
                  $ref: "#/components/schemas/xPaymentDetails"

                PayerInformation:
                  $ref: "#/components/schemas/xPayerInformation"

                PayerBillingAddress:
                  $ref: "#/components/schemas/xPayerBillingAddress"

                CreditCardDetails:
                  $ref: "#/components/schemas/xCreditCardDetails"

    FundCompanyWalletUsingCreditCardResponse:
      type: object
      required:
        - FundCompanyWalletUsingCreditCardResponse
      properties:
        FundCompanyWalletUsingCreditCardResponse:
          type: object
          required:
            - FundCompanyWalletUsingCreditCardResult
          properties:
            FundCompanyWalletUsingCreditCardResult:
              $ref: "#/components/schemas/xFundUsingCreditCardResult"

    FundUserWalletsUsingCreditCardRequest:
      type: object
      required:
        - FundUserWalletUsingCreditCardRequest
      properties:
        FundUserWalletUsingCreditCardRequest:
          type: object
          required:
            - Request
          properties:
            Request:
              type: object
              required:
                - PaymentDetails
                - PayerInformation
                - PayerBillingAddress
                - CreditCardDetails
              properties:
                PaymentDetails:
                  $ref: "#/components/schemas/xPaymentDetails"

                PayerInformation:
                  $ref: "#/components/schemas/xPayerInformation"

                PayerBillingAddress:
                  $ref: "#/components/schemas/xPayerBillingAddress"

                CreditCardDetails:
                  $ref: "#/components/schemas/xCreditCardDetails"

    FundUserWalletsUsingCreditCardResponse:
      type: object
      properties:
        FundUserWalletUsingCreditCardResponse:
          type: object
          properties:
            FundUserWalletUsingCreditCardResult:
              $ref: "#/components/schemas/xFundUsingCreditCardResult"

    FundWalletUsingACHDebitRequest:
      type: object
      required:
        - FundWalletUsingACHDebitRequest
      properties:
        FundWalletUsingACHDebitRequest:
          type: object
          required:
            - Request
          properties:
            Request:
              type: object
              required:
                - IssuerAccountNumber
                - Amount
                - CurrencyCode
                - WalletID
                - LinkedBankID
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                Amount:
                  $ref: "#/components/schemas/xAmount"
                CurrencyCode:
                  $ref: "#/components/schemas/xCurrencyCode"
                WalletID:
                  $ref: "#/components/schemas/xWalletID"
                LinkedBankID:
                  $ref: "#/components/schemas/xLinkedBankID"

    FundWalletUsingACHDebitResponse:
      type: object
      properties:
        FundWalletUsingACHDebitResponse:
          type: object
          properties:
            FundWalletUsingACHDebitResult:
              type: object
              properties:
                TransactionID:
                  $ref: "#/components/schemas/xTransactionID"
                Amount:
                  $ref: "#/components/schemas/xAmount"
                Fee:
                  $ref: "#/components/schemas/xAmount"
                TotalAmount:
                  $ref: "#/components/schemas/xAmount"
                Currency:
                  $ref: "#/components/schemas/xCurrencyCode"
                Status:
                  description: >
                    This payment is *pending* until the ACH
                    transaction has been completed.
                  type: string
                  example: "ACH debit fund request in process (Pending). You will be notified once completed."
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    GetUserWalletsRequest:
      type: object
      required:
        - GetUserWallets
      properties:
        GetUserWallets:
          type: object
          required:
            - Request
          properties:
            Request:
              type: object
              required:
                - IssuerAccountNumber
                - UserID
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                UserID:
                  $ref: "#/components/schemas/xIAN"

    GetUserWalletsResponse:
      type: object
      properties:
        GetUserWalletsResponse:
          type: object
          properties:
            GetUserWalletsResult:
              type: object
              properties:
                Wallets:
                  type: array
                  items:
                    $ref: "#/components/schemas/xWallet"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    GetExchangeRateRequest:
      type: object
      required:
        - GetExchangeRateRequest
      properties:
        GetExchangeRateRequest:
          type: object
          required:
            - Request
          properties:
            Request:
              type: object
              required:
                - IssuerAccountNumber
                - FromCurrency
                - ToCurrency
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                FromCurrency:
                  $ref: "#/components/schemas/xCurrencyCode"
                ToCurrency:
                  $ref: "#/components/schemas/xCurrencyCode"

    GetExchangeRateResponse:
      type: object
      properties:
        GetExchangeRateMethodsResponse:
          type: object
          properties:
            ExchangeRateMethodResult:
              type: object
              properties:
                ExchangeRateMethods:
                  type: object
                  properties:
                    ExchangeRateMethodsDetail:
                      type: array
                      items:
                        $ref: "#/components/schemas/xExchangeRateMethodsDetailItem"

    BookExchangeRequest:
      type: object
      required:
        - BookExchange
      properties:
        BookExchange:
          type: object
          required:
            - Request
          properties:
            Request:
              type: object
              required:
                - IssuerAccountNumber
                - AccountNumber
                - FromWalletID
                - ToWalletID
                - FromCurrency
                - ToCurrency
                - OTP
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                AccountNumber:
                  $ref: "#/components/schemas/xIAN"
                FromWalletID:
                  $ref: "#/components/schemas/xWalletID"
                ToWalletID:
                  $ref: "#/components/schemas/xWalletID"
                FromCurrency:
                  $ref: "#/components/schemas/xCurrencyCode"
                ToCurrency:
                  $ref: "#/components/schemas/xCurrencyCode"
                OTP:
                  $ref: "#/components/schemas/xOTP"

    BookExchangeResponse:
      type: object
      properties:
        BookExchangeMethodsResponse:
          type: object
          properties:
            BookExchangeMethodResult:
              type: object
              properties:
                BookExchangeMethodDetail:
                  type: object
                  properties:
                    DebitTransactionID:
                      $ref: "#/components/schemas/xTransactionID"
                    CreditTransactionID:
                      $ref: "#/components/schemas/xTransactionID"
                    PaymentDate:
                      $ref: "#/components/schemas/xDateString"
                    PaymentStatus:
                      $ref: "#/components/schemas/xAmount"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    GetUserWalletTransactionsByRemitterRequest:
      type: object
      required:
        - GetUserWalletTransactionsByRemitter
      properties:
        GetUserWalletTransactionsByRemitter:
          type: object
          required:
            - Request
          properties:
            Request:
              type: object
              required:
                - IssuerAccountNumber
                - UserID
                - RemitterAccountID
                - WalletCurrency
                - Pagination
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                UserID:
                  $ref: "#/components/schemas/xIAN"
                RemitterAccountID:
                  $ref: "#/components/schemas/xIAN"
                WalletCurrency:
                  $ref: "#/components/schemas/xCurrencyCode"
                Pagination:
                  $ref: "#/components/schemas/xPagination"

    GetUserWalletTransactionsByRemitterResponse:
      type: object
      properties:
        GetUserWalletTransactionsByRemitterResponse:
          type: object
          properties:
            GetUserWalletTransactionsByRemitterResult:
              type: object
              properties:
                Transactions:
                  type: array
                  items:
                    $ref: "#/components/schemas/xWalletTransaction"
                PaginationTotal:
                  $ref: "#/components/schemas/xPaginationTotal"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    TransferFundRequest:
      type: object
      required:
        - TransferFund
      properties:
        TransferFund:
          type: object
          required:
            - request
          properties:
            request:
              type: object
              required:
                - IssuerAccountNumber
                - PaymentType
                - PaymentMethodID
                - ProgramID
                - WalletID
                - PaymentDescription
                - PaymentCurrency
                - EmailNotification
                - TransactionDetails
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                PaymentType:
                  $ref: "#/components/schemas/xPaymentType"
                PaymentMethodID:
                  $ref: "#/components/schemas/xPaymentMethodID"
                ProgramID:
                  $ref: "#/components/schemas/xProgramID"
                WalletID:
                  $ref: "#/components/schemas/xWalletID"
                PaymentDescription:
                  $ref: "#/components/schemas/xDescription"
                PaymentCurrency:
                  $ref: "#/components/schemas/xCurrencyCode"
                EmailNotification:
                  $ref: "#/components/schemas/xEmailNotification"
                TransactionDetails:
                  type: array
                  items:
                    $ref: "#/components/schemas/xTransactionDetails"

    TransferFundResponse:
      type: object
      properties:
        TransferFundResponse:
          type: object
          properties:
            TransferFundResult:
              type: object
              properties:
                TransactionDetail:
                  type: array
                  items:
                    $ref: "#/components/schemas/xTransferFundTransactionDetail"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    UserWithdrawFundRequest:
      type: object
      required:
        - UserWithdrawFund
      properties:
        UserWithdrawFund:
          description: >
            `UserLinkedBandID` is mandatory for bank payments

            `UserPayPalEmailID` is mandatory for paypal transactions.
            **Paypal transactions have been suspended. Please contact
            support@xtrm.com for any questions regarding Paypal.**

            `UserPrepaidVisaEmailID` is mandatory for payment with a
            prepaid virtual cash card.

            `UserGiftCardEmailID` is mandatory for payment with a
            digital gift card.

          type: object
          required:
            - request
          properties:
            request:
              type: object
              required:
                - IssuerAccountNumber
                - UserID
                - Amount
                - Currency
                - PaymentMethodID
                - OTP

              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                UserID:
                  $ref: "#/components/schemas/xIAN"
                Amount:
                  $ref: "#/components/schemas/xAmount"
                Currency:
                  $ref: "#/components/schemas/xCurrencyCode"
                PaymentMethodID:
                  $ref: "#/components/schemas/xPaymentMethodID"
                UserLinkedBankID:
                  $ref: "#/components/schemas/xLinkedBankID"
                UserPayPalEmailID:
                  $ref: "#/components/schemas/xEmail"
                UserPrepaidVisaEmailID:
                  $ref: "#/components/schemas/xEmail"
                SKU:
                  $ref: "#/components/schemas/xSKU"
                UserGiftCardEmailID:
                  $ref: "#/components/schemas/xEmail"
                OTP:
                  $ref: "#/components/schemas/xOTP"
                SendTransferCodetoEmail:
                  description: >
                    True or false code
                  type: boolean
                SendTransferCodetoMobile:
                  description: >
                    True or false code
                  type: boolean

    UserWithdrawFundResponse:
      type: object
      properties:
        UserWithdrawFundResponse:
          type: object
          properties:
            PaymentDate:
              $ref: "#/components/schemas/xDateString"
            PaymentStatus:
              $ref: "#/components/schemas/xAmount"
            OperationStatus:
              $ref: "#/components/schemas/xOpStatus"

    TransferFundToCompanyRequest:
      type: object
      required:
        - TransferFundToCompany
      properties:
        TransferFundToCompany:
          type: object
          required:
            - request
          properties:
            request:
              description: >
                `BeneficiaryLinkedBankID` is required only if the payment method is via bank

                `BeneficiaryPayPalEmailID` is required only for Paypal payments.
                **Paypal payments have been suspended. Please contact support@xtrm.com
                for Paypal connectivity requirements.**

              type: object
              required:
                - IssuerAccountNumber
                - PaymentType
                - PaymentMethodID
                - WalletId
                - ProgramId
                - Description
                - Currency
                - EmailNotification
                - IssuerTransactionId
                - Amount
                - BeneficiaryAccountNumber
                - BeneficiaryWalletID

              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                PaymentType:
                  $ref: "#/components/schemas/xPaymentType"
                PaymentMethodID:
                  $ref: "#/components/schemas/xPaymentMethodID"
                ProgramId:
                  $ref: "#/components/schemas/xProgramID"
                WalletId:
                  $ref: "#/components/schemas/xWalletID"
                description:
                  $ref: "#/components/schemas/xDescription"
                Currency:
                  $ref: "#/components/schemas/xCurrencyCode"
                Amount:
                  $ref: "#/components/schemas/xAmount"
                EmailNotification:
                  $ref: "#/components/schemas/xEmailNotification"
                IssuerTransactionId:
                  $ref: "#/components/schemas/xIssuerTransactionId"
                BeneficiaryAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                BeneficiaryWalletID:
                  $ref: "#/components/schemas/xWalletID"
                BeneficiaryLinkedBankId:
                  $ref: "#/components/schemas/xLinkedBankID"
                BeneficiaryPayPayEmailId:
                  $ref: "#/components/schemas/xEmail"

    TransferFundToCompanyResponse:
      type: object
      properties:
        TransferFundToCompanyResponse:
          type: object
          properties:
            TransferFundToCompanyResult:
              type: object
              properties:
                IssuerTransactionId:
                  $ref: "#/components/schemas/xIssuerTransactionId"
                PaymentTransactionId:
                  $ref: "#/components/schemas/xTransactionID"
                BeneficiaryAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                PaymentDate:
                  $ref: "#/components/schemas/xDateString"
                PaymentStatus:
                  $ref: "#/components/schemas/xAmount"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    TransferFundUsertoCompanyUsingCCRequest:
      type: object
      required:
        - TransferFundUserToCompanyUsingCC
      properties:
        TransferFundUserToCompanyUsingCC:
          type: object
          required:
            - Request
          properties:
            Request:
              type: object
              required:
                - PaymentDetails
                - PayerInformation
                - PayerBillingAddress
                - CreditCardDetails
              properties:
                PaymentDetails:
                  $ref: "#/components/schemas/xUserPaymentDetails"
                PayerInformation:
                  $ref: "#/components/schemas/xPayerInformation"
                PayerBillingAddress:
                  $ref: "#/components/schemas/xPayerBillingAddress"
                CreditCardDetails:
                  $ref: "#/components/schemas/xCreditCardDetails"

    TransferFundUsertoCompanyUsingCCResponse:
      type: object
      properties:
        TransferFundUsertoCompanyUsingCCResponse:
          type: object
          properties:
            TransferFundUsertoCompanyUsingCCResult:
              type: object
              properties:
                TransactionID:
                  $ref: "#/components/schemas/xTransactionID"
                Amount:
                  $ref: "#/components/schemas/xAmount"
                Fee:
                  $ref: "#/components/schemas/xAmount"
                TotalAmount:
                  $ref: "#/components/schemas/xAmount"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    TransferFundWalletToWalletRequest:
      type: object
      required:
        - TransferFundWalletToWallet
      properties:
        TransferFundWalletToWallet:
          type: object
          required:
            - Request
          properties:
            Request:
              type: object
              required:
                - IssuerAccountNumber
                - FromAccountNumber
                - FromWalletID
                - ToAccountNumber
                - ToWalletID
                - Currency
                - Amount
                - Description
                - OTP
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                FromAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                FromWalletID:
                  $ref: "#/components/schemas/xWalletID"
                ToAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                ToWalletID:
                  $ref: "#/components/schemas/xWalletID"
                Currency:
                  $ref: "#/components/schemas/xCurrencyCode"
                Amount:
                  $ref: "#/components/schemas/xAmount"
                description:
                  $ref: "#/components/schemas/xDescription"
                OTP:
                  $ref: "#/components/schemas/xOTP"

    TransferFundWalletToWalletResponse:
      type: object
      properties:
        TransferFundWallettoWalletResponse:
          type: object
          properties:
            TransferFundWallettoWalletResult:
              type: object
              properties:
                TransactionID:
                  $ref: "#/components/schemas/xTransactionID"
                Amount:
                  $ref: "#/components/schemas/xAmount"
                Fee:
                  $ref: "#/components/schemas/xAmount"
                TotalAmount:
                  $ref: "#/components/schemas/xAmount"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    TransferFundDynamicAccountCreateUserRequest:
      type: object
      required:
        - TransferFundToDynamicAccountUser
      properties:
        TransferFundToDynamicAccountUser:
          type: object
          required:
            - Request
          properties:
            Request:
              type: object
              required:
                - IssuerAccountNumber
                - FromAccountNumber
                - FromWalletID
                - RecipientEmail
                - Currency
                - Amount
                - Description
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                FromAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                FromWalletID:
                  $ref: "#/components/schemas/xWalletID"
                RecipientFirstName:
                  $ref: "#/components/schemas/xName"
                RecipientLastName:
                  $ref: "#/components/schemas/xIAN"
                RecipientEmail:
                  $ref: "#/components/schemas/xEmail"
                Currency:
                  $ref: "#/components/schemas/xCurrencyCode"
                Amount:
                  $ref: "#/components/schemas/xAmount"
                description:
                  $ref: "#/components/schemas/xDescription"

    TransferFundDynamicAccountCreateUserResponse:
      type: object
      properties:
        TransferFundToDynamicAccountUserResponse:
          type: object
          properties:
            TransferFundToDynamicAccountUserResult:
              type: object
              properties:
                TransactionID:
                  $ref: "#/components/schemas/xTransactionID"
                Amount:
                  $ref: "#/components/schemas/xAmount"
                Fee:
                  $ref: "#/components/schemas/xAmount"
                TotalAmount:
                  $ref: "#/components/schemas/xAmount"
                Currency:
                  $ref: "#/components/schemas/xCurrencyCode"
                RecipientAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    GetDigitalGiftCardsRequest:
      type: object
      required:
        - GetGiftCards
      properties:
        GetGiftCards:
          type: object
          required:
            - Request
          properties:
            Request:
              type: object
              required:
                - IssuerAccountNumber
                - Currency
                - Pagination
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                Currency:
                  $ref: "#/components/schemas/xCurrencyCode"
                Pagination:
                  $ref: "#/components/schemas/xPagination"

    GetDigitalGiftCardsResponse:
      type: object
      properties:
        GetGiftCardResponse:
          type: object
          properties:
            GetGiftCardResult:
              type: object
              properties:
                GiftCard:
                  type: array
                  items:
                    $ref: "#/components/schemas/xGiftCardDetail"
            PaginationTotal:
              $ref: "#/components/schemas/xPaginationTotal"
            OperationStatus:
              $ref: "#/components/schemas/xOpStatus"

    GetLinkedBankAccountsRequest:
      type: object
      required:
        - GetLinkedBankAccounts
      properties:
        GetLinkedBankAccounts:
          type: object
          required:
            - request
          properties:
            request:
              type: object
              required:
                - IssuerAccountNumber
                - RecipientUserId
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                RecipientUserId:
                  $ref: "#/components/schemas/xIAN"

    GetLinkedBankAccountsResponse:
      type: object
      properties:
        GetLinkedBankAccountsResponse:
          type: object
          properties:
            GetLinkedBankAccountResult:
              type: object
              properties:
                Beneficiary:
                  type: object
                  properties:
                    BeneficiaryDetails:
                      type: array
                      items:
                        $ref: "#/components/schemas/xBeneficiaryDetails"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    GetACHDebitLinkedBankAccountsRequest:
      type: object
      required:
        - GetACHDebitLinkedBankAccounts
      properties:
        GetACHDebitLinkedBankAccounts:
          type: object
          required:
            - request
          properties:
            request:
              type: object
              required:
                - IssuerAccountNumber
                - RecipientUserId
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                RecipientUserId:
                  $ref: "#/components/schemas/xIAN"

    GetACHDebitLinkedBankAccountsResponse:
      type: object
      properties:
        GetACHDebitLinkedBankAccountsResponse:
          type: object
          properties:
            GetACHDebitLinkedBankAccountsResult:
              type: object
              properties:
                Beneficiary:
                  type: object
                  properties:
                    BeneficiaryDetails:
                      type: array
                      items:
                        $ref: "#/components/schemas/xBeneficiaryDetails"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    GetBankWithdrawTypesRequest:
      type: object
      required:
        - GetBankWithdrawTypes
      properties:
        GetBankWithdrawTypes:
          type: object
          required:
            - request
          properties:
            request:
              type: object
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                CountryISO2:
                  $ref: "#/components/schemas/xCountryISO2"

    GetBankWithdrawTypesResponse:
      type: object
      properties:
        GetBankWithdrawTypesResponse:
          type: object
          properties:
            GetBankWithdrawTypeResult:
              type: object
              properties:
                WithdrawTypes:
                  description: >
                    How funds are withdrawn (*ACH* or *WIRE*)
                  type: array
                  items:
                    description: >
                      How funds are withdrawn (*ACH* or *WIRE*)
                    type: string
                  example:
                    - WIRE
                    - ACH
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    LinkBankBeneficiaryRequest:
      type: object
      required:
        - LinkBankBeneficiary
      properties:
        LinkBankBeneficiary:
          type: object
          required:
            - request
          properties:
            request:
              type: object
              required:
                - IssuerAccountNumber
                - UserID
                - Beneficiary
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                UserID:
                  $ref: "#/components/schemas/xIAN"
                Beneficiary:
                  type: object
                  required:
                    - BeneficiaryDetails
                    - BankDetails
                  properties:
                    BeneficiaryDetails:
                      $ref: "#/components/schemas/xBankBeneficiary"
                    BankDetails:
                      type: object
                      required:
                        - BeneficiaryBankInformation
                      properties:
                        BeneficiaryBankInformation:
                          $ref: "#/components/schemas/xBeneficiaryBankInformation"

    LinkBankBeneficiaryResponse:
      type: object
      properties:
        LinkBankBeneficiaryResponse:
          type: object
          properties:
            LinkBankBeneficiaryResult:
              type: object
              properties:
                BeneficiaryId:
                  description: >
                    Bank’s identifier for customer
                  # cannot REF this because this field is NULLABLE
                  nullable: true
                  type: string
                  example: "12e6f8ac19804f90be485045f50ace57"
                BeneficiaryStatus:
                  description: >
                    Pending, Approved, Declined, etc.
                  nullable: true
                  type: string
                  example: Pending
                AccountIdentityLevel:
                  description: >
                    Known, partially known, etc.
                  nullable: true
                  type: string
                  example: Pending
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    LinkACHDebitBankBeneficiaryRequest:
      type: object
      required:
        - LinkACHDebitBankBeneficiary
      properties:
        LinkACHDebitBankBeneficiary:
          type: object
          required:
            - request
          properties:
            request:
              type: object
              required:
                - IssuerAccountNumber
                - UserID
                - Beneficiary
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                UserID:
                  $ref: "#/components/schemas/xIAN"
                Beneficiary:
                  type: object
                  required:
                    - BeneficiaryDetails
                    - BankDetails
                  properties:
                    BeneficiaryDetails:
                      $ref: "#/components/schemas/xBankBeneficiary"
                    BankDetails:
                      type: object
                      required:
                        - BeneficiaryBankInformation
                      properties:
                        BeneficiaryBankInformation:
                          $ref: "#/components/schemas/xBeneficiaryBankInformation"


    LinkACHDebitBankBeneficiaryResponse:
      type: object
      properties:
        LinkACHDebitBankBeneficiaryResponse:
          type: object
          properties:
            LinkACHDebitBankBeneficiaryResult:
              type: object
              properties:
                BeneficiaryId:
                  # cannot REF this because this field is NULLABLE
                  description: >
                    Bank-dependent ID for beneficiary
                  nullable: true
                  type: string
                  example: "12e6f8ac19804f90be485045f50ace57"
                BeneficiaryStatus:
                  description: >
                    Pending, Approved, Declined, etc.
                  nullable: true
                  type: string
                  example: Pending
                AccountIdentityLevel:
                  description: >
                    Known, partially known, etc.
                  nullable: true
                  type: string
                  example: Pending
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    DeleteBankBeneficiaryRequest:
      type: object
      required:
        - DeleteBankBeneficiary
      properties:
        DeleteBankBeneficiary:
          type: object
          required:
            - request
          properties:
            request:
              type: object
              required:
                - IssuerAccountNumber
                - RecipientAccountNumber
                - BeneficiaryBankID
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                RecipientAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                BeneficiaryBankID:
                  $ref: "#/components/schemas/xBankBeneficiaryID"

    DeleteBankBeneficiaryResponse:
      type: object
      properties:
        DeleteBankBeneficiary:
          type: object
          properties:
            DeleteBankBeneficiaryResult:
              type: object
              properties:
                BeneficiaryBankID:
                  $ref: "#/components/schemas/xBankBeneficiaryID"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    GetBeneficiariesRequest:
      type: object
      required:
        - GetBeneficiaries
      properties:
        GetBeneficiaries:
          type: object
          required:
            - request
          properties:
            request:
              type: object
              required:
                - IssuerAccountNumber
                - Pagination
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                Pagination:
                  $ref: "#/components/schemas/xPagination"

    GetBeneficiariesResponse:
      type: object
      properties:
        GetBeneficiariesResponse:
          type: object
          properties:
            BeneficiaryResult:
              type: object
              properties:
                BeneficiaryDetail:
                  type: array
                  items:
                    $ref: "#/components/schemas/xBasicBeneficiaryDetail"
            PaginationTotal:
              $ref: "#/components/schemas/xPaginationTotal"

    CreateBeneficiaryRequest:
      type: object
      required:
        - CreateBeneficiary
      properties:
        CreateBeneficiary:
          type: object
          required:
            - Request
          properties:
            Request:
              type: object
              required:
                - IssuerAccountNumber
                - BeneficiaryCompanyName
                - WebAddress
                - BeneficiaryCompanyAdminDetails
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                BeneficiaryCompanyName:
                  $ref: "#/components/schemas/xFiftyCharString"
                WebAddress:
                  $ref: "#/components/schemas/xFiftyCharString"
                BeneficiaryCompanyAdminDetails:
                  $ref: "#/components/schemas/xAdminDetails"

    CreateBeneficiaryResponse:
      type: object
      properties:
        CreateBeneficiaryResponse:
          type: object
          properties:
            CreateBeneficiaryResult:
              type: object
              properties:
                BeneficiaryID:
                  $ref: "#/components/schemas/xIAN"
                AccountIdentityLevel:
                  description: >
                    Known, partially known, etc.
                ClientID:
                  $ref: "#/components/schemas/xAPIID"
                SecretKey:
                  $ref: "#/components/schemas/xAPISecretKey"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    CheckBeneficiaryExistRequest:
      type: object
      required:
        - CheckBeneficiaryExist
      properties:
        CheckBeneficiaryExist:
          type: object
          required:
            - request
          properties:
            request:
              type: object
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                CompanyName:
                  $ref: "#/components/schemas/xHundredCharString"

    CheckBeneficiaryExistResponse:
      type: object
      properties:
        CheckBeneficiaryExistResponse:
          type: object
          properties:
            CheckBeneficiaryExistResult:
              type: object
              properties:
                Beneficary:
                  type: array
                  items:
                    $ref: "#/components/schemas/xBeneficiaryExistDetail"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

GetLinkeCardRequest:
type: object
required:
- GetLinkedCardAccounts
properties:
GetLinkedCardAccounts:
type: object
required:
- request
properties:
request:
type: object
required:
- IssuerAccountNumber
- RecipientUserId
properties:
IssuerAccountNumber:
$ref: "#/components/schemas/xIAN"
RecipientUserId:
$ref: "#/components/schemas/xIAN"

    GetLinkeCardResponse:
      type: object
      properties:
        GetLinkeCardResponse:
          type: object
          properties:
            GetLinkedCardAccountsResponse:
              type: object
              properties:
                Card:
                  type: object
                  properties:
                    CardDetails:
                      type: array
                      items:
                        $ref: "#/components/schemas/xCardDetails"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

LinkCardRequest:
type: object
required:
- LinkCard
properties:
LinkCard:
type: object
required:
- request
properties:
request:
type: object
required:
- IssuerAccountNumber
- UserID
- card
properties:
IssuerAccountNumber:
$ref: "#/components/schemas/xIAN"
UserID:
$ref: "#/components/schemas/xIAN"
Beneficiary:
type: object
required:
- Card
properties:
Card:
$ref: "#/components/schemas/xLinkCard"

    LinkCardResponse:
      type: object
      properties:
        LinkCardResponse:
          type: object
          properties:
            LinkCardResult:
              type: object
              properties:
                CardToken:
                  description: >
                    Bank’s identifier for customer
                  # cannot REF this because this field is NULLABLE
                  nullable: true
                  type: string
                  example: "224780E21B0D4A2"
                CardStatus:
                  description: >
                    Pending, Approved, Declined, etc.
                  nullable: true
                  type: string
                  example: Approved
                AccountIdentityLevel:
                  description: >
                    Known, partially known, etc.
                  nullable: true
                  type: string
                  example: Pending
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"


DeleteCardRequest:
type: object
required:
- DeleteCard
properties:
DeleteCard:
type: object
required:
- request
properties:
request:
type: object
required:
- IssuerAccountNumber
- RecipientAccountNumber
- CardToken
properties:
IssuerAccountNumber:
$ref: "#/components/schemas/xIAN"
RecipientAccountNumber:
$ref: "#/components/schemas/xIAN"
BeneficiaryBankID:
$ref: "#/components/schemas/xCardToken"

    DeleteCardResponse:
      type: object
      properties:
        DeleteCard:
          type: object
          properties:
            DeleteCardResult:
              type: object
              properties:
                CardToken:
                  $ref: "#/components/schemas/xCardToken"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    GetCompanyTypeResponse:
      type: object
      properties:
        GetCompanyTypeResponse:
          type: object
          properties:
            GetCompanyTypeResult:
              type: object
              properties:
                CompanyTypes:
                  type: object
                  properties:
                    CompanyTypeDetails:
                      type: array
                      items:
                        type: object
                        properties:
                          CompanyType:
                            description: >
                              The legal structure of the company (Incorporated, Limited Liability Company, etc)
                            type: string
                      example:
                        - CompanyType: Akciju sabiedrība (AS)
                        - CompanyType: Akcinė bendrovė (AB)
                        - CompanyType: Akciová společnost
                        - CompanyType: Akciová spoločnosť (a.s.)
                        - CompanyType: Aksjeselskap (AS)
                        - CompanyType: Aktiebolag (AB)
                        - CompanyType: Aktiengesellschaft
                        - CompanyType: Aktiengesellschaft (AG)
                        - CompanyType: Aktieselskab (A/S)
                        - CompanyType: Aktsiaselts (AS)
                        - CompanyType: aktsionerno drujestvo (AD)
                        - CompanyType: Allmennaksjeselskap (ASA)
                        - CompanyType: Anónimi Etería (A.E.)
                        - CompanyType: Anpartsselskab (ApS)
                        - CompanyType: Anstalt
                        - CompanyType: Ansvarlig selskap (ANS)
                        - CompanyType: Association
                        - CompanyType: Association/Non-Profit Organization
                        - CompanyType: Autre
                        - CompanyType: avoin yhtiö (Ay)
                        - CompanyType: Badan Usaha Milik Negara (BUMN)
                        - CompanyType: Beperkte Aansprakelijkheid (B.A.)
                        - CompanyType: Besloten vennootschap (bv)
                        - CompanyType: besloten vennootschap met beperkte aansprakelijkheid (BVBA)
                        - CompanyType: betéti társaság (bt.)
                        - CompanyType: Bezpeļņas organizācija (BO)
                        - CompanyType: C Corporation
                        - CompanyType: Commanditaire vennootschap (cv)
                        - CompanyType: commanditaire vennootschap op aandelen (Comm VA)
                        - CompanyType: Commanditaire Vennotschap (CV)
                        - CompanyType: Company limited by guarantee
                        - CompanyType: coöperatieve vennootschap met beperkte aansprakelijkheid (CVBA)
                        - CompanyType: coöperatieve vennootschap met onbeperkte aansprakelijkheid (CVOA)
                        - CompanyType: Cooperativa de Responsabilidade Limitada (CRL)
                        - CompanyType: Corporation
                        - CompanyType: Delniška družba (d.d.)
                        - CompanyType: Designated Activity Company (DAC)
                        - CompanyType: dioničko društvo (d.d.)
                        - CompanyType: drujestvo s ogranichena otgovornost (OOD)
                        - CompanyType: društvo s ograničenom odgovornošću (d.o.o.)
                        - CompanyType: Družba z neomejeno odgovornostjo (d.n.o.)
                        - CompanyType: Družba z omejeno odgovornostjo (d.o.o.)
                        - CompanyType: družstvo
                        - CompanyType: ednolichen turgovetz (ET)
                        - CompanyType: ednolichno aktsionerno druzhestvo (EAD)
                        - CompanyType: ednolichno druzhestvo s ogranichena otgovornost (EOOD)
                        - CompanyType: eenmanszaak
                        - CompanyType: Eenmanszaak
                        - CompanyType: eenpersoons besloten vennootschap met beperkte aansprakelijkheid (EBVBA)
                        - CompanyType: egyéni cég (e.c.)
                        - CompanyType: egyéni vállalkozó (e.v.)
                        - CompanyType: eingetragenes Einzelunternehmen (eU)
                        - CompanyType: einkahlutafélag (ehf)
                        - CompanyType: einstaklingsfyrirtæki
                        - CompanyType: Einzelunternehmen
                        - CompanyType: Empresa Individual de Responsabilidad Limitada (EIRL)
                        - CompanyType: Enkeltmandsvirksomhed
                        - CompanyType: Enkeltpersonforetak
                        - CompanyType: Enskild firma
                        - CompanyType: Entreprise individuelle (EI)
                        - CompanyType: Entreprise unipersonnelle à responsabilité limitée (EURL)
                        - CompanyType: Etería Periorisménis Euthínis (E.P.E.
                        - CompanyType: Eterórithmi Etería (E.E.)
                        - CompanyType: EURL
                        - CompanyType: Exempt Private Company (EPC)
                        - CompanyType: Firma (Fa)
                        - CompanyType: Fond commun de placement (FCP)
                        - CompanyType: Fonds commun de placement (FCP)
                        - CompanyType: Forening
                        - CompanyType: Füüsilisest isikust ettevõtja (FIE)
                        - CompanyType: General Partnership
                        - CompanyType: General Partnership (GP)
                        - CompanyType: Gesellschaft bürgerlichen Rechts (GbR)
                        - CompanyType: Gesellschaft des bürgerlichen Rechts (GesbR)
                        - CompanyType: Gesellschaft mit beschränkter Haftung (GmbH)
                        - CompanyType: gewone commanditaire vennootschap (Comm V)
                        - CompanyType: Government Entity
                        - CompanyType: Handelsbolag (HB)
                        - CompanyType: hlutafélag (hf.)
                        - CompanyType: Ideell förening
                        - CompanyType: Idiotiki kefaleouhiki Eteria
                        - CompanyType: Imprenditore
                        - CompanyType: Inc. (Incorporated)
                        - CompanyType: Incorporated (Inc)
                        - CompanyType: Individuālais komersants (IK)
                        - CompanyType: Individuali veikla
                        - CompanyType: Interessentskab (I/S)
                        - CompanyType: Investmentgesellschaft mit festem Kapital (SICAF)
                        - CompanyType: Investmentgesellschaft mit variablem Kapital (SICAV)
                        - CompanyType: Iværksætterselskab (IVS)
                        - CompanyType: javno trgovačko društvo (j.t.d.)
                        - CompanyType: jednoosobowa działalność gospodarcza
                        - CompanyType: jednostavno društvo s ograničenom odgovornošću (j.d.o.o.)
                        - CompanyType: julkinen osakeyhtiö (Oyj)
                        - CompanyType: Kollektivgesellschaft (GP)
                        - CompanyType: Komanditinė ūkinė bendrija (KUB)
                        - CompanyType: Komanditna družba (k.d.)
                        - CompanyType: Komanditná spoločnosť (k.s.)
                        - CompanyType: komanditní společnost (k.s.)
                        - CompanyType: komanditno društvo (k.d.)
                        - CompanyType: komanditno druzhestvo (KD)
                        - CompanyType: komanditno druzhestvo s aktzii (KDA)
                        - CompanyType: Komandītsabiedrība (KS)
                        - CompanyType: kommandiittiyhtiö (Ky)
                        - CompanyType: Kommanditbolag (KB)
                        - CompanyType: Kommanditgesellschaft (KG)
                        - CompanyType: Kommanditgesellschaft (LP)
                        - CompanyType: Kommanditgesellschaft auf Aktien (KgaA)
                        - CompanyType: Kommanditselskab (K/S)
                        - CompanyType: Kommandittselskap (KS)
                        - CompanyType: Kooperativ
                        - CompanyType: korlátolt felelősségű társaság (kft.)
                        - CompanyType: közkereseti társaság (kkt)
                        - CompanyType: Limitada (Lda.)
                        - CompanyType: Limited (Ltd.)
                        - CompanyType: Limited Liability Company (LLC)
                        - CompanyType: Limited Liability Corporation
                        - CompanyType: Limited Liability Limited Partnership (LLLP)
                        - CompanyType: Limited Liability Partnership
                        - CompanyType: Limited Liability Partnership (LLP)
                        - CompanyType: Limited Partnership
                        - CompanyType: Limited Partnership (LP)
                        - CompanyType: Look Through Company (LTC)
                        - CompanyType: Ltd. (Limited)
                        - CompanyType: Mažoji bendrija (MB)
                        - CompanyType: Monoprósopi Etería Periorisménis Euthínis
                        - CompanyType: naamloze vennootschap (NV)
                        - CompanyType: Naamloze vennootschap (nv)
                        - CompanyType: NL (No Liability)
                        - CompanyType: No-Liability Company (NL)
                        - CompanyType: nyilvánosan működő részvénytársaság (Nyrt.)
                        - CompanyType: obrt
                        - CompanyType: offene Gesellschaft (OG)
                        - CompanyType: Offene Handelsgesellschaft (OHG)
                        - CompanyType: Omórithmi Etería
                        - CompanyType: ortakluk
                        - CompanyType: osakeyhtiö (Oy)
                        - CompanyType: Osaühing (OU)
                        - CompanyType: osuuskunta (osk)
                        - CompanyType: Other
                        - CompanyType: Partnerselskab or Kommanditaktieselskab (P/S)
                        - CompanyType: Partnership
                        - CompanyType: Persekutuan Komanditer
                        - CompanyType: Persekutuan Perdata (Maatschap)
                        - CompanyType: Perseroan Terbatas (PT)
                        - CompanyType: Perseroan Terbatas Terbuka or Perseroan Terbuka (PT Tbk)
                        - CompanyType: persoana fizica autorizata (PFA)
                        - CompanyType: Persona Física
                        - CompanyType: Perusahaan Dagang (PD)
                        - CompanyType: Pilnsabiedrība (PS)
                        - CompanyType: Private Corporation
                        - CompanyType: Private Corporation (Pty Ltd)
                        - CompanyType: Private Corporation (PtyLtd)
                        - CompanyType: Private Fund Limited Partnership (PFLP)
                        - CompanyType: Private Liability Company
                        - CompanyType: Private Limited Company
                        - CompanyType: Private Limited Company/Sendirian Berhad (Pte Ltd/Sdn Bhd)
                        - CompanyType: Private Limited Liability Company
                        - CompanyType: Private Ltd Company (Ltd)
                        - CompanyType: Private stichting
                        - CompanyType: Privatstiftung
                        - CompanyType: Professional Limited Liability Company (PLLC)
                        - CompanyType: Proprietary Limited (Pty Ltd)
                        - CompanyType: Proprietary Limited Company (Pty. Ltd.)
                        - CompanyType: Pty. (Unlimited Proprietary) company with share capital
                        - CompanyType: Pty. Ltd. (Proprietary Limited Company)
                        - CompanyType: Public Company
                        - CompanyType: Public Company (Ltd).
                        - CompanyType: Public Company (Quoted)
                        - CompanyType: Public Company (Unquoted)
                        - CompanyType: Public Coproration (Ltd)
                        - CompanyType: Public Corporation
                        - CompanyType: Public Limited Company
                        - CompanyType: Public Limited Company (listed)
                        - CompanyType: Public Limited Company (unlisted)
                        - CompanyType: Public Limited Company/Berhad (Ltd/Bhd)
                        - CompanyType: Public Limited Liability Company
                        - CompanyType: Publikt aktiebolag (AB (publ)
                        - CompanyType: Registered Charity
                        - CompanyType: S Corporation
                        - CompanyType: säätiö
                        - CompanyType: Sabiedrība ar ierobežotu atbildību (SIA)
                        - CompanyType: sameignarfélag (sf.)
                        - CompanyType: samlagsfélag
                        - CompanyType: Samostojni podjetnik (s.p.)
                        - CompanyType: samvinnufélag
                        - CompanyType: SARL
                        - CompanyType: Single Member Company
                        - CompanyType: sjálfseignarstofnun
                        - CompanyType: Sociedad Anonima (SA)
                        - CompanyType: Sociedad Anónima (SA)
                        - CompanyType: Sociedad Colectiva (SC)
                        - CompanyType: Sociedad Comanditaria (S Cra)
                        - CompanyType: Sociedad Cooperativa (S Coop)
                        - CompanyType: Sociedad de Responsabilidad Limitada (Ltda)
                        - CompanyType: Sociedad de Responsabilidad Limitada (S de RL)
                        - CompanyType: Sociedad en Comandita por Acciones (S en C)
                        - CompanyType: Sociedad en Comandita Simple (S en C)
                        - CompanyType: Sociedad Limitada (SL)
                        - CompanyType: Sociedad Limitada Nueva Empresa (SLNE)
                        - CompanyType: Sociedad por Acciones (SpA)
                        - CompanyType: Sociedade Aberta (S.A.
                        - CompanyType: Sociedade Anónima (S.A.)
                        - CompanyType: Sociedade Fechada (S.F.)
                        - CompanyType: Sociedade Gestora de Participações Sociais (SGPS)
                        - CompanyType: Sociedades por Acciones Simplificada (SAS)
                        - CompanyType: Società a responsabilità limitata (Srl)
                        - CompanyType: Società cooperativa a responsabilità limitata (Scrl)
                        - CompanyType: Societa Europea
                        - CompanyType: Società in accomandita per azioni (Sapa)
                        - CompanyType: Società in accomandita semplice (Sas)
                        - CompanyType: Società in nome collettivo (snc)
                        - CompanyType: Società per azioni (Spa)
                        - CompanyType: Società semplice (Ss)
                        - CompanyType: societate cu răspundere limitată (S.R.L.)
                        - CompanyType: societate în comandită pe acţiuni (S.C.A.)
                        - CompanyType: societate în comandită simplă (S.C.S.)
                        - CompanyType: societate în nume colectiv (S.N.C.
                        - CompanyType: Societate pe Acţiuni (S.A.)
                        - CompanyType: Societatea cu răspundere limitată cu proprietar unic (SRL)
                        - CompanyType: Société à responsabilité limitée (S.A.R.L.)
                        - CompanyType: Société à responsabilité limitée (SARL)
                        - CompanyType: Société Anonyme
                        - CompanyType: Société anonyme (S.A.)
                        - CompanyType: Société anonyme (SA)
                        - CompanyType: Société cotée en bourse
                        - CompanyType: Société d’investissement à capital fixe (SICAF)
                        - CompanyType: Société d’Investissement à Capital variable (SICAV)
                        - CompanyType: Société d'investissement à capital fixe
                        - CompanyType: Société d'investissement à capital variable (SICAV)
                        - CompanyType: Société en commandite par actions (SCA)
                        - CompanyType: Société en commandite simple (SCS)
                        - CompanyType: Société en commandite simple (SECS)
                        - CompanyType: Société en nom collectif (SNC)
                        - CompanyType: Société par actions simplifiée (SAS)
                        - CompanyType: Société pas actions simplifiée
                        - CompanyType: Sole Proprietorship
                        - CompanyType: Sole Trader
                        - CompanyType: Sole Trader (Proprietorship)
                        - CompanyType: Spółdzielnia
                        - CompanyType: Společnost s ručením omezeným
                        - CompanyType: spółka akcyjna (S.A.)
                        - CompanyType: spółka jawna (sp.j.)
                        - CompanyType: spółka komandytowa (sp.k.)
                        - CompanyType: spółka komandytowo-akcyjna (S.K.A.)
                        - CompanyType: spółka partnerska (sp.p.)
                        - CompanyType: spółka z ograniczoną odpowiedzialnością (Sp. Z o.o.)
                        - CompanyType: Spoločnosť s ručením obmedzeným (s.r.o.)
                        - CompanyType: Stichting
                        - CompanyType: stichting van openbaar nut
                        - CompanyType: Stiftelse
                        - CompanyType: Stiftung
                        - CompanyType: Stiftung / Fondation / Fondazione
                        - CompanyType: stille Gesellschaft (stG)
                        - CompanyType: subiratelno druzhestvo (SD)
                        - CompanyType: Superannuation Fund
                        - CompanyType: szövetkezet (szov.)
                        - CompanyType: Täisühing (TU)
                        - CompanyType: Tikroji ūkinė bendrija (TUB)
                        - CompanyType: toiminimi (T mi)
                        - CompanyType: Toko
                        - CompanyType: Treuhandschaft
                        - CompanyType: Treuunternehmen
                        - CompanyType: U- unipersonnelle (EURL)
                        - CompanyType: Uitsluiting van Aansprakelijkheid (U.A.)
                        - CompanyType: Unipessoal Lda
                        - CompanyType: Unlimited Company
                        - CompanyType: Unlimited Company (UC)
                        - CompanyType: Unlimited Proprietary Company (Pty.)
                        - CompanyType: Unternehmergesellschaft (UG)
                        - CompanyType: Usaha Dagang (UD)
                        - CompanyType: Usaldusühing (UU)
                        - CompanyType: Uždaroji akcinė bendrovė (UAB)
                        - CompanyType: Variable Capital Company
                        - CompanyType: vennootschap onder firma (VOF)
                        - CompanyType: Vennootschap onder firma (vof)
                        - CompanyType: Verein
                        - CompanyType: Verein / Association / Associazione
                        - CompanyType: Vereinigung ohne Gewinnerzielungsabsicht (VoG)
                        - CompanyType: Verejná obchodná spoločnosť (v.o.s.)
                        - CompanyType: veřejná obchodní společnost (v.o.s.)
                        - CompanyType: Viešoji įstaiga (VsI)
                        - CompanyType: Wettelijke Aansprakelijkheid (W.A.)
                        - CompanyType: Yayasan
                        - CompanyType: zadruga
                        - CompanyType: živnost
                        - CompanyType: živnosť

    GetAdvancedContactJobTitlesResponse:
      type: object
      properties:
        JobTitlesResponse:
          type: object
          properties:
            JobTitlesResult:
              type: object
              properties:
                JobTitles:
                  type: object
                  properties:
                    JobTitlesDetails:
                      type: array
                      items:
                        type: object
                        properties:
                          JobTitle:
                            description: >
                              Typical job title
                            type: string
                  example:
                    JobTitles:
                      JobTitlesDetails:
                        - JobTitle: Accountant
                        - JobTitle: Accounts Manager
                        - JobTitle: Accounts Payable Clerk
                        - JobTitle: Accounts Payable Manager
                        - JobTitle: Accounts Receivable Clerk
                        - JobTitle: Accounts Receivable Manager
                        - JobTitle: Cash Manager
                        - JobTitle: CEO
                        - JobTitle: CFO
                        - JobTitle: Chairman
                        - JobTitle: Commercial Director
                        - JobTitle: Company Secretary
                        - JobTitle: Comptroller
                        - JobTitle: Controller
                        - JobTitle: COO
                        - JobTitle: Director
                        - JobTitle: Finance Manager
                        - JobTitle: Finance Officer
                        - JobTitle: Financial Controller
                        - JobTitle: Financial Director
                        - JobTitle: General Manager
                        - JobTitle: Group Accountant
                        - JobTitle: Group Finance Director
                        - JobTitle: Group Treasury Manager
                        - JobTitle: Head of Finance
                        - JobTitle: Head of Treasury
                        - JobTitle: Managing Director
                        - JobTitle: Office Manager
                        - JobTitle: Other
                        - JobTitle: Owner
                        - JobTitle: President
                        - JobTitle: Seafarer
                        - JobTitle: Treasury Manager
                        - JobTitle: VP of Finance
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    GetIdentificationTypeResponse:
      type: object
      properties:
        IdentificationTypeResponse:
          type: object
          properties:
            IndentificationTypeResult:
              type: object
              properties:
                IdentificationType:
                  type: object
                  properties:
                    IdentificationTypeDetails:
                      type: array
                      items:
                        type: object
                        properties:
                          IdentificationType:
                            description: >
                              Name and nature of the identity document
                            type: string
                  example:
                    Identification Type:
                      IdentificationTypeDetails:
                        - IdentificationType: CitizenshipCard
                        - IdentificationType: DriversLicense
                        - IdentificationType: IdentityCard
                        - IdentificationType: NationalIdentificationDocument
                        - IdentificationType: NationalRegistrationIdentityCard
                        - IdentificationType: Passport
                        - IdentificationType: VotingCard
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    GetAdvancedContactStateRequest:
      type: object
      required:
        - AdvancedContactState
      properties:
        AdvancedContactState:
          type: object
          required:
            - request
          properties:
            request:
              type: object
              properties:
                CountryCode:
                  $ref: "#/components/schemas/xCountryISO2"

    CompanyAdvancedProfileStatusRequest:
      type: object
      required:
        - AdvancedProfileStatus
      properties:
        AdvancedProfileStatus:
          type: object
          required:
            - request
          properties:
            request:
              type: object
              required:
                - IssuerAccountNumber
                - BeneficiaryAccountNumber
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                BeneficiaryAccountNumber:
                  $ref: "#/components/schemas/xIAN"

    CompanyAdvancedProfileStatusResponse:
      type: object
      properties:
        AdvancedProfileStatusResponse:
          type: object
          properties:
            AdvancedProfileStatusResult:
              type: object
              properties:
                BeneficiaryID:
                  $ref: "#/components/schemas/xIAN"
                Message:
                  $ref: "#/components/schemas/xProfileStatus"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    GetProgramsRequest:
      type: object
      required:
        - GetPrograms
      properties:
        GetPrograms:
          type: object
          required:
            - request
          properties:
            request:
              type: object
              required:
                - IssuerAccountNumber
                - Pagination
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                Pagination:
                  $ref: "#/components/schemas/xPagination"

    GetProgramsResponse:
      type: object
      properties:
        GetProgramsResponse:
          type: object
          properties:
            ProgramsResult:
              type: object
              properties:
                ProgramDetails:
                  type: array
                  items:
                    $ref: "#/components/schemas/xProgramDetails"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"
                PaginationTotal:
                  $ref: "#/components/schemas/xPaginationTotal"

    ProgramCategoryRequest:
      type: object
      required:
        - GetProgramCategory
      properties:
        GetProgramCategory:
          type: object
          required:
            - request
          properties:
            request:
              type: object
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                Pagination:
                  $ref: "#/components/schemas/xPagination"

    ProgramCategoryResponse:
      type: object
      properties:
        GetProgramCategoryResponse:
          type: object
          properties:
            ProgramCategoryResult:
              type: object
              properties:
                ProgramCategory:
                  $ref: "#/components/schemas/xProgramCategory"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"
                PaginationTotal:
                  $ref: "#/components/schemas/xPaginationTotal"

    xProgramCategory:
      type: object
      properties:
        ProgramCategoryDetails:
          type: array
          items:
            type: object
            properties:
              CategoryID:
                $ref: "#/components/schemas/xProgramCategoryID"
              CategoryName:
                $ref: "#/components/schemas/xCategoryProgramName"
          example:
            - CategoryID: 1
              CategoryName: Promotional
            - CategoryID: 2
              CategoryName: Performance
            - CategoryID: 7
              CategoryName: Project
            - CategoryID: 8
              CategoryName: Other

    xCategoryProgramName:
      description: >
        A user-friendly display name for the category or subcategory
      type: string

    xPaymentStatus:
      description: >
        Status of payment, *completed*, *pending*, etc.
      type: string
      example: "Pending"

    xProgramCategoryID:
      description: >
        Identification code for category or subcategory
      minLength: 1
      maxLength: 5
      type: string
      pattern: ([1-9]\d{0,4})
      example: "6"

    ProgramTypeRequest:
      type: object
      required:
        - GetProgramType
      properties:
        GetProgramType:
          type: object
          required:
            - request
          properties:
            request:
              type: object
              required:
                - IssuerAccountNumber
                - CategoryID
                - Pagination
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                CategoryID:
                  $ref: "#/components/schemas/xProgramCategoryID"
                Pagination:
                  $ref: "#/components/schemas/xPagination"

    ProgramTypeResponse:
      type: object
      properties:
        GetProgramTypeResponse:
          type: object
          properties:
            ProgramTypeResult:
              type: object
              properties:
                ProgramType:
                  type: object
                  properties:
                    ProgramTypeDetail:
                      type: array
                      items:
                        type: object
                        properties:
                          SubCategoryID:
                            $ref: "#/components/schemas/xProgramCategoryID"
                          SubCategoryName:
                            $ref: "#/components/schemas/xCategoryProgramName"
                  example:
                    ProgramType:
                      ProgramTypeDetail:
                        - SubCategoryID: 4
                          SubCategoryName: "SPIFF"
                        - SubCategoryID: 5
                          SubCategoryName: "Bonus"
                        - SubCategoryID: 7
                          SubCategoryName: "Rebates"
                        - SubCategoryID: 8
                          SubCategoryName: "MDF"
                        - SubCategoryID: 9
                          SubCategoryName: "Other"

                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"
                PaginationTotal:
                  $ref: "#/components/schemas/xPaginationTotal"

    CreateProgramRequest:
      type: object
      required:
        - CreateProgram
      properties:
        CreateProgram:
          type: object
          required:
            - request
          properties:
            request:
              type: object
              required:
                - IssuerAccountNumber
                - Description
                - TransactionCategoryID
                - TransactionSubCategoryID
                - ClaimAmount
                - CurrencyCode
                - IsClaim
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                description:
                  $ref: "#/components/schemas/xHundredCharString"
                TransactionCategoryID:
                  $ref: "#/components/schemas/xProgramCategoryID"
                TransactionSubCategoryID:
                  $ref: "#/components/schemas/xProgramCategoryID"
                ClaimAmount:
                  $ref: "#/components/schemas/xAmount"
                CurrencyCode:
                  $ref: "#/components/schemas/xCurrencyCode"
                IsClaim:
                  description: >
                    True or false code
                  type: boolean

    CreateProgramResponse:
      type: object
      properties:
        CreateProgramsResponse:
          type: object
          properties:
            CreateProgramsResult:
              type: object
              properties:
                ProgramID:
                  $ref: "#/components/schemas/xProgramID"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    UpdateProgramsRequest:
      type: object
      required:
        - UpdateProgram
      properties:
        UpdateProgram:
          type: object
          required:
            - request
          properties:
            request:
              type: object
              required:
                - IssuerAccountNumber
                - Description
                - TransactionCategoryID
                - TransactionSubCategoryID
                - ClaimAmount
                - CurrencyCode
                - IsClaim
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                ProgramID:
                  $ref: "#/components/schemas/xProgramID"
                description:
                  $ref: "#/components/schemas/xHundredCharString"
                TransactionCategoryID:
                  $ref: "#/components/schemas/xProgramCategoryID"
                TransactionSubCategoryID:
                  $ref: "#/components/schemas/xProgramCategoryID"
                ClaimAmount:
                  $ref: "#/components/schemas/xAmount"
                CurrencyCode:
                  $ref: "#/components/schemas/xCurrencyCode"
                IsClaim:
                  description: >
                    True or false code for request
                  type: boolean
                  example: true

    UpdateProgramsResponse:
      type: object
      properties:
        UpdateProgramsResponse:
          type: object
          properties:
            UpdateProgramResult:
              type: object
              properties:
                ProgramID:
                  $ref: "#/components/schemas/xProgramID"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    GetOTPAuthorizedVendorRequest:
      type: object
      required:
        - GetOTPAuthorizedVendor
      properties:
        GetOTPAuthorizedVendor:
          type: object
          required:
            - request
          properties:
            request:
              type: object
              required:
                - IssuerAccountNumber
                - RecipientUserId
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                RecipientUserId:
                  $ref: "#/components/schemas/xIAN"

    GetOTPAuthorizedVendorResponse:
      type: object
      properties:
        GetOTPAuthorizedVendorResponse:
          type: object
          properties:
            GetOTPAuthorizedVendorResult:
              type: object
              properties:
                Message:
                  description: >
                    Status or description
                  type: string
                  example: "OTP Send"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    ValidateOTPAuthorizeVendorRequest:
      type: object
      required:
        - ValidateOTPAuthorizeVendor
      properties:
        ValidateOTPAuthorizeVendor:
          type: object
          required:
            - request
          properties:
            request:
              type: object
              required:
                - IssuerAccountNumber
                - RecipientUserID
                - OneTimePassword
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                RecipientUserID:
                  $ref: "#/components/schemas/xIAN"
                OneTimePassword:
                  $ref: "#/components/schemas/xOTP"

    ValidateOTPAuthorizeVendorResponse:
      type: object
      properties:
        ValidateOTPAuthorizeVendorResponse:
          type: object
          properties:
            ValidateOTPAuthorizeVendorResult:
              type: object
              properties:
                Message:
                  description: >
                    Status or description of result (*Authorized*).
                  type: string
                  example: "Authorized"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    BeneficiaryCompanyWithdrawFundRequest:
      type: object
      required:
        - BeneficiaryCompanyWithdrawFund
      properties:
        BeneficiaryCompanyWithdrawFund:
          type: object
          required:
            - request
          properties:
            request:
              type: object
              required:
                - IssuerAccountNumber
                - Currency
                - Amount
                - PaymentMethodID
                - BeneficiaryAccountNumber
                - BeneficiaryWalletID
                - BeneficiaryLinkedBankID
                - OTP
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                Currency:
                  $ref: "#/components/schemas/xCurrencyCode"
                Amount:
                  $ref: "#/components/schemas/xAmount"
                PaymentMethodID:
                  $ref: "#/components/schemas/xPaymentMethodID"
                BeneficiaryAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                BeneficiaryWalletID:
                  $ref: "#/components/schemas/xWalletID"
                BeneficiaryLinkedBankID:
                  $ref: "#/components/schemas/xLinkedBankID"
                OTP:
                  $ref: "#/components/schemas/xOTP"
                SendTransferCodetoEmail:
                  $ref: "#/components/schemas/xBoolean"
                SendTransferCodetoMobile:
                  $ref: "#/components/schemas/xBoolean"

    xBoolean:
      description: >
        True / false
      type: boolean

    BeneficiaryCompanyWithdrawFundResponse:
      type: object
      properties:
        BeneficiaryCompanyWithdrawFundResponse:
          type: object
          properties:
            BeneficiaryCompanyWithdrawFundResult:
              type: object
              properties:
                PaymentDate:
                  $ref: "#/components/schemas/xDate"
                PaymentStatus:
                  $ref: "#/components/schemas/xPaymentStatus"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    GetGiftCardDetailsRequest:
      type: object
      required:
        - GetGiftCards
      properties:
        GetGiftCards:
          type: object
          required:
            - IssuerAccountNumber
            - SKU
          properties:
            IssuerAccountNumber:
              $ref: "#/components/schemas/xIAN"
            SKU:
              $ref: "#/components/schemas/xSKU"

    GetGiftCardDetailsResponse:
      type: object
      properties:
        GetGiftCardResponse:
          type: object
          properties:
            GetGiftCardResult:
              type: object
              properties:
                GiftCard:
                  type: array
                  items:
                    $ref: "#/components/schemas/xGiftCardDetail"
            OperationStatus:
              $ref: "#/components/schemas/xOpStatus"

    GetUserWalletTransactionDetailsRequest:
      type: object
      required:
        - GetUserTransactionDetails
      properties:
        GetUserTransactionDetails:
          type: object
          required:
            - Request
          properties:
            Request:
              type: object
              required:
                - IssuerAccountNumber
                - TransactionID
                - UserID
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                TransactionID:
                  $ref: "#/components/schemas/xTransactionID"
                UserID:
                  $ref: "#/components/schemas/xIAN"

    GetUserWalletTransactionDetailsResponse:
      type: object
      properties:
        GetUserTransactionDetailsReponse:
          type: object
          properties:
            GetUserTransactionDetailsResult:
              type: object
              properties:
                Field:
                  type: array
                  items:
                    properties:
                      Name:
                        description: >
                          The name of the field (equivalent to *name*:value)
                        type: string
                      Value:
                        description: >
                          The value of the field (equivalent to name:*value*)
                        type: string
                  example:
                    - Name: "Transfer Type"
                      Value: "WIRE"
                    - Name: "Transaction Status"
                      Value: "Submitted"
                    - Name: "Payment Amount"
                      Value: "521.43"
                    - Name: "Order Number"
                      Value: "25467"
                    - Name: "Date Booked"
                      Value: "20 Jan, 2021"
                    - Name: "Beneficiary Name"
                      Value: "WELLS FARGO BANK, NATIONAL"
                    - Name: "Account Number"
                      Value: "993993"
                    - Name: "Address"
                      Value: "1 Calle del Belmont"
                    - Name: "Beneficiary Bank Address"
                      Value: "ASSOCIATION 375 PARK AVE NEW YORK CITY"
                    - Name: "Beneficiary Bank Name"
                      Value: "Dawes, Tomes, Mousely, Grubbs Fidelity Fiduciary Bank"
                    - Name: "Beneficiary SWIFTBIC"
                      Value: "PNBPUS3NXXX"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    authorizationTokenRequest:
      type: object
      required:
        - grant_type
        - client_id
        - client_secret
      properties:
        grant_type:
          type: string
          description: >
            always has value `password` or `refresh_token`
          enum:
            - password
            - refresh_token
          example: "password"
        client_id:
          type: string
          description: >
            XTRM-supplied client id
          example: "{{xtrm_client_id}}"
        client_secret:
          type: string
          description: >
            XTRM-supplied client secret
          example: "{{xtrm_client_secret}}"
        refresh_token:
          nullable: true
          type: string
          description: >
            only present when renewing a token
          example: "66fa9c8ad2264077ab565aa1e38927b0"

    authorizationTokenResponse:
      type: object
      description: >
        Authorization and associated data
      allOf:
        - type: object
          properties:
            client_id:
              type: string
              description: >
                the client id
              example: 9999999_API_User
            access_token:
              type: string
              description: >
                Authorization token
              example: rlpRJOTSQxQt1QKs+VUKXKwYAdjuZXHk++P8fk+y/bYVeG77Xanucg8IhiXH2m6kF95YMN4fGfTNGIdp6201hIqOUKSCTsSGxgjtcrqkPchZ07wn20xt+55k6tWKeysdZdlulE9alPZWefSaBig+ekSITRO8cmndIRbRJBDIkF0C/fLFeGkMqNsVIcUU=
            token_type:
              type: string
              description: >
                Always 'bearer'
              enum:
                - bearer
              example: bearer
            expires_in:
              type: integer
              description: >
                Expiry time in seconds as string
              example: 86399
            refresh_token:
              type: string
              description: >
                Refresh token to get a new token to extend the session
              example: "8ad2264b06082ab565aa1e3896fa927c"
            .issued:
              type: string
              description: >
                time of creation
              example: "Mon, 22 Jan 2018 17:20:45 GMT"
            .expires:
              type: string
              description: >
                time of expiration
              example: "Tue, 23 Jan 2018 17:20:45 GMT"
          required:
            - access_token
            - token_type
            - expires_in
            - refresh_token
            - .issued
            - .expires

    GetCompanyAdvancedProfileDetailRequest:
      type: object
      required:
        - AdvancedProfileDetails
      properties:
        AdvancedProfileDetails:
          type: object
          required:
            - request
          properties:
            request:
              type: object
              required:
                - IssuerAccountNumber
                - BeneficiaryAccountNumber
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                BeneficiaryAccountNumber:
                  $ref: "#/components/schemas/xIAN"



    GetCompanyAdvancedProfileDetailResponse:
      type: object
      properties:
        AdvancedProfileDetailsResponse:
          type: object
          properties:
            AdvancedProfileDetailsResult:
              type: object
              properties:
                BeneficiaryAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                Country_Of_Registration_Code2:
                  $ref: "#/components/schemas/xCountryISO2"
                Status:
                  type: string
                  example: "Approved"
                BusinessEntityInformation:
                  $ref: "#/components/schemas/xBusinessEntityInformation"
                AuthorizedContactInformation:
                  $ref: "#/components/schemas/xAuthorizedContactInformation"
                DirectorInformation:
                  $ref: "#/components/schemas/xDirectorInformation"
                OwnerInformation:
                  $ref: "#/components/schemas/xOwnerInformation"


    GetConnectedStatusRequest:
      type: object
      required:
        - GetConnectedStatus
      properties:
        GetConnectedStatus:
          type: object
          required:
            - Request
          properties:
            Request:
              type: object
              required:
                - IssuerAccountNumber
                - RecipientUserID
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                RecipientUserID:
                  $ref: "#/components/schemas/xIAN"
                OneTimePassword:
                  $ref: "#/components/schemas/xOTP"

    GetConnectedStatusResponse:
      type: object
      properties:
        GetConnectedStatusResponse:
          type: object
          properties:
            GetConnectedStatusResult:
              type: object
              properties:
                Status:
                  description: >
                    Whether the beneficiary is connected, pending, or not connected
                  type: string
                  example: "Not Connected"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    CompanyAdvancedProfileRequest:
      type: object
      required:
        - AdvancedProfile
      properties:
        AdvancedProfile:
          type: object
          required:
            - request
          properties:
            request:
              type: object
              required:
                - IssuerAccountNumber
                - FromAccountNumber
                - Country_Of_Registration_Code2
                - BusinessEntityInformation
                - AuthorizedContactInformation
                - DirectorInformation
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                BeneficiaryAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                Country_Of_Registration_Code2:
                  $ref: "#/components/schemas/xCountryISO2"
                BusinessEntityInformation:
                  $ref: "#/components/schemas/xBusinessEntityInformation"
                AuthorizedContactInformation:
                  $ref: "#/components/schemas/xAuthorizedContactInformation"
                DirectorInformation:
                  $ref: "#/components/schemas/xDirectorInformation"
                OwnerInformation:
                  $ref: "#/components/schemas/xOwnerInformation"

    CompanyAdvancedProfileResponse:
      type: object
      properties:
        AdvancedProfileResponse:
          type: object
          properties:
            AdvancedProfileResponse:
              type: object
              properties:
                BeneficiaryID:
                  $ref: "#/components/schemas/xIAN"
                message:
                  type: string
                  example: "Advanced Profile Submitted Successfully"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    xDirectorInformation:
      type: object
      properties:
        Directors:
          type: array
          items:
            type: object
            properties:
              Job_Title:
                #description: Job title
                $ref: "#/components/schemas/xHundredCharString"
              First_Name:
                #description: First name
                $ref: "#/components/schemas/xFiftyCharString"
              Last_Name:
                #description: Last name
                $ref: "#/components/schemas/xFiftyCharString"
              Date_of_Birth:
                #description: Date of birth
                $ref: "#/components/schemas/xNumericDateString"
              Nationality:
                #description: Nationality
                $ref: "#/components/schemas/xHundredCharString"
              Identification_Type:
                #description: Identification Type
                $ref: "#/components/schemas/xHundredCharString"
              Identity_Document_Number:
                #description: Identification document number
                $ref: "#/components/schemas/xFiftyCharString"
              Identity_Document_Expiration:
                #description: Identification document expiry date
                $ref: "#/components/schemas/xNumericDateString"
              Jurisdiction:
                #description: Jurisdiction
                $ref: "#/components/schemas/xHundredCharString"
              Address_1:
                #description: Address line 1
                $ref: "#/components/schemas/xAddressLine"
              City:
                #description: City
                $ref: "#/components/schemas/xCity"
              Country_Code2:
                #description: Country ISO Code 2
                $ref: "#/components/schemas/xCountryISO2"
              Region_Code2:
                #description: Region Code
                $ref: "#/components/schemas/xRegion_Code2"
              Postal_Code:
                #description: Postal code
                type: string
      description: >


    xOwnerInformation:
      type: object
      properties:
        IsPublicallyTraded:
          type: string
        Owners:
          type: array
          items:
            type: object
            properties:
              Percentage_Owned:
                #description: Percentage owned by the owner
                $ref: "#/components/schemas/xNumericStringSix"
              First_Name:
                #description: First name of the owner
                $ref: "#/components/schemas/xFiftyCharString"
              Last_Name:
                #description: Last name of the owner
                $ref: "#/components/schemas/xFiftyCharString"
              Occupation:
                #description: Occupation of the owner
                $ref: "#/components/schemas/xHundredCharString"
              Source_of_Income:
                #description: Source of income
                $ref: "#/components/schemas/xHundredCharString"
              Date_of_Birth:
                #description: Date of birth
                $ref: "#/components/schemas/xNumericDateString"
              Nationality:
                #description: Nationality
                $ref: "#/components/schemas/xHundredCharString"
              Identity_Document_Type:
                #description: Identity document type
                $ref: "#/components/schemas/xFiftyCharString"
              Identity_Document_Number:
                #description: Identity document number
                $ref: "#/components/schemas/xHundredCharString"
              Identity_Document_Expiration:
                #description: Identity document expiry date
                $ref: "#/components/schemas/xNumericDateString"
              Jurisdiction:
                #description: Jurisdiction
                $ref: "#/components/schemas/xHundredCharString"
              Address_1:
                #description: Address line 1
                $ref: "#/components/schemas/xAddressLine"
              City:
                #description: City
                $ref: "#/components/schemas/xCity"
              Country_Code2:
                #description: Country ISO code 2
                $ref: "#/components/schemas/xCountryISO2"
              Region_Code2:
                #description: Region code
                $ref: "#/components/schemas/xRegion_Code2"
              Postal_Code:
                #description: Postal code
                $ref: "#/components/schemas/xPostalCode"


    GetUserPaymentMethodsResponse:
      type: object
      properties:
        GetUserPaymentMethodsResponse:
          type: object
          properties:
            GetUserPaymentMethodResult:
              type: object
              properties:
                UserPaymentMethods:
                  type: object
                  properties:
                    UserPaymentMethodDetails:
                      type: array
                      items:
                        type: object
                        properties:
                          UserPaymentMethodID:
                            description: >
                              An ID code (available via the API) for this payment method
                            type: string
                          UserPaymentMethodName:
                            description: >
                              A user-friendly description of the payment method (*Bank*, *Prepaid Virtual Debit Card*, etc.)
                            type: string
                      example:
                        - UserPaymentMethodID: "XTR94500"
                          UserPaymentMethodName: "Bank"
                        - UserPaymentMethodID: "XTR94503"
                          UserPaymentMethodName: "Prepaid Virtual Debit Card"
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    GetPrepaidCardsRequest:
      type: object
      required:
        - GetPrepaidCards
      properties:
        GetPrepaidCards:
          type: object
          required:
            - Request
          properties:
            Request:
              type: object
              required:
                - IssuerAccountNumber
                - Currency
                - Pagination
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                Currency:
                  $ref: "#/components/schemas/xCurrencyCode"
                Pagination:
                  $ref: "#/components/schemas/xPagination"

    GetPrepaidCardsResponse:
      type: object
      properties:
        GetPrepaidCardsResponse:
          type: object
          properties:
            GetPrepaidCardsResult:
              description: >
                *There is no documentation on the result of GetPrepaidCardsResponse*
              type: string
              example: "There is no documentation on the result of GetPrepaidCardsResponse"

    GetPrepaidCardDetailsRequest:
      type: object
      required:
        - GetPrepaidCardDetails
      properties:
        GetPrepaidCardDetails:
          type: object
          required:
            - Request
          properties:
            Request:
              type: object
              required:
                - IssuerAccountNumber
                - SKU
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                SKU:
                  $ref: "#/components/schemas/xSKU"

    GetPrepaidCardDetailsResponse:
      type: object
      properties:
        GetPrepaidCardDetailsResponse:
          type: object
          properties:
            GetPrepaidCardDetailsResult:
              description: >
                *There is no documentation on the result of GetPrepaidCardsResult*
              type: string
              example: "There is no documentation on the result of GetPrepaidCardsResult"

    UpdateBeneficiaryRequest:
      type: object
      required:
        - UpdateBeneficiary
      properties:
        UpdateBeneficiary:
          type: object
          required:
            - Request
          properties:
            Request:
              type: object
              required:
                - IssuerAccountNumber
                - BeneficiaryID
                - BeneficiaryCompanyName
                - WebAddress
                - BeneficiaryCompanyAdminDetails
                - SalesProgramDetails
              properties:
                IssuerAccountNumber:
                  $ref: "#/components/schemas/xIAN"
                BeneficiaryID:
                  $ref: "#/components/schemas/xIAN"
                BeneficiaryCompanyName:
                  $ref: "#/components/schemas/xCompanyName"
                WebAddress:
                  $ref: "#/components/schemas/xWebAddress"
                BeneficiaryCompanyAdminDetails:
                  $ref: "#/components/schemas/xBeneficiaryAdminDetails"
                SalesProgramDetails:
                  $ref: "#/components/schemas/xSalesProgramDetails"

    UpdateBeneficiaryResponse:
      type: string
      example: "No documentation exists for UpdateBeneficiaryRequest"

    #########################

    GetAdvancedContactCountryResponse:
      type: object
      properties:
        AdvancedContactCountryResponse:
          type: object
          properties:
            GetAdvancedContactCountryResult:
              type: object
              properties:
                AdvancedContactCountry:
                  type: object
                  properties:
                    AdvancedContactCountryDetails:
                      type: array
                      items:
                        type: object
                        properties:
                          CountryCode:
                            $ref: "#/components/schemas/xCountryISO2"
                          CountryName:
                            description: >
                              The full name of the country
                            type: string
                  example:
                    AdvancedContactCountry:
                      AdvancedContactCountryDetails:
                        - CountryCode: AF
                          CountryName: Afghanistan
                        - CountryCode: AX
                          CountryName: Aland Islands
                        - CountryCode: AL
                          CountryName: Albania
                        - CountryCode: DZ
                          CountryName: Algeria
                        - CountryCode: AS
                          CountryName: American Samoa
                        - CountryCode: AD
                          CountryName: Andorra
                        - CountryCode: AO
                          CountryName: Angola
                        - CountryCode: AI
                          CountryName: Anguilla
                        - CountryCode: AQ
                          CountryName: Antarctica
                        - CountryCode: AG
                          CountryName: Antigua And Barbuda
                        - CountryCode: AR
                          CountryName: Argentina
                        - CountryCode: AM
                          CountryName: Armenia
                        - CountryCode: AW
                          CountryName: Aruba
                        - CountryCode: AU
                          CountryName: Australia
                        - CountryCode: AT
                          CountryName: Austria
                        - CountryCode: AZ
                          CountryName: Azerbaijan
                        - CountryCode: BS
                          CountryName: Bahamas
                        - CountryCode: BH
                          CountryName: Bahrain
                        - CountryCode: BD
                          CountryName: Bangladesh
                        - CountryCode: BB
                          CountryName: Barbados
                        - CountryCode: BY
                          CountryName: Belarus
                        - CountryCode: BE
                          CountryName: Belgium
                        - CountryCode: BZ
                          CountryName: Belize
                        - CountryCode: BJ
                          CountryName: Benin
                        - CountryCode: BM
                          CountryName: Bermuda
                        - CountryCode: BT
                          CountryName: Bhutan
                        - CountryCode: BO
                          CountryName: Bolivia
                        - CountryCode: BQ
                          CountryName: Bonaire, Sint Eustatius And Saba
                        - CountryCode: BA
                          CountryName: Bosnia And Herzegovina
                        - CountryCode: BW
                          CountryName: Botswana
                        - CountryCode: BV
                          CountryName: Bouvet Island
                        - CountryCode: BR
                          CountryName: Brazil
                        - CountryCode: IO
                          CountryName: British Indian Ocean Territory
                        - CountryCode: BN
                          CountryName: Brunei Darussalam
                        - CountryCode: BG
                          CountryName: Bulgaria
                        - CountryCode: BF
                          CountryName: Burkina Faso
                        - CountryCode: BI
                          CountryName: Burundi
                        - CountryCode: KH
                          CountryName: Cambodia
                        - CountryCode: CM
                          CountryName: Cameroon
                        - CountryCode: CA
                          CountryName: Canada
                        - CountryCode: CV
                          CountryName: Cape Verde
                        - CountryCode: KY
                          CountryName: Cayman Islands
                        - CountryCode: CF
                          CountryName: Central African Republic
                        - CountryCode: TD
                          CountryName: Chad
                        - CountryCode: CL
                          CountryName: Chile
                        - CountryCode: CN
                          CountryName: China
                        - CountryCode: CX
                          CountryName: Christmas Island
                        - CountryCode: CC
                          CountryName: Cocos (Keeling) Islands
                        - CountryCode: CO
                          CountryName: Colombia
                        - CountryCode: KM
                          CountryName: Comoros
                        - CountryCode: CG
                          CountryName: Congo
                        - CountryCode: CD
                          CountryName: Congo, The Democratic Republic Of The
                        - CountryCode: CK
                          CountryName: Cook Islands
                        - CountryCode: CR
                          CountryName: Costa Rica
                        - CountryCode: CI
                          CountryName: Cote D'Ivoire
                        - CountryCode: HR
                          CountryName: Croatia
                        - CountryCode: CU
                          CountryName: Cuba
                        - CountryCode: CW
                          CountryName: Curaçao
                        - CountryCode: CY
                          CountryName: Cyprus
                        - CountryCode: CZ
                          CountryName: Czech Republic
                        - CountryCode: DK
                          CountryName: Denmark
                        - CountryCode: DJ
                          CountryName: Djibouti
                        - CountryCode: DM
                          CountryName: Dominica
                        - CountryCode: DO
                          CountryName: Dominican Republic
                        - CountryCode: TL
                          CountryName: East Timor
                        - CountryCode: EC
                          CountryName: Ecuador
                        - CountryCode: EG
                          CountryName: Egypt
                        - CountryCode: SV
                          CountryName: El Salvador
                        - CountryCode: GQ
                          CountryName: Equatorial Guinea
                        - CountryCode: ER
                          CountryName: Eritrea
                        - CountryCode: EE
                          CountryName: Estonia
                        - CountryCode: ET
                          CountryName: Ethiopia
                        - CountryCode: FK
                          CountryName: Falkland Islands (Malvinas)
                        - CountryCode: FO
                          CountryName: Faroe Islands
                        - CountryCode: FJ
                          CountryName: Fiji
                        - CountryCode: FI
                          CountryName: Finland
                        - CountryCode: FR
                          CountryName: France
                        - CountryCode: GF
                          CountryName: French Guiana
                        - CountryCode: PF
                          CountryName: French Polynesia
                        - CountryCode: TF
                          CountryName: French Southern Territories
                        - CountryCode: GA
                          CountryName: Gabon
                        - CountryCode: GM
                          CountryName: Gambia
                        - CountryCode: GE
                          CountryName: Georgia
                        - CountryCode: DE
                          CountryName: Germany
                        - CountryCode: GH
                          CountryName: Ghana
                        - CountryCode: GI
                          CountryName: Gibraltar
                        - CountryCode: GR
                          CountryName: Greece
                        - CountryCode: GL
                          CountryName: Greenland
                        - CountryCode: GD
                          CountryName: Grenada
                        - CountryCode: GP
                          CountryName: Guadeloupe
                        - CountryCode: GU
                          CountryName: Guam
                        - CountryCode: GT
                          CountryName: Guatemala
                        - CountryCode: GG
                          CountryName: Guernsey
                        - CountryCode: GN
                          CountryName: Guinea
                        - CountryCode: GW
                          CountryName: Guinea-Bissau
                        - CountryCode: GY
                          CountryName: Guyana
                        - CountryCode: HT
                          CountryName: Haiti
                        - CountryCode: HM
                          CountryName: Heard Island And Mcdonald Islands
                        - CountryCode: VA
                          CountryName: Holy See (Vatican City State)
                        - CountryCode: HN
                          CountryName: Honduras
                        - CountryCode: HK
                          CountryName: Hong Kong
                        - CountryCode: HU
                          CountryName: Hungary
                        - CountryCode: IS
                          CountryName: Iceland
                        - CountryCode: IN
                          CountryName: India
                        - CountryCode: ID
                          CountryName: Indonesia
                        - CountryCode: IR
                          CountryName: Iran, Islamic Republic Of
                        - CountryCode: IQ
                          CountryName: Iraq
                        - CountryCode: IE
                          CountryName: Ireland
                        - CountryCode: IM
                          CountryName: Isle Of Man
                        - CountryCode: IL
                          CountryName: Israel
                        - CountryCode: IT
                          CountryName: Italy
                        - CountryCode: JM
                          CountryName: Jamaica
                        - CountryCode: JP
                          CountryName: Japan
                        - CountryCode: JE
                          CountryName: Jersey
                        - CountryCode: JO
                          CountryName: Jordan
                        - CountryCode: KZ
                          CountryName: Kazakhstan
                        - CountryCode: KE
                          CountryName: Kenya
                        - CountryCode: KI
                          CountryName: Kiribati
                        - CountryCode: KP
                          CountryName: Korea, Democratic People'S Republic Of
                        - CountryCode: KR
                          CountryName: Korea, Republic Of
                        - CountryCode: KW
                          CountryName: Kuwait
                        - CountryCode: KG
                          CountryName: Kyrgyzstan
                        - CountryCode: LA
                          CountryName: Lao People'S Democratic Republic
                        - CountryCode: LV
                          CountryName: Latvia
                        - CountryCode: LB
                          CountryName: Lebanon
                        - CountryCode: LS
                          CountryName: Lesotho
                        - CountryCode: LR
                          CountryName: Liberia
                        - CountryCode: LY
                          CountryName: Libyan Arab Jamahiriya
                        - CountryCode: LI
                          CountryName: Liechtenstein
                        - CountryCode: LT
                          CountryName: Lithuania
                        - CountryCode: LU
                          CountryName: Luxembourg
                        - CountryCode: MO
                          CountryName: Macao
                        - CountryCode: MK
                          CountryName: Macedonia, The Former Yugoslav Republic Of
                        - CountryCode: MG
                          CountryName: Madagascar
                        - CountryCode: MW
                          CountryName: Malawi
                        - CountryCode: MY
                          CountryName: Malaysia
                        - CountryCode: MV
                          CountryName: Maldives
                        - CountryCode: ML
                          CountryName: Mali
                        - CountryCode: MT
                          CountryName: Malta
                        - CountryCode: MH
                          CountryName: Marshall Islands
                        - CountryCode: MQ
                          CountryName: Martinique
                        - CountryCode: MR
                          CountryName: Mauritania
                        - CountryCode: MU
                          CountryName: Mauritius
                        - CountryCode: YT
                          CountryName: Mayotte
                        - CountryCode: MX
                          CountryName: Mexico
                        - CountryCode: FM
                          CountryName: Micronesia, Federated States Of
                        - CountryCode: MD
                          CountryName: Moldova, Republic Of
                        - CountryCode: MC
                          CountryName: Monaco
                        - CountryCode: MN
                          CountryName: Mongolia
                        - CountryCode: ME
                          CountryName: Montenegro
                        - CountryCode: MS
                          CountryName: Montserrat
                        - CountryCode: MA
                          CountryName: Morocco
                        - CountryCode: MZ
                          CountryName: Mozambique
                        - CountryCode: MM
                          CountryName: Myanmar
                        - CountryCode: NA
                          CountryName: Namibia
                        - CountryCode: NR
                          CountryName: Nauru
                        - CountryCode: NP
                          CountryName: Nepal
                        - CountryCode: NL
                          CountryName: Netherlands
                        - CountryCode: AN
                          CountryName: Netherlands Antilles
                        - CountryCode: NC
                          CountryName: New Caledonia
                        - CountryCode: NZ
                          CountryName: New Zealand
                        - CountryCode: NI
                          CountryName: Nicaragua
                        - CountryCode: NE
                          CountryName: Niger
                        - CountryCode: NG
                          CountryName: Nigeria
                        - CountryCode: NU
                          CountryName: Niue
                        - CountryCode: NF
                          CountryName: Norfolk Island
                        - CountryCode: KP
                          CountryName: North Korea
                        - CountryCode: GB
                          CountryName: Northern Ireland
                        - CountryCode: MP
                          CountryName: Northern Mariana Islands
                        - CountryCode: NO
                          CountryName: Norway
                        - CountryCode: OM
                          CountryName: Oman
                        - CountryCode: PK
                          CountryName: Pakistan
                        - CountryCode: PW
                          CountryName: Palau
                        - CountryCode: PS
                          CountryName: Palestinian Territory, Occupied
                        - CountryCode: PA
                          CountryName: Panama
                        - CountryCode: PG
                          CountryName: Papua New Guinea
                        - CountryCode: PY
                          CountryName: Paraguay
                        - CountryCode: PE
                          CountryName: Peru
                        - CountryCode: PH
                          CountryName: Philippines
                        - CountryCode: PN
                          CountryName: Pitcairn
                        - CountryCode: PL
                          CountryName: Poland
                        - CountryCode: PT
                          CountryName: Portugal
                        - CountryCode: PR
                          CountryName: Puerto Rico
                        - CountryCode: QA
                          CountryName: Qatar
                        - CountryCode: RE
                          CountryName: Reunion
                        - CountryCode: RO
                          CountryName: Romania
                        - CountryCode: RU
                          CountryName: Russian Federation
                        - CountryCode: RW
                          CountryName: Rwanda
                        - CountryCode: BL
                          CountryName: Saint Barthélemy
                        - CountryCode: SH
                          CountryName: Saint Helena
                        - CountryCode: KN
                          CountryName: Saint Kitts And Nevis
                        - CountryCode: LC
                          CountryName: Saint Lucia
                        - CountryCode: MF
                          CountryName: Saint Martin
                        - CountryCode: PM
                          CountryName: Saint Pierre And Miquelon
                        - CountryCode: VC
                          CountryName: Saint Vincent And The Grenadines
                        - CountryCode: WS
                          CountryName: Samoa
                        - CountryCode: SM
                          CountryName: San Marino
                        - CountryCode: ST
                          CountryName: Sao Tome And Principe
                        - CountryCode: SA
                          CountryName: Saudi Arabia
                        - CountryCode: SN
                          CountryName: Senegal
                        - CountryCode: RS
                          CountryName: Serbia
                        - CountryCode: CS
                          CountryName: Serbia And Montenegro
                        - CountryCode: SC
                          CountryName: Seychelles
                        - CountryCode: SL
                          CountryName: Sierra Leone
                        - CountryCode: SG
                          CountryName: Singapore
                        - CountryCode: SX
                          CountryName: Sint Maarten
                        - CountryCode: SK
                          CountryName: Slovak Republic
                        - CountryCode: SK
                          CountryName: Slovakia
                        - CountryCode: SI
                          CountryName: Slovenia
                        - CountryCode: SB
                          CountryName: Solomon Islands
                        - CountryCode: SO
                          CountryName: Somalia
                        - CountryCode: ZA
                          CountryName: South Africa
                        - CountryCode: GS
                          CountryName: South Georgia And The South Sandwich Islands
                        - CountryCode: KR
                          CountryName: South Korea
                        - CountryCode: SS
                          CountryName: South Sudan
                        - CountryCode: ES
                          CountryName: Spain
                        - CountryCode: LK
                          CountryName: Sri Lanka
                        - CountryCode: SD
                          CountryName: Sudan
                        - CountryCode: SR
                          CountryName: Suriname
                        - CountryCode: SJ
                          CountryName: Svalbard And Jan Mayen
                        - CountryCode: SZ
                          CountryName: Swaziland
                        - CountryCode: SE
                          CountryName: Sweden
                        - CountryCode: CH
                          CountryName: Switzerland
                        - CountryCode: SY
                          CountryName: Syrian Arab Republic
                        - CountryCode: TW
                          CountryName: Taiwan, Province Of China
                        - CountryCode: TJ
                          CountryName: Tajikistan
                        - CountryCode: TZ
                          CountryName: Tanzania, United Republic Of
                        - CountryCode: TH
                          CountryName: Thailand
                        - CountryCode: TL
                          CountryName: Timor-Leste
                        - CountryCode: TG
                          CountryName: Togo
                        - CountryCode: TK
                          CountryName: Tokelau
                        - CountryCode: TO
                          CountryName: Tonga
                        - CountryCode: TT
                          CountryName: Trinidad And Tobago
                        - CountryCode: TN
                          CountryName: Tunisia
                        - CountryCode: TR
                          CountryName: Turkey
                        - CountryCode: TM
                          CountryName: Turkmenistan
                        - CountryCode: TC
                          CountryName: Turks And Caicos Islands
                        - CountryCode: TV
                          CountryName: Tuvalu
                        - CountryCode: UG
                          CountryName: Uganda
                        - CountryCode: UA
                          CountryName: Ukraine
                        - CountryCode: AE
                          CountryName: United Arab Emirates
                        - CountryCode: GB
                          CountryName: United Kingdom
                        - CountryCode: US
                          CountryName: United States
                        - CountryCode: UM
                          CountryName: United States Minor Outlying Islands
                        - CountryCode: UY
                          CountryName: Uruguay
                        - CountryCode: UZ
                          CountryName: Uzbekistan
                        - CountryCode: VU
                          CountryName: Vanuatu
                        - CountryCode: VE
                          CountryName: Venezuela
                        - CountryCode: VN
                          CountryName: Viet Nam
                        - CountryCode: VG
                          CountryName: Virgin Islands, British
                        - CountryCode: VI
                          CountryName: Virgin Islands, U.S.
                        - CountryCode: WF
                          CountryName: Wallis And Futuna
                        - CountryCode: EH
                          CountryName: Western Sahara
                        - CountryCode: YE
                          CountryName: Yemen
                        - CountryCode: ZM
                          CountryName: Zambia
                        - CountryCode: ZW
                          CountryName: Zimbabwe
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"

    GetAdvancedContactStateResponse:
      type: object
      properties:
        AdvancedContactStateResponse:
          type: object
          properties:
            AdvancedContactStateResult:
              type: object
              properties:
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"
                AdvancedContactState:
                  type: object
                  properties:
                    AdvancedContactStateDetails:
                      type: array
                      items:
                        type: object
                        properties:
                          StateCode:
                            description: >
                              Two letter abbreviation for state
                            type: string
                          StateName:
                            description: >
                              Full name of state
                            type: string
                  example:
                    AdvancedContactState:
                      AdvancedContactStateDetails:
                        - StateCode: AK
                          StateName: Alaska
                        - StateCode: AL
                          StateName: Alabama
                        - StateCode: AR
                          StateName: Arkansas
                        - StateCode: AS
                          StateName: American Samoa
                        - StateCode: AZ
                          StateName: Arizona
                        - StateCode: CA
                          StateName: California
                        - StateCode: CO
                          StateName: Colorado
                        - StateCode: CT
                          StateName: Connecticut
                        - StateCode: DC
                          StateName: District of Columbia
                        - StateCode: DE
                          StateName: Delaware
                        - StateCode: FL
                          StateName: Florida
                        - StateCode: GA
                          StateName: Georgia
                        - StateCode: GU
                          StateName: Guam
                        - StateCode: HI
                          StateName: Hawaii
                        - StateCode: IA
                          StateName: Iowa
                        - StateCode: ID
                          StateName: Idaho
                        - StateCode: IL
                          StateName: Illinois
                        - StateCode: IN
                          StateName: Indiana
                        - StateCode: KS
                          StateName: Kansas
                        - StateCode: KY
                          StateName: Kentucky
                        - StateCode: LA
                          StateName: Louisiana
                        - StateCode: MA
                          StateName: Massachusetts
                        - StateCode: MD
                          StateName: Maryland
                        - StateCode: ME
                          StateName: Maine
                        - StateCode: MI
                          StateName: Michigan
                        - StateCode: MN
                          StateName: Minnesota
                        - StateCode: MO
                          StateName: Missouri
                        - StateCode: MP
                          StateName: Northern Mariana Islands
                        - StateCode: MS
                          StateName: Mississippi
                        - StateCode: MT
                          StateName: Montana
                        - StateCode: NC
                          StateName: North Carolina
                        - StateCode: ND
                          StateName: North Dakota
                        - StateCode: NE
                          StateName: Nebraska
                        - StateCode: NH
                          StateName: New Hampshire
                        - StateCode: NJ
                          StateName: New Jersey
                        - StateCode: NM
                          StateName: New Mexico
                        - StateCode: NV
                          StateName: Nevada
                        - StateCode: NY
                          StateName: New York
                        - StateCode: OH
                          StateName: Ohio
                        - StateCode: OK
                          StateName: Oklahoma
                        - StateCode: OR
                          StateName: Oregon
                        - StateCode: PA
                          StateName: Pennsylvania
                        - StateCode: PR
                          StateName: Puerto Rico
                        - StateCode: RI
                          StateName: Rhode Island
                        - StateCode: SC
                          StateName: South Carolina
                        - StateCode: SD
                          StateName: South Dakota
                        - StateCode: TN
                          StateName: Tennessee
                        - StateCode: TX
                          StateName: Texas
                        - StateCode: UM
                          StateName: United States Minor Outlying Islands
                        - StateCode: UT
                          StateName: Utah
                        - StateCode: VA
                          StateName: Virginia
                        - StateCode: VI
                          StateName: Virgin Islands, U.S.
                        - StateCode: VT
                          StateName: Vermont
                        - StateCode: WA
                          StateName: Washington
                        - StateCode: WI
                          StateName: Wisconsin
                        - StateCode: WV
                          StateName: West Virginia
                        - StateCode: WY
                          StateName: Wyoming

    GetNAICSResponse:
      type: object
      properties:
        NAICSResponse:
          type: object
          properties:
            NAICSResult:
              type: object
              properties:
                OperationStatus:
                  $ref: "#/components/schemas/xOpStatus"
                NAICS:
                  type: object
                  properties:
                    NAICSDetails:
                      type: array
                      items:
                        type: object
                        properties:
                          Industry:
                            description: >
                              A general taxonomic identification of the business (Accommodation, utilities, etc.)
                            type: string

                  example:
                    NAICS:
                      NAICSDetails:
                        - Industry: Abrasive Product Manufacturing
                        - Industry: Accommodation
                        - Industry: Accommodation and Food Services
                        - Industry: Accounting, Tax Preparation, Bookkeeping, and Payroll Services
                        - Industry: Activities Related to Credit Intermediation
                        - Industry: Activities Related to Real Estate
                        - Industry: Adhesive Manufacturing
                        - Industry: Administration of Air and Water Resource and Solid Waste Management Programs
                        - Industry: Administration of Conservation Programs
                        - Industry: Administration of Economic Program
                        - Industry: Administration of Economic Programs
                        - Industry: Administration of Education Programs
                        - Industry: Administration of Environmental Quality Programs
                        - Industry: Administration of General Economic Programs
                        - Industry: Administration of Housing Programs
                        - Industry: Administration of Housing Programs, Urban Planning, and Community Development
                        - Industry: Administration of Human Resource Programs
                        - Industry: Administration of Human Resource Programs (except Education, Public Health, and Veterans' Affairs Programs)
                        - Industry: Administration of Public Health Programs
                        - Industry: Administration of Urban Planning and Community and Rural Development
                        - Industry: Administration of Veterans' Affairs
                        - Industry: Administrative and Support and Waste Management and Remediation Services
                        - Industry: Administrative and Support Services
                        - Industry: Administrative Management and General Management Consulting Services
                        - Industry: Advertising Agencies
                        - Industry: Advertising Material Distribution Services
                        - Industry: Advertising, Public Relations, and Related Services
                        - Industry: Aerospace Product and Parts Manufacturing
                        - Industry: Agencies, Brokerages, and Other Insurance Related Activities
                        - Industry: Agents and Managers for Artists, Athletes, Entertainers, and Other Public Figures
                        - Industry: Agricultural Implement Manufacturing
                        - Industry: Agriculture, Construction, and Mining Machinery Manufacturing
                        - Industry: Agriculture, Forestry, Fishing and Hunting
                        - Industry: Air and Gas Compressor Manufacturing
                        - Industry: Air Purification Equipment Manufacturing
                        - Industry: Air Traffic Control
                        - Industry: Air Transportation
                        - Industry: Air-Conditioning and Warm Air Heating Equipment and Commercial and Industrial Refrigeration Equipment Manufacturing
                        - Industry: Aircraft Engine and Engine Parts Manufacturing
                        - Industry: Aircraft Manufacturing
                        - Industry: Airport Operations
                        - Industry: Alkalies and Chlorine Manufacturing
                        - Industry: All Other Ambulatory Health Care Services
                        - Industry: All Other Amusement and Recreation Industries
                        - Industry: All Other Animal Production
                        - Industry: All Other Automotive Repair and Maintenance
                        - Industry: All Other Basic Inorganic Chemical Manufacturing
                        - Industry: All Other Basic Organic Chemical Manufacturing
                        - Industry: All Other Business Support Services
                        - Industry: All Other Chemical Product and Preparation Manufacturing
                        - Industry: All Other Consumer Goods Rental
                        - Industry: All Other Converted Paper Product Manufacturing
                        - Industry: All Other Crop Farming
                        - Industry: All Other Cut and Sew Apparel Manufacturing
                        - Industry: All Other Electrical Equipment and Component Manufacturing
                        - Industry: All Other Fabricated Metal Product Manufacturing
                        - Industry: All Other Financial Investment Activities
                        - Industry: All Other Food Manufacturing
                        - Industry: All Other General Merchandise Stores
                        - Industry: All Other General Purpose Machinery Manufacturing
                        - Industry: All Other Grain Farming
                        - Industry: All Other Health and Personal Care Stores
                        - Industry: All Other Home Furnishings Stores
                        - Industry: All Other Industrial Machinery Manufacturing
                        - Industry: All Other Information Services
                        - Industry: All Other Insurance Related Activities
                        - Industry: All Other Leather Good and Allied Product Manufacturing
                        - Industry: All Other Legal Services
                        - Industry: All Other Metal Ore Mining
                        - Industry: All Other Miscellaneous Ambulatory Health Care Services
                        - Industry: All Other Miscellaneous Chemical Product and Preparation Manufacturing
                        - Industry: All Other Miscellaneous Crop Farming
                        - Industry: All Other Miscellaneous Electrical Equipment and Component Manufacturing
                        - Industry: All Other Miscellaneous Fabricated Metal Product Manufacturing
                        - Industry: All Other Miscellaneous Food Manufacturing
                        - Industry: All Other Miscellaneous General Purpose Machinery Manufacturing
                        - Industry: All Other Miscellaneous Manufacturing
                        - Industry: All Other Miscellaneous Nonmetallic Mineral Product Manufacturing
                        - Industry: All Other Miscellaneous Schools and Instruction
                        - Industry: All Other Miscellaneous Store Retailers
                        - Industry: All Other Miscellaneous Store Retailers (except Tobacco Stores)
                        - Industry: All Other Miscellaneous Textile Product Mills
                        - Industry: All Other Miscellaneous Waste Management Services
                        - Industry: All Other Miscellaneous Wood Product Manufacturing
                        - Industry: All Other Motor Vehicle Dealers
                        - Industry: All Other Motor Vehicle Parts Manufacturing
                        - Industry: All Other Nondepository Credit Intermediation
                        - Industry: All Other Nonmetallic Mineral Mining
                        - Industry: All Other Nonmetallic Mineral Product Manufacturing
                        - Industry: All Other Outpatient Care Centers
                        - Industry: All Other Personal Services
                        - Industry: All Other Petroleum and Coal Products Manufacturing
                        - Industry: All Other Pipeline Transportation
                        - Industry: All Other Plastics Product Manufacturing
                        - Industry: All Other Professional, Scientific, and Technical Services
                        - Industry: All Other Publishers
                        - Industry: All Other Rubber Product Manufacturing
                        - Industry: All Other Schools and Instruction
                        - Industry: All Other Specialty Food Stores
                        - Industry: All Other Specialty Trade Contractors
                        - Industry: All Other Support Activities for Transportation
                        - Industry: All Other Support Services
                        - Industry: All Other Telecommunications
                        - Industry: All Other Textile Product Mills
                        - Industry: All Other Transit and Ground Passenger Transportation
                        - Industry: All Other Transportation Equipment Manufacturing
                        - Industry: All Other Travel Arrangement and Reservation Services
                        - Industry: All Other Traveler Accommodation
                        - Industry: All Other Waste Management Services
                        - Industry: All Other Wood Product Manufacturing
                        - Industry: Alumina and Aluminum Production and Processing
                        - Industry: Alumina Refining
                        - Industry: Aluminum Die-Casting Foundries
                        - Industry: Aluminum Extruded Product Manufacturing
                        - Industry: Aluminum Foundries (except Die-Casting)
                        - Industry: Aluminum Sheet, Plate, and Foil Manufacturing
                        - Industry: Ambulance Services
                        - Industry: Ambulatory Health Care Services
                        - Industry: American Indian and Alaska Native Tribal Governments
                        - Industry: Ammunition (except Small Arms) Manufacturing
                        - Industry: Amusement and Theme Parks
                        - Industry: Amusement Arcades
                        - Industry: Amusement Parks and Arcades
                        - Industry: Amusement, Gambling, and Recreation Industries
                        - Industry: Analytical Laboratory Instrument Manufacturing
                        - Industry: Animal (except Poultry) Slaughtering
                        - Industry: Animal Food Manufacturing
                        - Industry: Animal Production
                        - Industry: Animal Slaughtering and Processing
                        - Industry: Anthracite Mining
                        - Industry: Apiculture
                        - Industry: Apparel Accessories and Other Apparel Manufacturing
                        - Industry: Apparel Knitting Mills
                        - Industry: Apparel Manufacturing
                        - Industry: Apparel, Piece Goods, and Notions Merchant Wholesalers
                        - Industry: Apple Orchards
                        - Industry: Appliance Repair and Maintenance
                        - Industry: Appliance, Television, and Other Electronics Stores
                        - Industry: Apprenticeship Training
                        - Industry: Aquaculture
                        - Industry: Architectural and Structural Metals Manufacturing
                        - Industry: Architectural Services
                        - Industry: Architectural, Engineering, and Related Services
                        - Industry: Armored Car Services
                        - Industry: Art Dealers
                        - Industry: Artificial and Synthetic Fibers and Filaments Manufacturing
                        - Industry: Arts, Entertainment, and Recreation
                        - Industry: Asphalt Paving Mixture and Block Manufacturing
                        - Industry: Asphalt Paving, Roofing, and Saturated Materials Manufacturing
                        - Industry: Asphalt Shingle and Coating Materials Manufacturing
                        - Industry: Audio and Video Equipment Manufacturing
                        - Industry: Automatic Environmental Control Manufacturing for Residential, Commercial, and Appliance Use
                        - Industry: Automatic Vending Machine Manufacturing
                        - Industry: Automobile and Light Duty Motor Vehicle Manufacturing
                        - Industry: Automobile and Other Motor Vehicle Merchant Wholesalers
                        - Industry: Automobile Dealers
                        - Industry: Automobile Driving Schools
                        - Industry: Automobile Manufacturing
                        - Industry: Automotive Body, Paint, and Interior Repair and Maintenance
                        - Industry: Automotive Body, Paint, Interior, and Glass Repair
                        - Industry: Automotive Equipment Rental and Leasing
                        - Industry: Automotive Exhaust System Repair
                        - Industry: Automotive Glass Replacement Shops
                        - Industry: Automotive Mechanical and Electrical Repair and Maintenance
                        - Industry: Automotive Oil Change and Lubrication Shops
                        - Industry: Automotive Parts and Accessories Stores
                        - Industry: Automotive Parts, Accessories, and Tire Stores
                        - Industry: Automotive Repair and Maintenance
                        - Industry: Automotive Transmission Repair
                        - Industry: Baked Goods Stores
                        - Industry: Bakeries and Tortilla Manufacturing
                        - Industry: Ball and Roller Bearing Manufacturing
                        - Industry: Barber Shops
                        - Industry: Bare Printed Circuit Board Manufacturing
                        - Industry: Basic Chemical Manufacturing
                        - Industry: Battery Manufacturing
                        - Industry: Beauty Salons
                        - Industry: Bed-and-Breakfast Inns
                        - Industry: Beef Cattle Ranching and Farming
                        - Industry: Beef Cattle Ranching and Farming, including Feedlots
                        - Industry: Beer and Ale Merchant Wholesalers
                        - Industry: Beer, Wine, and Distilled Alcoholic Beverage Merchant Wholesalers
                        - Industry: Beer, Wine, and Liquor Stores
                        - Industry: Beet Sugar Manufacturing
                        - Industry: Berry (except Strawberry) Farming
                        - Industry: Beverage and Tobacco Product Manufacturing
                        - Industry: Beverage Manufacturing
                        - Industry: Biological Product (except Diagnostic) Manufacturing
                        - Industry: Bituminous Coal and Lignite Surface Mining
                        - Industry: Bituminous Coal Underground Mining
                        - Industry: Blankbook, Looseleaf Binders, and Devices Manufacturing
                        - Industry: Blind and Shade Manufacturing
                        - Industry: Blood and Organ Banks
                        - Industry: Boat Building
                        - Industry: Boat Dealers
                        - Industry: Boiler, Tank, and Shipping Container Manufacturing
                        - Industry: Bolt, Nut, Screw, Rivet, and Washer Manufacturing
                        - Industry: Book Publishers
                        - Industry: Book Stores
                        - Industry: Book Stores and News Dealers
                        - Industry: Book, Periodical, and Music Stores
                        - Industry: Book, Periodical, and Newspaper Merchant Wholesalers
                        - Industry: Books Printing
                        - Industry: Bottled Water Manufacturing
                        - Industry: Bowling Centers
                        - Industry: Bread and Bakery Product Manufacturing
                        - Industry: Breakfast Cereal Manufacturing
                        - Industry: Breweries
                        - Industry: Brick and Structural Clay Tile Manufacturing
                        - Industry: Brick, Stone, and Related Construction Material Merchant Wholesalers
                        - Industry: Broadcasting (except Internet)
                        - Industry: Broadwoven Fabric Finishing Mills
                        - Industry: Broadwoven Fabric Mills
                        - Industry: Broilers and Other Meat Type Chicken Production
                        - Industry: Broom, Brush, and Mop Manufacturing
                        - Industry: Building Equipment Contractors
                        - Industry: Building Finishing Contractors
                        - Industry: Building Inspection Services
                        - Industry: Building Material and Garden Equipment and Supplies Dealers
                        - Industry: Building Material and Supplies Dealers
                        - Industry: Burial Casket Manufacturing
                        - Industry: Bus and Other Motor Vehicle Transit Systems
                        - Industry: Business and Secretarial Schools
                        - Industry: Business Associations
                        - Industry: Business Schools and Computer and Management Training
                        - Industry: Business Service Centers
                        - Industry: Business Support Services
                        - Industry: Business to Business Electronic Markets
                        - Industry: Business, Professional, Labor, Political, and Similar Organizations
                        - Industry: Cable and Other Subscription Programming
                        - Industry: Cafeterias, Grill Buffets, and Buffets
                        - Industry: Camera and Photographic Supplies Stores
                        - Industry: Cane Sugar Refining
                        - Industry: Canvas and Related Product Mills
                        - Industry: Car Washes
                        - Industry: Carbon and Graphite Product Manufacturing
                        - Industry: Carbon Black Manufacturing
                        - Industry: Carbon Paper and Inked Ribbon Manufacturing
                        - Industry: Carburetor, Piston, Piston Ring, and Valve Manufacturing
                        - Industry: Carpet and Rug Mills
                        - Industry: Carpet and Upholstery Cleaning Services
                        - Industry: Casino Hotels
                        - Industry: Casinos (except Casino Hotels)
                        - Industry: Caterers
                        - Industry: Cattle Feedlots
                        - Industry: Cattle Ranching and Farming
                        - Industry: Cellulosic Organic Fiber Manufacturing
                        - Industry: Cement and Concrete Product Manufacturing
                        - Industry: Cement Manufacturing
                        - Industry: Cemeteries and Crematories
                        - Industry: Ceramic Wall and Floor Tile Manufacturing
                        - Industry: Charter Bus Industry
                        - Industry: Cheese Manufacturing
                        - Industry: Chemical and Allied Products Merchant Wholesalers
                        - Industry: Chemical Manufacturing
                        - Industry: Chicken Egg Production
                        - Industry: Child and Youth Services
                        - Industry: Child Day Care Services
                        - Industry: Children's and Infants' Clothing Stores
                        - Industry: Chocolate and Confectionery Manufacturing from Cacao Beans
                        - Industry: Cigarette Manufacturing
                        - Industry: Citrus (except Orange) Groves
                        - Industry: Civic and Social Organizations
                        - Industry: Claims Adjusting
                        - Industry: Clay and Ceramic and Refractory Minerals Mining
                        - Industry: Clay Building Material and Refractories Manufacturing
                        - Industry: Clay Product and Refractory Manufacturing
                        - Industry: Clay Refractory Manufacturing
                        - Industry: Clothing Accessories Stores
                        - Industry: Clothing and Clothing Accessories Stores
                        - Industry: Clothing Stores
                        - Industry: Coal and Other Mineral and Ore Merchant Wholesalers
                        - Industry: Coal Mining
                        - Industry: Coastal and Great Lakes Freight Transportation
                        - Industry: Coastal and Great Lakes Passenger Transportation
                        - Industry: Coated and Laminated Packaging Paper Manufacturing
                        - Industry: Coated and Laminated Paper Manufacturing
                        - Industry: Coated Paper Bag and Pouch Manufacturing
                        - Industry: Coating, Engraving, Heat Treating, and Allied Activities
                        - Industry: Coffee and Tea Manufacturing
                        - Industry: Coin-Operated Laundries and Drycleaners
                        - Industry: Collection Agencies
                        - Industry: Colleges, Universities, and Professional Schools
                        - Industry: Commercial Air, Rail, and Water Transportation Equipment Rental and Leasing
                        - Industry: Commercial and Industrial Machinery and Equipment (except Automotive and Electronic) Repair and Maintenance
                        - Industry: Commercial and Industrial Machinery and Equipment Rental and Leasing
                        - Industry: Commercial and Institutional Building Construction
                        - Industry: Commercial and Service Industry Machinery Manufacturing
                        - Industry: Commercial Bakeries
                        - Industry: Commercial Banking
                        - Industry: Commercial Flexographic Printing
                        - Industry: Commercial Gravure Printing
                        - Industry: Commercial Laundry, Drycleaning, and Pressing Machine Manufacturing
                        - Industry: Commercial Lithographic Printing
                        - Industry: Commercial Photography
                        - Industry: Commercial Screen Printing
                        - Industry: Commercial, Industrial, and Institutional Electric Lighting Fixture Manufacturing
                        - Industry: Commodity Contracts Brokerage
                        - Industry: Commodity Contracts Dealing
                        - Industry: Communication and Energy Wire and Cable Manufacturing
                        - Industry: Communication Equipment Repair and Maintenance
                        - Industry: Communications Equipment Manufacturing
                        - Industry: Community Care Facilities for the Elderly
                        - Industry: Community Food and Housing, and Emergency and Other Relief Services
                        - Industry: Community Food Services
                        - Industry: Community Housing Services
                        - Industry: Commuter Rail Systems
                        - Industry: Computer and Computer Peripheral Equipment and Software Merchant Wholesalers
                        - Industry: Computer and Electronic Product Manufacturing
                        - Industry: Computer and Office Machine Repair and Maintenance
                        - Industry: Computer and Peripheral Equipment Manufacturing
                        - Industry: Computer and Software Stores
                        - Industry: Computer Facilities Management Services
                        - Industry: Computer Storage Device Manufacturing
                        - Industry: Computer Systems Design and Related Services
                        - Industry: Computer Systems Design Services
                        - Industry: Computer Terminal Manufacturing
                        - Industry: Computer Training
                        - Industry: Concrete Block and Brick Manufacturing
                        - Industry: Concrete Pipe Manufacturing
                        - Industry: Concrete Pipe, Brick, and Block Manufacturing
                        - Industry: Confectionery and Nut Stores
                        - Industry: Confectionery Manufacturing from Purchased Chocolate
                        - Industry: Confectionery Merchant Wholesalers
                        - Industry: Construction
                        - Industry: Construction and Mining (except Oil Well) Machinery and Equipment Merchant Wholesalers
                        - Industry: Construction Machinery Manufacturing
                        - Industry: Construction of Buildings
                        - Industry: Construction Sand and Gravel Mining
                        - Industry: Construction, Mining, and Forestry Machinery and Equipment Rental and Leasing
                        - Industry: Construction, Transportation, Mining, and Forestry Machinery and Equipment Rental and Leasing
                        - Industry: Consumer Electronics and Appliances Rental
                        - Industry: Consumer Electronics Repair and Maintenance
                        - Industry: Consumer Goods Rental
                        - Industry: Consumer Lending
                        - Industry: Continuing Care Retirement Communities
                        - Industry: Convenience Stores
                        - Industry: Convention and Trade Show Organizers
                        - Industry: Convention and Visitors Bureaus
                        - Industry: Converted Paper Product Manufacturing
                        - Industry: Conveyor and Conveying Equipment Manufacturing
                        - Industry: Cookie and Cracker Manufacturing
                        - Industry: Cookie, Cracker, and Pasta Manufacturing
                        - Industry: Copper Foundries (except Die-Casting)
                        - Industry: Copper Ore and Nickel Ore Mining
                        - Industry: Copper Rolling, Drawing, and Extruding
                        - Industry: Copper Rolling, Drawing, Extruding, and Alloying
                        - Industry: Copper Wire (except Mechanical) Drawing
                        - Industry: Copper, Nickel, Lead, and Zinc Mining
                        - Industry: Corn Farming
                        - Industry: Corporate, Subsidiary, and Regional Managing Offices
                        - Industry: Correctional Institutions
                        - Industry: Corrugated and Solid Fiber Box Manufacturing
                        - Industry: Cosmetics, Beauty Supplies, and Perfume Stores
                        - Industry: Cosmetology and Barber Schools
                        - Industry: Costume Jewelry and Novelty Manufacturing
                        - Industry: Cotton Farming
                        - Industry: Cotton Ginning
                        - Industry: Couriers and Express Delivery Services
                        - Industry: Couriers and Messengers
                        - Industry: Court Reporting and Stenotype Services
                        - Industry: Courts
                        - Industry: Creamery Butter Manufacturing
                        - Industry: Credit Bureaus
                        - Industry: Credit Card Issuing
                        - Industry: Credit Intermediation and Related Activities
                        - Industry: Credit Unions
                        - Industry: Crop Harvesting, Primarily by Machine
                        - Industry: Crop Production
                        - Industry: Crown and Closure Manufacturing
                        - Industry: Crude Petroleum and Natural Gas Extraction
                        - Industry: Crushed and Broken Granite Mining and Quarrying
                        - Industry: Crushed and Broken Limestone Mining and Quarrying
                        - Industry: Current-Carrying Wiring Device Manufacturing
                        - Industry: Curtain and Drapery Mills
                        - Industry: Curtain and Linen Mills
                        - Industry: Custom Architectural Woodwork and Millwork Manufacturing
                        - Industry: Custom Compounding of Purchased Resins
                        - Industry: Custom Computer Programming Services
                        - Industry: Custom Roll Forming
                        - Industry: Cut and Sew Apparel Contractors
                        - Industry: Cut and Sew Apparel Manufacturing
                        - Industry: Cut Stock, Resawing Lumber, and Planing
                        - Industry: Cut Stone and Stone Product Manufacturing
                        - Industry: Cutlery and Flatware (except Precious) Manufacturing
                        - Industry: Cutlery and Handtool Manufacturing
                        - Industry: Cutting Tool and Machine Tool Accessory Manufacturing
                        - Industry: Cyclic Crude and Intermediate Manufacturing
                        - Industry: Dairy Cattle and Milk Production
                        - Industry: Dairy Product (except Dried or Canned) Merchant Wholesalers
                        - Industry: Dairy Product (except Frozen) Manufacturing
                        - Industry: Dairy Product Manufacturing
                        - Industry: Dance Companies
                        - Industry: Data Processing, Hosting and Related Services
                        - Industry: Death Care Services
                        - Industry: Deep Sea Freight Transportation
                        - Industry: Deep Sea Passenger Transportation
                        - Industry: Deep Sea, Coastal, and Great Lakes Water Transportation
                        - Industry: Dental Equipment and Supplies Manufacturing
                        - Industry: Dental Laboratories
                        - Industry: Department Stores
                        - Industry: Department Stores (except Discount Department Stores)
                        - Industry: Depository Credit Intermediation
                        - Industry: Diagnostic Imaging Centers
                        - Industry: Die-Cut Paper and Paperboard Office Supplies Manufacturing
                        - Industry: Diet and Weight Reducing Centers
                        - Industry: Digital Printing
                        - Industry: Dimension Stone Mining and Quarrying
                        - Industry: Direct Health and Medical Insurance Carriers
                        - Industry: Direct Insurance (except Life, Health, and Medical) Carriers
                        - Industry: Direct Life Insurance Carriers
                        - Industry: Direct Life, Health, and Medical Insurance Carriers
                        - Industry: Direct Mail Advertising
                        - Industry: Direct Property and Casualty Insurance Carriers
                        - Industry: Direct Selling Establishments
                        - Industry: Direct Title Insurance Carriers
                        - Industry: Directory and Mailing List Publishers
                        - Industry: Discount Department Stores
                        - Industry: Display Advertising
                        - Industry: Distilleries
                        - Industry: Document Preparation Services
                        - Industry: Dog and Cat Food Manufacturing
                        - Industry: Doll and Stuffed Toy Manufacturing
                        - Industry: Doll, Toy, and Game Manufacturing
                        - Industry: Drafting Services
                        - Industry: Dried and Dehydrated Food Manufacturing
                        - Industry: Drilling Oil and Gas Wells
                        - Industry: Drinking Places (Alcoholic Beverages)
                        - Industry: Drive-In Motion Picture Theaters
                        - Industry: Drugs and Druggists' Sundries Merchant Wholesalers
                        - Industry: Dry Pasta Manufacturing
                        - Industry: Dry Pea and Bean Farming
                        - Industry: Dry, Condensed, and Evaporated Dairy Product Manufacturing
                        - Industry: Drycleaning and Laundry Services
                        - Industry: Drycleaning and Laundry Services (except Coin-Operated)
                        - Industry: Drywall and Insulation Contractors
                        - Industry: Dual-Purpose Cattle Ranching and Farming
                        - Industry: Educational Services
                        - Industry: Educational Support Services
                        - Industry: Electric Bulk Power Transmission and Control
                        - Industry: Electric Housewares and Household Fan Manufacturing
                        - Industry: Electric Lamp Bulb and Part Manufacturing
                        - Industry: Electric Lighting Equipment Manufacturing
                        - Industry: Electric Power Distribution
                        - Industry: Electric Power Generation
                        - Industry: Electric Power Generation, Transmission and Distribution
                        - Industry: Electric Power Transmission, Control, and Distribution
                        - Industry: Electrical and Electronic Appliance, Television, and Radio Set Merchant Wholesalers
                        - Industry: Electrical and Electronic Goods Merchant Wholesalers
                        - Industry: Electrical Apparatus and Equipment, Wiring Supplies, and Related Equipment Merchant Wholesalers
                        - Industry: Electrical Contractors and Other Wiring Installation Contractors
                        - Industry: Electrical Equipment Manufacturing
                        - Industry: Electrical Equipment, Appliance, and Component Manufacturing
                        - Industry: Electromedical and Electrotherapeutic Apparatus Manufacturing
                        - Industry: Electrometallurgical Ferroalloy Product Manufacturing
                        - Industry: Electron Tube Manufacturing
                        - Industry: Electronic and Precision Equipment Repair and Maintenance
                        - Industry: Electronic Auctions
                        - Industry: Electronic Capacitor Manufacturing
                        - Industry: Electronic Coil, Transformer, and Other Inductor Manufacturing
                        - Industry: Electronic Computer Manufacturing
                        - Industry: Electronic Connector Manufacturing
                        - Industry: Electronic Resistor Manufacturing
                        - Industry: Electronic Shopping
                        - Industry: Electronic Shopping and Mail-Order Houses
                        - Industry: Electronics and Appliance Stores
                        - Industry: Electroplating, Plating, Polishing, Anodizing, and Coloring
                        - Industry: Elementary and Secondary Schools
                        - Industry: Elevator and Moving Stairway Manufacturing
                        - Industry: Emergency and Other Relief Services
                        - Industry: Employment Placement Agencies
                        - Industry: Employment Placement Agencies and Executive Search Services
                        - Industry: Employment Services
                        - Industry: Enameled Iron and Metal Sanitary Ware Manufacturing
                        - Industry: Engine, Turbine, and Power Transmission Equipment Manufacturing
                        - Industry: Engineered Wood Member (except Truss) Manufacturing
                        - Industry: Engineering Services
                        - Industry: Envelope Manufacturing
                        - Industry: Environment, Conservation and Wildlife Organizations
                        - Industry: Environmental Consulting Services
                        - Industry: Ethyl Alcohol Manufacturing
                        - Industry: Exam Preparation and Tutoring
                        - Industry: Executive and Legislative Offices, Combined
                        - Industry: Executive Offices
                        - Industry: Executive Search Services
                        - Industry: Executive, Legislative, and Other General Government Support
                        - Industry: Explosives Manufacturing
                        - Industry: Exterminating and Pest Control Services
                        - Industry: Fabric Coating Mills
                        - Industry: Fabric Mills
                        - Industry: Fabricated Metal Product Manufacturing
                        - Industry: Fabricated Pipe and Pipe Fitting Manufacturing
                        - Industry: Fabricated Structural Metal Manufacturing
                        - Industry: Facilities Support Services
                        - Industry: Family Clothing Stores
                        - Industry: Family Planning Centers
                        - Industry: Farm and Garden Machinery and Equipment Merchant Wholesalers
                        - Industry: Farm Labor Contractors and Crew Leaders
                        - Industry: Farm Machinery and Equipment Manufacturing
                        - Industry: Farm Management Services
                        - Industry: Farm Product Raw Material Merchant Wholesalers
                        - Industry: Farm Product Warehousing and Storage
                        - Industry: Farm Supplies Merchant Wholesalers
                        - Industry: Fastener, Button, Needle, and Pin Manufacturing
                        - Industry: Fats and Oils Refining and Blending
                        - Industry: Ferrous Metal Foundries
                        - Industry: Fertilizer (Mixing Only) Manufacturing
                        - Industry: Fertilizer Manufacturing
                        - Industry: Fiber Can, Tube, Drum, and Similar Products Manufacturing
                        - Industry: Fiber Optic Cable Manufacturing
                        - Industry: Fiber, Yarn, and Thread Mills
                        - Industry: Finance and Insurance
                        - Industry: Financial Transactions Processing, Reserve, and Clearinghouse Activities
                        - Industry: Fine Arts Schools
                        - Industry: Finfish Farming and Fish Hatcheries
                        - Industry: Finfish Fishing
                        - Industry: Finish Carpentry Contractors
                        - Industry: Fire Protection
                        - Industry: Fish and Seafood Markets
                        - Industry: Fish and Seafood Merchant Wholesalers
                        - Industry: Fishing
                        - Industry: Fishing, Hunting and Trapping
                        - Industry: Fitness and Recreational Sports Centers
                        - Industry: Flat Glass Manufacturing
                        - Industry: Flavoring Syrup and Concentrate Manufacturing
                        - Industry: Flight Training
                        - Industry: Floor Covering Stores
                        - Industry: Flooring Contractors
                        - Industry: Floriculture Production
                        - Industry: Florists
                        - Industry: Flour Milling
                        - Industry: Flour Milling and Malt Manufacturing
                        - Industry: Flour Mixes and Dough Manufacturing from Purchased Flour
                        - Industry: Flower, Nursery Stock, and Florists' Supplies Merchant Wholesalers
                        - Industry: Fluid Milk Manufacturing
                        - Industry: Fluid Power Cylinder and Actuator Manufacturing
                        - Industry: Fluid Power Pump and Motor Manufacturing
                        - Industry: Fluid Power Valve and Hose Fitting Manufacturing
                        - Industry: Folding Paperboard Box Manufacturing
                        - Industry: Food (Health) Supplement Stores
                        - Industry: Food and Beverage Stores
                        - Industry: Food Crops Grown Under Cover
                        - Industry: Food Manufacturing
                        - Industry: Food Product Machinery Manufacturing
                        - Industry: Food Service Contractors
                        - Industry: Food Services and Drinking Places
                        - Industry: Footwear and Leather Goods Repair
                        - Industry: Footwear Manufacturing
                        - Industry: Footwear Merchant Wholesalers
                        - Industry: Forest Nurseries and Gathering of Forest Products
                        - Industry: Forestry and Logging
                        - Industry: Forging and Stamping
                        - Industry: Formal Wear and Costume Rental
                        - Industry: Fossil Fuel Electric Power Generation
                        - Industry: Foundation, Structure, and Building Exterior Contractors
                        - Industry: Foundries
                        - Industry: Framing Contractors
                        - Industry: Freestanding Ambulatory Surgical and Emergency Centers
                        - Industry: Freight Transportation Arrangement
                        - Industry: Fresh and Frozen Seafood Processing
                        - Industry: Fresh Fruit and Vegetable Merchant Wholesalers
                        - Industry: Frozen Cakes, Pies, and Other Pastries Manufacturing
                        - Industry: Frozen Food Manufacturing
                        - Industry: Frozen Fruit, Juice, and Vegetable Manufacturing
                        - Industry: Frozen Specialty Food Manufacturing
                        - Industry: Fruit and Tree Nut Combination Farming
                        - Industry: Fruit and Tree Nut Farming
                        - Industry: Fruit and Vegetable Canning
                        - Industry: Fruit and Vegetable Canning, Pickling, and Drying
                        - Industry: Fruit and Vegetable Markets
                        - Industry: Fruit and Vegetable Preserving and Specialty Food Manufacturing
                        - Industry: Fuel Dealers
                        - Industry: Full-Service Restaurants
                        - Industry: Funds, Trusts, and Other Financial Vehicles
                        - Industry: Funeral Homes and Funeral Services
                        - Industry: Fur and Leather Apparel Manufacturing
                        - Industry: Fur-Bearing Animal and Rabbit Production
                        - Industry: Furniture and Home Furnishing Merchant Wholesalers
                        - Industry: Furniture and Home Furnishings Stores
                        - Industry: Furniture and Related Product Manufacturing
                        - Industry: Furniture Merchant Wholesalers
                        - Industry: Furniture Stores
                        - Industry: Gambling Industries
                        - Industry: Game, Toy, and Children's Vehicle Manufacturing
                        - Industry: Gasket, Packing, and Sealing Device Manufacturing
                        - Industry: Gasoline Engine and Engine Parts Manufacturing
                        - Industry: Gasoline Stations
                        - Industry: Gasoline Stations with Convenience Stores
                        - Industry: General Automotive Repair
                        - Industry: General Freight Trucking
                        - Industry: General Freight Trucking, Local
                        - Industry: General Freight Trucking, Long-Distance
                        - Industry: General Freight Trucking, Long-Distance, Less Than Truckload
                        - Industry: General Freight Trucking, Long-Distance, Truckload
                        - Industry: General Line Grocery Merchant Wholesalers
                        - Industry: General Medical and Surgical Hospitals
                        - Industry: General Merchandise Stores
                        - Industry: General Rental Centers
                        - Industry: General Warehousing and Storage
                        - Industry: Geophysical Surveying and Mapping Services
                        - Industry: Gift, Novelty, and Souvenir Stores
                        - Industry: Glass and Glass Product Manufacturing
                        - Industry: Glass and Glazing Contractors
                        - Industry: Glass Container Manufacturing
                        - Industry: Glass Product Manufacturing Made of Purchased Glass
                        - Industry: Glove and Mitten Manufacturing
                        - Industry: Goat Farming
                        - Industry: Gold Ore and Silver Ore Mining
                        - Industry: Gold Ore Mining
                        - Industry: Golf Courses and Country Clubs
                        - Industry: Grain and Field Bean Merchant Wholesalers
                        - Industry: Grain and Oilseed Milling
                        - Industry: Grantmaking and Giving Services
                        - Industry: Grantmaking Foundations
                        - Industry: Grape Vineyards
                        - Industry: Graphic Design Services
                        - Industry: Greenhouse, Nursery, and Floriculture Production
                        - Industry: Greeting Card Publishers
                        - Industry: Grocery and Related Product Merchant Wholesalers
                        - Industry: Grocery Stores
                        - Industry: Ground or Treated Mineral and Earth Manufacturing
                        - Industry: Guided Missile and Space Vehicle Manufacturing
                        - Industry: Guided Missile and Space Vehicle Propulsion Unit and Propulsion Unit Parts Manufacturing
                        - Industry: Gum and Wood Chemical Manufacturing
                        - Industry: Gypsum Product Manufacturing
                        - Industry: Hair, Nail, and Skin Care Services
                        - Industry: Hand and Edge Tool Manufacturing
                        - Industry: Hardware Manufacturing
                        - Industry: Hardware Merchant Wholesalers
                        - Industry: Hardware Stores
                        - Industry: Hardware, and Plumbing and Heating Equipment and Supplies Merchant Wholesalers
                        - Industry: Hardwood Veneer and Plywood Manufacturing
                        - Industry: Hat, Cap, and Millinery Manufacturing
                        - Industry: Hay Farming
                        - Industry: Hazardous Waste Collection
                        - Industry: Hazardous Waste Treatment and Disposal
                        - Industry: Health and Personal Care Stores
                        - Industry: Health and Welfare Funds
                        - Industry: Health Care and Social Assistance
                        - Industry: Heating Equipment (except Warm Air Furnaces) Manufacturing
                        - Industry: Heating Oil Dealers
                        - Industry: Heavy and Civil Engineering Construction
                        - Industry: Heavy Duty Truck Manufacturing
                        - Industry: Highway, Street, and Bridge Construction
                        - Industry: Historical Sites
                        - Industry: HMO Medical Centers
                        - Industry: Hobby, Toy, and Game Stores
                        - Industry: Hog and Pig Farming
                        - Industry: Home and Garden Equipment and Appliance Repair and Maintenance
                        - Industry: Home and Garden Equipment Repair and Maintenance
                        - Industry: Home Centers
                        - Industry: Home Furnishing Merchant Wholesalers
                        - Industry: Home Furnishings Stores
                        - Industry: Home Health Care Services
                        - Industry: Home Health Equipment Rental
                        - Industry: Homes for the Elderly
                        - Industry: Horses and Other Equine Production
                        - Industry: Hosiery and Sock Mills
                        - Industry: Hospitals
                        - Industry: Hotels (except Casino Hotels) and Motels
                        - Industry: House Slipper Manufacturing
                        - Industry: Household and Institutional Furniture and Kitchen Cabinet Manufacturing
                        - Industry: Household and Institutional Furniture Manufacturing
                        - Industry: Household Appliance Manufacturing
                        - Industry: Household Appliance Stores
                        - Industry: Household Cooking Appliance Manufacturing
                        - Industry: Household Furniture (except Wood and Metal) Manufacturing
                        - Industry: Household Laundry Equipment Manufacturing
                        - Industry: Household Refrigerator and Home Freezer Manufacturing
                        - Industry: Household Vacuum Cleaner Manufacturing
                        - Industry: Human Resources Consulting Services
                        - Industry: Human Rights Organizations
                        - Industry: Hunting and Trapping
                        - Industry: Hydroelectric Power Generation
                        - Industry: Ice Cream and Frozen Dessert Manufacturing
                        - Industry: Ice Manufacturing
                        - Industry: Independent Artists, Writers, and Performers
                        - Industry: Individual and Family Services
                        - Industry: Industrial and Commercial Fan and Blower Manufacturing
                        - Industry: Industrial and Personal Service Paper Merchant Wholesalers
                        - Industry: Industrial Building Construction
                        - Industry: Industrial Design Services
                        - Industry: Industrial Gas Manufacturing
                        - Industry: Industrial Launderers
                        - Industry: Industrial Machinery and Equipment Merchant Wholesalers
                        - Industry: Industrial Machinery Manufacturing
                        - Industry: Industrial Mold Manufacturing
                        - Industry: Industrial Pattern Manufacturing
                        - Industry: Industrial Process Furnace and Oven Manufacturing
                        - Industry: Industrial Sand Mining
                        - Industry: Industrial Supplies Merchant Wholesalers
                        - Industry: Industrial Truck, Tractor, Trailer, and Stacker Machinery Manufacturing
                        - Industry: Industrial Valve Manufacturing
                        - Industry: Infants' Cut and Sew Apparel Manufacturing
                        - Industry: Information
                        - Industry: Inland Water Freight Transportation
                        - Industry: Inland Water Passenger Transportation
                        - Industry: Inland Water Transportation
                        - Industry: Inorganic Dye and Pigment Manufacturing
                        - Industry: Institutional Furniture Manufacturing
                        - Industry: Instrument Manufacturing for Measuring and Testing Electricity and Electrical Signals
                        - Industry: Instruments and Related Products Manufacturing for Measuring, Displaying, and Controlling Industrial Process Variables
                        - Industry: Insurance Agencies and Brokerages
                        - Industry: Insurance and Employee Benefit Funds
                        - Industry: Insurance Carriers
                        - Industry: Insurance Carriers and Related Activities
                        - Industry: Integrated Record Production/Distribution
                        - Industry: Interior Design Services
                        - Industry: International Affairs
                        - Industry: International Trade Financing
                        - Industry: Internet Publishing and Broadcasting and Web Search Portals
                        - Industry: Interurban and Rural Bus Transportation
                        - Industry: Investigation and Security Services
                        - Industry: Investigation Services
                        - Industry: Investigation, Guard, and Armored Car Services
                        - Industry: Investment Advice
                        - Industry: Investment Banking and Securities Dealing
                        - Industry: In-Vitro Diagnostic Substance Manufacturing
                        - Industry: Iron and Steel Forging
                        - Industry: Iron and Steel Mills
                        - Industry: Iron and Steel Mills and Ferroalloy Manufacturing
                        - Industry: Iron and Steel Pipe and Tube Manufacturing from Purchased Steel
                        - Industry: Iron Foundries
                        - Industry: Iron Ore Mining
                        - Industry: Irradiation Apparatus Manufacturing
                        - Industry: Janitorial Services
                        - Industry: Jewelers' Material and Lapidary Work Manufacturing
                        - Industry: Jewelry (except Costume) Manufacturing
                        - Industry: Jewelry and Silverware Manufacturing
                        - Industry: Jewelry Stores
                        - Industry: Jewelry, Luggage, and Leather Goods Stores
                        - Industry: Jewelry, Watch, Precious Stone, and Precious Metal Merchant Wholesalers
                        - Industry: Junior Colleges
                        - Industry: Justice, Public Order, and Safety Activities
                        - Industry: Kaolin and Ball Clay Mining
                        - Industry: Kidney Dialysis Centers
                        - Industry: Kitchen Utensil, Pot, and Pan Manufacturing
                        - Industry: Knit Fabric Mills
                        - Industry: Labor Unions and Similar Labor Organizations
                        - Industry: Laminated Aluminum Foil Manufacturing for Flexible Packaging Uses
                        - Industry: Laminated Plastics Plate, Sheet (except Packaging), and Shape Manufacturing
                        - Industry: Land Subdivision
                        - Industry: Landscape Architectural Services
                        - Industry: Landscaping Services
                        - Industry: Language Schools
                        - Industry: Lawn and Garden Equipment and Supplies Stores
                        - Industry: Lawn and Garden Tractor and Home Lawn and Garden Equipment Manufacturing
                        - Industry: Lead Ore and Zinc Ore Mining
                        - Industry: Lead Pencil and Art Good Manufacturing
                        - Industry: Leather and Allied Product Manufacturing
                        - Industry: Leather and Hide Tanning and Finishing
                        - Industry: Legal Counsel and Prosecution
                        - Industry: Legal Services
                        - Industry: Legislative Bodies
                        - Industry: Lessors of Miniwarehouses and Self-Storage Units
                        - Industry: Lessors of Nonfinancial Intangible Assets (except Copyrighted Works)
                        - Industry: Lessors of Nonresidential Buildings (except Miniwarehouses)
                        - Industry: Lessors of Other Real Estate Property
                        - Industry: Lessors of Real Estate
                        - Industry: Lessors of Residential Buildings and Dwellings
                        - Industry: Libraries and Archives
                        - Industry: Light Truck and Utility Vehicle Manufacturing
                        - Industry: Lighting Fixture Manufacturing
                        - Industry: Lime and Gypsum Product Manufacturing
                        - Industry: Lime Manufacturing
                        - Industry: Limited-Service Eating Places
                        - Industry: Limited-Service Restaurants
                        - Industry: Limousine Service
                        - Industry: Line-Haul Railroads
                        - Industry: Linen and Uniform Supply
                        - Industry: Linen Supply
                        - Industry: Liquefied Petroleum Gas (Bottled Gas) Dealers
                        - Industry: Livestock Merchant Wholesalers
                        - Industry: Local Messengers and Local Delivery
                        - Industry: Locksmiths
                        - Industry: Logging
                        - Industry: Luggage and Leather Goods Stores
                        - Industry: Luggage Manufacturing
                        - Industry: Lumber and Other Construction Materials Merchant Wholesalers
                        - Industry: Lumber, Plywood, Millwork, and Wood Panel Merchant Wholesalers
                        - Industry: Machine Shops
                        - Industry: Machine Shops; Turned Product; and Screw, Nut, and Bolt Manufacturing
                        - Industry: Machine Tool (Metal Cutting Types) Manufacturing
                        - Industry: Machine Tool (Metal Forming Types) Manufacturing
                        - Industry: Machinery Manufacturing
                        - Industry: Machinery, Equipment, and Supplies Merchant Wholesalers
                        - Industry: Magnetic and Optical Recording Media Manufacturing
                        - Industry: Mail-Order Houses
                        - Industry: Major Appliance Manufacturing
                        - Industry: Malt Manufacturing
                        - Industry: Management Consulting Services
                        - Industry: Management of Companies and Enterprises
                        - Industry: Management, Scientific, and Technical Consulting Services
                        - Industry: Manifold Business Forms Printing
                        - Industry: Manufactured (Mobile) Home Dealers
                        - Industry: Manufactured Home (Mobile Home) Manufacturing
                        - Industry: Manufacturing
                        - Industry: Manufacturing and Reproducing Magnetic and Optical Media
                        - Industry: Marinas
                        - Industry: Marine Cargo Handling
                        - Industry: Marketing Consulting Services
                        - Industry: Marketing Research and Public Opinion Polling
                        - Industry: Marking Device Manufacturing
                        - Industry: Masonry Contractors
                        - Industry: Material Handling Equipment Manufacturing
                        - Industry: Materials Recovery Facilities
                        - Industry: Mattress Manufacturing
                        - Industry: Mayonnaise, Dressing, and Other Prepared Sauce Manufacturing
                        - Industry: Measuring and Dispensing Pump Manufacturing
                        - Industry: Meat and Meat Product Merchant Wholesalers
                        - Industry: Meat Markets
                        - Industry: Meat Processed from Carcasses
                        - Industry: Mechanical Power Transmission Equipment Manufacturing
                        - Industry: Media Buying Agencies
                        - Industry: Media Representatives
                        - Industry: Medical and Diagnostic Laboratories
                        - Industry: Medical Equipment and Supplies Manufacturing
                        - Industry: Medical Laboratories
                        - Industry: Medical, Dental, and Hospital Equipment and Supplies Merchant Wholesalers
                        - Industry: Medicinal and Botanical Manufacturing
                        - Industry: Men's and Boys' Clothing and Furnishings Merchant Wholesalers
                        - Industry: Men's and Boys' Cut and Sew Apparel Contractors
                        - Industry: Men's and Boys' Cut and Sew Apparel Manufacturing
                        - Industry: Men's and Boys' Cut and Sew Other Outerwear Manufacturing
                        - Industry: Men's and Boys' Cut and Sew Shirt (except Work Shirt) Manufacturing
                        - Industry: Men's and Boys' Cut and Sew Suit, Coat, and Overcoat Manufacturing
                        - Industry: Men's and Boys' Cut and Sew Trouser, Slack, and Jean Manufacturing
                        - Industry: Men's and Boys' Cut and Sew Underwear and Nightwear Manufacturing
                        - Industry: Men's and Boys' Cut and Sew Work Clothing Manufacturing
                        - Industry: Men's and Boys' Neckwear Manufacturing
                        - Industry: Men's Clothing Stores
                        - Industry: Men's Footwear (except Athletic) Manufacturing
                        - Industry: Merchant Wholesalers, Durable Goods
                        - Industry: Merchant Wholesalers, Nondurable Goods
                        - Industry: Metal and Mineral (except Petroleum) Merchant Wholesalers
                        - Industry: Metal Can Manufacturing
                        - Industry: Metal Can, Box, and Other Metal Container (Light Gauge) Manufacturing
                        - Industry: Metal Coating, Engraving (except Jewelry and Silverware), and Allied Services to Manufacturers
                        - Industry: Metal Heat Treating
                        - Industry: Metal Household Furniture Manufacturing
                        - Industry: Metal Ore Mining
                        - Industry: Metal Service Centers and Other Metal Merchant Wholesalers
                        - Industry: Metal Stamping
                        - Industry: Metal Tank (Heavy Gauge) Manufacturing
                        - Industry: Metal Valve Manufacturing
                        - Industry: Metal Window and Door Manufacturing
                        - Industry: Metalworking Machinery Manufacturing
                        - Industry: Military Armored Vehicle, Tank, and Tank Component Manufacturing
                        - Industry: Millwork
                        - Industry: Mineral Wool Manufacturing
                        - Industry: Mining (except Oil and Gas)
                        - Industry: Mining and Oil and Gas Field Machinery Manufacturing
                        - Industry: Mining Machinery and Equipment Manufacturing
                        - Industry: Mining, Quarrying, and Oil and Gas Extraction
                        - Industry: Miscellaneous Durable Goods Merchant Wholesalers
                        - Industry: Miscellaneous Financial Investment Activities
                        - Industry: Miscellaneous Intermediation
                        - Industry: Miscellaneous Manufacturing
                        - Industry: Miscellaneous Nondurable Goods Merchant Wholesalers
                        - Industry: Miscellaneous Store Retailers
                        - Industry: Mixed Mode Transit Systems
                        - Industry: Mobile Food Services
                        - Industry: Monetary Authorities-Central Bank
                        - Industry: Mortgage and Nonmortgage Loan Brokers
                        - Industry: Motion Picture and Sound Recording Industries
                        - Industry: Motion Picture and Video Distribution
                        - Industry: Motion Picture and Video Exhibition
                        - Industry: Motion Picture and Video Industries
                        - Industry: Motion Picture and Video Production
                        - Industry: Motion Picture Theaters (except Drive-Ins)
                        - Industry: Motor and Generator Manufacturing
                        - Industry: Motor Home Manufacturing
                        - Industry: Motor Vehicle Air-Conditioning Manufacturing
                        - Industry: Motor Vehicle and Motor Vehicle Parts and Supplies Merchant Wholesalers
                        - Industry: Motor Vehicle and Parts Dealers
                        - Industry: Motor Vehicle Body and Trailer Manufacturing
                        - Industry: Motor Vehicle Body Manufacturing
                        - Industry: Motor Vehicle Brake System Manufacturing
                        - Industry: Motor Vehicle Electrical and Electronic Equipment Manufacturing
                        - Industry: Motor Vehicle Gasoline Engine and Engine Parts Manufacturing
                        - Industry: Motor Vehicle Manufacturing
                        - Industry: Motor Vehicle Metal Stamping
                        - Industry: Motor Vehicle Parts (Used) Merchant Wholesalers
                        - Industry: Motor Vehicle Parts Manufacturing
                        - Industry: Motor Vehicle Seating and Interior Trim Manufacturing
                        - Industry: Motor Vehicle Steering and Suspension Components (except Spring) Manufacturing
                        - Industry: Motor Vehicle Supplies and New Parts Merchant Wholesalers
                        - Industry: Motor Vehicle Towing
                        - Industry: Motor Vehicle Transmission and Power Train Parts Manufacturing
                        - Industry: Motorcycle, ATV, and Personal Watercraft Dealers
                        - Industry: Motorcycle, Bicycle, and Parts Manufacturing
                        - Industry: Motorcycle, Boat, and Other Motor Vehicle Dealers
                        - Industry: Museums
                        - Industry: Museums, Historical Sites, and Similar Institutions
                        - Industry: Mushroom Production
                        - Industry: Music Publishers
                        - Industry: Musical Groups and Artists
                        - Industry: Musical Instrument and Supplies Stores
                        - Industry: Musical Instrument Manufacturing
                        - Industry: Nail Salons
                        - Industry: Narrow Fabric Mills
                        - Industry: Narrow Fabric Mills and Schiffli Machine Embroidery
                        - Industry: National Security
                        - Industry: National Security and International Affairs
                        - Industry: Natural Gas Distribution
                        - Industry: Natural Gas Liquid Extraction
                        - Industry: Nature Parks and Other Similar Institutions
                        - Industry: Navigational Services to Shipping
                        - Industry: Navigational, Measuring, Electromedical, and Control Instruments Manufacturing
                        - Industry: New Car Dealers
                        - Industry: New Housing Operative Builders
                        - Industry: New Multifamily Housing Construction (except Operative Builders)
                        - Industry: New Single-Family Housing Construction (except Operative Builders)
                        - Industry: News Dealers and Newsstands
                        - Industry: News Syndicates
                        - Industry: Newspaper Publishers
                        - Industry: Newspaper, Periodical, Book, and Directory Publishers
                        - Industry: Newsprint Mills
                        - Industry: Nitrogenous Fertilizer Manufacturing
                        - Industry: Noncellulosic Organic Fiber Manufacturing
                        - Industry: Nonchocolate Confectionery Manufacturing
                        - Industry: Noncitrus Fruit and Tree Nut Farming
                        - Industry: Nonclay Refractory Manufacturing
                        - Industry: Noncurrent-Carrying Wiring Device Manufacturing
                        - Industry: Nondepository Credit Intermediation
                        - Industry: Nonferrous (except Aluminum) Die-Casting Foundries
                        - Industry: Nonferrous Forging
                        - Industry: Nonferrous Metal (except Aluminum) Production and Processing
                        - Industry: Nonferrous Metal (except Aluminum) Smelting and Refining
                        - Industry: Nonferrous Metal (except Copper and Aluminum) Rolling, Drawing, and Extruding
                        - Industry: Nonferrous Metal (except Copper and Aluminum) Rolling, Drawing, Extruding, and Alloying
                        - Industry: Nonferrous Metal Foundries
                        - Industry: Nonfolding Sanitary Food Container Manufacturing
                        - Industry: Nonmetallic Mineral Mining and Quarrying
                        - Industry: Nonmetallic Mineral Product Manufacturing
                        - Industry: Nonresidential Building Construction
                        - Industry: Nonresidential Property Managers
                        - Industry: Nonscheduled Air Transportation
                        - Industry: Nonscheduled Chartered Freight Air Transportation
                        - Industry: Nonscheduled Chartered Passenger Air Transportation
                        - Industry: Nonstore Retailers
                        - Industry: Nonupholstered Wood Household Furniture Manufacturing
                        - Industry: Nonwoven Fabric Mills
                        - Industry: Nuclear Electric Power Generation
                        - Industry: Nursery and Floriculture Production
                        - Industry: Nursery and Tree Production
                        - Industry: Nursery, Garden Center, and Farm Supply Stores
                        - Industry: Nursing and Residential Care Facilities
                        - Industry: Nursing Care Facilities
                        - Industry: Office Administrative Services
                        - Industry: Office Equipment Merchant Wholesalers
                        - Industry: Office Furniture (except Wood) Manufacturing
                        - Industry: Office Furniture (including Fixtures) Manufacturing
                        - Industry: Office Machinery and Equipment Rental and Leasing
                        - Industry: Office Machinery Manufacturing
                        - Industry: Office Supplies (except Paper) Manufacturing
                        - Industry: Office Supplies and Stationery Stores
                        - Industry: Office Supplies, Stationery, and Gift Stores
                        - Industry: Offices of All Other Health Practitioners
                        - Industry: Offices of All Other Miscellaneous Health Practitioners
                        - Industry: Offices of Bank Holding Companies
                        - Industry: Offices of Certified Public Accountants
                        - Industry: Offices of Chiropractors
                        - Industry: Offices of Dentists
                        - Industry: Offices of Lawyers
                        - Industry: Offices of Mental Health Practitioners (except Physicians)
                        - Industry: Offices of Notaries
                        - Industry: Offices of Optometrists
                        - Industry: Offices of Other Health Practitioners
                        - Industry: Offices of Other Holding Companies
                        - Industry: Offices of Physical, Occupational and Speech Therapists, and Audiologists
                        - Industry: Offices of Physicians
                        - Industry: Offices of Physicians (except Mental Health Specialists)
                        - Industry: Offices of Physicians, Mental Health Specialists
                        - Industry: Offices of Podiatrists
                        - Industry: Offices of Real Estate Agents and Brokers
                        - Industry: Offices of Real Estate Appraisers
                        - Industry: Oil and Gas Extraction
                        - Industry: Oil and Gas Field Machinery and Equipment Manufacturing
                        - Industry: Oil and Gas Pipeline and Related Structures Construction
                        - Industry: Oilseed (except Soybean) Farming
                        - Industry: Oilseed and Grain Combination Farming
                        - Industry: Oilseed and Grain Farming
                        - Industry: One-Hour Photofinishing
                        - Industry: Open-End Investment Funds
                        - Industry: Ophthalmic Goods Manufacturing
                        - Industry: Ophthalmic Goods Merchant Wholesalers
                        - Industry: Optical Goods Stores
                        - Industry: Optical Instrument and Lens Manufacturing
                        - Industry: Orange Groves
                        - Industry: Ornamental and Architectural Metal Products Manufacturing
                        - Industry: Ornamental and Architectural Metal Work Manufacturing
                        - Industry: Other Accounting Services
                        - Industry: Other Activities Related to Credit Intermediation
                        - Industry: Other Activities Related to Real Estate
                        - Industry: Other Aircraft Parts and Auxiliary Equipment Manufacturing
                        - Industry: Other Airport Operations
                        - Industry: Other Aluminum Rolling and Drawing
                        - Industry: Other Ambulatory Health Care Services
                        - Industry: Other Amusement and Recreation Industries
                        - Industry: Other Animal Food Manufacturing
                        - Industry: Other Animal Production
                        - Industry: Other Apparel Accessories and Other Apparel Manufacturing
                        - Industry: Other Apparel Knitting Mills
                        - Industry: Other Aquaculture
                        - Industry: Other Automotive Mechanical and Electrical Repair and Maintenance
                        - Industry: Other Automotive Repair and Maintenance
                        - Industry: Other Basic Inorganic Chemical Manufacturing
                        - Industry: Other Basic Organic Chemical Manufacturing
                        - Industry: Other Building Equipment Contractors
                        - Industry: Other Building Finishing Contractors
                        - Industry: Other Building Material Dealers
                        - Industry: Other Business Service Centers (including Copy Shops)
                        - Industry: Other Business Support Services
                        - Industry: Other Chemical and Allied Products Merchant Wholesalers
                        - Industry: Other Chemical and Fertilizer Mineral Mining
                        - Industry: Other Chemical Product and Preparation Manufacturing
                        - Industry: Other Clothing Stores
                        - Industry: Other Commercial and Industrial Machinery and Equipment Rental and Leasing
                        - Industry: Other Commercial and Service Industry Machinery Manufacturing
                        - Industry: Other Commercial Equipment Merchant Wholesalers
                        - Industry: Other Commercial Printing
                        - Industry: Other Communication and Energy Wire Manufacturing
                        - Industry: Other Communications Equipment Manufacturing
                        - Industry: Other Community Housing Services
                        - Industry: Other Computer Peripheral Equipment Manufacturing
                        - Industry: Other Computer Related Services
                        - Industry: Other Concrete Product Manufacturing
                        - Industry: Other Construction Material Merchant Wholesalers
                        - Industry: Other Consumer Goods Rental
                        - Industry: Other Converted Paper Product Manufacturing
                        - Industry: Other Crop Farming
                        - Industry: Other Crushed and Broken Stone Mining and Quarrying
                        - Industry: Other Cut and Sew Apparel Manufacturing
                        - Industry: Other Depository Credit Intermediation
                        - Industry: Other Direct Insurance (except Life, Health, and Medical) Carriers
                        - Industry: Other Direct Selling Establishments
                        - Industry: Other Electric Power Generation
                        - Industry: Other Electrical Equipment and Component Manufacturing
                        - Industry: Other Electronic and Precision Equipment Repair and Maintenance
                        - Industry: Other Electronic Component Manufacturing
                        - Industry: Other Electronic Parts and Equipment Merchant Wholesalers
                        - Industry: Other Engine Equipment Manufacturing
                        - Industry: Other Fabricated Metal Product Manufacturing
                        - Industry: Other Fabricated Wire Product Manufacturing
                        - Industry: Other Farm Product Raw Material Merchant Wholesalers
                        - Industry: Other Financial Investment Activities
                        - Industry: Other Financial Vehicles
                        - Industry: Other Food Crops Grown Under Cover
                        - Industry: Other Food Manufacturing
                        - Industry: Other Footwear Manufacturing
                        - Industry: Other Foundation, Structure, and Building Exterior Contractors
                        - Industry: Other Fuel Dealers
                        - Industry: Other Furniture Related Product Manufacturing
                        - Industry: Other Gambling Industries
                        - Industry: Other Gasoline Stations
                        - Industry: Other General Government Support
                        - Industry: Other General Merchandise Stores
                        - Industry: Other General Purpose Machinery Manufacturing
                        - Industry: Other Grain Farming
                        - Industry: Other Grantmaking and Giving Services
                        - Industry: Other Grocery and Related Products Merchant Wholesalers
                        - Industry: Other Guided Missile and Space Vehicle Parts and Auxiliary Equipment Manufacturing
                        - Industry: Other Health and Personal Care Stores
                        - Industry: Other Heavy and Civil Engineering Construction
                        - Industry: Other Home Furnishings Stores
                        - Industry: Other Hosiery and Sock Mills
                        - Industry: Other Household Textile Product Mills
                        - Industry: Other Individual and Family Services
                        - Industry: Other Industrial Machinery Manufacturing
                        - Industry: Other Information Services
                        - Industry: Other Insurance Funds
                        - Industry: Other Insurance Related Activities
                        - Industry: Other Investment Pools and Funds
                        - Industry: Other Justice, Public Order, and Safety Activities
                        - Industry: Other Knit Fabric and Lace Mills
                        - Industry: Other Leather and Allied Product Manufacturing
                        - Industry: Other Legal Services
                        - Industry: Other Lighting Equipment Manufacturing
                        - Industry: Other Major Household Appliance Manufacturing
                        - Industry: Other Management Consulting Services
                        - Industry: Other Marine Fishing
                        - Industry: Other Measuring and Controlling Device Manufacturing
                        - Industry: Other Metal Container Manufacturing
                        - Industry: Other Metal Ore Mining
                        - Industry: Other Metal Valve and Pipe Fitting Manufacturing
                        - Industry: Other Metalworking Machinery Manufacturing
                        - Industry: Other Millwork (including Flooring)
                        - Industry: Other Miscellaneous Durable Goods Merchant Wholesalers
                        - Industry: Other Miscellaneous Manufacturing
                        - Industry: Other Miscellaneous Nondurable Goods Merchant Wholesalers
                        - Industry: Other Miscellaneous Store Retailers
                        - Industry: Other Motion Picture and Video Industries
                        - Industry: Other Motor Vehicle Dealers
                        - Industry: Other Motor Vehicle Electrical and Electronic Equipment Manufacturing
                        - Industry: Other Motor Vehicle Parts Manufacturing
                        - Industry: Other Noncitrus Fruit Farming
                        - Industry: Other Nondepository Credit Intermediation
                        - Industry: Other Nonferrous Foundries (except Die-Casting)
                        - Industry: Other Nonhazardous Waste Treatment and Disposal
                        - Industry: Other Nonmetallic Mineral Mining and Quarrying
                        - Industry: Other Nonmetallic Mineral Product Manufacturing
                        - Industry: Other Nonscheduled Air Transportation
                        - Industry: Other Oilseed Processing
                        - Industry: Other Ordnance and Accessories Manufacturing
                        - Industry: Other Outpatient Care Centers
                        - Industry: Other Performing Arts Companies
                        - Industry: Other Personal and Household Goods Repair and Maintenance
                        - Industry: Other Personal Care Services
                        - Industry: Other Personal Services
                        - Industry: Other Petroleum and Coal Products Manufacturing
                        - Industry: Other Pipeline Transportation
                        - Industry: Other Plastics Product Manufacturing
                        - Industry: Other Poultry Production
                        - Industry: Other Pressed and Blown Glass and Glassware Manufacturing
                        - Industry: Other Professional Equipment and Supplies Merchant Wholesalers
                        - Industry: Other Professional, Scientific, and Technical Services
                        - Industry: Other Publishers
                        - Industry: Other Residential Care Facilities
                        - Industry: Other Rubber Product Manufacturing
                        - Industry: Other Schools and Instruction
                        - Industry: Other Scientific and Technical Consulting Services
                        - Industry: Other Services (except Public Administration)
                        - Industry: Other Services Related to Advertising
                        - Industry: Other Services to Buildings and Dwellings
                        - Industry: Other Similar Organizations (except Business, Professional, Labor, and Political Organizations)
                        - Industry: Other Snack Food Manufacturing
                        - Industry: Other Social Advocacy Organizations
                        - Industry: Other Sound Recording Industries
                        - Industry: Other Specialized Design Services
                        - Industry: Other Specialty Food Stores
                        - Industry: Other Specialty Trade Contractors
                        - Industry: Other Spectator Sports
                        - Industry: Other Structural Clay Product Manufacturing
                        - Industry: Other Support Activities for Air Transportation
                        - Industry: Other Support Activities for Road Transportation
                        - Industry: Other Support Activities for Transportation
                        - Industry: Other Support Activities for Water Transportation
                        - Industry: Other Support Services
                        - Industry: Other Technical and Trade Schools
                        - Industry: Other Telecommunications
                        - Industry: Other Textile Product Mills
                        - Industry: Other Tobacco Product Manufacturing
                        - Industry: Other Transit and Ground Passenger Transportation
                        - Industry: Other Transportation Equipment Manufacturing
                        - Industry: Other Travel Arrangement and Reservation Services
                        - Industry: Other Traveler Accommodation
                        - Industry: Other Urban Transit Systems
                        - Industry: Other Vegetable (except Potato) and Melon Farming
                        - Industry: Other Warehousing and Storage
                        - Industry: Other Waste Collection
                        - Industry: Other Wood Product Manufacturing
                        - Industry: Outdoor Power Equipment Stores
                        - Industry: Outerwear Knitting Mills
                        - Industry: Outpatient Care Centers
                        - Industry: Outpatient Mental Health and Substance Abuse Centers
                        - Industry: Overhead Traveling Crane, Hoist, and Monorail System Manufacturing
                        - Industry: Packaged Frozen Food Merchant Wholesalers
                        - Industry: Packaging and Labeling Services
                        - Industry: Packaging Machinery Manufacturing
                        - Industry: Packing and Crating
                        - Industry: Paint and Coating Manufacturing
                        - Industry: Paint and Wallpaper Stores
                        - Industry: Paint, Coating, and Adhesive Manufacturing
                        - Industry: Paint, Varnish, and Supplies Merchant Wholesalers
                        - Industry: Painting and Wall Covering Contractors
                        - Industry: Paper (except Newsprint) Mills
                        - Industry: Paper and Paper Product Merchant Wholesalers
                        - Industry: Paper Bag and Coated and Treated Paper Manufacturing
                        - Industry: Paper Industry Machinery Manufacturing
                        - Industry: Paper Manufacturing
                        - Industry: Paper Mills
                        - Industry: Paperboard Container Manufacturing
                        - Industry: Paperboard Mills
                        - Industry: Parking Lots and Garages
                        - Industry: Parole Offices and Probation Offices
                        - Industry: Passenger Car Leasing
                        - Industry: Passenger Car Rental
                        - Industry: Passenger Car Rental and Leasing
                        - Industry: Payroll Services
                        - Industry: Peanut Farming
                        - Industry: Pen and Mechanical Pencil Manufacturing
                        - Industry: Pension Funds
                        - Industry: Performing Arts Companies
                        - Industry: Performing Arts, Spectator Sports, and Related Industries
                        - Industry: Periodical Publishers
                        - Industry: Perishable Prepared Food Manufacturing
                        - Industry: Personal and Household Goods Repair and Maintenance
                        - Industry: Personal and Laundry Services
                        - Industry: Personal Care Services
                        - Industry: Personal Leather Good (except Women's Handbag and Purse) Manufacturing
                        - Industry: Pesticide and Other Agricultural Chemical Manufacturing
                        - Industry: Pesticide, Fertilizer, and Other Agricultural Chemical Manufacturing
                        - Industry: Pet and Pet Supplies Stores
                        - Industry: Pet Care (except Veterinary) Services
                        - Industry: Petrochemical Manufacturing
                        - Industry: Petroleum and Coal Products Manufacturing
                        - Industry: Petroleum and Petroleum Products Merchant Wholesalers
                        - Industry: Petroleum and Petroleum Products Merchant Wholesalers (except Bulk Stations and Terminals)
                        - Industry: Petroleum Bulk Stations and Terminals
                        - Industry: Petroleum Lubricating Oil and Grease Manufacturing
                        - Industry: Petroleum Refineries
                        - Industry: Pharmaceutical and Medicine Manufacturing
                        - Industry: Pharmaceutical Preparation Manufacturing
                        - Industry: Pharmacies and Drug Stores
                        - Industry: Phosphate Rock Mining
                        - Industry: Phosphatic Fertilizer Manufacturing
                        - Industry: Photofinishing
                        - Industry: Photofinishing Laboratories (except One-Hour)
                        - Industry: Photographic and Photocopying Equipment Manufacturing
                        - Industry: Photographic Equipment and Supplies Merchant Wholesalers
                        - Industry: Photographic Film, Paper, Plate, and Chemical Manufacturing
                        - Industry: Photographic Services
                        - Industry: Photography Studios, Portrait
                        - Industry: Piece Goods, Notions, and Other Dry Goods Merchant Wholesalers
                        - Industry: Pipeline Transportation
                        - Industry: Pipeline Transportation of Crude Oil
                        - Industry: Pipeline Transportation of Natural Gas
                        - Industry: Pipeline Transportation of Refined Petroleum Products
                        - Industry: Plastics and Rubber Industry Machinery Manufacturing
                        - Industry: Plastics and Rubber Products Manufacturing
                        - Industry: Plastics Bag and Pouch Manufacturing
                        - Industry: Plastics Bottle Manufacturing
                        - Industry: Plastics Material and Resin Manufacturing
                        - Industry: Plastics Materials and Basic Forms and Shapes Merchant Wholesalers
                        - Industry: Plastics Packaging Film and Sheet (including Laminated) Manufacturing
                        - Industry: Plastics Packaging Materials and Unlaminated Film and Sheet Manufacturing
                        - Industry: Plastics Pipe and Pipe Fitting Manufacturing
                        - Industry: Plastics Pipe, Pipe Fitting, and Unlaminated Profile Shape Manufacturing
                        - Industry: Plastics Plumbing Fixture Manufacturing
                        - Industry: Plastics Product Manufacturing
                        - Industry: Plate Work and Fabricated Structural Product Manufacturing
                        - Industry: Plate Work Manufacturing
                        - Industry: Plumbing and Heating Equipment and Supplies (Hydronics) Merchant Wholesalers
                        - Industry: Plumbing Fixture Fitting and Trim Manufacturing
                        - Industry: Plumbing, Heating, and Air-Conditioning Contractors
                        - Industry: Police Protection
                        - Industry: Polish and Other Sanitation Good Manufacturing
                        - Industry: Political Organizations
                        - Industry: Polystyrene Foam Product Manufacturing
                        - Industry: Porcelain Electrical Supply Manufacturing
                        - Industry: Port and Harbor Operations
                        - Industry: Portfolio Management
                        - Industry: Postal Service
                        - Industry: Postharvest Crop Activities (except Cotton Ginning)
                        - Industry: Postproduction Services and Other Motion Picture and Video Industries
                        - Industry: Potash, Soda, and Borate Mineral Mining
                        - Industry: Potato Farming
                        - Industry: Pottery, Ceramics, and Plumbing Fixture Manufacturing
                        - Industry: Poultry and Egg Production
                        - Industry: Poultry and Poultry Product Merchant Wholesalers
                        - Industry: Poultry Hatcheries
                        - Industry: Poultry Processing
                        - Industry: Poured Concrete Foundation and Structure Contractors
                        - Industry: Powder Metallurgy Part Manufacturing
                        - Industry: Power and Communication Line and Related Structures Construction
                        - Industry: Power Boiler and Heat Exchanger Manufacturing
                        - Industry: Power, Distribution, and Specialty Transformer Manufacturing
                        - Industry: Power-Driven Handtool Manufacturing
                        - Industry: Precision Turned Product Manufacturing
                        - Industry: Prefabricated Metal Building and Component Manufacturing
                        - Industry: Prefabricated Wood Building Manufacturing
                        - Industry: Prepress Services
                        - Industry: Prerecorded Compact Disc (except Software), Tape, and Record Reproducing
                        - Industry: Prerecorded Tape, Compact Disc, and Record Stores
                        - Industry: Primary Aluminum Production
                        - Industry: Primary Battery Manufacturing
                        - Industry: Primary Metal Manufacturing
                        - Industry: Primary Smelting and Refining of Copper
                        - Industry: Primary Smelting and Refining of Nonferrous Metal (except Copper and Aluminum)
                        - Industry: Printed Circuit Assembly (Electronic Assembly) Manufacturing
                        - Industry: Printing
                        - Industry: Printing and Related Support Activities
                        - Industry: Printing and Writing Paper Merchant Wholesalers
                        - Industry: Printing Ink Manufacturing
                        - Industry: Printing Machinery and Equipment Manufacturing
                        - Industry: Private Households
                        - Industry: Private Mail Centers
                        - Industry: Process, Physical Distribution, and Logistics Consulting Services
                        - Industry: Professional and Commercial Equipment and Supplies Merchant Wholesalers
                        - Industry: Professional and Management Development Training
                        - Industry: Professional Employer Organizations
                        - Industry: Professional Organizations
                        - Industry: Professional, Scientific, and Technical Services
                        - Industry: Promoters of Performing Arts, Sports, and Similar Events
                        - Industry: Promoters of Performing Arts, Sports, and Similar Events with Facilities
                        - Industry: Promoters of Performing Arts, Sports, and Similar Events without Facilities
                        - Industry: Psychiatric and Substance Abuse Hospitals
                        - Industry: Public Administration
                        - Industry: Public Finance Activities
                        - Industry: Public Relations Agencies
                        - Industry: Publishing Industries (except Internet)
                        - Industry: Pulp Mills
                        - Industry: Pulp, Paper, and Paperboard Mills
                        - Industry: Pump and Compressor Manufacturing
                        - Industry: Pump and Pumping Equipment Manufacturing
                        - Industry: Quick Printing
                        - Industry: Racetracks
                        - Industry: Radio and Television Broadcasting
                        - Industry: Radio and Television Broadcasting and Wireless Communications Equipment Manufacturing
                        - Industry: Radio Broadcasting
                        - Industry: Radio Networks
                        - Industry: Radio Stations
                        - Industry: Radio, Television, and Other Electronics Stores
                        - Industry: Rail Transportation
                        - Industry: Railroad Rolling Stock Manufacturing
                        - Industry: Ready-Mix Concrete Manufacturing
                        - Industry: Real Estate
                        - Industry: Real Estate and Rental and Leasing
                        - Industry: Real Estate Credit
                        - Industry: Real Estate Property Managers
                        - Industry: Reconstituted Wood Product Manufacturing
                        - Industry: Record Production
                        - Industry: Recreational and Vacation Camps (except Campgrounds)
                        - Industry: Recreational Goods Rental
                        - Industry: Recreational Vehicle Dealers
                        - Industry: Recyclable Material Merchant Wholesalers
                        - Industry: Refrigerated Warehousing and Storage
                        - Industry: Refrigeration Equipment and Supplies Merchant Wholesalers
                        - Industry: Regulation and Administration of Communications, Electric, Gas, and Other Utilities
                        - Industry: Regulation and Administration of Transportation Programs
                        - Industry: Regulation of Agricultural Marketing and Commodities
                        - Industry: Regulation, Licensing, and Inspection of Miscellaneous Commercial Sectors
                        - Industry: Reinsurance Carriers
                        - Industry: Relay and Industrial Control Manufacturing
                        - Industry: Religious Organizations
                        - Industry: Religious, Grantmaking, Civic, Professional, and Similar Organizations
                        - Industry: Remediation and Other Waste Management Services
                        - Industry: Remediation Services
                        - Industry: Rendering and Meat Byproduct Processing
                        - Industry: Rental and Leasing Services
                        - Industry: Repair and Maintenance
                        - Industry: Repossession Services
                        - Industry: Research and Development in Biotechnology
                        - Industry: Research and Development in the Physical, Engineering, and Life Sciences
                        - Industry: Research and Development in the Physical, Engineering, and Life Sciences (except Biotechnology)
                        - Industry: Research and Development in the Social Sciences and Humanities
                        - Industry: Residential Building Construction
                        - Industry: Residential Electric Lighting Fixture Manufacturing
                        - Industry: Residential Mental Health and Substance Abuse Facilities
                        - Industry: Residential Mental Retardation Facilities
                        - Industry: Residential Mental Retardation, Mental Health and Substance Abuse Facilities
                        - Industry: Residential Property Managers
                        - Industry: Residential Remodelers
                        - Industry: Resilient Floor Covering Manufacturing
                        - Industry: Resin and Synthetic Rubber Manufacturing
                        - Industry: Resin, Synthetic Rubber, and Artificial Synthetic Fibers and Filaments Manufacturing
                        - Industry: Retail Bakeries
                        - Industry: Retail Trade
                        - Industry: Reupholstery and Furniture Repair
                        - Industry: Rice Farming
                        - Industry: Rice Milling
                        - Industry: Roasted Nuts and Peanut Butter Manufacturing
                        - Industry: Rolled Steel Shape Manufacturing
                        - Industry: Rolling and Drawing of Purchased Steel
                        - Industry: Rolling Mill Machinery and Equipment Manufacturing
                        - Industry: Roofing Contractors
                        - Industry: Roofing, Siding, and Insulation Material Merchant Wholesalers
                        - Industry: Rooming and Boarding Houses
                        - Industry: Rope, Cordage, and Twine Mills
                        - Industry: Rubber and Plastics Footwear Manufacturing
                        - Industry: Rubber and Plastics Hoses and Belting Manufacturing
                        - Industry: Rubber Product Manufacturing
                        - Industry: Rubber Product Manufacturing for Mechanical Use
                        - Industry: RV (Recreational Vehicle) Parks and Campgrounds
                        - Industry: RV (Recreational Vehicle) Parks and Recreational Camps
                        - Industry: Sales Financing
                        - Industry: Sand, Gravel, Clay, and Ceramic and Refractory Minerals Mining and Quarrying
                        - Industry: Sanitary Paper Product Manufacturing
                        - Industry: Satellite Telecommunications
                        - Industry: Savings Institutions
                        - Industry: Saw Blade and Handsaw Manufacturing
                        - Industry: Sawmill and Woodworking Machinery Manufacturing
                        - Industry: Sawmills
                        - Industry: Sawmills and Wood Preservation
                        - Industry: Scale and Balance Manufacturing
                        - Industry: Scenic and Sightseeing Transportation
                        - Industry: Scenic and Sightseeing Transportation, Land
                        - Industry: Scenic and Sightseeing Transportation, Other
                        - Industry: Scenic and Sightseeing Transportation, Water
                        - Industry: Scheduled Air Transportation
                        - Industry: Scheduled Freight Air Transportation
                        - Industry: Scheduled Passenger Air Transportation
                        - Industry: Schiffli Machine Embroidery
                        - Industry: School and Employee Bus Transportation
                        - Industry: Scientific Research and Development Services
                        - Industry: Seafood Canning
                        - Industry: Seafood Product Preparation and Packaging
                        - Industry: Search, Detection, Navigation, Guidance, Aeronautical, and Nautical System and Instrument Manufacturing
                        - Industry: Seasoning and Dressing Manufacturing
                        - Industry: Secondary Market Financing
                        - Industry: Secondary Smelting and Alloying of Aluminum
                        - Industry: Secondary Smelting, Refining, and Alloying of Copper
                        - Industry: Secondary Smelting, Refining, and Alloying of Nonferrous Metal (except Copper and Aluminum)
                        - Industry: Securities and Commodity Contracts Intermediation and Brokerage
                        - Industry: Securities and Commodity Exchanges
                        - Industry: Securities Brokerage
                        - Industry: Securities, Commodity Contracts, and Other Financial Investments and Related Activities
                        - Industry: Security Guards and Patrol Services
                        - Industry: Security Systems Services
                        - Industry: Security Systems Services (except Locksmiths)
                        - Industry: Semiconductor and Other Electronic Component Manufacturing
                        - Industry: Semiconductor and Related Device Manufacturing
                        - Industry: Semiconductor Machinery Manufacturing
                        - Industry: Septic Tank and Related Services
                        - Industry: Service Establishment Equipment and Supplies Merchant Wholesalers
                        - Industry: Services for the Elderly and Persons with Disabilities
                        - Industry: Services to Buildings and Dwellings
                        - Industry: Setup Paperboard Box Manufacturing
                        - Industry: Sewage Treatment Facilities
                        - Industry: Sewing, Needlework, and Piece Goods Stores
                        - Industry: Sheep and Goat Farming
                        - Industry: Sheep Farming
                        - Industry: Sheer Hosiery Mills
                        - Industry: Sheet Metal Work Manufacturing
                        - Industry: Shellfish Farming
                        - Industry: Shellfish Fishing
                        - Industry: Ship and Boat Building
                        - Industry: Ship Building and Repairing
                        - Industry: Shoe Stores
                        - Industry: Short Line Railroads
                        - Industry: Showcase, Partition, Shelving, and Locker Manufacturing
                        - Industry: Siding Contractors
                        - Industry: Sign Manufacturing
                        - Industry: Silver Ore Mining
                        - Industry: Silverware and Hollowware Manufacturing
                        - Industry: Site Preparation Contractors
                        - Industry: Skiing Facilities
                        - Industry: Small Arms Ammunition Manufacturing
                        - Industry: Small Arms Manufacturing
                        - Industry: Small Electrical Appliance Manufacturing
                        - Industry: Snack and Nonalcoholic Beverage Bars
                        - Industry: Snack Food Manufacturing
                        - Industry: Soap and Cleaning Compound Manufacturing
                        - Industry: Soap and Other Detergent Manufacturing
                        - Industry: Soap, Cleaning Compound, and Toilet Preparation Manufacturing
                        - Industry: Social Advocacy Organizations
                        - Industry: Social Assistance
                        - Industry: Soft Drink and Ice Manufacturing
                        - Industry: Soft Drink Manufacturing
                        - Industry: Software Publishers
                        - Industry: Software Reproducing
                        - Industry: Softwood Veneer and Plywood Manufacturing
                        - Industry: Soil Preparation, Planting, and Cultivating
                        - Industry: Solid Waste Collection
                        - Industry: Solid Waste Combustors and Incinerators
                        - Industry: Solid Waste Landfill
                        - Industry: Sound Recording Industries
                        - Industry: Sound Recording Studios
                        - Industry: Soybean Farming
                        - Industry: Soybean Processing
                        - Industry: Space Research and Technology
                        - Industry: Special Die and Tool, Die Set, Jig, and Fixture Manufacturing
                        - Industry: Special Food Services
                        - Industry: Special Needs Transportation
                        - Industry: Specialized Design Services
                        - Industry: Specialized Freight (except Used Goods) Trucking, Local
                        - Industry: Specialized Freight (except Used Goods) Trucking, Long-Distance
                        - Industry: Specialized Freight Trucking
                        - Industry: Specialty (except Psychiatric and Substance Abuse) Hospitals
                        - Industry: Specialty Canning
                        - Industry: Specialty Food Stores
                        - Industry: Specialty Trade Contractors
                        - Industry: Spectator Sports
                        - Industry: Speed Changer, Industrial High-Speed Drive, and Gear Manufacturing
                        - Industry: Spice and Extract Manufacturing
                        - Industry: Sporting and Athletic Goods Manufacturing
                        - Industry: Sporting and Recreational Goods and Supplies Merchant Wholesalers
                        - Industry: Sporting Goods Stores
                        - Industry: Sporting Goods, Hobby, and Musical Instrument Stores
                        - Industry: Sporting Goods, Hobby, Book, and Music Stores
                        - Industry: Sports and Recreation Instruction
                        - Industry: Sports Teams and Clubs
                        - Industry: Spring (Heavy Gauge) Manufacturing
                        - Industry: Spring (Light Gauge) Manufacturing
                        - Industry: Spring and Wire Product Manufacturing
                        - Industry: Starch and Vegetable Fats and Oils Manufacturing
                        - Industry: Stationery and Office Supplies Merchant Wholesalers
                        - Industry: Stationery Product Manufacturing
                        - Industry: Stationery, Tablet, and Related Product Manufacturing
                        - Industry: Steam and Air-Conditioning Supply
                        - Industry: Steel Foundries (except Investment)
                        - Industry: Steel Investment Foundries
                        - Industry: Steel Product Manufacturing from Purchased Steel
                        - Industry: Steel Wire Drawing
                        - Industry: Stone Mining and Quarrying
                        - Industry: Storage Battery Manufacturing
                        - Industry: Strawberry Farming
                        - Industry: Structural Steel and Precast Concrete Contractors
                        - Industry: Sugar and Confectionery Product Manufacturing
                        - Industry: Sugar Beet Farming
                        - Industry: Sugar Manufacturing
                        - Industry: Sugarcane Farming
                        - Industry: Sugarcane Mills
                        - Industry: Supermarkets and Other Grocery (except Convenience) Stores
                        - Industry: Support Activities for Agriculture and Forestry
                        - Industry: Support Activities for Air Transportation
                        - Industry: Support Activities for Animal Production
                        - Industry: Support Activities for Coal Mining
                        - Industry: Support Activities for Crop Production
                        - Industry: Support Activities for Forestry
                        - Industry: Support Activities for Metal Mining
                        - Industry: Support Activities for Mining
                        - Industry: Support Activities for Nonmetallic Minerals (except Fuels) Mining
                        - Industry: Support Activities for Oil and Gas Operations
                        - Industry: Support Activities for Printing
                        - Industry: Support Activities for Rail Transportation
                        - Industry: Support Activities for Road Transportation
                        - Industry: Support Activities for Transportation
                        - Industry: Support Activities for Water Transportation
                        - Industry: Surface Active Agent Manufacturing
                        - Industry: Surface-Coated Paperboard Manufacturing
                        - Industry: Surgical and Medical Instrument Manufacturing
                        - Industry: Surgical Appliance and Supplies Manufacturing
                        - Industry: Surveying and Mapping (except Geophysical) Services
                        - Industry: Switchgear and Switchboard Apparatus Manufacturing
                        - Industry: Synthetic Dye and Pigment Manufacturing
                        - Industry: Synthetic Organic Dye and Pigment Manufacturing
                        - Industry: Synthetic Rubber Manufacturing
                        - Industry: Tax Preparation Services
                        - Industry: Taxi and Limousine Service
                        - Industry: Taxi Service
                        - Industry: Technical and Trade Schools
                        - Industry: Telecommunications
                        - Industry: Telecommunications Resellers
                        - Industry: Telemarketing Bureaus and Other Contact Centers
                        - Industry: Telephone Answering Services
                        - Industry: Telephone Apparatus Manufacturing
                        - Industry: Telephone Call Centers
                        - Industry: Teleproduction and Other Postproduction Services
                        - Industry: Television Broadcasting
                        - Industry: Temporary Help Services
                        - Industry: Temporary Shelters
                        - Industry: Testing Laboratories
                        - Industry: Textile and Fabric Finishing (except Broadwoven Fabric) Mills
                        - Industry: Textile and Fabric Finishing and Fabric Coating Mills
                        - Industry: Textile and Fabric Finishing Mills
                        - Industry: Textile Bag and Canvas Mills
                        - Industry: Textile Bag Mills
                        - Industry: Textile Furnishings Mills
                        - Industry: Textile Machinery Manufacturing
                        - Industry: Textile Mills
                        - Industry: Textile Product Mills
                        - Industry: Theater Companies and Dinner Theaters
                        - Industry: Third Party Administration of Insurance and Pension Funds
                        - Industry: Thread Mills
                        - Industry: Tile and Terrazzo Contractors
                        - Industry: Timber Tract Operations
                        - Industry: Tire and Tube Merchant Wholesalers
                        - Industry: Tire Cord and Tire Fabric Mills
                        - Industry: Tire Dealers
                        - Industry: Tire Manufacturing
                        - Industry: Tire Manufacturing (except Retreading)
                        - Industry: Tire Retreading
                        - Industry: Title Abstract and Settlement Offices
                        - Industry: Tobacco and Tobacco Product Merchant Wholesalers
                        - Industry: Tobacco Farming
                        - Industry: Tobacco Manufacturing
                        - Industry: Tobacco Product Manufacturing
                        - Industry: Tobacco Stemming and Redrying
                        - Industry: Tobacco Stores
                        - Industry: Toilet Preparation Manufacturing
                        - Industry: Tortilla Manufacturing
                        - Industry: Totalizing Fluid Meter and Counting Device Manufacturing
                        - Industry: Tour Operators
                        - Industry: Toy and Hobby Goods and Supplies Merchant Wholesalers
                        - Industry: Tradebinding and Related Work
                        - Industry: Transit and Ground Passenger Transportation
                        - Industry: Translation and Interpretation Services
                        - Industry: Transportation and Warehousing
                        - Industry: Transportation Equipment and Supplies (except Motor Vehicle) Merchant Wholesalers
                        - Industry: Transportation Equipment Manufacturing
                        - Industry: Travel Agencies
                        - Industry: Travel Arrangement and Reservation Services
                        - Industry: Travel Trailer and Camper Manufacturing
                        - Industry: Traveler Accommodation
                        - Industry: Tree Nut Farming
                        - Industry: Truck Trailer Manufacturing
                        - Industry: Truck Transportation
                        - Industry: Truck, Utility Trailer, and RV (Recreational Vehicle) Rental and Leasing
                        - Industry: Truss Manufacturing
                        - Industry: Trust, Fiduciary, and Custody Activities
                        - Industry: Trusts, Estates, and Agency Accounts
                        - Industry: Turbine and Turbine Generator Set Units Manufacturing
                        - Industry: Turkey Production
                        - Industry: Turned Product and Screw, Nut, and Bolt Manufacturing
                        - Industry: Uncoated Paper and Multiwall Bag Manufacturing
                        - Industry: Underwear and Nightwear Knitting Mills
                        - Industry: Unlaminated Plastics Film and Sheet (except Packaging) Manufacturing
                        - Industry: Unlaminated Plastics Profile Shape Manufacturing
                        - Industry: Upholstered Household Furniture Manufacturing
                        - Industry: Uranium-Radium-Vanadium Ore Mining
                        - Industry: Urban Transit Systems
                        - Industry: Urethane and Other Foam Product (except Polystyrene) Manufacturing
                        - Industry: Used Car Dealers
                        - Industry: Used Household and Office Goods Moving
                        - Industry: Used Merchandise Stores
                        - Industry: Utilities
                        - Industry: Utility System Construction
                        - Industry: Vegetable and Melon Farming
                        - Industry: Vehicular Lighting Equipment Manufacturing
                        - Industry: Vending Machine Operators
                        - Industry: Veneer, Plywood, and Engineered Wood Product Manufacturing
                        - Industry: Ventilation, Heating, Air-Conditioning, and Commercial Refrigeration Equipment Manufacturing
                        - Industry: Veterinary Services
                        - Industry: Video Tape and Disc Rental
                        - Industry: Vitreous China Plumbing Fixture and China and Earthenware Bathroom Accessories Manufacturing
                        - Industry: Vitreous China, Fine Earthenware, and Other Pottery Product Manufacturing
                        - Industry: Vocational Rehabilitation Services
                        - Industry: Voluntary Health Organizations
                        - Industry: Warehouse Clubs and Supercenters
                        - Industry: Warehousing and Storage
                        - Industry: Warm Air Heating and Air-Conditioning Equipment and Supplies Merchant Wholesalers
                        - Industry: Waste Collection
                        - Industry: Waste Management and Remediation Services
                        - Industry: Waste Treatment and Disposal
                        - Industry: Watch, Clock, and Part Manufacturing
                        - Industry: Water and Sewer Line and Related Structures Construction
                        - Industry: Water Supply and Irrigation Systems
                        - Industry: Water Transportation
                        - Industry: Water, Sewage and Other Systems
                        - Industry: Weft Knit Fabric Mills
                        - Industry: Welding and Soldering Equipment Manufacturing
                        - Industry: Wet Corn Milling
                        - Industry: Wheat Farming
                        - Industry: Wholesale Electronic Markets and Agents and Brokers
                        - Industry: Wholesale Trade
                        - Industry: Wholesale Trade Agents and Brokers
                        - Industry: Window Treatment Stores
                        - Industry: Wine and Distilled Alcoholic Beverage Merchant Wholesalers
                        - Industry: Wineries
                        - Industry: Wired Telecommunications Carriers
                        - Industry: Wireless Telecommunications Carriers (except Satellite)
                        - Industry: Wiring Device Manufacturing
                        - Industry: Women's and Girls' Cut and Sew Apparel Manufacturing
                        - Industry: Women's and Girls' Cut and Sew Blouse and Shirt Manufacturing
                        - Industry: Women's and Girls' Cut and Sew Dress Manufacturing
                        - Industry: Women's and Girls' Cut and Sew Lingerie, Loungewear, and Nightwear Manufacturing
                        - Industry: Women's and Girls' Cut and Sew Other Outerwear Manufacturing
                        - Industry: Women's and Girls' Cut and Sew Suit, Coat, Tailored Jacket, and Skirt Manufacturing
                        - Industry: Women's Clothing Stores
                        - Industry: Women's Footwear (except Athletic) Manufacturing
                        - Industry: Women's Handbag and Purse Manufacturing
                        - Industry: Women's, Children's, and Infants' Clothing and Accessories Merchant Wholesalers
                        - Industry: Women's, Girls', and Infants' Cut and Sew Apparel Contractors
                        - Industry: Wood Container and Pallet Manufacturing
                        - Industry: Wood Kitchen Cabinet and Countertop Manufacturing
                        - Industry: Wood Office Furniture Manufacturing
                        - Industry: Wood Preservation
                        - Industry: Wood Product Manufacturing
                        - Industry: Wood Television, Radio, and Sewing Machine Cabinet Manufacturing
                        - Industry: Wood Window and Door Manufacturing
                        - Industry: Yarn Spinning Mills
                        - Industry: Yarn Texturizing, Throwing, and Twisting Mills
                        - Industry: Zoos and Botanical Gardens