package controllers

import (
	"github.com/mindoktor/mdbugs/app/models"
	"github.com/mindoktor/mdbugs/app/routes"
	"github.com/revel/revel"
)

type App struct {
	GorpController
	PageTitle string
}

func (c App) PageLoad() revel.Result {		
	c.PageTitle = "Correct page title"
	c.RenderArgs["pageTitle"] = c.PageTitle	
	return nil
}

func (c App) Index() revel.Result {	
	//c.RenderArgs["pageTitle"] = c.PageTitle
	return c.Render()
}

func (c App) IndexPost(message string) revel.Result {
	// required message
	c.Validation.Required(message).Message("Please input message!")
	if c.Validation.HasErrors(){
		c.Validation.Keep()		
		return c.Redirect(App.Index)
	}
	
	version := "0.0.1-alpha"
	newCase := models.NewCase(message, version)
	
	if err := c.Txn.Insert(newCase); err != nil {
		panic(err)
	}

	c.Flash.Success("Success")
	return c.Redirect(routes.App.Index())
}
