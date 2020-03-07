/*This file is computer-generated*/
package css

import "strconv"

type unitOrNoneValue interface {
	ruleable
	unitOrNoneValue()
}
type unitOrNoneType string

func (self unitOrNoneType) Rule() Rule  { return Rule(self) }
func (unitOrNoneType) unitOrNoneValue() {}

type colorValue interface {
	ruleable
	colorValue()
}
type colorType string

func (self colorType) Rule() Rule { return Rule(self) }
func (colorType) colorValue()     {}

type gridTemplateValue interface {
	ruleable
	gridTemplateValue()
}
type gridTemplateType string

func (self gridTemplateType) Rule() Rule    { return Rule(self) }
func (gridTemplateType) gridTemplateValue() {}

type overflowValue interface {
	ruleable
	overflowValue()
}
type overflowType string

func (self overflowType) Rule() Rule { return Rule(self) }
func (overflowType) overflowValue()  {}

type unitOrAutoValue interface {
	ruleable
	unitOrAutoValue()
}
type unitOrAutoType string

func (self unitOrAutoType) Rule() Rule  { return Rule(self) }
func (unitOrAutoType) unitOrAutoValue() {}

type sizeValue interface {
	ruleable
	sizeValue()
}
type sizeType string

func (self sizeType) Rule() Rule { return Rule(self) }
func (sizeType) sizeValue()      {}

type uintValue interface {
	ruleable
	uintValue()
}
type uintType string

func (self uintType) Rule() Rule { return Rule(self) }
func (uintType) uintValue()      {}

type unitValue interface {
	ruleable
	unitValue()
}
type unitType string

func (self unitType) Rule() Rule { return Rule(self) }
func (unitType) unitValue()      {}

type unitAndUnitValue interface {
	ruleable
	unitAndUnitValue()
}
type unitAndUnitType string

func (self unitAndUnitType) Rule() Rule   { return Rule(self) }
func (unitAndUnitType) unitAndUnitValue() {}

type shadowValue interface {
	ruleable
	shadowValue()
}
type shadowType string

func (self shadowType) Rule() Rule { return Rule(self) }
func (shadowType) shadowValue()    {}

type boxValue interface {
	ruleable
	boxValue()
}
type boxType string

func (self boxType) Rule() Rule { return Rule(self) }
func (boxType) boxValue()       {}

type borderStyleValue interface {
	ruleable
	borderStyleValue()
}
type borderStyleType string

func (self borderStyleType) Rule() Rule   { return Rule(self) }
func (borderStyleType) borderStyleValue() {}

type breakValue interface {
	ruleable
	breakValue()
}
type breakType string

func (self breakType) Rule() Rule { return Rule(self) }
func (breakType) breakValue()     {}

type normalOrAutoValue interface {
	ruleable
	normalOrAutoValue()
}
type normalOrAutoType string

func (self normalOrAutoType) Rule() Rule    { return Rule(self) }
func (normalOrAutoType) normalOrAutoValue() {}

type thicknessValue interface {
	ruleable
	thicknessValue()
}
type thicknessType string

func (self thicknessType) Rule() Rule { return Rule(self) }
func (thicknessType) thicknessValue() {}

type numberValue interface {
	ruleable
	numberValue()
}
type numberType string

func (self numberType) Rule() Rule { return Rule(self) }
func (numberType) numberValue()    {}

type gridStopValue interface {
	ruleable
	gridStopValue()
}
type gridStopType string

func (self gridStopType) Rule() Rule { return Rule(self) }
func (gridStopType) gridStopValue()  {}

type uintOrUnitValue interface {
	ruleable
	uintOrUnitValue()
}
type uintOrUnitType string

func (self uintOrUnitType) Rule() Rule  { return Rule(self) }
func (uintOrUnitType) uintOrUnitValue() {}

type imageValue interface {
	ruleable
	imageValue()
}
type imageType string

func (self imageType) Rule() Rule { return Rule(self) }
func (imageType) imageValue()     {}

type pageBreakValue interface {
	ruleable
	pageBreakValue()
}
type pageBreakType string

func (self pageBreakType) Rule() Rule { return Rule(self) }
func (pageBreakType) pageBreakValue() {}

type durationValue interface {
	ruleable
	durationValue()
}
type durationType string

func (self durationType) Rule() Rule { return Rule(self) }
func (durationType) durationValue()  {}

type gridAutoValue interface {
	ruleable
	gridAutoValue()
}
type gridAutoType string

func (self gridAutoType) Rule() Rule { return Rule(self) }
func (gridAutoType) gridAutoValue()  {}

type normalOrUnitOrAutoValue interface {
	ruleable
	normalOrUnitOrAutoValue()
}
type normalOrUnitOrAutoType string

func (self normalOrUnitOrAutoType) Rule() Rule          { return Rule(self) }
func (normalOrUnitOrAutoType) normalOrUnitOrAutoValue() {}

type integerOrAutoValue interface {
	ruleable
	integerOrAutoValue()
}
type integerOrAutoType string

func (self integerOrAutoType) Rule() Rule     { return Rule(self) }
func (integerOrAutoType) integerOrAutoValue() {}

type nameValue interface {
	ruleable
	nameValue()
}
type nameType string

func (self nameType) Rule() Rule { return Rule(self) }
func (nameType) nameValue()      {}

type animationDirectionValue interface {
	ruleable
	animationDirectionValue()
}
type animationDirectionType string

func (self animationDirectionType) Rule() Rule          { return Rule(self) }
func (animationDirectionType) animationDirectionValue() {}

func SetAnimationDirection(value animationDirectionValue) Rule {
	return "animation-direction:" + value.Rule() + ";"
}

type borderTopRightRadiusValue interface {
	ruleable
	borderTopRightRadiusValue()
}
type borderTopRightRadiusType string

func (self borderTopRightRadiusType) Rule() Rule            { return Rule(self) }
func (borderTopRightRadiusType) borderTopRightRadiusValue() {}

func SetBorderTopRightRadius(value borderTopRightRadiusValue) Rule {
	return "border-top-right-radius:" + value.Rule() + ";"
}

type gridRowValue interface {
	ruleable
	gridRowValue()
}
type gridRowType string

func (self gridRowType) Rule() Rule { return Rule(self) }
func (gridRowType) gridRowValue()   {}

func SetGridRow(value gridRowValue) Rule {
	return "grid-row:" + value.Rule() + ";"
}
func SetGridAutoRows(value gridAutoValue) Rule {
	return "grid-auto-rows:" + value.Rule() + ";"
}

type wordBreakValue interface {
	ruleable
	wordBreakValue()
}
type wordBreakType string

func (self wordBreakType) Rule() Rule { return Rule(self) }
func (wordBreakType) wordBreakValue() {}

func SetWordBreak(value wordBreakValue) Rule {
	return "word-break:" + value.Rule() + ";"
}

type animationNameValue interface {
	ruleable
	animationNameValue()
}
type animationNameType string

func (self animationNameType) Rule() Rule     { return Rule(self) }
func (animationNameType) animationNameValue() {}

func SetAnimationName(value animationNameValue) Rule {
	return "animation-name:" + value.Rule() + ";"
}

type boxDecorationBreakValue interface {
	ruleable
	boxDecorationBreakValue()
}
type boxDecorationBreakType string

func (self boxDecorationBreakType) Rule() Rule          { return Rule(self) }
func (boxDecorationBreakType) boxDecorationBreakValue() {}

func SetBoxDecorationBreak(value boxDecorationBreakValue) Rule {
	return "box-decoration-break:" + value.Rule() + ";"
}

type contentValue interface {
	ruleable
	contentValue()
}
type contentType string

func (self contentType) Rule() Rule { return Rule(self) }
func (contentType) contentValue()   {}

func SetContent(value contentValue) Rule {
	return "content:" + value.Rule() + ";"
}

type gridColumnValue interface {
	ruleable
	gridColumnValue()
}
type gridColumnType string

func (self gridColumnType) Rule() Rule  { return Rule(self) }
func (gridColumnType) gridColumnValue() {}

func SetGridColumn(value gridColumnValue) Rule {
	return "grid-column:" + value.Rule() + ";"
}

type imageRenderingValue interface {
	ruleable
	imageRenderingValue()
}
type imageRenderingType string

func (self imageRenderingType) Rule() Rule      { return Rule(self) }
func (imageRenderingType) imageRenderingValue() {}

func SetImageRendering(value imageRenderingValue) Rule {
	return "image-rendering:" + value.Rule() + ";"
}

type borderImageRepeatValue interface {
	ruleable
	borderImageRepeatValue()
}
type borderImageRepeatType string

func (self borderImageRepeatType) Rule() Rule         { return Rule(self) }
func (borderImageRepeatType) borderImageRepeatValue() {}

func SetBorderImageRepeat(value borderImageRepeatValue) Rule {
	return "border-image-repeat:" + value.Rule() + ";"
}

type fontValue interface {
	ruleable
	fontValue()
}
type fontType string

func (self fontType) Rule() Rule { return Rule(self) }
func (fontType) fontValue()      {}

func SetFont(value fontValue) Rule {
	return "font:" + value.Rule() + ";"
}
func SetMinWidth(value unitOrNoneValue) Rule {
	return "min-width:" + value.Rule() + ";"
}

type filterValue interface {
	ruleable
	filterValue()
}
type filterType string

func (self filterType) Rule() Rule { return Rule(self) }
func (filterType) filterValue()    {}

func SetFilter(value filterValue) Rule {
	return "filter:" + value.Rule() + ";"
}
func SetFlexShrink(value numberValue) Rule {
	return "flex-shrink:" + value.Rule() + ";"
}
func SetWidth(value unitOrAutoValue) Rule {
	return "width:" + value.Rule() + ";"
}
func SetBorderBottomRightRadius(value unitValue) Rule {
	return "border-bottom-right-radius:" + value.Rule() + ";"
}

type flexFlowValue interface {
	ruleable
	flexFlowValue()
}
type flexFlowType string

func (self flexFlowType) Rule() Rule { return Rule(self) }
func (flexFlowType) flexFlowValue()  {}

func SetFlexFlow(value flexFlowValue) Rule {
	return "flex-flow:" + value.Rule() + ";"
}

type gridGapValue interface {
	ruleable
	gridGapValue()
}
type gridGapType string

func (self gridGapType) Rule() Rule { return Rule(self) }
func (gridGapType) gridGapValue()   {}

func SetGridGap(value gridGapValue) Rule {
	return "grid-gap:" + value.Rule() + ";"
}

func SetWillChange(properties ...interface{}) Rule {
	var names string

	/*for i, property := range properties {
		var s = NewStyle()
		var catcher = propertyCatcher("")
		s.Stylable = &catcher

		reflect.ValueOf(property).Call([]reflect.Value{reflect.ValueOf(&s)})

		names += *((*string)(s.Stylable.(*propertyCatcher)))
		if i != len(properties) - 1 {
			names += ","
		}
	}*/

	return "will-change: " + unitType(names).Rule() + ";"
}

type captionSideValue interface {
	ruleable
	captionSideValue()
}
type captionSideType string

func (self captionSideType) Rule() Rule   { return Rule(self) }
func (captionSideType) captionSideValue() {}

func SetCaptionSide(value captionSideValue) Rule {
	return "caption-side:" + value.Rule() + ";"
}

type fontDisplayValue interface {
	ruleable
	fontDisplayValue()
}
type fontDisplayType string

func (self fontDisplayType) Rule() Rule   { return Rule(self) }
func (fontDisplayType) fontDisplayValue() {}

func SetFontDisplay(value fontDisplayValue) Rule {
	return "font-display:" + value.Rule() + ";"
}

type fontWeightValue interface {
	ruleable
	fontWeightValue()
}
type fontWeightType string

func (self fontWeightType) Rule() Rule  { return Rule(self) }
func (fontWeightType) fontWeightValue() {}

func SetFontWeight(value fontWeightValue) Rule {
	return "font-weight:" + value.Rule() + ";"
}

type backgroundRepeatValue interface {
	ruleable
	backgroundRepeatValue()
}
type backgroundRepeatType string

func (self backgroundRepeatType) Rule() Rule        { return Rule(self) }
func (backgroundRepeatType) backgroundRepeatValue() {}

func SetBackgroundRepeat(value backgroundRepeatValue) Rule {
	return "background-repeat:" + value.Rule() + ";"
}

type fontVariantValue interface {
	ruleable
	fontVariantValue()
}
type fontVariantType string

func (self fontVariantType) Rule() Rule   { return Rule(self) }
func (fontVariantType) fontVariantValue() {}

func SetFontVariant(value fontVariantValue) Rule {
	return "font-variant:" + value.Rule() + ";"
}
func SetGridRowStart(value gridStopValue) Rule {
	return "grid-row-start:" + value.Rule() + ";"
}

type marginValue interface {
	ruleable
	marginValue()
}
type marginType string

func (self marginType) Rule() Rule { return Rule(self) }
func (marginType) marginValue()    {}

func SetMargin(value marginValue) Rule {
	return "margin:" + value.Rule() + ";"
}

type clipValue interface {
	ruleable
	clipValue()
}
type clipType string

func (self clipType) Rule() Rule { return Rule(self) }
func (clipType) clipValue()      {}

func SetClip(value clipValue) Rule {
	return "clip:" + value.Rule() + ";"
}

type directionValue interface {
	ruleable
	directionValue()
}
type directionType string

func (self directionType) Rule() Rule { return Rule(self) }
func (directionType) directionValue() {}

func SetDirection(value directionValue) Rule {
	return "direction:" + value.Rule() + ";"
}

type fontSizeAdjustValue interface {
	ruleable
	fontSizeAdjustValue()
}
type fontSizeAdjustType string

func (self fontSizeAdjustType) Rule() Rule      { return Rule(self) }
func (fontSizeAdjustType) fontSizeAdjustValue() {}

func SetFontSizeAdjust(value fontSizeAdjustValue) Rule {
	return "font-size-adjust:" + value.Rule() + ";"
}

func SetTransformOrigin(p positionValue, z ...unitValue) Rule {
	if len(z) > 0 {
		return "transform-origin: " + p.Rule() + " " + z[0].Rule()
	} else {
		return "transform-origin: " + p.Rule()
	}
}

type animationFillModeValue interface {
	ruleable
	animationFillModeValue()
}
type animationFillModeType string

func (self animationFillModeType) Rule() Rule         { return Rule(self) }
func (animationFillModeType) animationFillModeValue() {}

func SetAnimationFillMode(value animationFillModeValue) Rule {
	return "animation-fill-mode:" + value.Rule() + ";"
}

func SetGridTemplateAreas(names []string) Rule {
	if len(names) == 0 {
		return "grid-template-areas: " + unitType("none").Rule() + ";"
	}
	var result string
	for _, name := range names {
		result += name + " "
	}
	return "grid-template-areas: " + unitType(result).Rule() + ";"
}

type textAlignValue interface {
	ruleable
	textAlignValue()
}
type textAlignType string

func (self textAlignType) Rule() Rule { return Rule(self) }
func (textAlignType) textAlignValue() {}

func SetTextAlign(value textAlignValue) Rule {
	return "text-align:" + value.Rule() + ";"
}
func SetAnimationDuration(value durationValue) Rule {
	return "animation-duration:" + value.Rule() + ";"
}

type animationIterationCountValue interface {
	ruleable
	animationIterationCountValue()
}
type animationIterationCountType string

func (self animationIterationCountType) Rule() Rule               { return Rule(self) }
func (animationIterationCountType) animationIterationCountValue() {}

func SetAnimationIterationCount(value animationIterationCountValue) Rule {
	return "animation-iteration-count:" + value.Rule() + ";"
}

type tableLayoutValue interface {
	ruleable
	tableLayoutValue()
}
type tableLayoutType string

func (self tableLayoutType) Rule() Rule   { return Rule(self) }
func (tableLayoutType) tableLayoutValue() {}

func SetTableLayout(value tableLayoutValue) Rule {
	return "table-layout:" + value.Rule() + ";"
}
func SetBorderRightWidth(value thicknessValue) Rule {
	return "border-right-width:" + value.Rule() + ";"
}

type columnsValue interface {
	ruleable
	columnsValue()
}
type columnsType string

func (self columnsType) Rule() Rule { return Rule(self) }
func (columnsType) columnsValue()   {}

func SetColumns(value columnsValue) Rule {
	return "columns:" + value.Rule() + ";"
}
func SetOutlineColor(value colorValue) Rule {
	return "outline-color:" + value.Rule() + ";"
}

type flexWrapValue interface {
	ruleable
	flexWrapValue()
}
type flexWrapType string

func (self flexWrapType) Rule() Rule { return Rule(self) }
func (flexWrapType) flexWrapValue()  {}

func SetFlexWrap(value flexWrapValue) Rule {
	return "flex-wrap:" + value.Rule() + ";"
}

type textAlignLastValue interface {
	ruleable
	textAlignLastValue()
}
type textAlignLastType string

func (self textAlignLastType) Rule() Rule     { return Rule(self) }
func (textAlignLastType) textAlignLastValue() {}

func SetTextAlignLast(value textAlignLastValue) Rule {
	return "text-align-last:" + value.Rule() + ";"
}

type allValue interface {
	ruleable
	allValue()
}
type allType string

func (self allType) Rule() Rule { return Rule(self) }
func (allType) allValue()       {}

func SetAll(value allValue) Rule {
	return "all:" + value.Rule() + ";"
}

type backfaceVisibilityValue interface {
	ruleable
	backfaceVisibilityValue()
}
type backfaceVisibilityType string

func (self backfaceVisibilityType) Rule() Rule          { return Rule(self) }
func (backfaceVisibilityType) backfaceVisibilityValue() {}

func SetBackfaceVisibility(value backfaceVisibilityValue) Rule {
	return "backface-visibility:" + value.Rule() + ";"
}

type fontVariantLigaturesValue interface {
	ruleable
	fontVariantLigaturesValue()
}
type fontVariantLigaturesType string

func (self fontVariantLigaturesType) Rule() Rule            { return Rule(self) }
func (fontVariantLigaturesType) fontVariantLigaturesValue() {}

func SetFontVariantLigatures(value fontVariantLigaturesValue) Rule {
	return "font-variant-ligatures:" + value.Rule() + ";"
}
func SetZIndex(value integerOrAutoValue) Rule {
	return "z-index:" + value.Rule() + ";"
}
func SetLetterSpacing(value normalOrUnitOrAutoValue) Rule {
	return "letter-spacing:" + value.Rule() + ";"
}
func SetMinHeight(value unitOrNoneValue) Rule {
	return "min-height:" + value.Rule() + ";"
}

type textJustifyValue interface {
	ruleable
	textJustifyValue()
}
type textJustifyType string

func (self textJustifyType) Rule() Rule   { return Rule(self) }
func (textJustifyType) textJustifyValue() {}

func SetTextJustify(value textJustifyValue) Rule {
	return "text-justify:" + value.Rule() + ";"
}
func SetOverflowY(value overflowValue) Rule {
	return "overflow-y:" + value.Rule() + ";"
}
func SetPaddingTop(value unitValue) Rule {
	return "padding-top:" + value.Rule() + ";"
}

type textDecorationStyleValue interface {
	ruleable
	textDecorationStyleValue()
}
type textDecorationStyleType string

func (self textDecorationStyleType) Rule() Rule           { return Rule(self) }
func (textDecorationStyleType) textDecorationStyleValue() {}

func SetTextDecorationStyle(value textDecorationStyleValue) Rule {
	return "text-decoration-style:" + value.Rule() + ";"
}

type clearValue interface {
	ruleable
	clearValue()
}
type clearType string

func (self clearType) Rule() Rule { return Rule(self) }
func (clearType) clearValue()     {}

func SetClear(value clearValue) Rule {
	return "clear:" + value.Rule() + ";"
}

type emptyCellsValue interface {
	ruleable
	emptyCellsValue()
}
type emptyCellsType string

func (self emptyCellsType) Rule() Rule  { return Rule(self) }
func (emptyCellsType) emptyCellsValue() {}

func SetEmptyCells(value emptyCellsValue) Rule {
	return "empty-cells:" + value.Rule() + ";"
}

type justifyContentValue interface {
	ruleable
	justifyContentValue()
}
type justifyContentType string

func (self justifyContentType) Rule() Rule      { return Rule(self) }
func (justifyContentType) justifyContentValue() {}

func SetJustifyContent(value justifyContentValue) Rule {
	return "justify-content:" + value.Rule() + ";"
}

type transitionValue interface {
	ruleable
	transitionValue()
}
type transitionType string

func (self transitionType) Rule() Rule  { return Rule(self) }
func (transitionType) transitionValue() {}

func SetTransition(value transitionValue) Rule {
	return "transition:" + value.Rule() + ";"
}

func SetTransitionProperty(properties ...interface{}) Rule {
	var names string

	/*for _, property := range properties {
		var s = NewStyle()
		reflect.ValueOf(property).Call([]reflect.Value{reflect.ValueOf(&s)})

		for i := range s.Stylable.(Implementation) {
			names += i
		}
	}*/

	return "transform-property: " + unitType(names).Rule()
}

type gridAreaValue interface {
	ruleable
	gridAreaValue()
}
type gridAreaType string

func (self gridAreaType) Rule() Rule { return Rule(self) }
func (gridAreaType) gridAreaValue()  {}

func SetGridArea(value gridAreaValue) Rule {
	return "grid-area:" + value.Rule() + ";"
}
func SetGridRowGap(value unitValue) Rule {
	return "grid-row-gap:" + value.Rule() + ";"
}
func SetObjectPosition(value unitAndUnitValue) Rule {
	return "object-position:" + value.Rule() + ";"
}
func SetTextShadow(value shadowValue) Rule {
	return "text-shadow:" + value.Rule() + ";"
}

type wordSpacingValue interface {
	ruleable
	wordSpacingValue()
}
type wordSpacingType string

func (self wordSpacingType) Rule() Rule   { return Rule(self) }
func (wordSpacingType) wordSpacingValue() {}

func SetWordSpacing(value wordSpacingValue) Rule {
	return "word-spacing:" + value.Rule() + ";"
}

type backgroundAttachmentValue interface {
	ruleable
	backgroundAttachmentValue()
}
type backgroundAttachmentType string

func (self backgroundAttachmentType) Rule() Rule            { return Rule(self) }
func (backgroundAttachmentType) backgroundAttachmentValue() {}

func SetBackgroundAttachment(value backgroundAttachmentValue) Rule {
	return "background-attachment:" + value.Rule() + ";"
}

type backgroundBlendModeValue interface {
	ruleable
	backgroundBlendModeValue()
}
type backgroundBlendModeType string

func (self backgroundBlendModeType) Rule() Rule           { return Rule(self) }
func (backgroundBlendModeType) backgroundBlendModeValue() {}

func SetBackgroundBlendMode(value backgroundBlendModeValue) Rule {
	return "background-blend-mode:" + value.Rule() + ";"
}

type flexDirectionValue interface {
	ruleable
	flexDirectionValue()
}
type flexDirectionType string

func (self flexDirectionType) Rule() Rule     { return Rule(self) }
func (flexDirectionType) flexDirectionValue() {}

func SetFlexDirection(value flexDirectionValue) Rule {
	return "flex-direction:" + value.Rule() + ";"
}

type columnRuleWidthValue interface {
	ruleable
	columnRuleWidthValue()
}
type columnRuleWidthType string

func (self columnRuleWidthType) Rule() Rule       { return Rule(self) }
func (columnRuleWidthType) columnRuleWidthValue() {}

func SetColumnRuleWidth(value columnRuleWidthValue) Rule {
	return "column-rule-width:" + value.Rule() + ";"
}

type fontLanguageOverrideValue interface {
	ruleable
	fontLanguageOverrideValue()
}
type fontLanguageOverrideType string

func (self fontLanguageOverrideType) Rule() Rule            { return Rule(self) }
func (fontLanguageOverrideType) fontLanguageOverrideValue() {}

func SetFontLanguageOverride(value fontLanguageOverrideValue) Rule {
	return "font-language-override:" + value.Rule() + ";"
}

type fontStretchValue interface {
	ruleable
	fontStretchValue()
}
type fontStretchType string

func (self fontStretchType) Rule() Rule   { return Rule(self) }
func (fontStretchType) fontStretchValue() {}

func SetFontStretch(value fontStretchValue) Rule {
	return "font-stretch:" + value.Rule() + ";"
}
func SetGridColumnGap(value unitValue) Rule {
	return "grid-column-gap:" + value.Rule() + ";"
}

type unicodeBidiValue interface {
	ruleable
	unicodeBidiValue()
}
type unicodeBidiType string

func (self unicodeBidiType) Rule() Rule   { return Rule(self) }
func (unicodeBidiType) unicodeBidiValue() {}

func SetUnicodeBidi(value unicodeBidiValue) Rule {
	return "unicode-bidi:" + value.Rule() + ";"
}

type borderBottomValue interface {
	ruleable
	borderBottomValue()
}
type borderBottomType string

func (self borderBottomType) Rule() Rule    { return Rule(self) }
func (borderBottomType) borderBottomValue() {}

func SetBorderBottom(value borderBottomValue) Rule {
	return "border-bottom:" + value.Rule() + ";"
}
func SetBorderImageSource(value imageValue) Rule {
	return "border-image-source:" + value.Rule() + ";"
}

type borderTopLeftRadiusValue interface {
	ruleable
	borderTopLeftRadiusValue()
}
type borderTopLeftRadiusType string

func (self borderTopLeftRadiusType) Rule() Rule           { return Rule(self) }
func (borderTopLeftRadiusType) borderTopLeftRadiusValue() {}

func SetBorderTopLeftRadius(value borderTopLeftRadiusValue) Rule {
	return "border-top-left-radius:" + value.Rule() + ";"
}

type columnWidthValue interface {
	ruleable
	columnWidthValue()
}
type columnWidthType string

func (self columnWidthType) Rule() Rule   { return Rule(self) }
func (columnWidthType) columnWidthValue() {}

func SetColumnWidth(value columnWidthValue) Rule {
	return "column-width:" + value.Rule() + ";"
}
func SetPageBreakAfter(value pageBreakValue) Rule {
	return "page-break-after:" + value.Rule() + ";"
}

type transformStyleValue interface {
	ruleable
	transformStyleValue()
}
type transformStyleType string

func (self transformStyleType) Rule() Rule      { return Rule(self) }
func (transformStyleType) transformStyleValue() {}

func SetTransformStyle(value transformStyleValue) Rule {
	return "transform-style:" + value.Rule() + ";"
}
func SetCounterReset(value nameValue) Rule {
	return "counter-reset:" + value.Rule() + ";"
}

type outlineValue interface {
	ruleable
	outlineValue()
}
type outlineType string

func (self outlineType) Rule() Rule { return Rule(self) }
func (outlineType) outlineValue()   {}

func SetOutline(value outlineValue) Rule {
	return "outline:" + value.Rule() + ";"
}
func SetTextIndent(value unitValue) Rule {
	return "text-indent:" + value.Rule() + ";"
}

type transitionTimingFunctionValue interface {
	ruleable
	transitionTimingFunctionValue()
}
type transitionTimingFunctionType string

func (self transitionTimingFunctionType) Rule() Rule                { return Rule(self) }
func (transitionTimingFunctionType) transitionTimingFunctionValue() {}

func SetTransitionTimingFunction(value transitionTimingFunctionValue) Rule {
	return "transition-timing-function:" + value.Rule() + ";"
}
func SetBackgroundSize(value sizeValue) Rule {
	return "background-size:" + value.Rule() + ";"
}
func SetBorderBottomColor(value colorValue) Rule {
	return "border-bottom-color:" + value.Rule() + ";"
}
func SetBorderRightStyle(value borderStyleValue) Rule {
	return "border-right-style:" + value.Rule() + ";"
}

type cursorValue interface {
	ruleable
	cursorValue()
}
type cursorType string

func (self cursorType) Rule() Rule { return Rule(self) }
func (cursorType) cursorValue()    {}

func SetCursor(value cursorValue) Rule {
	return "cursor:" + value.Rule() + ";"
}

func SetGridTemplateColumns(values []gridTemplateValue) Rule {
	if len(values) == 0 {
		return "grid-template-columns: " + unitType("none").Rule() + ";"
	}
	var result Rule
	for _, value := range values {
		result += value.Rule() + " "
	}
	return "grid-template-columns: " + unitType(result).Rule() + ";"
}

type lineHeightValue interface {
	ruleable
	lineHeightValue()
}
type lineHeightType string

func (self lineHeightType) Rule() Rule  { return Rule(self) }
func (lineHeightType) lineHeightValue() {}

func SetLineHeight(value lineHeightValue) Rule {
	return "line-height:" + value.Rule() + ";"
}

type textOrientationValue interface {
	ruleable
	textOrientationValue()
}
type textOrientationType string

func (self textOrientationType) Rule() Rule       { return Rule(self) }
func (textOrientationType) textOrientationValue() {}

func SetTextOrientation(value textOrientationValue) Rule {
	return "text-orientation:" + value.Rule() + ";"
}
func SetListStyleImage(value imageValue) Rule {
	return "list-style-image:" + value.Rule() + ";"
}

type alignContentValue interface {
	ruleable
	alignContentValue()
}
type alignContentType string

func (self alignContentType) Rule() Rule    { return Rule(self) }
func (alignContentType) alignContentValue() {}

func SetAlignContent(value alignContentValue) Rule {
	return "align-content:" + value.Rule() + ";"
}
func SetColumnRuleStyle(value borderStyleValue) Rule {
	return "column-rule-style:" + value.Rule() + ";"
}
func SetFlexGrow(value numberValue) Rule {
	return "flex-grow:" + value.Rule() + ";"
}

type fontVariantAlternatesValue interface {
	ruleable
	fontVariantAlternatesValue()
}
type fontVariantAlternatesType string

func (self fontVariantAlternatesType) Rule() Rule             { return Rule(self) }
func (fontVariantAlternatesType) fontVariantAlternatesValue() {}

func SetFontVariantAlternates(value fontVariantAlternatesValue) Rule {
	return "font-variant-alternates:" + value.Rule() + ";"
}

type listStylePositionValue interface {
	ruleable
	listStylePositionValue()
}
type listStylePositionType string

func (self listStylePositionType) Rule() Rule         { return Rule(self) }
func (listStylePositionType) listStylePositionValue() {}

func SetListStylePosition(value listStylePositionValue) Rule {
	return "list-style-position:" + value.Rule() + ";"
}

type pointerEventsValue interface {
	ruleable
	pointerEventsValue()
}
type pointerEventsType string

func (self pointerEventsType) Rule() Rule     { return Rule(self) }
func (pointerEventsType) pointerEventsValue() {}

func SetPointerEvents(value pointerEventsValue) Rule {
	return "pointer-events:" + value.Rule() + ";"
}

type verticalAlignValue interface {
	ruleable
	verticalAlignValue()
}
type verticalAlignType string

func (self verticalAlignType) Rule() Rule     { return Rule(self) }
func (verticalAlignType) verticalAlignValue() {}

func SetVerticalAlign(value verticalAlignValue) Rule {
	return "vertical-align:" + value.Rule() + ";"
}

type alignItemsValue interface {
	ruleable
	alignItemsValue()
}
type alignItemsType string

func (self alignItemsType) Rule() Rule  { return Rule(self) }
func (alignItemsType) alignItemsValue() {}

func SetAlignItems(value alignItemsValue) Rule {
	return "align-items:" + value.Rule() + ";"
}
func SetBorderSpacing(value unitValue) Rule {
	return "border-spacing:" + value.Rule() + ";"
}

type columnGapValue interface {
	ruleable
	columnGapValue()
}
type columnGapType string

func (self columnGapType) Rule() Rule { return Rule(self) }
func (columnGapType) columnGapValue() {}

func SetColumnGap(value columnGapValue) Rule {
	return "column-gap:" + value.Rule() + ";"
}

type textCombineUprightValue interface {
	ruleable
	textCombineUprightValue()
}
type textCombineUprightType string

func (self textCombineUprightType) Rule() Rule          { return Rule(self) }
func (textCombineUprightType) textCombineUprightValue() {}

func SetTextCombineUpright(value textCombineUprightValue) Rule {
	return "text-combine-upright:" + value.Rule() + ";"
}

type backgroundOriginValue interface {
	ruleable
	backgroundOriginValue()
}
type backgroundOriginType string

func (self backgroundOriginType) Rule() Rule        { return Rule(self) }
func (backgroundOriginType) backgroundOriginValue() {}

func SetBackgroundOrigin(value backgroundOriginValue) Rule {
	return "background-origin:" + value.Rule() + ";"
}

func SetGridTemplateRows(values []gridTemplateValue) Rule {
	if len(values) == 0 {
		return "grid-template-rows: " + unitType("none").Rule() + ";"
	}
	var result Rule
	for _, value := range values {
		result += value.Rule() + " "
	}
	return "grid-template-rows: " + unitType(result).Rule() + ";"
}

type overflowWrapValue interface {
	ruleable
	overflowWrapValue()
}
type overflowWrapType string

func (self overflowWrapType) Rule() Rule    { return Rule(self) }
func (overflowWrapType) overflowWrapValue() {}

func SetOverflowWrap(value overflowWrapValue) Rule {
	return "overflow-wrap:" + value.Rule() + ";"
}
func SetHeight(value unitOrAutoValue) Rule {
	return "height:" + value.Rule() + ";"
}

type orderValue interface {
	ruleable
	orderValue()
}
type orderType string

func (self orderType) Rule() Rule { return Rule(self) }
func (orderType) orderValue()     {}

func SetOrder(value orderValue) Rule {
	return "order:" + value.Rule() + ";"
}

type animationTimingFunctionValue interface {
	ruleable
	animationTimingFunctionValue()
}
type animationTimingFunctionType string

func (self animationTimingFunctionType) Rule() Rule               { return Rule(self) }
func (animationTimingFunctionType) animationTimingFunctionValue() {}

func SetAnimationTimingFunction(value animationTimingFunctionValue) Rule {
	return "animation-timing-function:" + value.Rule() + ";"
}

type fontFamilyValue interface {
	ruleable
	fontFamilyValue()
}
type fontFamilyType string

func (self fontFamilyType) Rule() Rule  { return Rule(self) }
func (fontFamilyType) fontFamilyValue() {}

func SetFontFamily(value fontFamilyValue) Rule {
	return "font-family:" + value.Rule() + ";"
}
func SetGridColumnStart(value gridStopValue) Rule {
	return "grid-column-start:" + value.Rule() + ";"
}
func SetPerspectiveOrigin(value unitAndUnitValue) Rule {
	return "perspective-origin:" + value.Rule() + ";"
}

type gridValue interface {
	ruleable
	gridValue()
}
type gridType string

func (self gridType) Rule() Rule { return Rule(self) }
func (gridType) gridValue()      {}

func SetGrid(value gridValue) Rule {
	return "grid:" + value.Rule() + ";"
}
func SetGridRowEnd(value gridStopValue) Rule {
	return "grid-row-end:" + value.Rule() + ";"
}
func SetOrphans(value uintValue) Rule {
	return "orphans:" + value.Rule() + ";"
}
func SetTextDecorationColor(value colorValue) Rule {
	return "text-decoration-color:" + value.Rule() + ";"
}
func SetBorderLeftColor(value sizeValue) Rule {
	return "border-left-color:" + value.Rule() + ";"
}

type borderRightValue interface {
	ruleable
	borderRightValue()
}
type borderRightType string

func (self borderRightType) Rule() Rule   { return Rule(self) }
func (borderRightType) borderRightValue() {}

func SetBorderRight(value borderRightValue) Rule {
	return "border-right:" + value.Rule() + ";"
}

type fontVariantNumericValue interface {
	ruleable
	fontVariantNumericValue()
}
type fontVariantNumericType string

func (self fontVariantNumericType) Rule() Rule          { return Rule(self) }
func (fontVariantNumericType) fontVariantNumericValue() {}

func SetFontVariantNumeric(value fontVariantNumericValue) Rule {
	return "font-variant-numeric:" + value.Rule() + ";"
}

type transitionDurationValue interface {
	ruleable
	transitionDurationValue()
}
type transitionDurationType string

func (self transitionDurationType) Rule() Rule          { return Rule(self) }
func (transitionDurationType) transitionDurationValue() {}

func SetTransitionDuration(value transitionDurationValue) Rule {
	return "transition-duration:" + value.Rule() + ";"
}

type columnCountValue interface {
	ruleable
	columnCountValue()
}
type columnCountType string

func (self columnCountType) Rule() Rule   { return Rule(self) }
func (columnCountType) columnCountValue() {}

func SetColumnCount(value columnCountValue) Rule {
	return "column-count:" + value.Rule() + ";"
}
func SetFlexBasis(value unitOrAutoValue) Rule {
	return "flex-basis:" + value.Rule() + ";"
}

type listStyleTypeValue interface {
	ruleable
	listStyleTypeValue()
}
type listStyleTypeType string

func (self listStyleTypeType) Rule() Rule     { return Rule(self) }
func (listStyleTypeType) listStyleTypeValue() {}

func SetListStyleType(value listStyleTypeValue) Rule {
	return "list-style-type:" + value.Rule() + ";"
}
func SetOutlineOffset(value unitValue) Rule {
	return "outline-offset:" + value.Rule() + ";"
}

type scrollBehaviorValue interface {
	ruleable
	scrollBehaviorValue()
}
type scrollBehaviorType string

func (self scrollBehaviorType) Rule() Rule      { return Rule(self) }
func (scrollBehaviorType) scrollBehaviorValue() {}

func SetScrollBehavior(value scrollBehaviorValue) Rule {
	return "scroll-behavior:" + value.Rule() + ";"
}
func SetBorderColor(value colorValue) Rule {
	return "border-color:" + value.Rule() + ";"
}
func SetBoxSizing(value boxValue) Rule {
	return "box-sizing:" + value.Rule() + ";"
}

type columnRuleValue interface {
	ruleable
	columnRuleValue()
}
type columnRuleType string

func (self columnRuleType) Rule() Rule  { return Rule(self) }
func (columnRuleType) columnRuleValue() {}

func SetColumnRule(value columnRuleValue) Rule {
	return "column-rule:" + value.Rule() + ";"
}
func SetBorderRightColor(value colorValue) Rule {
	return "border-right-color:" + value.Rule() + ";"
}
func SetMarginBottom(value unitOrAutoValue) Rule {
	return "margin-bottom:" + value.Rule() + ";"
}

type fontVariantCapsValue interface {
	ruleable
	fontVariantCapsValue()
}
type fontVariantCapsType string

func (self fontVariantCapsType) Rule() Rule       { return Rule(self) }
func (fontVariantCapsType) fontVariantCapsValue() {}

func SetFontVariantCaps(value fontVariantCapsValue) Rule {
	return "font-variant-caps:" + value.Rule() + ";"
}
func SetMaxHeight(value unitOrNoneValue) Rule {
	return "max-height:" + value.Rule() + ";"
}

type objectFitValue interface {
	ruleable
	objectFitValue()
}
type objectFitType string

func (self objectFitType) Rule() Rule { return Rule(self) }
func (objectFitType) objectFitValue() {}

func SetObjectFit(value objectFitValue) Rule {
	return "object-fit:" + value.Rule() + ";"
}

type hyphensValue interface {
	ruleable
	hyphensValue()
}
type hyphensType string

func (self hyphensType) Rule() Rule { return Rule(self) }
func (hyphensType) hyphensValue()   {}

func SetHyphens(value hyphensValue) Rule {
	return "hyphens:" + value.Rule() + ";"
}

type listStyleValue interface {
	ruleable
	listStyleValue()
}
type listStyleType string

func (self listStyleType) Rule() Rule { return Rule(self) }
func (listStyleType) listStyleValue() {}

func SetListStyle(value listStyleValue) Rule {
	return "list-style:" + value.Rule() + ";"
}

type fontVariantPositionValue interface {
	ruleable
	fontVariantPositionValue()
}
type fontVariantPositionType string

func (self fontVariantPositionType) Rule() Rule           { return Rule(self) }
func (fontVariantPositionType) fontVariantPositionValue() {}

func SetFontVariantPosition(value fontVariantPositionValue) Rule {
	return "font-variant-position:" + value.Rule() + ";"
}
func SetColor(value colorValue) Rule {
	return "color:" + value.Rule() + ";"
}

type fontSynthesisValue string

func (f fontSynthesisValue) Rule() Rule {
	return Rule(f)
}

func FontSynthesis(weight, style bool) fontSynthesisValue {
	if !weight && !style {
		return fontSynthesisValue("none")
	}
	var result fontSynthesisValue
	if weight {
		result += fontSynthesisValue("weight ")
	}
	if style {
		result += fontSynthesisValue("style")
	}
	return result
}

func SetFontSynthesis(value fontSynthesisValue) Rule {
	return "font-synthesis: " + value.Rule() + ";"
}

func SetGridTemplate(value gridTemplateValue) Rule {
	return "grid-template:" + value.Rule() + ";"
}
func SetOverflow(value overflowValue) Rule {
	return "overflow:" + value.Rule() + ";"
}

type textDecorationLineValue interface {
	ruleable
	textDecorationLineValue()
}
type textDecorationLineType string

func (self textDecorationLineType) Rule() Rule          { return Rule(self) }
func (textDecorationLineType) textDecorationLineValue() {}

func SetTextDecorationLine(value textDecorationLineValue) Rule {
	return "text-decoration-line:" + value.Rule() + ";"
}
func SetBorderTopColor(value colorValue) Rule {
	return "border-top-color:" + value.Rule() + ";"
}

type fontVariantEastAsianValue interface {
	ruleable
	fontVariantEastAsianValue()
}
type fontVariantEastAsianType string

func (self fontVariantEastAsianType) Rule() Rule            { return Rule(self) }
func (fontVariantEastAsianType) fontVariantEastAsianValue() {}

func SetFontVariantEastAsian(value fontVariantEastAsianValue) Rule {
	return "font-variant-east-asian:" + value.Rule() + ";"
}
func SetRight(value unitOrAutoValue) Rule {
	return "right:" + value.Rule() + ";"
}
func SetBorderImageWidth(value sizeValue) Rule {
	return "border-image-width:" + value.Rule() + ";"
}
func SetMaxWidth(value unitOrNoneValue) Rule {
	return "max-width:" + value.Rule() + ";"
}

type textDecorationValue interface {
	ruleable
	textDecorationValue()
}
type textDecorationType string

func (self textDecorationType) Rule() Rule      { return Rule(self) }
func (textDecorationType) textDecorationValue() {}

func SetTextDecoration(value textDecorationValue) Rule {
	return "text-decoration:" + value.Rule() + ";"
}
func SetWhiteSpace(value uintValue) Rule {
	return "white-space:" + value.Rule() + ";"
}

type borderRadiusValue interface {
	ruleable
	borderRadiusValue()
}
type borderRadiusType string

func (self borderRadiusType) Rule() Rule    { return Rule(self) }
func (borderRadiusType) borderRadiusValue() {}

func SetBorderRadius(value borderRadiusValue) Rule {
	return "border-radius:" + value.Rule() + ";"
}
func SetMarginTop(value unitOrAutoValue) Rule {
	return "margin-top:" + value.Rule() + ";"
}

type transformValue interface {
	ruleable
	transformValue()
}
type transformType string

func (self transformType) Rule() Rule { return Rule(self) }
func (transformType) transformValue() {}

func SetTransform(value transformValue) Rule {
	return "transform:" + value.Rule() + ";"
}
func SetColumnRuleColor(value colorValue) Rule {
	return "column-rule-color:" + value.Rule() + ";"
}
func SetPaddingRight(value unitValue) Rule {
	return "padding-right:" + value.Rule() + ";"
}
func SetBackgroundPosition(value unitAndUnitValue) Rule {
	return "background-position:" + value.Rule() + ";"
}

func SetQuotes(quotes []string) Rule {
	if len(quotes) == 0 {
		return "quotes: " + unitType("none").Rule() + ";"
	}
	var result string
	for _, quote := range quotes {
		result += strconv.Quote(quote)
	}
	return "quotes: " + unitType(result).Rule() + ";"
}

func SetBoxShadow(value shadowValue) Rule {
	return "box-shadow:" + value.Rule() + ";"
}

type columnFillValue interface {
	ruleable
	columnFillValue()
}
type columnFillType string

func (self columnFillType) Rule() Rule  { return Rule(self) }
func (columnFillType) columnFillValue() {}

func SetColumnFill(value columnFillValue) Rule {
	return "column-fill:" + value.Rule() + ";"
}

type flexValue interface {
	ruleable
	flexValue()
}
type flexType string

func (self flexType) Rule() Rule { return Rule(self) }
func (flexType) flexValue()      {}

func SetFlex(value flexValue) Rule {
	return "flex:" + value.Rule() + ";"
}
func SetOverflowX(value overflowValue) Rule {
	return "overflow-x:" + value.Rule() + ";"
}
func SetTop(value unitOrAutoValue) Rule {
	return "top:" + value.Rule() + ";"
}
func SetBackgroundClip(value boxValue) Rule {
	return "background-clip:" + value.Rule() + ";"
}

type borderLeftValue interface {
	ruleable
	borderLeftValue()
}
type borderLeftType string

func (self borderLeftType) Rule() Rule  { return Rule(self) }
func (borderLeftType) borderLeftValue() {}

func SetBorderLeft(value borderLeftValue) Rule {
	return "border-left:" + value.Rule() + ";"
}
func SetBorderStyle(value borderStyleValue) Rule {
	return "border-style:" + value.Rule() + ";"
}
func SetFontKerning(value normalOrAutoValue) Rule {
	return "font-kerning:" + value.Rule() + ";"
}
func SetBackgroundColor(value colorValue) Rule {
	return "background-color:" + value.Rule() + ";"
}
func SetBreakBefore(value breakValue) Rule {
	return "break-before:" + value.Rule() + ";"
}

type counterIncrementValue interface {
	ruleable
	counterIncrementValue()
}
type counterIncrementType string

func (self counterIncrementType) Rule() Rule        { return Rule(self) }
func (counterIncrementType) counterIncrementValue() {}

func SetCounterIncrement(value counterIncrementValue) Rule {
	return "counter-increment:" + value.Rule() + ";"
}

type floatValue interface {
	ruleable
	floatValue()
}
type floatType string

func (self floatType) Rule() Rule { return Rule(self) }
func (floatType) floatValue()     {}

func SetFloat(value floatValue) Rule {
	return "float:" + value.Rule() + ";"
}

type mixBlendModeValue interface {
	ruleable
	mixBlendModeValue()
}
type mixBlendModeType string

func (self mixBlendModeType) Rule() Rule    { return Rule(self) }
func (mixBlendModeType) mixBlendModeValue() {}

func SetMixBlendMode(value mixBlendModeValue) Rule {
	return "mix-blend-mode:" + value.Rule() + ";"
}

type transitionDelayValue interface {
	ruleable
	transitionDelayValue()
}
type transitionDelayType string

func (self transitionDelayType) Rule() Rule       { return Rule(self) }
func (transitionDelayType) transitionDelayValue() {}

func SetTransitionDelay(value transitionDelayValue) Rule {
	return "transition-delay:" + value.Rule() + ";"
}

type lineBreakValue interface {
	ruleable
	lineBreakValue()
}
type lineBreakType string

func (self lineBreakType) Rule() Rule { return Rule(self) }
func (lineBreakType) lineBreakValue() {}

func SetLineBreak(value lineBreakValue) Rule {
	return "line-break:" + value.Rule() + ";"
}

type writingModeValue interface {
	ruleable
	writingModeValue()
}
type writingModeType string

func (self writingModeType) Rule() Rule   { return Rule(self) }
func (writingModeType) writingModeValue() {}

func SetWritingMode(value writingModeValue) Rule {
	return "writing-mode:" + value.Rule() + ";"
}

type fontStyleValue interface {
	ruleable
	fontStyleValue()
}
type fontStyleType string

func (self fontStyleType) Rule() Rule { return Rule(self) }
func (fontStyleType) fontStyleValue() {}

func SetFontStyle(value fontStyleValue) Rule {
	return "font-style:" + value.Rule() + ";"
}

type hangingPunctuationValue interface {
	ruleable
	hangingPunctuationValue()
}
type hangingPunctuationType string

func (self hangingPunctuationType) Rule() Rule          { return Rule(self) }
func (hangingPunctuationType) hangingPunctuationValue() {}

func SetHangingPunctuation(value hangingPunctuationValue) Rule {
	return "hanging-punctuation:" + value.Rule() + ";"
}
func SetLeft(value unitOrAutoValue) Rule {
	return "left:" + value.Rule() + ";"
}

type borderValue interface {
	ruleable
	borderValue()
}
type borderType string

func (self borderType) Rule() Rule { return Rule(self) }
func (borderType) borderValue()    {}

func SetBorder(value borderValue) Rule {
	return "border:" + value.Rule() + ";"
}
func SetBorderWidth(value thicknessValue) Rule {
	return "border-width:" + value.Rule() + ";"
}

type userSelectValue interface {
	ruleable
	userSelectValue()
}
type userSelectType string

func (self userSelectType) Rule() Rule  { return Rule(self) }
func (userSelectType) userSelectValue() {}

func SetUserSelect(value userSelectValue) Rule {
	return "user-select:" + value.Rule() + ";"
}

type animationValue interface {
	ruleable
	animationValue()
}
type animationType string

func (self animationType) Rule() Rule { return Rule(self) }
func (animationType) animationValue() {}

func SetAnimation(value animationValue) Rule {
	return "animation:" + value.Rule() + ";"
}

type backgroundValue interface {
	ruleable
	backgroundValue()
}
type backgroundType string

func (self backgroundType) Rule() Rule  { return Rule(self) }
func (backgroundType) backgroundValue() {}

func SetBackground(value backgroundValue) Rule {
	return "background:" + value.Rule() + ";"
}
func SetOutlineStyle(value borderStyleValue) Rule {
	return "outline-style:" + value.Rule() + ";"
}
func SetColumnSpan(value unitOrAutoValue) Rule {
	return "column-span:" + value.Rule() + ";"
}
func SetMarginLeft(value unitOrAutoValue) Rule {
	return "margin-left:" + value.Rule() + ";"
}

type alignSelfValue interface {
	ruleable
	alignSelfValue()
}
type alignSelfType string

func (self alignSelfType) Rule() Rule { return Rule(self) }
func (alignSelfType) alignSelfValue() {}

func SetAlignSelf(value alignSelfValue) Rule {
	return "align-self:" + value.Rule() + ";"
}

type borderImageSliceValue interface {
	ruleable
	borderImageSliceValue()
}
type borderImageSliceType string

func (self borderImageSliceType) Rule() Rule        { return Rule(self) }
func (borderImageSliceType) borderImageSliceValue() {}

func SetBorderImageSlice(value borderImageSliceValue) Rule {
	return "border-image-slice:" + value.Rule() + ";"
}
func SetBorderLeftWidth(value thicknessValue) Rule {
	return "border-left-width:" + value.Rule() + ";"
}

type animationPlayStateValue interface {
	ruleable
	animationPlayStateValue()
}
type animationPlayStateType string

func (self animationPlayStateType) Rule() Rule          { return Rule(self) }
func (animationPlayStateType) animationPlayStateValue() {}

func SetAnimationPlayState(value animationPlayStateValue) Rule {
	return "animation-play-state:" + value.Rule() + ";"
}
func SetBorderBottomLeftRadius(value unitValue) Rule {
	return "border-bottom-left-radius:" + value.Rule() + ";"
}

type gridAutoFlowValue interface {
	ruleable
	gridAutoFlowValue()
}
type gridAutoFlowType string

func (self gridAutoFlowType) Rule() Rule    { return Rule(self) }
func (gridAutoFlowType) gridAutoFlowValue() {}

func SetGridAutoFlow(value gridAutoFlowValue) Rule {
	return "grid-auto-flow:" + value.Rule() + ";"
}

type visibilityValue interface {
	ruleable
	visibilityValue()
}
type visibilityType string

func (self visibilityType) Rule() Rule  { return Rule(self) }
func (visibilityType) visibilityValue() {}

func SetVisibility(value visibilityValue) Rule {
	return "visibility:" + value.Rule() + ";"
}
func SetBorderBottomWidth(value thicknessValue) Rule {
	return "border-bottom-width:" + value.Rule() + ";"
}

type borderCollapseValue interface {
	ruleable
	borderCollapseValue()
}
type borderCollapseType string

func (self borderCollapseType) Rule() Rule      { return Rule(self) }
func (borderCollapseType) borderCollapseValue() {}

func SetBorderCollapse(value borderCollapseValue) Rule {
	return "border-collapse:" + value.Rule() + ";"
}
func SetTabSize(value numberValue) Rule {
	return "tab-size:" + value.Rule() + ";"
}

type isolationValue interface {
	ruleable
	isolationValue()
}
type isolationType string

func (self isolationType) Rule() Rule { return Rule(self) }
func (isolationType) isolationValue() {}

func SetIsolation(value isolationValue) Rule {
	return "isolation:" + value.Rule() + ";"
}
func SetMarginRight(value unitOrAutoValue) Rule {
	return "margin-right:" + value.Rule() + ";"
}

type paddingValue interface {
	ruleable
	paddingValue()
}
type paddingType string

func (self paddingType) Rule() Rule { return Rule(self) }
func (paddingType) paddingValue()   {}

func SetPadding(value paddingValue) Rule {
	return "padding:" + value.Rule() + ";"
}
func SetBorderTopWidth(value thicknessValue) Rule {
	return "border-top-width:" + value.Rule() + ";"
}

type resizeValue interface {
	ruleable
	resizeValue()
}
type resizeType string

func (self resizeType) Rule() Rule { return Rule(self) }
func (resizeType) resizeValue()    {}

func SetResize(value resizeValue) Rule {
	return "resize:" + value.Rule() + ";"
}

type textOverflowValue interface {
	ruleable
	textOverflowValue()
}
type textOverflowType string

func (self textOverflowType) Rule() Rule    { return Rule(self) }
func (textOverflowType) textOverflowValue() {}

func SetTextOverflow(value textOverflowValue) Rule {
	return "text-overflow:" + value.Rule() + ";"
}
func SetBottom(value unitOrAutoValue) Rule {
	return "bottom:" + value.Rule() + ";"
}
func SetPaddingLeft(value unitValue) Rule {
	return "padding-left:" + value.Rule() + ";"
}

type pageBreakInsideValue interface {
	ruleable
	pageBreakInsideValue()
}
type pageBreakInsideType string

func (self pageBreakInsideType) Rule() Rule       { return Rule(self) }
func (pageBreakInsideType) pageBreakInsideValue() {}

func SetPageBreakInside(value pageBreakInsideValue) Rule {
	return "page-break-inside:" + value.Rule() + ";"
}
func SetBorderTopStyle(value borderStyleValue) Rule {
	return "border-top-style:" + value.Rule() + ";"
}
func SetBreakAfter(value breakValue) Rule {
	return "break-after:" + value.Rule() + ";"
}
func SetGridColumnEnd(value gridStopValue) Rule {
	return "grid-column-end:" + value.Rule() + ";"
}

type fontFeatureSettingsValue interface {
	ruleable
	fontFeatureSettingsValue()
}
type fontFeatureSettingsType string

func (self fontFeatureSettingsType) Rule() Rule           { return Rule(self) }
func (fontFeatureSettingsType) fontFeatureSettingsValue() {}

func SetFontFeatureSettings(value fontFeatureSettingsValue) Rule {
	return "font-feature-settings:" + value.Rule() + ";"
}

type widowsValue interface {
	ruleable
	widowsValue()
}
type widowsType string

func (self widowsType) Rule() Rule { return Rule(self) }
func (widowsType) widowsValue()    {}

func SetWidows(value widowsValue) Rule {
	return "widows:" + value.Rule() + ";"
}

type wordWrapValue interface {
	ruleable
	wordWrapValue()
}
type wordWrapType string

func (self wordWrapType) Rule() Rule { return Rule(self) }
func (wordWrapType) wordWrapValue()  {}

func SetWordWrap(value wordWrapValue) Rule {
	return "word-wrap:" + value.Rule() + ";"
}

type positionValue interface {
	ruleable
	positionValue()
}
type positionType string

func (self positionType) Rule() Rule { return Rule(self) }
func (positionType) positionValue()  {}

func SetPosition(value positionValue) Rule {
	return "position:" + value.Rule() + ";"
}

type textUnderlinePositionValue interface {
	ruleable
	textUnderlinePositionValue()
}
type textUnderlinePositionType string

func (self textUnderlinePositionType) Rule() Rule             { return Rule(self) }
func (textUnderlinePositionType) textUnderlinePositionValue() {}

func SetTextUnderlinePosition(value textUnderlinePositionValue) Rule {
	return "text-underline-position:" + value.Rule() + ";"
}
func SetBorderBottomStyle(value borderStyleValue) Rule {
	return "border-bottom-style:" + value.Rule() + ";"
}

type borderImageValue interface {
	ruleable
	borderImageValue()
}
type borderImageType string

func (self borderImageType) Rule() Rule   { return Rule(self) }
func (borderImageType) borderImageValue() {}

func SetBorderImage(value borderImageValue) Rule {
	return "border-image:" + value.Rule() + ";"
}
func SetBorderImageOutset(value uintOrUnitValue) Rule {
	return "border-image-outset:" + value.Rule() + ";"
}
func SetBackgroundImage(value imageValue) Rule {
	return "background-image:" + value.Rule() + ";"
}
func SetOpacity(value numberValue) Rule {
	return "opacity:" + value.Rule() + ";"
}
func SetPaddingBottom(value unitValue) Rule {
	return "padding-bottom:" + value.Rule() + ";"
}
func SetPageBreakBefore(value pageBreakValue) Rule {
	return "page-break-before:" + value.Rule() + ";"
}

type textTransformValue interface {
	ruleable
	textTransformValue()
}
type textTransformType string

func (self textTransformType) Rule() Rule     { return Rule(self) }
func (textTransformType) textTransformValue() {}

func SetTextTransform(value textTransformValue) Rule {
	return "text-transform:" + value.Rule() + ";"
}
func SetAnimationDelay(value durationValue) Rule {
	return "animation-delay:" + value.Rule() + ";"
}
func SetBorderLeftStyle(value borderStyleValue) Rule {
	return "border-left-style:" + value.Rule() + ";"
}

type fontSizeValue interface {
	ruleable
	fontSizeValue()
}
type fontSizeType string

func (self fontSizeType) Rule() Rule { return Rule(self) }
func (fontSizeType) fontSizeValue()  {}

func SetFontSize(value fontSizeValue) Rule {
	return "font-size:" + value.Rule() + ";"
}

type borderTopValue interface {
	ruleable
	borderTopValue()
}
type borderTopType string

func (self borderTopType) Rule() Rule { return Rule(self) }
func (borderTopType) borderTopValue() {}

func SetBorderTop(value borderTopValue) Rule {
	return "border-top:" + value.Rule() + ";"
}
func SetPerspective(value unitOrNoneValue) Rule {
	return "perspective:" + value.Rule() + ";"
}

type displayValue interface {
	ruleable
	displayValue()
}
type displayType string

func (self displayType) Rule() Rule { return Rule(self) }
func (displayType) displayValue()   {}

func SetDisplay(value displayValue) Rule {
	return "display:" + value.Rule() + ";"
}
func SetGridAutoColumns(value gridAutoValue) Rule {
	return "grid-auto-columns:" + value.Rule() + ";"
}

type breakInsideValue interface {
	ruleable
	breakInsideValue()
}
type breakInsideType string

func (self breakInsideType) Rule() Rule   { return Rule(self) }
func (breakInsideType) breakInsideValue() {}

func SetBreakInside(value breakInsideValue) Rule {
	return "break-inside:" + value.Rule() + ";"
}
func SetOutlineWidth(value thicknessValue) Rule {
	return "outline-width:" + value.Rule() + ";"
}
