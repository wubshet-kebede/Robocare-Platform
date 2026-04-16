export default defineNuxtRouteMiddleware((to) => {
  const user = useAuthUser();

  if (!user.value) return navigateTo("/login");

  const slug = to.params.slug;

  if (slug && slug !== user.value.hospital.slug) {
    return navigateTo(`/${user.value.hospital.slug}/dashboard`);
  }
});
