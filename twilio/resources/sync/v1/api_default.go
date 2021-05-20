/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.15.0
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
	. "github.com/twilio/twilio-go/rest/sync/v1"
)

func ResourceServices() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServices,
		ReadContext:   readServices,
		UpdateContext: updateServices,
		DeleteContext: deleteServices,
		Schema: map[string]*schema.Schema{
			"acl_enabled":                     AsString(SchemaOptional),
			"friendly_name":                   AsString(SchemaOptional),
			"reachability_debouncing_enabled": AsString(SchemaOptional),
			"reachability_debouncing_window":  AsString(SchemaOptional),
			"reachability_webhooks_enabled":   AsString(SchemaOptional),
			"webhook_url":                     AsString(SchemaOptional),
			"webhooks_from_rest_enabled":      AsString(SchemaOptional),
			"account_sid":                     AsString(SchemaComputed),
			"date_created":                    AsString(SchemaComputed),
			"date_updated":                    AsString(SchemaComputed),
			"links":                           AsString(SchemaComputed),
			"sid":                             AsString(SchemaComputed),
			"unique_name":                     AsString(SchemaComputed),
			"url":                             AsString(SchemaComputed),
		},
	}
}

func createServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateServiceParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.SyncV1.CreateService(&params)
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

	err := m.(*client.Config).Client.SyncV1.DeleteService(sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServices(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.FetchService(sid)
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

	r, err := m.(*client.Config).Client.SyncV1.UpdateService(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceServicesStreams() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesStreams,
		ReadContext:   readServicesStreams,
		UpdateContext: updateServicesStreams,
		DeleteContext: deleteServicesStreams,
		Schema: map[string]*schema.Schema{
			"service_sid":  AsString(SchemaRequired),
			"ttl":          AsString(SchemaOptional),
			"unique_name":  AsString(SchemaOptional),
			"account_sid":  AsString(SchemaComputed),
			"created_by":   AsString(SchemaComputed),
			"date_created": AsString(SchemaComputed),
			"date_expires": AsString(SchemaComputed),
			"date_updated": AsString(SchemaComputed),
			"links":        AsString(SchemaComputed),
			"sid":          AsString(SchemaComputed),
			"url":          AsString(SchemaComputed),
		},
	}
}

func createServicesStreams(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateSyncStreamParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.CreateSyncStream(serviceSid, &params)
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

func deleteServicesStreams(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.SyncV1.DeleteSyncStream(serviceSid, sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesStreams(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.FetchSyncStream(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesStreams(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateSyncStreamParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.UpdateSyncStream(serviceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceServicesDocuments() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesDocuments,
		ReadContext:   readServicesDocuments,
		UpdateContext: updateServicesDocuments,
		DeleteContext: deleteServicesDocuments,
		Schema: map[string]*schema.Schema{
			"service_sid":  AsString(SchemaRequired),
			"data":         AsString(SchemaOptional),
			"ttl":          AsString(SchemaOptional),
			"unique_name":  AsString(SchemaOptional),
			"account_sid":  AsString(SchemaComputed),
			"created_by":   AsString(SchemaComputed),
			"date_created": AsString(SchemaComputed),
			"date_expires": AsString(SchemaComputed),
			"date_updated": AsString(SchemaComputed),
			"links":        AsString(SchemaComputed),
			"revision":     AsString(SchemaComputed),
			"sid":          AsString(SchemaComputed),
			"url":          AsString(SchemaComputed),
		},
	}
}

func createServicesDocuments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateDocumentParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.CreateDocument(serviceSid, &params)
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

func deleteServicesDocuments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.SyncV1.DeleteDocument(serviceSid, sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesDocuments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.FetchDocument(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesDocuments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateDocumentParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.UpdateDocument(serviceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceServicesListsItems() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesListsItems,
		ReadContext:   readServicesListsItems,
		UpdateContext: updateServicesListsItems,
		DeleteContext: deleteServicesListsItems,
		Schema: map[string]*schema.Schema{
			"service_sid":    AsString(SchemaRequired),
			"list_sid":       AsString(SchemaRequired),
			"collection_ttl": AsString(SchemaOptional),
			"data":           AsString(SchemaOptional),
			"item_ttl":       AsString(SchemaOptional),
			"ttl":            AsString(SchemaOptional),
			"account_sid":    AsString(SchemaComputed),
			"created_by":     AsString(SchemaComputed),
			"date_created":   AsString(SchemaComputed),
			"date_expires":   AsString(SchemaComputed),
			"date_updated":   AsString(SchemaComputed),
			"index":          AsString(SchemaComputed),
			"revision":       AsString(SchemaComputed),
			"url":            AsString(SchemaComputed),
		},
	}
}

func createServicesListsItems(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateSyncListItemParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	listSid := d.Get("list_sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.CreateSyncListItem(serviceSid, listSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.ServiceSid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteServicesListsItems(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := DeleteSyncListItemParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	listSid := d.Get("list_sid").(string)
	index := d.Get("index").(string)

	err := m.(*client.Config).Client.SyncV1.DeleteSyncListItem(serviceSid, listSid, index, &params)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesListsItems(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	listSid := d.Get("list_sid").(string)
	index := d.Get("index").(string)

	r, err := m.(*client.Config).Client.SyncV1.FetchSyncListItem(serviceSid, listSid, index)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesListsItems(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateSyncListItemParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	listSid := d.Get("list_sid").(string)
	index := d.Get("index").(string)

	r, err := m.(*client.Config).Client.SyncV1.UpdateSyncListItem(serviceSid, listSid, index, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceServicesMapsItems() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesMapsItems,
		ReadContext:   readServicesMapsItems,
		UpdateContext: updateServicesMapsItems,
		DeleteContext: deleteServicesMapsItems,
		Schema: map[string]*schema.Schema{
			"service_sid":    AsString(SchemaRequired),
			"map_sid":        AsString(SchemaRequired),
			"collection_ttl": AsString(SchemaOptional),
			"data":           AsString(SchemaOptional),
			"item_ttl":       AsString(SchemaOptional),
			"key":            AsString(SchemaOptional),
			"ttl":            AsString(SchemaOptional),
			"account_sid":    AsString(SchemaComputed),
			"created_by":     AsString(SchemaComputed),
			"date_created":   AsString(SchemaComputed),
			"date_expires":   AsString(SchemaComputed),
			"date_updated":   AsString(SchemaComputed),
			"revision":       AsString(SchemaComputed),
			"url":            AsString(SchemaComputed),
		},
	}
}

func createServicesMapsItems(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateSyncMapItemParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	mapSid := d.Get("map_sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.CreateSyncMapItem(serviceSid, mapSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.ServiceSid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteServicesMapsItems(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := DeleteSyncMapItemParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	mapSid := d.Get("map_sid").(string)
	key := d.Get("key").(string)

	err := m.(*client.Config).Client.SyncV1.DeleteSyncMapItem(serviceSid, mapSid, key, &params)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesMapsItems(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	mapSid := d.Get("map_sid").(string)
	key := d.Get("key").(string)

	r, err := m.(*client.Config).Client.SyncV1.FetchSyncMapItem(serviceSid, mapSid, key)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesMapsItems(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateSyncMapItemParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	mapSid := d.Get("map_sid").(string)
	key := d.Get("key").(string)

	r, err := m.(*client.Config).Client.SyncV1.UpdateSyncMapItem(serviceSid, mapSid, key, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceServicesLists() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesLists,
		ReadContext:   readServicesLists,
		UpdateContext: updateServicesLists,
		DeleteContext: deleteServicesLists,
		Schema: map[string]*schema.Schema{
			"service_sid":    AsString(SchemaRequired),
			"collection_ttl": AsString(SchemaOptional),
			"ttl":            AsString(SchemaOptional),
			"unique_name":    AsString(SchemaOptional),
			"account_sid":    AsString(SchemaComputed),
			"created_by":     AsString(SchemaComputed),
			"date_created":   AsString(SchemaComputed),
			"date_expires":   AsString(SchemaComputed),
			"date_updated":   AsString(SchemaComputed),
			"links":          AsString(SchemaComputed),
			"revision":       AsString(SchemaComputed),
			"sid":            AsString(SchemaComputed),
			"url":            AsString(SchemaComputed),
		},
	}
}

func createServicesLists(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateSyncListParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.CreateSyncList(serviceSid, &params)
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

func deleteServicesLists(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.SyncV1.DeleteSyncList(serviceSid, sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesLists(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.FetchSyncList(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesLists(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateSyncListParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.UpdateSyncList(serviceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceServicesMaps() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesMaps,
		ReadContext:   readServicesMaps,
		UpdateContext: updateServicesMaps,
		DeleteContext: deleteServicesMaps,
		Schema: map[string]*schema.Schema{
			"service_sid":    AsString(SchemaRequired),
			"collection_ttl": AsString(SchemaOptional),
			"ttl":            AsString(SchemaOptional),
			"unique_name":    AsString(SchemaOptional),
			"account_sid":    AsString(SchemaComputed),
			"created_by":     AsString(SchemaComputed),
			"date_created":   AsString(SchemaComputed),
			"date_expires":   AsString(SchemaComputed),
			"date_updated":   AsString(SchemaComputed),
			"links":          AsString(SchemaComputed),
			"revision":       AsString(SchemaComputed),
			"sid":            AsString(SchemaComputed),
			"url":            AsString(SchemaComputed),
		},
	}
}

func createServicesMaps(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateSyncMapParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.CreateSyncMap(serviceSid, &params)
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

func deleteServicesMaps(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.SyncV1.DeleteSyncMap(serviceSid, sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesMaps(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.FetchSyncMap(serviceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesMaps(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateSyncMapParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.SyncV1.UpdateSyncMap(serviceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}