package main

func main(){
	gRPCServer := NewGRPCServer(":4000");
	gRPCServer.Run()
}