import { config } from './config';

const peerConnectionConfig = {
	iceServers: [{ urls: 'stun:stun.stunprotocol.org:3478' }, { urls: 'stun:stun.l.google.com:19302' }],
};

const constraints: MediaStreamConstraints = {
	preferCurrentTab: true,
	video: true,
	audio: true,
};

export class WebRTC {
	serverConnection: WebSocket | null = null;
	localStream: MediaStream | null = null;
	localVideo: HTMLVideoElement;
	peerConnection: RTCPeerConnection | null = null;
	userType: 'tetris' | 'wordguess' | '' = '';

	constructor(localVideo: HTMLVideoElement, session: string, userType: 'tetris' | 'wordguess' | '') {
		this.localVideo = localVideo;
		this.userType = userType;
		const address = config.apiWsHost + '/api/game/signaling/' + session;
		console.log('tries establish websocket connection on ' + address);
		this.serverConnection = new WebSocket(address);
		this.serverConnection.onmessage = this.gotMessageFromServer;
		this.serverConnection.onopen = async () => {
			await navigator.mediaDevices
				.getDisplayMedia(constraints)
				.then((stream) => {
					this.localStream = stream;
				})
				.catch((e) => {
					this.errorHandler(e);
				});
			this.serverConnection?.send(JSON.stringify({ userType: this.userType }));
		};
		this.peerConnection = new RTCPeerConnection(peerConnectionConfig);
		this.peerConnection.onicecandidate = (e) => {
			this.gotIceCandidate(e);
		};
		this.peerConnection.ontrack = (e) => {
			this.gotRemoteStream(e);
		};
		if (this.localStream?.getTracks()[0])
			this.peerConnection.addTrack(this.localStream?.getTracks()[0], this.localStream!);
	}

	async start() {
		console.log('establish connection with peer');
		if (this.userType == 'tetris') {
			this.peerConnection
				?.createOffer()
				.then((description) => {
					return this.createdDescription(description);
				})
				.catch((e) => this.errorHandler(e));
		}
	}

	gotIceCandidate(event: RTCPeerConnectionIceEvent) {
		if (event.candidate != null) {
			this.serverConnection?.send(JSON.stringify({ type: 'ice', data: { ice: event.candidate } }));
		}
	}

	async gotRemoteStream(event: RTCTrackEvent) {
		this.localVideo.srcObject = event.streams[0];
	}

	async gotMessageFromServer(message: MessageEvent) {
		const data = JSON.parse(message.data);

		console.log(data);

		if (data.type == 'join') {
			this.start();
		} else if (data.type == 'sdp') {
			await this.peerConnection?.setRemoteDescription(new RTCSessionDescription(data.data.sdp)).catch((e) => {
				this.errorHandler(e);
			});

			if (data.data.sdp.type == 'offer') {
				this.peerConnection
					?.createAnswer()
					.then((description) => {
						this.createdDescription(description);
					})
					.catch((e) => {
						this.errorHandler(e);
					});
			} else if (data.data.sdp.type == 'answer') {
				console.log('answer');
			}
		} else if (data.type == 'ice') {
			this.peerConnection?.addIceCandidate(new RTCIceCandidate(data.data.ice)).catch((e) => {
				this.errorHandler(e);
			});
		}
	}

	createdDescription(description: RTCSessionDescriptionInit) {
		this.peerConnection
			?.setLocalDescription(description)
			.then(() => {
				console.log('set local description');
				this.serverConnection?.send(
					JSON.stringify({
						type: 'sdp',
						data: { sdp: this.peerConnection?.localDescription },
					}),
				);
			})
			.catch(this.errorHandler);
	}

	errorHandler(error: Error) {
		console.log(error);
	}
}
