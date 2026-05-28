<script setup>
import { reactive, ref } from "vue";

const isOpen = ref(false);
const loadingSubmit = ref(false);

const values = reactive({
  roomNumber: "",
  floor: null,
  capacity: 1,
  status: "available",
  departmentId: "",
  locationName: "",
  x: 0,
  y: 0,
  yaw: 0,
});

const roomStatusOptions = [
  {
    id: "Available",
    name: "available",
  },
  {
    id: "Occupied",
    name: "occupied",
  },
  {
    id: "Cleaning",
    name: "Cleaning",
  },
];

const departmentOptions = [
  {
    id: "Emergency",
    name: "dept-1",
  },
  {
    id: "ICU",
    name: "dept-2",
  },
  {
    id: "Surgery",
    name: "dept-3",
  },
];

const resetForm = () => {
  values.roomNumber = "";
  values.floor = null;
  values.capacity = 1;
  values.status = "available";
  values.departmentId = "";
  values.locationName = "";
  values.x = 0;
  values.y = 0;
  values.yaw = 0;
};

const submit = async () => {
  try {
    loadingSubmit.value = true;

    const payload = {
      room_number: values.roomNumber,
      floor: Number(values.floor),
      capacity: Number(values.capacity),
      status: values.status,
      department_id: values.departmentId,
      location_name: values.locationName,
      x: Number(values.x),
      y: Number(values.y),
      yaw: Number(values.yaw),
    };

    console.log("Submitting Room:", payload);

    /**
     * Example API Call
     */
    // await $fetch("/api/rooms", {
    //   method: "POST",
    //   body: payload,
    // });

    resetForm();
    isOpen.value = false;
  } catch (error) {
    console.error("Failed to create room:", error);
  } finally {
    loadingSubmit.value = false;
  }
};
</script>
<template>
  <ModalsModal
    v-model="isOpen"
    title="Room Management"
    wrapperClass="max-w-3xl"
  >
    <template #content>
      <div class="p-8 space-y-8">
        <div>
          <h2 class="text-xl font-semibold mb-6">Room Information</h2>

          <div class="grid grid-cols-2 gap-6">
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
                <h1 class="text-md font-medium mb-2">Department</h1>
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

          <div class="grid grid-cols-3 gap-6">
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
