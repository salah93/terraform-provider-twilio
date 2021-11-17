/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/twilio/terraform-provider-twilio/client"
	. "github.com/twilio/terraform-provider-twilio/core"
	. "github.com/twilio/twilio-go/rest/verify/v2"
)

func ResourceServicesRateLimitsBuckets() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesRateLimitsBuckets,
		ReadContext:   readServicesRateLimitsBuckets,
		UpdateContext: updateServicesRateLimitsBuckets,
		DeleteContext: deleteServicesRateLimitsBuckets,
		Schema: map[string]*schema.Schema{
			"service_sid":    AsString(SchemaRequired),
			"rate_limit_sid": AsString(SchemaRequired),
			"interval":       AsInt(SchemaRequired),
			"max":            AsInt(SchemaRequired),
			"sid":            AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesRateLimitsBucketsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServicesRateLimitsBuckets(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateBucketParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	rateLimitSid := d.Get("rate_limit_sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.CreateBucket(serviceSid, rateLimitSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{serviceSid, rateLimitSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteServicesRateLimitsBuckets(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	rateLimitSid := d.Get("rate_limit_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.VerifyV2.DeleteBucket(serviceSid, rateLimitSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServicesRateLimitsBuckets(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	rateLimitSid := d.Get("rate_limit_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.FetchBucket(serviceSid, rateLimitSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesRateLimitsBucketsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected service_sid/rate_limit_sid/sid"

	if len(importParts) != 3 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("service_sid", importParts[0])
	d.Set("rate_limit_sid", importParts[1])
	d.Set("sid", importParts[2])

	return nil
}
func updateServicesRateLimitsBuckets(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateBucketParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	rateLimitSid := d.Get("rate_limit_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.UpdateBucket(serviceSid, rateLimitSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceServicesEntities() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesEntities,
		ReadContext:   readServicesEntities,
		DeleteContext: deleteServicesEntities,
		Schema: map[string]*schema.Schema{
			"service_sid": AsString(SchemaForceNewRequired),
			"identity":    AsString(SchemaForceNewRequired),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesEntitiesImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServicesEntities(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateEntityParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.CreateEntity(serviceSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{serviceSid}
	idParts = append(idParts, (*r.Identity))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteServicesEntities(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	identity := d.Get("identity").(string)

	err := m.(*client.Config).Client.VerifyV2.DeleteEntity(serviceSid, identity)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServicesEntities(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	identity := d.Get("identity").(string)

	r, err := m.(*client.Config).Client.VerifyV2.FetchEntity(serviceSid, identity)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesEntitiesImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected service_sid/identity"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("service_sid", importParts[0])
	d.Set("identity", importParts[1])

	return nil
}
func ResourceServicesMessagingConfigurations() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesMessagingConfigurations,
		ReadContext:   readServicesMessagingConfigurations,
		UpdateContext: updateServicesMessagingConfigurations,
		DeleteContext: deleteServicesMessagingConfigurations,
		Schema: map[string]*schema.Schema{
			"service_sid":           AsString(SchemaRequired),
			"country":               AsString(SchemaRequired),
			"messaging_service_sid": AsString(SchemaRequired),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesMessagingConfigurationsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServicesMessagingConfigurations(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateMessagingConfigurationParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.CreateMessagingConfiguration(serviceSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{serviceSid}
	idParts = append(idParts, (*r.Country))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteServicesMessagingConfigurations(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	country := d.Get("country").(string)

	err := m.(*client.Config).Client.VerifyV2.DeleteMessagingConfiguration(serviceSid, country)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServicesMessagingConfigurations(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	country := d.Get("country").(string)

	r, err := m.(*client.Config).Client.VerifyV2.FetchMessagingConfiguration(serviceSid, country)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesMessagingConfigurationsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected service_sid/country"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("service_sid", importParts[0])
	d.Set("country", importParts[1])

	return nil
}
func updateServicesMessagingConfigurations(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateMessagingConfigurationParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	country := d.Get("country").(string)

	r, err := m.(*client.Config).Client.VerifyV2.UpdateMessagingConfiguration(serviceSid, country, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceServicesEntitiesFactors() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesEntitiesFactors,
		ReadContext:   readServicesEntitiesFactors,
		UpdateContext: updateServicesEntitiesFactors,
		DeleteContext: deleteServicesEntitiesFactors,
		Schema: map[string]*schema.Schema{
			"service_sid":                  AsString(SchemaRequired),
			"identity":                     AsString(SchemaRequired),
			"factor_type":                  AsString(SchemaRequired),
			"friendly_name":                AsString(SchemaRequired),
			"binding_alg":                  AsString(SchemaComputedOptional),
			"binding_public_key":           AsString(SchemaComputedOptional),
			"binding_secret":               AsString(SchemaComputedOptional),
			"config_alg":                   AsString(SchemaComputedOptional),
			"config_app_id":                AsString(SchemaComputedOptional),
			"config_code_length":           AsInt(SchemaComputedOptional),
			"config_notification_platform": AsString(SchemaComputedOptional),
			"config_notification_token":    AsString(SchemaComputedOptional),
			"config_sdk_version":           AsString(SchemaComputedOptional),
			"config_skew":                  AsInt(SchemaComputedOptional),
			"config_time_step":             AsInt(SchemaComputedOptional),
			"sid":                          AsString(SchemaComputed),
			"auth_payload":                 AsString(SchemaComputedOptional),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesEntitiesFactorsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServicesEntitiesFactors(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateNewFactorParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	identity := d.Get("identity").(string)

	r, err := m.(*client.Config).Client.VerifyV2.CreateNewFactor(serviceSid, identity, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{serviceSid, identity}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))
	d.Set("sid", *r.Sid)

	return updateServicesEntitiesFactors(ctx, d, m)
}

func deleteServicesEntitiesFactors(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	identity := d.Get("identity").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.VerifyV2.DeleteFactor(serviceSid, identity, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServicesEntitiesFactors(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	identity := d.Get("identity").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.FetchFactor(serviceSid, identity, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesEntitiesFactorsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected service_sid/identity/sid"

	if len(importParts) != 3 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("service_sid", importParts[0])
	d.Set("identity", importParts[1])
	d.Set("sid", importParts[2])

	return nil
}
func updateServicesEntitiesFactors(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateFactorParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	identity := d.Get("identity").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.UpdateFactor(serviceSid, identity, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceServicesRateLimits() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesRateLimits,
		ReadContext:   readServicesRateLimits,
		UpdateContext: updateServicesRateLimits,
		DeleteContext: deleteServicesRateLimits,
		Schema: map[string]*schema.Schema{
			"service_sid": AsString(SchemaRequired),
			"unique_name": AsString(SchemaRequired),
			"description": AsString(SchemaComputedOptional),
			"sid":         AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesRateLimitsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServicesRateLimits(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateRateLimitParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.CreateRateLimit(serviceSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{serviceSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteServicesRateLimits(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.VerifyV2.DeleteRateLimit(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServicesRateLimits(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.FetchRateLimit(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesRateLimitsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected service_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("service_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func updateServicesRateLimits(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateRateLimitParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.UpdateRateLimit(serviceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceServices() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServices,
		ReadContext:   readServices,
		UpdateContext: updateServices,
		DeleteContext: deleteServices,
		Schema: map[string]*schema.Schema{
			"friendly_name":                AsString(SchemaRequired),
			"code_length":                  AsInt(SchemaComputedOptional),
			"custom_code_enabled":          AsBool(SchemaComputedOptional),
			"default_template_sid":         AsString(SchemaComputedOptional),
			"do_not_share_warning_enabled": AsBool(SchemaComputedOptional),
			"dtmf_input_required":          AsBool(SchemaComputedOptional),
			"lookup_enabled":               AsBool(SchemaComputedOptional),
			"psd2_enabled":                 AsBool(SchemaComputedOptional),
			"push_apn_credential_sid":      AsString(SchemaComputedOptional),
			"push_fcm_credential_sid":      AsString(SchemaComputedOptional),
			"push_include_date":            AsBool(SchemaComputedOptional),
			"skip_sms_to_landlines":        AsBool(SchemaComputedOptional),
			"totp_code_length":             AsInt(SchemaComputedOptional),
			"totp_issuer":                  AsString(SchemaComputedOptional),
			"totp_skew":                    AsInt(SchemaComputedOptional),
			"totp_time_step":               AsInt(SchemaComputedOptional),
			"tts_name":                     AsString(SchemaComputedOptional),
			"sid":                          AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateServiceParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.VerifyV2.CreateService(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.VerifyV2.DeleteService(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.FetchService(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func updateServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateServiceParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.UpdateService(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceServicesWebhooks() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesWebhooks,
		ReadContext:   readServicesWebhooks,
		UpdateContext: updateServicesWebhooks,
		DeleteContext: deleteServicesWebhooks,
		Schema: map[string]*schema.Schema{
			"service_sid":   AsString(SchemaRequired),
			"event_types":   AsList(AsString(SchemaRequired), SchemaRequired),
			"friendly_name": AsString(SchemaRequired),
			"webhook_url":   AsString(SchemaRequired),
			"status":        AsString(SchemaComputedOptional),
			"version":       AsString(SchemaComputedOptional),
			"sid":           AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesWebhooksImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServicesWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateWebhookParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.CreateWebhook(serviceSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{serviceSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteServicesWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.VerifyV2.DeleteWebhook(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServicesWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.FetchWebhook(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesWebhooksImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected service_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("service_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func updateServicesWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateWebhookParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.VerifyV2.UpdateWebhook(serviceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
