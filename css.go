package seed

//CSS is the bundled, default css that is packaged with Qlovaseed.
const CSS = `
	/* CSS reset */
	html, body, div, span, applet, object, iframe,
	h1, h2, h3, h4, h5, h6, p, blockquote, pre,
	a, abbr, acronym, address, big, cite, code,
	del, dfn, em, img, ins, kbd, q, s, samp,
	small, strike, strong, sub, sup, tt, var,
	b, u, i, center,
	dl, dt, dd, ol, ul, li,
	fieldset, form, label, legend,
	table, caption, tbody, tfoot, thead, tr, th, td,
	article, aside, canvas, details, embed, 
	figure, figcaption, footer, header, hgroup, 
	menu, nav, output, ruby, section, summary,
	time, mark, audio, video {
		margin: 0;
		padding: 0;
		border: 0;
		font-size: 100%;
		font: inherit;
		vertical-align: baseline;
	}
	/* HTML5 display-role reset for older browsers */
	article, aside, details, figcaption, figure, 
	footer, header, hgroup, menu, nav, section {
		display: block;
	}
	body {
		line-height: 1;
	}
	ol, ul {
		list-style: none;
	}
	blockquote, q {
		quotes: none;
	}
	blockquote:before, blockquote:after,
	q:before, q:after {
		content: '';
		content: none;
	}
	table {
		border-collapse: collapse;
		border-spacing: 0;
	}

	* {
		-webkit-tap-highlight-color: rgba(255, 255, 255, 0) !important; 
		-webkit-focus-ring-color: rgba(255, 255, 255, 0) !important; 
		outline: none !important;
		flex-shrink: 0;
		cursor: default;
	}

	input[type=range]::-moz-focus-outer {
		border: 0;
	}

	a {
		text-decoration: none;
	}
	
	img {
		object-fit: contain;
		max-width: 100%;
	}
	
	html {
		height: 100%;
		box-sizing: border-box;
	}
	*, *:before, *:after {
		box-sizing: border-box;
	}
	pre, hr {
		margin: 0;
	}

	div {
		position: relative;
	}

	#install_instructions {
		position: fixed;
		display: flex;
		left: 0px;
		top: 0px;
		width: 100vw;
		flex-wrap: wrap;
		height: 100vh;
		background-color: white;
		align-content: center;
		justify-content: center;
		align-items: center;
		justify-items: center;
		font-family: sans-serif;
		color: #343333;
	}

	
	body {
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		position: fixed;
		overscroll-behavior: none; 
		-webkit-overscroll-behavior: none; 
		-webkit-overflow-scrolling: none; 
		cursor: pointer; 
		margin: 0; 
		padding: 0;
		height: 100%;
		width: 100vw;
		
		/* TODO small screens only? */
		-webkit-touch-callout: none;
		-webkit-user-select: none;
		-khtml-user-select: none;
		-moz-user-select: none;
		-ms-user-select: none;
		user-select: none;
		-webkit-tap-highlight-color: transparent;
		
		/* Some nice defaults for centering content. */
		display: inline-flex;
		align-items: center;
		justify-content: center;
		flex-direction: row;
		overflow: hidden;
	}
`
