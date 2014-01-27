# String similarity algorithms

This repository provides implementations for the string similarity algorithms Jaro and Jaro-Winkler. For details of the algorithms please have a look at Wikipedia or somewhere else.

// Jaro-Winkler
* via Go
For Jaro-Winkler it is optionally possible to set custom boost threshold and prefix scale. The defaults are 0.7 and 0.1 as suggested by Jaro and Winkler.

```go
j := stringSimilarity.Jaro("string1", "string2")

stringSimilarity.SetBoostThreshold(0.7)
stringSimilarity.SetPrefixScale(0.1)
jw := stringSimilarity.JaroWinkler("string1", "string2")
```

* via command line
install commandLineInterface.go from within stringSimilarity/strsim
> go install

and run the compiled output $GOPATH/bin/strsim
> $GOPATH/bin/strsim metric string1 string2

where metric is either "jaro" or "jaroWinkler"
and string1 and string2 are the strings that should be compared
