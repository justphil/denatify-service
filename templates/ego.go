// Generated by ego on Fri May  1 17:20:18 2015.
// DO NOT EDIT

package templates
import (
"fmt"
"html"
"io"
"time"
"github.com/justphil/denatify-service/models"
)
//line templates/landing_page/landing_page.ego:1
 func LandingPage(w io.Writer) error  {
//line templates/landing_page/landing_page.ego:2
_, _ = fmt.Fprint(w, "\n\n<h1>Denatify empowers you to easily expose your Web app while being restricted by NAT.</h1>\n<h2>Dead simple web app exposure behind NAT routers.</h2>")
return nil
}
//line templates/layout.ego:1
 func Layout(w io.Writer, renderPartial func()) error  {
//line templates/layout.ego:2
_, _ = fmt.Fprint(w, "\n")
//line templates/layout.ego:3
_, _ = fmt.Fprint(w, "\n\n<!DOCTYPE html>\n<html>\n<head>\n\t<meta charset=\"utf-8\">\n\t<title>denatify</title>\n\t<link href=\"//maxcdn.bootstrapcdn.com/bootswatch/3.3.4/cosmo/bootstrap.min.css\" rel=\"stylesheet\">\n</head>\n<body>\n\n  <div class=\"container\">\n    <div class=\"row\">\n      <div class=\"col-xs-12\">\n        denatify # <em>rendered on ")
//line templates/layout.ego:16
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  time.Now().UTC().Format(time.RFC1123) )))
//line templates/layout.ego:16
_, _ = fmt.Fprint(w, "</em>\n      </div>\n    </div>\n  </div>\n\n  ")
//line templates/layout.ego:21
 renderPartial() 
//line templates/layout.ego:22
_, _ = fmt.Fprint(w, "\n\n</body>\n</html>")
return nil
}
//line templates/users/delete.ego:1
 func DeleteUser(w io.Writer, u *models.User, err error) error  {
//line templates/users/delete.ego:2
_, _ = fmt.Fprint(w, "\n\n")
//line templates/users/delete.ego:4
_, _ = fmt.Fprint(w, "\n\n<div class=\"container\">\n\t<div class=\"row\">\n\t\t\n\t\t")
//line templates/users/delete.ego:8
 if err != nil { 
//line templates/users/delete.ego:9
_, _ = fmt.Fprint(w, "\n\t\t\t<div class=\"col-md-6 col-md-offset-3\">\n\t\t\t\t<div class=\"alert alert-danger\" role=\"alert\">")
//line templates/users/delete.ego:10
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  err.Error() )))
//line templates/users/delete.ego:10
_, _ = fmt.Fprint(w, "</div>\n\t\t\t</div>\n\t\t")
//line templates/users/delete.ego:12
 } else { 
//line templates/users/delete.ego:13
_, _ = fmt.Fprint(w, "\n\t\t\t<div class=\"col-md-6 col-md-offset-3\">\n\t\t\t\tDo you really want to delete the user\n\t\t\t\t<em class=\"text-uppercase\">")
//line templates/users/delete.ego:15
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  u.Fullname )))
//line templates/users/delete.ego:15
_, _ = fmt.Fprint(w, "</em>\n\t\t\t\t(<span class=\"text-muted\">ID: ")
//line templates/users/delete.ego:16
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  u.Id.Hex() )))
//line templates/users/delete.ego:16
_, _ = fmt.Fprint(w, "</span>)?\n\t\t\t</div>\n\t\t\t\n\t\t\t<div class=\"col-md-3 col-md-offset-3\">\n\t\t\t\t<a class=\"btn btn-primary btn-block\" href=\"/users/")
//line templates/users/delete.ego:20
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  u.Id.Hex() )))
//line templates/users/delete.ego:20
_, _ = fmt.Fprint(w, "/perform-deletion\" role=\"button\">Yes</a>\n\t\t\t</div>\n\t\t\t\n\t\t\t<div class=\"col-md-3\">\n\t\t\t\t<a class=\"btn btn-warning btn-block\" href=\"/users\" role=\"button\">No</a>\n\t\t\t</div>\n\t\t")
//line templates/users/delete.ego:26
 } 
//line templates/users/delete.ego:27
_, _ = fmt.Fprint(w, "\n\n\t</div>\n</div>")
return nil
}
//line templates/users/user.ego:1
 func User(w io.Writer, u *models.User, err error) error  {
//line templates/users/user.ego:2
_, _ = fmt.Fprint(w, "\n\n")
//line templates/users/user.ego:4
_, _ = fmt.Fprint(w, "\n\n<div class=\"container\">\n\t<div class=\"row\">\n\n\t\t")
//line templates/users/user.ego:8
 if err != nil { 
//line templates/users/user.ego:9
_, _ = fmt.Fprint(w, "\n\t\t\t<div class=\"col-md-6 col-md-offset-3\">\n\t\t\t\t<div class=\"alert alert-danger\" role=\"alert\">")
//line templates/users/user.ego:10
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  err.Error() )))
//line templates/users/user.ego:10
_, _ = fmt.Fprint(w, "</div>\n\t\t\t</div>\n\t\t")
//line templates/users/user.ego:12
 } else { 
//line templates/users/user.ego:13
_, _ = fmt.Fprint(w, "\n\t\t\t<div class=\"col-md-6 col-md-offset-3\">\n\t\t\t\t<h2>")
//line templates/users/user.ego:14
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  u.Fullname )))
//line templates/users/user.ego:14
_, _ = fmt.Fprint(w, "</h2>\n\t\t\t\t<ul>\n\t  \t\t\t\t<li>ID: ")
//line templates/users/user.ego:16
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  u.Id.Hex() )))
//line templates/users/user.ego:16
_, _ = fmt.Fprint(w, "</li>\n\t\t  \t\t\t<li>Email: ")
//line templates/users/user.ego:17
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  u.Email )))
//line templates/users/user.ego:17
_, _ = fmt.Fprint(w, "</li>\n\t\t  \t\t</ul>\n\t  \t\t</div>\n\n\t  \t\t<div class=\"col-md-6 col-md-offset-3\">\n\t  \t\t\t<a class=\"btn btn-default btn-block\" href=\"/users/")
//line templates/users/user.ego:22
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  u.Id.Hex() )))
//line templates/users/user.ego:22
_, _ = fmt.Fprint(w, "/update\" role=\"button\">Update User Info...</a>\n\t  \t\t</div>\n\t\t")
//line templates/users/user.ego:24
 } 
//line templates/users/user.ego:25
_, _ = fmt.Fprint(w, "\n\n\t</div>\n</div>")
return nil
}
//line templates/users/user_form.ego:1
 func UserForm(w io.Writer, action string, errMsg string, d map[string]string) error  {
//line templates/users/user_form.ego:2
_, _ = fmt.Fprint(w, "\n\n<div class=\"container\">\n\t<div class=\"row\">\n\t  <div class=\"col-md-6 col-md-offset-3\">\n\t  \t")
//line templates/users/user_form.ego:6
 if action == "update" { 
//line templates/users/user_form.ego:7
_, _ = fmt.Fprint(w, "\n\t  \t\t<h1>Update User Info</h1>\n\t  \t")
//line templates/users/user_form.ego:8
 } else { 
//line templates/users/user_form.ego:9
_, _ = fmt.Fprint(w, "\n\t  \t\t<h1>New User</h1>\n\t  \t")
//line templates/users/user_form.ego:10
 } 
//line templates/users/user_form.ego:11
_, _ = fmt.Fprint(w, "\n\t  </div>\n\n\t  <div class=\"col-md-6 col-md-offset-3\">\n\t  \t")
//line templates/users/user_form.ego:14
 if len(errMsg) > 0 { 
//line templates/users/user_form.ego:15
_, _ = fmt.Fprint(w, "\n\t  \t\t<div class=\"alert alert-danger\" role=\"alert\">")
//line templates/users/user_form.ego:15
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  errMsg )))
//line templates/users/user_form.ego:15
_, _ = fmt.Fprint(w, "</div>\n\t  \t")
//line templates/users/user_form.ego:16
 } 
//line templates/users/user_form.ego:17
_, _ = fmt.Fprint(w, "\n\n\t  \t")
//line templates/users/user_form.ego:18
 if action == "update" { 
//line templates/users/user_form.ego:19
_, _ = fmt.Fprint(w, "\n\t  \t\t<form action=\"/users/")
//line templates/users/user_form.ego:19
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  d["userId"] )))
//line templates/users/user_form.ego:19
_, _ = fmt.Fprint(w, "/perform-update\" method=\"POST\">\n\t  \t")
//line templates/users/user_form.ego:20
 } else { 
//line templates/users/user_form.ego:21
_, _ = fmt.Fprint(w, "\n\t  \t\t<form action=\"/user/new\" method=\"POST\">\n\t  \t")
//line templates/users/user_form.ego:22
 } 
//line templates/users/user_form.ego:23
_, _ = fmt.Fprint(w, "\n\t  \t  <div class=\"form-group\">\n\t\t    <label for=\"fullname\">Name</label>\n\t\t    <input type=\"text\" class=\"form-control\"\n\t\t    \t   id=\"fullname\" name=\"fullname\" placeholder=\"Please enter your name...\"\n\t\t    \t   value=\"")
//line templates/users/user_form.ego:27
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  d["fullname"] )))
//line templates/users/user_form.ego:27
_, _ = fmt.Fprint(w, "\">\n\t\t  </div>\n\t\t  <div class=\"form-group\">\n\t\t    <label for=\"email\">Email</label>\n\t\t    <input type=\"email\" class=\"form-control\"\n\t\t    \t   id=\"email\" name=\"email\" placeholder=\"Please enter your email...\"\n\t\t    \t   value=\"")
//line templates/users/user_form.ego:33
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  d["email"] )))
//line templates/users/user_form.ego:33
_, _ = fmt.Fprint(w, "\">\n\t\t  </div>\n\t\t  <div class=\"form-group\">\n\t\t    <label for=\"password\">Password</label>\n\t\t    <input type=\"password\" class=\"form-control\"\n\t\t    \t   id=\"password\" name=\"password\" placeholder=\"Please enter a password...\"\n\t\t    \t   value=\"")
//line templates/users/user_form.ego:39
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  d["password"] )))
//line templates/users/user_form.ego:39
_, _ = fmt.Fprint(w, "\">\n\t\t  </div>\n\n\t\t  ")
//line templates/users/user_form.ego:42
 if action == "update" { 
//line templates/users/user_form.ego:43
_, _ = fmt.Fprint(w, "\n\t\t  \t<button type=\"submit\" class=\"btn btn-default btn-block\">Update User Info</button>\n\t\t  ")
//line templates/users/user_form.ego:44
 } else { 
//line templates/users/user_form.ego:45
_, _ = fmt.Fprint(w, "\n\t\t  \t<button type=\"submit\" class=\"btn btn-default btn-block\">Create User</button>\n\t\t  ")
//line templates/users/user_form.ego:46
 } 
//line templates/users/user_form.ego:47
_, _ = fmt.Fprint(w, "\n\n\t\t</form>\n\t  </div>\n\t</div>\n</div>")
return nil
}
//line templates/users/users.ego:1
 func Users(w io.Writer, users []*models.User, errMsg string) error  {
//line templates/users/users.ego:2
_, _ = fmt.Fprint(w, "\n\n")
//line templates/users/users.ego:4
_, _ = fmt.Fprint(w, "\n\n<div class=\"container\">\n\t<div class=\"row\">\n\t  <div class=\"col-md-12\">\n\t  \t<table class=\"table table-striped\">\n\t  \t\t<thead>\n\t  \t\t\t<tr>\n\t  \t\t\t\t<th>ID</th>\n\t  \t\t\t\t<th>Fullname</th>\n\t  \t\t\t\t<th>Email</th>\n\t  \t\t\t\t<th>Delete</th>\n\t  \t\t\t</tr>\n\t  \t\t</thead>\n\n\t  \t\t<tbody>\n\t  \t\t\t")
//line templates/users/users.ego:19
 for _, u := range users { 
//line templates/users/users.ego:20
_, _ = fmt.Fprint(w, "\n\t  \t\t\t\t<tr>\n\t  \t\t\t\t\t<td><a href=\"/users/")
//line templates/users/users.ego:21
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  u.Id.Hex() )))
//line templates/users/users.ego:21
_, _ = fmt.Fprint(w, "\">")
//line templates/users/users.ego:21
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  u.Id.Hex() )))
//line templates/users/users.ego:21
_, _ = fmt.Fprint(w, "</a></td>\n\t  \t\t\t\t\t<td>")
//line templates/users/users.ego:22
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  u.Fullname )))
//line templates/users/users.ego:22
_, _ = fmt.Fprint(w, "</td>\n\t  \t\t\t\t\t<td>")
//line templates/users/users.ego:23
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  u.Email )))
//line templates/users/users.ego:23
_, _ = fmt.Fprint(w, "</td>\n\t  \t\t\t\t\t<td><a href=\"/users/")
//line templates/users/users.ego:24
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  u.Id.Hex() )))
//line templates/users/users.ego:24
_, _ = fmt.Fprint(w, "/delete\">delete user</a></td>\n\t  \t\t\t\t</tr>\n\t  \t\t\t")
//line templates/users/users.ego:26
 } 
//line templates/users/users.ego:27
_, _ = fmt.Fprint(w, "\n\t  \t\t</tbody>\n\t  \t</table>\n\t  </div>\n\n\t  <div class=\"col-md-12\">\n\t  \t<a class=\"btn btn-default pull-right\" href=\"/user/new\" role=\"button\">Add User</a>\n\t  </div>\n\t</div>\n</div>")
return nil
}
