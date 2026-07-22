<script setup>
import { useForm } from "vee-validate";

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false,
  },

  patient: {
    type: Object,
    required: true,
  },
  assignableStaff: {
    type: Array,
    default: () => [],
  },
});
console.log("the props we accepted from the parent is", props.assignableStaff);
const emits = defineEmits(["update:modelValue", "success"]);

const isOpen = computed({
  get() {
    return props.modelValue;
  },

  set(value) {
    emits("update:modelValue", value);
  },
});

const loadingSubmit = ref(false);
const selectedStaffId = ref("");
const staffOptions = computed(() => {
  return props.assignableStaff.map((staff) => ({
    id: staff.id,
    name: staff.full_name,
  }));
});
const { admitPatient } = usePatientService();
const { getRooms } = useRoomService();
const rooms = ref([]);
const loadingRooms = ref(false);

const fetchRooms = async () => {
  try {
    loadingRooms.value = true;

    const res = await getRooms();

    console.log("RAW ROOMS:", res);
    rooms.value = Array.isArray(res) ? res : (res?.data ?? []);

    console.log("NORMALIZED ROOMS:", rooms.value);
  } catch (err) {
    console.error("Failed to fetch rooms:", err);
    rooms.value = [];
  } finally {
    loadingRooms.value = false;
  }
};

onMounted(() => {
  fetchRooms();
});
const roomOptions = computed(() => {
  if (!Array.isArray(rooms.value)) return [];

  return rooms.value.map((room) => ({
    id: room.id,
    name: room.room_number + " - " + room.location_name,
  }));
});
const { handleSubmit, resetForm, values } = useForm({
  initialValues: {
    selectedStaffId: "",
    roomId: "",
    bedNumber: "",
    diagnosis: "",
    reasonForAdmission: "",
    urgency: "Normal",
    status: "stable",
    admissionStatus: "Waiting",
    admissionDate: "",
  },
});

const submit = handleSubmit(async (formValues) => {
  try {
    loadingSubmit.value = true;

    const payload = {
      patient_id: props.patient.id,
      assigned_doctor_id: formValues.selectedStaffId,
      room_id: formValues.roomId,
      bed_number: formValues.bedNumber,
      diagnosis: formValues.diagnosis,
      reason_for_admission: formValues.reasonForAdmission,
      urgency: formValues.urgency,
      status: formValues.status,
      admission_status: formValues.admissionStatus,
      admission_date: new Date(formValues.admissionDate).toISOString(),
    };
    const response = await admitPatient(payload);

    console.log("Admission Created:", response);

    resetForm();

    isOpen.value = false;

    emits("success", response);
  } catch (error) {
    console.log("Admission Creation Error:", error);
  } finally {
    loadingSubmit.value = false;
  }
});

const urgencyOptions = [
  {
    id: "Normal",
    name: "Normal",
  },

  {
    id: "Urgent",
    name: "Urgent",
  },

  {
    id: "Emergency",
    name: "Emergency",
  },
];

const statusOptions = [
  {
    id: "stable",
    name: "Stable",
  },

  {
    id: "critical",
    name: "Critical",
  },

  {
    id: "warning",
    name: "Warning",
  },
  {
    id: "unknown",
    name: "Unknown",
  },
];
const admissionStatusOptions = [
  {
    id: "Waiting",
    name: "Waiting",
  },

  {
    id: "Admitted",
    name: "Admitted",
  },
  {
    id: "Discharged",
    name: "Discharged",
  },

  {
    id: "Deceased",
    name: "Deceased",
  },
];
</script>

<template>
  <ModalsModal
    v-model="isOpen"
    title="Patient Admission"
    wrapperClass="w-full max-w-4xl"
  >
    <template #content>
      <div class="p-4 sm:p-6 lg:p-8 space-y-6 lg:space-y-8">
        <div>
          <h2 class="text-lg sm:text-xl font-semibold mb-4 sm:mb-6">
            Patient Information
          </h2>

          <div
            class="rounded-2xl border border-gray-200 bg-gray-50 p-4 sm:p-5 lg:p-6"
          >
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 sm:gap-6">
              <div>
                <p class="text-sm text-gray-400">Patient Name</p>

                <p class="font-semibold text-lg">
                  {{ patient.fullName }}
                </p>
              </div>

              <div>
                <p class="text-sm text-gray-400">Gender</p>

                <p class="font-semibold text-lg capitalize">
                  {{ patient.gender }}
                </p>
              </div>

              <div>
                <p class="text-sm text-gray-400">Phone</p>

                <p class="font-semibold text-lg">
                  {{ patient.phone }}
                </p>
              </div>

              <div>
                <p class="text-sm text-gray-400">Blood Type</p>

                <p class="font-semibold text-lg">
                  {{ patient.bloodType }}
                </p>
              </div>
            </div>
          </div>
        </div>
        <div>
          <h2 class="text-lg sm:text-xl font-semibold mb-4 sm:mb-6">
            Admission Information
          </h2>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-4 sm:gap-6">
            <UiListSelect
              v-model="values.selectedStaffId"
              :items="staffOptions"
              name="selectedStaffId"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">Assigned Doctor</h1>
              </template>
            </UiListSelect>
            <UiListSelect
              v-model="values.roomId"
              :items="roomOptions"
              name="roomId"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">Room</h1>
              </template>
            </UiListSelect>
            <UiBaseInput
              v-model="values.bedNumber"
              name="bedNumber"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">Bed Number</h1>
              </template>
            </UiBaseInput>
            <UiBaseInput
              v-model="values.admissionDate"
              name="admissionDate"
              type="datetime-local"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">Admission Date</h1>
              </template>
            </UiBaseInput>
          </div>
        </div>
        <div>
          <h2 class="text-xl font-semibold mb-6">Medical Details</h2>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4 sm:gap-6">
            <UiListSelect
              v-model="values.status"
              :items="statusOptions"
              name="status"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">Patient Status</h1>
              </template>
            </UiListSelect>
            <UiListSelect
              v-model="values.urgency"
              :items="urgencyOptions"
              name="urgency"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">Urgency Level</h1>
              </template>
            </UiListSelect>
            <UiListSelect
              v-model="values.admissionStatus"
              :items="admissionStatusOptions"
              name="admissionStatus"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">Admission Status</h1>
              </template>
            </UiListSelect>
            <UiBaseInput
              v-model="values.diagnosis"
              name="diagnosis"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">Diagnosis</h1>
              </template>
            </UiBaseInput>
            <div class="md:col-span-2">
              <UiBaseInput
                v-model="values.reasonForAdmission"
                name="reasonForAdmission"
                rules="required"
              >
                <template #label>
                  <h1 class="text-md font-medium mb-2">Reason For Admission</h1>
                </template>
              </UiBaseInput>
            </div>
          </div>
        </div>
        <div
          class="flex flex-col-reverse sm:flex-row justify-end gap-3 sm:gap-4 pt-6"
        >
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
            {{ loadingSubmit ? "Creating Admission..." : "Admit Patient" }}
          </button>
        </div>
      </div>
    </template>
  </ModalsModal>
</template>
