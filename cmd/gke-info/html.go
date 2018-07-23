/**
# Copyright 2015 Google Inc. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
**/

package main

const (
	html = `<!doctype html>
<html>
<head>

<script>
	// addEventListener support for IE8
	function bindEvent(element, eventName, eventHandler) {
			if (element.addEventListener) {
					element.addEventListener(eventName, eventHandler, false);
			} else if (element.attachEvent) {
					element.attachEvent('on' + eventName, eventHandler);
			}
	}

	// Listen to messages from parent window
	bindEvent(window, 'message', function (e) {
			if (e.data=="refresh") {
				window.location.reload();	
			}
	});
</script>

<!-- Compiled and minified CSS -->
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/0.97.0/css/materialize.min.css">

<!-- Compiled and minified JavaScript -->
<title>Frontend Web Server</title>
</head>
<body>
<div class="container">
<div class="row">
<div>


<div class="card {{.Color}}">
<div class="card-content white-text">
<div class="card-title">Current App Version</div>
</div>
<div class="card-content white">
<table class="bordered">
  <tbody>
	<tr>
	  <td><center><h3><b>{{.Version}}</b></h3></center></td>
	</tr>
  </tbody>
</table>
</div>
</div>

</div>
</div>
</div>
</div>
</html>`
)
