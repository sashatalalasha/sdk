# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule Ory.Model.UpdateSubscriptionBody do
  @moduledoc """
  Update Subscription Request Body
  """

  @derive [Poison.Encoder]
  defstruct [
    :plan_or_price,
    :return_to
  ]

  @type t :: %__MODULE__{
    :plan_or_price => String.t,
    :return_to => String.t | nil
  }
end

defimpl Poison.Decoder, for: Ory.Model.UpdateSubscriptionBody do
  def decode(value, _options) do
    value
  end
end
