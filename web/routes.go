package web

import (
	"backend/accounts"
	"backend/structures"
)

func Route(app structures.SkyfallenCASApp){

	accounts.Route(app)

}