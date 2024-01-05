# NOTE: This file is auto generated by OpenAPI Generator 7.2.0 (https://openapi-generator.tech).
# Do not edit this file manually.

defmodule Ory.Model.OAuth2ConsentRequest do
  @moduledoc """
  
  """

  @derive Jason.Encoder
  defstruct [
    :acr,
    :amr,
    :challenge,
    :client,
    :context,
    :login_challenge,
    :login_session_id,
    :oidc_context,
    :request_url,
    :requested_access_token_audience,
    :requested_scope,
    :skip,
    :subject
  ]

  @type t :: %__MODULE__{
    :acr => String.t | nil,
    :amr => [String.t] | nil,
    :challenge => String.t,
    :client => Ory.Model.OAuth2Client.t | nil,
    :context => map() | nil,
    :login_challenge => String.t | nil,
    :login_session_id => String.t | nil,
    :oidc_context => Ory.Model.OAuth2ConsentRequestOpenIdConnectContext.t | nil,
    :request_url => String.t | nil,
    :requested_access_token_audience => [String.t] | nil,
    :requested_scope => [String.t] | nil,
    :skip => boolean() | nil,
    :subject => String.t | nil
  }

  alias Ory.Deserializer

  def decode(value) do
    value
     |> Deserializer.deserialize(:client, :struct, Ory.Model.OAuth2Client)
     |> Deserializer.deserialize(:oidc_context, :struct, Ory.Model.OAuth2ConsentRequestOpenIdConnectContext)
  end
end

