// GENERATED CODE - DO NOT EDIT
// This file provides a way of creating URL's based on all the actions
// found in all the controllers.
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) GetApplicationListById(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.GetApplicationListById", args).URL
}


type tApplication struct {}
var Application tApplication


func (_ tApplication) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.Index", args).URL
}

func (_ tApplication) Hello(
		username string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	return revel.MainRouter.Reverse("Application.Hello", args).URL
}

func (_ tApplication) Login(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.Login", args).URL
}

func (_ tApplication) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.Logout", args).URL
}

func (_ tApplication) PostLogin(
		email string,
		password string,
		remember bool,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "email", email)
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "remember", remember)
	return revel.MainRouter.Reverse("Application.PostLogin", args).URL
}


type tUser struct {}
var User tUser


func (_ tUser) ListApp(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.ListApp", args).URL
}

func (_ tUser) ViewApp(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.ViewApp", args).URL
}

func (_ tUser) NewApp(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.NewApp", args).URL
}

func (_ tUser) ViewUser(
		id int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("User.ViewUser", args).URL
}

func (_ tUser) Register(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Register", args).URL
}

func (_ tUser) SaveUser(
		utilisateur interface{},
		password interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "utilisateur", utilisateur)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("User.SaveUser", args).URL
}

func (_ tUser) Setting(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Setting", args).URL
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeDir(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeDir", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}

func (_ tStatic) ServeModuleDir(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModuleDir", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


