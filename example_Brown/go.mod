module github.com/example_Brown

go 1.16

replace github.com/evaluation => ../evaluation

replace github.com/smoothingMethod => ../smoothingMethod

require (
	github.com/evaluation v0.0.0-00010101000000-000000000000
	github.com/go-gota/gota v0.11.0
	github.com/smoothingMethod v0.0.0-00010101000000-000000000000
)
