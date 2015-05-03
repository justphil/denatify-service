// Generated by ego on Sun May  3 17:19:45 2015.
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
_, _ = fmt.Fprint(w, "\n\n<!DOCTYPE html>\n<html>\n<head>\n  <meta charset=\"utf-8\">\n  <title>denatify</title>\n  <link href=\"//maxcdn.bootstrapcdn.com/bootswatch/3.3.4/cosmo/bootstrap.min.css\" rel=\"stylesheet\">\n</head>\n<body>\n\n  <div class=\"container\">\n    <div class=\"row\">\n      <div class=\"col-xs-12\">\n        denatify # <em>rendered on ")
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
_, _ = fmt.Fprint(w, "\n\n<div class=\"container\">\n  <div class=\"row\">\n    \n    ")
//line templates/users/delete.ego:8
 if err != nil { 
//line templates/users/delete.ego:9
_, _ = fmt.Fprint(w, "\n      <div class=\"col-md-6 col-md-offset-3\">\n        <div class=\"alert alert-danger\" role=\"alert\">")
//line templates/users/delete.ego:10
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  err.Error() )))
//line templates/users/delete.ego:10
_, _ = fmt.Fprint(w, "</div>\n      </div>\n    ")
//line templates/users/delete.ego:12
 } else { 
//line templates/users/delete.ego:13
_, _ = fmt.Fprint(w, "\n      <div class=\"col-md-6 col-md-offset-3\">\n        Do you really want to delete the user\n        <em class=\"text-uppercase\">")
//line templates/users/delete.ego:15
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  u.Fullname )))
//line templates/users/delete.ego:15
_, _ = fmt.Fprint(w, "</em>\n        (<span class=\"text-muted\">ID: ")
//line templates/users/delete.ego:16
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  u.Id.Hex() )))
//line templates/users/delete.ego:16
_, _ = fmt.Fprint(w, "</span>)?\n      </div>\n      \n      <div class=\"col-md-3 col-md-offset-3\">\n        <a class=\"btn btn-primary btn-block\" href=\"./perform-deletion\" role=\"button\">Yes</a>\n      </div>\n      \n      <div class=\"col-md-3\">\n        <a class=\"btn btn-warning btn-block\" href=\"..\" role=\"button\">No</a>\n      </div>\n    ")
//line templates/users/delete.ego:26
 } 
//line templates/users/delete.ego:27
_, _ = fmt.Fprint(w, "\n\n  </div>\n</div>")
return nil
}
//line templates/users/user.ego:1
 func User(w io.Writer, u *models.User, err error) error  {
//line templates/users/user.ego:2
_, _ = fmt.Fprint(w, "\n\n")
//line templates/users/user.ego:4
_, _ = fmt.Fprint(w, "\n\n<div class=\"container\">\n  <div class=\"row\">\n\n    ")
//line templates/users/user.ego:8
 if err != nil { 
//line templates/users/user.ego:9
_, _ = fmt.Fprint(w, "\n      <div class=\"col-md-6 col-md-offset-3\">\n        <div class=\"alert alert-danger\" role=\"alert\">")
//line templates/users/user.ego:10
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  err.Error() )))
//line templates/users/user.ego:10
_, _ = fmt.Fprint(w, "</div>\n      </div>\n    ")
//line templates/users/user.ego:12
 } else { 
//line templates/users/user.ego:13
_, _ = fmt.Fprint(w, "\n      <div class=\"col-md-6 col-md-offset-3\">\n        <h2>")
//line templates/users/user.ego:14
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  u.Fullname )))
//line templates/users/user.ego:14
_, _ = fmt.Fprint(w, "</h2>\n        <ul>\n            <li>ID: ")
//line templates/users/user.ego:16
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  u.Id.Hex() )))
//line templates/users/user.ego:16
_, _ = fmt.Fprint(w, "</li>\n            <li>Email: ")
//line templates/users/user.ego:17
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  u.Email )))
//line templates/users/user.ego:17
_, _ = fmt.Fprint(w, "</li>\n          </ul>\n        </div>\n\n        <div class=\"col-md-6 col-md-offset-3\">\n          <a class=\"btn btn-default btn-block\" href=\"")
//line templates/users/user.ego:22
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  u.Id.Hex() )))
//line templates/users/user.ego:22
_, _ = fmt.Fprint(w, "/update\" role=\"button\">Update User Info...</a>\n        </div>\n    ")
//line templates/users/user.ego:24
 } 
//line templates/users/user.ego:25
_, _ = fmt.Fprint(w, "\n\n  </div>\n</div>")
return nil
}
//line templates/users/user_form.ego:1
 func UserForm(w io.Writer, action string, err error, d map[string]string) error  {
//line templates/users/user_form.ego:2
_, _ = fmt.Fprint(w, "\n\n<div class=\"container\">\n  <div class=\"row\">\n    <div class=\"col-md-6 col-md-offset-3\">\n      ")
//line templates/users/user_form.ego:6
 if action == "update" { 
//line templates/users/user_form.ego:7
_, _ = fmt.Fprint(w, "\n        <h1>Update User Info</h1>\n      ")
//line templates/users/user_form.ego:8
 } else { 
//line templates/users/user_form.ego:9
_, _ = fmt.Fprint(w, "\n        <h1>New User</h1>\n      ")
//line templates/users/user_form.ego:10
 } 
//line templates/users/user_form.ego:11
_, _ = fmt.Fprint(w, "\n    </div>\n\n    <div class=\"col-md-6 col-md-offset-3\">\n      ")
//line templates/users/user_form.ego:14
 if err != nil { 
//line templates/users/user_form.ego:15
_, _ = fmt.Fprint(w, "\n        <div class=\"alert alert-danger\" role=\"alert\">")
//line templates/users/user_form.ego:15
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  err.Error() )))
//line templates/users/user_form.ego:15
_, _ = fmt.Fprint(w, "</div>\n      ")
//line templates/users/user_form.ego:16
 } 
//line templates/users/user_form.ego:17
_, _ = fmt.Fprint(w, "\n\n      ")
//line templates/users/user_form.ego:18
 if action == "update" { 
//line templates/users/user_form.ego:19
_, _ = fmt.Fprint(w, "\n        <form action=\"./perform-update\" method=\"POST\">\n      ")
//line templates/users/user_form.ego:20
 } else { 
//line templates/users/user_form.ego:21
_, _ = fmt.Fprint(w, "\n        <form action=\"./new\" method=\"POST\">\n      ")
//line templates/users/user_form.ego:22
 } 
//line templates/users/user_form.ego:23
_, _ = fmt.Fprint(w, "\n        <div class=\"form-group\">\n        <label for=\"fullname\">Name</label>\n        <input type=\"text\" class=\"form-control\"\n             id=\"fullname\" name=\"fullname\" placeholder=\"Please enter your name...\"\n             value=\"")
//line templates/users/user_form.ego:27
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  d["fullname"] )))
//line templates/users/user_form.ego:27
_, _ = fmt.Fprint(w, "\">\n      </div>\n      <div class=\"form-group\">\n        <label for=\"email\">Email</label>\n        <input type=\"email\" class=\"form-control\"\n             id=\"email\" name=\"email\" placeholder=\"Please enter your email...\"\n             value=\"")
//line templates/users/user_form.ego:33
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  d["email"] )))
//line templates/users/user_form.ego:33
_, _ = fmt.Fprint(w, "\">\n      </div>\n      <div class=\"form-group\">\n        <label for=\"password\">Password</label>\n        <input type=\"password\" class=\"form-control\"\n             id=\"password\" name=\"password\" placeholder=\"Please enter a password...\"\n             value=\"")
//line templates/users/user_form.ego:39
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  d["password"] )))
//line templates/users/user_form.ego:39
_, _ = fmt.Fprint(w, "\">\n      </div>\n\n      ")
//line templates/users/user_form.ego:42
 if action == "update" { 
//line templates/users/user_form.ego:43
_, _ = fmt.Fprint(w, "\n        <button type=\"submit\" class=\"btn btn-default btn-block\">Update User Info</button>\n      ")
//line templates/users/user_form.ego:44
 } else { 
//line templates/users/user_form.ego:45
_, _ = fmt.Fprint(w, "\n        <button type=\"submit\" class=\"btn btn-default btn-block\">Create User</button>\n      ")
//line templates/users/user_form.ego:46
 } 
//line templates/users/user_form.ego:47
_, _ = fmt.Fprint(w, "\n\n    </form>\n    </div>\n  </div>\n</div>")
return nil
}
//line templates/users/users.ego:1
 func Users(w io.Writer, users []*models.User, err error) error  {
//line templates/users/users.ego:2
_, _ = fmt.Fprint(w, "\n\n")
//line templates/users/users.ego:4
_, _ = fmt.Fprint(w, "\n\n<div class=\"container\">\n  <div class=\"row\">\n    <div class=\"col-md-12\">\n      <table class=\"table table-striped\">\n        <thead>\n          <tr>\n            <th>ID</th>\n            <th>Fullname</th>\n            <th>Email</th>\n            <th>Delete</th>\n          </tr>\n        </thead>\n\n        <tbody>\n          ")
//line templates/users/users.ego:19
 for _, u := range users { 
//line templates/users/users.ego:20
_, _ = fmt.Fprint(w, "\n            <tr>\n              <td><a href=\"users/")
//line templates/users/users.ego:21
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  u.Id.Hex() )))
//line templates/users/users.ego:21
_, _ = fmt.Fprint(w, "\">")
//line templates/users/users.ego:21
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  u.Id.Hex() )))
//line templates/users/users.ego:21
_, _ = fmt.Fprint(w, "</a></td>\n              <td>")
//line templates/users/users.ego:22
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  u.Fullname )))
//line templates/users/users.ego:22
_, _ = fmt.Fprint(w, "</td>\n              <td>")
//line templates/users/users.ego:23
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  u.Email )))
//line templates/users/users.ego:23
_, _ = fmt.Fprint(w, "</td>\n              <td><a href=\"users/")
//line templates/users/users.ego:24
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  u.Id.Hex() )))
//line templates/users/users.ego:24
_, _ = fmt.Fprint(w, "/delete\">delete user</a></td>\n            </tr>\n          ")
//line templates/users/users.ego:26
 } 
//line templates/users/users.ego:27
_, _ = fmt.Fprint(w, "\n        </tbody>\n      </table>\n    </div>\n\n    <div class=\"col-md-12\">\n      <a class=\"btn btn-default pull-right\" href=\"user/new\" role=\"button\">Add User</a>\n    </div>\n  </div>\n</div>")
return nil
}
