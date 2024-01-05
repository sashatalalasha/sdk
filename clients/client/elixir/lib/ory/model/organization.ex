# NOTE: This file is auto generated by OpenAPI Generator 7.2.0 (https://openapi-generator.tech).
# Do not edit this file manually.

defmodule Ory.Model.Organization do
  @moduledoc """
  B2B SSO Organization
  """

  @derive Jason.Encoder
  defstruct [
    :created_at,
    :domains,
    :id,
    :label,
    :project_id,
    :updated_at
  ]

  @type t :: %__MODULE__{
    :created_at => DateTime.t,
    :domains => [String.t],
    :id => String.t,
    :label => String.t,
    :project_id => String.t,
    :updated_at => DateTime.t
  }

  alias Ory.Deserializer

  def decode(value) do
    value
     |> Deserializer.deserialize(:created_at, :datetime, nil)
     |> Deserializer.deserialize(:updated_at, :datetime, nil)
  end
end

