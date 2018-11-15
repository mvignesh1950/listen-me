'listen-me' the builder
===

Build go projects automatically when files are modified.

Getting
===

```
go get github.com/mvignesh1950/listen-me
```

Usage
===

With `GOPATH/bin` in `PATH`,

```
Usage: listen-me [options]
options:
	-p, -path       Directory               The directory to watch.
	-n, -name       Name                    The name for binary file.
	-e, -env        Environment file path   Path to file containing environment variables to be set for the service.
	-b, -buildonly  Build only mode         Just do a build when a change is detected.
	-c, -commands   Custom commands to run  Run custom command after build.
	-v, -version    Version                 Prints the version.
	-h, -help       Help                    Show this help.

```
