# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule Ory.Model.CourierMessageStatus do
  @moduledoc """
  A Message's Status
  """

  @derive [Poison.Encoder]
  defstruct [
    
  ]

  @type t :: %__MODULE__{
    
  }
end

defimpl Poison.Decoder, for: Ory.Model.CourierMessageStatus do
  def decode(value, _options) do
    value
  end
end

