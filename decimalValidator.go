package arg

import (
	"fmt"

	"github.com/Bofry/arg/internal"
	"github.com/shopspring/decimal"
)

var (
	_ DecimalPtrValidator = new(DecimalValidator).AssertPtr
	_ ValueValidator      = new(DecimalValidator).AssertValue
)

func (fn DecimalValidator) AssertPtr(v *Decimal, name string) error {
	if v != nil {
		return fn(*v, name)
	}
	return nil
}

func (fn DecimalValidator) AssertValue(v interface{}, name string) error {

	switch real := v.(type) {
	case *Decimal:
		if real != nil {
			return fn(*real, name)
		}
		return nil
	case Decimal:
		return fn(real, name)
	case *float64:
		if real != nil {
			d := decimal.NewFromFloat(*real)
			return fn(d, name)
		}
		return nil
	case float64:
		d := decimal.NewFromFloat(real)
		return fn(d, name)
	case *float32:
		if real != nil {
			d := decimal.NewFromFloat32(*real)
			return fn(d, name)
		}
		return nil
	case float32:
		d := decimal.NewFromFloat32(real)
		return fn(d, name)
	case *int:
		if real != nil {
			d := decimal.NewFromInt(int64(*real))
			return fn(d, name)
		}
		return nil
	case int:
		d := decimal.NewFromInt(int64(real))
		return fn(d, name)
	case *int64:
		if real != nil {
			d := decimal.NewFromInt(*real)
			return fn(d, name)
		}
		return nil
	case int64:
		d := decimal.NewFromInt(real)
		return fn(d, name)
	case *int32:
		if real != nil {
			d := decimal.NewFromInt32(*real)
			return fn(d, name)
		}
		return nil
	case int32:
		d := decimal.NewFromInt32(real)
		return fn(d, name)
	case *uint64:
		if real != nil {
			d := decimal.NewFromUint64(*real)
			return fn(d, name)
		}
		return nil
	case uint64:
		d := decimal.NewFromUint64(real)
		return fn(d, name)
	default:
		var str string
		switch p := v.(type) {
		case *string:
			str = *p
		case string:
			str = p
		case fmt.Stringer:
			str = p.String()
		}

		d, err := decimal.NewFromString(str)
		if err != nil {
			return &InvalidArgumentError{
				Name:   name,
				Reason: fmt.Sprintf(internal.ERR_INVALID_NUMBER, v),
				Err:    err,
			}
		}
		return fn(d, name)
	}
}
