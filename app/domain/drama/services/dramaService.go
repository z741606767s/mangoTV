package services

import "mangoTV/app/svc"

type DramaService struct {
	ctx *svc.ServiceContext
}

func NewDramaService(ctx *svc.ServiceContext) *DramaService {
	return &DramaService{
		ctx: ctx,
	}
}
