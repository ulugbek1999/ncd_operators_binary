package settings

import "github.com/gorilla/securecookie"

// Hash keys should be at least 32 bytes long
var hashKey = []byte("ncd")

// Block keys should be 16 bytes (AES-128) or 32 bytes (AES-256) long.
// Shorter keys may weaken the encryption used.
var blockKey = []byte("ncd-operator-module-written-ingo")
var SecureCookie = securecookie.New(hashKey, blockKey)
