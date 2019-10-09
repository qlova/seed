/*This file is computer-generated*/
package css

import "strconv"
import "reflect"

type gridAutoValue interface {
	stringable
	gridAutoValue()
}
type gridAutoType string

func (self gridAutoType) String() string { return string(self) }
func (gridAutoType) gridAutoValue()      {}

type imageValue interface {
	stringable
	imageValue()
}
type imageType string

func (self imageType) String() string { return string(self) }
func (imageType) imageValue()         {}

type durationValue interface {
	stringable
	durationValue()
}
type durationType string

func (self durationType) String() string { return string(self) }
func (durationType) durationValue()      {}

type thicknessValue interface {
	stringable
	thicknessValue()
}
type thicknessType string

func (self thicknessType) String() string { return string(self) }
func (thicknessType) thicknessValue()     {}

type colorValue interface {
	stringable
	colorValue()
}
type colorType string

func (self colorType) String() string { return string(self) }
func (colorType) colorValue()         {}

type unitOrAutoValue interface {
	stringable
	unitOrAutoValue()
}
type unitOrAutoType string

func (self unitOrAutoType) String() string { return string(self) }
func (unitOrAutoType) unitOrAutoValue()    {}

type gridStopValue interface {
	stringable
	gridStopValue()
}
type gridStopType string

func (self gridStopType) String() string { return string(self) }
func (gridStopType) gridStopValue()      {}

type sizeValue interface {
	stringable
	sizeValue()
}
type sizeType string

func (self sizeType) String() string { return string(self) }
func (sizeType) sizeValue()          {}

type unitValue interface {
	stringable
	unitValue()
}
type unitType string

func (self unitType) String() string { return string(self) }
func (unitType) unitValue()          {}

type gridTemplateValue interface {
	stringable
	gridTemplateValue()
}
type gridTemplateType string

func (self gridTemplateType) String() string { return string(self) }
func (gridTemplateType) gridTemplateValue()  {}

type breakValue interface {
	stringable
	breakValue()
}
type breakType string

func (self breakType) String() string { return string(self) }
func (breakType) breakValue()         {}

type borderStyleValue interface {
	stringable
	borderStyleValue()
}
type borderStyleType string

func (self borderStyleType) String() string { return string(self) }
func (borderStyleType) borderStyleValue()   {}

type uintValue interface {
	stringable
	uintValue()
}
type uintType string

func (self uintType) String() string { return string(self) }
func (uintType) uintValue()          {}

type shadowValue interface {
	stringable
	shadowValue()
}
type shadowType string

func (self shadowType) String() string { return string(self) }
func (shadowType) shadowValue()        {}

type numberValue interface {
	stringable
	numberValue()
}
type numberType string

func (self numberType) String() string { return string(self) }
func (numberType) numberValue()        {}

type integerOrAutoValue interface {
	stringable
	integerOrAutoValue()
}
type integerOrAutoType string

func (self integerOrAutoType) String() string { return string(self) }
func (integerOrAutoType) integerOrAutoValue() {}

type overflowValue interface {
	stringable
	overflowValue()
}
type overflowType string

func (self overflowType) String() string { return string(self) }
func (overflowType) overflowValue()      {}

type unitAndUnitValue interface {
	stringable
	unitAndUnitValue()
}
type unitAndUnitType string

func (self unitAndUnitType) String() string { return string(self) }
func (unitAndUnitType) unitAndUnitValue()   {}

type nameValue interface {
	stringable
	nameValue()
}
type nameType string

func (self nameType) String() string { return string(self) }
func (nameType) nameValue()          {}

type unitOrNoneValue interface {
	stringable
	unitOrNoneValue()
}
type unitOrNoneType string

func (self unitOrNoneType) String() string { return string(self) }
func (unitOrNoneType) unitOrNoneValue()    {}

type normalOrUnitOrAutoValue interface {
	stringable
	normalOrUnitOrAutoValue()
}
type normalOrUnitOrAutoType string

func (self normalOrUnitOrAutoType) String() string      { return string(self) }
func (normalOrUnitOrAutoType) normalOrUnitOrAutoValue() {}

type pageBreakValue interface {
	stringable
	pageBreakValue()
}
type pageBreakType string

func (self pageBreakType) String() string { return string(self) }
func (pageBreakType) pageBreakValue()     {}

type boxValue interface {
	stringable
	boxValue()
}
type boxType string

func (self boxType) String() string { return string(self) }
func (boxType) boxValue()           {}

type normalOrAutoValue interface {
	stringable
	normalOrAutoValue()
}
type normalOrAutoType string

func (self normalOrAutoType) String() string { return string(self) }
func (normalOrAutoType) normalOrAutoValue()  {}

type uintOrUnitValue interface {
	stringable
	uintOrUnitValue()
}
type uintOrUnitType string

func (self uintOrUnitType) String() string { return string(self) }
func (uintOrUnitType) uintOrUnitValue()    {}

func (style Style) GridAutoRows() gridAutoValue {
	return gridAutoType(style.Get("grid-auto-rows"))
}
func (style Style) SetGridAutoRows(value gridAutoValue) {
	style.set("grid-auto-rows", value)
}

type textDecorationStyleValue interface {
	stringable
	textDecorationStyleValue()
}
type textDecorationStyleType string

func (self textDecorationStyleType) String() string       { return string(self) }
func (textDecorationStyleType) textDecorationStyleValue() {}

func (style Style) TextDecorationStyle() textDecorationStyleValue {
	return textDecorationStyleType(style.Get("text-decoration-style"))
}
func (style Style) SetTextDecorationStyle(value textDecorationStyleValue) {
	style.set("text-decoration-style", value)
}

type transformStyleValue interface {
	stringable
	transformStyleValue()
}
type transformStyleType string

func (self transformStyleType) String() string  { return string(self) }
func (transformStyleType) transformStyleValue() {}

func (style Style) TransformStyle() transformStyleValue {
	return transformStyleType(style.Get("transform-style"))
}
func (style Style) SetTransformStyle(value transformStyleValue) {
	style.set("transform-style", value)
}

type alignSelfValue interface {
	stringable
	alignSelfValue()
}
type alignSelfType string

func (self alignSelfType) String() string { return string(self) }
func (alignSelfType) alignSelfValue()     {}

func (style Style) AlignSelf() alignSelfValue {
	return alignSelfType(style.Get("align-self"))
}
func (style Style) SetAlignSelf(value alignSelfValue) {
	style.set("align-self", value)
}

type flexFlowValue interface {
	stringable
	flexFlowValue()
}
type flexFlowType string

func (self flexFlowType) String() string { return string(self) }
func (flexFlowType) flexFlowValue()      {}

func (style Style) FlexFlow() flexFlowValue {
	return flexFlowType(style.Get("flex-flow"))
}
func (style Style) SetFlexFlow(value flexFlowValue) {
	style.set("flex-flow", value)
}

type fontVariantValue interface {
	stringable
	fontVariantValue()
}
type fontVariantType string

func (self fontVariantType) String() string { return string(self) }
func (fontVariantType) fontVariantValue()   {}

func (style Style) FontVariant() fontVariantValue {
	return fontVariantType(style.Get("font-variant"))
}
func (style Style) SetFontVariant(value fontVariantValue) {
	style.set("font-variant", value)
}
func (style Style) PaddingRight() unitValue {
	return unitType(style.Get("padding-right"))
}
func (style Style) SetPaddingRight(value unitValue) {
	style.set("padding-right", value)
}
func (style Style) PageBreakAfter() pageBreakValue {
	return pageBreakType(style.Get("page-break-after"))
}
func (style Style) SetPageBreakAfter(value pageBreakValue) {
	style.set("page-break-after", value)
}
func (style Style) Right() unitOrAutoValue {
	return unitOrAutoType(style.Get("right"))
}
func (style Style) SetRight(value unitOrAutoValue) {
	style.set("right", value)
}
func (style Style) BackgroundClip() boxValue {
	return boxType(style.Get("background-clip"))
}
func (style Style) SetBackgroundClip(value boxValue) {
	style.set("background-clip", value)
}

type columnRuleWidthValue interface {
	stringable
	columnRuleWidthValue()
}
type columnRuleWidthType string

func (self columnRuleWidthType) String() string   { return string(self) }
func (columnRuleWidthType) columnRuleWidthValue() {}

func (style Style) ColumnRuleWidth() columnRuleWidthValue {
	return columnRuleWidthType(style.Get("column-rule-width"))
}
func (style Style) SetColumnRuleWidth(value columnRuleWidthValue) {
	style.set("column-rule-width", value)
}

type fontDisplayValue interface {
	stringable
	fontDisplayValue()
}
type fontDisplayType string

func (self fontDisplayType) String() string { return string(self) }
func (fontDisplayType) fontDisplayValue()   {}

func (style Style) FontDisplay() fontDisplayValue {
	return fontDisplayType(style.Get("font-display"))
}
func (style Style) SetFontDisplay(value fontDisplayValue) {
	style.set("font-display", value)
}

type visibilityValue interface {
	stringable
	visibilityValue()
}
type visibilityType string

func (self visibilityType) String() string { return string(self) }
func (visibilityType) visibilityValue()    {}

func (style Style) Visibility() visibilityValue {
	return visibilityType(style.Get("visibility"))
}
func (style Style) SetVisibility(value visibilityValue) {
	style.set("visibility", value)
}
func (style Style) MarginBottom() unitOrAutoValue {
	return unitOrAutoType(style.Get("margin-bottom"))
}
func (style Style) SetMarginBottom(value unitOrAutoValue) {
	style.set("margin-bottom", value)
}
func (style Style) MarginRight() unitOrAutoValue {
	return unitOrAutoType(style.Get("margin-right"))
}
func (style Style) SetMarginRight(value unitOrAutoValue) {
	style.set("margin-right", value)
}

type transitionDelayValue interface {
	stringable
	transitionDelayValue()
}
type transitionDelayType string

func (self transitionDelayType) String() string   { return string(self) }
func (transitionDelayType) transitionDelayValue() {}

func (style Style) TransitionDelay() transitionDelayValue {
	return transitionDelayType(style.Get("transition-delay"))
}
func (style Style) SetTransitionDelay(value transitionDelayValue) {
	style.set("transition-delay", value)
}

type animationFillModeValue interface {
	stringable
	animationFillModeValue()
}
type animationFillModeType string

func (self animationFillModeType) String() string     { return string(self) }
func (animationFillModeType) animationFillModeValue() {}

func (style Style) AnimationFillMode() animationFillModeValue {
	return animationFillModeType(style.Get("animation-fill-mode"))
}
func (style Style) SetAnimationFillMode(value animationFillModeValue) {
	style.set("animation-fill-mode", value)
}
func (style Style) BorderImageOutset() uintOrUnitValue {
	return uintOrUnitType(style.Get("border-image-outset"))
}
func (style Style) SetBorderImageOutset(value uintOrUnitValue) {
	style.set("border-image-outset", value)
}

type directionValue interface {
	stringable
	directionValue()
}
type directionType string

func (self directionType) String() string { return string(self) }
func (directionType) directionValue()     {}

func (style Style) Direction() directionValue {
	return directionType(style.Get("direction"))
}
func (style Style) SetDirection(value directionValue) {
	style.set("direction", value)
}
func (style Style) MarginLeft() unitOrAutoValue {
	return unitOrAutoType(style.Get("margin-left"))
}
func (style Style) SetMarginLeft(value unitOrAutoValue) {
	style.set("margin-left", value)
}
func (style Style) PaddingBottom() unitValue {
	return unitType(style.Get("padding-bottom"))
}
func (style Style) SetPaddingBottom(value unitValue) {
	style.set("padding-bottom", value)
}

type flexValue interface {
	stringable
	flexValue()
}
type flexType string

func (self flexType) String() string { return string(self) }
func (flexType) flexValue()          {}

func (style Style) Flex() flexValue {
	return flexType(style.Get("flex"))
}
func (style Style) SetFlex(value flexValue) {
	style.set("flex", value)
}

type fontWeightValue interface {
	stringable
	fontWeightValue()
}
type fontWeightType string

func (self fontWeightType) String() string { return string(self) }
func (fontWeightType) fontWeightValue()    {}

func (style Style) FontWeight() fontWeightValue {
	return fontWeightType(style.Get("font-weight"))
}
func (style Style) SetFontWeight(value fontWeightValue) {
	style.set("font-weight", value)
}
func (style Style) Opacity() numberValue {
	return numberType(style.Get("opacity"))
}
func (style Style) SetOpacity(value numberValue) {
	style.set("opacity", value)
}

type textAlignValue interface {
	stringable
	textAlignValue()
}
type textAlignType string

func (self textAlignType) String() string { return string(self) }
func (textAlignType) textAlignValue()     {}

func (style Style) TextAlign() textAlignValue {
	return textAlignType(style.Get("text-align"))
}
func (style Style) SetTextAlign(value textAlignValue) {
	style.set("text-align", value)
}

type animationPlayStateValue interface {
	stringable
	animationPlayStateValue()
}
type animationPlayStateType string

func (self animationPlayStateType) String() string      { return string(self) }
func (animationPlayStateType) animationPlayStateValue() {}

func (style Style) AnimationPlayState() animationPlayStateValue {
	return animationPlayStateType(style.Get("animation-play-state"))
}
func (style Style) SetAnimationPlayState(value animationPlayStateValue) {
	style.set("animation-play-state", value)
}
func (style Style) BorderImageSource() imageValue {
	return imageType(style.Get("border-image-source"))
}
func (style Style) SetBorderImageSource(value imageValue) {
	style.set("border-image-source", value)
}

type borderRadiusValue interface {
	stringable
	borderRadiusValue()
}
type borderRadiusType string

func (self borderRadiusType) String() string { return string(self) }
func (borderRadiusType) borderRadiusValue()  {}

func (style Style) BorderRadius() borderRadiusValue {
	return borderRadiusType(style.Get("border-radius"))
}
func (style Style) SetBorderRadius(value borderRadiusValue) {
	style.set("border-radius", value)
}
func (style Style) GridAutoColumns() gridAutoValue {
	return gridAutoType(style.Get("grid-auto-columns"))
}
func (style Style) SetGridAutoColumns(value gridAutoValue) {
	style.set("grid-auto-columns", value)
}

type textDecorationLineValue interface {
	stringable
	textDecorationLineValue()
}
type textDecorationLineType string

func (self textDecorationLineType) String() string      { return string(self) }
func (textDecorationLineType) textDecorationLineValue() {}

func (style Style) TextDecorationLine() textDecorationLineValue {
	return textDecorationLineType(style.Get("text-decoration-line"))
}
func (style Style) SetTextDecorationLine(value textDecorationLineValue) {
	style.set("text-decoration-line", value)
}

func (style Style) SetWillChange(properties ...interface{}) {
	var names string

	for i, property := range properties {
		var s = NewStyle()
		var catcher = propertyCatcher("")
		s.Stylable = &catcher

		reflect.ValueOf(property).Call([]reflect.Value{reflect.ValueOf(&s)})

		names += *((*string)(s.Stylable.(*propertyCatcher)))
		if i != len(properties)-1 {
			names += ","
		}
	}
	style.set("will-change", unitType(names))
}

func (style Style) AnimationDelay() durationValue {
	return durationType(style.Get("animation-delay"))
}
func (style Style) SetAnimationDelay(value durationValue) {
	style.set("animation-delay", value)
}

type animationIterationCountValue interface {
	stringable
	animationIterationCountValue()
}
type animationIterationCountType string

func (self animationIterationCountType) String() string           { return string(self) }
func (animationIterationCountType) animationIterationCountValue() {}

func (style Style) AnimationIterationCount() animationIterationCountValue {
	return animationIterationCountType(style.Get("animation-iteration-count"))
}
func (style Style) SetAnimationIterationCount(value animationIterationCountValue) {
	style.set("animation-iteration-count", value)
}
func (style Style) BorderTopWidth() thicknessValue {
	return thicknessType(style.Get("border-top-width"))
}
func (style Style) SetBorderTopWidth(value thicknessValue) {
	style.set("border-top-width", value)
}
func (style Style) Color() colorValue {
	return colorType(style.Get("color"))
}
func (style Style) SetColor(value colorValue) {
	style.set("color", value)
}

type fontLanguageOverrideValue interface {
	stringable
	fontLanguageOverrideValue()
}
type fontLanguageOverrideType string

func (self fontLanguageOverrideType) String() string        { return string(self) }
func (fontLanguageOverrideType) fontLanguageOverrideValue() {}

func (style Style) FontLanguageOverride() fontLanguageOverrideValue {
	return fontLanguageOverrideType(style.Get("font-language-override"))
}
func (style Style) SetFontLanguageOverride(value fontLanguageOverrideValue) {
	style.set("font-language-override", value)
}

type allValue interface {
	stringable
	allValue()
}
type allType string

func (self allType) String() string { return string(self) }
func (allType) allValue()           {}

func (style Style) All() allValue {
	return allType(style.Get("all"))
}
func (style Style) SetAll(value allValue) {
	style.set("all", value)
}
func (style Style) ColumnSpan() unitOrAutoValue {
	return unitOrAutoType(style.Get("column-span"))
}
func (style Style) SetColumnSpan(value unitOrAutoValue) {
	style.set("column-span", value)
}

type fontFeatureSettingsValue interface {
	stringable
	fontFeatureSettingsValue()
}
type fontFeatureSettingsType string

func (self fontFeatureSettingsType) String() string       { return string(self) }
func (fontFeatureSettingsType) fontFeatureSettingsValue() {}

func (style Style) FontFeatureSettings() fontFeatureSettingsValue {
	return fontFeatureSettingsType(style.Get("font-feature-settings"))
}
func (style Style) SetFontFeatureSettings(value fontFeatureSettingsValue) {
	style.set("font-feature-settings", value)
}

type textCombineUprightValue interface {
	stringable
	textCombineUprightValue()
}
type textCombineUprightType string

func (self textCombineUprightType) String() string      { return string(self) }
func (textCombineUprightType) textCombineUprightValue() {}

func (style Style) TextCombineUpright() textCombineUprightValue {
	return textCombineUprightType(style.Get("text-combine-upright"))
}
func (style Style) SetTextCombineUpright(value textCombineUprightValue) {
	style.set("text-combine-upright", value)
}

type borderValue interface {
	stringable
	borderValue()
}
type borderType string

func (self borderType) String() string { return string(self) }
func (borderType) borderValue()        {}

func (style Style) Border() borderValue {
	return borderType(style.Get("border"))
}
func (style Style) SetBorder(value borderValue) {
	style.set("border", value)
}
func (style Style) GridRowEnd() gridStopValue {
	return gridStopType(style.Get("grid-row-end"))
}
func (style Style) SetGridRowEnd(value gridStopValue) {
	style.set("grid-row-end", value)
}

type backfaceVisibilityValue interface {
	stringable
	backfaceVisibilityValue()
}
type backfaceVisibilityType string

func (self backfaceVisibilityType) String() string      { return string(self) }
func (backfaceVisibilityType) backfaceVisibilityValue() {}

func (style Style) BackfaceVisibility() backfaceVisibilityValue {
	return backfaceVisibilityType(style.Get("backface-visibility"))
}
func (style Style) SetBackfaceVisibility(value backfaceVisibilityValue) {
	style.set("backface-visibility", value)
}
func (style Style) BackgroundSize() sizeValue {
	return sizeType(style.Get("background-size"))
}
func (style Style) SetBackgroundSize(value sizeValue) {
	style.set("background-size", value)
}

type breakInsideValue interface {
	stringable
	breakInsideValue()
}
type breakInsideType string

func (self breakInsideType) String() string { return string(self) }
func (breakInsideType) breakInsideValue()   {}

func (style Style) BreakInside() breakInsideValue {
	return breakInsideType(style.Get("break-inside"))
}
func (style Style) SetBreakInside(value breakInsideValue) {
	style.set("break-inside", value)
}

type clipValue interface {
	stringable
	clipValue()
}
type clipType string

func (self clipType) String() string { return string(self) }
func (clipType) clipValue()          {}

func (style Style) Clip() clipValue {
	return clipType(style.Get("clip"))
}
func (style Style) SetClip(value clipValue) {
	style.set("clip", value)
}
func (style Style) BorderBottomLeftRadius() unitValue {
	return unitType(style.Get("border-bottom-left-radius"))
}
func (style Style) SetBorderBottomLeftRadius(value unitValue) {
	style.set("border-bottom-left-radius", value)
}

type fontValue interface {
	stringable
	fontValue()
}
type fontType string

func (self fontType) String() string { return string(self) }
func (fontType) fontValue()          {}

func (style Style) Font() fontValue {
	return fontType(style.Get("font"))
}
func (style Style) SetFont(value fontValue) {
	style.set("font", value)
}
func (style Style) GridTemplate() gridTemplateValue {
	return gridTemplateType(style.Get("grid-template"))
}
func (style Style) SetGridTemplate(value gridTemplateValue) {
	style.set("grid-template", value)
}

type orderValue interface {
	stringable
	orderValue()
}
type orderType string

func (self orderType) String() string { return string(self) }
func (orderType) orderValue()         {}

func (style Style) Order() orderValue {
	return orderType(style.Get("order"))
}
func (style Style) SetOrder(value orderValue) {
	style.set("order", value)
}

type textAlignLastValue interface {
	stringable
	textAlignLastValue()
}
type textAlignLastType string

func (self textAlignLastType) String() string { return string(self) }
func (textAlignLastType) textAlignLastValue() {}

func (style Style) TextAlignLast() textAlignLastValue {
	return textAlignLastType(style.Get("text-align-last"))
}
func (style Style) SetTextAlignLast(value textAlignLastValue) {
	style.set("text-align-last", value)
}
func (style Style) AnimationDuration() durationValue {
	return durationType(style.Get("animation-duration"))
}
func (style Style) SetAnimationDuration(value durationValue) {
	style.set("animation-duration", value)
}
func (style Style) BorderBottomColor() colorValue {
	return colorType(style.Get("border-bottom-color"))
}
func (style Style) SetBorderBottomColor(value colorValue) {
	style.set("border-bottom-color", value)
}
func (style Style) BreakBefore() breakValue {
	return breakType(style.Get("break-before"))
}
func (style Style) SetBreakBefore(value breakValue) {
	style.set("break-before", value)
}

type columnWidthValue interface {
	stringable
	columnWidthValue()
}
type columnWidthType string

func (self columnWidthType) String() string { return string(self) }
func (columnWidthType) columnWidthValue()   {}

func (style Style) ColumnWidth() columnWidthValue {
	return columnWidthType(style.Get("column-width"))
}
func (style Style) SetColumnWidth(value columnWidthValue) {
	style.set("column-width", value)
}

type wordWrapValue interface {
	stringable
	wordWrapValue()
}
type wordWrapType string

func (self wordWrapType) String() string { return string(self) }
func (wordWrapType) wordWrapValue()      {}

func (style Style) WordWrap() wordWrapValue {
	return wordWrapType(style.Get("word-wrap"))
}
func (style Style) SetWordWrap(value wordWrapValue) {
	style.set("word-wrap", value)
}

type alignItemsValue interface {
	stringable
	alignItemsValue()
}
type alignItemsType string

func (self alignItemsType) String() string { return string(self) }
func (alignItemsType) alignItemsValue()    {}

func (style Style) AlignItems() alignItemsValue {
	return alignItemsType(style.Get("align-items"))
}
func (style Style) SetAlignItems(value alignItemsValue) {
	style.set("align-items", value)
}
func (style Style) BorderLeftColor() sizeValue {
	return sizeType(style.Get("border-left-color"))
}
func (style Style) SetBorderLeftColor(value sizeValue) {
	style.set("border-left-color", value)
}
func (style Style) Bottom() unitOrAutoValue {
	return unitOrAutoType(style.Get("bottom"))
}
func (style Style) SetBottom(value unitOrAutoValue) {
	style.set("bottom", value)
}

type fontSizeAdjustValue interface {
	stringable
	fontSizeAdjustValue()
}
type fontSizeAdjustType string

func (self fontSizeAdjustType) String() string  { return string(self) }
func (fontSizeAdjustType) fontSizeAdjustValue() {}

func (style Style) FontSizeAdjust() fontSizeAdjustValue {
	return fontSizeAdjustType(style.Get("font-size-adjust"))
}
func (style Style) SetFontSizeAdjust(value fontSizeAdjustValue) {
	style.set("font-size-adjust", value)
}
func (style Style) BorderRightStyle() borderStyleValue {
	return borderStyleType(style.Get("border-right-style"))
}
func (style Style) SetBorderRightStyle(value borderStyleValue) {
	style.set("border-right-style", value)
}

type fontStyleValue interface {
	stringable
	fontStyleValue()
}
type fontStyleType string

func (self fontStyleType) String() string { return string(self) }
func (fontStyleType) fontStyleValue()     {}

func (style Style) FontStyle() fontStyleValue {
	return fontStyleType(style.Get("font-style"))
}
func (style Style) SetFontStyle(value fontStyleValue) {
	style.set("font-style", value)
}

type isolationValue interface {
	stringable
	isolationValue()
}
type isolationType string

func (self isolationType) String() string { return string(self) }
func (isolationType) isolationValue()     {}

func (style Style) Isolation() isolationValue {
	return isolationType(style.Get("isolation"))
}
func (style Style) SetIsolation(value isolationValue) {
	style.set("isolation", value)
}
func (style Style) ListStyleImage() imageValue {
	return imageType(style.Get("list-style-image"))
}
func (style Style) SetListStyleImage(value imageValue) {
	style.set("list-style-image", value)
}
func (style Style) BorderImageWidth() sizeValue {
	return sizeType(style.Get("border-image-width"))
}
func (style Style) SetBorderImageWidth(value sizeValue) {
	style.set("border-image-width", value)
}
func (style Style) BorderLeftWidth() thicknessValue {
	return thicknessType(style.Get("border-left-width"))
}
func (style Style) SetBorderLeftWidth(value thicknessValue) {
	style.set("border-left-width", value)
}

type borderTopLeftRadiusValue interface {
	stringable
	borderTopLeftRadiusValue()
}
type borderTopLeftRadiusType string

func (self borderTopLeftRadiusType) String() string       { return string(self) }
func (borderTopLeftRadiusType) borderTopLeftRadiusValue() {}

func (style Style) BorderTopLeftRadius() borderTopLeftRadiusValue {
	return borderTopLeftRadiusType(style.Get("border-top-left-radius"))
}
func (style Style) SetBorderTopLeftRadius(value borderTopLeftRadiusValue) {
	style.set("border-top-left-radius", value)
}

type columnGapValue interface {
	stringable
	columnGapValue()
}
type columnGapType string

func (self columnGapType) String() string { return string(self) }
func (columnGapType) columnGapValue()     {}

func (style Style) ColumnGap() columnGapValue {
	return columnGapType(style.Get("column-gap"))
}
func (style Style) SetColumnGap(value columnGapValue) {
	style.set("column-gap", value)
}

type transitionDurationValue interface {
	stringable
	transitionDurationValue()
}
type transitionDurationType string

func (self transitionDurationType) String() string      { return string(self) }
func (transitionDurationType) transitionDurationValue() {}

func (style Style) TransitionDuration() transitionDurationValue {
	return transitionDurationType(style.Get("transition-duration"))
}
func (style Style) SetTransitionDuration(value transitionDurationValue) {
	style.set("transition-duration", value)
}
func (style Style) BorderWidth() thicknessValue {
	return thicknessType(style.Get("border-width"))
}
func (style Style) SetBorderWidth(value thicknessValue) {
	style.set("border-width", value)
}

type tableLayoutValue interface {
	stringable
	tableLayoutValue()
}
type tableLayoutType string

func (self tableLayoutType) String() string { return string(self) }
func (tableLayoutType) tableLayoutValue()   {}

func (style Style) TableLayout() tableLayoutValue {
	return tableLayoutType(style.Get("table-layout"))
}
func (style Style) SetTableLayout(value tableLayoutValue) {
	style.set("table-layout", value)
}

type textTransformValue interface {
	stringable
	textTransformValue()
}
type textTransformType string

func (self textTransformType) String() string { return string(self) }
func (textTransformType) textTransformValue() {}

func (style Style) TextTransform() textTransformValue {
	return textTransformType(style.Get("text-transform"))
}
func (style Style) SetTextTransform(value textTransformValue) {
	style.set("text-transform", value)
}

type backgroundRepeatValue interface {
	stringable
	backgroundRepeatValue()
}
type backgroundRepeatType string

func (self backgroundRepeatType) String() string    { return string(self) }
func (backgroundRepeatType) backgroundRepeatValue() {}

func (style Style) BackgroundRepeat() backgroundRepeatValue {
	return backgroundRepeatType(style.Get("background-repeat"))
}
func (style Style) SetBackgroundRepeat(value backgroundRepeatValue) {
	style.set("background-repeat", value)
}

type resizeValue interface {
	stringable
	resizeValue()
}
type resizeType string

func (self resizeType) String() string { return string(self) }
func (resizeType) resizeValue()        {}

func (style Style) Resize() resizeValue {
	return resizeType(style.Get("resize"))
}
func (style Style) SetResize(value resizeValue) {
	style.set("resize", value)
}

type textUnderlinePositionValue interface {
	stringable
	textUnderlinePositionValue()
}
type textUnderlinePositionType string

func (self textUnderlinePositionType) String() string         { return string(self) }
func (textUnderlinePositionType) textUnderlinePositionValue() {}

func (style Style) TextUnderlinePosition() textUnderlinePositionValue {
	return textUnderlinePositionType(style.Get("text-underline-position"))
}
func (style Style) SetTextUnderlinePosition(value textUnderlinePositionValue) {
	style.set("text-underline-position", value)
}
func (style Style) OutlineOffset() unitValue {
	return unitType(style.Get("outline-offset"))
}
func (style Style) SetOutlineOffset(value unitValue) {
	style.set("outline-offset", value)
}

type textOverflowValue interface {
	stringable
	textOverflowValue()
}
type textOverflowType string

func (self textOverflowType) String() string { return string(self) }
func (textOverflowType) textOverflowValue()  {}

func (style Style) TextOverflow() textOverflowValue {
	return textOverflowType(style.Get("text-overflow"))
}
func (style Style) SetTextOverflow(value textOverflowValue) {
	style.set("text-overflow", value)
}

type flexDirectionValue interface {
	stringable
	flexDirectionValue()
}
type flexDirectionType string

func (self flexDirectionType) String() string { return string(self) }
func (flexDirectionType) flexDirectionValue() {}

func (style Style) FlexDirection() flexDirectionValue {
	return flexDirectionType(style.Get("flex-direction"))
}
func (style Style) SetFlexDirection(value flexDirectionValue) {
	style.set("flex-direction", value)
}

type fontVariantEastAsianValue interface {
	stringable
	fontVariantEastAsianValue()
}
type fontVariantEastAsianType string

func (self fontVariantEastAsianType) String() string        { return string(self) }
func (fontVariantEastAsianType) fontVariantEastAsianValue() {}

func (style Style) FontVariantEastAsian() fontVariantEastAsianValue {
	return fontVariantEastAsianType(style.Get("font-variant-east-asian"))
}
func (style Style) SetFontVariantEastAsian(value fontVariantEastAsianValue) {
	style.set("font-variant-east-asian", value)
}

type gridAreaValue interface {
	stringable
	gridAreaValue()
}
type gridAreaType string

func (self gridAreaType) String() string { return string(self) }
func (gridAreaType) gridAreaValue()      {}

func (style Style) GridArea() gridAreaValue {
	return gridAreaType(style.Get("grid-area"))
}
func (style Style) SetGridArea(value gridAreaValue) {
	style.set("grid-area", value)
}
func (style Style) BorderRightColor() colorValue {
	return colorType(style.Get("border-right-color"))
}
func (style Style) SetBorderRightColor(value colorValue) {
	style.set("border-right-color", value)
}
func (style Style) BorderSpacing() unitValue {
	return unitType(style.Get("border-spacing"))
}
func (style Style) SetBorderSpacing(value unitValue) {
	style.set("border-spacing", value)
}

type columnsValue interface {
	stringable
	columnsValue()
}
type columnsType string

func (self columnsType) String() string { return string(self) }
func (columnsType) columnsValue()       {}

func (style Style) Columns() columnsValue {
	return columnsType(style.Get("columns"))
}
func (style Style) SetColumns(value columnsValue) {
	style.set("columns", value)
}

type imageRenderingValue interface {
	stringable
	imageRenderingValue()
}
type imageRenderingType string

func (self imageRenderingType) String() string  { return string(self) }
func (imageRenderingType) imageRenderingValue() {}

func (style Style) ImageRendering() imageRenderingValue {
	return imageRenderingType(style.Get("image-rendering"))
}
func (style Style) SetImageRendering(value imageRenderingValue) {
	style.set("image-rendering", value)
}
func (style Style) WhiteSpace() uintValue {
	return uintType(style.Get("white-space"))
}
func (style Style) SetWhiteSpace(value uintValue) {
	style.set("white-space", value)
}

type borderTopRightRadiusValue interface {
	stringable
	borderTopRightRadiusValue()
}
type borderTopRightRadiusType string

func (self borderTopRightRadiusType) String() string        { return string(self) }
func (borderTopRightRadiusType) borderTopRightRadiusValue() {}

func (style Style) BorderTopRightRadius() borderTopRightRadiusValue {
	return borderTopRightRadiusType(style.Get("border-top-right-radius"))
}
func (style Style) SetBorderTopRightRadius(value borderTopRightRadiusValue) {
	style.set("border-top-right-radius", value)
}
func (style Style) GridColumnEnd() gridStopValue {
	return gridStopType(style.Get("grid-column-end"))
}
func (style Style) SetGridColumnEnd(value gridStopValue) {
	style.set("grid-column-end", value)
}
func (style Style) BorderBottomWidth() thicknessValue {
	return thicknessType(style.Get("border-bottom-width"))
}
func (style Style) SetBorderBottomWidth(value thicknessValue) {
	style.set("border-bottom-width", value)
}
func (style Style) Height() unitOrAutoValue {
	return unitOrAutoType(style.Get("height"))
}
func (style Style) SetHeight(value unitOrAutoValue) {
	style.set("height", value)
}
func (style Style) TextShadow() shadowValue {
	return shadowType(style.Get("text-shadow"))
}
func (style Style) SetTextShadow(value shadowValue) {
	style.set("text-shadow", value)
}

type animationValue interface {
	stringable
	animationValue()
}
type animationType string

func (self animationType) String() string { return string(self) }
func (animationType) animationValue()     {}

func (style Style) Animation() animationValue {
	return animationType(style.Get("animation"))
}
func (style Style) SetAnimation(value animationValue) {
	style.set("animation", value)
}
func (style Style) FlexShrink() numberValue {
	return numberType(style.Get("flex-shrink"))
}
func (style Style) SetFlexShrink(value numberValue) {
	style.set("flex-shrink", value)
}

type fontVariantPositionValue interface {
	stringable
	fontVariantPositionValue()
}
type fontVariantPositionType string

func (self fontVariantPositionType) String() string       { return string(self) }
func (fontVariantPositionType) fontVariantPositionValue() {}

func (style Style) FontVariantPosition() fontVariantPositionValue {
	return fontVariantPositionType(style.Get("font-variant-position"))
}
func (style Style) SetFontVariantPosition(value fontVariantPositionValue) {
	style.set("font-variant-position", value)
}

type gridAutoFlowValue interface {
	stringable
	gridAutoFlowValue()
}
type gridAutoFlowType string

func (self gridAutoFlowType) String() string { return string(self) }
func (gridAutoFlowType) gridAutoFlowValue()  {}

func (style Style) GridAutoFlow() gridAutoFlowValue {
	return gridAutoFlowType(style.Get("grid-auto-flow"))
}
func (style Style) SetGridAutoFlow(value gridAutoFlowValue) {
	style.set("grid-auto-flow", value)
}
func (style Style) BoxShadow() shadowValue {
	return shadowType(style.Get("box-shadow"))
}
func (style Style) SetBoxShadow(value shadowValue) {
	style.set("box-shadow", value)
}

type listStyleTypeValue interface {
	stringable
	listStyleTypeValue()
}
type listStyleTypeType string

func (self listStyleTypeType) String() string { return string(self) }
func (listStyleTypeType) listStyleTypeValue() {}

func (style Style) ListStyleType() listStyleTypeValue {
	return listStyleTypeType(style.Get("list-style-type"))
}
func (style Style) SetListStyleType(value listStyleTypeValue) {
	style.set("list-style-type", value)
}
func (style Style) ZIndex() integerOrAutoValue {
	return integerOrAutoType(style.Get("z-index"))
}
func (style Style) SetZIndex(value integerOrAutoValue) {
	style.set("z-index", value)
}

type borderTopValue interface {
	stringable
	borderTopValue()
}
type borderTopType string

func (self borderTopType) String() string { return string(self) }
func (borderTopType) borderTopValue()     {}

func (style Style) BorderTop() borderTopValue {
	return borderTopType(style.Get("border-top"))
}
func (style Style) SetBorderTop(value borderTopValue) {
	style.set("border-top", value)
}

type fontVariantLigaturesValue interface {
	stringable
	fontVariantLigaturesValue()
}
type fontVariantLigaturesType string

func (self fontVariantLigaturesType) String() string        { return string(self) }
func (fontVariantLigaturesType) fontVariantLigaturesValue() {}

func (style Style) FontVariantLigatures() fontVariantLigaturesValue {
	return fontVariantLigaturesType(style.Get("font-variant-ligatures"))
}
func (style Style) SetFontVariantLigatures(value fontVariantLigaturesValue) {
	style.set("font-variant-ligatures", value)
}

type lineHeightValue interface {
	stringable
	lineHeightValue()
}
type lineHeightType string

func (self lineHeightType) String() string { return string(self) }
func (lineHeightType) lineHeightValue()    {}

func (style Style) LineHeight() lineHeightValue {
	return lineHeightType(style.Get("line-height"))
}
func (style Style) SetLineHeight(value lineHeightValue) {
	style.set("line-height", value)
}

type verticalAlignValue interface {
	stringable
	verticalAlignValue()
}
type verticalAlignType string

func (self verticalAlignType) String() string { return string(self) }
func (verticalAlignType) verticalAlignValue() {}

func (style Style) VerticalAlign() verticalAlignValue {
	return verticalAlignType(style.Get("vertical-align"))
}
func (style Style) SetVerticalAlign(value verticalAlignValue) {
	style.set("vertical-align", value)
}
func (style Style) Width() unitOrAutoValue {
	return unitOrAutoType(style.Get("width"))
}
func (style Style) SetWidth(value unitOrAutoValue) {
	style.set("width", value)
}
func (style Style) PerspectiveOrigin() unitAndUnitValue {
	return unitAndUnitType(style.Get("perspective-origin"))
}
func (style Style) SetPerspectiveOrigin(value unitAndUnitValue) {
	style.set("perspective-origin", value)
}

type textDecorationValue interface {
	stringable
	textDecorationValue()
}
type textDecorationType string

func (self textDecorationType) String() string  { return string(self) }
func (textDecorationType) textDecorationValue() {}

func (style Style) TextDecoration() textDecorationValue {
	return textDecorationType(style.Get("text-decoration"))
}
func (style Style) SetTextDecoration(value textDecorationValue) {
	style.set("text-decoration", value)
}

type borderRightValue interface {
	stringable
	borderRightValue()
}
type borderRightType string

func (self borderRightType) String() string { return string(self) }
func (borderRightType) borderRightValue()   {}

func (style Style) BorderRight() borderRightValue {
	return borderRightType(style.Get("border-right"))
}
func (style Style) SetBorderRight(value borderRightValue) {
	style.set("border-right", value)
}
func (style Style) BorderTopColor() colorValue {
	return colorType(style.Get("border-top-color"))
}
func (style Style) SetBorderTopColor(value colorValue) {
	style.set("border-top-color", value)
}
func (style Style) CounterReset() nameValue {
	return nameType(style.Get("counter-reset"))
}
func (style Style) SetCounterReset(value nameValue) {
	style.set("counter-reset", value)
}

type filterValue interface {
	stringable
	filterValue()
}
type filterType string

func (self filterType) String() string { return string(self) }
func (filterType) filterValue()        {}

func (style Style) Filter() filterValue {
	return filterType(style.Get("filter"))
}
func (style Style) SetFilter(value filterValue) {
	style.set("filter", value)
}
func (style Style) OverflowX() overflowValue {
	return overflowType(style.Get("overflow-x"))
}
func (style Style) SetOverflowX(value overflowValue) {
	style.set("overflow-x", value)
}
func (style Style) BorderBottomStyle() borderStyleValue {
	return borderStyleType(style.Get("border-bottom-style"))
}
func (style Style) SetBorderBottomStyle(value borderStyleValue) {
	style.set("border-bottom-style", value)
}

type fontVariantNumericValue interface {
	stringable
	fontVariantNumericValue()
}
type fontVariantNumericType string

func (self fontVariantNumericType) String() string      { return string(self) }
func (fontVariantNumericType) fontVariantNumericValue() {}

func (style Style) FontVariantNumeric() fontVariantNumericValue {
	return fontVariantNumericType(style.Get("font-variant-numeric"))
}
func (style Style) SetFontVariantNumeric(value fontVariantNumericValue) {
	style.set("font-variant-numeric", value)
}

type gridRowValue interface {
	stringable
	gridRowValue()
}
type gridRowType string

func (self gridRowType) String() string { return string(self) }
func (gridRowType) gridRowValue()       {}

func (style Style) GridRow() gridRowValue {
	return gridRowType(style.Get("grid-row"))
}
func (style Style) SetGridRow(value gridRowValue) {
	style.set("grid-row", value)
}
func (style Style) Perspective() unitOrNoneValue {
	return unitOrNoneType(style.Get("perspective"))
}
func (style Style) SetPerspective(value unitOrNoneValue) {
	style.set("perspective", value)
}
func (style Style) Top() unitOrAutoValue {
	return unitOrAutoType(style.Get("top"))
}
func (style Style) SetTop(value unitOrAutoValue) {
	style.set("top", value)
}
func (style Style) FlexBasis() unitOrAutoValue {
	return unitOrAutoType(style.Get("flex-basis"))
}
func (style Style) SetFlexBasis(value unitOrAutoValue) {
	style.set("flex-basis", value)
}

type borderImageRepeatValue interface {
	stringable
	borderImageRepeatValue()
}
type borderImageRepeatType string

func (self borderImageRepeatType) String() string     { return string(self) }
func (borderImageRepeatType) borderImageRepeatValue() {}

func (style Style) BorderImageRepeat() borderImageRepeatValue {
	return borderImageRepeatType(style.Get("border-image-repeat"))
}
func (style Style) SetBorderImageRepeat(value borderImageRepeatValue) {
	style.set("border-image-repeat", value)
}

type emptyCellsValue interface {
	stringable
	emptyCellsValue()
}
type emptyCellsType string

func (self emptyCellsType) String() string { return string(self) }
func (emptyCellsType) emptyCellsValue()    {}

func (style Style) EmptyCells() emptyCellsValue {
	return emptyCellsType(style.Get("empty-cells"))
}
func (style Style) SetEmptyCells(value emptyCellsValue) {
	style.set("empty-cells", value)
}

type gridGapValue interface {
	stringable
	gridGapValue()
}
type gridGapType string

func (self gridGapType) String() string { return string(self) }
func (gridGapType) gridGapValue()       {}

func (style Style) GridGap() gridGapValue {
	return gridGapType(style.Get("grid-gap"))
}
func (style Style) SetGridGap(value gridGapValue) {
	style.set("grid-gap", value)
}

type alignContentValue interface {
	stringable
	alignContentValue()
}
type alignContentType string

func (self alignContentType) String() string { return string(self) }
func (alignContentType) alignContentValue()  {}

func (style Style) AlignContent() alignContentValue {
	return alignContentType(style.Get("align-content"))
}
func (style Style) SetAlignContent(value alignContentValue) {
	style.set("align-content", value)
}

type positionValue interface {
	stringable
	positionValue()
}
type positionType string

func (self positionType) String() string { return string(self) }
func (positionType) positionValue()      {}

func (style Style) Position() positionValue {
	return positionType(style.Get("position"))
}
func (style Style) SetPosition(value positionValue) {
	style.set("position", value)
}

type widowsValue interface {
	stringable
	widowsValue()
}
type widowsType string

func (self widowsType) String() string { return string(self) }
func (widowsType) widowsValue()        {}

func (style Style) Widows() widowsValue {
	return widowsType(style.Get("widows"))
}
func (style Style) SetWidows(value widowsValue) {
	style.set("widows", value)
}
func (style Style) BackgroundImage() imageValue {
	return imageType(style.Get("background-image"))
}
func (style Style) SetBackgroundImage(value imageValue) {
	style.set("background-image", value)
}

type borderBottomValue interface {
	stringable
	borderBottomValue()
}
type borderBottomType string

func (self borderBottomType) String() string { return string(self) }
func (borderBottomType) borderBottomValue()  {}

func (style Style) BorderBottom() borderBottomValue {
	return borderBottomType(style.Get("border-bottom"))
}
func (style Style) SetBorderBottom(value borderBottomValue) {
	style.set("border-bottom", value)
}

type userSelectValue interface {
	stringable
	userSelectValue()
}
type userSelectType string

func (self userSelectType) String() string { return string(self) }
func (userSelectType) userSelectValue()    {}

func (style Style) UserSelect() userSelectValue {
	return userSelectType(style.Get("user-select"))
}
func (style Style) SetUserSelect(value userSelectValue) {
	style.set("user-select", value)
}
func (style Style) MaxWidth() unitOrNoneValue {
	return unitOrNoneType(style.Get("max-width"))
}
func (style Style) SetMaxWidth(value unitOrNoneValue) {
	style.set("max-width", value)
}

type outlineValue interface {
	stringable
	outlineValue()
}
type outlineType string

func (self outlineType) String() string { return string(self) }
func (outlineType) outlineValue()       {}

func (style Style) Outline() outlineValue {
	return outlineType(style.Get("outline"))
}
func (style Style) SetOutline(value outlineValue) {
	style.set("outline", value)
}

func (style Style) SetQuotes(quotes []string) {
	if len(quotes) == 0 {
		style.set("quotes", unitType("none"))
		return
	}
	var result string
	for _, quote := range quotes {
		result += strconv.Quote(quote)
	}
	style.set("quotes", unitType(result))
}

func (style Style) MarginTop() unitOrAutoValue {
	return unitOrAutoType(style.Get("margin-top"))
}
func (style Style) SetMarginTop(value unitOrAutoValue) {
	style.set("margin-top", value)
}
func (style Style) MinHeight() unitOrNoneValue {
	return unitOrNoneType(style.Get("min-height"))
}
func (style Style) SetMinHeight(value unitOrNoneValue) {
	style.set("min-height", value)
}
func (style Style) LetterSpacing() normalOrUnitOrAutoValue {
	return normalOrUnitOrAutoType(style.Get("letter-spacing"))
}
func (style Style) SetLetterSpacing(value normalOrUnitOrAutoValue) {
	style.set("letter-spacing", value)
}

type displayValue interface {
	stringable
	displayValue()
}
type displayType string

func (self displayType) String() string { return string(self) }
func (displayType) displayValue()       {}

func (style Style) Display() displayValue {
	return displayType(style.Get("display"))
}
func (style Style) SetDisplay(value displayValue) {
	style.set("display", value)
}
func (style Style) TextDecorationColor() colorValue {
	return colorType(style.Get("text-decoration-color"))
}
func (style Style) SetTextDecorationColor(value colorValue) {
	style.set("text-decoration-color", value)
}

func (style Style) SetTransitionProperty(properties ...interface{}) {
	var names string

	for _, property := range properties {
		var s = NewStyle()
		reflect.ValueOf(property).Call([]reflect.Value{reflect.ValueOf(&s)})

		for i := range s.Stylable.(Implementation) {
			names += i
		}
	}
	style.set("transform-property", unitType(names))
}

type backgroundAttachmentValue interface {
	stringable
	backgroundAttachmentValue()
}
type backgroundAttachmentType string

func (self backgroundAttachmentType) String() string        { return string(self) }
func (backgroundAttachmentType) backgroundAttachmentValue() {}

func (style Style) BackgroundAttachment() backgroundAttachmentValue {
	return backgroundAttachmentType(style.Get("background-attachment"))
}
func (style Style) SetBackgroundAttachment(value backgroundAttachmentValue) {
	style.set("background-attachment", value)
}
func (style Style) BorderColor() colorValue {
	return colorType(style.Get("border-color"))
}
func (style Style) SetBorderColor(value colorValue) {
	style.set("border-color", value)
}
func (style Style) BorderRightWidth() thicknessValue {
	return thicknessType(style.Get("border-right-width"))
}
func (style Style) SetBorderRightWidth(value thicknessValue) {
	style.set("border-right-width", value)
}
func (style Style) ColumnRuleStyle() borderStyleValue {
	return borderStyleType(style.Get("column-rule-style"))
}
func (style Style) SetColumnRuleStyle(value borderStyleValue) {
	style.set("column-rule-style", value)
}

type hangingPunctuationValue interface {
	stringable
	hangingPunctuationValue()
}
type hangingPunctuationType string

func (self hangingPunctuationType) String() string      { return string(self) }
func (hangingPunctuationType) hangingPunctuationValue() {}

func (style Style) HangingPunctuation() hangingPunctuationValue {
	return hangingPunctuationType(style.Get("hanging-punctuation"))
}
func (style Style) SetHangingPunctuation(value hangingPunctuationValue) {
	style.set("hanging-punctuation", value)
}

type fontStretchValue interface {
	stringable
	fontStretchValue()
}
type fontStretchType string

func (self fontStretchType) String() string { return string(self) }
func (fontStretchType) fontStretchValue()   {}

func (style Style) FontStretch() fontStretchValue {
	return fontStretchType(style.Get("font-stretch"))
}
func (style Style) SetFontStretch(value fontStretchValue) {
	style.set("font-stretch", value)
}

type gridColumnValue interface {
	stringable
	gridColumnValue()
}
type gridColumnType string

func (self gridColumnType) String() string { return string(self) }
func (gridColumnType) gridColumnValue()    {}

func (style Style) GridColumn() gridColumnValue {
	return gridColumnType(style.Get("grid-column"))
}
func (style Style) SetGridColumn(value gridColumnValue) {
	style.set("grid-column", value)
}

type listStyleValue interface {
	stringable
	listStyleValue()
}
type listStyleType string

func (self listStyleType) String() string { return string(self) }
func (listStyleType) listStyleValue()     {}

func (style Style) ListStyle() listStyleValue {
	return listStyleType(style.Get("list-style"))
}
func (style Style) SetListStyle(value listStyleValue) {
	style.set("list-style", value)
}
func (style Style) TabSize() numberValue {
	return numberType(style.Get("tab-size"))
}
func (style Style) SetTabSize(value numberValue) {
	style.set("tab-size", value)
}

type borderImageSliceValue interface {
	stringable
	borderImageSliceValue()
}
type borderImageSliceType string

func (self borderImageSliceType) String() string    { return string(self) }
func (borderImageSliceType) borderImageSliceValue() {}

func (style Style) BorderImageSlice() borderImageSliceValue {
	return borderImageSliceType(style.Get("border-image-slice"))
}
func (style Style) SetBorderImageSlice(value borderImageSliceValue) {
	style.set("border-image-slice", value)
}

type captionSideValue interface {
	stringable
	captionSideValue()
}
type captionSideType string

func (self captionSideType) String() string { return string(self) }
func (captionSideType) captionSideValue()   {}

func (style Style) CaptionSide() captionSideValue {
	return captionSideType(style.Get("caption-side"))
}
func (style Style) SetCaptionSide(value captionSideValue) {
	style.set("caption-side", value)
}
func (style Style) OutlineStyle() borderStyleValue {
	return borderStyleType(style.Get("outline-style"))
}
func (style Style) SetOutlineStyle(value borderStyleValue) {
	style.set("outline-style", value)
}

type pointerEventsValue interface {
	stringable
	pointerEventsValue()
}
type pointerEventsType string

func (self pointerEventsType) String() string { return string(self) }
func (pointerEventsType) pointerEventsValue() {}

func (style Style) PointerEvents() pointerEventsValue {
	return pointerEventsType(style.Get("pointer-events"))
}
func (style Style) SetPointerEvents(value pointerEventsValue) {
	style.set("pointer-events", value)
}

type animationDirectionValue interface {
	stringable
	animationDirectionValue()
}
type animationDirectionType string

func (self animationDirectionType) String() string      { return string(self) }
func (animationDirectionType) animationDirectionValue() {}

func (style Style) AnimationDirection() animationDirectionValue {
	return animationDirectionType(style.Get("animation-direction"))
}
func (style Style) SetAnimationDirection(value animationDirectionValue) {
	style.set("animation-direction", value)
}

type animationTimingFunctionValue interface {
	stringable
	animationTimingFunctionValue()
}
type animationTimingFunctionType string

func (self animationTimingFunctionType) String() string           { return string(self) }
func (animationTimingFunctionType) animationTimingFunctionValue() {}

func (style Style) AnimationTimingFunction() animationTimingFunctionValue {
	return animationTimingFunctionType(style.Get("animation-timing-function"))
}
func (style Style) SetAnimationTimingFunction(value animationTimingFunctionValue) {
	style.set("animation-timing-function", value)
}

type lineBreakValue interface {
	stringable
	lineBreakValue()
}
type lineBreakType string

func (self lineBreakType) String() string { return string(self) }
func (lineBreakType) lineBreakValue()     {}

func (style Style) LineBreak() lineBreakValue {
	return lineBreakType(style.Get("line-break"))
}
func (style Style) SetLineBreak(value lineBreakValue) {
	style.set("line-break", value)
}
func (style Style) OverflowY() overflowValue {
	return overflowType(style.Get("overflow-y"))
}
func (style Style) SetOverflowY(value overflowValue) {
	style.set("overflow-y", value)
}
func (style Style) PageBreakBefore() pageBreakValue {
	return pageBreakType(style.Get("page-break-before"))
}
func (style Style) SetPageBreakBefore(value pageBreakValue) {
	style.set("page-break-before", value)
}

type pageBreakInsideValue interface {
	stringable
	pageBreakInsideValue()
}
type pageBreakInsideType string

func (self pageBreakInsideType) String() string   { return string(self) }
func (pageBreakInsideType) pageBreakInsideValue() {}

func (style Style) PageBreakInside() pageBreakInsideValue {
	return pageBreakInsideType(style.Get("page-break-inside"))
}
func (style Style) SetPageBreakInside(value pageBreakInsideValue) {
	style.set("page-break-inside", value)
}

type textOrientationValue interface {
	stringable
	textOrientationValue()
}
type textOrientationType string

func (self textOrientationType) String() string   { return string(self) }
func (textOrientationType) textOrientationValue() {}

func (style Style) TextOrientation() textOrientationValue {
	return textOrientationType(style.Get("text-orientation"))
}
func (style Style) SetTextOrientation(value textOrientationValue) {
	style.set("text-orientation", value)
}

type borderCollapseValue interface {
	stringable
	borderCollapseValue()
}
type borderCollapseType string

func (self borderCollapseType) String() string  { return string(self) }
func (borderCollapseType) borderCollapseValue() {}

func (style Style) BorderCollapse() borderCollapseValue {
	return borderCollapseType(style.Get("border-collapse"))
}
func (style Style) SetBorderCollapse(value borderCollapseValue) {
	style.set("border-collapse", value)
}

type boxDecorationBreakValue interface {
	stringable
	boxDecorationBreakValue()
}
type boxDecorationBreakType string

func (self boxDecorationBreakType) String() string      { return string(self) }
func (boxDecorationBreakType) boxDecorationBreakValue() {}

func (style Style) BoxDecorationBreak() boxDecorationBreakValue {
	return boxDecorationBreakType(style.Get("box-decoration-break"))
}
func (style Style) SetBoxDecorationBreak(value boxDecorationBreakValue) {
	style.set("box-decoration-break", value)
}

type fontFamilyValue interface {
	stringable
	fontFamilyValue()
}
type fontFamilyType string

func (self fontFamilyType) String() string { return string(self) }
func (fontFamilyType) fontFamilyValue()    {}

func (style Style) FontFamily() fontFamilyValue {
	return fontFamilyType(style.Get("font-family"))
}
func (style Style) SetFontFamily(value fontFamilyValue) {
	style.set("font-family", value)
}
func (style Style) Left() unitOrAutoValue {
	return unitOrAutoType(style.Get("left"))
}
func (style Style) SetLeft(value unitOrAutoValue) {
	style.set("left", value)
}

type mixBlendModeValue interface {
	stringable
	mixBlendModeValue()
}
type mixBlendModeType string

func (self mixBlendModeType) String() string { return string(self) }
func (mixBlendModeType) mixBlendModeValue()  {}

func (style Style) MixBlendMode() mixBlendModeValue {
	return mixBlendModeType(style.Get("mix-blend-mode"))
}
func (style Style) SetMixBlendMode(value mixBlendModeValue) {
	style.set("mix-blend-mode", value)
}

type wordBreakValue interface {
	stringable
	wordBreakValue()
}
type wordBreakType string

func (self wordBreakType) String() string { return string(self) }
func (wordBreakType) wordBreakValue()     {}

func (style Style) WordBreak() wordBreakValue {
	return wordBreakType(style.Get("word-break"))
}
func (style Style) SetWordBreak(value wordBreakValue) {
	style.set("word-break", value)
}
func (style Style) BoxSizing() boxValue {
	return boxType(style.Get("box-sizing"))
}
func (style Style) SetBoxSizing(value boxValue) {
	style.set("box-sizing", value)
}

type contentValue interface {
	stringable
	contentValue()
}
type contentType string

func (self contentType) String() string { return string(self) }
func (contentType) contentValue()       {}

func (style Style) Content() contentValue {
	return contentType(style.Get("content"))
}
func (style Style) SetContent(value contentValue) {
	style.set("content", value)
}

func (style Style) SetGridTemplateRows(values []gridTemplateValue) {
	if len(values) == 0 {
		style.set("grid-template-rows", unitType("none"))
		return
	}
	var result string
	for _, value := range values {
		result += value.String() + " "
	}
	style.set("grid-template-rows", unitType(result))
}

func (style Style) SetTransformOrigin(p positionValue, z ...unitValue) {
	if len(z) > 0 {
		style.set("transform-origin", unitType(p.String()+" "+z[0].String()))
	} else {
		style.set("transform-origin", p)
	}
}

type counterIncrementValue interface {
	stringable
	counterIncrementValue()
}
type counterIncrementType string

func (self counterIncrementType) String() string    { return string(self) }
func (counterIncrementType) counterIncrementValue() {}

func (style Style) CounterIncrement() counterIncrementValue {
	return counterIncrementType(style.Get("counter-increment"))
}
func (style Style) SetCounterIncrement(value counterIncrementValue) {
	style.set("counter-increment", value)
}
func (style Style) FontKerning() normalOrAutoValue {
	return normalOrAutoType(style.Get("font-kerning"))
}
func (style Style) SetFontKerning(value normalOrAutoValue) {
	style.set("font-kerning", value)
}

type hyphensValue interface {
	stringable
	hyphensValue()
}
type hyphensType string

func (self hyphensType) String() string { return string(self) }
func (hyphensType) hyphensValue()       {}

func (style Style) Hyphens() hyphensValue {
	return hyphensType(style.Get("hyphens"))
}
func (style Style) SetHyphens(value hyphensValue) {
	style.set("hyphens", value)
}

type overflowWrapValue interface {
	stringable
	overflowWrapValue()
}
type overflowWrapType string

func (self overflowWrapType) String() string { return string(self) }
func (overflowWrapType) overflowWrapValue()  {}

func (style Style) OverflowWrap() overflowWrapValue {
	return overflowWrapType(style.Get("overflow-wrap"))
}
func (style Style) SetOverflowWrap(value overflowWrapValue) {
	style.set("overflow-wrap", value)
}

type wordSpacingValue interface {
	stringable
	wordSpacingValue()
}
type wordSpacingType string

func (self wordSpacingType) String() string { return string(self) }
func (wordSpacingType) wordSpacingValue()   {}

func (style Style) WordSpacing() wordSpacingValue {
	return wordSpacingType(style.Get("word-spacing"))
}
func (style Style) SetWordSpacing(value wordSpacingValue) {
	style.set("word-spacing", value)
}

type writingModeValue interface {
	stringable
	writingModeValue()
}
type writingModeType string

func (self writingModeType) String() string { return string(self) }
func (writingModeType) writingModeValue()   {}

func (style Style) WritingMode() writingModeValue {
	return writingModeType(style.Get("writing-mode"))
}
func (style Style) SetWritingMode(value writingModeValue) {
	style.set("writing-mode", value)
}

type borderLeftValue interface {
	stringable
	borderLeftValue()
}
type borderLeftType string

func (self borderLeftType) String() string { return string(self) }
func (borderLeftType) borderLeftValue()    {}

func (style Style) BorderLeft() borderLeftValue {
	return borderLeftType(style.Get("border-left"))
}
func (style Style) SetBorderLeft(value borderLeftValue) {
	style.set("border-left", value)
}

type columnCountValue interface {
	stringable
	columnCountValue()
}
type columnCountType string

func (self columnCountType) String() string { return string(self) }
func (columnCountType) columnCountValue()   {}

func (style Style) ColumnCount() columnCountValue {
	return columnCountType(style.Get("column-count"))
}
func (style Style) SetColumnCount(value columnCountValue) {
	style.set("column-count", value)
}

type floatValue interface {
	stringable
	floatValue()
}
type floatType string

func (self floatType) String() string { return string(self) }
func (floatType) floatValue()         {}

func (style Style) Float() floatValue {
	return floatType(style.Get("float"))
}
func (style Style) SetFloat(value floatValue) {
	style.set("float", value)
}
func (style Style) MinWidth() unitOrNoneValue {
	return unitOrNoneType(style.Get("min-width"))
}
func (style Style) SetMinWidth(value unitOrNoneValue) {
	style.set("min-width", value)
}
func (style Style) OutlineWidth() thicknessValue {
	return thicknessType(style.Get("outline-width"))
}
func (style Style) SetOutlineWidth(value thicknessValue) {
	style.set("outline-width", value)
}

type borderImageValue interface {
	stringable
	borderImageValue()
}
type borderImageType string

func (self borderImageType) String() string { return string(self) }
func (borderImageType) borderImageValue()   {}

func (style Style) BorderImage() borderImageValue {
	return borderImageType(style.Get("border-image"))
}
func (style Style) SetBorderImage(value borderImageValue) {
	style.set("border-image", value)
}

type fontVariantAlternatesValue interface {
	stringable
	fontVariantAlternatesValue()
}
type fontVariantAlternatesType string

func (self fontVariantAlternatesType) String() string         { return string(self) }
func (fontVariantAlternatesType) fontVariantAlternatesValue() {}

func (style Style) FontVariantAlternates() fontVariantAlternatesValue {
	return fontVariantAlternatesType(style.Get("font-variant-alternates"))
}
func (style Style) SetFontVariantAlternates(value fontVariantAlternatesValue) {
	style.set("font-variant-alternates", value)
}
func (style Style) GridColumnGap() unitValue {
	return unitType(style.Get("grid-column-gap"))
}
func (style Style) SetGridColumnGap(value unitValue) {
	style.set("grid-column-gap", value)
}

type listStylePositionValue interface {
	stringable
	listStylePositionValue()
}
type listStylePositionType string

func (self listStylePositionType) String() string     { return string(self) }
func (listStylePositionType) listStylePositionValue() {}

func (style Style) ListStylePosition() listStylePositionValue {
	return listStylePositionType(style.Get("list-style-position"))
}
func (style Style) SetListStylePosition(value listStylePositionValue) {
	style.set("list-style-position", value)
}

type scrollBehaviorValue interface {
	stringable
	scrollBehaviorValue()
}
type scrollBehaviorType string

func (self scrollBehaviorType) String() string  { return string(self) }
func (scrollBehaviorType) scrollBehaviorValue() {}

func (style Style) ScrollBehavior() scrollBehaviorValue {
	return scrollBehaviorType(style.Get("scroll-behavior"))
}
func (style Style) SetScrollBehavior(value scrollBehaviorValue) {
	style.set("scroll-behavior", value)
}

type backgroundValue interface {
	stringable
	backgroundValue()
}
type backgroundType string

func (self backgroundType) String() string { return string(self) }
func (backgroundType) backgroundValue()    {}

func (style Style) Background() backgroundValue {
	return backgroundType(style.Get("background"))
}
func (style Style) SetBackground(value backgroundValue) {
	style.set("background", value)
}

func (style Style) SetGridTemplateColumns(values []gridTemplateValue) {
	if len(values) == 0 {
		style.set("grid-template-columns", unitType("none"))
		return
	}
	var result string
	for _, value := range values {
		result += value.String() + " "
	}
	style.set("grid-template-columns", unitType(result))
}

func (style Style) ObjectPosition() unitAndUnitValue {
	return unitAndUnitType(style.Get("object-position"))
}
func (style Style) SetObjectPosition(value unitAndUnitValue) {
	style.set("object-position", value)
}
func (style Style) BorderLeftStyle() borderStyleValue {
	return borderStyleType(style.Get("border-left-style"))
}
func (style Style) SetBorderLeftStyle(value borderStyleValue) {
	style.set("border-left-style", value)
}
func (style Style) BreakAfter() breakValue {
	return breakType(style.Get("break-after"))
}
func (style Style) SetBreakAfter(value breakValue) {
	style.set("break-after", value)
}

type clearValue interface {
	stringable
	clearValue()
}
type clearType string

func (self clearType) String() string { return string(self) }
func (clearType) clearValue()         {}

func (style Style) Clear() clearValue {
	return clearType(style.Get("clear"))
}
func (style Style) SetClear(value clearValue) {
	style.set("clear", value)
}
func (style Style) ColumnRuleColor() colorValue {
	return colorType(style.Get("column-rule-color"))
}
func (style Style) SetColumnRuleColor(value colorValue) {
	style.set("column-rule-color", value)
}

type gridValue interface {
	stringable
	gridValue()
}
type gridType string

func (self gridType) String() string { return string(self) }
func (gridType) gridValue()          {}

func (style Style) Grid() gridValue {
	return gridType(style.Get("grid"))
}
func (style Style) SetGrid(value gridValue) {
	style.set("grid", value)
}

type transitionTimingFunctionValue interface {
	stringable
	transitionTimingFunctionValue()
}
type transitionTimingFunctionType string

func (self transitionTimingFunctionType) String() string            { return string(self) }
func (transitionTimingFunctionType) transitionTimingFunctionValue() {}

func (style Style) TransitionTimingFunction() transitionTimingFunctionValue {
	return transitionTimingFunctionType(style.Get("transition-timing-function"))
}
func (style Style) SetTransitionTimingFunction(value transitionTimingFunctionValue) {
	style.set("transition-timing-function", value)
}

type backgroundOriginValue interface {
	stringable
	backgroundOriginValue()
}
type backgroundOriginType string

func (self backgroundOriginType) String() string    { return string(self) }
func (backgroundOriginType) backgroundOriginValue() {}

func (style Style) BackgroundOrigin() backgroundOriginValue {
	return backgroundOriginType(style.Get("background-origin"))
}
func (style Style) SetBackgroundOrigin(value backgroundOriginValue) {
	style.set("background-origin", value)
}
func (style Style) BorderBottomRightRadius() unitValue {
	return unitType(style.Get("border-bottom-right-radius"))
}
func (style Style) SetBorderBottomRightRadius(value unitValue) {
	style.set("border-bottom-right-radius", value)
}

type paddingValue interface {
	stringable
	paddingValue()
}
type paddingType string

func (self paddingType) String() string { return string(self) }
func (paddingType) paddingValue()       {}

func (style Style) Padding() paddingValue {
	return paddingType(style.Get("padding"))
}
func (style Style) SetPadding(value paddingValue) {
	style.set("padding", value)
}

type cursorValue interface {
	stringable
	cursorValue()
}
type cursorType string

func (self cursorType) String() string { return string(self) }
func (cursorType) cursorValue()        {}

func (style Style) Cursor() cursorValue {
	return cursorType(style.Get("cursor"))
}
func (style Style) SetCursor(value cursorValue) {
	style.set("cursor", value)
}
func (style Style) GridColumnStart() gridStopValue {
	return gridStopType(style.Get("grid-column-start"))
}
func (style Style) SetGridColumnStart(value gridStopValue) {
	style.set("grid-column-start", value)
}
func (style Style) MaxHeight() unitOrNoneValue {
	return unitOrNoneType(style.Get("max-height"))
}
func (style Style) SetMaxHeight(value unitOrNoneValue) {
	style.set("max-height", value)
}
func (style Style) Orphans() uintValue {
	return uintType(style.Get("orphans"))
}
func (style Style) SetOrphans(value uintValue) {
	style.set("orphans", value)
}
func (style Style) OutlineColor() colorValue {
	return colorType(style.Get("outline-color"))
}
func (style Style) SetOutlineColor(value colorValue) {
	style.set("outline-color", value)
}

func (style Style) SetGridTemplateAreas(names []string) {
	if len(names) == 0 {
		style.set("grid-template-areas", unitType("none"))
		return
	}
	var result string
	for _, name := range names {
		result += name + " "
	}
	style.set("grid-template-areas", unitType(result))
}

func (style Style) PaddingLeft() unitValue {
	return unitType(style.Get("padding-left"))
}
func (style Style) SetPaddingLeft(value unitValue) {
	style.set("padding-left", value)
}

type textJustifyValue interface {
	stringable
	textJustifyValue()
}
type textJustifyType string

func (self textJustifyType) String() string { return string(self) }
func (textJustifyType) textJustifyValue()   {}

func (style Style) TextJustify() textJustifyValue {
	return textJustifyType(style.Get("text-justify"))
}
func (style Style) SetTextJustify(value textJustifyValue) {
	style.set("text-justify", value)
}

type unicodeBidiValue interface {
	stringable
	unicodeBidiValue()
}
type unicodeBidiType string

func (self unicodeBidiType) String() string { return string(self) }
func (unicodeBidiType) unicodeBidiValue()   {}

func (style Style) UnicodeBidi() unicodeBidiValue {
	return unicodeBidiType(style.Get("unicode-bidi"))
}
func (style Style) SetUnicodeBidi(value unicodeBidiValue) {
	style.set("unicode-bidi", value)
}

type flexWrapValue interface {
	stringable
	flexWrapValue()
}
type flexWrapType string

func (self flexWrapType) String() string { return string(self) }
func (flexWrapType) flexWrapValue()      {}

func (style Style) FlexWrap() flexWrapValue {
	return flexWrapType(style.Get("flex-wrap"))
}
func (style Style) SetFlexWrap(value flexWrapValue) {
	style.set("flex-wrap", value)
}

type transformValue interface {
	stringable
	transformValue()
}
type transformType string

func (self transformType) String() string { return string(self) }
func (transformType) transformValue()     {}

func (style Style) Transform() transformValue {
	return transformType(style.Get("transform"))
}
func (style Style) SetTransform(value transformValue) {
	style.set("transform", value)
}

type backgroundBlendModeValue interface {
	stringable
	backgroundBlendModeValue()
}
type backgroundBlendModeType string

func (self backgroundBlendModeType) String() string       { return string(self) }
func (backgroundBlendModeType) backgroundBlendModeValue() {}

func (style Style) BackgroundBlendMode() backgroundBlendModeValue {
	return backgroundBlendModeType(style.Get("background-blend-mode"))
}
func (style Style) SetBackgroundBlendMode(value backgroundBlendModeValue) {
	style.set("background-blend-mode", value)
}
func (style Style) BorderTopStyle() borderStyleValue {
	return borderStyleType(style.Get("border-top-style"))
}
func (style Style) SetBorderTopStyle(value borderStyleValue) {
	style.set("border-top-style", value)
}

type justifyContentValue interface {
	stringable
	justifyContentValue()
}
type justifyContentType string

func (self justifyContentType) String() string  { return string(self) }
func (justifyContentType) justifyContentValue() {}

func (style Style) JustifyContent() justifyContentValue {
	return justifyContentType(style.Get("justify-content"))
}
func (style Style) SetJustifyContent(value justifyContentValue) {
	style.set("justify-content", value)
}
func (style Style) FlexGrow() numberValue {
	return numberType(style.Get("flex-grow"))
}
func (style Style) SetFlexGrow(value numberValue) {
	style.set("flex-grow", value)
}
func (style Style) PaddingTop() unitValue {
	return unitType(style.Get("padding-top"))
}
func (style Style) SetPaddingTop(value unitValue) {
	style.set("padding-top", value)
}

type animationNameValue interface {
	stringable
	animationNameValue()
}
type animationNameType string

func (self animationNameType) String() string { return string(self) }
func (animationNameType) animationNameValue() {}

func (style Style) AnimationName() animationNameValue {
	return animationNameType(style.Get("animation-name"))
}
func (style Style) SetAnimationName(value animationNameValue) {
	style.set("animation-name", value)
}
func (style Style) BackgroundPosition() unitAndUnitValue {
	return unitAndUnitType(style.Get("background-position"))
}
func (style Style) SetBackgroundPosition(value unitAndUnitValue) {
	style.set("background-position", value)
}

type columnFillValue interface {
	stringable
	columnFillValue()
}
type columnFillType string

func (self columnFillType) String() string { return string(self) }
func (columnFillType) columnFillValue()    {}

func (style Style) ColumnFill() columnFillValue {
	return columnFillType(style.Get("column-fill"))
}
func (style Style) SetColumnFill(value columnFillValue) {
	style.set("column-fill", value)
}

type objectFitValue interface {
	stringable
	objectFitValue()
}
type objectFitType string

func (self objectFitType) String() string { return string(self) }
func (objectFitType) objectFitValue()     {}

func (style Style) ObjectFit() objectFitValue {
	return objectFitType(style.Get("object-fit"))
}
func (style Style) SetObjectFit(value objectFitValue) {
	style.set("object-fit", value)
}
func (style Style) GridRowGap() unitValue {
	return unitType(style.Get("grid-row-gap"))
}
func (style Style) SetGridRowGap(value unitValue) {
	style.set("grid-row-gap", value)
}

type marginValue interface {
	stringable
	marginValue()
}
type marginType string

func (self marginType) String() string { return string(self) }
func (marginType) marginValue()        {}

func (style Style) Margin() marginValue {
	return marginType(style.Get("margin"))
}
func (style Style) SetMargin(value marginValue) {
	style.set("margin", value)
}
func (style Style) BorderStyle() borderStyleValue {
	return borderStyleType(style.Get("border-style"))
}
func (style Style) SetBorderStyle(value borderStyleValue) {
	style.set("border-style", value)
}

type columnRuleValue interface {
	stringable
	columnRuleValue()
}
type columnRuleType string

func (self columnRuleType) String() string { return string(self) }
func (columnRuleType) columnRuleValue()    {}

func (style Style) ColumnRule() columnRuleValue {
	return columnRuleType(style.Get("column-rule"))
}
func (style Style) SetColumnRule(value columnRuleValue) {
	style.set("column-rule", value)
}

type fontVariantCapsValue interface {
	stringable
	fontVariantCapsValue()
}
type fontVariantCapsType string

func (self fontVariantCapsType) String() string   { return string(self) }
func (fontVariantCapsType) fontVariantCapsValue() {}

func (style Style) FontVariantCaps() fontVariantCapsValue {
	return fontVariantCapsType(style.Get("font-variant-caps"))
}
func (style Style) SetFontVariantCaps(value fontVariantCapsValue) {
	style.set("font-variant-caps", value)
}

type transitionValue interface {
	stringable
	transitionValue()
}
type transitionType string

func (self transitionType) String() string { return string(self) }
func (transitionType) transitionValue()    {}

func (style Style) Transition() transitionValue {
	return transitionType(style.Get("transition"))
}
func (style Style) SetTransition(value transitionValue) {
	style.set("transition", value)
}
func (style Style) BackgroundColor() colorValue {
	return colorType(style.Get("background-color"))
}
func (style Style) SetBackgroundColor(value colorValue) {
	style.set("background-color", value)
}
func (style Style) GridRowStart() gridStopValue {
	return gridStopType(style.Get("grid-row-start"))
}
func (style Style) SetGridRowStart(value gridStopValue) {
	style.set("grid-row-start", value)
}
func (style Style) Overflow() overflowValue {
	return overflowType(style.Get("overflow"))
}
func (style Style) SetOverflow(value overflowValue) {
	style.set("overflow", value)
}
func (style Style) TextIndent() unitValue {
	return unitType(style.Get("text-indent"))
}
func (style Style) SetTextIndent(value unitValue) {
	style.set("text-indent", value)
}

type fontSizeValue interface {
	stringable
	fontSizeValue()
}
type fontSizeType string

func (self fontSizeType) String() string { return string(self) }
func (fontSizeType) fontSizeValue()      {}

func (style Style) FontSize() fontSizeValue {
	return fontSizeType(style.Get("font-size"))
}
func (style Style) SetFontSize(value fontSizeValue) {
	style.set("font-size", value)
}

type fontSynthesisValue string

func (f fontSynthesisValue) String() string {
	return string(f)
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

func (style Style) SetFontSynthesis(value fontSynthesisValue) {
	style.set("font-synthesis", value)
}
