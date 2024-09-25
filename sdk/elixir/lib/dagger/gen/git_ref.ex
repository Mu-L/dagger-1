# This file generated by `dagger_codegen`. Please DO NOT EDIT.
defmodule Dagger.GitRef do
  @moduledoc "A git ref (tag, branch, or commit)."

  alias Dagger.Core.Client
  alias Dagger.Core.QueryBuilder, as: QB

  @derive Dagger.ID

  defstruct [:query_builder, :client]

  @type t() :: %__MODULE__{}

  @doc "The resolved commit id at this ref."
  @spec commit(t()) :: {:ok, String.t()} | {:error, term()}
  def commit(%__MODULE__{} = git_ref) do
    query_builder =
      git_ref.query_builder |> QB.select("commit")

    Client.execute(git_ref.client, query_builder)
  end

  @doc "A unique identifier for this GitRef."
  @spec id(t()) :: {:ok, Dagger.GitRefID.t()} | {:error, term()}
  def id(%__MODULE__{} = git_ref) do
    query_builder =
      git_ref.query_builder |> QB.select("id")

    Client.execute(git_ref.client, query_builder)
  end

  @doc "The filesystem tree at this ref."
  @spec tree(t(), [{:discard_git_dir, boolean() | nil}]) :: Dagger.Directory.t()
  def tree(%__MODULE__{} = git_ref, optional_args \\ []) do
    query_builder =
      git_ref.query_builder
      |> QB.select("tree")
      |> QB.maybe_put_arg("discardGitDir", optional_args[:discard_git_dir])

    %Dagger.Directory{
      query_builder: query_builder,
      client: git_ref.client
    }
  end
end
