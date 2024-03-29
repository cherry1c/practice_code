package main

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	grpc_transport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"net"
	book "rpc_frame/pb"
)

type BookServer struct {
	book.UnimplementedBookServiceServer
	bookListHandler grpc_transport.Handler
	bookInfoHandler grpc_transport.Handler
}

// 通过grpc调用GetBookInfo时,GetBookInfo只做数据透传, 调用BookServer中对应Handler.ServeGRPC转交给go-kit处理
func (s *BookServer) GetBookInfo(ctx context.Context, in *book.BookInfoParams) (*book.BookInfo, error) {
	_, rsp, err := s.bookInfoHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*book.BookInfo), err
}

// 通过grpc调用GetBookList时,GetBookList只做数据透传, 调用BookServer中对应Handler.ServeGRPC转交给go-kit处理
func (s *BookServer) GetBookList(ctx context.Context, in *book.BookListParams) (*book.BookList, error) {
	_, rsp, err := s.bookListHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*book.BookList), err
}

// 创建bookList的EndPoint
func makeGetBookListEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//请求列表时返回 书籍列表
		bl := new(book.BookList)
		bl.BookList = append(bl.BookList, &book.BookInfo{BookId: 1, BookName: "Go入门到精通"})
		bl.BookList = append(bl.BookList, &book.BookInfo{BookId: 2, BookName: "go-kit入门到精通"})
		bl.BookList = append(bl.BookList, &book.BookInfo{BookId: 2, BookName: "go-micro入门到精通"})
		return bl, nil
	}
}

// 创建bookInfo的EndPoint
func makeGetBookInfoEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//请求详情时返回 书籍信息
		req := request.(*book.BookInfoParams)
		b := new(book.BookInfo)
		b.BookId = req.BookId
		b.BookName = "Go与微服务"
		return b, nil
	}
}

func decodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func encodeResponse(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func main() {
	//包装BookServer

	bookServer := new(BookServer)
	//创建bookList的Handler
	bookListHandler := grpc_transport.NewServer(
		makeGetBookListEndpoint(),
		decodeRequest,
		encodeResponse,
	)
	//bookServer 增加 go-kit流程的 bookList处理逻辑
	bookServer.bookListHandler = bookListHandler

	//创建bookInfo的Handler
	bookInfoHandler := grpc_transport.NewServer(
		makeGetBookInfoEndpoint(),
		decodeRequest,
		encodeResponse,
	)
	//bookServer 增加 go-kit流程的 bookInfo处理逻辑
	bookServer.bookInfoHandler = bookInfoHandler

	//启动grpc服务
	serviceAddress := ":50052"
	ls, _ := net.Listen("tcp", serviceAddress)
	gs := grpc.NewServer()
	book.RegisterBookServiceServer(gs, bookServer)
	gs.Serve(ls)
}

// Todo go-kit参考：https://blog.csdn.net/lihao19910921/article/details/80166399
