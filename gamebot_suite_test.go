package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGamebot(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gamebot Suite")
}
