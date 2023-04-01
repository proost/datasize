# datasize

Go package for parsing human readable data sizes.

## Inspiration

This package is inspired by [https://github.com/c2h5oh/datasize]()

There are two differences
    * one is that this package divides binary units and decimal units into separate constants.
    * the other is that this package allows parsing of string which value is floating point number.

So, this package is almost same except for the above two points.

## Constants

Just like `time` package provides `time.Second`, `time.Day` constants `datasize` provides:

* `datasize.B` 1 byte
* `datasize.KB` 1 kilobyte
* `datasize.MB` 1 megabyte
* `datasize.GB` 1 gigabyte
* `datasize.TB` 1 terabyte
* `datasize.PB` 1 petabyte
* `datasize.EB` 1 exabyte
* `datasize.KiB` 1 kibibyte
* `datasize.MiB` 1 mebibyte
* `datasize.GiB` 1 gibibyte
* `datasize.TiB` 1 tebibyte
* `datasize.PiB` 1 pebibyte
* `datasize.EiB` 1 exbibyte

## Helpers

Just like `time` package provides `duration.Nanoseconds() uint64 `, `duration.Hours() float64` helpers `datasize` has.

* `ByteSize.Bytes() uint64`
* `ByteSize.Kilobytes() float64`
* `ByteSize.Megabytes() float64`
* `ByteSize.Gigabytes() float64`
* `ByteSize.Terabytes() float64`
* `ByteSize.Petabytes() float64`
* `ByteSize.Exabytes() float64`
* `ByteSize.Kibibytes() float64`
* `ByteSize.Mebibytes() float64`
* `ByteSize.Gibibytes() float64`
* `ByteSize.Tebibytes() float64`
* `ByteSize.Pebibytes() float64`
* `ByteSize.Exbibytes() float64`
