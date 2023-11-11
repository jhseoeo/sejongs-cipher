import{x as q,s as le,l as X,m as ue,i as M,p as ee,b as ce,c as de,u as fe,g as me,d as pe,q as ne}from"./scheduler.fdb56d66.js";import{S as he,i as be,q as W,r as V,e as U,u as ve,v as Z,a as F,p as ge,t as L,w as J,b as _e,d as k,f as C,g as ye,s as Le,h as Ae,j as Se,c as ke,k as Ce}from"./index.8124e3d0.js";import{c as H,g as K,a as v,f as De,S as Q}from"./SmuiElement.f3162490.js";function Be(n,e,t,i={bubbles:!0},r=!1){if(typeof Event>"u")throw new Error("Event not defined.");if(!n)throw new Error("Tried to dipatch event without element.");const o=new CustomEvent(e,Object.assign(Object.assign({},i),{detail:t}));if(n==null||n.dispatchEvent(o),r&&e.startsWith("SMUI")){const l=new CustomEvent(e.replace(/^SMUI/g,()=>"MDC"),Object.assign(Object.assign({},i),{detail:t}));n==null||n.dispatchEvent(l),l.defaultPrevented&&o.preventDefault()}return o}var z;function Ge(n,e){e===void 0&&(e=!1);var t=n.CSS,i=z;if(typeof z=="boolean"&&!e)return z;var r=t&&typeof t.supports=="function";if(!r)return!1;var o=t.supports("--css-vars","yes"),l=t.supports("(--css-vars: yes)")&&t.supports("color","#00000000");return i=o||l,e||(z=i),i}function Oe(n,e,t){if(!n)return{x:0,y:0};var i=e.x,r=e.y,o=i+t.left,l=r+t.top,a,s;if(n.type==="touchstart"){var u=n;a=u.changedTouches[0].pageX-o,s=u.changedTouches[0].pageY-l}else{var f=n;a=f.pageX-o,s=f.pageY-l}return{x:a,y:s}}var te=function(n,e){return te=Object.setPrototypeOf||{__proto__:[]}instanceof Array&&function(t,i){t.__proto__=i}||function(t,i){for(var r in i)Object.prototype.hasOwnProperty.call(i,r)&&(t[r]=i[r])},te(n,e)};function Ie(n,e){if(typeof e!="function"&&e!==null)throw new TypeError("Class extends value "+String(e)+" is not a constructor or null");te(n,e);function t(){this.constructor=n}n.prototype=e===null?Object.create(e):(t.prototype=e.prototype,new t)}var O=function(){return O=Object.assign||function(e){for(var t,i=1,r=arguments.length;i<r;i++){t=arguments[i];for(var o in t)Object.prototype.hasOwnProperty.call(t,o)&&(e[o]=t[o])}return e},O.apply(this,arguments)};function j(n){var e=typeof Symbol=="function"&&Symbol.iterator,t=e&&n[e],i=0;if(t)return t.call(n);if(n&&typeof n.length=="number")return{next:function(){return n&&i>=n.length&&(n=void 0),{value:n&&n[i++],done:!n}}};throw new TypeError(e?"Object is not iterable.":"Symbol.iterator is not defined.")}/**
 * @license
 * Copyright 2016 Google Inc.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */var ze=function(){function n(e){e===void 0&&(e={}),this.adapter=e}return Object.defineProperty(n,"cssClasses",{get:function(){return{}},enumerable:!1,configurable:!0}),Object.defineProperty(n,"strings",{get:function(){return{}},enumerable:!1,configurable:!0}),Object.defineProperty(n,"numbers",{get:function(){return{}},enumerable:!1,configurable:!0}),Object.defineProperty(n,"defaultAdapter",{get:function(){return{}},enumerable:!1,configurable:!0}),n.prototype.init=function(){},n.prototype.destroy=function(){},n}();/**
 * @license
 * Copyright 2019 Google Inc.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */function je(n){return n===void 0&&(n=window),Pe(n)?{passive:!0}:!1}function Pe(n){n===void 0&&(n=window);var e=!1;try{var t={get passive(){return e=!0,!1}},i=function(){};n.document.addEventListener("test",i,t),n.document.removeEventListener("test",i,t)}catch{e=!1}return e}const Ee=Object.freeze(Object.defineProperty({__proto__:null,applyPassive:je},Symbol.toStringTag,{value:"Module"}));/**
 * @license
 * Copyright 2018 Google Inc.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */function qe(n,e){if(n.closest)return n.closest(e);for(var t=n;t;){if(we(t,e))return t;t=t.parentElement}return null}function we(n,e){var t=n.matches||n.webkitMatchesSelector||n.msMatchesSelector;return t.call(n,e)}function Xe(n){var e=n;if(e.offsetParent!==null)return e.scrollWidth;var t=e.cloneNode(!0);t.style.setProperty("position","absolute"),t.style.setProperty("transform","translate(-9999px, -9999px)"),document.documentElement.appendChild(t);var i=t.scrollWidth;return document.documentElement.removeChild(t),i}const We=Object.freeze(Object.defineProperty({__proto__:null,closest:qe,estimateScrollWidth:Xe,matches:we},Symbol.toStringTag,{value:"Module"}));/**
 * @license
 * Copyright 2016 Google Inc.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */var Ve={BG_FOCUSED:"mdc-ripple-upgraded--background-focused",FG_ACTIVATION:"mdc-ripple-upgraded--foreground-activation",FG_DEACTIVATION:"mdc-ripple-upgraded--foreground-deactivation",ROOT:"mdc-ripple-upgraded",UNBOUNDED:"mdc-ripple-upgraded--unbounded"},Ze={VAR_FG_SCALE:"--mdc-ripple-fg-scale",VAR_FG_SIZE:"--mdc-ripple-fg-size",VAR_FG_TRANSLATE_END:"--mdc-ripple-fg-translate-end",VAR_FG_TRANSLATE_START:"--mdc-ripple-fg-translate-start",VAR_LEFT:"--mdc-ripple-left",VAR_TOP:"--mdc-ripple-top"},ie={DEACTIVATION_TIMEOUT_MS:225,FG_DEACTIVATION_MS:150,INITIAL_ORIGIN_SCALE:.6,PADDING:10,TAP_DELAY_MS:300};/**
 * @license
 * Copyright 2016 Google Inc.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */var re=["touchstart","pointerdown","mousedown","keydown"],se=["touchend","pointerup","mouseup","contextmenu"],P=[],Je=function(n){Ie(e,n);function e(t){var i=n.call(this,O(O({},e.defaultAdapter),t))||this;return i.activationAnimationHasEnded=!1,i.activationTimer=0,i.fgDeactivationRemovalTimer=0,i.fgScale="0",i.frame={width:0,height:0},i.initialSize=0,i.layoutFrame=0,i.maxRadius=0,i.unboundedCoords={left:0,top:0},i.activationState=i.defaultActivationState(),i.activationTimerCallback=function(){i.activationAnimationHasEnded=!0,i.runDeactivationUXLogicIfReady()},i.activateHandler=function(r){i.activateImpl(r)},i.deactivateHandler=function(){i.deactivateImpl()},i.focusHandler=function(){i.handleFocus()},i.blurHandler=function(){i.handleBlur()},i.resizeHandler=function(){i.layout()},i}return Object.defineProperty(e,"cssClasses",{get:function(){return Ve},enumerable:!1,configurable:!0}),Object.defineProperty(e,"strings",{get:function(){return Ze},enumerable:!1,configurable:!0}),Object.defineProperty(e,"numbers",{get:function(){return ie},enumerable:!1,configurable:!0}),Object.defineProperty(e,"defaultAdapter",{get:function(){return{addClass:function(){},browserSupportsCssVars:function(){return!0},computeBoundingRect:function(){return{top:0,right:0,bottom:0,left:0,width:0,height:0}},containsEventTarget:function(){return!0},deregisterDocumentInteractionHandler:function(){},deregisterInteractionHandler:function(){},deregisterResizeHandler:function(){},getWindowPageOffset:function(){return{x:0,y:0}},isSurfaceActive:function(){return!0},isSurfaceDisabled:function(){return!0},isUnbounded:function(){return!0},registerDocumentInteractionHandler:function(){},registerInteractionHandler:function(){},registerResizeHandler:function(){},removeClass:function(){},updateCssVariable:function(){}}},enumerable:!1,configurable:!0}),e.prototype.init=function(){var t=this,i=this.supportsPressRipple();if(this.registerRootHandlers(i),i){var r=e.cssClasses,o=r.ROOT,l=r.UNBOUNDED;requestAnimationFrame(function(){t.adapter.addClass(o),t.adapter.isUnbounded()&&(t.adapter.addClass(l),t.layoutInternal())})}},e.prototype.destroy=function(){var t=this;if(this.supportsPressRipple()){this.activationTimer&&(clearTimeout(this.activationTimer),this.activationTimer=0,this.adapter.removeClass(e.cssClasses.FG_ACTIVATION)),this.fgDeactivationRemovalTimer&&(clearTimeout(this.fgDeactivationRemovalTimer),this.fgDeactivationRemovalTimer=0,this.adapter.removeClass(e.cssClasses.FG_DEACTIVATION));var i=e.cssClasses,r=i.ROOT,o=i.UNBOUNDED;requestAnimationFrame(function(){t.adapter.removeClass(r),t.adapter.removeClass(o),t.removeCssVars()})}this.deregisterRootHandlers(),this.deregisterDeactivationHandlers()},e.prototype.activate=function(t){this.activateImpl(t)},e.prototype.deactivate=function(){this.deactivateImpl()},e.prototype.layout=function(){var t=this;this.layoutFrame&&cancelAnimationFrame(this.layoutFrame),this.layoutFrame=requestAnimationFrame(function(){t.layoutInternal(),t.layoutFrame=0})},e.prototype.setUnbounded=function(t){var i=e.cssClasses.UNBOUNDED;t?this.adapter.addClass(i):this.adapter.removeClass(i)},e.prototype.handleFocus=function(){var t=this;requestAnimationFrame(function(){return t.adapter.addClass(e.cssClasses.BG_FOCUSED)})},e.prototype.handleBlur=function(){var t=this;requestAnimationFrame(function(){return t.adapter.removeClass(e.cssClasses.BG_FOCUSED)})},e.prototype.supportsPressRipple=function(){return this.adapter.browserSupportsCssVars()},e.prototype.defaultActivationState=function(){return{activationEvent:void 0,hasDeactivationUXRun:!1,isActivated:!1,isProgrammatic:!1,wasActivatedByPointer:!1,wasElementMadeActive:!1}},e.prototype.registerRootHandlers=function(t){var i,r;if(t){try{for(var o=j(re),l=o.next();!l.done;l=o.next()){var a=l.value;this.adapter.registerInteractionHandler(a,this.activateHandler)}}catch(s){i={error:s}}finally{try{l&&!l.done&&(r=o.return)&&r.call(o)}finally{if(i)throw i.error}}this.adapter.isUnbounded()&&this.adapter.registerResizeHandler(this.resizeHandler)}this.adapter.registerInteractionHandler("focus",this.focusHandler),this.adapter.registerInteractionHandler("blur",this.blurHandler)},e.prototype.registerDeactivationHandlers=function(t){var i,r;if(t.type==="keydown")this.adapter.registerInteractionHandler("keyup",this.deactivateHandler);else try{for(var o=j(se),l=o.next();!l.done;l=o.next()){var a=l.value;this.adapter.registerDocumentInteractionHandler(a,this.deactivateHandler)}}catch(s){i={error:s}}finally{try{l&&!l.done&&(r=o.return)&&r.call(o)}finally{if(i)throw i.error}}},e.prototype.deregisterRootHandlers=function(){var t,i;try{for(var r=j(re),o=r.next();!o.done;o=r.next()){var l=o.value;this.adapter.deregisterInteractionHandler(l,this.activateHandler)}}catch(a){t={error:a}}finally{try{o&&!o.done&&(i=r.return)&&i.call(r)}finally{if(t)throw t.error}}this.adapter.deregisterInteractionHandler("focus",this.focusHandler),this.adapter.deregisterInteractionHandler("blur",this.blurHandler),this.adapter.isUnbounded()&&this.adapter.deregisterResizeHandler(this.resizeHandler)},e.prototype.deregisterDeactivationHandlers=function(){var t,i;this.adapter.deregisterInteractionHandler("keyup",this.deactivateHandler);try{for(var r=j(se),o=r.next();!o.done;o=r.next()){var l=o.value;this.adapter.deregisterDocumentInteractionHandler(l,this.deactivateHandler)}}catch(a){t={error:a}}finally{try{o&&!o.done&&(i=r.return)&&i.call(r)}finally{if(t)throw t.error}}},e.prototype.removeCssVars=function(){var t=this,i=e.strings,r=Object.keys(i);r.forEach(function(o){o.indexOf("VAR_")===0&&t.adapter.updateCssVariable(i[o],null)})},e.prototype.activateImpl=function(t){var i=this;if(!this.adapter.isSurfaceDisabled()){var r=this.activationState;if(!r.isActivated){var o=this.previousActivationEvent,l=o&&t!==void 0&&o.type!==t.type;if(!l){r.isActivated=!0,r.isProgrammatic=t===void 0,r.activationEvent=t,r.wasActivatedByPointer=r.isProgrammatic?!1:t!==void 0&&(t.type==="mousedown"||t.type==="touchstart"||t.type==="pointerdown");var a=t!==void 0&&P.length>0&&P.some(function(s){return i.adapter.containsEventTarget(s)});if(a){this.resetActivationState();return}t!==void 0&&(P.push(t.target),this.registerDeactivationHandlers(t)),r.wasElementMadeActive=this.checkElementMadeActive(t),r.wasElementMadeActive&&this.animateActivation(),requestAnimationFrame(function(){P=[],!r.wasElementMadeActive&&t!==void 0&&(t.key===" "||t.keyCode===32)&&(r.wasElementMadeActive=i.checkElementMadeActive(t),r.wasElementMadeActive&&i.animateActivation()),r.wasElementMadeActive||(i.activationState=i.defaultActivationState())})}}}},e.prototype.checkElementMadeActive=function(t){return t!==void 0&&t.type==="keydown"?this.adapter.isSurfaceActive():!0},e.prototype.animateActivation=function(){var t=this,i=e.strings,r=i.VAR_FG_TRANSLATE_START,o=i.VAR_FG_TRANSLATE_END,l=e.cssClasses,a=l.FG_DEACTIVATION,s=l.FG_ACTIVATION,u=e.numbers.DEACTIVATION_TIMEOUT_MS;this.layoutInternal();var f="",m="";if(!this.adapter.isUnbounded()){var g=this.getFgTranslationCoordinates(),_=g.startPoint,d=g.endPoint;f=_.x+"px, "+_.y+"px",m=d.x+"px, "+d.y+"px"}this.adapter.updateCssVariable(r,f),this.adapter.updateCssVariable(o,m),clearTimeout(this.activationTimer),clearTimeout(this.fgDeactivationRemovalTimer),this.rmBoundedActivationClasses(),this.adapter.removeClass(a),this.adapter.computeBoundingRect(),this.adapter.addClass(s),this.activationTimer=setTimeout(function(){t.activationTimerCallback()},u)},e.prototype.getFgTranslationCoordinates=function(){var t=this.activationState,i=t.activationEvent,r=t.wasActivatedByPointer,o;r?o=Oe(i,this.adapter.getWindowPageOffset(),this.adapter.computeBoundingRect()):o={x:this.frame.width/2,y:this.frame.height/2},o={x:o.x-this.initialSize/2,y:o.y-this.initialSize/2};var l={x:this.frame.width/2-this.initialSize/2,y:this.frame.height/2-this.initialSize/2};return{startPoint:o,endPoint:l}},e.prototype.runDeactivationUXLogicIfReady=function(){var t=this,i=e.cssClasses.FG_DEACTIVATION,r=this.activationState,o=r.hasDeactivationUXRun,l=r.isActivated,a=o||!l;a&&this.activationAnimationHasEnded&&(this.rmBoundedActivationClasses(),this.adapter.addClass(i),this.fgDeactivationRemovalTimer=setTimeout(function(){t.adapter.removeClass(i)},ie.FG_DEACTIVATION_MS))},e.prototype.rmBoundedActivationClasses=function(){var t=e.cssClasses.FG_ACTIVATION;this.adapter.removeClass(t),this.activationAnimationHasEnded=!1,this.adapter.computeBoundingRect()},e.prototype.resetActivationState=function(){var t=this;this.previousActivationEvent=this.activationState.activationEvent,this.activationState=this.defaultActivationState(),setTimeout(function(){return t.previousActivationEvent=void 0},e.numbers.TAP_DELAY_MS)},e.prototype.deactivateImpl=function(){var t=this,i=this.activationState;if(i.isActivated){var r=O({},i);i.isProgrammatic?(requestAnimationFrame(function(){t.animateDeactivation(r)}),this.resetActivationState()):(this.deregisterDeactivationHandlers(),requestAnimationFrame(function(){t.activationState.hasDeactivationUXRun=!0,t.animateDeactivation(r),t.resetActivationState()}))}},e.prototype.animateDeactivation=function(t){var i=t.wasActivatedByPointer,r=t.wasElementMadeActive;(i||r)&&this.runDeactivationUXLogicIfReady()},e.prototype.layoutInternal=function(){var t=this;this.frame=this.adapter.computeBoundingRect();var i=Math.max(this.frame.height,this.frame.width),r=function(){var l=Math.sqrt(Math.pow(t.frame.width,2)+Math.pow(t.frame.height,2));return l+e.numbers.PADDING};this.maxRadius=this.adapter.isUnbounded()?i:r();var o=Math.floor(i*e.numbers.INITIAL_ORIGIN_SCALE);this.adapter.isUnbounded()&&o%2!==0?this.initialSize=o-1:this.initialSize=o,this.fgScale=""+this.maxRadius/this.initialSize,this.updateLayoutCssVars()},e.prototype.updateLayoutCssVars=function(){var t=e.strings,i=t.VAR_FG_SIZE,r=t.VAR_LEFT,o=t.VAR_TOP,l=t.VAR_FG_SCALE;this.adapter.updateCssVariable(i,this.initialSize+"px"),this.adapter.updateCssVariable(l,this.fgScale),this.adapter.isUnbounded()&&(this.unboundedCoords={left:Math.round(this.frame.width/2-this.initialSize/2),top:Math.round(this.frame.height/2-this.initialSize/2)},this.adapter.updateCssVariable(r,this.unboundedCoords.left+"px"),this.adapter.updateCssVariable(o,this.unboundedCoords.top+"px"))},e}(ze);const{applyPassive:E}=Ee,{matches:Ke}=We;function x(n,{ripple:e=!0,surface:t=!1,unbounded:i=!1,disabled:r=!1,color:o,active:l,rippleElement:a,eventTarget:s,activeTarget:u,addClass:f=d=>n.classList.add(d),removeClass:m=d=>n.classList.remove(d),addStyle:g=(d,y)=>n.style.setProperty(d,y),initPromise:_=Promise.resolve()}={}){let d,y=q("SMUI:addLayoutListener"),A,h=l,w=s,D=u;function R(){t?(f("mdc-ripple-surface"),o==="primary"?(f("smui-ripple-surface--primary"),m("smui-ripple-surface--secondary")):o==="secondary"?(m("smui-ripple-surface--primary"),f("smui-ripple-surface--secondary")):(m("smui-ripple-surface--primary"),m("smui-ripple-surface--secondary"))):(m("mdc-ripple-surface"),m("smui-ripple-surface--primary"),m("smui-ripple-surface--secondary")),d&&h!==l&&(h=l,l?d.activate():l===!1&&d.deactivate()),e&&!d?(d=new Je({addClass:f,browserSupportsCssVars:()=>Ge(window),computeBoundingRect:()=>(a||n).getBoundingClientRect(),containsEventTarget:b=>n.contains(b),deregisterDocumentInteractionHandler:(b,p)=>document.documentElement.removeEventListener(b,p,E()),deregisterInteractionHandler:(b,p)=>(s||n).removeEventListener(b,p,E()),deregisterResizeHandler:b=>window.removeEventListener("resize",b),getWindowPageOffset:()=>({x:window.pageXOffset,y:window.pageYOffset}),isSurfaceActive:()=>l??Ke(u||n,":active"),isSurfaceDisabled:()=>!!r,isUnbounded:()=>!!i,registerDocumentInteractionHandler:(b,p)=>document.documentElement.addEventListener(b,p,E()),registerInteractionHandler:(b,p)=>(s||n).addEventListener(b,p,E()),registerResizeHandler:b=>window.addEventListener("resize",b),removeClass:m,updateCssVariable:g}),_.then(()=>{d&&(d.init(),d.setUnbounded(i))})):d&&!e&&_.then(()=>{d&&(d.destroy(),d=void 0)}),d&&(w!==s||D!==u)&&(w=s,D=u,d.destroy(),requestAnimationFrame(()=>{d&&(d.init(),d.setUnbounded(i))})),!e&&i&&f("mdc-ripple-upgraded--unbounded")}R(),y&&(A=y(B));function B(){d&&d.layout()}return{update(b){({ripple:e,surface:t,unbounded:i,disabled:r,color:o,active:l,rippleElement:a,eventTarget:s,activeTarget:u,addClass:f,removeClass:m,addStyle:g,initPromise:_}=Object.assign({ripple:!0,surface:!1,unbounded:!1,disabled:!1,color:void 0,active:void 0,rippleElement:void 0,eventTarget:void 0,activeTarget:void 0,addClass:p=>n.classList.add(p),removeClass:p=>n.classList.remove(p),addStyle:(p,S)=>n.style.setProperty(p,S),initPromise:Promise.resolve()},b)),R()},destroy(){d&&(d.destroy(),d=void 0,m("mdc-ripple-surface"),m("smui-ripple-surface--primary"),m("smui-ripple-surface--secondary")),A&&A()}}}function Qe(n){let e;const t=n[10].default,i=de(t,n,n[12],null);return{c(){i&&i.c()},l(r){i&&i.l(r)},m(r,o){i&&i.m(r,o),e=!0},p(r,o){i&&i.p&&(!e||o&4096)&&fe(i,t,r,r[12],e?pe(t,r[12],o,null):me(r[12]),null)},i(r){e||(k(i,r),e=!0)},o(r){L(i,r),e=!1},d(r){i&&i.d(r)}}}function Ne(n){let e,t,i;const r=[{tag:n[3]},{use:[n[5],...n[0]]},{class:H({[n[1]]:!0,"mdc-button__label":n[6]==="button","mdc-fab__label":n[6]==="fab","mdc-tab__text-label":n[6]==="tab","mdc-image-list__label":n[6]==="image-list","mdc-snackbar__label":n[6]==="snackbar","mdc-banner__text":n[6]==="banner","mdc-segmented-button__label":n[6]==="segmented-button","mdc-data-table__pagination-rows-per-page-label":n[6]==="data-table:pagination","mdc-data-table__header-cell-label":n[6]==="data-table:sortable-header-cell"})},n[6]==="snackbar"?{"aria-atomic":"false"}:{},{tabindex:n[7]},n[8]];var o=n[2];function l(a,s){let u={$$slots:{default:[Qe]},$$scope:{ctx:a}};if(s!==void 0&&s&491)u=K(r,[s&8&&{tag:a[3]},s&33&&{use:[a[5],...a[0]]},s&66&&{class:H({[a[1]]:!0,"mdc-button__label":a[6]==="button","mdc-fab__label":a[6]==="fab","mdc-tab__text-label":a[6]==="tab","mdc-image-list__label":a[6]==="image-list","mdc-snackbar__label":a[6]==="snackbar","mdc-banner__text":a[6]==="banner","mdc-segmented-button__label":a[6]==="segmented-button","mdc-data-table__pagination-rows-per-page-label":a[6]==="data-table:pagination","mdc-data-table__header-cell-label":a[6]==="data-table:sortable-header-cell"})},s&64&&v(a[6]==="snackbar"?{"aria-atomic":"false"}:{}),s&128&&{tabindex:a[7]},s&256&&v(a[8])]);else for(let f=0;f<r.length;f+=1)u=M(u,r[f]);return{props:u}}return o&&(e=W(o,l(n)),n[11](e)),{c(){e&&V(e.$$.fragment),t=U()},l(a){e&&ve(e.$$.fragment,a),t=U()},m(a,s){e&&Z(e,a,s),F(a,t,s),i=!0},p(a,[s]){if(s&4&&o!==(o=a[2])){if(e){ge();const u=e;L(u.$$.fragment,1,0,()=>{J(u,1)}),_e()}o?(e=W(o,l(a,s)),a[11](e),V(e.$$.fragment),k(e.$$.fragment,1),Z(e,t.parentNode,t)):e=null}else if(o){const u=s&491?K(r,[s&8&&{tag:a[3]},s&33&&{use:[a[5],...a[0]]},s&66&&{class:H({[a[1]]:!0,"mdc-button__label":a[6]==="button","mdc-fab__label":a[6]==="fab","mdc-tab__text-label":a[6]==="tab","mdc-image-list__label":a[6]==="image-list","mdc-snackbar__label":a[6]==="snackbar","mdc-banner__text":a[6]==="banner","mdc-segmented-button__label":a[6]==="segmented-button","mdc-data-table__pagination-rows-per-page-label":a[6]==="data-table:pagination","mdc-data-table__header-cell-label":a[6]==="data-table:sortable-header-cell"})},s&64&&v(a[6]==="snackbar"?{"aria-atomic":"false"}:{}),s&128&&{tabindex:a[7]},s&256&&v(a[8])]):{};s&4096&&(u.$$scope={dirty:s,ctx:a}),e.$set(u)}},i(a){i||(e&&k(e.$$.fragment,a),i=!0)},o(a){e&&L(e.$$.fragment,a),i=!1},d(a){a&&C(t),n[11](null),e&&J(e,a)}}}function Te(n,e,t){const i=["use","class","component","tag","getElement"];let r=X(e,i),{$$slots:o={},$$scope:l}=e;const a=De(ue());let{use:s=[]}=e,{class:u=""}=e,f,{component:m=Q}=e,{tag:g=m===Q?"span":void 0}=e;const _=q("SMUI:label:context"),d=q("SMUI:label:tabindex");function y(){return f.getElement()}function A(h){ce[h?"unshift":"push"](()=>{f=h,t(4,f)})}return n.$$set=h=>{e=M(M({},e),ee(h)),t(8,r=X(e,i)),"use"in h&&t(0,s=h.use),"class"in h&&t(1,u=h.class),"component"in h&&t(2,m=h.component),"tag"in h&&t(3,g=h.tag),"$$scope"in h&&t(12,l=h.$$scope)},[s,u,m,g,f,a,_,d,r,y,o,A,l]}class nt extends he{constructor(e){super(),be(this,e,Te,Ne,le,{use:0,class:1,component:2,tag:3,getElement:9})}get getElement(){return this.$$.ctx[9]}}function oe(n){let e;return{c(){e=ye("div"),this.h()},l(t){e=Ae(t,"DIV",{class:!0}),Se(e).forEach(C),this.h()},h(){Ce(e,"class","mdc-button__touch")},m(t,i){F(t,e,i)},d(t){t&&C(e)}}}function Ye(n){let e,t,i,r;const o=n[28].default,l=de(o,n,n[30],null);let a=n[6]&&oe();return{c(){e=ye("div"),t=Le(),l&&l.c(),a&&a.c(),i=U(),this.h()},l(s){e=Ae(s,"DIV",{class:!0}),Se(e).forEach(C),t=ke(s),l&&l.l(s),a&&a.l(s),i=U(),this.h()},h(){Ce(e,"class","mdc-button__ripple")},m(s,u){F(s,e,u),F(s,t,u),l&&l.m(s,u),a&&a.m(s,u),F(s,i,u),r=!0},p(s,u){l&&l.p&&(!r||u[0]&1073741824)&&fe(l,o,s,s[30],r?pe(o,s[30],u,null):me(s[30]),null),s[6]?a||(a=oe(),a.c(),a.m(i.parentNode,i)):a&&(a.d(1),a=null)},i(s){r||(k(l,s),r=!0)},o(s){L(l,s),r=!1},d(s){s&&(C(e),C(t),C(i)),l&&l.d(s),a&&a.d(s)}}}function xe(n){let e,t,i;const r=[{tag:n[10]},{use:[[x,{ripple:n[3],unbounded:!1,color:n[4],disabled:!!n[23].disabled,addClass:n[19],removeClass:n[20],addStyle:n[21]}],n[17],...n[0]]},{class:H({[n[1]]:!0,"mdc-button":!0,"mdc-button--raised":n[5]==="raised","mdc-button--unelevated":n[5]==="unelevated","mdc-button--outlined":n[5]==="outlined","smui-button--color-secondary":n[4]==="secondary","mdc-button--touch":n[6],"mdc-card__action":n[18]==="card:action","mdc-card__action--button":n[18]==="card:action","mdc-dialog__button":n[18]==="dialog:action","mdc-top-app-bar__navigation-icon":n[18]==="top-app-bar:navigation","mdc-top-app-bar__action-item":n[18]==="top-app-bar:action","mdc-snackbar__action":n[18]==="snackbar:actions","mdc-banner__secondary-action":n[18]==="banner"&&n[8],"mdc-banner__primary-action":n[18]==="banner"&&!n[8],"mdc-tooltip__action":n[18]==="tooltip:rich-actions",...n[12]})},{style:Object.entries(n[13]).map($).concat([n[2]]).join(" ")},n[16],n[15],n[14],{href:n[7]},n[23]];var o=n[9];function l(a,s){let u={$$slots:{default:[Ye]},$$scope:{ctx:a}};if(s!==void 0&&s[0]&12580351)u=K(r,[s[0]&1024&&{tag:a[10]},s[0]&12189721&&{use:[[x,{ripple:a[3],unbounded:!1,color:a[4],disabled:!!a[23].disabled,addClass:a[19],removeClass:a[20],addStyle:a[21]}],a[17],...a[0]]},s[0]&266610&&{class:H({[a[1]]:!0,"mdc-button":!0,"mdc-button--raised":a[5]==="raised","mdc-button--unelevated":a[5]==="unelevated","mdc-button--outlined":a[5]==="outlined","smui-button--color-secondary":a[4]==="secondary","mdc-button--touch":a[6],"mdc-card__action":a[18]==="card:action","mdc-card__action--button":a[18]==="card:action","mdc-dialog__button":a[18]==="dialog:action","mdc-top-app-bar__navigation-icon":a[18]==="top-app-bar:navigation","mdc-top-app-bar__action-item":a[18]==="top-app-bar:action","mdc-snackbar__action":a[18]==="snackbar:actions","mdc-banner__secondary-action":a[18]==="banner"&&a[8],"mdc-banner__primary-action":a[18]==="banner"&&!a[8],"mdc-tooltip__action":a[18]==="tooltip:rich-actions",...a[12]})},s[0]&8196&&{style:Object.entries(a[13]).map($).concat([a[2]]).join(" ")},s[0]&65536&&v(a[16]),s[0]&32768&&v(a[15]),s[0]&16384&&v(a[14]),s[0]&128&&{href:a[7]},s[0]&8388608&&v(a[23])]);else for(let f=0;f<r.length;f+=1)u=M(u,r[f]);return{props:u}}return o&&(e=W(o,l(n)),n[29](e),e.$on("click",n[22])),{c(){e&&V(e.$$.fragment),t=U()},l(a){e&&ve(e.$$.fragment,a),t=U()},m(a,s){e&&Z(e,a,s),F(a,t,s),i=!0},p(a,s){if(s[0]&512&&o!==(o=a[9])){if(e){ge();const u=e;L(u.$$.fragment,1,0,()=>{J(u,1)}),_e()}o?(e=W(o,l(a,s)),a[29](e),e.$on("click",a[22]),V(e.$$.fragment),k(e.$$.fragment,1),Z(e,t.parentNode,t)):e=null}else if(o){const u=s[0]&12580351?K(r,[s[0]&1024&&{tag:a[10]},s[0]&12189721&&{use:[[x,{ripple:a[3],unbounded:!1,color:a[4],disabled:!!a[23].disabled,addClass:a[19],removeClass:a[20],addStyle:a[21]}],a[17],...a[0]]},s[0]&266610&&{class:H({[a[1]]:!0,"mdc-button":!0,"mdc-button--raised":a[5]==="raised","mdc-button--unelevated":a[5]==="unelevated","mdc-button--outlined":a[5]==="outlined","smui-button--color-secondary":a[4]==="secondary","mdc-button--touch":a[6],"mdc-card__action":a[18]==="card:action","mdc-card__action--button":a[18]==="card:action","mdc-dialog__button":a[18]==="dialog:action","mdc-top-app-bar__navigation-icon":a[18]==="top-app-bar:navigation","mdc-top-app-bar__action-item":a[18]==="top-app-bar:action","mdc-snackbar__action":a[18]==="snackbar:actions","mdc-banner__secondary-action":a[18]==="banner"&&a[8],"mdc-banner__primary-action":a[18]==="banner"&&!a[8],"mdc-tooltip__action":a[18]==="tooltip:rich-actions",...a[12]})},s[0]&8196&&{style:Object.entries(a[13]).map($).concat([a[2]]).join(" ")},s[0]&65536&&v(a[16]),s[0]&32768&&v(a[15]),s[0]&16384&&v(a[14]),s[0]&128&&{href:a[7]},s[0]&8388608&&v(a[23])]):{};s[0]&1073741888&&(u.$$scope={dirty:s,ctx:a}),e.$set(u)}},i(a){i||(e&&k(e.$$.fragment,a),i=!0)},o(a){e&&L(e.$$.fragment,a),i=!1},d(a){a&&C(t),n[29](null),e&&J(e,a)}}}const $=([n,e])=>`${n}: ${e};`;function $e(n,e,t){let i,r,o;const l=["use","class","style","ripple","color","variant","touch","href","action","defaultAction","secondary","component","tag","getElement"];let a=X(e,l),{$$slots:s={},$$scope:u}=e;const f=De(ue());let{use:m=[]}=e,{class:g=""}=e,{style:_=""}=e,{ripple:d=!0}=e,{color:y="primary"}=e,{variant:A="text"}=e,{touch:h=!1}=e,{href:w=void 0}=e,{action:D="close"}=e,{defaultAction:R=!1}=e,{secondary:B=!1}=e,b,p={},S={},G=q("SMUI:button:context"),{component:N=Q}=e,{tag:ae=N===Q?w==null?"button":"a":void 0}=e,T=a.disabled;ne("SMUI:label:context","button"),ne("SMUI:icon:context","button");function Re(c){p[c]||t(12,p[c]=!0,p)}function Fe(c){(!(c in p)||p[c])&&t(12,p[c]=!1,p)}function He(c,I){S[c]!=I&&(I===""||I==null?(delete S[c],t(13,S)):t(13,S[c]=I,S))}function Me(){G==="banner"&&Be(Y(),B?"SMUIBannerButton:secondaryActionClick":"SMUIBannerButton:primaryActionClick")}function Y(){return b.getElement()}function Ue(c){ce[c?"unshift":"push"](()=>{b=c,t(11,b)})}return n.$$set=c=>{t(31,e=M(M({},e),ee(c))),t(23,a=X(e,l)),"use"in c&&t(0,m=c.use),"class"in c&&t(1,g=c.class),"style"in c&&t(2,_=c.style),"ripple"in c&&t(3,d=c.ripple),"color"in c&&t(4,y=c.color),"variant"in c&&t(5,A=c.variant),"touch"in c&&t(6,h=c.touch),"href"in c&&t(7,w=c.href),"action"in c&&t(24,D=c.action),"defaultAction"in c&&t(25,R=c.defaultAction),"secondary"in c&&t(8,B=c.secondary),"component"in c&&t(9,N=c.component),"tag"in c&&t(10,ae=c.tag),"$$scope"in c&&t(30,u=c.$$scope)},n.$$.update=()=>{if(t(16,i=G==="dialog:action"&&D!=null?{"data-mdc-dialog-action":D}:{action:e.action}),t(15,r=G==="dialog:action"&&R?{"data-mdc-dialog-button-default":""}:{default:e.default}),t(14,o=G==="banner"?{}:{secondary:e.secondary}),T!==a.disabled){const c=Y();"blur"in c&&c.blur(),t(27,T=a.disabled)}},e=ee(e),[m,g,_,d,y,A,h,w,B,N,ae,b,p,S,o,r,i,f,G,Re,Fe,He,Me,a,D,R,Y,T,s,Ue,u]}class it extends he{constructor(e){super(),be(this,e,$e,xe,le,{use:0,class:1,style:2,ripple:3,color:4,variant:5,touch:6,href:7,action:24,defaultAction:25,secondary:8,component:9,tag:10,getElement:26},null,[-1,-1])}get getElement(){return this.$$.ctx[26]}}export{it as B,nt as C,ze as M,x as R,Ie as _,O as a,j as b,Be as d,Ee as e};