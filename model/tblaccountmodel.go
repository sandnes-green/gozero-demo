package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tblAccountFieldNames          = builder.RawFieldNames(&TblAccount{})
	tblAccountRows                = strings.Join(tblAccountFieldNames, ",")
	tblAccountRowsExpectAutoSet   = strings.Join(stringx.Remove(tblAccountFieldNames, "`create_time`", "`update_time`"), ",")
	tblAccountRowsWithPlaceHolder = strings.Join(stringx.Remove(tblAccountFieldNames, "`fld_userid`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cache91cbAccountTblAccountFldUseridPrefix = "cache:91cbAccount:tblAccount:fldUserid:"
)

type (
	TblAccountModel interface {
		Insert(ctx context.Context, data *TblAccount) (sql.Result, error)
		FindOne(ctx context.Context, fldUserid int64) (*TblAccount, error)
		GetUserInfo(ctx context.Context, fldUserid int64) (*UserInfo, error)
		Update(ctx context.Context, data *TblAccount) error
		Delete(ctx context.Context, fldUserid int64) error
	}

	defaultTblAccountModel struct {
		sqlc.CachedConn
		table string
	}

	TblAccount struct {
		FldUserid             int64     `db:"fld_userid"`               // 用户帐号，可用于登录
		FldCbid               string    `db:"fld_cbid"`                 // 唯一，只用于定位用户，查找好友，不附加其他功能，不能用于登录。系统自动生成。
		FldLoginid            string    `db:"fld_loginid"`              // 个性化登录名，可为空
		FldBindMobile         string    `db:"fld_bind_mobile"`          // 绑定手机号
		FldBindEmail          string    `db:"fld_bind_email"`           // 绑定email
		FldIsSchoolTeacher    int64     `db:"fld_is_school_teacher"`    // 是否中小学教师
		FldIsSchoolPrincipal  int64     `db:"fld_is_school_principal"`  // 是否中小学校长
		FldIsSchoolStudent    int64     `db:"fld_is_school_student"`    // 是否中小学生
		FldIsTeachResearcher  int64     `db:"fld_is_teach_researcher"`  // 是否教研员
		FldIsEduManagement    int64     `db:"fld_is_edu_management"`    // 是否教育系统管理人员
		FldIsParents          int64     `db:"fld_is_parents"`           // 是否家长
		FldIsCollegeStudent   int64     `db:"fld_is_college_student"`   // 是否高职高校学生
		FldIsCollegeTeacher   int64     `db:"fld_is_college_teacher"`   // 是否高职高校教师
		FldIsCollegePrincipal int64     `db:"fld_is_college_principal"` // 是否高职高校校长
		FldIsCommercial       int64     `db:"fld_is_commercial"`        // 是否商业机构用户
		FldRegistTime         time.Time `db:"fld_regist_time"`          // 注册时间
		FldStatus             int64     `db:"fld_status"`               // 帐号状态   0=冻结 1=可用
	}

	UserInfo struct {
		FldUserid     int64     `db:"fld_userid"`      // 用户帐号，可用于登录
		FldBindMobile string    `db:"fld_bind_mobile"` // 绑定手机号
		FldBindEmail  string    `db:"fld_bind_email"`  // 绑定email
		FldRegistTime time.Time `db:"fld_regist_time"` // 注册时间
		FldStatus     int64     `db:"fld_status"`      // 帐号状态   0=冻结 1=可用
		FldName       string    `db:"fld_name"`        //用户名
		FldSex        int8      `db:"fld_sex"`         //性别
		FldCity       int32     `db:"fld_city"`        //所属城市
		FldSchoolId   int32     `db:"fld_schoolid"`    //学校id
		FldMood       string    `db:"fld_mood"`        //用户签名
	}
)

func NewTblAccountModel(conn sqlx.SqlConn, c cache.CacheConf) TblAccountModel {
	return &defaultTblAccountModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`tbl_account`",
	}
}

func (m *defaultTblAccountModel) GetUserInfo(ctx context.Context, fldUserid int64) (*UserInfo, error) {
	_91cbAccountTblAccountFldUseridKey := fmt.Sprintf("%s%v", cache91cbAccountTblAccountFldUseridPrefix, fldUserid)
	logx.Info("fldUserid===", fldUserid)
	var resp UserInfo
	err := m.QueryRowCtx(ctx, &resp, _91cbAccountTblAccountFldUseridKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select d.fld_userid as fld_userid,fld_bind_mobile,fld_bind_email,fld_regist_time,fld_status,fld_name,fld_sex,fld_city,fld_schoolid,fld_mood " +
			"from tbl_account as d left join tbl_user_info as b on d.fld_userid = b.fld_userid where d.`fld_userid` = ? limit 1")
		logx.Info("query===", query)
		return conn.QueryRowCtx(ctx, v, query, fldUserid)
	})

	logx.Info("resp===", resp)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTblAccountModel) Insert(ctx context.Context, data *TblAccount) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, tblAccountRowsExpectAutoSet)
	ret, err := m.ExecNoCacheCtx(ctx, query, data.FldUserid, data.FldCbid, data.FldLoginid, data.FldBindMobile, data.FldBindEmail, data.FldIsSchoolTeacher, data.FldIsSchoolPrincipal, data.FldIsSchoolStudent, data.FldIsTeachResearcher, data.FldIsEduManagement, data.FldIsParents, data.FldIsCollegeStudent, data.FldIsCollegeTeacher, data.FldIsCollegePrincipal, data.FldIsCommercial, data.FldRegistTime, data.FldStatus)

	return ret, err
}

func (m *defaultTblAccountModel) FindOne(ctx context.Context, fldUserid int64) (*TblAccount, error) {
	_91cbAccountTblAccountFldUseridKey := fmt.Sprintf("%s%v", cache91cbAccountTblAccountFldUseridPrefix, fldUserid)
	var resp TblAccount
	err := m.QueryRowCtx(ctx, &resp, _91cbAccountTblAccountFldUseridKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `fld_userid` = ? limit 1", tblAccountRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, fldUserid)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTblAccountModel) Update(ctx context.Context, data *TblAccount) error {
	_91cbAccountTblAccountFldUseridKey := fmt.Sprintf("%s%v", cache91cbAccountTblAccountFldUseridPrefix, data.FldUserid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `fld_userid` = ?", m.table, tblAccountRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.FldCbid, data.FldLoginid, data.FldBindMobile, data.FldBindEmail, data.FldIsSchoolTeacher, data.FldIsSchoolPrincipal, data.FldIsSchoolStudent, data.FldIsTeachResearcher, data.FldIsEduManagement, data.FldIsParents, data.FldIsCollegeStudent, data.FldIsCollegeTeacher, data.FldIsCollegePrincipal, data.FldIsCommercial, data.FldRegistTime, data.FldStatus, data.FldUserid)
	}, _91cbAccountTblAccountFldUseridKey)
	return err
}

func (m *defaultTblAccountModel) Delete(ctx context.Context, fldUserid int64) error {
	_91cbAccountTblAccountFldUseridKey := fmt.Sprintf("%s%v", cache91cbAccountTblAccountFldUseridPrefix, fldUserid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `fld_userid` = ?", m.table)
		return conn.ExecCtx(ctx, query, fldUserid)
	}, _91cbAccountTblAccountFldUseridKey)
	return err
}

func (m *defaultTblAccountModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cache91cbAccountTblAccountFldUseridPrefix, primary)
}

func (m *defaultTblAccountModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `fld_userid` = ? limit 1", tblAccountRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}
