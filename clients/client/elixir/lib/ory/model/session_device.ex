# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule Ory.Model.SessionDevice do
  @moduledoc """
  Device corresponding to a Session
  """

  @derive [Poison.Encoder]
  defstruct [
    :id,
    :ip_address,
    :location,
    :user_agent
  ]

  @type t :: %__MODULE__{
    :id => String.t,
    :ip_address => String.t | nil,
    :location => String.t | nil,
    :user_agent => String.t | nil
  }
end

defimpl Poison.Decoder, for: Ory.Model.SessionDevice do
  def decode(value, _options) do
    value
  end
end

