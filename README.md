# Brute-force Password Breaker

A straightforward password breaker in Go.

## How to use

To crack a password, add it to a new line in `./passwords.txt`. In the same line, add the alphabet that the password breaker will use for the attack.

Example of `./passwords.txt`:

```
87355 0123456789
zqcyb abcdefghijklmnopqrstuvwxyz
GpaTJ aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ
RE5Pd aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ0123456789
W9#j! aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ0123456789!@#$%%^&*
```

Then just execute the following command:

```bash
go run .\main.go
```

The output should look as follows:

```
87355 -> cracked in 5.2848ms
zqcyb -> cracked in 498.6085ms
GpaTJ -> cracked in 4.500829s
RE5Pd -> cracked in 23.3394005s
W9#j! -> cracked in 54.5428082s
```
