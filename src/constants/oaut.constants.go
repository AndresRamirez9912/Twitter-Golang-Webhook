package constants

// OAuth Links
const REQUEST_TOKEN_URL = "https://api.twitter.com/oauth/request_token"
const ACCESS_TOKEN_URL = "https://api.twitter.com/oauth/access_token"

// OAuth Parameters
const OAUTH_CONSUMER_KEY = "oauth_consumer_key"
const OAUTH_SIGNATURE_METHOD = "oauth_signature_method"
const OAUTH_TIMESTAMP = "oauth_timestamp"
const OAUTH_NONCE = "oauth_nonce"
const OAUTH_VERSION = "oauth_version"
const OAUTH_TOKEN = "oauth_token"
const OAUTH_METHOD = "HMAC-SHA1"
const OAUTH_VERIFIER = "oauth_verifier"
const OAUTH_TOKEN_SECRET = "oauth_token_secret"
const SCREEN_NAME = "screen_name"
const USER_ID = "user_id"

// Twitter Credentials
const API_KEY = "API_KEY"
const API_KEY_SECRET = "API_KEY_SECRET"
const ACCESS_TOKEN = "ACCESS_TOKEN"
const ACCESS_TOKEN_SECRET = "ACCESS_TOKEN_SECRET"

// Headers
const ACCEPT = "Accept"
const AUTHORIZATION = "Authorization"
const CONNECTION = "Connection"

// Rest Methods
const POST = "POST"
const GET = "GET"

// Templates
const VALIDATION_LINK_TEMPLATE = "https://api.twitter.com/oauth/authorize?oauth_token=%s"
const AUTHORIZATION_TEMPLATE = `OAuth oauth_consumer_key="%s", oauth_nonce="%s", oauth_signature="%s", oauth_signature_method="HMAC-SHA1", oauth_timestamp="%s", oauth_token="%s", oauth_version="1.0"`
