ifeq ($(OS),Windows_NT)
    FILESERVER_BIN := bin\fileserver.exe
    LB_BIN := bin\loadbalancer.exe
    MKDIR_BIN := if not exist bin mkdir bin
    RUN_FS_1 := set PORT=9001&& bin\fileserver.exe
    RUN_FS_2 := set PORT=9002&& bin\fileserver.exe
    RUN_FS_3 := set PORT=9003&& bin\fileserver.exe
    RUN_LB := bin\loadbalancer.exe
else
    FILESERVER_BIN := bin/fileserver
    LB_BIN := bin/loadbalancer
    MKDIR_BIN := mkdir -p bin
    RUN_FS_1 := PORT=9001 ./bin/fileserver
    RUN_FS_2 := PORT=9002 ./bin/fileserver
    RUN_FS_3 := PORT=9003 ./bin/fileserver
    RUN_LB := ./bin/loadbalancer
endif

.PHONY: build run-backend1 run-backend2 run-backend3 run-lb

build:
	$(MKDIR_BIN)
	go build -o $(FILESERVER_BIN) ./fileserver
	go build -o $(LB_BIN) ./loadbalancer

run-backend1: build
	$(RUN_FS_1)

run-backend2: build
	$(RUN_FS_2)

run-backend3: build
	$(RUN_FS_3)

run-lb: build
	$(RUN_LB)
