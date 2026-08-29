package dto

import "strconv"

// Uint64ToString 将 uint64 转为 string，用于 DTO 序列化，防止前端精度丢失
func Uint64ToString(v uint64) string {
	return strconv.FormatUint(v, 10)
}

// StringToUint64 将 string 转为 uint64，用于从 DTO 反序列化回实体
func StringToUint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}