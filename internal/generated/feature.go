package generated

import (
	inject "go.octolab.org/toolkit/config"

	"go.octolab.org/template/service/internal/config"
)

func init() {
	config.Defaults.Build.Features = append(
		config.Defaults.Build.Features,
		inject.Feature{
			// 15c05532-6136-4e25-aed5-0c73f975ea40
			ID: [16]byte{
				0x15, 0xc0, 0x55, 0x32,
				0x61, 0x36, 0x4e, 0x25,
				0xae, 0xd5, 0xc, 0x73,
				0xf9, 0x75, 0xea, 0x40,
			},
			Name:    "template",
			Brief:   "One possible example how to use features.",
			Docs:    "https://go-service.octolab.org/",
			RFC:     "https://github.com/octomation/go-service/discussions",
			Enabled: true,
		},
	)
}
