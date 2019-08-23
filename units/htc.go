package units

import (
	"encoding/xml"
	"fmt"
)

type HeatTransferCoefficient float64

func (htc HeatTransferCoefficient) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s := fmt.Sprintf("%.7f", htc)
	return e.EncodeElement(s, start)
}
