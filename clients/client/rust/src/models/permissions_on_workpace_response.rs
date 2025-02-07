/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.8.1
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// PermissionsOnWorkpaceResponse : Get Permissions on Project Request Parameters



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct PermissionsOnWorkpaceResponse {
    #[serde(rename = "permissions", skip_serializing_if = "Option::is_none")]
    pub permissions: Option<::std::collections::HashMap<String, bool>>,
}

impl Default for PermissionsOnWorkpaceResponse {
    fn default() -> Self {
        Self::new()
    }
}

impl PermissionsOnWorkpaceResponse {
    /// Get Permissions on Project Request Parameters
    pub fn new() -> PermissionsOnWorkpaceResponse {
        PermissionsOnWorkpaceResponse {
                permissions: None,
        }
    }
}


