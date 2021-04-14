module example.com/hello

go 1.15

replace example.com/greetings => ../greetings
replace github.com/hypebeast/go-osc => ../go-osc

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	github.com/hypebeast/go-osc v0.0.0-20210408213458-3287e1838f40
)
