package app

const builtinCSS = `
*, *:before, *:after {
    box-sizing: border-box;
    position: relative;
    max-width: 100vw;
    object-fit: contain;
}

body {
	width: 100vw;
	height: 100%;
	position: fixed;
	user-select: none;
	overflow: hidden;
	align-items: center;
	line-height: 1;
}
`
