# Example

```
✦2 ❯ make bug
rm -f go.mod
make go.mod 2>/dev/null 1>/dev/null
go run main.go
nomad.hcl:3,5-19: Unsupported argument; An argument named "distinct_hosts" is not expected here.

Wed Nov 15 19:54:53 2023 exit 0 🟢 took 2s
```

```
✦2 ❯ make nobug
sed -i '/replace/d' go.mod
go mod tidy 2>/dev/null 1>/dev/null
go run main.go
<nil>

Wed Nov 15 19:55:56 2023 exit 0 🟢
```
