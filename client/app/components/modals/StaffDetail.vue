<script setup>
const props = defineProps({
  modelValue: Boolean,
  staff: Object,
});

const emits = defineEmits(["update:modelValue"]);
</script>

<template>
  <ModalsModal
    :model-value="modelValue"
    @update:modelValue="emits('update:modelValue', $event)"
    wrapperClass="sm:max-w-2xl"
  >
    <template #Heading> Staff Details </template>

    <template #content>
      <div class="mt-6">
        <div class="flex items-center gap-4">
          <div
            class="h-16 w-16 rounded-full bg-primary/10 flex items-center justify-center"
          >
            <Icon name="mdi:user" class="h-8 w-8" />
          </div>

          <div>
            <h2 class="text-xl font-semibold">
              {{ staff.name }}
            </h2>

            <p class="text-sm text-gray-500">
              {{ staff.role }}
            </p>
          </div>
        </div>

        <div class="grid grid-cols-2 gap-4 mt-8">
          <div>
            <p class="text-sm text-gray-500">Department</p>
            <p class="font-medium">{{ staff.department }}</p>
          </div>

          <div>
            <p class="text-sm text-gray-500">Shift</p>
            <p class="font-medium">{{ staff.shift }}</p>
          </div>

          <div>
            <p class="text-sm text-gray-500">Email</p>
            <p class="font-medium">{{ staff.email }}</p>
          </div>

          <div>
            <p class="text-sm text-gray-500">Phone</p>
            <p class="font-medium">{{ staff.phone }}</p>
          </div>

          <div>
            <p class="text-sm text-gray-500">Status</p>
            <p class="font-medium">{{ staff.status }}</p>
          </div>

          <div>
            <p class="text-sm text-gray-500">Experience</p>
            <p class="font-medium">{{ staff.exp }} years</p>
          </div>
        </div>
      </div>
    </template>
  </ModalsModal>
</template>
