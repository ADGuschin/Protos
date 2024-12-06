# Dressy protoBuf library

This is the protocol buffer files for gRPC interface between all system services

## Installation

Clone the repository.

In addition to the *.proto files, there are two scripts for adding/updating system environment variables:

### SET_ENV_PROTO_PATH.ps1 (for Win)

#### ! Administrator rights required !

If you got SecurityError, try this:
```bash
Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass
SET_ENV_PROTO_PATH.ps1
```

### SET_ENV_PROTO_PATH.sh (for macOS/Linux not tested yet)

Using:
```bash
chmod +x SET_ENV_PROTO_PATH.sh
sudo ./SET_ENV_PROTO_PATH.sh
```

## IDE Setup
### Visual Studio 2022
In your *.csproj file add this:
```xml
	<ItemGroup>
		<Protobuf Include="$(DRESSY_LOCAL_PROTO_PATH)/**/*.proto" GrpcServices="Both">
			<Link>Protos/%(RecursiveDir)%(FileName)%(Extension)</Link>
		</Protobuf>
		<Protobuf Include="$(DRESSY_LOCAL_PROTO_PATH)/**/common.proto" GrpcServices="None">
			<Link>Protos/%(RecursiveDir)%(FileName)%(Extension)</Link>
		</Protobuf>
	</ItemGroup>
```

### Other IDEs

(Help with any info for it, please)
