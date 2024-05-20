package pkg

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func UploadZipFile(ctx *fiber.Ctx, url *string) (int16, error) {
	multipartFile, err := ctx.FormFile("file")
	if err != nil {
		return 500, err
	}

	file, err := multipartFile.Open()
	if err != nil {
		return 500, err
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return 500, err
	}

	req, err := http.NewRequest(http.MethodPut, *url, buf)
	if err != nil {
		return 500, err
	}

	req.Header.Set("Content-Type", "application/zip")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return 400, err
	}
	defer res.Body.Close()

	return int16(res.StatusCode), nil
}
