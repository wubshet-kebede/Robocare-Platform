<script setup>
import { useForm } from "vee-validate";

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false,
  },
});

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

const { registerPatient } = usePatientService();

const { handleSubmit, resetForm, values } = useForm({
  initialValues: {
    fullName: "",
    dateOfBirth: "",
    gender: "",
    email: "",
    phone: "",
    address: "",
    bloodType: "",
    allergies: "",
    chronicConditions: "",
    emergencyContactName: "",
    emergencyContactPhone: "",
  },
});
const submit = handleSubmit(async (formValues) => {
  try {
    loadingSubmit.value = true;

    const payload = {
      full_name: formValues.fullName,
      date_of_birth: new Date(formValues.dateOfBirth).toISOString(),
      gender: formValues.gender,

      email: formValues.email,
      phone: formValues.phone,
      address: formValues.address,

      blood_type: formValues.bloodType,
      allergies: formValues.allergies,
      chronic_conditions: formValues.chronicConditions,

      emergency_contact_name: formValues.emergencyContactName,
      emergency_contact_phone: formValues.emergencyContactPhone,
    };
    console.log(payload);
    const response = await registerPatient(payload);

    console.log("Patient Registered:", response);

    resetForm();

    isOpen.value = false;

    emits("success", response);
  } catch (error) {
    console.log("Patient Registration Error:", error);
  } finally {
    loadingSubmit.value = false;
  }
});

const genderOptions = [
  {
    id: "male",
    name: "Male",
  },

  {
    id: "female",
    name: "Female",
  },
];

const bloodTypeOptions = [
  {
    id: "A+",
    name: "A+",
  },

  {
    id: "A-",
    name: "A-",
  },

  {
    id: "B+",
    name: "B+",
  },

  {
    id: "B-",
    name: "B-",
  },

  {
    id: "AB+",
    name: "AB+",
  },

  {
    id: "AB-",
    name: "AB-",
  },

  {
    id: "O+",
    name: "O+",
  },

  {
    label: "O-",
    value: "O-",
  },
];
</script>

<template>
  <ModalsModal
    v-model="isOpen"
    title="Patient Registration"
    wrapperClass="w-full max-w-4xl"
  >
    <template #content>
      <div class="p-4 sm:p-6 lg:p-8 space-y-6 lg:space-y-8">
        <div>
          <h2 class="text-lg sm:text-xl font-semibold mb-4 sm:mb-6">
            Personal Information
          </h2>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-4 sm:gap-6">
            <UiBaseInput
              v-model="values.fullName"
              name="fullName"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  Full Name
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </UiBaseInput>
            <UiBaseInput
              v-model="values.dateOfBirth"
              name="dateOfBirth"
              type="date"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  Date of Birth
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </UiBaseInput>
            <UiListSelect
              v-model="values.gender"
              :items="genderOptions"
              name="gender"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">Gender</h1>
              </template>
            </UiListSelect>
            <UiListSelect
              v-model="values.bloodType"
              :items="bloodTypeOptions"
              name="bloodType"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">Blood Type</h1>
              </template>
            </UiListSelect>
          </div>
        </div>
        <div>
          <h2 class="text-xl font-semibold mb-6">Contact Information</h2>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-4 sm:gap-6">
            <UiBaseInput
              v-model="values.email"
              name="email"
              rules="required|email"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  Email Address
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </UiBaseInput>
            <UiBaseInput v-model="values.phone" name="phone" rules="required">
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  Phone Number
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </UiBaseInput>
            <div class="md:col-span-2">
              <UiBaseInput
                v-model="values.address"
                name="address"
                rules="required"
              >
                <template #label>
                  <h1 class="text-md font-medium mb-2">
                    Address
                    <span class="text-red-500">*</span>
                  </h1>
                </template>
              </UiBaseInput>
            </div>
          </div>
        </div>
        <div>
          <h2 class="text-xl font-semibold mb-6">Medical Information</h2>

          <div class="space-y-4 sm:space-y-6">
            <UiBaseInput v-model="values.allergies" name="allergies">
              <template #label>
                <h1 class="text-md font-medium mb-2">Allergies</h1>
              </template>
            </UiBaseInput>
            <UiBaseInput
              v-model="values.chronicConditions"
              name="chronicConditions"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">Chronic Conditions</h1>
              </template>
            </UiBaseInput>
          </div>
        </div>
        <div>
          <h2 class="text-xl font-semibold mb-6">Emergency Contact</h2>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-4 sm:gap-6">
            <UiBaseInput
              v-model="values.emergencyContactName"
              name="emergencyContactName"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  Contact Name
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </UiBaseInput>
            <UiBaseInput
              v-model="values.emergencyContactPhone"
              name="emergencyContactPhone"
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
            {{ loadingSubmit ? "Registering Patient..." : "Register Patient" }}
          </button>
        </div>
      </div>
    </template>
  </ModalsModal>
</template>
