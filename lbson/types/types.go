package types

const (
	BJTINT    = 1
	BJTUINT   = 2
	BJTFLOAT  = 3
	BJTSTRING = 4
	BJTBOOL   = 5
	BJTMAP    = 6
	BJTSLICE  = 7
)

func CodeToString(code int) string {
	switch code {
	case BJTINT:
		return "int"
	case BJTUINT:
		return "uint"
	case BJTFLOAT:
		return "float"
	case BJTSTRING:
		return "string"
	case BJTBOOL:
		return "bool"
	case BJTMAP:
		return "map"
	case BJTSLICE:
		return "slice"
	}
	return ""
}
