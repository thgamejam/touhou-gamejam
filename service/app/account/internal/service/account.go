package service

import (
	"context"
	"regexp"
	"service/app/account/internal/biz"

	pb "service/api/account/v1"
)

const (
	// 按照 RFC 5322 的邮箱地址正则规则
	matchEMail = `([!#-'*+/-9=?A-Z^-~-]+(\.[!#-'*+/-9=?A-Z^-~-]+)*|"([]!#-[^-~ \t]|(\\[\t -~]))+")@([!#-'*+/-9=?A-Z^-~-]+(\.[!#-'*+/-9=?A-Z^-~-]+)*|\[[\t -Z^-~]*])`
)

func (s *AccountService) GetKey(ctx context.Context, req *pb.GetKeyReq) (*pb.GetKeyReply, error) {
	key, err := s.uc.GetRandomlyKey(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetKeyReply{
		Hash: key.Hash,
		Key:  key.Key,
	}, nil
}

func (s *AccountService) ExistAccountEMail(
	ctx context.Context, req *pb.ExistAccountEMailReq) (*pb.ExistAccountEMailReply, error) {

	ok, err := s.uc.ExistAccountEMail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	return &pb.ExistAccountEMailReply{Ok: ok}, nil
}

func (s *AccountService) PrepareCreateEMailAccount(
	ctx context.Context, req *pb.PrepareCreateEMailAccountReq) (*pb.PrepareCreateEMailAccountReply, error) {

	passwdCT := &biz.PasswordCiphertext{
		KeyHash:    req.Hash,
		Ciphertext: req.Ciphertext,
	}

	sid, err := s.uc.PrepareCreateEMailAccount(ctx, req.Email, passwdCT)
	if err != nil {
		return nil, err
	}

	return &pb.PrepareCreateEMailAccountReply{Sid: sid}, nil
}

func (s *AccountService) FinishCreateEMailAccount(
	ctx context.Context, req *pb.FinishCreateEMailAccountReq) (*pb.FinishCreateEMailAccountReply, error) {

	id, err := s.uc.FinishCreateEMailAccount(ctx, req.Sid)
	if err != nil {
		return nil, err
	}
	return &pb.FinishCreateEMailAccountReply{Id: id}, nil
}

func (s *AccountService) GetAccount(ctx context.Context, req *pb.GetAccountReq) (*pb.GetAccountReply, error) {
	account, err := s.uc.GetAccount(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetAccountReply{
		Uuid:    account.UUID.String(),
		Email:   account.Email,
		TelCode: uint32(account.Phone.TelCode),
		Phone:   account.Phone.Phone,
		Status:  uint32(account.Status),
	}, nil
}

func (s *AccountService) VerifyPassword(
	ctx context.Context, req *pb.VerifyPasswordReq) (*pb.VerifyPasswordReply, error) {

	// 正则判断req.Username是否为邮箱
	isEMail, err := regexp.MatchString(matchEMail, req.Username)
	if err != nil {
		return nil, err
	}
	if isEMail {
		passwdCT := &biz.PasswordCiphertext{
			KeyHash:    req.Hash,
			Ciphertext: req.Ciphertext,
		}
		id, ok, err := s.uc.VerifyPasswordByEMail(ctx, req.Username, passwdCT)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, pb.ErrorIncorrectAccount("账号或密码错误")
		}
		return &pb.VerifyPasswordReply{Id: id}, nil
	}
	return nil, pb.ErrorContentMissing("账号参数错误")
}

func (s *AccountService) SavePassword(ctx context.Context, req *pb.SavePasswordReq) (*pb.SavePasswordReply, error) {
	passwdCT := &biz.PasswordCiphertext{
		KeyHash:    req.Hash,
		Ciphertext: req.Ciphertext,
	}
	err := s.uc.SavePassword(ctx, req.Id, passwdCT)
	if err != nil {
		return nil, err
	}
	return &pb.SavePasswordReply{}, nil
}

func (s *AccountService) CreateSession(ctx context.Context, req *pb.CreateSessionReq) (*pb.CreateSessionReply, error) {
	sid, err := s.uc.CreateSession(ctx, req.Id, req.Ip)
	if err != nil {
		return nil, err
	}
	return &pb.CreateSessionReply{Sid: sid}, nil
}

func (s *AccountService) VerifySession(ctx context.Context, req *pb.VerifySessionReq) (*pb.VerifySessionReply, error) {
	ok, err := s.uc.ExistSession(ctx, req.Id, req.Sid)
	if err != nil {
		return nil, err
	}
	return &pb.VerifySessionReply{Ok: ok}, nil
}

func (s *AccountService) CloseSession(ctx context.Context, req *pb.CloseSessionReq) (*pb.CloseSessionReply, error) {
	err := s.uc.CloseSession(ctx, req.Id, req.Sid)
	if err != nil {
		return nil, err
	}
	return &pb.CloseSessionReply{Ok: true}, nil
}
