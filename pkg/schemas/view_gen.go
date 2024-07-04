// Code generated by sdk-to-schema generator; DO NOT EDIT.

package schemas

import (
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ShowViewSchema represents output of SHOW query for the single View.
var ShowViewSchema = map[string]*schema.Schema{
	"created_on": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"name": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"kind": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"reserved": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"database_name": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"schema_name": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"owner": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"comment": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"text": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"is_secure": {
		Type:     schema.TypeBool,
		Computed: true,
	},
	"is_materialized": {
		Type:     schema.TypeBool,
		Computed: true,
	},
	"owner_role_type": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"change_tracking": {
		Type:     schema.TypeString,
		Computed: true,
	},
}

var _ = ShowViewSchema

func ViewToSchema(view *sdk.View) map[string]any {
	viewSchema := make(map[string]any)
	viewSchema["created_on"] = view.CreatedOn
	viewSchema["name"] = view.Name
	viewSchema["kind"] = view.Kind
	viewSchema["reserved"] = view.Reserved
	viewSchema["database_name"] = view.DatabaseName
	viewSchema["schema_name"] = view.SchemaName
	viewSchema["owner"] = view.Owner
	viewSchema["comment"] = view.Comment
	viewSchema["text"] = view.Text
	viewSchema["is_secure"] = view.IsSecure
	viewSchema["is_materialized"] = view.IsMaterialized
	viewSchema["owner_role_type"] = view.OwnerRoleType
	viewSchema["change_tracking"] = view.ChangeTracking
	return viewSchema
}

var _ = ViewToSchema