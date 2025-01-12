// Copyright 2023 The frp Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type ClientPluginOptions interface{}

type TypedClientPluginOptions struct {
	Type string `json:"type"`
	ClientPluginOptions
}

func (c *TypedClientPluginOptions) UnmarshalJSON(b []byte) error {
	if len(b) == 4 && string(b) == "null" {
		return errors.New("type is required")
	}

	typeStruct := struct {
		Type string `json:"type"`
	}{}
	if err := json.Unmarshal(b, &typeStruct); err != nil {
		return err
	}

	c.Type = typeStruct.Type

	v, ok := clientPluginOptionsTypeMap[typeStruct.Type]
	if !ok {
		return fmt.Errorf("unknown plugin type: %s", typeStruct.Type)
	}
	if err := json.Unmarshal(b, v); err != nil {
		return err
	}
	c.ClientPluginOptions = v
	return nil
}

const (
	PluginHTTP2HTTPS       = "http2https"
	PluginHTTPProxy        = "http_proxy"
	PluginHTTPS2HTTP       = "https2http"
	PluginHTTPS2HTTPS      = "https2https"
	PluginSocks5           = "socks5"
	PluginStaticFile       = "static_file"
	PluginUnixDomainSocket = "unix_domain_socket"
)

var clientPluginOptionsTypeMap = map[string]reflect.Type{
	PluginHTTP2HTTPS:       reflect.TypeOf(HTTP2HTTPSPluginOptions{}),
	PluginHTTPProxy:        reflect.TypeOf(HTTPProxyPluginOptions{}),
	PluginHTTPS2HTTP:       reflect.TypeOf(HTTPS2HTTPPluginOptions{}),
	PluginHTTPS2HTTPS:      reflect.TypeOf(HTTPS2HTTPSPluginOptions{}),
	PluginSocks5:           reflect.TypeOf(Socks5PluginOptions{}),
	PluginStaticFile:       reflect.TypeOf(StaticFilePluginOptions{}),
	PluginUnixDomainSocket: reflect.TypeOf(UnixDomainSocketPluginOptions{}),
}

type HTTP2HTTPSPluginOptions struct {
	LocalAddr         string           `json:"localAddr,omitempty"`
	HostHeaderRewrite string           `json:"hostHeaderRewrite,omitempty"`
	RequestHeaders    HeaderOperations `json:"requestHeaders,omitempty"`
}

type HTTPProxyPluginOptions struct {
	HTTPUser     string `json:"httpUser,omitempty"`
	HTTPPassword string `json:"httpPassword,omitempty"`
}

type HTTPS2HTTPPluginOptions struct {
	LocalAddr         string           `json:"localAddr,omitempty"`
	HostHeaderRewrite string           `json:"hostHeaderRewrite,omitempty"`
	RequestHeaders    HeaderOperations `json:"requestHeaders,omitempty"`
	CrtPath           string           `json:"crtPath,omitempty"`
	KeyPath           string           `json:"keyPath,omitempty"`
}

type HTTPS2HTTPSPluginOptions struct {
	LocalAddr         string           `json:"localAddr,omitempty"`
	HostHeaderRewrite string           `json:"hostHeaderRewrite,omitempty"`
	RequestHeaders    HeaderOperations `json:"requestHeaders,omitempty"`
	CrtPath           string           `json:"crtPath,omitempty"`
	KeyPath           string           `json:"keyPath,omitempty"`
}

type Socks5PluginOptions struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type StaticFilePluginOptions struct {
	LocalPath    string `json:"localPath,omitempty"`
	StripPrefix  string `json:"stripPrefix,omitempty"`
	HTTPUser     string `json:"httpUser,omitempty"`
	HTTPPassword string `json:"httpPassword,omitempty"`
}

type UnixDomainSocketPluginOptions struct {
	UnixPath string `json:"unixPath,omitempty"`
}
