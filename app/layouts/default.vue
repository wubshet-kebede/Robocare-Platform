<template>
  <div class="min-h-screen bg-slate-950 font-sans text-slate-200">
    <!-- Note the Navigation prefix -->
    <NavigationHeader />

    <main class="pt-20">
      <slot />
    </main>

    <NavigationFooter />
  </div>
</template>

<!-- <style>
/* Smooth scrolling for anchor links (About, Contact, etc.) */
html {
  scroll-behavior: smooth;
}
</style> -->
