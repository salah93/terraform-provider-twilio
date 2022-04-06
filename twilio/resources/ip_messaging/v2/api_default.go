/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.1
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
	. "github.com/twilio/twilio-go/rest/ip_messaging/v2"
)

func ResourceServicesChannels() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesChannels,
		ReadContext:   readServicesChannels,
		UpdateContext: updateServicesChannels,
		DeleteContext: deleteServicesChannels,
		Schema: map[string]*schema.Schema{
			"service_sid":              AsString(SchemaRequired),
			"x_twilio_webhook_enabled": AsString(SchemaComputedOptional),
			"attributes":               AsString(SchemaComputedOptional),
			"created_by":               AsString(SchemaComputedOptional),
			"date_created":             AsString(SchemaComputedOptional),
			"date_updated":             AsString(SchemaComputedOptional),
			"friendly_name":            AsString(SchemaComputedOptional),
			"type":                     AsString(SchemaComputedOptional),
			"unique_name":              AsString(SchemaComputedOptional),
			"sid":                      AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesChannelsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServicesChannels(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateChannelParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.CreateChannel(serviceSid, &params)
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

func deleteServicesChannels(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := DeleteChannelParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.IpMessagingV2.DeleteChannel(serviceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServicesChannels(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.FetchChannel(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesChannelsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected service_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("service_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func updateServicesChannels(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateChannelParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.UpdateChannel(serviceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceServicesChannelsWebhooks() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesChannelsWebhooks,
		ReadContext:   readServicesChannelsWebhooks,
		UpdateContext: updateServicesChannelsWebhooks,
		DeleteContext: deleteServicesChannelsWebhooks,
		Schema: map[string]*schema.Schema{
			"service_sid":               AsString(SchemaRequired),
			"channel_sid":               AsString(SchemaRequired),
			"type":                      AsString(SchemaRequired),
			"configuration_filters":     AsList(AsString(SchemaComputedOptional), SchemaComputedOptional),
			"configuration_flow_sid":    AsString(SchemaComputedOptional),
			"configuration_method":      AsString(SchemaComputedOptional),
			"configuration_retry_count": AsInt(SchemaComputedOptional),
			"configuration_triggers":    AsList(AsString(SchemaComputedOptional), SchemaComputedOptional),
			"configuration_url":         AsString(SchemaComputedOptional),
			"sid":                       AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesChannelsWebhooksImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServicesChannelsWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateChannelWebhookParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.CreateChannelWebhook(serviceSid, channelSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{serviceSid, channelSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteServicesChannelsWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.IpMessagingV2.DeleteChannelWebhook(serviceSid, channelSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServicesChannelsWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.FetchChannelWebhook(serviceSid, channelSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesChannelsWebhooksImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected service_sid/channel_sid/sid"

	if len(importParts) != 3 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("service_sid", importParts[0])
	d.Set("channel_sid", importParts[1])
	d.Set("sid", importParts[2])

	return nil
}
func updateServicesChannelsWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateChannelWebhookParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.UpdateChannelWebhook(serviceSid, channelSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceCredentials() *schema.Resource {
	return &schema.Resource{
		CreateContext: createCredentials,
		ReadContext:   readCredentials,
		UpdateContext: updateCredentials,
		DeleteContext: deleteCredentials,
		Schema: map[string]*schema.Schema{
			"type":          AsString(SchemaRequired),
			"api_key":       AsString(SchemaComputedOptional),
			"certificate":   AsString(SchemaComputedOptional),
			"friendly_name": AsString(SchemaComputedOptional),
			"private_key":   AsString(SchemaComputedOptional),
			"sandbox":       AsBool(SchemaComputedOptional),
			"secret":        AsString(SchemaComputedOptional),
			"sid":           AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseCredentialsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createCredentials(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateCredentialParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.IpMessagingV2.CreateCredential(&params)
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

func deleteCredentials(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.IpMessagingV2.DeleteCredential(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readCredentials(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.FetchCredential(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseCredentialsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func updateCredentials(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateCredentialParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.UpdateCredential(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceServicesChannelsInvites() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesChannelsInvites,
		ReadContext:   readServicesChannelsInvites,
		DeleteContext: deleteServicesChannelsInvites,
		Schema: map[string]*schema.Schema{
			"service_sid": AsString(SchemaForceNewRequired),
			"channel_sid": AsString(SchemaForceNewRequired),
			"identity":    AsString(SchemaForceNewRequired),
			"role_sid":    AsString(SchemaForceNewOptional),
			"sid":         AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesChannelsInvitesImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServicesChannelsInvites(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateInviteParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.CreateInvite(serviceSid, channelSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{serviceSid, channelSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteServicesChannelsInvites(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.IpMessagingV2.DeleteInvite(serviceSid, channelSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServicesChannelsInvites(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.FetchInvite(serviceSid, channelSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesChannelsInvitesImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected service_sid/channel_sid/sid"

	if len(importParts) != 3 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("service_sid", importParts[0])
	d.Set("channel_sid", importParts[1])
	d.Set("sid", importParts[2])

	return nil
}
func ResourceServicesChannelsMembers() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesChannelsMembers,
		ReadContext:   readServicesChannelsMembers,
		UpdateContext: updateServicesChannelsMembers,
		DeleteContext: deleteServicesChannelsMembers,
		Schema: map[string]*schema.Schema{
			"service_sid":                 AsString(SchemaRequired),
			"channel_sid":                 AsString(SchemaRequired),
			"identity":                    AsString(SchemaRequired),
			"x_twilio_webhook_enabled":    AsString(SchemaComputedOptional),
			"attributes":                  AsString(SchemaComputedOptional),
			"date_created":                AsString(SchemaComputedOptional),
			"date_updated":                AsString(SchemaComputedOptional),
			"last_consumed_message_index": AsInt(SchemaComputedOptional),
			"last_consumption_timestamp":  AsString(SchemaComputedOptional),
			"role_sid":                    AsString(SchemaComputedOptional),
			"sid":                         AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesChannelsMembersImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServicesChannelsMembers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateMemberParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.CreateMember(serviceSid, channelSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{serviceSid, channelSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteServicesChannelsMembers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := DeleteMemberParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.IpMessagingV2.DeleteMember(serviceSid, channelSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServicesChannelsMembers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.FetchMember(serviceSid, channelSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesChannelsMembersImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected service_sid/channel_sid/sid"

	if len(importParts) != 3 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("service_sid", importParts[0])
	d.Set("channel_sid", importParts[1])
	d.Set("sid", importParts[2])

	return nil
}
func updateServicesChannelsMembers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateMemberParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.UpdateMember(serviceSid, channelSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceServicesChannelsMessages() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesChannelsMessages,
		ReadContext:   readServicesChannelsMessages,
		UpdateContext: updateServicesChannelsMessages,
		DeleteContext: deleteServicesChannelsMessages,
		Schema: map[string]*schema.Schema{
			"service_sid":              AsString(SchemaRequired),
			"channel_sid":              AsString(SchemaRequired),
			"x_twilio_webhook_enabled": AsString(SchemaComputedOptional),
			"attributes":               AsString(SchemaComputedOptional),
			"body":                     AsString(SchemaComputedOptional),
			"date_created":             AsString(SchemaComputedOptional),
			"date_updated":             AsString(SchemaComputedOptional),
			"from":                     AsString(SchemaComputedOptional),
			"last_updated_by":          AsString(SchemaComputedOptional),
			"media_sid":                AsString(SchemaComputedOptional),
			"sid":                      AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesChannelsMessagesImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServicesChannelsMessages(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateMessageParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.CreateMessage(serviceSid, channelSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{serviceSid, channelSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteServicesChannelsMessages(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := DeleteMessageParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.IpMessagingV2.DeleteMessage(serviceSid, channelSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServicesChannelsMessages(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.FetchMessage(serviceSid, channelSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesChannelsMessagesImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected service_sid/channel_sid/sid"

	if len(importParts) != 3 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("service_sid", importParts[0])
	d.Set("channel_sid", importParts[1])
	d.Set("sid", importParts[2])

	return nil
}
func updateServicesChannelsMessages(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateMessageParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.UpdateMessage(serviceSid, channelSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceServicesRoles() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesRoles,
		ReadContext:   readServicesRoles,
		UpdateContext: updateServicesRoles,
		DeleteContext: deleteServicesRoles,
		Schema: map[string]*schema.Schema{
			"service_sid":   AsString(SchemaRequired),
			"friendly_name": AsString(SchemaRequired),
			"permission":    AsList(AsString(SchemaRequired), SchemaRequired),
			"type":          AsString(SchemaRequired),
			"sid":           AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesRolesImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServicesRoles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateRoleParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.CreateRole(serviceSid, &params)
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

func deleteServicesRoles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.IpMessagingV2.DeleteRole(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServicesRoles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.FetchRole(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesRolesImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected service_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("service_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func updateServicesRoles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateRoleParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.UpdateRole(serviceSid, sid, &params)
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
			"friendly_name":                                 AsString(SchemaRequired),
			"sid":                                           AsString(SchemaComputed),
			"consumption_report_interval":                   AsInt(SchemaComputedOptional),
			"default_channel_creator_role_sid":              AsString(SchemaComputedOptional),
			"default_channel_role_sid":                      AsString(SchemaComputedOptional),
			"default_service_role_sid":                      AsString(SchemaComputedOptional),
			"limits_channel_members":                        AsInt(SchemaComputedOptional),
			"limits_user_channels":                          AsInt(SchemaComputedOptional),
			"media_compatibility_message":                   AsString(SchemaComputedOptional),
			"notifications_added_to_channel_enabled":        AsBool(SchemaComputedOptional),
			"notifications_added_to_channel_sound":          AsString(SchemaComputedOptional),
			"notifications_added_to_channel_template":       AsString(SchemaComputedOptional),
			"notifications_invited_to_channel_enabled":      AsBool(SchemaComputedOptional),
			"notifications_invited_to_channel_sound":        AsString(SchemaComputedOptional),
			"notifications_invited_to_channel_template":     AsString(SchemaComputedOptional),
			"notifications_log_enabled":                     AsBool(SchemaComputedOptional),
			"notifications_new_message_badge_count_enabled": AsBool(SchemaComputedOptional),
			"notifications_new_message_enabled":             AsBool(SchemaComputedOptional),
			"notifications_new_message_sound":               AsString(SchemaComputedOptional),
			"notifications_new_message_template":            AsString(SchemaComputedOptional),
			"notifications_removed_from_channel_enabled":    AsBool(SchemaComputedOptional),
			"notifications_removed_from_channel_sound":      AsString(SchemaComputedOptional),
			"notifications_removed_from_channel_template":   AsString(SchemaComputedOptional),
			"post_webhook_retry_count":                      AsInt(SchemaComputedOptional),
			"post_webhook_url":                              AsString(SchemaComputedOptional),
			"pre_webhook_retry_count":                       AsInt(SchemaComputedOptional),
			"pre_webhook_url":                               AsString(SchemaComputedOptional),
			"reachability_enabled":                          AsBool(SchemaComputedOptional),
			"read_status_enabled":                           AsBool(SchemaComputedOptional),
			"typing_indicator_timeout":                      AsInt(SchemaComputedOptional),
			"webhook_filters":                               AsList(AsString(SchemaComputedOptional), SchemaComputedOptional),
			"webhook_method":                                AsString(SchemaComputedOptional),
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

	r, err := m.(*client.Config).Client.IpMessagingV2.CreateService(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))
	d.Set("sid", *r.Sid)

	return updateServices(ctx, d, m)
}

func deleteServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.IpMessagingV2.DeleteService(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.FetchService(sid)
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

	r, err := m.(*client.Config).Client.IpMessagingV2.UpdateService(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceServicesUsers() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesUsers,
		ReadContext:   readServicesUsers,
		UpdateContext: updateServicesUsers,
		DeleteContext: deleteServicesUsers,
		Schema: map[string]*schema.Schema{
			"service_sid":              AsString(SchemaRequired),
			"identity":                 AsString(SchemaRequired),
			"x_twilio_webhook_enabled": AsString(SchemaComputedOptional),
			"attributes":               AsString(SchemaComputedOptional),
			"friendly_name":            AsString(SchemaComputedOptional),
			"role_sid":                 AsString(SchemaComputedOptional),
			"sid":                      AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseServicesUsersImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createServicesUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateUserParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.CreateUser(serviceSid, &params)
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

func deleteServicesUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.IpMessagingV2.DeleteUser(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readServicesUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.FetchUser(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseServicesUsersImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected service_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("service_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func updateServicesUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateUserParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.IpMessagingV2.UpdateUser(serviceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
