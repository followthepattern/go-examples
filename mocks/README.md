## How the mock works in Golang
In this package you can find practical example of how to mock interfaces, which later you can use for writing unit tests.

`interface` package has interfaces I copied from different libraries as examples you can give it to mockgen to generate mocks.

[mocks.go](./mocks.go) file contains commands to specify the interfaces for mock generation as comments. It can be executed in terminal as well `mockgen -destination=./io.go -package=mocks -source=../interfaces/io.go Reader,Writer`.