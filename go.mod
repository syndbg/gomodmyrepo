module github.com/syndbg/gomodmyrepo

go 1.12

replace github.com/syndbg/gomodmyrepo => ./mig

require (
	github.com/syndbg/gomodanotherrepo v1.0.0
	github.com/syndbg/gomodmyrepo/mig v1.0.0
)
