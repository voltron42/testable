package pdf

import (
	"../adt"
	"github.com/jung-kurt/gofpdf"
	"io"
)

func NewPdfFactory() *ipdf.PdfFactory {
	return &PdfFactory{}

}

type PdfFactory struct {
}

func (p *PdfFactory) HTMLBasicTokenize(htmlStr string) (list []gofpdf.HTMLBasicSegmentType) {
	return gofpdf.HTMLBasicTokenize(htmlStr)
}

func (p *PdfFactory) MakeFont(fontFileStr, encodingFileStr, dstDirStr string, msgWriter io.Writer, embed bool) (err error) {
	return gofpdf.MakeFont(fontFileStr, encodingFileStr, dstDirStr, msgWriter, embed)
}

func (p *PdfFactory) UnicodeTranslator(r io.Reader) (f func(string) string, err error) {
	return gofpdf.UnicodeTranslator(r)
}

func (p *PdfFactory) UnicodeTranslatorFromFile(fileStr string) (f func(string) string, err error) {
	return gofpdf.UnicodeTranslatorFromFile(fileStr)
}

func (p *PdfFactory) New(orientationStr, unitStr, sizeStr, fontDirStr string) (f *ipdf.Pdf) {
	return &Pdf{gofpdf.New(orientationStr, unitStr, sizeStr, fontDirStr)}
}

func (p *PdfFactory) NewCustom(init *gofpdf.InitType) (f *ipdf.Pdf) {
	return &Pdf{gofpdf.NewCustom(init)}
}

func (p *PdfFactory) SVGBasicFileParse(svgFileStr string) (sig gofpdf.SVGBasicType, err error) {
	return gofpdf.SVGBasicFileParse(svgFileStr)
}

func (p *PdfFactory) SVGBasicParse(buf []byte) (sig gofpdf.SVGBasicType, err error) {
	return gofpdf.SVGBasicParse(buf)
}

func (p *PdfFactory) TtfParse(fileStr string) (TtfRec gofpdf.TtfType, err error) {
	return gofpdf.TtfParse(fileStr)
}

type Pdf struct {
	pdf *gofpdf.Fpdf
}

func (p *Pdf) AddFont(family, style, file string) {
	p.pdf.AddFont(family, style, file)
}

func (p *Pdf) AddFontFromReader(family, style string, r io.Reader) {
	p.pdf.AddFontFromReader(family, style, r)
}

func (p *Pdf) AddLayer(name string, visible bool) (layerID int) {
	return p.pdf.AddLayer(name, visible)
}

func (p *Pdf) AddLink() int {
	return p.pdf.AddLink()
}

func (p *Pdf) AddPage() {
	p.pdf.AddPage()
}

func (p *Pdf) AddPageFormat(orientation string, size gofpdf.SizeType) {

}

func (p *Pdf) AliasNbPages(alias string) {

}

func (p *Pdf) Arc(x, y, rx, ry, degRotate, degStart, degEnd float64, style string) {

}

func (p *Pdf) BeginLayer(id int) {

}

func (p *Pdf) Beziergon(points []gofpdf.PointType, style string) {

}

func (p *Pdf) Bookmark(text string, level int, y float64) {

}

func (p *Pdf) Cell(width, height float64, text string) {

}

func (p *Pdf) CellFormat(width, height float64, text, border string, ln int, align string, fill bool, link int, linkStr string) {

}

func (p *Pdf) Cellf(width, height float64, format string, args ...interface{}) {

}

func (p *Pdf) Circle(x, y, radius float64, style string) {

}

func (p *Pdf) ClipCircle(x, y, radius float64, outline bool) {

}

func (p *Pdf) ClipEllipse(x, y, rx, ry float64, outline bool) {

}

func (p *Pdf) ClipEnd() {

}

func (p *Pdf) ClipPolygon(outline bool, points ...gofpdf.PointType) {

}

func (p *Pdf) ClipRect(x, y, width, height float64, outline bool) {

}

func (p *Pdf) ClipRoundedRect(x, y, width, height, radius float64, outline bool) {

}

func (p *Pdf) ClipText(x, y, float64, text string, outline bool) {

}

func (p *Pdf) Close() {

}

func (p *Pdf) Curve(x0, y0, cx, cy, x1, y1 float64, style string) {

}

func (p *Pdf) CurveBezierCubic(x0, y0, cx0, cy0, cx1, cy1, x1, y1 float64, style string) {

}

func (p *Pdf) CurveCubic(x0, y0, cx0, cy0, x1, y1, cx1, cy1 float64, style string) {

}

func (p *Pdf) Ellipse(x, y, rx, ry, degRotate float64, style string) {

}

func (p *Pdf) EndLayer() {

}

func (p *Pdf) Err() bool {
	return false
}

func (p *Pdf) Error() error {
	return nil
}

func (p *Pdf) GetCellMargin() float64 {
	return 0.0
}

func (p *Pdf) GetDrawColor() (int, int, int) {
	return 0, 0, 0
}

func (p *Pdf) GetFillColor() (int, int, int) {
	return 0, 0, 0
}

func (p *Pdf) GetFontSize() (ptSize, unitSize float64) {
	return 0.0, 0.0
}

func (p *Pdf) GetLineWidth() float64 {
	return 0.0
}

func (p *Pdf) GetMargins() (left, top, right, bottom float64) {
	return 0.0, 0.0, 0.0, 0.0
}

func (p *Pdf) GetPageSize() (width, height float64) {
	return 0.0, 0.0
}

func (p *Pdf) GetStringWidth(s string) float64 {
	return 0.0
}

func (p *Pdf) GetTextColor() (int, int, int) {
	return 0, 0, 0
}

func (p *Pdf) GetX() float64 {
	return 0.0
}

func (p *Pdf) GetXY() (float64, float64) {
	return 0.0, 0.0
}

func (p *Pdf) GetY() float64 {
	return 0.0
}

func (p *Pdf) HTMLBasicNew() (html gofpdf.HTMLBasicType) {
	return gofpdf.HTMLBasicType{}
}

func (p *Pdf) Image(imageNameStr string, x, y, w, h float64, flow bool, tp string, link int, linkStr string) {

}

func (p *Pdf) ImageTypeFromMime(mimeStr string) (tp string) {

}

func (p *Pdf) Line(x1, y1, x2, y2 float64) {

}

func (p *Pdf) LinearGradient(x, y, w, h float64, r1, g1, b1 int, r2, g2, b2 int, x1, y1, x2, y2 float64) {

}

func (p *Pdf) Link(x, y, w, h float64, link int) {

}

func (p *Pdf) LinkString(x, y, w, h float64, linkStr string) {

}

func (p *Pdf) Ln(h float64) {

}

func (p *Pdf) MultiCell(w, h float64, txtStr, borderStr, alignStr string, fill bool) {

}

func (p *Pdf) Ok() bool {

}

func (p *Pdf) OpenLayerPane() {

}

func (p *Pdf) Output(w io.Writer) error {

}

func (p *Pdf) OutputAndClose(w io.WriteCloser) error {

}

func (p *Pdf) OutputFileAndClose(fileStr string) error {

}

func (p *Pdf) PageNo() int {

}

func (p *Pdf) PageSize(pageNum int) (wd, ht float64, unitStr string) {

}

func (p *Pdf) PointConvert(pt float64) float64 {

}

func (p *Pdf) Polygon(styleStr string, points ...gofpdf.PointType) {

}

func (p *Pdf) RadialGradient(x, y, w, h float64, r1, g1, b1 int, r2, g2, b2 int, x1, y1, x2, y2, r float64) {

}

func (p *Pdf) Rect(x, y, w, h float64, styleStr string) {

}

func (p *Pdf) RegisterImage(fileStr, tp string) (info *ipdf.ImageInfoType) {

}

func (p *Pdf) RegisterImageReader(imgName, tp string, r io.Reader) (info *ipdf.ImageInfoType) {

}

func (p *Pdf) SVGBasicWrite(sb *gofpdf.SVGBasicType, scale float64) {

}

func (p *Pdf) SetAcceptPageBreakFunc(fnc func() bool) {

}

func (p *Pdf) SetAlpha(alpha float64, blendModeStr string) {

}

func (p *Pdf) SetAuthor(authorStr string, isUTF8 bool) {

}

func (p *Pdf) SetAutoPageBreak(auto bool, margin float64) {

}

func (p *Pdf) SetCellMargin(margin float64) {

}

func (p *Pdf) SetCompression(compress bool) {

}

func (p *Pdf) SetCreator(creatorStr string, isUTF8 bool) {

}

func (p *Pdf) SetDisplayMode(zoomStr, layoutStr string) {

}

func (p *Pdf) SetDrawColor(r, g, b int) {

}

func (p *Pdf) SetError(err error) {

}

func (p *Pdf) SetErrorf(fmtStr string, args ...interface{}) {

}

func (p *Pdf) SetFillColor(r, g, b int) {

}

func (p *Pdf) SetFont(familyStr, styleStr string, size float64) {

}

func (p *Pdf) SetFontLocation(fontDirStr string) {

}

func (p *Pdf) SetFontSize(size float64) {

}

func (p *Pdf) SetFooterFunc(fnc func()) {

}

func (p *Pdf) SetHeaderFunc(fnc func()) {

}

func (p *Pdf) SetKeywords(keywordsStr string, isUTF8 bool) {

}

func (p *Pdf) SetLeftMargin(margin float64) {

}

func (p *Pdf) SetLineCapStyle(styleStr string) {

}

func (p *Pdf) SetLineWidth(width float64) {

}

func (p *Pdf) SetLink(link int, y float64, page int) {

}

func (p *Pdf) SetMargins(left, top, right float64) {

}

func (p *Pdf) SetProtection(actionFlag byte, userPassStr, ownerPassStr string) {

}

func (p *Pdf) SetRightMargin(margin float64) {

}

func (p *Pdf) SetSubject(subjectStr string, isUTF8 bool) {

}

func (p *Pdf) SetTextColor(r, g, b int) {

}

func (p *Pdf) SetTitle(titleStr string, isUTF8 bool) {

}

func (p *Pdf) SetTopMargin(margin float64) {

}

func (p *Pdf) SetX(x float64) {

}

func (p *Pdf) SetXY(x, y float64) {

}

func (p *Pdf) SetY(y float64) {

}

func (p *Pdf) SplitLines(txt []byte, w float64) [][]byte {

}

func (p *Pdf) String() string {

}

func (p *Pdf) Text(x, y float64, txtStr string) {

}

func (p *Pdf) Transform(tm gofpdf.TransformMatrix) {

}

func (p *Pdf) TransformBegin() {

}

func (p *Pdf) TransformEnd() {

}

func (p *Pdf) TransformMirrorHorizontal(x float64) {

}

func (p *Pdf) TransformMirrorLine(angle, x, y float64) {

}

func (p *Pdf) TransformMirrorPoint(x, y float64) {

}

func (p *Pdf) TransformMirrorVertical(y float64) {

}

func (p *Pdf) TransformRotate(angle, x, y float64) {

}

func (p *Pdf) TransformScale(scaleWd, scaleHt, x, y float64) {

}

func (p *Pdf) TransformScaleX(scaleWd, x, y float64) {

}

func (p *Pdf) TransformScaleXY(s, x, y float64) {

}

func (p *Pdf) TransformScaleY(scaleHt, x, y float64) {

}

func (p *Pdf) TransformSkew(angleX, angleY, x, y float64) {

}

func (p *Pdf) TransformSkewX(angleX, x, y float64) {

}

func (p *Pdf) TransformSkewY(angleY, x, y float64) {

}

func (p *Pdf) TransformTranslate(tx, ty float64) {

}

func (p *Pdf) TransformTranslateX(tx float64) {

}

func (p *Pdf) TransformTranslateY(ty float64) {

}

func (p *Pdf) UnicodeTranslatorFromDescriptor(cpStr string) (rep func(string) string) {

}

func (p *Pdf) Write(h float64, txtStr string) {

}

func (p *Pdf) WriteLinkID(h float64, displayStr string, linkID int) {

}

func (p *Pdf) WriteLinkString(h float64, displayStr, targetStr string) {

}

func (p *Pdf) Writef(h float64, fmtStr string, args ...interface{}) {

}

type ImageInfoType struct {
	imageInfoType gofpdf.ImageInfoType
}

func (i *ImageInfoType) Extent() (wd, ht float64) {
	return i.imageInfoType.Extent()
}

func (i *ImageInfoType) Height() float64 {
	return i.imageInfoType.Height()
}

func (i *ImageInfoType) Width() float64 {
	return i.imageInfoType.Width()
}
