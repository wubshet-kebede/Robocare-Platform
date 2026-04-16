<script setup>
defineProps({
  staff: {
    type: Object,
    required: true,
    // Expecting: { name, role, department, phone, email, rating, exp, shift, status }
  },
});

const getStatusClasses = (status) => {
  const statusMap = {
    "On Duty": "bg-emerald-100 text-emerald-700",
    "Off Duty": "bg-gray-100 text-gray-600",
    "On Leave": "bg-amber-100 text-amber-700",
  };
  return statusMap[status] || "bg-gray-100 text-gray-600";
};
</script>

<template>
  <div
    class="group relative rounded-2xl border border-gray-100 bg-white p-6 shadow-sm transition-all duration-200 hover:shadow-md"
  >
    <div class="flex items-start justify-between">
      <div class="flex items-center gap-4">
        <div
          class="flex h-12 w-12 items-center justify-center rounded-full bg-indigo-50 text-indigo-700 font-bold"
        >
          {{
            staff.name
              .split(" ")
              .map((n) => n[0])
              .join("")
          }}
        </div>
        <div>
          <h3 class="font-bold text-gray-900">{{ staff.name }}</h3>
          <p class="text-sm text-gray-500">{{ staff.role }}</p>
        </div>
      </div>
      <span
        :class="[
          'px-2.5 py-0.5 rounded-full text-xs font-medium',
          getStatusClasses(staff.status),
        ]"
      >
        {{ staff.status }}
      </span>
    </div>

    <div class="mt-6 space-y-3">
      <div class="flex items-center text-sm text-gray-600">
        <Icon name="lucide:shield" class="mr-2 h-4 w-4" />
        {{ staff.department }}
      </div>
      <div class="flex items-center text-sm text-gray-600">
        <Icon name="lucide:phone" class="mr-2 h-4 w-4" /> {{ staff.phone }}
      </div>
      <div class="flex items-center text-sm text-gray-600">
        <Icon name="lucide:mail" class="mr-2 h-4 w-4" /> {{ staff.email }}
      </div>
    </div>

    <div
      class="mt-6 flex items-center justify-between border-t pt-4 text-xs text-gray-500"
    >
      <div class="flex items-center gap-1 text-amber-500">
        <Icon name="lucide:star" class="h-4 w-4" />
        <span>{{ staff.rating }}</span>
      </div>
      <span>{{ staff.exp }} yrs exp.</span>
      <div class="flex items-center gap-1">
        <Icon name="lucide:clock" class="h-4 w-4" />
        <span>{{ staff.shift }}</span>
      </div>
    </div>
  </div>
</template>
