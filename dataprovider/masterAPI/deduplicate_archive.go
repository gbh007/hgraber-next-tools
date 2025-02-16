package masterAPI

import (
	"context"
	"fmt"
	"io"
	"net/url"

	"github.com/gbh007/hgraber-next-agent-core/entities"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/gbh007/hgraber-next/pkg"
)

func (c *Client) DeduplicateArchive(ctx context.Context, body io.Reader) ([]entities.DeduplicateArchiveResult, error) {
	res, err := c.rawClient.APIDeduplicateArchivePost(ctx, serverapi.APIDeduplicateArchivePostReq{
		Data: body,
	})
	if err != nil {
		return nil, fmt.Errorf("master api: %w", err)
	}

	switch typedRes := res.(type) {
	case *serverapi.APIDeduplicateArchivePostOKApplicationJSON:
		return pkg.Map(*typedRes, func(raw serverapi.APIDeduplicateArchivePostOKItem) entities.DeduplicateArchiveResult {
			var u *url.URL

			if raw.BookOriginURL.IsSet() {
				u = &raw.BookOriginURL.Value
			}

			return entities.DeduplicateArchiveResult{
				TargetBookID:           raw.BookID,
				OriginBookURL:          u,
				EntryPercentage:        raw.EntryPercentage,
				ReverseEntryPercentage: raw.ReverseEntryPercentage,
			}
		}), nil

	case *serverapi.APIDeduplicateArchivePostBadRequest:
		return nil, fmt.Errorf("%w: %s", entities.MasterAPIBadRequest, typedRes.Details.Value)

	case *serverapi.APIDeduplicateArchivePostUnauthorized:
		return nil, fmt.Errorf("%w: %s", entities.MasterAPIUnauthorized, typedRes.Details.Value)

	case *serverapi.APIDeduplicateArchivePostForbidden:
		return nil, fmt.Errorf("%w: %s", entities.MasterAPIForbidden, typedRes.Details.Value)

	case *serverapi.APIDeduplicateArchivePostInternalServerError:
		return nil, fmt.Errorf("%w: %s", entities.MasterAPIInternalError, typedRes.Details.Value)

	default:
		return nil, entities.MasterAPIUnknownResponse
	}
}
