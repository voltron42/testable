package ipdf

import (
	"github.com/jung-kurt/gofpdf"
	"io"
)

var Singleton PdfFactory

type PdfFactory interface {
	HTMLBasicTokenize(htmlStr string) (list []gofpdf.HTMLBasicSegmentType)
	MakeFont(fontFileStr, encodingFileStr, dstDirStr string, msgWriter io.Writer, embed bool) (err error)
	UnicodeTranslator(r io.Reader) (f func(string) string, err error)
	UnicodeTranslatorFromFile(fileStr string) (f func(string) string, err error)
	New(orientationStr, unitStr, sizeStr, fontDirStr string) (f Pdf)
	NewCustom(init *gofpdf.InitType) (f Pdf)
	SVGBasicFileParse(svgFileStr string) (sig gofpdf.SVGBasicType, err error)
	SVGBasicParse(buf []byte) (sig gofpdf.SVGBasicType, err error)
	TtfParse(fileStr string) (TtfRec gofpdf.TtfType, err error)
}

type Pdf interface {
	AddFont(family, style, file string)
	AddFontFromReader(family, style string, r io.Reader)
	AddLayer(name string, visible bool) (layerID int)
	AddLink() int
	AddPage()
	AddPageFormat(orientation string, size gofpdf.SizeType)
	AliasNbPages(alias string)
	Arc(x, y, rx, ry, degRotate, degStart, degEnd float64, style string)
	BeginLayer(id int)
	Beziergon(points []gofpdf.PointType, style string)
	Bookmark(text string, level int, y float64)
	Cell(width, height float64, text string)
	CellFormat(width, height float64, text, border string, ln int, align string, fill bool, link int, linkStr string)
	Cellf(width, height float64, format string, args ...interface{})
	Circle(x, y, radius float64, style string)
	ClipCircle(x, y, radius float64, outline bool)
	ClipEllipse(x, y, rx, ry float64, outline bool)
	ClipEnd()
	ClipPolygon(outline bool, points ...gofpdf.PointType)
	ClipRect(x, y, width, height float64, outline bool)
	ClipRoundedRect(x, y, width, height, radius float64, outline bool)
	ClipText(x, y float64, text string, outline bool)
	Close()
	Curve(x0, y0, cx, cy, x1, y1 float64, style string)
	CurveBezierCubic(x0, y0, cx0, cy0, cx1, cy1, x1, y1 float64, style string)
	CurveCubic(x0, y0, cx0, cy0, x1, y1, cx1, cy1 float64, style string)
	Ellipse(x, y, rx, ry, degRotate float64, style string)
	EndLayer()
	Err() bool
	Error() error
	GetCellMargin() float64
	GetDrawColor() (int, int, int)
	GetFillColor() (int, int, int)
	GetFontSize() (ptSize, unitSize float64)
	GetLineWidth() float64
	GetMargins() (left, top, right, bottom float64)
	GetPageSize() (width, height float64)
	GetStringWidth(s string) float64
	GetTextColor() (int, int, int)
	GetX() float64
	GetXY() (float64, float64)
	GetY() float64
	HTMLBasicNew() (html gofpdf.HTMLBasicType)
	Image(imageNameStr string, x, y, w, h float64, flow bool, tp string, link int, linkStr string)
	ImageTypeFromMime(mimeStr string) (tp string)
	Line(x1, y1, x2, y2 float64)
	LinearGradient(x, y, w, h float64, r1, g1, b1 int, r2, g2, b2 int, x1, y1, x2, y2 float64)
	Link(x, y, w, h float64, link int)
	LinkString(x, y, w, h float64, linkStr string)
	Ln(h float64)
	MultiCell(w, h float64, txtStr, borderStr, alignStr string, fill bool)
	Ok() bool
	OpenLayerPane()
	Output(w io.Writer) error
	OutputAndClose(w io.WriteCloser) error
	OutputFileAndClose(fileStr string) error
	PageNo() int
	PageSize(pageNum int) (wd, ht float64, unitStr string)
	PointConvert(pt float64) float64
	Polygon(styleStr string, points ...gofpdf.PointType)
	RadialGradient(x, y, w, h float64, r1, g1, b1 int, r2, g2, b2 int, x1, y1, x2, y2, r float64)
	Rect(x, y, w, h float64, styleStr string)
	RegisterImage(fileStr, tp string) (info ImageInfoType)
	RegisterImageReader(imgName, tp string, r io.Reader) (info ImageInfoType)
	SVGBasicWrite(sb *gofpdf.SVGBasicType, scale float64)
	SetAcceptPageBreakFunc(fnc func() bool)
	SetAlpha(alpha float64, blendModeStr string)
	SetAuthor(authorStr string, isUTF8 bool)
	SetAutoPageBreak(auto bool, margin float64)
	SetCellMargin(margin float64)
	SetCompression(compress bool)
	SetCreator(creatorStr string, isUTF8 bool)
	SetDisplayMode(zoomStr, layoutStr string)
	SetDrawColor(r, g, b int)
	SetError(err error)
	SetErrorf(fmtStr string, args ...interface{})
	SetFillColor(r, g, b int)
	SetFont(familyStr, styleStr string, size float64)
	SetFontLocation(fontDirStr string)
	SetFontSize(size float64)
	SetFooterFunc(fnc func())
	SetHeaderFunc(fnc func())
	SetKeywords(keywordsStr string, isUTF8 bool)
	SetLeftMargin(margin float64)
	SetLineCapStyle(styleStr string)
	SetLineWidth(width float64)
	SetLink(link int, y float64, page int)
	SetMargins(left, top, right float64)
	SetProtection(actionFlag byte, userPassStr, ownerPassStr string)
	SetRightMargin(margin float64)
	SetSubject(subjectStr string, isUTF8 bool)
	SetTextColor(r, g, b int)
	SetTitle(titleStr string, isUTF8 bool)
	SetTopMargin(margin float64)
	SetX(x float64)
	SetXY(x, y float64)
	SetY(y float64)
	SplitLines(txt []byte, w float64) [][]byte
	String() string
	Text(x, y float64, txtStr string)
	Transform(tm gofpdf.TransformMatrix)
	TransformBegin()
	TransformEnd()
	TransformMirrorHorizontal(x float64)
	TransformMirrorLine(angle, x, y float64)
	TransformMirrorPoint(x, y float64)
	TransformMirrorVertical(y float64)
	TransformRotate(angle, x, y float64)
	TransformScale(scaleWd, scaleHt, x, y float64)
	TransformScaleX(scaleWd, x, y float64)
	TransformScaleXY(s, x, y float64)
	TransformScaleY(scaleHt, x, y float64)
	TransformSkew(angleX, angleY, x, y float64)
	TransformSkewX(angleX, x, y float64)
	TransformSkewY(angleY, x, y float64)
	TransformTranslate(tx, ty float64)
	TransformTranslateX(tx float64)
	TransformTranslateY(ty float64)
	UnicodeTranslatorFromDescriptor(cpStr string) (rep func(string) string)
	Write(h float64, txtStr string)
	WriteLinkID(h float64, displayStr string, linkID int)
	WriteLinkString(h float64, displayStr, targetStr string)
	Writef(h float64, fmtStr string, args ...interface{})
}

type ImageInfoType interface {
	Extent() (wd, ht float64)
	Height() float64
	Width() float64
}
