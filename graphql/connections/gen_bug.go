// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package connections

import (
	"fmt"

	"github.com/MichaelMure/git-bug/bug"
	"github.com/MichaelMure/git-bug/graphql/models"
)

type BugSnapshotEdger func(value bug.Snapshot, offset int) Edge
type BugSnapshotConMaker func(edges []models.BugEdge, info models.PageInfo, totalCount int) models.BugConnection

func BugSnapshotCon(source []bug.Snapshot, edger BugSnapshotEdger, conMaker BugSnapshotConMaker, input models.ConnectionInput) (models.BugConnection, error) {
	var edges []models.BugEdge
	var pageInfo models.PageInfo

	emptyCon := conMaker(edges, pageInfo, 0)

	offset := 0

	if input.After != nil {
		for i, value := range source {
			edge := edger(value, i)
			if edge.GetCursor() == *input.After {
				// remove all previous element including the "after" one
				source = source[i+1:]
				offset = i + 1
				break
			}
		}
	}

	if input.Before != nil {
		for i, value := range source {
			edge := edger(value, i+offset)

			if edge.GetCursor() == *input.Before {
				// remove all after element including the "before" one
				break
			}

			edges = append(edges, edge.(models.BugEdge))
		}
	} else {
		edges = make([]models.BugEdge, len(source))

		for i, value := range source {
			edges[i] = edger(value, i+offset).(models.BugEdge)
		}
	}

	if input.First != nil {
		if *input.First < 0 {
			return emptyCon, fmt.Errorf("first less than zero")
		}

		if len(edges) > *input.First {
			// Slice result to be of length first by removing edges from the end
			edges = edges[:*input.First]
			pageInfo.HasNextPage = true
		}
	}

	if input.Last != nil {
		if *input.Last < 0 {
			return emptyCon, fmt.Errorf("last less than zero")
		}

		if len(edges) > *input.Last {
			// Slice result to be of length last by removing edges from the start
			edges = edges[len(edges)-*input.Last:]
			pageInfo.HasPreviousPage = true
		}
	}

	con := conMaker(edges, pageInfo, len(source))

	return con, nil
}
