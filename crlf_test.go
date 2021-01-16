package crlf

import (
	"bytes"
	"io"
	"log"
	"reflect"
	"testing"

	"golang.org/x/text/transform"
)

func TestWindows1252CRLF(t *testing.T) {
	type args struct {
		src []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Convert newlines",
			args: args{
				[]byte("\n\n"),
			},
			want: []byte{13, 10, 13, 10},
		},
		{
			name: "Handle CRLF in current",
			args: args{
				[]byte("\r\n"),
			},
			want: []byte{13, 10},
		},
		{
			name: "Registered Symbol Convert",
			args: args{
				[]byte("®\n"),
			},
			want: []byte{174, 13, 10},
		},
		{
			name: "unsupported runes",
			args: args{
				[]byte("�"),
			},
			want: []byte{26},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			byteReader := bytes.NewReader(tt.args.src)
			encoded := transformReader(byteReader)

			if !reflect.DeepEqual(encoded, tt.want) {
				t.Logf("final byte array does not match\n")
				t.Logf("src : %v\n", tt.args.src)
				t.Logf("enc : %v\n", encoded)
				t.Logf("want: %v\n", tt.want)
				t.Fail()
			}
		})
	}
}

func transformReader(reader io.Reader) []byte {
	transformedReader := transform.NewReader(reader, Windows1252Crlf())
	encodedBytes := make([]byte, 4096)
	read, err := transformedReader.Read(encodedBytes)
	if err != nil {
		log.Fatal(err)
	}
	return encodedBytes[:read]
}
