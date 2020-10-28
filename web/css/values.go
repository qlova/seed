/*This file is computer-generated*/
package css

const None none = "none"

type none string

func (none) Rule() Rule   { return "none" }
func (none) shadowValue() {}

const Unset unset = "unset"

type unset string

func (unset) Rule() Rule   { return "unset" }
func (unset) shadowValue() {}

const Initial initial = "initial"

type initial string

func (initial) Rule() Rule   { return "initial" }
func (initial) shadowValue() {}

const Inherit inherit = "inherit"

type inherit string

func (inherit) Rule() Rule   { return "inherit" }
func (inherit) shadowValue() {}

const Zero zero = "0"

type zero string

func (zero) Rule() Rule   { return "0" }
func (zero) numberValue() {}

func (unset) numberValue() {}

func (initial) numberValue() {}

func (inherit) numberValue() {}

func (none) unitOrNoneValue() {}

type lengthType string

func (s lengthType) String() string { return string(s) }
func (lengthType) unitOrNoneValue() {}

func (unset) unitOrNoneValue() {}

func (initial) unitOrNoneValue() {}

func (inherit) unitOrNoneValue() {}

func (unset) allValue() {}

func (initial) allValue() {}

func (inherit) allValue() {}

const Auto auto = "auto"

type auto string

func (auto) Rule() Rule       { return "auto" }
func (auto) unitOrAutoValue() {}

func (unitType) unitOrAutoValue() {}

func (zero) unitOrAutoValue() {}

func (unset) unitOrAutoValue() {}

func (initial) unitOrAutoValue() {}

func (inherit) unitOrAutoValue() {}

func (auto) columnCountValue() {}

type integerType string

func (s integerType) String() string  { return string(s) }
func (integerType) columnCountValue() {}

func (unset) columnCountValue() {}

func (initial) columnCountValue() {}

func (inherit) columnCountValue() {}

const Row row = "row"

type row string

func (row) Rule() Rule          { return "row" }
func (row) flexDirectionValue() {}

const RowReverse rowReverse = "row-reverse"

type rowReverse string

func (rowReverse) Rule() Rule          { return "row-reverse" }
func (rowReverse) flexDirectionValue() {}

const ColumnReverse columnReverse = "column-reverse"

type columnReverse string

func (columnReverse) Rule() Rule          { return "column-reverse" }
func (columnReverse) flexDirectionValue() {}

const Column column = "column"

type column string

func (column) Rule() Rule          { return "column" }
func (column) flexDirectionValue() {}

func (unset) flexDirectionValue() {}

func (initial) flexDirectionValue() {}

func (inherit) flexDirectionValue() {}

func (row) gridColumnValue() {}

func (column) gridColumnValue() {}

const Dense dense = "dense"

type dense string

func (dense) Rule() Rule       { return "dense" }
func (dense) gridColumnValue() {}

func (unset) gridColumnValue() {}

func (initial) gridColumnValue() {}

func (inherit) gridColumnValue() {}

func (auto) pageBreakInsideValue() {}

const Avoid avoid = "avoid"

type avoid string

func (avoid) Rule() Rule            { return "avoid" }
func (avoid) pageBreakInsideValue() {}

func (unset) pageBreakInsideValue() {}

func (initial) pageBreakInsideValue() {}

func (inherit) pageBreakInsideValue() {}

const All all = "all"

type all string

func (all) Rule() Rule               { return "all" }
func (all) transitionPropertyValue() {}

func (none) transitionPropertyValue() {}

func (unset) transitionPropertyValue() {}

func (initial) transitionPropertyValue() {}

func (inherit) transitionPropertyValue() {}

func (unset) borderImageValue() {}

func (initial) borderImageValue() {}

func (inherit) borderImageValue() {}

const Normal normal = "normal"

type normal string

func (normal) Rule() Rule         { return "normal" }
func (normal) mixBlendModeValue() {}

const SoftLight softLight = "soft-light"

type softLight string

func (softLight) Rule() Rule         { return "soft-light" }
func (softLight) mixBlendModeValue() {}

const Screen screen = "screen"

type screen string

func (screen) Rule() Rule         { return "screen" }
func (screen) mixBlendModeValue() {}

const Saturation saturation = "saturation"

type saturation string

func (saturation) Rule() Rule         { return "saturation" }
func (saturation) mixBlendModeValue() {}

const Overlay overlay = "overlay"

type overlay string

func (overlay) Rule() Rule         { return "overlay" }
func (overlay) mixBlendModeValue() {}

const Multiply multiply = "multiply"

type multiply string

func (multiply) Rule() Rule         { return "multiply" }
func (multiply) mixBlendModeValue() {}

const Luminosity luminosity = "luminosity"

type luminosity string

func (luminosity) Rule() Rule         { return "luminosity" }
func (luminosity) mixBlendModeValue() {}

const Lighten lighten = "lighten"

type lighten string

func (lighten) Rule() Rule         { return "lighten" }
func (lighten) mixBlendModeValue() {}

const Hue hue = "hue"

type hue string

func (hue) Rule() Rule         { return "hue" }
func (hue) mixBlendModeValue() {}

const HardLight hardLight = "hard-light"

type hardLight string

func (hardLight) Rule() Rule         { return "hard-light" }
func (hardLight) mixBlendModeValue() {}

const Exclusion exclusion = "exclusion"

type exclusion string

func (exclusion) Rule() Rule         { return "exclusion" }
func (exclusion) mixBlendModeValue() {}

const Difference difference = "difference"

type difference string

func (difference) Rule() Rule         { return "difference" }
func (difference) mixBlendModeValue() {}

const Darken darken = "darken"

type darken string

func (darken) Rule() Rule         { return "darken" }
func (darken) mixBlendModeValue() {}

const ColorDodge colorDodge = "color-dodge"

type colorDodge string

func (colorDodge) Rule() Rule         { return "color-dodge" }
func (colorDodge) mixBlendModeValue() {}

const ColorBurn colorBurn = "color-burn"

type colorBurn string

func (colorBurn) Rule() Rule         { return "color-burn" }
func (colorBurn) mixBlendModeValue() {}

const Color color = "color"

type color string

func (color) Rule() Rule         { return "color" }
func (color) mixBlendModeValue() {}

func (unset) mixBlendModeValue() {}

func (initial) mixBlendModeValue() {}

func (inherit) mixBlendModeValue() {}

const Fill fill = "fill"

type fill string

func (fill) Rule() Rule      { return "fill" }
func (fill) objectFitValue() {}

const ScaleDown scaleDown = "scale-down"

type scaleDown string

func (scaleDown) Rule() Rule      { return "scale-down" }
func (scaleDown) objectFitValue() {}

func (none) objectFitValue() {}

const Cover cover = "cover"

type cover string

func (cover) Rule() Rule      { return "cover" }
func (cover) objectFitValue() {}

const Contain contain = "contain"

type contain string

func (contain) Rule() Rule      { return "contain" }
func (contain) objectFitValue() {}

func (unset) objectFitValue() {}

func (initial) objectFitValue() {}

func (inherit) objectFitValue() {}

const Inline inline = "inline"

type inline string

func (inline) Rule() Rule    { return "inline" }
func (inline) displayValue() {}

const TableRowGroup tableRowGroup = "table-row-group"

type tableRowGroup string

func (tableRowGroup) Rule() Rule    { return "table-row-group" }
func (tableRowGroup) displayValue() {}

const TableRow tableRow = "table-row"

type tableRow string

func (tableRow) Rule() Rule    { return "table-row" }
func (tableRow) displayValue() {}

const TableHeaderGroup tableHeaderGroup = "table-header-group"

type tableHeaderGroup string

func (tableHeaderGroup) Rule() Rule    { return "table-header-group" }
func (tableHeaderGroup) displayValue() {}

const TableFooterGroup tableFooterGroup = "table-footer-group"

type tableFooterGroup string

func (tableFooterGroup) Rule() Rule    { return "table-footer-group" }
func (tableFooterGroup) displayValue() {}

const TableColumnGroup tableColumnGroup = "table-column-group"

type tableColumnGroup string

func (tableColumnGroup) Rule() Rule    { return "table-column-group" }
func (tableColumnGroup) displayValue() {}

const TableColumn tableColumn = "table-column"

type tableColumn string

func (tableColumn) Rule() Rule    { return "table-column" }
func (tableColumn) displayValue() {}

const TableCell tableCell = "table-cell"

type tableCell string

func (tableCell) Rule() Rule    { return "table-cell" }
func (tableCell) displayValue() {}

const TableCaption tableCaption = "table-caption"

type tableCaption string

func (tableCaption) Rule() Rule    { return "table-caption" }
func (tableCaption) displayValue() {}

const Table table = "table"

type table string

func (table) Rule() Rule    { return "table" }
func (table) displayValue() {}

const RunIn runIn = "run-in"

type runIn string

func (runIn) Rule() Rule    { return "run-in" }
func (runIn) displayValue() {}

func (none) displayValue() {}

const ListItem listItem = "list-item"

type listItem string

func (listItem) Rule() Rule    { return "list-item" }
func (listItem) displayValue() {}

const InlineTable inlineTable = "inline-table"

type inlineTable string

func (inlineTable) Rule() Rule    { return "inline-table" }
func (inlineTable) displayValue() {}

const InlineFlex inlineFlex = "inline-flex"

type inlineFlex string

func (inlineFlex) Rule() Rule    { return "inline-flex" }
func (inlineFlex) displayValue() {}

const InlineBlock inlineBlock = "inline-block"

type inlineBlock string

func (inlineBlock) Rule() Rule    { return "inline-block" }
func (inlineBlock) displayValue() {}

const Flex flex = "flex"

type flex string

func (flex) Rule() Rule    { return "flex" }
func (flex) displayValue() {}

const Container container = "container"

type container string

func (container) Rule() Rule    { return "container" }
func (container) displayValue() {}

const Compact compact = "compact"

type compact string

func (compact) Rule() Rule    { return "compact" }
func (compact) displayValue() {}

const Block block = "block"

type block string

func (block) Rule() Rule    { return "block" }
func (block) displayValue() {}

func (unset) displayValue() {}

func (initial) displayValue() {}

func (inherit) displayValue() {}

func (normal) fontVariantPositionValue() {}

const Sub sub = "sub"

type sub string

func (sub) Rule() Rule                { return "sub" }
func (sub) fontVariantPositionValue() {}

const Super super = "super"

type super string

func (super) Rule() Rule                { return "super" }
func (super) fontVariantPositionValue() {}

func (unset) fontVariantPositionValue() {}

func (initial) fontVariantPositionValue() {}

func (inherit) fontVariantPositionValue() {}

const Visible visible = "visible"

type visible string

func (visible) Rule() Rule     { return "visible" }
func (visible) overflowValue() {}

const Scroll scroll = "scroll"

type scroll string

func (scroll) Rule() Rule     { return "scroll" }
func (scroll) overflowValue() {}

const Hidden hidden = "hidden"

type hidden string

func (hidden) Rule() Rule     { return "hidden" }
func (hidden) overflowValue() {}

func (auto) overflowValue() {}

func (unset) overflowValue() {}

func (initial) overflowValue() {}

func (inherit) overflowValue() {}

func (boxType) backgroundOriginValue() {}

func (unset) backgroundOriginValue() {}

func (initial) backgroundOriginValue() {}

func (inherit) backgroundOriginValue() {}

func (unitType) borderTopLeftRadiusValue() {}

func (unset) borderTopLeftRadiusValue() {}

func (initial) borderTopLeftRadiusValue() {}

func (inherit) borderTopLeftRadiusValue() {}

const Maunal maunal = "maunal"

type maunal string

func (maunal) Rule() Rule    { return "maunal" }
func (maunal) hyphensValue() {}

func (none) hyphensValue() {}

func (auto) hyphensValue() {}

func (unset) hyphensValue() {}

func (initial) hyphensValue() {}

func (inherit) hyphensValue() {}

func (integerType) orderValue() {}

func (unset) orderValue() {}

func (initial) orderValue() {}

func (inherit) orderValue() {}

const Stretch stretch = "stretch"

type stretch string

func (stretch) Rule() Rule              { return "stretch" }
func (stretch) borderImageRepeatValue() {}

const Space space = "space"

type space string

func (space) Rule() Rule              { return "space" }
func (space) borderImageRepeatValue() {}

const Round round = "round"

type round string

func (round) Rule() Rule              { return "round" }
func (round) borderImageRepeatValue() {}

const Repeat repeat = "repeat"

type repeat string

func (repeat) Rule() Rule              { return "repeat" }
func (repeat) borderImageRepeatValue() {}

func (unset) borderImageRepeatValue() {}

func (initial) borderImageRepeatValue() {}

func (inherit) borderImageRepeatValue() {}

func (unset) thicknessValue() {}

func (initial) thicknessValue() {}

func (inherit) thicknessValue() {}

func (unset) textDecorationValue() {}

func (initial) textDecorationValue() {}

func (inherit) textDecorationValue() {}

func (unitAndUnitType) transformOriginValue() {}

func (unset) transformOriginValue() {}

func (initial) transformOriginValue() {}

func (inherit) transformOriginValue() {}

type timeType string

func (s timeType) String() string      { return string(s) }
func (timeType) transitionDelayValue() {}

func (unset) transitionDelayValue() {}

func (initial) transitionDelayValue() {}

func (inherit) transitionDelayValue() {}

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

func (zero) unitValue() {}

func (unset) unitValue() {}

func (initial) unitValue() {}

func (inherit) unitValue() {}

type pagebreakType string

func (s pagebreakType) String() string { return string(s) }
func (pagebreakType) pageBreakValue()  {}

func (unset) pageBreakValue() {}

func (initial) pageBreakValue() {}

func (inherit) pageBreakValue() {}

func (none) quotesValue() {}

func (unset) quotesValue() {}

func (initial) quotesValue() {}

func (inherit) quotesValue() {}

func (timeType) durationValue() {}

func (unset) durationValue() {}

func (initial) durationValue() {}

func (inherit) durationValue() {}

func (row) gridAutoFlowValue() {}

func (column) gridAutoFlowValue() {}

func (dense) gridAutoFlowValue() {}

func (unset) gridAutoFlowValue() {}

func (initial) gridAutoFlowValue() {}

func (inherit) gridAutoFlowValue() {}

const Disc disc = "disc"

type disc string

func (disc) Rule() Rule          { return "disc" }
func (disc) listStyleTypeValue() {}

const UpperRoman upperRoman = "upper-roman"

type upperRoman string

func (upperRoman) Rule() Rule          { return "upper-roman" }
func (upperRoman) listStyleTypeValue() {}

const UpperLatin upperLatin = "upper-latin"

type upperLatin string

func (upperLatin) Rule() Rule          { return "upper-latin" }
func (upperLatin) listStyleTypeValue() {}

const UpperAlpha upperAlpha = "upper-alpha"

type upperAlpha string

func (upperAlpha) Rule() Rule          { return "upper-alpha" }
func (upperAlpha) listStyleTypeValue() {}

const Square square = "square"

type square string

func (square) Rule() Rule          { return "square" }
func (square) listStyleTypeValue() {}

func (none) listStyleTypeValue() {}

const LowerRoman lowerRoman = "lower-roman"

type lowerRoman string

func (lowerRoman) Rule() Rule          { return "lower-roman" }
func (lowerRoman) listStyleTypeValue() {}

const LowerLatin lowerLatin = "lower-latin"

type lowerLatin string

func (lowerLatin) Rule() Rule          { return "lower-latin" }
func (lowerLatin) listStyleTypeValue() {}

const LowerGreek lowerGreek = "lower-greek"

type lowerGreek string

func (lowerGreek) Rule() Rule          { return "lower-greek" }
func (lowerGreek) listStyleTypeValue() {}

const LowerAlpha lowerAlpha = "lower-alpha"

type lowerAlpha string

func (lowerAlpha) Rule() Rule          { return "lower-alpha" }
func (lowerAlpha) listStyleTypeValue() {}

const Georgian georgian = "georgian"

type georgian string

func (georgian) Rule() Rule          { return "georgian" }
func (georgian) listStyleTypeValue() {}

const DecimalLeadingZero decimalLeadingZero = "decimal-leading-zero"

type decimalLeadingZero string

func (decimalLeadingZero) Rule() Rule          { return "decimal-leading-zero" }
func (decimalLeadingZero) listStyleTypeValue() {}

const Decimal decimal = "decimal"

type decimal string

func (decimal) Rule() Rule          { return "decimal" }
func (decimal) listStyleTypeValue() {}

const Circle circle = "circle"

type circle string

func (circle) Rule() Rule          { return "circle" }
func (circle) listStyleTypeValue() {}

const Armenian armenian = "armenian"

type armenian string

func (armenian) Rule() Rule          { return "armenian" }
func (armenian) listStyleTypeValue() {}

func (unset) listStyleTypeValue() {}

func (initial) listStyleTypeValue() {}

func (inherit) listStyleTypeValue() {}

func (unset) sizeValue() {}

func (initial) sizeValue() {}

func (inherit) sizeValue() {}

func (unset) borderTopValue() {}

func (initial) borderTopValue() {}

func (inherit) borderTopValue() {}

func (normal) fontVariantLigaturesValue() {}

func (none) fontVariantLigaturesValue() {}

func (unset) fontVariantLigaturesValue() {}

func (initial) fontVariantLigaturesValue() {}

func (inherit) fontVariantLigaturesValue() {}

func (unset) unitAndUnitValue() {}

func (initial) unitAndUnitValue() {}

func (inherit) unitAndUnitValue() {}

func (none) nameValue() {}

func (unset) nameValue() {}

func (initial) nameValue() {}

func (inherit) nameValue() {}

func (auto) columnWidthValue() {}

const Length length = "length"

type length string

func (length) Rule() Rule        { return "length" }
func (length) columnWidthValue() {}

func (unset) columnWidthValue() {}

func (initial) columnWidthValue() {}

func (inherit) columnWidthValue() {}

func (normal) fontVariantEastAsianValue() {}

func (unset) fontVariantEastAsianValue() {}

func (initial) fontVariantEastAsianValue() {}

func (inherit) fontVariantEastAsianValue() {}

func (normal) wordBreakValue() {}

const KeepAll keepAll = "keep-all"

type keepAll string

func (keepAll) Rule() Rule      { return "keep-all" }
func (keepAll) wordBreakValue() {}

const BreakAll breakAll = "break-all"

type breakAll string

func (breakAll) Rule() Rule      { return "break-all" }
func (breakAll) wordBreakValue() {}

func (unset) wordBreakValue() {}

func (initial) wordBreakValue() {}

func (inherit) wordBreakValue() {}

func (normal) fontVariantAlternatesValue() {}

const HistoricalForms historicalForms = "historical-forms"

type historicalForms string

func (historicalForms) Rule() Rule                  { return "historical-forms" }
func (historicalForms) fontVariantAlternatesValue() {}

func (unset) fontVariantAlternatesValue() {}

func (initial) fontVariantAlternatesValue() {}

func (inherit) fontVariantAlternatesValue() {}

const Mixed mixed = "mixed"

type mixed string

func (mixed) Rule() Rule            { return "mixed" }
func (mixed) textOrientationValue() {}

const UseGlyphOrientation useGlyphOrientation = "use-glyph-orientation"

type useGlyphOrientation string

func (useGlyphOrientation) Rule() Rule            { return "use-glyph-orientation" }
func (useGlyphOrientation) textOrientationValue() {}

const Upright upright = "upright"

type upright string

func (upright) Rule() Rule            { return "upright" }
func (upright) textOrientationValue() {}

const SidewaysRight sidewaysRight = "sideways-right"

type sidewaysRight string

func (sidewaysRight) Rule() Rule            { return "sideways-right" }
func (sidewaysRight) textOrientationValue() {}

const SidewaysLeft sidewaysLeft = "sideways-left"

type sidewaysLeft string

func (sidewaysLeft) Rule() Rule            { return "sideways-left" }
func (sidewaysLeft) textOrientationValue() {}

const Sideways sideways = "sideways"

type sideways string

func (sideways) Rule() Rule            { return "sideways" }
func (sideways) textOrientationValue() {}

func (unset) textOrientationValue() {}

func (initial) textOrientationValue() {}

func (inherit) textOrientationValue() {}

func (unset) colorValue() {}

func (initial) colorValue() {}

func (inherit) colorValue() {}

func (unset) gridAreaValue() {}

func (initial) gridAreaValue() {}

func (inherit) gridAreaValue() {}

type gridautoType string

func (s gridautoType) String() string { return string(s) }
func (gridautoType) gridAutoValue()   {}

func (unset) gridAutoValue() {}

func (initial) gridAutoValue() {}

func (inherit) gridAutoValue() {}

func (unset) outlineValue() {}

func (initial) outlineValue() {}

func (inherit) outlineValue() {}

const Medium medium = "medium"

type medium string

func (medium) Rule() Rule      { return "medium" }
func (medium) thicknessValue() {}

const Thin thin = "thin"

type thin string

func (thin) Rule() Rule      { return "thin" }
func (thin) thicknessValue() {}

const Thick thick = "thick"

type thick string

func (thick) Rule() Rule      { return "thick" }
func (thick) thicknessValue() {}

func (lengthType) thicknessValue() {}

func (zero) thicknessValue() {}

func (none) borderStyleValue() {}

const Solid solid = "solid"

type solid string

func (solid) Rule() Rule        { return "solid" }
func (solid) borderStyleValue() {}

const Ridge ridge = "ridge"

type ridge string

func (ridge) Rule() Rule        { return "ridge" }
func (ridge) borderStyleValue() {}

const Outset outset = "outset"

type outset string

func (outset) Rule() Rule        { return "outset" }
func (outset) borderStyleValue() {}

const Inset inset = "inset"

type inset string

func (inset) Rule() Rule        { return "inset" }
func (inset) borderStyleValue() {}

func (hidden) borderStyleValue() {}

const Groove groove = "groove"

type groove string

func (groove) Rule() Rule        { return "groove" }
func (groove) borderStyleValue() {}

const Double double = "double"

type double string

func (double) Rule() Rule        { return "double" }
func (double) borderStyleValue() {}

const Dotted dotted = "dotted"

type dotted string

func (dotted) Rule() Rule        { return "dotted" }
func (dotted) borderStyleValue() {}

const Dashed dashed = "dashed"

type dashed string

func (dashed) Rule() Rule        { return "dashed" }
func (dashed) borderStyleValue() {}

func (unset) borderStyleValue() {}

func (initial) borderStyleValue() {}

func (inherit) borderStyleValue() {}

func (solid) textDecorationStyleValue() {}

const Wavy wavy = "wavy"

type wavy string

func (wavy) Rule() Rule                { return "wavy" }
func (wavy) textDecorationStyleValue() {}

func (double) textDecorationStyleValue() {}

func (dotted) textDecorationStyleValue() {}

func (dashed) textDecorationStyleValue() {}

func (unset) textDecorationStyleValue() {}

func (initial) textDecorationStyleValue() {}

func (inherit) textDecorationStyleValue() {}

func (normal) backgroundBlendModeValue() {}

func (softLight) backgroundBlendModeValue() {}

func (screen) backgroundBlendModeValue() {}

func (saturation) backgroundBlendModeValue() {}

func (overlay) backgroundBlendModeValue() {}

func (multiply) backgroundBlendModeValue() {}

func (luminosity) backgroundBlendModeValue() {}

func (lighten) backgroundBlendModeValue() {}

func (hue) backgroundBlendModeValue() {}

func (hardLight) backgroundBlendModeValue() {}

func (exclusion) backgroundBlendModeValue() {}

func (difference) backgroundBlendModeValue() {}

func (darken) backgroundBlendModeValue() {}

func (colorDodge) backgroundBlendModeValue() {}

func (colorBurn) backgroundBlendModeValue() {}

func (color) backgroundBlendModeValue() {}

func (unset) backgroundBlendModeValue() {}

func (initial) backgroundBlendModeValue() {}

func (inherit) backgroundBlendModeValue() {}

func (medium) fontSizeValue() {}

const XxSmall xxSmall = "xx-small"

type xxSmall string

func (xxSmall) Rule() Rule     { return "xx-small" }
func (xxSmall) fontSizeValue() {}

const XxLarge xxLarge = "xx-large"

type xxLarge string

func (xxLarge) Rule() Rule     { return "xx-large" }
func (xxLarge) fontSizeValue() {}

const XSmall xSmall = "x-small"

type xSmall string

func (xSmall) Rule() Rule     { return "x-small" }
func (xSmall) fontSizeValue() {}

const XLarge xLarge = "x-large"

type xLarge string

func (xLarge) Rule() Rule     { return "x-large" }
func (xLarge) fontSizeValue() {}

const Smaller smaller = "smaller"

type smaller string

func (smaller) Rule() Rule     { return "smaller" }
func (smaller) fontSizeValue() {}

const Small small = "small"

type small string

func (small) Rule() Rule     { return "small" }
func (small) fontSizeValue() {}

const Larger larger = "larger"

type larger string

func (larger) Rule() Rule     { return "larger" }
func (larger) fontSizeValue() {}

const Large large = "large"

type large string

func (large) Rule() Rule     { return "large" }
func (large) fontSizeValue() {}

func (unitType) fontSizeValue() {}

func (unset) fontSizeValue() {}

func (initial) fontSizeValue() {}

func (inherit) fontSizeValue() {}

type gridstopType string

func (s gridstopType) String() string { return string(s) }
func (gridstopType) gridStopValue()   {}

func (unset) gridStopValue() {}

func (initial) gridStopValue() {}

func (inherit) gridStopValue() {}

const Transparent transparent = "transparent"

type transparent string

func (transparent) Rule() Rule  { return "transparent" }
func (transparent) colorValue() {}

const CurrentColor currentColor = "currentColor"

type currentColor string

func (currentColor) Rule() Rule  { return "currentColor" }
func (currentColor) colorValue() {}

func (unitType) borderTopRightRadiusValue() {}

func (unset) borderTopRightRadiusValue() {}

func (initial) borderTopRightRadiusValue() {}

func (inherit) borderTopRightRadiusValue() {}

func (unset) boxValue() {}

func (initial) boxValue() {}

func (inherit) boxValue() {}

func (unset) flexValue() {}

func (initial) flexValue() {}

func (inherit) flexValue() {}

type familynameType string

func (s familynameType) String() string { return string(s) }
func (familynameType) fontFamilyValue() {}

func (unset) fontFamilyValue() {}

func (initial) fontFamilyValue() {}

func (inherit) fontFamilyValue() {}

func (normal) lineHeightValue() {}

func (unitType) lineHeightValue() {}

func (numberType) lineHeightValue() {}

func (unset) lineHeightValue() {}

func (initial) lineHeightValue() {}

func (inherit) lineHeightValue() {}

const Outside outside = "outside"

type outside string

func (outside) Rule() Rule              { return "outside" }
func (outside) listStylePositionValue() {}

const Inside inside = "inside"

type inside string

func (inside) Rule() Rule              { return "inside" }
func (inside) listStylePositionValue() {}

func (unset) listStylePositionValue() {}

func (initial) listStylePositionValue() {}

func (inherit) listStylePositionValue() {}

func (zero) paddingValue() {}

func (unset) paddingValue() {}

func (initial) paddingValue() {}

func (inherit) paddingValue() {}

func (none) textTransformValue() {}

const Uppercase uppercase = "uppercase"

type uppercase string

func (uppercase) Rule() Rule          { return "uppercase" }
func (uppercase) textTransformValue() {}

const Lowercase lowercase = "lowercase"

type lowercase string

func (lowercase) Rule() Rule          { return "lowercase" }
func (lowercase) textTransformValue() {}

const FullWidth fullWidth = "full-width"

type fullWidth string

func (fullWidth) Rule() Rule          { return "full-width" }
func (fullWidth) textTransformValue() {}

const Capitalize capitalize = "capitalize"

type capitalize string

func (capitalize) Rule() Rule          { return "capitalize" }
func (capitalize) textTransformValue() {}

func (unset) textTransformValue() {}

func (initial) textTransformValue() {}

func (inherit) textTransformValue() {}

func (normal) fontStretchValue() {}

const UltraExpanded ultraExpanded = "ultra-expanded"

type ultraExpanded string

func (ultraExpanded) Rule() Rule        { return "ultra-expanded" }
func (ultraExpanded) fontStretchValue() {}

const UltraCondensed ultraCondensed = "ultra-condensed"

type ultraCondensed string

func (ultraCondensed) Rule() Rule        { return "ultra-condensed" }
func (ultraCondensed) fontStretchValue() {}

const SemiExpanded semiExpanded = "semi-expanded"

type semiExpanded string

func (semiExpanded) Rule() Rule        { return "semi-expanded" }
func (semiExpanded) fontStretchValue() {}

const SemiCondensed semiCondensed = "semi-condensed"

type semiCondensed string

func (semiCondensed) Rule() Rule        { return "semi-condensed" }
func (semiCondensed) fontStretchValue() {}

const ExtraExpanded extraExpanded = "extra-expanded"

type extraExpanded string

func (extraExpanded) Rule() Rule        { return "extra-expanded" }
func (extraExpanded) fontStretchValue() {}

const ExtraCondensed extraCondensed = "extra-condensed"

type extraCondensed string

func (extraCondensed) Rule() Rule        { return "extra-condensed" }
func (extraCondensed) fontStretchValue() {}

const Expanded expanded = "expanded"

type expanded string

func (expanded) Rule() Rule        { return "expanded" }
func (expanded) fontStretchValue() {}

const Condensed condensed = "condensed"

type condensed string

func (condensed) Rule() Rule        { return "condensed" }
func (condensed) fontStretchValue() {}

func (unset) fontStretchValue() {}

func (initial) fontStretchValue() {}

func (inherit) fontStretchValue() {}

func (unset) transitionTimingFunctionValue() {}

func (initial) transitionTimingFunctionValue() {}

func (inherit) transitionTimingFunctionValue() {}

func (auto) integerOrAutoValue() {}

func (integerType) integerOrAutoValue() {}

func (unset) integerOrAutoValue() {}

func (initial) integerOrAutoValue() {}

func (inherit) integerOrAutoValue() {}

func (none) fontSizeAdjustValue() {}

func (numberType) fontSizeAdjustValue() {}

func (unset) fontSizeAdjustValue() {}

func (initial) fontSizeAdjustValue() {}

func (inherit) fontSizeAdjustValue() {}

func (numberType) animationIterationCountValue() {}

const Infinite infinite = "infinite"

type infinite string

func (infinite) Rule() Rule                    { return "infinite" }
func (infinite) animationIterationCountValue() {}

func (unset) animationIterationCountValue() {}

func (initial) animationIterationCountValue() {}

func (inherit) animationIterationCountValue() {}

func (normal) contentValue() {}

const OpenQuote openQuote = "open-quote"

type openQuote string

func (openQuote) Rule() Rule    { return "open-quote" }
func (openQuote) contentValue() {}

func (none) contentValue() {}

const NoOpenQuote noOpenQuote = "no-open-quote"

type noOpenQuote string

func (noOpenQuote) Rule() Rule    { return "no-open-quote" }
func (noOpenQuote) contentValue() {}

const NoCloseQuote noCloseQuote = "no-close-quote"

type noCloseQuote string

func (noCloseQuote) Rule() Rule    { return "no-close-quote" }
func (noCloseQuote) contentValue() {}

const Icon icon = "icon"

type icon string

func (icon) Rule() Rule    { return "icon" }
func (icon) contentValue() {}

const CloseQuote closeQuote = "close-quote"

type closeQuote string

func (closeQuote) Rule() Rule    { return "close-quote" }
func (closeQuote) contentValue() {}

func (urlType) contentValue() {}

type stringType string

func (s stringType) String() string { return string(s) }
func (stringType) contentValue()    {}

type counterType string

func (s counterType) String() string { return string(s) }
func (counterType) contentValue()    {}

func (unset) contentValue() {}

func (initial) contentValue() {}

func (inherit) contentValue() {}

func (unset) uintOrUnitValue() {}

func (initial) uintOrUnitValue() {}

func (inherit) uintOrUnitValue() {}

const Baseline baseline = "baseline"

type baseline string

func (baseline) Rule() Rule          { return "baseline" }
func (baseline) verticalAlignValue() {}

const Top top = "top"

type top string

func (top) Rule() Rule          { return "top" }
func (top) verticalAlignValue() {}

const TextTop textTop = "text-top"

type textTop string

func (textTop) Rule() Rule          { return "text-top" }
func (textTop) verticalAlignValue() {}

const TextBottom textBottom = "text-bottom"

type textBottom string

func (textBottom) Rule() Rule          { return "text-bottom" }
func (textBottom) verticalAlignValue() {}

func (super) verticalAlignValue() {}

func (sub) verticalAlignValue() {}

const Middle middle = "middle"

type middle string

func (middle) Rule() Rule          { return "middle" }
func (middle) verticalAlignValue() {}

const Bottom bottom = "bottom"

type bottom string

func (bottom) Rule() Rule          { return "bottom" }
func (bottom) verticalAlignValue() {}

func (unitType) verticalAlignValue() {}

func (unset) verticalAlignValue() {}

func (initial) verticalAlignValue() {}

func (inherit) verticalAlignValue() {}

func (unset) fontValue() {}

func (initial) fontValue() {}

func (inherit) fontValue() {}

func (auto) marginValue() {}

func (unitType) marginValue() {}

func (zero) marginValue() {}

func (unset) marginValue() {}

func (initial) marginValue() {}

func (inherit) marginValue() {}

func (lengthType) unitValue() {}

func (unitType) unitOrNoneValue() {}

const Flat flat = "flat"

type flat string

func (flat) Rule() Rule           { return "flat" }
func (flat) transformStyleValue() {}

const Preserve3d preserve3d = "preserve-3d"

type preserve3d string

func (preserve3d) Rule() Rule           { return "preserve-3d" }
func (preserve3d) transformStyleValue() {}

func (unset) transformStyleValue() {}

func (initial) transformStyleValue() {}

func (inherit) transformStyleValue() {}

func (none) gridTemplateAreasValue() {}

func (unset) gridTemplateAreasValue() {}

func (initial) gridTemplateAreasValue() {}

func (inherit) gridTemplateAreasValue() {}

func (stretch) alignContentValue() {}

const SpaceBetween spaceBetween = "space-between"

type spaceBetween string

func (spaceBetween) Rule() Rule         { return "space-between" }
func (spaceBetween) alignContentValue() {}

const SpaceAround spaceAround = "space-around"

type spaceAround string

func (spaceAround) Rule() Rule         { return "space-around" }
func (spaceAround) alignContentValue() {}

const SpaceEvenly spaceEvenly = "space-evenly"

type spaceEvenly string

func (spaceEvenly) Rule() Rule         { return "space-evenly" }
func (spaceEvenly) alignContentValue() {}

const FlexStart flexStart = "flex-start"

type flexStart string

func (flexStart) Rule() Rule         { return "flex-start" }
func (flexStart) alignContentValue() {}

const FlexEnd flexEnd = "flex-end"

type flexEnd string

func (flexEnd) Rule() Rule         { return "flex-end" }
func (flexEnd) alignContentValue() {}

const Center center = "center"

type center string

func (center) Rule() Rule         { return "center" }
func (center) alignContentValue() {}

func (unset) alignContentValue() {}

func (initial) alignContentValue() {}

func (inherit) alignContentValue() {}

func (normal) fontVariantCapsValue() {}

const Unicase unicase = "unicase"

type unicase string

func (unicase) Rule() Rule            { return "unicase" }
func (unicase) fontVariantCapsValue() {}

const TitlingCaps titlingCaps = "titling-caps"

type titlingCaps string

func (titlingCaps) Rule() Rule            { return "titling-caps" }
func (titlingCaps) fontVariantCapsValue() {}

const SmallCaps smallCaps = "small-caps"

type smallCaps string

func (smallCaps) Rule() Rule            { return "small-caps" }
func (smallCaps) fontVariantCapsValue() {}

const PetiteCaps petiteCaps = "petite-caps"

type petiteCaps string

func (petiteCaps) Rule() Rule            { return "petite-caps" }
func (petiteCaps) fontVariantCapsValue() {}

const AllSmallCaps allSmallCaps = "all-small-caps"

type allSmallCaps string

func (allSmallCaps) Rule() Rule            { return "all-small-caps" }
func (allSmallCaps) fontVariantCapsValue() {}

const AllPetiteCaps allPetiteCaps = "all-petite-caps"

type allPetiteCaps string

func (allPetiteCaps) Rule() Rule            { return "all-petite-caps" }
func (allPetiteCaps) fontVariantCapsValue() {}

func (unset) fontVariantCapsValue() {}

func (initial) fontVariantCapsValue() {}

func (inherit) fontVariantCapsValue() {}

func (unset) gridRowValue() {}

func (initial) gridRowValue() {}

func (inherit) gridRowValue() {}

func (auto) imageRenderingValue() {}

const Pixelated pixelated = "pixelated"

type pixelated string

func (pixelated) Rule() Rule           { return "pixelated" }
func (pixelated) imageRenderingValue() {}

const CrispEdges crispEdges = "crisp-edges"

type crispEdges string

func (crispEdges) Rule() Rule           { return "crisp-edges" }
func (crispEdges) imageRenderingValue() {}

func (unset) imageRenderingValue() {}

func (initial) imageRenderingValue() {}

func (inherit) imageRenderingValue() {}

func (unset) gridValue() {}

func (initial) gridValue() {}

func (inherit) gridValue() {}

func (unset) willChangeValue() {}

func (initial) willChangeValue() {}

func (inherit) willChangeValue() {}

func (unset) borderValue() {}

func (initial) borderValue() {}

func (inherit) borderValue() {}

func (unset) borderBottomValue() {}

func (initial) borderBottomValue() {}

func (inherit) borderBottomValue() {}

func (auto) fontDisplayValue() {}

func (block) fontDisplayValue() {}

const Swap swap = "swap"

type swap string

func (swap) Rule() Rule        { return "swap" }
func (swap) fontDisplayValue() {}

const Fallback fallback = "fallback"

type fallback string

func (fallback) Rule() Rule        { return "fallback" }
func (fallback) fontDisplayValue() {}

const Optional optional = "optional"

type optional string

func (optional) Rule() Rule        { return "optional" }
func (optional) fontDisplayValue() {}

func (unset) fontDisplayValue() {}

func (initial) fontDisplayValue() {}

func (inherit) fontDisplayValue() {}

func (none) transformValue() {}

type transformationType string

func (s transformationType) String() string { return string(s) }
func (transformationType) transformValue()  {}

func (unset) transformValue() {}

func (initial) transformValue() {}

func (inherit) transformValue() {}

func (repeat) backgroundRepeatValue() {}

func (space) backgroundRepeatValue() {}

func (round) backgroundRepeatValue() {}

const RepeatY repeatY = "repeat-y"

type repeatY string

func (repeatY) Rule() Rule             { return "repeat-y" }
func (repeatY) backgroundRepeatValue() {}

const RepeatX repeatX = "repeat-x"

type repeatX string

func (repeatX) Rule() Rule             { return "repeat-x" }
func (repeatX) backgroundRepeatValue() {}

const NoRepeat noRepeat = "no-repeat"

type noRepeat string

func (noRepeat) Rule() Rule             { return "no-repeat" }
func (noRepeat) backgroundRepeatValue() {}

func (unset) backgroundRepeatValue() {}

func (initial) backgroundRepeatValue() {}

func (inherit) backgroundRepeatValue() {}

func (integerType) widowsValue() {}

func (unset) widowsValue() {}

func (initial) widowsValue() {}

func (inherit) widowsValue() {}

func (unset) borderRadiusValue() {}

func (initial) borderRadiusValue() {}

func (inherit) borderRadiusValue() {}

func (unset) borderRightValue() {}

func (initial) borderRightValue() {}

func (inherit) borderRightValue() {}

const Nowrap nowrap = "nowrap"

type nowrap string

func (nowrap) Rule() Rule     { return "nowrap" }
func (nowrap) flexWrapValue() {}

const Wrap wrap = "wrap"

type wrap string

func (wrap) Rule() Rule     { return "wrap" }
func (wrap) flexWrapValue() {}

const WrapReverse wrapReverse = "wrap-reverse"

type wrapReverse string

func (wrapReverse) Rule() Rule     { return "wrap-reverse" }
func (wrapReverse) flexWrapValue() {}

func (unset) flexWrapValue() {}

func (initial) flexWrapValue() {}

func (inherit) flexWrapValue() {}

func (visible) backfaceVisibilityValue() {}

func (hidden) backfaceVisibilityValue() {}

func (unset) backfaceVisibilityValue() {}

func (initial) backfaceVisibilityValue() {}

func (inherit) backfaceVisibilityValue() {}

func (auto) userSelectValue() {}

func (none) userSelectValue() {}

const Text text = "text"

type text string

func (text) Rule() Rule       { return "text" }
func (text) userSelectValue() {}

func (all) userSelectValue() {}

func (unset) userSelectValue() {}

func (initial) userSelectValue() {}

func (inherit) userSelectValue() {}

func (scroll) backgroundAttachmentValue() {}

const Local local = "local"

type local string

func (local) Rule() Rule                 { return "local" }
func (local) backgroundAttachmentValue() {}

const Fixed fixed = "fixed"

type fixed string

func (fixed) Rule() Rule                 { return "fixed" }
func (fixed) backgroundAttachmentValue() {}

func (unset) backgroundAttachmentValue() {}

func (initial) backgroundAttachmentValue() {}

func (inherit) backgroundAttachmentValue() {}

func (normal) fontVariantNumericValue() {}

func (unset) fontVariantNumericValue() {}

func (initial) fontVariantNumericValue() {}

func (inherit) fontVariantNumericValue() {}

func (lengthType) numberValue() {}

func (integerType) numberValue() {}

func (auto) tableLayoutValue() {}

func (fixed) tableLayoutValue() {}

func (unset) tableLayoutValue() {}

func (initial) tableLayoutValue() {}

func (inherit) tableLayoutValue() {}

type timingfunctionType string

func (s timingfunctionType) String() string              { return string(s) }
func (timingfunctionType) animationTimingFunctionValue() {}

func (unset) animationTimingFunctionValue() {}

func (initial) animationTimingFunctionValue() {}

func (inherit) animationTimingFunctionValue() {}

func (none) clearValue() {}

const Right right = "right"

type right string

func (right) Rule() Rule  { return "right" }
func (right) clearValue() {}

const Left left = "left"

type left string

func (left) Rule() Rule  { return "left" }
func (left) clearValue() {}

const Both both = "both"

type both string

func (both) Rule() Rule  { return "both" }
func (both) clearValue() {}

func (unset) clearValue() {}

func (initial) clearValue() {}

func (inherit) clearValue() {}

func (none) resizeValue() {}

const Vertical vertical = "vertical"

type vertical string

func (vertical) Rule() Rule   { return "vertical" }
func (vertical) resizeValue() {}

const Horizontal horizontal = "horizontal"

type horizontal string

func (horizontal) Rule() Rule   { return "horizontal" }
func (horizontal) resizeValue() {}

func (both) resizeValue() {}

func (unset) resizeValue() {}

func (initial) resizeValue() {}

func (inherit) resizeValue() {}

func (none) textCombineUprightValue() {}

func (all) textCombineUprightValue() {}

func (unset) textCombineUprightValue() {}

func (initial) textCombineUprightValue() {}

func (inherit) textCombineUprightValue() {}

func (unset) borderImageSliceValue() {}

func (initial) borderImageSliceValue() {}

func (inherit) borderImageSliceValue() {}

const Show show = "show"

type show string

func (show) Rule() Rule       { return "show" }
func (show) emptyCellsValue() {}

const Hide hide = "hide"

type hide string

func (hide) Rule() Rule       { return "hide" }
func (hide) emptyCellsValue() {}

func (unset) emptyCellsValue() {}

func (initial) emptyCellsValue() {}

func (inherit) emptyCellsValue() {}

const Weight weight = "weight"

type weight string

func (weight) Rule() Rule          { return "weight" }
func (weight) fontSynthesisValue() {}

const StyleProperty styleProperty = "style"

type styleProperty string

func (styleProperty) Rule() Rule          { return "style" }
func (styleProperty) fontSynthesisValue() {}

func (none) fontSynthesisValue() {}

func (unset) fontSynthesisValue() {}

func (initial) fontSynthesisValue() {}

func (inherit) fontSynthesisValue() {}

func (unset) gridGapValue() {}

func (initial) gridGapValue() {}

func (inherit) gridGapValue() {}

func (unset) flexFlowValue() {}

func (initial) flexFlowValue() {}

func (inherit) flexFlowValue() {}

func (auto) scrollBehaviorValue() {}

const Smooth smooth = "smooth"

type smooth string

func (smooth) Rule() Rule           { return "smooth" }
func (smooth) scrollBehaviorValue() {}

func (unset) scrollBehaviorValue() {}

func (initial) scrollBehaviorValue() {}

func (inherit) scrollBehaviorValue() {}

const HorizontalTb horizontalTb = "horizontal-tb"

type horizontalTb string

func (horizontalTb) Rule() Rule        { return "horizontal-tb" }
func (horizontalTb) writingModeValue() {}

const VerticalRl verticalRl = "vertical-rl"

type verticalRl string

func (verticalRl) Rule() Rule        { return "vertical-rl" }
func (verticalRl) writingModeValue() {}

const VerticalLr verticalLr = "vertical-lr"

type verticalLr string

func (verticalLr) Rule() Rule        { return "vertical-lr" }
func (verticalLr) writingModeValue() {}

func (unset) writingModeValue() {}

func (initial) writingModeValue() {}

func (inherit) writingModeValue() {}

func (normal) uintValue() {}

const PreWrap preWrap = "pre-wrap"

type preWrap string

func (preWrap) Rule() Rule { return "pre-wrap" }
func (preWrap) uintValue() {}

const PreLine preLine = "pre-line"

type preLine string

func (preLine) Rule() Rule { return "pre-line" }
func (preLine) uintValue() {}

const Pre pre = "pre"

type pre string

func (pre) Rule() Rule { return "pre" }
func (pre) uintValue() {}

func (nowrap) uintValue() {}

func (unset) uintValue() {}

func (initial) uintValue() {}

func (inherit) uintValue() {}

func (unset) animationValue() {}

func (initial) animationValue() {}

func (inherit) animationValue() {}

func (normal) overflowWrapValue() {}

const BreakWord breakWord = "break-word"

type breakWord string

func (breakWord) Rule() Rule         { return "break-word" }
func (breakWord) overflowWrapValue() {}

func (unset) overflowWrapValue() {}

func (initial) overflowWrapValue() {}

func (inherit) overflowWrapValue() {}

func (auto) alignSelfValue() {}

func (stretch) alignSelfValue() {}

func (flexStart) alignSelfValue() {}

func (flexEnd) alignSelfValue() {}

func (center) alignSelfValue() {}

func (baseline) alignSelfValue() {}

func (unset) alignSelfValue() {}

func (initial) alignSelfValue() {}

func (inherit) alignSelfValue() {}

const Balance balance = "balance"

type balance string

func (balance) Rule() Rule       { return "balance" }
func (balance) columnFillValue() {}

func (auto) columnFillValue() {}

func (unset) columnFillValue() {}

func (initial) columnFillValue() {}

func (inherit) columnFillValue() {}

func (unset) columnsValue() {}

func (initial) columnsValue() {}

func (inherit) columnsValue() {}

func (auto) isolationValue() {}

const Isolate isolate = "isolate"

type isolate string

func (isolate) Rule() Rule      { return "isolate" }
func (isolate) isolationValue() {}

func (unset) isolationValue() {}

func (initial) isolationValue() {}

func (inherit) isolationValue() {}

const Start start = "start"

type start string

func (start) Rule() Rule      { return "start" }
func (start) textAlignValue() {}

func (right) textAlignValue() {}

const MatchParent matchParent = "match-parent"

type matchParent string

func (matchParent) Rule() Rule      { return "match-parent" }
func (matchParent) textAlignValue() {}

func (left) textAlignValue() {}

const Justify justify = "justify"

type justify string

func (justify) Rule() Rule      { return "justify" }
func (justify) textAlignValue() {}

const End end = "end"

type end string

func (end) Rule() Rule      { return "end" }
func (end) textAlignValue() {}

func (center) textAlignValue() {}

func (stringType) textAlignValue() {}

func (unset) textAlignValue() {}

func (initial) textAlignValue() {}

func (inherit) textAlignValue() {}

func (flexStart) justifyContentValue() {}

func (spaceBetween) justifyContentValue() {}

func (spaceEvenly) justifyContentValue() {}

func (spaceAround) justifyContentValue() {}

func (flexEnd) justifyContentValue() {}

func (center) justifyContentValue() {}

func (unset) justifyContentValue() {}

func (initial) justifyContentValue() {}

func (inherit) justifyContentValue() {}

const Ltr ltr = "ltr"

type ltr string

func (ltr) Rule() Rule      { return "ltr" }
func (ltr) directionValue() {}

const Rtl rtl = "rtl"

type rtl string

func (rtl) Rule() Rule      { return "rtl" }
func (rtl) directionValue() {}

func (unset) directionValue() {}

func (initial) directionValue() {}

func (inherit) directionValue() {}

func (normal) fontLanguageOverrideValue() {}

func (stringType) fontLanguageOverrideValue() {}

func (unset) fontLanguageOverrideValue() {}

func (initial) fontLanguageOverrideValue() {}

func (inherit) fontLanguageOverrideValue() {}

func (unset) gridTemplateValue() {}

func (initial) gridTemplateValue() {}

func (inherit) gridTemplateValue() {}

type breakvalueType string

func (s breakvalueType) String() string { return string(s) }
func (breakvalueType) breakValue()      {}

func (unset) breakValue() {}

func (initial) breakValue() {}

func (inherit) breakValue() {}

func (auto) breakInsideValue() {}

const AvoidPage avoidPage = "avoid-page"

type avoidPage string

func (avoidPage) Rule() Rule        { return "avoid-page" }
func (avoidPage) breakInsideValue() {}

const AvoidColumn avoidColumn = "avoid-column"

type avoidColumn string

func (avoidColumn) Rule() Rule        { return "avoid-column" }
func (avoidColumn) breakInsideValue() {}

func (avoid) breakInsideValue() {}

func (unset) breakInsideValue() {}

func (initial) breakInsideValue() {}

func (inherit) breakInsideValue() {}

func (normal) normalOrUnitOrAutoValue() {}

func (lengthType) normalOrUnitOrAutoValue() {}

func (unset) normalOrUnitOrAutoValue() {}

func (initial) normalOrUnitOrAutoValue() {}

func (inherit) normalOrUnitOrAutoValue() {}

const Slice slice = "slice"

type slice string

func (slice) Rule() Rule               { return "slice" }
func (slice) boxDecorationBreakValue() {}

const Clone clone = "clone"

type clone string

func (clone) Rule() Rule               { return "clone" }
func (clone) boxDecorationBreakValue() {}

func (unset) boxDecorationBreakValue() {}

func (initial) boxDecorationBreakValue() {}

func (inherit) boxDecorationBreakValue() {}

func (none) gridTemplateValue() {}

func (gridautoType) gridTemplateValue() {}

func (auto) lineBreakValue() {}

const Strict strict = "strict"

type strict string

func (strict) Rule() Rule      { return "strict" }
func (strict) lineBreakValue() {}

func (normal) lineBreakValue() {}

const Loose loose = "loose"

type loose string

func (loose) Rule() Rule      { return "loose" }
func (loose) lineBreakValue() {}

func (unset) lineBreakValue() {}

func (initial) lineBreakValue() {}

func (inherit) lineBreakValue() {}

func (integerType) uintValue() {}

const Invert invert = "invert"

type invert string

func (invert) Rule() Rule  { return "invert" }
func (invert) colorValue() {}

func (auto) textJustifyValue() {}

func (none) textJustifyValue() {}

const InterWord interWord = "inter-word"

type interWord string

func (interWord) Rule() Rule        { return "inter-word" }
func (interWord) textJustifyValue() {}

const Distribute distribute = "distribute"

type distribute string

func (distribute) Rule() Rule        { return "distribute" }
func (distribute) textJustifyValue() {}

func (unset) textJustifyValue() {}

func (initial) textJustifyValue() {}

func (inherit) textJustifyValue() {}

func (unset) transitionValue() {}

func (initial) transitionValue() {}

func (inherit) transitionValue() {}

func (normal) columnGapValue() {}

func (lengthType) columnGapValue() {}

func (unset) columnGapValue() {}

func (initial) columnGapValue() {}

func (inherit) columnGapValue() {}

const Shadow shadow = "shadow"

type shadow string

func (shadow) Rule() Rule   { return "shadow" }
func (shadow) shadowValue() {}

func (none) animationNameValue() {}

type identifierType string

func (s identifierType) String() string    { return string(s) }
func (identifierType) animationNameValue() {}

func (unset) animationNameValue() {}

func (initial) animationNameValue() {}

func (inherit) animationNameValue() {}

func (auto) cursorValue() {}

const ZoomOut zoomOut = "zoom-out"

type zoomOut string

func (zoomOut) Rule() Rule   { return "zoom-out" }
func (zoomOut) cursorValue() {}

const ZoomIn zoomIn = "zoom-in"

type zoomIn string

func (zoomIn) Rule() Rule   { return "zoom-in" }
func (zoomIn) cursorValue() {}

const Wait wait = "wait"

type wait string

func (wait) Rule() Rule   { return "wait" }
func (wait) cursorValue() {}

const WResize wResize = "w-resize"

type wResize string

func (wResize) Rule() Rule   { return "w-resize" }
func (wResize) cursorValue() {}

const VerticalText verticalText = "vertical-text"

type verticalText string

func (verticalText) Rule() Rule   { return "vertical-text" }
func (verticalText) cursorValue() {}

func (urlType) cursorValue() {}

func (text) cursorValue() {}

const SwResize swResize = "sw-resize"

type swResize string

func (swResize) Rule() Rule   { return "sw-resize" }
func (swResize) cursorValue() {}

const SeResize seResize = "se-resize"

type seResize string

func (seResize) Rule() Rule   { return "se-resize" }
func (seResize) cursorValue() {}

const SResize sResize = "s-resize"

type sResize string

func (sResize) Rule() Rule   { return "s-resize" }
func (sResize) cursorValue() {}

const RowResize rowResize = "row-resize"

type rowResize string

func (rowResize) Rule() Rule   { return "row-resize" }
func (rowResize) cursorValue() {}

const Progress progress = "progress"

type progress string

func (progress) Rule() Rule   { return "progress" }
func (progress) cursorValue() {}

const Pointer pointer = "pointer"

type pointer string

func (pointer) Rule() Rule   { return "pointer" }
func (pointer) cursorValue() {}

const NwseResize nwseResize = "nwse-resize"

type nwseResize string

func (nwseResize) Rule() Rule   { return "nwse-resize" }
func (nwseResize) cursorValue() {}

const NwResize nwResize = "nw-resize"

type nwResize string

func (nwResize) Rule() Rule   { return "nw-resize" }
func (nwResize) cursorValue() {}

const NsResize nsResize = "ns-resize"

type nsResize string

func (nsResize) Rule() Rule   { return "ns-resize" }
func (nsResize) cursorValue() {}

const NotAllowed notAllowed = "not-allowed"

type notAllowed string

func (notAllowed) Rule() Rule   { return "not-allowed" }
func (notAllowed) cursorValue() {}

func (none) cursorValue() {}

const NoDrop noDrop = "no-drop"

type noDrop string

func (noDrop) Rule() Rule   { return "no-drop" }
func (noDrop) cursorValue() {}

const NeswResize neswResize = "nesw-resize"

type neswResize string

func (neswResize) Rule() Rule   { return "nesw-resize" }
func (neswResize) cursorValue() {}

const NeResize neResize = "ne-resize"

type neResize string

func (neResize) Rule() Rule   { return "ne-resize" }
func (neResize) cursorValue() {}

const NResize nResize = "n-resize"

type nResize string

func (nResize) Rule() Rule   { return "n-resize" }
func (nResize) cursorValue() {}

const Move move = "move"

type move string

func (move) Rule() Rule   { return "move" }
func (move) cursorValue() {}

const Help help = "help"

type help string

func (help) Rule() Rule   { return "help" }
func (help) cursorValue() {}

const EwResize ewResize = "ew-resize"

type ewResize string

func (ewResize) Rule() Rule   { return "ew-resize" }
func (ewResize) cursorValue() {}

const EResize eResize = "e-resize"

type eResize string

func (eResize) Rule() Rule   { return "e-resize" }
func (eResize) cursorValue() {}

const Default defaultValue = "default"

type defaultValue string

func (defaultValue) Rule() Rule   { return "default" }
func (defaultValue) cursorValue() {}

const Crosshair crosshair = "crosshair"

type crosshair string

func (crosshair) Rule() Rule   { return "crosshair" }
func (crosshair) cursorValue() {}

const Copy copy = "copy"

type copy string

func (copy) Rule() Rule   { return "copy" }
func (copy) cursorValue() {}

const ContextMenu contextMenu = "context-menu"

type contextMenu string

func (contextMenu) Rule() Rule   { return "context-menu" }
func (contextMenu) cursorValue() {}

const ColResize colResize = "col-resize"

type colResize string

func (colResize) Rule() Rule   { return "col-resize" }
func (colResize) cursorValue() {}

const Cell cell = "cell"

type cell string

func (cell) Rule() Rule   { return "cell" }
func (cell) cursorValue() {}

const AllScroll allScroll = "all-scroll"

type allScroll string

func (allScroll) Rule() Rule   { return "all-scroll" }
func (allScroll) cursorValue() {}

const Alias alias = "alias"

type alias string

func (alias) Rule() Rule   { return "alias" }
func (alias) cursorValue() {}

func (unset) cursorValue() {}

func (initial) cursorValue() {}

func (inherit) cursorValue() {}

func (none) floatValue() {}

func (left) floatValue() {}

func (right) floatValue() {}

func (unset) floatValue() {}

func (initial) floatValue() {}

func (inherit) floatValue() {}

func (normal) fontStyleValue() {}

const Oblique oblique = "oblique"

type oblique string

func (oblique) Rule() Rule      { return "oblique" }
func (oblique) fontStyleValue() {}

const Italic italic = "italic"

type italic string

func (italic) Rule() Rule      { return "italic" }
func (italic) fontStyleValue() {}

func (unset) fontStyleValue() {}

func (initial) fontStyleValue() {}

func (inherit) fontStyleValue() {}

func (normal) animationDirectionValue() {}

const Reverse reverse = "reverse"

type reverse string

func (reverse) Rule() Rule               { return "reverse" }
func (reverse) animationDirectionValue() {}

const AlternateReverse alternateReverse = "alternate-reverse"

type alternateReverse string

func (alternateReverse) Rule() Rule               { return "alternate-reverse" }
func (alternateReverse) animationDirectionValue() {}

const Alternate alternate = "alternate"

type alternate string

func (alternate) Rule() Rule               { return "alternate" }
func (alternate) animationDirectionValue() {}

func (unset) animationDirectionValue() {}

func (initial) animationDirectionValue() {}

func (inherit) animationDirectionValue() {}

func (none) unitOrAutoValue() {}

func (all) unitOrAutoValue() {}

func (none) filterValue() {}

type filterModeType string

func (s filterModeType) String() string { return string(s) }
func (filterModeType) filterValue()     {}

func (urlType) filterValue() {}

func (unset) filterValue() {}

func (initial) filterValue() {}

func (inherit) filterValue() {}

func (auto) pointerEventsValue() {}

func (none) pointerEventsValue() {}

func (unset) pointerEventsValue() {}

func (initial) pointerEventsValue() {}

func (inherit) pointerEventsValue() {}

func (unset) borderLeftValue() {}

func (initial) borderLeftValue() {}

func (inherit) borderLeftValue() {}

type featuretagvalueType string

func (s featuretagvalueType) String() string          { return string(s) }
func (featuretagvalueType) fontFeatureSettingsValue() {}

func (unset) fontFeatureSettingsValue() {}

func (initial) fontFeatureSettingsValue() {}

func (inherit) fontFeatureSettingsValue() {}

func (normal) fontWeightValue() {}

const Lighter lighter = "lighter"

type lighter string

func (lighter) Rule() Rule       { return "lighter" }
func (lighter) fontWeightValue() {}

const Bolder bolder = "bolder"

type bolder string

func (bolder) Rule() Rule       { return "bolder" }
func (bolder) fontWeightValue() {}

const Bold bold = "bold"

type bold string

func (bold) Rule() Rule       { return "bold" }
func (bold) fontWeightValue() {}

func (integerType) fontWeightValue() {}

func (unset) fontWeightValue() {}

func (initial) fontWeightValue() {}

func (inherit) fontWeightValue() {}

const Clip clip = "clip"

type clip string

func (clip) Rule() Rule         { return "clip" }
func (clip) textOverflowValue() {}

const Ellipsis ellipsis = "ellipsis"

type ellipsis string

func (ellipsis) Rule() Rule         { return "ellipsis" }
func (ellipsis) textOverflowValue() {}

func (stringType) textOverflowValue() {}

func (unset) textOverflowValue() {}

func (initial) textOverflowValue() {}

func (inherit) textOverflowValue() {}

func (medium) columnRuleWidthValue() {}

func (thin) columnRuleWidthValue() {}

func (thick) columnRuleWidthValue() {}

func (lengthType) columnRuleWidthValue() {}

func (zero) columnRuleWidthValue() {}

func (unset) columnRuleWidthValue() {}

func (initial) columnRuleWidthValue() {}

func (inherit) columnRuleWidthValue() {}

func (none) animationFillModeValue() {}

const Forwards forwards = "forwards"

type forwards string

func (forwards) Rule() Rule              { return "forwards" }
func (forwards) animationFillModeValue() {}

func (both) animationFillModeValue() {}

const Backwards backwards = "backwards"

type backwards string

func (backwards) Rule() Rule              { return "backwards" }
func (backwards) animationFillModeValue() {}

func (unset) animationFillModeValue() {}

func (initial) animationFillModeValue() {}

func (inherit) animationFillModeValue() {}

func (top) captionSideValue() {}

func (bottom) captionSideValue() {}

func (unset) captionSideValue() {}

func (initial) captionSideValue() {}

func (inherit) captionSideValue() {}

func (normal) wordWrapValue() {}

func (breakWord) wordWrapValue() {}

func (unset) wordWrapValue() {}

func (initial) wordWrapValue() {}

func (inherit) wordWrapValue() {}

const Running running = "running"

type running string

func (running) Rule() Rule               { return "running" }
func (running) animationPlayStateValue() {}

const Paused paused = "paused"

type paused string

func (paused) Rule() Rule               { return "paused" }
func (paused) animationPlayStateValue() {}

func (unset) animationPlayStateValue() {}

func (initial) animationPlayStateValue() {}

func (inherit) animationPlayStateValue() {}

const Seperate seperate = "seperate"

type seperate string

func (seperate) Rule() Rule           { return "seperate" }
func (seperate) borderCollapseValue() {}

const Collapse collapse = "collapse"

type collapse string

func (collapse) Rule() Rule           { return "collapse" }
func (collapse) borderCollapseValue() {}

func (unset) borderCollapseValue() {}

func (initial) borderCollapseValue() {}

func (inherit) borderCollapseValue() {}

func (none) textDecorationLineValue() {}

const Underline underline = "underline"

type underline string

func (underline) Rule() Rule               { return "underline" }
func (underline) textDecorationLineValue() {}

const Overline overline = "overline"

type overline string

func (overline) Rule() Rule               { return "overline" }
func (overline) textDecorationLineValue() {}

const LineThrough lineThrough = "line-through"

type lineThrough string

func (lineThrough) Rule() Rule               { return "line-through" }
func (lineThrough) textDecorationLineValue() {}

const Blink blink = "blink"

type blink string

func (blink) Rule() Rule               { return "blink" }
func (blink) textDecorationLineValue() {}

func (unset) textDecorationLineValue() {}

func (initial) textDecorationLineValue() {}

func (inherit) textDecorationLineValue() {}

func (timeType) transitionDurationValue() {}

func (unset) transitionDurationValue() {}

func (initial) transitionDurationValue() {}

func (inherit) transitionDurationValue() {}

func (normal) fontVariantValue() {}

func (unicase) fontVariantValue() {}

func (titlingCaps) fontVariantValue() {}

func (smallCaps) fontVariantValue() {}

func (petiteCaps) fontVariantValue() {}

func (allSmallCaps) fontVariantValue() {}

func (allPetiteCaps) fontVariantValue() {}

func (unset) fontVariantValue() {}

func (initial) fontVariantValue() {}

func (inherit) fontVariantValue() {}

func (visible) visibilityValue() {}

func (hidden) visibilityValue() {}

func (collapse) visibilityValue() {}

func (unset) visibilityValue() {}

func (initial) visibilityValue() {}

func (inherit) visibilityValue() {}

func (unset) backgroundValue() {}

func (initial) backgroundValue() {}

func (inherit) backgroundValue() {}

func (none) counterIncrementValue() {}

func (unset) counterIncrementValue() {}

func (initial) counterIncrementValue() {}

func (inherit) counterIncrementValue() {}

func (auto) normalOrAutoValue() {}

func (normal) normalOrAutoValue() {}

func (none) normalOrAutoValue() {}

func (unset) normalOrAutoValue() {}

func (initial) normalOrAutoValue() {}

func (inherit) normalOrAutoValue() {}

func (unset) listStyleValue() {}

func (initial) listStyleValue() {}

func (inherit) listStyleValue() {}

func (length) unitOrAutoValue() {}

func (none) hangingPunctuationValue() {}

const Last last = "last"

type last string

func (last) Rule() Rule               { return "last" }
func (last) hangingPunctuationValue() {}

const ForceEnd forceEnd = "force-end"

type forceEnd string

func (forceEnd) Rule() Rule               { return "force-end" }
func (forceEnd) hangingPunctuationValue() {}

const First first = "first"

type first string

func (first) Rule() Rule               { return "first" }
func (first) hangingPunctuationValue() {}

const AllowEnd allowEnd = "allow-end"

type allowEnd string

func (allowEnd) Rule() Rule               { return "allow-end" }
func (allowEnd) hangingPunctuationValue() {}

func (unset) hangingPunctuationValue() {}

func (initial) hangingPunctuationValue() {}

func (inherit) hangingPunctuationValue() {}

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

const Static static = "static"

type static string

func (static) Rule() Rule     { return "static" }
func (static) positionValue() {}

const Sticky sticky = "sticky"

type sticky string

func (sticky) Rule() Rule     { return "sticky" }
func (sticky) positionValue() {}

const Relative relative = "relative"

type relative string

func (relative) Rule() Rule     { return "relative" }
func (relative) positionValue() {}

const Page page = "page"

type page string

func (page) Rule() Rule     { return "page" }
func (page) positionValue() {}

func (fixed) positionValue() {}

func (center) positionValue() {}

const Absolute absolute = "absolute"

type absolute string

func (absolute) Rule() Rule     { return "absolute" }
func (absolute) positionValue() {}

func (unset) positionValue() {}

func (initial) positionValue() {}

func (inherit) positionValue() {}

func (auto) textUnderlinePositionValue() {}

const Under under = "under"

type under string

func (under) Rule() Rule                  { return "under" }
func (under) textUnderlinePositionValue() {}

func (right) textUnderlinePositionValue() {}

func (left) textUnderlinePositionValue() {}

func (unset) textUnderlinePositionValue() {}

func (initial) textUnderlinePositionValue() {}

func (inherit) textUnderlinePositionValue() {}

func (normal) wordSpacingValue() {}

func (unitType) wordSpacingValue() {}

func (unset) wordSpacingValue() {}

func (initial) wordSpacingValue() {}

func (inherit) wordSpacingValue() {}

func (stretch) alignItemsValue() {}

func (flexStart) alignItemsValue() {}

func (flexEnd) alignItemsValue() {}

func (center) alignItemsValue() {}

func (baseline) alignItemsValue() {}

func (unset) alignItemsValue() {}

func (initial) alignItemsValue() {}

func (inherit) alignItemsValue() {}

func (auto) clipValue() {}

func (unset) clipValue() {}

func (initial) clipValue() {}

func (inherit) clipValue() {}

func (unset) columnRuleValue() {}

func (initial) columnRuleValue() {}

func (inherit) columnRuleValue() {}

func (transparent) sizeValue() {}

func (colorType) sizeValue() {}

func (currentColor) sizeValue() {}

func (normal) unicodeBidiValue() {}

const Embed embed = "embed"

type embed string

func (embed) Rule() Rule        { return "embed" }
func (embed) unicodeBidiValue() {}

const BidiOverride bidiOverride = "bidi-override"

type bidiOverride string

func (bidiOverride) Rule() Rule        { return "bidi-override" }
func (bidiOverride) unicodeBidiValue() {}

func (unset) unicodeBidiValue() {}

func (initial) unicodeBidiValue() {}

func (inherit) unicodeBidiValue() {}
