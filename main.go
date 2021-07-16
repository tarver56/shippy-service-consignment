package main

// Import the generated protobuf code

const (
	port = ":50051"
)

type repository interface {
	Create(*pb.Consignment)
}
