package datasize

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestSize_UnmarshalText(t *testing.T) {
	testCases := []struct {
		in  string
		out ByteSize
		err *strconv.NumError
	}{
		{
			in:  "1B",
			out: B,
			err: nil,
		},
		{
			in:  "2kb",
			out: 2 * KB,
			err: nil,
		},
		{
			in:  "3Mega",
			out: 3 * MB,
			err: nil,
		},
		{
			in:  "4GigaByte",
			out: 4 * GB,
			err: nil,
		},
		{
			in:  "5TERABYTES",
			out: 5 * TB,
			err: nil,
		},
		{
			in:  "6p",
			out: 6 * PB,
			err: nil,
		},
		{
			in:  "7EB",
			out: 7 * EB,
			err: nil,
		},
		{
			in:  "8KIBI",
			out: 8 * KiB,
			err: nil,
		},
		{
			in:  "9Mebi",
			out: 9 * MiB,
			err: nil,
		},
		{
			in:  "10Gibibyte",
			out: 10 * GiB,
			err: nil,
		},
		{
			in:  "11TebiBytes",
			out: 11 * TiB,
			err: nil,
		},
		{
			in:  "12.13KB",
			out: 12130 * B,
			err: nil,
		},
		{
			in:  ".14MB",
			out: 140000 * B,
			err: nil,
		},
		{
			in:  "15.MiB",
			out: 15 * MiB,
			err: nil,
		},
		{
			in:  "16 MB",
			out: 16 * MB,
			err: nil,
		},
		{
			in:  "17 GiB",
			out: 17 * GiB,
			err: nil,
		},
		{
			in: "1",
			err: &strconv.NumError{
				Func: fnUnmarshalText,
				Num:  "1",
				Err:  strconv.ErrSyntax,
			},
		},
		{
			in: "1.2",
			err: &strconv.NumError{
				Func: fnUnmarshalText,
				Num:  "1.2",
				Err:  strconv.ErrSyntax,
			},
		},
		{
			in: "KB",
			err: &strconv.NumError{
				Func: fnUnmarshalText,
				Num:  "KB",
				Err:  strconv.ErrSyntax,
			},
		},
		{
			in: "1.2.3MB",
			err: &strconv.NumError{
				Func: fnUnmarshalText,
				Num:  "1.2.3MB",
				Err:  strconv.ErrSyntax,
			},
		},
		{
			in: "1Kb",
			err: &strconv.NumError{
				Func: fnUnmarshalText,
				Num:  "1Kb",
				Err:  strconv.ErrSyntax,
			},
		},
		{
			in: "1 Giga Bytes",
			err: &strconv.NumError{
				Func: fnUnmarshalText,
				Num:  "1 Giga Bytes",
				Err:  strconv.ErrSyntax,
			},
		},
		{
			in: "1.1KBS",
			err: &strconv.NumError{
				Func: fnUnmarshalText,
				Num:  "1.1KBS",
				Err:  strconv.ErrSyntax,
			},
		},
		{
			in: fmt.Sprintf("%dKB", uint64(math.MaxUint64)),
			err: &strconv.NumError{
				Func: fnUnmarshalText,
				Num:  fmt.Sprintf("%dKB", uint64(math.MaxUint64)),
				Err:  strconv.ErrRange,
			},
		},
		{
			in: "1.2B",
			err: &strconv.NumError{
				Func: fnUnmarshalText,
				Num:  "1.2B",
				Err:  strconv.ErrSyntax,
			},
		},
	}
	for _, tc := range testCases {
		t.Run("UnmarshalText "+tc.in, func(t *testing.T) {
			var s ByteSize

			err := s.UnmarshalText([]byte(tc.in))

			if tc.err == nil {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}

				if s != tc.out {
					t.Errorf("got %v, want %v", s, tc.out)
				}
			} else {
				var numError *strconv.NumError
				if errors.As(err, &numError) {
					if numError.Func != tc.err.Func {
						t.Errorf("got %q, want %q", numError.Func, tc.err.Func)
					}
					if numError.Num != tc.err.Num {
						t.Errorf("got %q, want %q", numError.Num, tc.err.Num)
					}
					if numError.Err != tc.err.Err {
						t.Errorf("got %q, want %q", numError.Err, tc.err.Err)
					}
				} else {
					t.Errorf("unexpected error: %v", err)
				}
			}
		})

		t.Run("Parse "+tc.in, func(t *testing.T) {
			s, err := Parse([]byte(tc.in))

			if tc.err == nil {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}

				if s != tc.out {
					t.Errorf("got %v, want %v", s, tc.out)
				}
			} else {
				var numError *strconv.NumError
				if errors.As(err, &numError) {
					if numError.Func != tc.err.Func {
						t.Errorf("got %q, want %q", numError.Func, tc.err.Func)
					}
					if numError.Num != tc.err.Num {
						t.Errorf("got %q, want %q", numError.Func, tc.err.Func)
					}
					if numError.Err != tc.err.Err {
						t.Errorf("got %q, want %q", numError.Func, tc.err.Func)
					}
				} else {
					t.Errorf("unexpected error: %v", err)
				}
			}
		})

		t.Run("MustParse "+tc.in, func(t *testing.T) {
			defer func() {
				r := recover()
				if tc.err != nil {
					if r == nil {
						t.Errorf("MustParse(%s) => no panic, want panic", tc.in)
					}

					err := r.(error)
					var numError *strconv.NumError
					if errors.As(err, &numError) {
						if numError.Func != tc.err.Func {
							t.Errorf("got %q, want %q", numError.Func, tc.err.Func)
						}
						if numError.Num != tc.err.Num {
							t.Errorf("got %q, want %q", numError.Func, tc.err.Func)
						}
						if numError.Err != tc.err.Err {
							t.Errorf("got %q, want %q", numError.Func, tc.err.Func)
						}
					} else {
						t.Errorf("unexpected error: %v", err)
					}
				} else if r != nil {
					t.Errorf("unexpected panic: %v", r)
				}
			}()

			s := MustParse([]byte(tc.in))

			if s != tc.out {
				t.Errorf("got %v, want %v", s, tc.out)
			}
		})

		t.Run("ParseString "+tc.in, func(t *testing.T) {
			s, err := ParseString(tc.in)

			if tc.err == nil {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}

				if s != tc.out {
					t.Errorf("got %v, want %v", s, tc.out)
				}
			} else {
				var numError *strconv.NumError
				if errors.As(err, &numError) {
					if numError.Func != tc.err.Func {
						t.Errorf("got %q, want %q", numError.Func, tc.err.Func)
					}
					if numError.Num != tc.err.Num {
						t.Errorf("got %q, want %q", numError.Func, tc.err.Func)
					}
					if numError.Err != tc.err.Err {
						t.Errorf("got %q, want %q", numError.Func, tc.err.Func)
					}
				} else {
					t.Errorf("unexpected error: %v", err)
				}
			}
		})

		t.Run("MustParseString "+tc.in, func(t *testing.T) {
			defer func() {
				r := recover()
				if tc.err != nil {
					if r == nil {
						t.Errorf("MustParse(%s) => no panic, want panic", tc.in)
					}

					err := r.(error)
					var numError *strconv.NumError
					if errors.As(err, &numError) {
						if numError.Func != tc.err.Func {
							t.Errorf("got %q, want %q", numError.Func, tc.err.Func)
						}
						if numError.Num != tc.err.Num {
							t.Errorf("got %q, want %q", numError.Func, tc.err.Func)
						}
						if numError.Err != tc.err.Err {
							t.Errorf("got %q, want %q", numError.Func, tc.err.Func)
						}
					} else {
						t.Errorf("unexpected error: %v", err)
					}
				} else if r != nil {
					t.Errorf("unexpected panic: %v", r)
				}
			}()

			s := MustParseString(tc.in)

			if s != tc.out {
				t.Errorf("got %v, want %v", s, tc.out)
			}
		})
	}
}

func TestByteSize_MarshalText(t *testing.T) {
	testCases := []struct {
		in  ByteSize
		out string
	}{
		{0, "0B"},
		{B, "1B"},
		{KB, "1KB"},
		{MB, "1MB"},
		{GB, "1GB"},
		{TB, "1TB"},
		{PB, "1PB"},
		{EB, "1EB"},
		{MB + 20*KB, "1020KB"},
		{KiB, "1.00KiB"},
		{MiB, "1.00MiB"},
		{GiB, "1.00GiB"},
		{TiB, "1.00TiB"},
		{PiB, "1.00PiB"},
		{EiB, "1.00EiB"},
		{123 * B, "123B"},
		{ByteSize(1.0625 * float64(KiB)), "1.06KiB"},
	}

	for _, tc := range testCases {
		t.Run("MarshalText "+tc.out, func(t *testing.T) {
			b, _ := tc.in.MarshalText()
			s := string(b)

			if s != tc.out {
				t.Errorf("got %q, want %q", s, tc.out)
			}
		})
	}
}
