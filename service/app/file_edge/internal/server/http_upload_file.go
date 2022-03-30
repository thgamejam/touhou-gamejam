package server

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"

	kratosHTTP "github.com/go-kratos/kratos/v2/transport/http"

	v1 "service/api/file_edge/v1"
	"service/app/file_edge/internal/conf"
	"service/app/file_edge/internal/service"
	pkgConf "service/pkg/conf"
	"service/pkg/jwt"
)

func uploadFile(c *pkgConf.Service, efs *service.FileEdgeService) func(ctx kratosHTTP.Context) error {
	env := conf.GetEnv()
	tmp := env.TemporaryFileDirectory
	return func(ctx kratosHTTP.Context) error {
		req := ctx.Request()

		// 验证 upload token
		signedToken := strings.TrimPrefix(req.URL.String(), PrefixURL)
		claims, success := jwt.ValidateUploadToken(signedToken, []byte(c.Data.ObjectStorage.SecretAccessKey))
		if !success {
			return v1.ErrorUnauthorizedUpload("upload authentication failed.")
		}

		// 从http请求中获取文件
		file, handler, err := req.FormFile("file")
		if err != nil {
			return v1.ErrorContentMissing("upload file error.")
		}
		defer file.Close()

		// 计算md5值
		{
			hash := md5.New()
			_, err = io.Copy(hash, file)
			if err != nil {
				return v1.ErrorContentMissing("damaged file..")
			}
			md5HashByte := hash.Sum(nil)
			md5Hash := hex.EncodeToString(md5HashByte[:])
			if md5Hash == "d41d8cd98f00b204e9800998ecf8427e" {
				return v1.ErrorInternalServerError("null hash.")
			}
			if claims.MD5 != md5Hash {
				return v1.ErrorContentMissing("damaged file.")
			}
			// 还原文件光标
			_, err = file.Seek(0, 0)
			if err != nil {
				return v1.ErrorInsufficientStorage("could not seek file.")
			}
		}

		// 计算sha1值
		{
			hash := sha1.New()
			_, err = io.Copy(hash, file)
			if err != nil {
				return v1.ErrorContentMissing("damaged file..")
			}
			sha1HashByte := hash.Sum(nil)
			sha1Hash := hex.EncodeToString(sha1HashByte[:])
			if claims.SHA1 != sha1Hash {
				return v1.ErrorContentMissing("damaged file.")
			}
			// 还原文件光标
			_, err = file.Seek(0, 0)
			if err != nil {
				return v1.ErrorInsufficientStorage("could not seek file.")
			}
		}

		// 建立请求数据类型
		var in service.UploadFileRequest
		in.ID = claims.FileID
		in.MD5 = claims.MD5
		in.SHA1 = claims.SHA1
		in.SliceID = claims.SliceID
		in.SliceLen = claims.SliceLen
		in.ExpiresAt = claims.ExpiresAt

		in.Name = handler.Filename

		// 拼接文件零时储存目录
		in.Path = fmt.Sprintf("%v/%v.%v", tmp, in.ID, in.SliceID)

		// 保存上传文件到零时目录
		in.File, err = os.Create(in.Path)
		if err != nil {
			return v1.ErrorInternalServerError("could not open file.")
		}
		defer in.File.Close()
		_, err = io.Copy(in.File, file)
		if err != nil {
			return v1.ErrorInsufficientStorage("could not save file.")
		}

		// 还原文件光标
		_, err = in.File.Seek(0, 0)
		if err != nil {
			return v1.ErrorInsufficientStorage("could not seek file.")
		}

		kratosHTTP.SetOperation(ctx, "/edge_file.v1.EdgeFile/UploadFile")
		h := ctx.Middleware(
			func(ctx context.Context, req interface{}) (interface{}, error) {
				return efs.UploadFile(ctx, req.(*service.UploadFileRequest))
			},
		)
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*service.UploadFileReply)
		return ctx.Result(200, reply)
	}
}
