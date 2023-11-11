import{s as I,n as x,r as D,o as E,b as V}from"../chunks/scheduler.fdb56d66.js";import{S as j,i as H,g as m,s as S,H as q,h as v,j as C,I as z,c as k,f as b,J as O,k as g,a as M,x as l,B as P}from"../chunks/index.8124e3d0.js";import{W as N}from"../chunks/webrtc.1240ba7f.js";const L=`<!DOCTYPE html>
<html lang="en-us">
	<head>
		<meta charset="utf-8" />
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<title>Unity WebGL Player | SejongsCipher</title>
		<link rel="shortcut icon" href="tetris/TemplateData/favicon.ico" />
		<link rel="stylesheet" href="tetris/TemplateData/style.css" />
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

			var buildUrl = 'tetris/Build';
			var loaderUrl = buildUrl + '/20231111_0919_HanSrangCodingClan_Build.loader.js';
			var config = {
				dataUrl: buildUrl + '/20231111_0919_HanSrangCodingClan_Build.data.br',
				frameworkUrl: buildUrl + '/20231111_0919_HanSrangCodingClan_Build.framework.js.br',
				codeUrl: buildUrl + '/20231111_0919_HanSrangCodingClan_Build.wasm.br',
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
`,R=async()=>({page:L}),Y=Object.freeze(Object.defineProperty({__proto__:null,load:R},Symbol.toStringTag,{value:"Module"}));function W(r){let n,e,a,d="시작",s,i,w="상태",y,t,p,u,f,_=r[0].page+"",B,T;return{c(){n=m("div"),e=m("div"),a=m("button"),a.textContent=d,s=S(),i=m("button"),i.textContent=w,y=S(),t=m("video"),p=S(),u=m("div"),f=new q(!1),this.h()},l(c){n=v(c,"DIV",{class:!0});var o=C(n);e=v(o,"DIV",{class:!0});var h=C(e);a=v(h,"BUTTON",{"data-svelte-h":!0}),z(a)!=="svelte-bx336s"&&(a.textContent=d),s=k(h),i=v(h,"BUTTON",{"data-svelte-h":!0}),z(i)!=="svelte-6pmyjs"&&(i.textContent=w),y=k(h),t=v(h,"VIDEO",{id:!0,width:!0,height:!0}),C(t).forEach(b),h.forEach(b),p=k(o),u=v(o,"DIV",{class:!0});var U=C(u);f=O(U,!1),U.forEach(b),o.forEach(b),this.h()},h(){g(t,"id","localVideo"),t.autoplay=!0,t.muted=!0,g(t,"width","600"),g(t,"height","400"),g(e,"class","videoArea svelte-1svz6r2"),f.a=null,g(u,"class","game"),g(n,"class","container svelte-1svz6r2")},m(c,o){M(c,n,o),l(n,e),l(e,a),l(e,s),l(e,i),l(e,y),l(e,t),r[5](t),l(n,p),l(n,u),f.m(_,u),B||(T=[P(a,"click",r[3]),P(i,"click",r[4])],B=!0)},p(c,[o]){o&1&&_!==(_=c[0].page+"")&&f.p(_)},i:x,o:x,d(c){c&&b(n),r[5](null),B=!1,D(T)}}}function A(r,n,e){let{data:a}=n,d,s;E(()=>{const p=new URLSearchParams(window.location.search).get("roomId");e(2,s=new N(d,p,"tetris"))});const i=()=>{s.start()},w=()=>{s.currentState()};function y(t){V[t?"unshift":"push"](()=>{d=t,e(1,d)})}return r.$$set=t=>{"data"in t&&e(0,a=t.data)},[a,d,s,i,w,y]}class K extends j{constructor(n){super(),H(this,n,A,W,I,{data:0})}}export{K as component,Y as universal};
