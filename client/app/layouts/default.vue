<template>
  <div class="min-h-screen bg-white font-sans text-slate-200">
    <!-- Note the Navigation prefix -->
    <NavigationHeader />

    <main class="pt-20">
      <slot />
    </main>

    <NavigationFooter />
  </div>
</template>
