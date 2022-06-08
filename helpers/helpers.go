package helpers

import (
	"net/http"
	"net/url"

	"github.com/thedevsaddam/govalidator"
)

func ValidatePayloadsArticle(post interface{}, r *http.Request) url.Values {
	rules := govalidator.MapData{
		"title": 			[]string{"required", "min:20"},
		"content": 		[]string{"required", "min:200"},
		"category": 	[]string{"required", "min:3"},
		"status": 		[]string{"required", "in:Publish,Draft,Thrash"},
	}

	messages := govalidator.MapData{
		"status": 		[]string{"required: Publish, Draft or Thrash"},
	}

	options := govalidator.Options{
		Request: r,
		Data: post,
		Rules: rules,
		Messages: messages,
	}

	validation := govalidator.New(options)
	err := validation.ValidateJSON()

	return err
}