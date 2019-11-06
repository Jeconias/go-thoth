package models_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/lab259/rlog/v2"
	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/gomega"

	"github.com/novln/macchiato"
)

func TestValidatorService(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	rlog.SetOutput(ginkgo.GinkgoWriter)

	dir, _ := os.Getwd()
	ginkgo.GinkgoWriter.Write([]byte(fmt.Sprintf("CWD: %s\n", dir)))
	ginkgo.GinkgoWriter.Write([]byte(fmt.Sprintf("ENV: %s\n", os.Getenv("ENV"))))
	ginkgo.GinkgoWriter.Write([]byte(fmt.Sprintf("RandomSeed: %d\n", config.GinkgoConfig.RandomSeed)))

	gomega.RegisterFailHandler(ginkgo.Fail)

	macchiato.RunSpecs(t, "GoThoth Test Suite")
}
