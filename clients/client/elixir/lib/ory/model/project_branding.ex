# NOTE: This file is auto generated by OpenAPI Generator 7.2.0 (https://openapi-generator.tech).
# Do not edit this file manually.

defmodule Ory.Model.ProjectBranding do
  @moduledoc """
  
  """

  @derive Jason.Encoder
  defstruct [
    :created_at,
    :default_theme,
    :id,
    :project_id,
    :themes,
    :updated_at
  ]

  @type t :: %__MODULE__{
    :created_at => DateTime.t,
    :default_theme => Ory.Model.ProjectBrandingTheme.t,
    :id => String.t,
    :project_id => String.t,
    :themes => [Ory.Model.ProjectBrandingTheme.t],
    :updated_at => DateTime.t
  }

  alias Ory.Deserializer

  def decode(value) do
    value
     |> Deserializer.deserialize(:created_at, :datetime, nil)
     |> Deserializer.deserialize(:default_theme, :struct, Ory.Model.ProjectBrandingTheme)
     |> Deserializer.deserialize(:themes, :list, Ory.Model.ProjectBrandingTheme)
     |> Deserializer.deserialize(:updated_at, :datetime, nil)
  end
end

