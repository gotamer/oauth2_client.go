package oauth2_client

type TwitterClient struct {
    OAuth1Client
}

func NewTwitterClient() *TwitterClient {
    return &TwitterClient{}
}

func (p *TwitterClient) Initialize(properties Properties) {
    if properties == nil {
        return
    }
    if v := properties.GetAsString("twitter.consumer.key"); len(v) > 0 {
        p.ConsumerKey = v
        //p.Credentials.Token = v
    }
    if v := properties.GetAsString("twitter.consumer.secret"); len(v) > 0 {
        p.ConsumerSecret = v
        //p.Credentials.Secret = v
    }
    if v := properties.GetAsString("twitter.callback_url"); len(v) > 0 {
        p.CallbackUrl = v
    }
    if v := properties.GetAsString("twitter.oauth1.request_token_path.url"); len(v) > 0 {
        p.RequestUrl = v
        //p.TemporaryCredentialRequestURI = v
    }
    if v := properties.GetAsString("twitter.oauth1.request_token_path.method"); len(v) > 0 {
        p.RequestUrlMethod = v
    }
    if v := properties.GetAsBool("twitter.oauth1.request_token_path.protected"); true {
        p.RequestUrlProtected = v
    }
    if v := properties.GetAsString("twitter.oauth1.authorization_path.url"); len(v) > 0 {
        p.AuthorizationUrl = v
        //p.ResourceOwnerAuthorizationURI = v
    }
    if v := properties.GetAsString("twitter.oauth1.access_token_path.url"); len(v) > 0 {
        p.AccessUrl = v
        //p.TokenRequestURI = v
    }
    if v := properties.GetAsString("twitter.oauth1.access_token_path.method"); len(v) > 0 {
        p.AccessUrlMethod = v
    }
    if v := properties.GetAsBool("twitter.oauth1.access_token_path.protected"); true {
        p.AccessUrlProtected = v
    }
    if v := properties.GetAsBool("twitter.oauth1.authorized_resource.protected"); true {
        p.AuthorizedResourceProtected = v
    }
    if v := properties.GetAsString("twitter.oauth1.scope"); len(v) > 0 {
        //p.Scope = v
    }
}


