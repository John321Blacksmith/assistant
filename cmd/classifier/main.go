// This package is an entrypoint
// For app startup and responding
// To client requests

package main

import (
	"classifier/engine"
	pb "classifier/proto"
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

type ClassifierServer struct {
	classifier *engine.Classifier
	pb.UnimplementedClassifierServer
}

func (s *ClassifierServer) GetMainContext(ctx context.Context, in *pb.BotRequest) (*pb.Response, error) {
	slog.Info("classifier ->", "GetMainContext", "PROCESSING DATA")
	fmt.Println(s.classifier.GetMainContext())
	sentences := s.classifier.ProcessInput(in.Input)
	s.classifier.RecognizeSentences(sentences)
	mainContext := s.classifier.GetMainContext()
	engine.DiscardCollection(s.classifier)
	return &pb.Response{Output: mainContext}, nil
}

func main() {
	dataManager := engine.NewDataManager("./categories.json")
	dataset, err := dataManager.LoadDataset()
	if err != nil {
		log.Fatalf("Couldn't load the dataset: %s", err)
	}
	classifier := engine.NewClassifier(dataset)
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error setting up listener: %s", err)
	}
	s := grpc.NewServer()
	pb.RegisterClassifierServer(
		s,
		&ClassifierServer{
			classifier: classifier,
		},
	)
	slog.Info("Classifier Server is running on port", "port", listener.Addr().String())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Error setting up the Classifier Server: %s", err)
	}
}
