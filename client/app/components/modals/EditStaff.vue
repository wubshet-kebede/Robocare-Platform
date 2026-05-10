<script setup>
const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false,
  },
  staff: {
    type: Object,
    required: true,
  },
});

const emits = defineEmits(["update:modelValue", "save"]);

const form = reactive({
  name: "",
  department: "",
  role: "",
  phone: "",
  shift: "",
  status: "",
});

watch(
  () => props.staff,
  (staff) => {
    if (!staff) return;

    form.name = staff.name || "";
    form.department = staff.department || "";
    form.role = staff.role || "";
    form.phone = staff.phone || "";
    form.shift = staff.shift || "";
    form.status = staff.status || "";
  },
  { immediate: true },
);

const loading = ref(false);

const submit = async () => {
  try {
    loading.value = true;

    const updatedStaff = {
      ...props.staff,
      ...form,
    };

    // later API call here

    emits("save", updatedStaff);

    emits("update:modelValue", false);
  } catch (error) {
    console.log(error);
  } finally {
    loading.value = false;
  }
};

const closeModal = () => {
  emits("update:modelValue", false);
};

const roles = [
  {
    id: "Doctors",
    name: "Doctors",
  },
  {
    id: "Nurses",
    name: "Nurses",
  },
  {
    id: "Administration",
    name: "Administration",
  },
];

const shifts = [
  {
    id: "Morning",
    name: "Morning",
  },
  {
    id: "Afternoon",
    name: "Afternoon",
  },
  {
    id: "Night",
    name: "Night",
  },
];

const statuses = [
  {
    id: "On Duty",
    name: "On Duty",
  },
  {
    id: "Off Duty",
    name: "Off Duty",
  },
  {
    id: "On Leave",
    name: "On Leave",
  },
  {
    id: "Suspended",
    name: "Suspended",
  },
];
</script>

<template>
  <ModalsModal
    :model-value="modelValue"
    @update:modelValue="closeModal"
    wrapperClass="sm:max-w-3xl"
  >
    <template #Heading>
      <div class="flex items-start gap-4">
        <div
          class="flex h-10 w-10 items-center justify-center rounded-xl bg-primary/10"
        >
          <Icon name="lucide:user-pen" class="h-5 w-5 text-primary" />
        </div>

        <div class="mb-5">
          <h2 class="text-lg font-semibold text-gray-900 dark:text-white">
            Edit Staff Member
          </h2>

          <p class="text-sm text-gray-500">
            Update staff information and permissions.
          </p>
        </div>
      </div>
    </template>

    <template #content>
      <form class="mt-6" @submit.prevent="submit">
        <div class="grid grid-cols-1 gap-6 md:grid-cols-2">
          <UiBaseInput
            v-model="form.name"
            name="name"
            placeholder="Enter full name"
            leadingIcon="lucide:user"
            rules="required|min:3"
          >
            <template #label>
              <label class="text-sm font-medium text-gray-900">
                Full Name
              </label>
            </template>
          </UiBaseInput>
          <UiBaseInput
            v-model="form.phone"
            name="phone"
            placeholder="Enter phone number"
            leadingIcon="lucide:phone"
            rules="required"
          >
            <template #label>
              <label class="text-sm font-medium text-gray-900">
                Phone Number
              </label>
            </template>
          </UiBaseInput>
          <UiBaseInput
            v-model="form.department"
            name="department"
            placeholder="Enter department"
            leadingIcon="lucide:building-2"
            rules="required"
          >
            <template #label>
              <label class="text-sm font-medium text-gray-900">
                Department
              </label>
            </template>
          </UiBaseInput>
          <div class="space-y-1">
            <UiListSelect
              v-model="form.role"
              name="role"
              :items="roles"
              placeholder="Select role"
            >
              <template #label>
                <label class="text-sm font-medium text-gray-900"> Role </label>
              </template>
            </UiListSelect>
          </div>
          <div class="space-y-1">
            <UiListSelect
              v-model="form.shift"
              name="shift"
              :items="shifts"
              placeholder="Select shift"
            >
              <template #label>
                <label class="text-sm font-medium text-gray-900"> Shift </label>
              </template>
            </UiListSelect>
          </div>
          <div class="space-y-1">
            <UiListSelect
              v-model="form.status"
              name="status"
              :items="statuses"
              placeholder="Select status"
            >
              <template #label>
                <label class="text-sm font-medium text-gray-900">
                  Status
                </label>
              </template>
            </UiListSelect>
          </div>
        </div>
        <div
          class="mt-10 flex flex-col-reverse gap-3 border-t border-gray-100 pt-6 sm:flex-row sm:justify-end"
        >
          <button
            type="button"
            @click="closeModal"
            class="inline-flex items-center justify-center rounded-xl border border-gray-300 bg-white px-5 py-3 text-sm font-medium text-gray-700 transition hover:bg-gray-50"
          >
            Cancel
          </button>

          <button
            type="submit"
            :disabled="loading"
            class="inline-flex items-center justify-center gap-2 rounded-xl bg-primary px-5 py-3 text-sm font-medium text-white transition hover:bg-primary/90 disabled:cursor-not-allowed disabled:opacity-50"
          >
            <Icon v-if="loading" name="svg-spinners:180-ring" class="h-4 w-4" />

            <Icon v-else name="lucide:save" class="h-4 w-4" />

            {{ loading ? "Saving..." : "Save Changes" }}
          </button>
        </div>
      </form>
    </template>
  </ModalsModal>
</template>
