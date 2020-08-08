// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package jsonbenmark

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson9478868cDecodeJsonbenmark(in *jlexer.Lexer, out *DemoEasy) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "a":
			out.A = int(in.Int())
		case "b":
			out.B = string(in.String())
		case "c":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.C).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson9478868cEncodeJsonbenmark(out *jwriter.Writer, in DemoEasy) {
	out.RawByte('{')
	first := true
	_ = first
	if in.A != 0 {
		const prefix string = ",\"a\":"
		first = false
		out.RawString(prefix[1:])
		out.Int(int(in.A))
	}
	if in.B != "" {
		const prefix string = ",\"b\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.B))
	}
	if true {
		const prefix string = ",\"c\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.C).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DemoEasy) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9478868cEncodeJsonbenmark(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DemoEasy) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9478868cEncodeJsonbenmark(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DemoEasy) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9478868cDecodeJsonbenmark(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DemoEasy) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9478868cDecodeJsonbenmark(l, v)
}