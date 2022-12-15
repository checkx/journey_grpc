package main

import (
	"fmt"
	pb "learngrpc/pb"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {

	products := &pb.Products{
		Data: []*pb.Product{
			{
				Id:    1,
				Name:  "Speaker",
				Price: 100000,
				Stock: 100,
				Category: &pb.Category{
					Id:   1,
					Name: "Electronics",
				},
			},
			{
				Id:    2,
				Name:  "Mouse",
				Price: 100000,
				Stock: 100,
				Category: &pb.Category{
					Id:   2,
					Name: "Computer",
				},
			},
		},
	}

	data, err := proto.Marshal(products)
	if err != nil {
		log.Fatal("Marshal Error", err)
	}

	fmt.Println(data)

	testProducts := &pb.Products{}
	if err = proto.Unmarshal(data, testProducts); err != nil {
		log.Fatal("Unmarshal error", err)
	}

	for _, product := range testProducts.GetData() {
		fmt.Println(product.GetName())
		fmt.Println(product.GetCategory().GetName())
	}
}
