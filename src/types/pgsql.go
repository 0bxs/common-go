package types

import (
	"strconv"
	"strings"
)

type PgI16Arr []int16

func (p *PgI16Arr) FromDB(data []byte) error {
	str := strings.Trim(string(data), "{}")
	if str == "" {
		*p = PgI16Arr{}
		return nil
	}
	parts := strings.Split(str, ",")
	var result PgI16Arr
	for _, p := range parts {
		n, err := strconv.ParseUint(p, 10, 16)
		if err != nil {
			return err
		}
		result = append(result, int16(n))
	}
	*p = result
	return nil
}

func (p PgI16Arr) ToDB() ([]byte, error) {
	var parts []string
	for _, v := range p {
		parts = append(parts, strconv.FormatUint(uint64(v), 10))
	}
	return []byte("{" + strings.Join(parts, ",") + "}"), nil
}

type PgI64Arr []int64

func (p *PgI64Arr) FromDB(data []byte) error {
	str := strings.Trim(string(data), "{}")
	if str == "" {
		*p = PgI64Arr{}
		return nil
	}
	parts := strings.Split(str, ",")
	var result PgI64Arr
	for _, p := range parts {
		n, err := strconv.ParseUint(p, 10, 16)
		if err != nil {
			return err
		}
		result = append(result, int64(n))
	}
	*p = result
	return nil
}

func (p PgI64Arr) ToDB() ([]byte, error) {
	var parts []string
	for _, v := range p {
		parts = append(parts, strconv.FormatUint(uint64(v), 10))
	}
	return []byte("{" + strings.Join(parts, ",") + "}"), nil
}
