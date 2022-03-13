package usecase

import (
	"encoding/base64"
	"fmt"

	"github.com/aed86/proof_of_work/internal/pkg/challenger/model"
)

var headerFormatTemplate = "%s:%d:%d:%d"

func (*powHeaderBuilder) Build(solution model.Solution) string {
	return fmt.Sprintf(
		headerFormatTemplate,
		base64.StdEncoding.EncodeToString(solution.Challenge),
		solution.Nonce,
		solution.LeadingZeros,
		solution.Timestamp,
	)
}
