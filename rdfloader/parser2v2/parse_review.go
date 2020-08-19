// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package parser2v2

import (
	"fmt"
	gordfParser "github.com/RishabhBhatnagar/gordf/rdfloader/parser"
	"github.com/spdx/tools-golang/spdx"
)

func (parser *rdfParser2_2) setReviewFromNode(reviewedNode *gordfParser.Node) error {
	review := spdx.Review2_2{}
	for _, triple := range parser.nodeToTriples[reviewedNode.String()] {
		switch triple.Predicate.ID {
		case RDF_TYPE:
			// cardinality: exactly 1
			continue
		case RDFS_COMMENT:
			// cardinality: max 1
			review.ReviewComment = triple.Object.ID
		case SPDX_REVIEW_DATE:
			// cardinality: exactly 1
			review.ReviewDate = triple.Object.ID
		case SPDX_REVIEWER:
			// cardinality: max 1
			review.Reviewer = triple.Object.ID
		default:
			return fmt.Errorf("unknown predicate %v for review triples", triple.Predicate)
		}
	}
	parser.doc.Reviews = append(parser.doc.Reviews, &review)
	return nil
}
