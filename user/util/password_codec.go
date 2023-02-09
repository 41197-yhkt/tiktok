package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPasswd(password string) (string, error) {
	var passwordHash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		return "", err
	}

	return string(passwordHash), nil
}

func ComparePasswd(dbPassword, loginPassword string) (isPassed bool, err error) {
	// 密码验证
	err = bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(loginPassword)) //验证（对比）
	if err != nil {
		fmt.Println("pwd wrong")
		return false, nil
	} else {
		fmt.Println("pwd ok")
		return true, nil
	}
}
