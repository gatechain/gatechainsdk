# gatechainsdk
An SDK for the GateChain blockchain in Go.

## Getting started

### Prerequisites
Before using the SDK, you'll need to configure these settings in `const.go` under the `common` directory:

``` go
EndPoint := "your_server_ip:port"  // Modify to your own server's IP and port
APIToken := "your_api_token"      // Set to the token used to connect to GateChain
```

Installation & Testing

```bash
make sodium
make test
```

## Where can I see examples?

Take a look at `examples/` for some examples of how to write clients.

