# Legacy

Collection of legacy decoder plugins for backward compatibility.

Decoder now doesn't have hard dependency of *vector packages, so need to provide support of old callbacks/modifiers for
parsing various formats.

List of supported callbacks:
* 

### Usage

Just add package to import section like that
```go
import (
	_ "github.com/koykov/decoder_legacy"
)
```
