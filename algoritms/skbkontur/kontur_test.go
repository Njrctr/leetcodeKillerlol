package skbkontur

import (
	"log"
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestCheckStringFunc(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	properties := gopter.NewProperties(parameters)

	randomCount := gen.Int()
	log.Println("randomCount: ", randomCount)
	// count := 15

	properties.Property("random data", prop.ForAll(
		func(randomStringSlice []string) bool {
			CheckStrings(randomStringSlice)
			return true
		},
		gen.ArrayOfN(5, gen.AnyString()),
	))

	properties.TestingRun(t)
}
