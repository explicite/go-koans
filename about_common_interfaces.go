package go_koans

import "bytes"
import "io"

func aboutCommonInterfaces() {
	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		io.Copy(out, in)

		assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
	}

	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		io.CopyN(out, in, 5)

		assert(out.String() == "hello") // duplicate only a portion of the io.Reader
	}
}
