<template>
  <div
    class="min-h-screen bg-gray-50 flex flex-col items-center justify-center px-6"
  >
    <!-- Hero Section -->
    <section class="text-center max-w-2xl">
      <h1 class="text-4xl font-bold text-gray-900 mb-4">Welcome to RoboCare</h1>
      <p class="text-lg text-gray-600 mb-6">
        A modular healthcare robotics platform designed to monitor, assist, and
        empower hospital workflows.
      </p>
      <div class="flex justify-center gap-4">
        <NuxtLink
          to="/robots"
          class="bg-blue-600 text-white px-5 py-2 rounded-lg shadow hover:bg-blue-700 transition"
        >
          View Robots
        </NuxtLink>
        <NuxtLink
          to="/patients"
          class="bg-gray-200 text-gray-800 px-5 py-2 rounded-lg shadow hover:bg-gray-300 transition"
        >
          Patient Dashboard
        </NuxtLink>
      </div>
    </section>

    <!-- Features Section -->
    <section
      class="mt-12 grid grid-cols-1 md:grid-cols-3 gap-6 w-full max-w-4xl"
    >
      <div class="bg-white p-6 rounded-lg shadow hover:shadow-lg transition">
        <h2 class="text-xl font-semibold mb-2">🤖 Robot Monitoring</h2>
        <p class="text-gray-600">
          Track robot status, battery, and live location in real-time.
        </p>
      </div>
      <div class="bg-white p-6 rounded-lg shadow hover:shadow-lg transition">
        <h2 class="text-xl font-semibold mb-2">🏥 Patient Care</h2>
        <p class="text-gray-600">
          Assign robots to patients for delivery, escort, and assistance.
        </p>
      </div>
      <div class="bg-white p-6 rounded-lg shadow hover:shadow-lg transition">
        <h2 class="text-xl font-semibold mb-2">⚙️ Admin Control</h2>
        <p class="text-gray-600">
          Configure hospital settings and manage robotic workflows.
        </p>
      </div>
    </section>
  </div>
</template>

<script setup lang="ts"></script>
