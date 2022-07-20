# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule Ory.Model.NormalizedProjectRevision do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :created_at,
    :id,
    :keto_namespaces,
    :keto_read_max_depth,
    :kratos_cookies_same_site,
    :kratos_courier_smtp_connection_uri,
    :kratos_courier_smtp_from_address,
    :kratos_courier_smtp_from_name,
    :kratos_courier_smtp_headers,
    :kratos_courier_templates_recovery_invalid_email_body_html,
    :kratos_courier_templates_recovery_invalid_email_body_plaintext,
    :kratos_courier_templates_recovery_invalid_email_subject,
    :kratos_courier_templates_recovery_valid_email_body_html,
    :kratos_courier_templates_recovery_valid_email_body_plaintext,
    :kratos_courier_templates_recovery_valid_email_subject,
    :kratos_courier_templates_verification_invalid_email_body_html,
    :kratos_courier_templates_verification_invalid_email_body_plaintext,
    :kratos_courier_templates_verification_invalid_email_subject,
    :kratos_courier_templates_verification_valid_email_body_html,
    :kratos_courier_templates_verification_valid_email_body_plaintext,
    :kratos_courier_templates_verification_valid_email_subject,
    :kratos_identity_schemas,
    :kratos_secrets_cipher,
    :kratos_secrets_cookie,
    :kratos_secrets_default,
    :kratos_selfservice_allowed_return_urls,
    :kratos_selfservice_default_browser_return_url,
    :kratos_selfservice_flows_error_ui_url,
    :kratos_selfservice_flows_hooks,
    :kratos_selfservice_flows_login_after_default_browser_return_url,
    :kratos_selfservice_flows_login_after_oidc_default_browser_return_url,
    :kratos_selfservice_flows_login_after_password_default_browser_return_url,
    :kratos_selfservice_flows_login_after_webauthn_default_browser_return_url,
    :kratos_selfservice_flows_login_lifespan,
    :kratos_selfservice_flows_login_ui_url,
    :kratos_selfservice_flows_logout_after_default_browser_return_url,
    :kratos_selfservice_flows_recovery_after_default_browser_return_url,
    :kratos_selfservice_flows_recovery_enabled,
    :kratos_selfservice_flows_recovery_lifespan,
    :kratos_selfservice_flows_recovery_ui_url,
    :kratos_selfservice_flows_registration_after_default_browser_return_url,
    :kratos_selfservice_flows_registration_after_oidc_default_browser_return_url,
    :kratos_selfservice_flows_registration_after_password_default_browser_return_url,
    :kratos_selfservice_flows_registration_after_webauthn_default_browser_return_url,
    :kratos_selfservice_flows_registration_enabled,
    :kratos_selfservice_flows_registration_lifespan,
    :kratos_selfservice_flows_registration_ui_url,
    :kratos_selfservice_flows_settings_after_default_browser_return_url,
    :kratos_selfservice_flows_settings_after_password_default_browser_return_url,
    :kratos_selfservice_flows_settings_after_profile_default_browser_return_url,
    :kratos_selfservice_flows_settings_lifespan,
    :kratos_selfservice_flows_settings_privileged_session_max_age,
    :kratos_selfservice_flows_settings_required_aal,
    :kratos_selfservice_flows_settings_ui_url,
    :kratos_selfservice_flows_verification_after_default_browser_return_url,
    :kratos_selfservice_flows_verification_enabled,
    :kratos_selfservice_flows_verification_lifespan,
    :kratos_selfservice_flows_verification_ui_url,
    :kratos_selfservice_methods_link_config_base_url,
    :kratos_selfservice_methods_link_config_lifespan,
    :kratos_selfservice_methods_link_enabled,
    :kratos_selfservice_methods_lookup_secret_enabled,
    :kratos_selfservice_methods_oidc_config_base_redirect_uri,
    :kratos_selfservice_methods_oidc_config_providers,
    :kratos_selfservice_methods_oidc_enabled,
    :kratos_selfservice_methods_password_config_haveibeenpwned_enabled,
    :kratos_selfservice_methods_password_config_identifier_similarity_check_enabled,
    :kratos_selfservice_methods_password_config_ignore_network_errors,
    :kratos_selfservice_methods_password_config_max_breaches,
    :kratos_selfservice_methods_password_config_min_password_length,
    :kratos_selfservice_methods_password_enabled,
    :kratos_selfservice_methods_profile_enabled,
    :kratos_selfservice_methods_totp_config_issuer,
    :kratos_selfservice_methods_totp_enabled,
    :kratos_selfservice_methods_webauthn_config_passwordless,
    :kratos_selfservice_methods_webauthn_config_rp_display_name,
    :kratos_selfservice_methods_webauthn_config_rp_icon,
    :kratos_selfservice_methods_webauthn_config_rp_id,
    :kratos_selfservice_methods_webauthn_config_rp_origin,
    :kratos_selfservice_methods_webauthn_enabled,
    :kratos_session_cookie_persistent,
    :kratos_session_cookie_same_site,
    :kratos_session_lifespan,
    :kratos_session_whoami_required_aal,
    :name,
    :project_id,
    :updated_at
  ]

  @type t :: %__MODULE__{
    :created_at => DateTime.t | nil,
    :id => String.t | nil,
    :keto_namespaces => [Ory.Model.KetoNamespace.t] | nil,
    :keto_read_max_depth => integer() | nil,
    :kratos_cookies_same_site => String.t | nil,
    :kratos_courier_smtp_connection_uri => String.t | nil,
    :kratos_courier_smtp_from_address => String.t | nil,
    :kratos_courier_smtp_from_name => String.t | nil,
    :kratos_courier_smtp_headers => map() | nil,
    :kratos_courier_templates_recovery_invalid_email_body_html => String.t | nil,
    :kratos_courier_templates_recovery_invalid_email_body_plaintext => String.t | nil,
    :kratos_courier_templates_recovery_invalid_email_subject => String.t | nil,
    :kratos_courier_templates_recovery_valid_email_body_html => String.t | nil,
    :kratos_courier_templates_recovery_valid_email_body_plaintext => String.t | nil,
    :kratos_courier_templates_recovery_valid_email_subject => String.t | nil,
    :kratos_courier_templates_verification_invalid_email_body_html => String.t | nil,
    :kratos_courier_templates_verification_invalid_email_body_plaintext => String.t | nil,
    :kratos_courier_templates_verification_invalid_email_subject => String.t | nil,
    :kratos_courier_templates_verification_valid_email_body_html => String.t | nil,
    :kratos_courier_templates_verification_valid_email_body_plaintext => String.t | nil,
    :kratos_courier_templates_verification_valid_email_subject => String.t | nil,
    :kratos_identity_schemas => [Ory.Model.NormalizedProjectRevisionIdentitySchema.t] | nil,
    :kratos_secrets_cipher => [String.t] | nil,
    :kratos_secrets_cookie => [String.t] | nil,
    :kratos_secrets_default => [String.t] | nil,
    :kratos_selfservice_allowed_return_urls => [String.t] | nil,
    :kratos_selfservice_default_browser_return_url => String.t | nil,
    :kratos_selfservice_flows_error_ui_url => String.t | nil,
    :kratos_selfservice_flows_hooks => [Ory.Model.NormalizedProjectRevisionHook.t] | nil,
    :kratos_selfservice_flows_login_after_default_browser_return_url => String.t | nil,
    :kratos_selfservice_flows_login_after_oidc_default_browser_return_url => String.t | nil,
    :kratos_selfservice_flows_login_after_password_default_browser_return_url => String.t | nil,
    :kratos_selfservice_flows_login_after_webauthn_default_browser_return_url => String.t | nil,
    :kratos_selfservice_flows_login_lifespan => String.t | nil,
    :kratos_selfservice_flows_login_ui_url => String.t | nil,
    :kratos_selfservice_flows_logout_after_default_browser_return_url => String.t | nil,
    :kratos_selfservice_flows_recovery_after_default_browser_return_url => String.t | nil,
    :kratos_selfservice_flows_recovery_enabled => boolean() | nil,
    :kratos_selfservice_flows_recovery_lifespan => String.t | nil,
    :kratos_selfservice_flows_recovery_ui_url => String.t | nil,
    :kratos_selfservice_flows_registration_after_default_browser_return_url => String.t | nil,
    :kratos_selfservice_flows_registration_after_oidc_default_browser_return_url => String.t | nil,
    :kratos_selfservice_flows_registration_after_password_default_browser_return_url => String.t | nil,
    :kratos_selfservice_flows_registration_after_webauthn_default_browser_return_url => String.t | nil,
    :kratos_selfservice_flows_registration_enabled => boolean() | nil,
    :kratos_selfservice_flows_registration_lifespan => String.t | nil,
    :kratos_selfservice_flows_registration_ui_url => String.t | nil,
    :kratos_selfservice_flows_settings_after_default_browser_return_url => String.t | nil,
    :kratos_selfservice_flows_settings_after_password_default_browser_return_url => String.t | nil,
    :kratos_selfservice_flows_settings_after_profile_default_browser_return_url => String.t | nil,
    :kratos_selfservice_flows_settings_lifespan => String.t | nil,
    :kratos_selfservice_flows_settings_privileged_session_max_age => String.t | nil,
    :kratos_selfservice_flows_settings_required_aal => String.t | nil,
    :kratos_selfservice_flows_settings_ui_url => String.t | nil,
    :kratos_selfservice_flows_verification_after_default_browser_return_url => String.t | nil,
    :kratos_selfservice_flows_verification_enabled => boolean() | nil,
    :kratos_selfservice_flows_verification_lifespan => String.t | nil,
    :kratos_selfservice_flows_verification_ui_url => String.t | nil,
    :kratos_selfservice_methods_link_config_base_url => String.t | nil,
    :kratos_selfservice_methods_link_config_lifespan => String.t | nil,
    :kratos_selfservice_methods_link_enabled => boolean() | nil,
    :kratos_selfservice_methods_lookup_secret_enabled => boolean() | nil,
    :kratos_selfservice_methods_oidc_config_base_redirect_uri => String.t | nil,
    :kratos_selfservice_methods_oidc_config_providers => [Ory.Model.NormalizedProjectRevisionThirdPartyProvider.t] | nil,
    :kratos_selfservice_methods_oidc_enabled => boolean() | nil,
    :kratos_selfservice_methods_password_config_haveibeenpwned_enabled => boolean() | nil,
    :kratos_selfservice_methods_password_config_identifier_similarity_check_enabled => boolean() | nil,
    :kratos_selfservice_methods_password_config_ignore_network_errors => boolean() | nil,
    :kratos_selfservice_methods_password_config_max_breaches => integer() | nil,
    :kratos_selfservice_methods_password_config_min_password_length => integer() | nil,
    :kratos_selfservice_methods_password_enabled => boolean() | nil,
    :kratos_selfservice_methods_profile_enabled => boolean() | nil,
    :kratos_selfservice_methods_totp_config_issuer => String.t | nil,
    :kratos_selfservice_methods_totp_enabled => boolean() | nil,
    :kratos_selfservice_methods_webauthn_config_passwordless => boolean() | nil,
    :kratos_selfservice_methods_webauthn_config_rp_display_name => String.t | nil,
    :kratos_selfservice_methods_webauthn_config_rp_icon => String.t | nil,
    :kratos_selfservice_methods_webauthn_config_rp_id => String.t | nil,
    :kratos_selfservice_methods_webauthn_config_rp_origin => String.t | nil,
    :kratos_selfservice_methods_webauthn_enabled => boolean() | nil,
    :kratos_session_cookie_persistent => boolean() | nil,
    :kratos_session_cookie_same_site => String.t | nil,
    :kratos_session_lifespan => String.t | nil,
    :kratos_session_whoami_required_aal => String.t | nil,
    :name => String.t,
    :project_id => String.t | nil,
    :updated_at => DateTime.t | nil
  }
end

defimpl Poison.Decoder, for: Ory.Model.NormalizedProjectRevision do
  import Ory.Deserializer
  def decode(value, options) do
    value
    |> deserialize(:keto_namespaces, :list, Ory.Model.KetoNamespace, options)
    |> deserialize(:kratos_identity_schemas, :list, Ory.Model.NormalizedProjectRevisionIdentitySchema, options)
    |> deserialize(:kratos_selfservice_flows_hooks, :list, Ory.Model.NormalizedProjectRevisionHook, options)
    |> deserialize(:kratos_selfservice_methods_oidc_config_providers, :list, Ory.Model.NormalizedProjectRevisionThirdPartyProvider, options)
  end
end
