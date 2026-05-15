<script setup>
import { ref, computed } from "vue";

const activeTab = ref("registered");
const patients = [
  {
    id: "P-1001",
    fullName: "Margaret Chen",
    initials: "MC",

    age: 64,
    gender: "Female",

    bloodType: "A+",
    allergies: "Penicillin",

    phone: "(555) 234-5678",
    emergencyContact: "David Chen",

    admission: {
      status: "active",
      urgency: "normal",
      diagnosis: "Type 2 Diabetes",
      assignedDoctor: "Dr. Sarah Mitchell",
      room: "204-A",
      bed: "B-12",
    },
  },

  {
    id: "P-1002",
    fullName: "James O'Sullivan",
    initials: "JO",

    age: 72,
    gender: "Male",

    bloodType: "B+",
    allergies: "None",

    phone: "(555) 888-2222",
    emergencyContact: "Emma Sullivan",

    admission: {
      status: "critical",
      urgency: "high",
      diagnosis: "Hypertension",
      assignedDoctor: "Dr. Robert Kim",
      room: "ICU-3",
      bed: "ICU-02",
    },
  },

  {
    id: "P-1003",
    fullName: "Aisha Rahman",
    initials: "AR",

    age: 45,
    gender: "Female",

    bloodType: "O+",
    allergies: "Latex",

    phone: "(555) 123-9876",
    emergencyContact: "Ahmed Rahman",

    admission: null,
  },

  {
    id: "P-1004",
    fullName: "Robert Nakamura",
    initials: "RN",

    age: 58,
    gender: "Male",

    bloodType: "AB+",
    allergies: "Aspirin",

    phone: "(555) 555-2222",
    emergencyContact: "Linda Nakamura",

    admission: {
      status: "discharged",
      urgency: "normal",
      diagnosis: "Chronic Heart Failure",
      assignedDoctor: "Dr. Sarah Mitchell",
      room: "ICU-7",
      bed: "ICU-08",
    },
  },
];
const selectedPatient = ref(patients[0]);
const tabs = [
  {
    label: "Registered",
    value: "registered",
  },

  {
    label: "Admitted",
    value: "admitted",
  },

  {
    label: "Critical",
    value: "critical",
  },

  {
    label: "Discharged",
    value: "discharged",
  },
];
const filteredPatients = computed(() => {
  if (activeTab.value === "registered") {
    return patients.filter((patient) => patient.admission === null);
  }
  if (activeTab.value === "admitted") {
    return patients.filter(
      (patient) =>
        patient.admission && patient.admission.status !== "discharged",
    );
  }

  if (activeTab.value === "critical") {
    return patients.filter(
      (patient) => patient.admission?.status === "critical",
    );
  }
  if (activeTab.value === "discharged") {
    return patients.filter(
      (patient) => patient.admission?.status === "discharged",
    );
  }

  return patients;
});
const getCount = (status) => {
  if (status === "registered") {
    return patients.filter((patient) => patient.admission === null).length;
  }

  if (status === "admitted") {
    return patients.filter(
      (patient) =>
        patient.admission && patient.admission.status !== "discharged",
    ).length;
  }

  if (status === "critical") {
    return patients.filter(
      (patient) => patient.admission?.status === "critical",
    ).length;
  }

  if (status === "discharged") {
    return patients.filter(
      (patient) => patient.admission?.status === "discharged",
    ).length;
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
const isAdmissionModalOpen = ref(false);

// const selectedPatient = ref(null);

const openAdmissionModal = (patient) => {
  selectedPatient.value = patient;

  isAdmissionModalOpen.value = true;
};
</script>

<template>
  <ModalsPatientAdmission
    v-if="selectedPatient"
    v-model="isAdmissionModalOpen"
    :patient="selectedPatient"
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
          :key="patient.id"
          @click="selectedPatient = patient"
          class="cursor-pointer rounded-2xl border bg-white p-5 shadow-sm transition-all duration-200 hover:shadow-md"
          :class="
            selectedPatient.id === patient.id
              ? 'border-red-200 ring-1 ring-red-100'
              : 'border-gray-200'
          "
        >
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-4">
              <div
                class="flex h-12 w-12 items-center justify-center rounded-full bg-[#f3eeee] text-sm font-bold"
              >
                {{ patient.initials }}
              </div>

              <div>
                <div class="flex items-center gap-2">
                  <h3 class="font-semibold text-xl">
                    {{ patient.fullName }}
                  </h3>
                  <span
                    v-if="patient.admission"
                    class="rounded-full px-3 py-1 text-xs font-medium capitalize"
                    :class="statusClasses[patient.admission.status]"
                  >
                    {{ patient.admission.status }}
                  </span>
                  <span
                    v-else
                    class="rounded-full bg-blue-100 text-blue-700 px-3 py-1 text-xs font-medium"
                  >
                    registered
                  </span>
                </div>
                <p v-if="patient.admission" class="text-sm text-gray-500">
                  {{ patient.admission.diagnosis }}
                </p>
                <p v-else class="text-sm text-gray-400">
                  Patient not admitted yet
                </p>
              </div>
            </div>
            <div class="flex items-center gap-10 text-sm">
              <div>
                <p class="text-gray-400 text-xs">Age</p>

                <p class="font-semibold">
                  {{ patient.age }} · {{ patient.gender }}
                </p>
              </div>
              <div v-if="patient.admission">
                <p class="text-gray-400 text-xs">Doctor</p>

                <p class="font-semibold">
                  {{ patient.admission.assignedDoctor }}
                </p>
              </div>
              <div v-if="patient.admission">
                <p class="text-gray-400 text-xs">Room</p>

                <p class="font-semibold">
                  {{ patient.admission.room }}
                </p>
              </div>
              <div v-else>
                <button
                  class="rounded-lg bg-primary px-4 py-2 text-xs font-medium text-white hover:opacity-90"
                  @click="openAdmissionModal(patient)"
                >
                  Admit Patient
                </button>
              </div>
              <div>
                <p class="text-gray-400 text-xs">ID</p>

                <p class="font-semibold">
                  {{ patient.id }}
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
              {{ selectedPatient.initials }}
            </div>

            <div>
              <h2 class="text-2xl font-bold">
                {{ selectedPatient.fullName }}
              </h2>

              <p class="text-sm text-gray-500">
                {{ selectedPatient.id }}
              </p>
            </div>
          </div>
          <div class="mt-8 grid grid-cols-2 gap-6">
            <div>
              <p class="text-sm text-gray-400">Age</p>

              <p class="font-semibold text-lg">
                {{ selectedPatient.age }} years
              </p>
            </div>

            <div>
              <p class="text-sm text-gray-400">Gender</p>

              <p class="font-semibold text-lg">
                {{ selectedPatient.gender }}
              </p>
            </div>

            <div>
              <p class="text-sm text-gray-400">Blood Type</p>

              <p class="font-semibold text-lg">
                {{ selectedPatient.bloodType }}
              </p>
            </div>

            <div>
              <p class="text-sm text-gray-400">Status</p>
              <span
                v-if="selectedPatient.admission"
                class="rounded-full px-3 py-1 text-xs font-medium capitalize"
                :class="statusClasses[selectedPatient.admission.status]"
              >
                {{ selectedPatient.admission.status }}
              </span>
              <span
                v-else
                class="rounded-full bg-blue-100 text-blue-700 px-3 py-1 text-xs font-medium"
              >
                registered
              </span>
            </div>
          </div>
          <template v-if="selectedPatient.admission">
            <div class="mt-8">
              <p class="text-sm text-gray-400">Diagnosis</p>

              <p class="mt-1 text-lg font-semibold">
                {{ selectedPatient.admission.diagnosis }}
              </p>
            </div>
            <div class="mt-8">
              <p class="text-sm text-gray-400">Assigned Doctor</p>

              <p class="mt-1 font-semibold">
                {{ selectedPatient.admission.assignedDoctor }}
              </p>
            </div>
            <div class="mt-8">
              <p class="text-sm text-gray-400">Room</p>

              <p class="mt-1 font-semibold">
                {{ selectedPatient.admission.room }}
              </p>
            </div>
            <div class="mt-8">
              <p class="text-sm text-gray-400">Bed</p>

              <p class="mt-1 font-semibold">
                {{ selectedPatient.admission.bed }}
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
                {{ selectedPatient.allergies }}
              </span>
            </div>
          </div>
          <div class="mt-8">
            <p class="text-sm text-gray-400">Emergency Contact</p>

            <p class="mt-1 font-semibold">
              {{ selectedPatient.emergencyContact }}
            </p>

            <p class="text-sm text-gray-500">
              {{ selectedPatient.phone }}
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
