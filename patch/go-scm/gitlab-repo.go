--- drone/vendor/github.com/drone/go-scm/scm/driver/gitlab/repo.go	2021-07-18 16:05:22.000000000 +0800
+++ drone/vendor/github.com/drone/go-scm/scm/driver/gitlab/repo-patched.go	2021-07-18 16:06:21.000000000 +0800
@@ -191,7 +191,7 @@
 			Admin: canAdmin(from),
 		},
 	}
-	if path := from.Namespace.FullPath; path != "" {
+	if path := strings.Replace(from.Namespace.FullPath, "/", "--", -1); path != "" {
 		to.Namespace = path
 	}
 	if to.Namespace == "" {
