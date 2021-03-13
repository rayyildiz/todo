package port

import "context"

const contextAuthKey = "auth:token"

func UserFromContext(ctx context.Context) string {
	if userId, ok := ctx.Value(contextAuthKey).(string); ok {
		return userId
	}
	return ""
}

func ContextWithUser(parent context.Context, user string) context.Context {
	return context.WithValue(parent, contextAuthKey, user)
}
