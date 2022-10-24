# NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
# https://openapi-generator.tech
# Do not edit the class manually.

defmodule Ory.Api.Write do
  @moduledoc """
  API calls for all endpoints tagged `Write`.
  """

  alias Ory.Connection
  import Ory.RequestBuilder


  @doc """
  Create a Relation Tuple
  Use this endpoint to create a relation tuple.

  ## Parameters

  - connection (Ory.Connection): Connection to server
  - opts (KeywordList): [optional] Optional parameters
    - :body (RelationQuery): 
  ## Returns

  {:ok, Ory.Model.RelationQuery.t} on success
  {:error, Tesla.Env.t} on failure
  """
  @spec create_relation_tuple(Tesla.Env.client, keyword()) :: {:ok, Ory.Model.GenericError.t} | {:ok, Ory.Model.RelationQuery.t} | {:error, Tesla.Env.t}
  def create_relation_tuple(connection, opts \\ []) do
    optional_params = %{
      :body => :body
    }
    %{}
    |> method(:put)
    |> url("/admin/relation-tuples")
    |> add_optional_params(optional_params, opts)
    |> ensure_body()
    |> Enum.into([])
    |> (&Connection.request(connection, &1)).()
    |> evaluate_response([
      { 201, %Ory.Model.RelationQuery{}},
      { 400, %Ory.Model.GenericError{}},
      { 500, %Ory.Model.GenericError{}}
    ])
  end

  @doc """
  Delete Relation Tuples
  Use this endpoint to delete relation tuples

  ## Parameters

  - connection (Ory.Connection): Connection to server
  - opts (KeywordList): [optional] Optional parameters
    - :namespace (String.t): Namespace of the Relation Tuple
    - :object (String.t): Object of the Relation Tuple
    - :relation (String.t): Relation of the Relation Tuple
    - :subject_id (String.t): SubjectID of the Relation Tuple
    - :subject_set_periodnamespace (String.t): Namespace of the Subject Set
    - :subject_set_periodobject (String.t): Object of the Subject Set
    - :subject_set_periodrelation (String.t): Relation of the Subject Set
  ## Returns

  {:ok, nil} on success
  {:error, Tesla.Env.t} on failure
  """
  @spec delete_relation_tuples(Tesla.Env.client, keyword()) :: {:ok, nil} | {:ok, Ory.Model.GenericError.t} | {:error, Tesla.Env.t}
  def delete_relation_tuples(connection, opts \\ []) do
    optional_params = %{
      :namespace => :query,
      :object => :query,
      :relation => :query,
      :subject_id => :query,
      :"subject_set.namespace" => :query,
      :"subject_set.object" => :query,
      :"subject_set.relation" => :query
    }
    %{}
    |> method(:delete)
    |> url("/admin/relation-tuples")
    |> add_optional_params(optional_params, opts)
    |> Enum.into([])
    |> (&Connection.request(connection, &1)).()
    |> evaluate_response([
      { 204, false},
      { 400, %Ory.Model.GenericError{}},
      { 500, %Ory.Model.GenericError{}}
    ])
  end

  @doc """
  Patch Multiple Relation Tuples
  Use this endpoint to patch one or more relation tuples.

  ## Parameters

  - connection (Ory.Connection): Connection to server
  - opts (KeywordList): [optional] Optional parameters
    - :body ([Ory.Model.PatchDelta.t]): 
  ## Returns

  {:ok, nil} on success
  {:error, Tesla.Env.t} on failure
  """
  @spec patch_relation_tuples(Tesla.Env.client, keyword()) :: {:ok, nil} | {:ok, Ory.Model.GenericError.t} | {:error, Tesla.Env.t}
  def patch_relation_tuples(connection, opts \\ []) do
    optional_params = %{
      :body => :body
    }
    %{}
    |> method(:patch)
    |> url("/admin/relation-tuples")
    |> add_optional_params(optional_params, opts)
    |> ensure_body()
    |> Enum.into([])
    |> (&Connection.request(connection, &1)).()
    |> evaluate_response([
      { 204, false},
      { 400, %Ory.Model.GenericError{}},
      { 404, %Ory.Model.GenericError{}},
      { 500, %Ory.Model.GenericError{}}
    ])
  end
end
