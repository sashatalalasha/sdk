# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule Ory.Model.SuccessfulOAuth2RequestResponse do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :redirect_to
  ]

  @type t :: %__MODULE__{
    :redirect_to => String.t
  }
end

defimpl Poison.Decoder, for: Ory.Model.SuccessfulOAuth2RequestResponse do
  def decode(value, _options) do
    value
  end
end

