.PHONY: c s crun srun
c: cbuild crun
s: sbuild srun
cbuild:
	mkdir -p build && cd build && go build ../udpclient.go
sbuild:
	mkdir -p build && cd build && go build ../udpserver.go
crun:
	build/udpclient -k "127.0.0.1:5001"
srun:
	build/udpserver --host 127.0.0.1 --port 5001
