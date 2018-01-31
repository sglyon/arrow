package math

//go:generate go run ../_tools/tmpl/main.go -i -data=float64.tmpldata type.go.tmpl=float64.go type_test.go.tmpl=float64_test.go
//go:generate go run ../_tools/tmpl/main.go -i -data=float64.tmpldata -d arch=avx2 type_amd64.go.tmpl=float64_avx2_amd64.go
//go:generate go run ../_tools/tmpl/main.go -i -data=float64.tmpldata -d arch=sse4 type_amd64.go.tmpl=float64_sse4_amd64.go

//go:generate go run ../_tools/tmpl/main.go -i -data=int64.tmpldata type.go.tmpl=int64.go type_test.go.tmpl=int64_test.go
//go:generate go run ../_tools/tmpl/main.go -i -data=int64.tmpldata -d arch=avx2 type_amd64.go.tmpl=int64_avx2_amd64.go
//go:generate go run ../_tools/tmpl/main.go -i -data=int64.tmpldata -d arch=sse4 type_amd64.go.tmpl=int64_sse4_amd64.go
