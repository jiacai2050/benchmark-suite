// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: metric.go

package jsonbenmark

import (
	"bytes"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *DemoFF) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *DemoFF) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ `)
	if j.A != 0 {
		buf.WriteString(`"a":`)
		fflib.FormatBits2(buf, uint64(j.A), 10, j.A < 0)
		buf.WriteByte(',')
	}
	if len(j.B) != 0 {
		buf.WriteString(`"b":`)
		fflib.WriteJsonString(buf, string(j.B))
		buf.WriteByte(',')
	}
	if true {
		buf.WriteString(`"c":`)

		{

			obj, err = j.C.MarshalJSON()
			if err != nil {
				return err
			}
			buf.Write(obj)

		}
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtDemoFFbase = iota
	ffjtDemoFFnosuchkey

	ffjtDemoFFA

	ffjtDemoFFB

	ffjtDemoFFC
)

var ffjKeyDemoFFA = []byte("a")

var ffjKeyDemoFFB = []byte("b")

var ffjKeyDemoFFC = []byte("c")

// UnmarshalJSON umarshall json - template of ffjson
func (j *DemoFF) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *DemoFF) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtDemoFFbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtDemoFFnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeyDemoFFA, kn) {
						currentKey = ffjtDemoFFA
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'b':

					if bytes.Equal(ffjKeyDemoFFB, kn) {
						currentKey = ffjtDemoFFB
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'c':

					if bytes.Equal(ffjKeyDemoFFC, kn) {
						currentKey = ffjtDemoFFC
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyDemoFFC, kn) {
					currentKey = ffjtDemoFFC
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyDemoFFB, kn) {
					currentKey = ffjtDemoFFB
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyDemoFFA, kn) {
					currentKey = ffjtDemoFFA
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtDemoFFnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtDemoFFA:
					goto handle_A

				case ffjtDemoFFB:
					goto handle_B

				case ffjtDemoFFC:
					goto handle_C

				case ffjtDemoFFnosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_A:

	/* handler: j.A type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.A = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_B:

	/* handler: j.B type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.B = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_C:

	/* handler: j.C type=time.Time kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.C.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}
