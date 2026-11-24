package main

type Platform int

const (
	PlatformNone Platform = iota
	Windows
	Darwin
	Debian
	Ubuntu
)
