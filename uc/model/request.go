package model

import (
	"io"
)

// PostRequest request interface for post
type PostRequest interface {
	Url() string
	Encode() []byte
}

// GetRequest request interface for get
type GetRequest interface {
	Url() string
	Encode() string
}

// UploadField multipart/form-data post request field struct
type UploadField struct {
	// Key field key
	Key string
	// Value field value
	Value string
	// Reader upload file reader
	Reader io.Reader
}

// UploadRequest multipart/form-data reqeust interface
type UploadRequest interface {
	Url() string
	// Encode encode request to UploadFields
	Encode() []UploadField
}

type Header struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}


type HeaderTarget struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Target   string `json:"target"`
}
