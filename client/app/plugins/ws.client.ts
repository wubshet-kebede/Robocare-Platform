import { useWebSocket } from "@vueuse/core";

let socketInitialized = false;

export default defineNuxtPlugin(() => {
  const vitalsMap = useState<Record<string, any>>("vitalsMap", () => ({}));
  const user = useAuthUser();

  const connect = (hospitalId: string) => {
    const { data } = useWebSocket(
      () => `ws://localhost:8082/ws/vitals?hospital_id=${hospitalId}`,
      {
        autoReconnect: {
          retries: 10,
          delay: 2000,
        },
      },
    );

    watch(data, (msg) => {
      if (!msg) return;

      try {
        const parsed = JSON.parse(msg);
        const patientId = parsed.patient_id;

        if (!patientId) return;

        vitalsMap.value[patientId] = {
          ...(vitalsMap.value[patientId] || {}),
          ...parsed,
        };
      } catch (e) {
        console.error("WS parse error:", e);
      }
    });
  };

  if (!socketInitialized) {
    socketInitialized = true;

    watch(
      () => user.value?.hospital?.id,
      (hospitalId) => {
        if (!hospitalId) return;

        connect(hospitalId);
      },
      { immediate: true },
    );
  }
});
