<script setup>
definePageMeta({
  layout: "dashboard",
});

const search = ref("");
const isSessionModalOpen = ref(false);
const admittedPatients = ref([]);
const selectedPatient = ref(null);

const { getAssignedPatients } = useAdmittedPatientService();
const localVideoRef = ref(null);
const { setRobot, connectWS, startWebRTC, videoRef, robotId } =
  useTelepresence();
const isStreamActive = ref(false);
const { getRobots } = useRobotService();
const loadingPatients = ref(false);

const getAdmittedPatients = async () => {
  try {
    loadingPatients.value = true;

    const response = await getAssignedPatients();

    admittedPatients.value = response.map((patient) => ({
      id: patient.patient_id,
      name: patient.full_name,
      gender: patient.gender,
      diagnosis: patient.diagnosis,
      room: patient.room_number,
      room_id: patient.room_id,
      doctorName: patient.assigned_doctor_name,
      urgency: patient.urgency,
      status: patient.admission_status,
      heart_rate: patient.heart_rate,
      spo2: patient.spo2,
      temperature: patient.temperature,
      // systolic_bp: patient.systolic_bp,
      // diastolic_bp: patient.diastolic_bp,

      robot: "AURA-01",
    }));

    if (admittedPatients.value.length > 0) {
      selectedPatient.value = admittedPatients.value[0];
    }

    console.log("Admitted Patients:", admittedPatients.value);
  } catch (error) {
    console.log("Fetch Patients Error:", error);
  } finally {
    loadingPatients.value = false;
  }
};
const loadingRobots = ref(false);
const robots = ref([]);
const getAvailableRobots = async () => {
  try {
    loadingRobots.value = true;

    const response = await getRobots();

    robots.value = response.map((robot) => ({
      id: robot.id,
      name: robot.name,
      status: robot.status,
    }));

    console.log("Available Robots:", robots.value);
  } catch (error) {
    console.log("Fetch Robots Error:", error);
  } finally {
    loadingRobots.value = false;
  }
};
const metrics = [
  {
    title: "Live Consultations",
    value: "3",
    icon: "lucide:video",
    colorTheme: "bg-rose-50 text-rose-500",
  },

  {
    title: "Patients Waiting",
    value: "5",
    icon: "mdi:patient",
    colorTheme: "bg-amber-50 text-amber-500",
  },

  {
    title: "Robots Available",
    value: "2",
    icon: "material-symbols:robot",
    colorTheme: "bg-emerald-50 text-emerald-500",
  },

  {
    title: "Emergency Requests",
    value: "1",
    icon: "lucide:siren",
    colorTheme: "bg-violet-50 text-violet-500",
  },
];
const vitals = ref({});
onMounted(() => {
  getAdmittedPatients();
  getAvailableRobots();
});
watch(
  admittedPatients,
  (list) => {
    if (list.length && !selectedPatient.value) {
      selectedPatient.value = list[0];
    }
  },
  { deep: true },
);
watch(
  selectedPatient,
  (patient) => {
    if (!patient) return;

    vitals.value = {
      heartRate: patient.heart_rate,
      spo2: patient.spo2,
      temperature: patient.temperature,
      // bloodPressure: patient.systolic_bp
      //   ? `${patient.systolic_bp}/${patient.diastolic_bp}`
      //   : "—",
    };
  },
  { immediate: true },
);
// const openSessionModal = () => {
//   isSessionModalOpen.value = true;
//   // pick robot from selected patient
//   robotId.value = selectedPatient.value?.robot || "robot-1";

//   // start telepresence
//   connect();
// };
const openSessionModal = async () => {
  isSessionModalOpen.value = true;
  robotId.value = selectedPatient.value?.robot || "robot-1";
  setRobot(robotId.value);

  try {
    console.log("[Page] Connecting to WebSocket...");
    await connectWS();
    console.log("[Page] WebSocket ready. Starting WebRTC handshake...");
    await startWebRTC();
    const checkStreamInterval = setInterval(() => {
      if (videoRef.value && videoRef.value.srcObject) {
        console.log("[Page] Live Video source detected! Revealing container.");
        isStreamActive.value = true;
        clearInterval(checkStreamInterval);
      }
    }, 200);
  } catch (error) {
    console.error("[Page] Failed to start telepresence session:", error);
  }
};
</script>

<template>
  <ModalsTelepresenceSession
    v-model="isSessionModalOpen"
    :patient="selectedPatient"
  />

  <div class="space-y-6">
    <div
      class="flex flex-col gap-4 lg:flex-row lg:items-center lg:justify-between"
    >
      <div>
        <h1 class="text-3xl font-bold tracking-tight">Robot Telepresence</h1>

        <p class="mt-1 text-sm text-gray-500">
          Remote patient consultation, robot monitoring, and live interaction
          center.
        </p>
      </div>

      <div class="flex items-center gap-3">
        <div class="relative">
          <UiBaseInput
            v-model="search"
            class="w-72 pl-10"
            placeholder="Search patients or sessions..."
            leading-icon="lucide:search"
            leadingIconClass="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-gray-400"
          />
        </div>

        <button
          type="button"
          @click="openSessionModal"
          class="inline-flex items-center gap-2 rounded-xl bg-primary px-5 py-3 text-sm font-medium text-white transition hover:opacity-90"
        >
          <Icon name="lucide:video" class="h-4 w-4" />
          Start Session
        </button>
      </div>
    </div>
    <div class="grid grid-cols-1 gap-5 md:grid-cols-2 xl:grid-cols-4">
      <UiMetricCard
        v-for="metric in metrics"
        :key="metric.title"
        :title="metric.title"
        :value="metric.value"
        :icon="metric.icon"
        :colorTheme="metric.colorTheme"
      />
    </div>
    <div class="grid grid-cols-12 gap-6">
      <div
        class="col-span-12 rounded-3xl border border-gray-200 bg-white p-5 shadow-sm xl:col-span-3"
      >
        <div class="mb-5 flex items-center justify-between">
          <h2 class="text-lg font-semibold">Assigned Patients</h2>

          <span
            class="rounded-full bg-primary/10 px-3 py-1 text-xs font-medium text-primary"
          >
            {{ admittedPatients.length }}
          </span>
        </div>

        <div class="space-y-4">
          <div
            v-for="patient in admittedPatients"
            :key="patient.id"
            @click="selectedPatient = patient"
            class="cursor-pointer rounded-2xl border p-4 transition-all duration-200 hover:shadow-md"
            :class="
              selectedPatient?.id === patient.id
                ? 'border-primary bg-primary/5'
                : 'border-gray-200'
            "
          >
            <div class="flex items-start justify-between">
              <div>
                <h3 class="font-semibold">
                  {{ patient.name }}
                </h3>

                <p class="mt-1 text-sm text-gray-500">
                  Room {{ patient.room }}
                </p>
              </div>

              <div
                class="h-3 w-3 rounded-full"
                :class="
                  patient.status === 'online'
                    ? 'bg-green-500'
                    : patient.status === 'busy'
                      ? 'bg-red-500'
                      : 'bg-yellow-500'
                "
              />
            </div>

            <div class="mt-4 flex items-center justify-between">
              <span
                class="rounded-full px-3 py-1 text-xs font-medium"
                :class="
                  patient.urgency === 'Critical'
                    ? 'bg-red-100 text-red-700'
                    : patient.urgency === 'Normal'
                      ? 'bg-blue-100 text-blue-700'
                      : 'bg-gray-100 text-gray-700'
                "
              >
                {{ patient?.urgency }}
              </span>

              <div class="flex items-center gap-2 text-xs text-gray-500">
                <Icon name="lucide:bot" class="h-4 w-4" />
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="col-span-12 xl:col-span-6">
        <div
          class="overflow-hidden rounded-3xl border border-gray-200 bg-black shadow-sm"
        >
          <div
            class="flex items-center justify-between border-b border-white/10 px-6 py-4"
          >
            <div>
              <h2 class="text-lg font-semibold text-white">
                {{ selectedPatient?.name || "No patient selected" }}
              </h2>

              <p class="text-sm text-gray-400">
                Room {{ selectedPatient?.room || "-" }}
              </p>
            </div>

            <div class="flex items-center gap-2">
              <div class="h-2.5 w-2.5 rounded-full bg-red-500 animate-pulse" />

              <span class="text-sm font-medium text-red-400"> LIVE </span>
            </div>
          </div>
          <div
            class="relative flex h-[500px] items-center justify-center bg-gradient-to-br from-gray-900 via-black to-gray-950"
          >
            <!-- 🔥 THE FIX: Bind directly to 'videoRef' from the composable, toggle visibility using 'isStreamActive' -->
            <video
              ref="videoRef"
              autoplay
              playsinline
              muted
              class="h-full w-full object-cover"
              :class="{ hidden: !isStreamActive }"
            />

            <!-- 🔥 THE FIX: Fallback overlay now tracks the 'isStreamActive' state flag -->
            <div
              v-if="!isStreamActive"
              class="text-center absolute inset-0 flex flex-col items-center justify-center bg-black/40"
            >
              <Icon
                name="lucide:video"
                class="mx-auto h-20 w-20 text-white/20"
              />

              <p class="mt-4 text-lg font-medium text-white/80">
                Live Robot Camera Stream
              </p>

              <p class="mt-1 text-sm text-gray-500">
                Waiting for robot connection...
              </p>
            </div>

            <div
              class="absolute bottom-5 left-5 rounded-2xl bg-white/10 px-4 py-3 backdrop-blur-md"
            >
              <div class="flex items-center gap-3">
                <div class="h-3 w-3 rounded-full bg-green-400" />
                <div>
                  <p class="text-xs text-gray-300">Connection Stable</p>
                  <p class="text-sm font-medium text-white">1080p · 40ms</p>
                </div>
              </div>
            </div>
          </div>
          <div
            class="flex flex-wrap items-center justify-center gap-3 border-t border-white/10 p-5"
          >
            <button
              class="rounded-xl bg-white/10 px-5 py-3 text-sm font-medium text-white backdrop-blur hover:bg-white/20"
            >
              <Icon name="lucide:mic" class="mr-2 inline h-4 w-4" />
              Audio
            </button>

            <button
              class="rounded-xl bg-white/10 px-5 py-3 text-sm font-medium text-white backdrop-blur hover:bg-white/20"
            >
              <Icon name="lucide:move" class="mr-2 inline h-4 w-4" />
              Move Robot
            </button>

            <button
              class="rounded-xl bg-emerald-500 px-5 py-3 text-sm font-medium text-white hover:bg-emerald-600"
            >
              <Icon name="lucide:activity" class="mr-2 inline h-4 w-4" />
              Measure Vitals
            </button>

            <button
              class="rounded-xl bg-red-500 px-5 py-3 text-sm font-medium text-white hover:bg-red-600"
            >
              <Icon name="lucide:siren" class="mr-2 inline h-4 w-4" />
              Emergency
            </button>
          </div>
        </div>
      </div>
      <div class="col-span-12 space-y-6 xl:col-span-3">
        <div class="rounded-3xl border border-gray-200 bg-white p-5 shadow-sm">
          <div class="mb-5 flex items-center justify-between">
            <h2 class="text-lg font-semibold">Live Vitals</h2>

            <span
              class="rounded-full bg-emerald-100 px-3 py-1 text-xs font-medium text-emerald-700"
            >
              Real-time
            </span>
          </div>

          <div class="space-y-4">
            <div
              class="flex items-center justify-between rounded-2xl bg-gray-50 p-4"
            >
              <div>
                <p class="text-sm text-gray-500">Heart Rate</p>

                <p class="text-xl font-bold">
                  {{ vitals?.heartRate || "—" }}
                </p>
              </div>

              <Icon name="lucide:heart-pulse" class="h-6 w-6 text-rose-500" />
            </div>

            <div
              class="flex items-center justify-between rounded-2xl bg-gray-50 p-4"
            >
              <div>
                <p class="text-sm text-gray-500">SpO2</p>

                <p class="text-xl font-bold">{{ vitals?.spo2 || "—" }}%</p>
              </div>

              <Icon name="lucide:activity" class="h-6 w-6 text-blue-500" />
            </div>

            <div
              class="flex items-center justify-between rounded-2xl bg-gray-50 p-4"
            >
              <div>
                <p class="text-sm text-gray-500">Temperature</p>

                <p class="text-xl font-bold">
                  {{ vitals?.temperature || "—" }}°C
                </p>
              </div>

              <Icon name="lucide:thermometer" class="h-6 w-6 text-orange-500" />
            </div>
          </div>
        </div>
        <div class="rounded-3xl border border-gray-200 bg-white p-5 shadow-sm">
          <h2 class="mb-5 text-lg font-semibold">Robot Status</h2>

          <div class="space-y-4">
            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-500">Robot</span>

              <span class="font-semibold"> </span>
            </div>

            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-500">Battery</span>

              <span class="font-semibold text-emerald-600">
                {{ vitals?.battery || "—" }}%
              </span>
            </div>

            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-500">Latency</span>

              <span class="font-semibold">
                {{ vitals?.latency || "—" }}
              </span>
            </div>

            <div class="flex items-center justify-between">
              <span class="text-sm text-gray-500">Connection</span>

              <span
                class="rounded-full bg-green-100 px-3 py-1 text-xs font-medium text-green-700"
              >
                Stable
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
