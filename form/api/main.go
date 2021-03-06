package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"mime"
	"mime/multipart"
	"strings"

	api "github.com/lack-io/vine/proto/api"
	"github.com/lack-io/vine/service"

	proto "github.com/lack-io/vine-example/form/api/proto"
)

type Form struct{}

func (f *Form) Submit(ctx context.Context, req *api.Request, rsp *api.Response) error {
	rsp.Body = fmt.Sprintf("got your values %+v", req.Post)
	return nil
}

func (f *Form) Multipart(ctx context.Context, req *api.Request, rsp *api.Response) error {
	ct := strings.Join(req.Header["Content-Type"].Values, ",")
	mt, p, err := mime.ParseMediaType(ct)
	if err != nil {
		return err
	}
	if !strings.HasPrefix(mt, "multipart/") {
		return fmt.Errorf("%v does not contain multipart", mt)
	}
	r := multipart.NewReader(bytes.NewReader([]byte(req.Body)), p["boundary"])
	form, err := r.ReadForm(32 << 20)
	if err != nil {
		return err
	}
	rsp.Body = fmt.Sprintf("got your values %+v", form)
	return nil
}

func main() {
	server := service.NewService(
		service.Name("go.vine.api.form"),
	)

	server.Init()

	proto.RegisterFormHandler(server.Server(), new(Form))

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
