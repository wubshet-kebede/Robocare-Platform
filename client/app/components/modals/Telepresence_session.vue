<script setup>
const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false,
  },
  patient: Object,
});
console.log(props.patient?.id);
console.log(props.patient?.name);
console.log("the props we accepted from the parent is", props.patient);
const emits = defineEmits(["update:modelValue"]);

const isOpen = computed({
  get() {
    return props.modelValue;
  },
  set(value) {
    emits("update:modelValue", value);
  },
});
const { getRobots } = useRobotService();
const { publishNavGoal } = useRobotService();
const loadingRobots = ref(false);
const robots = ref([]);
const getAvailableRobots = async () => {
  try {
    loadingRobots.value = true;

    const response = await getRobots();
    console.log("API Response:", response);
    robots.value = response.map((robot) => ({
      id: robot.id,
      name: robot.model,
      status: robot.status,
    }));

    console.log("Available Robots:", robots.value);
  } catch (error) {
    console.log("Fetch Robots Error:", error);
  } finally {
    loadingRobots.value = false;
  }
};

const loadingSubmit = ref(false);
const values = reactive({
  patient: "",
  roomId: "",
  robot: "",
  notes: "",
});
const robotOptions = computed(() =>
  robots.value.map((robot) => ({
    id: robot.id,
    name: robot.name,
  })),
);
console.log("the manual case", robotOptions.value);
// const sessionTypeOptions = [
//   {
//     id: "Routine Check",
//     name: "Routine Check",
//   },

//   {
//     id: "Follow-up",
//     name: "Follow-up",
//   },

//   {
//     id: "Emergency Consultation",
//     name: "Emergency Consultation",
//   },
// ];

// const priorityOptions = [
//   {
//     id: "Low",
//     name: "Low",
//   },

//   {
//     id: "Medium",
//     name: "Medium",
//   },

//   {
//     id: "High",
//     name: "High",
//   },

//   {
//     id: "Critical",
//     name: "Critical",
//   },
// ];

watch(
  () => props.patient,
  (patient) => {
    if (!patient) return;

    values.patient = patient.id;
    values.roomId = patient.room_id;
  },
  { immediate: true },
);
const submit = async () => {
  try {
    loadingSubmit.value = true;

    console.log("Patient ID:", values.patient);
    console.log("Robot ID:", values.robot);
    console.log("Room ID:", values.roomId);

    const response = await publishNavGoal({
      patient_id: values.patient,
      robot_id: values.robot,
      room_id: values.roomId,
    });

    console.log("Navigation Goal Response:", response);

    isOpen.value = false;
  } catch (error) {
    console.error(error);
  } finally {
    loadingSubmit.value = false;
  }
};
onMounted(() => {
  getAvailableRobots();
});
</script>

<template>
  <ModalsModal
    v-model="isOpen"
    title="Start Telepresence Session"
    wrapperClass="max-w-4xl"
  >
    <template #content>
      <div class="p-8 space-y-8">
        <div>
          <h2 class="text-xl font-semibold mb-6">Session Information</h2>

          <div class="grid grid-cols-2 gap-6">
            <!-- ROBOT -->
            <UiListSelect
              v-model="values.robot"
              :items="robotOptions"
              name="robot"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">Select Robot</h1>
              </template>
            </UiListSelect>
            <!-- <UiListSelect
              v-model="values.sessionType"
              :items="sessionTypeOptions"
              name="sessionType"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">Session Type</h1>
              </template>
            </UiListSelect>
            <UiListSelect
              v-model="values.priority"
              :items="priorityOptions"
              name="priority"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">Priority Level</h1>
              </template>
            </UiListSelect> -->
            <UiBaseInput
              modelValue="Ready for Navigation"
              name="status"
              disabled
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">Robot Status</h1>
              </template>
            </UiBaseInput>
          </div>
        </div>

        <div>
          <h2 class="text-xl font-semibold mb-6">Consultation Notes</h2>

          <textarea
            v-model="values.notes"
            rows="5"
            placeholder="Enter session notes or special instructions..."
            class="w-full rounded-xl border border-gray-300 px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-primary/20 focus:border-primary"
          />
        </div>

        <div class="flex justify-end gap-4 pt-6">
          <button
            @click="isOpen = false"
            class="rounded-xl border border-gray-300 px-6 py-3 hover:bg-gray-100 transition"
            :disabled="loadingSubmit"
          >
            Cancel
          </button>

          <button
            @click="submit"
            class="rounded-xl bg-primary px-6 py-3 text-white hover:opacity-90 transition"
            :disabled="loadingSubmit"
          >
            {{
              loadingSubmit
                ? "Dispatching Robot..."
                : "Dispatch Robot & Start Session"
            }}
          </button>
        </div>
      </div>
    </template>
  </ModalsModal>
</template>
