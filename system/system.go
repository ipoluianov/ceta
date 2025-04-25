package system

import (
	"archive/zip"
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/ipoluianov/ceta/httpserver"
)

type System struct {
	httpServer *httpserver.HttpServer
}

func NewSystem() *System {
	var c System
	return &c
}

func (c *System) Start() {
	c.httpServer = httpserver.NewHttpServer()
	c.httpServer.Start()

	go c.ThWork()
}

func (c *System) Stop() {
	c.httpServer.Stop()
}

func (c *System) CreateZipWithJSON(jsonData []byte) ([]byte, error) {
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	fileWriter, err := zipWriter.Create("data.json")
	if err != nil {
		return nil, err
	}

	_, err = fileWriter.Write(jsonData)
	if err != nil {
		return nil, err
	}

	if err := zipWriter.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c *System) extractFileFromZipImage(zipImage []byte) ([]byte, error) {
	zipReader, err := zip.NewReader(bytes.NewReader(zipImage), int64(len(zipImage)))
	if err != nil {
		return nil, err
	}

	for _, f := range zipReader.File {
		if f.Name == "data.json" {
			rc, err := f.Open()
			if err != nil {
				return nil, err
			}
			defer rc.Close()

			return io.ReadAll(rc)
		}
	}

	return nil, nil
}

func (c *System) fetchZipFileFromServer(url string) ([]byte, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	bs, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func (c *System) ThWork() {
	for {
		time.Sleep(1 * time.Second)
	}
}
