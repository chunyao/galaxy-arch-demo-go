package utils

import (
	"context"
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"reflect"
	"strconv"
)

var ErrInvalidNum = errors.New("Invalid num")

type RedisUtil struct {
	pool *redis.Pool
}

// 为了以后升级可以统一让所有key失效
func keyPatch(key string) string {
	return key
}
func isNum(i interface{}) (string, bool) {
	switch vi := i.(type) {
	case int8:
		v := int64(vi)
		s := strconv.FormatInt(v, 10)

		return s, true
	case int16:
		v := int64(vi)
		s := strconv.FormatInt(v, 10)

		return s, true
	case int32:
		v := int64(vi)
		s := strconv.FormatInt(v, 10)

		return s, true
	case int:
		v := int64(vi)
		s := strconv.FormatInt(v, 10)

		return s, true
	case int64:
		s := strconv.FormatInt(vi, 10)

		return s, true

	case uint8:
		v := uint64(vi)
		s := strconv.FormatUint(v, 10)

		return s, true
	case uint16:
		v := uint64(vi)
		s := strconv.FormatUint(v, 10)

		return s, true
	case uint32:
		v := uint64(vi)
		s := strconv.FormatUint(v, 10)

		return s, true
	case uint:
		v := uint64(vi)
		s := strconv.FormatUint(v, 10)

		return s, true
	case uint64:
		s := strconv.FormatUint(vi, 10)

		return s, true

	default:
		return "", false
	}
}

func isNumPtr(i interface{}) bool {
	val := reflect.ValueOf(i)
	if val.Kind() != reflect.Ptr {
		return false
	}

	v := val.Elem().Kind()

	switch v {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64:
		return true
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64:
		return true
	default:
		return false
	}
}

// 注：此方法溢出不会报error
func bytesToNum(b []byte, num interface{}) error {
	val := reflect.ValueOf(num)
	if val.Kind() != reflect.Ptr {
		return ErrInvalidNum
	}

	v := val.Elem().Kind()
	s := string(b)

	switch v {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64:
		n, err := strconv.ParseInt(s, 0, 64)
		if err != nil {
			return errors.WithMessagef(err, "%s convert to int err", s)
		}

		val.Elem().SetInt(n)
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64:
		n, err := strconv.ParseUint(s, 0, 64)
		if err != nil {
			return errors.WithMessagef(err, "%s convert to uint err", s)
		}

		val.Elem().SetUint(n)
	}

	return nil
}
func NewRedisUtil(pool *redis.Pool) *RedisUtil {
	return &RedisUtil{pool}
}

func (ru *RedisUtil) Set(key string, value interface{}, ttl int) (err error) {
	var bytesData []byte

	// 判断是否整数
	if s, ok := isNum(value); ok {
		bytesData = []byte(s)
	} else {
		bytesData, err = Serialize(value)

		if err != nil {
			return err
		}
	}

	err = ru.WrapDo(func(con redis.Conn) error {
		_, err = con.Do("SET", keyPatch(key), bytesData, "EX", ttl)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (ru *RedisUtil) Get(key string, value interface{}) (hit bool, err error) {
	if reflect.ValueOf(value).Kind() != reflect.Ptr {
		return false, errors.New("value must be ptr")
	}

	var replay []byte

	err = ru.WrapDo(func(con redis.Conn) error {
		replay, err = redis.Bytes(con.Do("GET", keyPatch(key)))

		if err != nil {
			return err
		}

		return nil
	})

	if err == redis.ErrNil {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	if isNumPtr(value) { // 数字
		err = bytesToNum(replay, value)
	} else {
		err = Decode(replay, value)
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func (ru *RedisUtil) Del(key string) (err error) {
	err = ru.WrapDo(func(con redis.Conn) error {
		_, err = con.Do("DEL", keyPatch(key))
		return err
	})

	return err
}

func (ru *RedisUtil) Expire(ctx context.Context, key string, ttl int) (err error) {
	err = ru.WrapDo(func(con redis.Conn) error {
		_, err = con.Do("EXPIRE", keyPatch(key), ttl)

		return err
	})

	return err
}

func (ru *RedisUtil) TTL(key string) (ttl int, err error) {
	err = ru.WrapDo(func(con redis.Conn) error {
		ttl, err = redis.Int(con.Do("TTL", keyPatch(key)))

		return err
	})

	return ttl, err
}

func (ru *RedisUtil) Incr(key string) (res int64, err error) {
	err = ru.WrapDo(func(con redis.Conn) error {
		res, err = redis.Int64(con.Do("INCR", keyPatch(key)))

		return err
	})

	return res, err
}

func (ru *RedisUtil) IncrBy(key string, diff int64) (res int64, err error) {
	err = ru.WrapDo(func(con redis.Conn) error {
		res, err = redis.Int64(con.Do("INCRBY", keyPatch(key), diff))

		return err
	})

	return res, err
}

func (ru *RedisUtil) Decr(key string) (res int64, err error) {
	err = ru.WrapDo(func(con redis.Conn) error {
		res, err = redis.Int64(con.Do("DECR", keyPatch(key)))

		return err
	})

	return res, err
}

func (ru *RedisUtil) DecrBy(key string, diff int64) (res int64, err error) {
	err = ru.WrapDo(func(con redis.Conn) error {
		res, err = redis.Int64(con.Do("DECRBY", keyPatch(key), diff))
		return err
	})

	return res, err
}

func (ru *RedisUtil) WrapDo(doFunction func(con redis.Conn) error) error {
	con := ru.pool.Get()
	defer con.Close()

	return doFunction(con)
}
