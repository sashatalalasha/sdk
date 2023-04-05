/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.1.24
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// SelfServiceFlowExpiredError : Is sent when a flow is expired



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct SelfServiceFlowExpiredError {
    /// The status code
    #[serde(rename = "code", skip_serializing_if = "Option::is_none")]
    pub code: Option<i64>,
    /// Debug information  This field is often not exposed to protect against leaking sensitive information.
    #[serde(rename = "debug", skip_serializing_if = "Option::is_none")]
    pub debug: Option<String>,
    /// Further error details
    #[serde(rename = "details", skip_serializing_if = "Option::is_none")]
    pub details: Option<::std::collections::HashMap<String, serde_json::Value>>,
    /// When the flow has expired
    #[serde(rename = "expired_at", skip_serializing_if = "Option::is_none")]
    pub expired_at: Option<String>,
    /// The error ID  Useful when trying to identify various errors in application logic.
    #[serde(rename = "id", skip_serializing_if = "Option::is_none")]
    pub id: Option<String>,
    /// Error message  The error's message.
    #[serde(rename = "message")]
    pub message: String,
    /// A human-readable reason for the error
    #[serde(rename = "reason", skip_serializing_if = "Option::is_none")]
    pub reason: Option<String>,
    /// The request ID  The request ID is often exposed internally in order to trace errors across service architectures. This is often a UUID.
    #[serde(rename = "request", skip_serializing_if = "Option::is_none")]
    pub request: Option<String>,
    /// A Duration represents the elapsed time between two instants as an int64 nanosecond count. The representation limits the largest representable duration to approximately 290 years.
    #[serde(rename = "since", skip_serializing_if = "Option::is_none")]
    pub since: Option<i64>,
    /// The status description
    #[serde(rename = "status", skip_serializing_if = "Option::is_none")]
    pub status: Option<String>,
    /// The flow ID that should be used for the new flow as it contains the correct messages.
    #[serde(rename = "use_flow_id", skip_serializing_if = "Option::is_none")]
    pub use_flow_id: Option<String>,
}


impl SelfServiceFlowExpiredError {
    /// Is sent when a flow is expired
    pub fn new(message: String) -> SelfServiceFlowExpiredError {
        SelfServiceFlowExpiredError {
                code: None,
                debug: None,
                details: None,
                expired_at: None,
                id: None,
                message,
                reason: None,
                request: None,
                since: None,
                status: None,
                use_flow_id: None,
        }
    }
}


