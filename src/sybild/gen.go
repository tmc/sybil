//go:generate protoc -I . -I ../../../../.. sybild.proto --go_out=plugins=grpc:pb

package sybild
