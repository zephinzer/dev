package assets

const OAuthTokenPage=`
<html>
  <head>
    <meta charset="utf-8" />
    <title>
      dev: ${OAUTH_PLATFORM} OAuth Token
    </title>
    <style>
      html, body {
        background-color: rgba(0,0,0,0.01);
        padding: 0;
        margin: 0;
      }
      .overlay {
        align-items: center;
        color: #444;
        display: flex;
        font-family: sans-serif;
        font-size: 48px;
        height: 100%;
        justify-content: center;
        width: 100%;
      }
      .information {
        display: flex;
        flex-direction: column;
        text-align: center;
        overflow: hidden;
      }
      .icon {
        display: block;
        height: 128px;
        margin: 40px 24px;
        max-height: 100%;
        max-width: 100%;
      }
      .icon img {
        background-color: #FFF;
        border-radius: 100%;
        border-color: #666;
        border-width: 4px;
        border-style: solid;
        box-shadow: 0px 0px 64px rgba(0,0,0,0.15);
        max-height: 100%;
        padding: 4px;
      }
      .token {
        background: rgba(0,0,0,0.01);
        border: none;
        border-radius: 64px;
        color: #AAA;
        font-family: monospace;
        font-size: 48px;
        text-align: center;
        transition: 0.3s;
        width: 100%;
      }
      .token:active, .token:focus, .token:hover {
        background: transparent;
        border: none;
        color: #888;
        outline: none !important;
        transition: 0.3s;
      }
      .token-caption {
        color: #AAA;
        font-size: 24px;
        font-style: italic;
      }
    </style>
  </head>
  <body>
    <div class="overlay">
      <div class="information">
        <div class="icon">
          <img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAgAAAAIACAMAAADDpiTIAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAA2hpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuMy1jMDExIDY2LjE0NTY2MSwgMjAxMi8wMi8wNi0xNDo1NjoyNyAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wTU09Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9tbS8iIHhtbG5zOnN0UmVmPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvc1R5cGUvUmVzb3VyY2VSZWYjIiB4bWxuczp4bXA9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC8iIHhtcE1NOk9yaWdpbmFsRG9jdW1lbnRJRD0ieG1wLmRpZDpGOTdGMTE3NDA3MjA2ODExQTAxRkUwN0JENDA5NjgxOCIgeG1wTU06RG9jdW1lbnRJRD0ieG1wLmRpZDpCRjQwQjY2QTg2MDExMUVBOTJGM0I2NDA3MDIzNEY0NyIgeG1wTU06SW5zdGFuY2VJRD0ieG1wLmlpZDpCRjQwQjY2OTg2MDExMUVBOTJGM0I2NDA3MDIzNEY0NyIgeG1wOkNyZWF0b3JUb29sPSJBZG9iZSBQaG90b3Nob3AgQ1M2IChNYWNpbnRvc2gpIj4gPHhtcE1NOkRlcml2ZWRGcm9tIHN0UmVmOmluc3RhbmNlSUQ9InhtcC5paWQ6Rjk3RjExNzQwNzIwNjgxMUEwMUZFMDdCRDQwOTY4MTgiIHN0UmVmOmRvY3VtZW50SUQ9InhtcC5kaWQ6Rjk3RjExNzQwNzIwNjgxMUEwMUZFMDdCRDQwOTY4MTgiLz4gPC9yZGY6RGVzY3Jp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" />
        </div>
        Your ${OAUTH_PLATFORM} token is:<br />
        <input
          class="token"
          id="token"
          onfocus="d()"
          onblur="e()"
          onkeydown="()=>{}"
          type="password"
          value="${OAUTH_ACCESS_TOKEN}"
        />
        <span class="token-caption">(click above to reveal/copy)</span>
      </div>
    </div>
    <script>
      function d() {
        var token = document.getElementById('token');
        token.setAttribute("type", "text");
      }
      function e() {
        var token = document.getElementById('token');
        token.setAttribute("type", "password");
      }
    </script>
  </body>
</html>
`