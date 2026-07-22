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

const { registerRoom } = useRoomService();
const { getDepartments } = useDepartmentService();
const departments = ref([]);
const loadingDepartments = ref(false);
const fetchDepartments = async () => {
  try {
    loadingDepartments.value = true;

    const res = await getDepartments();

    console.log("raw:", res);
    console.log("type:", typeof res);

    const data = typeof res === "string" ? JSON.parse(res) : res;

    departments.value = data;

    console.log("Normalized:", departments.value);
  } catch (err) {
    console.error("Failed to fetch departments:", err);
    departments.value = [];
  } finally {
    loadingDepartments.value = false;
  }
};

onMounted(() => {
  fetchDepartments();
});
const departmentOptions = computed(() => {
  if (!departments.value.length) return [];

  return departments.value.map((d) => ({
    id: d.id,
    name: d.name,
  }));
});
console.log("departmentOptions:", departmentOptions.value);
const { handleSubmit, resetForm, values } = useForm({
  initialValues: {
    roomNumber: "",
    locationName: "",
    floor: 1,
    capacity: 1,
    status: "available",
    departmentId: "",
    x: 0,
    y: 0,
    yaw: 0,
  },
});

const submit = handleSubmit(async (formValues) => {
  try {
    loadingSubmit.value = true;

    const payload = {
      room_number: formValues.roomNumber,
      location_name: formValues.locationName,

      floor: Number(formValues.floor),
      capacity: Number(formValues.capacity),
      department_id: formValues.departmentId,

      status: formValues.status,

      x: Number(formValues.x),
      y: Number(formValues.y),
      yaw: Number(formValues.yaw),
    };

    console.log("Room Payload:", payload);

    const response = await registerRoom(payload);

    console.log("Room Registered:", response);

    resetForm();

    isOpen.value = false;

    emits("success", response);
  } catch (error) {
    console.log("Room Registration Error:", error);
  } finally {
    loadingSubmit.value = false;
  }
});

const roomStatusOptions = [
  {
    id: "available",
    name: "Available",
  },

  {
    id: "occupied",
    name: "Occupied",
  },

  {
    id: "cleaning",
    name: "Cleaning",
  },
];
</script>
<template>
  <ModalsModal
    v-model="isOpen"
    title="Room Management"
    wrapperClass="w-[95vw] max-w-4xl"
  >
    <template #content>
      <div class="p-4 sm:p-6 lg:p-8 space-y-6 sm:space-y-8">
        <div>
          <h2 class="text-lg sm:text-xl font-semibold mb-4 sm:mb-6">
            Room Information
          </h2>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <UiBaseInput
              v-model="values.roomNumber"
              name="roomNumber"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  Room Number
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </UiBaseInput>

            <UiBaseInput
              v-model="values.floor"
              name="floor"
              type="number"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  Floor
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </UiBaseInput>

            <UiBaseInput
              v-model="values.capacity"
              name="capacity"
              type="number"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  Capacity
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </UiBaseInput>

            <UiListSelect
              v-model="values.status"
              :items="roomStatusOptions"
              name="status"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">Room Status</h1>
              </template>
            </UiListSelect>

            <UiListSelect
              v-model="values.departmentId"
              :items="departmentOptions"
              name="departmentId"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  Department
                  <!-- <span class="text-red-500">*</span> -->
                </h1>
              </template>
            </UiListSelect>

            <UiBaseInput v-model="values.locationName" name="locationName">
              <template #label>
                <h1 class="text-md font-medium mb-2">Room Name</h1>
              </template>
            </UiBaseInput>
          </div>
        </div>

        <div>
          <h2 class="text-xl font-semibold mb-6">
            Robot Navigation Coordinates
          </h2>
          <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
            <UiBaseInput
              v-model="values.x"
              name="x"
              type="number"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  X Coordinate
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </UiBaseInput>

            <UiBaseInput
              v-model="values.y"
              name="y"
              type="number"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  Y Coordinate
                  <span class="text-red-500">*</span>
                </h1>
              </template>
            </UiBaseInput>

            <UiBaseInput
              v-model="values.yaw"
              name="yaw"
              type="number"
              rules="required"
            >
              <template #label>
                <h1 class="text-md font-medium mb-2">
                  Yaw / Rotation
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
            {{ loadingSubmit ? "Saving Room..." : "Create Room" }}
          </button>
        </div>
      </div>
    </template>
  </ModalsModal>
</template>
