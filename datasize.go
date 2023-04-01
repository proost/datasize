package datasize

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type ByteSize uint64

const (
	B ByteSize = 1

	KB = 1000 * B
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB
	EB = 1000 * PB

	KiB = B << 10
	MiB = KiB << 10
	GiB = MiB << 10
	TiB = GiB << 10
	PiB = TiB << 10
	EiB = PiB << 10

	fnUnmarshalText string = "UnmarshalText"
)

var byteSizeRE = regexp.MustCompile(`([0-9]*)(\.[0-9]*)?([A-Za-z ]+)`)

// Bytes returns the number of bytes in the ByteSize.
func (b ByteSize) Bytes() uint64 {
	return uint64(b)
}

// Kilobytes returns the number of kilobytes in the ByteSize.
func (b ByteSize) Kilobytes() float64 {
	v := b / KB
	r := b % KB
	return float64(v) + float64(r)/float64(KB)
}

// Megabytes returns the number of megabytes in the ByteSize.
func (b ByteSize) Megabytes() float64 {
	v := b / MB
	r := b % MB
	return float64(v) + float64(r)/float64(MB)
}

// Gigabytes returns the number of gigabytes in the ByteSize.
func (b ByteSize) Gigabytes() float64 {
	v := b / GB
	r := b % GB
	return float64(v) + float64(r)/float64(GB)
}

// Terabytes returns the number of terabytes in the ByteSize.
func (b ByteSize) Terabytes() float64 {
	v := b / TB
	r := b % TB
	return float64(v) + float64(r)/float64(TB)
}

// Petabytes returns the number of petabytes in the ByteSize.
func (b ByteSize) Petabytes() float64 {
	v := b / PB
	r := b % PB
	return float64(v) + float64(r)/float64(PB)
}

// Exabytes returns the number of exabytes in the ByteSize.
func (b ByteSize) Exabytes() float64 {
	v := b / EB
	r := b % EB
	return float64(v) + float64(r)/float64(EB)
}

// Kibibytes returns the number of kibibytes in the ByteSize.
func (b ByteSize) Kibibytes() float64 {
	v := b / KiB
	r := b % KiB
	return float64(v) + float64(r)/float64(KiB)
}

// Mebibytes returns the number of mebibytes in the ByteSize.
func (b ByteSize) Mebibytes() float64 {
	v := b / MiB
	r := b % MiB
	return float64(v) + float64(r)/float64(MiB)
}

// Gibibytes returns the number of gibibytes in the ByteSize.
func (b ByteSize) Gibibytes() float64 {
	v := b / GiB
	r := b % GiB
	return float64(v) + float64(r)/float64(GiB)
}

// Tebibytes returns the number of tebibytes in the ByteSize.
func (b ByteSize) Tebibytes() float64 {
	v := b / TiB
	r := b % TiB
	return float64(v) + float64(r)/float64(TiB)
}

// Pebibytes returns the number of pebibytes in the ByteSize.
func (b ByteSize) Pebibytes() float64 {
	v := b / PiB
	r := b % PiB
	return float64(v) + float64(r)/float64(PiB)
}

// Exbibytes returns the number of exbibytes in the ByteSize.
func (b ByteSize) Exbibytes() float64 {
	v := b / EiB
	r := b % EiB
	return float64(v) + float64(r)/float64(EiB)
}

// String returns a string representation of the ByteSize.
func (b ByteSize) String() string {
	switch {
	case b == 0:
		return "0B"
	case b%EB == 0:
		return fmt.Sprintf("%dEB", b/EB)
	case b >= EiB:
		return fmt.Sprintf("%.2fEiB", b.Exbibytes())
	case b%PB == 0:
		return fmt.Sprintf("%dPB", b/PB)
	case b >= PiB:
		return fmt.Sprintf("%.2fPiB", b.Pebibytes())
	case b%TB == 0:
		return fmt.Sprintf("%dTB", b/TB)
	case b >= TiB:
		return fmt.Sprintf("%.2fTiB", b.Tebibytes())
	case b%GB == 0:
		return fmt.Sprintf("%dGB", b/GB)
	case b >= GiB:
		return fmt.Sprintf("%.2fGiB", b.Gibibytes())
	case b%MB == 0:
		return fmt.Sprintf("%dMB", b/MB)
	case b >= MiB:
		return fmt.Sprintf("%.2fMiB", b.Mebibytes())
	case b%KB == 0:
		return fmt.Sprintf("%dKB", b/KB)
	case b >= KiB:
		return fmt.Sprintf("%.2fKiB", b.Kibibytes())
	default:
		return fmt.Sprintf("%dB", b)
	}
}

func (b ByteSize) MarshalText() ([]byte, error) {
	return []byte(b.String()), nil
}

func (b *ByteSize) UnmarshalText(t []byte) error {
	s := string(t)
	ss := byteSizeRE.FindStringSubmatch(s)
	if len(ss) == 0 || ss[0] != s {
		return &strconv.NumError{Func: fnUnmarshalText, Num: s, Err: strconv.ErrSyntax}
	}

	num, err := strconv.ParseFloat(ss[1]+ss[2], 64)
	if err != nil {
		var numErr *strconv.NumError
		if errors.As(err, &numErr) {
			return &strconv.NumError{Func: fnUnmarshalText, Num: s, Err: numErr.Err}
		}

		return &strconv.NumError{Func: fnUnmarshalText, Num: s, Err: err}
	}

	unit := strings.TrimSpace(ss[3])
	switch unit {
	case "Kb", "Mb", "Gb", "Tb", "Pb", "Eb", "Kib", "Mib", "Gib", "Tib", "Pib", "Eib":
		return &strconv.NumError{Func: fnUnmarshalText, Num: s, Err: strconv.ErrSyntax}
	}

	unit = strings.ToLower(unit)
	switch unit {
	case "", "b", "byte":
	case "k", "kb", "kilo", "kilobyte", "kilobytes":
		if num > math.MaxUint64/float64(KB) {
			return &strconv.NumError{Func: fnUnmarshalText, Num: s, Err: strconv.ErrRange}
		}
		num *= float64(KB)

	case "m", "mb", "mega", "megabyte", "megabytes":
		if num > math.MaxUint64/float64(MB) {
			return &strconv.NumError{Func: fnUnmarshalText, Num: s, Err: strconv.ErrRange}
		}
		num *= float64(MB)

	case "g", "gb", "giga", "gigabyte", "gigabytes":
		if num > math.MaxUint64/float64(GB) {
			return &strconv.NumError{Func: fnUnmarshalText, Num: s, Err: strconv.ErrRange}
		}
		num *= float64(GB)

	case "t", "tb", "tera", "terabyte", "terabytes":
		if num > math.MaxUint64/float64(TB) {
			return &strconv.NumError{Func: fnUnmarshalText, Num: s, Err: strconv.ErrRange}
		}
		num *= float64(TB)

	case "p", "pb", "peta", "petabyte", "petabytes":
		if num > math.MaxUint64/float64(PB) {
			return &strconv.NumError{Func: fnUnmarshalText, Num: s, Err: strconv.ErrRange}
		}
		num *= float64(PB)

	case "e", "eb", "exa", "exabyte", "exabytes":
		if num > math.MaxUint64/float64(EB) {
			return &strconv.NumError{Func: fnUnmarshalText, Num: s, Err: strconv.ErrRange}
		}
		num *= float64(EB)

	case "ki", "kib", "kibi", "kibibyte", "kibibytes":
		if num > math.MaxUint64/float64(KiB) {
			return &strconv.NumError{Func: fnUnmarshalText, Num: s, Err: strconv.ErrRange}
		}
		num *= float64(KiB)

	case "mi", "mib", "mebi", "mebibyte", "mebibytes":
		if num > math.MaxUint64/float64(MiB) {
			return &strconv.NumError{Func: fnUnmarshalText, Num: s, Err: strconv.ErrRange}
		}
		num *= float64(MiB)

	case "gi", "gib", "gibi", "gibibyte", "gibibytes":
		if num > math.MaxUint64/float64(GiB) {
			return &strconv.NumError{Func: fnUnmarshalText, Num: s, Err: strconv.ErrRange}
		}
		num *= float64(GiB)

	case "ti", "tib", "tebi", "tebibyte", "tebibytes":
		if num > math.MaxUint64/float64(TiB) {
			return &strconv.NumError{Func: fnUnmarshalText, Num: s, Err: strconv.ErrRange}
		}
		num *= float64(TiB)

	case "pi", "pib", "pebi", "pebibyte", "pebibytes":
		if num > math.MaxUint64/float64(PiB) {
			return &strconv.NumError{Func: fnUnmarshalText, Num: s, Err: strconv.ErrRange}
		}
		num *= float64(PiB)

	case "ei", "eib", "exbi", "exbibyte", "exbibytes":
		if num > math.MaxUint64/float64(EiB) {
			return &strconv.NumError{Func: fnUnmarshalText, Num: s, Err: strconv.ErrRange}
		}
		num *= float64(EiB)

	default:
		return &strconv.NumError{Func: fnUnmarshalText, Num: s, Err: strconv.ErrSyntax}
	}

	v := ByteSize(num)
	if float64(v) != num {
		return &strconv.NumError{Func: fnUnmarshalText, Num: s, Err: strconv.ErrSyntax}
	}

	*b = v
	return nil
}

// Parse parses a byte slice and returns the byte size it represents.
func Parse(t []byte) (ByteSize, error) {
	var v ByteSize
	err := v.UnmarshalText(t)
	return v, err
}

// MustParse is like Parse but panics if the byte slice is invalid.
func MustParse(t []byte) ByteSize {
	v, err := Parse(t)
	if err != nil {
		panic(err)
	}
	return v
}

// ParseString parses a string and returns the byte size it represents.
func ParseString(s string) (ByteSize, error) {
	return Parse([]byte(s))
}

// MustParseString is like ParseString but panics if the string is invalid.
func MustParseString(s string) ByteSize {
	return MustParse([]byte(s))
}
