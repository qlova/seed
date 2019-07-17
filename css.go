package seed

const CSS = `
	* {
		-webkit-tap-highlight-color: rgba(255, 255, 255, 0) !important; 
		-webkit-focus-ring-color: rgba(255, 255, 255, 0) !important; 
		outline: none !important;
		flex-shrink: 0;
	}

	a {
		text-decoration: none;
	}
	
	p {
		margin-block-start: 0;
		margin-block-end: 0;
	}
	
	img {
		object-fit: contain;
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

	<!-- TODO These animations should be added dynamically by seed/script/animation.go -->
	@keyframes slideInFromLeft {
		from { transform: translateX(-100%); }
		to { transform: translateX(0); }
	}

	@keyframes slideInFromRight {
		from { transform: translateX(100%); }
		to { transform: translateX(0); }
	}

	@keyframes slideInFromTop {
		from { transform: translateY(-100%); }
		to { transform: translateY(0); }
	}

	@keyframes slideInFromBottom {
		from { transform: translateY(100%); }
		to { transform: translateY(0); }
	}

	@keyframes slideOutToLeft {
		from { transform: translateX(0); }
		to { transform: translateX(-100%); }
	}

	@keyframes slideOutToRight {
		from { transform: translateX(); }
		to { transform: translateX(100%); }
	}

	@keyframes slideOutToTop {
		from { transform: translateY(0); }
		to { transform: translateY(-100%); }
	}

	@keyframes slideOutToBottom {
		from { transform: translateY(0); }
		to { transform: translateY(100%); }
	}
	
	@keyframes fadeOut {
		from { opacity: 1; }
		to { opacity: 0; }
	}
	
	@keyframes fadeIn {
		from { opacity: 0; }
		to { opacity: 1; }
	}
	
	body {
		top: 0;
		left: 0
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
