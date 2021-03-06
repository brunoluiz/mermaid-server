package tmpl

import "html/template"

// IndexHTML Index page template
const IndexHTML = `<!doctype html>
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<title>Mermaid Live Server</title>
</head>
<body>
	<div class="mermaid">
	{{ .Content }}
	</div>
	<div class="error"></div>
	<script src="http://localhost:35729/livereload.js"></script>
	<script src="https://unpkg.com/mermaid@8.8.1/dist/mermaid.min.js"></script>
	<script>
		mermaid.initialize({
			startOnLoad:true,
			...{{ .Config }}
		});

		// mermaid.parseError = function(err, hash) {
		//   document.getElementByClassName('mermaid').innerHTML = err
		// }
	</script>
</body>
</html>`

// IndexHTMLSubs Index page template subs
type IndexHTMLSubs struct {
	Content string
	Config  template.JS
}
