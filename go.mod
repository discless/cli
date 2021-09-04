module github.com/discless/discless-cli

go 1.15

require (
	github.com/bwmarrin/discordgo v0.23.2
	github.com/containerd/containerd v1.5.5 // indirect
	github.com/discless/discless v0.0.0-20210818213551-a6fe55bdea77
	github.com/docker/docker v20.10.8+incompatible
	github.com/docker/go-connections v0.4.0
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/spf13/cobra v1.2.1
	golang.org/x/term v0.0.0-20210615171337-6886f2dfbf5b
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/discless/discless => /home/tris/Coding/Go/src/github.com/discless/discless
