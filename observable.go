package main

type Observable interface {
	sign(observer Observer)
	unsign(observer Observer)
	notifyAll()
}
