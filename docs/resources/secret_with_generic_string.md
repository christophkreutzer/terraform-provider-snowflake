---
page_title: "snowflake_secret_with_generic_string Resource - terraform-provider-snowflake"
subcategory: ""
description: |-
  Resource used to manage secret objects with Generic String. For more information, check secret documentation https://docs.snowflake.com/en/sql-reference/sql/create-secret.
---

!> **V1 release candidate** This resource is a release candidate for the V1. It is on the list of remaining GA objects for V1. We do not expect significant changes in it before the V1. We will welcome any feedback and adjust the resource if needed. Any errors reported will be resolved with a higher priority. We encourage checking this resource out before the V1 release. Please follow the [migration guide](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/MIGRATION_GUIDE.md#v0970--v0980) to use it.

# snowflake_secret_with_generic_string (Resource)

Resource used to manage secret objects with Generic String. For more information, check [secret documentation](https://docs.snowflake.com/en/sql-reference/sql/create-secret).

## Example Usage

```terraform
# basic resource
resource "snowflake_secret_with_generic_string" "test" {
  name          = "EXAMPLE_SECRET"
  database      = "EXAMPLE_DB"
  schema        = "EXAMPLE_SCHEMA"
  secret_string = "EXAMPLE_SECRET_STRING"
}

# resource with all fields set
resource "snowflake_secret_with_generic_string" "test" {
  name          = "EXAMPLE_SECRET"
  database      = "EXAMPLE_DB"
  schema        = "EXAMPLE_SCHEMA"
  secret_string = "EXAMPLE_SECRET_STRING"
  comment       = "EXAMPLE_COMMENT"
}
```
-> **Note** Instead of using fully_qualified_name, you can reference objects managed outside Terraform by constructing a correct ID, consult [identifiers guide](https://registry.terraform.io/providers/Snowflake-Labs/snowflake/latest/docs/guides/identifiers#new-computed-fully-qualified-name-field-in-resources).
<!-- TODO(SNOW-1634854): include an example showing both methods-->

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `database` (String) The database in which to create the secret Due to technical limitations (read more [here](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/docs/technical-documentation/identifiers_rework_design_decisions.md#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `(`, `)`, `"`
- `name` (String) String that specifies the identifier (i.e. name) for the secret, must be unique in your schema. Due to technical limitations (read more [here](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/docs/technical-documentation/identifiers_rework_design_decisions.md#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `(`, `)`, `"`
- `schema` (String) The schema in which to create the secret. Due to technical limitations (read more [here](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/docs/technical-documentation/identifiers_rework_design_decisions.md#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `(`, `)`, `"`
- `secret_string` (String, Sensitive) Specifies the string to store in the secret. The string can be an API token or a string of sensitive value that can be used in the handler code of a UDF or stored procedure. For details, see [Creating and using an external access integration](https://docs.snowflake.com/en/developer-guide/external-network-access/creating-using-external-network-access). You should not use this property to store any kind of OAuth token; use one of the other secret types for your OAuth use cases. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource manually using "terraform taint".

### Optional

- `comment` (String) Specifies a comment for the secret.

### Read-Only

- `describe_output` (List of Object) Outputs the result of `DESCRIBE SECRET` for the given secret. (see [below for nested schema](#nestedatt--describe_output))
- `fully_qualified_name` (String) Fully qualified name of the resource. For more information, see [object name resolution](https://docs.snowflake.com/en/sql-reference/name-resolution).
- `id` (String) The ID of this resource.
- `secret_type` (String) Specifies a type for the secret. This field is used for checking external changes and recreating the resources if needed.
- `show_output` (List of Object) Outputs the result of `SHOW SECRETS` for the given secret. (see [below for nested schema](#nestedatt--show_output))

<a id="nestedatt--describe_output"></a>
### Nested Schema for `describe_output`

Read-Only:

- `comment` (String)
- `created_on` (String)
- `database_name` (String)
- `integration_name` (String)
- `name` (String)
- `oauth_access_token_expiry_time` (String)
- `oauth_refresh_token_expiry_time` (String)
- `oauth_scopes` (Set of String)
- `owner` (String)
- `schema_name` (String)
- `secret_type` (String)
- `username` (String)


<a id="nestedatt--show_output"></a>
### Nested Schema for `show_output`

Read-Only:

- `comment` (String)
- `created_on` (String)
- `database_name` (String)
- `name` (String)
- `oauth_scopes` (Set of String)
- `owner` (String)
- `owner_role_type` (String)
- `schema_name` (String)
- `secret_type` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import snowflake_secret_with_generic_string.example '"<database_name>"."<schema_name>"."<secret_name>"'
```