package net

import (
	inet "github.com/libp2p/go-libp2p-net"
	protobufCodec "github.com/multiformats/go-multicodec/protobuf"
	"github.com/multiformats/go-multicodec"
	"bufio"
	"github.com/gangwn/dvc/pkg/protocol"
)

type InStream struct {
	Stream  inet.Stream
	decoder multicodec.Decoder
	r       *bufio.Reader
}

func NewInStream(s inet.Stream) *InStream {
	reader := bufio.NewReader(s)
	decoder := protobufCodec.Multicodec(nil).Decoder(reader)

	stream := &InStream{s,decoder, reader}
	return stream
}

func (s *InStream) DecodeMessage() *dvc_protocol.DVCMessage {
	message := &dvc_protocol.DVCMessage{}
	err := s.decoder.Decode(message)
	if err == nil {
		return message
	}

	return nil
}