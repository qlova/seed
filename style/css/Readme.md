# Type-safe CSS Bindings for Go
 
This package is a mostly complete set of type-safe CSS bindings for the Go programming language.  

Use this like you would use CSS, instead of dashes use captilisation.  
Eg. z-index becomes ZIndex, margin-width becomes MarginWidth.  

Example:  

```
	var style = css.NewStyle()
	style.SetMarginWidth(css.Number(100).Px()) //Equivalent to margin-width: 100px;
	fmt.Println(style.MarginWidth()) //100px
```

If you cannot figure out how to do something type-safe, report it as an issue.  
You can also use raw CSS setting if needed.

Example:  

```
	var style = css.NewStyle()
	style.Set("margin", "20% 20%") //Equivalent to margin: 20% 20%;
	fmt.Println(style.MarginWidth()) //20% 20%
```

The package works nicely with Go's new wasm support, you can retrieve the style of an element like this:  

Example:  

```
	var style = css.StyleOf(js.Global().Get("document").Get("body"))
	style.SetBackgroundColor(css.Blue) //Equivalent to background-color: blue;
	fmt.Println(style.BackgroundColor()) //This will return the computed style value of backgroundColor
```
