import { ref } from "vue";

export const useTelepresence = () => {
  const ws = ref<WebSocket | null>(null);
  const pc = ref<RTCPeerConnection | null>(null);
  const remoteVideo = ref<HTMLVideoElement | null>(null);

  const robotId = ref<string | null>(null);

  // =========================
  // CONNECT WEBSOCKET
  // =========================
  const connectWS = () => {
    if (!robotId.value) {
      console.error("[Telepresence] robotId is required");
      return;
    }

    ws.value = new WebSocket("ws://localhost:8082/ws");

    ws.value.onopen = () => {
      console.log("[Telepresence] WS connected");

      // STEP 1: bind session to robot
      ws.value?.send(
        JSON.stringify({
          type: "start_session",
          robotId: robotId.value,
        }),
      );

      // STEP 2: start WebRTC negotiation
      startWebRTC();
    };

    ws.value.onmessage = async (event) => {
      const msg = JSON.parse(event.data);

      switch (msg.type) {
        case "answer":
          await handleAnswer(msg.sdp);
          break;

        case "candidate":
          await handleCandidate(msg.candidate);
          break;
      }
    };

    ws.value.onclose = () => {
      console.log("[Telepresence] WS disconnected");
    };
  };

  // =========================
  // WEBRTC SETUP
  // =========================
  const startWebRTC = async () => {
    pc.value = new RTCPeerConnection({
      iceServers: [{ urls: "stun:stun.l.google.com:19302" }],
    });

    // receive robot video
    pc.value.ontrack = (event) => {
      const video = remoteVideo.value;
      const stream = event.streams?.[0];

      if (video && stream) {
        video.srcObject = stream;
      }
    };
    // send ICE candidates to backend
    pc.value.onicecandidate = (event) => {
      if (event.candidate) {
        ws.value?.send(
          JSON.stringify({
            type: "candidate",
            candidate: event.candidate,
          }),
        );
      }
    };

    // create offer
    const offer = await pc.value.createOffer();
    await pc.value.setLocalDescription(offer);

    ws.value?.send(
      JSON.stringify({
        type: "offer",
        sdp: offer.sdp,
      }),
    );
  };

  // =========================
  // HANDLE ANSWER
  // =========================
  const handleAnswer = async (sdp: string) => {
    if (!pc.value) return;

    await pc.value.setRemoteDescription(
      new RTCSessionDescription({
        type: "answer",
        sdp,
      }),
    );
  };

  // =========================
  // HANDLE ICE CANDIDATE
  // =========================
  const handleCandidate = async (candidate: any) => {
    try {
      await pc.value?.addIceCandidate(candidate);
    } catch (err) {
      console.error("[Telepresence] ICE error:", err);
    }
  };

  return {
    robotId,
    connectWS,
    remoteVideo,
  };
};
