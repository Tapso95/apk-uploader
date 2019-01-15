// GENERATED CODE - DO NOT EDIT
// This file is the run file for Revel.
// It registers all the controllers and provides details for the Revel server engine to
// properly inject parameters directly into the action endpoints.
package run

import (
	"reflect"
	"github.com/revel/revel"
	_ "apk-uploader/app"
	controllers "apk-uploader/app/controllers"
	models "apk-uploader/app/models"
	tests "apk-uploader/tests"
	controllers0 "github.com/revel/modules/static/app/controllers"
	_ "github.com/revel/modules/testrunner/app"
	controllers1 "github.com/revel/modules/testrunner/app/controllers"
	"github.com/revel/revel/testing"
)

var (
	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

// Register and run the application
func Run(port int) {
	Register()
	revel.Run(port)
}

// Register all the controllers
func Register() {
	revel.AppLog.Info("Running revel server")
	
	revel.RegisterController((*controllers.App)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "ListApp",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					42: []string{ 
						"user",
						"id",
						"applications",
					},
					44: []string{ 
						"user",
					},
				},
			},
			&revel.MethodType{
				Name: "ViewApp",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					79: []string{ 
						"app",
						"idApp",
					},
				},
			},
			&revel.MethodType{
				Name: "NewApp",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					99: []string{ 
						"user",
						"id",
						"categories",
					},
				},
			},
			&revel.MethodType{
				Name: "SaveApp",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "application", Type: reflect.TypeOf((**models.Application)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetApplicationListById",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					250: []string{ 
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Application)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					20: []string{ 
						"user",
					},
				},
			},
			&revel.MethodType{
				Name: "Hello",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					32: []string{ 
						"username",
					},
				},
			},
			&revel.MethodType{
				Name: "Login",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					48: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Logout",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "PostLogin",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "email", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "remember", Type: reflect.TypeOf((*bool)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "LoadTypeApp",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.User)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "ViewUser",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					36: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Register",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					40: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "SaveUser",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "utilisateur", Type: reflect.TypeOf((**models.Utilisateur)(nil)) },
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*models.Password)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Setting",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					92: []string{ 
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeDir",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModuleDir",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					76: []string{ 
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Suite",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					125: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
		"apk-uploader/app/controllers.Application.Hello": { 
			24: "username",
		},
		"apk-uploader/app/controllers.Application.PostLogin": { 
			73: "email",
			74: "password",
		},
		"apk-uploader/app/controllers.User.SaveUser": { 
			45: "utilisateur.NomUtilisateur",
			46: "utilisateur.PasswordUtilisateur",
			47: "utilisateur.EmailUtilisateur",
			55: "utilisateur.EmailUtilisateur",
		},
	}
	testing.TestSuites = []interface{}{ 
		(*tests.AppTest)(nil),
	}
}
