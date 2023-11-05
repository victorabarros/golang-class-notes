package grpc

import (
	"github.com/ricardoerikson/sgg/internal/transport/entities"
	"github.com/ricardoerikson/sgg/internal/transport/pb"
	"github.com/ricardoerikson/sgg/pkg/model"
)

// ConvertCreateArticleRequestFromProtoToEntity creates an entity model from a proto model
func ConvertCreateArticleRequestFromProtoToEntity(proto *pb.CreateArticleRequest) entities.CreateArticleRequest {
	entity := entities.CreateArticleRequest{
		Title:          proto.GetTitle(),
		Description:    proto.GetDescription(),
		ExpirationDate: proto.GetExpirationDate(),
	}

	for _, i := range proto.GetImages() {
		image := entities.AttachImageRequest{Path: i.GetPath()}
		entity.Images = append(entity.Images, image)
	}

	return entity
}

// ConvertArticleFromModelToProto converts an article model to proto
func ConvertArticleFromModelToProto(m model.Article) *pb.Article {
	article := &pb.Article{
		Id:             m.ID.String(),
		Title:          m.Title,
		Description:    m.Description,
		ExpirationDate: m.ExpirationDate.Format("2006-01-02"),
	}

	for _, i := range m.Images {
		image := pb.Image{Path: i.Path}
		article.Images = append(article.Images, &image)
	}

	return article
}

func ConvertAttachImageRequestFromProtoToEntity(proto *pb.AttachImageRequest) entities.AttachImageRequest {
	return entities.AttachImageRequest{Path: proto.Path}
}
