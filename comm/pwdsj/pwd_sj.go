package pwdsj

import (
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/shijiu-xf/go-base/constant/pwdconst"
	"github.com/shijiu-xf/go-base/constant/zero"
	"strings"
)

type EncodePwdInfo struct {
	Algorithm string // 算法
	HashFunc  string // hash算法
	encodePwd string // 密文
	salt      string // 盐值
}

type EncodePwdErr struct {
	errMsg string
}

func (e *EncodePwdErr) Error() string {
	return fmt.Sprintf(e.errMsg)
}
func (e *EncodePwdErr) getErrMsg() string {
	return e.getErrMsg()
}
func NewErr(msg string) *EncodePwdErr {
	return &EncodePwdErr{errMsg: msg}
}

// 设置加密配置
func encodeOpt() *password.Options {
	return &password.Options{SaltLen: 16, Iterations: 1000, KeyLen: 32, HashFunction: sha512.New}
}

// NewEncodedPwd 获取一个sha512加密的密码，返回盐值 和 密码密文
func NewEncodedPwd(pwd string) (string, string, error) {
	if pwd == "" {
		return zero.String, zero.String, NewErr("获取密码密文时，入参为空")
	}
	salt, encodedPwd := password.Encode(pwd, encodeOpt())
	return salt, encodedPwd, nil
}

// NewEncodedPwdSep 获取一个sha512加密的密码，盐值 ,算法拼接的字符串，方便存入数据库
func NewEncodedPwdSep(pwd string) (string, error) {
	salt, encodePwd, err := NewEncodedPwd(pwd)
	if err != nil {
		return zero.String, err
	}
	return GetDefaultPwdSep(salt, encodePwd), nil

}

// NewEncodedPwdByOpt 根据自定义的加密配生成密文
func NewEncodedPwdByOpt(code string, opt *password.Options) (string, string, error) {
	if code == zero.String {
		return zero.String, zero.String, NewErr("获取密码密文时，入参为空")
	}
	salt, encodedPwd := password.Encode(code, opt)
	return salt, encodedPwd, nil
}

// VerifyPwd 校验密码和密文是否对应
func VerifyPwd(pwd string, salt string, encodedPwd string) bool {
	return password.Verify(pwd, salt, encodedPwd, encodeOpt())
}

// VerifyPwdSep 校验密码和密文是否对应
func VerifyPwdSep(pwdSep string, pwd string) (bool, error) {
	salt, encodedPwd, err := SplitDefaultPwdSep(pwdSep)
	if err != nil {
		return false, NewErr("解析密码信息字符串时出错，解析后切片长度不为4")
	}
	return password.Verify(pwd, salt, encodedPwd, encodeOpt()), nil
}

// VerifyPwdByOpt 更加自定义的配置校验密码和密文是否对应
func VerifyPwdByOpt(pwd string, salt string, encodedPwd string, opt *password.Options) bool {
	return password.Verify(pwd, salt, encodedPwd, opt)
}

// GetDefaultPwdSep 获取一个默认的算法,盐值,密文,合并的字符串,方便存入数据库
func GetDefaultPwdSep(salt string, encodedPwd string) string {
	return pwdconst.DefaultAlgorithm + pwdconst.Separator + pwdconst.DefaultHashFunc + pwdconst.Separator + salt + pwdconst.Separator + encodedPwd
}

// SplitDefaultPwdSep 分割默认的算法,盐值,密文合并的密文字符串，取出盐值和密文
func SplitDefaultPwdSep(pwdSep string) (string, string, error) {
	split := strings.Split(pwdSep, pwdconst.Separator)
	if len(split) != 4 {
		return zero.String, zero.String, NewErr("解析密码信息字符串时出错，解析后切片长度不为4")
	}
	return split[2], split[3], nil
}
