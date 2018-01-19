package arrow

type BooleanType struct{}

func (t *BooleanType) ID() Type     { return BOOL }
func (t *BooleanType) Name() string { return "bool" }

// FixedWidth
func (t *BooleanType) BitWidth() int { return 1 }

var (
	FixedWidthTypes = struct {
		Boolean FixedWidthDataType
	}{
		Boolean: &BooleanType{},
	}
)
