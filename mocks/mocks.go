package mocks

//go:generate mockgen -destination=./io.go -package=mocks -source=../interfaces/io.go Reader,Writer
