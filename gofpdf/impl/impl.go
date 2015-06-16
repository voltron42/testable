package pdf

import (
	"../adt"
	"github.com/jung-kurt/gofpdf"
	"io"
)

func NewPdfFactory() ipdf.PdfFactory {
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

func (p *PdfFactory) New(orientationStr, unitStr, sizeStr, fontDirStr string) (f ipdf.Pdf) {
	return &Pdf{gofpdf.New(orientationStr, unitStr, sizeStr, fontDirStr)}
}

func (p *PdfFactory) NewCustom(init *gofpdf.InitType) (f ipdf.Pdf) {
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
	p.pdf.AddPageFormat(orientation, size)
}

func (p *Pdf) AliasNbPages(alias string) {
	p.pdf.AliasNbPages(alias)
}

func (p *Pdf) Arc(x, y, rx, ry, degRotate, degStart, degEnd float64, style string) {
	p.pdf.Arc(x, y, rx, ry, degRotate, degStart, degEnd, style)
}

func (p *Pdf) BeginLayer(id int) {
	p.pdf.BeginLayer(id)
}

func (p *Pdf) Beziergon(points []gofpdf.PointType, style string) {
	p.pdf.Beziergon(points, style)
}

func (p *Pdf) Bookmark(text string, level int, y float64) {
	p.pdf.Bookmark(text, level, y)
}

func (p *Pdf) Cell(width, height float64, text string) {
	p.pdf.Cell(width, height, text)
}

func (p *Pdf) CellFormat(width, height float64, text, border string, ln int, align string, fill bool, link int, linkStr string) {
	p.pdf.CellFormat(width, height, text, border, ln, align, fill, link, linkStr)
}

func (p *Pdf) Cellf(width, height float64, format string, args ...interface{}) {
	p.pdf.Cellf(width, height, format, args...)
}

func (p *Pdf) Circle(x, y, radius float64, style string) {
	p.pdf.Circle(x, y, radius, style)
}

func (p *Pdf) ClipCircle(x, y, radius float64, outline bool) {
	p.pdf.ClipCircle(x, y, radius, outline)
}

func (p *Pdf) ClipEllipse(x, y, rx, ry float64, outline bool) {
	p.pdf.ClipEllipse(x, y, rx, ry, outline)
}

func (p *Pdf) ClipEnd() {
	p.pdf.ClipEnd()
}

func (p *Pdf) ClipPolygon(outline bool, points ...gofpdf.PointType) {
	p.pdf.ClipPolygon(points, outline)
}

func (p *Pdf) ClipRect(x, y, width, height float64, outline bool) {
	p.pdf.ClipRect(x, y, width, height, outline)
}

func (p *Pdf) ClipRoundedRect(x, y, width, height, radius float64, outline bool) {
	p.pdf.ClipRoundedRect(x, y, width, height, radius, outline)
}

func (p *Pdf) ClipText(x, y float64, text string, outline bool) {
	p.pdf.ClipText(x, y, text, outline)
}

func (p *Pdf) Close() {
	p.pdf.Close()
}

func (p *Pdf) Curve(x0, y0, cx, cy, x1, y1 float64, style string) {
	p.pdf.Curve(x0, y0, cx, cy, x1, y1, style)
}

func (p *Pdf) CurveBezierCubic(x0, y0, cx0, cy0, cx1, cy1, x1, y1 float64, style string) {
	p.pdf.CurveBezierCubic(x0, y0, cx0, cy0, cx1, cy1, x1, y1, style)
}

func (p *Pdf) CurveCubic(x0, y0, cx0, cy0, x1, y1, cx1, cy1 float64, style string) {
	p.pdf.CurveCubic(x0, y0, cx0, cy0, x1, y1, cx1, cy1, style)
}

func (p *Pdf) Ellipse(x, y, rx, ry, degRotate float64, style string) {
	p.pdf.Ellipse(x, y, rx, ry, degRotate, style)
}

func (p *Pdf) EndLayer() {
	p.pdf.EndLayer()
}

func (p *Pdf) Err() bool {
	return p.pdf.Err()
}

func (p *Pdf) Error() error {
	return p.pdf.Error()
}

func (p *Pdf) GetCellMargin() float64 {
	return p.pdf.GetCellMargin()
}

func (p *Pdf) GetDrawColor() (int, int, int) {
	return p.pdf.GetDrawColor()
}

func (p *Pdf) GetFillColor() (int, int, int) {
	return p.pdf.GetFillColor()
}

func (p *Pdf) GetFontSize() (ptSize, unitSize float64) {
	return p.pdf.GetFontSize()
}

func (p *Pdf) GetLineWidth() float64 {
	return p.pdf.GetLineWidth()
}

func (p *Pdf) GetMargins() (left, top, right, bottom float64) {
	return p.pdf.GetMargins()
}

func (p *Pdf) GetPageSize() (width, height float64) {
	return p.pdf.GetPageSize()
}

func (p *Pdf) GetStringWidth(s string) float64 {
	return p.pdf.GetStringWidth(s)
}

func (p *Pdf) GetTextColor() (int, int, int) {
	return p.pdf.GetTextColor()
}

func (p *Pdf) GetX() float64 {
	return p.pdf.GetX()
}

func (p *Pdf) GetXY() (float64, float64) {
	return p.pdf.GetXY()
}

func (p *Pdf) GetY() float64 {
	return p.pdf.GetY()
}

func (p *Pdf) HTMLBasicNew() (html gofpdf.HTMLBasicType) {
	return p.pdf.HTMLBasicNew()
}

func (p *Pdf) Image(imageNameStr string, x, y, w, h float64, flow bool, tp string, link int, linkStr string) {
	p.pdf.Image(imageNameStr, x, y, w, h, flow, tp, link, linkStr)
}

func (p *Pdf) ImageTypeFromMime(mimeStr string) (tp string) {
	return p.pdf.ImageTypeFromMime(mimeStr)
}

func (p *Pdf) Line(x1, y1, x2, y2 float64) {
	p.pdf.Line(x1, y1, x2, y2)
}

func (p *Pdf) LinearGradient(x, y, w, h float64, r1, g1, b1 int, r2, g2, b2 int, x1, y1, x2, y2 float64) {
	p.pdf.LinearGradient(x, y, w, h, r1, g1, b1, r2, g2, b2, x1, y1, x2, y2)
}

func (p *Pdf) Link(x, y, w, h float64, link int) {
	p.pdf.Link(x, y, w, h, link)
}

func (p *Pdf) LinkString(x, y, w, h float64, link string) {
	p.pdf.LinkString(x, y, w, h, link)

}

func (p *Pdf) Ln(h float64) {
	p.pdf.Ln(h)
}

func (p *Pdf) MultiCell(w, h float64, txtStr, borderStr, alignStr string, fill bool) {
	p.pdf.MultiCell(w, h, txtStr, borderStr, alignStr, fill)
}

func (p *Pdf) Ok() bool {
	return p.pdf.Ok()
}

func (p *Pdf) OpenLayerPane() {
	p.pdf.OpenLayerPane()
}

func (p *Pdf) Output(w io.Writer) error {
	return p.pdf.Output(w)
}

func (p *Pdf) OutputAndClose(w io.WriteCloser) error {
	return p.pdf.OutputAndClose(w)
}

func (p *Pdf) OutputFileAndClose(fileStr string) error {
	return p.pdf.OutputFileAndClose(fileStr)
}

func (p *Pdf) PageNo() int {
	return p.pdf.PageNo()
}

func (p *Pdf) PageSize(pageNum int) (wd, ht float64, unitStr string) {
	return p.pdf.PageSize(pageNum)
}

func (p *Pdf) PointConvert(pt float64) float64 {
	return p.pdf.PointConvert(pt)
}

func (p *Pdf) Polygon(styleStr string, points ...gofpdf.PointType) {
	p.pdf.Polygon(points, styleStr)
}

func (p *Pdf) RadialGradient(x, y, w, h float64, r1, g1, b1 int, r2, g2, b2 int, x1, y1, x2, y2, r float64) {
	p.pdf.RadialGradient(x, y, w, h, r1, g1, b1, r2, g2, b2, x1, y1, x2, y2, r)
}

func (p *Pdf) Rect(x, y, w, h float64, styleStr string) {
	p.pdf.Rect(x, y, w, h, styleStr)
}

func (p *Pdf) RegisterImage(fileStr, tp string) (info ipdf.ImageInfoType) {
	return &ImageInfoType{p.pdf.RegisterImage(fileStr, tp)}
}

func (p *Pdf) RegisterImageReader(imgName, tp string, r io.Reader) (info ipdf.ImageInfoType) {
	return &ImageInfoType{p.pdf.RegisterImageReader(imgName, tp, r)}
}

func (p *Pdf) SVGBasicWrite(sb *gofpdf.SVGBasicType, scale float64) {
	p.pdf.SVGBasicWrite(sb, scale)
}

func (p *Pdf) SetAcceptPageBreakFunc(fnc func() bool) {
	p.pdf.SetAcceptPageBreakFunc(fnc)
}

func (p *Pdf) SetAlpha(alpha float64, blendModeStr string) {
	p.pdf.SetAlpha(alpha, blendModeStr)
}

func (p *Pdf) SetAuthor(authorStr string, isUTF8 bool) {
	p.pdf.SetAuthor(authorStr, isUTF8)
}

func (p *Pdf) SetAutoPageBreak(auto bool, margin float64) {
	p.pdf.SetAutoPageBreak(auto, margin)
}

func (p *Pdf) SetCellMargin(margin float64) {
	p.pdf.SetCellMargin(margin)
}

func (p *Pdf) SetCompression(compress bool) {
	p.pdf.SetCompression(compress)
}

func (p *Pdf) SetCreator(creatorStr string, isUTF8 bool) {
	p.pdf.SetCreator(creatorStr, isUTF8)
}

func (p *Pdf) SetDisplayMode(zoomStr, layoutStr string) {
	p.pdf.SetDisplayMode(zoomStr, layoutStr)
}

func (p *Pdf) SetDrawColor(r, g, b int) {
	p.pdf.SetDrawColor(r, g, b)
}

func (p *Pdf) SetError(err error) {
	p.pdf.SetError(err)
}

func (p *Pdf) SetErrorf(fmtStr string, args ...interface{}) {
	p.pdf.SetErrorf(fmtStr, args...)
}

func (p *Pdf) SetFillColor(r, g, b int) {
	p.pdf.SetFillColor(r, g, b)
}

func (p *Pdf) SetFont(familyStr, styleStr string, size float64) {
	p.pdf.SetFont(familyStr, styleStr, size)
}

func (p *Pdf) SetFontLocation(fontDirStr string) {
	p.pdf.SetFontLocation(fontDirStr)
}

func (p *Pdf) SetFontSize(size float64) {
	p.pdf.SetFontSize(size)
}

func (p *Pdf) SetFooterFunc(fnc func()) {
	p.pdf.SetFooterFunc(fnc)
}

func (p *Pdf) SetHeaderFunc(fnc func()) {
	p.pdf.SetHeaderFunc(fnc)
}

func (p *Pdf) SetKeywords(keywordsStr string, isUTF8 bool) {
	p.pdf.SetKeywords(keywordsStr, isUTF8)
}

func (p *Pdf) SetLeftMargin(margin float64) {
	p.pdf.SetLeftMargin(margin)
}

func (p *Pdf) SetLineCapStyle(styleStr string) {
	p.pdf.SetLineCapStyle(styleStr)
}

func (p *Pdf) SetLineWidth(width float64) {
	p.pdf.SetLineWidth(width)
}

func (p *Pdf) SetLink(link int, y float64, page int) {
	p.pdf.SetLink(link, y, page)
}

func (p *Pdf) SetMargins(left, top, right float64) {
	p.pdf.SetMargins(left, top, right)
}

func (p *Pdf) SetProtection(actionFlag byte, userPassStr, ownerPassStr string) {
	p.pdf.SetProtection(actionFlag, userPassStr, ownerPassStr)
}

func (p *Pdf) SetRightMargin(margin float64) {
	p.pdf.SetRightMargin(margin)
}

func (p *Pdf) SetSubject(subjectStr string, isUTF8 bool) {
	p.pdf.SetSubject(subjectStr, isUTF8)
}

func (p *Pdf) SetTextColor(r, g, b int) {
	p.pdf.SetTextColor(r, g, b)
}

func (p *Pdf) SetTitle(titleStr string, isUTF8 bool) {
	p.pdf.SetTitle(titleStr, isUTF8)
}

func (p *Pdf) SetTopMargin(margin float64) {
	p.pdf.SetTopMargin(margin)
}

func (p *Pdf) SetX(x float64) {
	p.pdf.SetX(x)
}

func (p *Pdf) SetXY(x, y float64) {
	p.pdf.SetXY(x, y)
}

func (p *Pdf) SetY(y float64) {
	p.pdf.SetY(y)
}

func (p *Pdf) SplitLines(txt []byte, w float64) [][]byte {
	return p.pdf.SplitLines(txt, w)
}

func (p *Pdf) String() string {
	return p.pdf.String()
}

func (p *Pdf) Text(x, y float64, txtStr string) {
	p.pdf.Text(x, y, txtStr)
}

func (p *Pdf) Transform(tm gofpdf.TransformMatrix) {
	p.pdf.Transform(tm)
}

func (p *Pdf) TransformBegin() {
	p.pdf.TransformBegin()
}

func (p *Pdf) TransformEnd() {
	p.pdf.TransformEnd()
}

func (p *Pdf) TransformMirrorHorizontal(x float64) {
	p.pdf.TransformMirrorHorizontal(x)
}

func (p *Pdf) TransformMirrorLine(angle, x, y float64) {
	p.pdf.TransformMirrorLine(angle, x, y)
}

func (p *Pdf) TransformMirrorPoint(x, y float64) {
	p.pdf.TransformMirrorPoint(x, y)
}

func (p *Pdf) TransformMirrorVertical(y float64) {
	p.pdf.TransformMirrorVertical(y)
}

func (p *Pdf) TransformRotate(angle, x, y float64) {
	p.pdf.TransformRotate(angle, x, y)
}

func (p *Pdf) TransformScale(scaleWd, scaleHt, x, y float64) {
	p.pdf.TransformScale(scaleWd, scaleHt, x, y)
}

func (p *Pdf) TransformScaleX(scaleWd, x, y float64) {
	p.pdf.TransformScaleX(scaleWd, x, y)
}

func (p *Pdf) TransformScaleXY(s, x, y float64) {
	p.pdf.TransformScaleXY(s, x, y)
}

func (p *Pdf) TransformScaleY(scaleHt, x, y float64) {
	p.pdf.TransformScaleY(scaleHt, x, y)
}

func (p *Pdf) TransformSkew(angleX, angleY, x, y float64) {
	p.pdf.TransformSkew(angleX, angleY, x, y)
}

func (p *Pdf) TransformSkewX(angleX, x, y float64) {
	p.pdf.TransformSkewX(angleX, x, y)
}

func (p *Pdf) TransformSkewY(angleY, x, y float64) {
	p.pdf.TransformSkewY(angleY, x, y)
}

func (p *Pdf) TransformTranslate(tx, ty float64) {
	p.pdf.TransformTranslate(tx, ty)
}

func (p *Pdf) TransformTranslateX(tx float64) {
	p.pdf.TransformTranslateX(tx)
}

func (p *Pdf) TransformTranslateY(ty float64) {
	p.pdf.TransformTranslateY(ty)
}

func (p *Pdf) UnicodeTranslatorFromDescriptor(cpStr string) (rep func(string) string) {
	return p.pdf.UnicodeTranslatorFromDescriptor(cpStr)
}

func (p *Pdf) Write(h float64, txtStr string) {
	p.pdf.Writef(h, txtStr)
}

func (p *Pdf) WriteLinkID(h float64, displayStr string, linkID int) {
	p.pdf.WriteLinkID(h, displayStr, linkID)
}

func (p *Pdf) WriteLinkString(h float64, displayStr, targetStr string) {
	p.pdf.WriteLinkString(h, displayStr, targetStr)
}

func (p *Pdf) Writef(h float64, fmtStr string, args ...interface{}) {
	p.pdf.Writef(h, fmtStr, args...)
}

type ImageInfoType struct {
	imageInfoType *gofpdf.ImageInfoType
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
