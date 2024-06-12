package main_test

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/zhenghaoz/gorse/client"
)

var _ = Describe("TestGorse", func() {
	It("test", func() {
		gorse := client.NewGorseClient("http://127.0.0.1:8087", "")

		scores, err := gorse.GetItemRecommendWithCategory(context.Background(), "", "", "read", "300s", 10, 0)
		Expect(err).ToNot(HaveOccurred())
		Expect(len(scores) > 0).To(BeTrue())

	})
})
