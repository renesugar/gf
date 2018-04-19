// Copyright 2018 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.
// 参数管理.

package ghttp

import (
    "time"
    "errors"
    "strconv"
    "strings"
    "net/http"
    "crypto/tls"
    "path/filepath"
)

// http server setting设置
// 注意使用该方法进行http server配置时，需要配置所有的配置项，否则没有配置的属性将会默认变量为空
func (s *Server)SetConfig(c ServerConfig) error {
    if s.status == 1 {
        return errors.New("server config cannot be changed while running")
    }
    if c.Handler == nil {
        c.Handler = http.HandlerFunc(s.defaultHttpHandle)
    }
    s.config = c
    // 需要处理server root最后的目录分隔符号
    if s.config.ServerRoot != "" {
        s.SetServerRoot(s.config.ServerRoot)
    }
    // 必需设置默认值的属性
    if len(s.config.IndexFiles) < 1 {
        s.SetIndexFiles(defaultServerConfig.IndexFiles)
    }
    if s.config.ServerAgent == "" {
        s.SetServerAgent(defaultServerConfig.ServerAgent)
    }
    return nil
}

// 设置http server参数 - Addr
func (s *Server)SetAddr(addr string) error {
    if s.status == 1 {
        return errors.New("server config cannot be changed while running")
    }
    s.config.Addr = addr
    return nil
}

// 设置http server参数 - Port
func (s *Server)SetPort(port int) error {
    if s.status == 1 {
        return errors.New("server config cannot be changed while running")
    }
    s.config.Addr = ":" + strconv.Itoa(port)
    return nil
}

// 设置http server参数 - TLSConfig
func (s *Server)SetTLSConfig(tls *tls.Config) error {
    if s.status == 1 {
        return errors.New("server config cannot be changed while running")
    }
    s.config.TLSConfig = tls
    return nil
}

// 设置http server参数 - ReadTimeout
func (s *Server)SetReadTimeout(t time.Duration) error {
    if s.status == 1 {
        return errors.New("server config cannot be changed while running")
    }
    s.config.ReadTimeout = t
    return nil
}

// 设置http server参数 - WriteTimeout
func (s *Server)SetWriteTimeout(t time.Duration) error {
    if s.status == 1 {
        return errors.New("server config cannot be changed while running")
    }
    s.config.WriteTimeout = t
    return nil
}

// 设置http server参数 - IdleTimeout
func (s *Server)SetIdleTimeout(t time.Duration) error {
    if s.status == 1 {
        return errors.New("server config cannot be changed while running")
    }
    s.config.IdleTimeout = t
    return nil
}

// 设置http server参数 - MaxHeaderBytes
func (s *Server)SetMaxHeaderBytes(b int) error {
    if s.status == 1 {
        return errors.New("server config cannot be changed while running")
    }
    s.config.MaxHeaderBytes = b
    return nil
}

// 设置http server参数 - IndexFiles
func (s *Server)SetIndexFiles(index []string) error {
    if s.status == 1 {
        return errors.New("server config cannot be changed while running")
    }
    s.config.IndexFiles = index
    return nil
}

// 设置http server参数 - IndexFolder
func (s *Server)SetIndexFolder(index bool) error {
    if s.status == 1 {
        return errors.New("server config cannot be changed while running")
    }
    s.config.IndexFolder = index
    return nil
}

// 设置http server参数 - ServerAgent
func (s *Server)SetServerAgent(agent string) error {
    if s.status == 1 {
        return errors.New("server config cannot be changed while running")
    }
    s.config.ServerAgent = agent
    return nil
}

// 设置http server参数 - ServerRoot
func (s *Server)SetServerRoot(root string) error {
    if s.status == 1 {
        return errors.New("server config cannot be changed while running")
    }
    s.config.ServerRoot  = strings.TrimRight(root, string(filepath.Separator))
    return nil
}

// 设置http server参数 - CookieMaxAge
func (s *Server)SetCookieMaxAge(maxage int) {
    s.cookieMaxAge.Set(maxage)
}

// 设置http server参数 - SessionMaxAge
func (s *Server)SetSessionMaxAge(maxage int) {
    s.sessionMaxAge.Set(maxage)
}

// 设置http server参数 - SessionIdName
func (s *Server)SetSessionIdName(name string) {
    s.sessionIdName.Set(name)
}

// 设置日志目录
func (s *Server)SetLogPath(path string) error {
    if err := s.logger.SetPath(path); err != nil {
        return err
    }
    s.logPath.Set(path)
    return nil
}

// 设置是否开启access log日志功能
func (s *Server)SetAccessLogEnabled(enabled bool) {
    s.accessLogEnabled.Set(enabled)
}

// 设置是否开启error log日志功能
func (s *Server)SetErrorLogEnabled(enabled bool) {
    s.errorLogEnabled.Set(enabled)
}

// 设置日志写入的回调函数
func (s *Server) SetLogHandler(handler func(r *Request, error ... interface{})) {
    s.logHandler.Set(handler)
}

// 获取日志写入的回调函数
func (s *Server) GetLogHandler() func(r *Request, error ... interface{}) {
    if v := s.logHandler.Val(); v != nil {
        return v.(func(r *Request, error ... interface{}))
    }
    return nil
}

// 获取日志目录
func (s *Server)GetLogPath() string {
    return s.logPath.Val()
}

// access log日志功能是否开启
func (s *Server)IsAccessLogEnabled() bool {
    return s.accessLogEnabled.Val()
}

// error log日志功能是否开启
func (s *Server)IsErrorLogEnabled() bool {
    return s.errorLogEnabled.Val()
}

// 获取
func (s *Server) GetName() string {
    return s.name
}

// 获取http server参数 - CookieMaxAge
func (s *Server)GetCookieMaxAge() int {
    return s.cookieMaxAge.Val()
}

// 获取http server参数 - SessionMaxAge
func (s *Server)GetSessionMaxAge() int {
    return s.sessionMaxAge.Val()
}

// 获取http server参数 - SessionIdName
func (s *Server)GetSessionIdName() string {
    return s.sessionIdName.Val()
}

