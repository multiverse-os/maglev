module github.com/multiverse-os/maglev

go 1.19

require (
	github.com/multiverse-os/cli v0.1.0
	github.com/multiverse-os/webframe v0.0.0-00010101000000-000000000000
)

require (
	github.com/multiverse-os/ansi v0.0.0-20230122075550-10efed87b9d4 // indirect
	github.com/multiverse-os/banner v0.0.0-20231006133835-80f8c892b073 // indirect
)

exclude (
	github.com/multiverse-os/ansi v0.1.0
	github.com/multiverse-os/banner v0.1.0
)

require (
	github.com/multiverse-os/cli/data v0.1.0 // indirect
	github.com/multiverse-os/cli/terminal/ansi v0.1.0 // indirect
	github.com/multiverse-os/cli/terminal/loading v0.1.0 // indirect
	github.com/multiverse-os/cli/terminal/text/banner v0.1.0 // indirect
)

replace (
	github.com/multiverse-os/cli/data => github.com/multiverse-os/data v0.1.0
	github.com/multiverse-os/cli/terminal/ansi => github.com/multiverse-os/ansi v0.1.0
	github.com/multiverse-os/cli/terminal/loading => github.com/multiverse-os/loading v0.1.0
	github.com/multiverse-os/cli/terminal/text/banner => github.com/multiverse-os/banner v0.1.0
)

replace github.com/multiverse-os/webframe => ../../webframe

require github.com/multiverse-os/muid v0.1.0 // indirect

require (
	git.mills.io/prologic/bitcask v1.0.2 // indirect
	github.com/abcum/lcp v0.0.0-20201209214815-7a3f3840be81 // indirect
	github.com/akrylysov/pogreb v0.10.2 // indirect
	github.com/gofrs/flock v0.8.0 // indirect
	github.com/multiverse-os/service v0.1.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/plar/go-adaptive-radix-tree v1.0.4 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/exp v0.0.0-20200228211341-fcea875c7e85 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
