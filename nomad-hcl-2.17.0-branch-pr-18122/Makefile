.PHONY: nobug bug

bug:
	rm -f go.mod
	$(MAKE) go.mod 2>/dev/null 1>/dev/null
	go run main.go nomad.hcl

nobug:
	sed -i '/replace/d' go.mod
	go mod tidy 2>/dev/null 1>/dev/null
	go run main.go nomad.hcl

nobug_aswell:
	rm -f go.mod
	$(MAKE) go.mod 2>/dev/null 1>/dev/null
	go run main.go ./nomad_nobug.hcl

go.mod:
	go mod init bugreport
	go get github.com/hashicorp/nomad@v1.6.3
	go get github.com/hashicorp/nomad/api@v0.0.0-20231115150718-f66eb83fc035
	echo 'replace github.com/hashicorp/hcl/v2 v2.9.2-0.20220525143345-ab3cae0737bc => github.com/hashicorp/hcl/v2 v2.17.1-0.20230725002108-58caf00be5aa' >> go.mod
	go mod tidy
