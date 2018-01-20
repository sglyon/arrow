package arrow

type BinaryType struct{}

func (t *BinaryType) ID() Type     { return BINARY }
func (t *BinaryType) Name() string { return "binary" }

var (
	BinaryTypes = struct {
		Binary DataType
	}{
		Binary: &BinaryType{},
	}
)
