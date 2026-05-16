export const useStaffService = () => {
  const { $api } = useNuxtApp();

  const getAssignableStaff = () => {
    return $api("/assignable-staff", {
      method: "GET",
    });
  };

  return { getAssignableStaff };
};
