package thoth_test

import (
	"testing"

	"github.com/novln/macchiato"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoThoth(t *testing.T) {
	RegisterFailHandler(Fail)
	macchiato.RunSpecs(t, "GoThoth Suite")
}
