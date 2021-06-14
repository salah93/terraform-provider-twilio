/*
 * Twilio - Notify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/twilio/terraform-provider-twilio/client"
	. "github.com/twilio/terraform-provider-twilio/twilio/common"
	. "github.com/twilio/twilio-go/rest/notify/v1"
)

func ResourceCredentials() *schema.Resource {
	return &schema.Resource{
		CreateContext: createCredentials,
		ReadContext:   readCredentials,
		UpdateContext: updateCredentials,
		DeleteContext: deleteCredentials,
		Schema: map[string]*schema.Schema{
			"api_key":       AsString(SchemaOptional),
			"certificate":   AsString(SchemaOptional),
			"friendly_name": AsString(SchemaOptional),
			"private_key":   AsString(SchemaOptional),
			"sandbox":       AsString(SchemaOptional),
			"secret":        AsString(SchemaOptional),
			"type":          AsString(SchemaOptional),
			"account_sid":   AsString(SchemaComputed),
			"date_created":  AsString(SchemaComputed),
			"date_updated":  AsString(SchemaComputed),
			"sid":           AsString(SchemaComputed),
			"url":           AsString(SchemaComputed),
		},
	}
}

func createCredentials(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateCredentialParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.NotifyV1.CreateCredential(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.Sid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteCredentials(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.NotifyV1.DeleteCredential(sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readCredentials(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.NotifyV1.FetchCredential(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateCredentials(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateCredentialParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.NotifyV1.UpdateCredential(sid, &params)
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
			"alexa_skill_id":     AsString(SchemaOptional),
			"apn_credential_sid": AsString(SchemaOptional),
			"default_alexa_notification_protocol_version": AsString(SchemaOptional),
			"default_apn_notification_protocol_version":   AsString(SchemaOptional),
			"default_fcm_notification_protocol_version":   AsString(SchemaOptional),
			"default_gcm_notification_protocol_version":   AsString(SchemaOptional),
			"delivery_callback_enabled":                   AsString(SchemaOptional),
			"delivery_callback_url":                       AsString(SchemaOptional),
			"facebook_messenger_page_id":                  AsString(SchemaOptional),
			"fcm_credential_sid":                          AsString(SchemaOptional),
			"friendly_name":                               AsString(SchemaOptional),
			"gcm_credential_sid":                          AsString(SchemaOptional),
			"log_enabled":                                 AsString(SchemaOptional),
			"messaging_service_sid":                       AsString(SchemaOptional),
			"account_sid":                                 AsString(SchemaComputed),
			"date_created":                                AsString(SchemaComputed),
			"date_updated":                                AsString(SchemaComputed),
			"links":                                       AsString(SchemaComputed),
			"sid":                                         AsString(SchemaComputed),
			"url":                                         AsString(SchemaComputed),
		},
	}
}

func createServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateServiceParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.NotifyV1.CreateService(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.Sid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.NotifyV1.DeleteService(sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.NotifyV1.FetchService(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateServiceParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.NotifyV1.UpdateService(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
