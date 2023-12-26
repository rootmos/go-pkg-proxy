package osext

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"fmt"
	"path/filepath"
	"os"

	"rootmos.io/go-pkg-proxy/internal/logging"
)

func Open(ctx context.Context, rawUrl string) (io.ReadCloser, error) {
	logger := logging.Get(ctx)

	u, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}

	switch u.Scheme {
	case "s3":
		return nil, fmt.Errorf("not implemented URL scheme: %s", u.Scheme)

		//s3c, err := S3(ctx)
		//if err != nil {
			//return nil, err
		//}

		//bucket, key := bucketKeyFromUrl(u)
		//logger, ctx = logging.WithAttrs(ctx, "bucket", bucket, "key", key)

		//logger.Debug("get object")
		//o, err := s3c.GetObject(ctx, &s3.GetObjectInput {
			//Bucket: aws.String(bucket),
			//Key: aws.String(key),
		//})

		//if err != nil {
			//return nil, err
		//}

		//logger.Debug("get object successful", "VersionId", aws.ToString(o.VersionId))

		//return o.Body, nil
	case "http", "https":
		rsp, err := http.Get(u.String())
		if err != nil {
			return nil, err
		}
		return rsp.Body, err
	case "", "file":
		path := filepath.Join(u.Host, u.Path)

		logger, _ = logging.WithAttrs(ctx, "path", path)

		logger.Debug("open")
		return os.Open(path)
	default:
		return nil, fmt.Errorf("unsupported URL scheme: %s", u.Scheme)
	}
}
