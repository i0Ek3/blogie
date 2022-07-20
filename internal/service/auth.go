package service

import "errors"

type AuthRequest struct {
	// NOTES: use header to replace form field here
	AppKey    string `header:"app_key" binding:"required"`
	AppSecret string `header:"app_secret" binding:"required"`
}

func (svc *Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(param.AppKey, param.AppSecret)
	if err != nil {
		return err
	}
	if auth.ID > 0 {
		return nil
	}
	return errors.New("auth info does dont exist")
}
