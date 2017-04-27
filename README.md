# logo

[![Build Status](https://travis-ci.org/fulldump/logo.svg?branch=master)](https://travis-ci.org/fulldump/logo)
[![Go report card](http://goreportcard.com/badge/fulldump/logo)](https://goreportcard.com/report/fulldump/logo)
[![GoDoc](https://godoc.org/github.com/fulldump/logo?status.svg)](https://godoc.org/github.com/fulldump/logo)

<sup>Tested for Go 1.5, 1.6, 1.7</sup>

Loging for go

<!-- MarkdownTOC autolink=true bracket=round depth=4 -->

- [Getting started](#getting-started)
- [Available levels for logging](#available-levels-for-logging)
- [Custom logger](#custom-logger)
- [Default logger](#default-logger)
- [Configure logo](#configure-logo)
	- [Level](#level)
	- [Component](#component)
	- [Output](#output)
	- [Showline](#showline)
	- [Print](#print)

<!-- /MarkdownTOC -->


## Getting started

The easiest way of logging is using the default logger:

```go
logo.Info("My name is ", name)
logo.Debug("DB status: ", dbStatus)
```

Will generate this output:

```txt
LOG:{"time":"2017-04-27T14:17:53.469122844Z","type":"LOG","level":"INFO","payload":"My name is Fulanez"}
LOG:{"time":"2017-04-27T14:17:53.469122844Z","type":"LOG","level":"DEBUG","payload":"DB status: ON (src/myproject/main.go:16)"}
```

Notice all levels other than `INFO` will show file and line number.


## Available levels for logging

* `Fatal`
* `Error`
* `Warning`
* `Info`
* `Debug`
* `Trace`


## Custom logger

You can create a new logger for your application:

```go
logger := logo.NewLogo()
logger.Log("hello!")
```


## Default logger

Default logger is accesible from:

```go
logo.Default
```

## Configure logo

Any logger (even the default one) can be configured

### Level

Configure level priority. All possible options are:

* `logo.LevelOff` - Disable all logs
* `logo.LevelFatal`
* `logo.LevelError`
* `logo.LevelWarning`
* `logo.LevelInfo`
* `logo.LevelDebug`
* `logo.LevelTrace`
* `logo.LevelAll` - Show all logs

### Component

Define a string to be added to all generated logs.

### Output

Output is any `Writer` interface, you can change it to the
stderr, a string buffer, etc.


### Showline

Disable file and line numbers in logs.


### Print

Is the function that knows the format to print the log.

```go

mylogger := logo.NewLogo()

mylogger.Print = func (l *Logo, lvl Level, a ...interface{}) {
	fmt.Fprint("my custom log ", a...)
}

```
