include $(GOROOT)/src/Make.inc

TARG=github.com/pomack/oauth2_client
GOFILES=\
    const.go\
    facebook.go\
    google.go\
    json_value.go\
    linkedin.go\
    oauth1_client.go\
    oauth2_client.go\
    smugmug.go\
    twitter.go\
    yahoo.go\


include $(GOROOT)/src/Make.pkg
