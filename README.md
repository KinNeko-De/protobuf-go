<picture>
  <img alt="Protobuf icon" src="/documentation/img/file-type-protobuf.svg" width="25%" height="25%">
</picture>
<picture>
  <img alt="Heart" src="https://github.githubassets.com/images/icons/emoji/unicode/2764.png">
</picture>
<picture>
  <img alt="Lua icon" src="/documentation/img/file-type-lua.svg" width="25%" height="25%">
</picture>
<picture>
  <img alt="Heart" src="https://github.githubassets.com/images/icons/emoji/unicode/2764.png">
</picture>
<picture>
  <img alt="Latex icon" src="/documentation/img/latex.svg" width="25%" height="25%">
</picture>

# Go related code for protocol buffers
This repository contains Go implementation for [protocol buffers](https://protobuf.dev).

The code design is based on [protocolbuffers/protobuf-go](https://github.com/protocolbuffers/protobuf-go). 

## Packages
*   [`encoding/protolua`]:  
    Package `protolua` converts protobuf messages to lua data types mostly using tables.  
    One way conversion is only supported for now.  
    Limited support of protobuf types because for it is only in usage in context of LuaTex.  

# Status
[![protobuf-go-ci](https://github.com/KinNeko-De/protobuf-go/actions/workflows/ci.yml/badge.svg)](https://github.com/KinNeko-De/protobuf-go/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/KinNeko-De/protobuf-go/branch/main/graph/badge.svg?token=yvQYJ6kpYr)](https://codecov.io/gh/KinNeko-De/protobuf-go)

