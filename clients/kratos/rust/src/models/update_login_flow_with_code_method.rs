/*
 * Ory Identities API
 *
 * This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more. 
 *
 * The version of the OpenAPI document: v1.1.0
 * Contact: office@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// UpdateLoginFlowWithCodeMethod : Update Login flow using the code method



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct UpdateLoginFlowWithCodeMethod {
    /// Code is the 6 digits code sent to the user
    #[serde(rename = "code", skip_serializing_if = "Option::is_none")]
    pub code: Option<String>,
    /// CSRFToken is the anti-CSRF token
    #[serde(rename = "csrf_token")]
    pub csrf_token: String,
    /// Identifier is the code identifier The identifier requires that the user has already completed the registration or settings with code flow.
    #[serde(rename = "identifier", skip_serializing_if = "Option::is_none")]
    pub identifier: Option<String>,
    /// Method should be set to \"code\" when logging in using the code strategy.
    #[serde(rename = "method")]
    pub method: String,
    /// Resend is set when the user wants to resend the code
    #[serde(rename = "resend", skip_serializing_if = "Option::is_none")]
    pub resend: Option<String>,
}


impl UpdateLoginFlowWithCodeMethod {
    /// Update Login flow using the code method
    pub fn new(csrf_token: String, method: String) -> UpdateLoginFlowWithCodeMethod {
        UpdateLoginFlowWithCodeMethod {
                code: None,
                csrf_token,
                identifier: None,
                method,
                resend: None,
        }
    }
}

