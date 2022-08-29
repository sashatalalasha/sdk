/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.2.0-alpha.22
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct RefreshTokenHookRequest {
    /// ClientID is the identifier of the OAuth 2.0 client.
    #[serde(rename = "client_id", skip_serializing_if = "Option::is_none")]
    pub client_id: Option<String>,
    /// GrantedAudience is the list of audiences granted to the OAuth 2.0 client.
    #[serde(rename = "granted_audience", skip_serializing_if = "Option::is_none")]
    pub granted_audience: Option<Vec<String>>,
    /// GrantedScopes is the list of scopes granted to the OAuth 2.0 client.
    #[serde(rename = "granted_scopes", skip_serializing_if = "Option::is_none")]
    pub granted_scopes: Option<Vec<String>>,
    #[serde(rename = "requester", skip_serializing_if = "Option::is_none")]
    pub requester: Option<Box<crate::models::OAuth2AccessRequest>>,
    #[serde(rename = "session", skip_serializing_if = "Option::is_none")]
    pub session: Option<Box<crate::models::OAuth2ConsentSession>>,
    /// Subject is the identifier of the authenticated end-user.
    #[serde(rename = "subject", skip_serializing_if = "Option::is_none")]
    pub subject: Option<String>,
}

impl Default for RefreshTokenHookRequest {
    fn default() -> Self {
        Self::new()
    }
}

impl RefreshTokenHookRequest {
    pub fn new() -> RefreshTokenHookRequest {
        RefreshTokenHookRequest {
                client_id: None,
                granted_audience: None,
                granted_scopes: None,
                requester: None,
                session: None,
                subject: None,
        }
    }
}

