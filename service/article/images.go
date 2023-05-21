package article

import (
	"blog-api/rest/req"
	"blog-api/tools/tokenmanager"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
)

const (
	FileNameLength = 16
)

type UploadImageRequest struct {
	Image string `json:"image"`
}

type UploadImageResponse struct {
	Path string `json:"path"`
}

func (s *Service) UploadImage(ctx *req.Ctx) error {
	var r UploadImageRequest

	ctx.ParseJSON(&r)

	fileName, err := tokenmanager.GenerateRandomSalt(FileNameLength)
	if err != nil {
		return err
	}

	base64Data := strings.TrimPrefix(r.Image, "data:image/png;base64,")

	img, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/%s", s.assetsDir, fileName)

	os.WriteFile(path, img, 0644)

	return ctx.JSON(UploadImageResponse{Path: "assets/" + fileName})
}

func (s *Service) SendFile(ctx *req.Ctx) error {
	fileName := chi.URLParam(ctx.Request, "file")

	fileBytes, err := os.ReadFile(fmt.Sprintf("%s/%s", s.assetsDir, fileName))
	if err != nil {
		return err
	}

	ctx.Writer.Header().Set("Content-Type", "application/octet-stream")
	ctx.Writer.Write(fileBytes)

	return nil
}
