package maybe_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMaybe(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Maybe Suite")
}
