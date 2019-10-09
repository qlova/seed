/*This file is computer-generated*/
package css

const Unset unset = 0

type unset byte

func (unset) String() string   { return "unset" }
func (unset) transitionValue() {}

const Initial initial = 0

type initial byte

func (initial) String() string   { return "initial" }
func (initial) transitionValue() {}

const Inherit inherit = 0

type inherit byte

func (inherit) String() string   { return "inherit" }
func (inherit) transitionValue() {}

func (numberType) animationIterationCountValue() {}

const Infinite infinite = 0

type infinite byte

func (infinite) String() string                { return "infinite" }
func (infinite) animationIterationCountValue() {}

func (unset) animationIterationCountValue() {}

func (initial) animationIterationCountValue() {}

func (inherit) animationIterationCountValue() {}

const None none = 0

type none byte

func (none) String() string    { return "none" }
func (none) borderStyleValue() {}

const Solid solid = 0

type solid byte

func (solid) String() string    { return "solid" }
func (solid) borderStyleValue() {}

const Ridge ridge = 0

type ridge byte

func (ridge) String() string    { return "ridge" }
func (ridge) borderStyleValue() {}

const Outset outset = 0

type outset byte

func (outset) String() string    { return "outset" }
func (outset) borderStyleValue() {}

const Inset inset = 0

type inset byte

func (inset) String() string    { return "inset" }
func (inset) borderStyleValue() {}

const Hidden hidden = 0

type hidden byte

func (hidden) String() string    { return "hidden" }
func (hidden) borderStyleValue() {}

const Groove groove = 0

type groove byte

func (groove) String() string    { return "groove" }
func (groove) borderStyleValue() {}

const Double double = 0

type double byte

func (double) String() string    { return "double" }
func (double) borderStyleValue() {}

const Dotted dotted = 0

type dotted byte

func (dotted) String() string    { return "dotted" }
func (dotted) borderStyleValue() {}

const Dashed dashed = 0

type dashed byte

func (dashed) String() string    { return "dashed" }
func (dashed) borderStyleValue() {}

func (unset) borderStyleValue() {}

func (initial) borderStyleValue() {}

func (inherit) borderStyleValue() {}

const Fill fill = 0

type fill byte

func (fill) String() string  { return "fill" }
func (fill) objectFitValue() {}

const ScaleDown scaleDown = 0

type scaleDown byte

func (scaleDown) String() string  { return "scale-down" }
func (scaleDown) objectFitValue() {}

func (none) objectFitValue() {}

const Cover cover = 0

type cover byte

func (cover) String() string  { return "cover" }
func (cover) objectFitValue() {}

const Contain contain = 0

type contain byte

func (contain) String() string  { return "contain" }
func (contain) objectFitValue() {}

func (unset) objectFitValue() {}

func (initial) objectFitValue() {}

func (inherit) objectFitValue() {}

const Zero zero = 0

type zero byte

func (zero) String() string { return "0" }
func (zero) unitValue()     {}

func (unset) unitValue() {}

func (initial) unitValue() {}

func (inherit) unitValue() {}

type lengthType string

func (s lengthType) String() string { return string(s) }
func (lengthType) numberValue()     {}

type integerType string

func (s integerType) String() string { return string(s) }
func (integerType) numberValue()     {}

func (unset) numberValue() {}

func (initial) numberValue() {}

func (inherit) numberValue() {}

const All all = 0

type all byte

func (all) String() string           { return "all" }
func (all) transitionPropertyValue() {}

func (none) transitionPropertyValue() {}

func (unset) transitionPropertyValue() {}

func (initial) transitionPropertyValue() {}

func (inherit) transitionPropertyValue() {}

func (unset) backgroundValue() {}

func (initial) backgroundValue() {}

func (inherit) backgroundValue() {}

func (unset) unitAndUnitValue() {}

func (initial) unitAndUnitValue() {}

func (inherit) unitAndUnitValue() {}

func (none) clearValue() {}

const Right right = 0

type right byte

func (right) String() string { return "right" }
func (right) clearValue()    {}

const Left left = 0

type left byte

func (left) String() string { return "left" }
func (left) clearValue()    {}

const Both both = 0

type both byte

func (both) String() string { return "both" }
func (both) clearValue()    {}

func (unset) clearValue() {}

func (initial) clearValue() {}

func (inherit) clearValue() {}

type gridautoType string

func (s gridautoType) String() string { return string(s) }
func (gridautoType) gridAutoValue()   {}

func (unset) gridAutoValue() {}

func (initial) gridAutoValue() {}

func (inherit) gridAutoValue() {}

func (none) gridTemplateValue() {}

func (gridautoType) gridTemplateValue() {}

func (unset) gridTemplateValue() {}

func (initial) gridTemplateValue() {}

func (inherit) gridTemplateValue() {}

const Auto auto = 0

type auto byte

func (auto) String() string       { return "auto" }
func (auto) imageRenderingValue() {}

const Pixelated pixelated = 0

type pixelated byte

func (pixelated) String() string       { return "pixelated" }
func (pixelated) imageRenderingValue() {}

const CrispEdges crispEdges = 0

type crispEdges byte

func (crispEdges) String() string       { return "crisp-edges" }
func (crispEdges) imageRenderingValue() {}

func (unset) imageRenderingValue() {}

func (initial) imageRenderingValue() {}

func (inherit) imageRenderingValue() {}

func (solid) textDecorationStyleValue() {}

const Wavy wavy = 0

type wavy byte

func (wavy) String() string            { return "wavy" }
func (wavy) textDecorationStyleValue() {}

func (double) textDecorationStyleValue() {}

func (dotted) textDecorationStyleValue() {}

func (dashed) textDecorationStyleValue() {}

func (unset) textDecorationStyleValue() {}

func (initial) textDecorationStyleValue() {}

func (inherit) textDecorationStyleValue() {}

const Mixed mixed = 0

type mixed byte

func (mixed) String() string        { return "mixed" }
func (mixed) textOrientationValue() {}

const UseGlyphOrientation useGlyphOrientation = 0

type useGlyphOrientation byte

func (useGlyphOrientation) String() string        { return "use-glyph-orientation" }
func (useGlyphOrientation) textOrientationValue() {}

const Upright upright = 0

type upright byte

func (upright) String() string        { return "upright" }
func (upright) textOrientationValue() {}

const SidewaysRight sidewaysRight = 0

type sidewaysRight byte

func (sidewaysRight) String() string        { return "sideways-right" }
func (sidewaysRight) textOrientationValue() {}

const SidewaysLeft sidewaysLeft = 0

type sidewaysLeft byte

func (sidewaysLeft) String() string        { return "sideways-left" }
func (sidewaysLeft) textOrientationValue() {}

const Sideways sideways = 0

type sideways byte

func (sideways) String() string        { return "sideways" }
func (sideways) textOrientationValue() {}

func (unset) textOrientationValue() {}

func (initial) textOrientationValue() {}

func (inherit) textOrientationValue() {}

const Visible visible = 0

type visible byte

func (visible) String() string           { return "visible" }
func (visible) backfaceVisibilityValue() {}

func (hidden) backfaceVisibilityValue() {}

func (unset) backfaceVisibilityValue() {}

func (initial) backfaceVisibilityValue() {}

func (inherit) backfaceVisibilityValue() {}

func (none) imageValue() {}

type urlType string

func (s urlType) String() string { return string(s) }
func (urlType) imageValue()      {}

type gradientType string

func (s gradientType) String() string { return string(s) }
func (gradientType) imageValue()      {}

func (unset) imageValue() {}

func (initial) imageValue() {}

func (inherit) imageValue() {}

const HorizontalTb horizontalTb = 0

type horizontalTb byte

func (horizontalTb) String() string    { return "horizontal-tb" }
func (horizontalTb) writingModeValue() {}

const VerticalRl verticalRl = 0

type verticalRl byte

func (verticalRl) String() string    { return "vertical-rl" }
func (verticalRl) writingModeValue() {}

const VerticalLr verticalLr = 0

type verticalLr byte

func (verticalLr) String() string    { return "vertical-lr" }
func (verticalLr) writingModeValue() {}

func (unset) writingModeValue() {}

func (initial) writingModeValue() {}

func (inherit) writingModeValue() {}

func (none) unitOrNoneValue() {}

func (lengthType) unitOrNoneValue() {}

func (unset) unitOrNoneValue() {}

func (initial) unitOrNoneValue() {}

func (inherit) unitOrNoneValue() {}

func (unset) sizeValue() {}

func (initial) sizeValue() {}

func (inherit) sizeValue() {}

const Normal normal = 0

type normal byte

func (normal) String() string             { return "normal" }
func (normal) fontVariantLigaturesValue() {}

func (none) fontVariantLigaturesValue() {}

func (unset) fontVariantLigaturesValue() {}

func (initial) fontVariantLigaturesValue() {}

func (inherit) fontVariantLigaturesValue() {}

func (auto) unitOrAutoValue() {}

func (unitType) unitOrAutoValue() {}

func (zero) unitOrAutoValue() {}

func (unset) unitOrAutoValue() {}

func (initial) unitOrAutoValue() {}

func (inherit) unitOrAutoValue() {}

type featuretagvalueType string

func (s featuretagvalueType) String() string          { return string(s) }
func (featuretagvalueType) fontFeatureSettingsValue() {}

func (unset) fontFeatureSettingsValue() {}

func (initial) fontFeatureSettingsValue() {}

func (inherit) fontFeatureSettingsValue() {}

func (auto) alignSelfValue() {}

const Stretch stretch = 0

type stretch byte

func (stretch) String() string  { return "stretch" }
func (stretch) alignSelfValue() {}

const FlexStart flexStart = 0

type flexStart byte

func (flexStart) String() string  { return "flex-start" }
func (flexStart) alignSelfValue() {}

const FlexEnd flexEnd = 0

type flexEnd byte

func (flexEnd) String() string  { return "flex-end" }
func (flexEnd) alignSelfValue() {}

const Center center = 0

type center byte

func (center) String() string  { return "center" }
func (center) alignSelfValue() {}

const Baseline baseline = 0

type baseline byte

func (baseline) String() string  { return "baseline" }
func (baseline) alignSelfValue() {}

func (unset) alignSelfValue() {}

func (initial) alignSelfValue() {}

func (inherit) alignSelfValue() {}

const Transparent transparent = 0

type transparent byte

func (transparent) String() string { return "transparent" }
func (transparent) colorValue()    {}

const CurrentColor currentColor = 0

type currentColor byte

func (currentColor) String() string { return "currentColor" }
func (currentColor) colorValue()    {}

func (unset) colorValue() {}

func (initial) colorValue() {}

func (inherit) colorValue() {}

func (normal) fontWeightValue() {}

const Lighter lighter = 0

type lighter byte

func (lighter) String() string   { return "lighter" }
func (lighter) fontWeightValue() {}

const Bolder bolder = 0

type bolder byte

func (bolder) String() string   { return "bolder" }
func (bolder) fontWeightValue() {}

const Bold bold = 0

type bold byte

func (bold) String() string   { return "bold" }
func (bold) fontWeightValue() {}

func (integerType) fontWeightValue() {}

func (unset) fontWeightValue() {}

func (initial) fontWeightValue() {}

func (inherit) fontWeightValue() {}

func (visible) overflowValue() {}

const Scroll scroll = 0

type scroll byte

func (scroll) String() string { return "scroll" }
func (scroll) overflowValue() {}

func (hidden) overflowValue() {}

func (auto) overflowValue() {}

func (unset) overflowValue() {}

func (initial) overflowValue() {}

func (inherit) overflowValue() {}

func (auto) textJustifyValue() {}

func (none) textJustifyValue() {}

const InterWord interWord = 0

type interWord byte

func (interWord) String() string    { return "inter-word" }
func (interWord) textJustifyValue() {}

const Distribute distribute = 0

type distribute byte

func (distribute) String() string    { return "distribute" }
func (distribute) textJustifyValue() {}

func (unset) textJustifyValue() {}

func (initial) textJustifyValue() {}

func (inherit) textJustifyValue() {}

func (unset) borderRadiusValue() {}

func (initial) borderRadiusValue() {}

func (inherit) borderRadiusValue() {}

const Slice slice = 0

type slice byte

func (slice) String() string           { return "slice" }
func (slice) boxDecorationBreakValue() {}

const Clone clone = 0

type clone byte

func (clone) String() string           { return "clone" }
func (clone) boxDecorationBreakValue() {}

func (unset) boxDecorationBreakValue() {}

func (initial) boxDecorationBreakValue() {}

func (inherit) boxDecorationBreakValue() {}

const Top top = 0

type top byte

func (top) String() string    { return "top" }
func (top) captionSideValue() {}

const Bottom bottom = 0

type bottom byte

func (bottom) String() string    { return "bottom" }
func (bottom) captionSideValue() {}

func (unset) captionSideValue() {}

func (initial) captionSideValue() {}

func (inherit) captionSideValue() {}

func (normal) fontVariantAlternatesValue() {}

const HistoricalForms historicalForms = 0

type historicalForms byte

func (historicalForms) String() string              { return "historical-forms" }
func (historicalForms) fontVariantAlternatesValue() {}

func (unset) fontVariantAlternatesValue() {}

func (initial) fontVariantAlternatesValue() {}

func (inherit) fontVariantAlternatesValue() {}

func (auto) isolationValue() {}

const Isolate isolate = 0

type isolate byte

func (isolate) String() string  { return "isolate" }
func (isolate) isolationValue() {}

func (unset) isolationValue() {}

func (initial) isolationValue() {}

func (inherit) isolationValue() {}

func (unset) borderImageValue() {}

func (initial) borderImageValue() {}

func (inherit) borderImageValue() {}

const Medium medium = 0

type medium byte

func (medium) String() string  { return "medium" }
func (medium) thicknessValue() {}

const Thin thin = 0

type thin byte

func (thin) String() string  { return "thin" }
func (thin) thicknessValue() {}

const Thick thick = 0

type thick byte

func (thick) String() string  { return "thick" }
func (thick) thicknessValue() {}

func (lengthType) thicknessValue() {}

func (zero) thicknessValue() {}

func (unset) thicknessValue() {}

func (initial) thicknessValue() {}

func (inherit) thicknessValue() {}

func (unset) transitionTimingFunctionValue() {}

func (initial) transitionTimingFunctionValue() {}

func (inherit) transitionTimingFunctionValue() {}

func (normal) wordSpacingValue() {}

func (unitType) wordSpacingValue() {}

func (unset) wordSpacingValue() {}

func (initial) wordSpacingValue() {}

func (inherit) wordSpacingValue() {}

const Balance balance = 0

type balance byte

func (balance) String() string   { return "balance" }
func (balance) columnFillValue() {}

func (auto) columnFillValue() {}

func (unset) columnFillValue() {}

func (initial) columnFillValue() {}

func (inherit) columnFillValue() {}

func (normal) fontVariantEastAsianValue() {}

func (unset) fontVariantEastAsianValue() {}

func (initial) fontVariantEastAsianValue() {}

func (inherit) fontVariantEastAsianValue() {}

const Maunal maunal = 0

type maunal byte

func (maunal) String() string { return "maunal" }
func (maunal) hyphensValue()  {}

func (none) hyphensValue() {}

func (auto) hyphensValue() {}

func (unset) hyphensValue() {}

func (initial) hyphensValue() {}

func (inherit) hyphensValue() {}

func (normal) lineHeightValue() {}

func (unitType) lineHeightValue() {}

func (numberType) lineHeightValue() {}

func (unset) lineHeightValue() {}

func (initial) lineHeightValue() {}

func (inherit) lineHeightValue() {}

const Outside outside = 0

type outside byte

func (outside) String() string          { return "outside" }
func (outside) listStylePositionValue() {}

const Inside inside = 0

type inside byte

func (inside) String() string          { return "inside" }
func (inside) listStylePositionValue() {}

func (unset) listStylePositionValue() {}

func (initial) listStylePositionValue() {}

func (inherit) listStylePositionValue() {}

func (unset) borderBottomValue() {}

func (initial) borderBottomValue() {}

func (inherit) borderBottomValue() {}

func (none) resizeValue() {}

const Vertical vertical = 0

type vertical byte

func (vertical) String() string { return "vertical" }
func (vertical) resizeValue()   {}

const Horizontal horizontal = 0

type horizontal byte

func (horizontal) String() string { return "horizontal" }
func (horizontal) resizeValue()   {}

func (both) resizeValue() {}

func (unset) resizeValue() {}

func (initial) resizeValue() {}

func (inherit) resizeValue() {}

type timeType string

func (s timeType) String() string { return string(s) }
func (timeType) durationValue()   {}

func (unset) durationValue() {}

func (initial) durationValue() {}

func (inherit) durationValue() {}

func (unitType) unitOrNoneValue() {}

func (normal) wordBreakValue() {}

const KeepAll keepAll = 0

type keepAll byte

func (keepAll) String() string  { return "keep-all" }
func (keepAll) wordBreakValue() {}

const BreakAll breakAll = 0

type breakAll byte

func (breakAll) String() string  { return "break-all" }
func (breakAll) wordBreakValue() {}

func (unset) wordBreakValue() {}

func (initial) wordBreakValue() {}

func (inherit) wordBreakValue() {}

func (integerType) orderValue() {}

func (unset) orderValue() {}

func (initial) orderValue() {}

func (inherit) orderValue() {}

func (stretch) borderImageRepeatValue() {}

const Space space = 0

type space byte

func (space) String() string          { return "space" }
func (space) borderImageRepeatValue() {}

const Round round = 0

type round byte

func (round) String() string          { return "round" }
func (round) borderImageRepeatValue() {}

const Repeat repeat = 0

type repeat byte

func (repeat) String() string          { return "repeat" }
func (repeat) borderImageRepeatValue() {}

func (unset) borderImageRepeatValue() {}

func (initial) borderImageRepeatValue() {}

func (inherit) borderImageRepeatValue() {}

func (none) shadowValue() {}

func (unset) shadowValue() {}

func (initial) shadowValue() {}

func (inherit) shadowValue() {}

func (medium) columnRuleWidthValue() {}

func (thin) columnRuleWidthValue() {}

func (thick) columnRuleWidthValue() {}

func (lengthType) columnRuleWidthValue() {}

func (zero) columnRuleWidthValue() {}

func (unset) columnRuleWidthValue() {}

func (initial) columnRuleWidthValue() {}

func (inherit) columnRuleWidthValue() {}

func (auto) columnWidthValue() {}

const Length length = 0

type length byte

func (length) String() string    { return "length" }
func (length) columnWidthValue() {}

func (unset) columnWidthValue() {}

func (initial) columnWidthValue() {}

func (inherit) columnWidthValue() {}

func (normal) fontVariantPositionValue() {}

const Sub sub = 0

type sub byte

func (sub) String() string            { return "sub" }
func (sub) fontVariantPositionValue() {}

const Super super = 0

type super byte

func (super) String() string            { return "super" }
func (super) fontVariantPositionValue() {}

func (unset) fontVariantPositionValue() {}

func (initial) fontVariantPositionValue() {}

func (inherit) fontVariantPositionValue() {}

type gridstopType string

func (s gridstopType) String() string { return string(s) }
func (gridstopType) gridStopValue()   {}

func (unset) gridStopValue() {}

func (initial) gridStopValue() {}

func (inherit) gridStopValue() {}

const Start start = 0

type start byte

func (start) String() string  { return "start" }
func (start) textAlignValue() {}

func (right) textAlignValue() {}

const MatchParent matchParent = 0

type matchParent byte

func (matchParent) String() string  { return "match-parent" }
func (matchParent) textAlignValue() {}

func (left) textAlignValue() {}

const Justify justify = 0

type justify byte

func (justify) String() string  { return "justify" }
func (justify) textAlignValue() {}

const End end = 0

type end byte

func (end) String() string  { return "end" }
func (end) textAlignValue() {}

func (center) textAlignValue() {}

type stringType string

func (s stringType) String() string { return string(s) }
func (stringType) textAlignValue()  {}

func (unset) textAlignValue() {}

func (initial) textAlignValue() {}

func (inherit) textAlignValue() {}

type timingfunctionType string

func (s timingfunctionType) String() string              { return string(s) }
func (timingfunctionType) animationTimingFunctionValue() {}

func (unset) animationTimingFunctionValue() {}

func (initial) animationTimingFunctionValue() {}

func (inherit) animationTimingFunctionValue() {}

func (unset) columnRuleValue() {}

func (initial) columnRuleValue() {}

func (inherit) columnRuleValue() {}

func (auto) pointerEventsValue() {}

func (none) pointerEventsValue() {}

func (unset) pointerEventsValue() {}

func (initial) pointerEventsValue() {}

func (inherit) pointerEventsValue() {}

func (normal) unicodeBidiValue() {}

const Embed embed = 0

type embed byte

func (embed) String() string    { return "embed" }
func (embed) unicodeBidiValue() {}

const BidiOverride bidiOverride = 0

type bidiOverride byte

func (bidiOverride) String() string    { return "bidi-override" }
func (bidiOverride) unicodeBidiValue() {}

func (unset) unicodeBidiValue() {}

func (initial) unicodeBidiValue() {}

func (inherit) unicodeBidiValue() {}

func (zero) paddingValue() {}

func (unset) paddingValue() {}

func (initial) paddingValue() {}

func (inherit) paddingValue() {}

const Clip clip = 0

type clip byte

func (clip) String() string     { return "clip" }
func (clip) textOverflowValue() {}

const Ellipsis ellipsis = 0

type ellipsis byte

func (ellipsis) String() string     { return "ellipsis" }
func (ellipsis) textOverflowValue() {}

func (stringType) textOverflowValue() {}

func (unset) textOverflowValue() {}

func (initial) textOverflowValue() {}

func (inherit) textOverflowValue() {}

func (normal) normalOrUnitOrAutoValue() {}

func (lengthType) normalOrUnitOrAutoValue() {}

func (unset) normalOrUnitOrAutoValue() {}

func (initial) normalOrUnitOrAutoValue() {}

func (inherit) normalOrUnitOrAutoValue() {}

func (unset) animationValue() {}

func (initial) animationValue() {}

func (inherit) animationValue() {}

func (none) unitOrAutoValue() {}

func (all) unitOrAutoValue() {}

func (normal) fontLanguageOverrideValue() {}

func (stringType) fontLanguageOverrideValue() {}

func (unset) fontLanguageOverrideValue() {}

func (initial) fontLanguageOverrideValue() {}

func (inherit) fontLanguageOverrideValue() {}

func (none) hangingPunctuationValue() {}

const Last last = 0

type last byte

func (last) String() string           { return "last" }
func (last) hangingPunctuationValue() {}

const ForceEnd forceEnd = 0

type forceEnd byte

func (forceEnd) String() string           { return "force-end" }
func (forceEnd) hangingPunctuationValue() {}

const First first = 0

type first byte

func (first) String() string           { return "first" }
func (first) hangingPunctuationValue() {}

const AllowEnd allowEnd = 0

type allowEnd byte

func (allowEnd) String() string           { return "allow-end" }
func (allowEnd) hangingPunctuationValue() {}

func (unset) hangingPunctuationValue() {}

func (initial) hangingPunctuationValue() {}

func (inherit) hangingPunctuationValue() {}

func (none) quotesValue() {}

func (unset) quotesValue() {}

func (initial) quotesValue() {}

func (inherit) quotesValue() {}

func (none) textCombineUprightValue() {}

func (all) textCombineUprightValue() {}

func (unset) textCombineUprightValue() {}

func (initial) textCombineUprightValue() {}

func (inherit) textCombineUprightValue() {}

func (normal) contentValue() {}

const OpenQuote openQuote = 0

type openQuote byte

func (openQuote) String() string { return "open-quote" }
func (openQuote) contentValue()  {}

func (none) contentValue() {}

const NoOpenQuote noOpenQuote = 0

type noOpenQuote byte

func (noOpenQuote) String() string { return "no-open-quote" }
func (noOpenQuote) contentValue()  {}

const NoCloseQuote noCloseQuote = 0

type noCloseQuote byte

func (noCloseQuote) String() string { return "no-close-quote" }
func (noCloseQuote) contentValue()  {}

const Icon icon = 0

type icon byte

func (icon) String() string { return "icon" }
func (icon) contentValue()  {}

const CloseQuote closeQuote = 0

type closeQuote byte

func (closeQuote) String() string { return "close-quote" }
func (closeQuote) contentValue()  {}

func (urlType) contentValue() {}

func (stringType) contentValue() {}

type counterType string

func (s counterType) String() string { return string(s) }
func (counterType) contentValue()    {}

func (unset) contentValue() {}

func (initial) contentValue() {}

func (inherit) contentValue() {}

func (none) fontSizeAdjustValue() {}

func (numberType) fontSizeAdjustValue() {}

func (unset) fontSizeAdjustValue() {}

func (initial) fontSizeAdjustValue() {}

func (inherit) fontSizeAdjustValue() {}

func (normal) animationDirectionValue() {}

const Reverse reverse = 0

type reverse byte

func (reverse) String() string           { return "reverse" }
func (reverse) animationDirectionValue() {}

const AlternateReverse alternateReverse = 0

type alternateReverse byte

func (alternateReverse) String() string           { return "alternate-reverse" }
func (alternateReverse) animationDirectionValue() {}

const Alternate alternate = 0

type alternate byte

func (alternate) String() string           { return "alternate" }
func (alternate) animationDirectionValue() {}

func (unset) animationDirectionValue() {}

func (initial) animationDirectionValue() {}

func (inherit) animationDirectionValue() {}

func (normal) fontVariantValue() {}

const Unicase unicase = 0

type unicase byte

func (unicase) String() string    { return "unicase" }
func (unicase) fontVariantValue() {}

const TitlingCaps titlingCaps = 0

type titlingCaps byte

func (titlingCaps) String() string    { return "titling-caps" }
func (titlingCaps) fontVariantValue() {}

const SmallCaps smallCaps = 0

type smallCaps byte

func (smallCaps) String() string    { return "small-caps" }
func (smallCaps) fontVariantValue() {}

const PetiteCaps petiteCaps = 0

type petiteCaps byte

func (petiteCaps) String() string    { return "petite-caps" }
func (petiteCaps) fontVariantValue() {}

const AllSmallCaps allSmallCaps = 0

type allSmallCaps byte

func (allSmallCaps) String() string    { return "all-small-caps" }
func (allSmallCaps) fontVariantValue() {}

const AllPetiteCaps allPetiteCaps = 0

type allPetiteCaps byte

func (allPetiteCaps) String() string    { return "all-petite-caps" }
func (allPetiteCaps) fontVariantValue() {}

func (unset) fontVariantValue() {}

func (initial) fontVariantValue() {}

func (inherit) fontVariantValue() {}

type familynameType string

func (s familynameType) String() string { return string(s) }
func (familynameType) fontFamilyValue() {}

func (unset) fontFamilyValue() {}

func (initial) fontFamilyValue() {}

func (inherit) fontFamilyValue() {}

func (unset) willChangeValue() {}

func (initial) willChangeValue() {}

func (inherit) willChangeValue() {}

func (none) animationFillModeValue() {}

const Forwards forwards = 0

type forwards byte

func (forwards) String() string          { return "forwards" }
func (forwards) animationFillModeValue() {}

func (both) animationFillModeValue() {}

const Backwards backwards = 0

type backwards byte

func (backwards) String() string          { return "backwards" }
func (backwards) animationFillModeValue() {}

func (unset) animationFillModeValue() {}

func (initial) animationFillModeValue() {}

func (inherit) animationFillModeValue() {}

const Running running = 0

type running byte

func (running) String() string           { return "running" }
func (running) animationPlayStateValue() {}

const Paused paused = 0

type paused byte

func (paused) String() string           { return "paused" }
func (paused) animationPlayStateValue() {}

func (unset) animationPlayStateValue() {}

func (initial) animationPlayStateValue() {}

func (inherit) animationPlayStateValue() {}

func (none) textTransformValue() {}

const Uppercase uppercase = 0

type uppercase byte

func (uppercase) String() string      { return "uppercase" }
func (uppercase) textTransformValue() {}

const Lowercase lowercase = 0

type lowercase byte

func (lowercase) String() string      { return "lowercase" }
func (lowercase) textTransformValue() {}

const FullWidth fullWidth = 0

type fullWidth byte

func (fullWidth) String() string      { return "full-width" }
func (fullWidth) textTransformValue() {}

const Capitalize capitalize = 0

type capitalize byte

func (capitalize) String() string      { return "capitalize" }
func (capitalize) textTransformValue() {}

func (unset) textTransformValue() {}

func (initial) textTransformValue() {}

func (inherit) textTransformValue() {}

func (unset) allValue() {}

func (initial) allValue() {}

func (inherit) allValue() {}

func (scroll) backgroundAttachmentValue() {}

const Local local = 0

type local byte

func (local) String() string             { return "local" }
func (local) backgroundAttachmentValue() {}

const Fixed fixed = 0

type fixed byte

func (fixed) String() string             { return "fixed" }
func (fixed) backgroundAttachmentValue() {}

func (unset) backgroundAttachmentValue() {}

func (initial) backgroundAttachmentValue() {}

func (inherit) backgroundAttachmentValue() {}

const Seperate seperate = 0

type seperate byte

func (seperate) String() string       { return "seperate" }
func (seperate) borderCollapseValue() {}

const Collapse collapse = 0

type collapse byte

func (collapse) String() string       { return "collapse" }
func (collapse) borderCollapseValue() {}

func (unset) borderCollapseValue() {}

func (initial) borderCollapseValue() {}

func (inherit) borderCollapseValue() {}

func (normal) columnGapValue() {}

func (lengthType) columnGapValue() {}

func (unset) columnGapValue() {}

func (initial) columnGapValue() {}

func (inherit) columnGapValue() {}

func (auto) textUnderlinePositionValue() {}

const Under under = 0

type under byte

func (under) String() string              { return "under" }
func (under) textUnderlinePositionValue() {}

func (right) textUnderlinePositionValue() {}

func (left) textUnderlinePositionValue() {}

func (unset) textUnderlinePositionValue() {}

func (initial) textUnderlinePositionValue() {}

func (inherit) textUnderlinePositionValue() {}

func (auto) integerOrAutoValue() {}

func (integerType) integerOrAutoValue() {}

func (unset) integerOrAutoValue() {}

func (initial) integerOrAutoValue() {}

func (inherit) integerOrAutoValue() {}

func (auto) lineBreakValue() {}

const Strict strict = 0

type strict byte

func (strict) String() string  { return "strict" }
func (strict) lineBreakValue() {}

func (normal) lineBreakValue() {}

const Loose loose = 0

type loose byte

func (loose) String() string  { return "loose" }
func (loose) lineBreakValue() {}

func (unset) lineBreakValue() {}

func (initial) lineBreakValue() {}

func (inherit) lineBreakValue() {}

func (timeType) transitionDurationValue() {}

func (unset) transitionDurationValue() {}

func (initial) transitionDurationValue() {}

func (inherit) transitionDurationValue() {}

func (boxType) backgroundOriginValue() {}

func (unset) backgroundOriginValue() {}

func (initial) backgroundOriginValue() {}

func (inherit) backgroundOriginValue() {}

func (auto) columnCountValue() {}

func (integerType) columnCountValue() {}

func (unset) columnCountValue() {}

func (initial) columnCountValue() {}

func (inherit) columnCountValue() {}

func (auto) tableLayoutValue() {}

func (fixed) tableLayoutValue() {}

func (unset) tableLayoutValue() {}

func (initial) tableLayoutValue() {}

func (inherit) tableLayoutValue() {}

const Shadow shadow = 0

type shadow byte

func (shadow) String() string { return "shadow" }
func (shadow) shadowValue()   {}

const Flat flat = 0

type flat byte

func (flat) String() string       { return "flat" }
func (flat) transformStyleValue() {}

const Preserve3d preserve3d = 0

type preserve3d byte

func (preserve3d) String() string       { return "preserve-3d" }
func (preserve3d) transformStyleValue() {}

func (unset) transformStyleValue() {}

func (initial) transformStyleValue() {}

func (inherit) transformStyleValue() {}

func (stretch) alignContentValue() {}

const SpaceBetween spaceBetween = 0

type spaceBetween byte

func (spaceBetween) String() string     { return "space-between" }
func (spaceBetween) alignContentValue() {}

const SpaceAround spaceAround = 0

type spaceAround byte

func (spaceAround) String() string     { return "space-around" }
func (spaceAround) alignContentValue() {}

const SpaceEvenly spaceEvenly = 0

type spaceEvenly byte

func (spaceEvenly) String() string     { return "space-evenly" }
func (spaceEvenly) alignContentValue() {}

func (flexStart) alignContentValue() {}

func (flexEnd) alignContentValue() {}

func (center) alignContentValue() {}

func (unset) alignContentValue() {}

func (initial) alignContentValue() {}

func (inherit) alignContentValue() {}

const Invert invert = 0

type invert byte

func (invert) String() string { return "invert" }
func (invert) colorValue()    {}

func (auto) scrollBehaviorValue() {}

const Smooth smooth = 0

type smooth byte

func (smooth) String() string       { return "smooth" }
func (smooth) scrollBehaviorValue() {}

func (unset) scrollBehaviorValue() {}

func (initial) scrollBehaviorValue() {}

func (inherit) scrollBehaviorValue() {}

func (stretch) alignItemsValue() {}

func (flexStart) alignItemsValue() {}

func (flexEnd) alignItemsValue() {}

func (center) alignItemsValue() {}

func (baseline) alignItemsValue() {}

func (unset) alignItemsValue() {}

func (initial) alignItemsValue() {}

func (inherit) alignItemsValue() {}

func (unset) borderLeftValue() {}

func (initial) borderLeftValue() {}

func (inherit) borderLeftValue() {}

const Weight weight = 0

type weight byte

func (weight) String() string      { return "weight" }
func (weight) fontSynthesisValue() {}

const StyleProperty styleProperty = 0

type styleProperty byte

func (styleProperty) String() string      { return "style" }
func (styleProperty) fontSynthesisValue() {}

func (none) fontSynthesisValue() {}

func (unset) fontSynthesisValue() {}

func (initial) fontSynthesisValue() {}

func (inherit) fontSynthesisValue() {}

func (integerType) uintValue() {}

func (unset) uintValue() {}

func (initial) uintValue() {}

func (inherit) uintValue() {}

const Row row = 0

type row byte

func (row) String() string     { return "row" }
func (row) gridAutoFlowValue() {}

const Column column = 0

type column byte

func (column) String() string     { return "column" }
func (column) gridAutoFlowValue() {}

const Dense dense = 0

type dense byte

func (dense) String() string     { return "dense" }
func (dense) gridAutoFlowValue() {}

func (unset) gridAutoFlowValue() {}

func (initial) gridAutoFlowValue() {}

func (inherit) gridAutoFlowValue() {}

func (auto) userSelectValue() {}

func (none) userSelectValue() {}

const Text text = 0

type text byte

func (text) String() string   { return "text" }
func (text) userSelectValue() {}

func (all) userSelectValue() {}

func (unset) userSelectValue() {}

func (initial) userSelectValue() {}

func (inherit) userSelectValue() {}

func (unset) gridRowValue() {}

func (initial) gridRowValue() {}

func (inherit) gridRowValue() {}

const Static static = 0

type static byte

func (static) String() string { return "static" }
func (static) positionValue() {}

const Sticky sticky = 0

type sticky byte

func (sticky) String() string { return "sticky" }
func (sticky) positionValue() {}

const Relative relative = 0

type relative byte

func (relative) String() string { return "relative" }
func (relative) positionValue() {}

const Page page = 0

type page byte

func (page) String() string { return "page" }
func (page) positionValue() {}

func (fixed) positionValue() {}

func (center) positionValue() {}

const Absolute absolute = 0

type absolute byte

func (absolute) String() string { return "absolute" }
func (absolute) positionValue() {}

func (unset) positionValue() {}

func (initial) positionValue() {}

func (inherit) positionValue() {}

func (unset) borderRightValue() {}

func (initial) borderRightValue() {}

func (inherit) borderRightValue() {}

func (baseline) verticalAlignValue() {}

func (top) verticalAlignValue() {}

const TextTop textTop = 0

type textTop byte

func (textTop) String() string      { return "text-top" }
func (textTop) verticalAlignValue() {}

const TextBottom textBottom = 0

type textBottom byte

func (textBottom) String() string      { return "text-bottom" }
func (textBottom) verticalAlignValue() {}

func (super) verticalAlignValue() {}

func (sub) verticalAlignValue() {}

const Middle middle = 0

type middle byte

func (middle) String() string      { return "middle" }
func (middle) verticalAlignValue() {}

func (bottom) verticalAlignValue() {}

func (unitType) verticalAlignValue() {}

func (unset) verticalAlignValue() {}

func (initial) verticalAlignValue() {}

func (inherit) verticalAlignValue() {}

func (transparent) sizeValue() {}

func (colorType) sizeValue() {}

func (currentColor) sizeValue() {}

func (unitAndUnitType) transformOriginValue() {}

func (unset) transformOriginValue() {}

func (initial) transformOriginValue() {}

func (inherit) transformOriginValue() {}

func (visible) visibilityValue() {}

func (hidden) visibilityValue() {}

func (collapse) visibilityValue() {}

func (unset) visibilityValue() {}

func (initial) visibilityValue() {}

func (inherit) visibilityValue() {}

func (unset) fontValue() {}

func (initial) fontValue() {}

func (inherit) fontValue() {}

func (auto) normalOrAutoValue() {}

func (normal) normalOrAutoValue() {}

func (none) normalOrAutoValue() {}

func (unset) normalOrAutoValue() {}

func (initial) normalOrAutoValue() {}

func (inherit) normalOrAutoValue() {}

func (lengthType) unitValue() {}

func (normal) backgroundBlendModeValue() {}

const SoftLight softLight = 0

type softLight byte

func (softLight) String() string            { return "soft-light" }
func (softLight) backgroundBlendModeValue() {}

const Screen screen = 0

type screen byte

func (screen) String() string            { return "screen" }
func (screen) backgroundBlendModeValue() {}

const Saturation saturation = 0

type saturation byte

func (saturation) String() string            { return "saturation" }
func (saturation) backgroundBlendModeValue() {}

const Overlay overlay = 0

type overlay byte

func (overlay) String() string            { return "overlay" }
func (overlay) backgroundBlendModeValue() {}

const Multiply multiply = 0

type multiply byte

func (multiply) String() string            { return "multiply" }
func (multiply) backgroundBlendModeValue() {}

const Luminosity luminosity = 0

type luminosity byte

func (luminosity) String() string            { return "luminosity" }
func (luminosity) backgroundBlendModeValue() {}

const Lighten lighten = 0

type lighten byte

func (lighten) String() string            { return "lighten" }
func (lighten) backgroundBlendModeValue() {}

const Hue hue = 0

type hue byte

func (hue) String() string            { return "hue" }
func (hue) backgroundBlendModeValue() {}

const HardLight hardLight = 0

type hardLight byte

func (hardLight) String() string            { return "hard-light" }
func (hardLight) backgroundBlendModeValue() {}

const Exclusion exclusion = 0

type exclusion byte

func (exclusion) String() string            { return "exclusion" }
func (exclusion) backgroundBlendModeValue() {}

const Difference difference = 0

type difference byte

func (difference) String() string            { return "difference" }
func (difference) backgroundBlendModeValue() {}

const Darken darken = 0

type darken byte

func (darken) String() string            { return "darken" }
func (darken) backgroundBlendModeValue() {}

const ColorDodge colorDodge = 0

type colorDodge byte

func (colorDodge) String() string            { return "color-dodge" }
func (colorDodge) backgroundBlendModeValue() {}

const ColorBurn colorBurn = 0

type colorBurn byte

func (colorBurn) String() string            { return "color-burn" }
func (colorBurn) backgroundBlendModeValue() {}

const Color color = 0

type color byte

func (color) String() string            { return "color" }
func (color) backgroundBlendModeValue() {}

func (unset) backgroundBlendModeValue() {}

func (initial) backgroundBlendModeValue() {}

func (inherit) backgroundBlendModeValue() {}

const Show show = 0

type show byte

func (show) String() string   { return "show" }
func (show) emptyCellsValue() {}

const Hide hide = 0

type hide byte

func (hide) String() string   { return "hide" }
func (hide) emptyCellsValue() {}

func (unset) emptyCellsValue() {}

func (initial) emptyCellsValue() {}

func (inherit) emptyCellsValue() {}

type pagebreakType string

func (s pagebreakType) String() string { return string(s) }
func (pagebreakType) pageBreakValue()  {}

func (unset) pageBreakValue() {}

func (initial) pageBreakValue() {}

func (inherit) pageBreakValue() {}

func (unset) textDecorationValue() {}

func (initial) textDecorationValue() {}

func (inherit) textDecorationValue() {}

func (normal) fontStyleValue() {}

const Oblique oblique = 0

type oblique byte

func (oblique) String() string  { return "oblique" }
func (oblique) fontStyleValue() {}

const Italic italic = 0

type italic byte

func (italic) String() string  { return "italic" }
func (italic) fontStyleValue() {}

func (unset) fontStyleValue() {}

func (initial) fontStyleValue() {}

func (inherit) fontStyleValue() {}

func (none) animationNameValue() {}

type identifierType string

func (s identifierType) String() string    { return string(s) }
func (identifierType) animationNameValue() {}

func (unset) animationNameValue() {}

func (initial) animationNameValue() {}

func (inherit) animationNameValue() {}

func (auto) fontDisplayValue() {}

const Block block = 0

type block byte

func (block) String() string    { return "block" }
func (block) fontDisplayValue() {}

const Swap swap = 0

type swap byte

func (swap) String() string    { return "swap" }
func (swap) fontDisplayValue() {}

const Fallback fallback = 0

type fallback byte

func (fallback) String() string    { return "fallback" }
func (fallback) fontDisplayValue() {}

const Optional optional = 0

type optional byte

func (optional) String() string    { return "optional" }
func (optional) fontDisplayValue() {}

func (unset) fontDisplayValue() {}

func (initial) fontDisplayValue() {}

func (inherit) fontDisplayValue() {}

func (timeType) transitionDelayValue() {}

func (unset) transitionDelayValue() {}

func (initial) transitionDelayValue() {}

func (inherit) transitionDelayValue() {}

func (unitType) borderTopRightRadiusValue() {}

func (unset) borderTopRightRadiusValue() {}

func (initial) borderTopRightRadiusValue() {}

func (inherit) borderTopRightRadiusValue() {}

type breakvalueType string

func (s breakvalueType) String() string { return string(s) }
func (breakvalueType) breakValue()      {}

func (unset) breakValue() {}

func (initial) breakValue() {}

func (inherit) breakValue() {}

const Ltr ltr = 0

type ltr byte

func (ltr) String() string  { return "ltr" }
func (ltr) directionValue() {}

const Rtl rtl = 0

type rtl byte

func (rtl) String() string  { return "rtl" }
func (rtl) directionValue() {}

func (unset) directionValue() {}

func (initial) directionValue() {}

func (inherit) directionValue() {}

const Inline inline = 0

type inline byte

func (inline) String() string { return "inline" }
func (inline) displayValue()  {}

const TableRowGroup tableRowGroup = 0

type tableRowGroup byte

func (tableRowGroup) String() string { return "table-row-group" }
func (tableRowGroup) displayValue()  {}

const TableRow tableRow = 0

type tableRow byte

func (tableRow) String() string { return "table-row" }
func (tableRow) displayValue()  {}

const TableHeaderGroup tableHeaderGroup = 0

type tableHeaderGroup byte

func (tableHeaderGroup) String() string { return "table-header-group" }
func (tableHeaderGroup) displayValue()  {}

const TableFooterGroup tableFooterGroup = 0

type tableFooterGroup byte

func (tableFooterGroup) String() string { return "table-footer-group" }
func (tableFooterGroup) displayValue()  {}

const TableColumnGroup tableColumnGroup = 0

type tableColumnGroup byte

func (tableColumnGroup) String() string { return "table-column-group" }
func (tableColumnGroup) displayValue()  {}

const TableColumn tableColumn = 0

type tableColumn byte

func (tableColumn) String() string { return "table-column" }
func (tableColumn) displayValue()  {}

const TableCell tableCell = 0

type tableCell byte

func (tableCell) String() string { return "table-cell" }
func (tableCell) displayValue()  {}

const TableCaption tableCaption = 0

type tableCaption byte

func (tableCaption) String() string { return "table-caption" }
func (tableCaption) displayValue()  {}

const Table table = 0

type table byte

func (table) String() string { return "table" }
func (table) displayValue()  {}

const RunIn runIn = 0

type runIn byte

func (runIn) String() string { return "run-in" }
func (runIn) displayValue()  {}

func (none) displayValue() {}

const ListItem listItem = 0

type listItem byte

func (listItem) String() string { return "list-item" }
func (listItem) displayValue()  {}

const InlineTable inlineTable = 0

type inlineTable byte

func (inlineTable) String() string { return "inline-table" }
func (inlineTable) displayValue()  {}

const InlineFlex inlineFlex = 0

type inlineFlex byte

func (inlineFlex) String() string { return "inline-flex" }
func (inlineFlex) displayValue()  {}

const InlineBlock inlineBlock = 0

type inlineBlock byte

func (inlineBlock) String() string { return "inline-block" }
func (inlineBlock) displayValue()  {}

const Flex flex = 0

type flex byte

func (flex) String() string { return "flex" }
func (flex) displayValue()  {}

const Container container = 0

type container byte

func (container) String() string { return "container" }
func (container) displayValue()  {}

const Compact compact = 0

type compact byte

func (compact) String() string { return "compact" }
func (compact) displayValue()  {}

func (block) displayValue() {}

func (unset) displayValue() {}

func (initial) displayValue() {}

func (inherit) displayValue() {}

func (unset) uintOrUnitValue() {}

func (initial) uintOrUnitValue() {}

func (inherit) uintOrUnitValue() {}

func (normal) fontStretchValue() {}

const UltraExpanded ultraExpanded = 0

type ultraExpanded byte

func (ultraExpanded) String() string    { return "ultra-expanded" }
func (ultraExpanded) fontStretchValue() {}

const UltraCondensed ultraCondensed = 0

type ultraCondensed byte

func (ultraCondensed) String() string    { return "ultra-condensed" }
func (ultraCondensed) fontStretchValue() {}

const SemiExpanded semiExpanded = 0

type semiExpanded byte

func (semiExpanded) String() string    { return "semi-expanded" }
func (semiExpanded) fontStretchValue() {}

const SemiCondensed semiCondensed = 0

type semiCondensed byte

func (semiCondensed) String() string    { return "semi-condensed" }
func (semiCondensed) fontStretchValue() {}

const ExtraExpanded extraExpanded = 0

type extraExpanded byte

func (extraExpanded) String() string    { return "extra-expanded" }
func (extraExpanded) fontStretchValue() {}

const ExtraCondensed extraCondensed = 0

type extraCondensed byte

func (extraCondensed) String() string    { return "extra-condensed" }
func (extraCondensed) fontStretchValue() {}

const Expanded expanded = 0

type expanded byte

func (expanded) String() string    { return "expanded" }
func (expanded) fontStretchValue() {}

const Condensed condensed = 0

type condensed byte

func (condensed) String() string    { return "condensed" }
func (condensed) fontStretchValue() {}

func (unset) fontStretchValue() {}

func (initial) fontStretchValue() {}

func (inherit) fontStretchValue() {}

func (unset) gridValue() {}

func (initial) gridValue() {}

func (inherit) gridValue() {}

func (unset) outlineValue() {}

func (initial) outlineValue() {}

func (inherit) outlineValue() {}

func (auto) cursorValue() {}

const ZoomOut zoomOut = 0

type zoomOut byte

func (zoomOut) String() string { return "zoom-out" }
func (zoomOut) cursorValue()   {}

const ZoomIn zoomIn = 0

type zoomIn byte

func (zoomIn) String() string { return "zoom-in" }
func (zoomIn) cursorValue()   {}

const Wait wait = 0

type wait byte

func (wait) String() string { return "wait" }
func (wait) cursorValue()   {}

const WResize wResize = 0

type wResize byte

func (wResize) String() string { return "w-resize" }
func (wResize) cursorValue()   {}

const VerticalText verticalText = 0

type verticalText byte

func (verticalText) String() string { return "vertical-text" }
func (verticalText) cursorValue()   {}

func (urlType) cursorValue() {}

func (text) cursorValue() {}

const SwResize swResize = 0

type swResize byte

func (swResize) String() string { return "sw-resize" }
func (swResize) cursorValue()   {}

const SeResize seResize = 0

type seResize byte

func (seResize) String() string { return "se-resize" }
func (seResize) cursorValue()   {}

const SResize sResize = 0

type sResize byte

func (sResize) String() string { return "s-resize" }
func (sResize) cursorValue()   {}

const RowResize rowResize = 0

type rowResize byte

func (rowResize) String() string { return "row-resize" }
func (rowResize) cursorValue()   {}

const Progress progress = 0

type progress byte

func (progress) String() string { return "progress" }
func (progress) cursorValue()   {}

const Pointer pointer = 0

type pointer byte

func (pointer) String() string { return "pointer" }
func (pointer) cursorValue()   {}

const NwseResize nwseResize = 0

type nwseResize byte

func (nwseResize) String() string { return "nwse-resize" }
func (nwseResize) cursorValue()   {}

const NwResize nwResize = 0

type nwResize byte

func (nwResize) String() string { return "nw-resize" }
func (nwResize) cursorValue()   {}

const NsResize nsResize = 0

type nsResize byte

func (nsResize) String() string { return "ns-resize" }
func (nsResize) cursorValue()   {}

const NotAllowed notAllowed = 0

type notAllowed byte

func (notAllowed) String() string { return "not-allowed" }
func (notAllowed) cursorValue()   {}

func (none) cursorValue() {}

const NoDrop noDrop = 0

type noDrop byte

func (noDrop) String() string { return "no-drop" }
func (noDrop) cursorValue()   {}

const NeswResize neswResize = 0

type neswResize byte

func (neswResize) String() string { return "nesw-resize" }
func (neswResize) cursorValue()   {}

const NeResize neResize = 0

type neResize byte

func (neResize) String() string { return "ne-resize" }
func (neResize) cursorValue()   {}

const NResize nResize = 0

type nResize byte

func (nResize) String() string { return "n-resize" }
func (nResize) cursorValue()   {}

const Move move = 0

type move byte

func (move) String() string { return "move" }
func (move) cursorValue()   {}

const Help help = 0

type help byte

func (help) String() string { return "help" }
func (help) cursorValue()   {}

const EwResize ewResize = 0

type ewResize byte

func (ewResize) String() string { return "ew-resize" }
func (ewResize) cursorValue()   {}

const EResize eResize = 0

type eResize byte

func (eResize) String() string { return "e-resize" }
func (eResize) cursorValue()   {}

const Default defaultValue = 0

type defaultValue byte

func (defaultValue) String() string { return "default" }
func (defaultValue) cursorValue()   {}

const Crosshair crosshair = 0

type crosshair byte

func (crosshair) String() string { return "crosshair" }
func (crosshair) cursorValue()   {}

const Copy copy = 0

type copy byte

func (copy) String() string { return "copy" }
func (copy) cursorValue()   {}

const ContextMenu contextMenu = 0

type contextMenu byte

func (contextMenu) String() string { return "context-menu" }
func (contextMenu) cursorValue()   {}

const ColResize colResize = 0

type colResize byte

func (colResize) String() string { return "col-resize" }
func (colResize) cursorValue()   {}

const Cell cell = 0

type cell byte

func (cell) String() string { return "cell" }
func (cell) cursorValue()   {}

const AllScroll allScroll = 0

type allScroll byte

func (allScroll) String() string { return "all-scroll" }
func (allScroll) cursorValue()   {}

const Alias alias = 0

type alias byte

func (alias) String() string { return "alias" }
func (alias) cursorValue()   {}

func (unset) cursorValue() {}

func (initial) cursorValue() {}

func (inherit) cursorValue() {}

func (unset) flexValue() {}

func (initial) flexValue() {}

func (inherit) flexValue() {}

func (normal) mixBlendModeValue() {}

func (softLight) mixBlendModeValue() {}

func (screen) mixBlendModeValue() {}

func (saturation) mixBlendModeValue() {}

func (overlay) mixBlendModeValue() {}

func (multiply) mixBlendModeValue() {}

func (luminosity) mixBlendModeValue() {}

func (lighten) mixBlendModeValue() {}

func (hue) mixBlendModeValue() {}

func (hardLight) mixBlendModeValue() {}

func (exclusion) mixBlendModeValue() {}

func (difference) mixBlendModeValue() {}

func (darken) mixBlendModeValue() {}

func (colorDodge) mixBlendModeValue() {}

func (colorBurn) mixBlendModeValue() {}

func (color) mixBlendModeValue() {}

func (unset) mixBlendModeValue() {}

func (initial) mixBlendModeValue() {}

func (inherit) mixBlendModeValue() {}

func (normal) overflowWrapValue() {}

const BreakWord breakWord = 0

type breakWord byte

func (breakWord) String() string     { return "break-word" }
func (breakWord) overflowWrapValue() {}

func (unset) overflowWrapValue() {}

func (initial) overflowWrapValue() {}

func (inherit) overflowWrapValue() {}

func (unset) borderValue() {}

func (initial) borderValue() {}

func (inherit) borderValue() {}

func (unset) borderTopValue() {}

func (initial) borderTopValue() {}

func (inherit) borderTopValue() {}

func (none) counterIncrementValue() {}

func (unset) counterIncrementValue() {}

func (initial) counterIncrementValue() {}

func (inherit) counterIncrementValue() {}

func (medium) fontSizeValue() {}

const XxSmall xxSmall = 0

type xxSmall byte

func (xxSmall) String() string { return "xx-small" }
func (xxSmall) fontSizeValue() {}

const XxLarge xxLarge = 0

type xxLarge byte

func (xxLarge) String() string { return "xx-large" }
func (xxLarge) fontSizeValue() {}

const XSmall xSmall = 0

type xSmall byte

func (xSmall) String() string { return "x-small" }
func (xSmall) fontSizeValue() {}

const XLarge xLarge = 0

type xLarge byte

func (xLarge) String() string { return "x-large" }
func (xLarge) fontSizeValue() {}

const Smaller smaller = 0

type smaller byte

func (smaller) String() string { return "smaller" }
func (smaller) fontSizeValue() {}

const Small small = 0

type small byte

func (small) String() string { return "small" }
func (small) fontSizeValue() {}

const Larger larger = 0

type larger byte

func (larger) String() string { return "larger" }
func (larger) fontSizeValue() {}

const Large large = 0

type large byte

func (large) String() string { return "large" }
func (large) fontSizeValue() {}

func (unitType) fontSizeValue() {}

func (unset) fontSizeValue() {}

func (initial) fontSizeValue() {}

func (inherit) fontSizeValue() {}

func (auto) textAlignLastValue() {}

func (start) textAlignLastValue() {}

func (right) textAlignLastValue() {}

func (left) textAlignLastValue() {}

func (justify) textAlignLastValue() {}

func (end) textAlignLastValue() {}

func (center) textAlignLastValue() {}

func (unset) textAlignLastValue() {}

func (initial) textAlignLastValue() {}

func (inherit) textAlignLastValue() {}

func (length) unitOrAutoValue() {}

const Nowrap nowrap = 0

type nowrap byte

func (nowrap) String() string { return "nowrap" }
func (nowrap) flexWrapValue() {}

const Wrap wrap = 0

type wrap byte

func (wrap) String() string { return "wrap" }
func (wrap) flexWrapValue() {}

const WrapReverse wrapReverse = 0

type wrapReverse byte

func (wrapReverse) String() string { return "wrap-reverse" }
func (wrapReverse) flexWrapValue() {}

func (unset) flexWrapValue() {}

func (initial) flexWrapValue() {}

func (inherit) flexWrapValue() {}

func (normal) uintValue() {}

const PreWrap preWrap = 0

type preWrap byte

func (preWrap) String() string { return "pre-wrap" }
func (preWrap) uintValue()     {}

const PreLine preLine = 0

type preLine byte

func (preLine) String() string { return "pre-line" }
func (preLine) uintValue()     {}

const Pre pre = 0

type pre byte

func (pre) String() string { return "pre" }
func (pre) uintValue()     {}

func (nowrap) uintValue() {}

func (unset) gridGapValue() {}

func (initial) gridGapValue() {}

func (inherit) gridGapValue() {}

func (normal) wordWrapValue() {}

func (breakWord) wordWrapValue() {}

func (unset) wordWrapValue() {}

func (initial) wordWrapValue() {}

func (inherit) wordWrapValue() {}

func (none) floatValue() {}

func (left) floatValue() {}

func (right) floatValue() {}

func (unset) floatValue() {}

func (initial) floatValue() {}

func (inherit) floatValue() {}

func (integerType) widowsValue() {}

func (unset) widowsValue() {}

func (initial) widowsValue() {}

func (inherit) widowsValue() {}

func (repeat) backgroundRepeatValue() {}

func (space) backgroundRepeatValue() {}

func (round) backgroundRepeatValue() {}

const RepeatY repeatY = 0

type repeatY byte

func (repeatY) String() string         { return "repeat-y" }
func (repeatY) backgroundRepeatValue() {}

const RepeatX repeatX = 0

type repeatX byte

func (repeatX) String() string         { return "repeat-x" }
func (repeatX) backgroundRepeatValue() {}

const NoRepeat noRepeat = 0

type noRepeat byte

func (noRepeat) String() string         { return "no-repeat" }
func (noRepeat) backgroundRepeatValue() {}

func (unset) backgroundRepeatValue() {}

func (initial) backgroundRepeatValue() {}

func (inherit) backgroundRepeatValue() {}

func (none) textDecorationLineValue() {}

const Underline underline = 0

type underline byte

func (underline) String() string           { return "underline" }
func (underline) textDecorationLineValue() {}

const Overline overline = 0

type overline byte

func (overline) String() string           { return "overline" }
func (overline) textDecorationLineValue() {}

const LineThrough lineThrough = 0

type lineThrough byte

func (lineThrough) String() string           { return "line-through" }
func (lineThrough) textDecorationLineValue() {}

const Blink blink = 0

type blink byte

func (blink) String() string           { return "blink" }
func (blink) textDecorationLineValue() {}

func (unset) textDecorationLineValue() {}

func (initial) textDecorationLineValue() {}

func (inherit) textDecorationLineValue() {}

func (unset) boxValue() {}

func (initial) boxValue() {}

func (inherit) boxValue() {}

func (none) nameValue() {}

func (unset) nameValue() {}

func (initial) nameValue() {}

func (inherit) nameValue() {}

func (normal) fontVariantCapsValue() {}

func (unicase) fontVariantCapsValue() {}

func (titlingCaps) fontVariantCapsValue() {}

func (smallCaps) fontVariantCapsValue() {}

func (petiteCaps) fontVariantCapsValue() {}

func (allSmallCaps) fontVariantCapsValue() {}

func (allPetiteCaps) fontVariantCapsValue() {}

func (unset) fontVariantCapsValue() {}

func (initial) fontVariantCapsValue() {}

func (inherit) fontVariantCapsValue() {}

func (auto) breakInsideValue() {}

const AvoidPage avoidPage = 0

type avoidPage byte

func (avoidPage) String() string    { return "avoid-page" }
func (avoidPage) breakInsideValue() {}

const AvoidColumn avoidColumn = 0

type avoidColumn byte

func (avoidColumn) String() string    { return "avoid-column" }
func (avoidColumn) breakInsideValue() {}

const Avoid avoid = 0

type avoid byte

func (avoid) String() string    { return "avoid" }
func (avoid) breakInsideValue() {}

func (unset) breakInsideValue() {}

func (initial) breakInsideValue() {}

func (inherit) breakInsideValue() {}

func (unset) columnsValue() {}

func (initial) columnsValue() {}

func (inherit) columnsValue() {}

func (row) gridColumnValue() {}

func (column) gridColumnValue() {}

func (dense) gridColumnValue() {}

func (unset) gridColumnValue() {}

func (initial) gridColumnValue() {}

func (inherit) gridColumnValue() {}

func (none) transformValue() {}

type transformationType string

func (s transformationType) String() string { return string(s) }
func (transformationType) transformValue()  {}

func (unset) transformValue() {}

func (initial) transformValue() {}

func (inherit) transformValue() {}

func (unset) listStyleValue() {}

func (initial) listStyleValue() {}

func (inherit) listStyleValue() {}

const Disc disc = 0

type disc byte

func (disc) String() string      { return "disc" }
func (disc) listStyleTypeValue() {}

const UpperRoman upperRoman = 0

type upperRoman byte

func (upperRoman) String() string      { return "upper-roman" }
func (upperRoman) listStyleTypeValue() {}

const UpperLatin upperLatin = 0

type upperLatin byte

func (upperLatin) String() string      { return "upper-latin" }
func (upperLatin) listStyleTypeValue() {}

const UpperAlpha upperAlpha = 0

type upperAlpha byte

func (upperAlpha) String() string      { return "upper-alpha" }
func (upperAlpha) listStyleTypeValue() {}

const Square square = 0

type square byte

func (square) String() string      { return "square" }
func (square) listStyleTypeValue() {}

func (none) listStyleTypeValue() {}

const LowerRoman lowerRoman = 0

type lowerRoman byte

func (lowerRoman) String() string      { return "lower-roman" }
func (lowerRoman) listStyleTypeValue() {}

const LowerLatin lowerLatin = 0

type lowerLatin byte

func (lowerLatin) String() string      { return "lower-latin" }
func (lowerLatin) listStyleTypeValue() {}

const LowerGreek lowerGreek = 0

type lowerGreek byte

func (lowerGreek) String() string      { return "lower-greek" }
func (lowerGreek) listStyleTypeValue() {}

const LowerAlpha lowerAlpha = 0

type lowerAlpha byte

func (lowerAlpha) String() string      { return "lower-alpha" }
func (lowerAlpha) listStyleTypeValue() {}

const Georgian georgian = 0

type georgian byte

func (georgian) String() string      { return "georgian" }
func (georgian) listStyleTypeValue() {}

const DecimalLeadingZero decimalLeadingZero = 0

type decimalLeadingZero byte

func (decimalLeadingZero) String() string      { return "decimal-leading-zero" }
func (decimalLeadingZero) listStyleTypeValue() {}

const Decimal decimal = 0

type decimal byte

func (decimal) String() string      { return "decimal" }
func (decimal) listStyleTypeValue() {}

const Circle circle = 0

type circle byte

func (circle) String() string      { return "circle" }
func (circle) listStyleTypeValue() {}

const Armenian armenian = 0

type armenian byte

func (armenian) String() string      { return "armenian" }
func (armenian) listStyleTypeValue() {}

func (unset) listStyleTypeValue() {}

func (initial) listStyleTypeValue() {}

func (inherit) listStyleTypeValue() {}

func (none) filterValue() {}

type filterModeType string

func (s filterModeType) String() string { return string(s) }
func (filterModeType) filterValue()     {}

func (urlType) filterValue() {}

func (unset) filterValue() {}

func (initial) filterValue() {}

func (inherit) filterValue() {}

func (auto) pageBreakInsideValue() {}

func (avoid) pageBreakInsideValue() {}

func (unset) pageBreakInsideValue() {}

func (initial) pageBreakInsideValue() {}

func (inherit) pageBreakInsideValue() {}

func (unset) flexFlowValue() {}

func (initial) flexFlowValue() {}

func (inherit) flexFlowValue() {}

func (normal) fontVariantNumericValue() {}

func (unset) fontVariantNumericValue() {}

func (initial) fontVariantNumericValue() {}

func (inherit) fontVariantNumericValue() {}

func (flexStart) justifyContentValue() {}

func (spaceBetween) justifyContentValue() {}

func (spaceEvenly) justifyContentValue() {}

func (spaceAround) justifyContentValue() {}

func (flexEnd) justifyContentValue() {}

func (center) justifyContentValue() {}

func (unset) justifyContentValue() {}

func (initial) justifyContentValue() {}

func (inherit) justifyContentValue() {}

func (unset) gridAreaValue() {}

func (initial) gridAreaValue() {}

func (inherit) gridAreaValue() {}

func (auto) marginValue() {}

func (unitType) marginValue() {}

func (zero) marginValue() {}

func (unset) marginValue() {}

func (initial) marginValue() {}

func (inherit) marginValue() {}

func (unset) borderImageSliceValue() {}

func (initial) borderImageSliceValue() {}

func (inherit) borderImageSliceValue() {}

func (unitType) borderTopLeftRadiusValue() {}

func (unset) borderTopLeftRadiusValue() {}

func (initial) borderTopLeftRadiusValue() {}

func (inherit) borderTopLeftRadiusValue() {}

func (auto) clipValue() {}

func (unset) clipValue() {}

func (initial) clipValue() {}

func (inherit) clipValue() {}

func (none) gridTemplateAreasValue() {}

func (unset) gridTemplateAreasValue() {}

func (initial) gridTemplateAreasValue() {}

func (inherit) gridTemplateAreasValue() {}

func (row) flexDirectionValue() {}

const RowReverse rowReverse = 0

type rowReverse byte

func (rowReverse) String() string      { return "row-reverse" }
func (rowReverse) flexDirectionValue() {}

const ColumnReverse columnReverse = 0

type columnReverse byte

func (columnReverse) String() string      { return "column-reverse" }
func (columnReverse) flexDirectionValue() {}

func (column) flexDirectionValue() {}

func (unset) flexDirectionValue() {}

func (initial) flexDirectionValue() {}

func (inherit) flexDirectionValue() {}
