# This file generated by `dagger_codegen`. Please DO NOT EDIT.
defmodule Dagger.Binding do
  @moduledoc """
  Dagger.Binding
  """

  use Dagger.Core.Base, kind: :object, name: "Binding"

  alias Dagger.Core.Client
  alias Dagger.Core.QueryBuilder, as: QB

  @derive Dagger.ID

  defstruct [:query_builder, :client]

  @type t() :: %__MODULE__{}

  @doc """
  Retrieve the binding value, as type CacheVolume
  """
  @spec as_cache_volume(t()) :: Dagger.CacheVolume.t()
  def as_cache_volume(%__MODULE__{} = binding) do
    query_builder =
      binding.query_builder |> QB.select("asCacheVolume")

    %Dagger.CacheVolume{
      query_builder: query_builder,
      client: binding.client
    }
  end

  @doc """
  Retrieve the binding value, as type Cloud
  """
  @spec as_cloud(t()) :: Dagger.Cloud.t()
  def as_cloud(%__MODULE__{} = binding) do
    query_builder =
      binding.query_builder |> QB.select("asCloud")

    %Dagger.Cloud{
      query_builder: query_builder,
      client: binding.client
    }
  end

  @doc """
  Retrieve the binding value, as type Container
  """
  @spec as_container(t()) :: Dagger.Container.t()
  def as_container(%__MODULE__{} = binding) do
    query_builder =
      binding.query_builder |> QB.select("asContainer")

    %Dagger.Container{
      query_builder: query_builder,
      client: binding.client
    }
  end

  @doc """
  Retrieve the binding value, as type Directory
  """
  @spec as_directory(t()) :: Dagger.Directory.t()
  def as_directory(%__MODULE__{} = binding) do
    query_builder =
      binding.query_builder |> QB.select("asDirectory")

    %Dagger.Directory{
      query_builder: query_builder,
      client: binding.client
    }
  end

  @doc """
  Retrieve the binding value, as type Env
  """
  @spec as_env(t()) :: Dagger.Env.t()
  def as_env(%__MODULE__{} = binding) do
    query_builder =
      binding.query_builder |> QB.select("asEnv")

    %Dagger.Env{
      query_builder: query_builder,
      client: binding.client
    }
  end

  @doc """
  Retrieve the binding value, as type File
  """
  @spec as_file(t()) :: Dagger.File.t()
  def as_file(%__MODULE__{} = binding) do
    query_builder =
      binding.query_builder |> QB.select("asFile")

    %Dagger.File{
      query_builder: query_builder,
      client: binding.client
    }
  end

  @doc """
  Retrieve the binding value, as type GitRef
  """
  @spec as_git_ref(t()) :: Dagger.GitRef.t()
  def as_git_ref(%__MODULE__{} = binding) do
    query_builder =
      binding.query_builder |> QB.select("asGitRef")

    %Dagger.GitRef{
      query_builder: query_builder,
      client: binding.client
    }
  end

  @doc """
  Retrieve the binding value, as type GitRepository
  """
  @spec as_git_repository(t()) :: Dagger.GitRepository.t()
  def as_git_repository(%__MODULE__{} = binding) do
    query_builder =
      binding.query_builder |> QB.select("asGitRepository")

    %Dagger.GitRepository{
      query_builder: query_builder,
      client: binding.client
    }
  end

  @doc """
  Retrieve the binding value, as type LLM
  """
  @spec as_llm(t()) :: Dagger.LLM.t()
  def as_llm(%__MODULE__{} = binding) do
    query_builder =
      binding.query_builder |> QB.select("asLLM")

    %Dagger.LLM{
      query_builder: query_builder,
      client: binding.client
    }
  end

  @doc """
  Retrieve the binding value, as type Module
  """
  @spec as_module(t()) :: Dagger.Module.t()
  def as_module(%__MODULE__{} = binding) do
    query_builder =
      binding.query_builder |> QB.select("asModule")

    %Dagger.Module{
      query_builder: query_builder,
      client: binding.client
    }
  end

  @doc """
  Retrieve the binding value, as type ModuleConfigClient
  """
  @spec as_module_config_client(t()) :: Dagger.ModuleConfigClient.t()
  def as_module_config_client(%__MODULE__{} = binding) do
    query_builder =
      binding.query_builder |> QB.select("asModuleConfigClient")

    %Dagger.ModuleConfigClient{
      query_builder: query_builder,
      client: binding.client
    }
  end

  @doc """
  Retrieve the binding value, as type ModuleSource
  """
  @spec as_module_source(t()) :: Dagger.ModuleSource.t()
  def as_module_source(%__MODULE__{} = binding) do
    query_builder =
      binding.query_builder |> QB.select("asModuleSource")

    %Dagger.ModuleSource{
      query_builder: query_builder,
      client: binding.client
    }
  end

  @doc """
  Retrieve the binding value, as type Secret
  """
  @spec as_secret(t()) :: Dagger.Secret.t()
  def as_secret(%__MODULE__{} = binding) do
    query_builder =
      binding.query_builder |> QB.select("asSecret")

    %Dagger.Secret{
      query_builder: query_builder,
      client: binding.client
    }
  end

  @doc """
  Retrieve the binding value, as type Service
  """
  @spec as_service(t()) :: Dagger.Service.t()
  def as_service(%__MODULE__{} = binding) do
    query_builder =
      binding.query_builder |> QB.select("asService")

    %Dagger.Service{
      query_builder: query_builder,
      client: binding.client
    }
  end

  @doc """
  Retrieve the binding value, as type Socket
  """
  @spec as_socket(t()) :: Dagger.Socket.t()
  def as_socket(%__MODULE__{} = binding) do
    query_builder =
      binding.query_builder |> QB.select("asSocket")

    %Dagger.Socket{
      query_builder: query_builder,
      client: binding.client
    }
  end

  @doc """
  The binding's string value
  """
  @spec as_string(t()) :: {:ok, String.t() | nil} | {:error, term()}
  def as_string(%__MODULE__{} = binding) do
    query_builder =
      binding.query_builder |> QB.select("asString")

    Client.execute(binding.client, query_builder)
  end

  @doc """
  The digest of the binding value
  """
  @spec digest(t()) :: {:ok, String.t()} | {:error, term()}
  def digest(%__MODULE__{} = binding) do
    query_builder =
      binding.query_builder |> QB.select("digest")

    Client.execute(binding.client, query_builder)
  end

  @doc """
  A unique identifier for this Binding.
  """
  @spec id(t()) :: {:ok, Dagger.BindingID.t()} | {:error, term()}
  def id(%__MODULE__{} = binding) do
    query_builder =
      binding.query_builder |> QB.select("id")

    Client.execute(binding.client, query_builder)
  end

  @doc """
  Returns true if the binding is null
  """
  @spec null?(t()) :: {:ok, boolean()} | {:error, term()}
  def null?(%__MODULE__{} = binding) do
    query_builder =
      binding.query_builder |> QB.select("isNull")

    Client.execute(binding.client, query_builder)
  end

  @doc """
  The binding name
  """
  @spec name(t()) :: {:ok, String.t()} | {:error, term()}
  def name(%__MODULE__{} = binding) do
    query_builder =
      binding.query_builder |> QB.select("name")

    Client.execute(binding.client, query_builder)
  end

  @doc """
  The binding type
  """
  @spec type_name(t()) :: {:ok, String.t()} | {:error, term()}
  def type_name(%__MODULE__{} = binding) do
    query_builder =
      binding.query_builder |> QB.select("typeName")

    Client.execute(binding.client, query_builder)
  end
end

defimpl Jason.Encoder, for: Dagger.Binding do
  def encode(binding, opts) do
    {:ok, id} = Dagger.Binding.id(binding)
    Jason.Encode.string(id, opts)
  end
end

defimpl Nestru.Decoder, for: Dagger.Binding do
  def decode_fields_hint(_struct, _context, id) do
    {:ok, Dagger.Client.load_binding_from_id(Dagger.Global.dag(), id)}
  end
end
