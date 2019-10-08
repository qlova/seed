package main

//<unit> = "<percentage>", "<length>"
//<box> = "padding-box", "content-box", "border-box"
//<style> = "none", "solid", "ridge", "outset", "inset", "hidden", "groove", "double", "dotted", "dashed"
//<thickness> = "medium", "thin", "thick", "<length>", "0"
//<breakvalue> = "auto", "right", "page", "left", "column", "avoid-page", "avoid-column", "avoid", "always"
//<filter> = "sepia(<number/percentage>)","saturate(<number/percentage>)","opacity(<number/percentage>)","invert(<number/percentage>)","hue-rotate(<angle>)","grayscale(<number/percentage>)","drop-shadow(<length> <color>)","contrast(<number/percentage>)","brightness(<number/percentage>)","blur(<length>)"
/*

<timingfunction> = "ease", "steps(<integer>, start)", "steps(<integer>, end)", "cubic-bezier(<number>, <number>, <number>, <number>)", "step-start", "step-end", "linear", "ease-out", "ease-in-out", "ease-in"

<east-asian-variant-values> = [ jis78 | jis83 | jis90 | jis04 | simplified | traditional ]
<east-asian-width-values> = [ full-width | proportional-width ]

<common-lig-values> = [ common-ligatures | no-common-ligatures ]
<discretionary-lig-values> = [ discretionary-ligatures | no-discretionary-ligatures ]
<historical-lig-values> = [ historical-ligatures | no-historical-ligatures ]
<contextual-alt-values> = [ contextual | no-contextual ]

<numeric-figure-values> = [ lining-nums | oldstyle-nums ]
<numeric-spacing-values> = [ proportional-nums | tabular-nums ]
<numeric-fraction-values> = [ diagonal-fractions | stacked-fractions ]

<gridauto> = "auto", "max-content", "min-content", "<length>"
<gridstop> = "auto", "span(<integer>)", "column-line"
<overflow> =
<pagebreak> = "auto","right","left","avoid","always"

<transformation> =
	<translateZ(<length>)
	translateY(<translation-value>)
	translateX(<translation-value>)
	translate3d(<translation-value>, <translation-value>, <length>)
	translate(<translation-value>, <translation-value>)
	skewY(<angle>)
	skewX(<angle>)
	scaleZ(<number>)
	scaleY(<number>)
	scaleX(<number>)
	scale3d(<number>, <number>, <number>)
	scale(<number>)
	rotateZ(<angle>)
	rotateY(<angle>)
	rotateX(<angle>)
	rotate3d(<number>, <number>, <number>, <angle>)
	rotate(<angle>)
	perspective(<length>)
	matrix3d([16 comma-separated <number> values])
	matrix([<number>, <number>, <number>, <number>, <number>, <number>])

*/

var Properties = map[string][]string{
	"align-content":              {"stretch", "space-between", "space-around", "space-evenly", "flex-start", "flex-end", "center"},
	"align-items":                {"stretch", "flex-start", "flex-end", "center", "baseline"},
	"align-self":                 {"auto", "stretch", "flex-start", "flex-end", "center", "baseline"},
	"all":                        {},
	"animation":                  {"<shorthand>"},
	"animation-delay":            {"<time>"},
	"animation-direction":        {"normal", "reverse", "alternate-reverse", "alternate"},
	"animation-duration":         {"<time>"},
	"animation-fill-mode":        {"none", "forwards", "both", "backwards"},
	"animation-iteration-count":  {"<number>", "infinite"},
	"animation-name":             {"none", "<identifier>"},
	"animation-play-state":       {"running", "paused"},
	"animation-timing-function":  {"<timingfunction>"},
	"backface-visibility":        {"visible", "hidden"},
	"background":                 {"<shorthand>"},
	"background-attachment":      {"scroll", "local", "fixed"},
	"background-blend-mode":      {"normal", "soft-light", "screen", "saturation", "overlay", "multiply", "luminosity", "lighten", "hue", "hard-light", "exclusion", "difference", "darken", "color-dodge", "color-burn", "color"},
	"background-clip":            {"<box>"},
	"background-color":           {"transparent", "<color>"},
	"background-image":           {"none", "<url>", "<gradient>", "<image>"},
	"background-origin":          {"<box>"},
	"background-position":        {"<position>"},
	"background-repeat":          {"repeat", "space", "round", "repeat-y", "repeat-x", "repeat no-repeat", "no-repeat"},
	"background-size":            {"<size>"},
	"border":                     {"<shorthand>"},
	"border-bottom":              {"<shorthand>"},
	"border-bottom-color":        {"transparent", "<color>", "currentColor"},
	"border-bottom-left-radius":  {"<unit>"},
	"border-bottom-right-radius": {"<unit>"},
	"border-bottom-style":        {"none", "solid", "ridge", "outset", "inset", "hidden", "groove", "double", "dotted", "dashed"},
	"border-bottom-width":        {"medium", "thin", "thick", "<length>", "0"},
	"border-collapse":            {"seperate", "collapse"},
	"border-color":               {"transparent", "<color>", "currentColor"},
	"border-image":               {"<shorthand>"},
	"border-image-outset":        {"[1-4]<number/percentage>"},
	"border-image-repeat":        {"stretch", "space", "round", "repeat"},
	"border-image-slice":         {"[1-4]<number/percentage>"},
	"border-image-source":        {"none", "<url>"},
	"border-image-width":         {"[1-4]<number/percentage>/length"},
	"border-left":                {"<shorthand>"},
	"border-left-color":          {"transparent", "<color>", "currentColor"},
	"border-left-style":          {"none", "solid", "ridge", "outset", "inset", "hidden", "groove", "double", "dotted", "dashed"},
	"border-left-width":          {"medium", "thin", "thick", "<length>", "0"},
	"border-radius":              {"<shorthand>"},
	"border-right":               {"<shorthand>"},
	"border-right-color":         {"transparent", "<color>", "currentColor"},
	"border-right-style":         {"none", "solid", "ridge", "outset", "inset", "hidden", "groove", "double", "dotted", "dashed"},
	"border-right-width":         {"medium", "thin", "thick", "<length>", "0"},
	"border-spacing":             {"[1-2]<length>"},
	"border-style":               {"<shorthand>"},
	"border-top":                 {"<shorthand>"},
	"border-top-color":           {"transparent", "<color>", "currentColor"},
	"border-top-left-radius":     {"<unit>"},
	"border-top-right-radius":    {"<unit>"},
	"border-top-style":           {"none", "solid", "ridge", "outset", "inset", "hidden", "groove", "double", "dotted", "dashed"},
	"border-top-width":           {"medium", "thin", "thick", "<length>", "0"},
	"border-width":               {"<shorthand>"},
	"bottom":                     {"auto", "<unit>", "0"},
	"box-decoration-break":       {"slice", "clone"},
	"box-shadow":                 {"none", "<shadow>"}, //inset?
	"box-sizing":                 {"<box>"},
	"break-after":                {"<breakvalue>"},
	"break-before":               {"<breakvalue>"},
	"break-inside":               {"auto", "avoid-page", "avoid-column", "avoid"},
	"caption-side":               {"top", "bottom"},
	"clear":                      {"none", "right", "left", "both"},
	"clip":                       {"auto", "rect(<length>, <length>, <length>, <length>)", "inset(<length>, <length>, <length>, <length>)"},
	"color":                      {"<color>"},
	"column-count":               {"auto", "<integer>"},
	"column-fill":                {"balance", "auto"},
	"column-gap":                 {"normal", "<length>"},
	"column-rule":                {"<shorthand>"},
	"column-rule-color":          {"<color>"},
	"column-rule-style":          {"none", "solid", "ridge", "outset", "inset", "hidden", "groove", "double", "dotted", "dashed"},
	"column-rule-width":          {"medium", "thin", "thick", "<length>", "0"},
	"column-span":                {"none", "all"},
	"column-width":               {"auto", "length"},
	"columns":                    {"<shorthand>"},
	"content":                    {"normal", "open-quote", "none", "no-open-quote", "no-close-quote", "icon", "close-quote", "attr(<identifier>)", "<url>", "<string>", "<counter>"},
	"counter-increment":          {"none", "<identifier> <integer>"},
	"counter-reset":              {"none", "<identifier> <integer>"},
	"cursor":                     {"auto", "zoom-out", "zoom-in", "wait", "w-resize", "vertical-text", "<url>", "text", "sw-resize", "se-resize", "s-resize", "row-resize", "progress", "pointer", "nwse-resize", "nw-resize", "ns-resize", "not-allowed", "none", "no-drop", "nesw-resize", "ne-resize", "n-resize", "move", "help", "ew-resize", "e-resize", "default", "crosshair", "copy", "context-menu", "col-resize", "cell", "all-scroll", "alias"},
	"direction":                  {"ltr", "rtl"},
	"display":                    {"inline", "table-row-group", "table-row", "table-header-group", "table-footer-group", "table-column-group", "table-column", "table-cell", "table-caption", "table", "run-in", "none", "list-item", "inline-table", "inline-flex", "inline-block", "flex", "container", "compact", "block"},
	"empty-cells":                {"show", "hide"},
	"filter":                     {"none", "<filter>", "<url>"},
	"flex":                       {"<shorthand>"},
	"flex-basis":                 {"auto", "length"},
	"flex-direction":             {"row", "row-reverse", "column-reverse", "column"},
	"flex-flow":                  {"<shorthand>"},
	"flex-grow":                  {"<number>"},
	"flex-shrink":                {"<number>"},
	"flex-wrap":                  {"nowrap", "wrap", "wrap-reverse"},
	"float":                      {"none", "left", "right"},
	"font":                       {"<shorthand>"},
	"@font-face":                 {},
	"font-family":                {"<familyname>"},
	"font-feature-settings":      {"<featuretagvalue>"},
	"@font-feature-values":       {},
	"font-display":               {"auto", "block", "swap", "fallback", "optional"},
	"font-kerning":               {"auto", "normal", "none"},
	"font-language-override":     {"normal", "<string>"},
	"font-size":                  {"medium", "xx-small", "xx-large", "x-small", "x-large", "smaller", "small", "larger", "large", "<unit>"},
	"font-size-adjust":           {"none", "<number>"},
	"font-stretch":               {"normal", "ultra-expanded", "ultra-condensed", "semi-expanded", "semi-condensed", "extra-expanded", "extra-condensed", "expanded", "condensed"},
	"font-style":                 {"normal", "oblique", "italic"},
	"font-synthesis":             {"weight style", "weight", "style", "none"},
	"font-variant":               {"normal", "unicase", "titling-caps", "small-caps", "petite-caps", "all-small-caps", "all-petite-caps"},
	"font-variant-alternates":    {"normal", "historical-forms", "stylistic(<feature-value-name>)", "styleset(<feature-value-name>)", "character-variant(<feature-value-name>)", "swash(<feature-value-name>)", "ornaments(<feature-value-name>)", "annotation(<feature-value-name>)"},
	"font-variant-caps":          {"normal", "unicase", "titling-caps", "small-caps", "petite-caps", "all-small-caps", "all-petite-caps"},
	"font-variant-east-asian":    {"normal", "<east-asian-variant-values>  <east-asian-width-values> ruby"},
	"font-variant-ligatures":     {"normal", "none", "<common-lig-values> <discretionary-lig-values> <historical-lig-values> <contextual-alt-values>"},
	"font-variant-numeric":       {"normal", "<numeric-figure-values> <numeric-spacing-values> <numeric-fraction-values> ordinal slashed-zero"},
	"font-variant-position":      {"normal", "sub", "super"},
	"font-weight":                {"normal", "lighter", "bolder", "bold", "<integer>"},
	"grid":                       {"<shorthand>"},
	"grid-area":                  {"<shorthand>"},
	"grid-auto-columns":          {"<gridauto>"},
	"grid-auto-flow":             {"row", "column", "dense", "row dense", "column dense"},
	"grid-auto-rows":             {"<gridauto>"},
	"grid-column":                {"row", "column", "dense", "row dense", "column dense"},
	"grid-column-end":            {"<gridstop>"},
	"grid-column-gap":            {"<length>"},
	"grid-column-start":          {"<gridstop>"},
	"grid-gap":                   {"<shorthand>"},
	"grid-row":                   {"<shorthand>"},
	"grid-row-end":               {"<gridstop>"},
	"grid-row-gap":               {"<length>"},
	"grid-row-start":             {"<gridstop>"},
	"grid-template":              {"<shorthand>"},
	"grid-template-areas":        {"none", "[1-n]<identifier>"},
	"grid-template-columns":      {"none", "<gridauto>"},
	"grid-template-rows":         {"none", "<gridauto>"},
	"hanging-punctuation":        {"none", "last force-end", "last allow-end", "last", "force-end", "first force-end", "first allow-end", "first", "allow-end"},
	"height":                     {"auto", "<unit>"},
	"hyphens":                    {"maunal", "none", "auto"},
	"image-rendering":            {"auto", "pixelated", "crisp-edges"},
	"@import":                    {},
	"isolation":                  {"auto", "isolate"},
	"justify-content":            {"flex-start", "space-between", "space-evenly", "space-around", "flex-end", "center"},
	"@keyframes":                 {},
	"left":                       {"auto", "<unit>", "0"},
	"letter-spacing":             {"normal", "<length>"},
	"line-break":                 {"auto", "strict", "normal", "loose"},
	"line-height":                {"normal", "<unit>", "<number>"},
	"list-style":                 {"<shorthand>"},
	"list-style-image":           {"none", "<url>"},
	"list-style-position":        {"outside", "inside"},
	"list-style-type":            {"disc", "upper-roman", "upper-latin", "upper-alpha", "square", "none", "lower-roman", "lower-latin", "lower-greek", "lower-alpha", "georgian", "decimal-leading-zero", "decimal", "circle", "armenian"},
	"margin":                     {"<shorthand>", "auto", "<unit>", "0"},
	"margin-bottom":              {"auto", "<unit>", "0"},
	"margin-left":                {"auto", "<unit>", "0"},
	"margin-right":               {"auto", "<unit>", "0"},
	"margin-top":                 {"auto", "<unit>", "0"},
	"max-height":                 {"none", "<unit>"},
	"max-width":                  {"none", "<unit>"},
	"@media":                     {},
	"min-height":                 {"none", "<unit>"},
	"min-width":                  {"none", "<unit>"},
	"mix-blend-mode":             {"normal", "soft-light", "screen", "saturation", "overlay", "multiply", "luminosity", "lighten", "hue", "hard-light", "exclusion", "difference", "darken", "color-dodge", "color-burn", "color"},
	"object-fit":                 {"fill", "scale-down", "none", "cover", "contain"},
	"object-position":            {"<position>"},
	"opacity":                    {"<number>"},
	"order":                      {"<integer>"},
	"orphans":                    {"<integer>"},
	"outline":                    {"<shorthand>"},
	"outline-color":              {"invert", "<color>"},
	"outline-offset":             {"<length>"},
	"outline-style":              {"none", "solid", "ridge", "outset", "inset", "hidden", "groove", "double", "dotted", "dashed"},
	"outline-width":              {"medium", "thin", "thick", "<length>", "0"},
	"overflow":                   {"visible", "scroll", "hidden", "auto"},
	"overflow-wrap":              {"normal", "break-word"},
	"overflow-x":                 {"visible", "scroll", "hidden", "auto"},
	"overflow-y":                 {"visible", "scroll", "hidden", "auto"},
	"padding":                    {"<shorthand>", "0"},
	"padding-bottom":             {"<unit>", "0"},
	"padding-left":               {"<unit>", "0"},
	"padding-right":              {"<unit>", "0"},
	"padding-top":                {"<unit>", "0"},
	"page-break-after":           {"<pagebreak>"},
	"page-break-before":          {"<pagebreak>"},
	"page-break-inside":          {"auto", "avoid"},
	"perspective":                {"none", "<length>"},
	"perspective-origin":         {"<position>"},
	"pointer-events":             {"auto", "none"},
	"position":                   {"static", "sticky", "relative", "page", "fixed", "center", "absolute"},
	"quotes":                     {"none", "[1-2]<string>"},
	"resize":                     {"none", "vertical", "horizontal", "both"},
	"right":                      {"auto", "<unit>", "0"},
	"scroll-behavior":            {"auto", "smooth"},
	"tab-size":                   {"<length>", "<integer>"},
	"table-layout":               {"auto", "fixed"},
	"text-align":                 {"start end", "start", "right", "match-parent", "left", "justify", "end", "center", "<string>"},
	"text-align-last":            {"auto", "start", "right", "left", "justify", "end", "center"},
	"text-combine-upright":       {"none", "digits(<integer>)", "all"},
	"text-decoration":            {"<shorthand>"},
	"text-decoration-color":      {"currentColor", "<color>"},
	"text-decoration-line":       {"none", "underline", "overline", "line-through", "blink"},
	"text-decoration-style":      {"solid", "wavy", "double", "dotted", "dashed"},
	"text-indent":                {"<unit>", "hanging-each-line(<length>)", "hanging(<length>)", "each-line(<length>)"},
	"text-justify":               {"auto", "none", "inter-word", "distribute"},
	"text-orientation":           {"mixed", "use-glyph-orientation", "upright", "sideways-right", "sideways-left", "sideways"},
	"text-overflow":              {"clip", "ellipsis", "<string>"},
	"text-shadow":                {"none", "shadow"},
	"text-transform":             {"none", "uppercase", "lowercase", "full-width", "capitalize"},
	"text-underline-position":    {"auto", "under right", "under left", "under", "right", "left"},
	"top":                        {"auto", "<unit>", "0"},
	"transform":                  {"none", "<transformation>"},
	"transform-origin":           {"<position>"},
	"transform-style":            {"flat", "preserve-3d"},
	"transition":                 {"<shorthand>"},
	"transition-delay":           {"<time>"},
	"transition-duration":        {"<time>"},
	"transition-property":        {"all", "none", "[1-n]<property>"},
	"transition-timing-function": {},
	"unicode-bidi":               {"normal", "embed", "bidi-override"},
	"user-select":                {"auto", "none", "text", "all"},
	"vertical-align":             {"baseline", "top", "text-top", "text-bottom", "super", "sub", "middle", "bottom", "<unit>"},
	"visibility":                 {"visible", "hidden", "collapse"},
	"white-space":                {"normal", "pre-wrap", "pre-line", "pre", "nowrap"},
	"widows":                     {"<integer>"},
	"width":                      {"auto", "<unit>"},
	"will-change":                {"[1-n]<property>"},
	"word-break":                 {"normal", "keep-all", "break-all"},
	"word-spacing":               {"normal", "<unit>"},
	"word-wrap":                  {"normal", "break-word"},
	"writing-mode":               {"horizontal-tb", "vertical-rl", "vertical-lr"},
	"z-index":                    {"auto", "<integer>"},
}
