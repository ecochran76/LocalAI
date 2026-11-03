package config

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Disabled hardware defaults GPU discovery", func() {
	It("does not probe the local GPU while hardware defaults are disabled", func() {
		GinkgoT().Setenv("LOCALAI_DISABLE_HARDWARE_DEFAULTS", "true")
		GinkgoT().Setenv("LOCALAI_DISABLE_GUESSING", "true")
		previous := localGPU
		DeferCleanup(func() { localGPU = previous })
		calls := 0
		localGPU = func() GPU { calls++; return GPU{} }

		cfg := &ModelConfig{}
		cfg.SetDefaults()

		Expect(calls).To(BeZero())
	})
})
