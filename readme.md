# urbit-gob
### A go port of [urbit-ob](https://github.com/urbit/urbit-ob)

---

#### Command line use
```
> go run cmd/main.go patp 0
~zod
> go run cmd/main.go clan ~marzod
star
> go 
Usage: ...main COMMAND args...

Valid commands:

    patp                : converts a number to a @p-encoded string

    patp2dec            : converts a @p-encoded string to a decimal-encoded string

    patp2hex            : converts a @p-encoded string to a hex-encoded string

    patq                : converts a number to a @q-encoded string

    patq2dec            : converts a @q-encoded string to a decimal-encoded string

    patq2hex            : converts a @q-encoded string to a hex-encoded string

    hex2patp            : converts a hex-encoded string to a @p-encoded string

    hex2patq            : converts a hex-encoded string to a @q-encoded string

    clan                : determines the ship class of a @p value

    sein                : determines the parent of a @p value

    eqpatq              : performs an equality comparison on @q values

    isvalidpat          : weakly checks if a string is a valid @p or @q value

    isvalidpatp         : validates a @p string

    isvalidpatq         : validates a @q string
```

#### Module use
```go
package main

import "github.com/deelawn/urbit-gob/co"

func main() {

	// name = ~fipfes
	name, err := co.Patp("65535")
	if err != nil {
		panic("patp")
	}

	// sponsor = ~fes
	sponsor, err := co.Sein(name)
	if err != nil {
		panic("sein")
	}
}
```