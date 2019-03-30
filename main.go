package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"model-service/models"
	pb "model-service/protos"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

var db *gorm.DB

type server struct{}

func (s *server) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterReply, error) {
	var m models.Model

	db.Where("email = ?", in.Email).First(&m)

	if m.ID != 0 {
		log.Printf("model already exists with email: %s", in.Email)

		return &pb.RegisterReply{
			Response: pb.Response_EMAIL_TAKEN,
		}, nil
	}

	password, err := bcrypt.GenerateFromPassword([]byte(in.Password), 10)

	if err != nil {
		log.Fatalf("failed to generate password: %v", err)

		return &pb.RegisterReply{
			Response: pb.Response_INTERNAL_ERROR,
		}, nil
	}

	m = models.Model{
		Email:    in.Email,
		Password: string(password),
	}

	db.Create(&m)

	return &pb.RegisterReply{}, nil
}

func (s *server) Activate(ctx context.Context, in *pb.ActivateRequest) (*pb.ActivateReply, error) {
	var m models.Model

	db.First(&m, in.Id)

	if m.ID == 0 {
		log.Printf("model could not be found with id: %d", in.Id)

		return &pb.ActivateReply{
			Response: pb.Response_NOT_FOUND,
		}, nil
	}

	m.IsActivated = true

	db.Save(&m)

	return &pb.ActivateReply{}, nil
}

func initDB() {
	var err error

	db, err = gorm.Open("mysql", fmt.Sprintf(
		"%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local",
		"admin",
		"root",
		"tcp(localhost:3333)",
		"model",
	))
	// db, err := gorm.Open("mysql", fmt.Sprintf(
	// 	"%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local",
	// 	os.Getenv("DB_USERNAME"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_NAME"),
	// ))

	// defer db.Close()

	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	db.AutoMigrate(&models.Model{})
}

func initGRPC() {
	listener, err := net.Listen("tcp", ":3000")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterModelServer(s, &server{})

	err = s.Serve(listener)

	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	initDB()
	initGRPC()
}
