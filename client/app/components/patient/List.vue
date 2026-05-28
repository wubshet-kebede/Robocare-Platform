<script setup>
import { ref, computed, onMounted } from "vue";

const activeTab = ref("registered");
const tabs = [
  { label: "Registered", value: "registered" },
  { label: "Admitted", value: "admitted" },
  { label: "Critical", value: "critical" },
  { label: "Discharged", value: "discharged" },
];
const patientsnew = ref([]);
const loadingPatients = ref(false);

const { fetchPatients } = usePatientService();
const { getAssignableStaff } = useStaffService();

const assignableStaff = ref([]);

const isAdmissionModalOpen = ref(false);
const selectedPatient = ref(null);

const calculateAge = (dob) => {
  const birth = new Date(dob);
  const today = new Date();

  let age = today.getFullYear() - birth.getFullYear();
  const m = today.getMonth() - birth.getMonth();

  if (m < 0 || (m === 0 && today.getDate() < birth.getDate())) {
    age--;
  }

  return age;
};
const getInitials = (fullName) => {
  if (!fullName) return "";

  return fullName
    .trim()
    .split(" ")
    .filter(Boolean)
    .map((word) => word[0])
    .join("")
    .toUpperCase();
};
const formatPatients1 = (data) => {
  return data.map((item) => {
    const p = item.patient;
    const admission = item.active_admission;

    return {
      id: p.id,
      fullName: p.full_name,
      initials: getInitials(p.full_name),
      age: calculateAge(p.date_of_birth),
      gender: p.gender,
      bloodType: p.blood_type,
      allergies: p.allergies,
      phone: p.phone,
      emergencyContact: p.emergency_contact_name,

      admission: admission
        ? {
            status: admission.admission_status,
            urgency: admission.urgency,
            diagnosis: admission.diagnosis,
            assignedDoctor: admission.assigned_doctor_name,
            room: admission.room_number,
            staffName: admission.staff_name,
            bed: admission.bed_number,
          }
        : null,
    };
  });
};
const formatPatients = (data) => {
  return data.map((p) => {
    const hasAdmission = p.admission_status !== null;

    return {
      id: p.patient_id,

      fullName: p.full_name,

      initials: getInitials(p.full_name),

      age: calculateAge(p.date_of_birth),

      gender: p.gender,

      bloodType: p.blood_type,

      allergies: p.allergies,

      phone: p.phone,

      emergencyContact: p.emergency_contact,

      admission: hasAdmission
        ? {
            status: p.admission_status,

            urgency: p.urgency,

            diagnosis: p.diagnosis,

            assignedDoctor: p.assigned_doctor_name,

            room: p.room_number,

            bed: p.bed_number,
          }
        : null,
    };
  });
};

const getPatients = async () => {
  try {
    loadingPatients.value = true;

    const response = await fetchPatients();
    console.log("Raw Patients Data:", response);

    patientsnew.value = formatPatients(response);

    console.log("Fetched Patients:", patientsnew.value);
  } catch (error) {
    console.log("Fetch Patients Error:", error);
  } finally {
    loadingPatients.value = false;
  }
};

const fetchAssignableStaff = async () => {
  try {
    const response = await getAssignableStaff();

    assignableStaff.value = response;

    console.log("Assignable Staff:", response);
  } catch (error) {
    console.log("Fetch Assignable Staff Error:", error);
  }
};
const openAdmissionModal = (patient) => {
  selectedPatient.value = patient;
  isAdmissionModalOpen.value = true;
};
const filteredPatients = computed(() => {
  console.log("patientsnew:", patientsnew.value);
  if (activeTab.value === "registered") {
    return patientsnew.value.filter((p) => p.admission === null);
  }

  if (activeTab.value === "admitted") {
    return patientsnew.value.filter(
      (p) => p.admission && p.admission.status !== "discharged",
    );
  }

  if (activeTab.value === "critical") {
    return patientsnew.value.filter((p) => p.admission?.status === "critical");
  }

  if (activeTab.value === "discharged") {
    return patientsnew.value.filter(
      (p) => p.admission?.status === "discharged",
    );
  }

  return patientsnew.value;
});
console.log("Filtered Patients:", filteredPatients.value);
const getCount = (status) => {
  if (status === "registered") {
    return patientsnew.value.filter((p) => p.admission === null).length;
  }

  if (status === "admitted") {
    return patientsnew.value.filter(
      (p) => p.admission && p.admission.status !== "discharged",
    ).length;
  }

  if (status === "critical") {
    return patientsnew.value.filter((p) => p.admission?.status === "critical")
      .length;
  }

  if (status === "discharged") {
    return patientsnew.value.filter((p) => p.admission?.status === "discharged")
      .length;
  }

  return 0;
};

const statusClasses = {
  active: "bg-green-100 text-green-700",
  critical: "bg-red-100 text-red-700",
  recovering: "bg-yellow-100 text-yellow-700",
  discharged: "bg-gray-100 text-gray-700",
  registered: "bg-blue-100 text-blue-700",
};

onMounted(() => {
  getPatients();
  fetchAssignableStaff();
});
const formatPatientId = (id) => {
  if (typeof id !== "string") return "";

  return `PID-${id.slice(0, 6).toUpperCase()}`;
};
</script>

<template>
  <ModalsPatientAdmission
    v-if="selectedPatient"
    v-model="isAdmissionModalOpen"
    :patient="selectedPatient"
    :assignable-staff="assignableStaff"
  />
  <div class="bg-[#faf7f7] min-h-screen p-6">
    <div class="grid grid-cols-12 gap-6">
      <div class="col-span-8 space-y-4">
        <div class="inline-flex rounded-xl bg-[#f1ecec] p-1 gap-1">
          <button
            v-for="tab in tabs"
            :key="tab.value"
            @click="activeTab = tab.value"
            class="flex items-center rounded-lg px-4 py-2 text-sm font-medium transition-all duration-200"
            :class="
              activeTab === tab.value
                ? 'bg-white shadow-sm text-black'
                : 'text-gray-500 hover:text-black'
            "
          >
            {{ tab.label }}

            <span
              class="ml-2 rounded-md px-2 py-0.5 text-[10px] font-semibold"
              :class="
                tab.value === 'critical'
                  ? 'bg-red-500 text-white'
                  : 'bg-gray-200 text-gray-700'
              "
            >
              {{ getCount(tab.value) }}
            </span>
          </button>
        </div>

        <div
          v-for="patient in filteredPatients"
          :key="patient?.id"
          @click="selectedPatient = patient"
          class="cursor-pointer rounded-2xl border bg-white px-6 py-5 shadow-sm transition-all duration-200 hover:shadow-md"
          :class="
            selectedPatient?.id === patient.id
              ? 'border-red-200 ring-1 ring-red-100'
              : 'border-gray-200'
          "
        >
          <div
            class="flex flex-col gap-4 xl:flex-row xl:items-center xl:justify-between"
          >
            <div class="flex items-center gap-4 min-w-0">
              <div
                class="flex h-14 w-14 shrink-0 items-center justify-center rounded-full bg-[#f5ecec] text-sm font-bold text-gray-700"
              >
                {{ patient?.initials }}
              </div>
              <div class="min-w-0">
                <div class="flex flex-wrap items-center gap-3">
                  <h3 class="truncate text-lg font-semibold text-gray-900">
                    {{ patient?.fullName }}
                  </h3>

                  <span
                    v-if="patient?.admission"
                    class="rounded-full px-3 py-1 text-xs font-semibold capitalize"
                    :class="statusClasses[patient.admission.status]"
                  >
                    {{ patient?.admission.status }}
                  </span>

                  <span
                    v-else
                    class="rounded-full bg-blue-100 px-3 py-1 text-xs font-semibold text-blue-700"
                  >
                    Registered
                  </span>
                </div>

                <p
                  v-if="patient?.admission"
                  class="mt-1 text-sm text-gray-500 truncate"
                >
                  {{ patient?.admission.diagnosis }}
                </p>

                <p v-else class="mt-1 text-sm text-gray-400">
                  Patient not admitted yet
                </p>
              </div>
            </div>
            <div
              class="grid grid-cols-2 gap-x-10 gap-y-4 xl:flex xl:items-center xl:gap-10"
            >
              <div>
                <p class="text-xs text-gray-400">Age</p>

                <p class="font-semibold text-gray-800">
                  {{ patient?.age }} · {{ patient?.gender }}
                </p>
              </div>
              <div v-if="patient?.admission">
                <p class="text-xs text-gray-400">Doctor</p>

                <p class="font-semibold text-gray-800">
                  {{ patient?.admission?.assignedDoctor }}
                </p>
              </div>
              <div v-if="patient?.admission">
                <p class="text-xs text-gray-400">Room</p>

                <p class="font-semibold text-gray-800">
                  {{ patient?.admission?.room }}
                </p>
              </div>
              <div v-else class="flex items-end">
                <button
                  class="rounded-xl bg-primary px-4 py-2 text-xs font-medium text-white transition hover:opacity-90"
                  @click.stop="openAdmissionModal(patient)"
                >
                  Admit Patient
                </button>
              </div>
              <div>
                <p class="text-xs text-gray-400">ID</p>

                <p class="font-semibold text-gray-800">
                  {{ formatPatientId(patient.id) }}
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="col-span-4">
        <div class="rounded-2xl border border-gray-200 bg-white p-6 shadow-sm">
          <div class="flex items-center gap-4">
            <div
              class="flex h-14 w-14 items-center justify-center rounded-full bg-[#f3eeee] text-lg font-bold"
            >
              {{ selectedPatient?.initials }}
            </div>

            <div>
              <h2 class="text-2xl font-bold">
                {{ selectedPatient?.fullName }}
              </h2>

              <p class="text-sm text-gray-500">
                {{ formatPatientId(selectedPatient?.id) }}
              </p>
            </div>
          </div>
          <div class="mt-8 grid grid-cols-2 gap-6">
            <div>
              <p class="text-sm text-gray-400">Age</p>

              <p class="font-semibold text-lg">
                {{ selectedPatient?.age }} years
              </p>
            </div>

            <div>
              <p class="text-sm text-gray-400">Gender</p>

              <p class="font-semibold text-lg">
                {{ selectedPatient?.gender }}
              </p>
            </div>

            <div>
              <p class="text-sm text-gray-400">Blood Type</p>

              <p class="font-semibold text-lg">
                {{ selectedPatient?.bloodType }}
              </p>
            </div>

            <div>
              <p class="text-sm text-gray-400">Status</p>
              <span
                v-if="selectedPatient?.admission"
                class="rounded-full px-3 py-1 text-xs font-medium capitalize"
                :class="statusClasses[selectedPatient?.admission?.status]"
              >
                {{ selectedPatient?.admission?.status }}
              </span>
              <span
                v-else
                class="rounded-full bg-blue-100 text-blue-700 px-3 py-1 text-xs font-medium"
              >
                registered
              </span>
            </div>
          </div>
          <template v-if="selectedPatient?.admission">
            <div class="mt-8">
              <p class="text-sm text-gray-400">Diagnosis</p>

              <p class="mt-1 text-lg font-semibold">
                {{ selectedPatient?.admission?.diagnosis }}
              </p>
            </div>
            <div class="mt-8">
              <p class="text-sm text-gray-400">Assigned Doctor</p>

              <p class="mt-1 font-semibold">
                {{ selectedPatient?.admission?.assignedDoctor }}
              </p>
            </div>
            <div class="mt-8">
              <p class="text-sm text-gray-400">Room</p>

              <p class="mt-1 font-semibold">
                {{ selectedPatient?.admission?.room }}
              </p>
            </div>
            <div class="mt-8">
              <p class="text-sm text-gray-400">Bed</p>

              <p class="mt-1 font-semibold">
                {{ selectedPatient?.admission?.bed }}
              </p>
            </div>
          </template>
          <template v-else>
            <div
              class="mt-8 rounded-2xl border border-dashed border-gray-300 p-6 text-center"
            >
              <div
                class="mx-auto mb-4 flex h-14 w-14 items-center justify-center rounded-full bg-blue-50"
              >
                <Icon name="lucide:bed-single" class="h-6 w-6 text-blue-600" />
              </div>

              <h3 class="text-lg font-semibold">Patient Not Admitted</h3>

              <p class="mt-2 text-sm text-gray-500">
                This patient is registered but has not yet been admitted.
              </p>

              <button
                @click="openAdmissionModal(patient)"
                class="mt-6 w-full rounded-xl bg-primary px-4 py-3 text-sm font-medium text-white hover:opacity-90"
              >
                Admit Patient
              </button>
            </div>
          </template>
          <div class="mt-8">
            <p class="text-sm text-gray-400">Allergies</p>

            <div class="mt-2">
              <span
                class="rounded-full bg-red-100 px-3 py-1 text-xs font-medium text-red-700"
              >
                {{ selectedPatient?.allergies }}
              </span>
            </div>
          </div>
          <div class="mt-8">
            <p class="text-sm text-gray-400">Emergency Contact</p>

            <p class="mt-1 font-semibold">
              {{ selectedPatient?.emergencyContact }}
            </p>

            <p class="text-sm text-gray-500">
              {{ selectedPatient?.phone }}
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
