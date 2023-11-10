import { config } from './config';

let;

const constraints = {
	video: true,
	audio: true,
};

export const pageReady = (localuuid, localStream, localVideo) => {
	const path = window.location.pathname.split('/')[1];
	const address = config.apiWsHost + '/api/game/signaling/ws/' + path;
	console.log('tries establish websocket connection on ' + address);

	serverConnection = new WebSocket(address);
	serverConnection.onmessage = gotMessageFromServer;
	serverConnection.onopen = async () => {
		if (navigator.mediaDevices.getUserMedia) {
			await navigator.mediaDevices
				.getUserMedia(constraints)
				.then(function getUserMediaSuccess(stream) {
					localStream = stream;
					localVideo.srcObject = stream;
				})
				.catch((e) => {
					errorHandler(e, null);
				});
		} else {
			alert('Your browser does not support getUserMedia API');
			return;
		}
		serverConnection.send(JSON.stringify({ type: 'join', data: null, srcuuid: localuuid, dstuuid: '' }));
	};
};

async function start(isCaller, uuid_) {
	if (isCaller) console.log('establish connection with existing users');
	else console.log('establish connection with new user');

	const peerConnection = new RTCPeerConnection(peerConnectionConfig);
	peerConnection.onicecandidate = (e) => {
		gotIceCandidate(e, uuid_);
	};
	peerConnection.ontrack = (e) => {
		gotRemoteStream(e, uuid_);
	};
	peerConnection.addStream(localStream);

	if (isCaller) {
		peerConnection
			.createOffer()
			.then((description) => {
				return createdDescription(description, uuid_);
			})
			.catch((e) => errorHandler(e, uuid_));
	}

	return peerConnection;
}

async function gotMessageFromServer(message) {
	const data = JSON.parse(message.data);

	console.log(data);

	if (data.type == 'users') {
		data.data.users.forEach(async (user) => {
			peers[user] = { rtc: await start(true, user) };
		});
	} else if (data.type == 'join') {
		``;
		peers[data.srcuuid] = { rtc: await start(false, data.srcuuid) };
	} else if (data.type == 'sdp') {
		peers[data.srcuuid].rtc
			.setRemoteDescription(new RTCSessionDescription(data.data.sdp))
			.then(function () {
				console.log('set remote description');
				// Only create answers in response to offers
				if (data.data.sdp.type == 'offer') {
					peers[data.srcuuid].rtc
						.createAnswer()
						.then((description) => {
							createdDescription(description, data.srcuuid);
						})
						.catch((e) => {
							errorHandler(e, data.srcuuid);
						});
				}
			})
			.catch((e) => {
				errorHandler(e, data.srcuuid);
			});
	} else if (data.type == 'ice') {
		peers[data.srcuuid].rtc.addIceCandidate(new RTCIceCandidate(data.data.ice)).catch((e) => {
			errorHandler(e, data.srcuuid);
		});
	} else if (data.type == 'leave') {
		peers[data.srcuuid].video.remove();
		delete peers[data.srcuuid];
	}
}

function gotIceCandidate(event, uuid_) {
	if (event.candidate != null) {
		serverConnection.send(
			JSON.stringify({ type: 'ice', data: { ice: event.candidate }, srcuuid: localuuid, dstuuid: uuid_ }),
		);
	}
}

function createdDescription(description, uuid_) {
	peers[uuid_].rtc
		.setLocalDescription(description)
		.then(function () {
			console.log('set local description');
			serverConnection.send(
				JSON.stringify({
					type: 'sdp',
					data: { sdp: peers[uuid_].rtc.localDescription },
					srcuuid: localuuid,
					dstuuid: uuid_,
				}),
			);
		})
		.catch(errorHandler);
}

function gotRemoteStream(event, uuid_) {
	console.log('got remote stream', event);
	if (!peers[uuid_].video) {
		peers[uuid_].mediastream = new MediaStream();
		peers[uuid_].video = document.createElement('video');
		peers[uuid_].video.style = 'width: 40%;';
		peers[uuid_].video.autoplay = true;
		peers[uuid_].video.srcObject = peers[uuid_].mediastream;
		videos.appendChild(peers[uuid_].video);
	}

	peers[uuid_].mediastream.addTrack(event.track);
}

function errorHandler(error, uuid_) {
	console.log(uuid_, error);
}
