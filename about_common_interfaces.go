package go_koans

import (
	"bytes"
	"io"
)

func aboutCommonInterfaces() {
	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		_, error := io.Copy(out, in)
		assert(error == nil)

		assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
	}

	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)
		_, error := io.CopyN(out, in, 5)
		assert(error == nil)

		assert(out.String() == "hello") // duplicate only a portion of the io.Reader
	}
}
