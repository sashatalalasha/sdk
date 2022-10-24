# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule Ory.Model.SourcePosition do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :Line,
    :column
  ]

  @type t :: %__MODULE__{
    :Line => integer() | nil,
    :column => integer() | nil
  }
end

defimpl Poison.Decoder, for: Ory.Model.SourcePosition do
  def decode(value, _options) do
    value
  end
end

