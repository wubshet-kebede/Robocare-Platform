<script setup>
const colorMode = useColorMode();

const isDark = computed(() => colorMode.value === "dark");

function toggleTheme() {
  colorMode.preference = isDark.value ? "light" : "dark";
}
const navSections = [
  {
    title: "Overview",
    icon: "mynaui:chevron-right",
    items: [
      {
        name: "Dashboard",
        href: "/dashboard",
        icon: "tabler:layout-dashboard",
        active: true,
      },
      {
        name: "Analytics",
        href: "/analytics",
        icon: "hugeicons:chart-column",
      },
      {
        name: "Charts",
        href: "/charts",
        icon: "lucide:chart-no-axes-combined",
      },
    ],
  },
  {
    title: "Clinical",
    icon: "material-symbols-light:chevron-right-rounded",
    items: [
      {
        name: "Patient Overview",
        href: "/patients",
        icon: "material-symbols:stethoscope",
      },
      {
        name: "Appointments",
        href: "/appointments",
        icon: "mdi-light:calendar",
        // badge: 8,
      },
      {
        name: "Health Records",
        href: "/health-records",
        icon: "solar:clipboard-linear",
      },
      {
        name: "Prescriptions",
        href: "/prescriptions",
        icon: "lucide:pill",
      },
      {
        name: "Lab Results",
        href: "/lab-results",
        icon: "lucide:test-tube",
      },
      {
        name: "Vaccinations",
        href: "/vaccinations",
        icon: "lucide:syringe",
      },
    ],
  },
  {
    title: "Operations",
    icon: "material-symbols-light:chevron-right-rounded",
    items: [
      {
        name: "Telemedicine",
        href: "/telemedicine",
        icon: "lucide:video",
      },
      {
        name: "Bed Management",
        href: "/bed-management",
        icon: "lucide:bed-double",
      },
      {
        name: "Staff",
        href: "/staff",
        icon: "lucide:user-check",
      },
      {
        name: "Task Board",
        href: "/kanban",
        icon: "lucide:kanban",
      },
      {
        name: "Vitals Monitor",
        href: "/vital_monitor",
        icon: "lucide:activity",
      },
    ],
  },
  {
    title: "Administration",
    icon: "lucide:chevron-right",
    items: [
      {
        name: "Patients List",
        href: "/customers",
        icon: "lucide:users",
      },
      {
        name: "Billing",
        href: "/billing",
        icon: "lucide:credit-card",
      },
      {
        name: "Invoices",
        href: "/invoices",
        icon: "lucide:file-text",
      },
      {
        name: "Onboarding",
        href: "/wizard",
        icon: "lucide:list-checks",
      },
    ],
  },
];
const user = useAuthUser();
const slug = user.value?.hospital.slug;
const getInitials = (name) => {
  if (!name) return "";

  return name
    .split(" ")
    .map((n) => n[0])
    .join("")
    .toUpperCase();
};
</script>

<template>
  <div class="flex min-h-screen dark:bg-background">
    <aside
      class="fixed ltr:left-0 rtl:right-0 top-0 z-40 hidden h-screen flex-col dark:bg-sidebar transition-all duration-300 ease-in-out lg:flex w-[260px]"
    >
      <div
        class="flex h-16 items-center gap-3 border-b border-sidebar-border px-4"
      >
        <div
          class="flex h-8 w-8 shrink-0 items-center justify-center rounded-lg bg-sidebar-primary"
        >
          <Icon
            name="bi:heart-pulse-fill"
            className="h-6 w-6 text-sidebar-robocare"
          />
        </div>
        <div class="flex flex-col">
          <span
            class="text-sm font-bold tracking-tight dark:text-sidebar-robocare"
            >Hospital
          </span>
          <span
            class="text-[10px] font-medium uppercase tracking-widest dark:text-sidebar-robocare/40"
            >Dashboard</span
          >
        </div>
      </div>

      <nav
        role="navigation"
        aria-label="Main navigation"
        class="flex-1 space-y-3 overflow-y-auto px-3 py-4"
      >
        <div v-for="section in navSections" :key="section.title">
          <!-- Section Header -->
          <button
            class="flex w-full items-center gap-2 rounded-md px-3 py-1.5 text-[10px] font-semibold uppercase tracking-widest dark:text-sidebar-robocare/30 transition-colors hover:text-sidebar-robocare/50"
          >
            <span class="flex-1 text-start">{{ section.title }}</span>
            <Icon
              :name="section.icon"
              class="size-3 transition-transform duration-200 rotate-90"
            />
          </button>

          <!-- Section Items -->
          <div
            class="grid transition-all duration-200 ease-in-out grid-rows-[1fr] opacity-100"
          >
            <div class="overflow-hidden">
              <div class="mt-1 space-y-0.5">
                <NuxtLink
                  v-for="item in section.items"
                  :key="item.name"
                  :to="`/${slug}${item.href}`"
                  class="group relative flex items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium transition-all duration-200 dark:text-sidebar-robocare/70 hover:bg-sidebar-accent/50 hover:text-sidebar-robocare"
                  :class="
                    item.active
                      ? 'dark:bg-sidebar-accent dark:text-sidebar-primary'
                      : ''
                  "
                >
                  <Icon
                    :name="item.icon"
                    class="h-[18px] w-[18px] shrink-0 transition-colors dark:text-sidebar-robocare/50 dark:group-hover:text-sidebar-robocare/80"
                  />

                  <span class="flex-1">{{ item.name }}</span>

                  <!-- Badge -->
                  <span
                    v-if="item.badge"
                    class="flex h-5 min-w-5 items-center justify-center rounded-full dark:bg-sidebar-primary/15 px-1.5 text-[10px] font-semibold text-sidebar-primary"
                  >
                    {{ item.badge }}
                  </span>
                </NuxtLink>
              </div>
            </div>
          </div>
        </div>
      </nav>

      <div class="border-t dark:border-sidebar-border p-3">
        <div class="flex items-center gap-2">
          <a
            class="flex flex-1 items-center gap-3 rounded-lg px-2 py-2 transition-colors dark:hover:bg-sidebar-accent/50"
            href="/profile"
          >
            <div
              class="flex h-8 w-8 shrink-0 items-center justify-center rounded-full bg-gradient-to-br from-sidebar-primary/80 to-sidebar-primary text-[11px] font-bold text-sidebar-primary-robocare"
            >
              {{ getInitials(user?.full_name) }}
            </div>
            <div class="flex flex-1 flex-col">
              <span class="text-sm font-medium dark:text-sidebar-foreground">{{
                user?.full_name
              }}</span>
              <span class="text-[11px] dark:text-sidebar-foreground/50">{{
                user?.role
              }}</span>
            </div> </a
          ><button
            aria-label="Log out"
            class="rounded-md p-1.5 dark:text-sidebar-foreground/40 transition-colors dark:hover:bg-sidebar-accent dark:hover:text-sidebar-robocare/70"
          >
            <Icon name="lucide:log-out" class="h-4 w-4" />
          </button>
        </div>
      </div>
      <button
        aria-label="Collapse sidebar"
        class="absolute ltr:-right-3 rtl:-left-3 top-20 flex h-6 w-6 items-center justify-center rounded-full border dark:border-sidebar-border dark:bg-sidebar dark:text-sidebar-robocare/50 shadow-sm transition-all dark:hover:bg-sidebar-accent dark:hover:text-sidebar-robocare"
      >
        <Icon
          name="lucide:chevron-left"
          class="h-3.5 w-3.5 transition-transform duration-300 ltr:rotate-0 rtl:rotate-180"
        />
      </button>
    </aside>
    <div class="flex flex-1 flex-col transition-all duration-300 lg:ms-[260px]">
      <div class="sticky top-0 z-30">
        <header
          class="flex h-16 items-center justify-between border-b dark:border-border dark:bg-background/80 px-4 backdrop-blur-xl sm:px-6"
        >
          <!-- LEFT -->
          <div class="flex items-center gap-3">
            <button
              aria-label="Open menu"
              class="flex h-8 w-8 items-center justify-center rounded-lg dark:text-muted-foreground transition-colors dark:hover:bg-accent dark:hover:text-foreground lg:hidden"
            >
              <Icon name="lucide:menu" class="h-5 w-5" />
            </button>

            <button
              class="relative hidden h-9 w-72 items-center rounded-lg border dark:border-input dark:bg-muted/40 ps-9 pe-4 text-start text-sm dark:text-muted-foreground/50 transition-colors hover:bg-muted/60 sm:flex"
            >
              <Icon
                name="lucide:search"
                class="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 dark:text-muted-foreground/60"
              />
              Search anything...
              <kbd
                class="absolute right-3 top-1/2 -translate-y-1/2 rounded border dark:border-border dark:bg-muted px-1.5 text-[10px]"
              >
                ⌘K
              </kbd>
            </button>
          </div>

          <!-- RIGHT -->
          <div class="flex items-center gap-2">
            <!-- New Appointment -->
            <button
              class="hidden sm:inline-flex items-center gap-1.5 h-8 px-3 text-xs rounded-md dark:bg-primary dark:text-primary-foreground dark:hover:bg-primary/90"
            >
              <Icon name="lucide:plus" class="h-3.5 w-3.5" />
              New Appointment
            </button>

            <div class="mx-1 hidden h-6 w-px dark:bg-border sm:block"></div>

            <!-- Theme -->
            <button
              @click="toggleTheme"
              class="h-8 w-8 flex items-center justify-center"
            >
              <span v-if="isDark">
                <Icon name="lucide:sun" class="h-4 w-4 dark:text-white"
              /></span>
              <span v-else
                ><Icon name="majesticons:moon-line" class="h-4 w-4"
              /></span>
            </button>

            <!-- Customize -->
            <button class="h-8 w-8 flex items-center justify-center">
              <Icon name="lucide:palette" class="h-4 w-4 dark:text-white" />
            </button>

            <!-- Notifications -->
            <button class="relative h-8 w-8 flex items-center justify-center">
              <Icon name="lucide:bell" class="h-4 w-4 dark:text-white" />
              <span
                class="absolute right-1.5 top-1.5 h-2 w-2 rounded-full dark:bg-destructive"
              ></span>
            </button>

            <!-- User -->
            <button
              class="ms-1 flex h-8 w-8 items-center justify-center rounded-full bg-primary/10 text-xs font-semibold text-primary transition-colors hover:bg-primary/20"
            >
              {{ getInitials(user?.full_name) }}
            </button>
          </div>
        </header>
      </div>
      <main class="flex-1 p-4 sm:p-6">
        <slot />
      </main>
    </div>
  </div>
</template>
