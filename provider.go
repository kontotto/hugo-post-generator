package main

type Provider interface {
	Data() (interface{}, error)
}
