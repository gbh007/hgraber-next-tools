package masterAPI

import (
	"context"
	"fmt"
	"io"

	"github.com/gbh007/hgraber-next-agent-core/entities"
	"github.com/gbh007/hgraber-next/openapi/serverapi"
	"github.com/google/uuid"
)

func (c *Client) UploadArchive(ctx context.Context, body io.Reader) (uuid.UUID, error) {
	res, err := c.rawClient.APISystemImportArchivePost(ctx, serverapi.APISystemImportArchivePostReq{
		Data: body,
	})
	if err != nil {
		return uuid.Nil, fmt.Errorf("master api: %w", err)
	}

	switch typedRes := res.(type) {
	case *serverapi.APISystemImportArchivePostOK:
		return typedRes.ID, nil

	case *serverapi.APISystemImportArchivePostBadRequest:
		return uuid.Nil, fmt.Errorf("%w: %s", entities.MasterAPIBadRequest, typedRes.Details.Value)

	case *serverapi.APISystemImportArchivePostUnauthorized:
		return uuid.Nil, fmt.Errorf("%w: %s", entities.MasterAPIUnauthorized, typedRes.Details.Value)

	case *serverapi.APISystemImportArchivePostForbidden:
		return uuid.Nil, fmt.Errorf("%w: %s", entities.MasterAPIForbidden, typedRes.Details.Value)

	case *serverapi.APISystemImportArchivePostInternalServerError:
		return uuid.Nil, fmt.Errorf("%w: %s", entities.MasterAPIInternalError, typedRes.Details.Value)

	default:
		return uuid.Nil, entities.MasterAPIUnknownResponse
	}
}
