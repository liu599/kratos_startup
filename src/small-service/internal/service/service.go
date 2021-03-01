package service

import (
	"context"
	"fmt"
    "strings"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	pb "small-service/api"
	"small-service/internal/dao"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.DemoServer), new(*Service)))

// Service service.
type Service struct {
	ac  *paladin.Map
	dao dao.Dao
}

// New new a service and return.
func New(d dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		ac:  &paladin.TOML{},
		dao: d,
	}
	cf = s.Close
	err = paladin.Watch("application.toml", s.ac)
	return
}

// SayHello grpc demo func.
func (s *Service) SayHello(ctx context.Context, req *pb.HelloReq) (reply *empty.Empty, err error) {
	reply = new(empty.Empty)
	fmt.Printf("hello %s", req.Name)
	return
}

// SayHelloURL bm demo func.
func (s *Service) SayHelloURL(ctx context.Context, req *pb.HelloReq) (reply *pb.HelloResp, err error) {
	reply = &pb.HelloResp{
		Content: "hello " + req.Name,
	}
	fmt.Printf("hello url %s", req.Name)
	return
}

func (s *Service) RequestItem(ctx context.Context, req *pb.HelloReq) (reply *pb.HelloResp, err error) {
	d, err2 := s.dao.CardInfo(ctx, int(req.Number))
	if err2!=nil {
		reply = &pb.HelloResp{
			Content: "Error: " + err2.Error(),
		}
		return
	}

	var color string
	switch d.Color {
		case "红":
			color = pb.CardColor(0).String()
			break
		case "蓝":
			color = pb.CardColor(1).String()
			break
		case "绿":
			color = pb.CardColor(2).String()
			break
		case "黄":
			color = pb.CardColor(3).String()
			break
		default:
			color = pb.CardColor(0).String()
	}

	wsid := strings.Replace(d.Wsid, "---", "/", -1)

	b1 := pb.CardResp{
		Wid:                  int32(d.Wid),
		Wsid:                 wsid,
		Cardname:             d.CardName,
		Cardcat:              d.CardCat,
		Cardcolor:            color,
		Prop:                 d.Prop,
		Rare:                 d.Rare,
		Level:                int32(d.Level),
		Cost:                 int32(d.Cost),
		Judge:                int32(d.Judge),
		Soul:                 int32(d.Soul),
		Attack:               int32(d.Attack),
		Series:               d.Series,
		Des1:                 d.Des1,
		Des2:                 d.Des2,
		Des3:                 d.Des3,
		Cover1:               d.Cover1,
		Cover2:               d.Cover2,
		Cover3:               d.Cover3,
		Rel1:                 d.Rel1,
		Rel2:                 d.Rel2,
		Cat:                  d.Cat,
	}
    var b2 []*pb.CardResp
	b2 = append(b2, &b1)

	reply = &pb.HelloResp{
		Content: "hello " + d.CardName,
		CardInfo: b2,
	}
	return
}



// Ping ping the resource.
func (s *Service) Ping(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
}
