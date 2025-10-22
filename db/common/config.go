package common

func GetDbDsn(target string) string {
	return "dev:dev@dev.com@tcp(127.0.0.1:3306)/dev?charset=utf8&parseTime=True&loc=Local"
}
