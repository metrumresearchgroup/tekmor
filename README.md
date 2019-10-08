# Tekmor

Tekmor was both a primitve Greek goddess and a word indicating a

> [sure sign, or token of some high and solemn kind"](https://en.wikipedia.org/wiki/Tekmor)

As such, Tekmor serves as a simple abstraction in golang for passing authentication to PAM and determining whether or not it was successful. 

## CGO and Requirements
Since we're using [msteinert's pam library](https://github.com/msteinert/pam), it has dependencies on GCC as well as the development libraries for PAM.

For Ubuntu systems, this would mean the following:
`apt-get install -y build-essential libpam0g-dev`

I have had little to no success on MAC as of yet, but these core dependencies will be necessary before a `go get github.com/msteinert/pam` will be successful.