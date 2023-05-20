package user

import "blog-api/rest/req"

func (s *Service) Me(ctx *req.Ctx) error {
	a, err := s.tokenManager.ValidateJWTToken(ctx.BearerToken())
	if err != nil {
		return err
	}
	ctx.Writer.Write([]byte(a.Username))
	return nil
}
