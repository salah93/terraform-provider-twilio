# Twilio Provider
The Twilio provider is used to interact with the resources supported by Twilio.
The provider needs to be configured with the proper credentials before it can be used.

- [Twilio Provider](#twilio-provider)
  * [Example Usage](#example-usage)
    + [Argument Reference](#argument-reference)
  * [Example: Available Phone Number Local Data Source](#example-available-phone-number-local-data-source)
  * [Example: Proxy Service](#example-proxy-service)

## Example Usage
```hcl-terraform
# Configure the Twilio provider
provider "twilio" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  auth_token = "YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY"
}

# Create a new chat service
resource "twilio_chat_service" "default" {
  # ...
}

# Create a new flex flow
resource "twilio_flex_flow" "default" {
  # ...
}
```

then:
```bash
$ terraform init
$ terraform apply
$ terraform destroy
```

### Argument Reference
The following arguments are supported:
- `account_sid` - (Required) Twilio Account SID. This can also be set via the `ACCOUNT_SID` environment variable. 
- `auth_token` - (Required) Auth token for the account. This can also be set via the `AUTH_TOKEN` environment variable.  

## Example: Available Phone Number Local Data Source
Use this data source to retrieve information about [Available Local Phone Numbers](https://www.twilio.com/docs/phone-numbers/api/availablephonenumberlocal-resource). 

### Usage
```hcl-terraform
data "twilio_available_phone_numbers_local" "default" {}
```

### Attributes Reference
- `area_code` - The area code of the phone numbers to read. Applies to only phone numbers in the US and Canada.
- `contains` - The pattern on which to match phone numbers. Valid characters are *, 0-9, a-z, and A-Z. The * character matches any single digit. If specified, this value must have at least two characters.
- `sms_enabled` - Whether the phone numbers can receive text messages. Can be: true or false.
- `mms_enabled` - Whether the phone numbers can receive MMS messages. Can be: true or false.
- `voice_enabled` - Whether the phone numbers can receive calls. Can be: true or false.
- `exclude_all_address_required` - Whether to exclude phone numbers that require an [Address](https://www.twilio.com/docs/usage/api/address). Can be: true or false and the default is false.
- `exclude_local_address_required` - Whether to exclude phone numbers that require a local [Address](https://www.twilio.com/docs/usage/api/address). Can be: true or false and the default is false.
- `exclude_foreign_address_required` - Whether to exclude phone numbers that require a foreign [Address](https://www.twilio.com/docs/usage/api/address). Can be: true or false and the default is false.
- `beta` - Whether to read phone numbers that are new to the Twilio platform. Can be: true or false and the default is true.
- `near_number` - Given a phone number, find a geographically close number within distance miles. Distance defaults to 25 miles. Applies to only phone numbers in the US and Canada.
- `near_lat_long` - Given a latitude/longitude pair lat,long find geographically close numbers within distance miles. Applies to only phone numbers in the US and Canada.
- `distance` - The search radius, in miles, for a near_ query. Can be up to 500 and the default is 25. Applies to only phone numbers in the US and Canada.
- `in_postal_code` - Limit results to a particular postal code. Given a phone number, search within the same postal code as that number. Applies to only phone numbers in the US and Canada.
- `in_region` - Limit results to a particular region, state, or province. Given a phone number, search within the same region as that number. Applies to only phone numbers in the US and Canada.
- `in_rate_center` - Limit results to a specific rate center, or given a phone number search within the same rate center as that number. Requires in_lata to be set as well. Applies to only phone numbers in the US and Canada.
- `in_lata` - Limit results to a specific local access and transport area (LATA). Given a phone number, search within the same LATA as that number. Applies to only phone numbers in the US and Canada.
- `in_locality` - Limit results to a particular locality or city. Given a phone number, search within the same Locality as that number.
- `fax_enabled` - Whether the phone numbers can receive faxes. Can be: true or false.

For more information see [the Available Phone Number Local documentation](https://www.twilio.com/docs/phone-numbers/api/availablephonenumberlocal-resource).

## Example: Proxy Service
`twilio_proxy_service` provides a Twilio Proxy Service resource.
Twilio Proxy Service is the top-level scope of all other resources in the Proxy Service REST API.

### Usage:
```hcl-terraform
resource "twilio_proxy_service" "default" {
  unique_name = "Unique Proxy Service"
  chat_instance_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  callback_url = "https://example.com/proxy/callback/url"
}
```

### Argument Reference
- `unique_name` - An application-defined string that uniquely identifies the resource. This value must be 191 characters or fewer in length and be unique. This value should not have PII.
- `default_ttl` - The default ttl value to set for Sessions created in the Service. The TTL (time to live) is measured in seconds after the Session's last create or last Interaction. The default value of 0 indicates an unlimited Session length. You can override a Session's default TTL value by setting its ttl value.
- `callback_url` - The URL we should call when the interaction status changes.
- `geo_match_level` - Where a proxy number must be located relative to the participant identifier. Can be: `country`, `area-code`, or `extended-area-code`. The default value is `country` and more specific areas than `country` are only available in North America.
- `number_selection_behavior` - The preference for Proxy Number selection in the Service instance. Can be: `prefer-sticky` or `avoid-sticky` and the default is `prefer-sticky`. `prefer-sticky` means that we will try and select the same Proxy Number for a given participant if they have previous Sessions, but we will not fail if that Proxy Number cannot be used. `avoid-sticky` means that we will try to use different Proxy Numbers as long as that is possible within a given pool rather than try and use a previously assigned number.
- `intercept_callback_url` - The URL we call on each interaction. If we receive a 403 status, we block the interaction; otherwise the interaction continues.
- `out_of_session_callback_url` - The URL we should call when an inbound call or SMS action occurs on a closed or non-existent Session. If your server (or a Twilio function) responds with valid TwiML, we will process it. This means it is possible, for example, to play a message for a call, send an automated text message response, or redirect a call to another Phone Number. See Out-of-Session Callback Response Guide for more information.
- `chat_instance_sid` - The SID of the Chat Service Instance managed by Proxy Service. The Chat Service enables Proxy to forward SMS and channel messages to this chat instance. This is a one-to-one relationship.

For more information see [the Proxy Service documentation](https://www.twilio.com/docs/proxy/api/service).