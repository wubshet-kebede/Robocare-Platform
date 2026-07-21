// import { useWebSocket } from "@vueuse/core";
import { ref, watch } from "vue";

let socketInstance: any = null;

export function useVitalsSocket(hospitalId: string) {
  if (socketInstance) return socketInstance;

  const vitalsMap = ref<Record<string, any>>({});

  const { data, status, send, open } = useWebSocket("ws://localhost:8082/ws", {
    autoReconnect: {
      retries: 10,
      delay: 2000,
    },
  });

  watch(status, (s) => {
    if (s === "OPEN") {
      // 🔥 IMPORTANT: register hospital
      send(
        JSON.stringify({
          type: "register_session",
          hospitalId: hospitalId,
        }),
      );
    }
  });

  watch(data, (newData) => {
    if (!newData) return;

    try {
      const parsed = JSON.parse(newData);

      const patientId = parsed.patient_id;
      if (!patientId) return;

      vitalsMap.value[patientId] = {
        ...(vitalsMap.value[patientId] || {}),
        ...parsed,
      };
    } catch (err) {
      console.error("Invalid WS data:", err);
    }
  });

  socketInstance = {
    vitalsMap,
    status,
  };

  return socketInstance;
}
