package request

import "time"

const (
	ApplicationJSON               ContentType = "application/json"
	ApplicationXWwwFormURLEncoded ContentType = "application/x-www-form-urlencoded"
	MultipartFormData             ContentType = "multipart/form-data"
)

type ContentType string

/*
     Method         = "OPTIONS"                ; Section 9.2
                    | "GET"                    ; Section 9.3
                    | "HEAD"                   ; Section 9.4
                    | "POST"                   ; Section 9.5
                    | "PUT"                    ; Section 9.6
                    | "DELETE"                 ; Section 9.7
                    | "TRACE"                  ; Section 9.8
                    | "CONNECT"                ; Section 9.9
                    | extension-method
   extension-method = token
     token          = 1*<any CHAR except CTLs or separators>
*/
type Client struct {
	URL         string
	Method      string
	Header      map[string]string
	Params      map[string]string
	Body        []byte
	Auth        Auth
	Timeout     time.Duration // second
	ContentType ContentType
}

type Auth struct {
	Username string
	Password string
}
