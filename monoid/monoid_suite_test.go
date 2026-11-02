package monoid_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMonoid(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Monoid Suite")
}
