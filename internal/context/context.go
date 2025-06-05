package context

import "context"

func SetUsername(ctx context.Context, username string) (context.Context, error) {
	ctx = context.WithValue(ctx, "username", username)
	if ctx == nil {
		return ctx, context.Canceled
	}
	return ctx, nil
}

func GetUsername(ctx context.Context) (string, error) {
	username, ok := ctx.Value("username").(string)
	if !ok {
		return "", context.Canceled
	}
	return username, nil
}
