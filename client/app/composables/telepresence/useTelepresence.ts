import { ref } from "vue";

export const useTelepresence = () => {
  const ws = ref<WebSocket | null>(null);
  const pc = ref<RTCPeerConnection | null>(null);
  const videoRef = ref<HTMLVideoElement | null>(null);

  const robotId = ref<string | null>(null);

  // =========================
  // STATE CONTROL (IMPORTANT)
  // =========================
  const sessionReady = ref(false);
  const isConnecting = ref(false);
  const isWSOpen = ref(false);

  // =========================
  // SET ROBOT ID
  // =========================
  const setRobot = (id: string) => {
    robotId.value = id;
  };

  // =========================
  // CONNECT WEBSOCKET
  // =========================
  const connectWS = (): Promise<void> => {
    return new Promise((resolve, reject) => {
      if (!robotId.value) {
        reject("robotId not set");
        return;
      }

      ws.value = new WebSocket("ws://localhost:8082/ws");

      ws.value.onopen = () => {
        console.log("[Telepresence] WS connected");

        ws.value?.send(
          JSON.stringify({
            type: "register_robot",
            robotId: robotId.value,
          }),
        );

        isWSOpen.value = true;
        sessionReady.value = true;

        console.log("[Telepresence] Robot registered");
        resolve();
      };

      ws.value.onerror = (err) => {
        console.error("[Telepresence] WS error", err);
        reject(err);
      };

      ws.value.onclose = () => {
        console.log("[Telepresence] WS closed");
        isWSOpen.value = false;
        sessionReady.value = false;
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
    });
  };

  // =========================
  // START WEBRTC
  // =========================
  const startWebRTC = async () => {
    // 🔥 HARD GUARDS
    if (!sessionReady.value || !isWSOpen.value) {
      console.error("[Telepresence] WS not ready");
      return;
    }

    if (isConnecting.value) {
      console.warn("[Telepresence] Already connecting...");
      return;
    }

    if (!robotId.value) {
      console.error("[Telepresence] Missing robotId");
      return;
    }

    isConnecting.value = true;

    try {
      // =========================
      // CLEAN OLD PEER CONNECTION
      // =========================
      if (pc.value) {
        pc.value.close();
        pc.value = null;
      }

      // =========================
      // CREATE NEW PEER CONNECTION
      // =========================
      pc.value = new RTCPeerConnection({
        iceServers: [{ urls: "stun:stun.l.google.com:19302" }],
      });
      pc.value.addTransceiver("video", { direction: "recvonly" });
      // =========================
      // RECEIVE VIDEO
      // =========================
      pc.value.ontrack = (event) => {
        const stream = event.streams?.[0];

        if (videoRef.value && stream) {
          videoRef.value.srcObject = stream;
        }
      };

      // =========================
      // SEND ICE CANDIDATES
      // =========================
      pc.value.onicecandidate = (event) => {
        if (event.candidate && ws.value) {
          ws.value.send(
            JSON.stringify({
              type: "candidate",
              robotId: robotId.value,
              candidate: event.candidate,
            }),
          );
        }
      };

      // =========================
      // CREATE OFFER
      // =========================
      const offer = await pc.value.createOffer();
      await pc.value.setLocalDescription(offer);

      if (!ws.value || ws.value.readyState !== WebSocket.OPEN) {
        console.error("[Telepresence] WebSocket not ready");
        return;
      }

      ws.value.send(
        JSON.stringify({
          type: "offer",
          robotId: robotId.value,
          sdp: offer.sdp,
        }),
      );

      console.log("[Telepresence] Offer sent");
    } catch (err) {
      console.error("[Telepresence] WebRTC error:", err);
    } finally {
      isConnecting.value = false;
    }
  };

  // =========================
  // HANDLE ANSWER
  // =========================
  const handleAnswer = async (sdp: string) => {
    if (!pc.value) return;

    try {
      await pc.value.setRemoteDescription(
        new RTCSessionDescription({
          type: "answer",
          sdp,
        }),
      );

      console.log("[Telepresence] Answer applied");
    } catch (err) {
      console.error("[Telepresence] Answer error:", err);
    }
  };

  // =========================
  // HANDLE ICE
  // =========================
  const handleCandidate = async (candidate: any) => {
    try {
      await pc.value?.addIceCandidate(candidate);
    } catch (err) {
      console.error("[Telepresence] ICE error:", err);
    }
  };

  // =========================
  // CLEANUP
  // =========================
  const disconnect = () => {
    ws.value?.close();
    pc.value?.close();

    ws.value = null;
    pc.value = null;

    sessionReady.value = false;
    isConnecting.value = false;
    isWSOpen.value = false;

    console.log("[Telepresence] Disconnected");
  };

  return {
    robotId,
    setRobot,
    connectWS,
    startWebRTC,
    videoRef,
    disconnect,
  };
};
