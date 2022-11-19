package consts

const (
	REDIS_KEY_LOGIN_REQUEST = "join:login:req:%s"  // Username, token as value
	REDIS_KEY_VALID_SESSION = "join:admin:sess:%s" // Session, UserID as value
)
