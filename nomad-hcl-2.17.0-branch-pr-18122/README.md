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

```
✦ ❯ make nobug_aswell 
rm -f go.mod
make go.mod 2>/dev/null 1>/dev/null
go run main.go ./nomad_nobug.hcl
<nil>

Wed Nov 15 20:22:19 2023 exit 0 🟢 took 2s
```

```diff
--- nomad.hcl	2023-11-15 19:41:37.354695393 -0500
+++ nomad_nobug.hcl	2023-11-15 20:19:14.026554975 -0500
@@ -1,6 +1,7 @@
 job "distinct" {
   constraint {
-    distinct_hosts = true
+    operator = "distinct_hosts"
+    value    = "true"
   }
 
   group "group" {
```
