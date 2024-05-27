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

		_, err := gorse.InsertItem(context.Background(), client.Item{
			ItemId:     "testnha",
			IsHidden:   false,
			Labels:     []string{"test"},
			Categories: []string{"ngu"},
			Timestamp:  "2022/06/05 07:06",
			Comment:    "hay",
		})
		Expect(err).ToNot(HaveOccurred())

	})
})
