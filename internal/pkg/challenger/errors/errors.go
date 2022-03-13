package errors

import "fmt"

var InvalidSolve = fmt.Errorf("challenge solution is not valid")
var OutdatedChallenge = fmt.Errorf("challenge is outdated")
