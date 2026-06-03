<script setup>
definePageMeta({
  layout: "dashboard",
});

import { Heart, Activity, Thermometer, Droplets } from "lucide-vue-next";

const user = useAuthUser();
const hospitalID = user.value?.hospital?.id;

const { vitalsMap, status } = useVitalsSocket(hospitalID);
const { fetchVitals } = useVitalSignService();

const loadingVitals = ref(false);
const vitals = ref([]);

/* ---------------------------
   MAP BACKEND → UI FORMAT
---------------------------- */
const buildPatientCard = (p) => {
  return {
    id: p.patient_id, // IMPORTANT FIX
    initials: p.patient_name
      .split(" ")
      .map((w) => w[0])
      .join("")
      .toUpperCase(),

    name: p.patient_name,
    room: "-",
    status: p.status || "Stable",

    vitals: [
      {
        label: "Heart Rate",
        value: p.heart_rate,
        unit: "bpm",
        trend: p.heart_rate > 100 ? "up" : "down",
        icon: Heart,
        iconClass: "text-red-500",
      },
      {
        label: "Blood Pressure",
        value: `${p.systolic_bp}/${p.diastolic_bp}`,
        unit: "mmHg",
        trend: "down",
        icon: Activity,
        iconClass: "text-blue-500",
      },
      {
        label: "Temperature",
        value: p.temperature,
        unit: "°C",
        trend: p.temperature > 37 ? "up" : "down",
        icon: Thermometer,
        iconClass: "text-orange-400",
      },
      {
        label: "SpO2",
        value: p.spo2,
        unit: "%",
        trend: p.spo2 < 95 ? "down" : "up",
        icon: Droplets,
        iconClass: "text-cyan-500",
      },
    ],
  };
};

/* ---------------------------
   FETCH VITALS
---------------------------- */
const getVitalSigns = async () => {
  try {
    loadingVitals.value = true;

    const res = await fetchVitals();

    const raw = Array.isArray(res) ? res : (res?.data ?? []);

    vitals.value = raw.map(buildPatientCard);

    console.log("NORMALIZED VITALS:", vitals.value);
  } catch (err) {
    console.error("Failed to fetch vitals:", err);
    vitals.value = [];
  } finally {
    loadingVitals.value = false;
  }
};

/* FIX: CALL CORRECT FUNCTION */
onMounted(() => {
  getVitalSigns();
});

/* ---------------------------
   LIVE PATIENTS (WS UPDATE)
---------------------------- */
const livePatients = computed(() => {
  return vitals.value.map((patient) => {
    const live = vitalsMap.value?.[patient.id];

    if (!live) return patient;

    return {
      ...patient,
      vitals: patient.vitals.map((v) => {
        if (v.label === "Heart Rate") {
          return {
            ...v,
            value: live.heart_rate ?? v.value,
            trend: live.heart_rate > 100 ? "up" : "down",
          };
        }

        if (v.label === "SpO2") {
          return {
            ...v,
            value: live.spo2 ?? v.value,
            trend: live.spo2 < 95 ? "down" : "up",
          };
        }

        if (v.label === "Temperature") {
          return {
            ...v,
            value: live.temperature ?? v.value,
          };
        }

        return v;
      }),
    };
  });
});
</script>

<template>
  <div
    class="mb-6 flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between"
  >
    <div>
      <h1 class="text-2xl font-bold tracking-tight">Vitals Monitor</h1>
      <p class="mt-1 text-sm text-muted-foreground">
        Real-time patient monitoring
      </p>
      <p>WS Status: {{ status }}</p>
    </div>
    <div
      class="inline-flex items-center rounded-full border font-medium transition-colors focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 border-transparent bg-slate-50 dark:text-secondary-foreground w-fit gap-1.5 px-3 py-1 text-xs"
    >
      <Icon name="lucide:activity" class="h-3.5 w-3.5" />

      12 Active Monitors
    </div>
  </div>
  <div class="grid grid-cols-1 gap-4 md:grid-cols-2 xl:grid-cols-3">
    <UiPatientVitals
      v-for="patient in livePatients"
      :key="patient.name"
      :patient="patient"
    />
  </div>
</template>
