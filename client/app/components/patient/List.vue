<script setup>
import { ref, computed } from "vue";

const activeTab = ref("all");

const patients = [
  {
    id: "P-1001",
    name: "Margaret Chen",
    initials: "MC",
    status: "active",
    condition: "Type 2 Diabetes",
    age: 64,
    gender: "Female",
    doctor: "Dr. Sarah Mitchell",
    room: "204-A",
    bloodType: "A+",
    allergies: "Penicillin",
    emergencyContact: "David Chen",
    phone: "(555) 234-5678",
  },

  {
    id: "P-1002",
    name: "James O'Sullivan",
    initials: "JO",
    status: "critical",
    condition: "Hypertension",
    age: 72,
    gender: "Male",
    doctor: "Dr. Robert Kim",
    room: "ICU-3",
    bloodType: "B+",
    allergies: "None",
    emergencyContact: "Emma Sullivan",
    phone: "(555) 888-2222",
  },

  {
    id: "P-1003",
    name: "Aisha Rahman",
    initials: "AR",
    status: "recovering",
    condition: "Post-op Recovery (Knee)",
    age: 45,
    gender: "Female",
    doctor: "Dr. Michael Torres",
    room: "312-B",
    bloodType: "O+",
    allergies: "Latex",
    emergencyContact: "Ahmed Rahman",
    phone: "(555) 123-9876",
  },

  {
    id: "P-1004",
    name: "Robert Nakamura",
    initials: "RN",
    status: "critical",
    condition: "Chronic Heart Failure",
    age: 58,
    gender: "Male",
    doctor: "Dr. Sarah Mitchell",
    room: "ICU-7",
    bloodType: "AB+",
    allergies: "Aspirin",
    emergencyContact: "Linda Nakamura",
    phone: "(555) 555-2222",
  },

  {
    id: "P-1005",
    name: "Elena Vasquez",
    initials: "EV",
    status: "active",
    condition: "Pneumonia",
    age: 34,
    gender: "Female",
    doctor: "Dr. Angela Park",
    room: "118-A",
    bloodType: "A-",
    allergies: "None",
    emergencyContact: "Carlos Vasquez",
    phone: "(555) 777-9999",
  },

  {
    id: "P-1006",
    name: "Thomas Bergstrom",
    initials: "TB",
    status: "active",
    condition: "Atrial Fibrillation",
    age: 81,
    gender: "Male",
    doctor: "Dr. Robert Kim",
    room: "205-C",
    bloodType: "O-",
    allergies: "Sulfa Drugs",
    emergencyContact: "Mia Bergstrom",
    phone: "(555) 444-1212",
  },

  {
    id: "P-1007",
    name: "Priya Patel",
    initials: "PP",
    status: "discharged",
    condition: "Appendectomy Recovery",
    age: 29,
    gender: "Female",
    doctor: "Dr. Michael Torres",
    room: "—",
    bloodType: "B-",
    allergies: "None",
    emergencyContact: "Raj Patel",
    phone: "(555) 222-3333",
  },
];

const selectedPatient = ref(patients[0]);

const tabs = [
  {
    label: "All Patients",
    value: "all",
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
  if (activeTab.value === "all") {
    return patients;
  }

  return patients.filter((patient) => patient.status === activeTab.value);
});

const getCount = (status) => {
  if (status === "all") {
    return patients.length;
  }

  return patients.filter((patient) => patient.status === status).length;
};

const statusClasses = {
  active: "bg-green-100 text-green-700",

  critical: "bg-red-100 text-red-700",

  recovering: "bg-yellow-100 text-yellow-700",

  discharged: "bg-gray-100 text-gray-700",
};
</script>

<template>
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
            <!-- LEFT -->
            <div class="flex items-center gap-4">
              <!-- AVATAR -->
              <div
                class="flex h-12 w-12 items-center justify-center rounded-full bg-[#f3eeee] text-sm font-bold"
              >
                {{ patient.initials }}
              </div>
              <div>
                <div class="flex items-center gap-2">
                  <h3 class="font-semibold text-xl">
                    {{ patient.name }}
                  </h3>

                  <span
                    class="rounded-full px-3 py-1 text-xs font-medium capitalize"
                    :class="statusClasses[patient.status]"
                  >
                    {{ patient.status }}
                  </span>
                </div>

                <p class="text-sm text-gray-500">
                  {{ patient.condition }}
                </p>
              </div>
            </div>

            <!-- RIGHT -->
            <div class="flex items-center gap-10 text-sm">
              <div>
                <p class="text-gray-400 text-xs">Age</p>

                <p class="font-semibold">
                  {{ patient.age }} · {{ patient.gender }}
                </p>
              </div>

              <div>
                <p class="text-gray-400 text-xs">Doctor</p>

                <p class="font-semibold">
                  {{ patient.doctor }}
                </p>
              </div>

              <div>
                <p class="text-gray-400 text-xs">Room</p>

                <p class="font-semibold">
                  {{ patient.room }}
                </p>
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
                {{ selectedPatient.name }}
              </h2>

              <p class="text-sm text-gray-500">
                {{ selectedPatient.id }}
                · Room {{ selectedPatient.room }}
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
                class="rounded-full px-3 py-1 text-xs font-medium capitalize"
                :class="statusClasses[selectedPatient.status]"
              >
                {{ selectedPatient.status }}
              </span>
            </div>
          </div>

          <div class="mt-8">
            <p class="text-sm text-gray-400">Diagnosis</p>

            <p class="mt-1 text-lg font-semibold">
              {{ selectedPatient.condition }}
            </p>
          </div>
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
