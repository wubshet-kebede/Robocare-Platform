<script setup>
definePageMeta({
  layout: "dashboard",
});
const metrics = [
  {
    title: "Total Patients",
    value: "1,247",
    trend: "+3.2%",
    trendText: "from last month",
    trendSuccess: true,
    icon: "lucide:users",
    colorTheme: "bg-rose-50 text-rose-500",
  },
  {
    title: "New This Month",
    value: "86",
    trend: "+12.5%",
    trendText: "vs last month",
    trendSuccess: true,
    icon: "lucide:user-plus",
    colorTheme: "bg-emerald-50 text-emerald-500",
  },
  {
    title: "Critical Cases",
    value: "12",
    trend: "-2",
    trendText: "since yesterday",
    trendSuccess: true,
    icon: "lucide:alert-triangle",
    colorTheme: "bg-rose-50 text-rose-500",
  },
  {
    title: "Avg Stay",
    value: "4.2 days",
    trend: "-0.5 days",
    trendText: "vs avg",
    trendSuccess: true,
    icon: "lucide:clock",
    colorTheme: "bg-violet-50 text-violet-500",
  },
];
</script>
<template>
  <div
    class="mb-6 flex flex-col gap-4 sm:flex-row items-center justify-between"
  >
    <div>
      <h1 class="text-2xl font-bold tracking-tight">Patient Overview</h1>
      <p class="mt-1 text-sm text-muted-foreground">
        Monitor and manage patient records, vitals, and care plans.
      </p>
    </div>
    <div class="flex items-center gap-3">
      <div class="relative">
        <UiBaseInput
          v-model="search"
          class="w-64 pl-9"
          placeholder="Search patients..."
          value=""
          leading-icon="lucide:search"
          leadingIconClass="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground"
        >
        </UiBaseInput>
      </div>
      <button
        type="button"
        @click="openInviteModal"
        class="inline-flex items-center gap-2 rounded-md bg-primary px-4 py-3 text-sm font-medium text-primary-foreground transition hover:bg-primary/90 focus-visible:ring-2 focus-visible:ring-ring focus-visible:outline-none disabled:opacity-50"
      >
        <Icon name="lucide:plus" class="h-3.5 w-3.5" />
        Add Patient
      </button>
    </div>
  </div>
  <div class="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-4">
    <div
      v-for="item in metrics"
      :key="item.title"
      class="rounded-2xl border border-gray-100 bg-white p-6 shadow-sm dark:border-gray-800 dark:bg-gray-900"
    >
      <div class="flex items-center justify-between">
        <div
          :class="[
            'flex h-12 w-12 items-center justify-center rounded-xl',
            item.colorTheme,
          ]"
        >
          <Icon :name="item.icon" class="h-6 w-6" />
        </div>

        <div
          :class="[
            'flex items-center gap-1 rounded-full px-2 py-1 text-xs font-medium',
            item.trendSuccess
              ? 'bg-emerald-50 text-emerald-600'
              : 'bg-rose-50 text-rose-600',
          ]"
        >
          <Icon
            :name="
              item.trend.startsWith('+')
                ? 'lucide:trending-up'
                : 'lucide:trending-down'
            "
            class="h-3 w-3"
          />
          {{ item.trend }}
        </div>
      </div>

      <div class="mt-4">
        <h3 class="text-sm font-medium text-gray-500">{{ item.title }}</h3>
        <div class="mt-1 flex items-baseline gap-2">
          <span class="text-2xl font-bold text-gray-900 dark:text-white">{{
            item.value
          }}</span>
          <span class="text-xs text-gray-400">{{ item.trendText }}</span>
        </div>
      </div>
    </div>
  </div>
</template>
