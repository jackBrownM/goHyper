package logic

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"goHyper/core/consts"
	req_admin "goHyper/internal/controller/admin/req"
	rsp_admin "goHyper/internal/controller/admin/rsp"
	"goHyper/internal/dao"
	"goHyper/internal/ent"
	"goHyper/libs/errLib"
	"goHyper/libs/jwtLib"
	"goHyper/libs/md5Lib"
	"goHyper/libs/resLib"
	"goHyper/libs/utilLib"
	"goHyper/svc/base"
	"strconv"
	"strings"
	"time"
)

type Admin struct {
	menu   *dao.Menu
	perm   *dao.Perm
	admin  *dao.Admin
	role   *dao.Role
	config *base.Config
}

func NewAdmin(admin *dao.Admin, role *dao.Role, config *base.Config) *Admin {
	return &Admin{
		admin:  admin,
		role:   role,
		config: config,
	}
}

func (l *Admin) Login(userName, passWord, ip string) (*rsp_admin.SystemLoginRsp, error) {
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
	claims := jwtLib.AdminClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  nowTime.Unix(),
			ExpiresAt: expiredTime.Unix(),
		},
		Id: sysAdmin.Id,
	}
	jwtObj := jwtLib.NewAdminJwt(l.config.Admin.JwtSignKey, l.config.Admin.JwtAesKey, &claims)
	jwtStr, err := jwtObj.Encode()
	if err != nil {
		return nil, err
	}
	// 更新登录信息
	err = l.admin.LoginUpdate(sysAdmin.Id, ip)
	if err != nil {
		return nil, err
	}
	return &rsp_admin.SystemLoginRsp{
		Token: jwtStr,
	}, nil
}

func (l *Admin) Logout(ctx *fiber.Ctx) {
	resLib.CookieRemove(ctx, consts.AdminTokenName)
}

func (l *Admin) Detail(id int) (rsp *rsp_admin.SystemAuthAdminRsp, err error) {
	systemAdmin, err := l.admin.GetById(id)
	if err != nil {
		return
	}
	if rsp.Dept == "" {
		rsp.Dept = strconv.FormatInt(int64(rsp.DeptId), 10)
	}
	resLib.Copy(rsp, systemAdmin)
	return
}

func (l *Admin) List(page req_admin.PageReq, listReq req_admin.SystemAuthAdminListReq) (*rsp_admin.PageRsp, error) {
	return l.admin.List(page, listReq)
}

func (l *Admin) Self(adminId int) (any, error) {
	// ===============================
	// 数据处理
	// ===============================
	// 管理员信息
	sysAdmin, err := l.admin.GetById(adminId)
	if err != nil {
		return nil, err
	}
	// 角色权限
	var auths []string
	if adminId > 1 {
		roleId, _ := strconv.Atoi(sysAdmin.Role)
		menuIds, err := l.perm.SelectMenuIdsByRoleId(roleId)
		if err != nil {
			return nil, err
		}
		if len(menuIds) > 0 {
			var menus []ent.SystemAuthMenu
			menus, err = l.menu.GetListByRoleId(menuIds)
			if err != nil {
				return nil, err
			}
			if len(menus) > 0 {
				for _, v := range menus {
					auths = append(auths, strings.Trim(v.Perms, " "))
				}
			}
		}
		if len(auths) > 0 {
			auths = append(auths, "")
		}
	} else {
		auths = append(auths, "*")
	}
	var admin rsp_admin.SystemAuthAdminSelfOneRsp
	resLib.Copy(&admin, sysAdmin)
	admin.Dept = strconv.FormatInt(int64(sysAdmin.DeptId), 10)
	return rsp_admin.SystemAuthAdminSelfRsp{User: admin, Permissions: auths}, nil
}

func (l *Admin) Create(addReq req_admin.SystemAuthAdminAddReq) error {
	var sysAdmin ent.SystemAuthAdmin
	// ===============================
	// 前置判断
	// ===============================
	// 判断用户名或称昵是否存在
	isExitAdmin := l.admin.IsExitAdmin(addReq.Username, addReq.Nickname)
	if isExitAdmin {
		return errLib.AccountExist
	}
	// ===============================
	// 整理数据
	// ===============================
	roleRsp, err := l.role.Detail(addReq.Role)
	if err != nil {
		return err
	}
	roleRsp.Member = l.admin.GetMemberCnt(addReq.Role)
	roleRsp.Menus, err = l.perm.SelectMenuIdsByRoleId(addReq.Role)
	if err != nil {
		return err
	}
	if roleRsp.IsDisable > 0 {
		return errLib.AccountDisabled
	}
	salt := utilLib.RandomString(5)
	resLib.Copy(&sysAdmin, addReq)
	sysAdmin.Role = strconv.FormatInt(int64(addReq.Role), 10)
	sysAdmin.Salt = salt
	sysAdmin.Password = utilLib.MakeMd5(strings.Trim(addReq.Password, " ") + salt)
	if addReq.Avatar == "" {
		addReq.Avatar = "/api/static/backend_avatar.png"
	}
	sysAdmin.Avatar = ""
	// ===============================
	// 创建数据
	// ===============================
	err = l.admin.Create(sysAdmin)
	if err != nil {
		return err
	}
	return nil
}

func (l *Admin) Update(editReq req_admin.SystemAuthAdminEditReq) error {
	// ===============================
	// 前置判断
	// ===============================
	// 判断用户名或称昵是否存在
	isExitAdmin := l.admin.IsExitAdmin(editReq.Username, editReq.Nickname)
	if isExitAdmin {
		return errLib.AccountExist
	}
	// 检查role
	if editReq.Role > 0 && editReq.ID != 1 {
		_, err := l.role.Detail(editReq.Role)
		if err != nil {
			return err
		}
	}
	// ===============================
	// 整理数据
	// ===============================
	var sysAdmin ent.SystemAuthAdmin
	resLib.Copy(&sysAdmin, editReq)
	role := editReq.Role
	if editReq.ID == 1 {
		role = 0
	}
	sysAdmin.Role = strconv.FormatUint(uint64(role), 10)
	if editReq.Password != "" {
		salt := utilLib.RandomString(5)
		sysAdmin.Salt = salt
		sysAdmin.Password = utilLib.MakeMd5(strings.Trim(editReq.Password, "") + salt)
	}
	// ===============================
	// 更新数据
	// ===============================
	err := l.admin.Update(sysAdmin)
	if err != nil {
		return err
	}
	return nil
}

func (l *Admin) UpInfo(id int, updateInfo req_admin.SystemAuthAdminUpdateReq) error {
	// ===============================
	// 前置判断
	// ===============================
	// 检查id
	admin, err := l.admin.GetById(id)
	if err != nil {
		return err
	}
	if admin == nil {
		return errLib.AccountNotExist
	}
	// ===============================
	// 数据处理
	// ===============================
	var sysAdmin ent.SystemAuthAdmin
	resLib.Copy(&sysAdmin, updateInfo)
	if updateInfo.Password != "" {
		salt := utilLib.RandomString(5)
		sysAdmin.Salt = salt
		sysAdmin.Password = utilLib.MakeMd5(strings.Trim(updateInfo.Password, "") + salt)
	}
	// ===============================
	// 更新数据
	// ===============================
	err = l.admin.Update(sysAdmin)
	if err != nil {
		return err
	}
	return nil
}

func (l *Admin) Delete(id, adminId int) error {
	// ===============================
	// 前置判断
	// ===============================
	// 检查id
	admin, err := l.admin.GetById(id)
	if err != nil {
		return err
	}
	if admin == nil {
		return errLib.AccountNotExist
	}
	// 系统管理员不能删除
	if id == 1 {
		return errLib.SystemAdminCannotDelete
	}
	// 不能删除自己
	if id == adminId {
		return errLib.CannotDeleteMySelf
	}
	// ===============================
	// 更新数据
	// ===============================
	err = l.admin.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (l *Admin) Disable(id, adminId int) error {
	// ===============================
	// 前置判断
	// ===============================
	// 检查id
	admin, err := l.admin.GetById(id)
	if err != nil {
		return err
	}
	if admin == nil {
		return errLib.AccountNotExist
	}
	// 不能禁用自己
	if id == adminId {
		return errLib.CannotDisableMySelf
	}
	var isDisable int
	if admin.IsDisable == 0 {
		isDisable = 1
	}
	// ===============================
	// 更新数据
	// ===============================
	err = l.admin.Disable(id, isDisable)
	if err != nil {
		return err
	}
	return nil
}
