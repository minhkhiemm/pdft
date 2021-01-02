package pdft

import gopdf "github.com/signintech/pdft/minigopdf"

// data type
type PDFLineData struct {
	contentObj gopdf.ContentObj
}

// create content object
func createLineObjFromBytes(lineType string, lineWidth, x1, y1, x2, y2 float64) gopdf.ContentObj {
	var cObj gopdf.ContentObj
	cObj.AppendStreamSetLineWidth(lineWidth)
	cObj.AppendStreamSetLineType(lineType)
	cObj.AppendStreamLine(x1, y1, x2, y2)
	return cObj
}

// set content object
func (p *PDFLineData) setLineObject(obj gopdf.ContentObj) {
	p.contentObj = obj
}
