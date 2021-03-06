package main

import (
	"context"
	micro "github.com/micro/go-micro/v2"
	pb "github.com/tanglei25/laracom/demo-service/proto/demo"
	"log"
)
func main() {
	service := micro.NewService(micro.Name("laracom.demo.cli"))
	service.Init()

	client := pb.NewDemoService("laracom.demo.service", service.Client())
	rsp, err := client.SayHello(context.TODO(), &pb.DemoRequest{Name: "学院君"})
	if err != nil {
		log.Fatalf("服务调用失败：%v", err)
		return
	}
	log.Println(rsp.Text)
}
