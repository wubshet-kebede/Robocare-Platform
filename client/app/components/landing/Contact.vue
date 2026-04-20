<template>
  <section id="contact" class="py-24 bg-slate-50 border-t border-slate-200">
    <div class="container mx-auto px-8">
      <div class="grid lg:grid-cols-2 gap-16 items-start">
        <div class="space-y-8">
          <div>
            <h2
              class="text-cyan-600 font-semibold text-sm uppercase tracking-[0.25em] mb-4"
            >
              Get In Touch
            </h2>

            <h3 class="text-4xl font-bold text-slate-900 mb-6 tracking-tight">
              Inquire About the <br />
              <span class="text-slate-400 italic"> Framework Deployment. </span>
            </h3>

            <p class="text-slate-600 text-lg leading-relaxed max-w-md">
              Are you interested in testing our IoT-SLAM framework in your
              medical facility? Reach out to our research team at the University
              of Gondar.
            </p>
          </div>
          <div class="space-y-6">
            <div class="flex items-center gap-4 group">
              <div
                class="w-12 h-12 bg-white border border-slate-200 rounded-xl flex items-center justify-center group-hover:border-cyan-400 transition-colors shadow-sm"
              >
                <Icon name="mdi:location" class="text-xl text-cyan-600" />
              </div>
              <div>
                <div class="text-slate-900 font-semibold text-sm">Location</div>
                <div class="text-slate-500 text-sm">
                  Institute of Technology, Gondar, Ethiopia
                </div>
              </div>
            </div>

            <div class="flex items-center gap-4 group">
              <div
                class="w-12 h-12 bg-white border border-slate-200 rounded-xl flex items-center justify-center group-hover:border-emerald-400 transition-colors shadow-sm"
              >
                <Icon
                  name="ic:baseline-email"
                  class="text-xl text-emerald-600"
                />
              </div>
              <div>
                <div class="text-slate-900 font-semibold text-sm">Email</div>
                <div class="text-slate-500 text-sm">wubshet2829@gmail.com</div>
              </div>
            </div>
          </div>
        </div>
        <div
          class="bg-white border border-slate-200 p-8 md:p-10 rounded-[2rem] shadow-md hover:shadow-lg transition-all"
        >
          <form @submit.prevent class="space-y-6">
            <div class="grid md:grid-cols-2 gap-6">
              <div class="space-y-2">
                <label
                  class="text-[11px] font-semibold uppercase tracking-wide text-slate-500 ml-1"
                >
                  Full Name
                </label>
                <input
                  type="text"
                  placeholder="John Doe"
                  class="w-full bg-white border border-slate-300 rounded-xl px-5 py-4 text-slate-900 text-sm focus:border-cyan-500 focus:ring-2 focus:ring-cyan-100 outline-none transition-all"
                />
              </div>

              <div class="space-y-2">
                <label
                  class="text-[11px] font-semibold uppercase tracking-wide text-slate-500 ml-1"
                >
                  Email Address
                </label>
                <input
                  type="email"
                  placeholder="john@hospital.com"
                  class="w-full bg-white border border-slate-300 rounded-xl px-5 py-4 text-slate-900 text-sm focus:border-cyan-500 focus:ring-2 focus:ring-cyan-100 outline-none transition-all"
                />
              </div>
            </div>

            <div class="space-y-2">
              <label
                class="text-[11px] font-semibold uppercase tracking-wide text-slate-500 ml-1"
              >
                Inquiry Type
              </label>
              <select
                class="w-full bg-white border border-slate-300 rounded-xl px-5 py-4 text-slate-600 text-sm focus:border-cyan-500 focus:ring-2 focus:ring-cyan-100 outline-none transition-all"
              >
                <option>Deployment Request</option>
                <option>Technical Collaboration</option>
                <option>Academic Inquiry</option>
              </select>
            </div>

            <div class="space-y-2">
              <label
                class="text-[11px] font-semibold uppercase tracking-wide text-slate-500 ml-1"
              >
                Message
              </label>
              <textarea
                rows="4"
                placeholder="How can we help your facility?"
                class="w-full bg-white border border-slate-300 rounded-xl px-5 py-4 text-slate-900 text-sm focus:border-cyan-500 focus:ring-2 focus:ring-cyan-100 outline-none transition-all resize-none"
              ></textarea>
            </div>

            <button
              class="w-full py-4 bg-gradient-to-r from-cyan-500 to-emerald-500 hover:from-cyan-600 hover:to-emerald-600 text-white font-semibold rounded-xl transition-all transform hover:-translate-y-0.5 shadow-md hover:shadow-lg"
            >
              Send Message
            </button>
          </form>
        </div>
      </div>
    </div>
  </section>
</template>
