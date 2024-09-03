package logic

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"goHyper/core/svc/base"
	"goHyper/internal/controller/admin/rsp"
	"goHyper/internal/dao"
	"goHyper/internal/infra/consts"
	"goHyper/libs/errLib"
	"goHyper/libs/jwtLib"
	"goHyper/libs/md5Lib"
	"goHyper/libs/resLib"
	"time"
)

type Admin struct {
	admin  *dao.Admin
	config *base.Config
}

func NewAdmin(admin *dao.Admin, config *base.Config) *Admin {
	return &Admin{
		admin:  admin,
		config: config,
	}
}

func (l *Admin) Login(ctx *fiber.Ctx, userName, passWord, ip string) (*rsp.SystemLoginRsp, error) {
	sysAdmin, err := l.admin.GetByUserName(userName)
	if err != nil {
		return nil, errLib.AccountNotExist
	}
	// ===============================
	// 前置判断
	// ===============================
	if sysAdmin.IsDisable == 1 {
		return nil, errLib.AccountDisabled
	}
	// 判断密码
	md5Pwd := md5Lib.MakeMd5(passWord + sysAdmin.Salt)
	if md5Pwd != sysAdmin.Password {
		return nil, errLib.PasswordError
	}
	// ****************************************
	// 生成jwt
	// ****************************************
	nowTime := time.Now()
	expireTime := time.Duration(180*24) * time.Hour // 180天
	expiredTime := nowTime.Add(expireTime)
	claims := jwtLib.UserClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  nowTime.Unix(),
			ExpiresAt: expiredTime.Unix(),
		},
		No: sysAdmin.No,
	}
	jwtObj := jwtLib.NewUserJwt(l.config.Admin.JwtSignKey, l.config.Admin.JwtAesKey, &claims)
	jwtStr, err := jwtObj.Encode()
	if err != nil {
		return nil, err
	}
	// 存cookie
	resLib.CookieAdd(ctx, consts.AdminTokenName, jwtStr, consts.MaxAge)
	// 更新登录信息
	err = l.admin.LoginUpdate(ip)
	if err != nil {
		return nil, err
	}
	return &rsp.SystemLoginRsp{
		Token: jwtStr,
	}, nil
}

func (l *Admin) Logout(ctx *fiber.Ctx) {
	resLib.CookieRemove(ctx, consts.AdminTokenName)
}
