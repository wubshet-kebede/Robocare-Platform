import { useWebSocket } from "@vueuse/core";
import { computed, ref, watch } from "vue";

let socketInstance: any = null;

export function useVitalsSocket(hospitalId: string) {
  if (socketInstance) {
    return socketInstance;
  }

  const vitalsMap = ref<Record<string, any>>({});

  const url = `ws://localhost:8082/ws/vitals?hospital_id=${hospitalId}`;

  const { data, status } = useWebSocket(url, {
    autoReconnect: {
      retries: 10,
      delay: 2000,
    },
  });

  watch(data, (newData) => {
    if (!newData) return;

    console.log("WS DATA:", newData);

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
