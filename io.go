package oprs

import (
	"fmt"
	"io"
	"os"
)

var DefaultPrinter = os.Stdout

// Fprinter returns a closure that prints to the given writer
// defaults to DefaultPrinter if writer is nil
func Fprinter(w io.Writer) func(...any) (int, error) {
	w = Ternary[io.Writer](w == nil, os.Stdout, w)
	return func(a ...any) (int, error) {
		return fmt.Fprint(w, a...)
	}
}

// Fprinterln returns a closure that prints a line to the given writer
// defaults to DefaultPrinter if writer is nil
func Fprinterln(w io.Writer) func(...any) (int, error) {
	w = Ternary[io.Writer](w == nil, os.Stdout, w)
	return func(a ...any) (int, error) {
		return fmt.Fprintln(w, a...)
	}
}

// Fprinterf returns a closure that prints a format string to the given writer
// defaults to DefaultPrinter if writer is nil
func Fprinterf(w io.Writer, format string) func(...any) (int, error) {
	w = Ternary[io.Writer](w == nil, os.Stdout, w)
	return func(a ...any) (int, error) {
		return fmt.Fprintf(w, format, a...)
	}
}

// Printer returns a closure that prints to the given writer
func Printer() func(...any) (int, error) {
	return Fprinter(nil)
}

// Printerln returns a closure that prints a line to the given writer
func Printerln() func(...any) (int, error) {
	return Fprinterln(nil)
}

// Printerf returns a closure that prints a format string to the given writer
func Printerf(format string) func(...any) (int, error) {
	return Fprinterf(nil, format)
}
