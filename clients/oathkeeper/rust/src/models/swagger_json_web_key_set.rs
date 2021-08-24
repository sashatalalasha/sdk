/*
 * ORY Oathkeeper
 *
 * ORY Oathkeeper is a reverse proxy that checks the HTTP Authorization for validity against a set of rules. This service uses Hydra to validate access tokens and policies.
 *
 * The version of the OpenAPI document: v0.0.0-alpha.54
 * Contact: hi@ory.am
 * Generated by: https://openapi-generator.tech
 */

/// SwaggerJsonWebKeySet : SwaggerJSONWebKeySet swagger JSON web key set



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct SwaggerJsonWebKeySet {
    /// The value of the \"keys\" parameter is an array of JWK values.  By default, the order of the JWK values within the array does not imply an order of preference among them, although applications of JWK Sets can choose to assign a meaning to the order for their purposes, if desired.
    #[serde(rename = "keys", skip_serializing_if = "Option::is_none")]
    pub keys: Option<Vec<crate::models::SwaggerJsonWebKey>>,
}

impl SwaggerJsonWebKeySet {
    /// SwaggerJSONWebKeySet swagger JSON web key set
    pub fn new() -> SwaggerJsonWebKeySet {
        SwaggerJsonWebKeySet {
            keys: None,
        }
    }
}


