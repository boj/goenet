# Go ENet

A Go wrapper for the C based ENet UDP network library.

## Install

    go get github.com/boj/goenet

## Documentation

Located at [GoDocs](http://godoc.org/github.com/boj/goenet).

## Existing Caveats

Currently when using _NewPacket_ to build a packet, the data parameter is required to be of type __[]byte__.  Data is also returned as such.

The _peer->data_ interface is strictly a stub until I can figure out how to easily put arbitrary data in there from Go.  This is somewhat related to the above issue.

## TODO

* Verify all ENet related methods work as intended.
* Optimizations.
* Benchmarks.
* Testing where applicable.

## Acknowledgements

[ENet](http://enet.bespin.org/) - MIT license.

## Author

Brian 'bojo' Jones mojobojo@gmail.com

## License

http://boj.mit-license.org/

