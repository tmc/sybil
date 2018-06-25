//go:generate protoc -I . -I ../../../../.. sybil.proto --go_out=plugins=grpc:../sybilpb/
//go:generate protoc -I . -I ../../../../.. internal.proto --gogofast_out=../internal/internalpb

package proto
