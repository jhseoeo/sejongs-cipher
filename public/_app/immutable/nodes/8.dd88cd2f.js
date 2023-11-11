import{s as U,n as C,o as T,b as x}from"../chunks/scheduler.fdb56d66.js";import{S as P,i as I,g as m,s as S,H as q,h as v,j as b,I as z,c as k,f as y,J as D,k as u,a as E,x as h,B as V}from"../chunks/index.8124e3d0.js";import{W as H}from"../chunks/webrtc.487b9709.js";const j=`<!DOCTYPE html>
<html lang="en-us">
	<head>
		<meta charset="utf-8" />
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<title>Unity WebGL Player | SejongsCipher</title>
		<link rel="shortcut icon" href="wordguess/TemplateData/favicon.ico" />
		<link rel="stylesheet" href="wordguess/TemplateData/style.css" />
	</head>
	<body>
		<div id="unity-container" class="unity-desktop">
			<canvas id="unity-canvas" width="1920" height="1080" tabindex="-1"></canvas>
			<div id="unity-loading-bar">
				<div id="unity-logo"></div>
				<div id="unity-progress-bar-empty">
					<div id="unity-progress-bar-full"></div>
				</div>
			</div>
			<div id="unity-warning"></div>
			<div id="unity-footer">
				<div id="unity-webgl-logo"></div>
				<div id="unity-fullscreen-button"></div>
				<div id="unity-build-title">SejongsCipher</div>
			</div>
		</div>
		<script>
			var container = document.querySelector('#unity-container');
			var canvas = document.querySelector('#unity-canvas');
			var loadingBar = document.querySelector('#unity-loading-bar');
			var progressBarFull = document.querySelector('#unity-progress-bar-full');
			var fullscreenButton = document.querySelector('#unity-fullscreen-button');
			var warningBanner = document.querySelector('#unity-warning');

			// Shows a temporary message banner/ribbon for a few seconds, or
			// a permanent error message on top of the canvas if type=='error'.
			// If type=='warning', a yellow highlight color is used.
			// Modify or remove this function to customize the visually presented
			// way that non-critical warnings and error messages are presented to the
			// user.
			function unityShowBanner(msg, type) {
				function updateBannerVisibility() {
					warningBanner.style.display = warningBanner.children.length ? 'block' : 'none';
				}
				var div = document.createElement('div');
				div.innerHTML = msg;
				warningBanner.appendChild(div);
				if (type == 'error') div.style = 'background: red; padding: 10px;';
				else {
					if (type == 'warning') div.style = 'background: yellow; padding: 10px;';
					setTimeout(function () {
						warningBanner.removeChild(div);
						updateBannerVisibility();
					}, 5000);
				}
				updateBannerVisibility();
			}

			var buildUrl = 'wordguess/Build';
			var loaderUrl = buildUrl + '/20231111_0931_HanSrangCodingClan_Build.loader.js';
			var config = {
				dataUrl: buildUrl + '/20231111_0931_HanSrangCodingClan_Build.data.br',
				frameworkUrl: buildUrl + '/20231111_0931_HanSrangCodingClan_Build.framework.js.br',
				codeUrl: buildUrl + '/20231111_0931_HanSrangCodingClan_Build.wasm.br',
				streamingAssetsUrl: 'StreamingAssets',
				companyName: 'HanSrangCodingClan',
				productName: 'SejongsCipher',
				productVersion: '1.0',
				showBanner: unityShowBanner,
			};

			// By default, Unity keeps WebGL canvas render target size matched with
			// the DOM size of the canvas element (scaled by window.devicePixelRatio)
			// Set this to false if you want to decouple this synchronization from
			// happening inside the engine, and you would instead like to size up
			// the canvas DOM size and WebGL render target sizes yourself.
			// config.matchWebGLToCanvasSize = false;

			if (/iPhone|iPad|iPod|Android/i.test(navigator.userAgent)) {
				// Mobile device style: fill the whole browser client area with the game canvas:

				var meta = document.createElement('meta');
				meta.name = 'viewport';
				meta.content =
					'width=device-width, height=device-height, initial-scale=1.0, user-scalable=no, shrink-to-fit=yes';
				document.getElementsByTagName('head')[0].appendChild(meta);
				container.className = 'unity-mobile';
				canvas.className = 'unity-mobile';

				// To lower canvas resolution on mobile devices to gain some
				// performance, uncomment the following line:
				// config.devicePixelRatio = 1;
			} else {
				// Desktop style: Render the game canvas in a window that can be maximized to fullscreen:

				canvas.style.width = '1920px';
				canvas.style.height = '1080px';
			}

			loadingBar.style.display = 'block';

			var script = document.createElement('script');
			script.src = loaderUrl;
			script.onload = () => {
				createUnityInstance(canvas, config, (progress) => {
					progressBarFull.style.width = 100 * progress + '%';
				})
					.then((unityInstance) => {
						loadingBar.style.display = 'none';
						fullscreenButton.onclick = () => {
							unityInstance.SetFullscreen(1);
						};
					})
					.catch((message) => {
						alert(message);
					});
			};

			document.body.appendChild(script);
		<\/script>
	</body>
</html>
`,M=async()=>({page:j}),A=Object.freeze(Object.defineProperty({__proto__:null,load:M},Symbol.toStringTag,{value:"Module"}));function O(s){let t,n,i,o="시작",l,a,g,e,d,p=s[0].page+"",w,_;return{c(){t=m("div"),n=m("div"),i=m("button"),i.textContent=o,l=S(),a=m("video"),g=S(),e=m("div"),d=new q(!1),this.h()},l(c){t=v(c,"DIV",{class:!0});var r=b(t);n=v(r,"DIV",{class:!0});var f=b(n);i=v(f,"BUTTON",{"data-svelte-h":!0}),z(i)!=="svelte-bx336s"&&(i.textContent=o),l=k(f),a=v(f,"VIDEO",{id:!0,width:!0,height:!0}),b(a).forEach(y),f.forEach(y),g=k(r),e=v(r,"DIV",{class:!0});var B=b(e);d=D(B,!1),B.forEach(y),r.forEach(y),this.h()},h(){u(a,"id","localVideo"),a.autoplay=!0,a.muted=!0,u(a,"width","600"),u(a,"height","400"),u(n,"class","videoArea svelte-1qhsord"),d.a=null,u(e,"class","game"),u(t,"class","container svelte-1qhsord")},m(c,r){E(c,t,r),h(t,n),h(n,i),h(n,l),h(n,a),s[4](a),h(t,g),h(t,e),d.m(p,e),w||(_=V(i,"click",s[3]),w=!0)},p(c,[r]){r&1&&p!==(p=c[0].page+"")&&d.p(p)},i:C,o:C,d(c){c&&y(t),s[4](null),w=!1,_()}}}function L(s,t,n){let{data:i}=t,o,l;T(()=>{const d=new URLSearchParams(window.location.search).get("roomId");n(2,l=new H(o,d,"wordguess"))});const a=()=>{l.start()};function g(e){x[e?"unshift":"push"](()=>{o=e,n(1,o)})}return s.$$set=e=>{"data"in e&&n(0,i=e.data)},[i,o,l,a,g]}class G extends P{constructor(t){super(),I(this,t,L,O,U,{data:0})}}export{G as component,A as universal};
