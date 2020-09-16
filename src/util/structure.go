package util

import "html/template"

type MetaData struct {
	ID 							string
	Title 						template.HTML
	Size 						string
	CTime 						string
	Icon 						string
}