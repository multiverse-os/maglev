module github.com/multiverse-os/maglev-app

go 1.19

require (
	github.com/multiverse-os/cli v0.1.0
	github.com/multiverse-os/maglev v0.1.0
)

replace (
	github.com/multiverse-os/cli/data => github.com/multiverse-os/data v0.1.0
	github.com/multiverse-os/cli/terminal/ansi => github.com/multiverse-os/ansi v0.1.0
	github.com/multiverse-os/cli/terminal/loading => github.com/multiverse-os/loading v0.1.0
	github.com/multiverse-os/cli/terminal/text/banner => github.com/multiverse-os/banner v0.1.0
)

require (
	/github.com/multiverse-os/cli/terminal/ansi v0.1.0
	github.com/multiverse-os/cli/data v0.1.0
	github.com/multiverse-os/cli/terminal/loading v0.1.0
	github.com/multiverse-os/cli/terminal/text/banner v0.1.0
)

replace github.com/multiverse-os/maglev-app v0.1.0 => ../app

replace github.com/multiverse-os/maglev v0.1.0 => ../

//require github.com/multiverse-os/banner v0.1.0 // indirect

//github.com/multiverse-os/banner/fonts/big v0.0.0-20231217224509-6c8731036547 // indirect
//github.com/multiverse-os/banner/fonts/giant v0.0.0-20231217224509-6c8731036547 // indirect
//require golang.org/x/text v0.14.0 // indirect
