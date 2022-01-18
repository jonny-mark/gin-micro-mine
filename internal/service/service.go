/**
 * @author jiangshangfang
 * @date 2022/1/10 11:43 AM
 **/
package service

import "gin/internal/repository"

var Svc Service

type Service interface {

}

// service struct
type service struct {
	repo repository.Repository
}

// New init service
func New(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}
