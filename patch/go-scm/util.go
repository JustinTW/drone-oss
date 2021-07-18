--- drone/vendor/github.com/drone/go-scm/scm/util.go	2021-07-18 16:05:22.000000000 +0800
+++ drone/vendor/github.com/drone/go-scm/scm/util-patched.go	2021-07-18 16:06:52.000000000 +0800
@@ -16,13 +16,10 @@
 
 // Split splits the full repository name into segments.
 func Split(s string) (owner, name string) {
-	parts := strings.SplitN(s, "/", 2)
-	switch len(parts) {
-	case 1:
-		name = parts[0]
-	case 2:
-		owner = parts[0]
-		name = parts[1]
+	parts := strings.Split(s, "/")
+	name = parts[len(parts)-1]
+	if len(parts) > 1 {
+		owner = strings.Join(parts[:len(parts)-1], "--")
 	}
 	return
 }
