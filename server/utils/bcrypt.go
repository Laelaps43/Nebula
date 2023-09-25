package utils
// 使用BCryp处理密码，自动加盐

import "golang.org/x/crypto/bcrypt"

// 对密码进行加密
func BcryptHash(password string) string{
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}


// 用于将加密的密码和传输过来的密码进行比较，为空匹配成功
func BcryptCheck(password, hash string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}