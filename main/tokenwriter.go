package main

import (
	"strings"

	pb "github.com/pawelstiasny/territory/indexers/go/pb"
)

type TokenWriter struct {
	tokens []*pb.Token
	text   strings.Builder
	offset uint32
}

func (tw *TokenWriter) AppendWithReal(typ pb.TokenType, text string, realOffset *uint32, realLine *uint32, href *pb.UniHref) {
	var h *pb.Token_UniHref
	if href != nil {
		h = &pb.Token_UniHref{UniHref: href}
	}

	tw.tokens = append(tw.tokens, &pb.Token{
		Offset:     tw.offset,
		Type:       typ,
		RealOffset: realOffset,
		RealLine:   realLine,
		Href:       h,
	})
	tw.offset += uint32(len(text))
	tw.text.WriteString(text)
}

func (tw *TokenWriter) Append(typ pb.TokenType, text string, href *pb.UniHref) {
	var h *pb.Token_UniHref
	if href != nil {
		h = &pb.Token_UniHref{UniHref: href}
	}

	tw.tokens = append(tw.tokens, &pb.Token{
		Offset: tw.offset,
		Type:   typ,
		Href:   h,
	})
	tw.offset += uint32(len(text))
	tw.text.WriteString(text)
}

func (tw *TokenWriter) Get() ([]*pb.Token, string) {
	return tw.tokens, tw.text.String()
}
