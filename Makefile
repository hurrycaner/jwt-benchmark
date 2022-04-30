
default:
	go test -bench=. -benchtime 1000000x

all: L0 L1 L2 L3 L4 L5 L6 L7 L8 L9

L0:
	go test -cpuprofile cpu.out -memprofile mem.out -bench ^BenchmarkVerify_L0\$
	go tool pprof -svg mem.out > L0_mem.svg
	go tool pprof -svg cpu.out > L0_cpu.svg
L1:
	go test -memprofile mem.out -cpuprofile cpu.out -bench ^BenchmarkVerify_L1\$
	go tool pprof -svg mem.out > L1_mem.svg
	go tool pprof -svg cpu.out > L1_cpu.svg
L2:
	go test -memprofile mem.out -cpuprofile cpu.out -bench ^BenchmarkVerify_L2\$
	go tool pprof -svg mem.out > L2_mem.svg
	go tool pprof -svg cpu.out > L2_cpu.svg
L3:
	go test -memprofile mem.out -cpuprofile cpu.out -bench ^BenchmarkVerify_L3\$
	go tool pprof -svg mem.out > L3_mem.svg
	go tool pprof -svg cpu.out > L3_cpu.svg
L4:
	go test -memprofile mem.out -cpuprofile cpu.out -bench ^BenchmarkVerify_L4\$
	go tool pprof -svg mem.out > L4_mem.svg
	go tool pprof -svg cpu.out > L4_cpu.svg
L5:
	go test -memprofile mem.out -cpuprofile cpu.out -bench ^BenchmarkVerify_L5\$
	go tool pprof -svg mem.out > L5_mem.svg
	go tool pprof -svg cpu.out > L5_cpu.svg
L6:
	go test -memprofile mem.out -cpuprofile cpu.out -bench ^BenchmarkVerify_L6\$
	go tool pprof -svg mem.out > L6_mem.svg
	go tool pprof -svg cpu.out > L6_cpu.svg
L7:
	go test -memprofile mem.out -cpuprofile cpu.out -bench ^BenchmarkVerify_L7\$
	go tool pprof -svg mem.out > L7_mem.svg
	go tool pprof -svg cpu.out > L7_cpu.svg
L8:
	go test -memprofile mem.out -cpuprofile cpu.out -bench ^BenchmarkVerify_L8\$
	go tool pprof -svg mem.out > L8_mem.svg
	go tool pprof -svg cpu.out > L8_cpu.svg
L9:
	go test -memprofile mem.out -cpuprofile cpu.out -bench ^BenchmarkVerify_L9\$
	go tool pprof -svg mem.out > L9_mem.svg
	go tool pprof -svg cpu.out > L9_cpu.svg
