package cachehttp

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CacheCheck 检查缓存命中接口
// @Summary 检查缓存命中接口
// @Description 检查缓存中是否有请求的值 有就返回没有将请求转发
// @Tags 缓存相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param infotype query string true "查询类型"
// @Param level query string true "查询等级"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /hit [get]
func (s *Server) CacheCheck(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	m := c.Request.Method

	if m == http.MethodGet {
		b, _ := s.Get(key)
		if len(b) == 0 {
			c.Set("miss", true) // 需要查数据库
			return
		}

		c.JSON(http.StatusOK, string(b))
		c.Set("miss", false) // 不需要查数据库
		return
	}

	c.JSON(http.StatusMethodNotAllowed, nil)
}

// CacheCheck 检查缓存命中接口
// @Summary 检查缓存命中接口
// @Description 检查缓存中是否有请求的值 有就返回没有将请求转发
// @Tags 缓存相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param infotype query string true "查询类型"
// @Param level query string true "查询等级"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /hit [get]
func (s *Server) UpdateHandler(c *gin.Context) {
	key := c.Param("key")

	if key == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	m := c.Request.Method
	if m == http.MethodPut {
		b, _ := io.ReadAll(c.Request.Body)
		if len(b) != 0 {
			e := s.Set(key, b)
			if e != nil {
				log.Println(e)
				c.JSON(http.StatusInternalServerError, nil)
			}
		}
		return
	}

	c.JSON(http.StatusMethodNotAllowed, nil)
}

// CacheCheck 检查缓存命中接口
// @Summary 检查缓存命中接口
// @Description 检查缓存中是否有请求的值 有就返回没有将请求转发
// @Tags 缓存相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param infotype query string true "查询类型"
// @Param level query string true "查询等级"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /hit [get]
func (s *Server) DeleteHandler(c *gin.Context) {
	key := c.Param("key")

	if key == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	m := c.Request.Method
	if m == http.MethodDelete {
		e := s.Del(key)
		if e != nil {
			log.Println(e)
			c.JSON(http.StatusInternalServerError, nil)
		}
		return
	}

	c.JSON(http.StatusMethodNotAllowed, nil)
}
