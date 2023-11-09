/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.2.17
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// ListOrganizationsResponse : B2B SSO Organization List



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct ListOrganizationsResponse {
    #[serde(rename = "organizations")]
    pub organizations: Vec<crate::models::Organization>,
}


impl ListOrganizationsResponse {
    /// B2B SSO Organization List
    pub fn new(organizations: Vec<crate::models::Organization>) -> ListOrganizationsResponse {
        ListOrganizationsResponse {
                organizations,
        }
    }
}


