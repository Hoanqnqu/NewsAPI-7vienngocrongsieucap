package main_test

import (
	"context"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/zhenghaoz/gorse/client"
	"time"
)

var _ = Describe("TestGorse", func() {
	It("test", func() {
		gorse := client.NewGorseClient("http://127.0.0.1:8087", "")

		_, err := gorse.InsertItem(context.Background(), client.Item{
			ItemId:     "hihi",
			IsHidden:   false,
			Labels:     []string{"test"},
			Categories: []string{"ngu"},
			Timestamp:  time.Now().String(),
			Comment:    "hay",
		})
		Expect(err).ToNot(HaveOccurred())

	})
})
