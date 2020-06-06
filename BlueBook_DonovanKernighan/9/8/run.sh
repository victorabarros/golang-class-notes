rm -rf *.log
GOMAXPROCS=1 go run main.go > logs_with_1_max_procs.log
GOMAXPROCS=2 go run main.go > logs_with_2_max_procs.log
