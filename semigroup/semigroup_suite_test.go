package semigroup_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSemigroup(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Semigroup Suite")
}
