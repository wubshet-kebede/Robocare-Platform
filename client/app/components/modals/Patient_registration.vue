<script setup>
const isOpen = ref(true);

const loadingSubmit = ref(false);

const values = ref({
  firstName: "",
  lastName: "",
  email: "",
  phone: "",
  gender: "",
  age: "",
  bloodType: "",
  status: "",
  doctor: "",
  room: "",
  diagnosis: "",
  allergies: "",
  emergencyContact: "",
  emergencyPhone: "",
});

const genderOptions = [
  {
    label: "Male",
    value: "male",
  },

  {
    label: "Female",
    value: "female",
  },
];

const bloodTypeOptions = [
  {
    label: "A+",
    value: "A+",
  },

  {
    label: "A-",
    value: "A-",
  },

  {
    label: "B+",
    value: "B+",
  },

  {
    label: "B-",
    value: "B-",
  },

  {
    label: "AB+",
    value: "AB+",
  },

  {
    label: "AB-",
    value: "AB-",
  },

  {
    label: "O+",
    value: "O+",
  },

  {
    label: "O-",
    value: "O-",
  },
];

const statusOptions = [
  {
    label: "Active",
    value: "active",
  },

  {
    label: "Critical",
    value: "critical",
  },

  {
    label: "Recovering",
    value: "recovering",
  },

  {
    label: "Discharged",
    value: "discharged",
  },
];

const doctorOptions = [
  {
    label: "Dr. Sarah Mitchell",
    value: "Dr. Sarah Mitchell",
  },

  {
    label: "Dr. Robert Kim",
    value: "Dr. Robert Kim",
  },

  {
    label: "Dr. Michael Torres",
    value: "Dr. Michael Torres",
  },

  {
    label: "Dr. Angela Park",
    value: "Dr. Angela Park",
  },
];

const submit = async () => {
  loadingSubmit.value = true;

  try {
    console.log(values.value);

    setTimeout(() => {
      loadingSubmit.value = false;
      isOpen.value = false;
    }, 1500);
  } catch (error) {
    loadingSubmit.value = false;
  }
};
</script>

<template>
  <ModalsModal
    v-model="isOpen"
    title="Patient Registration"
    wrapperClass="max-w-4xl"
  >
    <template #content>
      <div class="p-8 space-y-8">
        <!-- PERSONAL INFO -->
        <div>
          <h2 class="text-xl font-semibold mb-6">Personal Information</h2>

          <div class="grid grid-cols-2 gap-6">
            <!-- FIRST NAME -->
            <UiBaseInput
              v-model="values.firstName"
              name="firstName"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  First Name
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </UiBaseInput>

            <!-- LAST NAME -->
            <UiBaseInput
              v-model="values.lastName"
              name="lastName"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  Last Name
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </UiBaseInput>

            <!-- EMAIL -->
            <UiBaseInput
              v-model="values.email"
              name="email"
              rules="required|email"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">Email Address</h1>
              </template>
            </UiBaseInput>

            <!-- PHONE -->
            <UiBaseInput v-model="values.phone" name="phone" rules="required">
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  Phone Number
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </UiBaseInput>

            <!-- AGE -->
            <UiBaseInput
              v-model="values.age"
              name="age"
              rules="required"
              type="number"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  Age
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </UiBaseInput>

            <!-- GENDER -->
            <UiListSelect
              v-model="values.gender"
              :items="genderOptions"
              name="gender"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  Gender
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </UiListSelect>
          </div>
        </div>

        <!-- MEDICAL INFO -->
        <div>
          <h2 class="text-xl font-semibold mb-6">Medical Information</h2>

          <div class="grid grid-cols-2 gap-6">
            <!-- BLOOD TYPE -->
            <UiListSelect
              v-model="values.bloodType"
              :items="bloodTypeOptions"
              name="bloodType"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  Blood Type
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </UiListSelect>

            <!-- STATUS -->
            <UiListSelect
              v-model="values.status"
              :items="statusOptions"
              name="status"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  Patient Status
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </UiListSelect>

            <!-- DOCTOR -->
            <UiListSelect
              v-model="values.doctor"
              :items="doctorOptions"
              name="doctor"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  Assigned Doctor
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </UiListSelect>

            <!-- ROOM -->
            <UiBaseInput v-model="values.room" name="room" rules="required">
              <template #label>
                <h1 class="text-md font-medium mb-2">Room Number</h1>
              </template>
            </UiBaseInput>
          </div>
        </div>

        <!-- DIAGNOSIS -->
        <div>
          <h2 class="text-xl font-semibold mb-6">Diagnosis & Notes</h2>

          <div class="space-y-6">
            <!-- DIAGNOSIS -->
            <UiBaseInput
              v-model="values.diagnosis"
              name="diagnosis"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  Diagnosis
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </UiBaseInput>

            <!-- ALLERGIES -->
            <UiBaseInput v-model="values.allergies" name="allergies">
              <template #label>
                <h1 class="text-md font-medium mb-2">Allergies</h1>
              </template>
            </UiBaseInput>
          </div>
        </div>

        <!-- EMERGENCY CONTACT -->
        <div>
          <h2 class="text-xl font-semibold mb-6">Emergency Contact</h2>

          <div class="grid grid-cols-2 gap-6">
            <!-- CONTACT NAME -->
            <UiBaseInput
              v-model="values.emergencyContact"
              name="emergencyContact"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  Contact Name
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </UiBaseInput>

            <!-- CONTACT PHONE -->
            <UiBaseInput
              v-model="values.emergencyPhone"
              name="emergencyPhone"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  Contact Phone
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </UiBaseInput>
          </div>
        </div>

        <!-- ACTIONS -->
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
            {{ loadingSubmit ? "Saving Patient..." : "Register Patient" }}
          </button>
        </div>
      </div>
    </template>
  </ModalsModal>
</template>
