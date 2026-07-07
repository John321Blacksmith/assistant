go test -bench=. brain_uow_bench_test.go brain.go data_management.go data_structures.go entities.go -benchtime=1x
go test -bench=. brain_unit_bench_test.go brain.go data_structures.go entities.go -benchtime=1x
