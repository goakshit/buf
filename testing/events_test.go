package testing

import (
	"testing"

	"github.com/goakshit/buf/testing/events"
	"github.com/goakshit/greymatter/schema"
)

const (
	metadataFieldName           = "metadata"
	identifierMetadataFieldName = "identifier_metadata"
	payloadFieldName            = "payload"
	identifierFieldName         = "identifier"
)

func TestPayments(t *testing.T) {
	for _, e := range events.Events {
		reflectMsg := e.ProtoReflect()
		t.Log(reflectMsg.Descriptor().FullName())
		t.Log(schema.Schema{})

		// // Verify if metadata exists
		// rootMessageMetadataDescriptor := reflectMsg.Descriptor().Fields().ByName(metadataFieldName)
		// if rootMessageMetadataDescriptor == nil {
		// 	t.Errorf("Metadata field is missing in %s", string(reflectMsg.Descriptor().FullName()))
		// }
		// // Verify if metadata field number is 1
		// if !rootMessageMetadataDescriptor.Number().IsValid() || int(rootMessageMetadataDescriptor.Number()) != 1 {
		// 	t.Errorf("%s's metadata field number is not 1, it is %d instead", string(reflectMsg.Descriptor().FullName()), int(rootMessageMetadataDescriptor.Number()))
		// }

	}
}
