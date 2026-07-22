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
  get: () => props.modelValue,
  set: (value) => emits("update:modelValue", value),
});

const loadingSubmit = ref(false);

const { createDepartment } = useDepartmentService();
const { getAssignableStaff } = useStaffService();

const staffOptions = ref([]);
const loadingStaff = ref(false);

const fetchAssignableStaff = async () => {
  try {
    loadingStaff.value = true;

    const res = await getAssignableStaff();

    console.log("Raw Assignable Staff:", res);

    const data = typeof res === "string" ? JSON.parse(res) : res;

    staffOptions.value = [
      ...data.map((staff) => ({
        id: staff.id,
        name: staff.full_name,
      })),
    ];

    console.log("Staff Options:", staffOptions.value);
  } catch (err) {
    console.error("Failed to fetch assignable staff:", err);

    staffOptions.value = [
      {
        id: "",
        name: "No Department Head",
      },
    ];
  } finally {
    loadingStaff.value = false;
  }
};

onMounted(() => {
  fetchAssignableStaff();
});

const departmentTypeOptions = [
  {
    id: "clinical",
    name: "Clinical",
  },
  {
    id: "support",
    name: "Support",
  },
  {
    id: "administrative",
    name: "Administrative",
  },
  {
    id: "laboratory",
    name: "Laboratory",
  },
  {
    id: "emergency",
    name: "Emergency",
  },
  {
    id: "pharmacy",
    name: "Pharmacy",
  },
  {
    id: "custom",
    name: "Custom",
  },
];
const { handleSubmit, resetForm, values } = useForm({
  initialValues: {
    name: "",
    code: "",
    description: "",
    type: "clinical",
    maxBeds: 0,
    floor: 1,
    headId: "",
  },
});

const submit = handleSubmit(async (formValues) => {
  try {
    loadingSubmit.value = true;

    const payload = {
      name: formValues.name,
      code: formValues.code.toUpperCase(),
      description: formValues.description,
      type: formValues.type,
      max_beds: Number(formValues.maxBeds),
      floor: Number(formValues.floor),

      head_id: formValues.headId || null,
    };

    console.log("Department Payload:", payload);

    const response = await createDepartment(payload);

    console.log("Department Created:", response);

    resetForm();

    isOpen.value = false;

    emits("success", response);
  } catch (error) {
    console.error("Department Creation Error:", error);
  } finally {
    loadingSubmit.value = false;
  }
});
</script>
<template>
  <ModalsModal
    v-model="isOpen"
    title="Department Management"
    wrapperClass="w-[95vw] max-w-3xl"
  >
    <template #content>
      <div class="p-4 sm:p-6 lg:p-8 space-y-6 sm:space-y-8">
        <div>
          <h2 class="text-lg sm:text-xl font-semibold mb-4 sm:mb-6">
            Department Information
          </h2>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <!-- Department Name -->
            <UiBaseInput v-model="values.name" name="name" rules="required">
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  Department Name
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </UiBaseInput>

            <!-- Code -->
            <UiBaseInput v-model="values.code" name="code" rules="required">
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  Department Code
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </UiBaseInput>

            <!-- Type -->
            <UiListSelect
              v-model="values.type"
              :items="departmentTypeOptions"
              name="type"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">Department Type</h1>
              </template>
            </UiListSelect>

            <!-- Max Beds -->
            <UiBaseInput v-model="values.maxBeds" name="maxBeds" type="number">
              <template #label>
                <h1 class="text-md font-medium mb-2">Max Beds</h1>
              </template>
            </UiBaseInput>

            <!-- Floor -->
            <UiBaseInput v-model="values.floor" name="floor" type="number">
              <template #label>
                <h1 class="text-md font-medium mb-2">Floor</h1>
              </template>
            </UiBaseInput>

            <UiListSelect
              v-model="values.headId"
              :items="staffOptions"
              name="headId"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">Department Head</h1>
              </template>
            </UiListSelect>
          </div>

          <!-- Description -->
          <div class="mt-6">
            <UiBaseTextarea v-model="values.description" name="description">
              <template #label>
                <h1 class="text-md font-medium mb-2">Description</h1>
              </template>
            </UiBaseTextarea>
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
            {{ loadingSubmit ? "Creating Department..." : "Create Department" }}
          </button>
        </div>
      </div>
    </template>
  </ModalsModal>
</template>
