var p=Object.defineProperty;var g=(a,t,e)=>t in a?p(a,t,{enumerable:!0,configurable:!0,writable:!0,value:e}):a[t]=e;var n=(a,t,e)=>(g(a,typeof t!="symbol"?t+"":t,e),e);import{c as C}from"./config.27eae293.js";const m={iceServers:[{urls:["stun:stun.l.google.com:19302"]}]},f={preferCurrentTab:!0,video:!0,audio:!0},d={offerToReceiveAudio:!0,offerToReceiveVideo:!0};class y{constructor(t,e,o){n(this,"serverConnection",null);n(this,"localStream",null);n(this,"localVideo");n(this,"peerConnection",null);n(this,"userType","");n(this,"mediaStream",new MediaStream);n(this,"start",async()=>{var t;console.log("establish connection with peer"),this.userType=="tetris"&&((t=this.peerConnection)==null||t.createOffer(d).then(e=>this.createdDescription(e)).catch(e=>this.errorHandler(e)))});n(this,"gotIceCandidate",t=>{var e;console.log("got ice candidate"),t.candidate!=null&&((e=this.serverConnection)==null||e.send(JSON.stringify({type:"ice",data:{ice:t.candidate}})))});n(this,"gotRemoteStream",async t=>{console.log("got remote stream"),this.mediaStream.addTrack(t.track),this.localVideo.srcObject=this.mediaStream});n(this,"gotMessageFromServer",async t=>{var o,r,s;const e=JSON.parse(t.data);if(console.log(e),e.type=="join")this.start();else if(e.type=="sdp")if(await((o=this.peerConnection)==null?void 0:o.setRemoteDescription(new RTCSessionDescription(e.data.sdp)).catch(i=>{this.errorHandler(i)})),e.data.sdp.type=="offer"){console.log("offer",this.peerConnection);const i=await((r=this.peerConnection)==null?void 0:r.createAnswer(d).catch(c=>{this.errorHandler(c)}));console.log(i),this.createdDescription(i)}else e.data.sdp.type=="answer"&&console.log("answer");else e.type=="ice"&&((s=this.peerConnection)==null||s.addIceCandidate(new RTCIceCandidate(e.data.ice)).catch(i=>{this.errorHandler(i)}))});n(this,"createdDescription",t=>{var e;(e=this.peerConnection)==null||e.setLocalDescription(t).then(()=>{var o,r;console.log("set local description"),(r=this.serverConnection)==null||r.send(JSON.stringify({type:"sdp",data:{sdp:(o=this.peerConnection)==null?void 0:o.localDescription}}))}).catch(this.errorHandler)});n(this,"errorHandler",t=>{console.log(t)});n(this,"currentState",()=>{var t,e,o;return console.log((t=this.peerConnection)==null?void 0:t.signalingState,(e=this.peerConnection)==null?void 0:e.connectionState),(o=this.peerConnection)==null?void 0:o.signalingState});this.localVideo=t,this.localVideo.srcObject=this.mediaStream,this.userType=o;const r=C.apiWsHost+"/api/game/signaling/"+e;console.log("tries establish websocket connection on "+r),this.peerConnection=new RTCPeerConnection(m),this.peerConnection.onicecandidate=s=>{this.gotIceCandidate(s)},this.peerConnection.ontrack=s=>{this.gotRemoteStream(s)},this.serverConnection=new WebSocket(r),this.serverConnection.onmessage=this.gotMessageFromServer,this.serverConnection.onopen=async()=>{var s;await navigator.mediaDevices.getDisplayMedia(f).then(i=>{var c;this.localStream=i,(c=this.localStream)==null||c.getTracks().forEach(h=>{var l;(l=this.peerConnection)==null||l.addTrack(h,this.localStream)})}).catch(i=>{this.errorHandler(i)}),(s=this.serverConnection)==null||s.send(JSON.stringify({userType:this.userType}))}}}export{y as W};