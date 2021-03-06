# BlockIo

This Golang library is the official Block.IO SDK. To use the functions provided by this SDK, you will need a REST client of your choice, and the Bitcoin, Litecoin , or Dogecoin API key(s) from <a href="https://block.io" target="_blank">Block.io</a>. Go ahead, sign up :)

## Installation

```bash
  go get github.com/BlockIo/block_io-go
```

## Usage

It's easy to get started. In your code, do this:

```go
  import blockio "github.com/BlockIo/block_io-go"

  var withdrawResponse string // store json string response to /api/v2/withdraw here

  withdrawData, _ := blockio.ParseResponseData(withdrawResponse)
  signatureReq, _ := blockio.SignWithdrawRequest("YOUR_PIN", withdrawData)

  // post signatureReq to /api/v2/sign_and_finalize_withdrawal
```

##### For a more detailed guide on usage, check the examples folder in the repo

## Testing

```bash
  go test -v
```
