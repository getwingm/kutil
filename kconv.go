package kutil

import (
	"strconv"
)

func ToBool(val interface{}) (retVal bool, isOk bool) {
	isOk = true
	switch val.(type) {
	case bool:
		retVal = val.(bool)
	case int8:
		retVal = (val.(int8) != 0)
	case int16:
		retVal = (val.(int16) != 0)
	case int32:
		retVal = (val.(int32) != 0)
	case int64:
		retVal = (val.(int64) != 0)
	case int:
		retVal = (val.(int) != 0)
	case uint8:
		retVal = (val.(uint8) != 0)
	case uint16:
		retVal = (val.(uint16) != 0)
	case uint32:
		retVal = (val.(uint32) != 0)
	case uint64:
		retVal = (val.(uint64) != 0)
	case float32:
		retVal = !(val.(float32) > -0.0001 && val.(float32) < 0.0001)
	case float64:
		retVal = !(val.(float64) > -0.0001 && val.(float64) < 0.0001)
	case string:
		b, err := strconv.ParseBool(val.(string))
		//fmt.Println(b, err)
		retVal = b
		isOk = err == nil
	default:
		isOk = false
	}

	return retVal, isOk
}

func ToInt8(val interface{}) (retVal int8, isOk bool) {
	isOk = true
	switch val.(type) {
	case bool:
		if val.(bool) {
			retVal = 1
		} else {
			retVal = 0
		}
	case int8:
		retVal = int8(val.(int8))
	case int16:
		retVal = int8(val.(int16))
	case int32:
		retVal = int8(val.(int32))
	case int64:
		retVal = int8(val.(int64))
	case int:
		retVal = int8(val.(int))
	case uint8:
		retVal = int8(val.(uint8))
	case uint16:
		retVal = int8(val.(uint16))
	case uint32:
		retVal = int8(val.(uint32))
	case uint64:
		retVal = int8(val.(uint64))
	case float32:
		retVal = int8(val.(float32))
	case float64:
		retVal = int8(val.(float64))
	case string:
		i, err := strconv.ParseInt(val.(string), 10, 0)
		//fmt.Println(b, err)
		retVal = int8(i)
		isOk = err == nil
	default:
		isOk = false
	}

	return retVal, isOk
}

func ToInt16(val interface{}) (retVal int16, isOk bool) {
	isOk = true
	switch val.(type) {
	case bool:
		if val.(bool) {
			retVal = 1
		} else {
			retVal = 0
		}
	case int8:
		retVal = int16(val.(int8))
	case int16:
		retVal = int16(val.(int16))
	case int32:
		retVal = int16(val.(int32))
	case int64:
		retVal = int16(val.(int64))
	case int:
		retVal = int16(val.(int))
	case uint8:
		retVal = int16(val.(uint8))
	case uint16:
		retVal = int16(val.(uint16))
	case uint32:
		retVal = int16(val.(uint32))
	case uint64:
		retVal = int16(val.(uint64))
	case float32:
		retVal = int16(val.(float32))
	case float64:
		retVal = int16(val.(float64))
	case string:
		i, err := strconv.ParseInt(val.(string), 10, 0)
		//fmt.Println(b, err)
		retVal = int16(i)
		isOk = err == nil
	default:
		isOk = false
	}

	return retVal, isOk
}

func ToInt32(val interface{}) (retVal int32, isOk bool) {
	isOk = true
	switch val.(type) {
	case bool:
		if val.(bool) {
			retVal = 1
		} else {
			retVal = 0
		}
	case int8:
		retVal = int32(val.(int8))
	case int16:
		retVal = int32(val.(int16))
	case int32:
		retVal = int32(val.(int32))
	case int64:
		retVal = int32(val.(int64))
	case int:
		retVal = int32(val.(int))
	case uint8:
		retVal = int32(val.(uint8))
	case uint16:
		retVal = int32(val.(uint16))
	case uint32:
		retVal = int32(val.(uint32))
	case uint64:
		retVal = int32(val.(uint64))
	case float32:
		retVal = int32(val.(float32))
	case float64:
		retVal = int32(val.(float64))
	case string:
		i, err := strconv.ParseInt(val.(string), 10, 0)
		//fmt.Println(b, err)
		retVal = int32(i)
		isOk = err == nil
	default:
		isOk = false
	}

	return retVal, isOk
}

func ToInt64(val interface{}) (retVal int64, isOk bool) {
	isOk = true
	switch val.(type) {
	case bool:
		if val.(bool) {
			retVal = 1
		} else {
			retVal = 0
		}
	case int8:
		retVal = int64(val.(int8))
	case int16:
		retVal = int64(val.(int16))
	case int32:
		retVal = int64(val.(int32))
	case int64:
		retVal = int64(val.(int64))
	case int:
		retVal = int64(val.(int))
	case uint8:
		retVal = int64(val.(uint8))
	case uint16:
		retVal = int64(val.(uint16))
	case uint32:
		retVal = int64(val.(uint32))
	case uint64:
		retVal = int64(val.(uint64))
	case float32:
		retVal = int64(val.(float32))
	case float64:
		retVal = int64(val.(float64))
	case string:
		i, err := strconv.ParseInt(val.(string), 10, 0)
		//fmt.Println(b, err)
		retVal = int64(i)
		isOk = err == nil
	default:
		isOk = false
	}

	return retVal, isOk
}

func ToInt(val interface{}) (retVal int, isOk bool) {
	isOk = true
	switch val.(type) {
	case bool:
		if val.(bool) {
			retVal = 1
		} else {
			retVal = 0
		}
	case int8:
		retVal = int(val.(int8))
	case int16:
		retVal = int(val.(int16))
	case int32:
		retVal = int(val.(int32))
	case int64:
		retVal = int(val.(int64))
	case int:
		retVal = int(val.(int))
	case uint8:
		retVal = int(val.(uint8))
	case uint16:
		retVal = int(val.(uint16))
	case uint32:
		retVal = int(val.(uint32))
	case uint64:
		retVal = int(val.(uint64))
	case float32:
		retVal = int(val.(float32))
	case float64:
		retVal = int(val.(float64))
	case string:
		i, err := strconv.ParseInt(val.(string), 10, 0)
		//fmt.Println(b, err)
		retVal = int(i)
		isOk = err == nil
	default:
		isOk = false
	}

	return retVal, isOk
}

func ToUInt8(val interface{}) (retVal uint8, isOk bool) {
	isOk = true
	switch val.(type) {
	case bool:
		if val.(bool) {
			retVal = 1
		} else {
			retVal = 0
		}
	case int8:
		retVal = uint8(val.(int8))
	case int16:
		retVal = uint8(val.(int16))
	case int32:
		retVal = uint8(val.(int32))
	case int64:
		retVal = uint8(val.(int64))
	case int:
		retVal = uint8(val.(int))
	case uint8:
		retVal = uint8(val.(uint8))
	case uint16:
		retVal = uint8(val.(uint16))
	case uint32:
		retVal = uint8(val.(uint32))
	case uint64:
		retVal = uint8(val.(uint64))
	case float32:
		retVal = uint8(val.(float32))
	case float64:
		retVal = uint8(val.(float64))
	case string:
		i, err := strconv.ParseInt(val.(string), 10, 0)
		//fmt.Println(b, err)
		retVal = uint8(i)
		isOk = err == nil
	default:
		isOk = false
	}

	return retVal, isOk
}

func ToUInt16(val interface{}) (retVal uint16, isOk bool) {
	isOk = true
	switch val.(type) {
	case bool:
		if val.(bool) {
			retVal = 1
		} else {
			retVal = 0
		}
	case int8:
		retVal = uint16(val.(int8))
	case int16:
		retVal = uint16(val.(int16))
	case int32:
		retVal = uint16(val.(int32))
	case int64:
		retVal = uint16(val.(int64))
	case int:
		retVal = uint16(val.(int))
	case uint8:
		retVal = uint16(val.(uint8))
	case uint16:
		retVal = uint16(val.(uint16))
	case uint32:
		retVal = uint16(val.(uint32))
	case uint64:
		retVal = uint16(val.(uint64))
	case float32:
		retVal = uint16(val.(float32))
	case float64:
		retVal = uint16(val.(float64))
	case string:
		i, err := strconv.ParseInt(val.(string), 10, 0)
		//fmt.Println(b, err)
		retVal = uint16(i)
		isOk = err == nil
	default:
		isOk = false
	}

	return retVal, isOk
}

func ToUInt32(val interface{}) (retVal uint32, isOk bool) {
	isOk = true
	switch val.(type) {
	case bool:
		if val.(bool) {
			retVal = 1
		} else {
			retVal = 0
		}
	case int8:
		retVal = uint32(val.(int8))
	case int16:
		retVal = uint32(val.(int16))
	case int32:
		retVal = uint32(val.(int32))
	case int64:
		retVal = uint32(val.(int64))
	case int:
		retVal = uint32(val.(int))
	case uint8:
		retVal = uint32(val.(uint8))
	case uint16:
		retVal = uint32(val.(uint16))
	case uint32:
		retVal = uint32(val.(uint32))
	case uint64:
		retVal = uint32(val.(uint64))
	case float32:
		retVal = uint32(val.(float32))
	case float64:
		retVal = uint32(val.(float64))
	case string:
		i, err := strconv.ParseInt(val.(string), 10, 0)
		//fmt.Println(b, err)
		retVal = uint32(i)
		isOk = err == nil
	default:
		isOk = false
	}

	return retVal, isOk
}

func ToUInt64(val interface{}) (retVal uint64, isOk bool) {
	isOk = true
	switch val.(type) {
	case bool:
		if val.(bool) {
			retVal = 1
		} else {
			retVal = 0
		}
	case int8:
		retVal = uint64(val.(int8))
	case int16:
		retVal = uint64(val.(int16))
	case int32:
		retVal = uint64(val.(int32))
	case int64:
		retVal = uint64(val.(int64))
	case int:
		retVal = uint64(val.(int))
	case uint8:
		retVal = uint64(val.(uint8))
	case uint16:
		retVal = uint64(val.(uint16))
	case uint32:
		retVal = uint64(val.(uint32))
	case uint64:
		retVal = uint64(val.(uint64))
	case float32:
		retVal = uint64(val.(float32))
	case float64:
		retVal = uint64(val.(float64))
	case string:
		i, err := strconv.ParseInt(val.(string), 10, 0)
		//fmt.Println(b, err)
		retVal = uint64(i)
		isOk = err == nil
	default:
		isOk = false
	}

	return retVal, isOk
}

func ToFloat32(val interface{}) (retVal float32, isOk bool) {
	isOk = true
	switch val.(type) {
	case bool:
		if val.(bool) {
			retVal = 1.0
		} else {
			retVal = 0.0
		}
	case int8:
		retVal = float32(val.(int8))
	case int16:
		retVal = float32(val.(int16))
	case int32:
		retVal = float32(val.(int32))
	case int64:
		retVal = float32(val.(int64))
	case int:
		retVal = float32(val.(int))
	case uint8:
		retVal = float32(val.(uint8))
	case uint16:
		retVal = float32(val.(uint16))
	case uint32:
		retVal = float32(val.(uint32))
	case uint64:
		retVal = float32(val.(uint64))
	case float32:
		retVal = float32(val.(float32))
	case float64:
		retVal = float32(val.(float64))
	case string:
		i, err := strconv.ParseFloat(val.(string), 32)
		//fmt.Println(b, err)
		retVal = float32(i)
		isOk = err == nil
	default:
		isOk = false
	}

	return retVal, isOk
}

func ToFloat64(val interface{}) (retVal float64, isOk bool) {
	isOk = true
	switch val.(type) {
	case bool:
		if val.(bool) {
			retVal = 1.0
		} else {
			retVal = 0.0
		}
	case int8:
		retVal = float64(val.(int8))
	case int16:
		retVal = float64(val.(int16))
	case int32:
		retVal = float64(val.(int32))
	case int64:
		retVal = float64(val.(int64))
	case int:
		retVal = float64(val.(int))
	case uint8:
		retVal = float64(val.(uint8))
	case uint16:
		retVal = float64(val.(uint16))
	case uint32:
		retVal = float64(val.(uint32))
	case uint64:
		retVal = float64(val.(uint64))
	case float32:
		retVal = float64(val.(float32))
	case float64:
		retVal = float64(val.(float64))
	case string:
		i, err := strconv.ParseFloat(val.(string), 64)
		//fmt.Println(b, err)
		retVal = float64(i)
		isOk = err == nil
	default:
		isOk = false
	}

	return retVal, isOk
}
