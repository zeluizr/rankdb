package design

// Copyright 2019 Vivino. All rights reserved
//
// See LICENSE file for license details

import (
	"github.com/Vivino/rankdb"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("rankdb", func() {
	Title("The RankDB API")
	Description("API for controlling RankDB")
	Consumes("application/json") // Media types supported by the API
	Consumes("application/msgpack", func() {
		Package("github.com/goadesign/goa/encoding/msgpack")
	})
	Produces("application/json") // Media types generated by the API
	Produces("application/msgpack", func() {
		Package("github.com/goadesign/goa/encoding/msgpack")
	})
	Version("1.0")
})

var _ = Resource("health", func() {
	BasePath("/")
	Action("health", func() {
		Routing(
			GET("/health"),
		)
		Metadata("swagger:summary", "Return system information")
		Description("Return system information")
		Response(OK, SysInfo)
		JWTAPIManage()
	})
	Action("root", func() {
		Routing(
			GET(""),
		)
		Metadata("swagger:summary", "Root")
		Description("Ping server")
		Response(OK, "application/json")
		JWTAPIManage()
	})
})

var _ = Resource("backup", func() {
	BasePath("/backup")
	Action("status", func() {
		Routing(
			GET(":backup_id"),
		)
		Params(func() {
			Param("backup_id", String)
		})
		Metadata("swagger:summary", "Return backup progress")
		Description("Return backup progress")
		Response(NotFound, ErrorMedia)
		Response(OK, BackupInfo)
		JWTAPIManage()
	})
	Action("delete", func() {
		Routing(
			DELETE(":backup_id"),
		)
		Params(func() {
			Param("backup_id", String)
		})
		Metadata("swagger:summary", "cancel Backup")
		Description("cancel backup")
		Response(NotFound, ErrorMedia)
		Response(NoContent)
		JWTAPIManage()
	})
})

// BackupInfo backup information.
var BackupInfo = MediaType("application/vnd.rankdb.backup_status+json", func() {
	Description("Backup Information")
	Attributes(func() {
		Attribute("cancelled", Boolean, "Will be true if backup was cancelled")
		Attribute("done", Boolean, "Will be true when the backup has finished processing")
		Attribute("uri", String, "URI of backed up content. Used for restore.")
		Attribute("lists", Integer, "Number of lists to be backed up")
		Attribute("lists_done", Integer, "Number of lists backed up now")
		Attribute("storage", String, "Storage used for backup")
		Attribute("size", Integer, "Size of stored data", func() {
			Metadata("struct:field:type", "int64")
		})
		Attribute("custom", HashOf(String, String), "Custom information provided by backup")
		Attribute("started", DateTime, "Time backup was started")
		Attribute("finished", DateTime, "Time backup was finished")
		Attribute("success", ArrayOf(String), func() {
			Description("Successful operations, list IDs")
			Example([]string{"highscore-uk-all"})
		})
		Attribute("errors", HashOf(String, String), func() {
			Description("Failed operations, indexed by list IDs")
			Example(msi{"highscore-dk-all": rankdb.ErrNotFound.Error()})
		})

	})
	Required("done", "cancelled", "lists", "storage", "size", "uri", "started", "lists_done")
	tiny := func() {
		Attribute("cancelled")
		Attribute("done")
		Attribute("lists")
		Attribute("lists_done")
		Attribute("uri")
		Attribute("size")
		Attribute("started")
		Attribute("finished")
	}
	def := func() {
		tiny()
		Attribute("storage")
		Attribute("errors")
		Attribute("custom")
	}
	View("tiny", tiny)
	View("default", def)
	View("full", func() {
		def()
		Attribute("success")
	})
})

// SysInfo is the account resource media type.
var SysInfo = MediaType("application/vnd.rankdb.sysinfo+json", func() {
	Description("System Info. The model is sparse and may contain other information.")
	Attributes(func() {
		Attribute("memory", HashOf(String, Any), "Memory Information", func() {
			Example(msi{"Alloc": 621480, "TotalAlloc": 621480, "Sys": 4262136, "Lookups": 5, "Mallocs": 5940})
		})
		Attribute("element_cache", HashOf(String, Any), "Element Cache Information", func() {
			Example(msi{"current_entries": 5000, "max_size": 10000})
		})
		Attribute("lazy_saver", HashOf(String, Any), func() {
			Description("Lazy saver cache information.")
			Metadata("struct:field:type", "json.RawMessage", "encoding/json")
			Example(msi{"current_entries": 5000, "max_size": 10000})
		})
	})

	View("default", func() {
		Attribute("memory")
		Attribute("element_cache")
		Attribute("lazy_saver")
	})
})

var _ = Resource("jwt", func() {
	BasePath("/jwt")
	Action("jwt", func() {
		Routing(
			POST(""),
		)
		Metadata("swagger:summary", "JWT key generator")
		Description("JWT key generator.\n" +
			"If left disabled in config, Unauthorized is returned")
		Params(func() {
			Param("scope", String, func() {
				Description("Create key with scope")
				Example("api:read")
				Default("api:read")
			})
			Param("only_lists", String, func() {
				Description("Create key with list restrictions.\n" +
					"Use commas to separate multiple lists")
				Example("highscore-game-75,highscore-game-76")
			})
			Param("only_elements", String, func() {
				Description("Create key with list restrictions.\n" +
					"Use commas to separate multiple elements")
				Example("1234,9876")
			})
			Param("expire", Integer, func() {
				Description("Expire token in this many minutes. Default is 24 hours.")
				Example(10)
				Default(24 * 60)
				Minimum(1)
			})
			Required("scope")
		})
		Response(OK, "text/plain", func() {
			Headers(func() {
				Header("Authorization", String, "Generated JWT. Can be used as bearer.")
			})
		})
		JWTAPIManage()
	})
})

var _ = Resource("static", func() {
	Description("Documentation")
	BasePath("/")
	Origin("*", func() { // CORS policy that applies to all actions and file servers
		Methods("GET") // of "public" resource
	})
	Files("/doc/*filepath", "api/public/swagger", func() {
		Metadata("swagger:summary", "Swagger UX static files")
	})
	Files("/api/swagger/*filepath", "api/swagger")
})

// JWT defines a security scheme using JWT.  The scheme uses the "Authorization" header to lookup
// the token.  It also defines then scope "api".
var JWT = JWTSecurity("jwt", func() {
	Header("Authorization")
	Scope("api:read", "API Read Access")
	Scope("api:update", "API Update Access.")
	Scope("api:delete", "API Delete Access.")
	Scope("api:manage", "API Manager Access.")
})

var JWTAPIRead = func() {
	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:read") // Enforce presence of "api:read" scope in JWT claims.
	})
	Response(Unauthorized, ErrorMedia)
}

var JWTAPIUpdate = func() {
	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:update") // Enforce presence of "api:update" scope in JWT claims.
	})
	Response(Unauthorized, ErrorMedia)
}

var JWTAPIDelete = func() {
	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:delete") // Enforce presence of "api:delete" scope in JWT claims.
	})
	Response(Unauthorized, ErrorMedia)
}

var JWTAPIManage = func() {
	Security(JWT, func() { // Use JWT to auth requests to this endpoint
		Scope("api:manage") // Enforce presence of "api:manage" scope in JWT claims.
	})
	Response(Unauthorized, ErrorMedia)
}
