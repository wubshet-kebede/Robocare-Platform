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
const emits = defineEmits(["update:modelValue"]);

const isOpen = computed({
  get() {
    return props.modelValue;
  },
  set(value) {
    emits("update:modelValue", value);
  },
});

const loadingSubmit = ref(false);

const values = reactive({
  patient: "",
  room: "",
  robot: "",
  sessionType: "",
  priority: "",
  notes: "",
});

const patientOptions = [
  {
    id: "patient-1",
    name: "John Doe",
    room: "ICU-101",
    department: "ICU",
  },

  {
    id: "patient-2",
    name: "Sarah Johnson",
    room: "ER-203",
    department: "Emergency",
  },

  {
    id: "patient-3",
    name: "Michael Brown",
    room: "WARD-12",
    department: "Cardiology",
  },
];

const robotOptions = [
  {
    id: "robot-1",
    name: "MediRover-01",
  },

  {
    id: "robot-2",
    name: "MediRover-02",
  },

  {
    id: "robot-3",
    name: "MediRover-03",
  },
];

const sessionTypeOptions = [
  {
    id: "Routine Check",
    name: "Routine Check",
  },

  {
    id: "Follow-up",
    name: "Follow-up",
  },

  {
    id: "Emergency Consultation",
    name: "Emergency Consultation",
  },
];

const priorityOptions = [
  {
    id: "Low",
    name: "Low",
  },

  {
    id: "Medium",
    name: "Medium",
  },

  {
    id: "High",
    name: "High",
  },

  {
    id: "Critical",
    name: "Critical",
  },
];

watch(
  () => values.patient,
  (newPatient) => {
    const selectedPatient = patientOptions.find(
      (patient) => patient.id === newPatient,
    );

    if (selectedPatient) {
      values.room = selectedPatient.room;
    }
  },
);

const submit = async () => {
  try {
    loadingSubmit.value = true;

    console.log("Telepresence Session:", values);

    setTimeout(() => {
      loadingSubmit.value = false;
      isOpen.value = false;
    }, 1500);
  } catch (error) {
    console.log(error);
    loadingSubmit.value = false;
  }
};
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
            <UiListSelect
              v-model="values.patient"
              :items="patientOptions"
              name="patient"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  Select Patient
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </UiListSelect>

            <!-- ROBOT -->
            <UiListSelect
              v-model="values.robot"
              :items="robotOptions"
              name="robot"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  Select Robot
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </UiListSelect>
            <UiListSelect
              v-model="values.sessionType"
              :items="sessionTypeOptions"
              name="sessionType"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  Session Type
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </UiListSelect>
            <UiListSelect
              v-model="values.priority"
              :items="priorityOptions"
              name="priority"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  Priority Level
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </UiListSelect>
          </div>
        </div>
        <div>
          <h2 class="text-xl font-semibold mb-6">Robot Destination</h2>

          <div class="grid grid-cols-2 gap-6">
            <UiBaseInput v-model="values.room" name="room" disabled>
              <template #label>
                <h1 class="text-md font-medium mb-2">Patient Room</h1>
              </template>
            </UiBaseInput>

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
