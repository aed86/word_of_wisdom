package errors

import "fmt"

var InvalidChallenge = fmt.Errorf("challenge solution is not valid")
var OutdatedChallenge = fmt.Errorf("challenge is outdated")
