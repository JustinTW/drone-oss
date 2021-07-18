--- drone/vendor/github.com/drone/go-scm/scm/driver/gitlab/util.go	2021-07-18 16:05:22.000000000 +0800
+++ drone/vendor/github.com/drone/go-scm/scm/driver/gitlab/util-patched.go	2021-07-18 16:07:47.000000000 +0800
@@ -13,7 +13,9 @@
 )
 
 func encode(s string) string {
-	return strings.Replace(s, "/", "%2F", -1)
+	result := strings.Replace(s, "/", "%2F", -1)
+	result = strings.Replace(result, "--", "%2F", -1)
+	return result
 }
 
 func encodePath(s string) string {
