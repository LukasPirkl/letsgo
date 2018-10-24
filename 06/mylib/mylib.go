package mylib

type item struct {
	one string
	Two string
}

type Data struct {
	one string
	Two string
}

type Embedded struct {
	Data
	Three string
}

type Amount int

type Filter func([]item) []item
